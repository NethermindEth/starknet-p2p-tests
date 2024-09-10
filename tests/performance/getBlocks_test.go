package performance

import (
	"context"
	"fmt"
	"testing"
	"time"

	"starknet-p2p-tests/tests/performance/framework"
	synthetic_node "starknet-p2p-tests/tools"
)

const (
	expectedBlockNum = uint64(1)
	requestsPerPeer  = 100
	responseTimeout  = 3 * time.Second
	rampUpTime       = 30 * time.Second
)

// BenchmarkBlockHeaderRequestPerformance measures the performance of block header requests across varying peer counts,
// executing the requests in parallel to simulate concurrent load.
func BenchmarkBlockHeaderRequestPerformance(b *testing.B) {
	peerCounts := []int{1, 5, 10, 15, 20}
	allResults := make(map[string]framework.LatencyStats)

	for _, peerCount := range peerCounts {
		peerCount := peerCount // Capture range variable
		b.Run(fmt.Sprintf("Peers-%d", peerCount), func(b *testing.B) {
			latencyStats := framework.RunTest(b, peerCount, requestsPerPeer, rampUpTime, responseTimeout, requestBlockHeaders)
			allResults[fmt.Sprintf("Peers-%d", peerCount)] = latencyStats
		})
	}

	// Write all results to a file after all tests are complete using b.Cleanup
	b.Cleanup(func() {
		framework.WriteResultsToFile(b, allResults)
	})
}

// requestBlockHeaders sends a request to fetch block headers and measures the latency.
func requestBlockHeaders(ctx context.Context, syntheticNode *synthetic_node.SyntheticNode) (time.Duration, error) {
	// Apply a context timeout to prevent hanging requests
	ctx, cancel := context.WithTimeout(ctx, responseTimeout)
	defer cancel()

	start := time.Now()
	headers, err := syntheticNode.RequestBlockHeaders(ctx, expectedBlockNum, 1)
	if err != nil {
		return 0, err
	}
	if len(headers) == 0 {
		return 0, fmt.Errorf("received empty response for block number %d", expectedBlockNum)
	}
	return time.Since(start), nil
}
