title: Node Configuration Settings
Nodes can be configured with different options. These options will determine some of the capabilities of the node and whether it functions as a relay node or a non-relay node. This involves setting parameters in the configuration file for either the `algod` or `kmd` process. The configuration file (`config.json`) for the `algod` process is located in the node's `data` directory (rename `config.json.example`).  The configuration file (`kmd_config.json`) for `kmd` is located in the nodes `data/kmd-version` (rename `kmd_config.json.example') directory. See [Node Types](../../run-a-node/setup/types.md) for more information.

!!! info
    All changes require the node to be restarted to take effect.
 
# algod Configuration Settings
The `algod` process configuration parameters are shown in the table below.

| Property| Description | Default Value | 
|------|------|------|
| AnnounceParticipationKey	| AnnounceParticipationKey specifies that this node should announce its participation key (with the largest stake) to its gossip peers. This allows peers to prioritize our connection, if necessary, in case of a DoS attack. Disabling this means that the peers will not have any additional information to allow them to prioritize our connection.	| TRUE | 
| Archival | When true, the node is configured as an Archival node. Archival nodes retain a full copy of the ledger (blockchain). Non-Archival nodes will delete old blocks and only retain what's need to properly validate blockchain messages (currently the last 1000 blocks). This means non-Archival node ledgers can be significantly smaller than Archival node ledgers. Relays (nodes with a valid NetAddress) are always Archival, regardless of this setting. This may change in the future. If setting this to true for the first time, the existing ledger may need to be deleted to get the historical values stored as the setting only effects current blocks forward. To do this, shutdown the node and delete all .sqlite files within the data/testnet-version directory, except the crash.sqlite file. In most cases this will just be one file. Restart the node and wait for the node to sync. See [Node Types](../../run-a-node/setup/types.md) for more information. | FALSE |
| BaseLoggerDebugLevel	| Specifies the logging level for `algod` (node.log). The levels range from 0 (critical error / silent) to 5 (debug / verbose). This setting is currently ignored. In the future, it will default to 4 (‘Info’ - fairly verbose).	| 4 |
| BroadcastConnectionsLimit	| BroadcastConnectionsLimit specifies the number of connections that will receive broadcast (gossip) messages from this node. If the node has more connections than this number, it will send broadcasts to the top connections by priority (outgoing connections first, then by money held by peers based on their participation key). 0 means no outgoing messages (not even transaction broadcasting to outgoing peers). -1 means unbounded (default).	| -1 | 
| CadaverSizeTarget	| Specifies the maximum size of the agreement.cdv file in bytes. Once full the file will be renamed to agreement.archive.log and a new agreement.cdv will be created. Specified in bytes.	| 1073741824 | 
| CatchpointFileHistoryLength	| CatchpointFileHistoryLength defines how many catchpoint files to store back. 0 means don't store any, -1 mean unlimited and positive number suggest the number of most recent catchpoint files. |	365 | 
| CatchpointInterval	| CatchpointInterval set the interval at which catchpoint are being generated.	| 10000 | 
| CatchupBlockDownloadRetryAttempts	| CatchupLedgerDownloadRetryAttempts controls the number of attempt the block fetching would be attempted before giving up catching up to the provided catchpoint.	| 1000 | 
| CatchupFailurePeerRefreshRate	| Maximum number of retries a node will attempt to sync.	| 10 | 
| CatchupGossipBlockFetchTimeoutSec	| Controls how long the gossip query for fetching a block from a relay would take before giving up and trying another relay.	| 4 | 
| CatchupHTTPBlockFetchTimeoutSec	| Controls how long the http query for fetching a block from a relay would take before giving up and trying another relay.	| 4 | 
| CatchupLedgerDownloadRetryAttempts	| Controls the number of attempt the ledger fetching would be attempted before giving up catching up to the provided catchpoint.	| 50 | 
| CatchupParallelBlocks	| The maximal number of blocks that catchup will fetch in parallel. If less than Protocol.SeedLookback, then Protocol.SeedLookback will be used as to limit the catchup. | 16 | 
| ConnectionsRateLimitingCount	| ConnectionsRateLimitingCount is being used along with ConnectionsRateLimitingWindowSeconds to determine if a connection request should be accepted or not. The gossip network examine all the incoming requests in the past  ConnectionsRateLimitingWindowSeconds seconds that share the same origin. If the total count exceed the ConnectionsRateLimitingCount value, the connection is refused.	| 60 | 
| ConnectionsRateLimitingWindowSeconds	| ConnectionsRateLimitingWindowSeconds is being used in conjunction with ConnectionsRateLimitingCount; see ConnectionsRateLimitingCount description for further information. Providing a zero value in this variable disables the connection rate limiting | 1 | 
| DeadlockDetection	| Control enabling / disabling deadlock detection. Negative (-1) to disable, positive (1) to enable, 0 for default.	| 0 | 
| DisableOutgoingConnectionThrottling	| DisableOutgoingConnectionThrottling disables the connection throttling of the network library, which allow the network library to continuously  disconnect relays based on their relative ( and absolute ) performance.	| FALSE | 
| DNSBootstrapID	| Specifies the name of a set of DNS SRV records that identify the set of nodes available to connect to. &lt;network&gt; will be replaced by the genesis block’s network name. | &lt;network&gt;.algorand.network |  | 
| DNSSecurityFlags	| DNSSecurityFlags instructs `algod` validating DNS responses. Possible flag values 0 - disabled, 1 (dnssecSRV) - validate SRV response, 2 (dnssecRelayAddr) - validate relays' names to addresses resolution, 4 (dnssecTelemetryAddr) - validate telemetry and metrics names to addresses resolution.	| 1 | 
| EnableAgreementReporting	| Enable agreement reporting flag. Currently only prints additional period events.	| FALSE | 
| EnableAgreementTimeMetrics	| Enable agreement timing metrics flag	| FALSE | 
| EnableAssembleStats	| Generate AssembleBlockMetrics telemetry event	| FALSE | 
| EnableBlockService 	| EnableBlockService enables the block serving service. The functionality of this depends on NetAddress, which must also be provided. This functionality is required for the catchup.	| FALSE | 
| EnableDeveloperAPI 	| EnableDeveloperAPI enables teal/compile, teal/dryrun API endpoints.This functionlity is disabled by default.	| FALSE | 
| EnableGossipBlockService	| EnableGossipBlockService enables the block serving service over the gossip network. The functionality of this depends on NetAddress, which must also be provided. This functionality is required for the relays to perform catchup from nodes.	| TRUE | 
| EnableIncomingMessageFilter 	| Enable the filtering of incoming messages.	| FALSE | 
| EnableLedgerService	| EnableLedgerService enables the ledger serving service. The functionality of this depends on NetAddress, which must also be provided. This functionality is required for the catchpoint catchup.	| FALSE | 
| EnableMetricReporting	| Determines if the metrics service for a node is to be enabled. This setting controls metrics being collected from this specific instance of `algod`.  If any instance has metrics enabled, machine-wide metrics are also collected.	| FALSE | 
| EnableOutgoingNetworkMessageFiltering	| Enable the filtering of outgoing messages.	| TRUE | 
| EnablePingHandler	| EnablePingHandler controls whether the gossip node would respond to ping messages with a pong message.	| TRUE | 
| EnableProcessBlockStats	| Generate ProcessBlockMetrics telemetry event.	| FALSE | 
| EnableProfiler	| EnableProfiler enables the go pprof endpoints, should be false if, the `algod` api will be exposed to untrusted individuals.	| FALSE | 
| EnableRequestLogger	| EnableRequestLogger enabled the logging of the incoming requests to the telemetry server.	| FALSE | 
| EnableTopAccountsReporting	| Enables tracking the accounts with the highest balances and generating a telemetry event after each block. Should be disabled.	| FALSE | 
| EndpointAddress	| Configures the address the node listens to for REST API calls. Specify an IP and port or just port. For example, 127.0.0.1:0 will listen on a random port on the localhost (preferring 8080)	| 127.0.0.1:0 | 
| FallbackDNSResolverAddress	| The fallback DNS resolver address that would be used if the system resolver would fail to retrieve SRV records. |  | 	
| ForceRelayMessages	| ForceRelayMessages indicates whether the network library relay messages even in the case that no NetAddress was specified.	| FALSE | 
| GossipFanout	| The GossipFanout setting sets the maximum number of peers the node will connect to with outgoing connections. If the list of peers is less than this setting, fewer connections will be made. The node will not connect to the same peer multiple times (with outgoing connections).	| 4 | 
| IncomingConnectionsLimit	| Limits the number of incoming connections and is meant for relays. Normal nodes should not use this setting.	| 10000 | 
| IncomingMessageFilterBucketCount	| The number of incoming message hashes buckets.	| 5 | 
| IncomingMessageFilterBucketSize	| The size of each incoming message hash bucket.	| 512 | 
| IsIndexerActive	| This setting works in conjunction with the Archival setting. If Archival is set to false this setting does nothing. If it set to true, the node tracks all transactions stored on the node in an indexer and allows two additional REST calls for fast transaction searches. See [Node Types](../../run-a-node/setup/types.md) for more information.	| FALSE | 
| LogArchiveMaxAge	| LogArchiveMaxAge will be parsed by time.ParseDuration(). Valid units are 's' seconds, 'm' minutes, 'h' hours.	| | 
| LogArchiveName	| Text/template for creating log archive filename. Available template vars: Time at start of log: {{.Year}} {{.Month}} {{.Day}} {{.Hour}} {{.Minute}} {{.Second}} Time at end of log: {{.EndYear}} {{.EndMonth}} {{.EndDay}} {{.EndHour}} {{.EndMinute}} {{.EndSecond}}. If the filename ends with .gz or .bz2 it will be compressed. default: "node.archive.log" (no rotation, clobbers previous archive).	| node.archive.log | 
| LogSizeLimit	| Specifies the maximum size of the log file in bytes. Once full, the file will be renamed to node.archive.log and a new node.log will be created. Specified in bytes.	| 1073741824 | 
| MaxConnectionsPerIP 	| Maximum connections per IP address.	| 30 | 
| NetAddress	| The address and/or port on which the relay node listens for incoming connections, or blank to ignore incoming connections. Specify an IP and port or just a port. For example, 127.0.0.1:0 will listen on a random port on the localhost.	| | 
| NetworkProtocolVersion	| NetworkProtocolVersion overrides network protocol version ( if present ).	| | 
| NodeExporterListenAddress	| Used to set the specific address for publishing metrics; the Prometheus server connects to this incoming port to retrieve metrics.	| :9100 | 
| NodeExporterPath	| Specifies the path for the node_exporter binary; this binary is responsible for collecting metrics from the machine and each node instance with metrics enabled.	| ./node_exporter | 
| OutgoingMessageFilterBucketCount	| The number of outgoing message hashes buckets.	| 3 | 
| OutgoingMessageFilterBucketSize	| The size of each outgoing message hash bucket. | 128 | 
| PeerConnectionsUpdateInterval	| PeerConnectionsUpdateInterval defines the interval at which the peer connections information is being sent to the telemetry ( when enabled ). Defined in seconds.| 	3600 | 
| PeerPingPeriodSeconds	| Peer Ping Period Seconds. 0 == disable	| 0 | 
| PriorityPeers	| PriorityPeers specifies peer IP addresses that should always get outgoing broadcast messages from this node. |  | 	
| PublicAddress	| Public Address to connect to.	| | 
| ReconnectTime	| Interval between RPC calls to a random peer to sync from their transactions. Specified in seconds.	| 60 | 
| ReservedFDs	| To make sure the `algod` process does not run out of FDs, `algod` ensures that RLIMIT_NOFILE exceeds the max number of incoming connections (i.e., IncomingConnectionsLimit) by at least ReservedFDs.  ReservedFDs are meant to leave room for short-lived FDs like DNS queries, SQLite files, etc.	| 256 | 
| RestReadTimeoutSeconds	| Read timeouts passed to the rest http.Server implementation.	| 15 | 
| RestWriteTimeoutSeconds	| Write timeouts passed to the rest http.Server implementation.	| 120 | 
| RunHosted	| Prefer to run `algod` Hosted (under algoh). Observed by `goal` for now.	| FALSE | 
| SuggestedFeeBlockHistory	| The number of blocks to look back for computing the suggested fee. | 	3 | 
| SuggestedFeeSlidingWindowSize	| SuggestedFeeSlidingWindowSize is the number of past blocks that will be considered in computing the suggested fee.	| 50 | 
| TLSCertFile	| TLS Cert File for https serving. | | 	
| TLSKeyFile	| TLS Key File for https serving.	| | 
| TelemetryToLog	| TelemetryToLog records messages to node.log that are normally sent to remote event monitoring.	| TRUE | 
| TxPoolExponentialIncreaseFactor	| When the transaction pool is full, the priority of a new transaction must be at least TxPoolExponentialIncreaseFactor times greater than the minimum-priority of a transaction already in the pool (otherwise the new transaction is discarded).| 2 | 
| TxPoolSize	| The maximum number of transactions that a node's pending transaction pool can contain.	| 15000 | 
| TxSyncIntervalSeconds	| Interval between RPC calls to a random peer to sync from their transactions. Specified in seconds.	| 60 | 
| TxSyncServeResponseSize	| The max size the sync server would return.	| 1000000 | 
| TxSyncTimeoutSeconds	| Time to wait for a Transaction Sync RPC call to return. Specified in seconds.	| 30 | 
| UseXForwardedForAddressField	| UseXForwardedForAddress indicates whether or not the node should use the X-Forwarded-For HTTP Header when determining the source of a connection.  If used, it should be set to the string "X-Forwarded-For", unless the proxy vendor provides another header field.  In the case of CloudFlare proxy, the "CF-Connecting-IP" header field can be used.	| | 


# kmd Configuration Settings
The `kmd` process configuration parameters are shown in the table below.

| Property| Description | Default Value | 
|------|------|------|
| address | Configures the address the node listens to for REST API calls. Specify an IP and port or just port. For example, 127.0.0.1:0 will listen on a random port on the localhost | 127.0.0.1:0 |
| allowed_origins | Configures the whitelist for allowed domains which can access the kmd process. Specify an array of urls that will be white listed. ie {“allowed_origins”: [“https://othersite1.com“, “https://othersite2.com”]} | |
| session_lifetime_secs | Number of seconds for session expirations.| 60 |
