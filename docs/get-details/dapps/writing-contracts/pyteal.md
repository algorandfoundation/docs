title: PyTeal

!!! note
For a native Python experience, checkout our [Algorand Python](https://algorandfoundation.github.io/puya/) docs.

[PyTeal](https://github.com/algorand/pyteal) is a python library for generating [TEAL](/docs/get-details/dapps/avm/teal/) programs that provides a convenient and familiar syntax. 

Complete installation instructions and developer guides are available in the [PyTeal documentation](https://pyteal.readthedocs.io/en/latest/).



# Quick start videos

If you prefer videos, take a look at this playlist to learn about PyTeal. Most of the videos in the list are under 7 minutes each.

<iframe width="100%" style="aspect-ratio:16/9" src="https://www.youtube-nocookie.com/embed/yEFUv760I8A" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>

# PyTeal overview

This section assumes the reader is familiar with [Smart Contracts](/docs/get-details/dapps/smart-contracts/#smart-contracts) and [Smart Signatures](/docs/get-details/dapps/smart-contracts/#smart-signatures).

When building a dApp that makes use of smart contracts or smart signatures (smartsigs), PyTeal makes implementation of more complex logic much simpler than writing the TEAL manually. 

Generally, developers install PyTeal, write the contract in Python using their preferred editor, and then use PyTeal’s `compileProgram` method to produce the TEAL code. The TEAL source can be compiled into bytecode and deployed to the blockchain.

For most applications, these contracts will only be a portion of the dApp’s architecture. Typically, developers will build functionality in the dApp that resides on the blockchain and some front end to interact with the smart contracts. 

The PyTeal contracts can be pre-compiled into TEAL and used directly or the logic to generate TEAL dynamically may be integrated into the dApp front end workflow. For example, a front end may provide something like an exchange that allows limit orders to be created based on a template and then deployed once a user opens an order. In this case, the complete limit order may be implemented as part of a smart contract that is deployed when the order is opened by the dApp. 

# Building PyTeal smart contracts

On Algorand, smart contracts are small programs that are evaluated when the contract is deployed or called. These contracts make use of a set of functions [(opcodes)](../avm/teal/opcodes) to be evaluated against the context it was called with. This context includes the current state of its storage values, the transactions in the group, any arguments and references to accounts, assets, or other applications.

Besides evaluating the logic to approve or reject a transaction, the contracts may cause side effects. Side effects a contract may produce include changes to an Applications global or local state, or producing their own [transactions](../../transactions/transactions.md). These transactions are [atomic](../../atomic_transfers.md) with the outer group the application was invoked with.  The transactions produced during contract evaluation are by default sent from the Applications associated account.

For more information on smart contracts, see the [smart contract documentation](../smart-contracts/apps/index.md).


## Hello PyTeal

When building smart contracts in PyTeal it is important to realize that a smart contract actually consists of two programs. These are called the approval and the clear programs. In PyTeal both of these programs are generally created in the same Python file. So the beginning of a PyTeal program will contain logic similar to the following:

<!-- ===PYTEAL_APP_EMPTY_LOGIC=== -->
```python
    def approval_program():
        program = Return(Int(1))
        # Mode.Application specifies that this is a smart contract
        return compileTeal(program, Mode.Application, version=5)

    def clear_state_program():
        program = Return(Int(1))
        # Mode.Application specifies that this is a smart contract
        return compileTeal(program, Mode.Application, version=5)

    print(approval_program())
    print(clear_state_program())
```
[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/app_walkthrough.py#L27-L39)
<!-- ===PYTEAL_APP_EMPTY_LOGIC=== -->

In the above example, a function is defined to return each of the two programs. The `compileTeal` function compiles the program set to the `program` variable. In this case, we are just returning the integer 1 for both programs. The other arguments passed to `compileTeal` set the mode (Application or Signature) and version of TEAL to produce. The output from calling the `compileTeal` method is a string representing the compiled TEAL source.

Running the program, the output should look something like:
```teal
#pragma version 5
int 1
return
```

This output represents the TEAL source for the program. The TEAL source must be compiled to bytecode in order to make use of it on-chain.  This can be done using [goal](../../../clis/goal/clerk/compile.md) or one of the SDKs through the [REST](../../../rest-apis/algod.md#post-v2tealcompile) interface.


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

While _only_ PyTeal Expressions may be included in the Expression tree representing the program, the Expressions may be generated as part of running the Python program. This allows the author to make use of Python scripting logic to generate a PyTeal Expression tree based on more complex logic. 


## Writing a simple PyTeal Contract

Here we'll start updating our example to allow a more complex logical flow, showing how to build up a PyTeal contract using PyTeal Expressions.

Below is a simple logical switch statement used to route evaluation to different set of logic based on a Transaction's [`OnComplete`](../../transactions/transactions.md#application-call-transaction) value.  

<!-- ===PYTEAL_APP_MANUAL_ROUTER=== -->
```python
    def approval_program():

        handle_creation = Approve()
        handle_optin = Reject()
        handle_closeout = Reject()
        handle_update = Reject()
        handle_delete = Reject()
        handle_noop = Reject()

        program = Cond(
            [Txn.application_id() == Int(0), handle_creation],
            [Txn.on_completion() == OnComplete.OptIn, handle_optin],
            [Txn.on_completion() == OnComplete.CloseOut, handle_closeout],
            [Txn.on_completion() == OnComplete.UpdateApplication, handle_update],
            [Txn.on_completion() == OnComplete.DeleteApplication, handle_delete],
            [Txn.on_completion() == OnComplete.NoOp, handle_noop],
        )
        return program

    return compileTeal(approval_program(), Mode.Application, version=5)
```
[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/app_walkthrough.py#L44-L64)
<!-- ===PYTEAL_APP_MANUAL_ROUTER=== -->

The `program` variable is set to a [PyTeal `Cond` Expression](https://pyteal.readthedocs.io/en/latest/control_structures.html?highlight=Cond#chaining-tests-cond) which is an example of a Control Flow statement. Other Control flow statements are documented in the [PyTeal documentation](https://pyteal.readthedocs.io/en/latest/control_structures.html?highlight=seq#).

A `Cond` expression allows several conditions to be evaluated in order, taking a number of arguments as [`Condition Expression`, `Body Expression`].  The `Condition Expression` must evaluate to True or False. For a given invocation, the first Condition that evaluates True will pass flow of the program to its corresponding `Body Expression`.  If none of the conditions are true the smart contract will return an `err` and fail.  The body for each condition here should reference some other variable or method defined in your PyTeal contract.

!!! note
    For control flow statements like `Cond`, the Body Expression of each condition _MUST_ evaluate to the same type. The current types that are allowed are `none`, `any`, `uint64` or `bytes`

In the above example, most of the conditions check the transaction type using the `on_completion` transaction field. The result of a statement like `Txn.on_completion() == OnComplete.NoOp` is itself a PyTeal Expression, in this case evaluating to True or False.  

!!! note
    Using `Txn` is shorthand for `The transaction that invoked this contract` and should be used inspect transaction fields. Other transactions in the group including the one referenced by `Txn` are available using `Gtxn[n]` where n is the index of the Transaction in the group.  All other transaction fields can be examined using these transaction references, see the [PyTeal documentation](https://pyteal.readthedocs.io/en/latest/accessing_transaction_field.html) for a complete list.

The only condition above that does _not_ check the `on_completion` field is the first one, which only checks the `application_id` field. An `application_id` of `0` on the `Txn` tells us that this is meant to Create the Application. If the program succeeds, it will be assigned an application ID and further invocations will use that ID to call it. In this example the `handle_creation` Expression should handle any initialization necessary for this contract.

### First Router 

While the above method of constructing distinct Expression trees for both the approval and clear state programs works, it is often preferable to use the [Router](https://pyteal.readthedocs.io/en/stable/abi.html#registering-methods) class provided in `PyTeal`. The `Router` class abstracts much of the handling for method routing when constructing ABI compliant applications. This is especially useful for encoding and decoding ABI types and returning ABI types.

<!-- ===PYTEAL_APP_ROUTER_BASIC=== -->
```python
## Contract logic
count_key = Bytes("Count")

# Create an expression to store 0 in the `Count` global variable and return 1
handle_creation = Seq(App.globalPut(count_key, Int(0)), Approve())

# Main router class
router = Router(
    # Name of the contract
    "my-first-router",
    # What to do for each on-complete type when no arguments are passed (bare call)
    BareCallActions(
        # On create only, just approve
        no_op=OnCompleteAction.create_only(handle_creation),
        # Always let creator update/delete but only by the creator of this contract
        update_application=OnCompleteAction.always(Reject()),
        delete_application=OnCompleteAction.always(Reject()),
        # No local state, don't bother handling it.
        close_out=OnCompleteAction.never(),
        opt_in=OnCompleteAction.never(),
    ),
)
```
[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/simple_router.py#L81-L103)
<!-- ===PYTEAL_APP_ROUTER_BASIC=== -->

Here we defined the `handle_creation` variable to be a Sequence of Expressions using [`Seq`](https://pyteal.readthedocs.io/en/latest/control_structures.html?highlight=seq#chaining-expressions-seq).

The first expression stores a global variable named Count, and its value is set to 0. More information about storing state variables is available in the [PyTeal documentation](https://pyteal.readthedocs.io/en/latest/control_structures.html?highlight=seq#chaining-expressions-seq).

This second expression is just `Approve()` which is an alias for `Return(Int(1))` that gets passed back to the caller. 

The Router class is initialized with a string name (will be important below in [abi.json](#abi-specification)) and a set of BareCallActions. A Bare Call is an application call transaction with 0 application arguments.

Each OnComplete type may have an associated OnCompleteAction, the options are:

    - `create_only`: Only triggered on the initial creation of the application (Txn.application_id() == Int(0))
    - `call_only`: Triggered for any application calls that are not the initial creation (Txn.application_id() != Int(0))
    - `always`: Triggered any time this on complete is used Or(Txn.application_id() == Int(0), Txn.application_id() != Int(0))
    - `never`: Never triggered, application should reject the transaction if this OnComplete is specified in the transaction. 

The OnCompleteAction is passed an Expression to evaluate when it is triggered and it should `Return` something. 

In the first case we `Return` the result of the `handle_creation` expression. For update and delete we `Reject`, this is an alias for `Return(Int(0))`

To summarize, when this specific smart contract is first deployed it will store a global variable named Count with a value of 0 and return success. 

### Handle Method Calls 

Typically the OnCompletion value is set to `NoOp` to make calls to any application. Different logic handling is often achieved by inspecting other transaction fields, especially `application_args`.

Since our example will be an ARC4 contract, the application_args first argument will be the method selector but the router handles that for us.

Our example requires an increment function and a decrement function to be handle incrementing and decrementing the counter. 

<!-- ===PYTEAL_APP_ROUTER_METHODS=== -->
```python
@router.method
def increment():
    # Declare the ScratchVar as a Python variable _outside_ the expression tree
    scratchCount = ScratchVar(TealType.uint64)
    return Seq(
        Assert(Global.group_size() == Int(1)),
        # The initial `store` for the scratch var sets the value to
        # whatever is in the `Count` global state variable
        scratchCount.store(App.globalGet(count_key)),
        # Increment the value stored in the scratch var
        # and update the global state variable
        App.globalPut(count_key, scratchCount.load() + Int(1)),
    )


@router.method
def decrement():
    # Declare the ScratchVar as a Python variable _outside_ the expression tree
    scratchCount = ScratchVar(TealType.uint64)
    return Seq(
        Assert(Global.group_size() == Int(1)),
        # The initial `store` for the scratch var sets the value to
        # whatever is in the `Count` global state variable
        scratchCount.store(App.globalGet(count_key)),
        # Check if the value would be negative by decrementing
        If(
            scratchCount.load() > Int(0),
            # If the value is > 0, decrement the value stored
            # in the scratch var and update the global state variable
            App.globalPut(count_key, scratchCount.load() - Int(1)),
        ),
    )


```
[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/simple_router.py#L117-L151)
<!-- ===PYTEAL_APP_ROUTER_METHODS=== -->

In the example, we've implemented the `increment` and `decrement` methods as python methods that return a PyTeal `Expression`. They are attached to the router using the `@router.method` decorator which handles determining the method selector that should be used for routing when an application call is made.

The first line of both handlers creates a ScratchVar to be used later. Critically, the declaration of this `ScratchVar` happens _outside_ the Expression tree since the statement `scratchCount = ScratchVar()` is _not_ a valid PyTeal Expression. Smart contracts can hold up to 256 temporary variables in scratch space.  The scratch variable in this example happens to be an integer, byte arrays can also be stored. 

!!! note
    The `walrus operator` is very handy for these types of things. The above could have been `Seq( (sv := ScratchVar()).store(Int(0)), ...)` but would have complicated the example.

The first expression in the Sequence for both handlers is an `Assert` that will immediately Reject the Transaction if, in this case, the number of transactions in the group is not exactly 1. The `Global.group_size()` illustrates the use of a PyTeal global variables. See the [PyTeal documentation](https://pyteal.readthedocs.io/en/latest/accessing_transaction_field.html?highlight=global#global-parameters) for other global variables available.


Next, the current value of the global variable Count is read for the contract and placed in scratch space. 

Then, the contract either increments this number or decrements and then stores the result into the contract’s global variable. 

In the `decrement` method, an additional `If` expression is used to verify the current global variable is above 0. 

Finally, both methods exit the smart contract call. Assuming no errors were thrown, the `Router` will return 1, indicating success.

### Final product

Because no opt-in is allowed, a clear program need not do anything so we simply return 1, indicating success. The `Router` can now be used to compile our application to the `TEAL` programs and provide a `Contract` object to allow clients to call the `ABI` methods.

<!-- ===PYTEAL_APP_ROUTER_COMPILE=== -->
```python
# Compile the program
approval_program, clear_program, contract = router.compile_program(version=6)

# print out the results
print(approval_program)
print(clear_program)

print(json.dumps(contract.dictify()))
```
[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/simple_router.py#L106-L114)
<!-- ===PYTEAL_APP_ROUTER_COMPILE=== -->

The last bit to add is the `router.compile_program` which compiles the PyTeal into TEAL and returns the `approval` program, the `clear` program, and even the Python SDK `contract` object that can be used to make method calls or written to a file and shared.

### ABI specification

The contract as json can be be shared with callers and loaded into the SDKs, see the [abi](/docs/get-details/dapps/smart-contracts/ABI/) page for more. 

It will look line this:
```json
{
  "name": "my-first-router",
  "methods": [
    { "name": "increment", "args": [], "returns": { "type": "void" } },
    { "name": "decrement", "args": [], "returns": { "type": "void" } }
  ],
  "desc": null,
  "networks": {}
}
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

<!-- ===PYTEAL_APP_ROUTER_SETUP=== -->
```python
# user declared account mnemonics
creator_mnemonic = "TODO: Add your mnemonic"
# user declared algod connection parameters. Node must have EnableDeveloperAPI set to true in its config
algod_address = "http://localhost:4001"
algod_token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
```
[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/simple_router.py#L28-L33)
<!-- ===PYTEAL_APP_ROUTER_SETUP=== -->

The first is a creator mnemonic. This mnemonic is used to recover the private key for the funded account that will own and create the smart contract. 

!!!warning
    Placing a mnemonic like this in code should never be done in production. Typically applications will link to some protected wallet to sign transactions. Some examples of wallets are Pera wallet (mobile and web) and Exodus (web extension). Pera, along with some other wallets in the ecosystem, allow transcations be signed using the Wallet Connect API. The mnemonic is used here for learning purposes only.

The algod_address and algod_token values are the default values to connect to a sandbox installed node. Also note that in this example, the sandbox node is connected to the Algorand TestNet network (eg `./sandbox up testnet`).

In this example, the TEAL for the smart contract will be compiled programmatically by the node. The `EnableDeveloperAPI` configuration parameter must be set to `true` in the node’s configuration to allow this operation. For more information on changing node configuration parameters, see the [developer documentation](../../../run-a-node/reference/config.md). When using the sandbox install, this value is automatically set to true.

Next, a few helper functions need to be added to the sample.

<!-- ===PYTEAL_APP_ROUTER_HELPERS=== -->
```python
# helper function to compile program source
def compile_program(client: algod.AlgodClient, source_code: str):
    compile_response = client.compile(source_code)
    return base64.b64decode(compile_response["result"])


# helper function that converts a mnemonic passphrase into a private signing key
def get_private_key_from_mnemonic(mn: str):
    private_key = mnemonic.to_private_key(mn)
    return private_key


# helper function that formats global state for printing
def format_state(state: List[Dict[str, Any]]):
    formatted = {}
    for item in state:
        key = item["key"]
        value = item["value"]
        formatted_key = base64.b64decode(key).decode("utf-8")
        if value["type"] == 1:
            # byte string
            if formatted_key == "voted":
                formatted_value = base64.b64decode(value["bytes"]).decode("utf-8")
            else:
                formatted_value = value["bytes"]
            formatted[formatted_key] = formatted_value
        else:
            # integer
            formatted[formatted_key] = value["uint"]
    return formatted


# helper function to read app global state
def read_global_state(client: algod.AlgodClient, app_id: int):
    app = client.application_info(app_id)
    global_state = (
        app["params"]["global-state"] if "global-state" in app["params"] else []
    )
    return format_state(global_state)


```
[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/simple_router.py#L36-L77)
<!-- ===PYTEAL_APP_ROUTER_HELPERS=== -->

The `compile_program` function is a utility function that allows passing the generated TEAL code to a node that will compile and return the byte code. This returned byte code will be used with the application creation transaction (deploying the contract) later.

The `get_private_key_from_mnemonic` function is a utility function that takes a mnemonic (account backup phrase) and returns the private key of the specific account. This will be used in this sample to recover the private key of the funded account of the smart contract creator.

The `wait_for_confirmation` SDK function is a utility function that when called will wait until a specific transaction is confirmed on the Algorand blockchain. This will be used to confirm that the application creation transaction is successful and the smart contract is actively deployed.

As the sample smart contract manipulates global variables, a couple of helper functions are needed to display the contents of these values.


Global variables for smart contracts are actually stored in the creator account’s ledger entry on the blockchain. The location is referred to as global state and the SDKs provide a function to retrieve the application data including the global state. In this example, the function `read_global_state` uses the Python SDK function `application_info` to connect to the Algorand node and retrieve the application information. The function then extracts the global state values if they exist, otherwise returns an empty array. The `format_state` function takes the application data and formats the values for display. For more information on global and local state see the [smart contract documentation](../smart-contracts/apps/index.md).

As covered earlier in this guide, to deploy the contract an application creation transaction must be created and submitted to the blockchain. The SDKs provide a method for creating this transaction. The following code illustrates creating and submitting this transaction.

<!-- ===PYTEAL_APP_ROUTER_CREATOR=== -->
```python
# create new application
def create_app(
    client: algod.AlgodClient,
    private_key: str,
    approval_program: bytes,
    clear_program: bytes,
    global_schema: transaction.StateSchema,
    local_schema: transaction.StateSchema,
) -> int:
    # define sender as creator
    sender = account.address_from_private_key(private_key)

    # declare on_complete as NoOp
    on_complete = transaction.OnComplete.NoOpOC.real

    # get node suggested parameters
    params = client.suggested_params()

    # create unsigned transaction
    txn = transaction.ApplicationCreateTxn(
        sender,
        params,
        on_complete,
        approval_program,
        clear_program,
        global_schema,
        local_schema,
    )

    # sign transaction
    signed_txn = txn.sign(private_key)
    tx_id = signed_txn.transaction.get_txid()

    # send transaction
    client.send_transactions([signed_txn])

    # wait for confirmation
    try:
        transaction_response = transaction.wait_for_confirmation(client, tx_id, 5)
        print("TXID: ", tx_id)
        print(
            "Result confirmed in round: {}".format(
                transaction_response["confirmed-round"]
            )
        )
    except Exception as err:
        print(err)
        return 0

    # display results
    transaction_response = client.pending_transaction_info(tx_id)
    app_id = transaction_response["application-index"]
    print("Created new app-id:", app_id)

    return app_id


```
[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/simple_router.py#L155-L212)
<!-- ===PYTEAL_APP_ROUTER_CREATOR=== -->

This function is a simple example of creating an application creation transaction, which when submitted will deploy a smart contract. This example is very generic and can be used to deploy any smart contract. First, the creator’s address is resolved from the private key passed to the function, the transaction type is set to a NoOp application transaction, and the blockchain suggested parameters are retrieved from the connected node. These suggested parameters provide the default values that are required to submit a transaction, such as the expected fee for the transaction.

The Python SDK’s `ApplicationCreateTxn` function is called to create the transaction. This function takes the creator’s address, the approval and clear programs byte code, and a declaration of how much global and local state the smart contract will reserve. When creating a smart contract, the creation transaction has to specify how much state will be reserved. A contract can store up to 64 key-value pairs in global state and up to 16 key-value pairs per user who opts into the contract. Once these values are set, they can never be changed. The key is limited to 64 bytes. The key plus the value is limited to 128 bytes total. Using smaller keys to have more storage available for the value is possible. The keys are stored as byte slices (byte-array value) and the values are stored as either byte slices (byte-array value) or uint64s. More information on state values can be found in the [smart contract documentation](../smart-contracts/apps/index.md#modifying-state-in-smart-contract).

The passed-in private key is then used to sign the transaction and the ID of the transaction is retrieved. This ID is unique and can be used to look up the transaction later.

The transaction is then submitted to the connected node and the `wait_for_confirmation` SDK function is called to wait for the blockchain to process the transaction. Once the blockchain processes the transaction, a unique ID, called application ID, is returned for the smart contract. This can be used later to issue calls against the smart contract.

Now that all required functions are implemented, the main function can be created to deploy the contract.

<!-- ===PYTEAL_APP_ROUTER_DEPLOY=== -->
```python
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

    # Compile the program
    approval_program, clear_program, contract = router.compile_program(version=6)

    with open("./approval.teal", "w") as f:
        f.write(approval_program)

    with open("./clear.teal", "w") as f:
        f.write(clear_program)

    with open("./contract.json", "w") as f:
        import json

        f.write(json.dumps(contract.dictify()))

    # compile program to binary
    approval_program_compiled = compile_program(algod_client, approval_program)

    # compile program to binary
    clear_state_program_compiled = compile_program(algod_client, clear_program)

    print("Deploying Counter application......")

    # create new application
    app_id = create_app(
        algod_client,
        creator_private_key,
        approval_program_compiled,
        clear_state_program_compiled,
        global_schema,
        local_schema,
    )

    # read global state of application
    print("Global state:", read_global_state(algod_client, app_id))
```
[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/simple_router.py#L250-L298)
<!-- ===PYTEAL_APP_ROUTER_DEPLOY=== -->

First, a connection to the sandbox node is established. This is followed by recovering the account of the creator. Next, the amount of state to be used is defined. In this example, only one global integer is specified.

The SDK is then used to first convert the approval and clear programs to TEAL using the PyTeal library and both are written to local files. Each is then complied to byte code by the connected node. Finally, the smart contract is deployed using the `create_app` function created earlier and the current global state is then printed out for the contract. On deployment, this value will be set to 0.

## Calling the deployed smart contract

Now that the contract is deployed, the Add or Deduct functions can be called using a standard NoOp application transaction. The example created throughout this guide can be further modified to illustrate making a call to the smart contract.

To begin with, a function can be added to support calling the smart contract.

<!-- ===PYTEAL_APP_ROUTER_CALLER=== -->
```python
# call application
def call_app(client, private_key, index, contract):
    # get sender address
    sender = account.address_from_private_key(private_key)
    # create a Signer object
    signer = AccountTransactionSigner(private_key)

    # get node suggested parameters
    sp = client.suggested_params()

    # Create an instance of AtomicTransactionComposer
    atc = AtomicTransactionComposer()
    atc.add_method_call(
        app_id=index,
        method=contract.get_method_by_name("increment"),
        sender=sender,
        sp=sp,
        signer=signer,
        method_args=[],  # No method args needed here
    )

    # send transaction
    results = atc.execute(client, 2)

    # wait for confirmation
    print("TXID: ", results.tx_ids[0])
    print("Result confirmed in round: {}".format(results.confirmed_round))


```
[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/simple_router.py#L216-L245)
<!-- ===PYTEAL_APP_ROUTER_CALLER=== -->

This function constructs an AtomicTransactionComposer to handle adding the appropriate arguments. The only argument we need to pass in this case is the method selector for the method we wish to call (`increment`) but it is added automatically to the application args of the created transaction. Any arguments that the method specifies in the contract would be passed to the `method_args` array.

The `main` function can then be modified to call the smart contract after deploying by adding the following to the bottom of the `main` function.

<!-- ===PYTEAL_APP_ROUTER_CALL=== -->
```python
    print("Calling Counter application......")
    call_app(algod_client, creator_private_key, app_id, contract)

    # read global state of application
    print("Global state:", read_global_state(algod_client, app_id))
```
[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/simple_router.py#L301-L306)
<!-- ===PYTEAL_APP_CALL=== -->

In this example, the Add string is added to the application arguments array and the smart contract is called. The updated global state is then printed out. The value should now be set to 1.

For more information on using the SDKs to deploy and interact with smart contracts see the [developer documentation](../smart-contracts/frontend/smartsigs.md).

# Building PyTeal smart signatures

Smart signatures are small programs that are submitted as part of a transaction and evaluated at submission time. These types of signatures can be used as an escrow-type of account or can be used to delegate a portion of the authority for a specific account.

When used as an escrow, they can hold Algos or Algorand assets (ASAs). When used this way any transaction can send Algos or ASAs to the escrow but the logic in the signature determines when value leaves the escrow. In this respect, they act very similarly to smart contracts, but the logic must be supplied with every transaction.

When used as a delegate, the logic can be signed by a specific account. The logic is then evaluated when a transaction is submitted from the signing account that is signed by the logic and not the private key of the sender. This is often used to allow restricted access to an account. For example, a mortgage company may provide logic to an account to remove a certain number of Algos from the account once a month. The user then signs this logic and once a month the mortgage company can submit a transaction from the signing account, but the transaction is signed by the smart signature and not the private key of the account.

Any time a smart signature is used the complete logic must be submitted as part of the transaction where the logic is used. The logic is recorded as part of the transaction but this is after the fact.

PyTeal supports building smart signatures in Python. For example, assume an escrow account is needed. This escrow can be funded by anyone but only a specific account is the beneficiary of the escrow and that account can withdraw funds at any time.

<!-- ===PYTEAL_LSIG_SIMPLE_ESCROW=== -->
```python
def donation_escrow(benefactor):
    Fee = Int(1000)

    # Only the benefactor account can withdraw from this escrow
    program = And(
        Txn.type_enum() == TxnType.Payment,
        Txn.fee() <= Fee,
        Txn.receiver() == Addr(benefactor),
        Global.group_size() == Int(1),
        Txn.rekey_to() == Global.zero_address(),
    )

    # Mode.Signature specifies that this is a smart signature
    return compileTeal(program, Mode.Signature, version=5)


```
[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/lsig.py#L7-L23)
<!-- ===PYTEAL_LSIG_SIMPLE_ESCROW=== -->

This is a very simplistic smart signature. The code for the complete signature is defined in the `donation_escrow` function. This function takes an Algorand address as a parameter. This address represents the beneficiary of the escrow. The entire program is a set of conditions anded together using the [`And` logical expression](https://pyteal.readthedocs.io/en/latest/api.html#pyteal.And). This expression takes two or more arguments that are logically anded and produces a 0 (logically false) or 1 (logically true). In this sample, a set of transaction fields are compared to expected values. The transaction type is first verified to be a payment transaction, the transaction fee is compared to make sure it is less than 1000 microAlgos, the transaction receiver is compared to the benefactor’s address, the group size is verified to guarantee that this transaction is not submitted with other transactions in a group, and the rekey field of the transaction is verified to be the zero address. The zero address is used to verify that the rekey field is not set. This prevents the escrow from being rekeyed to another account. This sample uses transaction fields and global properties. See the PyTeal documentation for additional [transaction fields](https://pyteal.readthedocs.io/en/latest/accessing_transaction_field.html?highlight=global#id1) and [global properties](https://pyteal.readthedocs.io/en/latest/accessing_transaction_field.html?highlight=global#global-parameters). The entire program is compiled to TEAL using the `compileTeal` PyTeal function. This function compiles the program as defined by the program variable. The `compileTeal` method also sets the Mode.Signature. This lets PyTeal know this is for a smart signature and not a smart contract. The version parameter instructs PyTeal on which version of program version to produce when compiling. 

To test this sample, a sample address can be defined and a print command calling the `donation_escrow` function can be added to the sample.

<!-- ===PYTEAL_LSIG_SIMPLE_ESCROW_INIT=== -->
```python
rando_addr, _ = account.generate_account()
teal_program = donation_escrow(rando_addr)
```
[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/lsig.py#L26-L28)
<!-- ===PYTEAL_LSIG_SIMPLE_ESCROW_INIT=== -->


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

<!-- ===PYTEAL_LSIG_SIMPLE_SETUP=== -->
```python
# user declared account mnemonics
benefactor_mnemonic = "REPLACE WITH YOUR OWN MNEMONIC"
sender_mnemonic = "REPLACE WITH YOUR OWN MNEMONIC"


# user declared algod connection parameters. Node must have EnableDeveloperAPI set to true in its config
algod_address = "http://localhost:4001"
algod_token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
```
[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/lsig.py#L32-L40)
<!-- ===PYTEAL_LSIG_SIMPLE_SETUP=== -->

<!-- ===PYTEAL_LSIG_SIMPLE_HELPERS=== -->
```python
# helper function to compile program source
def compile_smart_signature(
    client: algod.AlgodClient, source_code: str
) -> Tuple[str, str]:
    compile_response = client.compile(source_code)
    return compile_response["result"], compile_response["hash"]


```
[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/lsig.py#L43-L51)
<!-- ===PYTEAL_LSIG_SIMPLE_HELPERS=== -->


The `compile_smart_contract` and `get_private_key_from_mnemonic` functions are explained in [Deploying and calling a smart contract](#deploying-the-contract).

A utility function is then added to create and submit a simple payment transaction.

<!-- ===PYTEAL_LSIG_SIMPLE_SEED_PAYMENT=== -->
```python
def payment_transaction(
    creator_mnemonic: str, amt: int, rcv: str, algod_client: algod.AlgodClient
) -> dict:
    creator_pk = mnemonic.to_private_key(creator_mnemonic)
    creator_address = account.address_from_private_key(creator_pk)

    params = algod_client.suggested_params()
    unsigned_txn = transaction.PaymentTxn(creator_address, params, rcv, amt)
    signed = unsigned_txn.sign(creator_pk)

    txid = algod_client.send_transaction(signed)
    pmtx = transaction.wait_for_confirmation(algod_client, txid, 5)
    return pmtx


```
[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/lsig.py#L55-L70)
<!-- ===PYTEAL_LSIG_SIMPLE_SEED_PAYMENT=== -->

This function takes a creator mnemonic of the address that is creating the payment transaction as the first parameter. The amount to send and the receiver of the payment transaction are the next two parameters. The final parameter is a connection to a valid Algorand node. In this example, the sandbox installed node is used.

In this function, the blockchain suggested parameters are retrieved from the connected node. These suggested parameters provide the default values that are required to submit a transaction, such as the expected fee for the transaction. The creator of the transaction’s address and private key are resolved from the mnemonic. The unsigned payment transaction is created using the Python SDK’s `PaymentTxn` method. This transaction is then signed with the recovered private key. As noted earlier, in a production application, the transaction should be signed by a valid wallet provider. The signed transaction is submitted to the node and the `wait_for_confirmation` SDK utility function is called, which will return when the transaction is finalized on the blockchain.

Another utility function is also added to create a payment transaction that is signed by the escrow logic. This function is very similar to the previous function.

<!-- ===PYTEAL_LSIG_SIMPLE_WITHDRAW=== -->
```python
def lsig_payment_txn(
    encoded_program: str, amt: int, rcv: str, algod_client: algod.AlgodClient
):
    # Create an lsig object using the compiled, b64 encoded program
    program = base64.b64decode(encoded_program)
    lsig = transaction.LogicSigAccount(program)

    # Create transaction with the lsig address as the sender
    params = algod_client.suggested_params()
    unsigned_txn = transaction.PaymentTxn(lsig.address(), params, rcv, amt)

    # sign the transaction using the logic
    stxn = transaction.LogicSigTransaction(unsigned_txn, lsig)
    tx_id = algod_client.send_transaction(stxn)
    pmtx = transaction.wait_for_confirmation(algod_client, tx_id, 10)
    return pmtx


```
[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/lsig.py#L74-L92)
<!-- ===PYTEAL_LSIG_SIMPLE_WITHDRAW=== -->

The primary difference is that the function is passed the base64 encoded string of the compiled bytecode for the smart signature and the escrow’s Algorand address. The program is then converted to a byte array and the Python SDK’s `LogicSigAccount` function is used to create a logic signature from the program bytes. The payment transaction is then signed with the logic using the SDKs `LogicSigTransaction` function. For more information on Logic Signatures and smart signatures see the [smart signatures documentation](../smart-contracts/smartsigs/index.md).

The solution can be completed by adding a main function to put the utility functions to use.

<!-- ===PYTEAL_LSIG_SIMPLE_USAGE=== -->
```python
def main():
    # initialize an algodClient
    algod_client = algod.AlgodClient(algod_token, algod_address)

    # define private keys
    private_key = mnemonic.to_private_key(benefactor_mnemonic)
    receiver_public_key = account.address_from_private_key(private_key)

    print("Compiling Donation Smart Signature......")

    stateless_program_teal = donation_escrow(receiver_public_key)
    escrow_result, escrow_address = compile_smart_signature(
        algod_client, stateless_program_teal
    )

    print("Program:", escrow_result)
    print("LSig Address: ", escrow_address)

    print("Activating Donation Smart Signature......")

    # Activate escrow contract by sending 2 algo and 1000 microalgo for transaction fee from creator
    amt = 2001000
    payment_transaction(sender_mnemonic, amt, escrow_address, algod_client)

    print("Withdraw from Donation Smart Signature......")

    # Withdraws 1 ALGO from smart signature using logic signature.
    withdrawal_amt = 1000000
    lsig_payment_txn(escrow_result, withdrawal_amt, receiver_public_key, algod_client)


```
[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/lsig.py#L96-L127)
<!-- ===PYTEAL_LSIG_SIMPLE_USAGE=== -->

The main function first makes a connection to the sandbox installed node, then the benefactor’s address is recovered. The `donation_escrow` built in the previous section is called to produce the TEAL for the smart signature. This TEAL is then compiled, returning both the base64 encoded bytes of the program and the address of the escrow.

A simple payment transaction is then created to fund the escrow with a little over 2 Algos. Finally, 1 Algo is dispensed from the escrow to the benefactor using a payment transaction signed by the smart signature. 

For more information on smart signatures, see the [developer documentation](../smart-contracts/smartsigs/index.md).