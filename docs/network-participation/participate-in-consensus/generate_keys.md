title: Generate a Participation Key

## Generate a new participation key

```zsh
$ goal account addpartkey -a EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4 --roundFirstValid=6000000 --roundLastValid=9000000 --keyDilution=1730 
Participation key generation successful
```

```zsh
$ ls $ALGORAND_DATA/testnet-v1.0/*.partkey
/home/ubuntu/node/data/testnet-v1.0/EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4.6000000.9000000.partkey
```

## View participation key info
```zsh
$ goal account partkeyinfo
Dumping participation key info from /home/ubuntu/node/data...
------------------------------------------------------------------
File: EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4.6000000.9000000.partkey
{
  "acct": "EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4",
  "first": 6000000,
  "last": 9000000,
  "sel": "X84ReKTmp+yfgmMCbbokVqeFFFrKQeFZKEXG89SXwm4=",
  "vote": "eXq34wzh2UIxCZaI1leALKyAvSz/+XOe0wqdHagM+bw=",
  "voteKD": 1730
}
```
This info must be given to the entity that will create the corresponding [key registration transaction](../../feature-guides/transactions.md#keyreg).