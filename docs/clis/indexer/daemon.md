title: indexer daemon
---
## indexer daemon



run indexer daemon



### Synopsis



run indexer daemon. Serve api on HTTP.



```

indexer daemon [flags]

```



### Options



```

      --api-config-file string                 supply an API config file to enable/disable parameters

  -c, --configfile string                      file path to configuration file (indexer.yml)

      --cpuprofile string                      file to record cpu profile to

  -i, --data-dir string                        path to indexer data dir, or $INDEXER_DATA

      --default-accounts-limit uint32          set the default Limit parameter for querying accounts, if none is provided (default 100)

      --default-applications-limit uint32      set the default Limit parameter for querying applications, if none is provided (default 100)

      --default-assets-limit uint32            set the default Limit parameter for querying assets, if none is provided (default 100)

      --default-balances-limit uint32          set the default Limit parameter for querying balances, if none is provided (default 1000)

      --default-boxes-limit uint32             set the default allowed Limit parameter for searching an app's boxes (default 1000)

      --default-transactions-limit uint32      set the default Limit parameter for querying transactions, if none is provided (default 1000)

      --dev-mode                               has no effect currently, reserved for future performance intensive operations

  -n, --dummydb                                use dummy indexer db

      --enable-all-parameters                  override default configuration and enable all parameters. Can't be used with --api-config-file

      --enable-private-network-access-header   respond to Private Network Access preflight requests

  -h, --help                                   help for daemon

  -f, --logfile string                         file to write logs to, if unset logs are written to standard out

  -l, --loglevel string                        verbosity of logs: [error, warn, info, debug, trace] (default "info")

      --max-accounts-limit uint32              set the maximum allowed Limit parameter for querying accounts (default 1000)

      --max-api-resources-per-account uint32   set the maximum total number of resources (created assets, created apps, asset holdings, and application local state) per account that will be allowed in REST API lookupAccountByID and searchForAccounts responses before returning a 400 Bad Request. Set zero for no limit (default 1000)

      --max-applications-limit uint32          set the maximum allowed Limit parameter for querying applications (default 1000)

      --max-assets-limit uint32                set the maximum allowed Limit parameter for querying assets (default 1000)

      --max-balances-limit uint32              set the maximum allowed Limit parameter for querying balances (default 10000)

      --max-boxes-limit uint32                 set the maximum allowed Limit parameter for searching an app's boxes (default 10000)

      --max-conn uint32                        set the maximum connections allowed in the connection pool, if the maximum is reached subsequent connections will wait until a connection becomes available, or timeout according to the read-timeout setting

      --max-transactions-limit uint32          set the maximum allowed Limit parameter for querying transactions (default 10000)

      --metrics-mode string                    configure the /metrics endpoint to [ON, OFF, VERBOSE] (default "OFF")

      --pidfile string                         file to write daemon's process id to

  -P, --postgres string                        connection string for postgres database

      --read-timeout duration                  set the maximum duration for reading the entire request (default 5s)

  -S, --server string                          host:port to serve API on (default :8980) (default ":8980")

  -t, --token string                           an optional auth token, when set REST calls must use this token in a bearer format, or in a 'X-Indexer-API-Token' header

  -v, --version                                print version and exit

      --write-timeout duration                 set the maximum duration to wait before timing out writes to a http response, breaking connection (default 30s)

```



### SEE ALSO



* [indexer](../../indexer/indexer/)	 - Algorand Indexer



