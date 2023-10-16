# App client

Application client that works with ARC-0032 application spec defined smart contracts (e.g. via Beaker).

App client is a high productivity application client that works with ARC-0032 application spec defined smart contracts, which you can use to create, update, delete, deploy and call a smart contract and access state data for it.

To see some usage examples check out the [automated tests](https://github.com/algorandfoundation/algokit-utils-py/blob/main/tests/test_app_client_call.py).

## Design

The design for the app client is based on a wrapper for parsing an [ARC-0032](https://github.com/algorandfoundation/ARCs/pull/150) application spec and wrapping the [App deployment](./app-deploy.md) functionality and corresponding [design](./app-deploy.md#design).

## Creating an application client

There are two key ways of instantiating an ApplicationClient:

1. By app ID - When needing to call an existing app by app ID or unconditionally create a new app.
The signature `ApplicationClient(algod_client, app_spec, app_id=..., ...)` requires:
   * `algod_client`: An `AlgodClient`
   * `app_spec`: An `ApplicationSpecification`
   * `app_id`: The app_id of an existing application, or 0 if creating a new app

2. By creator and app name - When needing to deploy or find an app associated with a specific creator account and app name.
The signature `ApplicationClient(algod_client, app_spec, creator=..., indexer=..., app_lookup)` requires:
   * `algod_client`: An `AlgodClient`
   * `app_spec`: An `ApplicationSpecification`
   * `creator`: The address or `Account` of the creator of the app for which to search for the deployed app under
   * `indexer`:
   * `app_lookup`: Optional if an indexer is provided,
   * `app_name`: An overridden name to identify the contract with, otherwise `contract.name` is used from the app spec

Both approaches also allow specifying the following parameters that will be used as defaults for all application calls:
* `signer`: `TransactionSigner` to sign transactions with.
* `sender`: Address to use for transaction signing, will be derived from the signer if not provided.
* `suggested_params`: Default `SuggestedParams` to use, will use current network suggested params by default
 
Both approaches also allow specifying a mapping of template values via the `template_values` parameter, this will be used before compiling the application to replace any 
`TMPL_` variables that may be in the TEAL. The `TMPL_UPDATABLE` and `TMPL_DELETABLE` variables used in some AlgoKit templates are handled by the `deploy` method, but should be included if 
using `create` or `update` directly.

## Calling methods on the app

There are various methods available on `ApplicationClient` that can be used to call an app:  

* `call`: Used to call methods with an on complete action of `no_op`
* `create`: Used to create an instance of the app, by using an `app_id` of 0, includes the approval and clear programs in the call
* `update`: Used to update an existing app, includes the approval and clear programs in the call, and is called with an on complete action of `update_application`
* `delete`: Used to remove an existing app, is called with an on complete action of `delete_application`
* `opt_in`: Used to opt in to an existing app, is called with an on complete action of `opt_in`
* `close_out`: Used to close out of an existing app, is called with an on complete action of `opt_in`
* `clear_state`: Used to unconditionally close out from an app, calls the clear program of an app

### Specifying which method

All methods for calling an app that support ABI methods (everything except `clear_state`) take a parameter `call_abi_method` which can be used to specify which method to call. 
The method selected can be specified explicitly, or allow the client to infer the method where possible, supported values are:

* `None`: The default value, when `None` is passed the client will attempt to find any ABI method or bare method that is compatible with the provided arguments
* `False`: Indicates that an ABI method should not be used, and instead a bare method call is made
* `True`: Indicates that an ABI method should be used, and the client will attempt to find an ABI method that is compatible with the provided arguments
* `str`: If a string is provided, it will be interpreted as either an ABI signature specifying a method, or as an ABI method name
* `algosdk.abi.Method`: The specified ABI method will be called
* `ABIReturnSubroutine`: Any type that has a `method_spec` function that returns an `algosd.abi.Method`

### ABI arguments

ABI arguments are passed as python keyword arguments e.g. to pass the ABI parameter `name` for the ABI method `hello` the following syntax is used `client.call("hello", name="world")`

### Transaction Parameters

All methods for calling an app take an optional `transaction_parameters` argument, with the following supported parameters:

* `signer`: The `TransactionSigner` to use on the call. This overrides any signer specified on the client
* `sender`: The address of the sender to use on the call, must be able to be signed for by the `signer`. This overrides any sender specified on the client
* `suggested_params`: `SuggestedParams` to use on the call. This overrides any suggested_params specified on the client
* `note`: Note to include in the transaction
* `lease`: Lease parameter for the transaction 
* `boxes`: A sequence of boxes to use in the transaction, this is a list of (app_index, box_name) tuples `[(0, "box_name"), (0, ...)]`
* `accounts`: Account references to include in the transaction
* `foreign_apps`: Foreign apps to include in the transaction
* `foreign_assets`: Foreign assets to include in the transaction
* `on_complete`: The on complete action to use for the transaction, only available when using `call` or `create`
* `extra_pages`: Additional pages to allocate when calling `create`, by default a sufficient amount will be calculated based on the current approval and clear. This can be overridden, if more is required
 for a future update

Parameters can be passed as one of the dataclasses `CommonCallParameters`, `OnCompleteCallParameters`, `CreateCallParameters` (exact type depends on method used) 
```python
client.call("hello", transaction_parameters=algokit_utils.OnCompleteCallParameters(signer=...))
```

Alternatively, parameters can be passed as a dictionary e.g.
```python
client.call("hello", transaction_parameters={"signer":...})
```

## Composing calls

If multiple calls need to be made in a single transaction, the `compose_` method variants can be used. All these methods take an `AtomicTransactionComposer` as their first argument.
Once all the calls have been added to the ATC, it can then be executed. For example:

```python
from algokit_utils import ApplicationClient
from algosdk.atomic_transaction_composer import AtomicTransactionComposer

client = ApplicationClient(...)
atc = AtomicTransactionComposer()
client.compose_call(atc, "hello", name="world")
... # additional compose calls

response = client.execute_atc(atc)
```


## Reading state

There are various methods defined that let you read state from the smart contract app:

* `get_global_state` - Gets the current global state of the app
* `get_local_state` - Gets the current local state for the given account address

## Handling logic errors and diagnosing errors

Often when calling a smart contract during development you will get logic errors that cause an exception to throw. This may be because of a failing assertion, a lack of fees, 
exhaustion of opcode budget, or any number of other reasons.

When this occurs, you will generally get an error that looks something like: `TransactionPool.Remember: transaction {TRANSACTION_ID}: logic eval error: {ERROR_MESSAGE}. Details: pc={PROGRAM_COUNTER_VALUE}, opcodes={LIST_OF_OP_CODES}`.

The information in that error message can be parsed and when combined with the [source map from compilation](./app-deploy.md#compilation-and-template-substitution) you can expose debugging 
information that makes it much easier to understand what's happening.

When an error is thrown then the resulting error that is re-thrown will be a `LogicError`, which has the following fields:

* `logic_error`: Original exception
* `program`: Program source (if available)
* `source_map`: Source map used (if available)
* `transaction_id`: Transaction ID of failing transaction
* `message`: The error message
* `line_no`: The line number in the TEAL program that

The function `trace()` will provide a formatted output of the surrounding TEAL where the error occurred.

```{note}
The extended information will only show if the Application Client has a source map. This will occur if:

1.) The ApplicationClient instance has already called, `create, `update` or `deploy` OR
2.) `template_values` are provided when creating the ApplicationClient, so a SourceMap can be obtained automatically OR
3.) `approval_source_map` on `ApplicationClient` has been set from a previously compiled approval program OR
4.) A source map has been exported/imported using `export_source_map`/`import_source_map`"""
```
