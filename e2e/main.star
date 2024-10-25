juno_from_juno = import_module("./tests/sync/juno_from_juno.star")
juno_from_pathfinder = import_module("./tests/sync/juno_from_pathfinder.star")
pathfinder_from_juno = import_module("./tests/sync/pathfinder_from_juno.star")
pathfinder_from_pathfinder = import_module("./tests/sync/pathfinder_from_pathfinder.star")
devnet_network = import_module("./tests/hive/devnet_network.star")

def run_juno_from_juno_sync(plan):
    return juno_from_juno.run(plan)

def run_juno_from_pathfinder_sync(plan):
    return juno_from_pathfinder.run(plan)

def run_pathfinder_from_juno_sync(plan):
    return pathfinder_from_juno.run(plan)

def run_pathfinder_from_pathfinder_sync(plan):
    return pathfinder_from_pathfinder.run(plan)

def run_devnet_network(plan):
    return devnet_network.run(plan)
