title: goal app clear
---
## goal app clear



Clear out an application's state in your account



### Synopsis



Remove any local state from your account associated with an application. The application does not need to exist anymore.



```

goal app clear [flags]

```



### Options



```

      --app-id uint                 Application ID

      --dryrun-accounts strings     additional accounts to include into dryrun request obj

      --dryrun-dump                 Dump in dryrun format acceptable by dryrun REST api

      --dryrun-dump-format string   Dryrun dump format: json, msgp (default "json")

      --fee uint                    The transaction fee (automatically determined by default), in microAlgos

      --firstvalid uint             The first round where the transaction may be committed to the ledger

  -f, --from string                 Account to clear app state for

  -h, --help                        help for clear

      --lastvalid uint              The last round where the transaction may be committed to the ledger

  -x, --lease string                Lease value (base64, optional): no transaction may also acquire this lease until lastvalid

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

      --access                     Put references into the transaction's access list, instead of foreign arrays.

      --app-account strings        Accounts that may be accessed from application logic

      --app-arg stringArray        Args to encode for application transactions (all will be encoded to a byte slice). For ints, use the form 'int:1234'. For raw bytes, use the form 'b64:A=='. For printable strings, use the form 'str:hello'. For addresses, use the form 'addr:XYZ...'.

  -i, --app-input string           JSON file containing encoded arguments and inputs (mutually exclusive with app-arg, app-account, foreign-app, foreign-asset, local, holding, and box)

      --approval-prog string       (Uncompiled) TEAL assembly program filename for approving/rejecting transactions

      --approval-prog-raw string   Compiled TEAL program filename for approving/rejecting transactions

      --box stringArray            A Box that may be accessed by this transaction. Use the same form as app-arg to name the box, preceded by an optional app-id and comma. Zero or omitted app-id indicates the box is accessible by the app being called.

      --clear-prog string          (Uncompiled) TEAL assembly program filename for updating application state when a user clears their local state

      --clear-prog-raw string      Compiled TEAL program filename for updating application state when a user clears their local state

  -d, --datadir stringArray        Data directory for the node

      --foreign-app strings        Indexes of other apps whose global state is read in this transaction

      --foreign-asset strings      Indexes of assets whose parameters are read in this transaction

      --holding strings            A Holding that may be accessed from application logic. An asset-id followed by a comma and an address

  -k, --kmddir string              Data directory for kmd

      --local strings              A Local State that may be accessed from application logic. An optional app-id and comma, followed by an address. Zero or omitted app-id indicates the local state for app being called.

      --reject-version uint        If set non-zero, reject for this app version or higher

  -w, --wallet string              Set the wallet to be used for the selected operation

```



### SEE ALSO



* [goal app](../../app/app/)	 - Manage applications



