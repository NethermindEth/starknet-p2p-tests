participants = import_module("../../clients/participants.star")

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
    plan.print("http://" + peer_node.ip_address + ":9545" + "/rpc/v0_7")
    plan.exec(tester_image.name, ExecRecipe(
        ["node", "index.mjs", "http://" + peer_node.ip_address + ":9545" + "/rpc/v0_7", "120", "10"]
    ))

    plan.print("Pathfinder from Pathfinder sync test completed")