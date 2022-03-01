title: goal node create
---
## goal node create



Create a node at the desired data directory for the desired network



### Synopsis



Create a node at the desired data directory for the desired network



```

goal node create [flags]

```



### Options



```

      --api string           REST API Endpoint

  -a, --archival             Make the new node archival, storing all blocks

      --destination string   Destination path for the new node

      --full-config          Store full config file

  -h, --help                 help for create

  -H, --hosted               Configure the new node to run hosted by algoh

  -i, --indexer              Configure the new node to enable the indexer feature (implies --archival)

      --network string       Network the new node should point to

      --relay string         Configure as a relay with specified listening address (NetAddress)

```



### Options inherited from parent commands



```

  -d, --datadir stringArray   Data directory for the node

  -k, --kmddir string         Data directory for kmd

```



### SEE ALSO



* [goal node](../../node/node/)	 - Manage a specified algorand node



