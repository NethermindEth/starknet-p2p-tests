participants = import_module("../../clients/participants.star")

def run(plan):
    # Run the Juno feeder node
    feeder_node = participants.run_participant(plan, "juno-feeder", {
        "type": "juno",
        "is_feeder": True,
        "private_key": "67f8eae550a5265238431d719c2b62163011ab2a3f2ebeee3bc8f3135e2e2500b9e2c2e9e4ebeea82cca787094d74ab6fcae8ec0367e866dc1130de89e37150b",
        "http_port": 6060,
    })

    # Run the Pathfinder node with the juno node as a peer
    peer_node = participants.run_participant(plan, "pathfinder-peer", {
        "type": "pathfinder",
        "is_feeder": False,
        "p2p_port": 20003,
        "peer_multiaddrs": ["/ip4/" + feeder_node.ip_address + "/tcp/7777/p2p/12D3KooWNKz9BJmyWVFUnod6SQYLG4dYZNhs3GrMpiot63Y1DLYS"]
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

    plan.print("Pathfinder from Juno sync test completed")