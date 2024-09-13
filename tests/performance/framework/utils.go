package framework

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

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
