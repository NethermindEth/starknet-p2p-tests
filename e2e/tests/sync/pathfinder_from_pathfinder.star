participants = import_module("../../clients/participants.star")
sync_utils = import_module("./sync_test_utils.star")

# Test configuration
SYNC_TIMEOUT_SECONDS = 1800
TARGET_BLOCK_NUMBER = 1000
RPC_PORT = 9545

def run(plan):
    # Run the Pathfinder feeder node
    feeder_node = participants.run_participant(plan, "pathfinder-feeder", {
        "type": "pathfinder",
        "is_feeder": True,
        "p2p_port": 20002,
    })

    # Run the Pathfinder peer node with the feeder node as a peer
    peer_node = participants.run_participant(plan, "pathfinder-peer", {
        "type": "pathfinder",
        "is_feeder": False,
        "p2p_port": 20003,
        "peer_multiaddrs": ["/ip4/" + feeder_node.ip_address + "/tcp/20002/p2p/12D3KooWFY6SaqJkRxJDepwvBi4Rw36iMUGZrejW69qkjYQQ2ydQ"]
    })

    sync_utils.run_sync_test(plan, feeder_node, peer_node, RPC_PORT, SYNC_TIMEOUT_SECONDS, TARGET_BLOCK_NUMBER)
    plan.print("Pathfinder from Pathfinder sync test completed")