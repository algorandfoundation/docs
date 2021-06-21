title: goal clerk send
---
## goal clerk send



Send money to an address



### Synopsis



Send money from one account to another. Note: by default, the money will be withdrawn from the default account. Creates a transaction sending amount tokens from fromAddr to toAddr. If the optional --fee is not provided, the transaction will use the recommended amount. If the optional --firstvalid and --lastvalid are provided, the transaction will only be valid from round firstValid to round lastValid. If broadcast of the transaction is successful, the transaction ID will be returned.



```

goal clerk send [flags]

```



### Options



```

  -a, --amount uint                 The amount to be transferred (required), in microAlgos

      --argb64 strings              base64 encoded args to pass to transaction logic

  -c, --close-to string             Close account and send remainder to this address

      --dryrun-dump                 Dump in dryrun format acceptable by dryrun REST api

      --dryrun-dump-format string   Dryrun dump format: json, msgp (default "json")

      --fee uint                    The transaction fee (automatically determined by default), in microAlgos

      --firstvalid uint             The first round where the transaction may be committed to the ledger

  -f, --from string                 Account address to send the money from (If not specified, uses default account)

  -F, --from-program string         Program source to use as account logic

  -P, --from-program-bytes string   Program binary to use as account logic

  -h, --help                        help for send

      --lastvalid uint              The last round where the transaction may be committed to the ledger

  -x, --lease string                Lease value (base64, optional): no transaction may also acquire this lease until lastvalid

  -L, --logic-sig string            LogicSig to apply to transaction

      --msig-params string          Multisig preimage parameters - [threshold] [Address 1] [Address 2] ...

                                    Used to add the necessary fields in case the account was rekeyed to a multisig account

  -N, --no-wait                     Don't wait for transaction to commit

  -n, --note string                 Note text (ignored if --noteb64 used also)

      --noteb64 string              Note (URL-base64 encoded)

  -o, --out string                  Write transaction to this file

      --rekey-to string             Rekey account to the given spending key/address. (Future transactions from this account will need to be signed with the new key.)

  -s, --sign                        Use with -o to indicate that the dumped transaction should be signed

  -t, --to string                   Address to send to money to (required)

      --validrounds uint            The number of rounds for which the transaction will be valid

```



### Options inherited from parent commands



```

  -d, --datadir stringArray   Data directory for the node

  -k, --kmddir string         Data directory for kmd

  -w, --wallet string         Set the wallet to be used for the selected operation

```



### SEE ALSO



* [goal clerk](../../clerk/clerk/)	 - Provides the tools to control transactions 



