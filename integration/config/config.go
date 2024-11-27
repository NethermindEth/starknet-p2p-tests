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
	TargetPeerAddress = getEnv("TARGET_PEER_ADDRESS", "/ip4/192.168.64.1/tcp/7777/p2p/12D3KooWLdURCjbp1D7hkXWk6ZVfcMDPtsNnPHuxoTcWXFtvrxGG")
	timeoutStr := getEnv("DEFAULT_TEST_TIMEOUT", "300s")
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
