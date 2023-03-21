title: Encoding and Decoding

When working with data from an algorand REST server or preparing transactions there is often a need to encode or decode fields. 


## Encoding Types

### JSON

The encoding most often returned when querying the state of the chain is [JSON](https://www.json.org/json-en.html).

It is easy to visually inspect but may be relatively slow to parse.

All [byte arrays](#byte-arrays) are base 64 encoded strings

### MessagePack

The encoding used when transmitting transactions to a node is [MessagePack](https://msgpack.org/index.html)

To inspect a given msgpack file contents a convenience commandline tool is provided:

```sh
    msgpacktool -d < file.msgp
```

### Base64

The encoding for byte arrays is [Base64](https://en.wikipedia.org/wiki/Base64).

This is to make it safe for the byte array to be transmitted as part of a json object.

### Base32

The encoding used for Addresses and Transaction Ids is [Base32](https://en.wikipedia.org/wiki/Base32)



## Individual Field encodings

### Address

In Algorand a [public key](../accounts/#transformation-public-key-to-algorand-address) is a 32 byte array. 

The Address developers or users are typically shown is a 58 byte string corresponding to a base32 encoding of the byte array of the public key + a checksum.

Given an address `4H5UNRBJ2Q6JENAXQ6HNTGKLKINP4J4VTQBEPK5F3I6RDICMZBPGNH6KD4`, encoding to and from the public key format can be done as follows:
=== "JavaScript"
	<!-- ===JSSDK_CODEC_ADDRESS=== -->
	```javascript
	const address = '4H5UNRBJ2Q6JENAXQ6HNTGKLKINP4J4VTQBEPK5F3I6RDICMZBPGNH6KD4';
	const pk = algosdk.decodeAddress(address);
	const addr = algosdk.encodeAddress(pk.publicKey);
	console.log(address, addr);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/encoding.ts#L16-L20)
	<!-- ===JSSDK_CODEC_ADDRESS=== -->

=== "Python"
	<!-- ===PYSDK_CODEC_ADDRESS=== -->
	```python
	address = "4H5UNRBJ2Q6JENAXQ6HNTGKLKINP4J4VTQBEPK5F3I6RDICMZBPGNH6KD4"
	pk = encoding.decode_address(address)
	addr = encoding.encode_address(pk)
	
	assert addr == address
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/codec.py#L7-L12)
	<!-- ===PYSDK_CODEC_ADDRESS=== -->

=== "Go"
	<!-- ===GOSDK_CODEC_ADDRESS=== -->
	```go
	address := "4H5UNRBJ2Q6JENAXQ6HNTGKLKINP4J4VTQBEPK5F3I6RDICMZBPGNH6KD4"
	pk, _ := types.DecodeAddress(address)
	addr := pk.String()
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/codec.go#L62-L65)
	<!-- ===GOSDK_CODEC_ADDRESS=== -->

=== "Java"
	<!-- ===JAVASDK_CODEC_ADDRESS=== -->
	```java
	String addrAsStr = "4H5UNRBJ2Q6JENAXQ6HNTGKLKINP4J4VTQBEPK5F3I6RDICMZBPGNH6KD4";
	// Instantiate a new Address object with string
	Address addr = new Address(addrAsStr);
	// Or with the bytes
	Address addrAgain = new Address(addr.getBytes());
	assert addrAgain.equals(addr);
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/CodecExamples.java#L25-L31)
	<!-- ===JAVASDK_CODEC_ADDRESS=== -->

### Byte Arrays

When transmitting an array of bytes over the network, byte arrays are base64 encoded.  The SDK will handle encoding from a byte array to base64 but may not decode some fields and you'll have to handle it yourself. For example compiled program results or the keys and values in a state delta in an application call will be returned as base64 encoded strings.

*Example:*

Given a base64 encoded byte array `SGksIEknbSBkZWNvZGVkIGZyb20gYmFzZTY0` it may be decoded as follows:

=== "JavaScript"
	<!-- ===JSSDK_CODEC_BASE64=== -->
	```javascript
	const b64Encoded = 'SGksIEknbSBkZWNvZGVkIGZyb20gYmFzZTY0';
	const b64Decoded = Buffer.from(b64Encoded, 'base64').toString();
	console.log(b64Encoded, b64Decoded);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/encoding.ts#L23-L26)
	<!-- ===JSSDK_CODEC_BASE64=== -->

=== "Python"
	<!-- ===PYSDK_CODEC_BASE64=== -->
	```python
	encoded_str = "SGksIEknbSBkZWNvZGVkIGZyb20gYmFzZTY0"
	decoded_str = base64.b64decode(encoded_str).decode("utf-8")
	print(decoded_str)
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/codec.py#L15-L18)
	<!-- ===PYSDK_CODEC_BASE64=== -->

=== "Go"
	<!-- ===GOSDK_CODEC_BASE64=== -->
	```go
	encoded := "SGksIEknbSBkZWNvZGVkIGZyb20gYmFzZTY0"
	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	reencoded := base64.StdEncoding.EncodeToString(decoded)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/codec.go#L69-L72)
	<!-- ===GOSDK_CODEC_BASE64=== -->

=== "Java"
	<!-- ===JAVASDK_CODEC_BASE64=== -->
	```java
	String encodedStr = "SGksIEknbSBkZWNvZGVkIGZyb20gYmFzZTY0";
	byte[] decodedBytes = Encoder.decodeFromBase64(encodedStr);
	String reEncodedStr = Encoder.encodeToBase64(decodedBytes);
	assert encodedStr.equals(reEncodedStr);
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/CodecExamples.java#L34-L38)
	<!-- ===JAVASDK_CODEC_BASE64=== -->

### Integers

Integers in algorand are almost always uint64, sometimes its required to encode them as bytes. For example when passing them as application arguments in an ApplicationCallTransaction. When encoding an integer to pass as an application argument, the integer should be encoded as the big endian 8 byte representation of the integer value.

*Example:*

Given an integer `1337`, you may encode it as:

=== "JavaScript"
	<!-- ===JSSDK_CODEC_UINT64=== -->
	```javascript
	const int = 1337;
	const encoded = algosdk.encodeUint64(int);
	const safeDecoded = algosdk.decodeUint64(encoded, 'safe');
	const mixedDecoded = algosdk.decodeUint64(encoded, 'bigint');
	console.log(int, encoded, safeDecoded, mixedDecoded);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/encoding.ts#L29-L34)
	<!-- ===JSSDK_CODEC_UINT64=== -->

=== "Python"
	<!-- ===PYSDK_CODEC_UINT64=== -->
	```python
	val = 1337
	encoded_uint = val.to_bytes(8, "big")
	decoded_uint = int.from_bytes(encoded_uint, byteorder="big")
	assert decoded_uint == val
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/codec.py#L21-L25)
	<!-- ===PYSDK_CODEC_UINT64=== -->

=== "Go"
	<!-- ===GOSDK_CODEC_UINT64=== -->
	```go
	val := 1337
	encodedInt := make([]byte, 8)
	binary.BigEndian.PutUint64(encodedInt, uint64(val))
	
	decodedInt := binary.BigEndian.Uint64(encodedInt)
	// decodedInt == val
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/codec.go#L76-L82)
	<!-- ===GOSDK_CODEC_UINT64=== -->

=== "Java"
	<!-- ===JAVASDK_CODEC_UINT64=== -->
	```java
	BigInteger val = BigInteger.valueOf(1337);
	byte[] encodedVal = Encoder.encodeUint64(val);
	BigInteger decodedVal = Encoder.decodeUint64(encodedVal);
	assert val.equals(decodedVal);
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/CodecExamples.java#L41-L45)
	<!-- ===JAVASDK_CODEC_UINT64=== -->


## Working with encoded structures

### Transaction

Sometimes an application needs to transmit a transaction or transaction group between the front end and back end. This can be done by msgpack encoding the transaction object on one side and msgpack decoding it on the other side. Often the msgpack'd bytes will be base64 encoded so that they can be safely transmitted in some json payload so we use that encoding here.

Essentially the encoding is: 

`tx_byte_str = base64encode(msgpack_encode(tx_obj))` 

and decoding is:

`tx_obj = msgpack_decode(base64decode(tx_byte_str))`


*Example*

Create a payment transaction from one account to another using suggested parameters and amount 10000, we write the msgpack encoded bytes

=== "JavaScript"
	<!-- ===JSSDK_CODEC_TRANSACTION_UNSIGNED=== -->
	```javascript
	const txn = algosdk.makePaymentTxnWithSuggestedParamsFromObject({
	  from: sender.addr,
	  to: receiver.addr,
	  amount: 1e6,
	  suggestedParams,
	});
	
	const txnBytes = txn.toByte();
	const txnB64 = Buffer.from(txnBytes).toString('base64');
	const restoredTxn = algosdk.decodeUnsignedTransaction(
	  Buffer.from(txnB64, 'base64')
	);
	console.log(restoredTxn);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/encoding.ts#L37-L50)
	<!-- ===JSSDK_CODEC_TRANSACTION_UNSIGNED=== -->
	<!-- ===JSSDK_CODEC_TRANSACTION_SIGNED=== -->
	```javascript
	const signedTxn = txn.signTxn(sender.privateKey);
	const signedB64Txn = Buffer.from(signedTxn).toString('base64');
	const restoredSignedTxn = algosdk.decodeSignedTransaction(
	  Buffer.from(signedB64Txn, 'base64')
	);
	console.log(restoredSignedTxn);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/encoding.ts#L53-L59)
	<!-- ===JSSDK_CODEC_TRANSACTION_SIGNED=== -->

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
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/codec.py#L31-L43)
	<!-- ===PYSDK_CODEC_TRANSACTION_UNSIGNED=== -->
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
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/codec.py#L48-L59)
	<!-- ===PYSDK_CODEC_TRANSACTION_SIGNED=== -->

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
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/codec.go#L25-L40)
	<!-- ===GOSDK_CODEC_TRANSACTION_UNSIGNED=== -->
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
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/codec.go#L43-L59)
	<!-- ===GOSDK_CODEC_TRANSACTION_SIGNED=== -->

=== "Java"
	<!-- ===JAVASDK_CODEC_TRANSACTION_UNSIGNED=== -->
	```java
	Response<TransactionParametersResponse> rsp = algodClient.TransactionParams().execute();
	TransactionParametersResponse sp = rsp.body();
	// Wipe the `reserve` address through an AssetConfigTransaction
	Transaction ptxn = Transaction.PaymentTransactionBuilder().suggestedParams(sp)
	        .sender(acct.getAddress()).receiver(acct.getAddress()).amount(100).build();
	
	byte[] encodedTxn = Encoder.encodeToMsgPack(ptxn);
	
	Transaction decodedTxn = Encoder.decodeFromMsgPack(encodedTxn, Transaction.class);
	assert decodedTxn.equals(ptxn);
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/CodecExamples.java#L48-L58)
	<!-- ===JAVASDK_CODEC_TRANSACTION_UNSIGNED=== -->
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

