# E2E Tests for P2P Syncing

This project contains end-to-end (E2E) tests for peer-to-peer (P2P) syncing with various node configurations. The planned configurations include:

| From \ To  | Juno | Pathfinder | Papyrus | Madara |
|------------|------|------------|---------|--------|
| Juno       | ✅    | ✅         | ❌      | ❌     |
| Pathfinder | ❌    | ✅         | ❌      | ❌     |
| Papyrus    | ❌    | ❌         | ❌      | ❌     |
| Madara     | ❌    | ❌         | ❌      | ❌     |

## Prerequisites

- [Kurtosis](https://docs.kurtosis.com/install) must be installed on your system.

## Running the Tests

To run the E2E tests, use the following command:

```
kurtosis run .
```

## Cleaning Up

After running the tests, you can clean up the environment with:

```
kurtosis clean --all
```

## Test Case Roadmap

We are planning to implement comprehensive E2E tests covering various P2P syncing scenarios across different node configurations. Below is the detailed roadmap:

### P2P Syncing Test Cases

- Juno syncing from Juno
- Juno syncing from Pathfinder
- Juno syncing from Papyrus
- Juno syncing from Madara

- Pathfinder syncing from Juno
- Pathfinder syncing from Pathfinder
- Pathfinder syncing from Papyrus
- Pathfinder syncing from Madara

- Papyrus syncing from Juno
- Papyrus syncing from Pathfinder
- Papyrus syncing from Papyrus
- Papyrus syncing from Madara

- Madara syncing from Juno
- Madara syncing from Pathfinder
- Madara syncing from Papyrus
- Madara syncing from Madara

### Advanced Test Scenarios

- **Network Partitioning**
  - Simulate network partitions to test node behavior and data consistency when parts of the network are isolated.
  - Verify recovery and data reconciliation once the network is restored.

- **Chaos Testing**
  - Introduce random failures and disruptions using tools like Chaos Monkey to test system resilience.
  - Assess how nodes handle unexpected shutdowns, network latency, and other disruptions.

- **High Latency and Packet Loss**
  - Test node performance and syncing capabilities under conditions of high network latency and packet loss.
  - Evaluate the impact on transaction throughput and data consistency.

- **Node Failover and Recovery**
  - Simulate node failures and test the system's ability to recover and maintain data integrity.
  - Verify that other nodes can take over the responsibilities of failed nodes.

- **Reorg Testing**
  - Create scenarios where blockchain reorgs occur to test node handling and data consistency.
  - Ensure nodes can correctly resolve forks and maintain a consistent state.

### Custom Network Setup

- **Madara Sequencer with Feeder Gateway**
  - Integrate Madara sequencer to run a private network, giving us full control over block production and network parameters.
  - Use the Feeder gateway to inject custom blocks and transactions.

- **Advantages of Custom Network**
  - Faster syncing due to control over block generation (e.g., fewer blocks).
  - Ability to create specific network conditions and scenarios.
  - Test reorgs and other edge cases that are hard to reproduce in a public network.

### Transaction Simulation

- **Gomu Gomu Gatling Tool**
  - Utilize the [Gomu Gomu Gatling](https://github.com/keep-starknet-strange/gomu-gomu-no-gatling) tool to simulate transactions on the network.
  - Stress-test nodes under high transaction volumes to assess performance and stability.

### Test Execution Considerations

- **Test Duration**
  - Acknowledge that some tests are long-running due to the nature of full node syncing.
  - Explore options to optimize test duration, such as adjusting block times or limiting the number of blocks.

- **Test Environment**
  - Investigate suitable environments for running tests (e.g., dedicated servers, cloud infrastructure).
  - Ensure resources are sufficient to handle multiple nodes and simulations concurrently.

### Publishing Test Results

- **Public Test Results Page**
  - Create a public-facing page or dashboard to publish test results.
  - Include statuses of each test case, logs, and any relevant metrics.
  - Update the page automatically after each test run to provide up-to-date information.
