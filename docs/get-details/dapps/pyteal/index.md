title: PyTeal

[PyTeal](https://github.com/algorand/pyteal) is a python library for generating [TEAL](../avm/teal/index.md) programs that provides a convenient and familiar syntax. 

To quickly get PyTeal installed and running, see the [Getting started tutorial](../../../get-started/dapps/pyteal.md) for PyTeal.

Complete installation instructions and developer guides are available in the [PyTeal documentation](https://pyteal.readthedocs.io/en/latest/).

# PyTeal overview

This section assumes the reader is familiar with [Smart Contracts](../smart-contracts/index.md#smart-contracts) and [Smart Signatures](../smart-contracts/index.md#smart-signatures).

When building a dApp that makes use of smart contracts or smart signatures (smartsigs), PyTeal makes implementation of more complex logic much simpler than writing the TEAL manually. 

Generally, developers install PyTeal, write the contract in Python using their preferred editor, and then use PyTeal’s `compileProgram` method to produce the TEAL code. The TEAL source can be compiled into bytecode and deployed to the blockchain.

For most applications, these contracts will only be a portion of the dApp’s architecture. Typically, developers will build functionality in the dApp that resides on the blockchain and some front end to interact with the smart contracts. 

The PyTeal contracts can be pre-compiled into TEAL and used directly or the logic to generate TEAL dynamically may be integrated into the dApp front end workflow. For example, a front end may provide something like an exchange that allows limit orders to be created based on a template and then deployed once a user opens an order. In this case, the complete limit order may be implemented as part of a smart contract that is deployed when the order is opened by the dApp. For more information on deployment models see [What is a dApp](../../../get-started/dapps/index.md) in the developer documentation.

# Building PyTeal smart contracts

On Algorand, smart contracts are small programs that are evaluated when the contract is deployed or called. These contracts make use of a set of functions [(opcodes)](../avm/teal/opcodes.md) to be evaluated against the context it was called with. This context includes the current state of its storage values, the transactions in the group, any arguments and references to accounts, assets, or other applications.

Besides evaluating the logic to approve or reject a transaction, the contracts may cause side effects. Side effects a contract may produce include changes to an Applications global or local state, or producing their own [transactions](../../transactions/transactions.md). These transactions are [atomic](../../atomic_transfers.md) with the outer group the application was invoked with.  The transactions produced during contract evaluation are by default sent from the Applications associated account.

For more information on smart contracts, see the [smart contract documentation](../smart-contracts/apps/index.md).


## Hello PyTeal

When building smart contracts in PyTeal it is important to realize that a smart contract actually consists of two programs. These are called the approval and the clear programs. In PyTeal both of these programs are generally created in the same Python file. So the beginning of a PyTeal program will contain logic similar to the following:

```python
#samplecontract.py
from pyteal import *

"""Basic Counter Application"""

def approval_program():
    program = Return(Int(1))
    # Mode.Application specifies that this is a smart contract
    return compileTeal(program, Mode.Application, version=5)

def clear_state_program():
    program = Return(Int(1))
    # Mode.Application specifies that this is a smart contract
    return compileTeal(program, Mode.Application, version=5)

print(approval_program())
```

In the above example, a function is defined to return each of the two programs. The `compileTeal` function compiles the program set to the `program` variable. In this case, we are just returning the integer 1 for both programs. The other arguments passed to `compileTeal` set the mode (Application or Signature) and version of TEAL to produce. The output from calling the `compileTeal` method is a string representing the compiled TEAL source.

Running the program, the output should look something like:
```teal
#pragma version 5
int 1
return
```

This output represents the TEAL source for the program. The TEAL source must be compiled to bytecode in order to make use of it on-chain.  This can be done using [goal](../../../clis/goal/clerk/compile.md) or one of the SDKs through the [REST](../../../rest-apis/algod/v2.md#post-v2tealcompile) interface.


## A note about PyTeal Expressions

Before going further, it is important to have the right mental model for thinking about PyTeal contracts.

    A PyTeal program is a PyTeal Expression composed of other PyTeal Expressions. 
    
Reread that sentence.

When first starting to write contracts using PyTeal, the most common point of confusion comes from trying to include _Python_ native expressions in the the _PyTeal_ Expression tree.  In the above example the entire program was `Return(Int(1))`.  Here `Return` is a PyTeal Expression that takes as an argument another PyTeal Expression. The argument we pass it is the PyTeal Expression `Int(1)`. 

If we were to call `print(program)` we'd see our tiny Expression tree:
```
(Return (Int: 1))
```

As a counter example, if instead of passing `Int(1)`, we make the program `Return(1)`, PyTeal would throw an error complaining that `1` is not a valid PyTeal Expression.  

!!! note
    If you ever see any python error like `...has no attribute 'type_of'` PyTeal is probably trying to tell you you included something that is _not_ a valid PyTeal Expression

While _only_ PyTeal Expressions may be included in the Expression tree representing the program, the Expressions may be generated as part of running the Python program. This allows the author to make use of Python scripting logic to generate a the PyTeal Expression tree based on more complex logic. 


## Writing a simple PyTeal Contract

Here we'll start updating our example to allow a more complex logical flow, showing how to build up a PyTeal contract using PyTeal Expressions.

Below is a simple logical switch statement used to route evaluation to different set of logic based on a Transaction's [`OnComplete`](../../transactions/transactions.md#application-call-transaction) value.  

```python
def approval_program():
    program = Cond(
        [Txn.application_id() == Int(0), handle_creation],
        [Txn.on_completion() == OnComplete.OptIn, handle_optin],
        [Txn.on_completion() == OnComplete.CloseOut, handle_closeout],
        [Txn.on_completion() == OnComplete.UpdateApplication, handle_updateapp],
        [Txn.on_completion() == OnComplete.DeleteApplication, handle_deleteapp],
        [Txn.on_completion() == OnComplete.NoOp, handle_noop]
    )

    return compileTeal(program, Mode.Application, version=5)
```

The `program` variable is set to a [PyTeal `Cond` Expression](https://pyteal.readthedocs.io/en/latest/control_structures.html?highlight=Cond#chaining-tests-cond) which is an example of a Control Flow statement. Other Control flow statements are documented in the [PyTeal documentation](https://pyteal.readthedocs.io/en/latest/control_structures.html?highlight=seq#).

A `Cond` expression allows several conditions to be evaluated in order, taking a number of arguments as [`Condition Expression`, `Body Expression`].  The `Condition Expression` must evaluate to True or False. For a given invocation, the first Condition that evaluates True will pass flow of the program to its corresponding `Body Expression`.  If none of the conditions are true the smart contract will return an `err` and fail.  The body for each condition here should reference some other variable or method defined in your PyTeal contract.

!!! note
    For control flow statements like `Cond`, the Body Expression of each condition _MUST_ evaluate to the same type. The current types that are allowed are `none`, `any`, `uint64` or `bytes`

In the above example, most of the conditions check the transaction type using the `on_completion` transaction field. The result of a statement like `Txn.on_completion() == OnComplete.NoOp` is itself a PyTeal Expression, in this case evaluating to True or False.  

!!! note
    Using `Txn` is shorthand for `The transaction that invoked this contract` and should be used inspect transaction fields. Other transactions in the group including the one referenced by `Txn` are available using `Gtxn[n]` where n is the index of the Transaction in the group.  All other transaction fields can be examined using these transaction references, see the [PyTeal documentation](https://pyteal.readthedocs.io/en/latest/accessing_transaction_field.html) for a complete list.

The only condition above that does _not_ check the `on_completion` field is the first one, which only checks the `application_id` field. An `application_id` of `0` on the `Txn` tells us that this is meant to Create the Application. If the program succeeds, it will be assigned an application ID and further invocations will use that ID to call it. In this example the `handle_creation` Expression should handle any initialization necessary for this contract.

### First handler

```python
def approval_program():
    handle_creation = Seq(
        App.globalPut(Bytes("Count"), Int(0)),
        Return(Int(1))
    )

    program = Cond(
        [Txn.application_id() == Int(0), handle_creation],
        [Txn.on_completion() == OnComplete.OptIn, handle_optin],
        [Txn.on_completion() == OnComplete.CloseOut, handle_closeout],
        [Txn.on_completion() == OnComplete.UpdateApplication, handle_updateapp],
        [Txn.on_completion() == OnComplete.DeleteApplication, handle_deleteapp],
        [Txn.on_completion() == OnComplete.NoOp, handle_noop]
    )
    return compileTeal(program, Mode.Application, version=5)
```

Here we defined the `handle_creation` variable to be a Sequence of Expressions using [`Seq`](https://pyteal.readthedocs.io/en/latest/control_structures.html?highlight=seq#chaining-expressions-seq).

The first expression stores a global variable named Count, and its value is set to 0. More information about storing state variables is available in the [PyTeal documentation](https://pyteal.readthedocs.io/en/latest/control_structures.html?highlight=seq#chaining-expressions-seq).

This second expression is the familiar `Return` expression which exits the program with the return value. In this case, it returns a value of 1, indicating success. 

When this specific smart contract is first deployed it will store a global variable named Count with a value of 0 and immediately return success. 

### Handle other OnComplete values

Because our program requires no Local State, opting-in to this contract is not required so we can reject any opt-ins. Similarly, since there is no need to opt-in, there is no reason to allow a close-out, so those can be rejected as well.  Finally, the contract can deny application transactions that attempt to delete or update the contract.

```python

def approval_program():
    handle_creation = Seq([
        App.globalPut(Bytes("Count"), Int(0)),
        Return(Int(1)) # Could also be Approve()
    ])

    handle_optin = Return(Int(0)) # Could also be Reject()

    handle_closeout = Return(Int(0))

    handle_updateapp = Return(Int(0))

    handle_deleteapp = Return(Int(0))


    program = Cond(
        [Txn.application_id() == Int(0), handle_creation],
        [Txn.on_completion() == OnComplete.OptIn, handle_optin],
        [Txn.on_completion() == OnComplete.CloseOut, handle_closeout],
        [Txn.on_completion() == OnComplete.UpdateApplication, handle_updateapp],
        [Txn.on_completion() == OnComplete.DeleteApplication, handle_deleteapp],
        [Txn.on_completion() == OnComplete.NoOp, handle_noop]
    )
    return compileTeal(program, Mode.Application, version=5)

```

All four of these transaction types simply return a 0, which will cause the transactions to fail. This contract now handles all application transactions but the standard NoOp type. 

### Handle NoOp 

Typically the OnCompletion value is set to `NoOp` to make calls to any application. Different logic handling is often achieved by inspecting other transaction fields, especially `application_args`.

Our example requires an add and a deduct function, to increment and decrement the counter respectively, to be handled for NoOp application transactions. 

Which of these two methods is executed will depend on the first element in the `application_args` passed.  For more information on passing parameters to smart contracts, see the [smart contract documentation](../smart-contracts/apps/index.md).

```python
from pyteal import *

"""Basic Counter Application"""

def approval_program():
    handle_creation = Seq([
        App.globalPut(Bytes("Count"), Int(0)),
        Return(Int(1))
    ])

    handle_optin = Return(Int(0))

    handle_closeout = Return(Int(0))

    handle_updateapp = Return(Int(0))

    handle_deleteapp = Return(Int(0))

    handle_noop = Seq(
        # First, lets fail immediately if this transaction is grouped with any others
        Assert(Global.group_size() == Int(1)), 
        Cond(
            [Txn.application_args[0] == Bytes("Add"), add], 
            [Txn.application_args[0] == Bytes("Deduct"), deduct]
        )
    )

    program = Cond(
        [Txn.application_id() == Int(0), handle_creation],
        [Txn.on_completion() == OnComplete.OptIn, handle_optin],
        [Txn.on_completion() == OnComplete.CloseOut, handle_closeout],
        [Txn.on_completion() == OnComplete.UpdateApplication, handle_updateapp],
        [Txn.on_completion() == OnComplete.DeleteApplication, handle_deleteapp],
        [Txn.on_completion() == OnComplete.NoOp, handle_noop]
    )
    return compileTeal(program, Mode.Application, version=5)

```

In the example, we've implemented the `handle_noop` body as a Seq containing 2 Expressions.

The first is an `Assert` that will immediately Reject the Transaction if, in this case, the number of transactions in the group is not exactly 1. The `Global.group_size()` illustrates the use of a PyTeal global variables. See the [PyTeal documentation](https://pyteal.readthedocs.io/en/latest/accessing_transaction_field.html?highlight=global#global-parameters) for other global variables available.

The second Expression in the Sequence is another `Cond` expression is used to route to the correct Expression based on the first argument passed with the application transaction. If neither of these conditions evaluate to True, the transaction is rejected.


### Implement functionality

The final step for the approval program is to implement the add and deduct functions for the smart contract.

```python
def approval_program():
    handle_creation = Seq([
        App.globalPut(Bytes("Count"), Int(0)),
        Return(Int(1))
    ])

    handle_optin = Return(Int(0))

    handle_closeout = Return(Int(0))

    handle_updateapp = Return(Int(0))

    handle_deleteapp = Return(Int(0))

    # Declare the ScratchVar as a Python variable _outside_ the expression tree
    scratchCount = ScratchVar(TealType.uint64)

    add = Seq(
        # The initial `store` for the scratch var sets the value to 
        # whatever is in the `Count` global state variable
        scratchCount.store(App.globalGet(Bytes("Count"))), 
        # Increment the value stored in the scratch var 
        # and update the global state variable 
        App.globalPut(Bytes("Count"), scratchCount.load() + Int(1)),
        Return(Int(1))
    )

     deduct = Seq(
        # The initial `store` for the scratch var sets the value to 
        # whatever is in the `Count` global state variable
        scratchCount.store(App.globalGet(Bytes("Count"))),
        # Check if the value would be negative by decrementing 
        If(scratchCount.load() > Int(0),
            # If the value is > 0, decrement the value stored 
            # in the scratch var and update the global state variable
            App.globalPut(Bytes("Count"), scratchCount.load() - Int(1)),
        ),
        Return(Int(1))
    )

    handle_noop = Seq(
        Assert(Global.group_size() == Int(1)), 
        Cond(
            [Txn.application_args[0] == Bytes("Add"), add], 
            [Txn.application_args[0] == Bytes("Deduct"), deduct]
        )
    )

    program = Cond(
        [Txn.application_id() == Int(0), handle_creation],
        [Txn.on_completion() == OnComplete.OptIn, handle_optin],
        [Txn.on_completion() == OnComplete.CloseOut, handle_closeout],
        [Txn.on_completion() == OnComplete.UpdateApplication, handle_updateapp],
        [Txn.on_completion() == OnComplete.DeleteApplication, handle_deleteapp],
        [Txn.on_completion() == OnComplete.NoOp, handle_noop]
    )
    return compileTeal(program, Mode.Application, version=5)

```

The contract is modified to create a temporary variable in scratch space. Critically, the declaration of this `ScratchVar` happens _outside_ the Expression tree since the statement `scratchCount = ScratchVar()` is _not_ a valid PyTeal Expression. Smart contracts can hold up to 256 temporary variables in scratch space.  The scratch variable in this example happens to be an integer, byte arrays can also be stored. 

First, the current value of the global variable Count is read for the contract and placed in scratch space. 
Then, the contract either increments this number or decrements and then stores the result into the contract’s global variable. 

On the deduct, an additional `If` expression is used to verify the current global variable is above 0. 

Finally, both methods exit the smart contract call, returning a 1, indicating approval.


### Final product

Because no opt-in is allowed, a clear program need not do anything so we simply return 1, indicating success. The full example is presented below. 

```python
#samplecontract.py
from pyteal import *

"""Basic Counter Application"""

def approval_program():
    handle_creation = Seq([
        App.globalPut(Bytes("Count"), Int(0)),
        Return(Int(1))
    ])

    handle_optin = Return(Int(0))
    handle_closeout = Return(Int(0))
    handle_updateapp = Return(Int(0))
    handle_deleteapp = Return(Int(0))
    scratchCount = ScratchVar(TealType.uint64)

    add = Seq([
        scratchCount.store(App.globalGet(Bytes("Count"))),
        App.globalPut(Bytes("Count"), scratchCount.load() + Int(1)),
        Return(Int(1))
    ])

    deduct = Seq([
        scratchCount.store(App.globalGet(Bytes("Count"))),
         If(scratchCount.load() > Int(0),
             App.globalPut(Bytes("Count"), scratchCount.load() - Int(1)),
         ),
         Return(Int(1))
    ])

    handle_noop = Seq(
        Assert(Global.group_size() == Int(1)), 
        Cond(
            [Txn.application_args[0] == Bytes("Add"), add], 
            [Txn.application_args[0] == Bytes("Deduct"), deduct]
        )
    )


    program = Cond(
        [Txn.application_id() == Int(0), handle_creation],
        [Txn.on_completion() == OnComplete.OptIn, handle_optin],
        [Txn.on_completion() == OnComplete.CloseOut, handle_closeout],
        [Txn.on_completion() == OnComplete.UpdateApplication, handle_updateapp],
        [Txn.on_completion() == OnComplete.DeleteApplication, handle_deleteapp],
        [Txn.on_completion() == OnComplete.NoOp, handle_noop]
    )

    return compileTeal(program, Mode.Application, version=5)


def clear_state_program():
    program = Return(Int(1))
    return compileTeal(program, Mode.Application, version=5)

# print out the results
print(approval_program())
print(clear_state_program())
```

This program can be executed to illustrate compiling the PyTeal and printing out the resultant TEAL code.

```bash
python3 samplecontract.py
```

This example of a smart contract is very simple. Using PyTeal, more sophisticated contracts can be created. To learn more about what can be done in smart contracts, see the [smart contract documentation](../smart-contracts/apps/index.md). The documentation also contains many PyTeal code snippets that can be used within smart contracts.

# Deploying and calling the smart contract

This section explains how to deploy and call the smart contract developed in the previous section.

## Deploying the contract

In the previous section, the development of a simple smart contract was explained. This smart contract can be deployed in many different ways, but generally, this will be done using one of the Algorand SDKs ([Python](../../../sdks/python/index.md), [JavaScript](../../../sdks/javascript/index.md), [Go](../../../sdks/go/index.md), and [Java](../../../sdks/java/index.md)). This section will add additional code to the previous section’s example using the Python SDK to illustrate deploying the example contract.

!!!note
    This example expects the developer to use the sandbox install. Additionally, one account should be set up and funded. See the [Python SDK](../../../sdks/python/index.md) getting started guide for more details.

Before getting into the details of deploying the contract, a couple of global variables must be added to the PyTeal Python example.

```python
# user declared account mnemonics
creator_mnemonic = "REPLACE WITH YOUR OWN MNEMONIC"
# user declared algod connection parameters.
# Node must have EnableDeveloperAPI set to true in its config
algod_address = "http://localhost:4001"
algod_token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
```

The first is a creator mnemonic. This mnemonic is used to recover the private key for the funded account that will own and create the smart contract. Placing a mnemonic like this in code should never be done in production. Typically applications will link to some protected wallet to sign transactions. Some examples of wallets are the Algorand mobile wallet, AlgoSigner, MyAlgo Wallet, and Aikon ORE. When using the Algorand mobile wallet, transactions can be signed using the [Wallet Connect API](../../walletconnect/index.md). The mnemonic is used here for learning purposes only.

The algod_address and algod_token values are the default values to connect to a sandbox installed node. Also note that in this example, the sandbox node is connected to the Algorand TestNet network (eg `./sandbox up testnet`).

In this example, the TEAL for the smart contract will be compiled programmatically by the node. The `EnableDeveloperAPI` configuration parameter must be set to `true` in the node’s configuration to allow this operation. For more information on changing node configuration parameters, see the [developer documentation](../../../run-a-node/reference/config.md). When using the sandbox install, this value is automatically set to true.

Next, a few helper functions need to be added to the sample.

```python
# helper function to compile program source
def compile_program(client, source_code):
    compile_response = client.compile(source_code)
    return base64.b64decode(compile_response['result'])

# helper function that converts a mnemonic passphrase into a private signing key
def get_private_key_from_mnemonic(mn) :
    private_key = mnemonic.to_private_key(mn)
    return private_key


```

The `compile_program` function is a utility function that allows passing the generated TEAL code to a node that will compile and return the byte code. This returned byte code will be used with the application creation transaction (deploying the contract) later.

The `get_private_key_from_mnemonic` function is a utility function that takes a mnemonic (account backup phrase) and returns the private key of the specific account. This will be used in this sample to recover the private key of the funded account of the smart contract creator.

The `wait_for_confirmation` SDK function is a utility function that when called will wait until a specific transaction is confirmed on the Algorand blockchain. This will be used to confirm that the application creation transaction is successful and the smart contract is actively deployed.

As the sample smart contract manipulates global variables, a couple of helper functions are needed to display the contents of these values.

```python
# helper function that formats global state for printing
def format_state(state):
    formatted = {}
    for item in state:
        key = item['key']
        value = item['value']
        formatted_key = base64.b64decode(key).decode('utf-8')
        if value['type'] == 1:
            # byte string
            if formatted_key == 'voted':
                formatted_value = base64.b64decode(value['bytes']).decode('utf-8')
            else:
                formatted_value = value['bytes']
            formatted[formatted_key] = formatted_value
        else:
            # integer
            formatted[formatted_key] = value['uint']
    return formatted

# helper function to read app global state
def read_global_state(client, app_id):
    app = client.application_info(app_id)
    global_state = app['params']['global-state'] if "global-state" in app['params'] else []
    return format_state(global_state)

```

Global variables for smart contracts are actually stored in the creator account’s ledger entry on the blockchain. The location is referred to as global state and the SDKs provide a function to retrieve the application data including the global state. In this example, the function `read_global_state` uses the Python SDK function `application_info` to connect to the Algorand node and retrieve the application information. The function then extracts the global state values if they exist, otherwise returns an empty array. The `format_state` function takes the application data and formats the values for display. For more information on global and local state see the [smart contract documentation](../smart-contracts/apps/index.md).

As covered earlier in this guide, to deploy the contract an application creation transaction must be created and submitted to the blockchain. The SDKs provide a method for creating this transaction. The following code illustrates creating and submitting this transaction.

```python
# create new application
def create_app(client, private_key, approval_program, clear_program, global_schema, local_schema):
    # define sender as creator
    sender = account.address_from_private_key(private_key)

    # declare on_complete as NoOp
    on_complete = transaction.OnComplete.NoOpOC.real

    # get node suggested parameters
    params = client.suggested_params()

    # create unsigned transaction
    txn = transaction.ApplicationCreateTxn(sender, params, on_complete, \
                                            approval_program, clear_program, \
                                            global_schema, local_schema)

    # sign transaction
    signed_txn = txn.sign(private_key)
    tx_id = signed_txn.transaction.get_txid()

    # send transaction
    client.send_transactions([signed_txn])

    # wait for confirmation
    try:
        transaction_response = transaction.wait_for_confirmation(client, tx_id, 4)
        print("TXID: ", tx_id)
        print("Result confirmed in round: {}".format(transaction_response['confirmed-round']))
       
    except Exception as err:
        print(err)
        return

    # display results
    transaction_response = client.pending_transaction_info(tx_id)
    app_id = transaction_response['application-index']
    print("Created new app-id:", app_id)

    return app_id
```

This function is a simple example of creating an application creation transaction, which when submitted will deploy a smart contract. This example is very generic and can be used to deploy any smart contract. First, the creator’s address is resolved from the private key passed to the function, the transaction type is set to a NoOp application transaction, and the blockchain suggested parameters are retrieved from the connected node. These suggested parameters provide the default values that are required to submit a transaction, such as the expected fee for the transaction.

The Python SDK’s `ApplicationCreateTxn` function is called to create the transaction. This function takes the creator’s address, the approval and clear programs byte code, and a declaration of how much global and local state the smart contract will reserve. When creating a smart contract, the creation transaction has to specify how much state will be reserved. A contract can store up to 64 key-value pairs in global state and up to 16 key-value pairs per user who opts into the contract. Once these values are set, they can never be changed. The key is limited to 64 bytes. The key plus the value is limited to 128 bytes total. Using smaller keys to have more storage available for the value is possible. The keys are stored as byte slices (byte-array value) and the values are stored as either byte slices (byte-array value) or uint64s. More information on state values can be found in the [smart contract documentation](../smart-contracts/apps/index.md#modifying-state-in-smart-contract).

The passed-in private key is then used to sign the transaction and the ID of the transaction is retrieved. This ID is unique and can be used to look up the transaction later.

The transaction is then submitted to the connected node and the `wait_for_confirmation` SDK function is called to wait for the blockchain to process the transaction. Once the blockchain processes the transaction, a unique ID, called application ID, is returned for the smart contract. This can be used later to issue calls against the smart contract.

Now that all required functions are implemented, the main function can be created to deploy the contract.

```python
def main() :
    # initialize an algodClient
    algod_client = algod.AlgodClient(algod_token, algod_address)

    # define private keys
    creator_private_key = get_private_key_from_mnemonic(creator_mnemonic)

    # declare application state storage (immutable)
    local_ints = 0
    local_bytes = 0
    global_ints = 1
    global_bytes = 0
    global_schema = transaction.StateSchema(global_ints, global_bytes)
    local_schema = transaction.StateSchema(local_ints, local_bytes)

    # compile program to TEAL assembly
    with open("./approval.teal", "w") as f:
        approval_program_teal = approval_program()
        f.write(approval_program_teal)

    # compile program to TEAL assembly
    with open("./clear.teal", "w") as f:
        clear_state_program_teal = clear_state_program()
        f.write(clear_state_program_teal)

    # compile program to binary
    approval_program_compiled = compile_program(algod_client, approval_program_teal)

    # compile program to binary
    clear_state_program_compiled = compile_program(algod_client, clear_state_program_teal)

    print("--------------------------------------------")
    print("Deploying Counter application......")

    # create new application
    app_id = create_app(algod_client, creator_private_key, approval_program_compiled, clear_state_program_compiled, global_schema, local_schema)

    # read global state of application
    print("Global state:", read_global_state(algod_client, app_id))
```

First, a connection to the sandbox node is established. This is followed by recovering the account of the creator. Next, the amount of state to be used is defined. In this example, only one global integer is specified.

The SDK is then used to first convert the approval and clear programs to TEAL using the PyTeal library and both are written to local files. Each is then complied to byte code by the connected node. Finally, the smart contract is deployed using the `create_app` function created earlier and the current global state is then printed out for the contract. On deployment, this value will be set to 0.

## Calling the deployed smart contract

Now that the contract is deployed, the Add or Deduct functions can be called using a standard NoOp application transaction. The example created throughout this guide can be further modified to illustrate making a call to the smart contract.

To begin with, a function can be added to support calling the smart contract.

```python
# call application
def call_app(client, private_key, index, app_args) :
    # declare sender
    sender = account.address_from_private_key(private_key)

    # get node suggested parameters
    params = client.suggested_params()

    # create unsigned transaction
    txn = transaction.ApplicationNoOpTxn(sender, params, index, app_args)

    # sign transaction
    signed_txn = txn.sign(private_key)
    tx_id = signed_txn.transaction.get_txid()

    # send transaction
    client.send_transactions([signed_txn])

    # wait for confirmation
    try:
        transaction_response = transaction.wait_for_confirmation(client, tx_id, 5)
        print("TXID: ", tx_id)
        print("Result confirmed in round: {}".format(transaction_response['confirmed-round']))
       
    except Exception as err:
        print(err)
        return
    print("Application called")
```

This function operates similarly to the `create_app` function we defined earlier. In this case, we use the Python SDK’s `ApplicationNoOpTxn` function to create a standard NoOp application transaction. The address of the account sending the call is specified, followed by the network suggested parameters, the application id of the smart contract, and any arguments to the call. The arguments will be used to specify either the Add or Deduct methods.

The `main` function can then be modified to call the smart contract after deploying by adding the following to the bottom of the `main` function.

```python
    print("--------------------------------------------")
    print("Calling Counter application......")
    app_args = ["Add"]
    call_app(algod_client, creator_private_key, app_id, app_args)

    # read global state of application
    print("Global state:", read_global_state(algod_client, app_id))
```

In this example, the Add string is added to the application arguments array and the smart contract is called. The updated global state is then printed out. The value should now be set to 1.

The complete example is shown below.

```python
import base64

from algosdk.future import transaction
from algosdk import account, mnemonic, logic
from algosdk.v2client import algod
from pyteal import *

# user declared account mnemonics
creator_mnemonic = "finger rigid hat room course salmon say detect avocado assault awake sea public curious exit valve donkey tired escape dash drink diagram section absent cruise"
# user declared algod connection parameters. Node must have EnableDeveloperAPI set to true in its config
algod_address = "http://localhost:4001"
algod_token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

# helper function to compile program source
def compile_program(client, source_code):
    compile_response = client.compile(source_code)
    return base64.b64decode(compile_response['result'])

# helper function that converts a mnemonic passphrase into a private signing key
def get_private_key_from_mnemonic(mn) :
    private_key = mnemonic.to_private_key(mn)
    return private_key


# helper function that formats global state for printing
def format_state(state):
    formatted = {}
    for item in state:
        key = item['key']
        value = item['value']
        formatted_key = base64.b64decode(key).decode('utf-8')
        if value['type'] == 1:
            # byte string
            if formatted_key == 'voted':
                formatted_value = base64.b64decode(value['bytes']).decode('utf-8')
            else:
                formatted_value = value['bytes']
            formatted[formatted_key] = formatted_value
        else:
            # integer
            formatted[formatted_key] = value['uint']
    return formatted

# helper function to read app global state
def read_global_state(client, app_id):
    app = client.application_info(app_id)
    global_state = app['params']['global-state'] if "global-state" in app['params'] else []
    return format_state(global_state)


"""Basic Counter Application in PyTeal"""

def approval_program():
    on_creation = Seq([
        App.globalPut(Bytes("Count"), Int(0)),
        Return(Int(1))
    ])

    handle_optin = Return(Int(0))

    handle_closeout = Return(Int(0))

    handle_updateapp = Return(Int(0))

    handle_deleteapp = Return(Int(0))

    scratchCount = ScratchVar(TealType.uint64)

    add = Seq([
        scratchCount.store(App.globalGet(Bytes("Count"))),
        App.globalPut(Bytes("Count"), scratchCount.load() + Int(1)),
        Return(Int(1))
    ])

    deduct = Seq([
       scratchCount.store(App.globalGet(Bytes("Count"))),
        If(scratchCount.load() > Int(0),
            App.globalPut(Bytes("Count"), scratchCount.load() - Int(1)),
        ),
        Return(Int(1))
   ])

    handle_noop = Cond(
        [And(
            Global.group_size() == Int(1),
            Txn.application_args[0] == Bytes("Add")
        ), add],
        [And(
            Global.group_size() == Int(1),
            Txn.application_args[0] == Bytes("Deduct")
        ), deduct],
    )

    program = Cond(
        [Txn.application_id() == Int(0), on_creation],
        [Txn.on_completion() == OnComplete.OptIn, handle_optin],
        [Txn.on_completion() == OnComplete.CloseOut, handle_closeout],
        [Txn.on_completion() == OnComplete.UpdateApplication, handle_updateapp],
        [Txn.on_completion() == OnComplete.DeleteApplication, handle_deleteapp],
        [Txn.on_completion() == OnComplete.NoOp, handle_noop]
    )
    # Mode.Application specifies that this is a smart contract
    return compileTeal(program, Mode.Application, version=5)

def clear_state_program():
    program = Return(Int(1))
    # Mode.Application specifies that this is a smart contract
    return compileTeal(program, Mode.Application, version=5)


# create new application
def create_app(client, private_key, approval_program, clear_program, global_schema, local_schema):
    # define sender as creator
    sender = account.address_from_private_key(private_key)

    # declare on_complete as NoOp
    on_complete = transaction.OnComplete.NoOpOC.real

    # get node suggested parameters
    params = client.suggested_params()

    # create unsigned transaction
    txn = transaction.ApplicationCreateTxn(sender, params, on_complete, \
                                            approval_program, clear_program, \
                                            global_schema, local_schema)

    # sign transaction
    signed_txn = txn.sign(private_key)
    tx_id = signed_txn.transaction.get_txid()

    # send transaction
    client.send_transactions([signed_txn])

    # wait for confirmation
    try:
        transaction_response = transaction.wait_for_confirmation(client, tx_id, 5)
        print("TXID: ", tx_id)
        print("Result confirmed in round: {}".format(transaction_response['confirmed-round']))
       
    except Exception as err:
        print(err)
        return

    # display results
    transaction_response = client.pending_transaction_info(tx_id)
    app_id = transaction_response['application-index']
    print("Created new app-id:", app_id)

    return app_id


# call application
def call_app(client, private_key, index, app_args) :
    # declare sender
    sender = account.address_from_private_key(private_key)

    # get node suggested parameters
    params = client.suggested_params()

    # create unsigned transaction
    txn = transaction.ApplicationNoOpTxn(sender, params, index, app_args)

    # sign transaction
    signed_txn = txn.sign(private_key)
    tx_id = signed_txn.transaction.get_txid()

    # send transaction
    client.send_transactions([signed_txn])


    # wait for confirmation
    try:
        transaction_response = transaction.wait_for_confirmation(client, tx_id, 4)
        print("TXID: ", tx_id)
        print("Result confirmed in round: {}".format(transaction_response['confirmed-round']))
       
    except Exception as err:
        print(err)
        return
    print("Application called")

def main() :
    # initialize an algodClient
    algod_client = algod.AlgodClient(algod_token, algod_address)

    # define private keys
    creator_private_key = get_private_key_from_mnemonic(creator_mnemonic)

    # declare application state storage (immutable)
    local_ints = 0
    local_bytes = 0
    global_ints = 1
    global_bytes = 0
    global_schema = transaction.StateSchema(global_ints, global_bytes)
    local_schema = transaction.StateSchema(local_ints, local_bytes)

    # compile program to TEAL assembly
    with open("./approval.teal", "w") as f:
        approval_program_teal = approval_program()
        f.write(approval_program_teal)


    # compile program to TEAL assembly
    with open("./clear.teal", "w") as f:
        clear_state_program_teal = clear_state_program()
        f.write(clear_state_program_teal)

    # compile program to binary
    approval_program_compiled = compile_program(algod_client, approval_program_teal)

    # compile program to binary
    clear_state_program_compiled = compile_program(algod_client, clear_state_program_teal)

    print("--------------------------------------------")
    print("Deploying Counter application......")

    # create new application
    app_id = create_app(algod_client, creator_private_key, approval_program_compiled, clear_state_program_compiled, global_schema, local_schema)

    # read global state of application
    print("Global state:", read_global_state(algod_client, app_id))

    print("--------------------------------------------")
    print("Calling Counter application......")
    app_args = ["Add"]
    call_app(algod_client, creator_private_key, app_id, app_args)

    # read global state of application
    print("Global state:", read_global_state(algod_client, app_id))

main()

```

For more information on using the SDKs to deploy and interact with smart contracts see the [developer documentation](../smart-contracts/frontend/smartsigs.md).

# Building PyTeal smart signatures

Smart signatures are small programs that are submitted as part of a transaction and evaluated at submission time. These types of signatures can be used as an escrow-type of account or can be used to delegate a portion of the authority for a specific account.

When used as an escrow, they can hold Algos or Algorand assets (ASAs). When used this way any transaction can send Algos or ASAs to the escrow but the logic in the signature determines when value leaves the escrow. In this respect, they act very similarly to smart contracts, but the logic must be supplied with every transaction.

When used as a delegate, the logic can be signed by a specific account. The logic is then evaluated when a transaction is submitted from the signing account that is signed by the logic and not the private key of the sender. This is often used to allow restricted access to an account. For example, a mortgage company may provide logic to an account to remove a certain number of Algos from the account once a month. The user then signs this logic and once a month the mortgage company can submit a transaction from the signing account, but the transaction is signed by the smart signature and not the private key of the account.

Any time a smart signature is used the complete logic must be submitted as part of the transaction where the logic is used. The logic is recorded as part of the transaction but this is after the fact.

PyTeal supports building smart signatures in Python. For example, assume an escrow account is needed. This escrow can be funded by anyone but only a specific account is the beneficiary of the escrow and that account can withdraw funds at any time.

```python
#sample_smart_sig.py
from pyteal import *

"""Basic Donation Escrow"""
def donation_escrow(benefactor):
    Fee = Int(1000)

    #Only the benefactor account can withdraw from this escrow
    program = And(
        Txn.type_enum() == TxnType.Payment,
        Txn.fee() <= Fee,
        Txn.receiver() == Addr(benefactor),
        Global.group_size() == Int(1),
        Txn.rekey_to() == Global.zero_address()
    )
    # Mode.Signature specifies that this is a smart signature
    return compileTeal(program, Mode.Signature, version=5)
```

This is a very simplistic smart signature. The code for the complete signature is defined in the `donation_escrow` function. This function takes an Algorand address as a parameter. This address represents the beneficiary of the escrow. The entire program is a set of conditions anded together using the [`And` logical expression](https://pyteal.readthedocs.io/en/latest/api.html#pyteal.And). This expression takes two or more arguments that are logically anded and produces a 0 (logically false) or 1 (logically true). In this sample, a set of transaction fields are compared to expected values. The transaction type is first verified to be a payment transaction, the transaction fee is compared to make sure it is less than 1000 microAlgos, the transaction receiver is compared to the benefactor’s address, the group size is verified to guarantee that this transaction is not submitted with other transactions in a group, and the rekey field of the transaction is verified to be the zero address. The zero address is used to verify that the rekey field is not set. This prevents the escrow from being rekeyed to another account. This sample uses transaction fields and global properties. See the PyTeal documentation for additional [transaction fields](https://pyteal.readthedocs.io/en/latest/accessing_transaction_field.html?highlight=global#id1) and [global properties](https://pyteal.readthedocs.io/en/latest/accessing_transaction_field.html?highlight=global#global-parameters). The entire program is compiled to TEAL using the `compileTeal` PyTeal function. This function compiles the program as defined by the program variable. The `compileTeal` method also sets the Mode.Signature. This lets PyTeal know this is for a smart signature and not a smart contract. The version parameter instructs PyTeal on which version of TEAL to produce when compiling. 

To test this sample, a sample address can be defined and a print command calling the `donation_escrow` function can be added to the sample.

```python
#sample_smart_sig.py
from pyteal import *

"""Basic Donation Escrow"""
def donation_escrow(benefactor):
    Fee = Int(1000)

    #Only the benefactor account can withdraw from this escrow
    program = And(
        Txn.type_enum() == TxnType.Payment,
        Txn.fee() <= Fee,
        Txn.receiver() == Addr(benefactor),
        Global.group_size() == Int(1),
        Txn.rekey_to() == Global.zero_address()
    )
    # Mode.Signature specifies that this is a smart signature
    return compileTeal(program, Mode.Signature, version=5)

test_benefactor = "CZHGG36RBYTTK36N3ZC7MENGFOL3R6D4NNEJQU3G43U5GH457SU34ZGRLY"
print( donation_escrow(test_benefactor))
```

This sample can be executed using the following command.

```bash
python3 sample_smart_sig.py
```

This will print out the compiled TEAL. The Algorand address of the escrow can be retrieved by first saving the produced TEAL to a file and then compiled to byte code using the `goal` command-line tool. In the next section, using this smart signature with a transaction will be demonstrated.

```bash
$ python3 smart_sig.py > test.teal
$ ./goal clerk compile test.teal
test.teal: ZNJNTBMZKTCSO2RF4AJ3TLVFCZ5ZTHKAUBGR5AHJ23IHRFGK6GRIUVH2MU
```

# Deploying the smart signature

As stated in the previous section, smart signatures are used in conjunction with a transaction submission. In the previous section, a sample escrow was created. With escrows, any account can fund these accounts. These funds can not leave the escrow unless the logic evaluates to true. Once you have the escrow address, simple payment transactions can be used to fund it. To remove funds, a payment transaction can also be used but the transaction needs to be signed with the logic, not a private key. The following example illustrates:

- Compiling the escrow
- Funding the escrow with a simple payment transaction
- Dispensing funds using a payment transaction to the beneficiary signed with the logic

A few global variables are created and some utility functions are added to the previous section’s sample. The `benefactor_mnemonic` is the backup phrase for the address of the benefactor and the `sender_mnemonic` represents the account that will fund the escrow. Mnemonics should never be included in the source of a production environment. It is done here for learning purposes only. Key management should be handled by a proper wallet.

```python
import base64

from algosdk.future import transaction
from algosdk import mnemonic
from algosdk.v2client import algod
from pyteal import *

# user declared account mnemonics
benefactor_mnemonic = "REPLACE WITH YOUR OWN MNEMONIC"
sender_mnemonic = "REPLACE WITH YOUR OWN MNEMONIC"

# user declared algod connection parameters. Node must have EnableDeveloperAPI set to true in its config
algod_address = "http://localhost:4001"
algod_token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

# helper function to compile program source
def compile_smart_signature(client, source_code):
    compile_response = client.compile(source_code)
    return compile_response['result'], compile_response['hash']

# helper function that converts a mnemonic passphrase into a private signing key
def get_private_key_from_mnemonic(mn) :
    private_key = mnemonic.to_private_key(mn)
    return private_key



The `compile_smart_contract` and `get_private_key_from_mnemonic` functions are explained in [Deploying and calling a smart contract](#deploying-the-contract).

A utility function is then added to create and submit a simple payment transaction.

```python
def payment_transaction(creator_mnemonic, amt, rcv, algod_client)->dict:
    params = algod_client.suggested_params()
    add = mnemonic.to_public_key(creator_mnemonic)
    key = mnemonic.to_private_key(creator_mnemonic)
    unsigned_txn = transaction.PaymentTxn(add, params, rcv, amt)
    signed = unsigned_txn.sign(key)
    txid = algod_client.send_transaction(signed)

    # wait for confirmation
    try:
        pmtx = transaction.wait_for_confirmation(algod_client, txid, 5)
        print("TXID: ", txid)
        print("Result confirmed in round: {}".format(pmtx['confirmed-round']))
       
    except Exception as err:
        print(err)
        return
    return pmtx
```

This function takes a creator mnemonic of the address that is creating the payment transaction as the first parameter. The amount to send and the receiver of the payment transaction are the next two parameters. The final parameter is a connection to a valid Algorand node. In this example, the sandbox installed node is used.

In this function, the blockchain suggested parameters are retrieved from the connected node. These suggested parameters provide the default values that are required to submit a transaction, such as the expected fee for the transaction. The creator of the transaction’s address and private key are resolved from the mnemonic. The unsigned payment transaction is created using the Python SDK’s `PaymentTxn` method. This transaction is then signed with the recovered private key. As noted earlier, in a production application, the transaction should be signed by a valid wallet provider. The signed transaction is submitted to the node and the `wait_for_confirmation` SDK utility function is called, which will return when the transaction is finalized on the blockchain.

Another utility function is also added to create a payment transaction that is signed by the escrow logic. This function is very similar to the previous function.

```python
def lsig_payment_txn(escrowProg, escrow_address, amt, rcv, algod_client):
    params = algod_client.suggested_params()
    unsigned_txn = transaction.PaymentTxn(escrow_address, params, rcv, amt)
    encodedProg = escrowProg.encode()
    program = base64.decodebytes(encodedProg)
    lsig = transaction.LogicSigAccount(program)
    stxn = transaction.LogicSigTransaction(unsigned_txn, lsig)
    tx_id = algod_client.send_transaction(stxn)
    # wait for confirmation
    try:
        pmtx = transaction.wait_for_confirmation(algod_client, tx_id, 10)
        print("TXID: ", tx_id)
        print("Result confirmed in round: {}".format(pmtx['confirmed-round']))
       
    except Exception as err:
        print(err)
        return    
    return pmtx
```

The primary difference is that the function is passed the base64 encoded string of the compiled bytecode for the smart signature and the escrow’s Algorand address. The program is then converted to a byte array and the Python SDK’s `LogicSigAccount` function is used to create a logic signature from the program bytes. The payment transaction is then signed with the logic using the SDKs `LogicSigTransaction` function. For more information on Logic Signatures and smart signatures see the [smart signatures documentation](../smart-contracts/smartsigs/index.md).

The solution can be completed by adding a main function to put the utility functions to use.

```python
def main() :
    # initialize an algodClient
    algod_client = algod.AlgodClient(algod_token, algod_address)

    # define private keys
    receiver_public_key = mnemonic.to_public_key(benefactor_mnemonic)

    print("--------------------------------------------")
    print("Compiling Donation Smart Signature......")

    stateless_program_teal = donation_escrow(receiver_public_key)
    escrow_result, escrow_address= compile_smart_signature(algod_client, stateless_program_teal)

    print("Program:", escrow_result)
    print("hash: ", escrow_address)

    print("--------------------------------------------")
    print("Activating Donation Smart Signature......")

    # Activate escrow contract by sending 2 algo and 1000 microalgo for transaction fee from creator
    amt = 2001000
    payment_transaction(sender_mnemonic, amt, escrow_address, algod_client)

    print("--------------------------------------------")
    print("Withdraw from Donation Smart Signature......")

    # Withdraws 1 ALGO from smart signature using logic signature.
    withdrawal_amt = 1000000
    lsig_payment_txn(escrow_result, escrow_address, withdrawal_amt, receiver_public_key, algod_client)

```

The main function first makes a connection to the sandbox installed node, then the benefactor’s address is recovered. The `donation_escrow` built in the previous section is called to produce the TEAL for the smart signature. This TEAL is then compiled, returning both the base64 encoded bytes of the program and the address of the escrow.

A simple payment transaction is then created to fund the escrow with a little over 2 Algos. Finally, 1 Algo is dispensed from the escrow to the benefactor using a payment transaction signed by the smart signature. The complete example is shown below.

```python
import base64

from algosdk.future import transaction
from algosdk import mnemonic
from algosdk.v2client import algod
from pyteal import *

# user declared account mnemonics
benefactor_mnemonic = "REPLACE WITH YOUR OWN MNEMONIC"
sender_mnemonic = "REPLACE WITH YOUR OWN MNEMONIC"


# user declared algod connection parameters. Node must have EnableDeveloperAPI set to true in its config
algod_address = "http://localhost:4001"
algod_token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

# helper function to compile program source
def compile_smart_signature(client, source_code):
    compile_response = client.compile(source_code)
    return compile_response['result'], compile_response['hash']

# helper function that converts a mnemonic passphrase into a private signing key
def get_private_key_from_mnemonic(mn) :
    private_key = mnemonic.to_private_key(mn)
    return private_key



def payment_transaction(creator_mnemonic, amt, rcv, algod_client)->dict:
    params = algod_client.suggested_params()
    add = mnemonic.to_public_key(creator_mnemonic)
    key = mnemonic.to_private_key(creator_mnemonic)
    unsigned_txn = transaction.PaymentTxn(add, params, rcv, amt)
    signed = unsigned_txn.sign(key)
    txid = algod_client.send_transaction(signed)
    pmtx = transaction.wait_for_confirmation(algod_client, txid , 5)
    return pmtx

def lsig_payment_txn(escrowProg, escrow_address, amt, rcv, algod_client):
    params = algod_client.suggested_params()
    unsigned_txn = transaction.PaymentTxn(escrow_address, params, rcv, amt)
    encodedProg = escrowProg.encode()
    program = base64.decodebytes(encodedProg)
    lsig = transaction.LogicSigAccount(program)
    stxn = transaction.LogicSigTransaction(unsigned_txn, lsig)
    tx_id = algod_client.send_transaction(stxn)
    pmtx = transaction.wait_for_confirmation(algod_client, tx_id, 10)
    return pmtx

"""Basic Donation Escrow"""

def donation_escrow(benefactor):
    Fee = Int(1000)

    #Only the benefactor account can withdraw from this escrow
    program = And(
        Txn.type_enum() == TxnType.Payment,
        Txn.fee() <= Fee,
        Txn.receiver() == Addr(benefactor),
        Global.group_size() == Int(1),
        Txn.rekey_to() == Global.zero_address()
    )

    # Mode.Signature specifies that this is a smart signature
    return compileTeal(program, Mode.Signature, version=5)

def main() :
    # initialize an algodClient
    algod_client = algod.AlgodClient(algod_token, algod_address)

    # define private keys
    receiver_public_key = mnemonic.to_public_key(benefactor_mnemonic)

    print("--------------------------------------------")
    print("Compiling Donation Smart Signature......")

    stateless_program_teal = donation_escrow(receiver_public_key)
    escrow_result, escrow_address= compile_smart_signature(algod_client, stateless_program_teal)

    print("Program:", escrow_result)
    print("hash: ", escrow_address)

    print("--------------------------------------------")
    print("Activating Donation Smart Signature......")

    # Activate escrow contract by sending 2 algo and 1000 microalgo for transaction fee from creator
    amt = 2001000
    payment_transaction(sender_mnemonic, amt, escrow_address, algod_client)

    print("--------------------------------------------")
    print("Withdraw from Donation Smart Signature......")

    # Withdraws 1 ALGO from smart signature using logic signature.
    withdrawal_amt = 1000000
    lsig_payment_txn(escrow_result, escrow_address, withdrawal_amt, receiver_public_key, algod_client)

main()

```

For more information on smart signatures, see the [developer documentation](../smart-contracts/smartsigs/index.md).
