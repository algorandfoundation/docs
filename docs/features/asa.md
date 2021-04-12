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

Specifying a reserve account signifies that non-minted assets will reside in that account instead of the default creator account. Assets transferred from this account are "minted" units of the asset. If you specify a new reserve address, you must make sure the new account has opted into the asset and then issue a transaction to transfer all assets to the new reserve.

!!! warning 
    The reserve account has no functional authority in the protocol. It is purely informational. 


[**Freeze Address**](../reference/transactions.md#freezeaddr)

The freeze account is allowed to freeze or unfreeze the asset holdings for a specific account. When an account is frozen it cannot send or receive the frozen asset. In traditional finance, freezing assets may be performed to restrict liquidation of company stock, to investigate suspected criminal activity or to deny-list certain accounts. If the DefaultFrozen state is set to True, you can use the unfreeze action to authorize certain accounts to trade the asset (such as after passing KYC/AML checks). 

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
# CREATE ASSET
# Get network params for transactions before every transaction.
params = algod_client.suggested_params()
# comment these two lines if you want to use suggested params
params.fee = 1000
params.flat_fee = True

# Account 1 creates an asset called latinum and
# sets Account 2 as the manager, reserve, freeze, and clawback address.
# Asset Creation transaction

txn = AssetConfigTxn(
    sender=accounts[1]['pk'],
    sp=params,
    total=1000,
    default_frozen=False,
    unit_name="LATINUM",
    asset_name="latinum",
    manager=accounts[2]['pk'],
    reserve=accounts[2]['pk'],
    freeze=accounts[2]['pk'],
    clawback=accounts[2]['pk'],
    url="https://path/to/my/asset/details", 
    decimals=0)
# Sign with secret key of creator
stxn = txn.sign(accounts[1]['sk'])

# Send the transaction to the network and retrieve the txid.
txid = algod_client.send_transaction(stxn)
print(txid)

# Retrieve the asset ID of the newly created asset by first
# ensuring that the creation transaction was confirmed,
# then grabbing the asset id from the transaction.

# Wait for the transaction to be confirmed
wait_for_confirmation(algod_client,txid)

try:
    # Pull account info for the creator
    # account_info = algod_client.account_info(accounts[1]['pk'])
    # get asset_id from tx
    # Get the new asset's information from the creator account
    ptx = algod_client.pending_transaction_info(txid)
    asset_id = ptx["asset-index"]
    print_created_asset(algod_client, accounts[1]['pk'], asset_id)
    print_asset_holding(algod_client, accounts[1]['pk'], asset_id)
except Exception as e:
    print(e)

```

``` java tab="Java"  
    // CREATE ASSET
    // get changing network parameters for each transaction
    TransactionParametersResponse params = client.TransactionParams().execute().body();
    params.fee = (long) 1000;

    // Create the Asset:
    BigInteger assetTotal = BigInteger.valueOf(10000);
    boolean defaultFrozen = false;
    String unitName = "myunit";
    String assetName = "my longer asset name";
    String url = "http://this.test.com";
    String assetMetadataHash = "16efaa3924a6fd9d3a4824799a4ac65d";
    Address manager = acct2.getAddress();
    Address reserve = acct2.getAddress();
    Address freeze = acct2.getAddress();
    Address clawback = acct2.getAddress();
    Integer decimals = 0;
    Transaction tx = Transaction.AssetCreateTransactionBuilder().sender(acct1.getAddress()).assetTotal(assetTotal)
            .assetDecimals(decimals).assetUnitName(unitName).assetName(assetName).url(url)
            .metadataHashUTF8(assetMetadataHash).manager(manager).reserve(reserve).freeze(freeze)
            .defaultFrozen(defaultFrozen).clawback(clawback).suggestedParams(params).build();

    // Sign the Transaction with creator account
    SignedTransaction signedTx = acct1.signTransaction(tx);
    Long assetID = null;
    try {
        String id = submitTransaction(signedTx);
        System.out.println("Transaction ID: " + id);
        waitForConfirmation(id);
        // Read the transaction
        PendingTransactionResponse pTrx = client.PendingTransactionInformation(id).execute().body();
        // Now that the transaction is confirmed we can get the assetID
        assetID = pTrx.assetIndex;
        System.out.println("AssetID = " + assetID);
        printCreatedAsset(acct1, assetID);
        printAssetHolding(acct1, assetID);

    } catch (Exception e) {
        e.printStackTrace();
        return;
    }
```

``` go tab="Go"  
	// CREATE ASSET

	// Construct the transaction
	// Set parameters for asset creation 
	creator := pks[1]
	assetName := "latinum"
	unitName := "latinum"
	assetURL := "https://path/to/my/asset/details"
	assetMetadataHash := "thisIsSomeLength32HashCommitment"
	defaultFrozen := false
	decimals := uint32(0)
	totalIssuance := uint64(1000)
	manager := pks[2]
	reserve := pks[2]
	freeze := pks[2]
	clawback := pks[2]
	note := []byte(nil)
	txn, err := transaction.MakeAssetCreateTxn(creator,
		note,
		txParams, totalIssuance, decimals,
		defaultFrozen, manager, reserve, freeze, clawback,
		unitName, assetName, assetURL, assetMetadataHash)

	if err != nil {
		fmt.Printf("Failed to make asset: %s\n", err)
		return
	}
	fmt.Printf("Asset created AssetName: %s\n", txn.AssetConfigTxnFields.AssetParams.AssetName)
	// sign the transaction
	txid, stx, err := crypto.SignTransaction(sks[1], txn)
	if err != nil {
		fmt.Printf("Failed to sign transaction: %s\n", err)
		return
	}
	fmt.Printf("Transaction ID: %s\n", txid)
	// Broadcast the transaction to the network
	sendResponse, err := algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Submitted transaction %s\n", sendResponse)
	// Wait for transaction to be confirmed
	waitForConfirmation(txid, algodClient)
	//    response := algodClient.PendingTransactionInformation(txid)
	//    prettyPrint(response)
	// Retrieve asset ID by grabbing the max asset ID
	// from the creator account's holdings.
	act, err := algodClient.AccountInformation(pks[1]).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to get account information: %s\n", err)
		return
	}

	assetID := uint64(0)
	//	find newest (highest) asset for this account
	for _, asset := range act.CreatedAssets {
		if asset.Index > assetID {
			assetID = asset.Index
		}
	}

	// print created asset and asset holding info for this asset
	fmt.Printf("Asset ID: %d\n", assetID)
	printCreatedAsset(assetID, pks[1], algodClient)
	printAssetHolding(assetID, pks[1], algodClient)
```

``` goal tab="goal"  
goal asset create --creator <address> --total 1000 --unitname <unit-name> --asseturl "https://path/to/my/asset/details" --decimals 0   -d data
```

[See complete code...](https://github.com/algorand/docs/tree/master/examples/assets/v2)

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
# CHANGE MANAGER

# The current manager(Account 2) issues an asset configuration transaction that assigns Account 1 as the new manager.
# Keep reserve, freeze, and clawback address same as before, i.e. account 2
params = algod_client.suggested_params()
# comment these two lines if you want to use suggested params
params.fee = 1000
params.flat_fee = True

# asset_id = 328952;

txn = AssetConfigTxn(
    sender=accounts[2]['pk'],
    sp=params,
    index=asset_id, 
    manager=accounts[1]['pk'],
    reserve=accounts[2]['pk'],
    freeze=accounts[2]['pk'],
    clawback=accounts[2]['pk'])
# sign by the current manager - Account 2
stxn = txn.sign(accounts[2]['sk'])
txid = algod_client.send_transaction(stxn)
print(txid)

# Wait for the transaction to be confirmed
wait_for_confirmation(algod_client, txid)

# Check asset info to view change in management. manager should now be account 1
print_created_asset(algod_client, accounts[1]['pk'], asset_id)
```

``` java tab="Java"  
    // CHANGE MANAGER
    // Change Asset Configuration:
    // assetID = Long.valueOf((your asset id));
    // get changing network parameters for each transaction
    params = client.TransactionParams().execute().body();
    params.fee = (long) 1000;
    // configuration changes must be done by
    // the manager account - changing manager of the asset

    tx = Transaction.AssetConfigureTransactionBuilder().sender(acct2.getAddress()).assetIndex(assetID)
            .manager(acct1.getAddress()).reserve(reserve).freeze(freeze).clawback(clawback).suggestedParams(params)
            .build();

    // the transaction must be signed by the current manager account
    signedTx = acct2.signTransaction(tx);
    // send the transaction to the network
    try {
        String id = submitTransaction(signedTx);
        System.out.println("Transaction ID: " + id);
        waitForConfirmation(signedTx.transactionID);
        // the manager should now be the same as the creator
        System.out.println("AssetID = " + assetID);
        printCreatedAsset(acct1, assetID);

    } catch (Exception e) {
        e.printStackTrace();
        return;
    }
```

``` go tab="Go"  
    // CHANGE MANAGER
	// Change Asset Manager from Account 2 to Account 1
	// assetID := uint64(332920)
	// Get network-related transaction parameters and assign
	txParams, err = algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	txParams.FlatFee = true
	txParams.Fee = 1000

	manager = pks[1]
	oldmanager := pks[2]
	strictEmptyAddressChecking := true
	txn, err = transaction.MakeAssetConfigTxn(oldmanager, note, txParams, assetID, manager, reserve, freeze, clawback, strictEmptyAddressChecking)
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
	sendResponse, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Transaction ID raw: %s\n", txid)

	// Wait for transaction to be confirmed
	waitForConfirmation(txid,algodClient )
	// print created assetinfo for this asset
	fmt.Printf("Asset ID: %d\n", assetID)
	printCreatedAsset(assetID, pks[1], algodClient)
```

``` goal tab="goal"  
goal asset config  --manager <address> --new-reserve <address> --assetid <asset-id> -d data 
```

[See complete code...](https://github.com/algorand/docs/tree/master/examples/assets/v2)

**See also**

- [Anatomy of an Asset Reconfiguration Transaction](./transactions/index.md#reconfigure-an-asset)


## Receiving an Asset

**Authorized by**: The account opting in

Before an account can receive a specific asset it must opt-in to receive it. An opt-in transaction places an asset holding of 0 into the account and increases its minimum balance by 100,000 microAlgos. An opt-in transaction is simply an asset transfer with an amount of 0, both to and from the account opting in. The following code illustrates this transaction.

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
# OPT-IN

# Check if asset_id is in account 3's asset holdings prior
# to opt-in
params = algod_client.suggested_params()
# comment these two lines if you want to use suggested params
params.fee = 1000
params.flat_fee = True

account_info = algod_client.account_info(accounts[3]['pk'])
holding = None
idx = 0
for my_account_info in account_info['assets']:
    scrutinized_asset = account_info['assets'][idx]
    idx = idx + 1    
    if (scrutinized_asset['asset-id'] == asset_id):
        holding = True
        break

if not holding:

    # Use the AssetTransferTxn class to transfer assets and opt-in
    txn = AssetTransferTxn(
        sender=accounts[3]['pk'],
        sp=params,
        receiver=accounts[3]["pk"],
        amt=0,
        index=asset_id)
    stxn = txn.sign(accounts[3]['sk'])
    txid = algod_client.send_transaction(stxn)
    print(txid)
    # Wait for the transaction to be confirmed
    wait_for_confirmation(algod_client, txid)
    # Now check the asset holding for that account.
    # This should now show a holding with a balance of 0.
    print_asset_holding(algod_client, accounts[3]['pk'], asset_id)
```

``` java tab="Java"  
    // OPT-IN
    // Opt in to Receiving the Asset
    // assetID = Long.valueOf((your asset id));
    // get changing network parameters for each transaction
    params = client.TransactionParams().execute().body();
    params.fee = (long) 1000;
    // configuration changes must be done by
    // the manager account - changing manager of the asset
    tx = Transaction.AssetAcceptTransactionBuilder().acceptingAccount(acct3.getAddress()).assetIndex(assetID)
            .suggestedParams(params).build();
    // The transaction must be signed by the current manager account
    signedTx = acct3.signTransaction(tx);
    // send the transaction to the network and
    try {
        String id = submitTransaction(signedTx);
        System.out.println("Transaction ID: " + id);
        waitForConfirmation(signedTx.transactionID);
        // We can now list the account information for acct3
        // and see that it can accept the new asset
        System.out.println("Account 3 = " + acct3.getAddress().toString());
        printAssetHolding(acct3, assetID);
    } catch (Exception e) {
        e.printStackTrace();
        return;
    }
```

``` go tab="Go"  
	// OPT-IN

	// Account 3 opts in to receive latinum
	// Use previously set transaction parameters and update sending address to account 3
	// assetID := uint64(332920)
	// Get network-related transaction parameters and assign
	txParams, err = algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	txParams.FlatFee = true
	txParams.Fee = 1000

	txn, err = transaction.MakeAssetAcceptanceTxn(pks[3], note, txParams, assetID)
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
	sendResponse, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Transaction ID raw: %s\n", txid)

	// Wait for transaction to be confirmed
	waitForConfirmation(txid, algodClient)

	// print created assetholding for this asset and Account 3, showing 0 balance
	fmt.Printf("Asset ID: %d\n", assetID)
	fmt.Printf("Account 3: %s\n", pks[3])
	printAssetHolding(assetID, pks[3], algodClient)
```

``` goal tab="goal"  
goal asset send -a 0 --asset <asset-name>  -f <opt-in-account> -t <opt-in-account> --creator <asset-creator>  -d data
```

[See complete code...](https://github.com/algorand/docs/tree/master/examples/assets/v2)

**See also**

- [Structure of an Asset Opt-In Transaction](./transactions/index.md#opt-in-to-an-asset)

## Transferring an Asset

**Authorized by**: The account that holds the asset to be transferred.

Assets can be transferred between accounts that have opted-in to receiving the asset. These are analogous to standard payment transactions but for Algorand Standard Assets. 

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
# TRANSFER ASSET

# transfer asset of 10 from account 1 to account 3
params = algod_client.suggested_params()
# comment these two lines if you want to use suggested params
params.fee = 1000
params.flat_fee = True
txn = AssetTransferTxn(
    sender=accounts[1]['pk'],
    sp=params,
    receiver=accounts[3]["pk"],
    amt=10,
    index=asset_id)
stxn = txn.sign(accounts[1]['sk'])
txid = algod_client.send_transaction(stxn)
print(txid)
# Wait for the transaction to be confirmed
wait_for_confirmation(algod_client, txid)
# The balance should now be 10.
print_asset_holding(algod_client, accounts[3]['pk'], asset_id)
```

``` java tab="Java"  
    // TRANSFER ASSET
    // Transfer the Asset:
    // assetID = Long.valueOf((your asset id));
    // get changing network parameters for each transaction
    params = client.TransactionParams().execute().body();
    params.fee = (long) 1000;
    // set asset xfer specific parameters
    BigInteger assetAmount = BigInteger.valueOf(10);
    Address sender = acct1.getAddress();
    Address receiver = acct3.getAddress();
    tx = Transaction.AssetTransferTransactionBuilder().sender(sender).assetReceiver(receiver)
            .assetAmount(assetAmount).assetIndex(assetID).suggestedParams(params).build();
    // The transaction must be signed by the sender account
    signedTx = acct1.signTransaction(tx);
    // send the transaction to the network
    try {
        String id = submitTransaction(signedTx);
        System.out.println("Transaction ID: " + id);
        waitForConfirmation(signedTx.transactionID);
        // list the account information for acct1 and acct3
        System.out.println("Account 3  = " + acct3.getAddress().toString());
        printAssetHolding(acct3, assetID);
        System.out.println("Account 1  = " + acct1.getAddress().toString());
        printAssetHolding(acct1, assetID);
    } catch (Exception e) {
        e.printStackTrace();
        return;
    }     
```

``` go tab="Go"  
	// TRANSFER ASSET
	
	// Send  10 latinum from Account 1 to Account 3
	// assetID := uint64(332920)
	// Get network-related transaction parameters and assign
	txParams, err = algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	txParams.FlatFee = true
	txParams.Fee = 1000

	sender := pks[1]
	recipient := pks[3]
	amount := uint64(10)
	closeRemainderTo := ""
	txn, err = transaction.MakeAssetTransferTxn(sender, recipient, amount, note, txParams, closeRemainderTo, 
		assetID)
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
	sendResponse, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Transaction ID raw: %s\n", txid)

	// Wait for transaction to be confirmed
	waitForConfirmation(txid,algodClient)

	// print created assetholding for this asset and Account 3 and Account 1
	// You should see amount of 10 in Account 3, and 990 in Account 1
	fmt.Printf("Asset ID: %d\n", assetID)
	fmt.Printf("Account 3: %s\n", pks[3])
	printAssetHolding(assetID, pks[3], algodClient)
	fmt.Printf("Account 1: %s\n", pks[1])
	printAssetHolding(assetID, pks[1], algodClient)
```

``` goal tab="goal"  
goal asset send -a <asset-amount> --asset <asset-name> -f <asset-sender> -t <asset-receiver> --creator <asset-creator> -d data
```

[See complete code...](https://github.com/algorand/docs/tree/master/examples/assets/v2)

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
# FREEZE ASSET

params = algod_client.suggested_params()
# comment these two lines if you want to use suggested params
params.fee = 1000
params.flat_fee = True
# The freeze address (Account 2) freezes Account 3's latinum holdings.

txn = AssetFreezeTxn(
    sender=accounts[2]['pk'],
    sp=params,
    index=asset_id,
    target=accounts[3]["pk"],
    new_freeze_state=True   
    )
stxn = txn.sign(accounts[2]['sk'])
txid = algod_client.send_transaction(stxn)
print(txid)
# Wait for the transaction to be confirmed
wait_for_confirmation(algod_client, txid)
# The balance should now be 10 with frozen set to true.
print_asset_holding(algod_client, accounts[3]['pk'], asset_id)
```

``` java tab="Java"  
    // FREEZE
    // Freeze the Asset:
    // assetID = Long.valueOf((your asset id));
    // get changing network parameters for each transaction
    params = client.TransactionParams().execute().body();
    params.fee = (long) 1000;
    // The asset was created and configured to allow freezing an account
    // set asset specific parameters
    boolean freezeState = true;
    // The sender should be freeze account
    tx = Transaction.AssetFreezeTransactionBuilder().sender(acct2.getAddress()).freezeTarget(acct3.getAddress())
            .freezeState(freezeState).assetIndex(assetID).suggestedParams(params).build();
    // The transaction must be signed by the freeze account
    signedTx = acct2.signTransaction(tx);
    // send the transaction to the network
    try {
        String id = submitTransaction(signedTx);
        System.out.println("Transaction ID: " + id);
        waitForConfirmation(signedTx.transactionID);
        System.out.println("Account 3 = " + acct3.getAddress().toString());
        printAssetHolding(acct3, assetID);

    } catch (Exception e) {
        e.printStackTrace();
        return;
    }
```

``` go tab="Go"  
	// FREEZE ASSET
	// The freeze address (Account 2) Freeze's asset for Account 3.
	// assetID := uint64(332920)
	// Get network-related transaction parameters and assign
	txParams, err = algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	txParams.FlatFee = true
	txParams.Fee = 1000
	newFreezeSetting := true
	target := pks[3]
	txn, err = transaction.MakeAssetFreezeTxn(freeze, note, txParams, assetID, target, newFreezeSetting)
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
	sendResponse, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Transaction ID raw: %s\n", txid)
	// Wait for transaction to be confirmed
	waitForConfirmation(txid,algodClient)
    // You should now see is-frozen value of true
	fmt.Printf("Asset ID: %d\n", assetID)
	fmt.Printf("Account 3: %s\n", pks[3])
	printAssetHolding(assetID, pks[3], algodClient)
```

``` goal tab="goal"  
goal asset freeze --freezer <asset-freeze-account> --freeze=true --account <account-to-freeze> --creator <asset-creator> --asset <asset-name> -d data
```

[See complete code...](https://github.com/algorand/docs/tree/master/examples/assets/v2)

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
# REVOKE ASSET

# The clawback address (Account 2) revokes 10 latinum from Account 3 and places it back with Account 1.
params = algod_client.suggested_params()
# comment these two lines if you want to use suggested params
params.fee = 1000
params.flat_fee = True

# Must be signed by the account that is the Asset's clawback address
txn = AssetTransferTxn(
    sender=accounts[2]['pk'],
    sp=params,
    receiver=accounts[1]["pk"],
    amt=10,
    index=asset_id,
    revocation_target=accounts[3]['pk']
    )
stxn = txn.sign(accounts[2]['sk'])
txid = algod_client.send_transaction(stxn)
print(txid)
# Wait for the transaction to be confirmed
wait_for_confirmation(algod_client, txid)
# The balance of account 3 should now be 0.
# account_info = algod_client.account_info(accounts[3]['pk'])
print("Account 3")
print_asset_holding(algod_client, accounts[3]['pk'], asset_id)

# The balance of account 1 should increase by 10 to 1000.
print("Account 1")
print_asset_holding(algod_client, accounts[1]['pk'], asset_id)
```

``` java tab="Java"  
    // REVOKE (or clawback)
    // Revoke the asset:
    // The asset was also created with the ability for it to be revoked by
    // clawbackaddress.
    // assetID = Long.valueOf((your asset id));
    // get changing network parameters for each transaction
    params = client.TransactionParams().execute().body();
    params.fee = (long) 1000;

    // set asset specific parameters
    assetAmount = BigInteger.valueOf(10);
    tx = Transaction.AssetClawbackTransactionBuilder().sender(acct2.getAddress())
            .assetClawbackFrom(acct3.getAddress()).assetReceiver(acct1.getAddress()).assetAmount(assetAmount)
            .assetIndex(assetID).suggestedParams(params).build();
    // The transaction must be signed by the clawback account
    signedTx = acct2.signTransaction(tx);
    // send the transaction to the network and
    // wait for the transaction to be confirmed
    try {
        String id = submitTransaction(signedTx);
        System.out.println("Transaction ID: " + id);
        waitForConfirmation(signedTx.transactionID);
        // list the account information for acct1 and acct3
        System.out.println("Account 3  = " + acct3.getAddress().toString());
        printAssetHolding(acct3, assetID);
        System.out.println("Account 1  = " + acct1.getAddress().toString());
        printAssetHolding(acct1, assetID);
    } catch (Exception e) {
        e.printStackTrace();
        return;
    }
```

``` go tab="Go"  
	// REVOKE ASSET
	// Revoke an Asset
	// The clawback address (Account 2) revokes 10 latinum from Account 3 (target)
	// and places it back with Account 1 (creator).
	// assetID := uint64(332920)
	// Get network-related transaction parameters and assign
	txParams, err = algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	txParams.FlatFee = true
	txParams.Fee = 1000
	target = pks[3]
	txn, err = transaction.MakeAssetRevocationTxn(clawback, target, amount, creator, note,
		txParams, assetID)
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
	sendResponse, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Transaction ID raw: %s\n", txid)
	// Wait for transaction to be confirmed
	waitForConfirmation( txid, algodClient)
	// print created assetholding for this asset and Account 3 and Account 1
	// You should see amount of 0 in Account 3, and 1000 in Account 1
	fmt.Printf("Asset ID: %d\n", assetID)
	fmt.Printf("recipient")
	fmt.Printf("Account 3: %s\n", pks[3])
	printAssetHolding(assetID, pks[3], algodClient)
	fmt.Printf("target")
	fmt.Printf("Account 1: %s\n", pks[1])
	printAssetHolding(assetID, pks[1], algodClient)
```

``` goal tab="goal"  
goal asset send -a <amount-to-revoke> --asset <asset-name> -f <address-of-revoke-target> -t <address-to-send-assets-to> --clawback <clawback-address> --creator <creator-address> -d data
```

[See complete code...](https://github.com/algorand/docs/tree/master/examples/assets/v2)

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
# DESTROY ASSET
# With all assets back in the creator's account,
# the manager (Account 1) destroys the asset.
params = algod_client.suggested_params()
# comment these two lines if you want to use suggested params
params.fee = 1000
params.flat_fee = True

# Asset destroy transaction
txn = AssetConfigTxn(
    sender=accounts[1]['pk'],
    sp=params,
    index=asset_id,
    strict_empty_address_check=False
    )

# Sign with secret key of creator
stxn = txn.sign(accounts[1]['sk'])
# Send the transaction to the network and retrieve the txid.
txid = algod_client.send_transaction(stxn)
print(txid)
# Wait for the transaction to be confirmed
wait_for_confirmation(algod_client, txid)

# Asset was deleted.
try:
    print("Account 3 must do a transaction for an amount of 0, " )
    print("with a close_assets_to to the creator account, to clear it from its accountholdings")
    print("For Account 1, nothing should print after this as the asset is destroyed on the creator account")
   
    print_asset_holding(algod_client, accounts[1]['pk'], asset_id)
    print_created_asset(algod_client, accounts[1]['pk'], asset_id)
    # asset_info = algod_client.asset_info(asset_id)
except Exception as e:
    print(e)
```

``` java tab="Java"  
    // DESTROY

    // Destroy the Asset:
    // All assets should now be back in
    // creators account
    // assetID = Long.valueOf((your asset id));
    // get changing network parameters for each transaction
    params = client.TransactionParams().execute().body();
    params.fee = (long) 1000;

    // set destroy asset specific parameters
    // The manager must sign and submit the transaction
    tx = Transaction.AssetDestroyTransactionBuilder().sender(acct1.getAddress()).assetIndex(assetID)
            .suggestedParams(params).build();
    // The transaction must be signed by the manager account
    signedTx = acct1.signTransaction(tx);
    // send the transaction to the network
    try {
        String id = submitTransaction(signedTx);
        System.out.println("Transaction ID: " + id);
        waitForConfirmation(signedTx.transactionID);
        // We list the account information for acct1
        // and check that the asset is no longer exist
        System.out.println("Account 3 must do a transaction for an amount of 0, ");
        System.out.println("with a assetCloseTo to the creator account, to clear it from its accountholdings");
        System.out.println("Account 1  = " + acct1.getAddress().toString());            
        System.out.println("Nothing should print after this, Account 1 asset is sucessfully deleted");
        printAssetHolding(acct1, assetID);
        printCreatedAsset(acct1, assetID);
    } catch (Exception e) {
        e.printStackTrace();
        return;
    }
```

``` go tab="Go"  
	// DESTROY ASSET
	// Destroy the asset
	// Make sure all funds are back in the creator's account. Then use the
	// Manager account (Account 1) to destroy the asset.

	// assetID := uint64(332920)
	// Get network-related transaction parameters and assign
	txParams, err = algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	txParams.FlatFee = true
	txParams.Fee = 1000

	txn, err = transaction.MakeAssetDestroyTxn(manager, note, txParams, assetID)
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
	sendResponse, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Transaction ID raw: %s\n", txid)
	// Wait for transaction to be confirmed
	waitForConfirmation(txid,algodClient)
	fmt.Printf("Asset ID: %d\n", assetID)	
	fmt.Printf("Account 3 must do a transaction for an amount of 0, \n" )
    fmt.Printf("with a closeRemainderTo to the creator account, to clear it from its accountholdings. \n")
    fmt.Printf("For Account 1, nothing should print after this as the asset is destroyed on the creator account \n")

	// print created asset and asset holding info for this asset (should not print anything)

	printCreatedAsset(assetID, pks[1], algodClient)
	printAssetHolding(assetID, pks[1], algodClient)
```

``` goal tab="goal"  
goal asset destroy --creator <creator-address> --manager <asset-manager-address> --asset <asset-name> -d data 
```

[See complete code...](https://github.com/algorand/docs/tree/master/examples/assets/v2)

**See also**

- [Anatomy of the Asset Destroy Transaction](./transactions/index.md#destroy-an-asset)

# Retrieve Asset Information
Retrieve an asset's configuration information from the network using the SDKs or `goal`. Additional details are also added to the accounts that own the specific asset and can be listed with standard account information calls.

!!! info
    The code below illustrates getting asset information without the Indexer. If you have the Indexer installed use the Indexer API to [search for asset](../features/indexer/#search-assets) information.

``` javascript tab="JavaScript"
// Function used to print created asset for account and assetid
const printCreatedAsset = async function (algodclient, account, assetid) {
    // note: if you have an indexer instance available it is easier to just search accounts for an asset
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
    // note: if you have an indexer instance available it is easier to just search accounts for an asset
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
#   note: if you have an indexer instance available it may be easier to just search accounts for an asset
#   Utility function used to print created asset for account and assetid
def print_created_asset(algodclient, account, assetid):    
    # note: if you have an indexer instance available it is easier to just use this
    # response = myindexer.accounts(asset_id = assetid)
    # then use 'account_info['created-assets'][0] to get info on the created asset
    account_info = algodclient.account_info(account)
    idx = 0;
    for my_account_info in account_info['created-assets']:
        scrutinized_asset = account_info['created-assets'][idx]
        idx = idx + 1       
        if (scrutinized_asset['index'] == assetid):
            print("Asset ID: {}".format(scrutinized_asset['index']))
            print(json.dumps(my_account_info['params'], indent=4))
            break

#   Utility function used to print asset holding for account and assetid
def print_asset_holding(algodclient, account, assetid):
    # note: if you have an indexer instance available it is easier to just use this
    # response = myindexer.accounts(asset_id = assetid)
    # then loop thru the accounts returned and match the account you are looking for
    account_info = algodclient.account_info(account)
    idx = 0
    for my_account_info in account_info['assets']:
        scrutinized_asset = account_info['assets'][idx]
        idx = idx + 1        
        if (scrutinized_asset['asset-id'] == assetid):
            print("Asset ID: {}".format(scrutinized_asset['asset-id']))
            print(json.dumps(scrutinized_asset, indent=4))
            break
...
    print_created_asset(algod_client, accounts[1]['pk'], asset_id)
    print_asset_holding(algod_client, accounts[1]['pk'], asset_id)
```

```java tab="Java"
    //note: if you have an indexer instance available it may be easier to just search accounts for an asset
    // utility function to print created asset
    public void printCreatedAsset(Account account, Long assetID) throws Exception {
        if (client == null)
            this.client = connectToNetwork();
        String accountInfo = client.AccountInformation(account.getAddress()).execute().toString();
        JSONObject jsonObj = new JSONObject(accountInfo.toString());
        JSONArray jsonArray = (JSONArray) jsonObj.get("created-assets");
        if (jsonArray.length() > 0) {
            try {
                for (Object o : jsonArray) {
                    JSONObject ca = (JSONObject) o;
                    Integer myassetIDInt = (Integer) ca.get("index");
                    if (assetID.longValue() == myassetIDInt.longValue()) {
                        System.out.println("Created Asset Info: " + ca.toString(2)); // pretty print
                        break;
                    }
                }
            } catch (Exception e) {
                throw (e);
            }
        }
    }

    // utility function to print asset holding
    public void printAssetHolding(Account account, Long assetID) throws Exception {
        if (client == null)
            this.client = connectToNetwork();
        String accountInfo = client.AccountInformation(account.getAddress()).execute().toString();
        JSONObject jsonObj = new JSONObject(accountInfo.toString());
        JSONArray jsonArray = (JSONArray) jsonObj.get("assets");
        if (jsonArray.length() > 0) {
            try {
                for (Object o : jsonArray) {
                    JSONObject ca = (JSONObject) o;
                    Integer myassetIDInt = (Integer) ca.get("asset-id");
                    if (assetID.longValue() == myassetIDInt.longValue()) {
                        System.out.println("Asset Holding Info: " + ca.toString(2)); // pretty print
                        break;
                    }
                }
            } catch (Exception e) {
                throw (e);
            }
        }
    }
    ...
    printCreatedAsset(acct1, assetID);
    printAssetHolding(acct1, assetID);
```

```go tab="Go"
    // note: if you have an indexer instance available it is easier to just search accounts for an asset
    // printAssetHolding utility to print asset holding for account
    func printAssetHolding(assetID uint64, account string, client *algod.Client) {

        act, err := client.AccountInformation(account).Do(context.Background())
        if err != nil {
            fmt.Printf("failed to get account information: %s\n", err)
            return
        }
        for _, assetholding := range act.Assets {
            if assetID == assetholding.AssetId {
                prettyPrint(assetholding)
                break
            }
        }
    }

    // printCreatedAsset utility to print created assert for account
    func printCreatedAsset(assetID uint64, account string, client *algod.Client) {

        act, err := client.AccountInformation(account).Do(context.Background())
        if err != nil {
            fmt.Printf("failed to get account information: %s\n", err)
            return
        }
        for _, asset := range act.CreatedAssets {
            if assetID == asset.Index {
                prettyPrint(asset)
                break
            }
        }
    }
    ...
	printCreatedAsset(assetID, pks[1], algodClient)
	printAssetHolding(assetID, pks[1], algodClient)    
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

!!! info
    Full running code examples for each SDK are available within the GitHub repo for V1 and V2 at [/examples/assets](https://github.com/algorand/docs/tree/master/examples/assets) and for [download](https://github.com/algorand/docs/blob/master/examples/assets/assets.zip?raw=true) (.zip).
