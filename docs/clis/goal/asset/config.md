title: goal asset config
---
## goal asset config



Configure an asset



### Synopsis



Change an asset configuration. This transaction must be issued by the asset manager. This allows any management address to be changed: manager, freezer, reserve, or clawback.



```

goal asset config [flags]

```



### Options



```

      --asset string                Unit name of asset to configure

      --assetid uint                Asset ID to configure

      --creator string              Account address for asset to configure (defaults to manager)

      --dryrun-accounts strings     additional accounts to include into dryrun request obj

      --dryrun-dump                 Dump in dryrun format acceptable by dryrun REST api

      --dryrun-dump-format string   Dryrun dump format: json, msgp (default "json")

      --fee uint                    The transaction fee (automatically determined by default), in microAlgos

      --firstvalid uint             The first round where the transaction may be committed to the ledger

  -h, --help                        help for config

      --lastvalid uint              The last round where the transaction may be committed to the ledger

  -x, --lease string                Lease value (base64, optional): no transaction may also acquire this lease until lastvalid

      --manager string              Manager account to issue the config transaction

      --new-clawback string         New clawback address

      --new-freezer string          New freeze address

      --new-manager string          New manager address

      --new-reserve string          New reserve address

  -N, --no-wait                     Don't wait for transaction to commit

  -n, --note string                 Note text (ignored if --noteb64 used also)

      --noteb64 string              Note (URL-base64 encoded)

  -o, --out string                  Write transaction to this file

  -s, --sign                        Use with -o to indicate that the dumped transaction should be signed

  -S, --signer string               Address of key to sign with, if different from transaction "from" address due to rekeying

      --validrounds uint            The number of rounds for which the transaction will be valid

```



### Options inherited from parent commands



```

  -d, --datadir stringArray   Data directory for the node

  -k, --kmddir string         Data directory for kmd

  -w, --wallet string         Set the wallet to be used for the selected operation

```



### SEE ALSO



* [goal asset](../../asset/asset/)	 - Manage assets



