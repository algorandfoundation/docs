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

Conduit can be installed by downloading the [built binaries](https://github.com/algorand/conduit/releases), using a [docker image](https://hub.docker.com/r/algorand/conduit) on docker hub, or [built from source](https://github.com/algorand/conduit/tree/master#install-from-source). Full instructions for Installation of Conduit are described in the [Conduit Github Repository](https://github.com/algorand/conduit/blob/master/README.md).


# Conduit Architecture

The framework consists of three primary plugin components, _Importers_, _Processors_, and _Exporters_. _Importer_ plugins are designed to source data into the pipeline, _Processors_ manipulate or filter the data, and _Exporter_ plugins act on the processed data. Within a given instance of Conduit, the pipeline supports one _importer_ and one _exporter_, while zero or more _processor_ plugins can be used. 

<center>
![Conduit Architecture](/docs/imgs/conduitarch.png){: style="width:700px" align=center }
</center>

Default installations of the conduit binaries include two types of _importers_ ([`algod`](https://github.com/algorand/conduit/tree/master/conduit/plugins/importers/algod) and [`file_reader`](https://github.com/algorand/conduit/tree/master/conduit/plugins/importers/filereader)). The `algod` plugin is used to source data from an [Algorand Follower node](https://github.com/algorand/conduit/blob/master/docs/tutorials/IndexerWriter.md#node-algod-with-follow-mode) or an [Archival node](https://developer.algorand.org/docs/run-a-node/setup/types/#archival-mode). Using a Follower node is the recommended approach as you have access to more data for use in the _processor_ and also have access to the `postgresql` _exporter_. The `file_reader` plugin can be used to source block data from the filesystem. Most dApps will most likely use the `algod` importer. 

The default installation provides one _processor_ that filters the data based on transaction properties. This plugin ([`filter_processor`](https://github.com/algorand/conduit/tree/master/conduit/plugins/processors/filterprocessor)) will be described in more detail in a subsequent section. 

The default installation provides two _exporters_ ([`postgresql`](https://github.com/algorand/conduit/tree/master/conduit/plugins/exporters/postgresql) and [`file_writer`](https://github.com/algorand/conduit/tree/master/conduit/plugins/exporters/filewriter)). The `file_writer` _exporter_, writes block data to the filesystem. The `postgresql` _exporter_ writes block data to a postgreSQL database. 

# Configuring Conduit

Conduit is configured by defining all the plugins the pipeline will support in a YAML file called `conduit.yml`. These plugins must be built and part of the binaries that the specific instance of Conduit is using. This file is used to define which _importer_, _processors_, and _exporter_ plugin a particular instance is using. In addition, individual plugin properties are also configured in the YAML file as well.  For more information on creating and configuring the YAML file see the [Conduit documentation](https://github.com/algorand/conduit/blob/master/README.md#create-conduityml-configuration-file). 

!!!note
	Multiple instances of Conduit can be running simultaneously, with different configurations. No two instances can use the same follower node though.

# Running Conduit

Once Installed and configured, start Conduit with your data directory as an argument:

`./conduit -d data`

!!!warning 
    This data directory should be unique to conduit. This is not the data directory that algod uses.


# Customizing Conduit

In addition to the default plugins described above, the Conduit framework allows custom plugins to be developed allowing dApp developers total customization of how the data is sourced, processed and stored or acted on. This process is described in detail with several tutorials available in the [Conduit Github repository](https://github.com/algorand/conduit/blob/master/docs/PluginDevelopment.md). The Conduit team also hosts a known list of [externally developed plugins](https://github.com/algorand/conduit/blob/master/docs/ExternalPlugins.md). 

!!!note
	The Conduit team is actively looking for sample plugins that illustrate interesting use cases. If you build a plugin and want to share your work, please file a PR on the Conduit repository to add it to that page. 

# Using the Indexer API with an Instance of Conduit

When using Conduit, some dApps may wish to continue to use the Indexer API that Algorand provides. If your application will require this API, you will need to setup the Conduit pipeline as described in the [Conduit documentation](https://github.com/algorand/conduit/blob/master/docs/tutorials/IndexerWriter.md). 

!!!note
	The 2.X versions of Indexer used all the blockchain data since its inception. With Conduit, you can decide how much data you really want to store and search with the Indexer API.

# Filtering Block Data

One of the primary use cases of Conduit is to filter data that a specific dApp is interested in examining or acting on. To accommodate this, the default installation of Conduit provides the `filter_processor`. This filter is configured similarly to the following in `config.yml`.

```sh
name: filter_processor
config:
    # Whether the expression searches inner transactions for matches.
    search-inner: true

    # Whether to include the entire transaction group when the filter
    # conditions are met.
    omit-group-transactions: true

    # The list of filter expressions to use when matching transactions.
    filters:
      - any:
          - tag: "txn.rcv"
            expression-type: "equal"
            expression: "ADDRESS"
```

After specifying the proper name and two basic parameters, one or more filters can be added to the YAML file to specify particular transactions your dApp is interested in. In the above example, Conduit will only store/act on transactions where a specific address is the receiver of a transaction. Developers can set any number of filters, specifying different tags, expression-types and expressions.  The `tag` property has access to all transaction properties defined in the [Developer documentation](https://developer.algorand.org/docs/get-details/transactions/transactions/). In addition, if your instance of Conduit is attached to a Follower node, the filter can also tag specific data in the `ApplyData` set. `ApplyData` contains data that needs to be applied to the ledger, based on the results of a transaction(s). This includes properties like changes to Closing Amount , Application State, etc. The full list of available properties in the `ApplyData` set is available in the [Conduit Documentation](https://github.com/algorand/conduit/blob/master/conduit/plugins/processors/filterprocessor/Filter_tags.md). For more information and examples of other filters see the [Conduit filter examples](https://github.com/algorand/conduit/tree/master/conduit/plugins/processors/filterprocessor#examples).


