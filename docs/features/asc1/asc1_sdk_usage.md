This guide covers using TEAL programs with contract acccounts or delegated signatures with the available SDKs. The methods covered in this documentation are used for custom TEAL code and provide general access to any TEAL program. Algorand Smart Contract Templates are also available for officially supported TEAL programs for common use case functionality like Hash Time-Lock Contracts, Split Payments, Limit Orders, etc. These templates can be used with specific SDK APIs that are detailed in the Templates documenation.<LINK>

Each SDK's install process is discussed in the SDK Reference documenation.<LINK>

Code snippets are abbreviated for conciseness and clarity. See the full code example for each SDK at the bottom of this guide.

# Accessing TEAL program from SDKs
Before a TEAL program can be used is the SDKs, it must be compiled using the `goal` tool. The goal teal walkthrough documenation <LINK> explains this process. Once a TEAL program is compiled, the bytes of the program can be retrieved in various ways. Most of the SDKs support the bytes encoded in base64 or in hexadecimal format. The following example illustrates using shell commands to export the binary to hexidecimal or a base64 encoded string.

``` bash
//simple.teal contains int=0
//hexdump 
$ hexdump -C simple.teal.tok
00000000  01 20 01 00 22                                    |. .."|
00000005
//base64 format
$ cat simple.teal.tok | base64
ASABACI=
```

The converted binary bytes are used in the SDKs as shown below.

```javascript tab="JavaScript"
let program = new Uint8Array(Buffer.from("ASABACI=", "base64"));
```

```python tab="Python"
program = b"\x01\x20\x01\x00\x22"  
```

```java tab="Java"
byte[] program = {
    0x01, 0x20, 0x01, 0x00, 0x22  // int 0
};
```

```go tab="Go"
program, err :=  base64.StdEncoding.DecodeString("ASABACI=")
```
# Contract Account SDK usage
ASC1 Contract accounts are used to allow TEAL logic to determine when outgoing account transactions are approved. The compiled TEAL program produces an Algorand Address, which is funded with Algos or Algorand Assets. As the reciever of a transaction, these accounts function as an other account. When the account is specified as the sender in a transaction, the TEAL logic is evaluated and determines if the transaction is appoved. The ASC1 Usage Modes documentation<LINK> explains ASC1 modes in more detail. 

TEAL contract account transactions where the sender is set to the contract account, function much in the same way as normal Algorand transactions<Link>. The major difference is that instead of the transaction being signed with a private key, the transaction is signed with a logic signature<LINK>. See Transaction documentaiton for details on setting up a payment transaction.<Link>

Contract Accounts are created by compiling the TEAL logic. Once the contract account is created, it can be used as any other address. To send tokens or assets from the account the transaction must be signed by a Logic Signature. From an SDK standpoint, the following process should be used.

* Load the Program Bytes into the SDK.
* Create a Logic Signature based on the program.
* Create the Transaction.
* Set the `from` transaction property to the contract address.
* Sign the Transaction with the Logic Signature.
* Send the Transaction to the network.

<center>![Transaction From Contract Account](../../imgs/asc1_sdk_usage-1.png)</center>
<center>*Transaction From Contract Account*</center>

The following example illustrates signing a transaction with a created logic signature.

```javascript tab="JavaScript"
    let program = new Uint8Array(Buffer.from("ASABACI=", "base64"));
    let lsig = algosdk.makeLogicSig(program);
    let txn = {
        "from": [CONTRACT ADDRESS],
        "to": [RECEIVER],
        "fee": [FEE],
        "amount": [AMOUNT],
        "firstRound": [FIRST VALID ROUND],
        "lastRound": [LAST VALID ROUND],
        "genesisID": [GENESIS ID],
        "genesisHash": [GENESIS HASH]
    };
    let rawSignedTxn = algosdk.signLogicSigTransaction(txn, lsig);
    let tx = (await algodclient.sendRawTransaction(rawSignedTxn.blob));
```

```python tab="Python"
	program = b"\x01\x20\x01\x00\x22"  
	lsig = transaction.LogicSig(program)
	txn = transaction.PaymentTxn([CONTRACT ADDRESS], [FEE], [FIRST VALID ROUND], [LAST VALID ROUND], [GENESIS HASH], [RECEIVER], [AMOUNT], [CLOSE REMAINDER TO ADDRESS])
	lstx = transaction.LogicSigTransaction(txn, lsig)
	txid = acl.send_transaction(lstx)
```

```java tab="Java"
    byte[] program = {
        0x01, 0x20, 0x01, 0x00, 0x22  // int 0
    };
    LogicsigSignature lsig = new LogicsigSignature(program, null);
    Transaction tx = new Transaction(lsig.toAddress(), new Address([RECIEVER ADDRESS]), BigInteger.valueOf([FEE]), [AMOUNT], [FIRST VALID ROUND], [LAST VALID ROUND], [GENESIS ID], [GENESIS HASH]);
    try {
        SignedTransaction stx = Account.signLogicsigTransaction(lsig, tx);
        byte[] encodedTxBytes = Encoder.encodeToMsgPack(stx);
        TransactionID id = algodApiInstance.rawTransaction(encodedTxBytes);
    } catch (ApiException e) {
		...
    }
```

```go tab="Go"
    program, err :=  base64.StdEncoding.DecodeString("ASABACI=")
    var args [][]byte
    lsig, err := crypto.MakeLogicSig(contract.program, nil, nil, nil)
	addr := crypto.LogicSigAddress(lsig).String()	
	tx, err := transaction.MakePaymentTxnWithFlatFee(
		addr, [RECEIVER], [FEE], [AMOUNT], [FIRST VALID ROUND], [LAST VALID ROUND],
		[NOTE], [CLOSE REAMAINDER TO ADDRESS], [GENESIS ID], [GENESIS HASH] )
	txid, stx, err := crypto.SignLogicsigTransaction(lsig, tx)
	transactionID, err := algodClient.SendRawTransaction(stx)
```

# Account Delegation SDK Usage
ASC1 allows TEAL logic to be used to delegate signature authority. This allows specific accounts or multi-signature accounts to sign logic that allows transactions from the account to be approved based on the TEAL logic. The ASC1 Usage Modes documentation<LINK> explains ASC1 modes in more detail. 

Delegation transactions where the sender is set to the account that signed the specific logic, function much in the same way as normal Algorand transactions<Link>. The major difference is that instead of the transaction being signed with a private key, the transaction is signed with a logic signature<LINK>. See Transaction documentaiton for details on setting up a payment transaction.<Link>

Delgated Logic Signatures require that the logic signature be signed from a specific account or a multi-signature account. The TEAL program is first loaded, then a Logic Signature is created and then the Logic Signature is signed by a specific account or multi-signaure account. The transaction is created as normal. The transaction is then signed with the Logic Signature. From an SDK standpoint, the following process should be used.

* Load the Program Bytes into the SDK.
* Create a Logic Signature based on the program.
* Sign The Logic Signature with a specific account
* Create the Transaction.
* Set the `from` transaction property to the Address that signed the logic.
* Sign the Transaction with the Logic Signature.
* Send the Transaction to the network.

<center>![Delegated Signature Transaction](../../imgs/asc1_sdk_usage-2.png)</center>
<center>*Delegated Signature Transactiont*</center>

The following example illustrates signing a transaction with a created logic signature.

```javascript tab="JavaScript"
    let program = new Uint8Array(Buffer.from("ASABACI=", "base64"));
    let lsig = algosdk.makeLogicSig(program);
    lsig.sign([SECRET KEY]);
    let txn = {
        "from": [CONTRACT ADDRESS],
        "to": [RECEIVER],
        "fee": [FEE],
        "amount": [AMOUNT],
        "firstRound": [FIRST VALID ROUND],
        "lastRound": [LAST VALID ROUND],
        "genesisID": [GENESIS ID],
        "genesisHash": [GENESIS HASH]
    };
    let rawSignedTxn = algosdk.signLogicSigTransaction(txn, lsig);
    let tx = (await algodclient.sendRawTransaction(rawSignedTxn.blob));
```

```python tab="Python"
	program = b"\x01\x20\x01\x00\x22"  
	lsig = transaction.LogicSig(program)
    lsig.sign([SECRET KEY])
	txn = transaction.PaymentTxn([CONTRACT ADDRESS], [FEE], [FIRST VALID ROUND], [LAST VALID ROUND], [GENESIS HASH], [RECEIVER], [AMOUNT], [CLOSE REMAINDER TO ADDRESS])
	lstx = transaction.LogicSigTransaction(txn, lsig)
	txid = acl.send_transaction(lstx)
```

```java tab="Java"
    byte[] program = {
        0x01, 0x20, 0x01, 0x00, 0x22  // int 0
    };
    LogicsigSignature lsig = new LogicsigSignature(program, null);
    Account account = new Account([PASSPHARASE]);
    account.signLogicsig(lsig);
    Transaction tx = new Transaction(lsig.toAddress(), new Address([RECIEVER ADDRESS]), BigInteger.valueOf([FEE]), [AMOUNT], [FIRST VALID ROUND], [LAST VALID ROUND], [GENESIS ID], [GENESIS HASH]);
    try {
        SignedTransaction stx = Account.signLogicsigTransaction(lsig, tx);
        byte[] encodedTxBytes = Encoder.encodeToMsgPack(stx);
        TransactionID id = algodApiInstance.rawTransaction(encodedTxBytes);
    } catch (ApiException e) {
		...
    }
```

```go tab="Go"
    program, err :=  base64.StdEncoding.DecodeString("ASABACI=")
	var args [][]byte
	lsig, err := crypto.MakeLogicSig(program, args, [SECRET KEY], nil)
	addr := crypto.LogicSigAddress(lsig).String()	
	tx, err := transaction.MakePaymentTxnWithFlatFee(
		addr, [RECEIVER], [FEE], [AMOUNT], [FIRST VALID ROUND], [LAST VALID ROUND],
		[NOTE], [CLOSE REAMAINDER TO ADDRESS], [GENESIS ID], [GENESIS HASH] )
	txid, stx, err := crypto.SignLogicsigTransaction(lsig, tx)
	transactionID, err := algodClient.SendRawTransaction(stx)
```

# Save Transaction Output for Debugging

# Passing Parameters using the SDKs
In SDKs you can use the native language functions. For example to pass in a string argument to the Python SDK you would use code similar to the example below.
arg_str = "my string"
arg1 = arg_str.encode()
lsig = transaction.LogicSig(program, args=[arg1])
For the Integer argument you would use (remember TEAL only supports positive values):
arg1 = (123).to_bytes(8, 'big')
lsig = transaction.LogicSig(program, args=[arg1])