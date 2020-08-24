title: goal node catchup
---
## goal node catchup



Catchup the Algorand node to a specific catchpoint



### Synopsis



Catchup allows making large jumps over round ranges without the need to incrementally validate each individual round.



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

  -x, --abort   Aborts the current catchup process

  -h, --help    help for catchup

```



### Options inherited from parent commands



```

  -d, --datadir stringArray   Data directory for the node

  -k, --kmddir string         Data directory for kmd

```



### SEE ALSO



* [goal node](../../node/node/)	 - Manage a specified algorand node



