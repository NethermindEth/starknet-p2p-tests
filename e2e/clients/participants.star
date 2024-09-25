juno = import_module("./juno.star")
pathfinder = import_module("./pathfinder.star")

def run_participant(plan, name, participant):
    module = import_module("./{}.star".format(participant["type"]))
    return module.run(plan, name, participant)

def run(plan, participants):
    return [run_participant(plan, "{}-{}".format(participant["type"], index), participant) for (index, participant) in enumerate(participants)]
