title: goal account importrootkey
---
## goal account importrootkey



Import .rootkey files from the data directory into a kmd wallet



### Synopsis



Import .rootkey files from the data directory into a kmd wallet. This is analogous to using the import command with an account seed mnemonic: the imported account will be displayed alongside your wallet-derived accounts, but will not be tied to your wallet mnemonic.



```

goal account importrootkey [flags]

```



### Options



```

  -h, --help                 help for importrootkey

  -u, --unencrypted-wallet   Import into the default unencrypted wallet, potentially creating it

```



### Options inherited from parent commands



```

  -d, --datadir stringArray   Data directory for the node

  -k, --kmddir string         Data directory for kmd

  -w, --wallet string         Set the wallet to be used for the selected operation

```



### SEE ALSO



* [goal account](../../account/account/)	 - Control and manage Algorand accounts



