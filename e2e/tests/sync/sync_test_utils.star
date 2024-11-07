def execute_sync_test(plan, peer_node, rpc_port, sync_timeout_seconds, target_block_number):
    """
    Executes the sync test against a peer node:
    1. Creates a tester service
    2. Runs the sync test
    3. Returns the test results
    """
    # Create and run tester service
    tester_service = plan.add_service(
        "sync-tester",
        config=ServiceConfig(
            image=ImageBuildSpec(
                image_name="sync-test",
                build_context_dir="./../../tester",
            ),
        )
    )

    plan.print("Starting the sync tester...")
    # Use the correct RPC port in the URL
    plan.exec(tester_service.name, ExecRecipe(
        ["node", "index.mjs", "http://" + peer_node.ip_address + ":" + str(rpc_port), str(sync_timeout_seconds), str(target_block_number)]
    ))

def cleanup_services(plan):
    """
    Cleans up all services after the test is complete
    """
    services = plan.get_services(description="Fetching running services")
    for service in services:
        plan.print("Found service: " + service.name + " at " + service.ip_address)
        plan.remove_service(
            name=service.name, 
            description="Removing service " + service.name
        )

def run_sync_test(plan, feeder_node, peer_node, rpc_port, sync_timeout_seconds, target_block_number):
    """
    Orchestrates the complete sync test process:
    1. Executes the sync test
    2. Cleans up all services
    """
    execute_sync_test(plan, peer_node, rpc_port, sync_timeout_seconds, target_block_number)
    cleanup_services(plan) 