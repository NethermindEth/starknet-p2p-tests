# Import individual test modules
juno_from_juno = import_module("tests/sync/juno_from_juno.star")
juno_from_pathfinder = import_module("tests/sync/juno_from_pathfinder.star")
pathfinder_from_pathfinder = import_module("tests/sync/pathfinder_from_pathfinder.star")
pathfinder_from_juno = import_module("tests/sync/pathfinder_from_juno.star")

def run_juno_from_juno_sync(plan):
    juno_from_juno.run(plan)

def run_juno_from_pathfinder_sync(plan):
    juno_from_pathfinder.run(plan)

def run_pathfinder_from_pathfinder_sync(plan):
    pathfinder_from_pathfinder.run(plan)

def run_pathfinder_from_juno_sync(plan):
    pathfinder_from_juno.run(plan)
