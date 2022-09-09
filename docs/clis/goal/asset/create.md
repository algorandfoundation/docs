title: goal asset create
---
## goal asset create



Create an asset



### Synopsis



Post a transaction declaring and issuing a new layer-one asset on the network.



```

goal asset create [flags]

```



### Options



```

      --assetmetadatab64 string     base-64 encoded 32-byte commitment to asset metadata

      --asseturl string             URL where user can access more information about the asset (max 32 bytes)

      --clawback string             Clawback account that is allowed to transfer assets from and to any asset holder

      --creator string              Account address for creating an asset

      --decimals uint32             The number of digits to use after the decimal point when displaying this asset. If set to 0, the asset is not divisible beyond its base unit. If set to 1, the base asset unit is tenths. If 2, the base asset unit is hundredths, and so on.

      --defaultfrozen               Freeze or not freeze holdings by default

      --dryrun-accounts strings     additional accounts to include into dryrun request obj

      --dryrun-dump                 Dump in dryrun format acceptable by dryrun REST api

      --dryrun-dump-format string   Dryrun dump format: json, msgp (default "json")

      --fee uint                    The transaction fee (automatically determined by default), in microAlgos

      --firstvalid uint             The first round where the transaction may be committed to the ledger

      --freezer string              Freezer account that can freeze or unfreeze the asset holdings for a specific account

  -h, --help                        help for create

      --lastvalid uint              The last round where the transaction may be committed to the ledger

  -x, --lease string                Lease value (base64, optional): no transaction may also acquire this lease until lastvalid

      --manager string              Manager account that can issue transactions to re-configure or destroy the asset

      --name string                 Name for the entire asset

      --no-clawback                 Explicitly declare the lack of clawback

      --no-freezer                  Explicitly declare the lack of freezer

      --no-manager                  Explicitly declare the lack of manager

      --no-reserve                  Explicitly declare the lack of reserve

  -N, --no-wait                     Don't wait for transaction to commit

  -n, --note string                 Note text (ignored if --noteb64 used also)

      --noteb64 string              Note (URL-base64 encoded)

  -o, --out string                  Write transaction to this file

      --reserve string              Reserve account that non-minted assets will reside in

  -s, --sign                        Use with -o to indicate that the dumped transaction should be signed

  -S, --signer string               Address of key to sign with, if different from transaction "from" address due to rekeying

      --total uint                  Total amount of tokens for created asset

      --unitname string             Name for the unit of asset

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



