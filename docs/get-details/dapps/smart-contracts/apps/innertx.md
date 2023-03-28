title: Inner transactions
# Issuing Transactions from an Application 

When a smart contract is deployed to the Algorand blockchain it is assigned a unique identifier, called the app id. Additionally, every smart contract has a unique Algorand address that is generated from this specific ID. The address allows the smart contract to function as an escrow account. To see the specific address of the smart contract the `goal app info` command can be used.

```bash
% ./goal app info --app-id 1 -d data
Application ID:        1
Application account:   WCS6TVPJRBSARHLN2326LRU5BYVJZUKI2VJ53CAWKYYHDE455ZGKANWMGM
Creator:               VJG4SF5O4DUQDK7MQYHHSJ2HB5LEAB3OQMM2XCUM7OVQKWC7E4GSNCM4FQ
Approval hash:         IYRFKZVCHBQW3UNKEWRQPMCXSXH4TDF4K74EAH65UYLJ5R5JQGVZJ4FF2M
Clear hash:            P7GEWDXXW5IONRW6XRIRVPJCT2XXEQGOBGG65VJPBUOYZEJCBZWTPHS3VQ
Max global byteslices: 0
Max global integers:   1
Max local byteslices:  0
Max local integers:    0
```

The application address will be shown as the Application account. This address can also be retrieved with the SDKs or calculated by using the following code and the smart contract’s application id.

```bash
# app ID of 1’s address
python3 -c "import algosdk.encoding as e; print(e.encode_address(e.checksum(b'appID'+(1).to_bytes(8, 'big'))))"
WCS6TVPJRBSARHLN2326LRU5BYVJZUKI2VJ53CAWKYYHDE455ZGKANWMGM
```

## Inner transactions

To fund this account, any other account in the Algorand network can send algos to the specified account. In order for funds to leave the smart contract, the logic within the contract must submit an inner transaction. In addition, the smart contract’s logic must return true. A smart contract can issue up to a total of 256 inner transactions with one call. If any of these transactions fail, then the smart contract will also fail. Groups of transactions can also be made using inner transactions, which are primarily used when calling other smart contracts that will verify the calling groups transactions. Inner transactions support all the same transaction types as a regular account can make. To generate an inner transaction the `itxn_begin`, `itxn_field`, `itxn_next` and `itxn_submit` opcodes are used. The `itxn_begin` opcode signifies the beginning of an inner transaction. The `itxn_field` opcode is used to set specific transaction properties. The `itxn_next` opcode moves to the next transaction in the same group as the previous, and the `itxn_submit` opcode is used to submit the transaction or transaction group. As an example, the following contract code generates a simple payment transaction.

=== "PyTeal"
	<!-- ===PYTEAL_ITXN_PAYMENT=== -->
	```python
	        # ...
	        InnerTxnBuilder.Begin(),
	        InnerTxnBuilder.SetFields(
	            {
	                TxnField.type_enum: TxnType.Payment,
	                TxnField.amount: Int(5000),
	                TxnField.receiver: Txn.sender(),
	            }
	        ),
	        InnerTxnBuilder.Submit(),
	        # ...
	        # The `Sender` for the above is implied to be Global.current_application_address().
	        # If a different sender is needed, it'd have to be an account that has been rekeyed to
	        # the application address.
	```
	[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/itxn.py#L22-L36)
	<!-- ===PYTEAL_ITXN_PAYMENT=== -->

=== "TEAL"
	<!-- ===TEAL_ITXN_PAYMENT=== -->
	```teal
	itxn_begin
	
	int pay
	itxn_field TypeEnum
	
	int 1000000
	itxn_field Amount
	
	txn Sender
	itxn_field Receiver
	
	itxn_submit
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/itxn_payment/approval.teal#L6-L18)
	<!-- ===TEAL_ITXN_PAYMENT=== -->

 Fees for these transactions are paid by the smart contract and are set automatically to the minimum transaction fee. Inner transaction fees are eligible for [fee pooling](https://developer.algorand.org/docs/get-details/transactions/#pooled-transaction-fees) similar to any other transaction. This allows either the application call or any other transaction in a group of transactions to pay the fee for inner transactions. Inner transactions are evaluated during AVM execution, allowing changes to be visible within the contract. For example, if the ‘balance’ opcode is used before and after a ‘pay’ transaction is submitted, the balance change would be visible to the executing contract.

!!!note
    Inner transactions also have access to the Sender field. It is not required to set this field as all inner transactions default the sender to the contract address. If another account is rekeyed to the smart contract address, setting sender to the address that has been rekeyed allows the contract to spend from that account. The recipient of an inner transaction must be in the accounts array. Additionally, if the sender of an inner transaction is not the contract, the sender must also be in the accounts array.

!!!note
    Clear state programs do _not_ support creating inner transactions.  However, clear state programs _can_ be called by an inner transaction.


## Allowed transaction properties
Since TEAL 6, all transaction types can be used within inner transactions. If you're using TEAL 5 you will only be able to make payment and asset transfer transactions, with some properties such as `RekeyTo` not being allowed.

## Asset transfer
If a smart contract wishes to transfer an asset it holds or needs to opt into an asset this can be done with an asset transfer inner transaction.

=== "PyTeal"
	<!-- ===PYTEAL_ITXN_ASSET_TRANSFER=== -->
	```python
	        # ...
	        InnerTxnBuilder.Begin(),
	        InnerTxnBuilder.SetFields(
	            {
	                TxnField.type_enum: TxnType.AssetTransfer,
	                TxnField.asset_amount: Int(5000),
	                TxnField.asset_receiver: Txn.sender(),
	                TxnField.xfer_asset: Txn.assets[0],
	            }
	        ),
	        # ...
	        InnerTxnBuilder.Submit(),
	```
	[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/itxn.py#L44-L56)
	<!-- ===PYTEAL_ITXN_ASSET_TRANSFER=== -->

=== "TEAL"
	<!-- ===TEAL_ITXN_ASSET_TRANSFER=== -->
	```teal
	itxn_begin
	
	int axfer
	itxn_field TypeEnum
	
	txn Assets 0
	itxn_field XferAsset
	
	txn Accounts 1
	itxn_field AssetReceiver
	
	txn ApplicationArgs 3
	btoi
	itxn_field AssetAmount
	
	itxn_submit
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/itxn_asset_management/approval.teal#L293-L309)
	<!-- ===TEAL_ITXN_ASSET_TRANSFER=== -->

Note that the asset must be in the assets array. If the smart contract is opting into an asset, the contract would send 0 units of the asset to itself. In this case, the receiver could be set to the `global CurrentApplicationAddress`. 

## Asset freeze
A smart contract can freeze any asset, where the smart contract is the freeze address. This can be done with the following TEAL.

=== "PyTeal"
	<!-- ===PYTEAL_ITXN_ASSET_FREEZE=== -->
	```python
	        # ...
	        InnerTxnBuilder.Begin(),
	        InnerTxnBuilder.SetFields(
	            {
	                TxnField.type_enum: TxnType.AssetFreeze,
	                TxnField.freeze_asset: Txn.assets[0],
	                TxnField.freeze_asset_account: Txn.accounts[1],
	                TxnField.freeze_asset_frozen: Int(1),
	            }
	        ),
	        InnerTxnBuilder.Submit(),
	        # ...
	```
	[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/itxn.py#L64-L76)
	<!-- ===PYTEAL_ITXN_ASSET_FREEZE=== -->

=== "TEAL"
	<!-- ===TEAL_ITXN_ASSET_FREEZE=== -->
	```teal
	itxn_begin
	
	int afrz
	itxn_field TypeEnum
	
	txn Assets 0
	itxn_field FreezeAsset
	
	txn Accounts 1
	itxn_field FreezeAssetAccount
	
	// Flip the current account frozen state
	txn Accounts 1
	txn Assets 0
	asset_holding_get AssetFrozen
	assert
	!
	itxn_field FreezeAssetFrozen
	
	itxn_submit
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/itxn_asset_management/approval.teal#L158-L178)
	<!-- ===TEAL_ITXN_ASSET_FREEZE=== -->

## Asset revoke
A smart contract can revoke or clawback any asset where the smart contract address is specified as the asset clawback address. 

=== "PyTeal"
	<!-- ===PYTEAL_ITXN_ASSET_REVOKE=== -->
	```python
	        # ...
	        InnerTxnBuilder.Begin(),
	        InnerTxnBuilder.SetFields(
	            {
	                TxnField.type_enum: TxnType.AssetTransfer,
	                TxnField.asset_receiver: Global.current_application_address(),
	                # AssetSender is _only_ used in the case of clawback
	                # Sender is implied to be current_application_address
	                TxnField.asset_sender: Txn.accounts[1],
	                TxnField.asset_amount: Int(1000),
	            }
	        ),
	        InnerTxnBuilder.Submit(),
	        # ...
	```
	[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/itxn.py#L84-L98)
	<!-- ===PYTEAL_ITXN_ASSET_REVOKE=== -->

=== "TEAL"
	<!-- ===TEAL_ITXN_ASSET_REVOKE=== -->
	```teal
	itxn_begin
	
	int axfer
	itxn_field TypeEnum
	
	txn Assets 0
	itxn_field XferAsset
	
	// Any amount lower or equal to their holding can be revoked
	// Here we use the accounts entire asset balance
	txn Accounts 1
	txn Assets 0
	asset_holding_get AssetBalance
	assert
	itxn_field AssetAmount
	
	txn Accounts 1
	itxn_field AssetSender
	
	global CurrentApplicationAddress
	itxn_field AssetReceiver
	
	itxn_submit
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/itxn_asset_management/approval.teal#L195-L218)
	<!-- ===TEAL_ITXN_ASSET_REVOKE=== -->

## Asset create
Assets can also be created by a smart contract. To create an asset with an inner transaction use the following contract code.

=== "PyTeal"
	<!-- ===PYTEAL_ITXN_ASSET_CREATE=== -->
	```python
	        # ...
	        InnerTxnBuilder.Begin(),
	        InnerTxnBuilder.SetFields(
	            {
	                TxnField.type_enum: TxnType.AssetConfig,
	                TxnField.config_asset_total: Int(1000000),
	                TxnField.config_asset_decimals: Int(3),
	                TxnField.config_asset_unit_name: Bytes("oz"),
	                TxnField.config_asset_name: Bytes("Gold"),
	                TxnField.config_asset_url: Bytes("https://gold.rush"),
	                TxnField.config_asset_manager: Global.current_application_address(),
	                TxnField.config_asset_reserve: Global.current_application_address(),
	                TxnField.config_asset_freeze: Global.current_application_address(),
	                TxnField.config_asset_clawback: Global.current_application_address(),
	            }
	        ),
	        InnerTxnBuilder.Submit(),
	        # ...
	```
	[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/itxn.py#L106-L124)
	<!-- ===PYTEAL_ITXN_ASSET_CREATE=== -->

=== "TEAL"
	<!-- ===TEAL_ITXN_ASSET_CREATE=== -->
	```teal
	itxn_begin
	
	int acfg
	itxn_field TypeEnum
	
	byte "Demo Asset"
	itxn_field ConfigAssetName
	
	byte "DA"
	itxn_field ConfigAssetUnitName
	
	int 100
	itxn_field ConfigAssetTotal
	
	int 2
	itxn_field ConfigAssetDecimals
	
	global CurrentApplicationAddress
	dupn 3
	itxn_field ConfigAssetManager
	itxn_field ConfigAssetReserve
	itxn_field ConfigAssetFreeze
	itxn_field ConfigAssetClawback
	
	itxn_submit
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/itxn_asset_management/approval.teal#L79-L104)
	<!-- ===TEAL_ITXN_ASSET_CREATE=== -->

In this example, a simple asset is created. Using the `itxn CreatedAssetID` opcode after the transaction is submitted allows the contract to get the asset id of the newly created asset.

##Asset configuration
As with all assets, the mutable addresses can be changed using contract code similar to the code below. 

=== "PyTeal"
	<!-- ===PYTEAL_ITXN_ASSET_CONFIG=== -->
	```python
	        # ...
	        InnerTxnBuilder.Begin(),
	        InnerTxnBuilder.SetFields(
	            {
	                TxnField.type_enum: TxnType.AssetConfig,
	                TxnField.config_asset: Txn.assets[0],
	                TxnField.config_asset_manager: Txn.sender(),
	                TxnField.config_asset_reserve: Txn.sender(),
	                TxnField.config_asset_freeze: Txn.sender(),
	                TxnField.config_asset_clawback: Txn.sender(),
	            }
	        ),
	        InnerTxnBuilder.Submit(),
	        # ...
	```
	[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/itxn.py#L132-L146)
	<!-- ===PYTEAL_ITXN_ASSET_CONFIG=== -->

=== "TEAL"
	<!-- ===TEAL_ITXN_ASSET_CONFIG=== -->
	```teal
	itxn_begin
	
	int acfg
	itxn_field TypeEnum
	
	txn Assets 0
	itxn_field ConfigAsset
	
	global CurrentApplicationAddress
	dupn 3
	itxn_field ConfigAssetManager
	itxn_field ConfigAssetReserve
	itxn_field ConfigAssetFreeze
	itxn_field ConfigAssetClawback
	
	itxn_submit
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/itxn_asset_management/approval.teal#L125-L141)
	<!-- ===TEAL_ITXN_ASSET_CONFIG=== -->

!!!Warning
	Note that when changing one address, all others must be reset or they will be cleared. Cleared addresses will be locked forever.

## Delete an asset
Assets managed by the contract can also be deleted. This can be done with the following contract code.

=== "PyTeal"
	<!-- ===PYTEAL_ITXN_ASSET_DESTROY=== -->
	```python
	        # ...
	        InnerTxnBuilder.Begin(),
	        InnerTxnBuilder.SetFields(
	            {
	                TxnField.type_enum: TxnType.AssetConfig,
	                TxnField.config_asset: Txn.assets[0],
	            }
	        ),
	        InnerTxnBuilder.Submit(),
	        # ...
	```
	[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/itxn.py#L154-L164)
	<!-- ===PYTEAL_ITXN_ASSET_DESTROY=== -->

=== "TEAL"
	<!-- ===TEAL_ITXN_ASSET_DESTROY=== -->
	```teal
	itxn_begin
	
	int acfg
	itxn_field TypeEnum
	
	txn Assets 0
	itxn_field ConfigAsset
	
	itxn_submit
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/itxn_asset_management/approval.teal#L235-L244)
	<!-- ===TEAL_ITXN_ASSET_DESTROY=== -->

## Grouped inner transaction

A smart contract can make inner transactions consisting of grouped transactions. The following example groups a payment transaction with a call to another smart contract.

=== "PyTeal"
	<!-- ===PYTEAL_GROUPED_ITXN=== -->
	```python
	        # This returns a `MaybeValue`, see pyteal docs
	        addr := AppParam.address(Int(1234)),
	        Assert(addr.hasValue()),
	        # ...
	        InnerTxnBuilder.Begin(),
	        InnerTxnBuilder.SetFields(
	            {
	                TxnField.type_enum: TxnType.Payment,
	                TxnField.receiver: addr.value(),
	                TxnField.amount: Int(1000000),
	            }
	        ),
	        InnerTxnBuilder.Next(),  # This indicates we're moving to constructing the next txn in the group
	        InnerTxnBuilder.SetFields(
	            {
	                TxnField.type_enum: TxnType.ApplicationCall,
	                TxnField.application_id: Int(1234),
	                TxnField.on_completion: OnComplete.NoOp,
	                # Note this is _not_ using the ABI to call the
	                # method in the other app
	                TxnField.application_args: [Bytes("buy")],
	            }
	        ),
	        InnerTxnBuilder.Submit(),
	        # ...
	```
	[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/itxn.py#L172-L197)
	<!-- ===PYTEAL_GROUPED_ITXN=== -->

=== "TEAL"
	<!-- ===TEAL_GROUPED_ITXN=== -->
	```teal
	itxn_begin
	
	int pay
	itxn_field TypeEnum
	
	int 1000000
	itxn_field Amount
	
	int 123
	app_params_get AppAddress
	assert
	itxn_field Receiver
	
	itxn_next
	
	int appl
	itxn_field TypeEnum
	
	int 123
	itxn_field ApplicationID
	
	int NoOp
	itxn_field OnCompletion
	
	byte "buy"
	itxn_field ApplicationArgs
	
	itxn_submit
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/itxn_groups/approval.teal#L3-L31)
	<!-- ===TEAL_GROUPED_ITXN=== -->

All inner transactions will be stored as inner transactions within the outer application transaction. These can be accessed by getting the transaction id as normal and looking for the `inner-txns` header in the transaction response.

# Contract To Contract Calls

With the release of TEAL 6 (AVM 1.1), Smart Contracts may issue inner transactions that invoke other Smart Contracts. This allows for composability across applications but comes with some limitations.

* An application may not call itself, even indirectly. This is referred to as `re-entrancy` and is explicitly forbidden. 
* An application may only call into other applications up to a stack depth of 8. In other words if app calls (`->`) look like 1->2->3->4->5->6->7->8, App 8 may _not_ call another application. This would violate the stack depth limit.
* An application may issue up to 256 inner transactions to increase its budget (max budget of 179.2k even for a group size of 1), but the max call budget is shared for all applications in the group. Meaning you can't have two app calls in the same group that _both_ try to issue 256 inner app calls. 
* An application of program version 6 or above may _not_ call contracts with a program version 3 or below. This limitation protects an older application from unexpected behavior introduced in newer program versions.

## Application call
A smart contract can call other smart contracts using any of the `OnComplete` types. This allows a smart contract to create, opt in, close out, clear state, delete, or just call (NoOp) other smart contracts. To call an existing smart contract the following contract code can be used.

=== "PyTeal"
	<!-- ===PYTEAL_ITXN_C2C=== -->
	```python
	        # ...
	        InnerTxnBuilder.Begin(),
	        InnerTxnBuilder.SetFields(
	            {
	                TxnField.type_enum: TxnType.ApplicationCall,
	                TxnField.application_id: Int(1234),
	                TxnField.on_completion: OnComplete.NoOp,
	            }
	        ),
	        InnerTxnBuilder.Submit(),
	        # ...
	```
	[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/itxn.py#L205-L216)
	<!-- ===PYTEAL_ITXN_C2C=== -->

=== "TEAL"
	<!-- ===TEAL_ITXN_C2C=== -->
	```teal
	itxn_begin
	
	int appl
	itxn_field TypeEnum
	
	txn Applications 1
	itxn_field ApplicationID
	
	int NoOp
	itxn_field OnCompletion
	
	itxn_submit
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/itxn_c2c/approval.teal#L8-L20)
	<!-- ===TEAL_ITXN_C2C=== -->

## Composability

When writing smart contracts that call other applications or expect to be called via Inner Transactions, an important consideration is composability.

With the finalization of the [ABI](/docs/get-details/dapps/smart-contracts/ABI/) an API may be defined for an application. This allows contracts to be written to take advantage of the ABI to provide structured calls to other applications. 

Additionally, when validating transactions, using relative position of transactions instead of absolute position will help to allow behavior to be composed. 

Since TEAL 6, all created assets and apps are available to be accessed by application calls in the same group which allows more dynamic behavior. For example an application can be created via Inner Transaction then funded immediately in the same transaction group since we have access to the created application id and address.   



