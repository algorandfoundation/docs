title: Beaker


Beaker is a framework for building Smart Contracts using PyTeal. 

This page shows some high level details but for more see [Beaker's documentation](https://beaker.algo.xyz)


# High Level Overview 

talk about what it provides

# Install

Make sure you have python3 installed then

`pip install beaker-pyteal`

# Initialize Application

To create an application simply initialize a new Beaker Application object

<!-- ===BEAKER_INIT_APP=== -->
```python
from beaker import Application

app = Application("MyRadApp", descr="This is a rad app")
```
[Snippet Source](https://github.com/algorand-devrel/beaker/blob/examples/examples/docs_app/app_handlers.py#L1-L4)
<!-- ===BEAKER_INIT_APP=== -->

This is enough to generate the app specification that can be exported for use by other tools.

<!-- ===BEAKER_APP_SPEC=== -->
```python
app_spec = app.build()
print(app_spec.to_json())
```
[Snippet Source](https://github.com/algorand-devrel/beaker/blob/examples/examples/docs_app/app_handlers.py#L7-L9)
<!-- ===BEAKER_APP_SPEC=== -->

# Add Method Handlers

Adding method handlers can be done in 2 ways...

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


Adding handlers using a blueprint allows code re-use and makes it easier to add behaviors especially for applications that wish to adhere to some ARC standard.

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

# Add State 

Lets go back and add some state to our application

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

We can interact with our application using the Beaker provided ApplicationClient

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

Now we can deploy the app and interact with it

<!-- ===BEAKER_APP_CLIENT_DEPLOY=== -->
```python
app_id, app_addr, txid = app_client.create()
print(f"Created app with id: {app_id} and address: {app_addr} in tx: {txid}")
```
[Snippet Source](https://github.com/algorand-devrel/beaker/blob/examples/examples/docs_app/app_client.py#L19-L21)
<!-- ===BEAKER_APP_CLIENT_DEPLOY=== -->

<!-- ===BEAKER_APP_CLIENT_CALL=== -->
```python
result = app_client.call("add", a=1, b=2)
print(result.return_value)  # 3
```
[Snippet Source](https://github.com/algorand-devrel/beaker/blob/examples/examples/docs_app/app_client.py#L25-L27)
<!-- ===BEAKER_APP_CLIENT_CALL=== -->

This is only a small sample of what Beaker can do. For more see [Beaker's documentation](https://beaker.algo.xyz)