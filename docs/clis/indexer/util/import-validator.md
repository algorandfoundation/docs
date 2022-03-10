title: indexer util import-validator
---
## indexer util import-validator



Import validator



### Synopsis



Run with the special import validator mode. A sqlite ledger will be maintained along with the standard Indexer database. Each will increment the round in lock-step, and compare the account state of each modified account before moving on to the next round. If a data discrepency is detected, an error will be printed to stderr and the program will terminate.



```

indexer util import-validator [flags]

```



### Options



```

      --algod-ledger string   path to algod ledger directory

      --algod-net string      host:port of algod

      --algod-token string    api access token for algod

  -h, --help                  help for import-validator

      --postgres string       connection string for postgres database

```



### SEE ALSO



* [indexer util](../../util/util/)	 - Utilities for testing Indexer operation and correctness.



