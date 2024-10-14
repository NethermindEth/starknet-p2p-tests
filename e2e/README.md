# E2E Tests for P2P Syncing

This project contains end-to-end (E2E) tests for peer-to-peer (P2P) syncing with various node configurations. The planned configurations include:

| From \ To | Juno | Pathfinder | Papyrus | Dexoyss |
|-----------|------|------------|---------|---------|
| Juno      | ✅    | ❌         | ❌      | ❌      |
| Pathfinder| ❌    | ❌         | ❌      | ❌      |
| Papyrus   | ❌    | ❌         | ❌      | ❌      |
| Dexoyss   | ❌    | ❌         | ❌      | ❌      |


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
