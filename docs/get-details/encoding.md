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
    ```js
    import algosdk from 'algosdk'

    const address = "4H5UNRBJ2Q6JENAXQ6HNTGKLKINP4J4VTQBEPK5F3I6RDICMZBPGNH6KD4"
    const pk = algosdk.decodeAddress(address)
    const addr = algosdk.encodeAddress(pk.publicKey)
    // addr === address
    ```
<!-- ===JSSDK_CODEC_ADDRESS=== -->

=== "Python"
<!-- ===PYSDK_CODEC_ADDRESS=== -->
```python
address = "4H5UNRBJ2Q6JENAXQ6HNTGKLKINP4J4VTQBEPK5F3I6RDICMZBPGNH6KD4"
pk = encoding.decode_address(address)
addr = encoding.encode_address(pk)

assert addr == address
```
<!-- ===PYSDK_CODEC_ADDRESS=== -->

=== "Go"
<!-- ===GOSDK_CODEC_ADDRESS=== -->
    ```go
    import "github.com/algorand/go-algorand-sdk/types"

    //...

    address := "4H5UNRBJ2Q6JENAXQ6HNTGKLKINP4J4VTQBEPK5F3I6RDICMZBPGNH6KD4"
    pk, _ := types.DecodeAddress(address)
    addr := pk.String()
    
    //addr == address
    ```
<!-- ===GOSDK_CODEC_ADDRESS=== -->

=== "Java"
<!-- ===JAVASDK_CODEC_ADDRESS=== -->
    ```java
    import com.algorand.algosdk.crypto.Address;

    //...

    String address = "4H5UNRBJ2Q6JENAXQ6HNTGKLKINP4J4VTQBEPK5F3I6RDICMZBPGNH6KD4";
    Address pk = new Address(address);
    String addr = new Address(pk.getBytes());

    // addr == address
    ```
<!-- ===JAVASDK_CODEC_ADDRESS=== -->

### Byte Arrays

When transmitting an array of bytes over the network, byte arrays are base64 encoded.  The SDK will handle encoding from a byte array to base64 but may not decode some fields and you'll have to handle it yourself. For example compiled program results or the keys and values in a state delta in an application call will be returned as base64 encoded strings.

*Example:*

Given a base64 encoded byte array `SGksIEknbSBkZWNvZGVkIGZyb20gYmFzZTY0` it may be decoded as follows:

=== "JavaScript"
<!-- ===JSSDK_CODEC_BASE64=== -->
    ```js
    const encoded = "SGksIEknbSBkZWNvZGVkIGZyb20gYmFzZTY0"
    const decoded = Buffer.from(encoded, "base64").toString()
    ```
<!-- ===JSSDK_CODEC_BASE64=== -->

=== "Python"
<!-- ===PYSDK_CODEC_BASE64=== -->
```python
encoded_str = "SGksIEknbSBkZWNvZGVkIGZyb20gYmFzZTY0"
decoded_str = base64.b64decode(encoded_str).decode("utf-8")
print(decoded_str)
```
<!-- ===PYSDK_CODEC_BASE64=== -->

=== "Go"
<!-- ===GOSDK_CODEC_BASE64=== -->
    ```go
    import "encoding/base64"

    //...

    encoded := "SGksIEknbSBkZWNvZGVkIGZyb20gYmFzZTY0"
    decoded, _ := base64.StdEncoding.DecodeString(encoded)
    ```
<!-- ===GOSDK_CODEC_BASE64=== -->

=== "Java"
<!-- ===JAVASDK_CODEC_BASE64=== -->
    ```java
    import java.util.Base64

    //...

    String encoded = "SGksIEknbSBkZWNvZGVkIGZyb20gYmFzZTY0";
    String decoded = new String(Base64.getDecoder().decode(encoded));
    ```
<!-- ===JAVASDK_CODEC_BASE64=== -->

### Integers

Integers in algorand are almost always uint64, sometimes its required to encode them as bytes. For example when passing them as application arguments in an ApplicationCallTransaction. When encoding an integer to pass as an application argument, the integer should be encoded as the big endian 8 byte representation of the integer value.

*Example:*

Given an integer `1337`, you may encode it as:

=== "JavaScript"
<!-- ===JSSDK_CODEC_UINT64=== -->
    ```js
    const val = 1337

    const encoded = algosdk.encodeUint64(val)
    const decoded = algosdk.decodeUint64(encoded)

    // val === decoded
    ```
<!-- ===JSSDK_CODEC_UINT64=== -->

=== "Python"
<!-- ===PYSDK_CODEC_UINT64=== -->
```python
val = 1337
encoded_uint = val.to_bytes(8, "big")
decoded_uint = int.from_bytes(encoded_uint, byteorder="big")
assert decoded_uint == val
```
<!-- ===PYSDK_CODEC_UINT64=== -->

=== "Go"
<!-- ===GOSDK_CODEC_UINT64=== -->
    ```go
    import "encoding/binary"

    val := 1337

    encoded := make([]byte, 8)
    binary.BigEndian.PutUint64(encoded, uint64(val)) 

    decoded := int64(binary.BigEndian.Uint64(encoded))

    // val == decoded
    ```
<!-- ===GOSDK_CODEC_UINT64=== -->

=== "Java"
<!-- ===JAVASDK_CODEC_UINT64=== -->
    ```java
    long val = 1337;

    byte[] encoded = Encoder.encodeUint64(val);

    BigInteger decoded = Encoder.decodeUint64(encoded);

    //decoded.toLong() == val
    ```
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
    ```js

    const sp = await client.getTransactionParams().do()
    const pay_txn = algosdk.makePaymentTxnWithSuggestedParams(acct1.addr, acct2.addr, 10000, undefined, undefined, sp)

    const pay_txn_bytes = algosdk.encodeObj(pay_txn.get_obj_for_encoding())
    fs.writeFileSync("pay.txn", Buffer.from(pay_txn_bytes).toString("base64"))


    const recovered_pay_txn = algosdk.decodeUnsignedTransaction(Buffer.from(fs.readFileSync("pay.txn").toString(), "base64"))
    console.log(recovered_pay_txn)

    ```
<!-- ===JSSDK_CODEC_TRANSACTION_UNSIGNED=== -->
<!-- ===JSSDK_CODEC_TRANSACTION_SIGNED=== -->
```js
    const spay_txn_bytes = pay_txn.signTxn(acct1.sk)
    fs.writeFileSync("signed_pay.txn", Buffer.from(spay_txn_bytes).toString("base64"))
    const recovered_signed_pay_txn = algosdk.decodeSignedTransaction(Buffer.from(fs.readFileSync("signed_pay.txn").toString(), "base64"))
    console.log(recovered_signed_pay_txn)
```
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
<!-- ===PYSDK_CODEC_TRANSACTION_SIGNED=== -->

=== "Go"
<!-- ===GOSDK_CODEC_TRANSACTION_UNSIGNED=== -->
```go
	// Error handling omitted for brevity
	sp, _ := client.SuggestedParams().Do(context.Background())

	pay_txn, _ := transaction.MakePaymentTxn(acct1.Address.String(), acct2.Address.String(), 10000, nil, "", sp)

	var pay_txn_bytes = make([]byte, 1e3)
	base64.StdEncoding.Encode(pay_txn_bytes, msgpack.Encode(pay_txn))
	f, _ := os.Create("pay.txn")
	f.Write(pay_txn_bytes)

    // ...

	var (
		recovered_pay_txn   = types.Transaction{}
		recovered_pay_bytes = make([]byte, 1e3)
	)
	b64_pay_bytes, _ := os.ReadFile("pay.txn")
	base64.StdEncoding.Decode(recovered_pay_bytes, b64_pay_bytes)

	msgpack.Decode(recovered_pay_bytes, &recovered_pay_txn)
	log.Printf("%+v", recovered_pay_txn)
```
<!-- ===GOSDK_CODEC_TRANSACTION_UNSIGNED=== -->

<!-- ===GOSDK_CODEC_TRANSACTION_SIGNED=== -->
    ```go

	_, spay_txn, _ := crypto.SignTransaction(acct1.PrivateKey, pay_txn)

	var spay_txn_bytes = make([]byte, 1e3)
	base64.StdEncoding.Encode(spay_txn_bytes, spay_txn)
	f2, _ := os.Create("signed_pay.txn")
	f2.Write(spay_txn_bytes)


	var (
		recovered_signed_pay_txn   = types.SignedTxn{}
		recovered_signed_pay_bytes = make([]byte, 1e3)
	)
	b64_signed_pay_bytes, _ := os.ReadFile("signed_pay.txn")
	base64.StdEncoding.Decode(recovered_signed_pay_bytes, b64_signed_pay_bytes)

	msgpack.Decode(recovered_signed_pay_bytes, &recovered_signed_pay_txn)
	log.Printf("%+v", recovered_signed_pay_txn)
    ```
<!-- ===GOSDK_CODEC_TRANSACTION_SIGNED=== -->

=== "Java"
<!-- ===JAVASDK_CODEC_TRANSACTION_UNSIGNED=== -->
<!-- ===JAVASDK_CODEC_TRANSACTION_UNSIGNED=== -->
<!-- ===JAVASDK_CODEC_TRANSACTION_SIGNED=== -->
<!-- ===JAVASDK_CODEC_TRANSACTION_SIGNED=== -->



### Blocks

One type that commonly needs to be decoded are the blocks themselves. Since some fields are raw byte arrays (like state deltas) the msgpack format should be used.

=== "Python"
<!-- ===PYSDK_CODEC_BLOCK=== -->
```python
```
<!-- ===PYSDK_CODEC_BLOCK=== -->

=== "JavaScript"
<!-- ===JSSDK_CODEC_BLOCK=== -->
<!-- ===JSSDK_CODEC_BLOCK=== -->

=== "Go"
<!-- ===GOSDK_CODEC_BLOCK=== -->
<!-- ===GOSDK_CODEC_BLOCK=== -->

=== "Java"
<!-- ===JAVASDK_CODEC_BLOCK=== -->
<!-- ===JAVASDK_CODEC_BLOCK=== -->