participants = import_module("../../clients/participants.star")
sync_utils = import_module("./sync_test_utils.star")

# Test configuration
SYNC_TIMEOUT_SECONDS = 1800
TARGET_BLOCK_NUMBER = 1000

def run(plan):
    # Run the Pathfinder as feeder node
    feeder_node = participants.run_participant(plan, "pathfinder-feeder", {
        "type": "pathfinder",
        "is_feeder": True,
        "p2p_port": 20002,
    })

    # Run the juno peer node with the feeder node as a peer
    peer_node = participants.run_participant(plan, "juno-peer", {
        "type": "juno",
        "is_feeder": False,
        "private_key": "a5a938ae6f012390fd68a10d3dd91038334fe5f0ed1c96753a3ee7bf0e8f1314e39307aea916c94e2d07e616fa20e315f4625c4f1e598ba2cc589410cc9c5cda",
        "http_port": 6061,
        "peer_multiaddrs": ["/ip4/" + feeder_node.ip_address + "/tcp/20002/p2p/12D3KooWFY6SaqJkRxJDepwvBi4Rw36iMUGZrejW69qkjYQQ2ydQ"],
        "network": "sepolia",
    })

    sync_utils.run_sync_test(plan, feeder_node, peer_node, SYNC_TIMEOUT_SECONDS, TARGET_BLOCK_NUMBER)
    plan.print("Juno from Pathfinder sync test completed")
