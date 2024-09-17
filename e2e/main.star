sync_test = import_module("tests/sync_test.star")

def run(plan):
    plan.print("Starting synchronization tests")
    sync_test.run_all_sync_tests(plan)