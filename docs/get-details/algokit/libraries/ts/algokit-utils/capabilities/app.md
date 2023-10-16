# App management

App management is a higher-order use case capability provided by AlgoKit Utils that builds on top of the core capabilities. It allows you to create, update, delete, call (ABI and otherwise) smart contract apps and the metadata associated with them (including state and boxes).

## `AppCallArgs`

All calls to smart contracts will allow you to optionally specify the arguments you want to pass in to the call.

This [type](../code/modules/types_app.md#appcallargs) is a union of two types: `RawAppCallArgs` and `ABIAppCallArgs`.

### `RawAppCallArgs`

[`RawAppCallArgs`](../code/interfaces/types_app.RawAppCallArgs.md) allows you to explicitly specify all of the arguments and has the following properties (all of which are optional):

- `accounts: (string | algosdk.Address)[]` - Any accounts to add to the [accounts array](https://developer.algorand.org/docs/get-details/dapps/smart-contracts/apps/#reference-arrays) either as the `string` or `algosdk.Address` representation of the public address of the account(s)
- `appArgs: (Uint8Array | string)[]` - Any [arguments to pass to the smart contract call](https://developer.algorand.org/docs/get-details/dapps/smart-contracts/apps/#reference-arrays), either as the binary representation or a string (that will be encoded to binary using `TextEncoder`)
- `boxes: (BoxReference | BoxIdentifier | algosdk.BoxReference)[]` - Any [boxes](#referencing-boxes) to load to the [boxes array](https://developer.algorand.org/docs/get-details/dapps/smart-contracts/apps/#reference-arrays)
- `apps: number[]`: The ID of any apps to load to the [foreign apps array](https://developer.algorand.org/docs/get-details/dapps/smart-contracts/apps/#reference-arrays)
- `assets: number[]`: The ID of any assets to load to the [foreign assets array](https://developer.algorand.org/docs/get-details/dapps/smart-contracts/apps/#reference-arrays)
- `lease: string | Uint8Array`: A [lease](https://developer.algorand.org/articles/leased-transactions-securing-advanced-smart-contract-design/) to assign to the transaction to enforce a mutually exclusive transaction (useful to prevent double-posting and other scenarios)

### `ABIAppCallArgs`

[`ABIAppCallArgs`](../code/modules/types_app.md#abiappcallargs) allows you to specify an [ARC-0004 ABI call](https://developer.algorand.org/docs/get-details/dapps/smart-contracts/ABI/)

- `method: algosdk.ABIMethodParams | algosdk.ABIMethod` - The ABI method to call
- `methodArgs: ABIAppCallArg[]` - The arguments to pass to the ABI call, which can be one of:
  - `algosdk.ABIArgument` - Which can be one of:
    - `boolean`
    - `number`
    - `bigint`
    - `string`
    - `Uint8Array`
    - An array of one of the above types
    - `algosdk.TransactionWithSigner`
  - [`TransactionToSign`](transaction.md#signing)
  - `algosdk.Transaction`
  - `Promise<SendTransactionResult>` - which allows you to use an AlgoKit Utils method call that [returns a transaction](transaction.md#sendtransactionresult) without needing to await the call and extract the transaction, if you do this be sure to use `skipWaiting: true` when specifying the [sending parameters](transaction.md#sendtransactionparams) so you get the transaction without sending it to the network
- `boxes: (BoxReference | BoxIdentifier | algosdk.BoxReference)[]` - Any [boxes](#referencing-boxes) to load to the [boxes array](https://developer.algorand.org/docs/get-details/dapps/smart-contracts/apps/#reference-arrays)
- `lease: string | Uint8Array`: A [lease](https://developer.algorand.org/articles/leased-transactions-securing-advanced-smart-contract-design/) to assign to the transaction to enforce a mutually exclusive transaction (useful to prevent double-posting and other scenarios)

### Utility methods

If you want to manually construct a transaction, but use these types to specify the arguments then you can use the following methods:

- [`algokit.getAppArgsForTransaction`](../code/modules/index.md#getappargsfortransaction) - Takes a `RawAppCallArgs` object and returns the corresponding fields ready to set onto an `algosdk.Transaction`
- [`algokit.getAppArgsForABICall`](../code/modules/index.md#getappargsforabicall) - Takes a `ABIAppCallArgs` object and returns the corresponding fields ready to pass into `addMethodCall` on `AtomicTransactionComposer`

## Referencing boxes

To reference a box in a transaction using AlgoKit Utils, you can specify one of the following types:

- `BoxIdentifier` - Which can be one of:
  - `string` - Which will be encoded as a box name using `TextEncoder`
  - `Uint8Array` - Which will be used directly
  - `SendTransactionFrom` - Which will be converted into the public address of the [sender corresponding to the account](account.md) and the public key encoded to binary
    - This type of encoding is compatible with directly referencing an account address within the smart contract (e.g. `(address := pt.abi.Address()).set(pt.Txn.sender())` in PyTEAL)
- `BoxReference` - Which is an interface that has two fields:
  - `appId: number` - The app ID
  - `name: BoxIdentifier` - The name, per the above type
- `algosdk.BoxReference` - The in-built algosdk `BoxReference` type, which has two properties:
  - `appIndex: number`
  - `name: UInt8Array` - The name in binary

If you specify a `BoxIdentifier` directly outside of the `BoxReference` type then the behaviour is to load the box from the current app the transaction that box identifier appears in. To see more about how box references work consult the [official documentation](https://developer.algorand.org/articles/smart-contract-storage-boxes/).

### `BoxName`

To get a box reference when reading box state there is a helpful [`BoxName`](../code/interfaces/types_app.BoxName.md) type that is exposed, which provides the following properties:

- `name: string`
- `nameBase64: string`
- `nameRaw: Uint8Array`

## Creating and updating apps

### `createApp`

To create an app you can call [`algokit.createApp(createPayload, algod)`](../code/modules/index.md#createapp). See the tests for an [example](../../src/app.spec.ts).

The payload to configure an app consists of a union of [`SendTransactionParams`](transaction.md#sendtransactionparams) and the [following properties](../code/interfaces/types_app.CreateAppParams.md):

- Required
  - `from: SendTransactionFrom` - The [account](account.md) (with private key loaded) that will send the transaction
  - `approvalProgram: Uint8Array | string` - The approval program as raw teal (string) or compiled teal, base 64 encoded as a byte array (Uint8Array)
  - `clearStateProgram: Uint8Array | string` - The clear state program as raw teal (string) or compiled teal, base 64 encoded as a byte array (Uint8Array)
  - `schema: AppStorageSchema` - The storage schema to request for the created app
  - `onCompleteAction?: algosdk.OnApplicationComplete` - The on-completion action to specify for the call; defaults to NoOp
- Optional:
  - `transactionParams: SuggestedParams` - Any [transaction parameters](transaction.md#transaction-params)
  - `note: TransactionNote` - A [transaction note](transaction.md#transaction-notes)
  - `args: AppCallArgs` - Any [arguments](#appcallargs) passed in to the app call

If you pass in `approvalProgram` or `clearProgram` as a string then it will automatically be compiled using Algod and the compilation result will be returned from the function (including the source map). To skip this behaviour you can pass in the compiled TEAL as `Uint8Array`.

If you pass in args that represent an ABI then it will use an `AtomicTransactionComposer` to construct and send the transaction(s). Because it's possible that other transactions may be present as ABI arguments, the full set of transactions that were sent are returned in `transactions` and the primary transaction for the create call will also be available in `transaction`. If you pass in the `atc` or `skipSending: true` then it won't execute the transaction and will simply return the transaction(s). The `return` value will have any ABI return value within it.

### `updateApp`

To update an app you can call [`algokit.updateApp(updatePayload, algod)`](../code/modules/index.md#updateapp).

The update payload and behaviour is the same as `createApp` with two payload differences:

- `schema` is not present, since it can only be set when creating a smart contract
- `appId: number` is present (required), which specifies the ID of the app to update

## Calling an app

To call an app outside of creation or update you can call [`algokit.callApp(callPayload, algod)`](../code/modules/index.md#callapp).

The payload to configure an app call consists of a union of [`SendTransactionParams`](transaction.md#sendtransactionparams) and the [following properties](../code/interfaces/types_app.AppCallParams.md):

- Required:

  - `appId: number` - The ID of the app to call
  - `callType: AppCallType | algosdk.OnApplicationComplete` - The [on-completion action](https://developer.algorand.org/docs/get-details/dapps/smart-contracts/apps/#the-lifecycle-of-a-smart-contract) for the call (either as an `algosdk.OnApplicationComplete` enum or a [string enum](../code/modules/types_app.md#appcalltype))
  - `from: SendTransactionFrom` - The [account](account.md) (with private key loaded) that will send the transaction

- Optional:
  - `transactionParams: SuggestedParams` - Any [transaction parameters](transaction.md#transaction-params)
  - `note: TransactionNote` - A [transaction note](transaction.md#transaction-notes)
  - `args: AppCallArgs` - Any [arguments](#appcallargs) passed in to the app call

If you pass in args that represent an ABI then it will use an `AtomicTransactionComposer` to construct and send the transaction(s). Because it's possible that other transactions may be present as ABI arguments, the full set of transactions that were sent are returned in `transactions` and the primary transaction for the create call will also be available in `transaction`. If you pass in the `atc` or `skipSending: true` then it won't execute the transaction and will simply return the transaction(s). The `return` value will have any ABI return value within it.

## Accessing state

### Global state

To access and parse global state you can use the following methods:

- [`algokit.getAppGlobalState(appId, algod)`](../code/modules/index.md#getappglobalstate) - Returns the current global state for the given app ID decoded into an object keyed by the UTF-8 representation of the state key with various parsed versions of the value (base64, UTF-8 and raw binary)
- [`decodeAppState(state)`](../code/modules/index.md#decodeappstate) - Takes the raw response from algod API for global state and returned a friendly decoded object (this is automatically used by `getAppGlobalState`)

### Local state

To access and parse local state you can use the following methods:

- [`algokit.getAppLocalState(appId, algod)`](../code/modules/index.md#getapplocalstate) - Returns the current local state for the given app ID decoded into an object keyed by the UTF-8 representation of the state key with various parsed versions of the value (base64, UTF-8 and raw binary)
- [`decodeAppState(state)`](../code/modules/index.md#decodeappstate) - Takes the raw response from algod API for local state and returned a friendly decoded object (this is automatically used by `getAppLocalState`)

### Boxes

To access and parse box values and names for an app you can use the following methods:

- [`algokit.getAppBoxNames(appId, algod)`](../code/modules/index.md#getappboxnames) - Returns the current [box names](#boxname) for the given app ID
- [`algokit.getAppBoxValue(appId, boxName, algod)`](../code/modules/index.md#getappboxvalue) - Returns the binary value of the given box name for the given app ID
- [`algokit.getAppBoxValues(appId, boxNames, algod)`](../code/modules/index.md#getappboxvalues) - Returns the binary values of the given box names for the given app ID
- [`algokit.getAppBoxValueFromABIType(request, algod)`](../code/modules/index.md#getappboxvaluefromabitype) - Returns the parsed ABI value of the given box name for the given app ID for the provided ABI type
- [`algokit.getAppBoxValuesFromABIType(request, algod)`](../code/modules/index.md#getappboxvaluesfromabitype) - Returns the parsed ABI values of the given box names for the given app ID for the provided ABI type
- [`algokit.getBoxReference(box)`](../code/modules/index.md#getboxreference) - Returns an `algosdk.BoxReference` representation of the given [box identifier](#referencing-boxes), which is useful when constructing a raw `algosdk.Transaction`

## Getting an app reference

To get reference information and metadata about an existing app you can use the following methods:

- [`algokit.getAppById(appId, algod)`](../code/modules/index.md#getappbyindex) - Returns an app reference by ID from algod
- [`algokit.lookupAccountCreatedApplicationByAddress(indexer, address, getAll?, paginationLimit?)`](../code/modules/index.md#lookupaccountcreatedapplicationbyaddress) - Returns all apps created by a given account from indexer
