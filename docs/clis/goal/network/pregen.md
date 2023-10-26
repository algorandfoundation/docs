title: goal network pregen
---
## goal network pregen



Pregenerate private network



### Synopsis



Pregenerates the root and participation keys for a private network. The pregen directory can then be passed to the 'goal network create' to start the network more quickly.



```

goal network pregen [flags]

```



### Options



```

  -h, --help               help for pregen

  -p, --pregendir string   Specify the path to the directory to export genesis.json, root and partkey files. This should only be used on private networks.

  -t, --template string    Specify the path to the template file for the network

```



### Options inherited from parent commands



```

  -d, --datadir stringArray   Data directory for the node

  -k, --kmddir string         Data directory for kmd

  -r, --rootdir string        Root directory for the private network directories

```



### SEE ALSO



* [goal network](../../network/network/)	 - Create and manage private, multi-node, locally-hosted networks



