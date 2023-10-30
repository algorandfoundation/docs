title: Overview

# Introduction

Algorand smart contracts are pieces of logic that reside on the Algorand blockchain and are remotely callable. These contracts are primarily responsible for implementing the logic associated with a distributed application. Smart contracts are referred to as stateful smart contracts or applications in the Algorand documentation. Smart contracts can generate asset and payment transactions allowing them to function as Escrow accounts on the Algorand blockchain. Smart contracts can also store values on the blockchain. This storage can be global, local, or box storage. Local storage refers to storing values in an accounts balance record if that account participates in the contract. Global storage is data that is specifically stored on the blockchain for the contract globally. Box storage is also global and allows contracts to use larger segments of storage. Like smart signatures, smart contracts are written in Python using PyTeal or TEAL and can be deployed to the blockchain using either the `goal` command-line tool or the SDKs. The recommended approach for writing smart contracts is to use the Python SDK with the PyTeal library.  

See the [*PyTeal Documentation*](/docs/get-details/dapps/writing-contracts/pyteal) for information on building smart contracts in Python.

See the [*TEAL Reference Guide*](../../avm/teal/specification.md) to understand how to write TEAL and the [*TEAL Opcodes*](../../avm/teal/opcodes) documentation that describes the opcodes available. This guide assumes that the reader is familiar with [TEAL](../../avm/teal/index.md).

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

When smart contracts execute, they may require data stored within the ledger for evaluation. The following sections define how data access by a smart contract is achieved as of version 9 of the AVM. For previous versions of the AVM see the [TEAL specification](https://github.com/algorandfoundation/specs/blob/master/dev/TEAL.md#resource-availability).

<center>![Smart Contract](/docs/imgs/refarray_1.png)</center>


## Resource availability
The AVM provides several opcodes(functions) that support reading ledger data. The primary ledger data items that can be read are accounts, assets(ASAs), applications(smart contracts), and boxes. These ledger items allow the contract to look at specific information about each of these components such as reading an accounts’s Algo balance or examining the immutable properties of an ASA. These items are stored in the ledger in what can be thought of as large lists.

<center>![Smart Contract](/docs/imgs/refarray_2.png)</center>


Two of these components ( accounts and applications ) also contain sublists of potentially large datasets. For example, an account can opt into potentially an unlimited set of assets or applications (which store user local state). Additionally, smart contracts allow boxes of data to be stored within the ledger and are also potentially unlimited. For example, a smart contract may create a unique box of arbitrary data for every user of the smart contract.

<center>![Smart Contract](/docs/imgs/refarray_3.png)</center>

To maintain a high level of performance, the AVM restricts how much of the ledger can be viewed within a given contract execution. This is implemented with a set of reference arrays passed with each application transaction that define the specific ledger items that are **available** during execution. These arrays are the Account, Asset, Application, and Boxes arrays. 

These four arrays are limited to eight total values combined (per application transaction). The accounts array can have no more than four accounts. The values passed in using these arrays can change per application transaction. 

Additionally, when accessing one of the sublists of items, the application transaction must include both the top level item and the nested list item within the same call. For example, to read an ASA balance for a specific account, the account and the asset must both be present in the respective accounts and assets arrays for the given transaction.

<center>![Smart Contract](/docs/imgs/refarray_4.png)</center>

By default the reference arrays are empty with the exception of the accounts array and the applications array. The Accounts array contains the transaction sender's address, and the Application array contains the called smart contract ID.

<center>![Smart Contract](/docs/imgs/refarray_5.png)</center>

If a specific ledger item (account, asset, application, account+asset, account+application, application+box) can be read during a smart contract’s execution, it is said to be **available**. These four arrays determine the **availability** of these six unique ledger items.

AVM opcodes that require an address, asset ID, or application ID can access the value of these from the reference arrays using slot referencing or by passing the actual address or ID. Using the actual address or ID is the preferred method as shown in the example below.

```python
@app.external
def asa_balance_with_reference_type(
    account: abi.Account, asa: abi.Asset, *, output: abi.Uint64
) -> Expr:
    asset_balance = AssetHolding.balance(account.address(), asa.asset_id())
    return Seq(
        asset_balance,
        Assert(asset_balance.hasValue()),
        output.set(asset_balance.value()),
    )

```
 
Slot referencing can be used as well, but will not benefit from resource sharing.

```python
@app.external
def asa_balance_with_slot_referencing(*, output: abi.Uint64) -> Expr:
    asset_balance = AssetHolding.balance(Txn.accounts[0], Txn.assets[0])
    return Seq(
        asset_balance,
        Assert(asset_balance.hasValue()),
        output.set(asset_balance.value()),
    )

```

!!! info
    Since V9 of the AVM, resource sharing is now available. Specific addresses/IDs should be used and not slot references. Slot referencing may be deprecated in the future.  Read more about resource sharing below.

!!! info
    If a smart contract uses an inner transaction to initiate a call to another smart contract, the inner contract automatically inherits the resource **availability** of the top level smart contract.

## Resource sharing
Resources are shared across atomically grouped transactions. This means that if two smart contracts are in the same atomic group, they share **availability**. 

For example, say you have two smart contract call transactions grouped together, transaction #1 and transaction #2. Transaction #1 has asset 123456 in its assets array, and transaction #2 has asset 555555 in its assets array. Both assets are **available** to both smart contract calls during evaluation. 

<center>![Smart Contract](/docs/imgs/refarray_6.png)</center>

Note that when accessing a sublist item (account+asa, account+application local state, application+box), both items need to be in the same transaction’s set of arrays. For example, you cannot have account A in transaction #1 and asset Z in transaction #2 and then try to get the balance of asset Z for account A. Asset Z and account A must be in the same application transaction. If both asset Z and account A are in transaction #1’s arrays, then A’s balance for Z is **available** to both transactions during evaluation. 

Because Algorand supports grouping up to 16 transactions simultaneously, this pushes the **available** resources up to 8x16 or 128 items, if all 16 transactions are application transactions.

If an application transaction is grouped with other types of transactions, other resources will be made **available** to the smart contract called in the application transaction. For example, if an application transaction is grouped with a payment transaction, the payment transaction’s sender and receiver accounts are **available** to the smart contract. Additionally, if the `CloseRemainderTo` field is set, that account will also be **available** to the smart contract. The table below summarizes what each transaction type adds to resource **availability**.

| Transaction | Type | Availability Notes | 
|------|------|------|
| Payment | `pay` | `txn.Sender`, `txn.Receiver`, and `txn.CloseRemainderTo` (if set)|
| Key Registration | `keyreg` | `txn.Sender` |
| Asset Config/Create | `acfg` | `txn.Sender`, `txn.ConfigAsset`, and the `txn.ConfigAsset` holding of `txn.Sender`|
| Asset Transfer | `axfer` | `txn.Sender`, `txn.AssetReceiver`, `txn.AssetSender` (if set), `txnAssetCloseTo` (if set), `txn.XferAsset`, and the `txn.XferAsset` holding of each of those accounts. |
| Asset Freeze | `afrz` | `txn.Sender`, `txn.FreezeAccount`, `txn.FreezeAsset`, and the `txn.FreezeAsset` holding of `txn.FreezeAccount`. The `txn.FreezeAsset` holding of `txn.Sender` is not made available. |

!!! info
    In future versions of the AVM, the Atomic Transaction Composer may be enhanced to automatically generate reference arrays for specific transaction groups

!!! info
	If any application or asset is created within a group of transactions, they are **available** resources as long as they are created before they are accessed.

## Boxes Array
Boxes function similar to the other arrays but differ is significant ways which are explained in detail in the [Boxes section of the documentation](state.md#box-details).

## Loading the Reference Arrays 
For information on adding entries to reference arrays when calling a smart contract see the [ATC documentation](https://developer.algorand.org/docs/get-details/atc/#foreign-references).

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
