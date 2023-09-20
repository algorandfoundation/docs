title: Offline signatures

This section explains how to authorize transactions with private keys that are kept **offline**. In particular, this guide shows how to create and save transactions to a file that can then be transferred to an offline device for signing. To learn about the structure of transactions and how to authorize them in general visit the [Transactions Structure](../../transactions/) and [Authorizing Transactions](../../transactions/signatures/) sections, respectively.

The same methodology described here can also be used to work with [LogicSignatures](../../transactions/signatures/#logic-signatures) and [Multisignatures](../../transactions/signatures/#multisignatures). All objects in the following examples use msgpack to store the transaction object ensuring interoperability with the SDKs and `goal`.

!!! info
    Storing keys _offline_ is also referred to as placing them in **cold storage**. An _online_ device that stores private keys is often referred to as a **hot wallet**.  

# Unsigned Transaction File Operations
Algorand SDK's and `goal` support writing and reading both signed and unsigned transactions to a file. Examples of these scenarios are shown in the following code snippets.

Unsigned transactions require the transaction object to be created before writing to a file.

=== "JavaScript"
	<!-- ===JSSDK_CODEC_TRANSACTION_UNSIGNED=== -->
	```javascript
	const txn = algosdk.makePaymentTxnWithSuggestedParamsFromObject({
	  from: sender.addr,
	  to: receiver.addr,
	  amount: 1e6,
	  suggestedParams,
	});
	
	const txnBytes = algosdk.encodeUnsignedTransaction(txn);
	const txnB64 = Buffer.from(txnBytes).toString('base64');
	// ...
	const restoredTxn = algosdk.decodeUnsignedTransaction(
	  Buffer.from(txnB64, 'base64')
	);
	console.log(restoredTxn);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/codec.ts#L37-L51)
	<!-- ===JSSDK_CODEC_TRANSACTION_UNSIGNED=== -->

=== "Python"
	<!-- ===PYSDK_CODEC_TRANSACTION_UNSIGNED=== -->
	```python
	sp = algod_client.suggested_params()
	pay_txn = transaction.PaymentTxn(acct.address, sp, acct.address, 10000)
	
	# Write message packed transaction to disk
	with open("pay.txn", "w") as f:
	    f.write(encoding.msgpack_encode(pay_txn))
	
	# Read message packed transaction and decode it to a Transaction object
	with open("pay.txn", "r") as f:
	    recovered_txn = encoding.msgpack_decode(f.read())
	
	print(recovered_txn.dictify())
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/codec.py#L40-L52)
	<!-- ===PYSDK_CODEC_TRANSACTION_UNSIGNED=== -->

=== "Java"
	<!-- ===JAVASDK_CODEC_TRANSACTION_UNSIGNED=== -->
	```java
	Response<TransactionParametersResponse> rsp = algodClient.TransactionParams().execute();
	TransactionParametersResponse sp = rsp.body();
	Transaction ptxn = Transaction.PaymentTransactionBuilder().suggestedParams(sp)
	        .sender(acct.getAddress()).receiver(acct.getAddress()).amount(100).build();
	
	byte[] encodedTxn = Encoder.encodeToMsgPack(ptxn);
	
	Transaction decodedTxn = Encoder.decodeFromMsgPack(encodedTxn, Transaction.class);
	assert decodedTxn.equals(ptxn);
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/CodecExamples.java#L49-L58)
	<!-- ===JAVASDK_CODEC_TRANSACTION_UNSIGNED=== -->

=== "Go"
	<!-- ===GOSDK_CODEC_TRANSACTION_UNSIGNED=== -->
	```go
	// Error handling omitted for brevity
	sp, _ := algodClient.SuggestedParams().Do(context.Background())
	ptxn, _ := transaction.MakePaymentTxn(
		acct1.Address.String(), acct1.Address.String(), 10000, nil, "", sp,
	)
	
	// Encode the txn as bytes,
	// if sending over the wire (like to a frontend) it should also be b64 encoded
	encodedTxn := msgpack.Encode(ptxn)
	os.WriteFile("pay.txn", encodedTxn, 0655)
	
	var recoveredPayTxn = types.Transaction{}
	
	msgpack.Decode(encodedTxn, &recoveredPayTxn)
	log.Printf("%+v", recoveredPayTxn)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/codec/main.go#L28-L43)
	<!-- ===GOSDK_CODEC_TRANSACTION_UNSIGNED=== -->

=== "goal"
	<!-- ===GOAL_CODEC_TRANSACTION_UNSIGNED=== -->
    ``` goal
    $ goal clerk send --from=<my-account> --to=GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A --fee=1000 --amount=1000000 --out="unsigned.txn"

    $ goal clerk sign --infile unsigned.txn --outfile signed.txn

    $ goal clerk rawsend --filename signed.txn

    ```
	<!-- ===GOAL_CODEC_TRANSACTION_UNSIGNED=== -->
# Signed Transaction File Operations 
Signed Transactions are similar, but require an account to sign the transaction before writing it to a file.

=== "JavaScript"
	<!-- ===JSSDK_CODEC_TRANSACTION_SIGNED=== -->
	```javascript
	const signedTxn = txn.signTxn(sender.privateKey);
	const signedB64Txn = Buffer.from(signedTxn).toString('base64');
	const restoredSignedTxn = algosdk.decodeSignedTransaction(
	  Buffer.from(signedB64Txn, 'base64')
	);
	console.log(restoredSignedTxn);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/codec.ts#L54-L60)
	<!-- ===JSSDK_CODEC_TRANSACTION_SIGNED=== -->

=== "Python"
	<!-- ===PYSDK_CODEC_TRANSACTION_SIGNED=== -->
	```python
	# Sign transaction
	spay_txn = pay_txn.sign(acct.private_key)
	# write message packed signed transaction to disk
	with open("signed_pay.txn", "w") as f:
	    f.write(encoding.msgpack_encode(spay_txn))
	
	# read message packed signed transaction into a SignedTransaction object
	with open("signed_pay.txn", "r") as f:
	    recovered_signed_txn = encoding.msgpack_decode(f.read())
	
	print(recovered_signed_txn.dictify())
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/codec.py#L57-L68)
	<!-- ===PYSDK_CODEC_TRANSACTION_SIGNED=== -->

=== "Java"
	<!-- ===JAVASDK_CODEC_TRANSACTION_SIGNED=== -->
	```java
	SignedTransaction signedTxn = acct.signTransaction(ptxn);
	byte[] encodedSignedTxn = Encoder.encodeToMsgPack(signedTxn);
	
	SignedTransaction decodedSignedTransaction = Encoder.decodeFromMsgPack(encodedSignedTxn,
	        SignedTransaction.class);
	assert decodedSignedTransaction.equals(signedTxn);
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/CodecExamples.java#L61-L67)
	<!-- ===JAVASDK_CODEC_TRANSACTION_SIGNED=== -->

=== "Go"
	<!-- ===GOSDK_CODEC_TRANSACTION_SIGNED=== -->
	```go
	// Assuming we already have a pay transaction `ptxn`
	
	// Sign the transaction
	_, signedTxn, err := crypto.SignTransaction(acct1.PrivateKey, ptxn)
	if err != nil {
		log.Fatalf("failed to sign transaction: %s", err)
	}
	
	// Save the signed transaction to file
	os.WriteFile("pay.stxn", signedTxn, 0644)
	
	signedPayTxn := types.SignedTxn{}
	err = msgpack.Decode(signedTxn, &signedPayTxn)
	if err != nil {
		log.Fatalf("failed to decode signed transaction: %s", err)
	}
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/codec/main.go#L46-L62)
	<!-- ===GOSDK_CODEC_TRANSACTION_SIGNED=== -->

=== "goal"
	<!-- ===GOAL_CODEC_TRANSACTION_SIGNED=== -->
    ``` goal
    $ goal clerk rawsend --filename signed.txn
    ```
	<!-- ===GOAL_CODEC_TRANSACTION_SIGNED=== -->

## Signature Verification

Sometimes a transaction is signed by a third party, and you want to verify that the signature is valid. This can be done by decoding the signed transaction into a `SignedTransaction` object using one of the SDKs and then running and ed25519 verify on the signature.

=== "JavaScript"
	<!-- ===JSSDK_OFFLINE_VERIFY_SIG=== -->
	```javascript
	const stxn = algosdk.decodeSignedTransaction(rawSignedTxn);
	if (stxn.sig === undefined) return false;
	
	// Get the txn object
	const txnObj = stxn.txn.get_obj_for_encoding();
	if (txnObj === undefined) return false;
	
	// Encode as msgpack
	const txnBytes = algosdk.encodeObj(txnObj);
	// Create byte array with TX prefix
	const msgBytes = new Uint8Array(txnBytes.length + 2);
	msgBytes.set(Buffer.from('TX'));
	msgBytes.set(txnBytes, 2);
	
	// Grab the other things we need to verify
	const pkBytes = stxn.txn.from.publicKey;
	const sigBytes = new Uint8Array(stxn.sig);
	
	// Return the result of the verification
	const valid = nacl.sign.detached.verify(msgBytes, sigBytes, pkBytes);
	console.log('Valid? ', valid);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/verify_signature.ts#L6-L27)
	<!-- ===JSSDK_OFFLINE_VERIFY_SIG=== -->

=== "Python"
	<!-- ===PYSDK_OFFLINE_VERIFY_SIG=== -->
	```python
	# decode the signed transaction
	stxn = encoding.msgpack_decode(raw_stxn)
	if stxn.signature is None or len(stxn.signature) == 0:
	    return False
	
	public_key = stxn.transaction.sender
	if stxn.authorizing_address is not None:
	    public_key = stxn.authorizing_address
	
	# Create a VerifyKey from nacl using the 32 byte public key
	verify_key = VerifyKey(encoding.decode_address(public_key))
	
	# Generate the message that was signed with TX prefix
	prefixed_message = b"TX" + base64.b64decode(
	    encoding.msgpack_encode(stxn.transaction)
	)
	
	try:
	    # Verify the signature
	    verify_key.verify(prefixed_message, base64.b64decode(stxn.signature))
	    return True
	except BadSignatureError:
	    return False
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/verify_signature.py#L9-L32)
	<!-- ===PYSDK_OFFLINE_VERIFY_SIG=== -->

=== "Go"
	<!-- ===GOSDK_OFFLINE_VERIFY_SIG=== -->
	```go
	signedTxn := types.SignedTxn{}
	msgpack.Decode(stxn, &signedTxn)
	
	from := signedTxn.Txn.Sender[:]
	
	encodedTx := msgpack.Encode(signedTxn.Txn)
	
	msgParts := [][]byte{txidPrefix, encodedTx}
	msg := bytes.Join(msgParts, nil)
	
	valid := ed25519.Verify(from, msg, signedTxn.Sig[:])
	
	log.Printf("Valid? %t", valid)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/verify_signature/main.go#L31-L44)
	<!-- ===GOSDK_OFFLINE_VERIFY_SIG=== -->

=== "Java"
	<!-- ===JAVASDK_OFFLINE_VERIFY_SIG=== -->
	```java
	// decode the signature
	SignedTransaction decodedSignedTransaction = Encoder.decodeFromMsgPack(rawSignedTxn,
	        SignedTransaction.class);
	Transaction txn = decodedSignedTransaction.tx;
	
	// get the bytes that were signed
	byte[] signedBytes = txn.bytesToSign();
	// get the pubkey that signed them
	PublicKey pk = txn.sender.toVerifyKey();
	
	// set up the sig checker
	java.security.Signature sigChecker = java.security.Signature.getInstance("Ed25519");
	sigChecker.initVerify(pk);
	sigChecker.update(signedBytes);
	// verify the signature 
	boolean valid = sigChecker.verify(decodedSignedTransaction.sig.getBytes());
	System.out.printf("Valid? %b\n", valid);
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/VerifySignature.java#L33-L50)
	<!-- ===JAVASDK_OFFLINE_VERIFY_SIG=== -->

    
??? example "Saving Signed and Unsigned Multisig Transactions to a File using goal"
    
    
    Create a multisig account by listing all of the accounts in the multisig and specifying the threshold number of accounts to sign with the -T flag

    ```bash
    goal account multisig new <my-account1> <my-account2> <my-account3> etc… -T 2    
    ```

    Create an unsigned transaction and write to file

    ```bash
    goal clerk send --from <my-multisig-account>  --to AZLR2XP4O2WFHLX6TX7AZVY23HLVLG3K5K3FRIKIYDOYN6ISIF54SA4RNY --fee=1000 --amount=1000000 --out="unsigned.txn"
    ```

    Sign by the required number of accounts to meet the threshold. 

    ```bash
    goal clerk multisig sign -a F <my-account1> -t=unsigned.txn
    goal clerk multisig sign -a F <my-account2> -t=unsigned.txn
    ```

    Merge signings 
    ```bash
    goal clerk multisig merge --out signed.txn unsigned.txn
    ```

    Broadcast

    ```bash
    goal clerk rawsend --filename signed.txn
    ```