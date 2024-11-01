participants = import_module("../../clients/participants.star")

# Test configuration
SYNC_TIMEOUT_SECONDS = 1200
TARGET_BLOCK_NUMBER = 1000

def run(plan):
    # Run the Pathfinder feeder node
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

    tester_image = plan.add_service(
        "sync-tester",
        config=ServiceConfig(
            image=ImageBuildSpec(
                image_name="sync-test",
                build_context_dir="./../../tester",
            ),
        )
    )

    # Run the tester
    plan.print("Starting the sync tester...")
    plan.exec(tester_image.name, ExecRecipe(
        ["node", "index.mjs", "http://" + peer_node.ip_address + ":6061", str(SYNC_TIMEOUT_SECONDS), str(TARGET_BLOCK_NUMBER)]
    ))

    plan.print("Juno from Pathfinder sync test completed")