package performance

import (
	"context"
	"fmt"
	"testing"
	"time"

	synthetic_node "starknet-p2p-tests/tools"
)

const (
	expectedBlockNum = uint64(1)
	requestsPerPeer  = 100
	responseTimeout  = 3 * time.Second
	rampUpTime       = 30 * time.Second
)

func BenchmarkBlockHeaderRequestPerformance(b *testing.B) {
	peerCounts := []int{1, 5}
	allResults := make(map[string]LatencyStats)

	for _, peerCount := range peerCounts {
		b.Run(fmt.Sprintf("Peers-%d", peerCount), func(b *testing.B) {
			latencyStats := runPerformanceTest(b, peerCount)
			allResults[fmt.Sprintf("Peers-%d", peerCount)] = latencyStats
		})
		time.Sleep(5 * time.Second) // Cool-down period between tests
	}

	// Write all results to a file after all tests are complete
	WriteResultsToFile(b, allResults)
}

func runPerformanceTest(b *testing.B, peerCount int) LatencyStats {
	testFunc := func(ctx context.Context, syntheticNode *synthetic_node.SyntheticNode) (time.Duration, error) {
		start := time.Now()
		headers, err := syntheticNode.RequestBlockHeaders(ctx, expectedBlockNum, 1)
		if err != nil {
			return 0, err
		}
		if len(headers) == 0 {
			return 0, fmt.Errorf("empty response")
		}
		return time.Since(start), nil
	}

	return RunTest(b, peerCount, requestsPerPeer, rampUpTime, responseTimeout, testFunc)
}
