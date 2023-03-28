title: conduit
---
## conduit



Run the Conduit framework.



### Synopsis



Conduit is a framework for ingesting blocks from the Algorand blockchain

into external applications. It is designed as a modular plugin system that

allows users to configure their own data pipelines.



You must provide a data directory containing a file named conduit.yml. The

file configures pipeline and all enabled plugins.



See other subcommands for further built in utilities and information.



Detailed documentation is online: https://github.com/algorand/conduit



```

conduit [flags]

```



### Options



```

  -d, --data-dir string            Set the Conduit data directory. If not set the CONDUIT_DATA_DIR environment variable is used.

  -h, --help                       help for conduit

  -r, --next-round-override uint   Set the starting round. Overrides next-round in metadata.json. Some exporters do not support overriding the starting round.

  -v, --version                    Print the Conduit version.

```



### SEE ALSO



* [conduit init](../init/)	 - initializes a Conduit data directory
* [conduit list](../list/list/)	 - List all available Conduit plugins



