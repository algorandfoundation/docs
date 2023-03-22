title: ABI details


The ABI (Application Binary Interface) is a specification that defines the encoding/decoding of data types and a standard for exposing and invoking methods in a smart contract.

The specification is defined in [ARC4](https://github.com/algorandfoundation/ARCs/blob/main/ARCs/arc-0004.md).

At a high level, the ABI allows contracts to define an API with rich types and offer an interface description so clients know exactly what the contract is expecting to be passed. 

To encode or decode ABI types with the SDKs, see the [ABI encoding and decoding](../../../encoding.md#abi-encoding) documentation.

!!! note
    When constructing Application Call transactions that are expected to encode arguments, using the [Atomic Transaction Composer](../../../atc.md) is recommended.

# Data Types

Encoding for the data types is specified [here](https://github.com/algorandfoundation/ARCs/blob/main/ARCs/arc-0004.md#encoding).

The data types specified are:

| Type | Description | 
| ---- | ----------- |
|uintN| An N-bit unsigned integer, where 8 <= N <= 512 and N % 8 = 0|
|byte| An alias for uint8|
|bool| A boolean value that is restricted to either 0 or 1. When encoded, up to 8 consecutive bool values will be packed into a single byte|
|ufixedNxM| An N-bit unsigned fixed-point decimal number with precision M, where 8 <= N <= 512, N % 8 = 0, and 0 < M <= 160, which denotes a value v as v / (10^M)|
|type[N]| A fixed-length array of length N, where N >= 0. type can be any other type|
|address| Used to represent a 32-byte Algorand address. This is equivalent to byte[32]|
|type[]| A variable-length array. type can be any other type|
|string| A variable-length byte array (byte[]) assumed to contain UTF-8 encoded content|
|(T1,T2,...,TN)| A tuple of the types T1, T2, …, TN, N >= 0|
|reference type | account, asset, application only for arguments, in which case they are an alias for uint8. See section "Reference Types" below|

The encoding and decoding of these types should be handled by the SDKs for calling methods and reading out return values.  

Because stack types in the AVM are limited to uint64 and bytes, a smart contract may rely on ABI types defined in PyTeal to encode or decode the data passed in the application args. 

## Reference Types

Reference types may be specified in the method signature referring to some transaction parameters that must be passed.  The value encoded is a uint8 reference to the index of element in the relevant array (i.e. for account, the index in the foreign accounts array).


# Methods

Methods may be exposed by the smart contract and called by submitting an ApplicationCall transaction to the existing application id. 

A *method signature* is defined as a name, argument types, and return type. The stringified version is then hashed and the first 4 bytes are taken as a *method selector*.

For example:

A *method signature* for an `add` method that takes 2 uint64s and returns 1 uint128:
```
Method signature: add(uint64,uint64)uint128
```

The string version of the *method signature* is hashed and the first 4 bytes are its *method selector*:
```
SHA-512/256 hash (in hex): 8aa3b61f0f1965c3a1cbfa91d46b24e54c67270184ff89dc114e877b1753254a
Method selector (in hex): 8aa3b61f
```

Once the method selector is known, it is used in the smart contract logic to route to the appropriate logic that implements the `add` method. 

The `method` pseudo-opcode can be used in a contract to do the above work and produce a *method selector* given the *method signature* string.

```
method "add(uint64,uint64)uint128"
```

## Implementing a method

[Implementing a method](https://github.com/algorandfoundation/ARCs/blob/main/ARCs/arc-0004.md#implementing-a-method) is done by handling an ApplicationCall transaction where the first element matches its method selector and the subsequent elements are used by the logic in the method body.

The initial handling logic of the contract should route to the correct method given a match against the method selector passed and the known method selector of the application method.

The return value of the method _must_ be logged with the prefix `151f7c75` which is the result of `sha256("return")[:4]`.  Only the last logged element with this prefix is considered the return value of this method call.


# API

The API of a smart contract can be published as an [interface description object](https://github.com/algorandfoundation/ARCs/blob/main/ARCs/arc-0004.md#interface-description). A user may read this object and instantiate a client that handles the encoding/decoding of the arguments and returns values using one of the SDKs.


A full example of a contract json file might look like:
```json
{
    "name":"super-awesome-contract",
    "networks":{
        "MainNet":{
            "appID": 123456
        }
    },
    "methods":[
        {
            "name":"add",
            "desc":"Add 2 integers",
            "args":[ { "type":"uint64" }, { "type":"uint64" } ],
            "returns": {"type":"uint64"}
        },
        {
            "name":"sub",
            "desc":"Subtract 2 integers",
            "args":[ { "type":"uint64" }, { "type":"uint64" } ],
            "returns": {"type":"uint64"}
        },
        {
            "name":"mul",
            "desc":"Multiply 2 integers",
            "args":[ { "type":"uint64" }, { "type":"uint64" } ],
            "returns": {"type":"uint64"}
        },
        {
            "name":"div",
            "desc":"Divide 2 integers, throw away the remainder",
            "args":[ { "type":"uint64" }, { "type":"uint64" } ],
            "returns": {"type":"uint64"}
        },
        {
            "name":"qrem",
            "desc":"Divide 2 integers, return both the quotient and remainder",
            "args":[ { "type":"uint64" }, { "type":"uint64" } ],
            "returns": {"type":"(uint64,uint64)"}
        },
        {
            "name":"reverse",
            "desc":"Reverses a string",
            "args":[ { "type":"string" } ],
            "returns": {"type":"string"}
        },
        {
            "name":"txntest",
            "desc":"just check it",
            "args":[{"type":"uint64"}, {"type": "pay"}, {"type":"uint64"}],
            "returns":{"type": "uint64"}
        },
        {
            "name":"concat_strings",
            "desc":"concat some strings",
            "args":[{"type":"string[]"}],
            "returns":{"type": "string"}
        },
        {
            "name":"manyargs",
            "desc":"Try to send 20 arguments",
            "args":[
                {"type":"uint64"}, {"type": "uint64"}, {"type":"uint64"},{"type":"uint64"},
                {"type":"uint64"}, {"type": "uint64"}, {"type":"uint64"},{"type":"uint64"},
                {"type":"uint64"}, {"type": "uint64"}, {"type":"uint64"},{"type":"uint64"},
                {"type":"uint64"}, {"type": "uint64"}, {"type":"uint64"},{"type":"uint64"},
                {"type":"uint64"}, {"type": "uint64"}, {"type":"uint64"},{"type":"uint64"}
            ],
            "returns":{"type": "uint64"}
        },
        {
            "name":"min_bal",
            "desc":"Get the minimum balance for given account",
            "args":[
                {"type":"account"}
            ],
            "returns":{"type":"uint64"}
        },
        {
            "name":"tupler",
            "desc":"",
            "args":[
                {"type":"(string,uint64,string)"}
            ],
            "returns":{"type":"uint64"}
        }
    ]
}
```
