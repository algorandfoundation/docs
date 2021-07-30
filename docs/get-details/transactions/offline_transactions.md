title: Offline authorization

This section explains how to authorize transactions with private keys that are kept **offline**. In particular, this guide shows how to create and save transactions to a file that can then be transferred to an offline device for signing. To learn about the structure of transactions and how to authorize them in general visit the [Transactions Structure](./index.md) and [Authorizing Transactions](./signatures.md) sections, respectively.

The same methodology described here can also be used to work with [LogicSignatures](../asc1/stateless/modes.md#logic-signatures) and [Multisignatures](./signatures.md#multisignatures). All objects in the following examples use msgpack to store the transaction object ensuring interoperability with the SDKs and `goal`.

!!! info
    Storing keys _offline_ is also referred to as placing them in **cold storage**. An _online_ device that stores private keys is often referred to as a **hot wallet**.  

# Unsigned Transaction File Operations
Algorand SDK's and `goal` support writing and reading both signed and unsigned transactions to a file. Examples of these scenarios are shown in the following code snippets.

Unsigned transactions require the transaction object to be created before writing to a file.


``` javascript tab="JavaScript"
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
    let confirmedTxn = await waitForConfirmation(algodClient, tx.txId, 4);
    var string = new TextDecoder().decode(confirmedTxn.txn.txn.note);
    console.log("Note field: ", string);       
```

``` python tab="Python"
    # build transaction
    params = algod_client.suggested_params()
    # comment out the next two (2) lines to use suggested fees
    params.flat_fee = True
    params.fee = 1000
    receiver = "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A"
    note = "Hello World".encode()
    unsigned_txn = PaymentTxn(
        my_address, params, receiver, 1000000, None, note)
    # write to file
    dir_path = os.path.dirname(os.path.realpath(__file__))
    transaction.write_to_file([unsigned_txn], dir_path + "/unsigned.txn")
	
    # read from file
    dir_path = os.path.dirname(os.path.realpath(__file__))
    txns = transaction.retrieve_from_file(dir_path + "/unsigned.txn")
    # sign and submit transaction
    txn = txns[0]
    signed_txn = txn.sign(mnemonic.to_private_key(passphrase))
    txid = signed_txn.transaction.get_txid()
    print("Signed transaction with txID: {}".format(txid))	
    try:
        txid = algod_client.send_transaction(signed_txn)
        # wait for confirmation              
        confirmed_txn = wait_for_confirmation(algod_client, txid, 4)
    except Exception as err:
        print(err)
        return
    print("Transaction information: {}".format(
        json.dumps(confirmed_txn, indent=4)))
    print("Decoded note: {}".format(base64.b64decode(
        confirmed_txn["txn"]["txn"]["note"]).decode()))  
```

``` java tab="Java"
    // Construct the transaction
    final String RECEIVER = "L5EUPCF4ROKNZMAE37R5FY2T5DF2M3NVYLPKSGWTUKVJRUGIW4RKVPNPD4";
    String note = "Hello World";
    Response < TransactionParametersResponse > resp = client.TransactionParams().execute();
    if (!resp.isSuccessful()) {
        throw new Exception(resp.message());
    }
    TransactionParametersResponse params = resp.body();
    if (params == null) {
        throw new Exception("Params retrieval error");
    }
    System.out.println("Algorand suggested parameters: " + params);
    Transaction tx = Transaction.PaymentTransactionBuilder()
        .sender(myAddress)
        .note(note.getBytes())
        .amount(100000)
        .receiver(new Address(RECEIVER))
        .suggestedParams(params)
        .build();

    // save as signed even though it has not been
    SignedTransaction stx = myAccount.signTransaction(tx);
    System.out.println("Signed transaction with txid: " + stx.transactionID);
    stx.tx = tx;  
    // Save transaction to a file 
    Files.write(Paths.get("./unsigned.txn"), Encoder.encodeToMsgPack(stx));
    System.out.println("Transaction written to a file");

    // read transaction from file
    SignedTransaction decodedTransaction = Encoder.decodeFromMsgPack(
        Files.readAllBytes(Paths.get("./unsigned.txn")), SignedTransaction.class);            
    Transaction tx = decodedTransaction.tx;           
    // Sign the transaction
    SignedTransaction signedTxn = myAccount.signTransaction(tx);
    System.out.println("Signed transaction with txid: " + signedTxn.transactionID);
    // Submit the transaction to the network
    byte[] encodedTxBytes = Encoder.encodeToMsgPack(signedTxn);
    Response < PostTransactionsResponse > rawtxresponse = client.RawTransaction().rawtxn(encodedTxBytes).execute();
    if (!rawtxresponse.isSuccessful()) {
        throw new Exception(rawtxresponse.message());
    }
    String id = rawtxresponse.body().txId;
    System.out.println("Successfully sent tx with ID: " + id);
    // Wait for transaction confirmation
    PendingTransactionResponse pTrx = waitForConfirmation(client, id, 4);
    System.out.println("Transaction " + id + " confirmed in round " + pTrx.confirmedRound);
    // Read the transaction
    JSONObject jsonObj = new JSONObject(pTrx.toString());
    System.out.println("Transaction information (with notes): " + jsonObj.toString(2));
    System.out.println("Decoded note: " + new String(pTrx.txn.tx.note));
    printBalance(myAccount);
```

``` go tab="Go"
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
	confirmedTxn, err := waitForConfirmation(txID, algodClient, 4)
	if err != nil {
		fmt.Printf("Error waiting for confirmation on txID: %s\n", txID)
		return
	}
	fmt.Printf("Decoded note: %s\n", string(confirmedTxn.Transaction.Txn.Note))
```


``` goal tab="goal"
$ goal clerk send --from=<my-account> --to=GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A --fee=1000 --amount=1000000 --out="unsigned.txn"

$ goal clerk sign --infile unsigned.txn --outfile signed.txn

$ goal clerk rawsend --filename signed.txn

```
# Signed Transaction File Operations 
Signed Transactions are similar, but require an account to sign the transaction before writing it to a file.

``` javascript tab="JavaScript"
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

``` python tab="Python"
    # build transaction
    params = algod_client.suggested_params()
    # comment out the next two (2) lines to use suggested fees
    params.flat_fee = True
    params.fee = 1000
    receiver = "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A"
    note = "Hello World".encode()
    unsigned_txn = PaymentTxn(
        my_address, params, receiver, 1000000, None, note)
    # sign transaction
    signed_txn = unsigned_txn.sign(mnemonic.to_private_key(passphrase))
    # write to file
    dir_path = os.path.dirname(os.path.realpath(__file__))
    transaction.write_to_file([signed_txn], dir_path + "/signed.txn")

    # read signed transaction from file
    dir_path = os.path.dirname(os.path.realpath(__file__))
    txns = transaction.retrieve_from_file(dir_path + "/signed.txn")
    signed_txn = txns[0]
    try:
        txid = algod_client.send_transaction(signed_txn)
        print("Signed transaction with txID: {}".format(txid))
    # wait for confirmation
        confirmed_txn = wait_for_confirmation(algod_client, txid, 4)
    except Exception as err:
        print(err)
        return
    print("Transaction information: {}".format(
        json.dumps(confirmed_txn, indent=4)))
    print("Decoded note: {}".format(base64.b64decode(
        confirmed_txn["txn"]["txn"]["note"]).decode())) 
```

``` java tab="Java"
    // Construct the transaction
    final String RECEIVER = "L5EUPCF4ROKNZMAE37R5FY2T5DF2M3NVYLPKSGWTUKVJRUGIW4RKVPNPD4";
    String note = "Hello World";
    Response < TransactionParametersResponse > resp = client.TransactionParams().execute();
    if (!resp.isSuccessful()) {
        throw new Exception(resp.message());
    }
    TransactionParametersResponse params = resp.body();
    if (params == null) {
        throw new Exception("Params retrieval error");
    }
    System.out.println("Algorand suggested parameters: " + params);
    Transaction txn = Transaction.PaymentTransactionBuilder()
        .sender(myAddress)
        .note(note.getBytes())
        .amount(100000)
        .receiver(new Address(RECEIVER))
        .suggestedParams(params)
        .build();

    // Sign the transaction
    SignedTransaction signedTx = myAccount.signTransaction(txn);
    System.out.println("Signed transaction with txid: " + signedTx.transactionID);
    // save signed transaction to  a file 
    Files.write(Paths.get("./signed.txn"), Encoder.encodeToMsgPack(signedTx));

    // read signed transaction
    SignedTransaction decodedSignedTransaction = Encoder.decodeFromMsgPack(
    Files.readAllBytes(Paths.get("./signed.txn")), SignedTransaction.class);     
    // Msgpack encode the signed transaction
    byte[] encodedTxBytes = Encoder.encodeToMsgPack(decodedSignedTransaction);
    // Submit the transaction to the network          
    Response < PostTransactionsResponse > rawtxresponse = client.RawTransaction().rawtxn(encodedTxBytes).execute();
    if (!rawtxresponse.isSuccessful()) {
        throw new Exception(rawtxresponse.message());
    }
    String id = rawtxresponse.body().txId;
    System.out.println("Successfully sent tx with ID: " + id);
    // Wait for transaction confirmation
    PendingTransactionResponse pTrx = waitForConfirmation(client, id, 4);
    System.out.println("Transaction " + id + " confirmed in round " + pTrx.confirmedRound);
    // Read the transaction
    JSONObject jsonObj = new JSONObject(pTrx.toString());
    System.out.println("Transaction information (with notes): " + jsonObj.toString(2));
    System.out.println("Decoded note: " + new String(pTrx.txn.tx.note));
    printBalance(myAccount); 
```

``` go tab="Go"
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
	confirmedTxn, err := waitForConfirmation(sendResponse, algodClient, 4)
	if err != nil {
		fmt.Printf("Error waiting for confirmation on txID: %s\n", sendResponse)
		return
	}
	fmt.Printf("Decoded note: %s\n", string(confirmedTxn.Transaction.Txn.Note))
```

``` goal tab="goal"
$ goal clerk rawsend --filename signed.txn
```
!!! info
    Example transaction code snippets are provided throughout this page. Full running code transaction examples as well as **offline multisig** for each SDK are available within the GitHub repo at [/examples/offline](https://github.com/algorand/docs/tree/master/examples/offline) and for [download](https://github.com/algorand/docs/blob/master/examples/offine/offline.zip?raw=true) (.zip).
    
??? example "Saving Signed and Unsigned Multisig Transactions to a File using goal"
    
    ```java tab="goal"
    Create a multisig account by listing all of the accounts in the multisig and specifying the threshold number of accounts to sign with the -T flag
    $ goal account multisig new <my-account1> <my-account2> <my-account3> etc… -T 2    

    Create an unsigned transaction and write to file
    $ goal clerk send --from <my-multisig-account>  --to AZLR2XP4O2WFHLX6TX7AZVY23HLVLG3K5K3FRIKIYDOYN6ISIF54SA4RNY --fee=1000 --amount=1000000 --out="unsigned.txn"

    Sign by the required number of accounts to meet the threshold. 
    $ goal clerk multisig sign -a F <my-account1> -t=unsigned.txn
    goal clerk multisig sign -a F <my-account2> -t=unsigned.txn

    Merge signings 
    $ goal clerk multisig merge --out signed.txn unsigned.txn

    Broadcast 
    $ goal clerk rawsend --filename signed.txn
    ```
