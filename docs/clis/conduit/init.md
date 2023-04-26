title: conduit init
---
## conduit init



initializes a Conduit data directory



### Synopsis



Initializes a conduit.yml file and writes it to stdout. By default

the config file uses an algod importer in follower mode and a block

file writer exporter. The plugin templates can be changed using the

different options.



Once initialized the conduit.yml file needs to be modified. Refer to the file

comments for details.



If the 'data' option is used, the file will be written to that

directory and additional help is written to stdout.



Once configured, launch conduit with './conduit -d /path/to/data'.



```

conduit init [flags]

```



### Examples



```

conduit init  -d /path/to/data -i importer -p processor1,processor2 -e exporter

```



### Options



```

  -d, --data string          Full path to new data directory to initialize. If not set, the conduit configuration YAML is written to stdout.

  -e, --exporter string      data exporter name.

  -h, --help                 help for init

  -i, --importer string      data importer name.

  -p, --processors strings   comma-separated list of processors.

```



### SEE ALSO



* [conduit](../../conduit/conduit/)	 - Run the Conduit framework.



