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

!!! info
    Archival nodes retain a full copy of the ledger (blockchain). Non-Archival nodes will delete old blocks and only retain what's needed to properly validate blockchain messages (currently the last 1000 blocks). Archival nodes can be used to populate indexer data. See chart below for more details.


See [Node Types](../../run-a-node/setup/types.md) for more information.

!!! info
    All changes require the node to be restarted to take effect.

!!! warning
    Changing some parameter values can have drastic negative impact on performance. In particular, never set `IsIndexerActive` to `true`. This activates the very slow deprecated V1 indexer. If indexer is required, use the [V2 indexer](../../../get-details/indexer).

# algod Configuration Settings
The `algod` process configuration parameters are shown in the table below.

| Property| Description | Default Value |
|------|------|------|
| Version | Version tracks the current version of the defaults so we can migrate old -> new<br>This is specifically important whenever we decide to change the default value<br>for an existing parameter. This field tag must be updated any time we add a new version. | 35 |
| Archival | Archival nodes retain a full copy of the block history. Non-Archival nodes will delete old blocks and only retain what's need to properly validate blockchain messages (the precise number of recent blocks depends on the consensus parameters. Currently the last 1321 blocks are required). This means that non-Archival nodes require significantly less storage than Archival nodes.  If setting this to true for the first time, the existing ledger may need to be deleted to get the historical values stored as the setting only affects current blocks forward. To do this, shutdown the node and delete all .sqlite files within the data/testnet-version directory, except the crash.sqlite file. Restart the node and wait for the node to sync. | false |
| GossipFanout | GossipFanout sets the maximum number of peers the node will connect to with outgoing connections. If the list of peers is less than this setting, fewer connections will be made. The node will not connect to the same peer multiple times (with outgoing connections). | 4 |
| NetAddress | NetAddress is the address and/or port on which a node listens for incoming connections, or blank to ignore incoming connections. Specify an IP and port or just a port. For example, 127.0.0.1:0 will listen on a random port on the localhost. |  |
| ReconnectTime | ReconnectTime is deprecated and unused. | 60000000000 |
| PublicAddress | PublicAddress is the public address to connect to that is advertised to other nodes.<br>For MainNet relays, make sure this entry includes the full SRV host name<br>plus the publicly-accessible port number.<br>A valid entry will avoid "self-gossip" and is used for identity exchange<br>to de-duplicate redundant connections |  |
| MaxConnectionsPerIP | MaxConnectionsPerIP is the maximum number of connections allowed per IP address. | 8 |
| PeerPingPeriodSeconds | PeerPingPeriodSeconds is deprecated and unused. | 0 |
| TLSCertFile | TLSCertFile is the certificate file used for the websocket network if povided. |  |
| TLSKeyFile | TLSKeyFile is the key file used for the websocket network if povided. |  |
| BaseLoggerDebugLevel | BaseLoggerDebugLevel specifies the logging level for algod (node.log). The levels range from 0 (critical error / silent) to 5 (debug / verbose). The default value is 4 (‘Info’ - fairly verbose). | 4 |
| CadaverSizeTarget | CadaverSizeTarget specifies the maximum size of the agreement.cfv file in bytes. Once full the file will be renamed to agreement.archive.log and a new agreement.cdv will be created. | 0 |
| CadaverDirectory | if this is not set, MakeService will attempt to use ColdDataDir instead |  |
| HotDataDir | HotDataDir is an optional directory to store data that is frequently accessed by the node.<br>For isolation, the node will create a subdirectory in this location, named by the genesis-id of the network.<br>If not specified, the node will use the runtime supplied datadir to store this data.<br>Individual resources may have their own override specified, which would override this setting for that resource.<br>Setting HotDataDir to a dedicated high performance disk allows for basic disc tuning. |  |
| ColdDataDir | ColdDataDir is an optional directory to store data that is infrequently accessed by the node.<br>For isolation, the node will create a subdirectory in this location, named by the genesis-id of the network.<br>If not specified, the node will use the runtime supplied datadir.<br>Individual resources may have their own override specified, which would override this setting for that resource.<br>Setting ColdDataDir to a less critical or cheaper disk allows for basic disc tuning. |  |
| TrackerDBDir | TrackerDbDir is an optional directory to store the tracker database.<br>For isolation, the node will create a subdirectory in this location, named by the genesis-id of the network.<br>If not specified, the node will use the HotDataDir. |  |
| BlockDBDir | BlockDBDir is an optional directory to store the block database.<br>For isolation, the node will create a subdirectory in this location, named by the genesis-id of the network.<br>If not specified, the node will use the ColdDataDir. |  |
| CatchpointDir | CatchpointDir is an optional directory to store catchpoint files,<br>except for the in-progress temp file, which will use the HotDataDir and is not separately configurable.<br>For isolation, the node will create a subdirectory in this location, named by the genesis-id of the network.<br>If not specified, the node will use the ColdDataDir. |  |
| StateproofDir | StateproofDir is an optional directory to persist state about observed and issued state proof messages.<br>For isolation, the node will create a subdirectory in this location, named by the genesis-id of the network.<br>If not specified, the node will use the HotDataDir. |  |
| CrashDBDir | CrashDBDir is an optional directory to persist agreement's consensus participation state.<br>For isolation, the node will create a subdirectory in this location, named by the genesis-id of the network.<br>If not specified, the node will use the HotDataDir |  |
| LogFileDir | LogFileDir is an optional directory to store the log, node.log<br>If not specified, the node will use the HotDataDir.<br>The -o command line option can be used to override this output location. |  |
| LogArchiveDir | LogArchiveDir is an optional directory to store the log archive.<br>If not specified, the node will use the ColdDataDir. |  |
| IncomingConnectionsLimit | IncomingConnectionsLimit specifies the max number of incoming connections<br>for the gossip protocol configured in NetAddress. 0 means no connections allowed. Must be non-negative.<br>Estimating 1.5MB per incoming connection, 1.5MB*2400 = 3.6GB | 2400 |
| P2PHybridIncomingConnectionsLimit | P2PHybridIncomingConnectionsLimit is used as IncomingConnectionsLimit for P2P connections in hybrid mode.<br>For pure P2P nodes IncomingConnectionsLimit is used. | 1200 |
| BroadcastConnectionsLimit | BroadcastConnectionsLimit specifies the number of connections that<br>will receive broadcast (gossip) messages from this node. If the<br>node has more connections than this number, it will send broadcasts<br>to the top connections by priority (outgoing connections first, then<br>by money held by peers based on their participation key). 0 means<br>no outgoing messages (not even transaction broadcasting to outgoing<br>peers). -1 means unbounded (default). | -1 |
| AnnounceParticipationKey | AnnounceParticipationKey specifies that this node should announce its<br>participation key (with the largest stake) to its gossip peers.  This<br>allows peers to prioritize our connection, if necessary, in case of a<br>DoS attack.  Disabling this means that the peers will not have any<br>additional information to allow them to prioritize our connection. | true |
| PriorityPeers | PriorityPeers specifies peer IP addresses that should always get<br>outgoing broadcast messages from this node. |  |
| ReservedFDs | ReservedFDs is used to make sure the algod process does not run out of file descriptors (FDs). Algod ensures<br>that RLIMIT_NOFILE >= IncomingConnectionsLimit + RestConnectionsHardLimit +<br>ReservedFDs. ReservedFDs are meant to leave room for short-lived FDs like<br>DNS queries, SQLite files, etc. This parameter shouldn't be changed.<br>If RLIMIT_NOFILE < IncomingConnectionsLimit + RestConnectionsHardLimit + ReservedFDs<br>then either RestConnectionsHardLimit or IncomingConnectionsLimit decreased. | 256 |
| EndpointAddress | EndpointAddress configures the address the node listens to for REST API calls. Specify an IP and port or just port. For example, 127.0.0.1:0 will listen on a random port on the localhost (preferring 8080). | 127.0.0.1 |
| EnablePrivateNetworkAccessHeader | Respond to Private Network Access preflight requests sent to the node. Useful when a public website is trying to access a node that's hosted on a local network. | false |
| RestReadTimeoutSeconds | RestReadTimeoutSeconds is passed to the API servers rest http.Server implementation. | 15 |
| RestWriteTimeoutSeconds | RestWriteTimeoutSeconds is passed to the API servers rest http.Server implementation. | 120 |
| DNSBootstrapID | DNSBootstrapID specifies the names of a set of DNS SRV records that identify the set of nodes available to connect to.<br>This is applicable to both relay and archival nodes - they are assumed to use the same DNSBootstrapID today.<br>When resolving the bootstrap ID <network> will be replaced by the genesis block's network name. This string uses a URL<br>parsing library and supports optional backup and dedup parameters. 'backup' is used to provide a second DNS entry to use<br>in case the primary is unavailable. dedup is intended to be used to deduplicate SRV records returned from the primary<br>and backup DNS address. If the <name> macro is used in the dedup mask, it must be at the beginning of the expression.<br>This is not typically something a user would configure. For more information see config/dnsbootstrap.go. | &lt;network&gt;.algorand.network?backup=&lt;network&gt;.algorand.net&amp;dedup=&lt;name&gt;.algorand-&lt;network&gt;.(network|net) |
| LogSizeLimit | LogSizeLimit is the log file size limit in bytes. When set to 0 logs will be written to stdout. | 1073741824 |
| LogArchiveName | LogArchiveName text/template for creating log archive filename.<br>Available template vars:<br>Time at start of log: {{.Year}} {{.Month}} {{.Day}} {{.Hour}} {{.Minute}} {{.Second}}<br>Time at end of log: {{.EndYear}} {{.EndMonth}} {{.EndDay}} {{.EndHour}} {{.EndMinute}} {{.EndSecond}}<br><br>If the filename ends with .gz or .bz2 it will be compressed.<br><br>default: "node.archive.log" (no rotation, clobbers previous archive) | node.archive.log |
| LogArchiveMaxAge | LogArchiveMaxAge will be parsed by time.ParseDuration().<br>Valid units are 's' seconds, 'm' minutes, 'h' hours |  |
| CatchupFailurePeerRefreshRate | CatchupFailurePeerRefreshRate is the maximum number of consecutive attempts to catchup after which we replace the peers we're connected to. | 10 |
| NodeExporterListenAddress | NodeExporterListenAddress is used to set the specific address for publishing metrics; the Prometheus server connects to this incoming port to retrieve metrics. |  |
| EnableMetricReporting | EnableMetricReporting determines if the metrics service for a node is to be enabled. This setting controls metrics being collected from this specific instance of algod. If any instance has metrics enabled, machine-wide metrics are also collected. | false |
| EnableTopAccountsReporting | EnableTopAccountsReporting enable top accounts reporting flag. Deprecated, do not use. | false |
| EnableAgreementReporting | EnableAgreementReporting controls the agreement reporting flag. Currently only prints additional period events. | false |
| EnableAgreementTimeMetrics | EnableAgreementTimeMetrics controls the agreement timing metrics flag. | false |
| NodeExporterPath | NodeExporterPath is the path to the node_exporter binary. | ./node_exporter |
| FallbackDNSResolverAddress | FallbackDNSResolverAddress defines the fallback DNS resolver address that would be used if the system resolver would fail to retrieve SRV records. |  |
| TxPoolExponentialIncreaseFactor | TxPoolExponentialIncreaseFactor exponential increase factor of transaction pool's fee threshold, should always be 2 in production. | 2 |
| SuggestedFeeBlockHistory | SuggestedFeeBlockHistory is deprecated and unused. | 3 |
| TxBacklogServiceRateWindowSeconds | TxBacklogServiceRateWindowSeconds is the window size used to determine the service rate of the txBacklog | 10 |
| TxBacklogReservedCapacityPerPeer | TxBacklogReservedCapacityPerPeer determines how much dedicated serving capacity the TxBacklog gives each peer | 20 |
| TxBacklogAppTxRateLimiterMaxSize | TxBacklogAppTxRateLimiterMaxSize denotes a max size for the tx rate limiter<br>calculated as "a thousand apps on a network of thousand of peers" | 1048576 |
| TxBacklogAppTxPerSecondRate | TxBacklogAppTxPerSecondRate determines a target app per second rate for the app tx rate limiter | 100 |
| TxBacklogRateLimitingCongestionPct | TxBacklogRateLimitingCongestionRatio determines the backlog filling threshold percentage at which the app limiter kicks in<br>or the tx backlog rate limiter kicks off. | 50 |
| EnableTxBacklogAppRateLimiting | EnableTxBacklogAppRateLimiting controls if an app rate limiter should be attached to the tx backlog enqueue process | true |
| TxBacklogAppRateLimitingCountERLDrops | TxBacklogAppRateLimitingCountERLDrops feeds messages dropped by the ERL congestion manager & rate limiter (enabled by<br>EnableTxBacklogRateLimiting) to the app rate limiter (enabled by EnableTxBacklogAppRateLimiting), so that all TX messages<br>are counted. This provides more accurate rate limiting for the app rate limiter, at the potential expense of additional<br>deserialization overhead. | false |
| EnableTxBacklogRateLimiting | EnableTxBacklogRateLimiting controls if a rate limiter and congestion manager should be attached to the tx backlog enqueue process<br>if enabled, the over-all TXBacklog Size will be larger by MAX_PEERS*TxBacklogReservedCapacityPerPeer | true |
| TxBacklogSize | TxBacklogSize is the queue size used for receiving transactions. default of 26000 to approximate 1 block of transactions<br>if EnableTxBacklogRateLimiting enabled, the over-all size will be larger by MAX_PEERS*TxBacklogReservedCapacityPerPeer | 26000 |
| TxPoolSize | TxPoolSize is the number of transactions in the transaction pool buffer. | 75000 |
| TxSyncTimeoutSeconds | number of seconds allowed for syncing transactions | 30 |
| TxSyncIntervalSeconds | TxSyncIntervalSeconds number of seconds between transaction synchronizations. | 60 |
| IncomingMessageFilterBucketCount | IncomingMessageFilterBucketCount is the number of incoming message hash buckets. | 5 |
| IncomingMessageFilterBucketSize | IncomingMessageFilterBucketSize is the size of each incoming message hash bucket. | 512 |
| OutgoingMessageFilterBucketCount | OutgoingMessageFilterBucketCount is the number of outgoing message hash buckets. | 3 |
| OutgoingMessageFilterBucketSize | OutgoingMessageFilterBucketSize is the size of each outgoing message hash bucket. | 128 |
| EnableOutgoingNetworkMessageFiltering | EnableOutgoingNetworkMessageFiltering enable the filtering of outgoing messages | true |
| EnableIncomingMessageFilter | EnableIncomingMessageFilter enable the filtering of incoming messages. | false |
| DeadlockDetection | DeadlockDetection controls enabling or disabling deadlock detection.<br>negative (-1) to disable, positive (1) to enable, 0 for default. | 0 |
| DeadlockDetectionThreshold | DeadlockDetectionThreshold is the threshold used for deadlock detection, in seconds. | 30 |
| RunHosted | RunHosted configures whether to run algod in Hosted mode (under algoh). Observed by `goal` for now. | false |
| CatchupParallelBlocks | CatchupParallelBlocks is the maximum number of blocks that catchup will fetch in parallel.<br>If less than Protocol.SeedLookback, then Protocol.SeedLookback will be used as to limit the catchup.<br>Setting this variable to 0 would disable the catchup | 16 |
| EnableAssembleStats | EnableAssembleStats specifies whether or not to emit the AssembleBlockMetrics telemetry event. |  |
| EnableProcessBlockStats | EnableProcessBlockStats specifies whether or not to emit the ProcessBlockMetrics telemetry event. |  |
| SuggestedFeeSlidingWindowSize | SuggestedFeeSlidingWindowSize is deprecated and unused. | 50 |
| TxSyncServeResponseSize | TxSyncServeResponseSize the max size the sync server would return. | 1000000 |
| UseXForwardedForAddressField | UseXForwardedForAddressField indicates whether or not the node should use the X-Forwarded-For HTTP Header when<br>determining the source of a connection.  If used, it should be set to the string "X-Forwarded-For", unless the<br>proxy vendor provides another header field.  In the case of CloudFlare proxy, the "CF-Connecting-IP" header<br>field can be used.<br>This setting does not support multiple X-Forwarded-For HTTP headers or multiple values in in the header and always uses the last value<br>from the last X-Forwarded-For HTTP header that corresponds to a single reverse proxy (even if it received the request from another reverse proxy or adversary node).<br><br>WARNING: By enabling this option, you are trusting peers to provide accurate forwarding addresses.<br>Bad actors can easily spoof these headers to circumvent this node's rate and connection limiting<br>logic. Do not enable this if your node is publicly reachable or used by untrusted parties. |  |
| ForceRelayMessages | ForceRelayMessages indicates whether the network library should relay messages even in the case that no NetAddress was specified. | false |
| ConnectionsRateLimitingWindowSeconds | ConnectionsRateLimitingWindowSeconds is being used along with ConnectionsRateLimitingCount;<br>see ConnectionsRateLimitingCount description for further information. Providing a zero value<br>in this variable disables the connection rate limiting. | 1 |
| ConnectionsRateLimitingCount | ConnectionsRateLimitingCount is being used along with ConnectionsRateLimitingWindowSeconds to determine if<br>a connection request should be accepted or not. The gossip network examines all the incoming requests in the past<br>ConnectionsRateLimitingWindowSeconds seconds that share the same origin. If the total count exceed the ConnectionsRateLimitingCount<br>value, the connection is refused. | 60 |
| EnableRequestLogger | EnableRequestLogger enabled the logging of the incoming requests to the telemetry server. | false |
| PeerConnectionsUpdateInterval | PeerConnectionsUpdateInterval defines the interval at which the peer connections information is sent to<br>telemetry (when enabled). Defined in seconds. | 3600 |
| HeartbeatUpdateInterval | HeartbeatUpdateInterval defines the interval at which the heartbeat information is being sent to the<br>telemetry (when enabled). Defined in seconds. Minimum value is 60. | 600 |
| EnableProfiler | EnableProfiler enables the go pprof endpoints, should be false if<br>the algod api will be exposed to untrusted individuals | false |
| EnableRuntimeMetrics | EnableRuntimeMetrics exposes Go runtime metrics in /metrics and via node_exporter. | false |
| EnableNetDevMetrics | EnableNetDevMetrics exposes network interface total bytes sent/received metrics in /metrics | false |
| TelemetryToLog | TelemetryToLog configures whether to record messages to node.log that are normally only sent to remote event monitoring. | true |
| DNSSecurityFlags | DNSSecurityFlags instructs algod validating DNS responses.<br>Possible fla values<br>0x00 - disabled<br>0x01 (dnssecSRV) - validate SRV response<br>0x02 (dnssecRelayAddr) - validate relays' names to addresses resolution<br>0x04 (dnssecTelemetryAddr) - validate telemetry and metrics names to addresses resolution<br>0x08 (dnssecTXT) - validate TXT response<br>... | 9 |
| EnablePingHandler | EnablePingHandler controls whether the gossip node would respond to ping messages with a pong message. | true |
| DisableOutgoingConnectionThrottling | DisableOutgoingConnectionThrottling disables the connection throttling of the network library, which<br>allow the network library to continuously disconnect relays based on their relative (and absolute) performance. | false |
| NetworkProtocolVersion | NetworkProtocolVersion overrides network protocol version ( if present ) |  |
| CatchpointInterval | CatchpointInterval sets the interval at which catchpoint are being generated. Setting this to 0 disables the catchpoint from being generated.<br>See CatchpointTracking for more details. | 10000 |
| CatchpointFileHistoryLength | CatchpointFileHistoryLength defines how many catchpoint files to store.<br>0 means don't store any, -1 mean unlimited and positive number suggest the maximum number of most recent catchpoint files to store. | 365 |
| EnableGossipService | EnableGossipService enables the gossip network HTTP websockets endpoint. The functionality of this depends on NetAddress, which must also be provided.<br>This functionality is required for serving gossip traffic. | true |
| EnableLedgerService | EnableLedgerService enables the ledger serving service. The functionality of this depends on NetAddress, which must also be provided.<br>This functionality is required for the catchpoint catchup. | false |
| EnableBlockService | EnableBlockService controls whether to enables the block serving service. The functionality of this depends on NetAddress, which must also be provided.<br>This functionality is required for catchup. | false |
| EnableGossipBlockService | EnableGossipBlockService enables the block serving service over the gossip network. The functionality of this depends on NetAddress, which must also be provided.<br>This functionality is required for the relays to perform catchup from nodes. | true |
| CatchupHTTPBlockFetchTimeoutSec | CatchupHTTPBlockFetchTimeoutSec controls how long the http query for fetching a block from a relay would take before giving up and trying another relay. | 4 |
| CatchupGossipBlockFetchTimeoutSec | CatchupGossipBlockFetchTimeoutSec controls how long the gossip query for fetching a block from a relay would take before giving up and trying another relay. | 4 |
| CatchupLedgerDownloadRetryAttempts | CatchupLedgerDownloadRetryAttempts controls the number of attempt the ledger fetching would be attempted before giving up catching up to the provided catchpoint. | 50 |
| CatchupBlockDownloadRetryAttempts | CatchupBlockDownloadRetryAttempts controls the number of attempts the block fetcher would make before giving up on a provided catchpoint. | 1000 |
| EnableDeveloperAPI | EnableDeveloperAPI enables teal/compile and teal/dryrun API endpoints.<br>This functionality is disabled by default. | false |
| OptimizeAccountsDatabaseOnStartup | OptimizeAccountsDatabaseOnStartup controls whether the accounts database would be optimized<br>on algod startup. | false |
| CatchpointTracking | CatchpointTracking determines if catchpoints are going to be tracked. The value is interpreted as follows:<br>A value of -1 means "don't track catchpoints".<br>A value of 1 means "track catchpoints as long as CatchpointInterval > 0".<br>A value of 2 means "track catchpoints and always generate catchpoint files as long as CatchpointInterval > 0".<br>A value of 0 means automatic, which is the default value. In this mode, a non archival node would not track the catchpoints, and an archival node would track the catchpoints as long as CatchpointInterval > 0.<br>Other values of CatchpointTracking would behave as if the default value was provided. | 0 |
| LedgerSynchronousMode | LedgerSynchronousMode defines the synchronous mode used by the ledger database. The supported options are:<br>0 - SQLite continues without syncing as soon as it has handed data off to the operating system.<br>1 - SQLite database engine will still sync at the most critical moments, but less often than in FULL mode.<br>2 - SQLite database engine will use the xSync method of the VFS to ensure that all content is safely written to the disk surface prior to continuing. On Mac OS, the data is additionally syncronized via fullfsync.<br>3 - In addition to what being done in 2, it provides additional durability if the commit is followed closely by a power loss.<br>for further information see the description of SynchronousMode in dbutil.go | 2 |
| AccountsRebuildSynchronousMode | AccountsRebuildSynchronousMode defines the synchronous mode used by the ledger database while the account database is being rebuilt. This is not a typical operational use-case,<br>and is expected to happen only on either startup (after enabling the catchpoint interval, or on certain database upgrades) or during fast catchup. The values specified here<br>and their meanings are identical to the ones in LedgerSynchronousMode. | 1 |
| MaxCatchpointDownloadDuration | MaxCatchpointDownloadDuration defines the maximum duration a client will be keeping the outgoing connection of a catchpoint download request open for processing before<br>shutting it down. Networks that have large catchpoint files, slow connection or slow storage could be a good reason to increase this value. Note that this is a client-side only<br>configuration value, and it's independent of the actual catchpoint file size. | 43200000000000 |
| MinCatchpointFileDownloadBytesPerSecond | MinCatchpointFileDownloadBytesPerSecond defines the minimal download speed that would be considered to be "acceptable" by the catchpoint file fetcher, measured in bytes per seconds. If the<br>provided stream speed drops below this threshold, the connection would be recycled. Note that this field is evaluated per catchpoint "chunk" and not on it's own. If this field is zero,<br>the default of 20480 would be used. | 20480 |
| NetworkMessageTraceServer | NetworkMessageTraceServer is a host:port address to report graph propagation trace info to. |  |
| VerifiedTranscationsCacheSize | VerifiedTranscationsCacheSize defines the number of transactions that the verified transactions cache would hold before cycling the cache storage in a round-robin fashion. | 150000 |
| DisableLocalhostConnectionRateLimit | DisableLocalhostConnectionRateLimit controls whether the incoming connection rate limit would apply for<br>connections that are originating from the local machine. Setting this to "true", allow to create large<br>local-machine networks that won't trip the incoming connection limit observed by relays. | true |
| BlockServiceCustomFallbackEndpoints | BlockServiceCustomFallbackEndpoints is a comma delimited list of endpoints which the block service uses to<br>redirect the http requests to in case it does not have the round. If empty, the block service will return<br>StatusNotFound (404) |  |
| CatchupBlockValidateMode | CatchupBlockValidateMode is a development and testing configuration used by the catchup service.<br>It can be used to omit certain validations to speed up the catchup process, or to apply extra validations which are redundant in normal operation.<br>This field is a bit-field with:<br>bit 0: (default 0) 0: verify the block certificate; 1: skip this validation<br>bit 1: (default 0) 0: verify payset committed hash in block header matches payset hash; 1: skip this validation<br>bit 2: (default 0) 0: don't verify the transaction signatures on the block are valid; 1: verify the transaction signatures on block<br>bit 3: (default 0) 0: don't verify that the hash of the recomputed payset matches the hash of the payset committed in the block header; 1: do perform the above verification<br>Note: not all permutations of the above bitset are currently functional. In particular, the ones that are functional are:<br>0  : default behavior.<br>3  : speed up catchup by skipping necessary validations<br>12 : perform all validation methods (normal and additional). These extra tests helps to verify the integrity of the compiled executable against<br>     previously used executabled, and would not provide any additional security guarantees. | 0 |
| EnableAccountUpdatesStats | EnableAccountUpdatesStats specifies whether or not to emit the AccountUpdates telemetry event. | false |
| AccountUpdatesStatsInterval | AccountUpdatesStatsInterval is the time interval in nanoseconds between accountUpdates telemetry events. | 5000000000 |
| ParticipationKeysRefreshInterval | ParticipationKeysRefreshInterval is the duration between two consecutive checks to see if new participation<br>keys have been placed on the genesis directory. Deprecated and unused. | 60000000000 |
| DisableNetworking | DisableNetworking disables all the incoming and outgoing communication a node would perform. This is useful<br>when we have a single-node private network, where there are no other nodes that need to be communicated with.<br>Features like catchpoint catchup would be rendered completely non-operational, and many of the node inner<br>working would be completely dis-functional. | false |
| ForceFetchTransactions | ForceFetchTransactions allows to explicitly configure a node to retrieve all the transactions<br>into it's transaction pool, even if those would not be required as the node doesn't<br>participate in consensus and is not used to relay transactions. | false |
| EnableVerbosedTransactionSyncLogging | EnableVerbosedTransactionSyncLogging enables the transaction sync to write extensive<br>message exchange information to the log file. This option is disabled by default,<br>so that the log files would not grow too rapidly. | false |
| TransactionSyncDataExchangeRate | TransactionSyncDataExchangeRate overrides the auto-calculated data exchange rate between each<br>two peers. The unit of the data exchange rate is in bytes per second. Setting the value to<br>zero implies allowing the transaction sync to dynamically calculate the value. | 0 |
| TransactionSyncSignificantMessageThreshold | TransactionSyncSignificantMessageThreshold define the threshold used for a transaction sync<br>message before it can be used for calculating the data exchange rate. Setting this to zero<br>would use the default values. The threshold is defined in units of bytes. | 0 |
| ProposalAssemblyTime | ProposalAssemblyTime is the max amount of time to spend on generating a proposal block. | 500000000 |
| RestConnectionsSoftLimit | RestConnectionsSoftLimit is the maximum number of active requests the API server<br>When the number of http connections to the REST layer exceeds the soft limit,<br>we start returning http code 429 Too Many Requests. | 1024 |
| RestConnectionsHardLimit | RestConnectionsHardLimit is the maximum number of active connections the API server will accept before closing requests with no response. | 2048 |
| MaxAPIResourcesPerAccount | MaxAPIResourcesPerAccount sets the maximum total number of resources (created assets, created apps,<br>asset holdings, and application local state) per account that will be allowed in AccountInformation<br>REST API responses before returning a 400 Bad Request. Set zero for no limit. | 100000 |
| AgreementIncomingVotesQueueLength | AgreementIncomingVotesQueueLength sets the size of the buffer holding incoming votes. | 20000 |
| AgreementIncomingProposalsQueueLength | AgreementIncomingProposalsQueueLength sets the size of the buffer holding incoming proposals. | 50 |
| AgreementIncomingBundlesQueueLength | AgreementIncomingBundlesQueueLength sets the size of the buffer holding incoming bundles. | 15 |
| MaxAcctLookback | MaxAcctLookback sets the maximum lookback range for account states,<br>i.e. the ledger can answer account states questions for the range Latest-MaxAcctLookback...Latest | 4 |
| MaxBlockHistoryLookback | BlockHistoryLookback sets the max lookback range for block information.<br>i.e. the block DB can return transaction IDs for questions for the range Latest-MaxBlockHistoryLookback...Latest | 0 |
| EnableUsageLog | EnableUsageLog enables 10Hz log of CPU and RAM usage.<br>Also adds 'algod_ram_usage` (number of bytes in use) to /metrics | false |
| MaxAPIBoxPerApplication | MaxAPIBoxPerApplication defines the maximum total number of boxes per application that will be returned<br>in GetApplicationBoxes REST API responses. | 100000 |
| TxIncomingFilteringFlags | TxIncomingFilteringFlags instructs algod filtering incoming tx messages<br>Flag values:<br>0x00 - disabled<br>0x01 (txFilterRawMsg) - check for raw tx message duplicates<br>0x02 (txFilterCanonical) - check for canonical tx group duplicates | 1 |
| EnableExperimentalAPI | EnableExperimentalAPI enables experimental API endpoint. Note that these endpoints have no<br>guarantees in terms of functionality or future support. | false |
| DisableLedgerLRUCache | DisableLedgerLRUCache disables LRU caches in ledger.<br>Setting it to TRUE might result in significant performance degradation<br>and SHOULD NOT be used for other reasons than testing. | false |
| EnableFollowMode | EnableFollowMode launches the node in "follower" mode. This turns off the agreement service,<br>and APIs related to broadcasting transactions, and enables APIs which can retrieve detailed information<br>from ledger caches and can control the ledger round. | false |
| EnableTxnEvalTracer | EnableTxnEvalTracer turns on features in the BlockEvaluator which collect data on transactions, exposing them via algod APIs.<br>It will store txn deltas created during block evaluation, potentially consuming much larger amounts of memory, | false |
| StorageEngine | StorageEngine allows to control which type of storage to use for the ledger.<br>Available options are:<br>- sqlite (default)<br>- pebbledb (experimental, in development) | sqlite |
| TxIncomingFilterMaxSize | TxIncomingFilterMaxSize sets the maximum size for the de-duplication cache used by the incoming tx filter<br>only relevant if TxIncomingFilteringFlags is non-zero | 500000 |
| BlockServiceMemCap | BlockServiceMemCap is the memory capacity in bytes which is allowed for the block service to use for HTTP block requests.<br>When it exceeds this capacity, it redirects the block requests to a different node | 500000000 |
| EnableP2P | EnableP2P turns on the peer to peer network.<br>When both EnableP2P and EnableP2PHybridMode (below) are set, EnableP2PHybridMode takes precedence. | false |
| EnableP2PHybridMode | EnableP2PHybridMode turns on both websockets and P2P networking.<br>Enabling this setting also requires PublicAddress to be set. | false |
| P2PHybridNetAddress | P2PHybridNetAddress sets the listen address used for P2P networking, if hybrid mode is set. |  |
| EnableDHTProviders | EnableDHT will turn on the hash table for use with capabilities advertisement | false |
| P2PPersistPeerID | P2PPersistPeerID will write the private key used for the node's PeerID to the P2PPrivateKeyLocation.<br>This is only used when P2PEnable is true. If P2PPrivateKey is not specified, it uses the default location. | false |
| P2PPrivateKeyLocation | P2PPrivateKeyLocation allows the user to specify a custom path to the private key used for the node's PeerID.<br>The private key provided must be an ed25519 private key.<br>This is only used when P2PEnable is true. If the parameter is not set, it uses the default location. |  |
| DisableAPIAuth | DisableAPIAuth turns off authentication for public (non-admin) API endpoints. | false |
| GoMemLimit | GoMemLimit provides the Go runtime with a soft memory limit. The default behavior is no limit,<br>unless the GOMEMLIMIT environment variable is set. | 0 |




# kmd Configuration Settings
The `kmd` process configuration parameters are shown in the table below.

| Property| Description | Default Value |
|------|------|------|
| address | Configures the address the node listens to for REST API calls. Specify an IP and port or just port. For example, 127.0.0.1:0 will listen on a random port on the localhost | 127.0.0.1:0 |
| allowed_origins | Configures the whitelist for allowed domains which can access the kmd process. Specify an array of urls that will be white listed. ie {“allowed_origins”: [“https://othersite1.com“, “https://othersite2.com”]} | |
| session_lifetime_secs | Number of seconds for session expirations.| 60 |
