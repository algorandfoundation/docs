title: goal clerk dryrun-remote
---
## goal clerk dryrun-remote



Test a program with algod's dryrun REST endpoint



### Synopsis

!!! warning
    As of AVMv8, `dryrun` will no longer work with any contract that uses box storage. A new endpoint that will replace `dryrun` is current in development.

Test a TEAL program with algod's dryrun REST endpoint under various conditions and verbosity.



```

goal clerk dryrun-remote [flags]

```



### Options



```

  -D, --dryrun-state string   dryrun request object to run

  -h, --help                  help for dryrun-remote

  -r, --raw                   output raw response from algod

  -v, --verbose               print more info

```



### Options inherited from parent commands



```

  -d, --datadir stringArray   Data directory for the node

  -k, --kmddir string         Data directory for kmd

  -w, --wallet string         Set the wallet to be used for the selected operation

```



### SEE ALSO



* [goal clerk](../../clerk/clerk/)	 - Provides the tools to control transactions 



