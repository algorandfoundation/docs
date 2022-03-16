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

  -d, --algod string             path to algod data dir, or $ALGORAND_DATA

      --algod-net string         host:port of algod

      --algod-token string       api access token for algod

      --allow-migration          allow migrations to happen even when no algod connected

  -c, --configfile string        file path to configuration file (indexer.yml)

      --cpuprofile string        file to record cpu profile to

      --dev-mode                 allow performance intensive operations like searching for accounts at a particular round

  -n, --dummydb                  use dummy indexer db

  -g, --genesis string           path to genesis.json (defaults to genesis.json in algod data dir if that was set)

  -h, --help                     help for daemon

  -f, --logfile string           file to write logs to, if unset logs are written to standard out

  -l, --loglevel string          verbosity of logs: [error, warn, info, debug, trace] (default "info")

      --max-conn uint32          set the maximum connections allowed in the connection pool, if the maximum is reached subsequent connections will wait until a connection becomes available, or timeout according to the read-timeout setting

      --metrics-mode string      configure the /metrics endpoint to [ON, OFF, VERBOSE] (default "OFF")

      --no-algod                 disable connecting to algod for block following

      --pidfile string           file to write daemon's process id to

  -P, --postgres string          connection string for postgres database

      --read-timeout duration    set the maximum duration for reading the entire request (default 5s)

  -S, --server string            host:port to serve API on (default :8980) (default ":8980")

  -t, --token string             an optional auth token, when set REST calls must use this token in a bearer format, or in a 'X-Indexer-API-Token' header

  -v, --version                  print version and exit

      --write-timeout duration   set the maximum duration to wait before timing out writes to a http response, breaking connection (default 30s)

```



### SEE ALSO



* [indexer](../../indexer/indexer/)	 - Algorand Indexer



