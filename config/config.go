package config

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var (
	TargetPeerAddress    string
	DefaultTestTimeout   time.Duration
	SyntheticListenAddrs []string
	NetworkName          string
)

func init() {
	//TargetPeerAddress = getEnv("TARGET_PEER_ADDRESS", "/ip4/35.237.66.77/tcp/7777/p2p/12D3KooWR8ikUDiinyE5wgdYiqsdLfJRsBDYKGii6L3oyoipVEaV")
	TargetPeerAddress = getEnv("TARGET_PEER_ADDRESS", "/ip4/127.0.0.1/tcp/7777/p2p/12D3KooWFY7F7BrsVXFkfivcsFnDx7nGZu9G6YDNJ9oscohm5z8n")
	timeoutStr := getEnv("DEFAULT_TEST_TIMEOUT", "30s")
	var err error
	DefaultTestTimeout, err = time.ParseDuration(timeoutStr)
	if err != nil {
		fmt.Printf("Error parsing DEFAULT_TEST_TIMEOUT: %v. Defaulting to 30 seconds.\n", err)
		DefaultTestTimeout = 30 * time.Second
	}

	SyntheticListenAddrs = strings.Split(getEnv("SYNTHETIC_LISTEN_ADDRS", "/ip4/0.0.0.0/tcp/0"), ",")
	NetworkName = getEnv("NETWORK_NAME", "sepolia")
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
