package tests

import (
	"context"
	"testing"
	"time"

	"starknet-p2p-tests/config"
	synthetic_node "starknet-p2p-tests/tools"

	"github.com/stretchr/testify/require"
)

// TestDiscovery tests the discovery functionality between two synthetic nodes.
// It creates two synthetic nodes, connects them to a target node, and then
// waits for them to discover each other within a specified timeout.
func TestDiscovery(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), config.DefaultTestTimeout)
	defer cancel()

	// Creating synthetic nodes for testing discovery functionality
	node1, err := synthetic_node.New(ctx, t)
	require.NoError(t, err)
	node2, err := synthetic_node.New(ctx, t)
	require.NoError(t, err)

	// Connect both synthetic nodes to the target node
	require.NoError(t, node1.Connect(ctx, config.TargetPeerAddress))
	require.NoError(t, node2.Connect(ctx, config.TargetPeerAddress))

	// Wait for nodes to discover each other
	require.NoError(t, node1.WaitForPeerDiscovery(ctx, node2.Host.ID(), 10*time.Second),
		"Node 1 failed to discover Node 2 within the timeout")
	require.NoError(t, node2.WaitForPeerDiscovery(ctx, node1.Host.ID(), 10*time.Second),
		"Node 2 failed to discover Node 1 within the timeout")
}
