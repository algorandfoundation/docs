title: Node Configuration Settings
Nodes can be configured with different options. These options will determine some of the capabilities of the node and whether it functions as a relay node or a non-relay node. This involves setting parameters in the configuration file for either the `algod` or `kmd` process. The configuration file (`config.json`) for the `algod` process is located in the nodes `data` directory (rename `config.json.example`).  The configuration file (`kmd_config.json`) for `kmd` is located in the nodes `data/kmd-version` (rename `kmd_config.json.example') directory. See [Node Types](../../run-a-node/setup/types.md) for more information.

!!! info
    All changes require the node to be restarted to take effect.
 
# algod Configuration Settings
The `algod` process configuration parameters are shown in the table below.

| Property| Description | Default Value | 
|------|------|------|
| Archival | When true, the node is configured as an Archival node. Archival nodes retain a full copy of the ledger (blockchain). Non-Archival nodes will delete old blocks and only retain what's need to properly validate blockchain messages (currently the last 1000 blocks). This means non-Archival node ledgers can be significantly smaller than Archival node ledgers. Relays (nodes with a valid NetAddress) are always Archival, regardless of this setting. This may change in the future. If setting this to true for the first time, the existing ledger may need to be deleted to get the historical values stored as the setting only effects current blocks forward. To do this, shutdown the node and delete all .sqlite files within the data/testnet-version directory, except the crash.sqlite file. In most cases this will just be one file. Restart the node and wait for the node to sync. See [Node Types](../../run-a-node/setup/types.md) for more information. | FALSE |
| isIndexerActive | This setting works in conjunction with the Archival setting. If Archival is set to false this setting does nothing. If it set to true, the node tracks all transactions stored on the node in an indexer and allows two additional REST calls for fast transaction searches. See [Node Types](../../run-a-node/setup/types.md) for more information. | FALSE |
| GossipFanout | The GossipFanout setting sets the maximum number of peers the node will connect to with outgoing connections. If the list of peers is less than this setting, fewer connections will be made. The node will not connect to the same peer multiple times (with outgoing connections). | 4 |
| NetAddress | The address and/or port on which the relay node listens for incoming connections, or blank to ignore incoming connections. Specify an IP and port or just a port. For example, 127.0.0.1:0 will listen on a random port on the localhost. |  |
| ReconnectTime | Specifies, in seconds, how long the system should wait before trying to reconnect to a peer from which we have disconnected. Specified in seconds. | 60 |
| BaseLoggerDebugLevel | Specifies the logging level for `algod` (node.log). The levels range from 0 (critical error / silent) to 5 (debug / verbose). This setting is currently ignored. In the future, it will default to 4 ('Info' - fairly verbose). | 1 |
| IncomingConnectionsLimit | Limits the number of incoming connections and is meant for relays. Normal nodes should not use this setting. | 10000 |
| EndpointAddress | Configures the address the node listens to for REST API calls. Specify an IP and port or just port. For example, 127.0.0.1:0 will listen on a random port on the localhost (preferring 8080) | 127.0.0.1:0 |
| DNSBootstrapID | Specifies the name of a set of DNS SRV records that identify the set of nodes available to connect to. &lt;network&gt; will be replaced by the genesis block’s network name. | &lt;network&gt;.algorand.network |
| LogSizeLimit | Specifies the maximum size of the log file in bytes. Once full, the file will be renamed to node.archive.log and a new node.log will be created. Specified in bytes. | 1073741824 |
|CatchupFailurePeerRefreshRate | Maximum number of retries a node will attempt to sync. | 10 |
| EnableMetricReporting | Determines if the metrics service for a node is to be enabled. This setting controls metrics being collected from this specific instance of `algod`. If any instance has metrics enabled, machine-wide metrics are also collected. | FALSE |
| NodeExporterListenAddress | Used to set the specific address for publishing metrics; the Prometheus server connects to this incoming port to retrieve metrics. | :9100 |
| NodeExporterPath | Specifies the path for the node_exporter binary; this binary is responsible for collecting metrics from the machine and each node instance with metrics enabled. | ./node_exporter |
| CadaverSizeTarget | Specifies the maximum size of the agreement.cdv file in bytes. Once full the file will be renamed to agreement.archive.log and a new agreement.cdv will be created. Specified in bytes. | 1073741824 |
| EnableTopAccountsReporting | Enables tracking the accounts with the highest balances and generating a telemetry event after each block. Should be disabled. | FALSE |
| TxPoolExponentialIncreaseFactor | When the transaction pool is full, the priority of a new transaction must be at least TxPoolExponentialIncreaseFactor times greater than the minimum-priority of a transaction already in the pool (otherwise the new transaction is discarded). | 2 |
| SuggestedFeeBlockHistory | The number of blocks to look back for computing the suggested fee. | 3 |
| TxPoolSize | The maximum number of transactions that a node's pending transaction pool can contain. | 50000 |
| TxSyncTimeoutSeconds | Time to wait for a Transaction Sync RPC call to return. Specified in seconds. | 30 |
| TxSyncIntervalSeconds | Interval between RPC calls to a random peer to sync from their transactions. Specified in seconds. | 60 | 

# kmd Configuration Settings
The `kmd` process configuration parameters are shown in the table below.

| Property| Description | Default Value | 
|------|------|------|
| address | Configures the address the node listens to for REST API calls. Specify an IP and port or just port. For example, 127.0.0.1:0 will listen on a random port on the localhost | 127.0.0.1:0 |
| allowed_origins | Configures the whitelist for allowed domains which can access the kmd process. Specify an array of urls that will be white listed. ie {“allowed_origins”: [“https://othersite1.com“, “https://othersite2.com”]} | |
