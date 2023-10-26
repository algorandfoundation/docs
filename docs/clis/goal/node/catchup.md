title: goal node catchup
---
## goal node catchup



Catchup the Algorand node to a specific catchpoint



### Synopsis



Catchup allows making large jumps over round ranges without the need to incrementally validate each individual round. Using external catchpoints is not a secure practice and should not be done for consensus participating nodes.

If no catchpoint is provided, this command attempts to lookup the latest catchpoint from algorand-catchpoints.s3.us-east-2.amazonaws.com.



```

goal node catchup [flags]

```



### Examples



```

goal node catchup 6500000#1234567890ABCDEF01234567890ABCDEF0	Start catching up to round 6500000 with the provided catchpoint

goal node catchup --abort					Abort the current catchup

```



### Options



```

  -x, --abort      Aborts the current catchup process

      --force      Forces fast catchup with implicit catchpoint to start without a consent prompt

  -h, --help       help for catchup

  -m, --min uint   Catchup only if the catchpoint would advance the node by the specified minimum number of rounds

```



### Options inherited from parent commands



```

  -d, --datadir stringArray   Data directory for the node

  -k, --kmddir string         Data directory for kmd

```



### SEE ALSO



* [goal node](../../node/node/)	 - Manage a specified algorand node



