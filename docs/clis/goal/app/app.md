title: goal app
---
## goal app



Manage applications



```

goal app [flags]

```



### Options



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

      --foreign-app strings        Indexes of other apps whose global state is read in this transaction

      --foreign-asset strings      Indexes of assets whose parameters are read in this transaction

  -h, --help                       help for app

      --holding strings            A Holding that may be accessed from application logic. An asset-id followed by a comma and an address

      --local strings              A Local State that may be accessed from application logic. An optional app-id and comma, followed by an address. Zero or omitted app-id indicates the local state for app being called.

      --reject-version uint        If set non-zero, reject for this app version or higher

  -w, --wallet string              Set the wallet to be used for the selected operation

```



### Options inherited from parent commands



```

  -d, --datadir stringArray   Data directory for the node

  -k, --kmddir string         Data directory for kmd

```



### SEE ALSO



* [goal](../../../goal/goal/)	 - CLI for interacting with Algorand
* [goal app box](../box/box/)	 - Read application box data
* [goal app call](../call/)	 - Call an application
* [goal app clear](../clear/)	 - Clear out an application's state in your account
* [goal app closeout](../closeout/)	 - Close out of an application
* [goal app create](../create/)	 - Create an application
* [goal app delete](../delete/)	 - Delete an application
* [goal app info](../info/)	 - Look up current parameters for an application
* [goal app interact](../interact/interact/)	 - Interact with an application
* [goal app method](../method/)	 - Invoke an ABI method
* [goal app optin](../optin/)	 - Opt in to an application
* [goal app read](../read/)	 - Read local or global state for an application
* [goal app update](../update/)	 - Update an application's programs



