title: Algorand node types

The Algorand network is comprised of two distinct types of nodes, **relay nodes**, and **non-relay nodes**. Relay nodes are primarily used for communication routing to a set of connected non-relay nodes. Relay nodes communicate with other relay nodes and route blocks to all connected non-relay nodes. Non-relay nodes only connect to relay nodes and can also participate in consensus. Non-relay nodes may connect to several relay nodes but never connect to another non-relay node.

In addition to the two node types, nodes can be configured to be [**archival**](#archival-mode). Archival nodes store the entire ledger, as opposed to the last 1000 blocks for non-archival blocks. Relay nodes are necessarily archival. Non-relay archival nodes are often used to feed an [**indexer**](../indexer/) that allows more advanced queries on the history of the blockchain.

Finally, a node may either participate in consensus or not. Participation nodes do not need to be archival. In addition, to reduce attack surfaces and outage risks, it is strongly recommended that participation nodes are only used for the purpose in participating in consensus. In particular, participation nodes should not be relays.

All node types use the same install procedure. To setup a node for a specific type, requires a few configuration parameter changes or operations as described below. The default install will set the node up as a non-relay non-archival non-participation mode.

The table below is a summary of possible configurations:

| Relay? | Archival? | Participation? | |
|-|-|-|-|
| relay | archival | participation | ❌ insecure / strongly discouraged |
| relay | archival | non-participation | ✅ normal configuration for a **relay** |
| relay | non-archival | * | ❌ impossible |
| non-relay | archival | participation | ❓ discouraged, as participation nodes do not need to be archival and should only be used for participation, and not used for any other purpose |
| non-relay | archival | non-participation | ✅ normal configuration for an **archival node**, often connected to an [**indexer**](../indexer/) |
| non-relay | non-archival | participation | ✅ normal configuration for a **participation node** |
| non-relay | non-archival | non-particiption | ✅ normal configuration for a node used to submit transactions to the blockchain and access current state (current balances, smart contract state, ...) but no historical state |


# Which mode do I need?

Here are some common use cases:

* I want to participate in consensus and help secure the Algorand network.
    * ➥ non-relay non-archival participation node.
    * Note: I need to have some Algos for that purpose and I need to monitor my node 24/7 to ensure it is working properly.
* I want to send transactions and read current state of smart contracts/applications:
    * ➥ non-relay non-archival non-participation node.
    * Example: a dApp website that does not use any historical information (past transaction/operation), a website displaying balances of a list of important accounts.
* I want full access to historical data (blocks, transactions) with advanced querying:
    * ➥ non-relay archival non-participation node, together with an [**indexer**](../indexer/).
* I want to get state proofs for any block:
    * ➥ non-relay archival non-participation node.

# Participation Node 
How to install a node is described [here](../install/).  Classifying a node as a participation node is not a configuration parameter but a dynamic operation where the node is hosting participation keys for one or more online accounts. This process is described in [Participate in Consensus](../participate/index.md). Technically both non-relay and relay nodes can participate in consensus, but Algorand recommends *only* non-relay nodes participate in consensus. 

!!! info
    Non-relay nodes do not have to participate in consensus. They still have access to the ledger and can be used with applications that need to connect to the network to submit transactions and read block data. 


# Archival Mode

By default non-relay nodes only store a limited number of blocks (approximately up to the last 1000 blocks) locally. Older blocks are dropped from the local copy of the ledger. This reduces the disk space requirement of the node. These nodes can still participate in consensus and applications can connect to these nodes for transaction submission and reading block data. The primary drawback for this type of operation is that older block data will not be available.

The archival property must be set to true to run in archival mode, which will then set the node to store the entire ledger.

!!! warning
     Setting a node to run in archival mode on MainNet/TestNet/BetaNet will significantly increase the disk space requirements for the node. For example, in September 2023, a MainNet non-archival node uses around 20GB of storage, while an archival node requires approximately 2TB of storage.


!!! info
    Relay nodes are always set to Archival mode. Non-relay nodes have the option to run in either configuration.

# Relay Node

A relay node uses the same software install as a non-relay node and only requires setting a few additional configuration parameters.

A node is a valid relay node if two things are true:

1. The node is configured to accept incoming connections on a publicly-accessible port (4160 by convention on MainNet, and 4161 on TestNet, except if behind a proxy in which case port 80 is used.).
2. The node's public IP address (or a valid DNS name) and assigned port are registered in Algorand's SRV records for a specific network (MainNet/TestNet).
   
Relay nodes are where other nodes connect. Therefore, a relay node must be able to support a large number of connections and handle the processing load associated with all the data flowing to and from these connections. Thus, relay nodes require significantly more power than non-relay nodes. Relay nodes are always configured in archival mode.

Relay nodes also require very high egress bandwidth. In October 2022, MainNet relay nodes egress bandwidth is between 10TB to 30TB per month.

See [Configuring Node as a Relay](../../reference/relay) for more information on setting up a relay.





