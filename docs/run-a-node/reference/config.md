title: Node configuration settings
Nodes can be configured with different options. These options will determine some of the capabilities of the node and whether it functions as a relay node or a non-relay node. This involves setting parameters in the configuration file for either the `algod` or `kmd` process. 

The configuration file (`config.json`) for the `algod` process is located in the node's `data` directory. 
If it does not exist, it needs to be created.
A full example is provided as `config.json.example`.
However, it is strongly recommended to only specify the parameters with non-default values in a custom `config.json` file, otherwise, when the algod software is updated, you may be using older non-recommended values for some of the parameters.

Concretely, the `config.json` for an archival node should usually just be:
```json
{
    "Archival": true
}
```

The configuration file (`kmd_config.json`) for `kmd` is located in the nodes `data/kmd-version` (rename `kmd_config.json.example') directory. 

See [Node Types](../../run-a-node/setup/types.md) for more information.

!!! info
    All changes require the node to be restarted to take effect.

!!! warning
    Changing some parameter values can have drastic negative impact on performance. In particular, never set `IsIndexerActive` to `true`. This activates the very slow deprecated V1 indexer. If indexer is required, use the [V2 indexer](../../../get-details/indexer).
 
# algod Configuration Settings
The `algod` process configuration parameters are shown in the table below.

| Property| Description | Default Value | 
|------|------|------|
| AccountUpdatesStatsInterval | Time interval in nanoseconds for generating accountUpdates telemetry event. | 5000000000 |
| AccountsRebuildSynchronousMode	| 	The synchronous mode used by the ledger database while the account database is being rebuilt. This is not a typical operational usecase, and is expected to happen only on either startup ( after enabling the catchpoint interval, or on certain database upgrades ) or during fast catchup. The values specified here and their meanings are identical to the ones in LedgerSynchronousMode.	| 1 |
| AgreementIncomingBundlesQueueLength | The size of the buffer holding incoming bundles. | 15 |
| AgreementIncomingProposalsQueueLength | The size of the buffer holding incoming proposals. | 50 |
| AgreementIncomingVotesQueueLength | The size of the buffer holding incoming votes. | 20000 |
| AnnounceParticipationKey	| AnnounceParticipationKey specifies that this node should announce its participation key (with the largest stake) to its gossip peers. This allows peers to prioritize our connection, if necessary, in case of a DoS attack. Disabling this means that the peers will not have any additional information to allow them to prioritize our connection.	| TRUE | 
| Archival | When true, the node is configured as an Archival node. Archival nodes retain a full copy of the ledger (blockchain). Non-Archival nodes will delete old blocks and only retain what's need to properly validate blockchain messages (currently the last 1000 blocks). This means non-Archival node ledgers can be significantly smaller than Archival node ledgers. Relays (nodes with a valid NetAddress) are always Archival, regardless of this setting. This may change in the future. If setting this to true for the first time, the existing ledger may need to be deleted to get the historical values stored as the setting only effects current blocks forward. To do this, shutdown the node and delete all .sqlite files within the data/testnet-version directory, except the crash.sqlite file. In most cases this will just be one file. Restart the node and wait for the node to sync. See [Node Types](../../run-a-node/setup/types.md) for more information. | FALSE |
| BaseLoggerDebugLevel	| Specifies the logging level for `algod` (node.log). The levels range from 0 (critical error / silent) to 5 (debug / verbose). The default value is 4 (‘Info’ - fairly verbose).	| 4 |
| BlockDBDir    | BlockDBDir is an optional directory to store the block database. For isolation, the node will create a subdirectory in this location, named by the genesis-id of the network. If not specified, the node will use the ColdDataDir.    |   |
| BlockServiceCustomFallbackEndpoints | A comma delimited list of endpoints which the block service uses to redirect the http requests to in case it does not have the round. If it is not specified, will check EnableBlockServiceFallbackToArchiver. |  |
| BlockServiceMemCap    | BlockServiceMemCap is the memory capacity in bytes which is allowed for the block service to use for HTTP block requests. When it exceeds this capacity, it redirects the block requests to a different node. | 500000000 |
| BroadcastConnectionsLimit	| BroadcastConnectionsLimit specifies the number of connections that will receive broadcast (gossip) messages from this node. If the node has more connections than this number, it will send broadcasts to the top connections by priority (outgoing connections first, then by money held by peers based on their participation key). 0 means no outgoing messages (not even transaction broadcasting to outgoing peers). -1 means unbounded (default).	| -1 | 
| CadaverDirectory  | if this is not set, MakeService will attempt to use ColdDataDir instead.  |   |
| CadaverSizeTarget	| Specifies the maximum size of the agreement.cdv file in bytes. Once full the file will be renamed to agreement.archive.log and a new agreement.cdv will be created. Specified in bytes.	| 0 |
| CatchpointDir | CatchpointDir is an optional directory to store catchpoint files, except for the in-progress temp file, which will use the HotDataDir and is not separately configurable. For isolation, the node will create a subdirectory in this location, named by the genesis-id of the network. If not specified, the node will use the ColdDataDir.   |   |
| CatchpointFileHistoryLength	| CatchpointFileHistoryLength defines how many catchpoint files to store back. 0 means don't store any, -1 mean unlimited and positive number suggest the number of most recent catchpoint files. |	365 | 
| CatchpointInterval	| CatchpointInterval set the interval at which catchpoint are being generated.	| 10000 | 
| CatchpointTracking	| Determines if catchpoints are going to be tracked. The value is interpreted as follows: A value of -1 means "don't track catchpoints". A value of 1 means "track catchpoints as long as CatchpointInterval is also set to a positive non-zero value". If CatchpointInterval <= 0, no catchpoint tracking would be performed. A value of 0 means automatic, which is the default value. In this mode, a non archival node would not track the catchpoints, and an archival node would track the catchpoints as long as CatchpointInterval > 0. Other values of CatchpointTracking would give a warning in the log file, and would behave as if the default value was provided.	| 0 | 
| CatchupBlockDownloadRetryAttempts	| CatchupLedgerDownloadRetryAttempts controls the number of attempt the block fetching would be attempted before giving up catching up to the provided catchpoint.	| 1000 | 
| CatchupBlockValidateMode | A development and testing configuration used by the catchup service. It can be used to omit certain validations to speed up the catchup process, or to apply extra validations which are redundant in normal operation. This field is a bit-field with: bit 0: (default 0) 0: verify the block certificate; 1: skip this validation. bit 1: (default 0) 0: verify payset committed hash in block header matches payset hash; 1: skip this validation. bit 2: (default 0) 0: don't verify the transaction signatures on the block are valid; 1: verify the transaction signatures on block. bit 3: (default 0) 0: don't verify that the hash of the recomputed payset matches the hash of the payset committed in the block header; 1: do perform the above verification. Note: not all permutations of the above bitset are currently functional. In particular, the ones that are functional are: 0  : default behavior. 3  : speed up catchup by skipping necessary validations 12 : perform all validation methods (normal and additional). These extra tests helps to verify the integrity of the compiled executable against previously used executabled, and would not provide any additional security guarantees. | 0 |
| CatchupFailurePeerRefreshRate	| Maximum number of retries a node will attempt to sync.	| 10 | 
| CatchupGossipBlockFetchTimeoutSec	| Controls how long the gossip query for fetching a block from a relay would take before giving up and trying another relay.	| 4 | 
| CatchupHTTPBlockFetchTimeoutSec	| Controls how long the http query for fetching a block from a relay would take before giving up and trying another relay.	| 4 | 
| CatchupLedgerDownloadRetryAttempts	| Controls the number of attempt the ledger fetching would be attempted before giving up catching up to the provided catchpoint.	| 50 | 
| CatchupParallelBlocks	| The maximal number of blocks that catchup will fetch in parallel. If less than Protocol.SeedLookback, then Protocol.SeedLookback will be used as to limit the catchup. | 16 | 
| ColdDataDir   | ColdDataDir is an optional directory to store data that is infrequently accessed by the node. For isolation, the node will create a subdirectory in this location, named by the genesis-id of the network. If not specified, the node will use the runtime supplied datadir. Individual resources may have their own override specified, which would override this setting for that resource. Setting ColdDataDir to a less critical or cheaper disk allows for basic disc tuning.    |   |
| ConnectionsRateLimitingCount	| ConnectionsRateLimitingCount is being used along with ConnectionsRateLimitingWindowSeconds to determine if a connection request should be accepted or not. The gossip network examine all the incoming requests in the past  ConnectionsRateLimitingWindowSeconds seconds that share the same origin. If the total count exceed the ConnectionsRateLimitingCount value, the connection is refused.	| 60 | 
| ConnectionsRateLimitingWindowSeconds	| ConnectionsRateLimitingWindowSeconds is being used in conjunction with ConnectionsRateLimitingCount; see ConnectionsRateLimitingCount description for further information. Providing a zero value in this variable disables the connection rate limiting | 1 | 
| CrashDBDir    | CrashDBDir is an optional directory to store the crash database. For isolation, the node will create a subdirectory in this location, named by the genesis-id of the network. If not specified, the node will use the ColdDataDir.    |   |
| DNSBootstrapID    | DNSBootstrapID specifies the names of a set of DNS SRV records that identify the set of nodes available to connect to. When resolving the bootstrap ID <network> will be replaced by the genesis block's network name. This string uses a URL parsing library and supports optional backup and dedup parameters. 'backup' is used to provide a second DNS entry to use in case the primary is unavailable. dedup is intended to be used to deduplicate SRV records returned from the primary and backup DNS address. If the <name> macro is used in the dedup mask, it must be at the beginning of the expression. This is not typically something a user would configure. For more information see config/dnsbootstrap.go.   | &lt;network&gt;.algorand.network?backup=&lt;network&gt;.algorand.net&dedup=&lt;name&gt;.algorand-&lt;network&gt;.(network\|net)    |
| DNSSecurityFlags	| DNSSecurityFlags instructs `algod` validating DNS responses. Possible flag values 0 - disabled, 1 (dnssecSRV) - validate SRV response, 2 (dnssecRelayAddr) - validate relays' names to addresses resolution, 4 (dnssecTelemetryAddr) - validate telemetry and metrics names to addresses resolution.	| 1 | 
| DeadlockDetection	| Control enabling / disabling deadlock detection. Negative (-1) to disable, positive (1) to enable, 0 for default.	| 0 | 
| DeadlockDetectionThreshold	| The threshold used for deadlock detection, in seconds.	| 30 | 
| DisableAPIAuth    | DisableAPIAuth turns off authentication for public (non-admin) API endpoints. | FALSE |
| DisableLocalhostConnectionRateLimit | Controls whether the incoming connection rate limit would apply for connections that are originating from the local machine. Setting this to "true", allow to create large local-machine networks that won't trip the incoming connection limit observed by relays. | TRUE |
| DisableLedgerLRUCache | DisableLedgerLRUCache disables LRU caches in ledger. Setting it to TRUE might result in significant performance degradation and SHOULD NOT be used for other reasons than testing.    | FALSE |
| DisableNetworking | Disables all the incoming and outgoing communication a node would perform. This is useful when we have a single-node private network, where there is no other nodes that need to be communicated with features like catchpoint catchup would be rendered completly non-operational, and many of the node inner working would be completly dis-functional. | FALSE |
| DisableOutgoingConnectionThrottling	| DisableOutgoingConnectionThrottling disables the connection throttling of the network library, which allow the network library to continuously  disconnect relays based on their relative ( and absolute ) performance.	| FALSE | 
| EnableAccountUpdatesStats | Generate AccountUpdates telemetry event. | FALSE |
| EnableAgreementReporting	| Enable agreement reporting flag. Currently only prints additional period events.	| FALSE | 
| EnableAgreementTimeMetrics	| Enable agreement timing metrics flag	| FALSE | 
| EnableAssembleStats	| Generate AssembleBlockMetrics telemetry event	| FALSE | 
| EnableBlockService 	| EnableBlockService enables the block serving service. The functionality of this depends on NetAddress, which must also be provided. This functionality is required for the catchup.	| FALSE |
| EnableBlockServiceFallbackToArchiver | Controls whether the block service redirects the http requests to an archiver or return StatusNotFound (404) when in does not have the requested round, and BlockServiceCustomFallbackEndpoints is empty. The archiver is randomly selected, if none is available, will return StatusNotFound (404). | FALSE |
| EnableCatchupFromArchiveServers	| 	Controls which peers the catchup service would use in order to catchup. When enabled, the catchup service would use the archive servers before falling back to the relays. On networks that doesn't have archive servers, this becomes a no-op, as the catchup service would have no archive server to pick from, and therefore automatically selects one of the relay nodes.	| FALSE |
| EnableDeveloperAPI 	| EnableDeveloperAPI enables teal/compile, teal/dryrun API endpoints. This functionality is disabled by default.	| FALSE | 
| EnableExperimentalAPI | EnableExperimentalAPI enables experimental API endpoint. Note that these endpoints have no guarantees in terms of functionality or future support.    | FALSE |
| EnableFollowMode  | EnableFollowMode launches the node in "follower" mode. This turns off the agreement service, and APIs related to broadcasting transactions, and enables APIs which can retrieve detailed information from ledger caches and can control the ledger round. | FALSE |
| EnableGossipBlockService	| EnableGossipBlockService enables the block serving service over the gossip network. The functionality of this depends on NetAddress, which must also be provided. This functionality is required for the relays to perform catchup from nodes.	| TRUE | 
| EnableIncomingMessageFilter 	| Enable the filtering of incoming messages.	| FALSE | 
| EnableLedgerService	| EnableLedgerService enables the ledger serving service. The functionality of this depends on NetAddress, which must also be provided. This functionality is required for the catchpoint catchup.	| FALSE | 
| EnableMetricReporting	| Determines if the metrics service for a node is to be enabled. This setting controls metrics being collected from this specific instance of `algod`.  If any instance has metrics enabled, machine-wide metrics are also collected.	| FALSE | 
| EnableOutgoingNetworkMessageFiltering	| Enable the filtering of outgoing messages.	| TRUE | 
| EnableP2P | EnableP2P turns on the peer to peer network.  | FALSE |
| EnablePingHandler	| EnablePingHandler controls whether the gossip node would respond to ping messages with a pong message.	| TRUE | 
| EnableProcessBlockStats	| Generate ProcessBlockMetrics telemetry event.	| FALSE | 
| EnableProfiler	| EnableProfiler enables the go pprof endpoints, should be false if, the `algod` api will be exposed to untrusted individuals.	| FALSE | 
| EnableRequestLogger	| EnableRequestLogger enabled the logging of the incoming requests to the telemetry server.	| FALSE | 
| EnableRuntimeMetrics  | EnableRuntimeMetrics exposes Go runtime metrics in /metrics and via node_exporter. | FALSE |
| EnableTopAccountsReporting    | Enables tracking the accounts with the highest balances and generating a telemetry event after each block. Should be disabled.	| FALSE | 
| EnableTxBacklogRateLimiting   | EnableTxBacklogRateLimiting controls if a rate limiter and congestion manager shouild be attached to the tx backlog enqueue process. If enabled, the over-all TXBacklog Size will be larger by MAX\_PEERS\*TxBacklogReservedCapacityPerPeer.  | TRUE  |
| EnableTxnEvalTracer   | EnableTxnEvalTracer turns on features in the BlockEvaluator which collect data on transactions, exposing them via algod APIs. It will store txn deltas created during block evaluation, potentially consuming much larger amounts of memory.  | FALSE |
| EnableUsageLog    | EnableUsageLog enables 10Hz log of CPU and RAM usage. Also adds 'algod_ram_usage` (number of bytes in use) to /metrics   | FALSE |
| EnableVerbosedTransactionSyncLogging | The transaction sync to write extensive message exchange information to the log file. This option is disabled by default, so that the log files would not grow too rapidly. | FALSE |
| EndpointAddress	| Configures the address the node listens to for REST API calls. Specify an IP and port or just port. For example, 127.0.0.1:0 will listen on a random port on the localhost (preferring 8080)	| 127.0.0.1:0 | 
| FallbackDNSResolverAddress	| The fallback DNS resolver address that would be used if the system resolver would fail to retrieve SRV records. |  | 	
| ForceFetchTransactions | Explicitly configure a node to retrieve all the transactions into it's transaction pool, even if those would not be required as the node doesn't participate in the consensus or used to relay transactions. | FALSE |
| ForceRelayMessages	| ForceRelayMessages indicates whether the network library relay messages even in the case that no NetAddress was specified.	| FALSE | 
| GossipFanout	| The GossipFanout setting sets the maximum number of peers the node will connect to with outgoing connections. If the list of peers is less than this setting, fewer connections will be made. The node will not connect to the same peer multiple times (with outgoing connections).	| 4 | 
| HeartbeatUpdateInterval   | HeartbeatUpdateInterval defines the interval at which the heartbeat information is being sent to the telemetry (when enabled). Defined in seconds. Minimum value is 60.   | 600   |
| HotDataDir    | HotDataDir is an optional directory to store data that is frequently accessed by the node. For isolation, the node will create a subdirectory in this location, named by the genesis-id of the network. If not specified, the node will use the runtime supplied datadir to store this data. Individual resources may have their own override specified, which would override this setting for that resource. Setting HotDataDir to a dedicated high performance disk allows for basic disc tuning.   |   |
| IncomingConnectionsLimit  | IncomingConnectionsLimit specifies the max number of long-lived incoming connections. 0 means no connections allowed. Must be non-negative. Estimating 1.5MB per incoming connection, 1.5MB\*2400 = 3.6GB  | 2400  |
| IncomingMessageFilterBucketCount	| The number of incoming message hashes buckets.	| 5 | 
| IncomingMessageFilterBucketSize	| The size of each incoming message hash bucket.	| 512 | 
| IsIndexerActive	| This setting works in conjunction with the Archival setting. If Archival is set to false this setting does nothing. If it set to true, the node tracks all transactions stored on the node in an indexer and allows two additional REST calls for fast transaction searches. See [Node Types](../../run-a-node/setup/types.md) for more information.	| FALSE | 
| LedgerSynchronousMode	| 	The synchronous mode used by the ledger database. The supported options are: 0 - SQLite continues without syncing as soon as it has handed data off to the operating system. 1 - SQLite database engine will still sync at the most critical moments, but less often than in FULL mode. 2 - SQLite database engine will use the xSync method of the VFS to ensure that all content is safely written to the disk surface prior to continuing. On Mac OS, the data is additionally syncronized via fullfsync. 3 - In addition to what being done in 2, it provides additional durability if the commit is followed closely by a power loss. for further information see the description of SynchronousMode in dbutil.go	| 2 | 
| LogArchiveDir | LogArchiveDir is an optional directory to store the log archive. If not specified, the node will use the ColdDataDir. |   |
| LogArchiveMaxAge	| LogArchiveMaxAge will be parsed by time.ParseDuration(). Valid units are 's' seconds, 'm' minutes, 'h' hours.	| | 
| LogArchiveName	| Text/template for creating log archive filename. Available template vars: Time at start of log: {{.Year}} {{.Month}} {{.Day}} {{.Hour}} {{.Minute}} {{.Second}} Time at end of log: {{.EndYear}} {{.EndMonth}} {{.EndDay}} {{.EndHour}} {{.EndMinute}} {{.EndSecond}}. If the filename ends with .gz or .bz2 it will be compressed. default: "node.archive.log" (no rotation, clobbers previous archive).	| node.archive.log | 
| LogFileDir    | LogFileDir is an optional directory to store the log, node.log If not specified, the node will use the HotDataDir. The -o command line option can be used to override this output location.   |   |
| LogSizeLimit	| Specifies the maximum size of the log file in bytes. Once full, the file will be renamed to node.archive.log and a new node.log will be created. Specified in bytes.	| 1073741824 | 
| MaxAPIBoxPerApplication   | MaxAPIBoxPerApplication defines the maximum total number of boxes per application that will be returned in GetApplicationBoxes REST API responses.    | 100000    |
| MaxAPIResourcesPerAccount | The maximum total number of resources (created assets, created apps, asset holdings, and application local state) per account that will be allowed in AccountInformation REST API responses before returning a 400 Bad Request. Set zero for no limit. | 100000 |
| MaxAcctLookback   | MaxAcctLookback sets the maximum lookback range for account states, i.e. the ledger can answer account states questions for the range Latest-MaxAcctLookback...Latest | 8 |
| MaxBlockHistoryLookback   | BlockHistoryLookback sets the max lookback range for block information. i.e. the block DB can return transaction IDs for questions for the range Latest-MaxBlockHistoryLookback...Latest  | 0 |
| MaxCatchpointDownloadDuration	| 	The maximum duration a client will be keeping the outgoing connection of a catchpoint download request open for processing before shutting it down. Networks that have large catchpoint files, slow connection or slow storage could be a good reason to increase this value. Note that this is a client-side only configuration value, and it's independent of the actual catchpoint file size.	| 43200000000000    |
| MaxConnectionsPerIP 	| MaxAPIBoxPerApplication defines the maximum total number of boxes per application that will be returned in GetApplicationBoxes REST API responses.    | 15    |
| MinCatchpointFileDownloadBytesPerSecond	| 	The minimal download speed that would be considered to be "acceptable" by the catchpoint file fetcher, measured in bytes per seconds. If the provided stream speed drops below this threshold, the connection would be recycled. Note that this field is evaluated per catchpoint "chunk" and not on it's own. If this field is zero, the default of 20480 would be used.	| 20480 | 
| NetAddress	| The address and/or port on which the relay node listens for incoming connections, or blank to ignore incoming connections. Specify an IP and port or just a port. For example, 127.0.0.1:0 will listen on a random port on the localhost.	| | 
| NetworkMessageTraceServer	| 	TraceServer is a host:port to report graph propagation trace info to.	|  | 
| NetworkProtocolVersion	| NetworkProtocolVersion overrides network protocol version ( if present ).	| | 
| NodeExporterListenAddress	| Used to set the specific address for publishing metrics; the Prometheus server connects to this incoming port to retrieve metrics.	| :9100 | 
| NodeExporterPath	| Specifies the path for the node\_exporter binary; this binary is responsible for collecting metrics from the machine and each node instance with metrics enabled.	| ./node\_exporter |
| OptimizeAccountsDatabaseOnStartup	| Controls whether the accounts database would be optimized on algod startup.	| FALSE |
| OutgoingMessageFilterBucketCount	| The number of outgoing message hashes buckets.	| 3 | 
| OutgoingMessageFilterBucketSize	| The size of each outgoing message hash bucket. | 128 | 
| P2PPersistPeerID    | P2PPersistPeerID will write the private key used for the node's PeerID to the P2PPrivateKeyLocation. This is only used when P2PEnable is true. If P2PPrivateKey is not specified, it uses the default location. | FALSE |
| P2PPrivateKeyLocation | P2PPrivateKeyLocation allows the user to specify a custom path to the private key used for the node's PeerID. The private key provided must be an ed25519 private key. This is only used when P2PEnable is true. If the parameter is not set, it uses the default location.   |   |
| ParticipationKeysRefreshInterval | The duration between two consecutive checks to see if new participation keys have been placed on the genesis directory. | 60000000000 |
| PeerConnectionsUpdateInterval	| PeerConnectionsUpdateInterval defines the interval at which the peer connections information is being sent to the telemetry ( when enabled ). Defined in seconds.| 	3600 | 
| PeerPingPeriodSeconds	| Peer Ping Period Seconds. 0 == disable	| 0 | 
| PriorityPeers	| PriorityPeers specifies peer IP addresses that should always get outgoing broadcast messages from this node. |  | 	
| ProposalAssemblyTime | The max amount of time to spend on generating a proposal block. | 250000000 |
| PublicAddress	| The public address to connect to that is advertised to other nodes. For MainNet relays, make sure this entry includes the full SRV host name plus the publicly-accessible port number. A valid entry will avoid "self-gossip" errors in the log.	| | 
| ReconnectTime	| Interval between RPC calls to a random peer to sync from their transactions. Specified in nanoseconds.	| 60000000000 | 
| ReservedFDs	| To make sure the `algod` process does not run out of FDs, `algod` ensures that RLIMIT\_NOFILE exceeds the max number of incoming connections (i.e., IncomingConnectionsLimit) by at least ReservedFDs.  ReservedFDs are meant to leave room for short-lived FDs like DNS queries, SQLite files, etc.	| 256 |
| RestConnectionsHardLimit | The http server does not accept new connections as long we have this many (hard limit) connections already. | 2048 |
| RestConnectionsSoftLimit | When the number of http connections to the REST layer exceeds the soft limit, we start returning http code 429 Too Many Requests. | 1024 |
| RestReadTimeoutSeconds	| Read timeouts passed to the rest http.Server implementation.	| 15 | 
| RestWriteTimeoutSeconds	| Write timeouts passed to the rest http.Server implementation.	| 120 | 
| RunHosted	| Prefer to run `algod` Hosted (under algoh). Observed by `goal` for now.	| FALSE | 
| StateproofDir | StateproofDir is an optional directory to store stateproof data. For isolation, the node will create a subdirectory in this location, named by the genesis-id of the network. If not specified, the node will use the ColdDataDir.    |   |
| StorageEngine | StorageEngine allows to control which type of storage to use for the ledger. Available options are: sqlite (default), pebbledb (experimental, in development).    | sqlite    |
| SuggestedFeeBlockHistory	| The number of blocks to look back for computing the suggested fee. | 	3 | 
| SuggestedFeeSlidingWindowSize	| SuggestedFeeSlidingWindowSize is the number of past blocks that will be considered in computing the suggested fee.	| 50 | 
| TLSCertFile	| TLS Cert File for https serving. | | 	
| TLSKeyFile	| TLS Key File for https serving.	| | 
| TelemetryToLog	| TelemetryToLog records messages to node.log that are normally sent to remote event monitoring.	| TRUE | 
| TrackerDBDir  | TrackerDbDir is an optional directory to store the tracker database. For isolation, the node will create a subdirectory in this location, named by the genesis-id of the network. If not specified, the node will use the HotDataDir. |   |
| TransactionSyncDataExchangeRate | The auto-calculated data exchange rate between each two peers. The unit of the data exchange rate is in bytes per second. Setting the value to zero implies allowing the transaction sync to dynamically calculate the value. | 0 |
| TransactionSyncSignificantMessageThreshold | The threshold used for a transaction sync message before it can be used for calculating the data exchange rate. Setting this to zero would use the default values. The threshold is defined in units of bytes. | 0 |
| TxBacklogReservedCapacityPerPeer  | TxBacklogReservedCapacityPerPeer determines how much dedicated serving capacity the TxBacklog gives each peer.    | 20    |
| TxBacklogServiceRateWindowSeconds | TxBacklogServiceRateWindowSeconds is the window size used to determine the service rate of the txBacklog. | 10    |
| TxBacklogSize | TxBacklogSize is the queue size used for receiving transactions. default of 26000 to approximate 1 block of transactions if EnableTxBacklogRateLimiting enabled, the over-all size will be larger by MAX\_PEERS\*TxBacklogReservedCapacityPerPeer.    | 26000 |
| TxIncomingFilteringFlags  | TxIncomingFilteringFlags instructs algod filtering incoming tx messages. Flag values: 0x00 - disabled. 0x01 (txFilterRawMsg) - check for raw tx message duplicates. 0x02 (txFilterCanonical) - check for canonical tx group duplicates.   | 1 |
| TxIncomingFilterMaxSize   | TxIncomingFilterMaxSize sets the maximum size for the de-duplication cache used by the incoming tx filter only relevant if TxIncomingFilteringFlags is non-zero.  | 500000    |
| TxPoolExponentialIncreaseFactor	| When the transaction pool is full, the priority of a new transaction must be at least TxPoolExponentialIncreaseFactor times greater than the minimum-priority of a transaction already in the pool (otherwise the new transaction is discarded).| 2 | 
| TxPoolSize	| The maximum number of transactions that a node's pending transaction pool can contain.	| 75000 | 
| TxSyncIntervalSeconds	| Interval between RPC calls to a random peer to sync from their transactions. Specified in seconds.	| 60 | 
| TxSyncServeResponseSize	| The max size the sync server would return.	| 1000000 | 
| TxSyncTimeoutSeconds	| Time to wait for a Transaction Sync RPC call to return. Specified in seconds.	| 30 | 
| UseXForwardedForAddressField	| Indicates whether or not the node should use the X-Forwarded-For HTTP Header when determining the source of a connection.  If used, it should be set to the string "X-Forwarded-For", unless the proxy vendor provides another header field.  In the case of CloudFlare proxy, the "CF-Connecting-IP" header field can be used.	| | 
| VerifiedTranscationsCacheSize	| 	The number of transactions that the verified transactions cache would hold before cycling the cache storage in a round-robin fashion.	| 30000 | 



# kmd Configuration Settings
The `kmd` process configuration parameters are shown in the table below.

| Property| Description | Default Value | 
|------|------|------|
| address | Configures the address the node listens to for REST API calls. Specify an IP and port or just port. For example, 127.0.0.1:0 will listen on a random port on the localhost | 127.0.0.1:0 |
| allowed_origins | Configures the whitelist for allowed domains which can access the kmd process. Specify an array of urls that will be white listed. ie {“allowed_origins”: [“https://othersite1.com“, “https://othersite2.com”]} | |
| session_lifetime_secs | Number of seconds for session expirations.| 60 |
