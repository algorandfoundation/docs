title: goal clerk group
---
## goal clerk group



Group transactions together



### Synopsis



Form a transaction group.  The input file must contain one or more unsigned transactions that will form a group.  The output file will contain the same transactions, in order, with a group flag added to each transaction, which requires that the transactions must be committed together. The group command would retain the logic signature, if present, as the TEAL program could verify the group using a logic signature argument.



```

goal clerk group [flags]

```



### Options



```

  -h, --help             help for group

  -i, --infile string    File storing transactions to be grouped

  -o, --outfile string   Filename for writing the grouped transactions

```



### Options inherited from parent commands



```

  -d, --datadir stringArray   Data directory for the node

  -k, --kmddir string         Data directory for kmd

  -w, --wallet string         Set the wallet to be used for the selected operation

```



### SEE ALSO



* [goal clerk](../../clerk/clerk/)	 - Provides the tools to control transactions 



