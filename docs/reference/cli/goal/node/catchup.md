title: goal node catchup
---
## goal node catchup



Catchup the Algorand node to a specific catchpoint.



### Synopsis



Catchup allows making large jumps over round ranges without the need to incremently validate each individual round.



```

goal node catchup [flags]

```


### Usage

Once you aquired a catchpoint from a secured source, you can instruct your node to catchup to that catchpoint. To do so, type the following:

```

goal node catchup 6500000#1234567890ABCDEF01234567890ABCDEF0 -d <data directory>

```

The above would start catching up to the given catchpoint. When the catchpoint catchup would complete, the node would be at round 6500000, and the account data would match the hash of 1234567890ABCDEF01234567890ABCDEF0.

|   |
|---|
| :information_source: In this example, the provided hash was invalid : hashes are base32 encoded, and the hash above contains the numbers 0,1,8,9 which aren't part of base32. |

If you decided to abort the catchup before it completes ( for instance, you maybe want to apply a different catchpoint ) type the following:


```

goal node catchup --abort -d <data directory>

```

|   |
|---|
| :information_source: The [goal node status](../../node/status/) command could help you monitor the catchpoint catchup progress. Use the *-w* option of the command to continuesly monitor the node's status |


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



