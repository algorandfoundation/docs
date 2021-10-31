title: goal clerk tealsign
---
## goal clerk tealsign



Sign data to be verified in a TEAL program



### Synopsis



Sign data to be verified in a TEAL program.



Data verified by the ed25519verify TEAL opcode must be domain separated. As part of this process, the signed payload includes the hash of the program logic. This hash must be specified. To do this, provide a transaction whose logic sig contains the program via --lsig-txn, or provide a contract address directly with --contract-addr. These options are mutually exclusive.



Next, you must specify the data to be signed. When using --lsig-txn, you can use the --sign-txid flag to sign that transaction's txid. Alternatively, arbitrary data can be signed with the --data-file, --data-b64, or --data-b32 options. These options are mutually exclusive.



The base64 encoding of the signature will always be printed to stdout. Optionally, when using --lsig-txn, you may specify that the signature be used as a TEAL argument for that transaction. Specify the argument index with the --set-lsig-arg-idx flag. The --lsig-txn file will be updated in place, and any existing argument at that index will be overwritten.



```

goal clerk tealsign [flags]

```



### Options



```

      --account string         Address of account to sign with

      --contract-addr string   Contract address to sign data for. not necessary if --lsig-txn is provided

      --data-b32 string        base32 data to sign

      --data-b64 string        base64 data to sign

      --data-file string       Data file to sign

  -h, --help                   help for tealsign

      --keyfile string         algokey private key file to sign with

      --lsig-txn string        Transaction with logicsig to sign data for

      --set-lsig-arg-idx int   If --lsig-txn is also specified, set the lsig arg at this index to the raw signature bytes. Overwrites any existing argument at this index. Updates --lsig-txn file in place. nil args will be appended until index is valid. (default -1)

      --sign-txid              Use the txid of --lsig-txn as the data to sign

```



### Options inherited from parent commands



```

  -d, --datadir stringArray   Data directory for the node

  -k, --kmddir string         Data directory for kmd

  -w, --wallet string         Set the wallet to be used for the selected operation

```



### SEE ALSO



* [goal clerk](../../clerk/clerk/)	 - Provides the tools to control transactions 



