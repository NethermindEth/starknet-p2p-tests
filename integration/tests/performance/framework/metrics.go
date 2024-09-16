package framework

import (
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

func CalculateStats(latencies, connectTimes []float64, successfulRequests, totalRequests int64, totalTime float64, peerCount int, errorCounts map[string]int) LatencyStats {
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
