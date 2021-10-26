title: Generate a participation key

Algorand provides a set of keys for voting and proposing blocks separate from account spending keys. These are called **participation keys** (sometimes referred to as **partkeys**).  At a high-level, participation keys are a specialized set of keys located on a single node. Once this participation key set is associated with an account, the account has the ability to participate in consensus. 

_[Read more about how Participation Keys function in the Algorand Consensus Protocol](../../../get-details/algorand_consensus#participation-keys)._

!!! info "Important"
	- The account’s private spending key does not need to be on the node to generate a participation key. Technically, anyone can generate a participation key for a particular account, but only the private spending key of the account can authorize the transaction that would register the account to go online with a particular participation key. This distinction allows you to keep private keys in cold storage.
	- For security, the individual keys for each round are deleted from the key file as each round is completed. It is critical for the safety of the Algorand blockchain to avoid storing backups of participation key files that have been registered for an account.  
	- There is no theoretical limit to the range you can specify for a partkey validity period, only disk space constraints. A recommended range is 3,000,000 rounds.

!!! warning
	All the `goal` commands in this page must be run under the user that owns the `data` directory and runs `algod`. They should never run under the user `root`, as this may compromise permissions in the `data` folder. For example, with a default installation of the Debian package, `goal account ...` should be replaced by `sudo -u algorand -E goal account ...`.

# Generate the participation key with `goal`

To generate a participation key, use the [`goal account addpartkey`](../../../clis/goal/account/addpartkey) command on the node where the participation key will reside. This command takes the address of the participating account, a range of rounds, and an optional key dilution parameter.  It then generates a [VRF key pair](../../../get-details/algorand_consensus#verifiable-random-function) and, using optimizations, generates a set of single-round voting keys for each round of the range specified. The VRF private key is what is passed into the VRF to determine if you are selected to propose or vote on a block in any given round. 

=== "goal"
    ```zsh 
    $ goal account addpartkey -a <address-of-participating-account> --roundFirstValid=<partkey-first-round> --roundLastValid=<partkey-last-round> --keyDilution=<key-dilution-value> 
    Participation key generation successful
    ```

This creates a participation key in the ledger directory of the node, which is where it should ultimately live. Use the `-o` flag to specify a different directory in the case where you will eventually transfer your key to a different node's ledger directory.

!!! tip
	The Key Dilution parameter defaults to 10,000. To reduce the size of the participation key, set the key dilution value to roughly the square root of the range that the partkey is valid for.

# Check that the key exists in the node's ledger directory

The [`goal account listpartkeys`](../../../clis/goal/account/listpartkeys) command will check for any participation keys that live on the node and display pertinent information about them. 

=== "goal"
    ```zsh
    $ goal account listpartkeys
    Registered	Filename                                                                        	Parent address                                              	 First round	  Last round	   First key
    no        	BH55E5RMBD4GYWXGX5W5PJ5JAHPGM5OXKDQH5DC4O2MGI7NW4H6Q.10000.10111.partkey        	BH55E5RMBD4GYWXGX5W5PJ5JAHPGM5OXKDQH5DC4O2MGI7NW4H6VOE4CP4  	       10000	       10111	    240821.0
    no        	EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4.4595158.6000000.partkey	EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4  	     4595158	     6000000	    478.2927
    ```


The output above is an example of `goal account listpartkeys` run from a particular node. It displays all partkeys and whether or not each key has been **registered**, the **filename** of the participation key, the **first** and **last** rounds of validity for the partkey, the **parent address** (i.e. the address of the participating account) and the **first key**. The first key refers to the key batch and the index in that batch (`<key-batch>.<index>`) of the latest key that has not been deleted. This is useful in verifying that your node is participating (i.e. the batch should continue to increment as keys are deleted). It can also help ensure that you don't store extra copies of registered participation keys that have past round keys intact. 

If the key you generated does not show up, check that it was correctly placed in the ledger directory `ls $ALGORAND_DATA/<network-id>/ | grep <participating-account-address>`. If not, move it there and run this command again to confirm.


!!! warning
	It is okay to have multiple participation keys on a single node. However, if you generate multiple participation keys for the same account with overlapping rounds make sure you are aware of which one is the active one. It is recommended that you only keep one key per account - the active one - _except_ during partkey renewal when you switch from the old key to the new key. Renewing participation keys is discussed in detail in the [Renew Participation Keys](./renew.md) section.

# View participation key info

Use [`goal account partkeyinfo`](../../../clis/goal/account/partkeyinfo) to dump all the information about each participation key that lives on the node. This information is used to generate the online key registration transaction [described in the next section](./online.md).

=== "goal"
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
    ...
    ```


Above is the example output from a particular node. Use these values to create the [key registration transaction](../../../get-details/transactions#register-account-online) that will place the account online.
