# App client

Application client that works with ARC-0032 application spec defined smart contracts (e.g. via Beaker).

App client is a higher-order use case capability provided by AlgoKit Utils that builds on top of the core capabilities, particularly [App deployment](./app-deploy.md) and [App management](./app.md). It allows you to access a high productivity application client that works with ARC-0032 application spec defined smart contracts, which you can use to create, update, delete, deploy and call a smart contract and access state data for it.

To see some usage examples check out the [automated tests](../../src/types/app-client.spec.ts).

## Design

The design for the app client is based on a wrapper for parsing an [ARC-0032](https://github.com/algorandfoundation/ARCs/pull/150) application spec and wrapping the [App deployment](./app-deploy.md) functionality and corresponding [design](./app-deploy.md#design). It's also heavily inspired by [beaker-ts](https://github.com/algorand-devrel/beaker-ts), which this library aims to eventually replace.

## Creating an application client

To create an application you can either use `algokit.getAppClient(appDetails, algod)` or `import { ApplicationClient } from '@algorandfoundation/algokit-utils/types/app-client'` and `new ApplicationClient(appDetails, algod)`

The `appDetails` parameter is of type [`AppSpecAppDetails`](../code/modules/types_app_client.md#appspecappdetails), which contains some core properties and then one of two key mechanisms to specify the app to target.

- Core parameters - Core parameters that can always be applied
  - `app: AppSpec | string` - Either the parsed ARC-0032 `AppSpec`, or a raw JSON `string` which will get parsed as an `AppSpec`
  - `sender?: SendTransactionFrom` - Optional [sender](./account.md#sendtransactionfrom) to send/sign all transactions with (if left out then individual methods must have a sender passed to them)
  - `params?: SuggestedParams` - Optional [sending parameters](./transaction.md#transaction-params) if you want to avoid an extra call to algod
- App target - How to resolve an existing app (if one exists), which can either be:
  1. [`ResolveAppById`](../code/interfaces/types_app_client.ResolveAppById.md) - When you want to resolve an existing app by app ID, which consists of the following parameters:
     - `id: number` - The app ID, which should be set as `0` if you have yet to deploy the contract
     - `name? string` - The optional name to mark the contract with if you are deploying it, otherwise `contract.name` is used from the app spec
  2. [`ResolveAppByCreatorAndName`](../code/modules/types_app_client.md#resolveappbycreatorandname) - When you want to resolve an existing app by name for a given creator account, which consists of the following parameters:
     - `creatorAddress: string` - The address of the creator account of the app for which to search for the deployed app under
     - `name?: string` - An overridden name to identify the contract with, otherwise `contract.name` is used from the app spec
     - And either:
       1. `indexer: Indexer` - An indexer instance so the existing app deployments can be queried
       2. `existingDeployments: AppLookup` - The result of an existing indexer lookup to generate an [app lookup](./app-deploy.md#getcreatorappsbyname), which avoids extra indexer calls from being made

## Creating, updating, deploying and deleting the app

Once you have an application client you can perform the following actions related to creating and managing the lifecycle of an app:

- `compile(compilationParams?)` - Allows you to compile the application (approval and clear program)), including [deploy-time parameter replacements and deploy-time immutability and permanence control](./app-deploy.md#compilation-and-template-substitution); it returns the compiled AVM code and source maps
- `deploy(deploymentParams?)` - Allows you to perform an idempotent (safely retryable) deployment of the smart contract app per the design of [`deployApp`](app-deploy.md#deployapp)
- `create(createParams?)` - Allows you to perform a creation of the smart contract app
- `update(updateParams?)` - Allows you to perform an update of the (existing) smart contract app
- `delete(deleteParams?)` - Allows you to delete the (existing) smart contract app

The input payload for `create` and `update` are the same and are a union of [`AppClientCallParams`](#appclientcallparams) and [`AppClientCompilationParams`](#appclientcompilationparams). The input payload for `delete` is [`AppClientCallParams`](#appclientcallparams). The input payload for `deploy` is [`AppClientDeployParams`](#appclientdeployparams).

The return payload for these methods directly matches the equivalent underlying [App management](./app.md) / [App deployment](./app-deploy.md) methods (since these methods are wrappers):

- `create` -> [`createApp`](./app.md#createapp)
- `update` -> [`updateApp`](./app.md#updateapp)
- `delete` -> [`deleteApp`](./app.md#deleteapp)
- `deploy` -> [`deployApp`](./app-deploy.md#deployapp)

## Calling the app

To make a call to a smart contract you can use the following methods (which determine the [on complete action](https://developer.algorand.org/docs/get-details/dapps/smart-contracts/apps/#the-lifecycle-of-a-smart-contract) that the call will use):

- `call(call?)` - A normal (`noop` on completion action) call
- `optIn(call?)` - An opt-in call
- `closeOut(call?)` - A close-out call
- `clearState(call?)` - A clear state call (note: calls the clear program)
- `callOfType(call, callType)` - Make a call with a specified call type

These calls will only work if the Application Client knows the ID of the app, which will occur if:

- The app ID is passed into the constructor;
- You have passed `creatorAccount` and the smart contract name to the constructor and the contract already exists; or
- You have called `create` or `deploy` using that Application Client.

The input payload for all of these calls is the same as `delete`; [`AppClientCallParams`](#appclientcallparams).

The return payload for all of these is the same as [`callApp`](./app.md#calling-an-app).

## Getting a reference to the app

To get reference information for the app from outside the Application Client you can call `getAppReference()`. If you passed the `creatorAddress` and app name to the constructor then this method will return the full [`AppMetadata`](../code/interfaces/types_app.AppMetadata.md) per [`getCreatorAppsByName`](./app-deploy.md#getcreatorappsbyname). If you just passed in the app ID or used `create` rather than `deploy` then you will just receive an [`AppReference`](../code/interfaces/types_app.AppReference.md) (which is also a sub-type of the `AppMetadata`):

- `appId: number` - The id of the app
- `appAddress: string` - The Algorand address of the account associated with the app

## Funding the app account

Often there is a need to fund an app account to cover minimum balance requirements for boxes and other scenarios. There is a helper method that will do this for you `fundAppAccount(fundParams)`.

The input parameters can either be:

- An [`AlgoAmount`](./amount.md) value (note: requires `sender` to be passed into the constructor)
- A [`FundAppAccountParams`](../code/interfaces/types_app_client.FundAppAccountParams.md), which has the following properties:
  - `amount: AlgoAmount` - The amount to fund
  - `sender?: SendTransactionFrom` - The [sender/signer](./account.md#sendtransactionfrom) to use; if unspecified then the sender that was passed into the constructor of the Application Client is used
  - `note?: TransactionNote` - The [transaction note](./transaction.md#transaction-notes) to use when issuing the transaction
  - `sendParams?: SendTransactionParams` - The [transaction sending configuration](./transaction.md#sendtransactionparams)

This call will only work if the Application Client knows the ID of the app, which will occur if:

- The app ID is passed into the constructor;
- You have passed `creatorAccount` and the smart contract name to the constructor and the contract already exists; or
- You have called `create` or `deploy` using that Application Client.

Note: If you are passing the funding payment in as an ABI argument so it can be validated by the ABI method then you'll want to issue the `skipSending` configuration. That might look something like this as an example:

```typescript
const result = await appClient.call({
  method: 'bootstrap',
  methodArgs: {
    args: [
      appClient.fundAppAccount({
        amount: algokit.microAlgos(200_000),
        sendParams: { skipSending: true },
      }),
    ],
    boxes: ['Box1'],
  },
})
```

## Reading state

There are various methods defined that let you read state from the smart contract app:

- `getGlobalState()` - Gets the current global state using [`algokit.getAppGlobalState](./app.md#global-state)
- `getLocalState(account: string | SendTransactionFrom)` - Gets the current local state for the given account address using [`algokit.getAppLocalState](./app.md#local-state).
- `getBoxNames()` - Gets the current box names using [`algokit.getAppBoxNames`](./app.md#boxes)
- `getBoxValue(name)` - Gets the current value of the given box using [`algokit.getAppBoxValue](./app.md#boxes)
- `getBoxValueFromABIType(name)` - Gets the current value of the given box from an ABI type using [`algokit.getAppBoxValueFromABIType](./app.md#boxes)
- `getBoxValues(filter)` - Gets the current values of the boxes using [`algokit.getAppBoxValues](./app.md#boxes)
- `getBoxValuesFromABIType(type, filter)` - Gets the current values of the boxes from an ABI type using [`algokit.getAppBoxValuesFromABIType](./app.md#boxes)

These calls will only work if the Application Client knows the ID of the app, which will occur if:

- The app ID is passed into the constructor;
- You have passed `creatorAccount` and the smart contract name to the constructor and the contract already exists; or
- You have called `create` or `deploy` using that Application Client.

## Handling logic errors and diagnosing errors

Often when calling a smart contract during development you will get logic errors that cause an exception to throw. This may be because of a failing assertion, a lack of fees, exhaustion of opcode budget, or any number of other reasons.

When this occurs, you will generally get an error that looks something like: `TransactionPool.Remember: transaction {TRANSACTION_ID}: logic eval error: {ERROR_MESSAGE}. Details: pc={PROGRAM_COUNTER_VALUE}, opcodes={LIST_OF_OP_CODES}`.

The information in that error message can be parsed and when combined with the [source map from compilation](./app-deploy.md#compilation-and-template-substitution) you can expose debugging information that makes it much easier to understand what's happening.

The Application Client automatically provides this functionality for all smart contract calls. It also exposes a function that can be used for any custom calls you manually construct and need to add into your own try/catch `exposeLogicError(e: Error, isClear?: boolean)`.

When an error is thrown then the resulting error that is re-thrown will be a [`LogicError` object](../code/classes/types_logic_error.LogicError.md), which has the following fields:

- `message: string` - The formatted error message `{ERROR_MESSAGE}. at:{TEAL_LINE}. {ERROR_DESCRIPTION}`
- `stack: string` - A stack trace of the TEAL code showing where the error was with the 5 lines either side of it
- `led: LogicErrorDetails` - The parsed [logic error details](../code/interfaces/types_logic_error.LogicErrorDetails.md) from the error message, with the following properties:
  - `txId: string` - The transaction ID that triggered the error
  - `pc: number` - The program counter
  - `msg: string` - The raw error message
  - `desc: string` - The full error description
  - `traces: Record<string, unknown>[]` - Any traces that were included in the error
- `program: string[]` - The TEAL program split by line
- `teal_line: number` - The line number in the TEAL program that triggered the error

Note: This information will only show if the Application Client has a source map. This will occur if:

- You have called `create`, `update` or `deploy`
- You have called `importSourceMaps(sourceMaps)` and provided the source maps (which you can get by calling `exportSourceMaps()` after calling `create`, `update` or `deploy` and it returns a serialisable value)

If you want to go a step further and automatically issue a [dry run transaction](https://developer.algorand.org/docs/get-details/dapps/smart-contracts/debugging/#dryrun-rest-endpoint) when there is an error when an ABI method is called you can turn on debug mode:

```typescript
algokit.Config.configure({ debug: true })
```

> ⚠️ **Note:** The "dry run" feature has been deprecated and is now replaced by the "simulation" feature. Please refer to the [Simulation Documentation](https://algorand.github.io/js-algorand-sdk/classes/modelsv2.SimulateTransactionResult.html) for more details.

If you do that then the exception will have the `traces` property within the underlying exception will have key information from the simulation within it and this will get populated into the `led.traces` property of the thrown error.

## `AppClientCallParams`

All methods that call the smart contract apart from `deploy` make use of [this type](../code/modules/types_app_client.md#appclientcallparams). It consists of the following core properties, all of which are optional:

- `sender?: SendTransactionFrom` - The [sender/signer](./account.md#sendtransactionfrom) to use; if unspecified then the sender that was passed into the constructor of the Application Client is used
- `note?: TransactionNote` - The [transaction note](./transaction.md#transaction-notes) to use when issuing the transaction
- `sendParams?: SendTransactionParams` - The [transaction sending configuration](./transaction.md#sendtransactionparams)

In addition to these parameters, it may specify [call arguments](#appclientcallargs) parameters (`AppClientCallArgs`).

## `AppClientCallArgs`

Whenever an app call is specified, including within `deploy` [this type](../code/modules/types_app_client.md#appclientcallargs) specifies the arguments. There are two forms you can use:

- **Raw call** - Directly specifies the [values](./app.md#rawappcallargs) that will be populated onto an `algosdk.Transaction`
- **ABI call** - Specifies a [ARC-0004 ABI call](https://developer.algorand.org/docs/get-details/dapps/smart-contracts/ABI/) along with relevant values (like boxes) that will be directly populated onto an `algosdk.Transaction`. Consists of the [ABI app call args type](./app.md#abiappcallargs) with the `method` parameter replaced with a string (since the Application Client only needs it as a string):
  - `method: string` - The name of the method (e.g. `hello`) or the ABI signature of the method (e.g. `hello(string)string`) for when you have multiple methods with the same name and need to differentiate between them
  - `methodArgs?: ABIAppCallArg[]` - An array of arguments to pass into the ABI method
  - `boxes: (BoxReference | BoxIdentifier | algosdk.BoxReference)[]` - Any [boxes](./app.md#referencing-boxes) to load to the [boxes array](https://developer.algorand.org/docs/get-details/dapps/smart-contracts/apps/#reference-arrays)
  - `lease: string | Uint8Array`: A [lease](https://developer.algorand.org/articles/leased-transactions-securing-advanced-smart-contract-design/) to assign to the transaction to enforce a mutually exclusive transaction (useful to prevent double-posting and other scenarios)

If you want to get call args for manually populating into an `algosdk.Transaction` you can use the `getCallArgs` method on Application Client.

Also, if you want to manually construct an ABI call (e.g. to use with AtomicTransactionComposer directly) then you can use `getABIMethod` and/or `getABIMethodParams`

## `AppClientCompilationParams`

When calling `create` or `update` there are extra parameter that need to be passed to facilitate the compilation of the code in addition to the other parameters in [`AppClientCallParams`](#appclientcallparams):

- `deployTimeParams?: TealTemplateParams` - Any [deploy-time parameters](./app-deploy.md#compilation-and-template-substitution) to replace in the TEAL code
- `updatable?: boolean` - Whether or not the contract should have [deploy-time immutability control](./app-deploy.md#compilation-and-template-substitution) set, undefined = ignore
- `deletable?: boolean` - Whether or not the contract should have [deploy-time permanence control](./app-deploy.md#compilation-and-template-substitution) set, undefined = ignore

## `AppClientDeployParams`

When calling `deploy` the [`AppClientDeployParams`](../code/interfaces/types_app_client.AppClientDeployParams.md) type defines the input parameters, all of which are optional. This type closely models the semantics of [`deployApp`](./app-deploy.md#deployapp).

- `version?: string` - The version of the contract, uses "1.0" by default
- `sender?: SendTransactionFrom` - The [sender/signer](./account.md#sendtransactionfrom) to use; if unspecified then the sender that was passed into the constructor of the Application Client is used
- `sendParams?: SendTransactionParams` - The [transaction sending configuration](./transaction.md#sendtransactionparams)
- `deployTimeParams?: TealTemplateParams` - Any [deploy-time parameters](./app-deploy.md#compilation-and-template-substitution) to replace in the TEAL code
- `allowUpdate?: boolean` - Whether or not to allow updates in the contract using the [deploy-time updatability](./app-deploy.md#compilation-and-template-substitution) control if present in your contract; if this is not specified then it will automatically be determined based on the AppSpec definition (if there is an update method)
- `allowDelete?: boolean` - Whether or not to allow deletes in the contract using the [deploy-time deletability](./app-deploy.md#compilation-and-template-substitution) control if present in your contract; if this is not specified then it will automatically be determined based on the AppSpec definition (if there is a delete method)
- `onSchemaBreak?: 'replace' | 'fail' | OnSchemaBreak` determines what should happen if a breaking change to the schema is detected (e.g. if you need more global or local state that was previously requested when the contract was originally created)
- `onUpdate?: 'update' | 'replace' | 'fail' | OnUpdate` determines what should happen if an update to the smart contract is detected (e.g. the TEAL code has changed since last deployment)
- `createArgs?: AppClientCallArgs` - The args to use if a create is needed
- `updateArgs?: AppClientCallArgs` - The args to use if an update is needed
- `deleteArgs?: AppClientCallArgs` - The args to use if a delete is needed
