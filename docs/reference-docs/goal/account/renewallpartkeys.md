title: goal account renewallpartkeys
---
## goal account renewallpartkeys



Renew all existing participation keys



### Synopsis



Generate new participation keys for all existing accounts with participation keys and issue the necessary transactions to register them.



```

goal account renewallpartkeys [flags]

```



### Options



```

  -f, --fee uint              The Fee to set on the status change transactions (defaults to suggested fee)

  -h, --help                  help for renewallpartkeys

      --keyDilution uint      Key dilution for two-level participation keys

  -N, --no-wait               Don't wait for transaction to commit

      --roundLastValid uint   The last round for which the generated partkeys will be valid

```



### Options inherited from parent commands



```

  -d, --datadir stringArray   Data directory for the node

  -k, --kmddir string         Data directory for kmd

  -w, --wallet string         Set the wallet to be used for the selected operation

```



### SEE ALSO



* [goal account](../../account/account/)	 - Control and manage Algorand accounts



