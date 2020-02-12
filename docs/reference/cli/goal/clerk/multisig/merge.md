title: goal clerk multisig merge
---
## goal clerk multisig merge



Merge multisig signatures on transactions



### Synopsis



Combine multiple partially-signed multisig transactions, and write out transactions with a single merged multisig signature.



```

goal clerk multisig merge -o MERGEDTXFILE TXFILE1 TXFILE2 ... [flags]

```



### Options



```

  -h, --help         help for merge

  -o, --out string   Output file for merged transactions

```



### Options inherited from parent commands



```

  -d, --datadir stringArray   Data directory for the node

  -k, --kmddir string         Data directory for kmd

  -w, --wallet string         Set the wallet to be used for the selected operation

```



### SEE ALSO



* [goal clerk multisig](../../multisig/multisig/)	 - Provides tools working with multisig transactions 



