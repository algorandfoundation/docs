# Client management

Client management is one of the core capabilities provided by AlgoKit Utils. 
It allows you to create [algod](https://developer.algorand.org/docs/rest-apis/algod), [indexer](https://developer.algorand.org/docs/rest-apis/indexer) 
and [kmd](https://developer.algorand.org/docs/rest-apis/kmd) clients against various networks resolved from environment or specified configuration.

Any AlgoKit Utils function that needs one of these clients will take the underlying `algosdk` classes (`algosdk.v2client.algod.AlgodClient`, `algosdk.v2client.indexer.IndexerClient`, 
`algosdk.kmd.KMDClient`) so inline with the [Modularity](../index.md#core-principles) principle you can use existing logic to get instances of these clients without needing to use the 
Client management capability if you prefer.

To see some usage examples check out the [automated tests](https://github.com/algorandfoundation/algokit-utils-py/blob/main/tests/test_network_clients.py).

## Network configuration

The network configuration is specified using the `AlgoClientConfig` class. This same interface is used to specify the config for algod, indexer and kmd clients.

There are a number of ways to produce one of these configuration objects:

* Manually creating the object, e.g. `AlgoClientConfig(server="https://myalgodnode.com", token="SECRET_TOKEN")`
* `algokit_utils.get_algonode_config(network, config, token)`: Loads an Algod or indexer config against [AlgoNode](https://algonode.io/api/) to either MainNet or TestNet
* `algokit_utils.get_purestake_config(network, config, token)`: Loads an Algod or indexer config against [PureStake](https://developer.purestake.io/) to either MainNet or TestNet
* `algokit_utils.get_default_localnet_config(configOrPort)`: Loads an Algod, Indexer or Kmd config against [LocalNet](https://github.com/algorandfoundation/algokit-cli/blob/main/docs/features/localnet.md) using the default configuration

## Clients

Once you have the configuration for a client, to get the client you can use the following functions:

- `algokit_utils.get_algod_client(config)`: Returns an Algod client for the given configuration or if none is provided retrieves a configuration from the environment using `ALGOD_SERVER`, `ALGOD_TOKEN` and optionally `ALGOD_PORT`.
- `algokit_utils.get_indexer_client(config)`: Returns an Indexer client for given configuration or if none is provided retrieves a configuration from the environment using `INDEXER_SERVER`, `INDEXER_TOKEN` and optionally `INDEXER_PORT`
- `algokit_utils.get_kmd_client_from_algod_client(config)`: - Returns a Kmd client based on the provided algod client configuration, with the assumption the KMD services is running on the same host but a different port (either `KMD_PORT` environment variable or `4002` by default)
