title: goal clerk compile
---
## goal clerk compile



Compile a contract program



### Synopsis



Reads a TEAL contract program and compiles it to binary output and contract address.



```

goal clerk compile [input file 1] [input file 2]... [flags]

```



### Options



```

  -a, --account string   Account address to sign the program (If not specified, uses default account)

  -D, --disassemble      disassemble a compiled program

  -h, --help             help for compile

  -m, --map              write out source map

  -n, --no-out           don't write contract program binary

  -o, --outfile string   Filename to write program bytes or signed LogicSig to

  -s, --sign             sign program, output is a binary signed LogicSig record

```



### Options inherited from parent commands



```

  -d, --datadir stringArray   Data directory for the node

  -k, --kmddir string         Data directory for kmd

  -w, --wallet string         Set the wallet to be used for the selected operation

```



### SEE ALSO



* [goal clerk](../../clerk/clerk/)	 - Provides the tools to control transactions 



