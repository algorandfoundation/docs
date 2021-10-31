title: Register offline

To mark an account **offline** send a key registration transaction to the network authorized by the account to be marked offline. The signal to mark the sending account offline is the issuance of a `"type": "keyreg"` transaction that does not contain any participation key-related fields (i.e. they are all set to null values)

!!! info "Important"
	Just like with online keyreg transactions. The moment a key registration transaction is confirmed by the network it takes 320 rounds for the change to take effect. So, if a key registration is confirmed in round 5000, the account will stop participating at round 5320.

# Create an offline key registration transaction

Create an offline key registration transaction for the address: `EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4` by inserting the following code snippet into the construction portion of the example shown in [Authorizing Transactions Offline](../../../get-details/transactions/offline_transactions#unsigned-transaction-file-operations). The file produced and displayed with `goal clerk inspect` should look almost exactly the same as the output shown in the [constructing a register offline transaction example](../../../get-details/transactions#register-account-offline). 

=== "Python"
    ```python
    ...
    def write_unsigned():
        # setup none connection
        algod_client = connect_to_network()

        # get suggested parameters
        params = algod_client.suggested_params()
        # create transaction
        data = {
            "sender": "EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4",
            "votekey": None,
            "selkey": None,
            "votefst": None,
            "votelst": None,
            "votekd": None,
            "fee": 1000,
            "flat_fee": True,
            "first": 7000000,
            "last": 7001000,
            "gen": params.get('genesisID'),
            "gh": params.get('genesishashb64')
        }
        txn = transaction.KeyregTxn(**data)
    ...
    ```

=== "Java"
	```java

    ...
            final String SRC_ADDR = "EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4";

            try {
                // Get suggested parameters from the node
                TransactionParams params = algodApiInstance.transactionParams();

                // create transaction
                String genId = params.getGenesisID();
                Digest genesisHash = new Digest(params.getGenesishashb64());
                BigInteger firstRound = BigInteger.valueOf(7000000);
                BigInteger lastRound = BigInteger.valueOf(7001000);
                BigInteger fee = BigInteger.valueOf(1000);
                Transaction tx = new Transaction(new Address(SRC_ADDR), fee, firstRound, lastRound,
                        null, genId, genesisHash, null, null,  null, null, null);
    ...
    ```

=== "goal"
    ```zsh
    $ goal account changeonlinestatus --address=EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4 --fee=1000 --firstvalid=7000000 --lastvalid=7001000 --online=false --txfile=offline.txn
    ```

# Authorize and Send the Transaction
Use the appropriate [authorization method](../../../get-details/transactions/signatures) to sign the transaction. 

!!! tip
    It is recommended that you authorize the transaction offline to protect the security of your private keys, especially for high-value participating accounts. 

Once the transaction is signed, wait for the network to reach the transaction's first valid round and then [submit](../../../archive/build-apps/hello_world#submit-the-transaction) and [wait for confirmation](../../../archive/build-apps/hello_world#wait-for-confirmation). 

**See also**

- [Key Registration Transactions](../../../get-details/transactions#key-registration-transaction)
- [Register account offline](../../../get-details/transactions#register-account-offline)
