participants = import_module("../clients/participants.star")

def run_juno_to_juno_sync(plan):
    # Run the Juno feeder node
    feeder_node = participants.run_participant(plan, "juno-feeder", {
        "type": "juno",
        "is_feeder": True,
        "private_key": "67f8eae550a5265238431d719c2b62163011ab2a3f2ebeee3bc8f3135e2e2500b9e2c2e9e4ebeea82cca787094d74ab6fcae8ec0367e866dc1130de89e37150b",
        "http_port": 6060,
        "image": "nethermindeth/juno:fix-p2p-issues-for-0-13-2-block-sync"
    })

    # Run the juno peer node with the feeder node as a peer
    peer_node = participants.run_participant(plan, "juno-peer", {
        "type": "juno",
        "is_feeder": False,
        "private_key": "a5a938ae6f012390fd68a10d3dd91038334fe5f0ed1c96753a3ee7bf0e8f1314e39307aea916c94e2d07e616fa20e315f4625c4f1e598ba2cc589410cc9c5cda",
        "http_port": 6061,
        "image": "nethermindeth/juno:fix-p2p-issues-for-0-13-2-block-sync",
        "peer_multiaddrs": ["/ip4/" + feeder_node.ip_address + "/tcp/7777/p2p/12D3KooWNKz9BJmyWVFUnod6SQYLG4dYZNhs3GrMpiot63Y1DLYS"]
    })

    tester_image = plan.add_service(
        "sync-tester",
        config=ServiceConfig(
            image=ImageBuildSpec(
                image_name="sync-test",
                build_context_dir="./../tester",
            ),
        )
    )

    # Run the tester
    plan.print("Starting the sync tester...")
    plan.exec(tester_image.name, ExecRecipe(
        ["node", "index.mjs", "http://" + feeder_node.ip_address + ":6060", "http://" + peer_node.ip_address + ":6061"]
    ))

    plan.print("Juno to Juno sync test completed")

def run_all_sync_tests(plan):
    run_juno_to_juno_sync(plan)
