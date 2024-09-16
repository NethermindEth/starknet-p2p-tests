package conformance

import (
	"context"
	"encoding/hex"
	"starknet-p2p-tests/config"
	"starknet-p2p-tests/protocol/p2p/starknet/spec"
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
		{name: "Basic Consecutive Blocks", startBlock: 1, limit: 5, step: 1},
		{name: "Non-unit Step", startBlock: 1, limit: 5, step: 2},
		{name: "Non-standard Start and Limit", startBlock: 10, limit: 3, step: 1},
		{name: "Large Step", startBlock: 1, limit: 5, step: 10},
		{name: "High Start Block", startBlock: 100, limit: 5, step: 1},
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

	syntheticNode, err := synthetic_node.New(ctx, t)
	require.NoError(t, err, "Failed to create synthetic node")
	require.NoError(t, syntheticNode.Connect(ctx, config.TargetPeerAddress), "Failed to connect to target peer")
	headers, err := syntheticNode.RequestBlockHeaders(ctx, tc.startBlock, tc.limit, tc.step)
	require.NoError(t, err, "Failed to request block headers")

	validateBlockHeaders(t, headers, tc)
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

	assert.NotNil(t, header.BlockHash, "BlockHash should not be nil")
	assert.Len(t, header.BlockHash.Elements, 32, "BlockHash should be 32 bytes long")

	assert.NotNil(t, header.ParentHash, "ParentHash should not be nil")
	assert.Len(t, header.ParentHash.Elements, 32, "ParentHash should be 32 bytes long")

	assert.Greater(t, header.Number, uint64(0), "Block number should be greater than 0")

	assert.Greater(t, header.Time, uint64(0), "Block time should be greater than 0")
	assert.LessOrEqual(t, header.Time, uint64(time.Now().Unix()), "Block time should not be in the future")

	assert.NotNil(t, header.SequencerAddress, "SequencerAddress should not be nil")
	assert.Len(t, header.SequencerAddress.Elements, 32, "SequencerAddress should be 32 bytes long")

	assert.NotNil(t, header.StateRoot, "StateRoot should not be nil")
	assert.Len(t, header.StateRoot.Elements, 32, "StateRoot should be 32 bytes long")

	assert.NotNil(t, header.StateDiffCommitment, "StateDiffCommitment should not be nil")
	assert.Greater(t, header.StateDiffCommitment.StateDiffLength, uint64(0), "StateDiffLength should be greater than 0")
	assert.NotNil(t, header.StateDiffCommitment.Root, "StateDiffCommitment Root should not be nil")

	assert.NotNil(t, header.Transactions, "Transactions should not be nil")
	assert.Greater(t, header.Transactions.NLeaves, uint64(0), "Transactions NLeaves should be greater than 0")

	assert.NotNil(t, header.Events, "Events should not be nil")

	assert.NotEmpty(t, header.ProtocolVersion, "ProtocolVersion should not be empty")

	assert.NotNil(t, header.GasPriceFri, "GasPriceFri should not be nil")
	assert.NotNil(t, header.GasPriceWei, "GasPriceWei should not be nil")
	assert.NotNil(t, header.DataGasPriceFri, "DataGasPriceFri should not be nil")
	assert.NotNil(t, header.DataGasPriceWei, "DataGasPriceWei should not be nil")

	assert.NotEmpty(t, header.Signatures, "Signatures should not be empty")
	for i, sig := range header.Signatures {
		assert.NotNil(t, sig, "Signature %d should not be nil", i)
		assert.NotNil(t, sig.R, "Signature %d R should not be nil", i)
		assert.NotNil(t, sig.S, "Signature %d S should not be nil", i)
	}

	t.Logf("Block Number: %d, Time: %s, Protocol Version: %s",
		header.Number,
		time.Unix(int64(header.Time), 0).UTC(),
		header.ProtocolVersion)
}
