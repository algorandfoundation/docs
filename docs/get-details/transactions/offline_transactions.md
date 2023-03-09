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
    ``` javascript
        // get network suggested parameters
        let params = await algodClient.getTransactionParams().do();
        const enc = new TextEncoder();
        const note = enc.encode("Hello World");
        console.log(note);
        let txn = algosdk.makePaymentTxnWithSuggestedParams(myAccount.addr, receiver, 1000000, undefined, note, params);        
        // Save transaction to file
        fs.writeFileSync('./unsigned.txn', algosdk.encodeUnsignedTransaction( txn ));     
    
        // read transaction from file and sign it
        let txn = algosdk.decodeUnsignedTransaction(fs.readFileSync('./unsigned.txn'));  
        let signedTxn = algosdk.signTransaction(txn, myAccount.sk);
        let txId = signedTxn.txID;
        console.log("Signed transaction with txID: %s", txId);
        // send signed transaction to node
        await algodClient.sendRawTransaction(signedTxn.blob).do();
        // Wait for transaction to be confirmed
        let confirmedTxn = await algosdk.waitForConfirmation(algodClient, txId, 4);
        //Get the completed Transaction
        console.log("Transaction " + txId + " confirmed in round " + confirmedTxn["confirmed-round"]);
        var string = new TextDecoder().decode(confirmedTxn.txn.txn.note);
        console.log("Note field: ", string);       
    ```
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
[Snippet Source](https://github.com/barnjamin/py-algorand-sdk/blob/doc-examples/_examples/codec.py#L31-L43)
<!-- ===PYSDK_CODEC_TRANSACTION_UNSIGNED=== -->

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
[Snippet Source](https://github.com/barnjamin/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/CodecExamples.java#L48-L58)
<!-- ===JAVASDK_CODEC_TRANSACTION_UNSIGNED=== -->

=== "Go"
<!-- ===GOSDK_CODEC_TRANSACTION_UNSIGNED=== -->
    ``` go
        // Construct the transaction
        txParams, err := algodClient.SuggestedParams().Do(context.Background())
        if err != nil {
            fmt.Printf("Error getting suggested tx params: %s\n", err)
            return
        }
        // comment out the next two (2) lines to use suggested fees
        txParams.FlatFee = true
        txParams.Fee = 1000
        fromAddr := myAddress
        toAddr := "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A"
        var amount uint64 = 1000000
        var minFee uint64 = 1000
        note := []byte("Hello World")
        genID := txParams.GenesisID
        genHash := txParams.GenesisHash
        firstValidRound := uint64(txParams.FirstRoundValid)
        lastValidRound := uint64(txParams.LastRoundValid)
        tx, err := transaction.MakePaymentTxnWithFlatFee(fromAddr, toAddr, minFee, amount, firstValidRound, lastValidRound, note, "", genID, genHash)
        if err != nil {
            fmt.Printf("Error creating transaction: %s\n", err)
            return
        }
        unsignedTx := types.SignedTxn{
            Txn: tx,
        }
        // save unsigned Transaction to file
        err = ioutil.WriteFile("./unsigned.txn", msgpack.Encode(unsignedTx), 0644)
        if err == nil {
            fmt.Printf("Saved unsigned transaction to file\n")
            return
        }
        fmt.Printf("Failed in saving trx to file, error %s\n", err) 


        // read unsigned transaction from file
        dat, err := ioutil.ReadFile("./unsigned.txn")
        if err != nil {
            fmt.Printf("Error reading transaction from file: %s\n", err)
            return
        }
        var unsignedTxRaw types.SignedTxn
        var unsignedTxn types.Transaction
        msgpack.Decode(dat, &unsignedTxRaw)
        unsignedTxn = unsignedTxRaw.Txn
        // recover account
        myAddress, privateKey := recoverAccount()
        fmt.Printf("Address is: %s\n", myAddress)
        // Check account balance
        accountInfo, err := algodClient.AccountInformation(myAddress).Do(context.Background())
        if err != nil {
            fmt.Printf("Error getting account info: %s\n", err)
            return
        }
        fmt.Printf("Account balance: %d microAlgos\n", accountInfo.Amount)
        // Sign the transaction
        txID, signedTxn, err := crypto.SignTransaction(privateKey, unsignedTxn)
        if err != nil {
            fmt.Printf("Failed to sign transaction: %s\n", err)
            return
        }
        fmt.Printf("Signed txid: %s\n", txID)
        // Submit the transaction
        sendResponse, err := algodClient.SendRawTransaction(signedTxn).Do(context.Background())
        if err != nil {
            fmt.Printf("failed to send transaction: %s\n", err)
            return
        }
        fmt.Printf("Submitted transaction %s\n", sendResponse)

        // Wait for confirmation
        confirmedTxn, err := transaction.WaitForConfirmation(algodClient,txID,  4, context.Background())
        if err != nil {
            fmt.Printf("Error waiting for confirmation on txID: %s\n", sendResponse)
            return
        }
        fmt.Printf("Confirmed Transaction: %s in Round %d\n", txID ,confirmedTxn.ConfirmedRound)


        fmt.Printf("Decoded note: %s\n", string(confirmedTxn.Transaction.Txn.Note))
    ```
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
    ``` javascript
        // get network suggested parameters
        let params = await algodClient.getTransactionParams().do();
        // setup a transaction
        const enc = new TextEncoder();
        const note = enc.encode("Hello World");
        console.log(note);
        let txn = algosdk.makePaymentTxnWithSuggestedParams(myAccount.addr, receiver, 1000000, undefined, note, params);        
        // sign transaction and write to file
        let signedTxn = txn.signTxn(myAccount.sk);
        fs.writeFileSync('./signed.stxn', signedTxn );  
        
        // read signed transaction from file
        let stx = fs.readFileSync("./signed.stxn");
        // send signed transaction to node
        let tx = await algodClient.sendRawTransaction(stx).do();
        console.log("Signed transaction with txID: %s", tx.txId);
        // Wait for confirmation
        let confirmedTxn = await waitForConfirmation(algodClient, tx.txId, 4);
        //Get the completed Transaction
        console.log("Transaction " + tx.txId + " confirmed in round " + confirmedTxn["confirmed-round"]);
        let mytxinfo = JSON.stringify(confirmedTxn.txn.txn, undefined, 2);
        console.log("Transaction information: %o", mytxinfo);
        var string = new TextDecoder().decode(confirmedTxn.txn.txn.note);
        console.log("Note field: ", string);    
    ```
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
[Snippet Source](https://github.com/barnjamin/py-algorand-sdk/blob/doc-examples/_examples/codec.py#L48-L59)
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
[Snippet Source](https://github.com/barnjamin/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/CodecExamples.java#L61-L67)
<!-- ===JAVASDK_CODEC_TRANSACTION_SIGNED=== -->

=== "Go"
<!-- ===GOSDK_CODEC_TRANSACTION_SIGNED=== -->
    ``` go
        // Construct the transaction
        txParams, err := algodClient.SuggestedParams().Do(context.Background())
        if err != nil {
            fmt.Printf("Error getting suggested tx params: %s\n", err)
            return
        }
        // comment out the next two (2) lines to use suggested fees
        txParams.FlatFee = true
        txParams.Fee = 1000
        fromAddr := myAddress
        toAddr := "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A"
        var amount uint64 = 1000000
        var minFee uint64 = 1000
        note := []byte("Hello World")
        genID := txParams.GenesisID
        genHash := txParams.GenesisHash
        firstValidRound := uint64(txParams.FirstRoundValid)
        lastValidRound := uint64(txParams.LastRoundValid)
        tx, err := transaction.MakePaymentTxnWithFlatFee(fromAddr, toAddr, minFee, amount, firstValidRound, lastValidRound, note, "", genID, genHash)
        if err != nil {
            fmt.Printf("Error creating transaction: %s\n", err)
            return
        }
        // Sign the transaction
        txID, signedTxn, err := crypto.SignTransaction(privateKey, tx)
        if err != nil {
            fmt.Printf("Failed to sign transaction: %s\n", err)
            return
        }
        fmt.Printf("Signed txid: %s\n", txID)
        // Save the signed transaction to file
        err = ioutil.WriteFile("./signed.stxn", signedTxn, 0644)
        if err == nil {
            fmt.Printf("Saved signed transaction to file\n")
            return
        }
        fmt.Printf("Failed in saving trx to file, error %s\n", err)

        
        // Read signed transaction from file
        dat, err := ioutil.ReadFile("./signed.stxn")
        if err != nil {
            fmt.Printf("Error reading signed transaction from file: %s\n", err)
            return
        }
        // Submit the transaction
        sendResponse, err := algodClient.SendRawTransaction(dat).Do(context.Background())
        if err != nil {
            fmt.Printf("failed to send transaction: %s\n", err)
            return
        }
        fmt.Printf("Submitted transaction %s\n", sendResponse)

        // Wait for confirmation
        confirmedTxn, err := transaction.WaitForConfirmation(algodClient,sendResponse,  4, context.Background())
        if err != nil {
            fmt.Printf("Error waiting for confirmation on txID: %s\n", sendResponse)
            return
        }
        fmt.Printf("Confirmed Transaction: %s in Round %d\n", sendResponse ,confirmedTxn.ConfirmedRound)


        fmt.Printf("Decoded note: %s\n", string(confirmedTxn.Transaction.Txn.Note))
    ```
<!-- ===GOSDK_CODEC_TRANSACTION_SIGNED=== -->

=== "goal"
<!-- ===GOAL_CODEC_TRANSACTION_SIGNED=== -->
    ``` goal
    $ goal clerk rawsend --filename signed.txn
    ```
<!-- ===GOAL_CODEC_TRANSACTION_SIGNED=== -->

!!! info
    Example transaction code snippets are provided throughout this page. Full running code transaction examples as well as **offline multisig** for each SDK are available within the GitHub repo at [/examples/offline](https://github.com/algorand/docs/tree/master/examples/offline) and for [download](https://github.com/algorand/docs/blob/master/examples/offine/offline.zip?raw=true) (.zip).
    
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