title: Beaker

!!! note
For a native Python experience, checkout our [Algorand Python](https://algorandfoundation.github.io/puya/) docs.

Beaker is a framework for building Smart Contracts using PyTeal. Beaker is designed to simplify writing, testing and deploying Algorand smart contracts. The Beaker source code available on [github](https://github.com/algorand-devrel/beaker). 

This page provides an overview of the features available in Beaker. For complete details see the [Beaker's documentation](https://beaker.algo.xyz). 


# Quick start videos

If you prefer videos, take a look at this playlist to learn about Beaker. Most of the videos in the list are under 12 minutes each.

<iframe width="100%" style="aspect-ratio:16/9" src="https://www.youtube-nocookie.com/embed/RhfC5Xd25dk" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>

# High Level Overview 
Beaker provides several packages that extend PyTeal and provide convenience functionality for testing and deploying smart contracts. 

The  `Application` class is Beaker's primary class. It is used to create [ABI](/docs/get-details/dapps/smart-contracts/ABI/) compliant Algorand smart contracts. Beaker also provides decorators to route specific [application transactions](/docs/get-details/dapps/smart-contracts/apps/) to the proper functionality within a smart contract. Beaker facilitates management of [local](/docs/get-details/dapps/smart-contracts/apps/state#local-storage) and [global](/docs/get-details/dapps/smart-contracts/apps/state#global-storage) state, and [box storage](/docs/get-details/dapps/smart-contracts/apps/state#box-storage). 

The `ApplicationSpecification` class is used to generate a JSON manifest that describes the contract methods, source, and state schema used. This manifest can be used by other modules and utilities to deploy the smart contract. 

The `ApplicationClient` class can be used to connect to an Algorand node and interact with a specific `Application`.

Beaker's sandbox module can be used to quickly connect to the default docker sandbox installation to deploy and call a contract.  

# Install

Beaker can be installed using pip package manager.

`pip install beaker-pyteal`

Alternatively, Beaker can be installed with [AlgoKit](/docs/get-details/algokit) using the beaker project template.

`algokit init --template beaker`

Either of these methods will also install PyTeal in addition to Beaker.

!!!note
    Beaker requires python version 3.10 or higher


# Initialize Application

To create an application simply initialize a new Beaker Application object, supplying the name and description.

<!-- ===BEAKER_INIT_APP=== -->
```python
from beaker import Application

app = Application("MyRadApp", descr="This is a rad app")
```
[Snippet Source](https://github.com/algorand-devrel/beaker/blob/examples/examples/docs_app/app_handlers.py#L1-L4)
<!-- ===BEAKER_INIT_APP=== -->

This is enough to generate the `ApplicationSpecification` that can be exported for use by other tools. This spec can be generated using the `Application` `build` method. Optionally you can use the `export` method to export the approval and clear TEAL programs, the ABI contract manifest, and the application specification.

!!!note
    At this point a complete smart contract has been written, albeit with no utility.

<!-- ===BEAKER_APP_SPEC=== -->
```python
app_spec = app.build()
print(app_spec.to_json())
```
[Snippet Source](https://github.com/algorand-devrel/beaker/blob/examples/examples/docs_app/app_handlers.py#L7-L9)
<!-- ===BEAKER_APP_SPEC=== -->



# Add Method Handlers

Method handlers can be added to provide functionality within the smart contract. This can be accomplished using the `external` decorator or by using a blueprint.



## external

To provide a method that can be invoked by an application call transaction, Beaker provides the `external` decorator. This instructs the framework to expose the method publicly for incoming transactions. The method is then defined with its required [method signature](/docs/get-details/dapps/smart-contracts/ABI/#methods), where the parameter types in the method signature describe the input types and output type.  These types of the arguments must be valid ABI data types, using [PyTeal's ABI](https://pyteal.readthedocs.io/en/stable/abi.html) package. Arguments and output types are optional, omitting any arguments is perfectly valid.

!!! note
    Note that input types come first, and if a value is returned it should be denoted in the method signature at the end using the notation `*, output: abi.ValidABIType`, which provides a variable to write the output into.

<!-- ===BEAKER_HANDLERS_DIRECT=== -->
```python
import pyteal as pt


# use the decorator provided on the `app` object to register a handler
@app.external
def add(a: pt.abi.Uint64, b: pt.abi.Uint64, *, output: pt.abi.Uint64) -> pt.Expr:
    return output.set(a.get() + b.get())


```
[Snippet Source](https://github.com/algorand-devrel/beaker/blob/examples/examples/docs_app/app_handlers.py#L13-L22)
<!-- ===BEAKER_HANDLERS_DIRECT=== -->

In the example above the add method is defined to take two `Uint64` arguments (`a`, `b`) and return a `Uint64` (`output`).

The full method signature for the above is `add(uint64,uint64)uint64` and will, by default, field only application call transactions with an `OnComplete` of `NoOp`.

## On Complete handlers

There are other decorators that can be used to modify the behavior of the method handler including `create`, `optin`, `closeout`, `clear`, `update`, and `delete`. These decorators can be used to register handlers for specific `OnComplete` values. See the [full docs](https://beaker.algo.xyz) for more details.

## Blueprints

Beaker allows for a pattern called `blueprints` apply a set of method handlers to an Application. Adding handlers using a blueprint allows for code re-use and makes it easier to add behaviors, especially for applications that wish to adhere to some ARC standard. 

Blueprints can be defined using a standard python method definition that accepts an Application as an argument and applies the handlers.

The code below defines a calculator blueprint that applies method handlers for a set of functions to implement a simple calculator. The blueprint must take an `Application` argument and optionally other arguments to modify the behavior of the blueprint

An instantiated `Application` can apply this blueprint using the `.apply` method passing blueprint method as an argument. If other arguments in the blueprint method are defined, they can be passed with standard python kwarg format (i.e. `.apply(bp, arg1="hello")`)

<!-- ===BEAKER_HANDLERS_BLUEPRINT=== -->
```python
# passing the app to this method will register the handlers on the app
def calculator_blueprint(app: Application) -> Application:
    @app.external
    def add(a: pt.abi.Uint64, b: pt.abi.Uint64, *, output: pt.abi.Uint64) -> pt.Expr:
        return output.set(a.get() + b.get())

    @app.external
    def sub(a: pt.abi.Uint64, b: pt.abi.Uint64, *, output: pt.abi.Uint64) -> pt.Expr:
        return output.set(a.get() - b.get())

    @app.external
    def div(a: pt.abi.Uint64, b: pt.abi.Uint64, *, output: pt.abi.Uint64) -> pt.Expr:
        return output.set(a.get() / b.get())

    @app.external
    def mul(a: pt.abi.Uint64, b: pt.abi.Uint64, *, output: pt.abi.Uint64) -> pt.Expr:
        return output.set(a.get() * b.get())

    return app


calculator_app = Application("CalculatorApp", descr="This is a calculator app")
calculator_app.apply(calculator_blueprint)

calculator_app_spec = calculator_app.build()
print(calculator_app_spec.to_json())
```
[Snippet Source](https://github.com/algorand-devrel/beaker/blob/examples/examples/docs_app/app_handlers.py#L25-L51)
<!-- ===BEAKER_HANDLERS_BLUEPRINT=== -->

An application that has a blueprint applied can also implement additional handlers or apply additional blueprints. Identical method signatures are not allowed. However, if necessary, an identical method signature still be registered by adding the override attribute to the decorator, `external(override=True)`.  

# Add State 

An Application can define the state it uses to store data. This is done by defining a class that contains some number of `StateValue` objects as attributes and passing an instance of that class to the Application constructor. 


Beaker's `GlobalStateValue` class can be used to define and alter a contract's global state values. Global state values are defined by passing the `TealType` and a description the `GlobalStateValue` constructor. 

!!! note
    `TealType` is specific to the AVM and only `bytes`, `unit64` are acceptable values for state.


The code below illustrates creating global integer counter that is stored in state. First, a `CounterState` class is created with an instance of a `GlobalStateValue` as an attribute. This class is then instantiated and passed in the `Application` constructor. Two method handlers are also added to the app to increment and decrement the counter.    


<!-- ===BEAKER_STATE_GLOBAL=== -->
```python
import pyteal as pt
from beaker import Application, GlobalStateValue


class CounterState:
    counter = GlobalStateValue(
        stack_type=pt.TealType.uint64,
        descr="A counter for showing how to use application state",
    )


app = Application(
    "CounterApp", descr="An app that holds a counter", state=CounterState()
)


@app.external
def increment() -> pt.Expr:
    return app.state.counter.set(app.state.counter + pt.Int(1))


@app.external
def decrement() -> pt.Expr:
    return app.state.counter.set(app.state.counter - pt.Int(1))


app_spec = app.build()
print(app_spec.global_state_schema.dictify())
```
[Snippet Source](https://github.com/algorand-devrel/beaker/blob/examples/examples/docs_app/app_state.py#L3-L31)
<!-- ===BEAKER_STATE_GLOBAL=== -->



Similarly, a `LocalStateValue` can be used to alter and store local state values. The code below is identical to the previous example, except the counter is stored locally.

<!-- ===BEAKER_STATE_LOCAL=== -->
```python
import pyteal as pt
from beaker import Application, LocalStateValue


class LocalCounterState:
    local_counter = LocalStateValue(
        stack_type=pt.TealType.uint64,
        descr="A counter for showing how to use application state",
    )


local_app = Application(
    "CounterApp", descr="An app that holds a counter", state=LocalCounterState()
)


@local_app.external
def user_increment() -> pt.Expr:
    return local_app.state.local_counter.set(local_app.state.local_counter + pt.Int(1))


@local_app.external
def user_decrement() -> pt.Expr:
    return local_app.state.local_counter.set(local_app.state.local_counter - pt.Int(1))


local_app_spec = local_app.build()
print(local_app_spec.local_state_schema.dictify())
```
[Snippet Source](https://github.com/algorand-devrel/beaker/blob/examples/examples/docs_app/app_state.py#L34-L62)
<!-- ===BEAKER_STATE_LOCAL=== -->



Beaker provides the `BoxMapping` and `BoxList` classes to work in conjunction with existing PyTeal [box functionality](/docs/get-details/dapps/smart-contracts/apps/state#box-storage).

In the example below a `BoxMapping` instance is defined in the `MappingState` class. Each entry in the map is keyed using the type of `Address` and stores a `Uint64` value. 

The method handler we define allows us to set an integer for the specific application caller's address. 


<!-- ===BEAKER_STATE_MAPPING=== -->
```python
import pyteal as pt
from beaker.lib.storage import BoxMapping


class MappingState:
    users = BoxMapping(pt.abi.Address, pt.abi.Uint64)


mapping_app = Application(
    "MappingApp", descr="An app that holds a mapping", state=MappingState()
)


@mapping_app.external
def store_user_value(value: pt.abi.Uint64) -> pt.Expr:
    # access an element in the mapping by key
    return mapping_app.state.users[pt.Txn.sender()].set(value)


```
[Snippet Source](https://github.com/algorand-devrel/beaker/blob/examples/examples/docs_app/app_state.py#L65-L84)
<!-- ===BEAKER_STATE_MAPPING=== -->

The `BoxList` class can be used to store a list of specific _static_ ABI types. The example below creates a box named `users` that stores a list of five addresses. The `store_user` method is passed an address and an index. The passed-in address is then stored in the `BoxList` at the specified index.

<!-- ===BEAKER_STATE_LIST=== -->
```python
import pyteal as pt
from beaker.lib.storage import BoxList


class ListState:
    users = BoxList(pt.abi.Address, 5)


list_app = Application("ListApp", descr="An app that holds a list", state=ListState())


@list_app.external
def store_user(user: pt.abi.Address, index: pt.abi.Uint64) -> pt.Expr:
    # access an element in the list by index
    return list_app.state.users[index.get()].set(user)


```
[Snippet Source](https://github.com/algorand-devrel/beaker/blob/examples/examples/docs_app/app_state.py#L88-L105)
<!-- ===BEAKER_STATE_LIST=== -->

# Interacting with the Application

The contract can be deployed and tested using Beaker's sandbox module and the `ApplicationClient` class.
The code below first retrieves the accounts from the currently running sandbox instance. A `ApplicationClient` (app_client) is then instantiated with an algod client, the `Application` class that is going to be used, and the first sandbox account (sandbox default starts with a couple of predefined accounts) which will be used to sign transactions.


<!-- ===BEAKER_APP_CLIENT_INIT=== -->
```python
from beaker import sandbox, client

# grab funded accounts from the sandbox KMD
accts = sandbox.get_accounts()

# get a client for the sandbox algod
algod_client = sandbox.get_algod_client()

# create an application client for the calculator app
app_client = client.ApplicationClient(
    algod_client, calculator_app, signer=accts[0].signer
)
```
[Snippet Source](https://github.com/algorand-devrel/beaker/blob/examples/examples/docs_app/app_client.py#L3-L15)
<!-- ===BEAKER_APP_CLIENT_INIT=== -->

The instance of `AppliationClient` can deploy the calculator app now using the `create` method.

<!-- ===BEAKER_APP_CLIENT_DEPLOY=== -->
```python
app_id, app_addr, txid = app_client.create()
print(f"Created app with id: {app_id} and address: {app_addr} in tx: {txid}")
```
[Snippet Source](https://github.com/algorand-devrel/beaker/blob/examples/examples/docs_app/app_client.py#L19-L21)
<!-- ===BEAKER_APP_CLIENT_DEPLOY=== -->

The contract can then be used to call the contract. In this example the contracts `add` method is called, and two integers are passed as method arguments. Finally, the return value is printed.

<!-- ===BEAKER_APP_CLIENT_CALL=== -->
```python
result = app_client.call("add", a=1, b=2)
print(result.return_value)  # 3
```
[Snippet Source](https://github.com/algorand-devrel/beaker/blob/examples/examples/docs_app/app_client.py#L25-L27)
<!-- ===BEAKER_APP_CLIENT_CALL=== -->

This is only a small sample of what Beaker can do. For more see [Beaker's documentation](https://beaker.algo.xyz)