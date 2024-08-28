package synthetic_node

import (
	"context"
	"crypto/rand"
	"fmt"
	"iter"
	"starknet-p2p-tests/config"
	"starknet-p2p-tests/protocol/p2p/starknet"
	"starknet-p2p-tests/protocol/p2p/starknet/spec"
	"starknet-p2p-tests/protocol/p2p/utils"

	"github.com/libp2p/go-libp2p"
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
	logger         *utils.TestSimpleLogger
	targetPeer     peer.ID
}

func New(ctx context.Context, logger *utils.TestSimpleLogger) (*SyntheticNode, error) {
	priv, _, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate key: %w", err)
	}

	opts := []libp2p.Option{
		libp2p.Identity(priv),
		libp2p.ListenAddrStrings(config.SyntheticListenAddrs...),
	}

	h, err := libp2p.New(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create libp2p node: %w", err)
	}

	logger.Infow("Created new synthetic node", "address", h.Addrs(), "id", h.ID())

	return &SyntheticNode{
		Host:   h,
		logger: logger,
	}, nil
}

func (sn *SyntheticNode) Connect(ctx context.Context, targetAddress string) error {
	targetPeerInfo, err := ParsePeerAddress(targetAddress)
	if err != nil {
		return fmt.Errorf("failed to parse peer address: %w", err)
	}

	sn.logger.Infow("Connecting to peer: ", "address", targetAddress)
	if err := sn.Host.Connect(ctx, targetPeerInfo); err != nil {
		return fmt.Errorf("failed to connect to target peer: %w", err)
	}

	sn.targetPeer = targetPeerInfo.ID
	networkInfo := &utils.Network{Name: config.NetworkName}
	newStreamFunc := func(ctx context.Context, pids ...protocol.ID) (network.Stream, error) {
		return sn.Host.NewStream(ctx, targetPeerInfo.ID, pids...)
	}

	sn.StarknetClient = starknet.NewClient(newStreamFunc, networkInfo, sn.logger)
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

	sn.logger.Infow("Requesting block headers",
		"start", startBlock,
		"limit", limit,
		"step", stepValue)

	headersIt, err := sn.StarknetClient.RequestBlockHeaders(ctx, &spec.BlockHeadersRequest{Iteration: iteration})
	if err != nil {
		return nil, fmt.Errorf("failed to request block headers:", err)
	}

	var headers []*spec.BlockHeadersResponse
	for header := range headersIt {
		headers = append(headers, header)
	}
	sn.logger.Infow("Received block headers", "count", len(headers))
	return headers, nil
}

func (sn *SyntheticNode) RequestEvents(ctx context.Context, req *spec.EventsRequest) (iter.Seq[*spec.EventsResponse], error) {
	sn.logger.Infow("Requesting events: %+v", req)
	events, err := sn.StarknetClient.RequestEvents(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to request events: %w", err)
	}
	return events, nil
}

func (sn *SyntheticNode) Close() error {
	sn.logger.Infow("Closing synthetic node")
	if sn.targetPeer != "" {
		sn.Host.Network().ClosePeer(sn.targetPeer)
	}
	return sn.Host.Close()
}

func ParsePeerAddress(address string) (peer.AddrInfo, error) {
	maddr, err := multiaddr.NewMultiaddr(address)
	if err != nil {
		return peer.AddrInfo{}, fmt.Errorf("invalid multiaddr: %w", err)
	}

	addrInfo, err := peer.AddrInfoFromP2pAddr(maddr)
	if err != nil {
		return peer.AddrInfo{}, fmt.Errorf("invalid peer address: %w", err)
	}

	return *addrInfo, nil
}
