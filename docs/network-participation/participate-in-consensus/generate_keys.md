title: Generate a Participation Key

Algorand provides a set of keys for voting and proposing blocks separate from account spending keys. These are called participation keys.  At a high-level, participation keys are a specialized set of keys located on a single node. Once this participation key set is associated with an account, the account has the ability to participate in consensus. 

_[Read more about how Participation Keys function in the Algorand Consensus Protocol](../../learn/algorand_consensus.md#participation-keys)._


# Generate using `goal`
 Generate participation keys using the `goal account addpartkey` command. This command takes the address of the participating account, a range of rounds, and an optional key dilution parameter.  It then generates a [VRF key pair](../../learn/algorand_consensus.md#verifiable-random-function) and, using optimizations, generates a set of single-round voting keys for each round of the range specified. There is no range limit for the participation key. The key dilution parameter defaults to 10,000. To optimize for disk space, set the key dilution value to roughly the square root of the total number of rounds. 

!!! info
	The account’s private spending key does not need to be on the server that goal is running on as it is not used at all to generate participation keys.

!!! info
	For security, the individual keys for each round are deleted from the key file as each round is completed. It is critical for the safety of the Algorand blockchain to avoid storing backups of participation key files that have been registered for an account.  




The VRF private key is what is passed into the VRF to determine if you are selected to propose or vote on a block in any given round. As each round passes, voting keys are removed from the system. This prevents past votes from being modified if the node is ever compromised. Technically, the voting keys will never be used if an account has not been taken online. For an account to be taken online it first must have tokens in the account, participation keys generated for it and a “take online” transaction completed on the blockchain. To take an account online you can use the goal account changeonlinestatus command to create the transaction. This command provides a -t option that allows the transactions to be written to a file and signed offline (see signing offline). Once the changeonlinestatus transaction is processed by the blockchain, the VRF public key is written into the account’s data and the account will start participating in consensus with that key. This VRF public key is how the account is associated with the specific participation keys. Changing an account’s status takes 320 rounds after the transaction is processed by the blockchain.

A participation key 
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