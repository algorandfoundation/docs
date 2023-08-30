title: Conduit

Conduit is a plugin based framework that functions like a data pipeline. Conduit is primarily intended to allow dApp developers to customize the data the dApp wants to monitor/aggregate/act on from the Algorand Blockchain. 

Conduit can be used to:

- Build a notification system for on chain events.
- Power a next generation block explorer.
- Select dApp specific data and act on specific dApp transactions.
- Identify and act on dApp specific Token actions/movements.
- Identify and act on specific Algorand transactions based on Address or other transaction properties.
- Build a custom Indexer for a new [ARC](https://github.com/algorandfoundation/ARCs).
- Send blockchain data to another streaming data platform for additional processing (e.g. RabbitMQ, Kafka, ZeroMQ).
- Build an NFT catalog based on different standards.
- Have Searchability on a massively reduced set of transactions, requiring much less disk space. 

!!!note
	Full documentation for the Conduit framework is available in the [Conduit Github repository](https://github.com/algorand/conduit/blob/master/README.md).


# Installing Conduit

Conduit can be installed by downloading the [built binaries](https://github.com/algorand/conduit/tree/master#download), using a [docker image](https://hub.docker.com/r/algorand/conduit) on docker hub, or [built from source](https://github.com/algorand/conduit/tree/master#install-from-source). Full instructions for Installation of Conduit are described in the [Conduit Github Repository](https://github.com/algorand/conduit/blob/master/README.md).


# Conduit Architecture

The framework consists of three primary plugin components, _Importers_, _Processors_, and _Exporters_. _Importer_ plugins are designed to source data into the pipeline, _Processors_ manipulate or filter the data, and _Exporter_ plugins act on the processed data. Within a given instance of Conduit, the pipeline supports one _importer_ and one _exporter_, while zero or more _processor_ plugins can be used. 

<center>
![Conduit Architecture](/docs/imgs/conduitarch.png){: style="width:700px" align=center }
</center>

Default installations of the conduit binaries include two types of _importers_ ([algod](https://github.com/algorand/conduit/tree/master/conduit/plugins/importers/algod) and [file_reader](https://github.com/algorand/conduit/tree/master/conduit/plugins/importers/filereader)). The Algod plugin is used to source data from an [Algorand Follower node](https://github.com/algorand/conduit/blob/master/docs/tutorials/IndexerWriter.md#node-algod-with-follow-mode) or an [Archival node](https://developer.algorand.org/docs/run-a-node/setup/types/#archival-mode). Using a Follower node is the recommended approach as you have access to more data for use in the _processor_ and also have access to the postgreSQL _exporter_. The file_reader plugin can be used to source block data from the filesystem. Most dApps will most likely use the algod importer. 

The default installation provides one _processor_ that filters the data based on transaction properties. This plugin ([filter_processor](https://github.com/algorand/conduit/tree/master/conduit/plugins/processors/filterprocessor)) will be described in more detail in a subsequent section. 



# CLI Tool

The [AlgoKit CLI Tool](https://github.com/algorandfoundation/algokit-cli) provides a command line interface developers building applications on Algorand.

Using the CLI tool a developer can, among other things, spin up or manage a [local network](#localnet) or initialize a new project from a [Template](#templates).

## LocalNet

The AlgoKit `localnet` sub-command allows a developer to quickly spin up a private network in a local docker container. This is useful for testing and development. This command replaces the need to install a local node manually or with the Sandbox.

Simply running
```sh
$ algokit localnet start
```

