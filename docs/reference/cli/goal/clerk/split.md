title: goal clerk split
---
## goal clerk split



Split a file containing many transactions into one transaction per file



### Synopsis



Split a file containing many transactions.  The input file must contain one or more transactions.  These transactions will be written to individual files.



```

goal clerk split [flags]

```



### Options



```

  -h, --help             help for split

  -i, --infile string    File storing transactions to be split

  -o, --outfile string   Base filename for writing the individual transactions; each transaction will be written to filename-N.ext

```



### Options inherited from parent commands



```

  -d, --datadir stringArray   Data directory for the node

  -k, --kmddir string         Data directory for kmd

  -w, --wallet string         Set the wallet to be used for the selected operation

```



### SEE ALSO



* [goal clerk](../../clerk/clerk/)	 - Provides the tools to control transactions 



