title: goal account renewpartkey
---
## goal account renewpartkey



Renew an account's participation key



### Synopsis



Generate a participation key for the specified account and issue the necessary transaction to register it.



```

goal account renewpartkey [flags]

```



### Options



```

  -a, --address string        Account address to update (required)

  -f, --fee uint              The Fee to set on the status change transaction (defaults to suggested fee)

  -h, --help                  help for renewpartkey

      --keyDilution uint      Key dilution for two-level participation keys

  -N, --no-wait               Don't wait for transaction to commit

      --roundLastValid uint   The last round for which the generated partkey will be valid

```



### Options inherited from parent commands



```

  -d, --datadir stringArray   Data directory for the node

  -k, --kmddir string         Data directory for kmd

  -w, --wallet string         Set the wallet to be used for the selected operation

```



### SEE ALSO



* [goal account](../../account/account/)	 - Control and manage Algorand accounts



