title: algokey part keyreg
---
## algokey part keyreg



Make key registration transaction



```

algokey part keyreg [flags]

```



### Options



```

      --account string      account address to bring offline; only specify when taking an account offline from voting in Algorand consensus

      --fee uint            transaction fee (default 1000)

      --firstvalid uint     first round where the transaction may be committed to the ledger

  -h, --help                help for keyreg

      --keyfile string      participation keys to register, file is opened to fetch metadata for the transaction; only specify when bringing an account online to vote in Algorand consensus

      --lastvalid uint      last round where the generated transaction may be committed to the ledger, defaults to firstvalid + 1000

      --network string      the network where the provided keys will be registered, one of mainnet/testnet/betanet (default "mainnet")

      --offline             set to bring an account offline

  -o, --outputFile string   write signed transaction to this file, or '-' to write to stdout

```



### SEE ALSO



* [algokey part](../../part/part/)	 - Manage participation keys



