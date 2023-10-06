title: Evaluating transactions


# Checking the transaction type in a smart contract

The `ApplicationCall` transaction types defined in [The Lifecycle of a Smart Contract](#the-lifecycle-of-a-smart-contract) can be checked within the TEAL code by examining the `OnCompletion` transaction property. 

=== "PyTeal"
	<!-- ===PYTEAL_TXN_ONCOMPLETE=== -->
	```python
	    program = OnComplete.NoOp == Txn.on_completion()
	    print(compileTeal(program, Mode.Application))
	```
	[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/txn.py#L20-L22)
	<!-- ===PYTEAL_TXN_ONCOMPLETE=== -->

=== "TEAL"
	<!-- ===TEAL_TXN_ONCOMPLETE=== -->
	```teal
	txn OnCompletion
	int NoOp // OptIn, CloseOut, UpdateApplication, or DeleteApplication
	==
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/oncomplete/approval.teal#L3-L6)
	<!-- ===TEAL_TXN_ONCOMPLETE=== -->

# Global values in smart contracts
Smart contracts have access to many global variables. These variables are set for the blockchain, like the minimum transaction fee (MinTxnFee). As another example of Global variable use, in the [Atomic Transfers and Transaction Properties](#atomic-transfers-and-transaction-properties) section of this guide, `GroupSize` is used to show how to get the number of transactions that are grouped within a smart contract call. Smart contracts also have access to the `LatestTimestamp` global which represents the latest confirmed block's Unix timestamp. This is not the current time, but the time when the last block was confirmed. This can be used to set times on when the contract is allowed to do certain things. For example, a contract may only allow accounts to opt in after a start date, which is set when the contract is created and stored in global storage.

=== "PyTeal"
	<!-- ===PYTEAL_GLOBAL_LATEST_TIMESTAMP=== -->
	```python
	    program = Global.latest_timestamp() >= App.globalGet(Bytes("StartDate"))
	    print(compileTeal(program, Mode.Application))
	```
	[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/application.py#L25-L27)
	<!-- ===PYTEAL_GLOBAL_LATEST_TIMESTAMP=== -->

=== "TEAL"
	<!-- ===TEAL_GLOBAL_LATEST_TIMESTAMP=== -->
	```teal
	global LatestTimestamp
	byte "StateDate"
	app_global_get
	>=
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/misc/global.teal#L3-L7)
	<!-- ===TEAL_GLOBAL_LATEST_TIMESTAMP=== -->

# Atomic transfers and transaction properties
The [TEAL opcodes](../../avm/teal/opcodes) documentation describes all transaction properties that are available within a TEAL program. These properties can be retrieved using the following contract code.


=== "PyTeal"
	<!-- ===PYTEAL_TXN_AMOUNT=== -->
	```python
	    program = Txn.amount()
	    print(compileTeal(program, Mode.Application))
	```
	[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/txn.py#L48-L50)
	<!-- ===PYTEAL_TXN_AMOUNT=== -->

=== "TEAL"
	<!-- ===TEAL_TXN_AMOUNT=== -->
	```teal
	txn Amount
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/misc/global.teal#L13-L14)
	<!-- ===TEAL_TXN_AMOUNT=== -->

In many common patterns, the smart contract will be combined with other Algorand technologies such as assets, atomic transfers, or smart signatures to build a complete application. In the case of atomic transfers, more than one transaction’s properties can be checked within the smart contract. The number of transactions can be checked using the `GroupSize` global property. If the value is greater than 1, then the call to the smart contract is grouped with more than one transaction.

=== "PyTeal"
	<!-- ===PYTEAL_TXN_GROUP_SIZE=== -->
	```python
	    program = Global.group_size() == Int(2)
	    print(compileTeal(program, Mode.Application))
	```
	[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/txn.py#L55-L57)
	<!-- ===PYTEAL_TXN_GROUP_SIZE=== -->

=== "TEAL"
	<!-- ===TEAL_TXN_GROUP_SIZE=== -->
	```teal
	global GroupSize
	int 2
	==
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/misc/global.teal#L18-L21)
	<!-- ===TEAL_TXN_GROUP_SIZE=== -->

The above contract code will be true if there are two transactions submitted at once using an atomic transfer. To access the properties of a specific transaction in the atomic group use the `gtxn` opcode.

=== "PyTeal"
	<!-- ===PYTEAL_GTXN_TYPE_ENUM=== -->
	```python
	    program = Gtxn[1].type_enum() == TxnType.Payment
	    print(compileTeal(program, Mode.Application))
	```
	[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/txn.py#L62-L64)
	<!-- ===PYTEAL_GTXN_TYPE_ENUM=== -->

=== "TEAL"
	<!-- ===TEAL_GTXN_TYPE_ENUM=== -->
	```teal
	gtxn 1 TypeEnum
	int pay
	==
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/misc/global.teal#L25-L28)
	<!-- ===TEAL_GTXN_TYPE_ENUM=== -->

In the above example, the second transaction’s type is checked, where the `int pay` references a payment transaction. See the [opcodes](../../avm/teal/opcodes) documentation for all transaction types. Note that the `gtxn` call is a zero-based index into the atomic group of transactions. The `gtxns` opcode could also have been used to retrieve the index into the atomic group from the top of the stack instead of hard coding the index. If the TEAL program fails, all transactions in the group will fail.

If any transaction in a group of transactions is a call to a smart contract, the opcodes `gtxna` and `gtxnsa` can be used to access any of the transactions array values.

=== "PyTeal"
	<!-- ===PYTEAL_GTXN_APP_ARGS=== -->
	```python
	    program = Gtxn[Txn.group_index() - Int(1)].application_args[0]
	    print(compileTeal(program, Mode.Application))
	```
	[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/txn.py#L69-L71)
	<!-- ===PYTEAL_GTXN_APP_ARGS=== -->

=== "TEAL"
	<!-- ===TEAL_GTXN_APP_ARGS=== -->
	```teal
	txn GroupIndex
	int 1
	-
	gtxnsa ApplicationArgs 0
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/misc/global.teal#L32-L36)
	<!-- ===TEAL_GTXN_APP_ARGS=== -->

# Using assets in smart contracts
Smart contract applications can work in conjunction with assets. In addition to normal asset transaction properties, such as asset amount, sender, and receiver, TEAL provides an opcode to interrogate an account’s asset balance and whether the asset is frozen. This opcode `asset_holding_get` can be used to retrieve an asset balance or check whether the asset is frozen for any account in the transaction accounts array. The asset must also be in the assets array. See [Reference arrays](index.md#reference-arrays) for more details.

=== "PyTeal"
	```python
	    asset_balance = AssetHolding.balance(Txn.sender(), Int(123456))
	    program = Seq(
	        asset_balance, If(asset_balance.hasValue(), Return(Int(1)), Return(Int(0)))
	    )
	    print(compileTeal(program, Mode.Application))
	```

=== "TEAL"
	```teal
	txn Sender
	int 123456
	asset_holding_get AssetBalance
	bnz has_balance
	
	// Reject transaction if no asset balance
	int 0
	return
	
	has_balance:
	//balance value is now on top of the stack
	```

This opcode takes two parameters. The first parameter represents the account to check. The second parameter is the Asset ID of the asset to examine. The asset must be in the assets array and the account in the accounts array for the call to be successful. See [Reference arrays](index.md#reference-arrays) for more details. 

This opcode supports getting the asset balance and the frozen state of the asset for the specific account. To get the frozen state, replace `AssetBalance` above with `AssetFrozen`. This opcode also returns two values to the top of the stack. The first is a 0 or  1, where 0 means the asset balance was not found and 1 means an asset balance was found in the accounts balance record.

It is also possible to get an Asset’s configuration information within a smart contract if the asset ID is passed with the transaction in the assets array. To read the configuration, the `asset_params_get` opcode must be used. This opcode should be supplied with one parameter, which is the index into the assets array or the actual asset ID.

=== "PyTEAL"
	```python
	    program = AssetParam.total(Int(123456))
	    print(compileTeal(program, Mode.Application))
	```

=== "TEAL"
	<!-- ===TEAL_APPL_ASSET_PARAM=== -->
	```teal
	int 123456
	asset_params_get AssetTotal
	```

This call returns two values. The first is a 0 or 1 indicating if the parameter was found and the second contains the value of the parameter. See the [opcodes](../../avm/teal/opcodes) documentation for more details on what additional parameters can be read.

# Creating an asset or contract within a group of transactions
The Algorand Protocol assigns an identifier (ID) when creating an asset (ASA) or a smart contract. These IDs are used to refer to the asset or the contract later when either is used in a transaction or a call to the smart contract. Because these IDs are assigned when the asset or the contract is created, the ID is not available until after the creation transaction is fully executed. In an atomic group, TEAL can retrieve the ID of a previous group transaction which created an asset or contract, enabling a smart contract to store the asset ID or another smart contract ID in its state for later usage. 

The ID retrieval operation can be performed by using one of two opcodes (`gaid` and `gaids`). With the `gaid` opcode, the specific transaction to read must be passed to the command. The `gaids` opcode will use the last value on the stack as the transaction index.

=== "TEAL"
	<!-- ===TEAL_APPL_CREATED_ASSET_ID=== -->
	```teal
	// Get the created id of the asset created in the first tx
	gaid 0
	// Get the created id of the asset created in the second tx
	int 1
	gaids
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/appl_asset/approval.teal#L25-L30)
	<!-- ===TEAL_APPL_CREATED_ASSET_ID=== -->

# Sharing data between contracts
In addition to reading another contract's state, data can be passed between contracts. Algorand’s AVM gives smart contracts access to scratch space that can be used to temporarily store values, while the contract is executing. TEAL allows other contracts to read this scratch space. Scratch space can only be read by other transactions that the specific transaction is grouped with atomically. Grouped transactions execute in the order they are grouped, contracts can not read scratch space for transactions that occur after the current contract transaction. 

This operation can be performed by using one of two opcodes (`gload` and `gloads`). With the `gload` opcode, the specific transaction to read and the slot number must be passed to the command. The `gloads` opcode will use the last value on the stack as the transaction index and must be passed the slot number to read.

For example, with two grouped smart contracts the following code can be used.

Store an integer in scratch space in the first transaction.

=== "TEAL"
	<!-- ===TEAL_APPL_GLOAD_T1=== -->
	```teal
	int 777
	store 10
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/misc/scratch_t1.teal#L3-L5)
	<!-- ===TEAL_APPL_GLOAD_T1=== -->

In the second transaction read the stored value.

=== "TEAL"
	<!-- ===TEAL_APPL_GLOAD_T2=== -->
	```teal
	// read the first
	// transaction's 10th
	// slot of scratch space
	gload 0 10
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/misc/scratch_t2.teal#L3-L7)
	<!-- ===TEAL_APPL_GLOAD_T2=== -->
