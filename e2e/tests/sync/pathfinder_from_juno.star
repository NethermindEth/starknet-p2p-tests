participants = import_module("../../clients/participants.star")
sync_utils = import_module("./sync_test_utils.star")

# Test configuration
SYNC_TIMEOUT_SECONDS = 1800
TARGET_BLOCK_NUMBER = 1000
SOURCE_RPC_PORT = 6060  # Juno's RPC port
TARGET_RPC_PORT = 9545  # Pathfinder's RPC port

def run(plan):
    # Run the Juno feeder node
    feeder_node = participants.run_participant(plan, "juno-feeder", {
        "type": "juno",
        "is_feeder": True,
        "private_key": "67f8eae550a5265238431d719c2b62163011ab2a3f2ebeee3bc8f3135e2e2500b9e2c2e9e4ebeea82cca787094d74ab6fcae8ec0367e866dc1130de89e37150b",
        "http_port": SOURCE_RPC_PORT,
        "network": "sepolia"
    })

    # Run the Pathfinder node with the juno node as a peer
    peer_node = participants.run_participant(plan, "pathfinder-peer", {
        "type": "pathfinder",
        "is_feeder": False,
        "http_port": TARGET_RPC_PORT,
        "p2p_port": 20003,
        "peer_multiaddrs": ["/ip4/" + feeder_node.ip_address + "/tcp/7777/p2p/12D3KooWNKz9BJmyWVFUnod6SQYLG4dYZNhs3GrMpiot63Y1DLYS"]
    })

    sync_utils.run_sync_test(
        plan,
        feeder_node,
        peer_node,
        SOURCE_RPC_PORT,
        TARGET_RPC_PORT,
        SYNC_TIMEOUT_SECONDS,
        TARGET_BLOCK_NUMBER
    )
    plan.print("Pathfinder from Juno sync test completed")