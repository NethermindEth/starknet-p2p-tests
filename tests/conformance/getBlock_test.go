package tests

import (
	"context"
	"encoding/hex"
	"fmt"
	"starknet-p2p-tests/config"
	"starknet-p2p-tests/protocol/p2p/starknet/spec"
	"starknet-p2p-tests/protocol/p2p/utils"
	synthetic_node "starknet-p2p-tests/tools"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type blockHeaderTestCase struct {
	name       string
	startBlock uint64
	limit      uint64
	step       uint64
}

func TestSyntheticNodeMultipleBlockHeadersRequest(t *testing.T) {
	testCases := []blockHeaderTestCase{
		{"Basic Consecutive Blocks", 1, 5, 1},
		{"Non-unit Step", 1, 5, 2},
		{"Non-standard Start and Limit", 10, 3, 1},
		{"Large Step", 1, 5, 10},
		{"High Start Block", 100, 5, 1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			runBlockHeaderTest(t, tc)
		})
	}
}

func runBlockHeaderTest(t *testing.T, tc blockHeaderTestCase) {
	ctx, cancel := context.WithTimeout(context.Background(), config.DefaultTestTimeout)
	defer cancel()

	syntheticNode, err := setupSyntheticNode(t)
	require.NoError(t, err, "Failed to set up synthetic node")
	defer syntheticNode.Close()

	headers, err := requestBlockHeaders(ctx, t, syntheticNode, tc)
	require.NoError(t, err, "Error requesting block headers")

	validateBlockHeaders(t, headers, tc)
}

func setupSyntheticNode(t *testing.T) (*synthetic_node.SyntheticNode, error) {
	logger := &utils.TestSimpleLogger{Logger: t.Logf}
	syntheticNode, err := synthetic_node.New(context.Background(), logger)
	if err != nil {
		return nil, fmt.Errorf("failed to create synthetic node: %w", err)
	}

	err = syntheticNode.Connect(context.Background(), config.TargetPeerAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to target peer: %w", err)
	}

	return syntheticNode, nil
}

func requestBlockHeaders(ctx context.Context, t *testing.T, syntheticNode *synthetic_node.SyntheticNode, tc blockHeaderTestCase) ([]*spec.BlockHeadersResponse, error) {
	t.Logf("Requesting headers from block %d, limit %d, step %d", tc.startBlock, tc.limit, tc.step)
	headers, err := syntheticNode.RequestBlockHeaders(ctx, tc.startBlock, tc.limit, tc.step)
	if err != nil {
		return nil, fmt.Errorf("failed to request block headers: %w", err)
	}
	t.Logf("Received %d responses", len(headers))
	return headers, nil
}

func validateBlockHeaders(t *testing.T, headers []*spec.BlockHeadersResponse, tc blockHeaderTestCase) {
	actualHeaders := 0
	var lastBlockNum uint64

	for i, header := range headers {
		switch msg := header.HeaderMessage.(type) {
		case *spec.BlockHeadersResponse_Header:
			actualHeaders++
			if msg.Header != nil {
				validateHeader(t, msg.Header, tc, i)
				lastBlockNum = msg.Header.Number
			} else {
				t.Errorf("Header %d is nil", i)
			}
		case *spec.BlockHeadersResponse_Fin:
			t.Logf("Received Fin message at position %d", i)
		default:
			t.Errorf("Unknown message type at position %d: %T", i, msg)
		}
	}

	assert.Equal(t, int(tc.limit), actualHeaders, "Expected to receive %d headers", tc.limit)
	assert.Equal(t, tc.startBlock+(tc.limit-1)*tc.step, lastBlockNum, "Unexpected last block number")
}

func validateHeader(t *testing.T, header *spec.SignedBlockHeader, tc blockHeaderTestCase, index int) {
	blockNum := header.Number
	blockHash := hex.EncodeToString(header.BlockHash.Elements)
	t.Logf("Header %d: Block Number: %d, BlockHash: 0x%s", index, blockNum, blockHash)

	expectedBlockNum := tc.startBlock + uint64(index)*tc.step
	assert.Equal(t, expectedBlockNum, blockNum, "Unexpected block number for step %d", tc.step)

	assertValidHeader(t, header)
}

func assertValidHeader(t *testing.T, header *spec.SignedBlockHeader) {
	t.Helper()

	// Check BlockHash
	assert.NotNil(t, header.BlockHash, "BlockHash should not be nil")
	assert.Len(t, header.BlockHash.Elements, 32, "BlockHash should be 32 bytes long")

	// Check ParentHash
	assert.NotNil(t, header.ParentHash, "ParentHash should not be nil")
	assert.Len(t, header.ParentHash.Elements, 32, "ParentHash should be 32 bytes long")

	// Check Number
	assert.Greater(t, header.Number, uint64(0), "Block number should be greater than 0")

	// Check Time
	assert.Greater(t, header.Time, uint64(0), "Block time should be greater than 0")
	assert.LessOrEqual(t, header.Time, uint64(time.Now().Unix()), "Block time should not be in the future")

	// Check SequencerAddress
	assert.NotNil(t, header.SequencerAddress, "SequencerAddress should not be nil")
	assert.Len(t, header.SequencerAddress.Elements, 32, "SequencerAddress should be 32 bytes long")

	// Check StateRoot
	assert.NotNil(t, header.StateRoot, "StateRoot should not be nil")
	assert.Len(t, header.StateRoot.Elements, 32, "StateRoot should be 32 bytes long")

	// Check StateDiffCommitment
	assert.NotNil(t, header.StateDiffCommitment, "StateDiffCommitment should not be nil")
	assert.Greater(t, header.StateDiffCommitment.StateDiffLength, uint64(0), "StateDiffLength should be greater than 0")
	assert.NotNil(t, header.StateDiffCommitment.Root, "StateDiffCommitment Root should not be nil")

	// Check Transactions
	assert.NotNil(t, header.Transactions, "Transactions should not be nil")
	assert.Greater(t, header.Transactions.NLeaves, uint64(0), "Transactions NLeaves should be greater than 0")

	// Check Events
	assert.NotNil(t, header.Events, "Events should not be nil")

	// Check ProtocolVersion
	assert.NotEmpty(t, header.ProtocolVersion, "ProtocolVersion should not be empty")

	// Check GasPrices
	assert.NotNil(t, header.GasPriceFri, "GasPriceFri should not be nil")
	assert.NotNil(t, header.GasPriceWei, "GasPriceWei should not be nil")
	assert.NotNil(t, header.DataGasPriceFri, "DataGasPriceFri should not be nil")
	assert.NotNil(t, header.DataGasPriceWei, "DataGasPriceWei should not be nil")

	// Check Signatures
	assert.NotEmpty(t, header.Signatures, "Signatures should not be empty")
	for i, sig := range header.Signatures {
		assert.NotNil(t, sig, "Signature %d should not be nil", i)
		assert.NotNil(t, sig.R, "Signature %d R should not be nil", i)
		assert.NotNil(t, sig.S, "Signature %d S should not be nil", i)
	}

	// Log some key information
	t.Logf("Block Number: %d, Time: %s, Protocol Version: %s",
		header.Number,
		time.Unix(int64(header.Time), 0).UTC(),
		header.ProtocolVersion)
}
