package synthetic_node

import (
	"context"
	"crypto/rand"
	"errors"
	"iter"
	"log"
	"os"
	"starknet-p2p-tests/config"
	"starknet-p2p-tests/protocol/p2p/starknet"
	"starknet-p2p-tests/protocol/p2p/starknet/spec"
	"starknet-p2p-tests/protocol/p2p/utils"
	"time"

	"testing"

	"github.com/avast/retry-go"
	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/multiformats/go-multiaddr"
)

type SyntheticNode struct {
	Host           host.Host
	StarknetClient *starknet.Client
	logger         utils.SimpleLogger
	dht            *dht.IpfsDHT
}

func New(ctx context.Context, t *testing.T) (*SyntheticNode, error) {
	stdLogger := log.New(os.Stdout, "[SYNTHETIC-NODE] ", log.Ldate|log.Ltime|log.Lshortfile)
	logger := &utils.TestSimpleLogger{Logger: stdLogger.Printf}

	priv, _, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		logger.Errorw("Failed to generate key", "error", err)
		return nil, errors.New("failed to generate key")
	}

	opts := []libp2p.Option{
		libp2p.Identity(priv),
		libp2p.ListenAddrStrings(config.SyntheticListenAddrs...),
	}

	h, err := libp2p.New(opts...)
	if err != nil {
		logger.Errorw("Failed to create libp2p node", "error", err)
		return nil, errors.New("failed to create libp2p node")
	}

	kadDHT, err := dht.New(ctx, h,
		dht.ProtocolPrefix(starknet.Prefix),
		dht.Mode(dht.ModeServer),
	)

	if err := kadDHT.Bootstrap(ctx); err != nil {
		logger.Errorw("Failed to bootstrap DHT", "error", err)
		return nil, errors.New("failed to bootstrap DHT")
	}

	if err != nil {
		logger.Errorw("Failed to create DHT", "error", err)
		return nil, errors.New("failed to create DHT")
	}

	node := &SyntheticNode{
		Host:   h,
		logger: logger,
		dht:    kadDHT,
	}

	if t != nil {
		t.Cleanup(func() {
			if err := node.Close(); err != nil {
				t.Logf("Error closing node: %v", err)
			}
		})
	}

	return node, nil
}

func (sn *SyntheticNode) Connect(ctx context.Context, targetAddress string) error {
	targetPeerInfo, err := ParsePeerAddress(targetAddress)
	if err != nil {
		sn.logger.Errorw("Failed to parse peer address", "error", err, "address", targetAddress)
		return errors.New("failed to parse peer address")
	}

	sn.logger.Infow("Connecting to peer", "address", targetAddress)
	if err := sn.Host.Connect(ctx, targetPeerInfo); err != nil {
		sn.logger.Errorw("Failed to connect to target peer", "error", err, "address", targetAddress)
		return errors.New("failed to connect to target peer")
	}

	newStreamFunc := func(ctx context.Context, pids ...protocol.ID) (network.Stream, error) {
		return sn.Host.NewStream(ctx, targetPeerInfo.ID, pids...)
	}

	sn.StarknetClient = starknet.NewClient(newStreamFunc, &utils.Network{Name: config.NetworkName}, sn.logger)
	sn.logger.Infow("Successfully connected to peer", "id", targetPeerInfo.ID)
	return nil
}

func (sn *SyntheticNode) RequestBlockHeaders(ctx context.Context, startBlock uint64, limit uint64, step ...uint64) ([]*spec.BlockHeadersResponse, error) {
	stepValue := uint64(1) // Default step value
	if len(step) > 0 {
		stepValue = step[0]
	}
	iteration := &spec.Iteration{
		Start:     &spec.Iteration_BlockNumber{BlockNumber: startBlock},
		Direction: spec.Iteration_Forward,
		Limit:     limit,
		Step:      stepValue,
	}

	sn.logger.Infow("Requesting block headers", "start", startBlock, "limit", limit, "step", stepValue)

	headersIt, err := sn.StarknetClient.RequestBlockHeaders(ctx, &spec.BlockHeadersRequest{Iteration: iteration})
	if err != nil {
		sn.logger.Errorw("Failed to request block headers", "error", err)
		return nil, errors.New("failed to request block headers")
	}

	var headers []*spec.BlockHeadersResponse
	for header := range iter.Seq[*spec.BlockHeadersResponse](headersIt) {
		headers = append(headers, header)
	}
	sn.logger.Infow("Received block headers", "count", len(headers))
	return headers, nil
}

func (sn *SyntheticNode) RequestEvents(ctx context.Context, req *spec.EventsRequest) (iter.Seq[*spec.EventsResponse], error) {
	sn.logger.Infow("Requesting events", "request", req)
	events, err := sn.StarknetClient.RequestEvents(ctx, req)
	if err != nil {
		sn.logger.Errorw("Failed to request events", "error", err)
		return nil, errors.New("failed to request events")
	}
	return events, nil
}

func (sn *SyntheticNode) Close() error {
	sn.logger.Infow("Closing synthetic node")

	return sn.Host.Close()
}

func ParsePeerAddress(address string) (peer.AddrInfo, error) {
	maddr, err := multiaddr.NewMultiaddr(address)
	if err != nil {
		return peer.AddrInfo{}, errors.New("invalid multiaddr")
	}

	addrInfo, err := peer.AddrInfoFromP2pAddr(maddr)
	if err != nil {
		return peer.AddrInfo{}, errors.New("invalid peer address")
	}

	return *addrInfo, nil
}

func (sn *SyntheticNode) WaitForPeerDiscovery(ctx context.Context, peerID peer.ID, timeout time.Duration) error {
	start := time.Now()

	err := retry.Do(
		func() error {
			if sn.isPeerConnected(peerID) {
				return nil
			}
			return errors.New("peer not connected")
		},
		retry.Attempts(uint(timeout/time.Second)),
		retry.Delay(1*time.Second),
		retry.DelayType(retry.FixedDelay),
		retry.Context(ctx),
	)

	duration := time.Since(start)
	if err == nil {
		sn.logger.Infow("Peer discovered", "peerID", peerID, "duration", duration)
	} else {
		sn.logger.Warnw("Peer discovery timed out", "peerID", peerID, "timeout", timeout)
	}

	return err
}

func (sn *SyntheticNode) isPeerConnected(peerID peer.ID) bool {
	return sn.Host.Network().Connectedness(peerID) == network.Connected
}

func (sn *SyntheticNode) contains(peers []peer.ID, id peer.ID) bool {
	for _, p := range peers {
		if p == id {
			return true
		}
	}
	return false
}
