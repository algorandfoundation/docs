title: goal app method
---
## goal app method



Invoke an ABI method



### Synopsis



Invoke an ARC-4 ABI method on an App (stateful contract) with an application call transaction



```

goal app method [flags]

```



### Options



```

      --app-id uint                 Application ID

      --arg stringArray             Args to pass in for calling a method

      --create                      Create an application in this method call

      --dryrun-accounts strings     additional accounts to include into dryrun request obj

      --dryrun-dump                 Dump in dryrun format acceptable by dryrun REST api

      --dryrun-dump-format string   Dryrun dump format: json, msgp (default "json")

      --extra-pages uint32          Additional program space for supporting larger TEAL assembly program. A maximum of 3 extra pages is allowed. A page is 1024 bytes. Only valid when passed with --create.

      --fee uint                    The transaction fee (automatically determined by default), in microAlgos

      --firstvalid uint             The first round where the transaction may be committed to the ledger

  -f, --from string                 Account to call method from

      --global-byteslices uint      Maximum number of byte slices that may be stored in the global key/value store. Immutable, only valid when passed with --create.

      --global-ints uint            Maximum number of integer values that may be stored in the global key/value store. Immutable, only valid when passed with --create.

  -h, --help                        help for method

      --lastvalid uint              The last round where the transaction may be committed to the ledger

  -x, --lease string                Lease value (base64, optional): no transaction may also acquire this lease until lastvalid

      --local-byteslices uint       Maximum number of byte slices that may be stored in local (per-account) key/value stores for this app. Immutable, only valid when passed with --create.

      --local-ints uint             Maximum number of integer values that may be stored in local (per-account) key/value stores for this app. Immutable, only valid when passed with --create.

      --method string               Method to be called

  -N, --no-wait                     Don't wait for transaction to commit

  -n, --note string                 Note text (ignored if --noteb64 used also)

      --noteb64 string              Note (URL-base64 encoded)

      --on-completion string        OnCompletion action for application transaction (default "NoOp")

  -o, --out string                  Write transaction to this file

  -s, --sign                        Use with -o to indicate that the dumped transaction should be signed

  -S, --signer string               Address of key to sign with, if different from transaction "from" address due to rekeying

      --validrounds uint            The number of rounds for which the transaction will be valid

```



### Options inherited from parent commands



```

      --app-account strings        Accounts that may be accessed from application logic

      --app-arg stringArray        Args to encode for application transactions (all will be encoded to a byte slice). For ints, use the form 'int:1234'. For raw bytes, use the form 'b64:A=='. For printable strings, use the form 'str:hello'. For addresses, use the form 'addr:XYZ...'.

  -i, --app-input string           JSON file containing encoded arguments and inputs (mutually exclusive with app-arg, app-account, foreign-app, foreign-asset, and box)

      --approval-prog string       (Uncompiled) TEAL assembly program filename for approving/rejecting transactions

      --approval-prog-raw string   Compiled TEAL program filename for approving/rejecting transactions

      --box stringArray            Boxes that may be accessed by this transaction. Use the same form as app-arg to name the box, preceded by an optional app-id and comma. No app-id indicates the box is accessible by the app being called.

      --clear-prog string          (Uncompiled) TEAL assembly program filename for updating application state when a user clears their local state

      --clear-prog-raw string      Compiled TEAL program filename for updating application state when a user clears their local state

  -d, --datadir stringArray        Data directory for the node

      --foreign-app strings        Indexes of other apps whose global state is read in this transaction

      --foreign-asset strings      Indexes of assets whose parameters are read in this transaction

  -k, --kmddir string              Data directory for kmd

  -w, --wallet string              Set the wallet to be used for the selected operation

```



### SEE ALSO



* [goal app](../../app/app/)	 - Manage applications



