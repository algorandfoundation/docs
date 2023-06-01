Title: Creation 


# Creating the smart contract
Before creating a smart contract, the code for the `ApprovalProgram` and the `ClearStateProgram` program should be written. The SDKs and the `goal` CLI tool can be used to create a smart contract application. To create the application with `goal` use a command similar to the following.

```bash
$ goal app create --creator [address]  --approval-prog [approval_program.teal] --clear-prog [clear_state_program.teal] --global-byteslices [number-of-global-byteslices] --global-ints [number-of-global-ints] --local-byteslices [number-of-local-byteslices] --local-ints [number-local-ints] --extra-pages [number of extra 2KB pages]
```

!!!note
	See [Creating the smart contract](../frontend/apps.md#create) for details on using the SDKs to deploy a smart contract.

The creator is the account that is creating the application and this transaction is signed by this account. The approval program and the clear state program should also be provided. The number of global and local byte slices (byte-array value) and integers also needs to be specified. These represent the absolute on-chain amount of space that the smart contract will use. Once set, these values can never be changed. The key is limited to 64 bytes. The key plus the value is limited to 128 bytes total. When the smart contract is created the network will return a unique ApplicationID. This ID can then be used to make `ApplicationCall` transactions to the smart contract. The smart contract will also have a unique Algorand address that is generated from this ID. This address allows the contract to function as an escrow account.  

When creating a smart contract, there is a limit of 64 key-value pairs that can be used by the contract for global storage and 16 key-value pairs that can be used for local storage. When creating the smart contract the amount of storage can never be changed once the contract is created. Additionally, the minimum balance is raised for any account that participates in the contract. See [Minimum Balance Requirement for Smart Contracts](index.md#minimum-balance-requirement-for-a-smart-contract) described below for more detail.

Smart contracts are limited to 2KB total for the compiled approval and clear programs. This size can be increased up to 3 additional 2KB pages, which would result in an 8KB limit for both programs. Note the size increases will also increase the minimum balance requirement for creating the application. To request additional pages, the setting (`extra-pages`) is available when creating the smart contract using `goal`. These extra pages can also be requested using the SDKs. This setting allows setting up to 3 additional 2KB pages.

# Opt into the smart contract
Before any account, including the creator of the smart contract, can begin to make Application Transaction calls that use local state, it must first opt into the smart contract. This prevents accounts from being spammed with smart contracts. To opt in, an `ApplicationCall` transaction of type `OptIn` needs to be signed and submitted by the account desiring to opt into the smart contract. This can be done with the `goal` CLI or the SDKs.

```bash
$ goal app optin  --app-id [ID-of-Contract] --from [ADDRESS]
```

!!!note
	See [Opt-in](../frontend/apps.md#opt-in) for details on using the SDKs to opt into a smart contract.

When this transaction is submitted, the `ApprovalProgram` of the smart contract is called and if the call succeeds the account will be opted into the smart contract. The simplest program to handle this call would just put 1 on the stack and return. 

=== "PyTeal"
	<!-- ===PYTEAL_APPL_OPTIN=== -->
	```python
	    # this would reject _ANY_ transaction that isn't an opt-in
	    # and approve _ANY_ transaction that is an opt-in
	    program = If(OnComplete.OptIn == Txn.on_completion(), Approve(), Reject())
	    print(compileTeal(program, Mode.Application))
	```
	[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/txn.py#L108-L112)
	<!-- ===PYTEAL_APPL_OPTIN=== -->

=== "TEAL"
	<!-- ===TEAL_APPL_OPTIN=== -->
	```teal
	txn OnCompletion
	int OptIn
	==
	bz not_optin
	
	// Allow OptIn
	int 1
	return
	
	not_optin:
	// additional checks...
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/opting_in/approval.teal#L13-L24)
	<!-- ===TEAL_APPL_OPTIN=== -->

Other contracts may have much more complex opt in logic. TEAL also provides an opcode to check whether an account has already opted into the contract.

=== "PyTeal"

	```python
	    program = App.optedIn(Txn.sender(), Txn.application_id())
	    print(compileTeal(program, Mode.Application))
	```

=== "TEAL"
	<!-- ===TEAL_APPL_CHECK_OPTEDIN=== -->
	```teal
	txn Sender
	txn ApplicationID
	app_opted_in
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/opting_in/approval.teal#L6-L9)
	<!-- ===TEAL_APPL_CHECK_OPTEDIN=== -->

In the above example, `txn Sender` is the address of the transaction sender. The address can be any account in the accounts array. The `txn ApplicationID` refers to the current application ID, but technically any application ID could be used as long as its ID is in the applications array. See [Reference arrays](index.md#reference-arrays) for more details.

!!! info
    Applications that only use global state do not require accounts to opt in.

# Passing arguments to smart contracts
Arguments can be passed to any of the supported application transaction calls, including create. The number and type can also be different for any subsequent calls to the smart contract. The `goal` CLI supports passing strings, ints, base64 encoded data, and addresses as parameters. To pass a parameter supply the `--app-arg` option to the call and supply the value according to the format shown below.

Argument Type | Example
------------ | ------------- 
String | `goal app call --app-arg "str:mystring".....` 
Integer | `goal app create --app-arg "int:5".....` 
Address | `goal app call --app-arg "addr:address-string".....`
Base64 | `goal app call --app-arg "b64:A==".....`

!!!note
	See [Call with arguments](../frontend/apps.md#call-with-arguments), for more information on passing parmeters with SDKs.

These parameters are loaded into the arguments array. TEAL opcodes are available to get the values within the array. The primary argument opcode is the `ApplicationArgs` opcode and can be used as shown below.

=== "PyTeal"
	<!-- ===PYTEAL_TXN_APP_ARGS=== -->
	```python
	    program = Txn.application_args[1] == Bytes("claim")
	    print(compileTeal(program, Mode.Application))
	```
	[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/txn.py#L27-L29)
	<!-- ===PYTEAL_TXN_APP_ARGS=== -->

=== "TEAL"
	<!-- ===TEAL_TXN_APP_ARGS=== -->
	```teal
	txna ApplicationArgs 1
	byte "claim"
	==
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/appargs/approval.teal#L16-L19)
	<!-- ===TEAL_TXN_APP_ARGS=== -->

This call gets the second passed in argument and compares it to the string "claim".

A global variable is also available to check the size of the transaction argument array. This size can be checked with the following contract code.


=== "PyTeal"
	<!-- ===PYTEAL_TXN_NUM_APP_ARGS=== -->
	```python
	    program = Txn.application_args.length() == Int(4)
	    print(compileTeal(program, Mode.Application))
	```
	[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/txn.py#L34-L36)
	<!-- ===PYTEAL_TXN_NUM_APP_ARGS=== -->

=== "TEAL"
	<!-- ===TEAL_TXN_NUM_APP_ARGS=== -->
	```teal
	txn NumAppArgs
	int 4
	==
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/appargs/approval.teal#L3-L6)
	<!-- ===TEAL_TXN_NUM_APP_ARGS=== -->

The above contract code will push a 0 on the top of the stack if the number of parameters in this specific transaction is anything other than 4, else it will push a 1 on the top of the stack. Internally all transaction parameters are stored as byte slices (byte-array value). Integers can be converted using the `btoi` opcode.

=== "PyTeal"
	<!-- ===PYTEAL_TXN_APP_ARG_TO_INT=== -->
	```python
	    program = Btoi(Txn.application_args[0])
	    print(compileTeal(program, Mode.Application))
	```
	[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/txn.py#L41-L43)
	<!-- ===PYTEAL_TXN_APP_ARG_TO_INT=== -->

=== "TEAL"
	<!-- ===TEAL_TXN_APP_ARG_TO_INT=== -->
	```teal
	txna ApplicationArgs 0
	btoi
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/appargs/approval.teal#L10-L12)
	<!-- ===TEAL_TXN_APP_ARG_TO_INT=== -->


!!! info
    Argument passing for smart contracts is very different from passing arguments to smart signatures. 

The total size of all parameters is limited to 2KB in size.


# Call the smart contract
Any account can make a call to the smart contract. These calls will be in the form of `ApplicationCall` transactions that can be submitted with `goal` or the SDKs. Depending on the individual type of transaction as described in [The Lifecycle of a Smart Contract](index.md#the-lifecycle-of-a-smart-contract), either the `ApprovalProgram` or the `ClearStateProgram` will be called. Generally, individual calls will supply application arguments. See [Passing Arguments to a Smart Contract](#passing-arguments-to-smart-contracts) for details on passing arguments.

```bash
$ goal app call --app-id 1 --app-arg "str:myparam"  --from [ADDRESS]
```

!!!note
    See [Call(NoOp)](../frontend/apps.md#call-noop) for details on using the SDKs to call a smart contract. If using [ABI](/docs/get-details/dapps/smart-contracts/ABI/) compliant contracts, use the [AtomicTransactionComposer](../../../../atc) to interact with the smart contract.


The call must specify the intended contract using the `--app-id` option. Additionally, the `--from` option specifies the sender’s address. 


=== "PyTeal"
	<!-- ===PYTEAL_APPL_CALL=== -->
	```python
	    # this would approve _ANY_ transaction that has its
	    # first app arg set to the byte string "myparam"
	    # and reject all others
	    program = If(Bytes("myparm") == Txn.application_args[0], Approve(), Reject())
	    print(compileTeal(program, Mode.Application))
	```
	[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/txn.py#L76-L81)
	<!-- ===PYTEAL_APPL_CALL=== -->

=== "Teal"
	<!-- ===TEAL_APPL_CALL=== -->
	```teal
	byte "myparam"
	txna ApplicationArgs 0
	==
	bz not_myparam
	// handle my_param
	
	not_myparam:
	// handle not_myparam
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/calling/approval.teal#L6-L14)
	<!-- ===TEAL_APPL_CALL=== -->

# Update smart contract
A smart contract’s programs can be updated at any time. This is done by an `ApplicationCall` transaction type of `UpdateApplication`. This operation can be done with `goal` or the SDKs and requires passing the new programs and specifying the application ID.

```bash
goal app update --app-id=[APPID] --from [ADDRESS]  --approval-prog [new_approval_program.teal]   --clear-prog [new_clear_state_program.teal]
```

!!!note
    See [Update](../frontend/apps.md#update) for details on using the SDKs to update a smart contract.

The one caveat to this operation is that global or local state requirements for the smart contract can never be updated. Updating a smart contract's programs does not affect any values currently in state.

As stated earlier, anyone can update the program. If this is not desired and you want only the original creator to be able to update the programs, code must be added to your `ApprovalProgram` to handle this situation. This can be done by comparing the global `CreatorAddress` to the sender address.

=== "PyTeal"
	<!-- ===PYTEAL_APPL_UPDATE=== -->
	```python
	    program = Assert(
	        Txn.on_completion() == OnComplete.UpdateApplication,
	        Global.creator_address() == Txn.sender(),
	    )
	    print(compileTeal(program, Mode.Application))
	```
	[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/txn.py#L86-L91)
	<!-- ===PYTEAL_APPL_UPDATE=== -->

=== "TEAL"
	<!-- ===TEAL_APPL_UPDATE=== -->
	```teal
	byte "update"
	txna ApplicationArgs 0
	==
	bz not_update
	
	// Only Creator may update
	global CreatorAddress
	txn Sender
	==
	return
	
	not_update:
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/calling/approval.teal#L17-L29)
	<!-- ===TEAL_APPL_UPDATE=== -->

Or alternatively, the contract code can always return a 0 when an `UpdateApplication` application call is made to prevent anyone from ever updating the application code.

=== "PyTeal"
	<!-- ===PYTEAL_APPL_UPDATE_REJECT=== -->
	```python
	    program = If(
	        OnComplete.UpdateApplication == Txn.on_completion(),
	        Reject(),
	        # placeholder, update with actual logic
	        Approve(),
	    )
	    print(compileTeal(program, Mode.Application))
	```
	[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/txn.py#L96-L103)
	<!-- ===PYTEAL_APPL_UPDATE_REJECT=== -->

=== "TEAL"
	<!-- ===TEAL_APPL_UPDATE_REJECT=== -->
	```teal
	txn OnCompletion
	int UpdateApplication
	==
	bz not_update
	
	// Reject Update
	int 0
	return
	
	not_update:
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/calling/approval.teal#L32-L42)
	<!-- ===TEAL_APPL_UPDATE_REJECT=== -->

# Delete smart contract
To delete a smart contract, an `ApplicationCall` transaction of type `DeleteApplication` must be submitted to the blockchain. The `ApprovalProgram` handles this transaction type and if the call returns true, the application will be deleted. This can be done using `goal` or the SDKs. 

```bash
$ goal app delete --app-id=[APPID] --from [ADDRESS]
```

!!!note
    See [Delete](../frontend/apps.md#delete) for details on using the SDKs to delete a smart contract.

When making this call the `--app-id` and the `--from` options are required. Anyone can delete a smart contract. If this is not desired, logic in the program must reject the call. Using a method described in [Update Smart Contract](#update-smart-contract) must be supplied. 

# Boilerplate smart contract

As a way of getting started writing smart contracts, the following boilerplate template is supplied. The code provides labels or handling different `ApplicationCall` transactions and also prevents updating and deleting the smart contract.


=== "PyTEAL"
	<!-- ===PYTEAL_BOILERPLATE=== -->
	```python
	    # Handle each possible OnCompletion type. We don't have to worry about
	    # handling ClearState, because the ClearStateProgram will execute in that
	    # case, not the ApprovalProgram.
	    def approval_program():
	        handle_noop = Seq([Return(Int(1))])
	
	        handle_optin = Seq([Return(Int(1))])
	
	        handle_closeout = Seq([Return(Int(1))])
	
	        handle_updateapp = Err()
	
	        handle_deleteapp = Err()
	
	        program = Cond(
	            [Txn.on_completion() == OnComplete.NoOp, handle_noop],
	            [Txn.on_completion() == OnComplete.OptIn, handle_optin],
	            [Txn.on_completion() == OnComplete.CloseOut, handle_closeout],
	            [Txn.on_completion() == OnComplete.UpdateApplication, handle_updateapp],
	            [Txn.on_completion() == OnComplete.DeleteApplication, handle_deleteapp],
	        )
	        return program
	
	    with open("boilerplate_approval_pyteal.teal", "w") as f:
	        compiled = compileTeal(approval_program(), Mode.Application, version=5)
	        f.write(compiled)
	```
	[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/application.py#L32-L58)
	<!-- ===PYTEAL_BOILERPLATE=== -->

=== "TEAL"
	<!-- ===TEAL_BOILERPLATE=== -->
	```teal
	#pragma version 8
	
	// Handle each possible OnCompletion type. We don't have to worry about
	// handling ClearState, because the ClearStateProgram will execute in that
	// case, not the ApprovalProgram.
	
	txn OnCompletion
	int NoOp
	==
	bnz handle_noop
	
	txn OnCompletion
	int OptIn
	==
	bnz handle_optin
	
	txn OnCompletion
	int CloseOut
	==
	bnz handle_closeout
	
	txn OnCompletion
	int UpdateApplication
	==
	bnz handle_updateapp
	
	txn OnCompletion
	int DeleteApplication
	==
	bnz handle_deleteapp
	
	// Unexpected OnCompletion value. Should be unreachable.
	err
	
	handle_noop:
	// Handle NoOp
	int 1
	return
	
	handle_optin:
	// Handle OptIn
	int 1
	return
	
	handle_closeout:
	// Handle CloseOut
	int 1
	return
	
	// By default, disallow updating or deleting the app. Add custom authorization
	// logic below to allow updating or deletion in certain circumstances.
	handle_updateapp:
	handle_deleteapp:
	err
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/misc/boilerplate.teal#L1-L55)
	<!-- ===TEAL_BOILERPLATE=== -->
