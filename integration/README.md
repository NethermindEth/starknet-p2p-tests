# Starknet P2P Tests

This project contains peer-to-peer (P2P) tests for the Starknet network. It is inspired by Ziggurat (https://github.com/runziggurat) but written in Go.

## Purpose

The main purposes of these tests are:

1. To ensure reliable and efficient communication between nodes in the Starknet network.
2. To validate the correct implementation of the P2P protocol used in Starknet.
3. To identify and address any potential issues or vulnerabilities in the P2P layer.
4. To assess the conformance, performance, and resilience of the Starknet P2P network.

## Test Case List & Coverage

The project outlines the following test cases, based on the protocols described in the [Starknet p2p spec](https://github.com/starknet-io/starknet-p2p-specs/blob/main/p2p/proto/protocols.md#protocols-briefing). These are intended to guide the development of comprehensive tests:

### Conformance Tests

- **Headers Protocol** (`/starknet/headers/0.1.0-rc.0`):
  - Requesting block headers with valid parameters. ✅
  - Handling invalid or malformed `BlockHeadersRequest` messages. ✅
  - Requesting headers for non-existent blocks. ✅
  - Testing step and limit parameters in header requests. ✅
  - Validating the consistency of returned block headers.

- **StateDiffs Protocol** (`/starknet/state_diffs/0.1.0-rc.0`):
  - Fetching state diffs for specific block ranges.
  - Handling requests with invalid block numbers or hashes.
  - Verifying the integrity and correctness of `StateDiffsResponse`.
  - Testing response to requests for future blocks.

- **Classes Protocol** (`/starknet/classes/0.1.0-rc.0`):
  - Retrieving class definitions using valid class hashes.
  - Handling requests with unknown or invalid class hashes.
  - Validating the structure and content of `ClassesResponse`.
  - Testing responses to duplicate requests.

- **Transactions Protocol** (`/starknet/transactions/0.1.0-rc.0`):
  - Requesting transactions by block number, block hash, and transaction hash.
  - Handling invalid transaction requests or unknown identifiers.
  - Verifying support for all transaction types.
  - Testing batch transaction retrieval.

- **Events Protocol** (`/starknet/events/0.1.0-rc.0`):
  - Subscribing to events with valid filters.
  - Handling invalid subscription requests.
  - Receiving and validating event data streams.
  - Testing unsubscription and cleanup procedures.

- **Kademlia (Discovery) Protocol** (`/starknet/kad/<chain_id>/1.0.0`):
  - Discovering peers using Kademlia routing. ✅

- **Identify Protocol** (`/ipfs/id/1.0.0`):
  - Retrieving peer identification information.
  - Validating protocol versions and agent strings.
  - Handling malformed or unexpected identify responses.
  - Testing compatibility with different node implementations.

### Performance Tests

- **Headers Protocol**:
  - Benchmarking header retrieval for large block ranges.
  - Measuring latency and throughput under high request rates.
  - Testing performance impact of varying `step` and `limit` parameters.

- **StateDiffs Protocol**:
  - Assessing performance when fetching large state diffs.
  - Evaluating resource usage during intensive state diff requests.
  - Testing concurrency handling with multiple simultaneous requests.

- **Transactions Protocol**:
  - Evaluating transaction retrieval under high load.
  - Measuring response times for bulk transaction requests.
  - Testing scalability with increasing transaction history.

### Resilience Tests

- **Protocol Version Mismatch**:
  - Testing interactions between nodes running different protocol versions.
  - Validating proper negotiation and graceful degradation.

- **Malicious Input Handling**:
  - Sending malformed or malicious messages to test security.
  - Observing node responses to potential attack vectors.

- **Resource Limits**:
  - Testing node behavior under resource constraints (e.g., memory, CPU).
  - Assessing recovery from resource exhaustion scenarios.

## Getting Started

### Prerequisites

- Go 1.23 or later

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/NethermindEth/starknet-p2p-tests.git
   ```

2. Navigate to the project directory:
   ```
   cd starknet-p2p-tests
   ```

3. Install dependencies:
   ```
   go mod tidy
   ```

## Configuration

The project uses environment variables for configuration, which are defined in the `config/config.go` file. 

[More info](#environment-variables)

## Running Tests

Before running the tests, ensure that your environment is properly configured:

1. Set the necessary environment variables (e.g., `TARGET_PEER_ADDRESS`) if you want to override the defaults.

2. To run all tests:
   ```
   go test ./tests/...
   ```
   For verbose output:
   ```
   go test -v ./tests/...
   ```

3. To run specific test categories:
   ```
   go test ./tests/conformance
   go test ./tests/performance
   ```

4. For performance tests, you may need to increase the default timeout:
   ```
   go test -timeout 5m ./tests/performance
   ```
   
## Environment Variables

The following environment variables are available for configuration:

- `TARGET_PEER_ADDRESS`: The address of the target Starknet node (default: "/ip4/35.237.66.77/tcp/7777/p2p/12D3KooWR8ikUDiinyE5wgdYiqsdLfJRsBDYKGii6L3oyoipVEaV")
- `DEFAULT_TEST_TIMEOUT`: The default timeout for tests (default: "30s")
- `SYNTHETIC_LISTEN_ADDRS`: The listen addresses for the synthetic node (default: "/ip4/0.0.0.0/tcp/0")
- `NETWORK_NAME`: The name of the Starknet network being tested (default: "sepolia")


Note: For performance tests, you may need to increase the default timeout. Use the -timeout flag, e.g., `go test -timeout 5m ./tests/performance`

## Sample Test Output

```
=== RUN   TestSyntheticNodeMultipleBlockHeadersRequest/Basic_Consecutive_Blocks
    utils.go:48: INFO: Created new synthetic node address=[/ip4/127.0.0.1/tcp/63979 /ip4/192.168.215.4/tcp/63979] id=12D3KooWC2xdz7kxnsLMCDY851hqvtNcsmMuS8xQRf2KG1MjKByb
    utils.go:48: INFO: Connecting to peer:  address=/ip4/35.237.66.77/tcp/7777/p2p/12D3KooWR8ikUDiinyE5wgdYiqsdLfJRsBDYKGii6L3oyoipVEaV
    utils.go:48: INFO: Successfully connected to peer id=12D3KooWR8ikUDiinyE5wgdYiqsdLfJRsBDYKGii6L3oyoipVEaV
    getBlock_test.go:71: Requesting headers from block 1, limit 5, step 1
    utils.go:48: INFO: Requesting block headers start=1 limit=5 step=1
    utils.go:48: INFO: Received block headers count=6
    getBlock_test.go:76: Received 6 responses
    getBlock_test.go:108: Header 0: Block Number: 1, BlockHash: 0x065f21e3bf9acb006889a6a0cfa53811204cb2a08e830bf61238547a454471af     
    getBlock_test.go:113: Block Number: 1, Time: 2023-11-20 10:05:24 +0000 UTC, Protocol Version: 0.12.3
    getBlock_test.go:108: Header 1: Block Number: 2, BlockHash: 0x06a592792d61b9eddbf7686bff0cdc318b5875efcbb363fd0cd4308a1a07660d     
    getBlock_test.go:113: Block Number: 2, Time: 2023-11-20 10:19:41 +0000 UTC, Protocol Version: 0.12.3
    getBlock_test.go:108: Header 2: Block Number: 3, BlockHash: 0x00f8648fbf10da5adca08c3307dd221fa19da3fe8e339ade0452545e9e643c31     
    getBlock_test.go:113: Block Number: 3, Time: 2023-11-20 10:57:44 +0000 UTC, Protocol Version: 0.12.3
    getBlock_test.go:108: Header 3: Block Number: 4, BlockHash: 0x05001dd2dd55c1cdfc347f7d353f63f245ad7dbba56b4bc720a6c80841156cb6     
    getBlock_test.go:113: Block Number: 4, Time: 2023-11-20 11:34:31 +0000 UTC, Protocol Version: 0.12.3
    getBlock_test.go:108: Header 4: Block Number: 5, BlockHash: 0x058b18d1ff10100a677e324d1b26dc2aa07e07f8e0d44e18ea3fc3e8aec1e323     
    getBlock_test.go:113: Block Number: 5, Time: 2023-11-20 12:08:02 +0000 UTC, Protocol Version: 0.12.3
    getBlock_test.go:95: Received Fin message at position 5
    utils.go:34: INFO: Closing synthetic node

```