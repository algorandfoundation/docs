title: goal asset send
---
## goal asset send



Transfer assets



### Synopsis



Transfer asset holdings. An account can begin accepting an asset by issuing a zero-amount asset transfer to itself.



```

goal asset send [flags]

```



### Options



```

  -a, --amount uint        The amount to be transferred (required), in base units of the asset.

      --asset string       Unit name of the asset being transferred

      --assetid uint       ID of the asset being transferred

      --clawback string    Address to issue a clawback transaction from (defaults to no clawback)

  -c, --close-to string    Close asset account and send remainder to this address

      --creator string     Account address for asset creator

      --fee uint           The transaction fee (automatically determined by default), in microAlgos

      --firstvalid uint    The first round where the transaction may be committed to the ledger

  -f, --from string        Account address to send the money from (if not specified, uses default account)

  -h, --help               help for send

      --lastvalid uint     The last round where the transaction may be committed to the ledger

  -x, --lease string       Lease value (base64, optional): no transaction may also acquire this lease until lastvalid

  -N, --no-wait            Don't wait for transaction to commit

  -n, --note string        Note text (ignored if --noteb64 used also)

      --noteb64 string     Note (URL-base64 encoded)

  -o, --out string         Write transaction to this file

  -s, --sign               Use with -o to indicate that the dumped transaction should be signed

  -t, --to string          Address to send to money to (required)

      --validrounds uint   The number of rounds for which the transaction will be valid

```



### Options inherited from parent commands



```

  -d, --datadir stringArray   Data directory for the node

  -k, --kmddir string         Data directory for kmd

  -w, --wallet string         Set the wallet to be used for the selected operation

```



### SEE ALSO



* [goal asset](../../asset/asset/)	 - Manage assets



