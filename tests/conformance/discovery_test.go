package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"starknet-p2p-tests/config"
	synthetic_node "starknet-p2p-tests/tools"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDiscovery(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
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

	fmt.Printf("Node 1 ID: %s\n", node1.Host.ID())
	fmt.Printf("Node 2 ID: %s\n", node2.Host.ID())

	// Wait for the DHT to work its magic
	time.Sleep(100 * time.Second)

	peers1 := node1.Host.Peerstore().Peers()
	peers2 := node2.Host.Peerstore().Peers()

	fmt.Printf("Node 1 peers: %v\n", peers1)
	fmt.Printf("Node 2 peers: %v\n", peers2)

	assert.Contains(t, peers1, node2.Host.ID(), "Node 1 should have discovered Node 2")
	assert.Contains(t, peers2, node1.Host.ID(), "Node 2 should have discovered Node 1")

	err = node1.Close()
	require.NoError(t, err)
	err = node2.Close()
	require.NoError(t, err)
}
