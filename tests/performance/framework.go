package performance

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"starknet-p2p-tests/config"
	synthetic_node "starknet-p2p-tests/tools"

	"github.com/montanaflynn/stats"
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

type TestFunc func(ctx context.Context, syntheticNode *synthetic_node.SyntheticNode) (time.Duration, error)

func RunTest(b *testing.B, peerCount, requestsPerPeer int, rampUpTime, responseTimeout time.Duration, testFunc TestFunc) LatencyStats {
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

	b.ResetTimer()
	startTime := time.Now()

	for i := 0; i < peerCount; i++ {
		wg.Add(1)
		go func(peerIndex int) {
			defer wg.Done()
			peerLatencies, peerSuccesses, connectTime, peerErrors := simulatePeer(b, ctx, peerIndex, peerCount, requestsPerPeer, responseTimeout, testFunc)
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

		time.Sleep(rampUpTime / time.Duration(peerCount))
	}

	wg.Wait()
	totalTime := time.Since(startTime).Seconds()

	b.StopTimer()

	return calculateStats(latencies, connectTimes, successfulRequests, totalRequests, totalTime, peerCount, errorCounts)
}

func simulatePeer(b *testing.B, ctx context.Context, peerIndex, totalPeers, requestsPerPeer int, responseTimeout time.Duration, testFunc TestFunc) ([]float64, int, float64, map[string]int) {
	syntheticNode, err := synthetic_node.New(ctx, b)
	if err != nil {
		b.Fatalf("Failed to create synthetic node: %v", err)
	}

	connectStart := time.Now()
	err = syntheticNode.Connect(ctx, config.TargetPeerAddress)
	connectTime := time.Since(connectStart).Seconds() * 1000
	if err != nil {
		b.Fatalf("Failed to connect to target peer: %v", err)
	}

	latencies := make([]float64, 0, requestsPerPeer)
	successfulRequests := 0
	errorCounts := make(map[string]int)

	for i := 0; i < requestsPerPeer; i++ {
		responseCtx, cancel := context.WithTimeout(ctx, responseTimeout)
		latency, err := testFunc(responseCtx, syntheticNode)
		cancel()

		if err != nil {
			errorCounts[err.Error()]++
			continue
		}

		latencies = append(latencies, float64(latency.Milliseconds()))
		successfulRequests++

		time.Sleep(time.Duration(50+peerIndex*10) * time.Millisecond)
	}

	return latencies, successfulRequests, connectTime, errorCounts
}

func calculateStats(latencies, connectTimes []float64, successfulRequests, totalRequests int64, totalTime float64, peerCount int, errorCounts map[string]int) LatencyStats {
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

func WriteResultsToFile(b *testing.B, results map[string]LatencyStats) {
	timestamp := time.Now().Format("20060102_150405")
	filename := filepath.Join("benchmark_results", fmt.Sprintf("benchmark_results_%s.json", timestamp))

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
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(results); err != nil {
		b.Fatalf("Failed to write results to file: %v", err)
	}
}
