participants = import_module("../../clients/participants.star")
sync_utils = import_module("./sync_test_utils.star")

# Test configuration
SYNC_TIMEOUT_SECONDS = 1800
TARGET_BLOCK_NUMBER = 1000
FEEDER_RPC_PORT = 6060
PEER_RPC_PORT = 6061

def run(plan):
    # Run the Juno feeder node
    feeder_node = participants.run_participant(plan, "juno-feeder", {
        "type": "juno",
        "is_feeder": True,
        "private_key": "67f8eae550a5265238431d719c2b62163011ab2a3f2ebeee3bc8f3135e2e2500b9e2c2e9e4ebeea82cca787094d74ab6fcae8ec0367e866dc1130de89e37150b",
        "http_port": FEEDER_RPC_PORT,
        "network": "sepolia"
    })

    # Run the juno peer node with the feeder node as a peer
    peer_node = participants.run_participant(plan, "juno-peer", {
        "type": "juno",   
        "is_feeder": False,
        "private_key": "a5a938ae6f012390fd68a10d3dd91038334fe5f0ed1c96753a3ee7bf0e8f1314e39307aea916c94e2d07e616fa20e315f4625c4f1e598ba2cc589410cc9c5cda",
        "http_port": PEER_RPC_PORT,
        "peer_multiaddrs": ["/ip4/" + feeder_node.ip_address + "/tcp/7777/p2p/12D3KooWNKz9BJmyWVFUnod6SQYLG4dYZNhs3GrMpiot63Y1DLYS"],
        "network": "sepolia",
    })

    sync_utils.run_sync_test(
        plan, 
        feeder_node, 
        peer_node, 
        FEEDER_RPC_PORT,
        PEER_RPC_PORT, 
        SYNC_TIMEOUT_SECONDS, 
        TARGET_BLOCK_NUMBER
    )
    plan.print("Juno to Juno sync test completed")