sync_test = import_module("tests/sync_test.star")

def run_juno_to_juno_sync(plan):
    sync_test.run_juno_to_juno_sync(plan)

def run_juno_from_pathfinder_sync(plan):
    sync_test.run_juno_from_pathfinder_sync(plan)

def run_pathfinder_from_pathfinder_sync(plan):
    sync_test.run_pathfinder_from_pathfinder_sync(plan)