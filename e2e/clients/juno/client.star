base = import_module("../common/base.star")

def run(plan, name, participant):
    image = participant.get("image", "nethermind/juno:latest")
    is_feeder = participant.get("is_feeder", False)
    private_key = participant.get("private_key", "")
    http_port = participant.get("http_port", 6060)
    p2p_port = participant.get("p2p_port", 7777)
    peer_multiaddrs = participant.get("peer_multiaddrs", [])

    cmd = [
        "--network", "sepolia",
        "--db-path", "/var/lib/juno",
        "--log-level", "info",
        "--http",
        "--http-port", str(http_port),
        "--http-host", "0.0.0.0",
        "--p2p",
        "--p2p-addr", "/ip4/0.0.0.0/tcp/" + str(p2p_port),
        "--p2p-private-key", private_key
    ]
    if is_feeder:
        cmd.append("--p2p-feeder-node")
    for peer_multiaddr in peer_multiaddrs:
        cmd.extend(["--p2p-peers", peer_multiaddr])
    
    ports = {
        "rpc": PortSpec(
            number=http_port,
            transport_protocol="TCP",
            application_protocol="http",
        ),
        "p2p": PortSpec(
            number=p2p_port,
            transport_protocol="TCP",
        )
    }

    return base.run(plan, name, image, cmd, ports, participant)
