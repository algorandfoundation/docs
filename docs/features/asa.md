title: Assets

The Algorand protocol supports the creation of on-chain assets that benefit from the same security, compatibility, speed and ease of use as the Algo. The official name for assets on Algorand is **Algorand Standard Assets (ASA)**.

With Algorand Standard Assets you can represent stablecoins, loyalty points, system credits, and in-game points, just to name a few examples. You can also represent single, unique assets like a deed for a house, collectable items, unique parts on a supply chain, etc. There is also optional functionality to place transfer restrictions on an asset that help support securities, compliance, and certification use cases.

!!! info
    Assets that represent many of the same type, like a stablecoin, may be referred to as **fungible assets**. Single, unique assets are referred to as **non-fungible assets**. 


This section begins with an [overview](#assets-overview) of the asset implementation on Algorand including a review of all [asset parameters](#asset-parameters). This is followed by [how-tos](#asset-functions) in the SDKs and `goal` for all on-chain asset functions.

!!! info
    Example code snippets are provided throughout this page. Full running code examples for each SDK are available within the GitHub repo for V1 and V2 at [/examples/assets](https://github.com/algorand/docs/tree/master/examples/assets) and for [download](https://github.com/algorand/docs/blob/master/examples/assets/assets.zip?raw=true) (.zip).

# Assets Overview

Here are several things to be aware of before getting started with assets.

- A single Algorand account is permitted to create up to 1000 assets. 
- For every asset an account creates or owns, its minimum balance is increased by 0.1 Algos (100,000 microAlgos). 
- Before a new asset can be transferred to a specific account the receiver must opt-in to receive the asset. This process is described below in [Receiving an Asset](#receiving-an-asset). 
- If any transaction is issued that would violate the maximum number of assets for an account or not meet the minimum balance requirements, the transaction will fail.

## Asset Parameters
The type of asset that is created will depend on the parameters that are passed during asset creation and sometimes during asset re-configuration. View the full list of asset parameters in the [Asset Parameters Reference](../reference/transactions.md#asset-parameters).

### Immutable Asset Parameters

These eight parameters can *only* be specified when an asset is created.  

- [Creator](../reference/transactions.md#creator) (*required*)
- [AssetName](../reference/transactions.md#assetname) (*optional, but recommended*)
- [UnitName](../reference/transactions.md#unitname) (*optional, but recommended*)
- [Total](../reference/transactions.md#total) (*required*)
- [Decimals](../reference/transactions.md#decimals) (*required*)
- [DefaultFrozen](../reference/transactions.md#defaultfrozen) (*required*)
- [URL](../reference/transactions.md#url) (*optional*)
- [MetaDataHash](../reference/transactions.md#metadatahash) (*optional*)

### Mutable Asset Parameters
There are four parameters that correspond to addresses that can authorize specific functionality for an asset. These addresses must be specified on creation but they can also be modified after creation. Alternatively, these addresses can be set as empty strings, which will irrevocably lock the function that they would have had authority over. 

Here are the four address types.

[**Manager Address**](../reference/transactions.md#manageraddr)

The manager account is the only account that can authorize transactions to [re-configure](#modifying-an-asset) or [destroy](#destroying-an-asset) an asset. 

!!! warning
    Never set this address to empty if you want to be able to re-configure or destroy the asset.

[**Reserve Address**](../reference/transactions.md#reserveaddr)

Specifying a reserve account signifies that non-minted assets will reside in that account instead of the default creator account. Assets transferred from this account are "minted" units of the asset. If you specify a new reserve address, you must make sure the new account has opted in to the asset and then issue a transaction to transfer all assets to the new reserve.

!!! warning 
    The reserve account has no functional authority in the protocol. It is purely informational. 


[**Freeze Address**](../reference/transactions.md#freezeaddr)

The freeze account is allowed to freeze or unfreeze the asset holdings for a specific account. When an account is frozen it cannot send or receive the frozen asset. In traditional finance, freezing assets may be performed to restrict liquidation of company stock, to investigate suspected criminal activity or to blacklist certain accounts. If the DefaultFrozen state is set to True, you can use the unfreeze action to authorize certain accounts to trade the asset (such as after passing KYC/AML checks). 

!!! tip
    Set this address to `""` if you want to prove to asset holders that the asset can never be frozen.

[**Clawback Address**](../reference/transactions.md#clawbackaddr)

The clawback address represents an account that is allowed to transfer assets from and to any asset holder (assuming they have opted-in).  Use this if you need the option to revoke assets from an account (like if they breach certain contractual obligations tied to holding the asset). In traditional finance, this sort of transaction is referred to as a clawback.

!!! tip
    Set this address to `""` if you want to ensure to asset holders that assets can never be revoked.

If any of these four addresses is set to `""` that address will be cleared and can never be reset for the life of the asset. This will also effectively disable the feature of that address. For example setting the freeze address to `""` will prevent the asset from ever being frozen.

# Asset Functions

## Creating an Asset
**Transaction Authorizer**: Any account with sufficient Algo balance

Create assets using either the SDKs or `goal`. When using the SDKs supply all creation parameters. With `goal`, managing the various addresses associated with the asset must be done after executing an asset creation. See Modifying an Asset in the next section for more details on changing addresses for the asset.

``` javascript tab="JavaScript" 
    let note = undefined; // arbitrary data to be stored in the transaction; here, none is stored
    // Asset creation specific parameters
    // The following parameters are asset specific
    // Throughout the example these will be re-used. 
    // We will also change the manager later in the example
    let addr = recoveredAccount1.addr;
    // Whether user accounts will need to be unfrozen before transacting    
    let defaultFrozen = false;
    // integer number of decimals for asset unit calculation
    let decimals = 0;
    // total number of this asset available for circulation   
    let totalIssuance = 1000;
    // Used to display asset units to user    
    let unitName = "LATINUM";
    // Friendly name of the asset    
    let assetName = "latinum";
    // Optional string pointing to a URL relating to the asset
    let assetURL = "http://someurl";
    // Optional hash commitment of some sort relating to the asset. 32 character length.
    let assetMetadataHash = "16efaa3924a6fd9d3a4824799a4ac65d";
    // The following parameters are the only ones
    // that can be changed, and they have to be changed
    // by the current manager
    // Specified address can change reserve, freeze, clawback, and manager
    let manager = recoveredAccount2.addr;
    // Specified address is considered the asset reserve
    // (it has no special privileges, this is only informational)
    let reserve = recoveredAccount2.addr;
    // Specified address can freeze or unfreeze user asset holdings 
    let freeze = recoveredAccount2.addr;
    // Specified address can revoke user asset holdings and send 
    // them to other addresses    
    let clawback = recoveredAccount2.addr;

    // signing and sending "txn" allows "addr" to create an asset
    let txn = algosdk.makeAssetCreateTxnWithSuggestedParams(addr, note,
         totalIssuance, decimals, defaultFrozen, manager, reserve, freeze,
        clawback, unitName, assetName, assetURL, assetMetadataHash, params);

    let rawSignedTxn = txn.signTxn(recoveredAccount1.sk)
    let tx = (await algodclient.sendRawTransaction(rawSignedTxn).do());
    console.log("Transaction : " + tx.txId);
    let assetID = null;
    // wait for transaction to be confirmed
    await waitForConfirmation(algodclient, tx.txId);
    // Get the new asset's information from the creator account
    let ptx = await algodclient.pendingTransactionInformation(tx.txId).do();
    assetID = ptx["asset-index"];
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
        "unit_name": <unit-name>,
        "asset_name": <asset-name>,
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
    String unitName = <unit-name>;
    String  assetName = <asset-name>;
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

``` goal tab="goal"  
goal asset create --creator <address> --total 1000 --unitname <unit-name> --asseturl "https://path/to/my/asset/details" --decimals 0   -d data
```

[See complete code...](#complete-code-example)

**See also**

- [Anatomy of an Asset Creation Transaction](./transactions/index.md#create-an-asset)


## Modifying an Asset

**Authorized by**: [Asset Manager Account](../reference/transactions.md#manageraddr)

After an asset has been created only the manager, reserve, freeze and reserve accounts can be changed. All other parameters are locked for the life of the asset. If any of these addresses are set to `""` that address will be cleared and can never be reset for the life of the asset. Only the manager account can make configuration changes and must authorize the transaction.

``` javascript tab="JavaScript"
    // Asset configuration specific parameters
    // all other values are the same so we leave 
    // Them set.
    // specified address can change reserve, freeze, clawback, and manager
    manager = recoveredAccount1.addr;
    // Note that the change has to come from the existing manager
    let ctxn = algosdk.makeAssetConfigTxnWithSuggestedParams(recoveredAccount2.addr, note, 
        assetID, manager, reserve, freeze, clawback, params);
    // This transaction must be signed by the current manager
    rawSignedTxn = ctxn.signTxn(recoveredAccount2.sk)
    let ctx = (await algodclient.sendRawTransaction(rawSignedTxn).do());
    console.log("Transaction : " + ctx.txId);
    // wait for transaction to be confirmed
    await waitForConfirmation(algodclient, ctx.txId);
    // Get the asset information for the newly changed asset
    // use indexer or utiltiy function for Account info
    // The manager should now be the same as the creator
    await printCreatedAsset(algodclient, recoveredAccount1.addr, assetID);
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

``` goal tab="goal"  
goal asset config  --manager <address> --new-reserve <address> --assetid <asset-id> -d data 
```

[See complete code...](#complete-code-example)

**See also**

- [Anatomy of an Asset Reconfiguration Transaction](./transactions/index.md#reconfigure-an-asset)


## Receiving an Asset

**Authorized by**: The account opting in

Before an account can receive a specific asset it must opt-in to receive it. An opt-in transaction places an asset holding of 0 into the account increases its minimum balance by 100,000 microAlgos. An opt-in transaction is simply an asset transfer with an amount of 0, both to and from the account opting in. The following code illustrates this transaction.

``` javascript tab="JavaScript"  

    // Opting in to transact with the new asset
    // Allow accounts that want recieve the new asset
    // Have to opt in. To do this they send an asset transfer
    // of the new asset to themseleves 
    // In this example we are setting up the 3rd recovered account to 
    // receive the new asset
    let sender = recoveredAccount3.addr;
    let recipient = sender;
    // We set revocationTarget to undefined as 
    // This is not a clawback operation
    let revocationTarget = undefined;
    // CloseReaminerTo is set to undefined as
    // we are not closing out an asset
    let closeRemainderTo = undefined;
    // We are sending 0 assets
    amount = 0;
    // signing and sending "txn" allows sender to begin accepting asset specified by creator and index
    let opttxn = algosdk.makeAssetTransferTxnWithSuggestedParams(sender, recipient, closeRemainderTo, revocationTarget,
         amount, note, assetID, params);
    // Must be signed by the account wishing to opt in to the asset    
    rawSignedTxn = opttxn.signTxn(recoveredAccount3.sk);
    let opttx = (await algodclient.sendRawTransaction(rawSignedTxn).do());
    console.log("Transaction : " + opttx.txId);
    // wait for transaction to be confirmed
    await waitForConfirmation(algodclient, opttx.txId);
    //You should now see the new asset listed in the account information
    console.log("Account 3 = " + recoveredAccount3.addr);
    await printAssetHolding(algodclient, recoveredAccount3.addr, assetID);
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

``` goal tab="goal"  
goal asset send -a 0 --asset <asset-name>  -f <opt-in-account> -t <opt-in-account> --creator <asset-creator>  -d data
```

[See complete code...](#complete-code-example)

**See also**

- [Structure of an Asset Opt-In Transaction](./transactions/index.md#opt-in-to-an-asset)

## Transferring an Asset

**Authorized by**: The account that holds the asset to be transferred.

Assets can be transferred between accounts that have opted-in to receiving the asset. These are analagous to standard payment transactions but for Algorand Standard Assets. 

``` javascript tab="JavaScript"  
    // Transfer New Asset:
    // Now that account3 can recieve the new tokens 
    // we can tranfer tokens in from the creator
    // to account3
    sender = recoveredAccount1.addr;
    recipient = recoveredAccount3.addr;
    revocationTarget = undefined;
    closeRemainderTo = undefined;
    //Amount of the asset to transfer
    amount = 10;

    // signing and sending "txn" will send "amount" assets from "sender" to "recipient"
    let xtxn = algosdk.makeAssetTransferTxnWithSuggestedParams(sender, recipient, closeRemainderTo, revocationTarget,
         amount,  note, assetID, params);
    // Must be signed by the account sending the asset  
    rawSignedTxn = xtxn.signTxn(recoveredAccount1.sk)
    let xtx = (await algodclient.sendRawTransaction(rawSignedTxn).do());
    console.log("Transaction : " + xtx.txId);
    // wait for transaction to be confirmed
    await waitForConfirmation(algodclient, xtx.txId);

    // You should now see the 10 assets listed in the account information
    console.log("Account 3 = " + recoveredAccount3.addr);
    await printAssetHolding(algodclient, recoveredAccount3.addr, assetID);
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

``` goal tab="goal"  
goal asset send -a <asset-amount> --asset <asset-name> -f <asset-sender> -t <asset-receiver> --creator <asset-creator> -d data
```

[See complete code...](#complete-code-example)

**See also**

- [Anatomy of an Asset Transfer Transaction](./transactions/index.md#transfer-an-asset)

## Freezing an Asset

**Authorized by**: [Asset Freeze Address](../reference/transactions.md#freezeaddr)

Freezing or unfreezing an asset for an account requires a transaction that is signed by the freeze account. The code below illustrates the freeze transaction.

``` javascript tab="JavaScript"  
    // The asset was created and configured to allow freezing an account
    // If the freeze address is set "", it will no longer be possible to do this.
    // In this example we will now freeze account3 from transacting with the 
    // The newly created asset. 
    // The freeze transaction is sent from the freeze acount
    // Which in this example is account2 
    from = recoveredAccount2.addr;
    freezeTarget = recoveredAccount3.addr;
    freezeState = true;

    // The freeze transaction needs to be signed by the freeze account
    let ftxn = algosdk.makeAssetFreezeTxnWithSuggestedParams(from, note,
        assetID, freezeTarget, freezeState, params)

    // Must be signed by the freeze account   
    rawSignedTxn = ftxn.signTxn(recoveredAccount2.sk)
    let ftx = (await algodclient.sendRawTransaction(rawSignedTxn).do());
    console.log("Transaction : " + ftx.txId);
    // wait for transaction to be confirmed
    await waitForConfirmation(algodclient, ftx.txId);

    // You should now see the asset is frozen listed in the account information
    console.log("Account 3 = " + recoveredAccount3.addr);
    await printAssetHolding(algodclient, recoveredAccount3.addr, assetID);
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

``` goal tab="goal"  
goal asset freeze --freezer <asset-freeze-account> --freeze=true --account <account-to-freeze> --creator <asset-creator> --asset <asset-name> -d data
```

[See complete code...](#complete-code-example)

**See also**

- [Anatomy of an Asset Freeze Transaction](./transactions/index.md#freeze-an-asset)

## Revoking an Asset

**Authorized by**: [Asset Clawback Address](../reference/transactions.md#clawbackaddr)

Revoking an asset for an account removes a specific number of the asset from the revoke target account. Revoking an asset from an account requires specifying an asset sender (the revoke target account) and an asset receiver (the account to transfer the funds back to). The code below illustrates the clawback transaction.

``` javascript tab="JavaScript"  
    // The asset was also created with the ability for it to be revoked by 
    // the clawbackaddress. If the asset was created or configured by the manager
    // to not allow this by setting the clawbackaddress to "" then this would 
    // not be possible.
    // We will now clawback the 10 assets in account3. account2
    // is the clawbackaccount and must sign the transaction
    // The sender will be be the clawback adress.
    // the recipient will also be be the creator in this case
    // that is account3
       sender = recoveredAccount2.addr;
    recipient = recoveredAccount1.addr;
    revocationTarget = recoveredAccount3.addr;
    closeRemainderTo = undefined;
    amount = 10;
    // signing and sending "txn" will send "amount" assets from "revocationTarget" to "recipient",
    // if and only if sender == clawback manager for this asset
    
    let rtxn = algosdk.makeAssetTransferTxnWithSuggestedParams(sender, recipient, closeRemainderTo, revocationTarget,
       amount, note, assetID, params);
    // Must be signed by the account that is the clawback address    
    rawSignedTxn = rtxn.signTxn(recoveredAccount2.sk)
    let rtx = (await algodclient.sendRawTransaction(rawSignedTxn).do());
    console.log("Transaction : " + rtx.txId);
    // wait for transaction to be confirmed
    await waitForConfirmation(algodclient, rtx.txId);

    // You should now see 0 assets listed in the account information
    // for the third account
    console.log("Account 3 = " + recoveredAccount3.addr);
    await printAssetHolding(algodclient, recoveredAccount3.addr, assetID);
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

``` goal tab="goal"  
goal asset send -a <amount-to-revoke> --asset <asset-name> -f <address-of-revoke-target> -t <address-to-send-assets-to> --clawback <clawback-address> --creator <creator-address> -d data
```

[See complete code...](#complete-code-example)

**See also**

- [Anatomy of an Asset Clawback Transaction](./transactions/index.md#revoke-an-asset)

## Destroying an Asset

**Authorized by**: [Asset Manager](../reference/transactions.md#manageraddr)

Created assets can be destroyed only by the asset manager account. All of the assets must be owned by the creator of the asset before the asset can be deleted. 

``` javascript tab="JavaScript"  
    // All of the created assets should now be back in the creators
    // Account so we can delete the asset.
    // If this is not the case the asset deletion will fail
   // The address for the from field must be the manager account
    // Which is currently the creator addr1
    addr = recoveredAccount1.addr;
    note = undefined;
    // if all assets are held by the asset creator,
    // the asset creator can sign and issue "txn" to remove the asset from the ledger. 
    let dtxn = algosdk.makeAssetDestroyTxnWithSuggestedParams(addr, note, assetID, params);
    // The transaction must be signed by the manager which 
    // is currently set to account1
    rawSignedTxn = dtxn.signTxn(recoveredAccount1.sk)
    let dtx = (await algodclient.sendRawTransaction(rawSignedTxn).do());
    console.log("Transaction : " + dtx.txId);
    // wait for transaction to be confirmed
    await waitForConfirmation(algodclient, dtx.txId);

    // The account3 and account1 should no longer contain the asset as it has been destroyed
    console.log("Asset ID: " + assetID);
    console.log("Account 1 = " + recoveredAccount1.addr);
    await printCreatedAsset(algodclient, recoveredAccount1.addr, assetID);
    await printAssetHolding(algodclient, recoveredAccount1.addr, assetID);
    console.log("Account 3 = " + recoveredAccount3.addr);
    await printAssetHolding(algodclient, recoveredAccount3.addr, assetID); 
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

``` goal tab="goal"  
goal asset destroy --creator <creator-address> --manager <asset-manager-address> --asset <asset-name> -d data 
```

[See complete code...](#complete-code-example)

**See also**

- [Anatomy of the Asset Destroy Transaction](./transactions/index.md#destroy-an-asset)

# Retrieve Asset Information
Retrieve an asset's configuration information from the network using the SDKs or `goal`. Additional details are also added to the accounts that own the specific asset and can be listed with standard account information calls.

``` javascript tab="JavaScript"
// Function used to print created asset for account and assetid
const printCreatedAsset = async function (algodclient, account, assetid) {
    // note: if you have an indexer instance available it is easier to just use this
    //     let accountInfo = await indexerClient.searchAccounts()
    //    .assetID(assetIndex).do();
    // and in the loop below use this to extract the asset for a particular account
    // accountInfo['accounts'][idx][account]);
    let accountInfo = await algodclient.accountInformation(account).do();
    for (idx = 0; idx < accountInfo['created-assets'].length; idx++) {
        let scrutinizedAsset = accountInfo['created-assets'][idx];
        if (scrutinizedAsset['index'] == assetid) {
            console.log("AssetID = " + scrutinizedAsset['index']);
            let myparms = JSON.stringify(scrutinizedAsset['params'], undefined, 2);
            console.log("parms = " + myparms);
            break;
        }
    }
};
// Function used to print asset holding for account and assetid
const printAssetHolding = async function (algodclient, account, assetid) {
    // note: if you have an indexer instance available it is easier to just use this
    //     let accountInfo = await indexerClient.searchAccounts()
    //    .assetID(assetIndex).do();
    // and in the loop below use this to extract the asset for a particular account
    // accountInfo['accounts'][idx][account]);
    let accountInfo = await algodclient.accountInformation(account).do();
    for (idx = 0; idx < accountInfo['assets'].length; idx++) {
        let scrutinizedAsset = accountInfo['assets'][idx];
        if (scrutinizedAsset['asset-id'] == assetid) {
            let myassetholding = JSON.stringify(scrutinizedAsset, undefined, 2);
            console.log("assetholdinginfo = " + myassetholding);
            break;
        }
    }
};
...
    await printCreatedAsset(algodclient, recoveredAccount1.addr, assetID);
    await printAssetHolding(algodclient, recoveredAccount1.addr, assetID);

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

``` goal tab="goal"  
goal asset info --creator <creator-address> --asset unitname  -d ~/node/data -w testwall
Asset ID:         <created-asset-id>
Creator:          <creator-address>
Asset name:       testtoken
Unit name:        unitname
Maximum issue:    12 unitname
Reserve amount:   12 unitname
Issued:           0 unitname
Decimals:         0
Default frozen:   false
Manager address:  <creator-address>
Reserve address:  <reserve-address>
Freeze address:   <freeze-address>
Clawback address: <clawback-address>
```

# Complete Code Example

??? example "Complete Example - Asset Options"
    
    ```javascript tab="JavaScript"
    const algosdk = require('algosdk');
    // Retrieve the token, server and port values for your installation in the 
    // algod.net and algod.token files within the data directory

    // UPDATE THESE VALUES
    // const token = "TOKEN";
    // const server = "SERVER";
    // const port = PORT;

    // sandbox
    const token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
    const server = "http://localhost";
    const port = 4001;


    // Function used to wait for a tx confirmation
    const waitForConfirmation = async function (algodclient, txId) {
        let response = await algodclient.status().do();
        let lastround = response["last-round"];
        while (true) {
            const pendingInfo = await algodclient.pendingTransactionInformation(txId).do();
            if (pendingInfo["confirmed-round"] !== null && pendingInfo["confirmed-round"] > 0) {
                //Got the completed Transaction
                console.log("Transaction " + txId + " confirmed in round " + pendingInfo["confirmed-round"]);
                break;
            }
            lastround++;
            await algodclient.statusAfterBlock(lastround).do();
        }
    };


    // Function used to print created asset for account and assetid
    const printCreatedAsset = async function (algodclient, account, assetid) {
        // note: if you have an indexer instance available it is easier to just use this
        //     let accountInfo = await indexerClient.searchAccounts()
        //    .assetID(assetIndex).do();
        // and in the loop below use this to extract the asset for a particular account
        // accountInfo['accounts'][idx][account]);
        let accountInfo = await algodclient.accountInformation(account).do();
        for (idx = 0; idx < accountInfo['created-assets'].length; idx++) {
            let scrutinizedAsset = accountInfo['created-assets'][idx];
            if (scrutinizedAsset['index'] == assetid) {
                console.log("AssetID = " + scrutinizedAsset['index']);
                let myparms = JSON.stringify(scrutinizedAsset['params'], undefined, 2);
                console.log("parms = " + myparms);
                break;
            }
        }
    };
    // Function used to print asset holding for account and assetid
    const printAssetHolding = async function (algodclient, account, assetid) {
        // note: if you have an indexer instance available it is easier to just use this
        //     let accountInfo = await indexerClient.searchAccounts()
        //    .assetID(assetIndex).do();
        // and in the loop below use this to extract the asset for a particular account
        // accountInfo['accounts'][idx][account]);
        let accountInfo = await algodclient.accountInformation(account).do();
        for (idx = 0; idx < accountInfo['assets'].length; idx++) {
            let scrutinizedAsset = accountInfo['assets'][idx];
            if (scrutinizedAsset['asset-id'] == assetid) {
                let myassetholding = JSON.stringify(scrutinizedAsset, undefined, 2);
                console.log("assetholdinginfo = " + myassetholding);
                break;
            }
        }
    };


    // Recover accounts
    // paste in mnemonic phrases here for each account

    var account1_mnemonic = "<25-word-passphrase>";
    var account2_mnemonic = "<25-word-passphrase>";
    var account3_mnemonic = "<25-word-passphrase>";

    var recoveredAccount1 = algosdk.mnemonicToSecretKey(account1_mnemonic);
    var recoveredAccount2 = algosdk.mnemonicToSecretKey(account2_mnemonic);
    var recoveredAccount3 = algosdk.mnemonicToSecretKey(account3_mnemonic);
    console.log(recoveredAccount1.addr);
    console.log(recoveredAccount2.addr);
    console.log(recoveredAccount3.addr);

    // Instantiate the algod wrapper

        let algodclient = new algosdk.Algodv2(token, server, port);


    (async () => {
        // Asset Creation:
        // The first transaciton is to create a new asset
        // These parameters will be required before every 
        // Transaction
        // We will account for changing transaction parameters
        // before every transaction in this example
        let params = await algodclient.getTransactionParams().do();
        //comment out the next two lines to use suggested fee
        params.fee = 1000;
        params.flatFee = true;
        console.log(params);
        let note = undefined; // arbitrary data to be stored in the transaction; here, none is stored

        // Asset creation specific parameters
        // The following parameters are asset specific
        // Throughout the example these will be re-used. 
        // We will also change the manager later in the example
        let addr = recoveredAccount1.addr;
        // Whether user accounts will need to be unfrozen before transacting    
        let defaultFrozen = false;
        // integer number of decimals for asset unit calculation
        let decimals = 0;
        // total number of this asset available for circulation   
        let totalIssuance = 1000;
        // Used to display asset units to user    
        let unitName = "LATINUM";
        // Friendly name of the asset    
        let assetName = "latinum";
        // Optional string pointing to a URL relating to the asset
        let assetURL = "http://someurl";
        // Optional hash commitment of some sort relating to the asset. 32 character length.
        let assetMetadataHash = "16efaa3924a6fd9d3a4824799a4ac65d";
        // The following parameters are the only ones
        // that can be changed, and they have to be changed
        // by the current manager
        // Specified address can change reserve, freeze, clawback, and manager
        let manager = recoveredAccount2.addr;
        // Specified address is considered the asset reserve
        // (it has no special privileges, this is only informational)
        let reserve = recoveredAccount2.addr;
        // Specified address can freeze or unfreeze user asset holdings 
        let freeze = recoveredAccount2.addr;
        // Specified address can revoke user asset holdings and send 
        // them to other addresses    
        let clawback = recoveredAccount2.addr;

        // signing and sending "txn" allows "addr" to create an asset
        let txn = algosdk.makeAssetCreateTxnWithSuggestedParams(addr, note,
            totalIssuance, decimals, defaultFrozen, manager, reserve, freeze,
            clawback, unitName, assetName, assetURL, assetMetadataHash, params);

        let rawSignedTxn = txn.signTxn(recoveredAccount1.sk)
        let tx = (await algodclient.sendRawTransaction(rawSignedTxn).do());
        console.log("Transaction : " + tx.txId);
        let assetID = null;
        // wait for transaction to be confirmed
        await waitForConfirmation(algodclient, tx.txId);
        // Get the new asset's information from the creator account
        let ptx = await algodclient.pendingTransactionInformation(tx.txId).do();
        assetID = ptx["asset-index"];
        // console.log("AssetID = " + assetID);
        
        await printCreatedAsset(algodclient, recoveredAccount1.addr, assetID);
        await printAssetHolding(algodclient, recoveredAccount1.addr, assetID);

        // Change Asset Configuration:
        // Change the manager using an asset configuration transaction

        // First update changing transaction parameters
        // We will account for changing transaction parameters
        // before every transaction in this example
        
        params = await algodclient.getTransactionParams().do();
        //comment out the next two lines to use suggested fee
        params.fee = 1000;
        params.flatFee = true;
        // Asset configuration specific parameters
        // all other values are the same so we leave 
        // Them set.
        // specified address can change reserve, freeze, clawback, and manager
        manager = recoveredAccount1.addr;

        // Note that the change has to come from the existing manager
        let ctxn = algosdk.makeAssetConfigTxnWithSuggestedParams(recoveredAccount2.addr, note, 
            assetID, manager, reserve, freeze, clawback, params);

        // This transaction must be signed by the current manager
        rawSignedTxn = ctxn.signTxn(recoveredAccount2.sk)
        let ctx = (await algodclient.sendRawTransaction(rawSignedTxn).do());
        console.log("Transaction : " + ctx.txId);
        // wait for transaction to be confirmed
        await waitForConfirmation(algodclient, ctx.txId);

        // Get the asset information for the newly changed asset
        // use indexer or utiltiy function for Account info
    
        // The manager should now be the same as the creator
        await printCreatedAsset(algodclient, recoveredAccount1.addr, assetID);
    

        // Opting in to an Asset:
        // Opting in to transact with the new asset
        // Allow accounts that want recieve the new asset
        // Have to opt in. To do this they send an asset transfer
        // of the new asset to themseleves 
        // In this example we are setting up the 3rd recovered account to 
        // receive the new asset
        let sender = recoveredAccount3.addr;
        let recipient = sender;
        // We set revocationTarget to undefined as 
        // This is not a clawback operation
        let revocationTarget = undefined;
        // CloseReaminerTo is set to undefined as
        // we are not closing out an asset
        let closeRemainderTo = undefined;
        // We are sending 0 assets
        amount = 0;

        // First update changing transaction parameters
        // We will account for changing transaction parameters
        // before every transaction in this example
        params = await algodclient.getTransactionParams().do();
        //comment out the next two lines to use suggested fee
        params.fee = 1000;
        params.flatFee = true;
        // signing and sending "txn" allows sender to begin accepting asset specified by creator and index
        let opttxn = algosdk.makeAssetTransferTxnWithSuggestedParams(sender, recipient, closeRemainderTo, revocationTarget,
            amount, note, assetID, params);

        // Must be signed by the account wishing to opt in to the asset    
        rawSignedTxn = opttxn.signTxn(recoveredAccount3.sk);
        let opttx = (await algodclient.sendRawTransaction(rawSignedTxn).do());
        console.log("Transaction : " + opttx.txId);
        // wait for transaction to be confirmed
        await waitForConfirmation(algodclient, opttx.txId);

        //You should now see the new asset listed in the account information
        console.log("Account 3 = " + recoveredAccount3.addr);
        await printAssetHolding(algodclient, recoveredAccount3.addr, assetID);

        // Transfer New Asset:
        // Now that account3 can recieve the new tokens 
        // we can tranfer tokens in from the creator
        // to account3

        sender = recoveredAccount1.addr;
        recipient = recoveredAccount3.addr;
        revocationTarget = undefined;
        closeRemainderTo = undefined;
        //Amount of the asset to transfer
        amount = 10;

        // First update changing transaction parameters
        // We will account for changing transaction parameters
        // before every transaction in this example

        params = await algodclient.getTransactionParams().do();
        //comment out the next two lines to use suggested fee
        params.fee = 1000;
        params.flatFee = true;
        // signing and sending "txn" will send "amount" assets from "sender" to "recipient"
        let xtxn = algosdk.makeAssetTransferTxnWithSuggestedParams(sender, recipient, closeRemainderTo, revocationTarget,
            amount,  note, assetID, params);
        // Must be signed by the account sending the asset  
        rawSignedTxn = xtxn.signTxn(recoveredAccount1.sk)
        let xtx = (await algodclient.sendRawTransaction(rawSignedTxn).do());
        console.log("Transaction : " + xtx.txId);
        // wait for transaction to be confirmed
        await waitForConfirmation(algodclient, xtx.txId);

        // You should now see the 10 assets listed in the account information
        console.log("Account 3 = " + recoveredAccount3.addr);
        await printAssetHolding(algodclient, recoveredAccount3.addr, assetID);


        // freeze asset

        // The asset was created and configured to allow freezing an account
        // If the freeze address is set "", it will no longer be possible to do this.
        // In this example we will now freeze account3 from transacting with the 
        // The newly created asset. 
        // The freeze transaction is sent from the freeze acount
        // Which in this example is account2 
        from = recoveredAccount2.addr;
        freezeTarget = recoveredAccount3.addr;
        freezeState = true;

        // First update changing transaction parameters
        // We will account for changing transaction parameters
        // before every transaction in this example
        // await getChangingParms(algodclient);
        params = await algodclient.getTransactionParams().do();
        //comment out the next two lines to use suggested fee
        params.fee = 1000;
        params.flatFee = true;

        // The freeze transaction needs to be signed by the freeze account
        let ftxn = algosdk.makeAssetFreezeTxnWithSuggestedParams(from, note,
            assetID, freezeTarget, freezeState, params)

        // Must be signed by the freeze account   
        rawSignedTxn = ftxn.signTxn(recoveredAccount2.sk)
        let ftx = (await algodclient.sendRawTransaction(rawSignedTxn).do());
        console.log("Transaction : " + ftx.txId);
        // wait for transaction to be confirmed
        await waitForConfirmation(algodclient, ftx.txId);

        // You should now see the asset is frozen listed in the account information
        console.log("Account 3 = " + recoveredAccount3.addr);
        await printAssetHolding(algodclient, recoveredAccount3.addr, assetID);

        // Revoke an Asset:
        // The asset was also created with the ability for it to be revoked by 
        // the clawbackaddress. If the asset was created or configured by the manager
        // to not allow this by setting the clawbackaddress to "" then this would 
        // not be possible.
        // We will now clawback the 10 assets in account3. account2
        // is the clawbackaccount and must sign the transaction
        // The sender will be be the clawback adress.
        // the recipient will also be be the creator in this case
        // that is account3

        sender = recoveredAccount2.addr;
        recipient = recoveredAccount1.addr;
        revocationTarget = recoveredAccount3.addr;
        closeRemainderTo = undefined;
        amount = 10;

        // First update changing transaction parameters
        // We will account for changing transaction parameters
        // before every transaction in this example
        params = await algodclient.getTransactionParams().do();
        //comment out the next two lines to use suggested fee
        params.fee = 1000;
        params.flatFee = true;
        // signing and sending "txn" will send "amount" assets from "revocationTarget" to "recipient",
        // if and only if sender == clawback manager for this asset
        

        let rtxn = algosdk.makeAssetTransferTxnWithSuggestedParams(sender, recipient, closeRemainderTo, revocationTarget,
        amount, note, assetID, params);
        // Must be signed by the account that is the clawback address    
        rawSignedTxn = rtxn.signTxn(recoveredAccount2.sk)
        let rtx = (await algodclient.sendRawTransaction(rawSignedTxn).do());
        console.log("Transaction : " + rtx.txId);
        // wait for transaction to be confirmed
        await waitForConfirmation(algodclient, rtx.txId);

        // You should now see 0 assets listed in the account information
        // for the third account
        console.log("Account 3 = " + recoveredAccount3.addr);
        await printAssetHolding(algodclient, recoveredAccount3.addr, assetID);
    

        // Destroy and Asset:
        // All of the created assets should now be back in the creators
        // Account so we can delete the asset.
        // If this is not the case the asset deletion will fail

        // First update changing transaction parameters
        // We will account for changing transaction parameters
        // before every transaction in this example

        params = await algodclient.getTransactionParams().do();
        //comment out the next two lines to use suggested fee
        params.fee = 1000;
        params.flatFee = true;

        // The address for the from field must be the manager account
        // Which is currently the creator addr1
        addr = recoveredAccount1.addr;
        note = undefined;
        // if all assets are held by the asset creator,
        // the asset creator can sign and issue "txn" to remove the asset from the ledger. 
        let dtxn = algosdk.makeAssetDestroyTxnWithSuggestedParams(addr, note, assetID, params);
        // The transaction must be signed by the manager which 
        // is currently set to account1
        rawSignedTxn = dtxn.signTxn(recoveredAccount1.sk)
        let dtx = (await algodclient.sendRawTransaction(rawSignedTxn).do());
        console.log("Transaction : " + dtx.txId);
        // wait for transaction to be confirmed
        await waitForConfirmation(algodclient, dtx.txId);

        // The account3 and account1 should no longer contain the asset as it has been destroyed
        console.log("Asset ID: " + assetID);
        console.log("Account 1 = " + recoveredAccount1.addr);
        await printCreatedAsset(algodclient, recoveredAccount1.addr, assetID);
        await printAssetHolding(algodclient, recoveredAccount1.addr, assetID);
        console.log("Account 3 = " + recoveredAccount3.addr);
        await printAssetHolding(algodclient, recoveredAccount3.addr, assetID);  

        
        // Notice that although the asset was destroyed, the asset id and associated 
        // metadata still exists in account holdings for Account 3. 
        // When you destroy an asset, the global parameters associated with that asset
        // (manager addresses, name, etc.) are deleted from the creator's balance record (Account 1).
        // However, holdings are not deleted automatically -- users still need to close out of the deleted asset.
        // This is necessary for technical reasons because we currently can't have a single transaction touch potentially 
        // thousands of accounts (all the holdings that would need to be deleted).


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
    mnemonic1 = <25-word-passphrase>
    mnemonic2 = <25-word-passphrase>
    mnemonic3 = <25-word-passphrase>

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
    algod_address = <algod-address>
    algod_token = <algod-token>
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
        "unit_name": <unit-name>,
        "asset_name": <asset-name>,
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

            final String ALGOD_API_ADDR = <algod-address>;
            final String ALGOD_API_TOKEN = <algod-token>;

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
            
            final String account1_mnemonic = <your-25-word-mnemonic>             
            final String account2_mnemonic = <your-25-word-mnemonic>             
            final String account3_mnemonic = <your-25-word-mnemonic>                    
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
            String unitName = <unit-name>;
            String  assetName = <asset-name>;
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
    const algodAddress = <algod-address>
    const algodToken = <algod-token>


    var txHeaders = append([]*algod.Header{}, &algod.Header{"Content-Type", "application/json"})

    // Accounts to be used through examples
    func loadAccounts() (map[int][]byte, map[int]string){
        var pks = map[int]string {
        	1: "<account1-address>",
        	2: "<account1-address>",
        	3: "<account1-address>",
        }
        mnemonic1 := <your-25-word-mnemonic>
        mnemonic2 := <your-25-word-mnemonic>
        mnemonic3 := <your-25-word-mnemonic>
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
        assetName := <asset-name>
        unitName := <unit-name>
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
