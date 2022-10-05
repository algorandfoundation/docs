title: goal app interact execute
---
## goal app interact execute



Execute a procedure on an application



### Synopsis



Execute a procedure on an application



```

goal app interact execute [flags]

```



### Options



```

      --app-id uint     Application ID (if omitted, zero, which creates an application)

  -f, --from string     Account to execute interaction from

  -h, --help            help for execute

  -S, --signer string   Address of key to sign with, if different from "from" address due to rekeying

```



### Options inherited from parent commands



```

      --app-account strings        Accounts that may be accessed from application logic

  -i, --app-input string           JSON file containing encoded arguments and inputs (mutually exclusive with app-arg-b64 and app-account)

      --approval-prog string       (Uncompiled) TEAL assembly program filename for approving/rejecting transactions

      --approval-prog-raw string   Compiled TEAL program filename for approving/rejecting transactions

      --clear-prog string          (Uncompiled) TEAL assembly program filename for updating application state when a user clears their local state

      --clear-prog-raw string      Compiled TEAL program filename for updating application state when a user clears their local state

  -d, --datadir stringArray        Data directory for the node

      --foreign-app strings        Indexes of other apps whose global state is read in this transaction

      --foreign-asset strings      Indexes of assets whose parameters are read in this transaction

      --header string              Application header

  -k, --kmddir string              Data directory for kmd

  -w, --wallet string              Set the wallet to be used for the selected operation

```



### SEE ALSO



* [goal app interact](../../interact/interact/)	 - Interact with an application



