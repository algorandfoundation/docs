title: indexer util validator
---
## indexer util validator



validator



### Synopsis



Compare algod and indexer to each other and report any discrepencies.



```

indexer util validator [flags]

```



### Options



```

      --addr string             If provided validate a single address instead of reading Stdin.

      --algod-token string      Algod token.

      --algod-url string        Algod url.

      --box string              If provided validate a single box (in the format "appid,b64boxname") instead of reading Stdin.

  -e, --error-log-file string   When specified, error messages are written to this file instead of to stderr.

  -h, --help                    help for validator

      --indexer-token string    Indexer token.

      --indexer-url string      Indexer url.

      --print-commands          Print curl commands, including tokens, to query algod and indexer contents.

      --print-skipped           Include accounts which were skipped in the error log.

      --processor int           Choose compare algorithm [0 = Struct, 1 = Reflection]

      --retries int             Number of retry attempts when a difference is detected. (default 5)

      --retry-delay int         Time in milliseconds to sleep between retries. (default 1000)

      --threads int             Number of worker threads to initialize. (default 4)

```



### SEE ALSO



* [indexer util](../../util/util/)	 - Utilities for testing Indexer operation and correctness.



