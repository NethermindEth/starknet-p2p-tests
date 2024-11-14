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

def execute_sync_test(plan, source_node, target_node, source_rpc_port, target_rpc_port, timeout_seconds, target_block_number):
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
    # Pass all required parameters to index.mjs
    plan.exec(
        service_name=tester_service.name,
        recipe=ExecRecipe(
            command=[
                "node",
                "index.mjs",
                "http://" + source_node.ip_address + ":" + str(source_rpc_port),
                "http://" + target_node.ip_address + ":" + str(target_rpc_port),
                str(timeout_seconds),
                str(target_block_number),
                source_node.name,
                target_node.name,
            ]
        )
    )

def run_sync_test(plan, source_node, target_node, source_rpc_port, target_rpc_port, timeout_seconds, target_block_number):
    """
    Orchestrates the complete sync test process:
    1. Executes the sync test
    2. Cleans up all services
    """
    execute_sync_test(
        plan,
        source_node,
        target_node,
        source_rpc_port,
        target_rpc_port,
        timeout_seconds,
        target_block_number
    )
    cleanup_services(plan) 