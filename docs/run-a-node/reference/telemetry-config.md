title: Telemetry config settings

#Initialization

When a node is run using the algod command, before the script starts the server, it configures its telemetry based on the appropriate logging.config file. The algoh command, which hosts algod, configures logging and telemetry before calling algod. Both of these commands can override the config file’s telemetry enable field’s value using the -t flag.
When a node’s telemetry is enabled, a telemetry state (which wraps the node’s hook for the elasticsearch server to which the logs are saved) is added to the node’s logger reflecting the fields contained within the appropriate config file.

#Configuration

A node’s telemetry status can be managed using the diagcfg CLI, which modifies the node’s logging.config file. This file instructs the node’s construction of its TelemetryConfig struct, defining the following fields:

| Key | Data type | Description |
| --- | --------- | ----------- |
| Enable | bool | Determines whether Algorand remote logging is enabled for this node. |
| SendToLog | bool | Determines whether telemetry entries should be logged locally as well. |
| URI | string | The URI for the elastic search server to be logged to. Leave blank for default. |
| Name | string | The name of the machine for remote logging purposes. |
| GUID | string | A unique identifier for the node’s telemetry logging. Except in contrived circumstances, there should exist one GUID across all nodes running on a machine. |
| MinLogLevel | logrus.LogLevel | The lowest event significance that should be logged. |
| ReportHistoryLevel | logrus.LogLevel | The logrus importance level at which the node’s history will be reported. Must be greater than or equal to MinLogLevel. |
| FilePath | string | The location to which the logging.config file instance of the struct will be saved. |
| UserName | string | The username credential for establishing an elastic telemetry hook. |
| Password | string | The password credential for establishing an elastic telemetry hook. |

An Algorand node host can configure their node’s telemetry before running it by modifying the logging.config file in their node’s data directory, or deleting this file and modifying their ~/.algorand/logging.config file. In addition, the user can alter a running node’s telemetry status using the diagcfg CLI

#Config File Location

The file named ‘logging.config’ informs the initial configuration of a node’s telemetry. There will typically be at least two logging.config files on a machine running a node: one for each node the machine is running, stored in that node’s data directory, and a global config file stored in ~/.algorand/. This global file is generally only accessed when a node-specific config file cannot be found.

However, the “diagcfg telemetry” command tree, which replaces the functionality of “goal logging”, updates or creates both the local and global config files when executing any command that changes the node’s telemetry state, only failing to create the local file if no dataDir is provided, in which case there’s presumably also no node running.