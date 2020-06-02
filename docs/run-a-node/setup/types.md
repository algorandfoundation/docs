title: Algorand Node Types

The Algorand network is comprised of two distinct types of nodes, **relay nodes**, and **non-relay nodes**. Relay nodes are primarily used for communication routing to a set of connected non-relay nodes. Relay nodes communicate with other relay nodes and route blocks to all connected non-relay nodes. Non-relay nodes only connect to relay nodes and can also participate in consensus. Non-relay nodes may connect to several relay nodes but never connect to another non-relay node.

In addition to the two node types, nodes can be configured to be [**archival**](#archival-mode) and [**indexed**](#indexer-mode). Archival nodes store the entire ledger and if the indexer is turned on, the search range via the API REST endpoint is increased. These additional configuration options are described below.

Both node types use the same install procedure. To setup a node for a specific type, requires a few configuration parameter changes as described below. The default install will set the node up as a non-relay node in non-archival and non-indexed mode.

# Participation Node 
Classifying a node as a participation node is not a configuration parameter but a dynamic operation where the node is hosting participation keys for one or more online accounts. This process is described in [Participate in Consensus](../participate/index.md). Technically both non-relay and relay nodes can participate in consensus, but Algorand recommends *only* non-relay nodes participate in consensus. 

!!! info
    Non-relays nodes do not have to participate in consensus. They still have access to the ledger and can be used with applications that need to connect to the network to submit transactions and read block data. 


# Archival Mode
 By default non-relay nodes only store a limited number of blocks (approximately up to the last 1000 blocks) locally. Older blocks are dropped from the local copy of the ledger. This reduces the disk space requirement of the node. These nodes are can still participate in consensus and applications can connect to these nodes for transaction submission and reading block data. The primary drawback for this type of operation is that older block data will not be available. 
 
 The archival property must be set to true to run in archival mode, which will then set the node to store the entire ledger. Visit the [Node Configuration](../../reference/node/config.md) guide for details on configuring your node. 
 
!!! warning
     Setting a node to run in archival mode will increase the disk space requirements for the node. For example, after 36 hours, the TestNet archival ledger was 212.5MB, whereas the non-archival ledger was 9MB.
 

!!! info
    Relay nodes are always set to Archival mode. Non-relay nodes have the option to run in either configuration.

# Relay Node
A relay node uses the same software install as a non-relay node and only requires setting a few additional configuration parameters.

A node is a valid relay node if two things are true:

1. The node is configured to accept incoming connections on a publicly-accessible port (4161 by convention).
2. The node's public IP address (or a valid DNS name) and assigned port are registered in Algorand's SRV records for a specific network (MainNet/TestNet).
   
Relay nodes are where other nodes connect. Therefore, a relay node must be able to support a large number of connections and handle the processing load associated with all the data flowing to and from these connections. Thus, relay nodes require significantly more power than non-relay nodes. Relay nodes are always configured in archival mode.

See [Configuring Node as a Relay](../../reference/node/relay.md) for more information on setting up a relay.





