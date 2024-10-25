juno = import_module("./juno/client.star")
pathfinder = import_module("./pathfinder/client.star")
madara = import_module("./madara/client.star")

def run_participant(plan, name, participant):
    if participant["type"] == "juno":
        return juno.run(plan, name, participant)
    elif participant["type"] == "pathfinder":
        return pathfinder.run(plan, name, participant)
    elif participant["type"] == "madara":
        return madara.run(plan, name, participant)
    else:
        fail("Unknown client type: " + participant["type"])

def run(plan, participants):
    return [run_participant(plan, "{}-{}".format(participant["type"], index), participant) for (index, participant) in enumerate(participants)]
