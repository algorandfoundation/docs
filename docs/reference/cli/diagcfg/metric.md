title: diagcfg metric
---
## diagcfg metric



Enable/disable and configure Algorand metric collection, as indicated in the config.json file



### Synopsis



Collection of commands configure metric collection of an Algorand node instance.
Without any commands or flags, calls 'metric status'.


```

diagcfg metric [command] [flags]

```



### Options

## Commands

```

  disable                Disable metric collection on node, setting the the EnableMetricReporting field in config.json to true

  enable                 Enable metric collection on node, setting the the EnableMetricReporting field in config.json to false

  status                 Print the node's metric status, as indicated by teh EnableMetricReporting field in config.json.

```


### Options inherited from parent commands



```

  -d, --dataDir string        Data directory for the node


```



### SEE ALSO



* [diagcfg](../../diagcfg/diagcfg/)	 - CLI for interacting with Algorand logging.

