title: indexer util block-generator runner
---
## indexer util block-generator runner



Run test suite and collect results.



### Synopsis



Run an automated test suite using the block-generator daemon and a provided algorand-indexer binary. Results are captured to a specified output directory.



```

indexer util block-generator runner [flags]

```



### Options



```

      --cpuprofile string                   Path where Indexer writes its CPU profile.

  -h, --help                                help for runner

  -i, --indexer-binary string               Path to indexer binary.

  -p, --indexer-port uint                   Port to start the server at. This is useful if you have a prometheus server for collecting additional data. (default 4010)

  -l, --log-level string                    LogLevel to use when starting Indexer. [error, warn, info, debug, trace] (default "error")

  -c, --postgres-connection-string string   Postgres connection string.

  -r, --report-directory string             Location to place test reports.

      --reset                               If set any existing report directory will be deleted before running tests.

  -s, --scenario string                     Directory containing scenarios, or specific scenario file.

  -d, --test-duration duration              Duration to use for each scenario. (default 5m0s)

      --validate                            If set the validator will run after test-duration has elapsed to verify data is correct. An extra line in each report indicates validator success or failure.

```



### SEE ALSO



* [indexer util block-generator](../../block-generator/block-generator/)	 - Block generator testing tools.



