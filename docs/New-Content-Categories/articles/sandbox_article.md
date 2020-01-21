# Introducing Sandbox: The quick way to get started on Algorand

Running a node on Algorand is a [very fast and straightforward process](../getting-started/setup.md).

Sandbox makes the process of node creation and configuration seamless with several helpful commands for containerization and process management. It is simply a quick and easy way to create and configure an Algorand node using Docker.

<center><iframe width="560" height="315" src="https://www.youtube.com/embed/IqPAHa-YjW0" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe></center>


# Why use Sandbox?
Sandbox simplifies the process of node configuration and maintenance allowing you to quickly spin up a containerized node on any of the three Algorand networks as well as managing it through the sandbox environment.

!!! warning 
    Sandbox is directed at learning and not meant for production.

# How to Use Sandbox
First, make sure you have Docker installed. Clone the repository, which can be found here: https://github.com/algorand/sandbox and if you are using a Mac, like I am, make sure you have `wget` installed. If you use homebrew as your package manager, simply run:
``` shell
$ brew install wget

```

To get started `run` the up command
``` shell
$ ./sandbox up
```
This will spin up a docker container defaulting to the testnet binaries.

Use docker commands to list the container

``` shell
$ docker container ls
```
This should return the one container that was initialized from the previous sandbox command. Sandbox manages one active container at a time. Note that you will see multiple containers if you have other containers initialized from other docker processes.

You can configure sandbox to run a node on any of the three networks by passing in the network name:
``` shell
$ ./sandbox up [mainnet||testnet||betanet]
```
Note that when you run `./sandbox up [NETWORK_NAME]` it will initialize with a snapshot and your node will sync up to the latest round. To skip the snapshot initialization and begin downloading the blockchain from the first round, simply pass in the `-s` flag.

``` shell
$ ./sandbox up [mainnet||testnet||betanet] [-s]
```
If you’ve already configured sandbox for a specific network (you spun up a docker container using sandbox and are running a node) you can tear down the environment and start it back up at any time.

```shell
$ ./sandbox down

```
Running `docker container ls` should not return a container ID. Spinning it back up is done using `./sandbox up` . In this case, the node will pick it up from the round it was in when the the sandbox environment was taken down.

You can also run `./sandbox clean` to stop and delete the container and data directory. This is different from `./sandbox down` as the container in this case is removed.

The `./sandbox test` command is helpful because it sends REST calls that hit algod and kmd <LINK> as well as running some goal <LINK> commands.

```shell
$ ./sandbox test
Test command forwarding...
$ docker exec -it sandbox uname -a
Linux 13ad4f9fd7b8 4.9.184-linuxkit #1 SMP Tue Jul 2 22:58:16 UTC 2019 x86_64 x86_64 x86_64 GNU/Linux
```

The result for `goal node status`

```shell
Test algod...
$ docker exec -it sandbox /opt/algorand/node/goal node status -d /opt/algorand/node/data
Last committed block: 4146633
Time since last block: 0.6s
Sync Time: 0.0s
Last consensus protocol: https://github.com/algorandfoundation/specs/tree/4a9db6a25595c6fd097cf9cc137cc83027787eaa
Next consensus protocol: https://github.com/algorandfoundation/specs/tree/4a9db6a25595c6fd097cf9cc137cc83027787eaa
Round for next consensus protocol: 4146634
Next consensus protocol supported: true
Has Synced Since Startup: false
Genesis ID: testnet-v1.0
Genesis hash: SGO1GKSzyE7IEPItTxCByw9x8FmnrCDexi9/cOUJOiI=
```

The result of hitting the algod API:

```shell
Test Algod REST API...
$ curl localhost:4001/v1/status -H "X-Algo-API-Token: $(cat data/algod.token)"
{"lastRound":4146365,"lastConsensusVersion":"https://github.com/algorandfoundation/specs/tree/4a9db6a25595c6fd097cf9cc137cc83027787eaa","nextConsensusVersion":"https://github.com/algorandfoundation/specs/tree/4a9db6a25595c6fd097cf9cc137cc83027787eaa","nextConsensusVersionRound":4146366,"nextConsensusVersionSupported":true,"timeSinceLastRound":1243605287,"catchupTime":0,"hasSyncedSinceStartup":false}
```

The `logs` command in sandbox is the same as `./carpenter -d data` which is Algorand’s debugging tool that reads the node log file and formats the output.

```shell
$ ./sandbox logs
```

This returns an output that looks like this
```shell
4146536.0.2:                BlockPipelined NXY5Z-4146537.0.0|
4146536.0.2:        VoteAccepted(114/2700) RD2VU-4146536.0.1|
4146536.0.2:          VoteAccepted(66/369) RD2VU-4146536.0.2|
4146536.0.2:          VoteAccepted(60/429) RD2VU-4146536.0.2|
4146536.0.2:         VoteAccepted(19/2719) RD2VU-4146536.0.1|
4146536.0.2:          VoteAccepted(70/499) RD2VU-4146536.0.2|
4146536.0.2:              ProposalAccepted YD2W5-4146537.0.0|
4146536.0.2:                BlockPipelined YD2W5-4146537.0.0|
4146536.0.2:          VoteAccepted(59/558) RD2VU-4146536.0.2|
4146536.0.2:          VoteAccepted(39/597) RD2VU-4146536.0.2|
4146536.0.2:          VoteAccepted(14/611) RD2VU-4146536.0.2|
4146536.0.2:          VoteAccepted(51/662) RD2VU-4146536.0.2|
4146536.0.2:          VoteAccepted(56/718) RD2VU-4146536.0.2|
4146536.0.2:          VoteAccepted(10/728) RD2VU-4146536.0.2|
4146536.0.2:        VoteAccepted(141/2860) RD2VU-4146536.0.1|
4146536.0.2:        VoteAccepted(116/2976) RD2VU-4146536.0.1|
4146536.0.2:          VoteAccepted(56/784) RD2VU-4146536.0.2|
4146536.0.2:          VoteAccepted(63/847) RD2VU-4146536.0.2|
4146536.0.2:          VoteAccepted(59/906) RD2VU-4146536.0.2|
4146536.0.2:          VoteAccepted(66/972) RD2VU-4146536.0.2|
4146536.0.2:         VoteAccepted(79/1051) RD2VU-4146536.0.2|
4146536.0.2:         VoteAccepted(27/1078) RD2VU-4146536.0.2|
4146536.0.2:         VoteAccepted(67/1145) RD2VU-4146536.0.2|
4146536.0.2:   ThresholdReached(1145/1112) RD2VU-4146536.0.2|
4146536.0.0:                RoundConcluded RD2VU-       |
4146536.0.0:                    RoundStart RD2VU-       |
4146537.0.0:                BlockAssembled YD2W5-4146537.0.0|
```

Using this tool, you can watch the agreement service write blocks to the ledger in realtime.

Sandbox is highly flexible and you can interact with goal in a few different ways.

```shell
$ ./sandbox status
```
This will return the same output as `./sandbox goal node status`

If you are familiar with running a node on Algorand, you would normally use the Algorand cli "goal" to manage your node. You can do the same thing with sandbox by running `./sandbox goal (args)` or by running `./sandbox enter` which basically puts a shell in sandbox allowing you to interact with the node and run commands from within the container.

```shell
$ ./sandbox enter
Entering /bin/bash session in the sandbox container...
root@5296a82a0746:/opt/algorand/node$ ls
COPYING  algokey       catchupsrv   find-nodes.sh  msgpacktool    update.sh
algocfg  algorand@.service.template  data   genesisfiles  node_exporter    updater
algod  backup        ddconfig.sh  goal   sudoers.template
algoh  carpenter       diagcfg   kmd   systemd-setup.sh
root@5296a82a0746:/opt/algorand/node$
```

The final, more experimental feature in Sandbox is the tutorial. Running `./sandbox introduction` gives you a great step by step walkthrough of working with an Algorand node which includes creating a wallet, creating accounts, funding accounts and broadcasting transactions to the network.

# Conclusion
Please check out sandbox and submit features requests to improve this great tool!

https://github.com/algorand/sandbox