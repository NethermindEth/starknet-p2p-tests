base = import_module("../common/base.star")

def run(plan, name, participant):
    image = participant.get("image", "eqlabs/pathfinder:latest-p2p")
    is_feeder = participant.get("is_feeder", False)
    p2p_port = participant.get("p2p_port", 20002)
    rpc_port = 9545  # Fixed RPC port, hardcoded in Pathfinder Dockerfile and not configurable
    ethereum_url = participant.get("ethereum_url", "wss://sepolia.infura.io/ws/v3/2ba63046038749aeadc99d0520cdaecb")
    peer_multiaddrs = participant.get("peer_multiaddrs", [])
   
    env_vars = {
        "RUST_LOG": "debug"
    }

    cmd = [
        "--network", "sepolia-testnet",
        "--p2p.listen-on", "/ip4/0.0.0.0/tcp/" + str(p2p_port),
        "--debug.restart-delay", "5",
        "--debug.pretty-log", "true",
        "--rpc.enable", "true",
        "--ethereum.url", ethereum_url
    ]

    files = {}
    if is_feeder:
        identity_artifact = plan.upload_files(src="identity.json", name="pathfinder_identity")
        files["/app/"] = identity_artifact
        cmd.extend([
            "--p2p.identity-config-file", "/app/identity.json",
            "--p2p.proxy", "true"
        ])

    for peer_multiaddr in peer_multiaddrs:
        cmd.extend(["--p2p.bootstrap-addresses", peer_multiaddr])

    ports = {
        "p2p": PortSpec(
            number=p2p_port,
            transport_protocol="TCP",
        ),
        "rpc": PortSpec(
            number=rpc_port,
            transport_protocol="TCP",
            application_protocol="http",
        )
    }

    return base.run(plan, name, image, cmd, ports, participant, files=files, env_vars=env_vars)
