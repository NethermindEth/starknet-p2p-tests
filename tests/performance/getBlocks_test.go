package tests

import (
	"context"
	"fmt"
	"starknet-p2p-tests/config"
	synthetic_node "starknet-p2p-tests/tools"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/montanaflynn/stats"
	"github.com/stretchr/testify/require"
)

const (
	testTimeout      = 5 * time.Minute
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

func TestBlockHeaderRequestPerformance(t *testing.T) {
	peerCounts := []int{1, 5, 10, 15, 20}
	stats := make([]LatencyStats, 0, len(peerCounts))

	for _, peerCount := range peerCounts {
		t.Logf("Starting test for %d peers", peerCount)
		latencyStats := runPerformanceTest(t, peerCount)
		stats = append(stats, latencyStats)
		time.Sleep(5 * time.Second) // Cool-down period between tests
	}

	printPerformanceTable(t, stats)
}

func runPerformanceTest(t *testing.T, peerCount int) LatencyStats {
	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()

	var wg sync.WaitGroup
	latencies := make([]float64, 0, peerCount*requestsPerPeer)
	connectTimes := make([]float64, 0, peerCount)
	var latenciesMutex sync.Mutex

	startTime := time.Now()

	var successfulRequests int64
	totalRequests := int64(peerCount * requestsPerPeer)

	errorCounts := make(map[string]int)
	var errorCountsMutex sync.Mutex

	for i := 0; i < peerCount; i++ {
		wg.Add(1)
		go func(peerIndex int) {
			defer wg.Done()
			peerLatencies, peerSuccesses, connectTime, peerErrors := simulatePeer(t, ctx, peerIndex, peerCount)
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

func simulatePeer(t *testing.T, ctx context.Context, peerIndex, totalPeers int) ([]float64, int, float64, map[string]int) {
	syntheticNode, err := synthetic_node.New(ctx, t)
	require.NoError(t, err, "Failed to create synthetic node")
	// Remove the defer syntheticNode.Close() line, as it's now handled by t.Cleanup()

	connectStart := time.Now()
	err = syntheticNode.Connect(ctx, config.TargetPeerAddress)
	connectTime := time.Since(connectStart).Seconds() * 1000
	require.NoError(t, err, "Failed to connect to target peer")

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

func printPerformanceTable(t *testing.T, stats []LatencyStats) {
	t.Log("\nBlock Header Request Performance Test Results:")
	t.Logf("%-10s %-10s %-10s %-10s %-10s %-10s %-10s %-10s %-10s %-10s %-10s %-20s",
		"Peers", "Requests", "Min (ms)", "Max (ms)", "Mean (ms)", "Median (ms)", "P95 (ms)", "Total (s)", "Success %", "Throughput", "Conn (ms)", "Errors")
	t.Log(strings.Repeat("-", 140))

	for _, s := range stats {
		errStr := fmt.Sprintf("Req:%d,TO:%d,Empty:%d",
			s.ErrorCounts["request_error"],
			s.ErrorCounts["timeout"],
			s.ErrorCounts["empty_response"])
		t.Logf("%-10d %-10d %-10.2f %-10.2f %-10.2f %-10.2f %-10.2f %-10.2f %-10.2f %-10.2f %-10.2f %s",
			s.PeerCount, s.PeerCount*requestsPerPeer, s.MinLatency, s.MaxLatency,
			s.MeanLatency, s.MedianLatency, s.P95Latency, s.TotalTime, s.SuccessRate,
			s.Throughput, s.AvgConnectTime, errStr)
	}
}
