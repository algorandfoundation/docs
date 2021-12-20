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

In Algorand a [public key](./accounts/#transformation-public-key-to-algorand-address) is a 32 byte array. 

The Address developers or users are typically showin is a 58 byte string corresponding to a base32 encoding of the byte array of the public key + a checksum.

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

Sometimes an application needs to transmit a transaction or transaction group between the front end and back end.

This can be done by msgpack encoding the transaction object on one side and msgpack decoding it on the other side.

### Signed Transaction

After making a request for a transaction to be signed by a wallet, the returned data structure is 

```ts
{
    TxId string // base32 encoded transaction id
    blob string // base64 encoded byte array
}
```

Sometimes you may want to inspect the transaction to, for example, validate that it was signed by the correct public key as in a web3 login.

The blob can be decoded with the same base64 decoding of a [byte array](#byte-arrays).  From the resulting byte array, an object may representing the signed transaction may be created using the msgpack decode method for each sdk.

From there you may inspect the signature field to check the transaction was signed correctly.

*Example:*

Given a blob ``, you may decode it as:

=== "JavaScript"
    ```js
    import algosdk from 'algosdk'

    const blob = ""
    const decoded =  Buffer.from(blob, "base64")
    const signedTxn = algosdk.decodeSignedTransaction(decoded)
    ```

=== "Python"
    ```python
    blob = ""
    signedTxn = encoding.future_msgpack_decode(blob)
    ```