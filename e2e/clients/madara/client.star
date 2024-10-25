base = import_module("../common/base.star")

def run(plan, name, participant):
    image = participant.get("image", "ghcr.io/madara-alliance/madara:latest")
    rpc_port = participant.get("rpc_port", 9944)
    gateway_port = participant.get("gateway_port", 8080)
    
    cmd = [
        "--rpc-port", str(rpc_port),
        "--rpc-external",
        "--rpc-cors=*"
    ]

    ports = {
        "rpc": PortSpec(
            number=rpc_port,
            transport_protocol="TCP",
            application_protocol="http",
        ),
        "gateway": PortSpec(
            number=gateway_port,
            transport_protocol="TCP",
            application_protocol="http",
        )
    }

    return base.run(plan, name, image, cmd, ports, participant)
