title: diagcfg metric
---
## diagcfg metric



Enable/diable and configure Algorand metric collection



### Synopsis



Collection of commands configure metric collection of an Algorand node instance.



```

diagcfg metric [command] [flags]

```



### Options

## Commands

```

  disable                Disable metric collection on node

  enable                 Enable metric collection on node

  status                 Print the node's metric status. Same functionality as the bare 'diagcfg metric' command

```

## Flags
```

  -e, --externalHostName string   External host name, such as relay-us-ea-3.algodev.network; will default to external IP Address if not specified

```



### Options inherited from parent commands



```

  -d, --dataDir string        Data directory for the node


```



### SEE ALSO



* [diagcfg](../../diagcfg/diagcfg/)	 - CLI for interacting with Algorand logging.



