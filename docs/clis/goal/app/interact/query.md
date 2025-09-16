title: goal app interact query
---
## goal app interact query



Query local or global state from an application



```

goal app interact query [flags]

```



### Options



```

      --app-id uint   Application ID

  -f, --from string   Account to query state for (if omitted, query from global state)

  -h, --help          help for query

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

      --header string              Application header

      --holding strings            A Holding that may be accessed from application logic. An asset-id followed by a comma and an address

  -k, --kmddir string              Data directory for kmd

      --local strings              A Local State that may be accessed from application logic. An optional app-id and comma, followed by an address. Zero or omitted app-id indicates the local state for app being called.

      --reject-version uint        If set non-zero, reject for this app version or higher

  -w, --wallet string              Set the wallet to be used for the selected operation

```



### SEE ALSO



* [goal app interact](../../interact/interact/)	 - Interact with an application



