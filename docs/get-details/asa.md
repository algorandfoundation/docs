title: Algorand Standard Assets (ASAs)

The Algorand protocol supports the creation of on-chain assets that benefit from the same security, compatibility, speed and ease of use as the Algo. The official name for assets on Algorand is **Algorand Standard Assets (ASA)**.

With Algorand Standard Assets you can represent stablecoins, loyalty points, system credits, and in-game points, just to name a few examples. You can also represent single, unique assets like a deed for a house, collectable items, unique parts on a supply chain, etc. There is also ABI_CODEC functionality to place transfer restrictions on an asset that help support securities, compliance, and certification use cases.



!!! info
    Assets that represent many of the same type, like a stablecoin, may be referred to as **fungible assets**. Single, unique assets are referred to as **non-fungible assets**. 


This section begins with an [overview](#assets-overview) of the asset implementation on Algorand including a review of all [asset parameters](#asset-parameters). This is followed by [how-tos](#asset-functions) in the SDKs and `goal` for all on-chain asset functions.

# Quick start videos

If you prefer videos, take a look at this 7 minute guide to learn about Introduction to Assets.

<iframe width="100%" style="aspect-ratio:16/9" src="https://www.youtube-nocookie.com/embed/94XMLuTt8nM" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>

# Assets overview

Here are several things to be aware of before getting started with assets.

- For every asset an account creates or owns, its minimum balance is increased by 0.1 Algos (100,000 microAlgos). 
- This minimum balance requirement will be placed on the original creator as long as the asset has not been destroyed. Transferring the asset does not alleviate the creator's minimum balance requirement.
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

# Quick start videos

If you prefer videos, take a look at this 8 minute guide to learn about Building Solutions Using ASAs.

<iframe width="100%" style="aspect-ratio:16/9" src="https://www.youtube-nocookie.com/embed/_Q37eysrRh4" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>

## Creating an asset
**Transaction Authorizer**: Any account with sufficient Algo balance

Create assets using either the SDKs or `goal`. When using the SDKs supply all creation parameters. With `goal`, managing the various addresses associated with the asset must be done after executing an asset creation. See Modifying an Asset in the next section for more details on changing addresses for the asset.

=== "JavaScript"
    <!-- ===JSSDK_ASSET_CREATE=== -->
	```javascript
	const suggestedParams = await algodClient.getTransactionParams().do();
	const txn = algosdk.makeAssetCreateTxnWithSuggestedParamsFromObject({
	  from: creator.addr,
	  suggestedParams,
	  defaultFrozen: false,
	  unitName: 'rug',
	  assetName: 'Really Useful Gift',
	  manager: creator.addr,
	  reserve: creator.addr,
	  freeze: creator.addr,
	  clawback: creator.addr,
	  assetURL: 'http://path/to/my/asset/details',
	  total: 1000,
	  decimals: 0,
	});
	
	const signedTxn = txn.signTxn(creator.privateKey);
	await algodClient.sendRawTransaction(signedTxn).do();
	const result = await algosdk.waitForConfirmation(
	  algodClient,
	  txn.txID().toString(),
	  3
	);
	
	const assetIndex = result['asset-index'];
	console.log(`Asset ID created: ${assetIndex}`);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/asa.ts#L17-L43)
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
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/asa.py#L13-L43)
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
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/ASAExamples.java#L47-L77)
    <!-- ===JAVASDK_ASSET_CREATE=== -->

=== "Go"
    <!-- ===GOSDK_ASSET_CREATE=== -->
	```go
	// Configure parameters for asset creation
	var (
		creatorAddr       = creator.Address.String()
		assetName         = "Really Useful Gift"
		unitName          = "rug"
		assetURL          = "https://path/to/my/asset/details"
		assetMetadataHash = "thisIsSomeLength32HashCommitment"
		defaultFrozen     = false
		decimals          = uint32(0)
		totalIssuance     = uint64(1000)
	
		manager  = creatorAddr
		reserve  = creatorAddr
		freeze   = creatorAddr
		clawback = creatorAddr
	
		note []byte
	)
	
	// Get network-related transaction parameters and assign
	txParams, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		log.Fatalf("error getting suggested tx params: %s", err)
	}
	
	// Construct the transaction
	txn, err := transaction.MakeAssetCreateTxn(
		creatorAddr, note, txParams, totalIssuance, decimals,
		defaultFrozen, manager, reserve, freeze, clawback,
		unitName, assetName, assetURL, assetMetadataHash,
	)
	
	if err != nil {
		log.Fatalf("failed to make transaction: %s", err)
	}
	
	// sign the transaction
	txid, stx, err := crypto.SignTransaction(creator.PrivateKey, txn)
	if err != nil {
		log.Fatalf("failed to sign transaction: %s", err)
	}
	
	// Broadcast the transaction to the network
	_, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		log.Fatalf("failed to send transaction: %s", err)
	}
	
	// Wait for confirmation
	confirmedTxn, err := transaction.WaitForConfirmation(algodClient, txid, 4, context.Background())
	if err != nil {
		log.Fatalf("error waiting for confirmation:  %s", err)
	}
	
	log.Printf("Create Transaction: %s confirmed in Round %d with new asset id: %d\n",
		txid, confirmedTxn.ConfirmedRound, confirmedTxn.AssetIndex)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/asa/main.go#L42-L98)
    <!-- ===GOSDK_ASSET_CREATE=== -->

=== "goal"
    <!-- ===GOAL_ASSET_CREATE=== -->
    ``` goal
        goal asset create --creator <address> --total 1000 --unitname <unit-name> --asseturl "https://path/to/my/asset/details" --decimals 0   -d data
    ```
    <!-- ===GOAL_ASSET_CREATE=== -->


!!!note
	The Algorand Foundation [hosts many standards (ARCs)](https://arc.algorand.foundation/) associated with asset creation. Conforming to these standards allows your apps/assets to work well with existing community tools that support them.

**See also**

- [Anatomy of an Asset Creation Transaction](../transactions#create-an-asset)


## Modifying an asset

**Authorized by**: [Asset Manager Account](../transactions/transactions#manageraddr)

After an asset has been created only the manager, reserve, freeze and clawback accounts can be changed. All other parameters are locked for the life of the asset. If any of these addresses are set to `""` that address will be cleared and can never be reset for the life of the asset. Only the manager account can make configuration changes and must authorize the transaction.

=== "JavaScript"
    <!-- ===JSSDK_ASSET_CONFIG=== -->
	```javascript
	const manager = accounts[1];
	
	const configTxn = algosdk.makeAssetConfigTxnWithSuggestedParamsFromObject({
	  from: creator.addr,
	  manager: manager.addr,
	  freeze: manager.addr,
	  clawback: manager.addr,
	  reserve: undefined,
	  suggestedParams,
	  assetIndex,
	  // don't throw error if freeze, clawback, or manager are empty
	  strictEmptyAddressChecking: false,
	});
	
	const signedConfigTxn = configTxn.signTxn(creator.privateKey);
	await algodClient.sendRawTransaction(signedConfigTxn).do();
	const configResult = await algosdk.waitForConfirmation(
	  algodClient,
	  txn.txID().toString(),
	  3
	);
	console.log(`Result confirmed in round: ${configResult['confirmed-round']}`);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/asa.ts#L60-L82)
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
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/asa.py#L46-L66)
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
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/ASAExamples.java#L83-L96)
    <!-- ===JAVASDK_ASSET_CONFIG=== -->

=== "Go"
    <!-- ===GOSDK_ASSET_CONFIG=== -->
	```go
	creatorAddr := creator.Address.String()
	var (
		newManager  = creatorAddr
		newFreeze   = creatorAddr
		newClawback = creatorAddr
		newReserve  = ""
	
		strictAddrCheck = false
		note            []byte
	)
	
	// Get network-related transaction parameters and assign
	sp, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		log.Fatalf("error getting suggested tx params: %s", err)
	}
	
	txn, err := transaction.MakeAssetConfigTxn(creatorAddr, note, sp, assetID, newManager, newReserve, newFreeze, newClawback, strictAddrCheck)
	if err != nil {
		log.Fatalf("failed to make  txn: %s", err)
	}
	// sign the transaction
	txid, stx, err := crypto.SignTransaction(creator.PrivateKey, txn)
	if err != nil {
		log.Fatalf("failed to sign transaction: %s", err)
	}
	
	// Broadcast the transaction to the network
	_, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		log.Fatalf("failed to send transaction: %s", err)
	}
	
	// Wait for confirmation
	confirmedTxn, err := transaction.WaitForConfirmation(algodClient, txid, 4, context.Background())
	if err != nil {
		log.Fatalf("error waiting for confirmation:  %s", err)
	}
	
	log.Printf("Asset Config Transaction: %s confirmed in Round %d\n", txid, confirmedTxn.ConfirmedRound)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/asa/main.go#L104-L144)
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
	```javascript
	
	// opt-in is simply a 0 amount transfer of the asset to oneself
	const optInTxn = algosdk.makeAssetTransferTxnWithSuggestedParamsFromObject({
	  from: receiver.addr,
	  to: receiver.addr,
	  suggestedParams,
	  assetIndex,
	  amount: 0,
	});
	
	const signedOptInTxn = optInTxn.signTxn(receiver.privateKey);
	await algodClient.sendRawTransaction(signedOptInTxn).do();
	await algosdk.waitForConfirmation(algodClient, optInTxn.txID().toString(), 3);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/asa.ts#L86-L99)
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
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/asa.py#L79-L92)
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
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/ASAExamples.java#L103-L112)
    <!-- ===JAVASDK_ASSET_OPTIN=== -->

=== "Go"
    <!-- ===GOSDK_ASSET_OPTIN=== -->
	```go
	userAddr := user.Address.String()
	
	sp, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		log.Fatalf("error getting suggested tx params: %s", err)
	}
	
	txn, err := transaction.MakeAssetAcceptanceTxn(userAddr, nil, sp, assetID)
	if err != nil {
		log.Fatalf("failed to make txn: %s", err)
	}
	// sign the transaction
	txid, stx, err := crypto.SignTransaction(user.PrivateKey, txn)
	if err != nil {
		log.Fatalf("failed to sign transaction: %s", err)
	}
	
	// Broadcast the transaction to the network
	_, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		log.Fatalf("failed to send transaction: %s", err)
	}
	
	// Wait for confirmation
	confirmedTxn, err := transaction.WaitForConfirmation(algodClient, txid, 4, context.Background())
	if err != nil {
		log.Fatalf("error waiting for confirmation:  %s", err)
	}
	
	log.Printf("OptIn Transaction: %s confirmed in Round %d\n", txid, confirmedTxn.ConfirmedRound)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/asa/main.go#L149-L179)
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
	```javascript
	const xferTxn = algosdk.makeAssetTransferTxnWithSuggestedParamsFromObject({
	  from: creator.addr,
	  to: receiver.addr,
	  suggestedParams,
	  assetIndex,
	  amount: 1,
	});
	
	const signedXferTxn = xferTxn.signTxn(creator.privateKey);
	await algodClient.sendRawTransaction(signedXferTxn).do();
	await algosdk.waitForConfirmation(algodClient, xferTxn.txID().toString(), 3);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/asa.ts#L102-L113)
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
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/asa.py#L105-L120)
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
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/ASAExamples.java#L120-L131)
    <!-- ===JAVASDK_ASSET_XFER=== -->

=== "Go"
    <!-- ===GOSDK_ASSET_XFER=== -->
	```go
	var (
		creatorAddr = creator.Address.String()
		userAddr    = user.Address.String()
	)
	
	sp, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		log.Fatalf("error getting suggested tx params: %s", err)
	}
	
	txn, err := transaction.MakeAssetTransferTxn(creatorAddr, userAddr, 1, nil, sp, "", assetID)
	if err != nil {
		log.Fatalf("failed to make asset txn: %s", err)
	}
	// sign the transaction
	txid, stx, err := crypto.SignTransaction(creator.PrivateKey, txn)
	if err != nil {
		log.Fatalf("failed to sign transaction: %s", err)
	}
	
	// Broadcast the transaction to the network
	_, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		log.Fatalf("failed to send transaction: %s", err)
	}
	
	// Wait for confirmation
	confirmedTxn, err := transaction.WaitForConfirmation(algodClient, txid, 4, context.Background())
	if err != nil {
		log.Fatalf("error waiting for confirmation:  %s", err)
	}
	
	log.Printf("Asset Transfer Transaction: %s confirmed in Round %d\n", txid, confirmedTxn.ConfirmedRound)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/asa/main.go#L219-L252)
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
	```javascript
	const freezeTxn = algosdk.makeAssetFreezeTxnWithSuggestedParamsFromObject({
	  from: manager.addr,
	  suggestedParams,
	  assetIndex,
	  // freezeState: false would unfreeze the account's asset holding
	  freezeState: true,
	  // freezeTarget is the account that is being frozen or unfrozen
	  freezeTarget: receiver.addr,
	});
	
	const signedFreezeTxn = freezeTxn.signTxn(manager.privateKey);
	await algodClient.sendRawTransaction(signedFreezeTxn).do();
	await algosdk.waitForConfirmation(
	  algodClient,
	  freezeTxn.txID().toString(),
	  3
	);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/asa.ts#L116-L133)
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
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/asa.py#L131-L146)
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
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/ASAExamples.java#L139-L151)
    <!-- ===JAVASDK_ASSET_FREEZE=== -->

=== "Go"
    <!-- ===GOSDK_ASSET_FREEZE=== -->
	```go
	var (
		creatorAddr = creator.Address.String()
		userAddr    = user.Address.String()
	)
	
	sp, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		log.Fatalf("error getting suggested tx params: %s", err)
	}
	
	// Create a freeze asset transaction with the target of the user address
	// and the new freeze setting of `true`
	txn, err := transaction.MakeAssetFreezeTxn(creatorAddr, nil, sp, assetID, userAddr, true)
	if err != nil {
		log.Fatalf("failed to make txn: %s", err)
	}
	// sign the transaction
	txid, stx, err := crypto.SignTransaction(creator.PrivateKey, txn)
	if err != nil {
		log.Fatalf("failed to sign transaction: %s", err)
	}
	
	// Broadcast the transaction to the network
	_, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		log.Fatalf("failed to send transaction: %s", err)
	}
	
	// Wait for confirmation
	confirmedTxn, err := transaction.WaitForConfirmation(algodClient, txid, 4, context.Background())
	if err != nil {
		log.Fatalf("error waiting for confirmation:  %s", err)
	}
	
	log.Printf("Freeze Transaction: %s confirmed in Round %d\n", txid, confirmedTxn.ConfirmedRound)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/asa/main.go#L257-L292)
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
	```javascript
	const clawbackTxn = algosdk.makeAssetTransferTxnWithSuggestedParamsFromObject(
	  {
	    from: manager.addr,
	    to: creator.addr,
	    // revocationTarget is the account that is being clawed back from
	    revocationTarget: receiver.addr,
	    suggestedParams,
	    assetIndex,
	    amount: 1,
	  }
	);
	
	const signedClawbackTxn = clawbackTxn.signTxn(manager.privateKey);
	await algodClient.sendRawTransaction(signedClawbackTxn).do();
	await algosdk.waitForConfirmation(
	  algodClient,
	  clawbackTxn.txID().toString(),
	  3
	);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/asa.ts#L136-L155)
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
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/asa.py#L157-L173)
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
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/ASAExamples.java#L177-L190)
    <!-- ===JAVASDK_ASSET_CLAWBACK=== -->

=== "Go"
    <!-- ===GOSDK_ASSET_CLAWBACK=== -->
	```go
	var (
		creatorAddr = creator.Address.String()
		userAddr    = user.Address.String()
	)
	
	sp, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		log.Fatalf("error getting suggested tx params: %s", err)
	}
	
	// Create a new clawback transaction with the target of the user address and the recipient as the creator
	// address, being sent from the address marked as `clawback` on the asset, in this case the same as creator
	txn, err := transaction.MakeAssetRevocationTxn(creatorAddr, userAddr, 1, creatorAddr, nil, sp, assetID)
	if err != nil {
		log.Fatalf("failed to make txn: %s", err)
	}
	// sign the transaction
	txid, stx, err := crypto.SignTransaction(creator.PrivateKey, txn)
	if err != nil {
		log.Fatalf("failed to sign transaction: %s", err)
	}
	
	// Broadcast the transaction to the network
	_, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		log.Fatalf("failed to send transaction: %s", err)
	}
	
	// Wait for confirmation
	confirmedTxn, err := transaction.WaitForConfirmation(algodClient, txid, 4, context.Background())
	if err != nil {
		log.Fatalf("error waiting for confirmation:  %s", err)
	}
	
	log.Printf("Clawback Transaction: %s confirmed in Round %d\n", txid, confirmedTxn.ConfirmedRound)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/asa/main.go#L297-L332)
    <!-- ===GOSDK_ASSET_CLAWBACK=== -->

=== "goal"
    <!-- ===GOAL_ASSET_CLAWBACK=== -->
	``` goal  
    goal asset send -a <amount-to-revoke> --asset <asset-name> -f <address-of-revoke-target> -t <address-to-send-assets-to> --clawback <clawback-address> --creator <creator-address> -d data
    ```
    <!-- ===GOAL_ASSET_CLAWBACK=== -->

**See also**

- [Anatomy of an Asset Clawback Transaction](../transactions#revoke-an-asset)

## Opting Out of an Asset

**Authorized by**: The account opting out 

An account can opt out of an asset at any time. This means that the account will no longer hold the asset, and the account will no longer be able to receive the asset. The account also recovers the Minimum Balance Requirement for the asset (0.1A).

=== "JavaScript"
	<!-- ===JSSDK_ASSET_OPT_OUT=== -->
	```javascript
	
	// opt-out is an amount transfer with the `closeRemainderTo` field set to
	// any account that can receive the asset.
	// note that closing to the asset creator will always succeed
	const optOutTxn = algosdk.makeAssetTransferTxnWithSuggestedParamsFromObject({
	  from: receiver.addr,
	  to: creator.addr,
	  closeRemainderTo: creator.addr,
	  suggestedParams,
	  assetIndex,
	  amount: 0,
	});
	
	const signedOptOutTxn = optOutTxn.signTxn(receiver.privateKey);
	await algodClient.sendRawTransaction(signedOptOutTxn).do();
	await algosdk.waitForConfirmation(
	  algodClient,
	  optOutTxn.txID().toString(),
	  3
	);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/asa.ts#L158-L178)
	<!-- ===JSSDK_ASSET_OPT_OUT=== -->

=== "Python"
	<!-- ===PYSDK_ASSET_OPT_OUT=== -->
	```python
	sp = algod_client.suggested_params()
	opt_out_txn = transaction.AssetTransferTxn(
	    sender=acct2.address,
	    sp=sp,
	    index=created_asset,
	    receiver=acct1.address,
	    # an opt out transaction sets its close_asset_to parameter
	    # it is always possible to close an asset to the creator
	    close_assets_to=acct1.address,
	    amt=0,
	)
	signed_opt_out = opt_out_txn.sign(acct2.private_key)
	txid = algod_client.send_transaction(signed_opt_out)
	print(f"Sent opt out transaction with txid: {txid}")
	
	results = transaction.wait_for_confirmation(algod_client, txid, 4)
	print(f"Result confirmed in round: {results['confirmed-round']}")
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/asa.py#L185-L202)
	<!-- ===PYSDK_ASSET_OPT_OUT=== -->

=== "Go"
	<!-- ===GOSDK_ASSET_OPT_OUT=== -->
	```go
	userAddr := user.Address.String()
	
	sp, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		log.Fatalf("error getting suggested tx params: %s", err)
	}
	
	txn, err := transaction.MakeAssetTransferTxn(userAddr, creator.Address.String(), 0, nil, sp, creator.Address.String(), assetID)
	if err != nil {
		log.Fatalf("failed to make txn: %s", err)
	}
	// sign the transaction
	txid, stx, err := crypto.SignTransaction(user.PrivateKey, txn)
	if err != nil {
		log.Fatalf("failed to sign transaction: %s", err)
	}
	
	// Broadcast the transaction to the network
	_, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		log.Fatalf("failed to send transaction: %s", err)
	}
	
	// Wait for confirmation
	confirmedTxn, err := transaction.WaitForConfirmation(algodClient, txid, 4, context.Background())
	if err != nil {
		log.Fatalf("error waiting for confirmation:  %s", err)
	}
	
	log.Printf("OptOut Transaction: %s confirmed in Round %d\n", txid, confirmedTxn.ConfirmedRound)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/asa/main.go#L184-L214)
	<!-- ===GOSDK_ASSET_OPT_OUT=== -->

=== "Java"
	<!-- ===JAVASDK_ASSET_OPT_OUT=== -->
	```java
	Response<TransactionParametersResponse> rsp = algodClient.TransactionParams().execute();
	TransactionParametersResponse sp = rsp.body();
	// Opt out of the asset by setting the assetCloseTo parameter
	Transaction optOutTxn = Transaction.AssetTransferTransactionBuilder().suggestedParams(sp)
	        .sender(sender.getAddress())
	        .assetReceiver(creator.getAddress())
	        .assetCloseTo(creator.getAddress())
	        .assetIndex(asaId)
	        .assetAmount(0)
	        .build();
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/ASAExamples.java#L159-L169)
	<!-- ===JAVASDK_ASSET_OPT_OUT=== -->


## Destroying an asset

**Authorized by**: [Asset Manager](../transactions/transactions#manageraddr)

Created assets can be destroyed only by the asset manager account. All of the assets must be owned by the creator of the asset before the asset can be deleted. 

=== "JavaScript"
    <!-- ===JSSDK_ASSET_DELETE=== -->
	```javascript
	const deleteTxn = algosdk.makeAssetDestroyTxnWithSuggestedParamsFromObject({
	  from: manager.addr,
	  suggestedParams,
	  assetIndex,
	});
	
	const signedDeleteTxn = deleteTxn.signTxn(manager.privateKey);
	await algodClient.sendRawTransaction(signedDeleteTxn).do();
	await algosdk.waitForConfirmation(
	  algodClient,
	  deleteTxn.txID().toString(),
	  3
	);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/asa.ts#L181-L194)
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
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/asa.py#L206-L225)
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
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/ASAExamples.java#L197-L210)
    <!-- ===JAVASDK_ASSET_DELETE=== -->

=== "Go"
    <!-- ===GOSDK_ASSET_DELETE=== -->
	```go
	var (
		creatorAddr = creator.Address.String()
	)
	
	sp, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		log.Fatalf("error getting suggested tx params: %s", err)
	}
	
	// Create a new clawback transaction with the target of the user address and the recipient as the creator
	// address, being sent from the address marked as `clawback` on the asset, in this case the same as creator
	txn, err := transaction.MakeAssetDestroyTxn(creatorAddr, nil, sp, assetID)
	if err != nil {
		log.Fatalf("failed to make txn: %s", err)
	}
	// sign the transaction
	txid, stx, err := crypto.SignTransaction(creator.PrivateKey, txn)
	if err != nil {
		log.Fatalf("failed to sign transaction: %s", err)
	}
	
	// Broadcast the transaction to the network
	_, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		log.Fatalf("failed to send transaction: %s", err)
	}
	
	// Wait for confirmation
	confirmedTxn, err := transaction.WaitForConfirmation(algodClient, txid, 4, context.Background())
	if err != nil {
		log.Fatalf("error waiting for confirmation:  %s", err)
	}
	
	log.Printf("Destroy Transaction: %s confirmed in Round %d\n", txid, confirmedTxn.ConfirmedRound)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/asa/main.go#L337-L371)
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
	```javascript
	const assetInfo = await algodClient.getAssetByID(assetIndex).do();
	console.log(`Asset Name: ${assetInfo.params.name}`);
	console.log(`Asset Params: ${assetInfo.params}`);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/asa.ts#L46-L49)
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
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/asa.py#L70-L75)
    <!-- ===PYSDK_ASSET_INFO=== -->

=== "Java"
    <!-- ===JAVASDK_ASSET_INFO=== -->
	```java
	// Retrieve the asset info of the newly created asset
	Response<Asset> assetResp = algodClient.GetAssetByID(asaId).execute();
	Asset assetInfo = assetResp.body();
	System.out.printf("Asset Name: %s\n", assetInfo.params.name);
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/ASAExamples.java#L38-L42)
    <!-- ===JAVASDK_ASSET_INFO=== -->

=== "Go"
    <!-- ===GOSDK_ASSET_INFO=== -->
	```go
	info, err := algodClient.GetAssetByID(assetID).Do(context.Background())
	if err != nil {
		log.Fatalf("failed to get asset info: %s", err)
	}
	log.Printf("Asset info for %d: %+v", assetID, info)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/asa/main.go#L25-L30)
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