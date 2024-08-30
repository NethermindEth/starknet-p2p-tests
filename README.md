# Starknet P2P Tests

This project contains peer-to-peer (P2P) tests for the Starknet network. It is inspired by Ziggurat (https://github.com/runziggurat) but written in Go.

## Overview

Starknet is a permissionless decentralized Validity-Rollup (often referred to as ZK-Rollup). It operates as a Layer 2 network over Ethereum, enabling any dApp to achieve unlimited scale for its computation without compromising Ethereum's composability and security.

This project focuses on testing the peer-to-peer communication aspects of the Starknet network, with a particular emphasis on conformance, performance, and resilience tests.

## Purpose

The main purposes of these tests are:

1. To ensure reliable and efficient communication between nodes in the Starknet network.
2. To validate the correct implementation of the P2P protocol used in Starknet.
3. To identify and address any potential issues or vulnerabilities in the P2P layer.
4. To assess the conformance, performance, and resilience of the Starknet P2P network.

## Project Structure

This project is written in Go and organized as follows:

- `tests/`: Contains all test files
  - `conformance/`: Conformance tests
  - `performance/`: Performance tests
  - `resilience/`: Resilience tests (to be implemented)
- `config/`: Contains configuration files
  - `config.go`: Defines configuration options and environment variables
- `tools/`: Contains utility tools
  - `synthetic_node.go`: Implements the synthetic node for testing

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

## Environment Variables

- `TARGET_PEER_ADDRESS`: The address of the target Starknet node to connect to for testing.
- `DEFAULT_TEST_TIMEOUT`: The default timeout for tests.
- `SYNTHETIC_LISTEN_ADDRS`: The listen addresses for the synthetic node.
- `NETWORK_NAME`: The name of the Starknet network being tested.

## Synthetic Node

The project implements a synthetic node in `tools/synthetic_node.go`. This synthetic node is used to simulate a Starknet node for testing purposes. It provides capabilities such as:

- Connecting to a target Starknet node
- Requesting block headers
- ...
- ...
- ...

The synthetic node is a key component in our testing infrastructure, allowing us to interact with the Starknet network in a controlled manner.

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

You can set these variables in your shell before running the tests to override the default values.

- To run all tests:
  ```
  go test ./tests/...
  ```
  For verbose output:
  ```
  go test -v ./tests/...
  ```

- To run specific test categories:
  ```
  go test ./tests/conformance
  go test ./tests/performance
  ```

Note: Some tests may require additional setup or configuration. Please check the output and individual test files for any specific requirements.

Note: For performance tests, you may need to increase the default timeout. Use the -timeout flag, e.g., `go test -timeout 5m ./tests/performance`

## Test Cases

The project currently includes the following test cases:

1. Conformance Tests:
   - `getBlock_test.go`: Tests the conformance of the getBlock functionality.
   - ...
   - ...
   - ...
2. Performance Tests:
   - `getBlocks_test.go`: Evaluates the performance of retrieving multiple blocks.
   - ...
   - ...
   - ...

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
## Contributing

We welcome contributions to improve and expand these tests. Please read our contributing guidelines before submitting pull requests.



## Contact

For any questions or concerns regarding these Starknet P2P tests written in Go, please open an issue in this repository.