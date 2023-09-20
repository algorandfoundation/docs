title: Register offline

To mark an account **offline** send a key registration transaction to the network authorized by the account to be marked offline. The signal to mark the sending account offline is the issuance of a `"type": "keyreg"` transaction that does not contain any participation key-related fields (i.e. they are all set to null values)

!!! info "Important"
	Just like with online keyreg transactions. The moment a key registration transaction is confirmed by the network it takes 320 rounds for the change to take effect. So, if a key registration is confirmed in round 5000, the account will stop participating at round 5320.

# Create an offline key registration transaction

Create an offline key registration transaction for the address: `EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4` by inserting the following code snippet into the construction portion of the example shown in [Authorizing Transactions Offline](../../../get-details/transactions/offline_transactions#unsigned-transaction-file-operations). The file produced and displayed with `goal clerk inspect` should look almost exactly the same as the output shown in the [constructing a register offline transaction example](../../../get-details/transactions#register-account-offline). 

=== "Python"
    <!-- ===PYSDK_TRANSACTION_KEYREG_OFFLINE_CREATE=== -->
	```python
	# get suggested parameters
	params = algod_client.suggested_params()
	
	# create keyreg transaction to take this account offline
	offline_keyreg = transaction.KeyregTxn(
	    sender="EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4",
	    sp=params,
	    votekey=None,
	    selkey=None,
	    votefst=None,
	    votelst=None,
	    votekd=None,
	)
	print(online_keyreg.dictify())
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/participation.py#L29-L43)
    <!-- ===PYSDK_TRANSACTION_KEYREG_OFFLINE_CREATE=== -->

=== "JavaScript"
    <!-- ===JSSDK_TRANSACTION_KEYREG_OFFLINE_CREATE=== -->
	```javascript
	// get suggested parameters
	const suggestedParams = await algodClient.getTransactionParams().do();
	// create keyreg transaction to take this account offline
	const offlineKeyReg = algosdk.makeKeyRegistrationTxnWithSuggestedParamsFromObject(
	  {
	    from: addr,
	    suggestedParams,
	  }
	);
	console.log(offlineKeyReg.get_obj_for_encoding());
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/participation.ts#L44-L54)
    <!-- ===JSSDK_TRANSACTION_KEYREG_OFFLINE_CREATE=== -->

=== "Go"
    <!-- ===GOSDK_TRANSACTION_KEYREG_OFFLINE_CREATE=== -->
    <!-- ===GOSDK_TRANSACTION_KEYREG_OFFLINE_CREATE=== -->

=== "Java"
    <!-- ===JAVASDK_TRANSACTION_KEYREG_OFFLINE_CREATE=== -->
	```java
	// create keyreg transaction to take this account offline
	Transaction keyRegOfflineTxn = Transaction.KeyRegistrationTransactionBuilder().suggestedParams(sp)
	        .sender(address)
	        .build();
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/Participation.java#L37-L41)
    <!-- ===JAVASDK_TRANSACTION_KEYREG_OFFLINE_CREATE=== -->

=== "goal"
    <!-- ===GOAL_TRANSACTION_KEYREG_OFFLINE_CREATE=== -->
    ```zsh
    $ goal account changeonlinestatus --address=EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4 --fee=1000 --firstvalid=7000000 --lastvalid=7001000 --online=false --txfile=offline.txn
    ```
    <!-- ===GOAL_TRANSACTION_KEYREG_OFFLINE_CREATE=== -->

# Authorize and Send the Transaction
Use the appropriate [authorization method](../../../get-details/transactions/signatures) to sign the transaction. 

!!! tip
    It is recommended that you authorize the transaction offline to protect the security of your private keys, especially for high-value participating accounts. 

Once the transaction is signed, wait for the network to reach the transaction's first valid round and then [submit](../../../archive/build-apps/hello_world#submit-the-transaction) and the SDK Method "wait for confirmation". 

**See also**

- [Key Registration Transactions](../../../get-details/transactions#key-registration-transaction)
- [Register account offline](../../../get-details/transactions#register-account-offline)