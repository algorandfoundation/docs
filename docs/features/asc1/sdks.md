title: Using the SDKs

This guide covers using TEAL programs with contract accounts or delegated signatures with the available SDKs. The methods covered in this documentation are used for custom TEAL code and provide general access to any TEAL program. Algorand Smart Contract Templates are also available for common use case functionality like Hash Time-Lock Contracts, Split Payments, Limit Orders, etc. Developer documentation describing the process for using these templates will be available soon.

Each SDK's install process is discussed in the [SDK Reference](../../reference/sdks/index.md) documentation.

!!! info
    The example code snippets are provided throughout this page and are abbreviated for conciseness and clarity. Full running code examples for each SDK are available within the GitHub repo for V1 and V2 at [/examples/smart_contracts](https://github.com/algorand/docs/tree/master/examples/smart_contracts) and for [download](https://github.com/algorand/docs/blob/master/examples/smart_contracts/smart_contracts.zip?raw=true) (.zip).


# Accessing TEAL program from SDKs
Before a TEAL program can be used is the SDKs, it must be compiled using the `goal` tool. The [goal TEAL walkthrough](./goal_teal_walkthrough.md) documentation explains this process. Once a TEAL program is compiled, the bytes of the program can be retrieved in various ways. Most of the SDKs support the bytes encoded in base64 or hexadecimal format. The following example illustrates using shell commands to export the binary to hexadecimal or a base64 encoded string.

``` bash
//simple.teal contains int 0
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
ASC1 Contract accounts are used to allow TEAL logic to determine when outgoing account transactions are approved. The compiled TEAL program produces an Algorand Address, which is funded with Algos or Algorand Assets. As the receiver of a transaction, these accounts function as any other account. When the account is specified as the sender in a transaction, the TEAL logic is evaluated and determines if the transaction is approved. The [ASC1 Usage Modes](./modes.md) documentation explains ASC1 modes in more detail. 

TEAL contract account transactions where the sender is set to the contract account, function much in the same way as normal Algorand [transactions](../transactions/index.md). The major difference is that instead of the transaction being signed with a private key, the transaction is signed with a [logic signature](./modes.md#logic-signatures). See [Transaction](../transactions/index.md) documentation for details on setting up a payment transaction.

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
// get suggested parameters
let params = await algodclient.getTransactionParams().do();
// comment out the next two lines to use suggested fee
params.fee = 1000;
params.flatFee = true;
console.log(params);
// create logic sig
// b64 example "ASABACI=" is `int 0`
// see more info here: https://developer.algorand.org/docs/features/asc1/sdks/#accessing-teal-program-from-sdks
let program = new Uint8Array(Buffer.from("base64-encoded-program" < PLACEHOLDER >, "base64"));
// let program = new Uint8Array(Buffer.from("ASABACI=" , "base64"));

let lsig = algosdk.makeLogicSig(program);

// create a transaction
let sender = lsig.address();
let receiver = "< PLACEHOLDER >";
let amount = 10000;
let closeToRemaninder = undefined;
let note = undefined;
let txn = algosdk.makePaymentTxnWithSuggestedParams(sender, receiver, amount, closeToRemaninder, note, params)

// Create the LogicSigTransaction with contract account LogicSig 
let rawSignedTxn = algosdk.signLogicSigTransactionObject(txn, lsig);

// send raw LogicSigTransaction to network
let tx = (await algodclient.sendRawTransaction(rawSignedTxn.blob).do());
console.log("Transaction : " + tx.txId);   
await waitForConfirmation(algodclient, tx.txId);
```

```python tab="Python"
    # create logic sig
    # program = b"hex-encoded-program"
    # b"\x01\x20\x01\x00\x22 is `int 0`
    # see more info here: https://developer.algorand.org/docs/features/asc1/sdks/#accessing-teal-program-from-sdks
    program = b"\x01\x20\x01\x00\x22"
    lsig = LogicSig(program)
    sender = lsig.address()
    
    # get suggested parameters
    params = algod_client.suggested_params()
    # comment out the next two (2) lines to use suggested fees
    params.flat_fee = True
    params.fee = 1000
    
    # build transaction  
    amount = amount < PLACEHOLDER >
    closeremainderto = None
    receiver = < PLACEHOLDER > 
    # create a transaction
    txn = PaymentTxn(
        sender, params, receiver, amount, closeremainderto)

    # Create the LogicSigTransaction with contract account LogicSig
    lstx = transaction.LogicSigTransaction(txn, lsig)

    # send raw LogicSigTransaction to network
    print("This transaction is expected to fail as it is int 0 , always false")
    txid = algod_client.send_transaction(lstx)
    print("Transaction ID: " + txid)    
    wait_for_confirmation(algod_client, txid) 
```

```java tab="Java"
    public void contractAccountExample() throws Exception {
        // Initialize an algod client
        if (client == null)
            this.client = connectToNetwork();
        // import your private key mnemonic and address
        
        // Set the receiver
        final String RECEIVER = "< receiver PLACEHOLDER >";
        // create logic sig
        // hex example 0x01, 0x20, 0x01, 0x00, 0x22 int 0 returns false, so rawTransaction will fail below
        byte[] program = { 0x01, 0x20, 0x01, 0x00, 0x22 };
             
        LogicsigSignature lsig = new LogicsigSignature(program, null);
        System.out.println("lsig address: " + lsig.toAddress());
        TransactionParametersResponse params = client.TransactionParams().execute().body();
        // create a transaction

        String note = "Hello World";
        Transaction txn = Transaction.PaymentTransactionBuilder()
                .sender(lsig
                        .toAddress())
                .note(note.getBytes())
                .amount(<PLACEHOLDER>)
                .receiver(new Address(RECEIVER))
                .suggestedParams(params)
                .build();   

        try {
            // create the LogicSigTransaction with contract account LogicSig
            SignedTransaction stx = Account.signLogicsigTransaction(lsig, txn);

            // send raw LogicSigTransaction to network
            byte[] encodedTxBytes = Encoder.encodeToMsgPack(stx);
            Response<PostTransactionsResponse> rp = client.RawTransaction().rawtxn(encodedTxBytes).execute();
            String id = null;
            if (rp.body() != null) {
                id = rp.body().txId;
            } else {
                System.out.println(rp.message());
            }
            // Wait for transaction confirmation
            waitForConfirmation(id);
            System.out.println("Successfully sent tx with id: " + id);
            // Read the transaction
            PendingTransactionResponse pTrx = client.PendingTransactionInformation(id).execute().body();
  
            JSONObject jsonObj = new JSONObject(pTrx.toString());
            System.out.println("Transaction information (with notes): " + jsonObj.toString(2)); // pretty print
            System.out.println("Decoded note: " + new String(pTrx.txn.tx.note));
        } catch (ApiException e) {
            System.err.println("Exception when calling algod#rawTransaction: " + e.getResponseBody());
        }
    }
```

```go tab="Go"
	// Create logic signature
	// example base64 encoded program "ASABACI=" int 0
	var sk ed25519.PrivateKey
	var ma crypto.MultisigAccount
	program, err :=  base64.StdEncoding.DecodeString("base64-encoded-program<PLACEHOLDER>")

	// program, err := base64.StdEncoding.DecodeString("ASABACI=")
	var args [][]byte
	lsig, err := crypto.MakeLogicSig(program, args, sk, ma)
	addr := crypto.LogicSigAddress(lsig).String()
	fmt.Printf("Escrow Address: %s\n", addr)

	// Create an algod client
	algodClient, err := algod.MakeClient(algodAddress, algodToken)
	if err != nil {
		fmt.Printf("failed to make algod client: %s\n", err)
		return
	}
	// Get suggested params for the transaction
	txParams, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	txParams.FlatFee = true
	txParams.Fee = 1000

	// Make transaction
	const receiver = "transaction-receiver"<PLACEHOLDER>
	const fee = fee<PLACEHOLDER>
	const amount = amount<PLACEHOLDER>

	var minFee uint64 = 1000
	note := []byte("Hello World")
	genID := txParams.GenesisID
	genHash := txParams.GenesisHash
	firstValidRound := uint64(txParams.FirstRoundValid)
	lastValidRound := uint64(txParams.LastRoundValid)
	tx, err := transaction.MakePaymentTxnWithFlatFee(
		addr, receiver, minFee, amount, firstValidRound, lastValidRound, note, "", genID, genHash)

	txID, stx, err := crypto.SignLogicsigTransaction(lsig, tx)
	if err != nil {
		fmt.Printf("Signing failed with %v", err)
		return
	}
	fmt.Printf("Signed tx: %v\n", txID)

	// Submit the raw transaction to network
	fmt.Printf("expected to fail with int 0 program, always false %s\n", err)
	transactionID, err := algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		fmt.Printf("Sending failed with %v\n", err)
	}
	fmt.Printf("Transaction ID: %v\n", transactionID)

	// Read transaction
	confirmedTxn, stxn, err := algodClient.PendingTransactionInformation(txID).Do(context.Background())
	if err != nil {
		fmt.Printf("Error retrieving transaction %s\n", txID)
		return
	}
	txnJSON, err := json.MarshalIndent(confirmedTxn.Transaction.Txn, "", "\t")
	if err != nil {
		fmt.Printf("Cannot marshal txn data: %s\n", err)
	}
	fmt.Printf("Transaction information: %s\n", txnJSON)
	fmt.Printf("Decoded note: %s\n", string(stxn.Txn.Note))
```

# Account Delegation SDK Usage
ASC1 allows TEAL logic to be used to delegate signature authority. This allows specific accounts or multi-signature accounts to sign logic that allows transactions from the account to be approved based on the TEAL logic. The [ASC1 Usage Modes](./modes.md) documentation explains ASC1 modes in more detail. 

Delegated transactions are special transactions where the `sender` also signs the logic and the transaction is then signed with the [logic signature](./modes.md#logic-signature). In all other aspects, the transaction functions as any other transaction. See [Transaction](../transactions/index.md) documentation for details on setting up a payment transaction.

Delegated Logic Signatures require that the logic signature be signed from a specific account or a multi-signature account. The TEAL program is first loaded, then a Logic Signature is created and then the Logic Signature is signed by a specific account or multi-signature account. The transaction is created as normal. The transaction is then signed with the Logic Signature. From an SDK standpoint, the following process should be used.

* Load the Program Bytes into the SDK.
* Create a Logic Signature based on the program.
* Sign The Logic Signature with a specific account
* Create the Transaction.
* Set the `from` transaction property to the Address that signed the logic.
* Sign the Transaction with the Logic Signature.
* Send the Transaction to the network.

<center>![Delegated Signature Transaction](../../imgs/asc1_sdk_usage-2.png)</center>
<center>*Delegated Signature Transactiont*</center>

The following example illustrates signing a transaction with a created logic signature that is signed by a specific account.

```javascript tab="JavaScript"
// get suggested parameters
let params = await algodclient.getTransactionParams().do();
// comment out the next two lines to use suggested fee 
params.fee = 1000;
params.flatFee = true;
console.log(params);
// create logic sig

// b64 example "ASABACI=" is `int 0`
// see more info here: https://developer.algorand.org/docs/features/asc1/sdks/#accessing-teal-program-from-sdks
let program = new Uint8Array(Buffer.from("base64-encoded-program" < PLACEHOLDER >, "base64"));
// let program = new Uint8Array(Buffer.from("ASABACI=" , "base64"));

let lsig = algosdk.makeLogicSig(program);

// sign the logic signature with an account sk
lsig.sign(myAccount.sk);

// create a transaction
let sender = myAccount.addr;
let receiver = "< PLACEHOLDER >";
let amount = 10000;
let closeToRemaninder = undefined;
let note = undefined;
let txn = algosdk.makePaymentTxnWithSuggestedParams(sender, receiver, amount, closeToRemaninder, note, params)

// Create the LogicSigTransaction with contract account LogicSig
let rawSignedTxn = algosdk.signLogicSigTransactionObject(txn, lsig);

// send raw LogicSigTransaction to network    
let tx = (await algodclient.sendRawTransaction(rawSignedTxn.blob).do());
console.log("Transaction : " + tx.txId);    
await waitForConfirmation(algodclient, tx.txId);
```

```python tab="Python"
    # create logic sig
    # program = b"hex-encoded-program"
    # b"\x01\x20\x01\x00\x22 is `int 0`
    # see more info here: https://developer.algorand.org/docs/features/asc1/sdks/#accessing-teal-program-from-sdks
    program = b"\x01\x20\x01\x00\x22"

    # program = b"hex-encoded-program"
    lsig = LogicSig(program)
    sender = lsig.address()

    #Recover the account that is wanting to delegate signature

    passphrase = "25-word-mnemonic<PLACEHOLDER>"
    sk = mnemonic.to_private_key(passphrase)
    addr = account.address_from_private_key(sk)
    print("Address of Sender/Delegator: " + addr)

    # sign the logic signature with an account sk
    lsig.sign(sk)

    # get suggested parameters
    params = algod_client.suggested_params()
    # comment out the next two (2) lines to use suggested fees
    params.flat_fee = True
    params.fee = 1000

    # build transaction
    amount = amount <PLACEHOLDER> 
    closeremainderto = None
    receiver = "<PLACEHOLDER>"

    # create a transaction
    txn = PaymentTxn(
        addr, params, receiver, amount, closeremainderto)
    # Create the LogicSigTransaction with contract account LogicSig
    lstx = transaction.LogicSigTransaction(txn, lsig)

    # send raw LogicSigTransaction to network
    print("This transaction is expected to fail as it is int 0 , always false")
    txid = algod_client.send_transaction(lstx)
    print("Transaction ID: " + txid)
    wait_for_confirmation(algod_client, txid)
```

```java tab="Java"
    public void accountDelegationExample() throws Exception {
        // Initialize an algod client
        if (client == null)
            this.client = connectToNetwork();
        // import your private key mnemonic and address
        
        final String SRC_ACCOUNT = "25-word-mnemonic<PLACEHOLDER>";
 
        Account src = new Account(SRC_ACCOUNT);
        // Set the receiver
        final String RECEIVER = "<PLACEHOLDER>";
        // create logic sig
        // hex example 0x01, 0x20, 0x01, 0x00, 0x22 int 0 returns false, so rawTransaction will fail below
        byte[] program = { 0x01, 0x20, 0x01, 0x00, 0x22 };
        LogicsigSignature lsig = new LogicsigSignature(program, null);
        System.out.println("lsig address: " + lsig.toAddress());
        // sign the logic signature with an account sk
        src.signLogicsig(lsig);
        TransactionParametersResponse params = client.TransactionParams().execute().body();
        // create a transaction

        String note = "Hello World";
        Transaction txn = Transaction.PaymentTransactionBuilder()
                .sender(src.getAddress())
                .note(note.getBytes())
                .amount(100000)
                .receiver(new Address(RECEIVER))
                .suggestedParams(params)
                .build();   

        try {
            // create the LogicSigTransaction with contract account LogicSig
            SignedTransaction stx = Account.signLogicsigTransaction(lsig, txn);

            // send raw LogicSigTransaction to network
            byte[] encodedTxBytes = Encoder.encodeToMsgPack(stx);
            Response<PostTransactionsResponse> rp = client.RawTransaction().rawtxn(encodedTxBytes).execute();
            String id = null;
            if (rp.body() != null) {
                id = rp.body().txId;
            } else {
                System.out.println(rp.message());
            }
            // Wait for transaction confirmation
            waitForConfirmation(id);
            System.out.println("Successfully sent tx with id: " + id);
            // Read the transaction
            PendingTransactionResponse pTrx = client.PendingTransactionInformation(id).execute().body();
  
            JSONObject jsonObj = new JSONObject(pTrx.toString());
            System.out.println("Transaction information (with notes): " + jsonObj.toString(2)); // pretty print
            System.out.println("Decoded note: " + new String(pTrx.txn.tx.note));
        } catch (ApiException e) {
            System.err.println("Exception when calling algod#rawTransaction: " + e.getResponseBody());
        }

    }
```

```go tab="Go"
    // Get private key for sender address
    PASSPHRASE := "25-word-mnemonic<PLACEHOLDER>"
	
    sk, err := mnemonic.ToPrivateKey(PASSPHRASE)	
    pk := sk.Public()
    var a types.Address
    cpk := pk.(ed25519.PublicKey)
    copy(a[:], cpk[:])
    fmt.Printf("Address: %s\n", a.String())	
    sender := a.String()

    // Create logic signature
    // example base64 encoded program "ASABACI="   int 0
    var ma crypto.MultisigAccount
    // program, err :=  base64.StdEncoding.DecodeString("base64-encoded-program"<PLACEHOLDER>)
    program, err :=  base64.StdEncoding.DecodeString("ASABACI=")		
    var args [][]byte
    lsig, err := crypto.MakeLogicSig(program, args, sk, ma)
    addr := crypto.LogicSigAddress(lsig).String()
    fmt.Printf("Escrow Address: %s\n" , addr )
    

    // Create an algod client
    algodClient, err := algod.MakeClient(algodAddress, algodToken)
    if err != nil {
        fmt.Printf("failed to make algod client: %s\n", err)
        return
    }	
    // Construct the transaction
    txParams, err := algodClient.SuggestedParams().Do(context.Background())
    if err != nil {
        fmt.Printf("Error getting suggested tx params: %s\n", err)
        return
    }
    // comment out the next two (2) lines to use suggested fees
    txParams.FlatFee = true
    txParams.Fee = 1000

    // Make transaction
    const receiver = "transaction-receiver"<PLACEHOLDER>	
    const fee = fee<PLACEHOLDER>
    const amount = amount<PLACEHOLDER>

    note := []byte("Hello World")
    genID := txParams.GenesisID
    genHash := txParams.GenesisHash
    firstValidRound := uint64(txParams.FirstRoundValid)
    lastValidRound := uint64(txParams.LastRoundValid)
    tx, err := transaction.MakePaymentTxnWithFlatFee(
        sender, receiver, fee, amount, firstValidRound, lastValidRound,
        note, "", genID, genHash )

    txID, stx, err := crypto.SignLogicsigTransaction(lsig, tx)
    if err != nil {
        fmt.Printf("Signing failed with %v", err)
        return
    }
    fmt.Printf("Signed tx: %v\n", txID)

    // Submit the raw transaction as normal - expected to fail with int 0 program, always false
    fmt.Printf("expected to fail with int 0 program, always false %s\n", err)
    transactionID, err := algodClient.SendRawTransaction(stx).Do(context.Background())
    if err != nil {
        fmt.Printf("Sending failed with %v\n", err)
    }
    fmt.Printf("Transaction ID: %v\n", transactionID)
    // Read transaction
    confirmedTxn, stxn, err := algodClient.PendingTransactionInformation(txID).Do(context.Background())
    if err != nil {
        fmt.Printf("Error retrieving transaction %s\n", txID)
        return
    }
    txnJSON, err := json.MarshalIndent(confirmedTxn.Transaction.Txn, "", "\t")
    if err != nil {
        fmt.Printf("Cannot marshal txn data: %s\n", err)
    }
    fmt.Printf("Transaction information: %s\n", txnJSON)
    fmt.Printf("Decoded note: %s\n", string(stxn.Txn.Note))
```

# Save Transaction Output for Debugging
The goal command-line tool provides functionality to do a test run of a TEAL program using the `goal clerk dryrun` command. This process is described in the [goal TEAL Walkthrough(goal_teal_walkthrough.md)] documentation. From the SDK a logic signature transaction can be written to a file to be used with the `goal clerk dryrun` command. The following code details how this is done. The goal tab illustrates run the `dryrun` on the generated file.

```javascript tab="JavaScript"
    let rawSignedTxn = algosdk.signLogicSigTransaction(txn, lsig);
    fs.writeFileSync("simple.stxn", rawSignedTxn.blob);
```

```python tab="Python"
    logicsig_txn = transaction.LogicSigTransaction(txn, lsig)
    transaction.write_to_file([logicsig_txn], "simple.stxn")
```

```java tab="Java"
    SignedTransaction stx = Account.signLogicsigTransaction(lsig, tx);
    byte[] outBytes = Encoder.encodeToMsgPack(stx);
    try {
        String FILEPATH = "./simple.stxn";
        File file = new File(FILEPATH);
        OutputStream os = new FileOutputStream(file);
        os.write(outBytes);
        os.close();
    }
    catch (Exception e) {
        System.out.println("Exception: " + e);
    }    
```

```go tab="Go"
	txid, stx, err := crypto.SignLogicsigTransaction(lsig, tx)
	if err != nil {
        ...
    }
    f, err := os.Create("simple.stxn")
    if err != nil {
        ...
    }
    defer f.Close()
    if _, err := f.Write(stx); err != nil {
        ...
    }
    if err := f.Sync(); err != nil {
        ...
    }    
```

```text tab="Goal"
$ goal clerk dryrun -t simple.stxn
tx[0] cost=2 trace:
  1 intcblock => <empty stack>
  4 intc_0 => 0 0x0

REJECT
```


# Passing Parameters using the SDKs
The SDKs require that parameters to a TEAL program be in byte arrays. This byte array is passed to the method that creates the logic signature. Currently, TEAL parameters must be either unsigned integers or binary strings. If comparing a constant string in TEAL, the constant within the TEAL program must be encoded in hex or base64. See the TEAL tab below for a simple example of comparing the string argument used in the other examples. SDK native language functions can be used to encode the parameters to the TEAL program correctly. The example below illustrates both a string parameter and an integer.


```javascript tab="JavaScript"
    //string parameter
    let args = ["my string"];
    let lsig = algosdk.makeLogicSig(program, args);
    //integer parameter
    let args = [[123]];
    let lsig = algosdk.makeLogicSig(program, args);
```

```python tab="Python"
    #string parameter
    arg_str = "my string"
    arg1 = arg_str.encode()
    lsig = transaction.LogicSig(program, args=[arg1])
    #integer parameter
    arg1 = (123).to_bytes(8, 'big')
    lsig = transaction.LogicSig(program, args=[arg1])
```

```java tab="Java"
    // string parameter
    ArrayList<byte[]> teal_args = new ArrayList<byte[]>();
    String orig = "my string";
    teal_args.add(orig.getBytes());
    LogicsigSignature lsig = new LogicsigSignature(program, teal_args);    

    // integer parameter
    ArrayList<byte[]> teal_args = new ArrayList<byte[]>();
    byte[] arg1 = {123};
    teal_args.add(arg1);
    LogicsigSignature lsig = new LogicsigSignature(program, teal_args);
```

```go tab="Go"
    // string parameter
	args := make([][]byte, 1)
	args[0] = []byte("my string")
    lsig, err := crypto.MakeLogicSig(program, args, sk, ma)
    
    // integer parameter
    args := make([][]byte, 1)
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], 123)
	args[0] = buf[:]
```

```text tab="TEAL"
//Never use this code for a real tansaction
//for educational purposes only
//string compare
arg_0
byte b64 bXkgc3RyaW5n
==
//integer compare
arg_0
btoi
int 123
==
```


!!! info
    The example code snippets are provided throughout this page and are abbreviated for conciseness and clarity. Full running code examples for each SDK are available within the GitHub repo for V1 and V2 at [/examples/smart_contracts](https://github.com/algorand/docs/tree/master/examples/smart_contracts) and for [download](https://github.com/algorand/docs/blob/master/examples/smart_contracts/smart_contracts.zip?raw=true) (.zip).


