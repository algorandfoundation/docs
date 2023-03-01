title: Register online

This section assumes that you have already [generated a participation key](generate_keys.md) for the account you plan to mark online. 

Registering an account online requires authorizing a [key registration transaction](../../get-details/transactions/index.md#key-registration-transaction) with details of the participation key that will vote on the account's behalf. Once the transaction is processed by the blockchain, the Verifiable Random Function public key (referred to as the VRF public key) is written into the account’s data and the account will start participating in consensus with that key. This VRF public key is how the account is associated with the specific participation keys.

!!! info "Important"
	The moment a key registration transaction is confirmed by the network it takes 320 rounds for the change to take effect. In other words, if a key registration is confirmed in round 1000, the account will not start participating until round 1320.

# Create an online key registration transaction

Create a key registration transaction for the address: `EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4` by inserting the following code snippet into the construction portion of the example shown in [Authorizing Transactions Offline](../../get-details/transactions/offline_transactions.md#unsigned-transaction-file-operations). The file produced and displayed with `goal clerk inspect` should look almost exactly the same as the output shown in the [constructing a register online transaction example](../../get-details/transactions/index.md#register-account-online). 


=== "Python"
    <!-- ===PYSDK_TRANSACTION_KEYREG_ONLINE_CREATE=== -->
```python
# get suggested parameters
params = algod_client.suggested_params()

votekey = "eXq34wzh2UIxCZaI1leALKyAvSz/+XOe0wqdHagM+bw="
selkey = "X84ReKTmp+yfgmMCbbokVqeFFFrKQeFZKEXG89SXwm4="

num_rounds = int(1e5)  # sets up keys for 100000 rounds
key_dilution = int(num_rounds**0.5)  # dilution default is sqrt num rounds

# create transaction
online_keyreg = transaction.KeyregTxn(
    sender="EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4",
    votekey=votekey,
    selkey=selkey,
    votefst=params.first,
    votelst=params.first + num_rounds,
    votekd=key_dilution,
    sp=params,
)
print(online_keyreg.dictify())
```
    <!-- ===PYSDK_TRANSACTION_KEYREG_ONLINE_CREATE=== -->

=== "Java"
    <!-- ===JAVASDK_TRANSACTION_KEYREG_ONLINE_CREATE=== -->
```java
        // get suggested parameters
        Response<TransactionParametersResponse> rsp = algodClient.TransactionParams().execute();
        TransactionParametersResponse sp = rsp.body();

        String address = "EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4";

        String votekey = "eXq34wzh2UIxCZaI1leALKyAvSz/+XOe0wqdHagM+bw=";
        String skey = "X84ReKTmp+yfgmMCbbokVqeFFFrKQeFZKEXG89SXwm4=";

        Long numRounds = 100000l;  // sets up keys for 100000 rounds
        Long keyDilution = (long) Math.sqrt(numRounds);  // dilution default is sqrt num rounds

        Transaction keyRegTxn = Transaction.KeyRegistrationTransactionBuilder().suggestedParams(sp)
                .sender(address)
                .selectionPublicKeyBase64(skey)
                .participationPublicKeyBase64(votekey)
                .voteFirst(sp.lastRound)
                .voteLast(sp.lastRound + numRounds)
                .voteKeyDilution(keyDilution)
                .build();
        // ... sign and send to network
```
    <!-- ===JAVASDK_TRANSACTION_KEYREG_ONLINE_CREATE=== -->

=== "Go"
    <!-- ===GOSDK_TRANSACTION_KEYREG_ONLINE_CREATE=== -->
    ```go 
    func saveUnsignedTransaction() {

        // setup connection
        algodClient := setupConnection()

        // get network suggested parameters
        txParams, err := algodClient.SuggestedParams()
        if err != nil {
            fmt.Printf("error getting suggested tx params: %s\n", err)
            return
        }

        // create transaction
        fromAddr := "EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4"
        genID := txParams.GenesisID
        genesisHash := base64.StdEncoding.EncodeToString(txParams.GenesisHash)
        voteKey := "eXq34wzh2UIxCZaI1leALKyAvSz/+XOe0wqdHagM+bw="
        selKey := "X84ReKTmp+yfgmMCbbokVqeFFFrKQeFZKEXG89SXwm4="
        voteFirst := uint64(6000000)
        voteLast := uint64(9000000)
        keyDilution := uint64(1730)
        tx, err := transaction.MakeKeyRegTxnWithFlatFee(fromAddr, 2000, 6002000,
            6003000, nil, genID, genesisHash, voteKey, selKey, voteFirst, voteLast,
            keyDilution)
        if err != nil {
            fmt.Printf("Error creating transaction: %s\n", err)
            return
        }
        unsignedTx := types.SignedTxn{
            Txn: tx,
        }
    ```
    <!-- ===GOSDK_TRANSACTION_KEYREG_ONLINE_CREATE=== -->

=== "goal"
    <!-- ===GOAL_TRANSACTION_KEYREG_ONLINE_CREATE=== -->
    ```zsh 
    # WARNING: This command must be run on the node where the partkey lives and the node
    # must only have a single partkey for the account. Otherwise the command will
    # choose one at random.
    $ goal account changeonlinestatus --address=EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4 --fee=2000 --firstvalid=6002000 --lastvalid=6003000 --online=true --txfile=online.txn
    ```
    <!-- ===GOAL_TRANSACTION_KEYREG_ONLINE_CREATE=== -->

# Authorize and Send the Transaction
Use the appropriate [authorization method](../../../get-details/transactions/signatures) to sign the transaction. 

!!! tip
    It is recommended that you authorize the transaction offline to protect the security of your private keys, especially for high-value participating accounts. 

[Verify that the participation key is in the node's ledger directory](../generate_keys#check-that-the-key-exists-in-the-nodes-ledger-directory) prior to submitting the signed transaction. Once verified, wait for the network to reach the transaction's first valid round, then [submit](../../../archive/build-apps/hello_world#submit-the-transaction) the transaction and the SDK method "wait for confirmation". 

# Check that the node is participating

At any time, you can validate whether your node is participating by `grep`-ing the `node.log` file in your data directory, looking for a `"VoteBroadcast"` messages where the `"Sender"` is your public key participation address.

```
$ grep 'VoteBroadcast' node.log
...
{"Context":"Agreement","Hash":"QJADVNJZDXYEQUPHITB6REFDGBY4AHBPPBIPVXLOPOASZA4T3PIA","ObjectPeriod":0,"ObjectRound":896659,"ObjectStep":2,"Period":0,"Round":0,"Sender":"3IE2GDYYSI56U53AQ6UUWRGAIGG5D4RHWLMCXJOPWQJA2ABF2X2A","Step":0,"Type":"VoteBroadcast","Weight":1,"WeightTotal":1,"file":"pseudonode.go","function":"github.com/algorand/go-algorand/agreement.pseudonodeVotesTask.execute","level":"info","line":344,"msg":"vote created for broadcast (weight 1, total weight 1)","time":"2019-05-10T18:38:54.137592-04:00"}
...
```

**See also**

- [Key Registration Transactions](../../../get-details/transactions/#key-registration-transaction)
- [Register account online](../../../get-details/transactions/#register-account-online)