title: conduit init
---
## conduit init



initializes a sample data directory



### Synopsis



initializes a Conduit data directory and conduit.yml file configured with the file_writer plugin. The config file needs to be modified slightly to include an algod address and token. Once ready, launch conduit with './conduit -d /path/to/data'.



```

conduit init [flags]

```



### Options



```

  -d, --data string   Full path to new data directory. If not set, a directory named 'data' will be created in the current directory.

  -h, --help          help for init

```



### SEE ALSO



* [conduit](../../conduit/conduit/)	 - run the conduit framework



