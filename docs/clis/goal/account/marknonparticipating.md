title: goal account marknonparticipating
---
## goal account marknonparticipating



Permanently mark an account as not participating (i.e. offline and earns no rewards)



### Synopsis



Permanently mark an account as not participating (as opposed to Online or Offline). Once marked, the account can never go online or offline, it is forever nonparticipating, and it will never earn rewards on its balance.



```

goal account marknonparticipating [flags]

```



### Options



```

  -a, --address string     Account address to change

  -f, --fee uint           The Fee to set on the status change transaction (defaults to suggested fee)

      --firstvalid uint    FirstValid for the status change transaction (0 for current)

  -h, --help               help for marknonparticipating

      --lastvalid uint     The last round where the transaction may be committed to the ledger

  -N, --no-wait            Don't wait for transaction to commit

  -t, --txfile string      Write status change transaction to this file, rather than posting to network

  -v, --validrounds uint   The validity period for the status change transaction

```



### Options inherited from parent commands



```

  -d, --datadir stringArray   Data directory for the node

  -k, --kmddir string         Data directory for kmd

  -w, --wallet string         Set the wallet to be used for the selected operation

```



### SEE ALSO



* [goal account](../../account/account/)	 - Control and manage Algorand accounts



