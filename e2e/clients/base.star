def run(plan, name, default_image, default_cmd, ports, participant, files={}):
    return plan.add_service(
        name=name,
        config=ServiceConfig(
            image=participant.get("image", default_image),
            cmd=default_cmd + participant.get("extra_args", []),
            ports=ports,
            files=files
        ),
    )