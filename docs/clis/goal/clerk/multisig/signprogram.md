title: goal clerk multisig signprogram
---
## goal clerk multisig signprogram



Add a signature to a multisig LogicSig



### Synopsis



Start a multisig LogicSig, or add a signature to an existing multisig, for a given program.



```

goal clerk multisig signprogram -a [address] [flags]

```



### Options



```

  -a, --address string         Address of the key to sign with

  -h, --help                   help for signprogram

      --legacy-msig            Use legacy multisig (if not specified, auto-detect consensus params from algod)

  -L, --lsig string            Partial LogicSig to add signature to

  -o, --lsig-out string        File to write partial Lsig to

  -A, --msig-address string    Multi-Sig Address that signing address is part of

  -p, --program string         Program source to be compiled and signed

  -P, --program-bytes string   Program binary to be signed

```



### Options inherited from parent commands



```

  -d, --datadir stringArray   Data directory for the node

  -k, --kmddir string         Data directory for kmd

  -w, --wallet string         Set the wallet to be used for the selected operation

```



### SEE ALSO



* [goal clerk multisig](../../multisig/multisig/)	 - Provides tools working with multisig transactions 



