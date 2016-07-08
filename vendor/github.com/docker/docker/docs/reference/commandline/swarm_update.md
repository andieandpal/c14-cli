<!--[metadata]>
+++
title = "swarm update"
description = "The swarm update command description and usage"
keywords = ["swarm, update"]
advisory = "rc"
[menu.main]
parent = "smn_cli"
+++
<![end-metadata]-->

# swarm update

```markdown
Usage:  docker swarm update [OPTIONS]

Update the Swarm

Options:
      --auto-accept value               Auto acceptance policy (worker, manager or none)
      --cert-expiry duration            Validity period for node certificates (default 2160h0m0s)
      --dispatcher-heartbeat duration   Dispatcher heartbeat period (default 5s)
      --external-ca value               Specifications of one or more certificate signing endpoints
      --help                            Print usage
      --secret string                   Set secret value needed to accept nodes into cluster
      --task-history-limit int          Task history retention limit (default 10)
```

Updates a Swarm cluster with new parameter values. This command must target a manager node.


```bash
$ docker swarm update --auto-accept manager
```

## Related information

* [swarm init](swarm_init.md)
* [swarm join](swarm_join.md)
* [swarm leave](swarm_leave.md)