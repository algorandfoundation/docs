title: Overview

Algorand Smart Contracts can be written in either a stateless or stateful manner. To be stateful means that some amount of storage on the chain is used to store values. This storage can be either global or local. Local storage refers to storing values in an accounts balance record if that account participates in the contract. Global storage is data that is specifically stored on the blockchain for the contract globally. Like stateless smart contracts, stateful contracts are written in TEAL and can be deployed to the blockchain using either the `goal` command-line tool or the SDKs. Stateless smart contracts’ primary purpose is to approve or reject spending transactions. Stateful contracts do not approve spending transactions but provide logic that allows the state (globally or locally) of the contract to be manipulated. Most often, these contracts will be paired with other Algorand capabilities or each other using atomic transfers to form a complete application.

See the [*TEAL Reference Guide*](../../../reference/teal/specification.md) to understand how to write TEAL and the [*TEAL Opcodes*](../../../reference/teal/opcodes.md) documentation that describes the opcodes available. This guide assumes that the reader is familiar with [TEAL](../teal/index.md).

!!! important "A note about PyTeal"
    Where possible, TEAL code snippets are accompanied by their counterparts in PyTeal. Here are a few things to be aware of when comparing across these two languages:
     
    - Each PyTeal code snippet ends with the action to compile the program and write it to a file so the user can view the underlying TEAL code.
    - Sometimes the compiled version of a PyTeal code snippet will differ slightly from the TEAL version. However, the resulting function should be equivalent for the particular area of focus in the documentation. In larger more complex programs, this may not always be the case.
    - When a TEAL code snippet includes comments as placeholders for code, the PyTeal example will often use a placeholder of `Seq([Return(Int(1))])` with a comment describing this as a placeholder. This allows the user to compile the program for learning purposes. However, returning 1 is a very permissive action and should be carefully updated when used in a real application.

# The Lifecycle of a Stateful Smart Contract
Stateful smart contracts are implemented using two TEAL programs:

* The `ApprovalProgram` is responsible for processing all application calls to the contract, with the exception of the clear call (described in the next bullet). This program is responsible for implementing most of the logic of an application. Like stateless contracts, this program will succeed only if one nonzero value is left on the stack upon program completion or the `return` opcode is called with a positive value on the top of the stack.
* The `ClearStateProgram` is used to handle accounts using the clear call to remove the smart contract from their balance record. This program will pass or fail the same way the `ApprovalProgram` does. 

In either program, if a global or local state variable is modified and the program fails, the state changes will not be applied. 

Having two programs allows an account to clear the contract from its state, whether the logic passes or not. When the clear call is made to the contract, whether the logic passes or fails, the contract will still be removed from the balance record for the specified account. Note the similarity to the CloseOut transaction call which can fail to remove the contract from the account, which is described below.

Calls to stateful smart contracts are implemented using `ApplicationCall` transactions. These transactions types are as follows:

* NoOp - Generic application calls to execute the `ApprovalProgram`
* OptIn - Accounts use this transaction to opt into the smart contract to participate (local storage usage).
* DeleteApplication - Transaction to delete the application.
* UpdateApplication - Transaction to update TEAL Programs for a contract.
* CloseOut - Accounts use this transaction to close out their participation in the contract. This call can fail based on the TEAL logic, preventing the account from removing the contract from its balance record.
* ClearState - Similar to CloseOut, but the transaction will always clear a contract from the account’s balance record whether the program succeeds or fails.

The `ClearStateProgram` handles the `ClearState` transaction and the `ApprovalProgram` handles all other `ApplicationCall` transactions. These transaction types can be created with either `goal` or the SDKs. The overall architecture of a stateful TEAL program is shown below. In the following sections, details on the individual capabilities of a stateful smart contract will be explained.

<center>![Stateful Smart Contract](../../../imgs/stateful-1.png)</center>
<center>*Stateful Smart Contract*</center>

The `goal` calls shown above in the orange boxes represent all the specific calls that can be made against a stateful smart contract and are described later in this document. These calls are also available in the SDKs. The teal-colored boxes represent the two required TEAL programs, the blue boxes are the state variables (local and global), and the yellow boxes represent the TEAL opcodes used to modify state. Modifying state is detailed in the next section.

# Stateful Contract Arrays
A set of arrays can be passed with any application transaction, which instructs the protocol to load additional data for use in the contract. These arrays are the *applications* array, the *accounts* array, the *assets* array, and the *arguments* array. The application array is used to pass other stateful smart contract IDs that can be used to read state for the specific contracts. The accounts array allows additional accounts to be passed to the contract for balance information and local storage. The assets array is used to pass a list of asset IDs that can be used to retrieve configuration and asset balance information. 

The arguments array is used to pass standard arguments to the contract. The arguments array is limited to 16 arguments with a size limit of 2KB for the total of arguments. See [Passing Arguments To Stateful Smart Contracts](#passing-arguments-to-stateful-smart-contracts) for more details on arguments.

The other three arrays are limited to 8 total values combined, and of those, the accounts array can have no more than four values. The values passed within these arrays can change per Application Transaction. Many opcodes that make use of these arrays take an integer parameter as an index into these arrays. The accounts and applications arrays contain the transaction sender and current application ID in the 0th position of the respective array. This shifts the contents of these two arrays by one slot. Most of the opcodes that use an index into these arrays also allow passing the actual value. For example, an address can be specified for an opcode that uses the accounts array. IDs can be specified for contracts and assets for an opcode that uses the applications or assets arrays, respectively. These opcodes will fail if the specified value does not exist in the corresponding array. The use of each of these arrays is detailed throughout this guide.

<center>![Stateful Smart Contract](../../../imgs/stateful-2.png)</center>
<center>*Stateful Smart Contract Arrays*</center>

# Modifying State in Smart Contract
Stateful smart contracts can create, update, and delete values in global or local state. The number of values that can be written is limited based on how the contract was first created. See [Creating the Smart Contract](#creating-the-smart-contract) for details on configuring the initial global and local storage. State is represented with key-value pairs. The key is limited to 64 bytes. The key plus the value is limited to 128 bytes total. Using smaller keys to have more storage available for the value is possible. The keys are stored as byte slices (byte-array value) and the values are stored as either byte slices (byte-array value) or uint64s. The TEAL language provides several opcodes for facilitating reading and writing to state.

## Reading Local State from other Accounts
Local storage values are stored in the account's balance record. Any account that sends a transaction to the smart contract can have its local storage modified by the smart contract as long as the account has opted into the smart contract. In addition, any call to the smart contract can also reference up to four additional accounts which can also have their local storage manipulated for the current smart contract as long as the account has opted into the contract. These five accounts can also have their storage values for any smart contract on Algorand read by specifying the application id of the smart contract, if the additional contract is in the applications array for the transaction. This is a read-only operation and does not allow one smart contract to modify the local state of another smart contract. The additionally referenced accounts can be changed per smart contract call (transaction). The process for reading local state from another account is described in the following sections.

## Reading Global State from other Smart Contracts
Global storage for the current contract can also be modified by the smart contract code. In addition, the global storage of any contract in the applications array can be read. This is a read-only operation. The global state can not be changed for other smart contracts. The external smart contracts can be changed per smart contract call (transaction). The process for reading global state from another smart contract is described in the following sections.


## Write to State
To write to either local or global state, the opcodes `app_global_put` and `app_local_put` should be used. These calls are similar but with local storage, you provide an additional account parameter. This determines what account should have its local storage modified. In addition to the sender of the transaction, any call to the smart contract can reference up to four additional accounts. Below is an example of doing a global write with TEAL.

```text tab="TEAL"
byte "Mykey"
int 50
app_global_put
```

```python tab="PyTeal"
program = App.globalPut(Bytes("Mykey"), Int(50))

print(compileTeal(program, Mode.Application))
```

To store a value in local storage, the following TEAL can be used.

```text tab="TEAL"
int 0
byte "MyLocalKey"
int 50
app_local_put
```

```python tab="PyTeal"
program = App.localPut(Int(0), Bytes("MyLocalKey"), Int(50))

print(compileTeal(program, Mode.Application))
```

In this example, the `int 0` represents the sender of the transaction. This is a reference into the accounts array that is passed with the transaction. With `goal` you can pass additional accounts using the `--app-account` option. The address can be also be specified instead of the index. If using an address, it still must exist in the accounts array.

```
$ goal app call --app-account account1 --app-account account2
```

To store a value into account2, the TEAL would be as follows.

```text tab="TEAL"
int 2
byte “MyLocalKey”
int 50
app_local_put
```

```python tab="PyTeal"
program = App.localPut(Int(2), Bytes("MyLocalKey"), Int(50))

print(compileTeal(program, Mode.Application))
```

Where 0 is the sender, 1 is the first additional account passed in and 2 is the second additional account passed with the application call.

!!! info
    Local storage writes are only allowed if the account has opted into the smart contract.

## Read From State
TEAL provides calls to read global and local state values for the current smart contract.  To read from local or global state TEAL provides the `app_local_get`, `app_global_get`, `app_local_get_ex` , and `app_global_get_ex` opcodes. The following TEAL code reads a value from global state for the current smart contract.

```text tab="TEAL"
byte "MyGlobalKey"
app_global_get
```

```python tab="PyTeal"
program = App.globalGet(Bytes("MyGlobalKey"))

print(compileTeal(program, Mode.Application))
```

The following TEAL code reads the local state of the sender account for the specific call to the current smart contract.

```text tab="TEAL"
int 0
byte "MyLocalKey"
app_local_get
```

```python tab="PyTeal"
program = App.localGet(Int(0), Bytes("MyLocalKey"))

print(compileTeal(program, Mode.Application))
```

In this example, the `int 0` represents the sender of the transaction. This is a reference into the accounts array that is passed with the transaction. The address can be specified instead of the index as long as the account is in the accounts array. With `goal` you can pass additional accounts using the `--app-account` option. 

`int 0` represents the sender, 1 is the first additional account passed in and 2 is the second additional account passed with the application call. 

The `_ex` opcodes return two values to the stack. The first value is a 0 or a 1 indicating the value was returned successfully or not, and the second value on the stack contains the actual value. These calls allow local and global states to be read from other accounts and applications (stateful smart contracts) as long as the account and the contract are in the accounts and applications arrays. To read a local storage value with the `app_local_get_ex` opcode the following TEAL should be used.

```text tab="TEAL"
int 0 // sender
txn ApplicationID // current smart contract
byte "MyAmountGiven"
app_local_get_ex
```

```python tab="PyTeal"
program = App.localGetEx(Int(0), Txn.application_id(), Bytes("MyAmountGiven"))

print(compileTeal(program, Mode.Application))
```

!!! note
    The PyTeal code snippet preemptively stores the return values from `localGetEx` in scratch space for later reference. 

The `int 0` is the index into the accounts array. The actual address could also be specified as long as the account is in the accounts array. The `txn ApplicationID` line refers to the current application, but could be any application that exists on Algorand as long as the contract's ID is in the applications array. Instead of specifying the application ID, the index into the application array can be used as well. The top value on the stack will either return 0 or 1 depending on if the variable was found.  Most likely branching logic will be used after a call to the `_ex` opcode. The following example illustrates this concept.

```text tab="TEAL"
int 0 // sender
txn ApplicationID
byte "MyAmountGiven"
app_local_get_ex
bz new-giver

// logic to deal with an existing giver
// stored value is on the top of the stack
// return

new-giver:

// logic to deal with a new giver
```

```python tab="PyTeal"
get_amount_given = App.localGetEx(Int(0), Txn.application_id(), Bytes("MyAmountGiven"))

# Change these to appropriate logic for new and previous givers.
new_giver_logic = Seq([
    Return(Int(1))
])

previous_giver_logic = Seq([
    Return(Int(1))
]) 

program = Seq([
            get_amount_given, 
            If(get_amount_given.hasValue(), 
                previous_giver_logic, 
                new_giver_logic),
    ])

print(compileTeal(program, Mode.Application))
```

The `app_global_get_ex` is used to read not only the global state of the current contract but any contract that is in the applications array. To access these foreign apps, they must be passed in with the application using the `--foreign-app` option. 

```bash
$ goal app call --foreign-app APP1ID --foreign-app APP2ID
```

To read from the global state with the `app_global_get_ex` opcode, use the following TEAL.

```text tab="TEAL"
int 0
byte "MyGlobalKey"
app_global_get_ex
bnz increment_existing //found value
```

```python tab="PyTeal"
get_global_key = App.globalGetEx(Int(0), Bytes("MyGlobalKey"))

# Update with appropriate logic for use case
increment_existing = Seq([
    Return(Int(1))
])

program = Seq([
        get_global_key, 
        If(get_global_key.hasValue(), 
            increment_existing, 
            Return(Int(1))),
    ])

print(compileTeal(program, Mode.Application))
```

The `int 0` represents the current application and `int 1` would reference the first passed in foreign app. Likewise, `int 2` would represent the second passed in foreign application. The actual contract IDs can also be specified as long as the contract is in the contracts array. Similar to the `app_local_get_ex` opcode, generally, there will be branching logic testing whether the value was found or not. 

## Summary of State Operations

| Context            | Write            | Read                | Delete           | Check If Exists     |
| ---                | ---              | ---                 | ---              | ---                 |
| Current App Global | `app_global_put` | `app_global_get`    | `app_global_del` | `app_global_get_ex` |
| Current App Local  | `app_local_put`  | `app_local_get`     | `app_local_del`  | `app_local_get_ex`  |
| Other App Global   |                  | `app_global_get_ex` |                  | `app_global_get_ex` |
| Other App Local    |                  | `app_local_get_ex`  |                  | `app_local_get_ex`  |


# Checking the Transaction Type in a Smart Contract
The `ApplicationCall` transaction types defined in [The Lifecycle of a Stateful Smart Contract](#the-lifecycle-of-a-stateful-smart-contract) can be checked within the TEAL code by examining the `OnCompletion` transaction property. 

```text tab="TEAL"
int NoOp 
//or OptIn, UpdateApplication, DeleteApplication, CloseOut, ClearState 
txn OnCompletion
==
```

```python tab="PyTeal"
program = OnComplete.NoOp == Txn.on_completion()

print(compileTeal(program, Mode.Application))
```

# Passing Arguments To Stateful Smart Contracts
Arguments can be passed to any of the supported application transaction calls, including create. The number and type can also be different for any subsequent calls to the stateful smart contract. The `goal` CLI supports passing strings, ints, base64 encoded data, and addresses as parameters. To pass a parameter supply the `--app-arg` option to the call and supply the value according to the format shown below.

Argument Type | Example
------------ | ------------- 
String | `goal app call --app-arg "str:mystring".....` 
Integer | `goal app create --app-arg "int:5".....` 
Address | `goal app call --app-arg "addr:address-string".....`
Base64 | `goal app call --app-arg "b64:A==".....`

These parameters are loaded into the arguments array. TEAL opcodes are available to get the values within the array. The primary argument opcode is the `ApplicationArgs` opcode and can be used as shown below.

```text tab="TEAL"
txna ApplicationArgs 1
byte "claim" 
==
```

```python tab="PyTeal"
program = Txn.application_args[1] == Bytes("claim")

print(compileTeal(program, Mode.Application))
```

This call gets the second passed in argument and compares it to the string "claim".

A global variable is also available to check the size of the transaction argument array. This size can be checked with a simple TEAL call.

```text tab="TEAL"
txn NumAppArgs
int 4
==
```

```python tab="PyTeal"
program = Txn.application_args.length() == Int(4)

print(compileTeal(program, Mode.Application))
```

The above TEAL code will push a 0 on the top of the stack if the number of parameters in this specific transaction is anything other than 4, else it will push a 1 on the top of the stack. Internally all transaction parameters are stored as byte slices (byte-array value). Integers can be converted using the `btoi` opcode.

```text tab="TEAL"
txna ApplicationArgs 0
btoi
```

```python tab="PyTeal"
program = Btoi(Txn.application_args[0])

print(compileTeal(program, Mode.Application))
```

!!! info
    Argument passing for stateful smart contracts is very different from passing arguments to stateless smart contracts. 

The total size of all parameters is limited to 2KB in size.


# Creating the Smart Contract
Before creating a stateful smart contract, the code for the `ApprovalProgram` and the `ClearStateProgram` program should be written. The SDKs and the `goal` CLI tool can be used to create a smart contract application. To create the application with `goal` use a command similar to the following.

```
$ goal app create --creator [address]  --approval-prog [approval_program.teal] --clear-prog [clear_state_program.teal] --global-byteslices [number-of-global-byteslices] --global-ints [number-of-global-ints] --local-byteslices [number-of-local-byteslices] --local-ints [number-local-ints] --extra-pages [number of extra 2KB pages]
```

The creator is the account that is creating the application and this transaction is signed by this account. The approval program and the clear state program should also be provided. The number of global and local byte slices (byte-array value) and integers also needs to be specified. These represent the absolute on-chain amount of space that the smart contract will use. Once set, these values can never be changed. The key is limited to 64 bytes. The key plus the value is limited to 128 bytes total. When the smart contract is created the network will return a unique ApplicationID. This ID can then be used to make `ApplicationCall` transactions to the smart contract. 

When creating a stateful smart contract, there is a limit of 64 key-value pairs that can be used by the contract for global storage and 16 key-value pairs that can be used for local storage. When creating the smart contract the amount of storage can never be changed once the contract is created. Additionally, the minimum balance is raised for any account that participates in the contract. See [Minimum Balance Requirement for Smart Contracts](#minimum-balance-requirement-for-a-smart-contract) described below for more detail.

Stateful smart contracts are limited to 2KB total for the compiled approval and clear programs. This size can be increased up to 3 additional 2KB pages, which would result in an 8KB limit for both programs. Note the size increases will also increase the minimum balance requirement for creating the application. To request additional pages, the setting (`extra-pages`) is available when creating the stateful smart contract using `goal`. These extra pages can also be requested using the SDKs. This setting allows setting up to 3 additional 2KB pages.

!!! warning
	Currently, applications up to 2KB in size can be updated, while applications between 2KB and 8KB in size can not. This discrepancy will be resolved in the next consensus upgrade and all applications of any size, including ones that have already been created, will be updateable. Please note - 2KB applications can still be updated. 

!!! info    
    Accounts can only opt into or create up to 10 stateful smart contracts.

# Opt into the Smart Contract
Before any account, including the creator of the smart contract, can begin to make Application Transaction calls that use local state, it must first opt into the smart contract. This prevents accounts from being spammed with smart contracts. To opt in, an `ApplicationCall` transaction of type `OptIn` needs to be signed and submitted by the account desiring to opt into the smart contract. This can be done with the `goal` CLI or the SDKs.

```
$ goal app optin  --app-id [ID-of-Contract] --from [ADDRESS]
```

When this transaction is submitted, the `ApprovalProgram` of the smart contract is called and if the call succeeds the account will be opted into the smart contract. The simplest TEAL program to handle this call would just put 1 on the stack and return. 

```text tab="TEAL"
int OptIn
txn OnCompletion
==
bz notoptingin
int 1
Return
notoptingin:
.
.
```

```python tab="PyTeal"
# just a placeholder; change this
not_opting_in = Seq([
    Return(Int(0))
])

program = If(OnComplete.OptIn == Txn.on_completion(), 
            Return(Int(1)), 
            not_opting_in)

print(compileTeal(program, Mode.Application))
```

Other contracts may have much more complex opt in logic. TEAL also provides an opcode to check whether an account has already opted into the contract.

```text tab="TEAL"
int 0 
txn ApplicationID
app_opted_in
```

```python tab="PyTeal"
program = App.optedIn(Int(0), Txn.application_id())

print(compileTeal(program, Mode.Application))
```

In the above example, the int 0 is a reference index into the accounts array, where 0 is the sender. A 1 would be the first account passed into the call and so on. The actual address may also be specified as long as it is in the accounts array. The `txn ApplicationID` refers to the current application ID, but technically any application ID could be used as long as its ID is in the applications array.

!!! info
    Accounts can only opt into or create up to 10 stateful smart contracts

!!! info
    Applications that only use global state do not require accounts to opt in.

# Call the Stateful Smart Contract
Once an account has opted into a stateful smart contract it can begin to make calls to the contract. These calls will be in the form of `ApplicationCall` transactions that can be submitted with `goal` or the SDKs. Depending on the individual type of transaction as described in [The Lifecycle of a Stateful Smart Contract](#the-lifecycle-of-a-stateful-smart-contract), either the `ApprovalProgram` or the `ClearStateProgram` will be called. Generally, individual calls will supply application arguments. See [Passing Arguments to a Smart Contract](#passing-arguments-to-stateful-smart-contracts) for details on passing arguments.

```
$ goal app call --app-id 1 --app-arg "str:myparam"  --from [ADDRESS]
```

The call must specify the intended contract using the `--app-id` option. Additionally, the `--from` option specifies the sender’s address. In this example, a string parameter is passed with this call. TEAL can use these parameters to make decisions on how to handle the call.

```text tab="TEAL"
byte "myparm" 
txna ApplicationArgs 0
==
bz not_my_parm
//handle my_parm
return
not_my_parm:
//handle not_my_parm
```

```python tab="PyTeal"
# placeholder; change this logic in both cases
is_my_parm = Seq([
        Return(Int(1))
    ])

is_not_my_parm = Seq([
        Return(Int(1))
    ])

program = If(Bytes("myparm") == Txn.application_args[0], 
            is_my_parm, 
            is_not_my_parm)

print(compileTeal(program, Mode.Application))
```

# Update Stateful Smart Contract
A stateful smart contract’s programs can be updated at any time. This is done by an `ApplicationCall` transaction type of `UpdateApplication`. This operation can be done with `goal` or the SDKs and requires passing the new programs and specifying the application ID.

```
goal app update --app-id=[APPID] --from [ADDRESS]  --approval-prog [new_approval_program.teal]   --clear-prog [new_clear_state_program.teal]
```

The one caveat to this operation is that global or local state requirements for the smart contract can never be updated.

As stated earlier, anyone can update the program. If this is not desired and you want only the original creator to be able to update the programs, code must be added to your `ApprovalProgram` to handle this situation. This can be done by comparing the global `CreatorAddress` to the sender address.

```text tab="TEAL"
global CreatorAddress
txn Sender
==
assert
```

```python tab="PyTeal"
program = Assert(Global.creator_address() == Txn.sender())

print(compileTeal(program, Mode.Application))
```


Or alternatively, the TEAL code can always return a 0 when an `UpdateApplication` application call is made to prevent anyone from ever updating the application code.

```text tab="TEAL"
int UpdateApplication
txn OnCompletion
==
bz not_update
int 0
return
```

```python tab="PyTeal"
# placeholder logic; change this
is_not_update = Seq([
    Return(Int(1))
])

program = If(OnComplete.UpdateApplication == Txn.on_completion(), 
            Return(Int(0)), 
            is_not_update)

print(compileTeal(program, Mode.Application))
```

# Delete Stateful Smart Contract
To delete a smart contract, an `ApplicationCall` transaction of type `DeleteApplication` must be submitted to the blockchain. The `ApprovalProgram` handles this transaction type and if the call returns true the application will be deleted. This can be done using `goal` or the SDKs. 

```
$ goal app delete --app-id=[APPID] --from [ADDRESS]
```

When making this call the `--app-id` and the `--from` options are required. Anyone can delete a smart contract. If this is not desired, logic in the program must reject the call. Using a method described in [Update Stateful Smart Contract](#update-stateful-smart-contract) must be supplied. 

# Global Values in Smart Contracts
Smart contracts have access to many global variables. These variables are set for the blockchain, like the minimum transaction fee (MinTxnFee). As another example of Global variable use, in the [Atomic Transfers and Transaction Properties](#atomic-transfers-and-transaction-properties) section of this guide, `GroupSize` is used to show how to get the number of transactions that are grouped within a smart contract call. Stateful smart contracts also have access to the `LatestTimestamp` global which represents the latest confirmed block's Unix timestamp. This is not the current time but the time when the last block was confirmed. This can be used to set times on when the contract is allowed to do certain things. For example, a contract may only allow accounts to opt in after a start date, which is set when the contract is created and stored in global storage.

```text tab="TEAL"
global LatestTimestamp
byte "StartDate"
app_global_get
>=
```

```python tab="PyTeal"
program = Global.latest_timestamp() >= App.globalGet(Bytes("StartDate"))

print(compileTeal(program, Mode.Application))
```

# Atomic Transfers and Transaction Properties
The [TEAL opcodes](../../../reference/teal/opcodes.md) documentation describes all transaction properties that are available within a TEAL program. These properties can be retrieved using TEAL.

```text tab="TEAL"
txn Amount
```

```python tab="PyTeal"
program = Txn.amount()

print(compileTeal(program, Mode.Application))
```

In many common patterns, the stateful TEAL contract will be combined with other Algorand technologies such as Algorand Assets, Atomic Transfers, or Stateless Smart Contracts to build a complete application. In the case of Atomic transfers, more than one transaction’s properties can be checked within the stateful smart contract. The number of transactions can be checked using the `GroupSize` global property. If the value is greater than 1, then the call to the stateful smart contract is grouped with more than one transaction.

```text tab="TEAL"
global GroupSize
int 2
==
```

```python tab="PyTeal"
program = Global.group_size() == Int(2)

print(compileTeal(program, Mode.Application))
```

The above TEAL will be true if there are two transactions submitted at once using an Atomic transfer. To access the properties of a specific transaction in the atomic group use the `gtxn` opcode.

```text tab="TEAL"
gtxn 1 TypeEnum
int pay
==
```

```python tab="PyTeal"
program = Gtxn[1].type_enum() == TxnType.Payment

print(compileTeal(program, Mode.Application))
```

In the above example, the second transaction’s type is checked, where the `int pay` references a payment transaction. See the [opcodes](../../../reference/teal/opcodes.md) documentation for all transaction types. Note that the `gtxn` call is a zero-based index into the atomic group of transactions. The `gtxns` opcode could also have been used to retrieve the index into the atomic group from the top of the stack instead of hard coding the index. If the TEAL program fails, all transactions in the group will fail.

If any transaction in a group of transactions is a call to a stateful smart contract, the opcodes `gtxna` and `gtxnsa` can be used to access any of the transactions array values.

```text tab="TEAL"
// get the first argument of the previous transaction
// in a stateful smart contract
txn GroupIndex
int 1
-
gtxnsa ApplicationArgs 0
```

```python tab="PyTeal"
program = Gtxn[Txn.group_index() - Int(1)].application_args[0]

print(compileTeal(program, Mode.Application))
```

# Using Assets in Smart Contracts
Stateful contract applications can work in conjunction with Algorand Assets. In addition to normal asset transaction properties, such as asset amount, sender, and receiver, TEAL provides an opcode to interrogate an account’s asset balance and whether the asset is frozen. This opcode `asset_holding_get` can be used to retrieve an asset balance or check whether the asset is frozen for any account in the transaction accounts array. The asset must also be in the assets array.

```text tab="TEAL"
int 0
int 2
asset_holding_get AssetBalance
bnz has_balance
int 0 
return
has_balance:
//balance value is now on top of the stack
```

```python tab="PyTeal"
asset_balance = AssetHolding.balance(Int(0), Int(2))

program = Seq([
        asset_balance,
        If(asset_balance.hasValue(), 
            Return(Int(1)), 
            Return(Int(0)))
    ])

print(compileTeal(program, Mode.Application))
```

This opcode takes two parameters. The first represents an index into the accounts array where `int 0` is the sender of the transaction’s address. If additional accounts are passed in using the `--app-account` `goal` option then higher index numbers would be used to retrieve values. The actual address can also be specified as long as is it in the accounts array. The second parameter is the Asset ID of the asset to examine. This can be either an index into the assets array or the actual asset ID. The asset must be in the assets array for the call to be successful. In this example, the asset ID is 2. This opcode supports getting the asset balance and the frozen state of the asset for the specific account. To get the frozen state, replace `AssetBalance` above with `AssetFrozen`. This opcode also returns two values to the top of the stack. The first is a 0 or  1, where 0 means the asset balance was not found and 1 means an asset balance was found in the accounts balance record.

It is also possible to get an Asset’s configuration information within a smart contract if the asset ID is passed with the transaction in the assets array. This can be done with `goal` by supplying the `--foreign-asset` parameter. The value of the parameter should be the asset ID. To read the configuration, the `asset_params_get` opcode must be used. This opcode should be supplied with one parameter, which is the index into the assets array or the actual asset ID.

```text tab="TEAL"
int 0
asset_params_get AssetTotal
```

```python tab="PyTeal"
program = AssetParam.total(Int(0))

print(compileTeal(program, Mode.Application))
```

This call returns two values. The first is a 0 or 1 indicating if the parameter was found and the second contains the value of the parameter. See the [opcodes](../../../reference/teal/opcodes.md) documentation for more details on what additional parameters can be read.

# Creating An Asset or Contract Within A Group Of Transactions
The Algorand Protocol assigns an identifier (ID) when creating Algorand Standard Assets (ASA) or Stateful Smart Contracts. These IDs are used to refer to the asset or the contract later when either is used in a transaction or a call to the stateful smart contract. Because these IDs are assigned when the asset or the contract is created, the ID is not available until after the creation transaction is fully executed. When creating either in an atomic transfer TEAL can be used to retrieve these IDs. For example you may want to store the asset ID or another smart contract ID in the contract’s state for later usage. 

This operation can be performed by using one of two opcodes (`gaid` and `gaids`). With the `gaid` opcode, the specific transaction to read must be passed to the command. The `gaids` opcode will use the last value on the stack as the transaction index.

```tab="TEAL"
// Get the created id of the asset created in the first tx
gaid 0
// Get the created id of the asset created in the second tx
int 1
gaids
```

# Sharing Data Between Contracts
In additon to reading another contract's state, data can be passed between contracts. Algorand’s AVM gives smart contracts access to scratch space that can be used to temporarily store values, while the contract is executing. TEAL allows other contracts to read this scratch space. Scratch space can only be read by other transactions that the specific transaction is grouped with atomically. Additionally, because grouped transactions execute in the order they are grouped, contracts can not read scratch space for transactions that occur after the current contract transaction. 

This operation can be performed by using one of two opcodes (`gload` and `gloads`). With the `gload` opcode, the specific transaction to read and the slot number must be passed to the command. The `gloads` opcode will use the last value on the stack as the transaction index and must be passed the slot number to read.

For example, with two grouped stateful smart contracts the following code can be used.

Store an integer in scratch space in the first transaction.

```tab="TEAL"
int 777
store 10
```

In the second transaction read the stored value.

```tab="TEAL"
// read the first 
// transactions 10th
// slot of scratch space
gload 0 10
```

# Reading a Smart Contracts State
In addition to being able to read the state of a smart contract using TEAL, these global and local values can be read externally with the SDKs and `goal`. These reads are not transactions and just query the current state of the contract. 

```
$ goal app read --app-id 1 --guess-format --global --from [ADDRESS]
```

In the above example, the global state of the smart contract with the application ID of 1 is returned. The `--guess-format` opt in the above example tries programmatically to display the properly formatted values of the state variables. To get the local state, replace `--global` with `--local` and note that this call will only return the local state of the `--from` account.

Here is an example output with 3 keys/values:
```json
{
  "Creator": {
    "tb": "FRYCPGH25DHCYQGXEB54NJ6LHQG6I2TWMUV2P3UWUU7RWP7BQ2BMBBDPD4",
    "tt": 1
  },
  "MyBytesKey": {
    "tb": "hello",
    "tt": 1
  },
  "MyUintKey": {
    "tt": 2,
    "ui": 50
  }
}
```

Interpretation:

* the keys are `Creator`, `MyBytesKey`, `MyUintKey`.
* the field `tt` is the type of the value: 1 for byte slices (byte-array value), 2 for uint.
* when `tt=1`, the value is in the field `tb`. Note that because of `--guess-format`, the value for `Creator` is automatically converted to an Algorand address with checksum (as opposed to a 32-byte public key, see [the accounts overview](../../accounts#transformation-public-key-to-algorand-address)).
* when `tt=2`, the value is in the field `ui`.

# Differences Between Stateful and Stateless Smart Contracts
Smart Contracts in Algorand are either stateful or stateless, where stateful contracts actually store values on blockchain and stateless are used to approve spending transactions. In addition to this primary difference, several of the TEAL opcodes are restricted to only be used with specific smart contract types. For example, the `ed25519verify` opcode can only be used in stateless smart contracts, and the `app_opted_in` opcode can only be used in stateful smart contracts. To easily determine where an opcode is valid, the [TEAL opcodes](../../../reference/teal/opcodes.md) documentation supplies a `Mode` field that will either designate `Signature ` or `Application`. The `Signature` mode refers to only stateless smart contracts and the `Application` mode refers to the stateful smart contracts. If the `Mode` attribute is not present on the opcode reference, the call can be used in either smart contract type.

# Boilerplate Stateful Smart Contract
As a way of getting started writing stateful smart contracts, the following boilerplate template is supplied. The code provides labels or handling different `ApplicationCall` transactions and also prevents updating and deleting the smart contract.

```text tab="TEAL"
#pragma version 4

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

handle_optin:
// Handle OptIn

handle_closeout:
// Handle CloseOut

// By default, disallow updating or deleting the app. Add custom authorization
// logic below to allow updating or deletion in certain circumstances.
handle_updateapp:
handle_deleteapp:
err
```

```python tab="PyTeal"
from pyteal import *
# Handle each possible OnCompletion type. We don't have to worry about
# handling ClearState, because the ClearStateProgram will execute in that
# case, not the ApprovalProgram.
def approval_program():
    handle_noop = Seq([
        Return(Int(1))
    ])

    handle_optin = Seq([
        Return(Int(1))
    ])

    handle_closeout = Seq([
        Return(Int(1))
    ])

    handle_updateapp = Err()

    handle_deleteapp = Err()

    program = Cond(
        [Txn.on_completion() == OnComplete.NoOp, handle_noop],
        [Txn.on_completion() == OnComplete.OptIn, handle_optin],
        [Txn.on_completion() == OnComplete.CloseOut, handle_closeout],
        [Txn.on_completion() == OnComplete.UpdateApplication, handle_updateapp],
        [Txn.on_completion() == OnComplete.DeleteApplication, handle_deleteapp]
    )
    return program

with open('boilerplate_approval_pyteal.teal', 'w') as f:
    compiled = compileTeal(approval_program(), Mode.Application)
    f.write(compiled)
```

# Minimum Balance Requirement for a Smart Contract
When creating or opting into a stateful smart contract your minimum balance will be raised. The amount at which it is raised will depend on the amount of on-chain storage that is used and additional 2KB pages requested when deploying the contract. The calculation for the amount for creation is as follows.

100000*(1+ExtraProgramPages) + (28500)*schema.NumUint + (50000)*schema.NumByteSlice

100000 microAlgo base fee for opting in or creating multiplied by 1 plus the number of extra pages. 50000 microAlgos for each key byte-slice (byte-array value) and 28500 microAlgos for each integer value.

Any user that opts into the stateful smart contract will have their minimum balance increased by the following amount.

100000 + (28500)*schema.NumUint + (50000*schema.NumByteSlice)

100000 microAlgo base fee for opting in. 50000 microAlgos for each key byte-slice (byte-array value) and 28500 microAlgos for each integer value.

Note that global storage is actually stored in the creator account, so that account is responsible for the global storage minimum balance and when a user options in (including the creator) that account is responsible for the minimum balance of local storage. As an example, suppose a smart contract has one 1 global key-value-pair of type byteslice and one 1 local storage key-value pair of type integer. The creator would need an additional 150000 microAlgos in its balance.

100000 to create the contract +  50000 for the global byte slice.

Any account that opts into the contract would have its balance requirement raised 128500 microAlgos.

100000 to opt into the contract + 28500 for the locally stored integer.
