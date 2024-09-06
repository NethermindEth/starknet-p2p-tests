package tests

import (
	"context"
	"testing"
	"time"

	"starknet-p2p-tests/config"
	synthetic_node "starknet-p2p-tests/tools"

	"github.com/stretchr/testify/require"
)

func TestDiscovery(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	node1, err := synthetic_node.New(ctx)
	require.NoError(t, err)
	node2, err := synthetic_node.New(ctx)
	require.NoError(t, err)

	// Connect both synthetic nodes to the target node
	err = node1.Connect(ctx, config.TargetPeerAddress)
	require.NoError(t, err)
	err = node2.Connect(ctx, config.TargetPeerAddress)
	require.NoError(t, err)

	// Wait for the DHT to work its magic
	time.Sleep(60 * time.Second)

	peers1 := node1.Host.Peerstore().Peers()
	peers2 := node2.Host.Peerstore().Peers()

	require.Contains(t, peers1, node2.Host.ID(), "Node 1 should have discovered Node 2")
	require.Contains(t, peers2, node1.Host.ID(), "Node 2 should have discovered Node 1")

	err = node1.Close()
	require.NoError(t, err)
	err = node2.Close()
	require.NoError(t, err)
}
