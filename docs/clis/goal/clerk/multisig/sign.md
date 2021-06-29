title: goal clerk multisig sign
---
## goal clerk multisig sign



Add a signature to a multisig transaction



### Synopsis



Start a multisig, or add a signature to an existing multisig, for a given transaction.



```

goal clerk multisig sign -t [transaction file] -a [address] [flags]

```



### Options



```

  -a, --address string   Address of the key to sign with

  -h, --help             help for sign

  -n, --no-sig           Fill in the transaction's multisig field with public keys and threshold information, but don't produce a signature

  -t, --tx string        Partially-signed transaction file to add signature to

```



### Options inherited from parent commands



```

  -d, --datadir stringArray   Data directory for the node

  -k, --kmddir string         Data directory for kmd

  -w, --wallet string         Set the wallet to be used for the selected operation

```



### SEE ALSO



* [goal clerk multisig](../../multisig/multisig/)	 - Provides tools working with multisig transactions 



