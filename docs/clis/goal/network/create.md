title: goal network create
---
## goal network create



Create a private named network from a template



### Synopsis



Creates a collection of folders under the specified root directory that make up the entire private network named 'private' (simplifying cleanup).



```

goal network create [flags]

```



### Options



```

      --devMode            Forces the configuration to enable DevMode, returns an error if the template is not compatible with DevMode.

  -h, --help               help for create

  -n, --network string     Specify the name to use for the private network

      --noclean            Prevents auto-cleanup on error - for diagnosing problems

  -K, --noimportkeys       Do not import root keys when creating the network (by default will import)

  -p, --pregendir string   Specify the path to the directory with pregenerated genesis.json, root and partkeys to import into the network directory. By default, the genesis.json and keys will be generated on start. This should only be used on private networks.

  -s, --start              Automatically start the network after creating it.

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



