title: diagcfg telemetry
---
## diagcfg telemetry



Enable/disable and configure Algorand remote logging



### Synopsis


With no commands or flags, displays status of node's telemetry reporting. Commands allow configuration of a node's remote logging.



```
diagcfg telemetry [command]
diagcfg telemetry [flags]
```



### Options

## Commands

```
  disable                Disable remote logging for a node
  enable                 Enable remote logging for a node
  endpoint -e <url>      Sets the URI property in the telemetry configuration
  
  name -n <nodeName>     Enable remote logging with a specified node name
  
  status                 Print the node's remote logging status. Same functionality as the bare 'diagcfg telemetry' command
```

### Options inherited from parent commands



```
  -d, --dataDir string        Data directory for the node
  -h, --help                  Help for telemetry commands
```



### SEE ALSO


* [diagcfg](../../diagcfg/diagcfg/)	 - CLI for interacting with Algorand logging.

