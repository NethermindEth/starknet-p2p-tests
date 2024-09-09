package tests

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"starknet-p2p-tests/config"
	synthetic_node "starknet-p2p-tests/tools"

	"encoding/json"

	"github.com/montanaflynn/stats"
	"github.com/stretchr/testify/require"
)

const (
	expectedBlockNum = uint64(1)
	requestsPerPeer  = 100
	responseTimeout  = 3 * time.Second
	rampUpTime       = 30 * time.Second
)

type LatencyStats struct {
	PeerCount      int
	MinLatency     float64
	MaxLatency     float64
	MeanLatency    float64
	MedianLatency  float64
	P95Latency     float64
	TotalTime      float64
	SuccessRate    float64
	Throughput     float64
	AvgConnectTime float64
	ErrorCounts    map[string]int
}

func BenchmarkBlockHeaderRequestPerformance(b *testing.B) {
	peerCounts := []int{1, 5, 10, 15, 20}
	allResults := make(map[string]LatencyStats)

	for _, peerCount := range peerCounts {
		b.Run(fmt.Sprintf("Peers-%d", peerCount), func(b *testing.B) {
			latencyStats := runPerformanceTest(b, peerCount)
			allResults[fmt.Sprintf("Peers-%d", peerCount)] = latencyStats
		})
		time.Sleep(5 * time.Second) // Cool-down period between tests
	}

	// Write all results to a file after all tests are complete
	writeResultsToFile(b, allResults)
}

func runPerformanceTest(b *testing.B, peerCount int) LatencyStats {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	latencies := make([]float64, 0, peerCount*requestsPerPeer)
	connectTimes := make([]float64, 0, peerCount)
	var latenciesMutex sync.Mutex

	var successfulRequests int64
	totalRequests := int64(peerCount * requestsPerPeer)

	errorCounts := make(map[string]int)
	var errorCountsMutex sync.Mutex

	b.ResetTimer() // Start timing here
	startTime := time.Now()

	for i := 0; i < peerCount; i++ {
		wg.Add(1)
		go func(peerIndex int) {
			defer wg.Done()
			peerLatencies, peerSuccesses, connectTime, peerErrors := simulatePeer(b, ctx, peerIndex, peerCount)
			atomic.AddInt64(&successfulRequests, int64(peerSuccesses))
			latenciesMutex.Lock()
			latencies = append(latencies, peerLatencies...)
			connectTimes = append(connectTimes, connectTime)
			latenciesMutex.Unlock()

			errorCountsMutex.Lock()
			for errType, count := range peerErrors {
				errorCounts[errType] += count
			}
			errorCountsMutex.Unlock()
		}(i)

		// Gradual ramp-up
		time.Sleep(rampUpTime / time.Duration(peerCount))
	}

	wg.Wait()
	totalTime := time.Since(startTime).Seconds()

	b.StopTimer() // Stop timing here

	min, _ := stats.Min(latencies)
	max, _ := stats.Max(latencies)
	mean, _ := stats.Mean(latencies)
	median, _ := stats.Median(latencies)
	p95, _ := stats.Percentile(latencies, 95)
	successRate := float64(successfulRequests) / float64(totalRequests) * 100
	throughput := float64(successfulRequests) / totalTime
	avgConnectTime, _ := stats.Mean(connectTimes)

	return LatencyStats{
		PeerCount:      peerCount,
		MinLatency:     min,
		MaxLatency:     max,
		MeanLatency:    mean,
		MedianLatency:  median,
		P95Latency:     p95,
		TotalTime:      totalTime,
		SuccessRate:    successRate,
		Throughput:     throughput,
		AvgConnectTime: avgConnectTime,
		ErrorCounts:    errorCounts,
	}
}

func simulatePeer(b *testing.B, ctx context.Context, peerIndex, totalPeers int) ([]float64, int, float64, map[string]int) {
	syntheticNode, err := synthetic_node.New(ctx, b)
	require.NoError(b, err, "Failed to create synthetic node")

	connectStart := time.Now()
	err = syntheticNode.Connect(ctx, config.TargetPeerAddress)
	connectTime := time.Since(connectStart).Seconds() * 1000
	require.NoError(b, err, "Failed to connect to target peer")

	latencies := make([]float64, 0, requestsPerPeer)
	successfulRequests := 0
	errorCounts := make(map[string]int)

	for i := 0; i < requestsPerPeer; i++ {
		start := time.Now()
		headers, err := syntheticNode.RequestBlockHeaders(ctx, expectedBlockNum, 1)
		if err != nil {
			errorCounts["request_error"]++
			continue
		}

		responseCtx, cancel := context.WithTimeout(ctx, responseTimeout)
		defer cancel()

		select {
		case <-responseCtx.Done():
			errorCounts["timeout"]++
		case <-ctx.Done():
			return latencies, successfulRequests, connectTime, errorCounts
		default:
			if len(headers) > 0 {
				latency := time.Since(start).Seconds() * 1000 // Convert to milliseconds
				latencies = append(latencies, latency)
				successfulRequests++
			} else {
				errorCounts["empty_response"]++
			}
		}

		// Simulate some thinking time between requests
		time.Sleep(time.Duration(50+peerIndex*10) * time.Millisecond)
	}

	return latencies, successfulRequests, connectTime, errorCounts
}

func writeResultsToFile(b *testing.B, results map[string]LatencyStats) {
	timestamp := time.Now().Format("20060102_150405")
	filename := filepath.Join("benchmark_results", fmt.Sprintf("benchmark_results_%s.json", timestamp))

	// Ensure the directory exists
	err := os.MkdirAll(filepath.Dir(filename), 0755)
	if err != nil {
		b.Fatalf("Failed to create directory: %v", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		b.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty print the JSON
	if err := encoder.Encode(results); err != nil {
		b.Fatalf("Failed to write results to file: %v", err)
	}
}
