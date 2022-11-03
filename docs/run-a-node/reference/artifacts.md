title: Node artifacts

These files run as part of the Algorand node or are CLI utilities that will help to diagnose or interact with a currently running node. The primary files are described below.

# goal
Goal is the command line utility used to interact with the Algorand node. It communicates with algod and the kmd process to do things like: create an account, list the ledger, examine the status of the network, or create a transaction. Goal documentation is available in the [Goal](../../../clis/goal/goal) guide.

# algod
Algod is the main Algorand process for handling the blockchain. Messages between nodes are processed, the protocol steps are executed, and the blocks are written to disk. The algod process also exposes a REST API server that developers can use to communicate with the node and the network. Algod uses the data directory for storage and configuration information.

# algoh
Algoh is an optional hosting process for algod, whose use is encouraged to help catch fatal runtime errors and proactively report with logs. Algoh also monitors for stalls and proactively reports them with diagnostics data in case the problem is local to the instance.  We’ll be adding optional layer-2 telemetry to some nodes running algoh, but it will be disabled for most nodes (and will be configurable).

# kmd
Kmd is the key management daemon. This process handles interacting with clients’ private keys for Algorand accounts. The process is responsible for generating and importing spending keys, signing transactions, and interacting with key storage mechanisms like hardware wallets. This process can also be executed on a separate machine, isolating the spending keys from the network. This process also uses a data directory for wallet configurations. In the default configuration, this will be in the data directory for algod but will contain its own folder labeled `kmd-version`. The kmd process also hosts a REST endpoint for integration.

# algokey
algokey is a command line utility for generating, exporting and importing keys. The tool can be used to sign single and multi-signature transactions as well. See [algokey](../../../clis/algokey/algokey) documentation for more details.

# carpenter
Carpenter is a debug tool that helps to visualize the protocol and how it is proceeding. The tool reads the node log file and formats the output. Generally, every entry displayed in the tool starts with the round number, with the period and step. Each line displayed in the tool will have a message that is relative to the step and will display when a proposal or a vote is accepted or rejected. When accepted votes are displayed, the number of votes is also included in parentheses. When the threshold for the vote is reached a message will be displayed letting the user know the next step in the protocol is about to start. Messages are color-coded by round and user.

When running carpenter, just specify the latest.log file from the data directory.

```
./carpenter -file data/latest.log
```

Or use the standard data directory specifier syntax:

```
./carpenter -d data
```

If you have the `ALGORAND_DATA` environment variable set, you can just run:

```
./carpenter -D
```

# update.sh, updater, and updatekey.json
These are the primary files used for installing and updating the node software. They can be executed manually or put them in a CRON job and have it execute more regularly. The process for doing this is explained in the installation guide.

# Node Data files
As part of the installation process, a data directory is created. The node stores its local copy of the blockchain in this directory. Log files and configuration files are also stored here. Each of the configuration files is described below.

# config.json.example
This file is just an example configuration file that lets you configure a node with specific parameters. If this file is renamed to config.json, the settings in this file will take effect on the node when it is started. Each of the settings is defined in the Node configuration settings guide.

# phonebook.json
As discussed in [Algorand Node Types](../../setup/types) there are two types of nodes; relay and non-relay. Any node can function as either but it is not good practice for nodes that are relays to run the kmd process or manage accounts. Relay nodes have higher requirements and open more ports than standard non-relay nodes. Non-relay nodes only connect to relay nodes. They never connect to other non-relay nodes. The relay nodes are published in Algorand SRV records. You can create a relay node that is not in a SRV record. You can specify one or more of these non-published relays to be placed in a pool of available relay nodes for a node starting up. The node will randomly pick from this pool of nodes to open connections to communicate with the network. Before specifying how to change the relay pool, note that setting some options will replace the default SRV records from the pool. This means that if you intend to connect to the Algorand MainNet or TestNet network, the relays in your pool must connect to a published Algorand SRV relay.

One method for replacing the pool entirely is to use the goal node start -p command. The -p parameter expects a list of relays that will be added to a pool. Note that this does away with the default pool. Here is an example of using the -p option to connect to TestNet.

```
./goal node start -d data -p "r1.algorand.network:4161;r2.algorand.network:4161"
```

In this example, the two specified relays are the only ones used.

Another option is to create a file named phonebook.json in your binary directory and add a list of relays similar to the following:

```
{ "Include": [
    "r1.algorand.network:4161",
    "r2.algorand.network:4161" ] }
```

In this case, the entries are added to the default pool. If you want these entries to override the default pool set the `DNSBootstrapID` in config.json to "". Setting a configuration value for a node is described in the Node Configuration Settings guide.

# node.archive.log and node.log
The node.log file is the log file of the current node. This file contains a set of JSON entries for various steps processed by the node. The carpenter utility can be used to view these entries in a much more digestible format. Once the node log file has reached its maximum size, it is copied to node.archive.log and a new version of node.log is created. 

# algod-err.log, algod-out.log
When started through goal, algod redirects the error and output streams to algod-err.log and algod-out.log respectively.

algod.net, algod.pid and algod.token, algod-listen.net
The algod.net, algo.pid and algod.token files are created in the node’s data folder when the node starts. The algod.net file contains the IP and port the node is serving REST API calls on. The algod.pid file contains the process id of the algod daemon of the running node. The algod.token file contains the API token that must be used to communicate with the node’s REST APIs. This is passed to the REST API using the `X-Algo-API-Token` header.  algod-listen.net contains the IP and port that the node is listening on for incoming connections if any.

# host.log
Contains logging output from algoh.

# wallet-genesis.id and genesis.json
These files are associated with the genesis block and the associated wallet ids at creation of the network. The wallet-genesis.id file contains the unique id for the genesis block that was most recently installed. This is also the name of the directory in which the blockchain’s SQLite file is stored.

The genesis.json file specifies the initial state of the blockchain - its ‘genesis block’. This is a JSON formatted file with the schema for the blockchain. It contains the network name and id, the protocol version and list of allocated addresses to start the chain with. Each address contains a list of things like their status and the amount of Algos they own.

# Nested Directories
The data directory will contain a couple of sub-directories depending on what has been done with that instance.  One directory, named after the specific Algorand network and genesis version, will contain the SQLite files associated with the blockchain ledger and any wallets. There may be additional directories if the node has been updated. The backup directory contains a backup of the data directory. This allows the node to revert to the previous version in case of a failure during an update. A kmd directory may also exist with a SQLite instance for that process. It will also contain REST endpoint and API token files, similar to the algod files.
