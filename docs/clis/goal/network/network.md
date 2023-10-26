title: goal network
---
## goal network



Create and manage private, multi-node, locally-hosted networks



### Synopsis



Collection of commands to support the creation and management of 'private networks'. These are fully-formed Algorand networks with private, custom Genesis ledgers running the current build of Algorand software. Rather than creating a node instance based on the released genesis.json, these networks have their own and need to be manually connected.



The basic idea is that we create one or more data directories and wallets to form this network, specify which node owns which wallets, and can start/stop the network as a unit. Each node is just like any other node running on TestNet or DevNet.



```

goal network [flags]

```



### Options



```

  -h, --help             help for network

  -r, --rootdir string   Root directory for the private network directories

```



### Options inherited from parent commands



```

  -d, --datadir stringArray   Data directory for the node

  -k, --kmddir string         Data directory for kmd

```



### SEE ALSO



* [goal](../../../goal/goal/)	 - CLI for interacting with Algorand
* [goal network create](../create/)	 - Create a private named network from a template
* [goal network delete](../delete/)	 - Stops and Deletes a deployed private network
* [goal network pregen](../pregen/)	 - Pregenerate private network
* [goal network restart](../restart/)	 - Restart a deployed private network
* [goal network start](../start/)	 - Start a deployed private network
* [goal network status](../status/)	 - Prints status for all nodes in a deployed private network
* [goal network stop](../stop/)	 - Stop a deployed private network



