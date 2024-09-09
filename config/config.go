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
	TargetPeerAddress = getEnv("TARGET_PEER_ADDRESS", "/ip4/127.0.0.1/tcp/7777/p2p/12D3KooWLBUjEPyTiACzQZ3K1oqBXRqHwRFvAUHrm561pWWbJkYf")
	timeoutStr := getEnv("DEFAULT_TEST_TIMEOUT", "3000s")
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

// P2P Private Key: 3d59fe5117449e06bd4b64e789e78b97a17ed2703d840cf81d3ab298a999904599fd7bdb269c2b86267a04b3867a78352a96f97b331d980cde7d83eb7d0eace0
// P2P Public Key: 99fd7bdb269c2b86267a04b3867a78352a96f97b331d980cde7d83eb7d0eace0
// P2P PeerID: 12D3KooWLBUjEPyTiACzQZ3K1oqBXRqHwRFvAUHrm561pWWbJkYf
// ./juno-v0.12.2-8-gb89e0786-macOS-x86_64 --network sepolia --p2p --p2p-feeder-node --p2p-addr /ip4/127.0.0.1/tcp/7777 --p2p-private-key 3d59fe5117449e06bd4b64e789e78b97a17ed2703d840cf81d3ab298a999904599fd7bdb269c2b86267a04b3867a78352a96f97b331d980cde7d83eb7d0eace0
