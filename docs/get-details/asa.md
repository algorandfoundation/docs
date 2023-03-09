title: Algorand Standard Assets (ASAs)

The Algorand protocol supports the creation of on-chain assets that benefit from the same security, compatibility, speed and ease of use as the Algo. The official name for assets on Algorand is **Algorand Standard Assets (ASA)**.

With Algorand Standard Assets you can represent stablecoins, loyalty points, system credits, and in-game points, just to name a few examples. You can also represent single, unique assets like a deed for a house, collectable items, unique parts on a supply chain, etc. There is also optional functionality to place transfer restrictions on an asset that help support securities, compliance, and certification use cases.



!!! info
    Assets that represent many of the same type, like a stablecoin, may be referred to as **fungible assets**. Single, unique assets are referred to as **non-fungible assets**. 


This section begins with an [overview](#assets-overview) of the asset implementation on Algorand including a review of all [asset parameters](#asset-parameters). This is followed by [how-tos](#asset-functions) in the SDKs and `goal` for all on-chain asset functions.

!!! info
    Example code snippets are provided throughout this page. Full running code examples for each SDK are available within the GitHub repo at [/examples/assets](https://github.com/algorand/docs/tree/master/examples/assets) and for [download](https://github.com/algorand/docs/blob/master/examples/assets/assets.zip?raw=true) (.zip).

# Assets overview

Here are several things to be aware of before getting started with assets.

- For every asset an account creates or owns, its minimum balance is increased by 0.1 Algos (100,000 microAlgos). 
- Before a new asset can be transferred to a specific account the receiver must opt-in to receive the asset. This process is described below in [Receiving an asset](#receiving-an-asset). 
- If any transaction is issued that would violate the minimum balance requirements, the transaction will fail.

!!! info
    Prior to AVM 1.1 (go-algorand 3.5.0), a single Algorand account was only permitted to create and optin to 1000 assets. These limits are now removed allowing an unlimited number of assets to be created and optin to by a single account.

## Asset parameters
The type of asset that is created will depend on the parameters that are passed during asset creation and sometimes during asset re-configuration. View the full list of asset parameters in the [Asset Parameters Reference](../transactions/transactions#asset-parameters).

### Immutable asset parameters

These eight parameters can *only* be specified when an asset is created.  

- [Creator](../transactions/transactions#creator) (*required*)
- [AssetName](../transactions/transactions#assetname) (*optional, but recommended*)
- [UnitName](../transactions/transactions#unitname) (*optional, but recommended*)
- [Total](../transactions/transactions#total) (*required*)
- [Decimals](../transactions/transactions#decimals) (*required*)
- [DefaultFrozen](../transactions/transactions#defaultfrozen) (*required*)
- [URL](../transactions/transactions#url) (*optional*)
- [MetaDataHash](../transactions/transactions#metadatahash) (*optional*)

### Mutable asset parameters
There are four parameters that correspond to addresses that can authorize specific functionality for an asset. These addresses must be specified on creation but they can also be modified after creation. Alternatively, these addresses can be set as empty strings, which will irrevocably lock the function that they would have had authority over. 

Here are the four address types.

[**Manager Address**](../transactions/transactions#manageraddr)

The manager account is the only account that can authorize transactions to [re-configure](#modifying-an-asset) or [destroy](#destroying-an-asset) an asset. 

!!! warning
    Never set this address to empty if you want to be able to re-configure or destroy the asset.

[**Reserve Address**](../transactions/transactions#reserveaddr)

Specifying a reserve account signifies that non-minted assets will reside in that account instead of the default creator account. Assets transferred from this account are "minted" units of the asset. If you specify a new reserve address, you must make sure the new account has opted into the asset and then issue a transaction to transfer all assets to the new reserve.

!!! warning 
    The reserve account has no functional authority in the protocol. It is purely informational. 


[**Freeze Address**](../transactions/transactions#freezeaddr)

The freeze account is allowed to freeze or unfreeze the asset holdings for a specific account. When an account is frozen it cannot send or receive the frozen asset. In traditional finance, freezing assets may be performed to restrict liquidation of company stock, to investigate suspected criminal activity or to deny-list certain accounts. If the DefaultFrozen state is set to True, you can use the unfreeze action to authorize certain accounts to trade the asset (such as after passing KYC/AML checks).

!!! note
    Just remember that DefaultFrozen is an [immutable parameter](#immutable-asset-parameters) and cannot be changed after creation.

!!! tip
    Set this address to `""` if you want to prove to asset holders that the asset can never be frozen.

[**Clawback Address**](../transactions/transactions#clawbackaddr)

The clawback address represents an account that is allowed to transfer assets from and to any asset holder (assuming they have opted-in).  Use this if you need the option to revoke assets from an account (like if they breach certain contractual obligations tied to holding the asset). In traditional finance, this sort of transaction is referred to as a clawback.

!!! tip
    Set this address to `""` if you want to ensure to asset holders that assets can never be revoked.

If any of these four addresses is set to `""` that address will be cleared and can never be reset for the life of the asset. This will also effectively disable the feature of that address. For example setting the freeze address to `""` will prevent the asset from ever being frozen.

# Asset functions

## Creating an asset
**Transaction Authorizer**: Any account with sufficient Algo balance

Create assets using either the SDKs or `goal`. When using the SDKs supply all creation parameters. With `goal`, managing the various addresses associated with the asset must be done after executing an asset creation. See Modifying an Asset in the next section for more details on changing addresses for the asset.

=== "JavaScript"
    <!-- ===JSSDK_ASSET_CREATE=== -->

    ``` javascript
        let params = await algodclient.getTransactionParams().do();
        // comment out the next two lines to use suggested fee
        // params.fee = 1000;
        // params.flatFee = true;
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
        let txn = algosdk.makeAssetCreateTxnWithSuggestedParams(
            addr, 
            note,
            totalIssuance, 
            decimals, 
            defaultFrozen, 
            manager, 
            reserve, 
            freeze,
            clawback, 
            unitName, 
            assetName, 
            assetURL, 
            assetMetadataHash, 
            params);

        let rawSignedTxn = txn.signTxn(recoveredAccount1.sk)
        let tx = (await algodclient.sendRawTransaction(rawSignedTxn).do());

        let assetID = null;
        // wait for transaction to be confirmed
        const ptx = await algosdk.waitForConfirmation(algodclient, tx.txId, 4);
        // Get the new asset's information from the creator account
        assetID = ptx["asset-index"];
        //Get the completed Transaction
        console.log("Transaction " + tx.txId + " confirmed in round " + ptx["confirmed-round"]);
        
    ```
    <!-- ===JSSDK_ASSET_CREATE=== -->

=== "Python"
    <!-- ===PYSDK_ASSET_CREATE=== -->
```python
# Account 1 creates an asset called `rug` with a total supply
# of 1000 units and sets itself to the freeze/clawback/manager/reserve roles
sp = algod_client.suggested_params()
txn = transaction.AssetConfigTxn(
    sender=acct1.address,
    sp=sp,
    default_frozen=False,
    unit_name="rug",
    asset_name="Really Useful Gift",
    manager=acct1.address,
    reserve=acct1.address,
    freeze=acct1.address,
    clawback=acct1.address,
    url="https://path/to/my/asset/details",
    total=1000,
    decimals=0,
)

# Sign with secret key of creator
stxn = txn.sign(acct1.private_key)
# Send the transaction to the network and retrieve the txid.
txid = algod_client.send_transaction(stxn)
print(f"Sent asset create transaction with txid: {txid}")
# Wait for the transaction to be confirmed
results = transaction.wait_for_confirmation(algod_client, txid, 4)
print(f"Result confirmed in round: {results['confirmed-round']}")

# grab the asset id for the asset we just created
created_asset = results["asset-index"]
print(f"Asset ID created: {created_asset}")
```
[Snippet Source](https://github.com/barnjamin/py-algorand-sdk/blob/doc-examples/_examples/asa.py#L13-L43)
    <!-- ===PYSDK_ASSET_CREATE=== -->

=== "Java"
    <!-- ===JAVASDK_ASSET_CREATE=== -->
```java
        // Account 1 creates an asset called `rug` with a total supply
        // of 1000 units and sets itself to the freeze/clawback/manager/reserve roles
        Response<TransactionParametersResponse> rsp = algodClient.TransactionParams().execute();
        TransactionParametersResponse sp = rsp.body();

        // Under the covers, this is an AssetConfig with asset id set to 0
        Transaction createTxn = Transaction.AssetCreateTransactionBuilder().suggestedParams(sp)
                .sender(acct.getAddress())
                .assetTotal(1000)
                .assetDecimals(0)
                .defaultFrozen(false)
                .assetUnitName("rug")
                .assetName("Really Useful Gift")
                .url("https://path/to/my/asset/details")
                .manager(acct.getAddress())
                .reserve(acct.getAddress())
                .freeze(acct.getAddress())
                .clawback(acct.getAddress())
                .build();

        SignedTransaction signedCreateTxn = acct.signTransaction(createTxn);
        Response<PostTransactionsResponse> submitResult = algodClient.RawTransaction()
                .rawtxn(Encoder.encodeToMsgPack(signedCreateTxn)).execute();
        String txId = submitResult.body().txId;
        PendingTransactionResponse result = Utils.waitForConfirmation(algodClient, txId, 4);

        // Grab the asset id for the asset we just created
        Long asaId = result.assetIndex;
        System.out.printf("Created asset with id: %d\n", asaId);

```
[Snippet Source](https://github.com/barnjamin/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/ASAExamples.java#L48-L78)
    <!-- ===JAVASDK_ASSET_CREATE=== -->

=== "Go"
    <!-- ===GOSDK_ASSET_CREATE=== -->
    ``` go
	// Get network-related transaction parameters and assign
	txParams, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	// txParams.FlatFee = true
	// txParams.Fee = 1000
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
	// Broadcast the transaction to the network
	sendResponse, err := algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Submitted transaction %s\n", sendResponse)
	// Wait for confirmation
	confirmedTxn, err := transaction.WaitForConfirmation(algodClient,txid,  4, context.Background())
	if err != nil {
		fmt.Printf("Error waiting for confirmation on txID: %s\n", txid)
		return
	}
	fmt.Printf("Confirmed Transaction: %s in Round %d\n", txid ,confirmedTxn.ConfirmedRound)
	assetID := confirmedTxn.AssetIndex
	// print created asset and asset holding info for this asset
	fmt.Printf("Asset ID: %d\n", assetID)
	printCreatedAsset(assetID, pks[1], algodClient)
	printAssetHolding(assetID, pks[1], algodClient)
    ```
    <!-- ===GOSDK_ASSET_CREATE=== -->

=== "goal"
    <!-- ===GOAL_ASSET_CREATE=== -->
    ``` goal
        goal asset create --creator <address> --total 1000 --unitname <unit-name> --asseturl "https://path/to/my/asset/details" --decimals 0   -d data
    ```
    <!-- ===GOAL_ASSET_CREATE=== -->


**See also**

- [Anatomy of an Asset Creation Transaction](../transactions#create-an-asset)


## Modifying an asset

**Authorized by**: [Asset Manager Account](../transactions/transactions#manageraddr)

After an asset has been created only the manager, reserve, freeze and clawback accounts can be changed. All other parameters are locked for the life of the asset. If any of these addresses are set to `""` that address will be cleared and can never be reset for the life of the asset. Only the manager account can make configuration changes and must authorize the transaction.

=== "JavaScript"
    <!-- ===JSSDK_ASSET_CONFIG=== -->
	```javascript
    params = await algodclient.getTransactionParams().do();
    // comment out the next two lines to use suggested fee
    // params.fee = 1000;
    // params.flatFee = true;
    // Asset configuration specific parameters
    // all other values are the same so we leave 
    // them set.
    // specified address can change reserve, freeze, clawback, and manager
    manager = recoveredAccount1.addr;

    // Note that the change has to come from the existing manager
    let ctxn = algosdk.makeAssetConfigTxnWithSuggestedParams(
        recoveredAccount2.addr, 
        note, 
        assetID, 
        manager, 
        reserve, 
        freeze, 
        clawback, 
        params);

    // This transaction must be signed by the current manager
    rawSignedTxn = ctxn.signTxn(recoveredAccount2.sk)
    let ctx = (await algodclient.sendRawTransaction(rawSignedTxn).do());
    // Wait for confirmation
    let confirmedTxn = await algosdk.waitForConfirmation(algodclient, ctx.txId, 4);
    //Get the completed Transaction
    console.log("Transaction " + ctx.txId + " confirmed in round " + confirmedTxn["confirmed-round"]);
    
    // Get the asset information for the newly changed asset
    // use indexer or utiltiy function for Account info
    // The manager should now be the same as the creator
    await printCreatedAsset(algodclient, recoveredAccount1.addr, assetID);
 
    ```
    <!-- ===JSSDK_ASSET_CONFIG=== -->

=== "Python"
    <!-- ===PYSDK_ASSET_CONFIG=== -->
```python
sp = algod_client.suggested_params()
# Create a config transaction that wipes the
# reserve address for the asset
txn = transaction.AssetConfigTxn(
    sender=acct1.address,
    sp=sp,
    manager=acct1.address,
    reserve=None,
    freeze=acct1.address,
    clawback=acct1.address,
    strict_empty_address_check=False,
)
# Sign with secret key of manager
stxn = txn.sign(acct1.private_key)
# Send the transaction to the network and retrieve the txid.
txid = algod_client.send_transaction(stxn)
print(f"Sent asset config transaction with txid: {txid}")
# Wait for the transaction to be confirmed
results = transaction.wait_for_confirmation(algod_client, txid, 4)
print(f"Result confirmed in round: {results['confirmed-round']}")
```
[Snippet Source](https://github.com/barnjamin/py-algorand-sdk/blob/doc-examples/_examples/asa.py#L46-L66)
    <!-- ===PYSDK_ASSET_CONFIG=== -->

=== "Java"
    <!-- ===JAVASDK_ASSET_CONFIG=== -->
```java
        Response<TransactionParametersResponse> rsp = algodClient.TransactionParams().execute();
        TransactionParametersResponse sp = rsp.body();
        // Wipe the `reserve` address through an AssetConfigTransaction
        Transaction reconfigureTxn = Transaction.AssetConfigureTransactionBuilder().suggestedParams(sp)
                .sender(acct.getAddress())
                .assetIndex(asaId)
                .manager(acct.getAddress())
                .freeze(acct.getAddress())
                .clawback(acct.getAddress())
                .strictEmptyAddressChecking(false)
                .reserve(new byte[32])
                .build();

```
[Snippet Source](https://github.com/barnjamin/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/ASAExamples.java#L84-L97)
    <!-- ===JAVASDK_ASSET_CONFIG=== -->

=== "Go"

    <!-- ===GOSDK_ASSET_CONFIG=== -->
	``` go  
    // CHANGE MANAGER
	// Change Asset Manager from Account 2 to Account 1
	// Get network-related transaction parameters and assign
	txParams, err = algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	// txParams.FlatFee = true
	// txParams.Fee = 1000

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

	// Broadcast the transaction to the network
	sendResponse, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}

	confirmedTxn, err = transaction.WaitForConfirmation(algodClient,txid,  4, context.Background())
	if err != nil {
		fmt.Printf("Error waiting for confirmation on txID: %s\n", txid)
		return
	}
	fmt.Printf("Confirmed Transaction: %s in Round %d\n", txid ,confirmedTxn.ConfirmedRound)

	// print created assetinfo for this asset
	fmt.Printf("Asset ID: %d\n", assetID)
	printCreatedAsset(assetID, pks[1], algodClient)

    ```
    <!-- ===GOSDK_ASSET_CONFIG=== -->

=== "goal"
    <!-- ===GOAL_ASSET_CONFIG=== -->
	``` goal  
    goal asset config  --manager <address> --new-reserve <address> --assetid <asset-id> -d data 
    ```
    <!-- ===GOAL_ASSET_CONFIG=== -->


**See also**

- [Anatomy of an Asset Reconfiguration Transaction](../transactions#reconfigure-an-asset)


## Receiving an asset

**Authorized by**: The account opting in

Before an account can receive a specific asset it must opt-in to receive it. An opt-in transaction places an asset holding of 0 into the account and increases its minimum balance by 100,000 microAlgos. An opt-in transaction is simply an asset transfer with an amount of 0, both to and from the account opting in. The following code illustrates this transaction.

=== "JavaScript"
    <!-- ===JSSDK_ASSET_OPTIN=== -->
	``` javascript  
    // Opting in to transact with the new asset
    // Allow accounts that want recieve the new asset
    // Have to opt in. To do this they send an asset transfer
    // of the new asset to themseleves 
    // In this example we are setting up the 3rd recovered account to 
    // receive the new asset

    // First update changing transaction parameters
    // We will account for changing transaction parameters
    // before every transaction in this example
    params = await algodclient.getTransactionParams().do();
    //comment out the next two lines to use suggested fee
    // params.fee = 1000;
    // params.flatFee = true;

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
    let opttxn = algosdk.makeAssetTransferTxnWithSuggestedParams(
        sender, 
        recipient, 
        closeRemainderTo, 
        revocationTarget,
        amount, 
        note, 
        assetID, 
        params);

    // Must be signed by the account wishing to opt in to the asset    
    rawSignedTxn = opttxn.signTxn(recoveredAccount3.sk);
    let opttx = (await algodclient.sendRawTransaction(rawSignedTxn).do());
    // Wait for confirmation
    confirmedTxn = await algosdk.waitForConfirmation(algodclient, opttx.txId, 4);
    //Get the completed Transaction
    console.log("Transaction " + opttx.txId + " confirmed in round " + confirmedTxn["confirmed-round"]);

    //You should now see the new asset listed in the account information
    console.log("Account 3 = " + recoveredAccount3.addr);
    await printAssetHolding(algodclient, recoveredAccount3.addr, assetID);

    ```
    <!-- ===JSSDK_ASSET_OPTIN=== -->

=== "Python"
    <!-- ===PYSDK_ASSET_OPTIN=== -->
```python
sp = algod_client.suggested_params()
# Create opt-in transaction
# asset transfer from me to me for asset id we want to opt-in to with amt==0
optin_txn = transaction.AssetOptInTxn(
    sender=acct2.address, sp=sp, index=created_asset
)
signed_optin_txn = optin_txn.sign(acct2.private_key)
txid = algod_client.send_transaction(signed_optin_txn)
print(f"Sent opt in transaction with txid: {txid}")

# Wait for the transaction to be confirmed
results = transaction.wait_for_confirmation(algod_client, txid, 4)
print(f"Result confirmed in round: {results['confirmed-round']}")

acct_info = algod_client.account_info(acct2.address)
matching_asset = [
    asset
    for asset in acct_info["assets"]
    if asset["asset-id"] == created_asset
].pop()
assert matching_asset["amount"] == 0
assert matching_asset["is-frozen"] is False
```
[Snippet Source](https://github.com/barnjamin/py-algorand-sdk/blob/doc-examples/_examples/asa.py#L79-L101)
    <!-- ===PYSDK_ASSET_OPTIN=== -->

=== "Java"
    <!-- ===JAVASDK_ASSET_OPTIN=== -->
```java
        Response<TransactionParametersResponse> rsp = algodClient.TransactionParams().execute();
        TransactionParametersResponse sp = rsp.body();
        // Under the covers, this is an AssetTransfer from me to me for amount 0
        // with asset id set to the asset we wish to start accepting
        Transaction optInTxn = Transaction.AssetAcceptTransactionBuilder().suggestedParams(sp)
                .sender(acct.getAddress())
                .assetIndex(asaId)
                .build();

```
[Snippet Source](https://github.com/barnjamin/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/ASAExamples.java#L104-L113)
    <!-- ===JAVASDK_ASSET_OPTIN=== -->

=== "Go"
    <!-- ===GOSDK_ASSET_OPTIN=== -->
	``` go  
	// OPT-IN
	// Account 3 opts in to receive latinum
	// Use previously set transaction parameters and update sending address to account 3
	// Get network-related transaction parameters and assign
	txParams, err = algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	// txParams.FlatFee = true
	// txParams.Fee = 1000
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
	// Broadcast the transaction to the network
	sendResponse, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	confirmedTxn, err = transaction.WaitForConfirmation(algodClient,txid,  4, context.Background())
	if err != nil {
		fmt.Printf("Error waiting for confirmation on txID: %s\n", txid)
		return
	}
	fmt.Printf("Confirmed Transaction: %s in Round %d\n", txid ,confirmedTxn.ConfirmedRound)
	// print created assetholding for this asset and Account 3, showing 0 balance
	fmt.Printf("Asset ID: %d\n", assetID)
	fmt.Printf("Account 3: %s\n", pks[3])
	printAssetHolding(assetID, pks[3], algodClient)
    ```
    <!-- ===GOSDK_ASSET_OPTIN=== -->

=== "goal"
    <!-- ===GOAL_ASSET_OPTIN=== -->
	``` goal  
    goal asset send -a 0 --asset <asset-name>  -f <opt-in-account> -t <opt-in-account> --creator <asset-creator>  -d data
    ```
    <!-- ===GOAL_ASSET_OPTIN=== -->

**See also**

- [Structure of an Asset Opt-In Transaction](../transactions#opt-in-to-an-asset)

## Transferring an asset

**Authorized by**: The account that holds the asset to be transferred.

Assets can be transferred between accounts that have opted-in to receiving the asset. These are analogous to standard payment transactions but for Algorand Standard Assets. 

=== "JavaScript"
    <!-- ===JSSDK_ASSET_XFER=== -->
	``` javascript  
    // Transfer New Asset:
    // Now that account3 can recieve the new tokens 
    // we can tranfer tokens in from the creator
    // to account3
    // First update changing transaction parameters
    // We will account for changing transaction parameters
    // before every transaction in this example

    params = await algodclient.getTransactionParams().do();
    //comment out the next two lines to use suggested fee
    // params.fee = 1000;
    // params.flatFee = true;

    sender = recoveredAccount1.addr;
    recipient = recoveredAccount3.addr;
    revocationTarget = undefined;
    closeRemainderTo = undefined;
    //Amount of the asset to transfer
    amount = 10;

    // signing and sending "txn" will send "amount" assets from "sender" to "recipient"
    let xtxn = algosdk.makeAssetTransferTxnWithSuggestedParams(
        sender, 
        recipient, 
        closeRemainderTo, 
        revocationTarget,
        amount,  
        note, 
        assetID, 
        params);
    // Must be signed by the account sending the asset  
    rawSignedTxn = xtxn.signTxn(recoveredAccount1.sk)
    let xtx = (await algodclient.sendRawTransaction(rawSignedTxn).do());

    // Wait for confirmation
    confirmedTxn = await algosdk.waitForConfirmation(algodclient, xtx.txId, 4);
    //Get the completed Transaction
    console.log("Transaction " + xtx.txId + " confirmed in round " + confirmedTxn["confirmed-round"]);

    // You should now see the 10 assets listed in the account information
    console.log("Account 3 = " + recoveredAccount3.addr);
    await printAssetHolding(algodclient, recoveredAccount3.addr, assetID);

    ```
    <!-- ===JSSDK_ASSET_XFER=== -->

=== "Python"
    <!-- ===PYSDK_ASSET_XFER=== -->
```python
sp = algod_client.suggested_params()
# Create transfer transaction
xfer_txn = transaction.AssetTransferTxn(
    sender=acct1.address,
    sp=sp,
    receiver=acct2.address,
    amt=1,
    index=created_asset,
)
signed_xfer_txn = xfer_txn.sign(acct1.private_key)
txid = algod_client.send_transaction(signed_xfer_txn)
print(f"Sent transfer transaction with txid: {txid}")

results = transaction.wait_for_confirmation(algod_client, txid, 4)
print(f"Result confirmed in round: {results['confirmed-round']}")

acct_info = algod_client.account_info(acct2.address)
matching_asset = [
    asset
    for asset in acct_info["assets"]
    if asset["asset-id"] == created_asset
].pop()
assert matching_asset["amount"] == 1
```
[Snippet Source](https://github.com/barnjamin/py-algorand-sdk/blob/doc-examples/_examples/asa.py#L105-L128)
    <!-- ===PYSDK_ASSET_XFER=== -->

=== "Java"
    <!-- ===JAVASDK_ASSET_XFER=== -->
```java
        Response<TransactionParametersResponse> rsp = algodClient.TransactionParams().execute();
        TransactionParametersResponse sp = rsp.body();
        // Under the covers, this is an AssetTransfer from me to me for amount 0
        // with asset id set to the asset we wish to start accepting
        Transaction xferTxn = Transaction.AssetTransferTransactionBuilder().suggestedParams(sp)
                .sender(sender.getAddress())
                .assetReceiver(receiver.getAddress())
                .assetIndex(asaId)
                .assetAmount(1)
                .build();

```
[Snippet Source](https://github.com/barnjamin/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/ASAExamples.java#L121-L132)
    <!-- ===JAVASDK_ASSET_XFER=== -->

=== "Go"
    <!-- ===GOSDK_ASSET_XFER=== -->
	``` go  
	// TRANSFER ASSET
	
	// Send 10 latinum from Account 1 to Account 3
	// Get network-related transaction parameters and assign
	txParams, err = algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	// txParams.FlatFee = true
	// txParams.Fee = 1000
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
	// Broadcast the transaction to the network
	sendResponse, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	// Wait for transaction to be confirmed
	confirmedTxn, err = transaction.WaitForConfirmation(algodClient,txid,  4, context.Background())
	if err != nil {
		fmt.Printf("Error waiting for confirmation on txID: %s\n", txid)
		return
	}
	fmt.Printf("Confirmed Transaction: %s in Round %d\n", txid ,confirmedTxn.ConfirmedRound)
	// print created assetholding for this asset and Account 3 and Account 1
	// You should see amount of 10 in Account 3, and 990 in Account 1
	fmt.Printf("Asset ID: %d\n", assetID)
	fmt.Printf("Account 3: %s\n", pks[3])
	printAssetHolding(assetID, pks[3], algodClient)
	fmt.Printf("Account 1: %s\n", pks[1])
	printAssetHolding(assetID, pks[1], algodClient)
    ```
    <!-- ===GOSDK_ASSET_XFER=== -->

=== "goal"
    <!-- ===GOAL_ASSET_XFER=== -->
	``` goal  
    goal asset send -a <asset-amount> --asset <asset-name> -f <asset-sender> -t <asset-receiver> --creator <asset-creator> -d data
    ```
    <!-- ===GOAL_ASSET_XFER=== -->

[See complete code...](https://github.com/algorand/docs/tree/master/examples/assets/v2)

**See also**

- [Anatomy of an Asset Transfer Transaction](../transactions#transfer-an-asset)

## Freezing an asset

**Authorized by**: [Asset Freeze Address](../transactions/transactions#freezeaddr)

Freezing or unfreezing an asset for an account requires a transaction that is signed by the freeze account. The code below illustrates the freeze transaction.

=== "JavaScript"
    <!-- ===JSSDK_ASSET_FREEZE=== -->
	``` javascript  
    // freeze asset
    // The asset was created and configured to allow freezing an account
    // If the freeze address is set "", it will no longer be possible to do this.
    // In this example we will now freeze account3 from transacting with the 
    // The newly created asset. 
    // The freeze transaction is sent from the freeze acount
    // Which in this example is account2 

    // First update changing transaction parameters
    // We will account for changing transaction parameters
    // before every transaction in this example
    // await getChangingParms(algodclient);
    params = await algodclient.getTransactionParams().do();
    //comment out the next two lines to use suggested fee
    // params.fee = 1000;
    // params.flatFee = true;

    from = recoveredAccount2.addr;
    freezeTarget = recoveredAccount3.addr;
    freezeState = true;

    // The freeze transaction needs to be signed by the freeze account
    let ftxn = algosdk.makeAssetFreezeTxnWithSuggestedParams(
        from, 
        note,
        assetID, 
        freezeTarget, 
        freezeState, 
        params)

    // Must be signed by the freeze account   
    rawSignedTxn = ftxn.signTxn(recoveredAccount2.sk)
    let ftx = (await algodclient.sendRawTransaction(rawSignedTxn).do());

    // Wait for confirmation
    confirmedTxn = await algosdk.waitForConfirmation(algodclient, ftx.txId, 4);
    //Get the completed Transaction
    console.log("Transaction " + ftx.txId + " confirmed in round " + confirmedTxn["confirmed-round"]);

    // You should now see the asset is frozen listed in the account information
    console.log("Account 3 = " + recoveredAccount3.addr);
    await printAssetHolding(algodclient, recoveredAccount3.addr, assetID);


    ```
    <!-- ===JSSDK_ASSET_FREEZE=== -->

=== "Python"
    <!-- ===PYSDK_ASSET_FREEZE=== -->
```python
sp = algod_client.suggested_params()
# Create freeze transaction to freeze the asset in acct2 balance
freeze_txn = transaction.AssetFreezeTxn(
    sender=acct1.address,
    sp=sp,
    index=created_asset,
    target=acct2.address,
    new_freeze_state=True,
)
signed_freeze_txn = freeze_txn.sign(acct1.private_key)
txid = algod_client.send_transaction(signed_freeze_txn)
print(f"Sent freeze transaction with txid: {txid}")

results = transaction.wait_for_confirmation(algod_client, txid, 4)
print(f"Result confirmed in round: {results['confirmed-round']}")

acct_info = algod_client.account_info(acct2.address)
matching_asset = [
    asset
    for asset in acct_info["assets"]
    if asset["asset-id"] == created_asset
].pop()
assert matching_asset["is-frozen"] is True
```
[Snippet Source](https://github.com/barnjamin/py-algorand-sdk/blob/doc-examples/_examples/asa.py#L131-L154)
    <!-- ===PYSDK_ASSET_FREEZE=== -->

=== "Java"
    <!-- ===JAVASDK_ASSET_FREEZE=== -->
```java
        Response<TransactionParametersResponse> rsp = algodClient.TransactionParams().execute();
        TransactionParametersResponse sp = rsp.body();
        // Set the freeze state on the account, only the account that is set to the
        // freeze role
        // on the asset may issue this transaction
        Transaction freezeTxn = Transaction.AssetFreezeTransactionBuilder().suggestedParams(sp)
                .sender(sender.getAddress())
                .freezeTarget(receiver.getAddress())
                .freezeState(true)
                .assetIndex(asaId)
                .build();

```
[Snippet Source](https://github.com/barnjamin/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/ASAExamples.java#L140-L152)
    <!-- ===JAVASDK_ASSET_FREEZE=== -->

=== "Go"
    <!-- ===GOSDK_ASSET_FREEZE=== -->
	``` go  
	// FREEZE ASSET
	// The freeze address (Account 2) Freeze's asset for Account 3.
	// Get network-related transaction parameters and assign
	txParams, err = algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	// txParams.FlatFee = true
	// txParams.Fee = 1000
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
	// Broadcast the transaction to the network
	sendResponse, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	// Wait for transaction to be confirmed
	confirmedTxn, err = transaction.WaitForConfirmation(algodClient,txid,  4, context.Background())
	if err != nil {
		fmt.Printf("Error waiting for confirmation on txID: %s\n", txid)
		return
	}
	fmt.Printf("Confirmed Transaction: %s in Round %d\n", txid ,confirmedTxn.ConfirmedRound)

    // You should now see is-frozen value of true
	fmt.Printf("Asset ID: %d\n", assetID)
	fmt.Printf("Account 3: %s\n", pks[3])
	printAssetHolding(assetID, pks[3], algodClient)
    ```
    <!-- ===GOSDK_ASSET_FREEZE=== -->

=== "goal"
    <!-- ===GOAL_ASSET_FREEZE=== -->
	``` goal  
    goal asset freeze --freezer <asset-freeze-account> --freeze=true --account <account-to-freeze> --creator <asset-creator> --asset <asset-name> -d data
    ```
    <!-- ===GOAL_ASSET_FREEZE=== -->

**See also**

- [Anatomy of an Asset Freeze Transaction](../transactions#freeze-an-asset)

## Revoking an asset

**Authorized by**: [Asset Clawback Address](../transactions/transactions#clawbackaddr)

Revoking an asset for an account removes a specific number of the asset from the revoke target account. Revoking an asset from an account requires specifying an asset sender (the revoke target account) and an asset receiver (the account to transfer the funds back to). The code below illustrates the clawback transaction.

=== "JavaScript"
    <!-- ===JSSDK_ASSET_CLAWBACK=== -->
	``` javascript  
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
    // First update changing transaction parameters
    // We will account for changing transaction parameters
    // before every transaction in this example
    params = await algodclient.getTransactionParams().do();
    //comment out the next two lines to use suggested fee
    // params.fee = 1000;
    // params.flatFee = true;   

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
    // Wait for confirmation
    confirmedTxn = await algosdk.waitForConfirmation(algodclient, rtx.txId, 4);
    //Get the completed Transaction
    console.log("Transaction " + rtx.txId + " confirmed in round " + confirmedTxn["confirmed-round"]);

    // You should now see 0 assets listed in the account information
    // for the third account
    console.log("Account 3 = " + recoveredAccount3.addr);
    await printAssetHolding(algodclient, recoveredAccount3.addr, assetID);

    ```
    <!-- ===JSSDK_ASSET_CLAWBACK=== -->

=== "Python"
    <!-- ===PYSDK_ASSET_CLAWBACK=== -->
```python
sp = algod_client.suggested_params()
# Create clawback transaction to freeze the asset in acct2 balance
clawback_txn = transaction.AssetTransferTxn(
    sender=acct1.address,
    sp=sp,
    receiver=acct1.address,
    amt=1,
    index=created_asset,
    revocation_target=acct2.address,
)
signed_clawback_txn = clawback_txn.sign(acct1.private_key)
txid = algod_client.send_transaction(signed_clawback_txn)
print(f"Sent clawback transaction with txid: {txid}")

results = transaction.wait_for_confirmation(algod_client, txid, 4)
print(f"Result confirmed in round: {results['confirmed-round']}")

acct_info = algod_client.account_info(acct2.address)
matching_asset = [
    asset
    for asset in acct_info["assets"]
    if asset["asset-id"] == created_asset
].pop()
assert matching_asset["amount"] == 0
assert matching_asset["is-frozen"] is True
```
[Snippet Source](https://github.com/barnjamin/py-algorand-sdk/blob/doc-examples/_examples/asa.py#L157-L182)
    <!-- ===PYSDK_ASSET_CLAWBACK=== -->

=== "Java"
    <!-- ===JAVASDK_ASSET_CLAWBACK=== -->
```java
        Response<TransactionParametersResponse> rsp = algodClient.TransactionParams().execute();
        TransactionParametersResponse sp = rsp.body();
        // revoke an asset from an account, only the account that is set to the clawback
        // role
        // on the asset may issue this transaction
        Transaction clawbackTxn = Transaction.AssetClawbackTransactionBuilder().suggestedParams(sp)
                .sender(sender.getAddress())
                .assetClawbackFrom(receiver.getAddress())
                .assetReceiver(sender.getAddress())
                .assetIndex(asaId)
                .assetAmount(1)
                .build();

```
[Snippet Source](https://github.com/barnjamin/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/ASAExamples.java#L160-L173)
    <!-- ===JAVASDK_ASSET_CLAWBACK=== -->

=== "Go"
    <!-- ===GOSDK_ASSET_CLAWBACK=== -->
	``` go  
	// REVOKE ASSET
	// Revoke an Asset
	// The clawback address (Account 2) revokes 10 latinum from Account 3 (target)
	// and places it back with Account 1 (creator).
	// Get network-related transaction parameters and assign
	txParams, err = algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	// txParams.FlatFee = true
	// txParams.Fee = 1000
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
	// Broadcast the transaction to the network
	sendResponse, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	// Wait for transaction to be confirmed
	confirmedTxn, err = transaction.WaitForConfirmation(algodClient,txid,  4, context.Background())
	if err != nil {
		fmt.Printf("Error waiting for confirmation on txID: %s\n", txid)
		return
	}
	fmt.Printf("Confirmed Transaction: %s in Round %d\n", txid ,confirmedTxn.ConfirmedRound)
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
    <!-- ===GOSDK_ASSET_CLAWBACK=== -->

=== "goal"
    <!-- ===GOAL_ASSET_CLAWBACK=== -->
	``` goal  
    goal asset send -a <amount-to-revoke> --asset <asset-name> -f <address-of-revoke-target> -t <address-to-send-assets-to> --clawback <clawback-address> --creator <creator-address> -d data
    ```
    <!-- ===GOAL_ASSET_CLAWBACK=== -->

**See also**

- [Anatomy of an Asset Clawback Transaction](../transactions#revoke-an-asset)

## Destroying an asset

**Authorized by**: [Asset Manager](../transactions/transactions#manageraddr)

Created assets can be destroyed only by the asset manager account. All of the assets must be owned by the creator of the asset before the asset can be deleted. 

=== "JavaScript"
    <!-- ===JSSDK_ASSET_DELETE=== -->
	``` javascript  
    // Destroy an Asset:
    // All of the created assets should now be back in the creators
    // Account so we can delete the asset.
    // If this is not the case the asset deletion will fail

    // First update changing transaction parameters
    // We will account for changing transaction parameters
    // before every transaction in this example

    params = await algodclient.getTransactionParams().do();
    //comment out the next two lines to use suggested fee
    // params.fee = 1000;
    // params.flatFee = true;

    // The address for the from field must be the manager account
    // Which is currently the creator addr1
    addr = recoveredAccount1.addr;
    note = undefined;
    // if all assets are held by the asset creator,
    // the asset creator can sign and issue "txn" to remove the asset from the ledger. 
    let dtxn = algosdk.makeAssetDestroyTxnWithSuggestedParams(
        addr, 
        note, 
        assetID, 
        params);
    // The transaction must be signed by the manager which 
    // is currently set to account1
    rawSignedTxn = dtxn.signTxn(recoveredAccount1.sk)
    let dtx = (await algodclient.sendRawTransaction(rawSignedTxn).do());

    // Wait for confirmation
    confirmedTxn = await algosdk.waitForConfirmation(algodclient, dtx.txId, 4);
    //Get the completed Transaction
    console.log("Transaction " + dtx.txId + " confirmed in round " + confirmedTxn["confirmed-round"]);

    // The account3 and account1 should no longer contain the asset as it has been destroyed
    console.log("Asset ID: " + assetID);
    console.log("Account 1 = " + recoveredAccount1.addr);
    await printCreatedAsset(algodclient, recoveredAccount1.addr, assetID);
    await printAssetHolding(algodclient, recoveredAccount1.addr, assetID);
    console.log("Account 3 = " + recoveredAccount3.addr);
    await printAssetHolding(algodclient, recoveredAccount3.addr, assetID);  

    ```
    <!-- ===JSSDK_ASSET_DELETE=== -->

=== "Python"
    <!-- ===PYSDK_ASSET_DELETE=== -->
```python
sp = algod_client.suggested_params()
# Create asset destroy transaction to destroy the asset
destroy_txn = transaction.AssetDestroyTxn(
    sender=acct1.address,
    sp=sp,
    index=created_asset,
)
signed_destroy_txn = destroy_txn.sign(acct1.private_key)
txid = algod_client.send_transaction(signed_destroy_txn)
print(f"Sent destroy transaction with txid: {txid}")

results = transaction.wait_for_confirmation(algod_client, txid, 4)
print(f"Result confirmed in round: {results['confirmed-round']}")

# now, trying to fetch the asset info should result in an error
try:
    info = algod_client.asset_info(created_asset)
except Exception as e:
    print("Expected Error:", e)
```
[Snippet Source](https://github.com/barnjamin/py-algorand-sdk/blob/doc-examples/_examples/asa.py#L185-L204)
    <!-- ===PYSDK_ASSET_DELETE=== -->

=== "Java"
    <!-- ===JAVASDK_ASSET_DELETE=== -->
```java
        Response<TransactionParametersResponse> rsp = algodClient.TransactionParams().execute();
        TransactionParametersResponse sp = rsp.body();
        // Under the covers, an AssetDestroyTransaction is an AssetConfig with all of
        // its
        // configurable fields set to empty
        // All units of the asset _must_ be owned by the creator account and this
        // transaction _must_
        // be issued by the account set to the manager role on the asset
        Transaction destroyTxn = Transaction.AssetDestroyTransactionBuilder().suggestedParams(sp)
                .sender(acct.getAddress())
                .assetIndex(asaId)
                .build();

```
[Snippet Source](https://github.com/barnjamin/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/ASAExamples.java#L180-L193)
    <!-- ===JAVASDK_ASSET_DELETE=== -->

=== "Go"
    <!-- ===GOSDK_ASSET_DELETE=== -->
	``` go  
	// DESTROY ASSET
	// Destroy the asset
	// Make sure all funds are back in the creator's account. Then use the
	// Manager account (Account 1) to destroy the asset.
	// Get network-related transaction parameters and assign
	txParams, err = algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	// txParams.FlatFee = true
	// txParams.Fee = 1000
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
	// Broadcast the transaction to the network
	sendResponse, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	// Wait for transaction to be confirmed
	confirmedTxn, err = transaction.WaitForConfirmation(algodClient,txid,  4, context.Background())
	if err != nil {
		fmt.Printf("Error waiting for confirmation on txID: %s\n", txid)
		return
	}
	fmt.Printf("Confirmed Transaction: %s in Round %d\n", txid ,confirmedTxn.ConfirmedRound)
	fmt.Printf("Asset ID: %d\n", assetID)	
	fmt.Printf("Account 3 must do a transaction for an amount of 0, \n" )
    fmt.Printf("with a closeRemainderTo to the creator account, to clear it from its accountholdings. \n")
    fmt.Printf("For Account 1, nothing should print after this as the asset is destroyed on the creator account \n")
	// print created asset and asset holding info for this asset (should not print anything)
	printCreatedAsset(assetID, pks[1], algodClient)
	printAssetHolding(assetID, pks[1], algodClient)
    ```
    <!-- ===GOSDK_ASSET_DELETE=== -->

=== "goal"
    <!-- ===GOAL_ASSET_DELETE=== -->
	``` goal  
    goal asset destroy --creator <creator-address> --manager <asset-manager-address> --asset <asset-name> -d data 
    ```
    <!-- ===GOAL_ASSET_DELETE=== -->


**See also**

- [Anatomy of the Asset Destroy Transaction](../transactions#destroy-an-asset)

# Retrieve asset information
Retrieve an asset's configuration information from the network using the SDKs or `goal`. Additional details are also added to the accounts that own the specific asset and can be listed with standard account information calls.

!!! info
    The code below illustrates getting asset information without the Indexer. If you have the Indexer installed use the Indexer API to [search for asset](../../rest-apis/indexer/#search-assets) information.

=== "JavaScript"
    <!-- ===JSSDK_ASSET_INFO=== -->
	``` javascript
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
    <!-- ===JSSDK_ASSET_INFO=== -->

=== "Python"
    <!-- ===PYSDK_ASSET_INFO=== -->
```python
# Retrieve the asset info of the newly created asset
asset_info = algod_client.asset_info(created_asset)
asset_params: Dict[str, Any] = asset_info["params"]
print(f"Asset Name: {asset_params['name']}")
print(f"Asset params: {list(asset_params.keys())}")
```
[Snippet Source](https://github.com/barnjamin/py-algorand-sdk/blob/doc-examples/_examples/asa.py#L70-L75)
    <!-- ===PYSDK_ASSET_INFO=== -->

=== "Java"
    <!-- ===JAVASDK_ASSET_INFO=== -->
```java
        // Retrieve the asset info of the newly created asset
        Response<Asset> assetResp = algodClient.GetAssetByID(asaId).execute();
        Asset assetInfo = assetResp.body();
        System.out.printf("Asset Name: %s\n", assetInfo.params.name);
```
[Snippet Source](https://github.com/barnjamin/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/ASAExamples.java#L39-L43)
    <!-- ===JAVASDK_ASSET_INFO=== -->

=== "Go"
    <!-- ===GOSDK_ASSET_INFO=== -->
	```go
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
    <!-- ===GOSDK_ASSET_INFO=== -->

=== "goal"
    <!-- ===GOAL_ASSET_INFO=== -->
	``` goal  
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
    <!-- ===GOAL_ASSET_INFO=== -->