title: Creating Assets

REVIEW IN PROGRESS

The Algorand protocol supports the creation of on-chain assets that benefit from the same security, compatibility, speed and ease of use as the Algo. The official name for assets on Algorand is **Algorand Standard Assets (ASA)**.

Examples of assets you can represent include stablecoins, loyalty points, system credits, and in-game points, just to name a few. You can also represent single, unique assets like a deed for a house, collectable items, unique parts on a supply chain, and countless more. 

!!! info
    Assets that represent many of the same type, like a stablecoin, may be referred to as **fungible assets**. Single, unique assets are referred to as **non-fungible assets**.

Finally, assets can be configured with certain restrictions that help support securities, compliance, and certification use cases. 

# Asset Requirements
A single Algorand account is permitted to create up to 1000 assets. For every asset and account creates or owns, its minimum balance is increased by 0.1 Algos. Before a new asset can be transfered to a specific account the receiver must option in to receive the asset. This process is described below in Receiving an Asset. If any transaction is issued that would violate the maximum number of assets for an account or not meet the minimum balance requirements, the transaction will fail.

# Asset Parameters
The type of asset that is created will depend on the parameters that are passed during asset creation or configuration. The primary Asset parameters are listed below. Once set these parameters can never be changed.


* The creator address specifies the address of the asset creator. The transaction to create the asset must be signed by this account.
* The name parameter represents the name of the entire asset. This is a string containing up to 32 characters. For example, a currency name may be used as the asset name. 
* The unitname parameter is used to give a specific name to a unit of the asset. This is an up to eight character name for the asset. If a currency was referenced for the name, the unitname may contain is ISO code or symbol.
* The defaultfrozen parameter specifies whether the asset is initially frozen, meaning transactions in this asset are currently not allowed for any account until unfrozen by the asset freeze account. More details on freezing and unfreezing an asset are shown below.
* The total property specifies the maximum amount of this specific asset type. This number can not be changed once the asset is created. For non-fungible assets, this total will be most likely be one, where fungible assets will have much larger total counts. 
* The decimals parameter specifies the divisibility of an asset. A value of 0 represents an asset that is not divisible, while a value of 1 represents an asset that is divisible into tenths and so on, i.e, the number of digits to display after the decimal place when displaying the asset. This value must be between 0 and 19. 
* The assetmetadatab64 property allows an optional hash commitment of some sort relating to the asset. You could use it to specify specific asset details like title deed location, court record, or specification location. 
* The asseturl parameter provides an option to reference a link that includes more details of the asset off-chain. 

Each Asset also has a set of addresses that are associated with the Asset that allow additional functionality.

 |Address| **Description** | **Required** | **Note** | 
:-- |:-------------:| :-------------: | :-------------: | 
**Manager**  | Address of the manager of the asset | Yes | Can be set to blank 
**Reserve**  | Address of the reserve account of the asset | Yes | Can be set to blank
**Freeze** | Address of the freeze account for the asset | Yes | Can be set to blank
**Clawback**  | Address of the clawback account for the asset | Yes | Can be set to blank

These addresses represent the only properites of an asset that can be changed after creation. They can also be locked for the life of the asset as described below.

* Manager Address - The manager account is the only account that can manipulate the configuration or deletion of an asset. An asset can only be deleted if the entire asset supply is returned to the creator address. 
* Reserve Address - The reserve account can be used to act as a primary reserve account for minting. By default the assets are created in the creator account. If the reserve account is used as a reserve all tokens must be transfered from the creator to the reserve account. This also requires the reserve account to be setup to recieve the new asset. See Receiving an Asset described below.
* Freeze Address - The freeze account is the only account that is allowed to freeze or unfreeze specific accounts.  When an account is frozen it cannot send or receive transactions for that asset.
* Clawback Address - The clawback address is an account that is allowed to take assets from any account and give them to another account.

If any of these four addresses is set to `""` that address will be cleared and can never be reset for the life of the asset. This will also effectively disable the feature of that address. For example setting the freeze address to `""` will prevent the asset from ever being frozen.

Fungible assets will have asset totals of more than one, where non-fungilbe tokens will usually have a count of one. For restricted assets the proper feeze, clawback and manager accounts will be set, where if they are unrestricted these addresses will be set to `""` forever preventing any modification to ownership of the asset.


# Creating an Asset
Assets can be created using either the SDKs or `goal`. When using the SDKs all creation parameters can be supplied. With `goal`, managing the various addresses associated with the asset must be done after executing an asset creation. See Modifying an Asset in the next section for more details on changing addesses for the asset.

``` javascript tab="JavaScript" 

    let addr = recoveredAccount1.addr; 
    let defaultFrozen = false;
    let totalIssuance = 100; 
    let unitName = "<var>unit-name</var>"; 
    let assetName = "<var>asset-name</var>"; 
    let assetURL = "http://someurl"; 
    let assetMetadataHash = "16efaa3924a6fd9d3a4824799a4ac65d"; 
    let manager = recoveredAccount2.addr; 
    let reserve = recoveredAccount2.addr;
    let freeze = recoveredAccount2.addr; 
    let clawback = recoveredAccount2.addr; 
    let decimals = 0;

    // signing and sending "txn" allows "addr" to create an asset
    let txn = algosdk.makeAssetCreateTxn(addr, cp.fee, cp.firstRound, cp.lastRound, note,
        cp.genHash, cp.genID, totalIssuance, decimals, defaultFrozen, manager, reserve, freeze, clawback,
        unitName, assetName, assetURL, assetMetadataHash);

    let rawSignedTxn = txn.signTxn(recoveredAccount1.sk)
    let tx = (await algodclient.sendRawTransaction(rawSignedTxn));
    console.log("Transaction : " + tx.txId);
    await waitForConfirmation(algodclient, tx.txId);
    let ptx = await algodclient.pendingTransactionInformation(tx.txId);
    let assetID = ptx.txresults.createdasset;
```

``` python tab="Python"  
    # Configure fields for creating the asset.
    data = {
        "sender": accounts[1]['pk'],
        "fee": min_fee,
        "first": first,
        "last": last,
        "gh": gh,
        "total": 1000,
        "decimals": 0,
        "default_frozen": False,
        "unit_name": "LATINUM",
        "asset_name": "latinum",
        "manager": accounts[1]['pk'],
        "reserve": accounts[1]['pk'],
        "freeze": accounts[1]['pk'],
        "clawback": accounts[1]['pk'],
        "url": "https://path/to/my/asset/details",
        "flat_fee": True
    }

    # Construct Asset Creation transaction
    txn = transaction.AssetConfigTxn(**data)

    # Sign with secret key of creator
    stxn = txn.sign(accounts[1]['sk'])

    print("Asset Creation")
    # Send the transaction to the network and retrieve the txid.
    txid = algod_client.send_transaction(stxn, headers={'content-type': 'application/x-binary'})
    # Retrieve the asset ID of the newly created asset by first
    # ensuring that the creation transaction was confirmed,
    # then pulling account info of the creator and grabbing the 
    # asset with the max asset ID. 
    # Wait for the transaction to be confirmed
    txinfo = wait_for_confirmation(txid)
    print(txinfo.keys())
    print(txinfo)
    asset_id = txinfo["txresults"]["createdasset"]
    account_info = algod_client.account_info(accounts[1]['pk'])
```

``` java tab="Java"  
       // Create the Asset:
        BigInteger assetTotal = BigInteger.valueOf(10000);
        boolean defaultFrozen = false;
        String unitName = "LTRCEX1";
        String  assetName = "ltrc-teset";
        String url = "http://this.test.com";
        String assetMetadataHash = "16efaa3924a6fd9d3a4824799a4ac65d";
        Address manager  = acct2.getAddress();
        Address reserve = acct2.getAddress();
        Address freeze = acct2.getAddress();
        Address clawback = acct2.getAddress();
        Integer decimals = 0;
        Transaction tx = Transaction.createAssetCreateTransaction(acct1.getAddress(), 
            BigInteger.valueOf( 1000 ), cp.firstRound, cp.lastRound, null, cp.genID, 
            cp.genHash, assetTotal, decimals, defaultFrozen, unitName, assetName, url, 
            assetMetadataHash.getBytes(), manager, reserve, freeze, clawback);
        // Update the fee as per what the BlockChain is suggesting
        Account.setFeeByFeePerByte(tx, cp.fee);

        // Sign the Transaction with creator account
        SignedTransaction signedTx = acct1.signTransaction(tx);
        BigInteger assetID = null;
        try{
            TransactionID id = ex.submitTransaction( signedTx );
            System.out.println( "Transaction ID: " + id );
            ex.waitForConfirmation( signedTx.transactionID);
            // Now that the transaction is confirmed we can get the assetID
            com.algorand.algosdk.algod.client.model.Transaction ptx = 
                algodApiInstance.pendingTransactionInformation(id.getTxId());
            assetID = ptx.getTxresults().getCreatedasset();

        } catch (Exception e){
            e.printStackTrace();
            return;
        }
        System.out.println( "AssetID = " +  assetID);
```

``` go tab="Go"  
	// Create an asset
	// Set parameters for asset creation transaction
	creator := pks[1]
	assetName := "latinum"
	unitName := "latinum"
	assetURL := "https://path/to/my/asset/details"
	assetMetadataHash := "thisIsSomeLength32HashCommitment"
	defaultFrozen := false
    decimals := uint32(0)
	totalIssuance := uint64(1000)
	manager := pks[1]
	reserve := pks[1]
	freeze := pks[1]
	clawback := pks[1]
	note := []byte(nil)
	txn, err := transaction.MakeAssetCreateTxn(creator, fee, firstRound, lastRound, note,
	genID, genHash, totalIssuance, decimals, defaultFrozen, manager, reserve, freeze, clawback,
	unitName, assetName, assetURL, assetMetadataHash)
	if err != nil {
		fmt.Printf("Failed to make asset: %s\n", err)
		return
	}
	fmt.Printf("Asset created AssetName: %s\n", txn.AssetConfigTxnFields.AssetParams.AssetName)
	
	txid, stx, err := crypto.SignTransaction(sks[1], txn)
	if err != nil {
		fmt.Printf("Failed to sign transaction: %s\n", err)
		return
	}
	fmt.Printf("Transaction ID: %s\n", txid)
		// Broadcast the transaction to the network
	sendResponse, err := algodClient.SendRawTransaction(stx)
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	
	// Wait for transaction to be confirmed
	waitForConfirmation(algodClient, sendResponse.TxID)
		
	// Retrieve asset ID by grabbing the max asset ID
	// from the creator account's holdings. 
	act, err := algodClient.AccountInformation(pks[1], txHeaders...)
	if err != nil {
		fmt.Printf("failed to get account information: %s\n", err)
		return
	}
	assetID := uint64(0)
	for i := range act.AssetParams {
		if i > assetID {
			assetID = i
		}
	}
	fmt.Printf("Asset ID from AssetParams: %d\n", assetID)
	// Retrieve asset info.
	assetInfo, err := algodClient.AssetInformation(assetID, txHeaders...)
```

``` goal tab="Goal"  
goal asset create --creator <var>address</var> --total 1000 --unitname <var>unit-name</var> --asseturl "https://path/to/my/asset/details" --decimals 0   -d data
```

# Modifying an Asset
After an asset has been created only the manager, reserve, freeze and reserve accounts can be changed. All other parametes are locked for the life of the asset. If any of these addresses are set to `""` that address will be cleared and can never be reset for the life of the asset. Only the manager account can make configuration changes and must sting the transaction.

``` javascript tab="JavaScript"
    // change the manager for the asset
    manager = recoveredAccount1.addr;

    // Note that the change has to come from the existing manager
    let ctxn = algosdk.makeAssetConfigTxn(recoveredAccount2.addr, cp.fee, 
        cp.firstRound, cp.lastRound, note, cp.genHash, cp.genID,
        assetID, manager, reserve, freeze, clawback);

    // This transaction must be signed by the current manager
    rawSignedTxn = ctxn.signTxn(recoveredAccount2.sk)
    let ctx = (await algodclient.sendRawTransaction(rawSignedTxn));
    console.log("Transaction : " + ctx.txId);
    // wait for transaction to be confirmed
    await waitForConfirmation(algodclient, ctx.txId);
 
    //Get the asset information for the newly changed asset
    let assetInfo = await algodclient.assetInformation(assetID);
    //The manager should now be the same as the creator
    console.log(assetInfo);  
```

``` python tab="Python"  
    # Update manager address.
    # Keep reserve, freeze, and clawback address same as before, i.e. account 1
    data = {
        "sender": accounts[1]['pk'],
        "fee": min_fee,
        "first": first,
        "last": last,
        "gh": gh,
        "index": asset_id,
        "manager": accounts[2]['pk'],
        "reserve": accounts[1]['pk'],
        "freeze": accounts[1]['pk'],
        "clawback": accounts[1]['pk'],
        "flat_fee": True
    }
    txn = transaction.AssetConfigTxn(**data)
    stxn = txn.sign(accounts[1]['sk'])
    print("Asset Modification")
    txid = algod_client.send_transaction(stxn, headers={'content-type': 'application/x-binary'})
    print(txid)

    # Wait for the transaction to be confirmed
    pendinginfo = algod_client.pending_transaction_info(txid)
    while pendinginfo['round'] == 0:
        pendinginfo = algod_client.pending_transaction_info(txid)

    # Check asset info to view change in management.
    asset_info = algod_client.asset_info(asset_id)
    print(json.dumps(asset_info, indent=4))
```

``` java tab="Java"  
    // configuration changes must be done by
    // the manager account - changing manager of the asset
    tx = Transaction.createAssetConfigureTransaction(acct2.getAddress(), 
            BigInteger.valueOf( 1000 ),cp.firstRound, cp.lastRound, null, 
            cp.genID, cp.genHash, assetID, acct1.getAddress(), reserve, 
            freeze, clawback, false);
    // update the fee as per what the BlockChain is suggesting
    Account.setFeeByFeePerByte(tx, cp.fee);
    // the transaction must be signed by the current manager account   
    signedTx = acct2.signTransaction(tx);
    // send the transaction to the network
    try{
        TransactionID id = ex.submitTransaction( signedTx );
        System.out.println( "Transaction ID: " + id );
        ex.waitForConfirmation( signedTx.transactionID);
    } catch (Exception e){
        e.printStackTrace();
        return;
    }  
    // list the asset
    AssetParams assetInfo = algodApiInstance.assetInformation(assetID);
    // The manager should now be the same as the creator
    System.out.println(assetInfo);
```

``` go tab="Go"  
	// Change Asset Manager from Account 1 to Account 2
	manager = pks[2]
	oldmanager := pks[1]

	txn, err = transaction.MakeAssetConfigTxn(oldmanager, fee, 
		firstRound, lastRound, note, genID, genHash, assetID, 
		manager, reserve, freeze, clawback, true)
	if err != nil {
        fmt.Printf("Failed to send txn: %s\n", err)
        return
    }
    txid, stx, err = crypto.SignTransaction(sks[1], txn)
    if err != nil {
        fmt.Printf("Failed to sign transaction: %s\n", err)
        return
	}
    fmt.Printf("Transaction ID: %s\n", txid)
    // Broadcast the transaction to the network
    sendResponse, err = algodClient.SendRawTransaction(stx)
    if err != nil {
        fmt.Printf("failed to send transaction: %s\n", err)
        return
    }
    fmt.Printf("Transaction ID raw: %s\n", sendResponse.TxID)

	// Wait for transaction to be confirmed
	waitForConfirmation(algodClient, sendResponse.TxID)
	// Retrieve asset info.
	assetInfo, err = algodClient.AssetInformation(assetID, txHeaders...)
	// Print asset info showing updated manager address.
	PrettyPrint(assetInfo)
```

``` goal tab="Goal"  
goal asset config  --manager <var>address</var> --new-reserve <var>address</var> --assetid <var>asset-id</var> -d data 
```

# Receiving an Asset
Before any account can receive a created asset it must option in to receiving it. This requires a transaction and the account minimum rises by 0.1 Algo. The account optioning in to receiving the asset type must send this transaction to themselves. This involves setting the sender and receiver of the transaction to be the same address. The following code illustrates this transaction.

``` javascript tab="JavaScript"  
    // Opting in to an Asset:
    // Transaction from and sender must be the same
    let sender = recoveredAccount3.addr;
    let recipient = sender;
    let revocationTarget = undefined;
    let closeRemainderTo = undefined;
    // We are sending 0 of new assets
    amount = 0;

    // update changing transaction parameters
    await getChangingParms(algodclient);

    // signing and sending "txn" allows sender to begin accepting asset specified by assetid
    let opttxn = algosdk.makeAssetTransferTxn(sender, recipient, closeRemainderTo, revocationTarget,
        cp.fee, amount, cp.firstRound, cp.lastRound, note, cp.genHash, cp.genID, assetID);

    // Must be signed by the account wishing to opt in to the asset    
    rawSignedTxn = opttxn.signTxn(recoveredAccount3.sk);
    let opttx = (await algodclient.sendRawTransaction(rawSignedTxn));
    console.log("Transaction : " + opttx.txId);
    // wait for transaction to be confirmed
    await waitForConfirmation(algodclient, opttx.txId);

    // the new asset listed in the account information
    act = await algodclient.accountInformation(recoveredAccount3.addr);
    console.log("Account Information for: " + JSON.stringify(act.assets));
```

``` python tab="Python"  
    # Check if asset_id is in account 2's asset holdings prior to opt-in
    account_info = algod_client.account_info(accounts[2]['pk'])
    holding = None
    if 'assets' in account_info:
        holding = account_info['assets'].get(str(asset_id))

    if not holding:
        # Get latest network parameters
        data = {
            "sender": accounts[2]['pk'],
            "fee": min_fee,
            "first": first,
            "last": last,
            "gh": gh,
            "receiver": accounts[2]["pk"],
            "amt": 0,
            "index": asset_id,
            "flat_fee": True
        }
        print("Asset Option In")
        # Use the AssetTransferTxn class to transfer assets
        txn = transaction.AssetTransferTxn(**data)
        stxn = txn.sign(accounts[2]['sk'])
        txid = algod_client.send_transaction(stxn, headers={'content-type': 'application/x-binary'})
        print(txid)
        # Wait for the transaction to be confirmed
        wait_for_confirmation(txid)
        # Now check the asset holding for that account. 
        # This should now show a holding with a balance of 0.
        account_info = algod_client.account_info(accounts[2]['pk'])
        print(json.dumps(account_info['assets'][str(asset_id)], indent=4))
```

``` java tab="Java"  
    // Opt in to Receiving the Assetn
    try {
        cp = ex.getChangingParms(algodApiInstance);
    } catch (ApiException e) {
        e.printStackTrace();
        return;
        }
    tx = Transaction.createAssetAcceptTransaction(acct3.getAddress(), 
        BigInteger.valueOf( 1000 ), cp.firstRound, 
        cp.lastRound, null, cp.genID, cp.genHash, assetID);
    // Update the fee based on the network suggested fee
    Account.setFeeByFeePerByte(tx, cp.fee);
    // The transaction must be signed by the current manager account  
    signedTx = acct3.signTransaction(tx);
    com.algorand.algosdk.algod.client.model.Account act;
	// send the transaction to the network and
    try{
        TransactionID id = ex.submitTransaction( signedTx );
        System.out.println( "Transaction ID: " + id );
        ex.waitForConfirmation( signedTx.transactionID);
        // We can now list the account information for acct3 
        // and see that it can accept the new asseet
        act = algodApiInstance.accountInformation(acct3.getAddress().toString());
        AssetHolding ah = act.getHolding(assetID);
        System.out.println( "Asset Holding: " + ah.getAmount() );

    } catch (Exception e){
        e.printStackTrace();
        return;
    }  
```

``` go tab="Go"  
	// Account 3 opts in to receive asset
	txn, err = transaction.MakeAssetAcceptanceTxn(pks[3], fee, firstRound, 
		lastRound, note, genID, genHash, assetID)
	if err != nil {
        fmt.Printf("Failed to send transaction MakeAssetAcceptanceTxn: %s\n", err)
        return
    }
    txid, stx, err = crypto.SignTransaction(sks[3], txn)
    if err != nil {
        fmt.Printf("Failed to sign transaction: %s\n", err)
        return
	}

    fmt.Printf("Transaction ID: %s\n", txid)
    // Broadcast the transaction to the network
    sendResponse, err = algodClient.SendRawTransaction(stx)
    if err != nil {
        fmt.Printf("failed to send transaction: %s\n", err)
        return
    }
    fmt.Printf("Transaction ID raw: %s\n", sendResponse.TxID)

	// Wait for transaction to be confirmed
	waitForConfirmation(algodClient, sendResponse.TxID)

	act, err = algodClient.AccountInformation(pks[3], txHeaders...)
    if err != nil {
        fmt.Printf("failed to get account information: %s\n", err)
        return
	}
	PrettyPrint(act.Assets[assetID])
```

``` goal tab="Goal"  
goal asset send -a 0 --asset <var>asset-name</var>  -f <var>opt-in-account</var> -t <var>opt-in-account</var> --creator <var>asset-creator</var>  -d data
```

# Transferring an Asset
Assets can be transferred between accounts that have optioned into receiving the asset and function similar to standard payment transactions that are signed by the sender account.

``` javascript tab="JavaScript"  
    // Transfer New Asset:
    sender = recoveredAccount1.addr;
    recipient = recoveredAccount3.addr;
    revocationTarget = undefined;
    closeRemainderTo = undefined;
    // amount of the asset to transfer
    amount = 10;

    // update changing transaction parameters
    await getChangingParms(algodclient);

    // signing and sending "txn"
    let xtxn = algosdk.makeAssetTransferTxn(sender, recipient, 
        closeRemainderTo, revocationTarget,cp.fee, amount, 
        cp.firstRound, cp.lastRound, note, cp.genHash, cp.genID, assetID);
    // Must be signed by the account sending the asset  
    rawSignedTxn = xtxn.signTxn(recoveredAccount1.sk)
    let xtx = (await algodclient.sendRawTransaction(rawSignedTxn));
    console.log("Transaction : " + xtx.txId);
    
    // wait for transaction to be confirmed
    await waitForConfirmation(algodclient, xtx.txId);

    // 10 assets listed in the account information
    act = await algodclient.accountInformation(recoveredAccount3.addr);
    console.log("Account Information for: " + JSON.stringify(act.assets));
```

``` python tab="Python"  
    # send 10 
    data = {
        "sender": accounts[1]['pk'],
        "fee": min_fee,
        "first": first,
        "last": last,
        "gh": gh,
        "receiver": accounts[2]["pk"],
        "amt": 10,
        "index": asset_id,
        "flat_fee": True
    }
    print("Asset Transfer")
    txn = transaction.AssetTransferTxn(**data)
    stxn = txn.sign(accounts[1]['sk'])
    txid = algod_client.send_transaction(stxn, headers={'content-type': 'application/x-binary'})
    print(txid)
    # Wait for the transaction to be confirmed
    wait_for_confirmation(txid)
    # The balance should now be 10.
    account_info = algod_client.account_info(accounts[2]['pk'])
    print(json.dumps(account_info['assets'][str(asset_id)], indent=4))
```

``` java tab="Java"  
    // set asset xfer specific parameters
    // We set the assetCloseTo to null so we do not close the asset out
    Address assetCloseTo = new Address();
    BigInteger assetAmount = BigInteger.valueOf(10);
    tx = Transaction.createAssetTransferTransaction(acct1.getAddress(), 
        acct3.getAddress(), assetCloseTo, assetAmount, BigInteger.valueOf( 1000 ), 
        cp.firstRound, cp.lastRound, null, cp.genID, cp.genHash, assetID);        
    // Update the fee based on the network suggested fee
    Account.setFeeByFeePerByte(tx, cp.fee);
    // The transaction must be signed by the sender account  
    signedTx = acct1.signTransaction(tx);
    // send the transaction to the network 
    try{
        TransactionID id = ex.submitTransaction( signedTx );
        System.out.println( "Transaction ID: " + id );
        ex.waitForConfirmation( signedTx.transactionID);
        // list the account information for acct3 
        // and see that it now has 5 of the new asset
        act = algodApiInstance.accountInformation(acct3.getAddress().toString());
        System.out.println( act.getHolding(assetID).getAmount() );
    } catch (Exception e){
        e.printStackTrace();
        return;
    }        
```

``` go tab="Go"  
	// Send  10 of asset from Account 1 to Account 3
	sender := pks[1]
	recipient := pks[3]
	amount := uint64(10)
	closeRemainderTo := ""
	txn, err = transaction.MakeAssetTransferTxn(sender, recipient, 
		closeRemainderTo, amount, fee, firstRound, lastRound, note,
        genID, genHash, assetID)
	if err != nil {
        fmt.Printf("Failed to send transaction MakeAssetTransfer Txn: %s\n", err)
        return
    }
    txid, stx, err = crypto.SignTransaction(sks[1], txn)
    if err != nil {
        fmt.Printf("Failed to sign transaction: %s\n", err)
        return
	}
    fmt.Printf("Transaction ID: %s\n", txid)
    // Broadcast the transaction to the network
    sendResponse, err = algodClient.SendRawTransaction(stx)
    if err != nil {
        fmt.Printf("failed to send transaction: %s\n", err)
        return
    }
    fmt.Printf("Transaction ID raw: %s\n", sendResponse.TxID)

	// Wait for transaction to be confirmed
	waitForConfirmation(algodClient, sendResponse.TxID)

	act, err = algodClient.AccountInformation(pks[3], txHeaders...)
    if err != nil {
        fmt.Printf("failed to get account information: %s\n", err)
        return
	}
	PrettyPrint(act.Assets[assetID])
```

``` goal tab="Goal"  
goal asset send -a <var>asset-amount</var> --asset <var>asset-name</var> -f <var>asset-sender</var> -t <var>asset-receiver</var> --creator <var>asset-creator</var> -d data
```

# Freezing an Asset
Assets configured with a freeze account can be frozen. If the freeze address is ever set to `""` by the asset manager, this capability is removed from the asset. Freezing an asset for an account prevents that account from either receiving or sending the specific asset. Freezing or unfreezing and asset for an account requires a transaction that is signed by the freeze account. The code below illustrates the freeze transaction.

``` javascript tab="JavaScript"  
    // The asset was created and 
    // configured to allow freezing an account
    from = recoveredAccount2.addr;
    freezeTarget = recoveredAccount3.addr;
    freezeState = true;

    // update changing transaction parameters
    await getChangingParms(algodclient);

    // The freeze transaction needs to be signed by the freeze account
    let ftxn = algosdk.makeAssetFreezeTxn(from, cp.fee, 
        cp.firstRound, cp.lastRound, note, cp.genHash, cp.genID,
        assetID, freezeTarget, freezeState)
  
    rawSignedTxn = ftxn.signTxn(recoveredAccount2.sk)
    let ftx = (await algodclient.sendRawTransaction(rawSignedTxn));
    console.log("Transaction : " + ftx.txId);
    // wait for transaction to be confirmed
    await waitForConfirmation(algodclient, ftx.txId);

    // the asset is frozen listed in the account information
    act = await algodclient.accountInformation(recoveredAccount3.addr);
    console.log("Account Information for: " + JSON.stringify(act.assets));
```

``` python tab="Python"  
    # Freezing an Asset
    data = {
        "sender": accounts[1]['pk'],
        "fee": min_fee,
        "first": first,
        "last": last,
        "gh": gh,
        "index": asset_id,
        "target": accounts[2]['pk'],
        "new_freeze_state": True,
        "flat_fee": True
    }
    print("Asset Freeze")
    txn = transaction.AssetFreezeTxn(**data)
    stxn = txn.sign(accounts[1]['sk'])
    txid = algod_client.send_transaction(stxn, headers={'content-type': 'application/x-binary'})
    print(txid)
    # Wait for the transaction to be confirmed
    wait_for_confirmation(txid)
    # The balance should now be 10.
    account_info = algod_client.account_info(accounts[2]['pk'])
    print(json.dumps(account_info['assets'][str(asset_id)], indent=4))
```

``` java tab="Java"  
    // Freeze the Asset:
    // The asset was created and configured to allow freezing an account
    try {
        cp = ex.getChangingParms(algodApiInstance);
    } catch (ApiException e) {
        e.printStackTrace();
        return;
    }
    // set asset specific parameters
    boolean freezeState = true;
    // The sender should be freeze account
    tx = Transaction.createAssetFreezeTransaction(acct2.getAddress(), 
        acct3.getAddress(), freezeState, BigInteger.valueOf( 1000 ), 
        cp.firstRound, cp.lastRound, null, cp.genHash, assetID);
    // Update the fee based on the network suggested fee
    Account.setFeeByFeePerByte(tx, cp.fee);
    // The transaction must be signed by the freeze account   
    signedTx = acct2.signTransaction(tx);
    // send the transaction to the network
    try{
        TransactionID id = ex.submitTransaction( signedTx );
        System.out.println( "Transaction ID: " + id );
        ex.waitForConfirmation( signedTx.transactionID);
    } catch (Exception e){
        e.printStackTrace();
        return;
    }
```

``` go tab="Go"  
	// Freeze asset for Account 3.
	newFreezeSetting := true
	target := pks[3]
	txn, err = transaction.MakeAssetFreezeTxn(freeze, fee, firstRound, 
		lastRound, note, genID, genHash, assetID, target, 
		newFreezeSetting)
	if err != nil {
        fmt.Printf("Failed to send txn: %s\n", err)
        return
    }
    txid, stx, err = crypto.SignTransaction(sks[1], txn)
    if err != nil {
        fmt.Printf("Failed to sign transaction: %s\n", err)
        return
	}
    fmt.Printf("Transaction ID: %s\n", txid)
    // Broadcast the transaction to the network
    sendResponse, err = algodClient.SendRawTransaction(stx)
    if err != nil {
        fmt.Printf("failed to send transaction: %s\n", err)
        return
    }
	fmt.Printf("Transaction ID raw: %s\n", sendResponse.TxID)
	// Wait for transaction to be confirmed
	waitForConfirmation(algodClient, sendResponse.TxID)

	act, err = algodClient.AccountInformation(pks[3], txHeaders...)
    if err != nil {
        fmt.Printf("failed to get account information: %s\n", err)
        return
	}
	PrettyPrint(act.Assets[assetID])
```

``` goal tab="Goal"  
goal asset freeze --freezer <var>asset-freeze-account</var> --freeze=true --account <var>account-to-freeze</var> --creator <var>asset-creator</var> --asset <var>asset-name</var> -d data
```

# Revoking an Asset
Assets configured with a clawback account can be revoked. If the clawback address is ever set to `""` by the asset manager, this capability is removed from the asset. Revoking an asset for an account removes a specific number of the asset from the revoke target account. Revoking an asset from an account requires a transaction that is signed by the clawback account which also specifies the ammount of the asset to revoke. The code below illustrates the clawback transaction.

``` javascript tab="JavaScript"  
    // Revoke and Asset:
    // sender must be clawback
    sender = recoveredAccount2.addr;
    recipient = recoveredAccount1.addr;
    revocationTarget = recoveredAccount3.addr;
    closeRemainderTo = undefined;
    // amount of asset to revoke
    amount = 10;
   
    // update changing transaction parameters
    await getChangingParms(algodclient);

    // signing and sending from clawback address
    let rtxn = algosdk.makeAssetTransferTxn(sender, 
        recipient, closeRemainderTo, revocationTarget,
        cp.fee, amount, cp.firstRound, cp.lastRound, 
        note, cp.genHash, cp.genID, assetID);
    // Must be signed by the account that is the clawback address    
    rawSignedTxn = rtxn.signTxn(recoveredAccount2.sk)
    let rtx = (await algodclient.sendRawTransaction(rawSignedTxn));
    console.log("Transaction : " + rtx.txId);
    // wait for transaction to be confirmed
    await waitForConfirmation(algodclient, rtx.txId);

    // 0 assets listed in the account information
    console.log("Asset ID: " + assetID);
    act = await algodclient.accountInformation(recoveredAccount3.addr);
    console.log("Account Information for: " + JSON.stringify(act.assets));
```

``` python tab="Python"  
    # Revoking an Asset
    data = {
        "sender": accounts[1]['pk'],
        "fee": min_fee,
        "first": first,
        "last": last,
        "gh": gh,
        "receiver": accounts[1]["pk"],
        "amt": 10,
        "index": asset_id,
        "revocation_target": accounts[2]['pk'],
        "flat_fee": True
    }
    print("Asset Revoke")
    txn = transaction.AssetTransferTxn(**data)
    stxn = txn.sign(accounts[1]['sk'])
    txid = algod_client.send_transaction(stxn, headers={'content-type': 'application/x-binary'})
    print(txid)
    # Wait for the transaction to be confirmed
    wait_for_confirmation(txid)
    # The balance of account 2 should now be 0.
    account_info = algod_client.account_info(accounts[2]['pk'])
    print(json.dumps(account_info['assets'][str(asset_id)], indent=4))
    # The balance of account 1 should increase by 10 to 1000.
    account_info = algod_client.account_info(accounts[1]['pk'])
    print(json.dumps(account_info['assets'][str(asset_id)], indent=4))
```

``` java tab="Java"  
    // Revoke the asset:
    // The asset was also created with the ability for it to be revoked by 
    // clawbackaddress. 
    try {
        cp = ex.getChangingParms(algodApiInstance);
    } catch (ApiException e) {
        e.printStackTrace();
        return;
    }
    // set asset specific parameters
    assetAmount = BigInteger.valueOf( 10 );
    tx = Transaction.createAssetRevokeTransaction(acct2.getAddress(), 
        acct3.getAddress(), acct1.getAddress(), assetAmount, 
        BigInteger.valueOf( 1000 ), cp.firstRound, 
    cp.lastRound, null, cp.genID, cp.genHash, assetID);
    // Update the fee based on the network suggested fee
    Account.setFeeByFeePerByte(tx, cp.fee);
    // The transaction must be signed by the clawback account  
    signedTx = acct2.signTransaction(tx);
    // send the transaction to the network and
    // wait for the transaction to be confirmed
    try{
        TransactionID id = ex.submitTransaction( signedTx );
        System.out.println( "Transaction ID: " + id );
        ex.waitForConfirmation( signedTx.transactionID);
        // list the account information 
        act = algodApiInstance.accountInformation(acct3.getAddress().toString());
        System.out.println( act.getHolding(assetID).getAmount() );
    } catch (Exception e){
        e.printStackTrace();
        return;
    }  
```

``` go tab="Go"  
	// Revoke an asset
	// The clawback account (Account 1) revokes 10 from Account 3.
	target = pks[3]
	txn, err = transaction.MakeAssetRevocationTxn(clawback, target, creator, amount, fee, firstRound, lastRound, note,
	genID, genHash, assetID)
	if err != nil {
        fmt.Printf("Failed to send txn: %s\n", err)
        return
    }
    txid, stx, err = crypto.SignTransaction(sks[1], txn)
    if err != nil {
        fmt.Printf("Failed to sign transaction: %s\n", err)
        return
	}
    fmt.Printf("Transaction ID: %s\n", txid)
    // Broadcast the transaction to the network
    sendResponse, err = algodClient.SendRawTransaction(stx)
    if err != nil {
        fmt.Printf("failed to send transaction: %s\n", err)
        return
    }
	fmt.Printf("Transaction ID raw: %s\n", sendResponse.TxID)
	// Wait for transaction to be confirmed
	waitForConfirmation(algodClient, sendResponse.TxID)

	act, err = algodClient.AccountInformation(pks[3], txHeaders...)
    if err != nil {
        fmt.Printf("failed to get account information: %s\n", err)
        return
	}
	PrettyPrint(act.Assets[assetID])
```

``` goal tab="Goal"  
goal asset send -a <var>amount-to-revoke</var> --asset <var>asset-name</var> -f <var>address-of-revoke-target</var> -t <var>address-to-send-assets-to</var> --clawback <var>clawback-address</var> --creator <var>creator-address</var> -d data
```

# Destroying an Asset
Created assets can be destroyed only by the asset manager account. All of the assets must be owned by the creator of the asset before the asset can be deleted. Deleting an asset requires a transaction that is signed by the manager as shown below.

``` javascript tab="JavaScript"  
    // Destroy and Asset:
    // All of the created assets should now be back in the creators
    // Account so we can delete the asset.
    await getChangingParms(algodclient);


    // The address for the from field 
    // must be the manager account
    addr = recoveredAccount1.addr;

    // if all assets are held by the asset creator,
    // the asset creator can sign and issue "txn" 
    // to remove the asset from the ledger. 
    let dtxn = algosdk.makeAssetDestroyTxn(addr, cp.fee,
         cp.firstRound, cp.lastRound, note, cp.genHash,
        cp.genID, assetID);
    // The transaction must be signed by the manager
    rawSignedTxn = dtxn.signTxn(recoveredAccount1.sk)
    let dtx = (await algodclient.sendRawTransaction(rawSignedTxn));
    console.log("Transaction : " + dtx.txId);
    // wait for transaction to be confirmed
    await waitForConfirmation(algodclient, dtx.txId);

    // asset is not shown in account information    
    act = await algodclient.accountInformation(recoveredAccount3.addr);
    console.log("Account Information for: " + JSON.stringify(act.assets));
```

``` python tab="Python"  
    # Destroy Asset
    data = {
        "sender": accounts[2]['pk'],
        "fee": min_fee,
        "first": first,
        "last": last,
        "gh": gh,
        "index": asset_id,
        "flat_fee": True,
        "strict_empty_address_check": False
    }
    print("Destroying Asset")
    # Construct Asset Creation transaction
    txn = transaction.AssetConfigTxn(**data)

    # Sign with secret key of creator
    stxn = txn.sign(accounts[2]['sk'])

    # Send the transaction to the network and retrieve the txid.
    txid = algod_client.send_transaction(stxn, headers={'content-type': 'application/x-binary'})
    print(txid)

    # Wait for the transaction to be confirmed
    wait_for_confirmation(txid)

    # This should raise an exception since the asset was deleted.
    try:
        asset_info = algod_client.asset_info(asset_id)
    except Exception as e:
        print(e)
```

``` java tab="Java"  
    // Destroy the Asset:
    // All assets should now be back in
    // creators account
    try {
        cp = ex.getChangingParms(algodApiInstance);
    } catch (ApiException e) {
        e.printStackTrace();
        return;
    }
    // set asset specific parameters
    // The manager must sign and submit the transaction
    tx = Transaction.createAssetDestroyTransaction(acct1.getAddress(), 
            BigInteger.valueOf( 1000 ), cp.firstRound, cp.lastRound, 
            null, cp.genHash, assetID);
    // Update the fee based on the network suggested fee
    Account.setFeeByFeePerByte(tx, cp.fee);
    // The transaction must be signed by the manager account  
    signedTx = acct1.signTransaction(tx);
    // send the transaction to the network 
    try{
        TransactionID id = ex.submitTransaction( signedTx );
        System.out.println( "Transaction ID: " + id );
        ex.waitForConfirmation( signedTx.transactionID);
        // We list the account information for acct1 
        // and check that the asset is no longer exist
        act = algodApiInstance.accountInformation(acct1.getAddress().toString());
        System.out.println( "Does AssetID: " + assetID + " exist? " + 
        act.getThisassettotal().containsKey(assetID) );
    } catch (Exception e){
        e.printStackTrace();
        return;
    } 
```

``` go tab="Go"  
	// Destroy the asset
	// all funds are back in the creator's account.
	// Manager account used to destroy the asset.
	txn, err = transaction.MakeAssetDestroyTxn(manager, fee, 
		firstRound, lastRound, note, genID, genHash, assetID)
	if err != nil {
        fmt.Printf("Failed to send txn: %s\n", err)
        return
    }
    txid, stx, err = crypto.SignTransaction(sks[2], txn)
    if err != nil {
        fmt.Printf("Failed to sign transaction: %s\n", err)
        return
	}
    fmt.Printf("Transaction ID: %s\n", txid)
    // Broadcast the transaction to the network
    sendResponse, err = algodClient.SendRawTransaction(stx)
    if err != nil {
        fmt.Printf("failed to send transaction: %s\n", err)
        return
    }
	fmt.Printf("Transaction ID raw: %s\n", sendResponse.TxID)
	// Wait for transaction to be confirmed
	waitForConfirmation(algodClient, sendResponse.TxID)
	// Retrieve asset info. This should now throw an error.
	assetInfo, err = algodClient.AssetInformation(assetID, txHeaders...)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
```

``` goal tab="Goal"  
goal asset destroy --creator <var>creator-address</var> --manager <var>asset-manager-address</var> --asset <var>asset-name</var> -d data 
```

# Get Asset Configuration Information
An asset's configuration information can be retreived from the network using the SDKs or `goal`. Additional details are also added to the accounts that own the specific asset and can be listed with standard account information calls.

``` javascript tab="JavaScript"
    //Get the asset information for anasset
    let assetInfo = await algodclient.assetInformation(assetID);
    console.log(assetInfo);
```

```python tab="Python"
    # Check asset info.
    asset_info = algod_client.asset_info(asset_id)
    print(json.dumps(asset_info, indent=4))
```

```java tab="Java"
    // list the asset
    AssetParams assetInfo = algodApiInstance.assetInformation(assetID);
    System.out.println(assetInfo);
```

```go tab="Go"
    assetInfo, err = algodClient.AssetInformation(assetID, txHeaders...)
    // Print asset info
    PrettyPrint(assetInfo)
```

``` goal tab="Goal"  
goal asset info --creator <var>creator-address</var> --asset unitname  -d ~/node/data -w testwall
Asset ID:         <var>created-asset-id</var>
Creator:          <var>creator-address</var>
Asset name:       testtoken
Unit name:        unitname
Maximum issue:    12 unitname
Reserve amount:   12 unitname
Issued:           0 unitname
Decimals:         0
Default frozen:   false
Manager address:  <var>creator-address</var>
Reserve address:  <var>reserve-address</var>
Freeze address:   <var>freeze-address</var>
Clawback address: <var>clawback-address</var>
```

??? example "Complete Example = Asset Options"
    
    ```javascript tab="JavaScript"
    const algosdk = require('algosdk');

    //Retrieve the token, server and port values for your installation in the algod.net
    //and algod.token files within the data directory
    const token = "<var>algod-address</var>";
    const server = "<var>algod-token</var>";
    const port = <var>port-number</var>;

    // Structure for changing blockchain params
    var cp = {
        fee: 0, 
        firstRound: 0,  
        lastRound: 0, 
        genID: "",
        genHash: ""    
    }
    //Utility function to update params from blockchain
    var getChangingParms = async function( algodclient ) {
        let params = await algodclient.getTransactionParams();
        cp.firstRound = params.lastRound;
        cp.lastRound = cp.firstRound + parseInt(1000);
        let sfee = await algodclient.suggestedFee();
        cp.fee = sfee.fee;
        cp.genID = params.genesisID;
        cp.genHash = params.genesishashb64;
    }
    // Function used to wait for a tx confirmation
    var waitForConfirmation = async function(algodclient, txId) {
        while (true) {
            b3 = await algodclient.pendingTransactionInformation(txId);
            if (b3.round != null && b3.round > 0) {
                //Got the completed Transaction
                console.log("Transaction " + b3.tx + " confirmed in round " + b3.round);
                break;
            }
        }
    };

    //Recover accounts used in example
    var account1_mnemonic ="<var>your-25-word-mnemonic</var>";    
    var account2_mnemonic ="<var>your-25-word-mnemonic</var>";    
    var account3_mnemonic ="<var>your-25-word-mnemonic</var>";
    var recoveredAccount1 = algosdk.mnemonicToSecretKey(account1_mnemonic);
    var recoveredAccount2 = algosdk.mnemonicToSecretKey(account2_mnemonic);
    var recoveredAccount3 = algosdk.mnemonicToSecretKey(account3_mnemonic);
    console.log(recoveredAccount1.addr);
    console.log(recoveredAccount2.addr);
    console.log(recoveredAccount3.addr);


    //instantiate the algod wrapper
    let algodclient = new algosdk.Algod(token, server, port);
    //submit the transaction
    (async() => {
        
        // Asset Creation:
        await getChangingParms(algodclient);
        let note = undefined; 

        // create the asset
        let addr = recoveredAccount1.addr; 
        let defaultFrozen = false;
        let totalIssuance = 100; 
        let unitName = "<var>asset-unit-name</var>"; 
        let assetName = "<var>asset-name</var>"; 
        let assetURL = "http://someurl"; 
        let assetMetadataHash = "16efaa3924a6fd9d3a4824799a4ac65d"; 
        let manager = recoveredAccount2.addr; 
        let reserve = recoveredAccount2.addr;
        let freeze = recoveredAccount2.addr; 
        let clawback = recoveredAccount2.addr; 
        let decimals = 0;

        // signing and sending "txn" allows "addr" to create an asset
        let txn = algosdk.makeAssetCreateTxn(addr, cp.fee, cp.firstRound, cp.lastRound, note,
            cp.genHash, cp.genID, totalIssuance, decimals, defaultFrozen, manager, reserve, freeze, clawback,
            unitName, assetName, assetURL, assetMetadataHash);

        let rawSignedTxn = txn.signTxn(recoveredAccount1.sk)
        let tx = (await algodclient.sendRawTransaction(rawSignedTxn));
        console.log("Transaction : " + tx.txId);

        // wait for transaction to be confirmed and get the assetid
        await waitForConfirmation(algodclient, tx.txId);
        let ptx = await algodclient.pendingTransactionInformation(tx.txId);
        console.log( ptx.txresults.createdasset );
        let assetID = ptx.txresults.createdasset;


        // Change Asset Configuration:
        // update network parameters
        await getChangingParms(algodclient);

        // Asset configuration specific parameters
        // change the manager for the asset
        manager = recoveredAccount1.addr;

        // Note that the change has to come from the existing manager
        let ctxn = algosdk.makeAssetConfigTxn(recoveredAccount2.addr, cp.fee, 
            cp.firstRound, cp.lastRound, note, cp.genHash, cp.genID,
            assetID, manager, reserve, freeze, clawback);

        // This transaction must be signed by the current manager
        rawSignedTxn = ctxn.signTxn(recoveredAccount2.sk)
        let ctx = (await algodclient.sendRawTransaction(rawSignedTxn));
        console.log("Transaction : " + ctx.txId);
        // wait for transaction to be confirmed
        await waitForConfirmation(algodclient, ctx.txId);
    
        //Get the asset information for the newly changed asset
        let assetInfo = await algodclient.assetInformation(assetID);
        console.log(assetInfo);

        // Opting in to an Asset:
        // Transaction from and sender must be the same
        let sender = recoveredAccount3.addr;
        let recipient = sender;
        let revocationTarget = undefined;
        let closeRemainderTo = undefined;
        // We are sending 0 of new assets
        amount = 0;

        // update changing transaction parameters
        await getChangingParms(algodclient);

        // signing and sending "txn" allows sender to begin accepting asset specified by assetid
        let opttxn = algosdk.makeAssetTransferTxn(sender, recipient, closeRemainderTo, revocationTarget,
            cp.fee, amount, cp.firstRound, cp.lastRound, note, cp.genHash, cp.genID, assetID);

        // Must be signed by the account wishing to opt in to the asset    
        rawSignedTxn = opttxn.signTxn(recoveredAccount3.sk);
        let opttx = (await algodclient.sendRawTransaction(rawSignedTxn));
        console.log("Transaction : " + opttx.txId);
        // wait for transaction to be confirmed
        await waitForConfirmation(algodclient, opttx.txId);

        // the new asset listed in the account information
        act = await algodclient.accountInformation(recoveredAccount3.addr);
        console.log("Account Information for: " + JSON.stringify(act.assets));

        // Transfer New Asset:
        sender = recoveredAccount1.addr;
        recipient = recoveredAccount3.addr;
        revocationTarget = undefined;
        closeRemainderTo = undefined;
        // amount of the asset to transfer
        amount = 10;

        // update changing transaction parameters
        await getChangingParms(algodclient);

        // signing and sending "txn" will send "amount" assets from "sender" to "recipient"
        let xtxn = algosdk.makeAssetTransferTxn(sender, recipient, 
            closeRemainderTo, revocationTarget,cp.fee, amount, 
            cp.firstRound, cp.lastRound, note, cp.genHash, cp.genID, assetID);
        // Must be signed by the account sending the asset  
        rawSignedTxn = xtxn.signTxn(recoveredAccount1.sk)
        let xtx = (await algodclient.sendRawTransaction(rawSignedTxn));
        console.log("Transaction : " + xtx.txId);
        
        // wait for transaction to be confirmed
        await waitForConfirmation(algodclient, xtx.txId);

        // 10 assets listed in the account information
        act = await algodclient.accountInformation(recoveredAccount3.addr);
        console.log("Account Information for: " + JSON.stringify(act.assets));

        // The asset was created and 
        // configured to allow freezing an account
        from = recoveredAccount2.addr;
        freezeTarget = recoveredAccount3.addr;
        freezeState = true;

        // update changing transaction parameters
        await getChangingParms(algodclient);

        // The freeze transaction needs to be signed by the freeze account
        let ftxn = algosdk.makeAssetFreezeTxn(from, cp.fee, 
            cp.firstRound, cp.lastRound, note, cp.genHash, cp.genID,
            assetID, freezeTarget, freezeState)
    
        rawSignedTxn = ftxn.signTxn(recoveredAccount2.sk)
        let ftx = (await algodclient.sendRawTransaction(rawSignedTxn));
        console.log("Transaction : " + ftx.txId);
        // wait for transaction to be confirmed
        await waitForConfirmation(algodclient, ftx.txId);

        // the asset is frozen listed in the account information
        act = await algodclient.accountInformation(recoveredAccount3.addr);
        console.log("Account Information for: " + JSON.stringify(act.assets));

        // Revoke and Asset:
        // sender must be clawback
        sender = recoveredAccount2.addr;
        recipient = recoveredAccount1.addr;
        revocationTarget = recoveredAccount3.addr;
        closeRemainderTo = undefined;
        // amount of asset to revoke
        amount = 10;
    
        // update changing transaction parameters
        await getChangingParms(algodclient);

        // signing and sending from clawback address
        let rtxn = algosdk.makeAssetTransferTxn(sender, 
            recipient, closeRemainderTo, revocationTarget,
            cp.fee, amount, cp.firstRound, cp.lastRound, 
            note, cp.genHash, cp.genID, assetID);
        // Must be signed by the account that is the clawback address    
        rawSignedTxn = rtxn.signTxn(recoveredAccount2.sk)
        let rtx = (await algodclient.sendRawTransaction(rawSignedTxn));
        console.log("Transaction : " + rtx.txId);
        // wait for transaction to be confirmed
        await waitForConfirmation(algodclient, rtx.txId);

        // 0 assets listed in the account information
        console.log("Asset ID: " + assetID);
        act = await algodclient.accountInformation(recoveredAccount3.addr);
        console.log("Account Information for: " + JSON.stringify(act.assets));

        // Destroy and Asset:
        // All of the created assets should now be back in the creators
        // Account so we can delete the asset.
        await getChangingParms(algodclient);


        // The address for the from field 
        // must be the manager account
        addr = recoveredAccount1.addr;

        // if all assets are held by the asset creator,
        // the asset creator can sign and issue "txn" 
        // to remove the asset from the ledger. 
        let dtxn = algosdk.makeAssetDestroyTxn(addr, cp.fee,
            cp.firstRound, cp.lastRound, note, cp.genHash,
            cp.genID, assetID);
        // The transaction must be signed by the manager
        rawSignedTxn = dtxn.signTxn(recoveredAccount1.sk)
        let dtx = (await algodclient.sendRawTransaction(rawSignedTxn));
        console.log("Transaction : " + dtx.txId);
        // wait for transaction to be confirmed
        await waitForConfirmation(algodclient, dtx.txId);

        // asset is not shown in account information    
        act = await algodclient.accountInformation(recoveredAccount3.addr);
        console.log("Account Information for: " + JSON.stringify(act.assets));


    })().catch(e => {
        console.log(e);
        console.trace();
    });
    ```

    ```python tab="Python"
    import json
    import time
    from algosdk import account, algod, mnemonic, transaction

    # Shown for demonstration purposes. NEVER reveal secret mnemonics in practice. 
    # Change these values if you want to use different accounts.
    mnemonic1 = "<var>25-word-passphrase</var>"
    mnemonic2 = "<var>25-word-passphrase</var>"
    mnemonic3 = "<var>25-word-passphrase</var>"

    # For ease of reference, add account public and private keys to 
    # an accounts dict.
    accounts = {}
    counter = 1
    for m in [mnemonic1, mnemonic2, mnemonic3]:
        accounts[counter] = {}
        accounts[counter]['pk'] = mnemonic.to_public_key(m)
        accounts[counter]['sk'] = mnemonic.to_private_key(m)
        counter += 1

    # Specify your node address and token. This must be updated.
    algod_address = "<var>algod-address</var>"
    algod_token = "<var>algod-token</var>"
    algod_client = algod.AlgodClient(algod_token, algod_address)

    # Get network params for transaction
    params = algod_client.suggested_params()
    first = params.get("lastRound")
    last = first + 1000
    gen = params.get("genesisID")
    gh = params.get("genesishashb64")
    min_fee = params.get("minFee")

    # utility for waiting on a transaction confirmation
    def wait_for_confirmation(  txid ):
        while True:
            txinfo = algod_client.pending_transaction_info(txid)
            if txinfo.get('round') and txinfo.get('round') > 0:
                print("Transaction {} confirmed in round {}.".format(txid, txinfo.get('round')))
                break
            else:
                print("Waiting for confirmation...")
                algod_client.status_after_block(algod_client.status().get('lastRound') +1)

        return txinfo
                
    # Configure fields for creating the asset.
    data = {
        "sender": accounts[1]['pk'],
        "fee": min_fee,
        "first": first,
        "last": last,
        "gh": gh,
        "total": 1000,
        "decimals": 0,
        "default_frozen": False,
        "unit_name": "LATINUM",
        "asset_name": "latinum",
        "manager": accounts[1]['pk'],
        "reserve": accounts[1]['pk'],
        "freeze": accounts[1]['pk'],
        "clawback": accounts[1]['pk'],
        "url": "https://path/to/my/asset/details",
        "flat_fee": True
    }

    # Construct Asset Creation transaction
    txn = transaction.AssetConfigTxn(**data)

    # Sign with secret key of creator
    stxn = txn.sign(accounts[1]['sk'])

    print("Asset Creation")
    # Send the transaction to the network and retrieve the txid.
    txid = algod_client.send_transaction(stxn, headers={'content-type': 'application/x-binary'})
    # Retrieve the asset ID of the newly created asset by first
    # ensuring that the creation transaction was confirmed,
    # then pulling account info of the creator and grabbing the 
    # asset with the max asset ID. 
    # Wait for the transaction to be confirmed
    txinfo = wait_for_confirmation(txid)
    print(txinfo.keys())
    print(txinfo)
    asset_id = txinfo["txresults"]["createdasset"]
    account_info = algod_client.account_info(accounts[1]['pk'])


    # Update manager address.
    # Keep reserve, freeze, and clawback address same as before, i.e. account 1
    data = {
        "sender": accounts[1]['pk'],
        "fee": min_fee,
        "first": first,
        "last": last,
        "gh": gh,
        "index": asset_id,
        "manager": accounts[2]['pk'],
        "reserve": accounts[1]['pk'],
        "freeze": accounts[1]['pk'],
        "clawback": accounts[1]['pk'],
        "flat_fee": True
    }
    txn = transaction.AssetConfigTxn(**data)
    stxn = txn.sign(accounts[1]['sk'])
    print("Asset Modification")
    txid = algod_client.send_transaction(stxn, headers={'content-type': 'application/x-binary'})
    print(txid)

    # Wait for the transaction to be confirmed
    wait_for_confirmation(txid)

    # Check asset info to view change in management.
    asset_info = algod_client.asset_info(asset_id)
    print(json.dumps(asset_info, indent=4))

    # Check if asset_id is in account 2's asset holdings prior to opt-in
    account_info = algod_client.account_info(accounts[2]['pk'])
    holding = None
    if 'assets' in account_info:
        holding = account_info['assets'].get(str(asset_id))

    if not holding:
        # Get latest network parameters
        data = {
            "sender": accounts[2]['pk'],
            "fee": min_fee,
            "first": first,
            "last": last,
            "gh": gh,
            "receiver": accounts[2]["pk"],
            "amt": 0,
            "index": asset_id,
            "flat_fee": True
        }
        print("Asset Option In")
        # Use the AssetTransferTxn class to transfer assets
        txn = transaction.AssetTransferTxn(**data)
        stxn = txn.sign(accounts[2]['sk'])
        txid = algod_client.send_transaction(stxn, headers={'content-type': 'application/x-binary'})
        print(txid)
        # Wait for the transaction to be confirmed
        wait_for_confirmation(txid)
        # Now check the asset holding for that account. 
        # This should now show a holding with a balance of 0.
        account_info = algod_client.account_info(accounts[2]['pk'])
        print(json.dumps(account_info['assets'][str(asset_id)], indent=4))

    # send 10 
    data = {
        "sender": accounts[1]['pk'],
        "fee": min_fee,
        "first": first,
        "last": last,
        "gh": gh,
        "receiver": accounts[2]["pk"],
        "amt": 10,
        "index": asset_id,
        "flat_fee": True
    }
    print("Asset Transfer")
    txn = transaction.AssetTransferTxn(**data)
    stxn = txn.sign(accounts[1]['sk'])
    txid = algod_client.send_transaction(stxn, headers={'content-type': 'application/x-binary'})
    print(txid)
    # Wait for the transaction to be confirmed
    wait_for_confirmation(txid)
    # The balance should now be 10.
    account_info = algod_client.account_info(accounts[2]['pk'])
    print(json.dumps(account_info['assets'][str(asset_id)], indent=4))

    # Freezing an Asset
    data = {
        "sender": accounts[1]['pk'],
        "fee": min_fee,
        "first": first,
        "last": last,
        "gh": gh,
        "index": asset_id,
        "target": accounts[2]['pk'],
        "new_freeze_state": True,
        "flat_fee": True
    }
    print("Asset Freeze")
    txn = transaction.AssetFreezeTxn(**data)
    stxn = txn.sign(accounts[1]['sk'])
    txid = algod_client.send_transaction(stxn, headers={'content-type': 'application/x-binary'})
    print(txid)
    # Wait for the transaction to be confirmed
    wait_for_confirmation(txid)
    # The balance should now be 10.
    account_info = algod_client.account_info(accounts[2]['pk'])
    print(json.dumps(account_info['assets'][str(asset_id)], indent=4))

    # Revoking an Asset
    data = {
        "sender": accounts[1]['pk'],
        "fee": min_fee,
        "first": first,
        "last": last,
        "gh": gh,
        "receiver": accounts[1]["pk"],
        "amt": 10,
        "index": asset_id,
        "revocation_target": accounts[2]['pk'],
        "flat_fee": True
    }
    print("Asset Revoke")
    txn = transaction.AssetTransferTxn(**data)
    stxn = txn.sign(accounts[1]['sk'])
    txid = algod_client.send_transaction(stxn, headers={'content-type': 'application/x-binary'})
    print(txid)
    # Wait for the transaction to be confirmed
    wait_for_confirmation(txid)
    # The balance of account 2 should now be 0.
    account_info = algod_client.account_info(accounts[2]['pk'])
    print(json.dumps(account_info['assets'][str(asset_id)], indent=4))
    # The balance of account 1 should increase by 10 to 1000.
    account_info = algod_client.account_info(accounts[1]['pk'])
    print(json.dumps(account_info['assets'][str(asset_id)], indent=4))

    # Destroy Asset
    data = {
        "sender": accounts[2]['pk'],
        "fee": min_fee,
        "first": first,
        "last": last,
        "gh": gh,
        "index": asset_id,
        "flat_fee": True,
        "strict_empty_address_check": False
    }
    print("Destroying Asset")
    # Construct Asset Creation transaction
    txn = transaction.AssetConfigTxn(**data)

    # Sign with secret key of creator
    stxn = txn.sign(accounts[2]['sk'])

    # Send the transaction to the network and retrieve the txid.
    txid = algod_client.send_transaction(stxn, headers={'content-type': 'application/x-binary'})
    print(txid)

    # Wait for the transaction to be confirmed
    wait_for_confirmation(txid)

    # This should raise an exception since the asset was deleted.
    try:
        asset_info = algod_client.asset_info(asset_id)
    except Exception as e:
        print(e)    

    ```

    ```java tab="Java"
    package com.algorand.algosdk.asset;

    import java.math.BigInteger;

    import com.algorand.algosdk.account.Account;
    import com.algorand.algosdk.algod.client.AlgodClient;
    import com.algorand.algosdk.algod.client.ApiException;
    import com.algorand.algosdk.algod.client.api.AlgodApi;
    import com.algorand.algosdk.algod.client.auth.ApiKeyAuth;
    import com.algorand.algosdk.algod.client.model.AssetHolding;
    import com.algorand.algosdk.algod.client.model.AssetParams;
    import com.algorand.algosdk.algod.client.model.TransactionID;
    import com.algorand.algosdk.algod.client.model.TransactionParams;
    import com.algorand.algosdk.crypto.Address;
    import com.algorand.algosdk.crypto.Digest;
    import com.algorand.algosdk.transaction.SignedTransaction;
    import com.algorand.algosdk.transaction.Transaction;
    import com.algorand.algosdk.util.Encoder;

    /**
    * Show Creating, modifying, sending and listing assets 
    */
    public class AssetExample 
    {   

        public AlgodApi algodApiInstance = null;
        
        // utility function to connect to a node
        private AlgodApi connectToNetwork(){

            final String ALGOD_API_ADDR = "<var>algod-address</var>";
            final String ALGOD_API_TOKEN = "<var>algod-token</var>";

            AlgodClient client = (AlgodClient) new AlgodClient().setBasePath(ALGOD_API_ADDR);
            ApiKeyAuth api_key = (ApiKeyAuth) client.getAuthentication("api_key");
            api_key.setApiKey(ALGOD_API_TOKEN);
            algodApiInstance = new AlgodApi(client);   
            return algodApiInstance;
        }     
        // Inline class to handle changing block parameters
        // Throughout the example
        public class ChangingBlockParms
        {
            public BigInteger fee; 
            public BigInteger firstRound;  
            public BigInteger lastRound; 
            public String genID;
            public Digest genHash;
            public ChangingBlockParms() {
                this.fee = BigInteger.valueOf(0);
                this.firstRound = BigInteger.valueOf(0);
                this.lastRound = BigInteger.valueOf(0);
                this.genID = "";
                this.genHash = null;
            }
        };
        // Utility function to wait on a transaction to be confirmed    
        public void waitForConfirmation( String txID ) throws Exception{
            if( algodApiInstance == null ) connectToNetwork();
            while(true) {
                try {
                    //Check the pending tranactions
                    com.algorand.algosdk.algod.client.model.Transaction pendingInfo = algodApiInstance.pendingTransactionInformation(txID);
                    if (pendingInfo.getRound() != null && pendingInfo.getRound().longValue() > 0) {
                        //Got the completed Transaction
                        System.out.println("Transaction " + pendingInfo.getTx() + " confirmed in round " + pendingInfo.getRound().longValue());
                        break;
                    } 
                    algodApiInstance.waitForBlock(BigInteger.valueOf( algodApiInstance.getStatus().getLastRound().longValue() +1 ) );
                } catch (Exception e) {
                    throw( e );
                }
            }

        }

        // Utility function to update changing block parameters 
        public ChangingBlockParms getChangingParms(AlgodApi algodApiInstance) throws Exception{
            ChangingBlockParms cp = new AssetExample.ChangingBlockParms(); 
            try {
                TransactionParams params = algodApiInstance.transactionParams();
                cp.fee = params.getFee();
                cp.firstRound = params.getLastRound();
                cp.lastRound = cp.firstRound.add(BigInteger.valueOf(1000));
                cp.genID = params.getGenesisID();
                cp.genHash = new Digest(params.getGenesishashb64());

            } catch (ApiException e) {
            throw( e );
            }
            return( cp );
        }

        // Utility function for sending a raw signed transaction to the network
        public  TransactionID submitTransaction( SignedTransaction signedTx ) throws Exception{
            try {
                // Msgpack encode the signed transaction
                byte[] encodedTxBytes = Encoder.encodeToMsgPack(signedTx);
                TransactionID id = algodApiInstance.rawTransaction(encodedTxBytes);
                return( id );
            } catch (ApiException e) {
                throw( e );
            }
        }


        public static void main(String args[]) throws Exception {

            AssetExample ex = new AssetExample();
            AlgodApi algodApiInstance= ex.connectToNetwork();

            // recover example accounts
            
            final String account1_mnemonic = <var>your-25-word-mnemonic</var>             final String account2_mnemonic = <var>your-25-word-mnemonic</var>             final String account3_mnemonic = <var>your-25-word-mnemonic</var>                    
            Account acct1  = new Account(account1_mnemonic); 
            Account acct2  = new Account(account2_mnemonic);
            Account acct3  = new Account(account3_mnemonic);                           
            
            // get changing network parameters
            ChangingBlockParms cp = null;
            try {
                cp = ex.getChangingParms(algodApiInstance);
            } catch (ApiException e) {
                e.printStackTrace();
                return;
            }
            // Create the Asset:
            BigInteger assetTotal = BigInteger.valueOf(10000);
            boolean defaultFrozen = false;
            String unitName = "LTRCEX1";
            String  assetName = "ltrc-teset";
            String url = "http://this.test.com";
            String assetMetadataHash = "16efaa3924a6fd9d3a4824799a4ac65d";
            Address manager  = acct2.getAddress();
            Address reserve = acct2.getAddress();
            Address freeze = acct2.getAddress();
            Address clawback = acct2.getAddress();
            Integer decimals = 0;
            Transaction tx = Transaction.createAssetCreateTransaction(acct1.getAddress(), 
                BigInteger.valueOf( 1000 ), cp.firstRound, cp.lastRound, null, cp.genID, 
                cp.genHash, assetTotal, decimals, defaultFrozen, unitName, assetName, url, 
                assetMetadataHash.getBytes(), manager, reserve, freeze, clawback);
            // Update the fee as per what the BlockChain is suggesting
            Account.setFeeByFeePerByte(tx, cp.fee);

            // Sign the Transaction with creator account
            SignedTransaction signedTx = acct1.signTransaction(tx);
            BigInteger assetID = null;
            try{
                TransactionID id = ex.submitTransaction( signedTx );
                System.out.println( "Transaction ID: " + id );
                ex.waitForConfirmation( signedTx.transactionID);
                // Now that the transaction is confirmed we can get the assetID
                com.algorand.algosdk.algod.client.model.Transaction ptx = 
                    algodApiInstance.pendingTransactionInformation(id.getTxId());
                assetID = ptx.getTxresults().getCreatedasset();

            } catch (Exception e){
                e.printStackTrace();
                return;
            }
            System.out.println( "AssetID = " +  assetID);

            // Change Asset Configuration:
            // Get changing network parameters
            try {
                cp = ex.getChangingParms(algodApiInstance);
            } catch (ApiException e) {
                e.printStackTrace();
                return;
            }
            // configuration changes must be done by
            // the manager account - changing manager of the asset
            tx = Transaction.createAssetConfigureTransaction(acct2.getAddress(), 
                    BigInteger.valueOf( 1000 ),cp.firstRound, cp.lastRound, null, 
                    cp.genID, cp.genHash, assetID, acct1.getAddress(), reserve, 
                    freeze, clawback, false);
            // update the fee as per what the BlockChain is suggesting
            Account.setFeeByFeePerByte(tx, cp.fee);
            // the transaction must be signed by the current manager account   
            signedTx = acct2.signTransaction(tx);
            // send the transaction to the network
            try{
                TransactionID id = ex.submitTransaction( signedTx );
                System.out.println( "Transaction ID: " + id );
                ex.waitForConfirmation( signedTx.transactionID);
            } catch (Exception e){
                e.printStackTrace();
                return;
            }  
            // list the asset
            AssetParams assetInfo = algodApiInstance.assetInformation(assetID);
            // the manager should now be the same as the creator
            System.out.println(assetInfo);
        
            // Opt in to Receiving the Asset
            try {
                cp = ex.getChangingParms(algodApiInstance);
            } catch (ApiException e) {
                e.printStackTrace();
                return;
            }
            tx = Transaction.createAssetAcceptTransaction(acct3.getAddress(), 
                BigInteger.valueOf( 1000 ), cp.firstRound, 
                cp.lastRound, null, cp.genID, cp.genHash, assetID);
            // Update the fee based on the network suggested fee
            Account.setFeeByFeePerByte(tx, cp.fee);
            // The transaction must be signed by the current manager account  
            signedTx = acct3.signTransaction(tx);
            com.algorand.algosdk.algod.client.model.Account act;
            // send the transaction to the network and
            try{
                TransactionID id = ex.submitTransaction( signedTx );
                System.out.println( "Transaction ID: " + id );
                ex.waitForConfirmation( signedTx.transactionID);
                // We can now list the account information for acct3 
                // and see that it can accept the new asseet
                act = algodApiInstance.accountInformation(acct3.getAddress().toString());
                AssetHolding ah = act.getHolding(assetID);
                System.out.println( "Asset Holding: " + ah.getAmount() );

            } catch (Exception e){
                e.printStackTrace();
                return;
            }  

            // Transfer the Asset:
            // get changing network parameters
            try {
                cp = ex.getChangingParms(algodApiInstance);
            } catch (ApiException e) {
                e.printStackTrace();
                return;
            }
            // set asset xfer specific parameters
            // We set the assetCloseTo to null so we do not close the asset out
            Address assetCloseTo = new Address();
            BigInteger assetAmount = BigInteger.valueOf(10);
            tx = Transaction.createAssetTransferTransaction(acct1.getAddress(), 
                acct3.getAddress(), assetCloseTo, assetAmount, BigInteger.valueOf( 1000 ), 
                cp.firstRound, cp.lastRound, null, cp.genID, cp.genHash, assetID);        
            // Update the fee based on the network suggested fee
            Account.setFeeByFeePerByte(tx, cp.fee);
            // The transaction must be signed by the sender account  
            signedTx = acct1.signTransaction(tx);
            // send the transaction to the network 
            try{
                TransactionID id = ex.submitTransaction( signedTx );
                System.out.println( "Transaction ID: " + id );
                ex.waitForConfirmation( signedTx.transactionID);
                // list the account information for acct3 
                // and see that it now has 5 of the new asset
                act = algodApiInstance.accountInformation(acct3.getAddress().toString());
                System.out.println( act.getHolding(assetID).getAmount() );
            } catch (Exception e){
                e.printStackTrace();
                return;
            }
            
            // Freeze the Asset:
            // The asset was created and configured to allow freezing an account
            try {
                cp = ex.getChangingParms(algodApiInstance);
            } catch (ApiException e) {
                e.printStackTrace();
                return;
            }
            // set asset specific parameters
            boolean freezeState = true;
            // The sender should be freeze account
            tx = Transaction.createAssetFreezeTransaction(acct2.getAddress(), 
                acct3.getAddress(), freezeState, BigInteger.valueOf( 1000 ), 
                cp.firstRound, cp.lastRound, null, cp.genHash, assetID);
            // Update the fee based on the network suggested fee
            Account.setFeeByFeePerByte(tx, cp.fee);
            // The transaction must be signed by the freeze account   
            signedTx = acct2.signTransaction(tx);
            // send the transaction to the network
            try{
                TransactionID id = ex.submitTransaction( signedTx );
                System.out.println( "Transaction ID: " + id );
                ex.waitForConfirmation( signedTx.transactionID);
            } catch (Exception e){
                e.printStackTrace();
                return;
            }

            // Revoke the asset:
            // The asset was also created with the ability for it to be revoked by 
            // clawbackaddress. 
            try {
                cp = ex.getChangingParms(algodApiInstance);
            } catch (ApiException e) {
                e.printStackTrace();
                return;
            }
            // set asset specific parameters
            assetAmount = BigInteger.valueOf( 10 );
            tx = Transaction.createAssetRevokeTransaction(acct2.getAddress(), 
                acct3.getAddress(), acct1.getAddress(), assetAmount, 
                BigInteger.valueOf( 1000 ), cp.firstRound, 
            cp.lastRound, null, cp.genID, cp.genHash, assetID);
            // Update the fee based on the network suggested fee
            Account.setFeeByFeePerByte(tx, cp.fee);
            // The transaction must be signed by the clawback account  
            signedTx = acct2.signTransaction(tx);
            // send the transaction to the network and
            // wait for the transaction to be confirmed
            try{
                TransactionID id = ex.submitTransaction( signedTx );
                System.out.println( "Transaction ID: " + id );
                ex.waitForConfirmation( signedTx.transactionID);
                // list the account information 
                act = algodApiInstance.accountInformation(acct3.getAddress().toString());
                System.out.println( act.getHolding(assetID).getAmount() );
            } catch (Exception e){
                e.printStackTrace();
                return;
            }  

            // Destroy the Asset:
            // All assets should now be back in
            // creators account
            try {
                cp = ex.getChangingParms(algodApiInstance);
            } catch (ApiException e) {
                e.printStackTrace();
                return;
            }
            // set asset specific parameters
            // The manager must sign and submit the transaction
            tx = Transaction.createAssetDestroyTransaction(acct1.getAddress(), 
                    BigInteger.valueOf( 1000 ), cp.firstRound, cp.lastRound, 
                    null, cp.genHash, assetID);
            // Update the fee based on the network suggested fee
            Account.setFeeByFeePerByte(tx, cp.fee);
            // The transaction must be signed by the manager account  
            signedTx = acct1.signTransaction(tx);
            // send the transaction to the network 
            try{
                TransactionID id = ex.submitTransaction( signedTx );
                System.out.println( "Transaction ID: " + id );
                ex.waitForConfirmation( signedTx.transactionID);
                // We list the account information for acct1 
                // and check that the asset is no longer exist
                act = algodApiInstance.accountInformation(acct1.getAddress().toString());
                System.out.println( "Does AssetID: " + assetID + " exist? " + 
                    act.getThisassettotal().containsKey(assetID) );
            } catch (Exception e){
                e.printStackTrace();
                return;
            }  
        }
    }
    ```

    ```go tab="Go"
    package main

    import (
        "fmt"
        json "encoding/json"
        b64 "encoding/base64"
        "github.com/algorand/go-algorand-sdk/transaction"
        "github.com/algorand/go-algorand-sdk/client/algod"
        "github.com/algorand/go-algorand-sdk/mnemonic"
        "github.com/algorand/go-algorand-sdk/crypto"
    )
    const algodAddress = "<var>algod-address</var>"
    const algodToken = "<var>algod-token</var>"


    var txHeaders = append([]*algod.Header{}, &algod.Header{"Content-Type", "application/json"})

    // Accounts to be used through examples
    func loadAccounts() (map[int][]byte, map[int]string){
        var pks = map[int]string {
        	1: "<var>account1-address</var>",
        	2: "<var>account1-address</var>",
        	3: "<var>account1-address</var>",
        }
        mnemonic1 = "<var>your-25-word-mnemonic</var>"
        mnemonic2 = "<var>your-25-word-mnemonic</var>"
        mnemonic3 = "<var>your-25-word-mnemonic</var>"
        mnemonics := []string{mnemonic1, mnemonic2, mnemonic3}
        var sks = make(map[int][]byte) 
        for i, m := range mnemonics {
            var err error
            sks[i+1], err = mnemonic.ToPrivateKey(m)
            if err != nil {
                fmt.Printf("Issue with account %d private key conversion.", i+1)
            } else {
                fmt.Printf("Loaded Key %d: %s\n", i+1, pks[i+1])
            }
        }
        return sks, pks
    }

    // Function that waits for a given txId to be confirmed by the network
    func waitForConfirmation(algodClient algod.Client, txID string) {
        for {
            pt, err := algodClient.PendingTransactionInformation(txID)
            if err != nil {
                fmt.Printf("waiting for confirmation... (pool error, if any): %s\n", err)
                continue
            }
            if pt.ConfirmedRound > 0 {
                fmt.Printf("Transaction "+pt.TxID+" confirmed in round %d\n", pt.ConfirmedRound)
                break
            }
            nodeStatus, err := algodClient.Status()
            if err != nil {
                fmt.Printf("error getting algod status: %s\n", err)
                return
            }
            algodClient.StatusAfterBlock( nodeStatus.LastRound + 1)
        }
    }

    // PrettyPrint prints Go structs
    func PrettyPrint(data interface{}) {
        var p []byte
        //    var err := error
        p, err := json.MarshalIndent(data, "", "\t")
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Printf("%s \n", p)
    }

    // Main function to demonstrate ASA examples
    func main() {
        // Get pre-defined set of keys for example
        sks, pks := loadAccounts()
        
        // Initialize an algodClient
        algodClient, err := algod.MakeClient(algodAddress, algodToken)
        if err != nil {
            return
        }

        // Get network-related transaction parameters and assign
        txParams, err := algodClient.SuggestedParams()
        if err != nil {
            fmt.Printf("error getting suggested tx params: %s\n", err)
            return
        }	
        
        // Initialize transaction parameters for the following examples
        fee := txParams.Fee
        firstRound := txParams.LastRound
        lastRound := txParams.LastRound + 1000 
        genHash := b64.StdEncoding.EncodeToString(txParams.GenesisHash)
        genID := txParams.GenesisID 
        // Create an asset
        // Set parameters for asset creation transaction
        creator := pks[1]
        assetName := "latinum"
        unitName := "latinum"
        assetURL := "https://path/to/my/asset/details"
        assetMetadataHash := "thisIsSomeLength32HashCommitment"
        defaultFrozen := false
        decimals := uint32(0)
        totalIssuance := uint64(1000)
        manager := pks[1]
        reserve := pks[1]
        freeze := pks[1]
        clawback := pks[1]
        note := []byte(nil)
        txn, err := transaction.MakeAssetCreateTxn(creator, fee, firstRound, lastRound, note,
        genID, genHash, totalIssuance, decimals, defaultFrozen, manager, reserve, freeze, clawback,
        unitName, assetName, assetURL, assetMetadataHash)
        if err != nil {
            fmt.Printf("Failed to make asset: %s\n", err)
            return
        }
        fmt.Printf("Asset created AssetName: %s\n", txn.AssetConfigTxnFields.AssetParams.AssetName)
        
        txid, stx, err := crypto.SignTransaction(sks[1], txn)
        if err != nil {
            fmt.Printf("Failed to sign transaction: %s\n", err)
            return
        }
        fmt.Printf("Transaction ID: %s\n", txid)
            // Broadcast the transaction to the network
        sendResponse, err := algodClient.SendRawTransaction(stx)
        if err != nil {
            fmt.Printf("failed to send transaction: %s\n", err)
            return
        }
        
        // Wait for transaction to be confirmed
        waitForConfirmation(algodClient, sendResponse.TxID)
            
        // Retrieve asset ID by grabbing the max asset ID
        // from the creator account's holdings. 
        act, err := algodClient.AccountInformation(pks[1], txHeaders...)
        if err != nil {
            fmt.Printf("failed to get account information: %s\n", err)
            return
        }
        assetID := uint64(0)
        for i := range act.AssetParams {
            if i > assetID {
                assetID = i
            }
        }
        fmt.Printf("Asset ID from AssetParams: %d\n", assetID)
        // Retrieve asset info.
        assetInfo, err := algodClient.AssetInformation(assetID, txHeaders...)

        // Print asset info for newly created asset.
        PrettyPrint(assetInfo)
        // Change Asset Manager from Account 1 to Account 2
        manager = pks[2]
        oldmanager := pks[1]

        txn, err = transaction.MakeAssetConfigTxn(oldmanager, fee, 
            firstRound, lastRound, note, genID, genHash, assetID, 
            manager, reserve, freeze, clawback, true)
        if err != nil {
            fmt.Printf("Failed to send txn: %s\n", err)
            return
        }
        txid, stx, err = crypto.SignTransaction(sks[1], txn)
        if err != nil {
            fmt.Printf("Failed to sign transaction: %s\n", err)
            return
        }
        fmt.Printf("Transaction ID: %s\n", txid)
        // Broadcast the transaction to the network
        sendResponse, err = algodClient.SendRawTransaction(stx)
        if err != nil {
            fmt.Printf("failed to send transaction: %s\n", err)
            return
        }
        fmt.Printf("Transaction ID raw: %s\n", sendResponse.TxID)

        // Wait for transaction to be confirmed
        waitForConfirmation(algodClient, sendResponse.TxID)
        // Retrieve asset info.
        assetInfo, err = algodClient.AssetInformation(assetID, txHeaders...)
        // Print asset info showing updated manager address.
        PrettyPrint(assetInfo)

        // Account 3 opts in to receive asset
        txn, err = transaction.MakeAssetAcceptanceTxn(pks[3], fee, firstRound, 
            lastRound, note, genID, genHash, assetID)
        if err != nil {
            fmt.Printf("Failed to send transaction MakeAssetAcceptanceTxn: %s\n", err)
            return
        }
        txid, stx, err = crypto.SignTransaction(sks[3], txn)
        if err != nil {
            fmt.Printf("Failed to sign transaction: %s\n", err)
            return
        }

        fmt.Printf("Transaction ID: %s\n", txid)
        // Broadcast the transaction to the network
        sendResponse, err = algodClient.SendRawTransaction(stx)
        if err != nil {
            fmt.Printf("failed to send transaction: %s\n", err)
            return
        }
        fmt.Printf("Transaction ID raw: %s\n", sendResponse.TxID)

        // Wait for transaction to be confirmed
        waitForConfirmation(algodClient, sendResponse.TxID)

        act, err = algodClient.AccountInformation(pks[3], txHeaders...)
        if err != nil {
            fmt.Printf("failed to get account information: %s\n", err)
            return
        }
        PrettyPrint(act.Assets[assetID])

        // Send  10 of asset from Account 1 to Account 3
        sender := pks[1]
        recipient := pks[3]
        amount := uint64(10)
        closeRemainderTo := ""
        txn, err = transaction.MakeAssetTransferTxn(sender, recipient, 
            closeRemainderTo, amount, fee, firstRound, lastRound, note,
            genID, genHash, assetID)
        if err != nil {
            fmt.Printf("Failed to send transaction MakeAssetTransfer Txn: %s\n", err)
            return
        }
        txid, stx, err = crypto.SignTransaction(sks[1], txn)
        if err != nil {
            fmt.Printf("Failed to sign transaction: %s\n", err)
            return
        }
        fmt.Printf("Transaction ID: %s\n", txid)
        // Broadcast the transaction to the network
        sendResponse, err = algodClient.SendRawTransaction(stx)
        if err != nil {
            fmt.Printf("failed to send transaction: %s\n", err)
            return
        }
        fmt.Printf("Transaction ID raw: %s\n", sendResponse.TxID)

        // Wait for transaction to be confirmed
        waitForConfirmation(algodClient, sendResponse.TxID)

        act, err = algodClient.AccountInformation(pks[3], txHeaders...)
        if err != nil {
            fmt.Printf("failed to get account information: %s\n", err)
            return
        }
        PrettyPrint(act.Assets[assetID])
        // Freeze asset for Account 3.
        newFreezeSetting := true
        target := pks[3]
        txn, err = transaction.MakeAssetFreezeTxn(freeze, fee, firstRound, 
            lastRound, note, genID, genHash, assetID, target, 
            newFreezeSetting)
        if err != nil {
            fmt.Printf("Failed to send txn: %s\n", err)
            return
        }
        txid, stx, err = crypto.SignTransaction(sks[1], txn)
        if err != nil {
            fmt.Printf("Failed to sign transaction: %s\n", err)
            return
        }
        fmt.Printf("Transaction ID: %s\n", txid)
        // Broadcast the transaction to the network
        sendResponse, err = algodClient.SendRawTransaction(stx)
        if err != nil {
            fmt.Printf("failed to send transaction: %s\n", err)
            return
        }
        fmt.Printf("Transaction ID raw: %s\n", sendResponse.TxID)
        // Wait for transaction to be confirmed
        waitForConfirmation(algodClient, sendResponse.TxID)

        act, err = algodClient.AccountInformation(pks[3], txHeaders...)
        if err != nil {
            fmt.Printf("failed to get account information: %s\n", err)
            return
        }
        PrettyPrint(act.Assets[assetID])
        // Revoke an asset
        // The clawback account (Account 1) revokes 10 from Account 3.
        target = pks[3]
        txn, err = transaction.MakeAssetRevocationTxn(clawback, target, creator, amount, fee, firstRound, lastRound, note,
        genID, genHash, assetID)
        if err != nil {
            fmt.Printf("Failed to send txn: %s\n", err)
            return
        }
        txid, stx, err = crypto.SignTransaction(sks[1], txn)
        if err != nil {
            fmt.Printf("Failed to sign transaction: %s\n", err)
            return
        }
        fmt.Printf("Transaction ID: %s\n", txid)
        // Broadcast the transaction to the network
        sendResponse, err = algodClient.SendRawTransaction(stx)
        if err != nil {
            fmt.Printf("failed to send transaction: %s\n", err)
            return
        }
        fmt.Printf("Transaction ID raw: %s\n", sendResponse.TxID)
        // Wait for transaction to be confirmed
        waitForConfirmation(algodClient, sendResponse.TxID)

        act, err = algodClient.AccountInformation(pks[3], txHeaders...)
        if err != nil {
            fmt.Printf("failed to get account information: %s\n", err)
            return
        }
        PrettyPrint(act.Assets[assetID])
        // Destroy the asset
        // all funds are back in the creator's account.
        // Manager account used to destroy the asset.
        txn, err = transaction.MakeAssetDestroyTxn(manager, fee, 
            firstRound, lastRound, note, genID, genHash, assetID)
        if err != nil {
            fmt.Printf("Failed to send txn: %s\n", err)
            return
        }
        txid, stx, err = crypto.SignTransaction(sks[2], txn)
        if err != nil {
            fmt.Printf("Failed to sign transaction: %s\n", err)
            return
        }
        fmt.Printf("Transaction ID: %s\n", txid)
        // Broadcast the transaction to the network
        sendResponse, err = algodClient.SendRawTransaction(stx)
        if err != nil {
            fmt.Printf("failed to send transaction: %s\n", err)
            return
        }
        fmt.Printf("Transaction ID raw: %s\n", sendResponse.TxID)
        // Wait for transaction to be confirmed
        waitForConfirmation(algodClient, sendResponse.TxID)
        // Retrieve asset info. This should now throw an error.
        assetInfo, err = algodClient.AssetInformation(assetID, txHeaders...)
        if err != nil {
            fmt.Printf("%s\n", err)
        }
    }
    ```
