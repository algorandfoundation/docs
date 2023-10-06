title: Register online

This section assumes that you have already [generated a participation key](generate_keys.md) for the account you plan to mark online. 

Registering an account online requires authorizing a [key registration transaction](../../get-details/transactions/index.md#key-registration-transaction) with details of the participation key that will vote on the account's behalf. Once the transaction is processed by the blockchain, the Verifiable Random Function public key (referred to as the VRF public key) is written into the account’s data and the account will start participating in consensus with that key. This VRF public key is how the account is associated with the specific participation keys.

!!! info "Important"
	The moment a key registration transaction is confirmed by the network it takes 320 rounds for the change to take effect. In other words, if a key registration is confirmed in round 1000, the account will not start participating until round 1320.

# Create an online key registration transaction

Create a key registration transaction for the address: `EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4` by inserting the following code snippet into the construction portion of the example shown in [Authorizing Transactions Offline](../../get-details/transactions/offline_transactions.md#unsigned-transaction-file-operations). The file produced and displayed with `goal clerk inspect` should look almost exactly the same as the output shown in the [constructing a register online transaction example](../../get-details/transactions/index.md#register-account-online). 


=== "JavaScript"
    <!-- ===JSSDK_TRANSACTION_KEYREG_ONLINE_CREATE=== -->
	```javascript
	// get suggested parameters
	const params = await algodClient.getTransactionParams().do();
	
	// Parent addr
	const addr = 'MWAPNXBDFFD2V5KWXAHWKBO7FO4JN36VR4CIBDKDDE7WAUAGZIXM3QPJW4';
	// VRF public key
	const selectionKey = 'LrpLhvzr+QpN/bivh6IPpOaKGbGzTTB5lJtVfixmmgk=';
	// Voting pub key
	const voteKey = 'G/lqTV6MKspW6J8wH2d8ZliZ5XZVZsruqSBJMwLwlmo=';
	// State proof key
	const stateProofKey =
	  'RpUpNWfZMjZ1zOOjv3MF2tjO714jsBt0GKnNsw0ihJ4HSZwci+d9zvUi3i67LwFUJgjQ5Dz4zZgHgGduElnmSA==';
	
	// sets up keys for 100000 rounds
	const numRounds = 1e5;
	
	// dilution default is sqrt num rounds
	const keyDilution = numRounds ** 0.5;
	
	// create transaction
	const onlineKeyreg = algosdk.makeKeyRegistrationTxnWithSuggestedParamsFromObject(
	  {
	    from: addr,
	    voteKey,
	    selectionKey,
	    stateProofKey,
	    voteFirst: params.firstRound,
	    voteLast: params.firstRound + numRounds,
	    voteKeyDilution: keyDilution,
	    suggestedParams: params,
	  }
	);
	
	console.log(onlineKeyreg.get_obj_for_encoding());
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/participation.ts#L7-L41)
    <!-- ===JSSDK_TRANSACTION_KEYREG_ONLINE_CREATE=== -->

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
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/participation.py#L6-L26)
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
	
	Long numRounds = 100000l; // sets up keys for 100000 rounds
	Long keyDilution = (long) Math.sqrt(numRounds); // dilution default is sqrt num rounds
	
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
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/Participation.java#L13-L34)
    <!-- ===JAVASDK_TRANSACTION_KEYREG_ONLINE_CREATE=== -->

=== "Go"
    <!-- ===GOSDK_TRANSACTION_KEYREG_ONLINE_CREATE=== -->
	```go
	fromAddr := "MWAPNXBDFFD2V5KWXAHWKBO7FO4JN36VR4CIBDKDDE7WAUAGZIXM3QPJW4"
	voteKey := "87iBW46PP4BpTDz6+IEGvxY6JqEaOtV0g+VWcJqoqtc="
	selKey := "1V2BE2lbFvS937H7pJebN0zxkqe1Nrv+aVHDTPbYRlw="
	sProofKey := "f0CYOA4yXovNBFMFX+1I/tYVBaAl7VN6e0Ki5yZA3H6jGqsU/LYHNaBkMQ/rN4M4F3UmNcpaTmbVbq+GgDsrhQ=="
	voteFirst := uint64(16532750)
	voteLast := uint64(19532750)
	keyDilution := uint64(1732)
	nonpart := false
	tx, err := transaction.MakeKeyRegTxnWithStateProofKey(
		fromAddr,
		[]byte{},
		sp,
		voteKey,
		selKey,
		sProofKey,
		voteFirst,
		voteLast,
		keyDilution,
		nonpart,
	)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/participation/main.go#L27-L47)
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

[Verify that the participation key is on the node](../generate_keys#check-that-the-key-exists) prior to submitting the signed transaction. Once verified, wait for the network to reach the transaction's first valid round, then [submit](../../../archive/build-apps/hello_world#submit-the-transaction) the transaction and the SDK method "wait for confirmation".

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
