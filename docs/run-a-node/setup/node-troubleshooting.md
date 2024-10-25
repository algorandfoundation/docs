title: Node Troubleshooting

# Do you want to run a node yourself?

If you are a developer, running a private network using [AlgoKit](/docs/get-started/algokit) is simpler and provide more flexibility.

Running a production node for MainNet is very beneficial for decentralization.
However, as any unmanaged system (and any blockchain node/indexer), to achieve high SLA, running a production node has many requirements: appropriate redundancy (some upgrades create downtime on nodes), 24/7 monitoring, regular maintenance, use of a staging environment for testing updates, ...

!!! tip
    Subscribe to notifications on [Forum](https://forum.algorand.org/c/announcements/5) to be aware of new releases. Read carefully all release notes as some updates may create downtimes or may require some configuration changes. It is recommended to systematically update to new releases after testing in a non-production environment.

You can consider using [third-party API services](https://developer.algorand.org/ecosystem-projects/?tags=api-services) or a third-party provider to help setting/maintaining a node.

# First steps

Be sure to have your `$PATH` and `$ALGORAND_DATA` environment variables set properly and that your node is running. In particular:

```bash
goal node status
```

should return something like:

```
Last committed block: 23119736
Time since last block: 0.1s
Sync Time: 2.7s
Last consensus protocol: https://github.com/algorandfoundation/specs/tree/d5ac876d7ede07367dbaa26e149aa42589aac1f7
Next consensus protocol: https://github.com/algorandfoundation/specs/tree/d5ac876d7ede07367dbaa26e149aa42589aac1f7
Round for next consensus protocol: 23119737
Next consensus protocol supported: true
Last Catchpoint:
Genesis ID: testnet-v1.0
Genesis hash: SGO1GKSzyE7IEPItTxCByw9x8FmnrCDexi9/cOUJOiI=
```

Read [Install a Node](../install) to set these variables properly.

If you see:

* `Data directory not specified.  Please use -d or set $ALGORAND_DATA in your environment. Exiting.`: `$ALGORAND_DATA` is not properly set up.
* `command not found: goal`: `$PATH` is not properly set up.
* `Cannot contact Algorand node: open ...: no such file or directory`: the node is not started. Starting a node is done differently depending on the installation method. See sections "Start the node" from [Install a Node](/docs/run-a-node/setup/install).
    
# Common Issues for algod

## Most common issues: wrong version, wrong network, not caught up

One of the most common issues are that the node is on the wrong network, has the wrong algod version, or is not fully synced.

* **Check that the node is synced/caught up** following [Catchup and Status](/docs/run-a-node/operations/catchup). See below if the node is not syncing.
* **Check that the node is on the right network**: when running `goal node status`, `Genesis ID` must be `mainnet-v1.0` for MainNet, `testnet-v1.0` for TestNet, `betanet-v1.0` for BetaNet. See [Switch Networks](/docs/run-a-node/operations/switch_networks) to solve this issue.
* **Check that the version** reported of `algod -v` and `goal version -v` are the latest stable release (if using MainNet or TestNet) or the latest beta release (if using BetaNet). See the [official repo](https://github.com/algorand/go-algorand/releases) for the list of all releases. Beta releases are clearly marked as such.

## My node is not syncing/catching up at all (Last Committed Block is 0)

As soon as the node is started, the `Last Committed Block` from `goal node status` should be increasing.
If it stays 0, the node is not syncing/catching up at all.
This is usually due to a connectivity issue or a DNS restriction imposed by the ISP.

### No connectivity

First, check that the node has proper internet access.
To check it with the command line, you may use `curl https://example.com`.

### DNS restrictions

By default, a node gets the list of relays it needs to connect to by reading DNS SRV records.
To ensure those records are not tampered with, the node uses DNSSec.

The node will first try to use the system DNS. If it fails, it will use the fallback DNS from `config.json` (if provided). If this fails again, it will try hardcoded DNS from [tools/network/dnssec/config.go](https://github.com/algorand/go-algorand/blob/master/tools/network/dnssec/config.go).

Some ISP, enterprise networks, or public networks only allows DNS queries to their own DNS server that may not support DNSSec.
In that case, you will need to set `"DNSSecurityFlags": 0` inside `$ALGORAND_DATA/config.json`.

!!! warning
    Setting `DNSSecurityFlags` to `0` lowers the security of the node and may allow attackers to make your node connect to untrustworthy relays. While untrustworthy relays cannot make your node accept invalid blocks or create invalid transactions, they may censor all transactions from your nodes or may prevent your node from syncing by not providing the latest blocks. It is strongly recommended to enable DNSSec in production.

!!! tip
    Do not forget to restart `algod` after any change to the configuration.
    
To check if your node's DNS access is correct, run the following commands in the terminal:

```
dig -t SRV _algobootstrap._tcp.mainnet.algorand.network +dnssec
dig -t SRV _algobootstrap._tcp.mainnet.algorand.network @8.8.8.8 +dnssec
```

at least one of them must return a list of relays without any error nor warning.
(The first command uses the system DNS while the second one uses Google DNS.)
    
### Other issues

Here are other less common reasons for a node not being able to catch up at all.

* Check the presence of the right `genesis.json` file in `$ALGORAND_DATA`. See the documentation on how to [switch networks](/docs/run-a-node/operations/switch_networks/).
* If using BetaNet, check that `$ALGORAND_DATA/config.json` contains the right `DNSBootstrapID`. See the [configuration for BetaNet](/docs/run-a-node/operations/switch_networks/#dns-configuration-for-betanet).

## My node is syncing/catching up very slowly (without fast catchup)

Syncing without fast catchup is expecting to take a 2 to 4 weeks in November 2022, due to the amount of data on the Algorand blockchain.
In addition, syncing will slow down as rounds increases, as newer blocks usually contain many more transactions than older blocks.

!!! tip
    Remember that non-archival node can be synced faster using [fast catchup](/docs/run-a-node/setup/install#sync-node-network-using-fast-catchup). Archival nodes cannot use fast catchup but algonode.io, a third-party service, provides [snapshots of archival nodes](https://algonode.io/extras/). This tip is NOT an endorsement of these snapshots and using these snapshots require careful weighting of associated risks: `algod` cannot verify that these snapshots are valid. In particular, `algod` cannot verify that these snapshots do not contain invalid data that would allow for double spending.

If it looks like your node will take much longer than 4 weeks for syncing, then:

1. Check that your node has the right [hardware requirements](/docs/run-a-node/setup/install/#hardware-requirements). In particular, if your node has less than 4GB (or ideally 8GB) of RAM and/or an HDD/slow-SATA-SSD/SD, then you will not be able to catch up on MainNet.
2. Check that your node is not overused:
    1. Check RAM and CPU use using `top` or `htop`
    2. Check free disk using `df -h`
3. Check that your Internet connection is fast enough using a speed test. You need to have at least a 100Mbps connection and low enough latency (below 100ms is most likely needed).  
4. While there are relays all over the world, some regions may have a few number of relays which may slow down catching up. Check latency to the best relays using [algonode.io](https://algonode.io/extras/) scripts (in the folder `utils` of https://snap.algonode.cloud/ --- copied below for completeness):
    ```
    #!/bin/bash

    # needs dig from dnsutils
    N=$(dig +short srv _algobootstrap._tcp.mainnet.algorand.network @1.1.1.1 |wc -l)
    echo "Querying $N nodes, be patient..."
    echo "" > report.txt
    for relay in $(dig +short srv _algobootstrap._tcp.mainnet.algorand.network @1.1.1.1|awk '{print $4 ":" $3}');
    do
      echo -n .
      curl -s -o /dev/null --max-time 1 "http://$relay/v1/urtho/ledger/0"
      echo -ne '\bo'
      curl -s -o /dev/null --max-time 1 "http://$relay/v1/urtho/ledger/0" -w %{time_total} >> report.txt
      echo -ne '\b+'
      echo "s;$relay" >> report.txt
    done

    echo "Top 20 nodes"
    sort -n report.txt | head -20
    ```
    
    Latency above 100ms to the top 20 relays may most likely cause issues.
5. Check that `$ALGORAND_DATA/config.json` is either absent or only contain the non-default parameters you actually need to change. Only change parameters if you understand the consequences, some parameter changes may significantly slow down syncing. 


## My node is not syncing/catching up with fast catchup

See [troubleshooting for fast catchup](/docs/run-a-node/setup/install#troubleshooting-for-fast-catchup).

## Other issues

### I get an `overspend` error when sending a transaction

If sending a transaction results in an `overspend` error (i.e., the word `overspend` appears in the error):
1. Check that the account has enough Algos on a [block explorer](https://developer.algorand.org/ecosystem-projects/?tags=block-explorers).
2. Check that your node is [synced and on the right network](#Most-common-issues-wrong-version-wrong-network-not-caught-up).
3. Do not forget to take into account:
    1. The [minimum balance requirement](/docs/get-details/parameter_tables/#minimum-balance) of 0.1 Algo for a basic account (more if ASA or applications are created or used).
    2. The [fee](/docs/get-details/transactions/#fees) paid by the sender of the transaction.


## None of the above works

If none of the above works, you can look for additional information on the [third-party algonode.io FAQ](https://algonode.io/faq/#algorand-faq).

If this still does not solve your issue, please open a new post on [Forum](https://forum.algorand.org) with the following information:

* what you are trying to do and what fails (with full command lines tried and outputs written in triple backquotes)
* OS version
* machine specs: number of CPU, size of RAM, disk type (NVMe SSD, SATA SSD, ...)
* actual use: memory available, disk available, ...
* algod version: `algod -v`
* goal version: `goal version -v`
* goal status: `goal node status`
* content of `config.json` (in the data folder `$ALGORAND_DATA`)
* link to the files `algod-out.log`, `algod-err.log`, `node.log` (in the data folder `$ALGORAND_DATA`) uploaded in a Github gist or equivalent system


!!! tip
    Always write code/long outputs between triple backquotes or use the "code" button to facilitate the user reading
    
        ```
        $ goal node status

        Last committed block: 23119736
        Time since last block: 0.1s
        Sync Time: 2.7s
        Last consensus protocol: https://github.com/algorandfoundation/specs/tree/d5ac876d7ede07367dbaa26e149aa42589aac1f7
        Next consensus protocol: https://github.com/algorandfoundation/specs/tree/d5ac876d7ede07367dbaa26e149aa42589aac1f7
        Round for next consensus protocol: 23119737
        Next consensus protocol supported: true
        Last Catchpoint:
        Genesis ID: testnet-v1.0
        Genesis hash: SGO1GKSzyE7IEPItTxCByw9x8FmnrCDexi9/cOUJOiI=
        ```