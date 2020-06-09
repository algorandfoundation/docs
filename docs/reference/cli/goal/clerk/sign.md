title: goal clerk sign
---
## goal clerk sign



Sign a transaction file



### Synopsis



Sign the passed transaction file, which may contain one or more transactions. If the infile and the outfile are the same, this overwrites the file with the new, signed data.



```

goal clerk sign -i INFILE -o OUTFILE [flags]

```



### Options



```

      --argb64 strings     base64 encoded args to pass to transaction logic

  -h, --help               help for sign

  -i, --infile string      Partially-signed transaction file to add signature to

  -L, --logic-sig string   LogicSig to apply to transaction

  -o, --outfile string     Filename for writing the signed transaction

  -p, --program string     Program source to use as account logic

  -P, --proto string       consensus protocol version id string

  -S, --signer string      Address of key to sign with, if different from transaction "from" address due to rekeying

```



### Options inherited from parent commands



```

  -d, --datadir stringArray   Data directory for the node

  -k, --kmddir string         Data directory for kmd

  -w, --wallet string         Set the wallet to be used for the selected operation

```



### SEE ALSO



* [goal clerk](../../clerk/clerk/)	 - Provides the tools to control transactions 



