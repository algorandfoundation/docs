title: Overview

# Introduction

Algorand smart contracts are pieces of logic that reside on the Algorand blockchain and are remotely callable. These contracts are primarily responsible for implementing the logic associated with a distributed application. Smart contracts are referred to as stateful smart contracts or applications in the Algorand documentation. Smart contracts can generate asset and payment transactions allowing them to function as Escrow accounts on the Algorand blockchain. Smart contracts can also store values on the blockchain. This storage can be global, local, or box storage. Local storage refers to storing values in an accounts balance record if that account participates in the contract. Global storage is data that is specifically stored on the blockchain for the contract globally. Box storage is also global and allows contracts to use larger segments of storage. Like smart signatures, smart contracts are written in Python using PyTeal or TEAL and can be deployed to the blockchain using either the `goal` command-line tool or the SDKs. The recommended approach for writing smart contracts is to use the Python SDK with the PyTeal library.  

See the [*PyTeal Documentation*](/docs/get-details/dapps/writing-contracts/pyteal) for information on building smart contracts in Python.

See the [*TEAL Reference Guide*](../../avm/teal/specification.md) to understand how to write TEAL and the [*TEAL Opcodes*](../../avm/teal/opcodes.md) documentation that describes the opcodes available. This guide assumes that the reader is familiar with [TEAL](../../avm/teal/index.md).

!!! important "A note about PyTeal"
    Where possible, TEAL code snippets are accompanied by their counterparts in PyTeal. Here are a few things to be aware of when comparing across these two languages:
     
    - Each PyTeal code snippet ends with the action to compile the program and write it to a file so the user can view the underlying TEAL code.
    - Sometimes the compiled version of a PyTeal code snippet will differ slightly from the TEAL version. However, the resulting function should be equivalent for the particular area of focus in the documentation. In larger more complex programs, this may not always be the case.
    - When a TEAL code snippet includes comments as placeholders for code, the PyTeal example will often use a placeholder of `Seq([Return(Int(1))])` with a comment describing this as a placeholder. This allows the user to compile the program for learning purposes. However, returning 1 is a very permissive action and should be carefully updated when used in a production application.


# The lifecycle of a smart contract
Smart contracts are implemented using two programs:

* The `ApprovalProgram` is responsible for processing all application calls to the contract, with the exception of the clear call (described in the next bullet). This program is responsible for implementing most of the logic of an application. Like smart signatures, this program will succeed only if one nonzero value is left on the stack upon program completion or the `return` opcode is called with a positive value on the top of the stack.
* The `ClearStateProgram` is used to handle accounts using the clear call to remove the smart contract from their balance record. This program will pass or fail the same way the `ApprovalProgram` does. 

In either program, if a global, box, or local state variable is modified and the program fails, the state changes will not be applied. 

Having two programs allows an account to clear the contract from its state, whether the logic passes or not. When the clear call is made to the contract, whether the logic passes or fails, the contract will be removed from the account's balance record. Note the similarity to the CloseOut transaction call which can fail to remove the contract from the account, which is described below.

Calls to smart contracts are implemented using `ApplicationCall` transactions. These transaction types are as follows:

* NoOp - Generic application calls to execute the `ApprovalProgram`.
* OptIn - Accounts use this transaction to begin participating in a smart contract.  Participation enables local storage usage.
* DeleteApplication - Transaction to delete the application.
* UpdateApplication - Transaction to update TEAL Programs for a contract.
* CloseOut - Accounts use this transaction to close out their participation in the contract. This call can fail based on the TEAL logic, preventing the account from removing the contract from its balance record.
* ClearState - Similar to CloseOut, but the transaction will always clear a contract from the account’s balance record whether the program succeeds or fails.

The `ClearStateProgram` handles the `ClearState` transaction and the `ApprovalProgram` handles all other `ApplicationCall` transactions. These transaction types can be created with either `goal` or the SDKs. In the following sections, details on the individual capabilities of a smart contract will be explained.

<center>![Smart Contract](/docs/imgs/stateful-1.png)</center>
<center>*Application Transaction Types*</center>

# Reference arrays
A set of arrays can be passed with any application transaction, which instructs the protocol to load additional data for use in the contract. These arrays are the *arguments* array, the *applications* array, the *assets* array, the *accounts* array, and the *boxes* array. These arrays are used to restrict how much of the ledger can be accessed in one specific call to the smart contract. This restriction is in place to maintaing blockchain performance. These arrays can be changed per application transaction (even within an atomic group).

The arguments array is used to pass standard arguments to the contract. The arguments array is limited to 16 arguments with a 2KB total size limit. See [Passing Arguments To Smart Contracts](create.md#passing-arguments-to-smart-contracts) for more details on arguments.

The other arrays are used to load data from the blockchain:

* The application array is used to pass other smart contract IDs that can be used to read state for those specific contracts. 
* The assets array is used to pass a list of asset IDs that can be used to retrieve configuration and asset balance information. 
* The accounts array allows additional accounts to be passed to the contract for balance information and local state. Note that to access an account's asset balance, both the account and the asset ID must be specified in their respective arrays. Similarly, to access an account's local state for a specific application, both the account and the smart contract ID must be specified in their respective arrays.
* The boxes array defines which boxes can be manipulated in a particular call to the smart contract. The box array is an array of pairs: the first element of each pair is an integer specifying the index into the application array, and the second element is the key name of the box to be accessed.   

These four arrays (*applications*, *assets*, *accounts*, and *boxes*) are limited to eight total values combined, and of those, the accounts array can have no more than four values. The values passed within these arrays can change per Application Transaction.  The opcodes that make use of these arrays take an integer parameter as an index into these arrays. 

The accounts and applications arrays contain the transaction sender and current application ID in the 0th position of the respective array. This shifts the contents of these two arrays by one slot. The opcodes that use an index into these arrays also allow passing the actual value. For example, an address can be specified for an opcode that uses the accounts array. IDs can be specified for contracts and assets for an opcode that uses the applications or assets arrays, respectively. These opcodes will fail if the specified value does not exist in the corresponding array. The use of each of these arrays is detailed throughout this guide.

The following examples illustrate populating these arrays before calling a smart contract when using the Atomic Transaction Composer(ATC).

Boxes function similar to the other arrays but differ is significant ways which are explained in detail in the [Boxes section of the documentation](state.md#box-details).

<center>![Smart Contract](/docs/imgs/stateful-2.png)</center>
<center>*Reference Arrays*</center>


# Application Account
Since September 2021 all deployed smart contracts are given their own application account with an associated Algorand public address. These accounts are used by issuing [inner transactions](/docs/get-details/dapps/smart-contracts/apps/innertx/) from within the smart contract.

The public address is devised by taking the application ID as an 8-byte big-endian array, prefixing it with `appID`, and then encoding it using the standard encoding method for Algorand addresses. The Algorand SDKs provide helpful utilities to do this for you given just the application ID.

=== "Python"
<!-- ===PYSDK_CODEC_APPLICATION_ACCOUNT=== -->
	```python
	app_id = 123
	app_addr = logic.get_application_address(app_id)
	
	print(f"Application ID:   {app_id}")
	print(f"Application Addr: {app_addr}")
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/codec.py#L16-L21)
<!-- ===PYSDK_CODEC_APPLICATION_ACCOUNT=== -->

=== "JavaScript"
<!-- ===JSSDK_CODEC_APPLICATION_ACCOUNT=== -->
	```javascript
	const appId = 123;
	const appAddr = algosdk.getApplicationAddress(appId);
	
	console.log(`Application ID:      ${  appId}`);
	console.log(`Application Address: ${  appAddr}`);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/codec.ts#L23-L28)
<!-- ===JSSDK_CODEC_APPLICATION_ACCOUNT=== -->

=== "Go"
<!-- ===GOSDK_CODEC_APPLICATION_ACCOUNT=== -->
	```go
	var appId uint64 = 123
	var appAddr types.Address = crypto.GetApplicationAddress(appId)
	
	fmt.Println("Application ID:      ", appId)
	fmt.Println("Application Address: ", appAddr)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/codec/main.go#L76-L81)
<!-- ===GOSDK_CODEC_APPLICATION_ACCOUNT=== -->


# Minimum balance requirement for a smart contract
When creating or opting into a smart contract your minimum balance will be raised. The amount at which it is raised will depend on the amount of on-chain storage that is used. This minimum balance requirement is associated with the account that creates or opts into the smart contract.

Calculation for increase in min-balance during creation or opt-in is as follows:

=== "Application Create"

    ```
    100,000*(1+ExtraProgramPages) + (25,000+3,500)*schema.NumUint + (25,000+25,000)*schema.NumByteSlice
    ```

    In words:

    - 100,000 microAlgo base fee for each page requested. 
    - 25,000 + 3,500 = 28,500 for each Uint in the *Global State* schema
    - 25,000 + 25,000 = 50,000 for each byte-slice in the *Global State* schema

=== "Application Opt-In"

    ```
    100,000 + (25,000+3,500)*schema.NumUint + (25,000+25,000)*schema.NumByteSlice
    ```

    In words:

    - 100,000 microAlgo base fee of opt-in
    - 25,000 + 3,500 = 28,500 for each Uint in the *Local State* schema
    - 25,000 + 25,000 = 50,000 for each byte-slice in the *Local State* schema


!!! note 
    Global storage is actually stored in the creator account, so that account is responsible for the global storage minimum balance.  When an account opts-in, it is responsible for the minimum balance of local storage. 
## Example
Given a smart contract with 1 *global* key-value-pair of type byteslice and 1 *local* storage key-value pair of type integer and no Extra Pages. 


=== "Creator"

    The creator of the Application would have its minimum balance raised by *150,000 microAlgos*.
    ```
    100,000 + 50,000  = 150,000
    ```

=== "Account"

    The account opting into the Application would have minimum balance raised *128,500 microAlgos*.

    ```
    100,000 + 28,500 = 128,500
    ```

# Minimum Balance Requirement For Boxes
Boxes are created by a smart contract and raise the minimum balance requirement (MBR) that must be in the contract's ledger balance. This means that a contract that intends to use boxes, must be funded beforehand.

When a box with name `n` and size `s` is created, the MBR is raised by `2500 + 400 * (len(n)+s)` microAlgos. When the box is destroyed, the minimum balance requirement is decremented by the same amount.

Notice that the key (name) is included as part of the MBR calculation. 

For example, if a box is created with the name “BoxA” (a 4 byte long key) and with a size of 1024 bytes, the MBR for the app account increases by 413,700 microAlgos:

    (2500 per box) + (400 * (box size + key size))
    (2500) + (400 * (1024+4)) = 413,700 microAlgos 