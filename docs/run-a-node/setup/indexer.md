title: Install the indexer

The Algorand Indexer is a feature that enables searching the blockchain for transactions, assets, accounts, and blocks with various criteria. Currently, Algorand has a V1 and V2 Indexer. The V1 Indexer is deprecated and can significantly slow down nodes. Users should now use the V2 Indexer. The V2 Indexer runs as an independent process that must connect to a [PostgreSQL](https://www.postgresql.org/) compatible database that contains the ledger data. The PostgreSQL database is populated by the indexer which connects to an Algorand node and processes all the ledger data and loads the database. The node the Indexer connects to must be an archival node to get all the ledger data. Alternatively, the Indexer can just connect to a PostgresSQL database that is populated by another instance of Indexer. This allows reader instances to be set up that provide the [REST APIs](../../../rest-apis/indexer) for searching the database and another Indexer to be responsible for loading the ledger data.


The V2 Indexer is network agnostic, meaning it can point at BetaNet, TestNet, or MainNet. 

The source code for the Indexer is provided on [github](https://github.com/algorand/indexer).

For details on Indexer usage, read the [Searching the Blockchain](../../../get-details/indexer) feature guide and the [REST API Indexer reference](../../../rest-apis/indexer). See [Indexer README](https://github.com/algorand/indexer) for more details on running the Indexer.


# Indexer V2
## Download the Indexer Binaries
To Install the new Indexer follow the instructions below. The Indexer binaries are available on [github](https://github.com/algorand/indexer/releases).

!!! info
    Additional install methods will be available in the near future.

## Extract the binaries to a specific directory
The binary can be placed in any directory you choose. In these instructions, an indexer folder is used which is located in the current user's home directory.

```bash
$ mkdir ~/indexer
$ cd /path/to/download-dir
$ tar -xf <your-os-tarfile> -C ~/indexer
$ cd ~/indexer/<tarfile-name>
```

## Run the Indexer
The Indexer primarily provides two services, loading a PostgreSQL database with ledger data and supplying a REST API to search this ledger data. You can set the Indexer to point at a database that was loaded by another instance of the Indexer. The database does not have to be on the current node. In fact, you can have one Indexer that loads the database and many Indexers that share this data through their REST APIs. How the Indexer operates is determined with parameters that are passed to the Indexer as it is started.

The Indexer has many options which can be seen using the -h option when running the [Indexer binary](../../clis/indexer/indexer.md).

To start the Indexer as a reader (ie not connecting to an Algorand node), supply the `--postgres` or `-P` option when running the indexer. The value should be a valid connection string for a postgres database.

```bash
$ ./algorand-indexer daemon --data-dir /tmp -P "host=[your-host] port=[your-port] user=[uname] password=[password] dbname=[ledgerdb] sslmode=disable"  --no-algod
```

To start the Indexer so it populates the PostgreSQL database, supply the Algorand Archival node connection details. This can be done by either specifying the Algorand Node data directory (with `--algod`), if the node is on the same machine as the Indexer, or by supplying the algod network host and port string (`--algod-net`) and the proper API token (`--algod-token`). The database needs to be created and running prior to starting the Indexer.

!!! note
    The indexer has a flag `--data-dir` for where to write it's data, which is distinct from the algod data directory in the above paragraph.

```bash
# start with local data directory
$ ./algorand-indexer daemon --data-dir /tmp -P "host=[your-host] port=[your-port] user=[uname] password=[password] dbname=[ledgerdb] sslmode=disable" --algod=~/node/data

# start with networked Algorand node
$ ./algorand-indexer daemon --data-dir /tmp -P "host=[your-host] port=[your-port] user=[uname] password=[password] dbname=[ledgerdb] sslmode=disable" --algod-net="http://[your-host]:[your-port]" --algod-token="[your-api-token]"

```

!!! info
    The initial loading of the Indexer Database will take a considerable amount of time.


## REST API Token and Server

When starting the Indexer, a REST API is exposed. To control access to this API you can you use the `--token` parameter, which allows specifying any desired token. REST API clients will be required to pass this token in their calls in order to return successful searches. The REST API defaults to serving on port 8980. This can be changed by supply a [host:port] value to the Indexer with the `--server` option.

[Click here for Indexer Rest Endpoint specfiles.](../../../rest-apis/restendpoints/)

To turn on indexing for a node, the `isIndexerActive` configuration parameter must be set to `true`. The [Node Configuration](../../reference/config) guide describes setting node configuration properties.

!!! warning
     Turning on indexing with a node will increase the disk space required by the node.

!!! info
    Indexing on a node is only allowed with nodes that have archival mode turned on. 
