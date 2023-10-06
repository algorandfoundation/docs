title: Generate a participation key

Algorand provides a set of keys for voting and proposing blocks separate from account spending keys. These are called **participation keys** (sometimes referred to as **partkeys**).  At a high-level, participation keys are a specialized set of keys located on a single node. Once this participation key set is associated with an account, the account has the ability to participate in consensus. 

_[Read more about how Participation Keys function in the Algorand Consensus Protocol](../../../get-details/algorand_consensus#participation-keys)._

!!! info "Important"
	- The account’s private spending key does not need to be on the node to generate a participation key. Technically, anyone can generate a participation key for a particular account, but only the private spending key of the account can authorize the transaction that would register the account to go online with a particular participation key. This distinction allows you to keep private keys in cold storage.
	- For security, the individual keys for each round are deleted from the key file as each round is completed. It is critical for the safety of the Algorand blockchain to avoid storing backups of participation key files that have been registered for an account.  
	- The limit to the range you can specify for a partkey validity period is 2^24 - 1 (16,777,215). For security, it is recommended that the range does not exceed 3,000,000 rounds.

# Generate the participation key with `goal`

To generate a participation key, use the [`goal account addpartkey`](../../../clis/goal/account/addpartkey) command on the node where the participation key will reside. This command takes the address of the participating account, a range of rounds, and an optional key dilution parameter.  It then generates a [VRF key pair](../../../get-details/algorand_consensus#verifiable-random-function) and, using optimizations, generates a set of single-round voting keys for each round of the range specified. The VRF private key is what is passed into the VRF to determine if you are selected to propose or vote on a block in any given round. 

=== "goal"
    ```zsh 
    $ goal account addpartkey -a <address-of-participating-account> --roundFirstValid=<partkey-first-round> --roundLastValid=<partkey-last-round>
    Participation key generation successful
    ```

This creates a participation key on the node. You can use the `-o` flag to specify a different location in the case where you will eventually transfer your key to a different node to construct the keyreg transaction.

!!! tip
	To optimize storage, the Key Dilution parameter defaults to the square root of the participation period length but this can be overridden with the flag `--keyDilution`. The Key Dilution determines how many ephemeral keys will be stored on an algorand node, as they are generated in batches. For example, if your participation period is set to 3,000,000 rounds, a batch of 1,732 ephemeral keys will be generated upfront, with additional batches getting generated after each set is used.

# Add participation key

If you chose to save the participation key and now want to add it to the server, you can use the following command to add the partkey file to the node.

```bash
$ goal account installpartkey --partkey ALICE...VWXYZ.0.30000.partkey --delete-input
```

# Check that the key exists

The [`goal account listpartkeys`](../../../clis/goal/account/listpartkeys) command will check for any participation keys that live on the node and display pertinent information about them. 

=== "goal"
    ```zsh
    $ goal account listpartkeys
    Registered  Account      ParticipationID   Last Used  First round  Last round
    yes         TUQ4...NLQQ  GOWHR456...              27            0     3000000
    ```


The output above is an example of `goal account listpartkeys` run from a particular node. It displays all partkeys and whether or not each key has been **registered**, the **filename** of the participation key, the **first** and **last** rounds of validity for the partkey, the **parent address** (i.e. the address of the participating account) and the **first key**. The first key refers to the key batch and the index in that batch (`<key-batch>.<index>`) of the latest key that has not been deleted. This is useful in verifying that your node is participating (i.e. the batch should continue to increment as keys are deleted). It can also help ensure that you don't store extra copies of registered participation keys that have past round keys intact. 


!!! warning
	It is okay to have multiple participation keys on a single node. However, if you generate multiple participation keys for the same account with overlapping rounds make sure you are aware of which one is the active one. It is recommended that you only keep one key per account - the active one - _except_ during partkey renewal when you switch from the old key to the new key. Renewing participation keys is discussed in detail in the [Renew Participation Keys](./renew.md) section.

# View participation key info

Use [`goal account partkeyinfo`](../../../clis/goal/account/partkeyinfo) to dump all the information about each participation key that lives on the node. This information is used to generate the online key registration transaction [described in the next section](./online.md).

=== "goal"
    ```zsh
    $ goal account partkeyinfo
    Dumping participation key info from /opt/data...

    Participation ID:          GOWHR456IK3LPU5KIJ66CRDLZM55MYV2OGNW7QTZYF5RNZEVS33A
    Parent address:            TUQ4HOIR3G5Z3BZUN2W2XTWVJ3AUUME4OKLINJFAGKBO4Y76L4UT5WNLQQ
    Last vote round:           11
    Last block proposal round: 12
    Effective first round:     1
    Effective last round:      3000000
    First round:               0
    Last round:                3000000
    Key dilution:              10000
    Selection key:             l6MsaTt7AiCAdG+69LG/wjaprsI1vImZuGN6gQ1jS88=
    Voting key:                Rleu99r3UqlwuuhaxCTrTQUuq1C9qk5uJd2WQQEG+6U=
    ```


Above is the example output from a particular node. Use these values to create the [key registration transaction](../../../get-details/transactions#register-account-online) that will place the account online.
