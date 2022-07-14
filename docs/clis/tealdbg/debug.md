title: tealdbg debug
---
## tealdbg debug



Debug TEAL program(s) off-chain



### Synopsis



Debug TEAL program(s) in controlled environment using a local TEAL evaluator



```

tealdbg debug [program.tok [program.teal ...]] [flags]

```



### Options



```

  -a, --app-id uint            Application ID for stateful TEAL if not set in transaction(s) (default 1380011588)

  -b, --balance string         Balance records to evaluate stateful TEAL on in form of json or msgpack file

  -d, --dryrun-req string      Program(s) and state(s) in dryrun REST request format

  -g, --group-index int        Transaction index in a txn group

  -h, --help                   help for debug

      --indexer-token string   API token for indexer to fetch Balance records from to evaluate stateful TEAL

  -i, --indexer-url string     URL for indexer to fetch Balance records from to evaluate stateful TEAL

  -l, --latest-timestamp int   Latest confirmed timestamp to evaluate stateful TEAL on

  -q, --listen-dr-req          Listen for upcoming debugging dryrun request objects instead of taking program(s) from command line

  -m, --mode string            TEAL evaluation mode: auto, signature, application (default "auto")

      --painless               Automatically create balance record for all accounts and applications

  -p, --proto string           Consensus protocol version for TEAL evaluation

  -r, --round uint             Ledger round number to evaluate stateful TEAL on

  -t, --txn string             Transaction(s) to evaluate TEAL on in form of json or msgpack file

```



### Options inherited from parent commands



```

  -f, --frontend string             Frontend to use: cdt, web (default "cdt")

      --listen string               Network interface to listen on (default "127.0.0.1")

      --no-source-map               Do not generate source maps

      --remote-debugging-port int   Port to listen on (default 9392)

  -v, --verbose                     Verbose output

```



### SEE ALSO



* [tealdbg](../../tealdbg/tealdbg/)	 - Algorand TEAL Debugger



