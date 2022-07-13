title: goal account addpartkey
---
## goal account addpartkey



Generate and install participation key for the specified account



### Synopsis



Generate and install participation key for the specified account. This participation key can then be used for going online and participating in consensus.



```

goal account addpartkey [flags]

```



### Options



```

  -a, --address string         Account to associate with the generated partkey

  -h, --help                   help for addpartkey

      --keyDilution uint       Key dilution for two-level participation keys (defaults to sqrt of validity window)

  -o, --outdir string          Save participation key file to specified output directory to (for offline creation)

      --roundFirstValid uint   The first round for which the generated partkey will be valid

      --roundLastValid uint    The last round for which the generated partkey will be valid

```



### Options inherited from parent commands



```

  -d, --datadir stringArray   Data directory for the node

  -k, --kmddir string         Data directory for kmd

  -w, --wallet string         Set the wallet to be used for the selected operation

```



### SEE ALSO



* [goal account](../../account/account/)	 - Control and manage Algorand accounts



