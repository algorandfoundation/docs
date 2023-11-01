# Client management

Client management is one of the core capabilities provided by AlgoKit Utils. It allows you to create (auto-retry) [algod](https://developer.algorand.org/docs/rest-apis/algod), [indexer](https://developer.algorand.org/docs/rest-apis/indexer) and [kmd](https://developer.algorand.org/docs/rest-apis/kmd) clients against various networks resolved from environment or specified configuration.

Any AlgoKit Utils function that needs one of these clients will take the underlying algosdk classes (`algosdk.Algodv2`, `algosdk.Indexer`, `algosdk.Kmd`) so inline with the [Modularity](../index.md#core-principles) principle you can use existing logic to get instances of these clients without needing to use the Client management capability if you prefer, including use of libraries like [useWallet](https://github.com/TxnLab/use-wallet) that have their own configuration mechanism.

To see some usage examples check out the [automated tests](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/network-client.spec.ts).

## Network configuration

The network configuration is specified using the [`AlgoClientConfig`](../code/interfaces/types_network_client.AlgoClientConfig.md) interface. This same interface is used to specify the config for algod, indexer and kmd clients.

There are a number of ways to produce one of these configuration objects:

- Manually specifying an object that conforms with the interface, e.g.
  ```typescript
  {
    server: 'https://myalgodnode.com'
  }
  // Or with the optional values:
  {
    server: 'https://myalgodnode.com',
    port: 443,
    token: 'SECRET_TOKEN'
  }
  ```
- [`algokit.getConfigFromEnvOrDefaults()`](../code/modules/index.md#getconfigfromenvordefaults) - Loads the Algod client config, the Indexer client config and the Kmd config from well-known environment variables; useful to have code that can work across multiple blockchain environments (including LocalNet), without having to change
- [`algokit.getAlgodConfigFromEnvironment()`](../code/modules/index.md#getalgodconfigfromenvironment) - Loads an Algod client config from well-known environment variables; useful to have code that can work across multiple blockchain environments (including LocalNet), without having to change
- [`algokit.getIndexerConfigFromEnvironment()`](../code/modules/index.md#getindexerconfigfromenvironment) - Loads an Indexer client config from well-known environment variables; useful to have code that can work across multiple blockchain environments (including LocalNet), without having to change
- [`algokit.getAlgoNodeConfig(network, config)`](../code/modules/index.md#getalgo) - Loads an Algod or indexer config against [AlgoNode](https://algonode.io/api/) to either MainNet or TestNet
- [`getDefaultLocalNetConfig(configOrPort)`](../code/modules/index.md#getdefaultlocalnetconfig) - Loads an Algod, Indexer or Kmd config against [LocalNet](https://github.com/algorandfoundation/algokit-cli/blob/main/docs/features/localnet.md) using the default configuration

## Clients

Once you have the configuration for a client, to get the client you can use the following functions:

- [`algokit.getAlgoClient(config)`](../code/modules/index.md#getalgoclient) - Returns an Algod client for the given configuration; the client automatically retries on transient HTTP errors; if one isn't provided it retrieves it from the environment
- [`algokit.getAlgoIndexerClient(config)`](../code/modules/index.md#getalgoindexerclient) - Returns an Indexer client for given configuration; if one isn't provided it retrieves it from the environment
- [`algokit.getAlgoKmdClient(config)`](../code/modules/index.md#getalgokmdclient) - Returns a Kmd client for the given configuration; if one isn't provided it retrieves it from the environment

## Automatic retry

When receiving an Algod client from AlgoKit Utils, it will be a special wrapper client that has transient failure retries in there. This is done via the [`AlgoHttpClientWithRetry`](../code/classes/types_algo_http_client_with_retry.AlgoHttpClientWithRetry.md) class.
