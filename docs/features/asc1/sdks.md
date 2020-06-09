title: Using the SDKs

This guide covers using TEAL programs with contract accounts or delegated signatures with the available SDKs. The methods covered in this documentation are used for custom TEAL code and provide general access to any TEAL program. Algorand Smart Contract Templates are also available for common use case functionality like Hash Time-Lock Contracts, Split Payments, Limit Orders, etc. Developer documentation describing the process for using these templates will be available soon.

Each SDK's install process is discussed in the [SDK Reference](../../reference/sdks/index.md) documentation.

Code snippets are abbreviated for conciseness and clarity. See the full code example for each SDK at the bottom of this guide.

# Accessing TEAL program from SDKs
Before a TEAL program can be used is the SDKs, it must be compiled using the `goal` tool. The [goal TEAL walkthrough](./goal_teal_walkthrough.md) documentation explains this process. Once a TEAL program is compiled, the bytes of the program can be retrieved in various ways. Most of the SDKs support the bytes encoded in base64 or hexadecimal format. The following example illustrates using shell commands to export the binary to hexadecimal or a base64 encoded string.

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
    let program = new Uint8Array(Buffer.from("ASABACI=", "base64"));
    let lsig = algosdk.makeLogicSig(program);
    let txn = {
        "from": contract-address<PLACEHOLDER>,
        "to": receiver<PLACEHOLDER>,
        "fee": fee<PLACEHOLDER>,
        "amount": amount<PLACEHOLDER>,
        "firstRound": first-valid-round<PLACEHOLDER>,
        "lastRound": last-valid-round<PLACEHOLDER>,
        "genesisID": genesis-id<PLACEHOLDER>,
        "genesisHash": genesis-hash<PLACEHOLDER>
    };
    let rawSignedTxn = algosdk.signLogicSigTransaction(txn, lsig);
    let tx = (await algodclient.sendRawTransaction(rawSignedTxn.blob));
```

```python tab="Python"
	program = b"\x01\x20\x01\x00\x22"  
	lsig = transaction.LogicSig(program)
	txn = transaction.PaymentTxn(contract-address<PLACEHOLDER>, fee<PLACEHOLDER>, first-valid-round<PLACEHOLDER>, last-valid-round<PLACEHOLDER>, genesis-hash<PLACEHOLDER>, receiver<PLACEHOLDER>, amount<PLACEHOLDER>, close-remainder-to-address<PLACEHOLDER>)
	lstx = transaction.LogicSigTransaction(txn, lsig)
	txid = acl.send_transaction(lstx)
```

```java tab="Java"
    byte[] program = {
        0x01, 0x20, 0x01, 0x00, 0x22  // int 0
    };
    LogicsigSignature lsig = new LogicsigSignature(program, null);
    Transaction tx = new Transaction(lsig.toAddress(), new Address(receiver-address<PLACEHOLDER>), BigInteger.valueOf(fee<PLACEHOLDER>), amount<PLACEHOLDER>, first-valid-round<PLACEHOLDER>, last-valid-round<PLACEHOLDER>, genesis-id<PLACEHOLDER>, genesis-hash<PLACEHOLDER>);
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
		addr, receiver<PLACEHOLDER>, fee<PLACEHOLDER>, amount<PLACEHOLDER>, first-valid-round<PLACEHOLDER>, last-valid-round<PLACEHOLDER>,
		note<PLACEHOLDER>, close-remainder-to-address<PLACEHOLDER>, genesis-id<PLACEHOLDER>, genesis-hash<PLACEHOLDER> )
	txid, stx, err := crypto.SignLogicsigTransaction(lsig, tx)
	transactionID, err := algodClient.SendRawTransaction(stx)
```

# Account Delegation SDK Usage
ASC1 allows TEAL logic to be used to delegate signature authority. This allows specific accounts or multi-signature accounts to sign logic that allows transactions from the account to be approved based on the TEAL logic. The [ASC1 Usage Modes](./modes.md) documentation explains ASC1 modes in more detail. 

Delegated transactions are special transactions where the `sender` also signs the logic and the transaction is then signed with the [logic signature](./modes.md#logic-signature). In all other aspects, the transaction functions as any other transaction. See [Transaction](../transactions/index.md) documentation for details on setting up a payment transaction.

Delegated Logic Signatures require that the logic signature be signed from a specific account or a multi-signature account. The TEAL program is first loaded, then a Logic Signature is created and then the Logic Signature is signed by a specific account or multi-signaure account. The transaction is created as normal. The transaction is then signed with the Logic Signature. From an SDK standpoint, the following process should be used.

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
    let program = new Uint8Array(Buffer.from("ASABACI=", "base64"));
    let lsig = algosdk.makeLogicSig(program);
    lsig.sign(secret-key<PLACEHOLDER>);
    let txn = {
        "from": contract-address<PLACEHOLDER>,
        "to": receiver<PLACEHOLDER>,
        "fee": fee<PLACEHOLDER>,
        "amount": amount<PLACEHOLDER>,
        "firstRound": first-valid-round<PLACEHOLDER>,
        "lastRound": last-valid-round<PLACEHOLDER>,
        "genesisID": genesis-id<PLACEHOLDER>,
        "genesisHash": genesis-hash<PLACEHOLDER>
    };
    let rawSignedTxn = algosdk.signLogicSigTransaction(txn, lsig);
    let tx = (await algodclient.sendRawTransaction(rawSignedTxn.blob));
```

```python tab="Python"
	program = b"\x01\x20\x01\x00\x22"  
	lsig = transaction.LogicSig(program)
    lsig.sign(secret-key<PLACEHOLDER>)
	txn = transaction.PaymentTxn(contract-address<PLACEHOLDER>, fee<PLACEHOLDER>, first-valid-round<PLACEHOLDER>, last-valid-round<PLACEHOLDER>, genesis-hash<PLACEHOLDER>, receiver<PLACEHOLDER>, amount<PLACEHOLDER>, close-remainder-to-address<PLACEHOLDER>)
	lstx = transaction.LogicSigTransaction(txn, lsig)
	txid = acl.send_transaction(lstx)
```

```java tab="Java"
    byte[] program = {
        0x01, 0x20, 0x01, 0x00, 0x22  // int 0
    };
    LogicsigSignature lsig = new LogicsigSignature(program, null);
    Account account = new Account(secret-key<PLACEHOLDER>);
    account.signLogicsig(lsig);
    Transaction tx = new Transaction(lsig.toAddress(), new Address(receiver-address<PLACEHOLDER>), BigInteger.valueOf(fee<PLACEHOLDER>), amount<PLACEHOLDER>, first-valid-round<PLACEHOLDER>, last-valid-round<PLACEHOLDER>, genesis-id<PLACEHOLDER>, genesis-hash<PLACEHOLDER>);
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
	lsig, err := crypto.MakeLogicSig(program, args, secret-key<PLACEHOLDER>, nil)

    PASSPHRASE := "25-word-mnemonic<PLACEHOLDER>"
    sk, err := mnemonic.ToPrivateKey(PASSPHRASE)	
    pk := sk.Public()
    var a types.Address
    cpk := pk.(ed25519.PublicKey)
    copy(a[:], cpk[:])
    fmt.Printf("Address: %s\n", a.String())	
    sender := a.String()
    tx, err := transaction.MakePaymentTxnWithFlatFee(	
		sender, receiver<PLACEHOLDER>, fee<PLACEHOLDER>, amount<PLACEHOLDER>, first-valid-round<PLACEHOLDER>, last-valid-round<PLACEHOLDER>,
		note<PLACEHOLDER>, close-remainder-to-address<PLACEHOLDER>, genesis-id<PLACEHOLDER>, genesis-hash<PLACEHOLDER> )
	txid, stx, err := crypto.SignLogicsigTransaction(lsig, tx)
	transactionID, err := algodClient.SendRawTransaction(stx)
```

# Save Transaction Output for Debugging
The goal command-line tool provides functionality to do a test run of a TEAL program using the `goal clerk dryun` command. This process is described in the [goal TEAL Walkthrough(goal_teal_walkthrough.md)] documentation. From the SDK a logic signature transaction can be written to a file to be used with the `goal clerk dryrun` command. The following code details how this is done. The goal tab illustrates run the dryrun on the generated file.

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

??? example "Complete Example = Contract Account"
    
    ```javascript tab="JavaScript"
    const algosdk = require('algosdk');

    const token = "algod-token"<PLACEHOLDER>;
    const server = "algod-address"<PLACEHOLDER>;
    const port = algod-port<PLACEHOLDER>;

    // create an algod client
    let algodclient = new algosdk.Algod(token, server, port);
    (async() => {

        // get suggested parameters
        let params = await algodclient.getTransactionParams();
        let endRound = params.lastRound + parseInt(1000);
        let fee = await algodclient.suggestedFee();

        // create logic sig
        // b64 example "ASABACI=" 
        let program = new Uint8Array(Buffer.from("base64-encoded-program"<PLACEHOLDER>, "base64"));
        let lsig = algosdk.makeLogicSig(program);

        // create a transaction
        let txn = {
            "from": lsig.address(),
            "to": "receiver-address"<PLACEHOLDER>,
            "fee": params.fee,
            "amount": amount<PLACEHOLDER>,
            "firstRound": params.lastRound,
            "lastRound": endRound,
            "genesisID": params.genesisID,
            "genesisHash": params.genesishashb64
        };

        // Create the LogicSigTransaction with contract account LogicSig 
        let rawSignedTxn = algosdk.signLogicSigTransaction(txn, lsig);

        // send raw LogicSigTransaction to network
        let tx = (await algodclient.sendRawTransaction(rawSignedTxn.blob));
        console.log("Transaction : " + tx.txId);
    })().catch(e => {
        console.log(e);
    });
    ```

    ```python tab="Python"
    from algosdk import algod, transaction, account, mnemonic
    
    try:

        # create logic sig
        # hex example b"\x01\x20\x01\x00\x22" 
        program = b"hex-encoded-program"  
        lsig = transaction.LogicSig(program)
        sender = lsig.address()

        # create an algod client
        algod_token = "algod-token"<PLACEHOLDER>
        algod_address = "algod-address"<PLACEHOLDER>
        receiver = "receiver-address"<PLACEHOLDER>
        acl = algod.AlgodClient(algod_token, algod_address)

        # get suggested parameters
        params = acl.suggested_params()
        gen = params["genesisID"]
        gh = params["genesishashb64"]
        last_round = params["lastRound"]
        fee = params["fee"]
        amount = amount<PLACEHOLDER>
        closeremainderto = None

        # create a transaction
        txn = transaction.PaymentTxn(sender, fee, last_round, last_round+100, gh, receiver, amount, closeremainderto)
        # Create the LogicSigTransaction with contract account LogicSig
        lstx = transaction.LogicSigTransaction(txn, lsig)

        # send raw LogicSigTransaction to network
        txid = acl.send_transaction(lstx)
        print("Transaction ID: " + txid )
    except Exception as e:
        print(e)    
    ```

    ```java tab="Java"
    package com.algorand.algosdk.teal;

    import java.math.BigInteger;
    import java.util.ArrayList;

    import com.algorand.algosdk.account.Account;
    import com.algorand.algosdk.algod.client.AlgodClient;
    import com.algorand.algosdk.algod.client.ApiException;
    import com.algorand.algosdk.algod.client.api.AlgodApi;
    import com.algorand.algosdk.algod.client.auth.ApiKeyAuth;
    import com.algorand.algosdk.algod.client.model.TransactionID;
    import com.algorand.algosdk.algod.client.model.TransactionParams;
    import com.algorand.algosdk.crypto.Address;
    import com.algorand.algosdk.crypto.Digest;
    import com.algorand.algosdk.crypto.LogicsigSignature;
    import com.algorand.algosdk.transaction.SignedTransaction;
    import com.algorand.algosdk.transaction.Transaction;
    import com.algorand.algosdk.util.Encoder;

    public class SubmitTealContract 
    {
        public static void main(final String args[]) throws Exception {
            // Initialize an algod client
            final String ALGOD_API_ADDR = "algod-address"<PLACEHOLDER>;
            final String ALGOD_API_TOKEN = "algod-token"<PLACEHOLDER>;

            //Create an instance of the algod API client
            AlgodClient client = (AlgodClient) new AlgodClient().setBasePath(ALGOD_API_ADDR);
            ApiKeyAuth api_key = (ApiKeyAuth) client.getAuthentication("api_key");
            api_key.setApiKey(ALGOD_API_TOKEN);
            AlgodApi algodApiInstance = new AlgodApi(client);

            // Set the reciever
            final String DEST_ADDR = "receiver-address"<PLACEHOLDER>;
    

            // get suggested parameters
            BigInteger suggestedFeePerByte = BigInteger.valueOf(1);
            BigInteger firstRound = BigInteger.valueOf(301);
            String genId = null;
            Digest genesisHash = null;
            try {
                // Get suggested parameters from the node
                TransactionParams params = algodApiInstance.transactionParams();
                suggestedFeePerByte = params.getFee();
                firstRound = params.getLastRound();
                genId = params.getGenesisID();
                genesisHash = new Digest(params.getGenesishashb64());

            } catch (ApiException e) {
                System.err.println("Exception when calling algod#transactionParams");
                e.printStackTrace();
            }

            // create logic sig
            // hex example 0x01, 0x20, 0x01, 0x00, 0x22
            byte[] program = {
                hex-encoded-program<PLACEHOLDER>
                };
            LogicsigSignature lsig = new LogicsigSignature(program, teal_args);        

            // create a transaction
            BigInteger amount = BigInteger.valueOf(amount<PLACEHOLDER>);
            BigInteger lastRound = firstRound.add(BigInteger.valueOf(1000)); 
            Transaction tx = new Transaction(lsig.toAddress(), new Address(DEST_ADDR), BigInteger.valueOf(1000), amount, firstRound, lastRound, genId, genesisHash);

            try {
                // create the LogicSigTransaction with contract account LogicSig
                SignedTransaction stx = Account.signLogicsigTransaction(lsig, tx);
                
                // send raw LogicSigTransaction to network
                byte[] encodedTxBytes = Encoder.encodeToMsgPack(stx);
                TransactionID id = algodApiInstance.rawTransaction(encodedTxBytes);
                System.out.println("Successfully sent tx with id: " + id);

            } catch (ApiException e) {
                System.err.println("Exception when calling algod#rawTransaction: " + e.getResponseBody());
            }

        }
    }
    ```

    ```go tab="Go"
    package main

    import (
        "fmt"
        "encoding/base64"

        "golang.org/x/crypto/ed25519"

        "github.com/algorand/go-algorand-sdk/client/algod"
        "github.com/algorand/go-algorand-sdk/crypto"
        "github.com/algorand/go-algorand-sdk/transaction"
    )

    func main() {

        const algodToken = "algod-token<PLACEHOLDER>"
        const algodAddress = "algod-address<PLACEHOLDER>"

        // Create logic signature
        // example base64 encoded program "ASABACI="
        var sk ed25519.PrivateKey
        var ma crypto.MultisigAccount
        program, err :=  base64.StdEncoding.DecodeString("base64-encoded-program<PLACEHOLDER>")
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
        // Get suggested params for the transaction
        txParams, err := algodClient.SuggestedParams()
        if err != nil {
                fmt.Printf("error getting suggested tx params: %s\n", err)
                return
        }

        // Make transaction
        const receiver = "transaction-receiver"<PLACEHOLDER>
        const fee = fee<PLACEHOLDER>
        const amount = amount<PLACEHOLDER>
        var note []byte	
        genID := txParams.GenesisID
        genHash := txParams.GenesisHash	
        tx, err := transaction.MakePaymentTxnWithFlatFee(
            addr, receiver, fee, amount, txParams.LastRound, (txParams.LastRound + 1000),
            note, "", genID, genHash )

        txid, stx, err := crypto.SignLogicsigTransaction(lsig, tx)
        if err != nil {
            fmt.Printf("Signing failed with %v", err)
            return
        }
        fmt.Printf("Signed tx: %v\n", txid)

        // Submit the raw transaction to network
        transactionID, err := algodClient.SendRawTransaction(stx)
        if err != nil {
            fmt.Printf("Sending failed with %v\n", err)
        }
        fmt.Printf("Transaction ID: %v\n", transactionID)

    }
    ```

??? example "Complete Example = Account Delegation"
    
    ```javascript tab="JavaScript"
    const algosdk = require('algosdk');

    const token = "algod-token"<PLACEHOLDER>;
    const server = "algod-address"<PLACEHOLDER>;
    const port = algod-port<PLACEHOLDER>;

    // import your private key mnemonic and address
    let PASSPHRASE = "25-word-mnemonic<PLACEHOLDER>";
    var myAccount = algosdk.mnemonicToSecretKey(PASSPHRASE);
    console.log("My Address: " + myAccount.addr);

    // create an algod client
    let algodclient = new algosdk.Algod(token, server, port);
    (async() => {

        // get suggested parameters
        let params = await algodclient.getTransactionParams();
        let endRound = params.lastRound + parseInt(1000);
        let fee = await algodclient.suggestedFee();

        // create logic sig
        // b64 example "ASABACI="
        let program = new Uint8Array(Buffer.from("base64-encoded-program"<PLACEHOLDER>, "base64"));
        let lsig = algosdk.makeLogicSig(program);

        // sign the logic signature with an account sk
        lsig.sign(myAccount.sk);

        // create a transaction
        let txn = {
            "from": myAccount.addr,
            "to": "receiver-address"<PLACEHOLDER>,
            "fee": params.fee,
            "amount": amount<PLACEHOLDER>,
            "firstRound": params.lastRound,
            "lastRound": endRound,
            "genesisID": params.genesisID,
            "genesisHash": params.genesishashb64
        };
    
        // Create the LogicSigTransaction with contract account LogicSig 
        let rawSignedTxn = algosdk.signLogicSigTransaction(txn, lsig);

        // send raw LogicSigTransaction to network    
        let tx = (await algodclient.sendRawTransaction(rawSignedTxn.blob));
        console.log("Transaction : " + tx.txId);
    })().catch(e => {
        console.log(e);
    });
    ```

    ```python tab="Python"
    from algosdk import algod, transaction, account, mnemonic
    
    try:

        # create logic sig
        # hex example b"\x01\x20\x01\x00\x22" 
        program = b"hex-encoded-program"  
        lsig = transaction.LogicSig(program)
        sender = lsig.address()

        #Recover the account that is wanting to delegate signature
        passphrase = "25-word-mnemonic<PLACEHOLDER>";
        sk = mnemonic.to_private_key(passphrase)
        addr = account.address_from_private_key(sk)
        print( "Address of Sender/Delgator: " + addr )

        # sign the logic signature with an account sk
        lsig.sign(sk)

        # create an algod client
        algod_token = "algod-token"<PLACEHOLDER>
        algod_address = "algod-address"<PLACEHOLDER>
        receiver = "receiver-address"<PLACEHOLDER>
        acl = algod.AlgodClient(algod_token, algod_address)

        # get suggested parameters
        params = acl.suggested_params()
        gen = params["genesisID"]
        gh = params["genesishashb64"]
        last_round = params["lastRound"]
        fee = params["fee"]
        amount = amount<PLACEHOLDER>
        closeremainderto = None

        # create a transaction
        txn = transaction.PaymentTxn(addr, fee, last_round, last_round+100, gh, receiver, amount, closeremainderto)
        # Create the LogicSigTransaction with contract account LogicSig
        lstx = transaction.LogicSigTransaction(txn, lsig)

        # send raw LogicSigTransaction to network
        txid = acl.send_transaction(lstx)
        print("Transaction ID: " + txid )
    except Exception as e:
        print(e)    
    ```

    ```java tab="Java"
    package com.algorand.algosdk.teal;

    import java.math.BigInteger;
    import java.util.ArrayList;

    import com.algorand.algosdk.account.Account;
    import com.algorand.algosdk.algod.client.AlgodClient;
    import com.algorand.algosdk.algod.client.ApiException;
    import com.algorand.algosdk.algod.client.api.AlgodApi;
    import com.algorand.algosdk.algod.client.auth.ApiKeyAuth;
    import com.algorand.algosdk.algod.client.model.TransactionID;
    import com.algorand.algosdk.algod.client.model.TransactionParams;
    import com.algorand.algosdk.crypto.Address;
    import com.algorand.algosdk.crypto.Digest;
    import com.algorand.algosdk.crypto.LogicsigSignature;
    import com.algorand.algosdk.transaction.SignedTransaction;
    import com.algorand.algosdk.transaction.Transaction;
    import com.algorand.algosdk.util.Encoder;

    public class SubmitTealDelegate 
    {
        public static void main(final String args[]) throws Exception {
            // Initialize an algod client
            final String ALGOD_API_ADDR = "algod-address"<PLACEHOLDER>;
            final String ALGOD_API_TOKEN = "algod-token"<PLACEHOLDER>;

            //Create an instance of the algod API client
            AlgodClient client = (AlgodClient) new AlgodClient().setBasePath(ALGOD_API_ADDR);
            ApiKeyAuth api_key = (ApiKeyAuth) client.getAuthentication("api_key");
            api_key.setApiKey(ALGOD_API_TOKEN);
            AlgodApi algodApiInstance = new AlgodApi(client);

            // import your private key mnemonic and address
            final String SRC_ACCOUNT = "25-word-mnemonic<PLACEHOLDER>";
            Account src = new Account(SRC_ACCOUNT);
            // Set the reciever
            final String DEST_ADDR = "receiver-address"<PLACEHOLDER>;

            // get suggested parameters
            BigInteger suggestedFeePerByte = BigInteger.valueOf(1);
            BigInteger firstRound = BigInteger.valueOf(301);
            String genId = null;
            Digest genesisHash = null;
            try {
                // Get suggested parameters from the node
                TransactionParams params = algodApiInstance.transactionParams();
                suggestedFeePerByte = params.getFee();
                firstRound = params.getLastRound();
                genId = params.getGenesisID();
                genesisHash = new Digest(params.getGenesishashb64());

            } catch (ApiException e) {
                System.err.println("Exception when calling algod#transactionParams");
                e.printStackTrace();
            }

            // create logic sig
            // hex example 0x01, 0x20, 0x01, 0x00, 0x22
            byte[] program = {
                hex-encoded-program<PLACEHOLDER>
                };
            LogicsigSignature lsig = new LogicsigSignature(program, teal_args);        

            // sign the logic signature with an account sk
            src.signLogicsig(lsig);

            // create a transaction
            BigInteger amount = BigInteger.valueOf(amount<PLACEHOLDER>);
            BigInteger lastRound = firstRound.add(BigInteger.valueOf(1000)); 
            Transaction tx = new Transaction(src.getAddress(), new Address(DEST_ADDR), BigInteger.valueOf(1000), amount, firstRound, lastRound, genId, genesisHash);

            try {
                // create the LogicSigTransaction with contract account LogicSig
                SignedTransaction stx = Account.signLogicsigTransaction(lsig, tx);
                
                // send raw LogicSigTransaction to network
                byte[] encodedTxBytes = Encoder.encodeToMsgPack(stx);
                TransactionID id = algodApiInstance.rawTransaction(encodedTxBytes);
                System.out.println("Successfully sent tx with id: " + id);

            } catch (ApiException e) {
                System.err.println("Exception when calling algod#rawTransaction: " + e.getResponseBody());
            }

        }
    }
    ```

    ```go tab="Go"
    package main

    import (
        "fmt"
        "encoding/base64"

        "github.com/algorand/go-algorand-sdk/client/algod"
        "github.com/algorand/go-algorand-sdk/crypto"
        "github.com/algorand/go-algorand-sdk/mnemonic"
        "github.com/algorand/go-algorand-sdk/transaction"
        "github.com/algorand/go-algorand-sdk/types"
    )

    func main() {

        const algodToken = "algod-token"<PLACEHOLDER>
        const algodAddress = "algod-address"<PLACEHOLDER>


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
        // example base64 encoded program "ASABACI="
        var ma crypto.MultisigAccount
        program, err :=  base64.StdEncoding.DecodeString("base64-encoded-program"<PLACEHOLDER>)
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
        // Get suggested params for the transaction
        txParams, err := algodClient.SuggestedParams()
        if err != nil {
                fmt.Printf("error getting suggested tx params: %s\n", err)
                return
        }

        // Make transaction
        const receiver = "transaction-receiver"<PLACEHOLDER>
        const fee = fee<PLACEHOLDER>
        const amount = amount<PLACEHOLDER>
        var note []byte	
        genID := txParams.GenesisID
        genHash := txParams.GenesisHash
        tx, err := transaction.MakePaymentTxnWithFlatFee(
            sender.String(), receiver, fee, amount, txParams.LastRound, (txParams.LastRound + 1000),
            note, "", genID, genHash )

        txid, stx, err := crypto.SignLogicsigTransaction(lsig, tx)
        if err != nil {
            fmt.Printf("Signing failed with %v", err)
            return
        }
        fmt.Printf("Signed tx: %v\n", txid)

        // Submit the raw transaction as normal
        transactionID, err := algodClient.SendRawTransaction(stx)
        if err != nil {
            fmt.Printf("Sending failed with %v\n", err)
        }
        fmt.Printf("Transaction ID: %v\n", transactionID)

    }
    ```
