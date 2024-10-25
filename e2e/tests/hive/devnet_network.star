participants = import_module("../../clients/participants.star")

# Network and node configuration
CHAIN_ID = "SN_WOJO"
BLOCK_TIME = "30s"
L1_CHAIN_ID = "31337"
JUNO_P2P_PORT = "7777"
JUNO_API_PORT = "6060"
MADARA_PORT = "8080"
# Private key for Juno feeder node
TEST_PRIVATE_KEY = "67f8eae550a5265238431d719c2b62163011ab2a3f2ebeee3bc8f3135e2e2500b9e2c2e9e4ebeea82cca787094d74ab6fcae8ec0367e866dc1130de89e37150b"
# Peer ID that second Juno node will use to discover and connect to the feeder node
TEST_PEER_ID = "12D3KooWNKz9BJmyWVFUnod6SQYLG4dYZNhs3GrMpiot63Y1DLYS"

def setup_madara_node(plan):
    # Start Madara as a devnet sequencer with feeder gateway enabled
    # This node will produce blocks that Juno nodes will sync
    return participants.run_participant(plan, "madara-devnet", {
        "type": "madara",
        "extra_args": [
            "--devnet",
            "--override-devnet-chain-id",
            "--chain-config-override=chain_id={},block_time={}".format(CHAIN_ID, BLOCK_TIME),
            "--feeder-gateway-enable",
            "--gateway-external"
        ]
    })

def setup_juno_feeder_node(plan, madara_node):
    # Start first Juno node that syncs directly from Madara's feeder gateway
    # This node will propagate blocks to other Juno nodes via P2P
    madara_url = "http://{}:{}".format(madara_node.ip_address, MADARA_PORT)
    return participants.run_participant(plan, "juno-p2p-feeder-node", {
        "type": "juno",
        "is_feeder": True,
        "private_key": TEST_PRIVATE_KEY,
        "extra_args": [
            "--cn-name", "devnet",
            "--cn-feeder-url", "{}/feeder_gateway/".format(madara_url),
            "--cn-gateway-url", "{}/gateway/".format(madara_url),
            "--cn-l1-chain-id", L1_CHAIN_ID,
            "--cn-l2-chain-id", CHAIN_ID,
            "--cn-core-contract-address", "0x0000000000000000000000000000000000000000",
            "--cn-unverifiable-range=0,0"
        ]
    })

def setup_tester_service(plan):
    # Setup service that will verify block synchronization on Juno nodes
    return plan.add_service(
        "devnet-tester",
        config=ServiceConfig(
            image=ImageBuildSpec(
                image_name="sync-test",
                build_context_dir="./../../tester",
            ),
        )
    )

def run_sync_test(plan, tester_image, node_ip, message=""):
    # Verify that a Juno node has properly synced blocks
    if message:
        plan.print(message)
    plan.exec(tester_image.name, ExecRecipe(
        ["node", "index.mjs", "http://{}:{}".format(node_ip, JUNO_API_PORT), "10", "0"]
    ))

def setup_juno_peer_node(plan, feeder_node):
    # Start second Juno node that syncs blocks via P2P from the feeder node
    # This node doesn't connect to Madara directly
    return participants.run_participant(plan, "juno-peer-node", {
        "type": "juno",
        "peer_multiaddrs": ["/ip4/{}/tcp/{}/p2p/{}".format(
            feeder_node.ip_address, 
            JUNO_P2P_PORT, 
            TEST_PEER_ID
        )]
    })

def run(plan):
    # Test flow:
    # 1. Start Madara sequencer with feeder gateway
    madara_node = setup_madara_node(plan)
    
    # 2. Start first Juno node that syncs from Madara and shares blocks via P2P
    juno_feeder_node = setup_juno_feeder_node(plan, madara_node)
    tester_image = setup_tester_service(plan)

    # 3. Verify first Juno node is syncing from Madara
    run_sync_test(plan, tester_image, juno_feeder_node.ip_address, 
                  "Waiting for juno feeder p2p to sync...")

    # 4. Start second Juno node that syncs via P2P and verify its synchronization
    peer_node = setup_juno_peer_node(plan, juno_feeder_node)
    run_sync_test(plan, tester_image, peer_node.ip_address, 
                  "Starting the peer sync test...")

    plan.print("Devnet network test completed successfully")
