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

	return CalculateStats(latencies, connectTimes, successfulRequests, totalRequests, totalTime, peerCount, errorCounts)
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
