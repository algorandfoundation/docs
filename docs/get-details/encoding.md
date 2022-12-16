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
    ```js
    import algosdk from 'algosdk'

    const address = "4H5UNRBJ2Q6JENAXQ6HNTGKLKINP4J4VTQBEPK5F3I6RDICMZBPGNH6KD4"
    const pk = algosdk.decodeAddress(address)
    const addr = algosdk.encodeAddress(pk.publicKey)
    // addr === address
    ```

=== "Python"
    ```python
    from algosdk.encoding import decode_address, encode_address

    address = "4H5UNRBJ2Q6JENAXQ6HNTGKLKINP4J4VTQBEPK5F3I6RDICMZBPGNH6KD4"
    pk = decode_address(address)
    addr = encode_address(pk)

    assert addr == address
    ```

=== "Go"
    ```go
    import "github.com/algorand/go-algorand-sdk/types"

    //...

    address := "4H5UNRBJ2Q6JENAXQ6HNTGKLKINP4J4VTQBEPK5F3I6RDICMZBPGNH6KD4"
    pk, _ := types.DecodeAddress(address)
    addr := pk.String()
    
    //addr == address
    ```

=== "Java"
    ```java
    import com.algorand.algosdk.crypto.Address;

    //...

    String address = "4H5UNRBJ2Q6JENAXQ6HNTGKLKINP4J4VTQBEPK5F3I6RDICMZBPGNH6KD4";
    Address pk = new Address(address);
    String addr = new Address(pk.getBytes());

    // addr == address
    ```

### Byte Arrays

When transmitting an array of bytes over the network, byte arrays are base64 encoded.  The SDK will handle encoding from a byte array to base64 but may not decode some fields and you'll have to handle it yourself. For example compiled program results or the keys and values in a state delta in an application call will be returned as base64 encoded strings.

*Example:*

Given a base64 encoded byte array `SGksIEknbSBkZWNvZGVkIGZyb20gYmFzZTY0` it may be decoded as follows:

=== "JavaScript"
    ```js
    const encoded = "SGksIEknbSBkZWNvZGVkIGZyb20gYmFzZTY0"
    const decoded = Buffer.from(encoded, "base64").toString()
    ```

=== "Python"
    ```python
    import base64

    encoded = "SGksIEknbSBkZWNvZGVkIGZyb20gYmFzZTY0"
    decoded = base64.b64decode(encoded).encode('ascii')
    ```

=== "Go"
    ```go
    import "encoding/base64"

    //...

    encoded := "SGksIEknbSBkZWNvZGVkIGZyb20gYmFzZTY0"
    decoded, _ := base64.StdEncoding.DecodeString(encoded)
    ```

=== "Java"
    ```java
    import java.util.Base64

    //...

    String encoded = "SGksIEknbSBkZWNvZGVkIGZyb20gYmFzZTY0";
    String decoded = new String(Base64.getDecoder().decode(encoded));
    ```

### Integers

Integers in algorand are almost always uint64, sometimes its required to encode them as bytes. For example when passing them as application arguments in an ApplicationCallTransaction. When encoding an integer to pass as an application argument, the integer should be encoded as the big endian 8 byte representation of the integer value.

*Example:*

Given an integer `1337`, you may encode it as:

=== "JavaScript"
    ```js
    const val = 1337

    const encoded = algosdk.encodeUint64(val)
    const decoded = algosdk.decodeUint64(encoded)

    // val === decoded
    ```

=== "Python"
    ```python
    val = 1337

    encoded = (val).to_bytes(8,'big')
    decoded = int.from_bytes(encoded, byteorder='big)

    assert decoded == val
    ```

=== "Go"
    ```go
    import "encoding/binary"

    val := 1337

    encoded := make([]byte, 8)
    binary.BigEndian.PutUint64(encoded, uint64(val)) 

    decoded := int64(binary.BigEndian.Uint64(encoded))

    // val == decoded
    ```

=== "Java"
    ```java
    long val = 1337;

    byte[] encoded = Encoder.encodeUint64(val);

    BigInteger decoded = Encoder.decodeUint64(encoded);

    //decoded.toLong() == val
    ```


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
    ```js

    const sp = await client.getTransactionParams().do()
    const pay_txn = algosdk.makePaymentTxnWithSuggestedParams(acct1.addr, acct2.addr, 10000, undefined, undefined, sp)

    const pay_txn_bytes = algosdk.encodeObj(pay_txn.get_obj_for_encoding())
    fs.writeFileSync("pay.txn", Buffer.from(pay_txn_bytes).toString("base64"))

    const spay_txn_bytes = pay_txn.signTxn(acct1.sk)
    fs.writeFileSync("signed_pay.txn", Buffer.from(spay_txn_bytes).toString("base64"))


    const recovered_pay_txn = algosdk.decodeUnsignedTransaction(Buffer.from(fs.readFileSync("pay.txn").toString(), "base64"))
    console.log(recovered_pay_txn)

    const recovered_signed_pay_txn = algosdk.decodeSignedTransaction(Buffer.from(fs.readFileSync("signed_pay.txn").toString(), "base64"))
    console.log(recovered_signed_pay_txn)

    ```

=== "Python"
    ```py

    sp = client.suggested_params()
    pay_txn = transaction.PaymentTxn(addr1, sp, addr2, 10000)
    with open("pay.txn", "w") as f:
        f.write(encoding.msgpack_encode(pay_txn))

    spay_txn = pay_txn.sign(pk1)
    with open("signed_pay.txn", "w") as f:
        f.write(encoding.msgpack_encode(spay_txn))


    with open("pay.txn", "r") as f:
        recovered_txn = encoding.future_msgpack_decode(f.read())

    print(recovered_txn)

    with open("signed_pay.txn", "r") as f:
        recovered_signed_txn = encoding.future_msgpack_decode(f.read())

    print(recovered_signed_txn)

    ```

=== "Go"
    ```go

	// Error handling omitted for brevity
	sp, _ := client.SuggestedParams().Do(context.Background())

	pay_txn, _ := transaction.MakePaymentTxn(acct1.Address.String(), acct2.Address.String(), 10000, nil, "", sp)

	var pay_txn_bytes = make([]byte, 1e3)
	base64.StdEncoding.Encode(pay_txn_bytes, msgpack.Encode(pay_txn))
	f, _ := os.Create("pay.txn")
	f.Write(pay_txn_bytes)

	_, spay_txn, _ := crypto.SignTransaction(acct1.PrivateKey, pay_txn)

	var spay_txn_bytes = make([]byte, 1e3)
	base64.StdEncoding.Encode(spay_txn_bytes, spay_txn)
	f2, _ := os.Create("signed_pay.txn")
	f2.Write(spay_txn_bytes)

	var (
		recovered_pay_txn   = types.Transaction{}
		recovered_pay_bytes = make([]byte, 1e3)
	)
	b64_pay_bytes, _ := os.ReadFile("pay.txn")
	base64.StdEncoding.Decode(recovered_pay_bytes, b64_pay_bytes)

	msgpack.Decode(recovered_pay_bytes, &recovered_pay_txn)
	log.Printf("%+v", recovered_pay_txn)

	var (
		recovered_signed_pay_txn   = types.SignedTxn{}
		recovered_signed_pay_bytes = make([]byte, 1e3)
	)
	b64_signed_pay_bytes, _ := os.ReadFile("signed_pay.txn")
	base64.StdEncoding.Decode(recovered_signed_pay_bytes, b64_signed_pay_bytes)

	msgpack.Decode(recovered_signed_pay_bytes, &recovered_signed_pay_txn)
	log.Printf("%+v", recovered_signed_pay_txn)
    ```



### Blocks

One type that commonly needs to be decoded are the blocks themselves. Since some fields are raw byte arrays (like state deltas) the msgpack format should be used.

=== "Python"
    ```py

    from algosdk import encoding
    from algosdk.v2client.algod import AlgodClient
    import msgpack

    client = AlgodClient("a"*64, "http://localhost:4001")

    # Get the block in msgpack format
    block = client.block_info(round_num=1234, response_format='msgpack')

    # Be sure to specify `raw=True` or msgpack will try to decode as utf8
    res = msgpack.unpackb(block, raw=True)

    # Lets get the 4th transaction
    txn = res[b'block'][b'txns'][3]

    # Grab the byte value for the key `KEY`
    address = encoding.encode_address(txn[b'dt'][b'gd'][b'KEY'][b'bs'])

    ```