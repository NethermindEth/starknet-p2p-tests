package framework

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"starknet-p2p-tests/config"
	synthetic_node "starknet-p2p-tests/tools"
)

type PeerStats struct {
	latencies   []float64
	successes   int
	connectTime float64
	errorCounts map[string]int
}

type TestFunc func(ctx context.Context, syntheticNode *synthetic_node.SyntheticNode) (time.Duration, error)

// RunTest starts multiple peers, runs the specified test function for each peer,
// and collects performance statistics. It simulates a specified number of peers
// making requests to a target node and measures latencies and error rates.
func RunTest(b *testing.B, peerCount, requestsPerPeer int, rampUpTime, responseTimeout time.Duration, testFunc TestFunc) LatencyStats {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var (
		wg                 sync.WaitGroup
		peerStatsList      = make([]PeerStats, 0, peerCount) // slice to hold stats for each peer
		latenciesMutex     sync.Mutex
		successfulRequests int64
		totalRequests      = int64(peerCount * requestsPerPeer)
	)

	// Helper function to collect stats from each peer
	collectStats := func(stats PeerStats) {
		atomic.AddInt64(&successfulRequests, int64(stats.successes))
		latenciesMutex.Lock()
		peerStatsList = append(peerStatsList, stats)
		latenciesMutex.Unlock()
	}

	b.ResetTimer()
	startTime := time.Now()

	for i := 0; i < peerCount; i++ {
		wg.Add(1)
		go func(peerIndex int) {
			defer wg.Done()
			// Simulate peer and collect stats
			peerStats := simulatePeer(b, ctx, peerIndex, requestsPerPeer, responseTimeout, testFunc)
			collectStats(peerStats)
		}(i)

		time.Sleep(rampUpTime / time.Duration(peerCount))
	}

	wg.Wait()
	totalTime := time.Since(startTime).Seconds()

	b.StopTimer()

	// Gather all latencies and error counts from each peer to calculate stats
	var allLatencies []float64
	var allErrorCounts = make(map[string]int)
	var connectTimes []float64

	for _, stats := range peerStatsList {
		allLatencies = append(allLatencies, stats.latencies...)
		connectTimes = append(connectTimes, stats.connectTime)
		for errType, count := range stats.errorCounts {
			allErrorCounts[errType] += count
		}
	}

	return CalculateStats(allLatencies, connectTimes, successfulRequests, totalRequests, totalTime, peerCount, allErrorCounts)
}

// simulatePeer simulates the behavior of a single peer, collects latencies, and tracks success/error counts.
func simulatePeer(b *testing.B, ctx context.Context, peerIndex, requestsPerPeer int, responseTimeout time.Duration, testFunc TestFunc) PeerStats {
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

	return PeerStats{
		latencies:   latencies,
		successes:   successfulRequests,
		connectTime: connectTime,
		errorCounts: errorCounts,
	}
}
