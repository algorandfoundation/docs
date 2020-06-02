title: Install the Indexer

# Indexer Mode V2
[TO DO]

# Indexer Mode V1
All transaction searching by default is limited to a 1000 round range. If a node is configured in archival mode, an additional configuration option can be used to turn on a node indexer and remove this restriction. With the indexer turned on, searching for specific transactions will be quicker. Two additional REST calls are also made available for more refined searching. 

The two Additional REST calls are:

```
/v1/transaction/{txid}
```
This call allows quickly locating a transaction using the txid
See [REST API Reference](../../reference/rest-apis/algod.md#get-v1transactiontxid) for more details.

```
/v1/account/{account}/transactions?fromDate=YYYY-MM-DD&toDate=YYYY-MM-DD) 
```

This call allows locating all transactions within a date range. Date parameters support RFC3339 (ie 2006-01-02T15:04:05Z07:00).
See [REST API Reference](../../reference/rest-apis/algod.md#get-v1accountaddresstransactions) for more details.

To turn on indexing for a node, the `isIndexerActive` configuration parameter must be set to `true`. The [Node Configuration](../../reference/node/config.md) guide describes setting node configuration properties.

!!! warning
     Turning on indexing with a node will increase the disk space required by the node.

!!! info
    Indexing on a node is only allowed with nodes that have archival mode turned on.