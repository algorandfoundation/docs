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

The Address developers or users are typically shown is a 58 character long string corresponding to a base32 encoding of the byte array of the public key + a checksum.

Given an address `4H5UNRBJ2Q6JENAXQ6HNTGKLKINP4J4VTQBEPK5F3I6RDICMZBPGNH6KD4`, encoding to and from the public key format can be done as follows:
=== "JavaScript"
	<!-- ===JSSDK_CODEC_ADDRESS=== -->
	```javascript
	const address = '4H5UNRBJ2Q6JENAXQ6HNTGKLKINP4J4VTQBEPK5F3I6RDICMZBPGNH6KD4';
	const pk = algosdk.decodeAddress(address);
	const addr = algosdk.encodeAddress(pk.publicKey);
	console.log(address, addr);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/codec.ts#L16-L20)
	<!-- ===JSSDK_CODEC_ADDRESS=== -->

=== "Python"
	<!-- ===PYSDK_CODEC_ADDRESS=== -->
	```python
	address = "4H5UNRBJ2Q6JENAXQ6HNTGKLKINP4J4VTQBEPK5F3I6RDICMZBPGNH6KD4"
	pk = encoding.decode_address(address)
	addr = encoding.encode_address(pk)
	
	assert addr == address
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/codec.py#L8-L13)
	<!-- ===PYSDK_CODEC_ADDRESS=== -->

=== "Go"
	<!-- ===GOSDK_CODEC_ADDRESS=== -->
	```go
	address := "4H5UNRBJ2Q6JENAXQ6HNTGKLKINP4J4VTQBEPK5F3I6RDICMZBPGNH6KD4"
	pk, _ := types.DecodeAddress(address)
	addr := pk.String()
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/codec/main.go#L69-L72)
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
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/CodecExamples.java#L26-L32)
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
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/codec.ts#L23-L26)
	<!-- ===JSSDK_CODEC_BASE64=== -->

=== "Python"
	<!-- ===PYSDK_CODEC_BASE64=== -->
	```python
	encoded_str = "SGksIEknbSBkZWNvZGVkIGZyb20gYmFzZTY0"
	decoded_str = base64.b64decode(encoded_str).decode("utf-8")
	print(decoded_str)
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/codec.py#L24-L27)
	<!-- ===PYSDK_CODEC_BASE64=== -->

=== "Go"
	<!-- ===GOSDK_CODEC_BASE64=== -->
	```go
	encoded := "SGksIEknbSBkZWNvZGVkIGZyb20gYmFzZTY0"
	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	reencoded := base64.StdEncoding.EncodeToString(decoded)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/codec/main.go#L84-L87)
	<!-- ===GOSDK_CODEC_BASE64=== -->

=== "Java"
	<!-- ===JAVASDK_CODEC_BASE64=== -->
	```java
	String encodedStr = "SGksIEknbSBkZWNvZGVkIGZyb20gYmFzZTY0";
	byte[] decodedBytes = Encoder.decodeFromBase64(encodedStr);
	String reEncodedStr = Encoder.encodeToBase64(decodedBytes);
	assert encodedStr.equals(reEncodedStr);
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/CodecExamples.java#L35-L39)
	<!-- ===JAVASDK_CODEC_BASE64=== -->

### Integers

Integers in algorand are almost always uint64, sometimes it's required to encode them as bytes. For example when passing them as application arguments in an ApplicationCallTransaction. When encoding an integer to pass as an application argument, the integer should be encoded as the big endian 8 byte representation of the integer value.

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
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/codec.ts#L29-L34)
	<!-- ===JSSDK_CODEC_UINT64=== -->

=== "Python"
	<!-- ===PYSDK_CODEC_UINT64=== -->
	```python
	val = 1337
	encoded_uint = val.to_bytes(8, "big")
	decoded_uint = int.from_bytes(encoded_uint, byteorder="big")
	assert decoded_uint == val
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/codec.py#L30-L34)
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
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/codec/main.go#L91-L97)
	<!-- ===GOSDK_CODEC_UINT64=== -->

=== "Java"
	<!-- ===JAVASDK_CODEC_UINT64=== -->
	```java
	BigInteger val = BigInteger.valueOf(1337);
	byte[] encodedVal = Encoder.encodeUint64(val);
	BigInteger decodedVal = Encoder.decodeUint64(encodedVal);
	assert val.equals(decodedVal);
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/CodecExamples.java#L42-L46)
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

# ABI Encoding

All the SDKs support encoding and decoding of ABI values. The encoding is done using the [Algorand ABI specification](https://developer.algorand.org/docs/get-details/dapps/smart-contracts/ABI/). 


=== "JavaScript"
	<!-- ===JSSDK_CODEC_ABI==== -->
	```javascript
	const stringTupleCodec = algosdk.ABIType.from('(string,string)');
	
	const stringTupleData = ['hello', 'world'];
	const encodedTuple = stringTupleCodec.encode(stringTupleData);
	console.log(Buffer.from(encodedTuple).toString('hex'));
	
	const decodedTuple = stringTupleCodec.decode(encodedTuple);
	console.log(decodedTuple); // ['hello', 'world']
	
	const uintArrayCodec = algosdk.ABIType.from('uint64[]');
	
	const uintArrayData = [1, 2, 3, 4, 5];
	const encodedArray = uintArrayCodec.encode(uintArrayData);
	console.log(Buffer.from(encodedArray).toString('hex'));
	
	const decodeArray = uintArrayCodec.decode(encodedArray);
	console.log(decodeArray); // [1, 2, 3, 4, 5]
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/codec.ts#L63-L80)
	<!-- ===JSSDK_CODEC_ABI==== -->

=== "Python"
	<!-- ===PYSDK_CODEC_ABI==== -->
	```python
	from algosdk import abi
	
	# generate a codec from the string representation of the ABI type
	# in this case, a tuple of two strings
	codec = abi.ABIType.from_string("(string,string)")
	
	# encode the value to its ABI encoding with the codec
	to_encode = ["hello", "world"]
	encoded = codec.encode(to_encode)
	print(encoded.hex())
	
	# decode the value from its ABI encoding with the codec
	decoded = codec.decode(encoded)
	print(decoded)  # prints ["hello", "world"]
	
	# generate a codec for a uint64 array
	uint_array_codec = abi.ABIType.from_string("uint64[]")
	uint_array = [1, 2, 3, 4, 5]
	encoded_array = uint_array_codec.encode(uint_array)
	print(encoded_array.hex())
	
	decoded_array = uint_array_codec.decode(encoded_array)
	print(decoded_array)  # prints [1, 2, 3, 4, 5]
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/codec.py#L74-L97)
	<!-- ===PYSDK_CODEC_ABI=== -->

=== "Go"
	<!-- ===GOSDK_CODEC_ABI=== -->
	```go
	tupleCodec, _ := abi.TypeOf("(string,string)")
	
	tupleVal := []string{"hello", "world"}
	encodedTuple, _ := tupleCodec.Encode(tupleVal)
	log.Printf("%x", encodedTuple)
	
	decodedTuple, _ := tupleCodec.Decode(encodedTuple)
	log.Printf("%v", decodedTuple) // [hello world]
	
	arrCodec, _ := abi.TypeOf("uint64[]")
	arrVal := []uint64{1, 2, 3, 4, 5}
	encodedArr, _ := arrCodec.Encode(arrVal)
	log.Printf("%x", encodedArr)
	
	decodedArr, _ := arrCodec.Decode(encodedArr)
	log.Printf("%v", decodedArr) // [1 2 3 4 5]
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/codec/main.go#L101-L117)
	<!-- ===GOSDK_CODEC_ABI=== -->

=== "Java"
	<!-- ===JAVASDK_CODEC_ABI=== -->
	```java
	ABIType tupleCodec = ABIType.valueOf("(string,string)");
	Object[] tupleData = new Object[] { "hello", "world" };
	byte[] tupleEncoded = tupleCodec.encode(tupleData);
	System.out.printf("Encoded: '%s'\n", Hex.encodeHexString(tupleEncoded));
	Object tupleDecoded = tupleCodec.decode(tupleEncoded);
	// prints [hello, world]
	System.out.printf("Decoded: %s\n", Arrays.toString((Object[]) tupleDecoded));
	
	ABIType arrCodec = ABIType.valueOf("uint64[]");
	Object[] arrData = new Object[] { 1, 2, 3, 4, 5 };
	byte[] arrEncoded = arrCodec.encode(arrData);
	System.out.printf("Encoded: '%s'\n", Hex.encodeHexString(arrEncoded));
	
	Object arrDecoded = arrCodec.decode(arrEncoded);
	// prints [1, 2, 3, 4, 5]
	System.out.printf("Decoded: %s\n", Arrays.toString((Object[]) arrDecoded));
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/CodecExamples.java#L70-L86)
	<!-- ===JAVASDK_CODEC_ABI=== -->