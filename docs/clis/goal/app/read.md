title: goal app read
---
## goal app read



Read local or global state for an application



### Synopsis



Read global or local (account-specific) state for an application



```

goal app read [flags]

```



### Options



```

      --app-id uint    Application ID

  -f, --from string    Account to fetch state from

      --global         Fetch global state for this application.

      --guess-format   Format application state using heuristics to guess data encoding.

  -h, --help           help for read

      --local --from   Fetch account-specific state for this application. --from address is required when using this flag

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

  -k, --kmddir string              Data directory for kmd

  -w, --wallet string              Set the wallet to be used for the selected operation

```



### SEE ALSO



* [goal app](../../app/app/)	 - Manage applications



