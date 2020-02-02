title: Node Types

The Algorand network is comprised of two distinct types of nodes, relay nodes, and non-relay nodes. Relay nodes are primarily used for communication routing to a set of connected non-relay nodes. Relay nodes communicate with other relay nodes and route blocks to to all connected non-relay nodes. Non-relay nodes only connect to relay nodes and can also participate in consensus. Non-relay nodes may connect to several relay nodes but never connect to another non-relay node.

In addition to the two node types, nodes can be configured to be archival and indexed. Archival nodes store the entire ledger and if the indexer is turn on, the range and speed of searching for specific data can be increased. These addition configuration options are described below.

Both node types use the same insall proceedure. To setup a node for a specify type, requires a few configuration parameter changes as described below. The default install will set the node up as a non-relay node in non-archival and non-indexed mode as described later in this guide. 

# Participation Node 
Technically both non-relay and relay nodes can participate in consensus, but Algorand only recommends non-relay nodes participate in consensus. Classifying a node as a participation node is not a configuration parameter but a dynamic operation where the node is hosting participation keys for one or more online accounts. This process is described in [Participate in Consensus](/docs/network-participation/participate-in-consensus/overview/).

!!! info
    Non-relays nodes do not have to participate in consensus. They still have access to the ledger and can be used with applications that need to connect to the network to submit transactions and read block data. 


# Archival mode
 By default non-relay nodes only store a limited number of blocks (approximately up to the last 1000 blocks) locally. Older blocks are dropped from the local copy of the ledger. This reduces the disk space requirement of the node. These nodes are can still participate in consensus and applications can connect to these nodes for transaction submission and reading block data. The primary drawback for this type of operation is that older block data will not be available. 
 
 The [Node Configuration](/docs/network-participation/run-a-node/config/) guide describes setting node configuration properties. The archival property must be set to true to run in archival mode, which will then set the node to store the entire ledger. 
 
!!! warning
     Setting a node to run in archival mode will increase the disk space requirements for the node. For example, after 36 hours, the TestNet archival ledger was 212.5MB, whereas the non-archival ledger was 9MB.
 

!!! info
    Relays nodes are always set to Archival mode. Non-relay nodes have the option to run in either configuration.

# Indexer Mode
All transaction searching by default is limited to a 1000 round range. If a node is configured in archival mode, an addtional configuration option can be used to turn on a node indexer to remove this restriction. With the indexer turned on, searching for specific transactions will be quicker. Two additional rest calls are also made available for more refined searching. 

The two Additional REST calls are:

```
/v1/transaction/{txid}
```
This call allows quickly locating a transaction using the txid
See [REST API Reference](/docs/reference-docs/rest-apis/algod/#get-v1transactiontxid) for more details.

```
/v1/account/{account}/transactions?fromDate=YYYY-MM-DD&toDate=YYYY-MM-DD) 
```

This call allows locating all transactions within a date range. Date parameters support RFC3339 (ie 2006-01-02T15:04:05Z07:00).
See [REST API Reference](/docs/reference-docs/rest-apis/algod/#get-v1accountaddresstransactions) for more details.

To turn on indexing for a node, the isIndexerActive configuration parameter must be set to true. The [Node Configuration](/docs/network-participation/run-a-node/config/) guide describes setting node configuration properties.

!!! warning
     Turning on indexing with a node will increase the disk space required by the node.

!!! info
    Indexing on a node is only allowed with nodes that have archival mode turned on.


# Relay Node
A relay node uses the same software install as a non-relay node and only requires setting a few additional configuration parameters.

A node is a valid relay node if two things are true:

1. The node is configured to accept incoming connections on a publicly-accessible port (4161 by convention).
2. The node's public IP address (or a valid DNS name) and assigned port are registered in Algorand's SRV records for a specific network (MainNet/TestNet).
   
Relay nodes are where other nodes connect. Therefore, a relay node must be able to support a large number of connections and handle the processing load associated with all the data flowing to and from these connections. Thus, relay nodes require significantly more power than non-relay nodes. Relay nodes are always configured in archival mode.

See [Configuring Node as a Relay]() for more information on setting up a relay.





