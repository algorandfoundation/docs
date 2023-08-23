title: algokey multisig append-auth-addr
---
## algokey multisig append-auth-addr



Adds the necessary fields to a transaction that is sent from an account that was rekeyed to a multisig account



```

algokey multisig append-auth-addr -t [transaction file] -p "[threshold] [Address 1] [Address 2] ..." [flags]

```



### Options



```

  -h, --help             help for append-auth-addr

  -o, --outfile string   Transaction output filename. If not specified, the original file will be modified

  -p, --params string    Multisig pre image parameters - [threshold] [Address 1] [Address 2] ...

  -t, --txfile string    Transaction input filename

```



### SEE ALSO



* [algokey multisig](../../multisig/multisig/)	 - Add a multisig signature to transactions from a file using a private key



