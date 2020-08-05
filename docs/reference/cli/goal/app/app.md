title: goal app
---
## goal app



Manage applications



### Synopsis



Manage applications



```

goal app [flags]

```



### Options



```

      --app-account strings        Accounts that may be accessed from application logic

      --app-arg strings            Args to encode for application transactions (all will be encoded to a byte slice). For ints, use the form 'int:1234'. For raw bytes, use the form 'b64:A=='. For printable strings, use the form 'str:hello'. For addresses, use the form 'addr:XYZ...'.

  -i, --app-input string           JSON file containing encoded arguments and inputs (mutually exclusive with app-arg-b64 and app-account)

      --approval-prog string       (Uncompiled) TEAL assembly program filename for approving/rejecting transactions

      --approval-prog-raw string   Compiled TEAL program filename for approving/rejecting transactions

      --clear-prog string          (Uncompiled) TEAL assembly program filename for updating application state when a user clears their local state

      --clear-prog-raw string      Compiled TEAL program filename for updating application state when a user clears their local state

      --foreign-app strings        Indexes of other apps whose global state is read in this transaction

  -h, --help                       help for app

  -w, --wallet string              Set the wallet to be used for the selected operation

```



### Options inherited from parent commands



```

  -d, --datadir stringArray   Data directory for the node

  -k, --kmddir string         Data directory for kmd

```



### SEE ALSO



* [goal](../../../goal/goal/)	 - CLI for interacting with Algorand
* [goal](../../../goal/goal/)	 - CLI for interacting with Algorand.
* [goal app call](../call/)	 - Call an application
* [goal app clear](../clear/)	 - Clear out an application's state in your account
* [goal app closeout](../closeout/)	 - Close out of an application
* [goal app create](../create/)	 - Create an application
* [goal app delete](../delete/)	 - Delete an application
* [goal app info](../info/)	 - Look up current parameters for an application
* [goal app interact](../interact/interact/)	 - Interact with an application
* [goal app optin](../optin/)	 - Opt in to an application
* [goal app read](../read/)	 - Read local or global state for an application
* [goal app update](../update/)	 - Update an application's programs



