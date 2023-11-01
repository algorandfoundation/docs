title: Overview

### Variables

- [Config](index.md#config)

### Functions

- [algos](index.md#algos)
- [callApp](index.md#callapp)
- [capTransactionFee](index.md#captransactionfee)
- [compileTeal](index.md#compileteal)
- [controlFees](index.md#controlfees)
- [createApp](index.md#createapp)
- [decodeAppState](index.md#decodeappstate)
- [deployApp](index.md#deployapp)
- [encodeTransactionNote](index.md#encodetransactionnote)
- [ensureFunded](index.md#ensurefunded)
- [executePaginatedRequest](index.md#executepaginatedrequest)
- [getABIMethodSignature](index.md#getabimethodsignature)
- [getABIReturn](index.md#getabireturn)
- [getAccount](index.md#getaccount)
- [getAccountAddressAsString](index.md#getaccountaddressasstring)
- [getAccountAddressAsUint8Array](index.md#getaccountaddressasuint8array)
- [getAccountConfigFromEnvironment](index.md#getaccountconfigfromenvironment)
- [getAlgoClient](index.md#getalgoclient)
- [getAlgoIndexerClient](index.md#getalgoindexerclient)
- [getAlgoKmdClient](index.md#getalgokmdclient)
- [getAlgoNodeConfig](index.md#getalgonodeconfig)
- [getAlgodConfigFromEnvironment](index.md#getalgodconfigfromenvironment)
- [getAppArgsForABICall](index.md#getappargsforabicall)
- [getAppArgsForTransaction](index.md#getappargsfortransaction)
- [getAppBoxNames](index.md#getappboxnames)
- [getAppBoxValue](index.md#getappboxvalue)
- [getAppBoxValueFromABIType](index.md#getappboxvaluefromabitype)
- [getAppBoxValues](index.md#getappboxvalues)
- [getAppBoxValuesFromABIType](index.md#getappboxvaluesfromabitype)
- [getAppById](index.md#getappbyid)
- [getAppClient](index.md#getappclient)
- [getAppClientByCreatorAndName](index.md#getappclientbycreatorandname)
- [getAppClientById](index.md#getappclientbyid)
- [getAppDeploymentTransactionNote](index.md#getappdeploymenttransactionnote)
- [getAppGlobalState](index.md#getappglobalstate)
- [getAppLocalState](index.md#getapplocalstate)
- [getAppOnCompleteAction](index.md#getapponcompleteaction)
- [getAtomicTransactionComposerTransactions](index.md#getatomictransactioncomposertransactions)
- [getBoxReference](index.md#getboxreference)
- [getConfigFromEnvOrDefaults](index.md#getconfigfromenvordefaults)
- [getCreatorAppsByName](index.md#getcreatorappsbyname)
- [getDefaultLocalNetConfig](index.md#getdefaultlocalnetconfig)
- [getDispenserAccount](index.md#getdispenseraccount)
- [getIndexerConfigFromEnvironment](index.md#getindexerconfigfromenvironment)
- [getKmdWalletAccount](index.md#getkmdwalletaccount)
- [getLocalNetDispenserAccount](index.md#getlocalnetdispenseraccount)
- [getOrCreateKmdWalletAccount](index.md#getorcreatekmdwalletaccount)
- [getSenderAddress](index.md#getsenderaddress)
- [getSenderTransactionSigner](index.md#getsendertransactionsigner)
- [getTestNetDispenserApiClient](index.md#gettestnetdispenserapiclient)
- [getTransactionParams](index.md#gettransactionparams)
- [getTransactionWithSigner](index.md#gettransactionwithsigner)
- [isLocalNet](index.md#islocalnet)
- [isMainNet](index.md#ismainnet)
- [isSchemaIsBroken](index.md#isschemaisbroken)
- [isTestNet](index.md#istestnet)
- [lookupAccountByAddress](index.md#lookupaccountbyaddress)
- [lookupAccountCreatedApplicationByAddress](index.md#lookupaccountcreatedapplicationbyaddress)
- [lookupTransactionById](index.md#lookuptransactionbyid)
- [microAlgos](index.md#microalgos)
- [mnemonicAccount](index.md#mnemonicaccount)
- [mnemonicAccountFromEnvironment](index.md#mnemonicaccountfromenvironment)
- [multisigAccount](index.md#multisigaccount)
- [optIn](index.md#optin)
- [optOut](index.md#optout)
- [performAtomicTransactionComposerDryrun](index.md#performatomictransactioncomposerdryrun)
- [performAtomicTransactionComposerSimulate](index.md#performatomictransactioncomposersimulate)
- [performTemplateSubstitution](index.md#performtemplatesubstitution)
- [performTemplateSubstitutionAndCompile](index.md#performtemplatesubstitutionandcompile)
- [randomAccount](index.md#randomaccount)
- [rekeyedAccount](index.md#rekeyedaccount)
- [replaceDeployTimeControlParams](index.md#replacedeploytimecontrolparams)
- [searchTransactions](index.md#searchtransactions)
- [sendAtomicTransactionComposer](index.md#sendatomictransactioncomposer)
- [sendGroupOfTransactions](index.md#sendgroupoftransactions)
- [sendTransaction](index.md#sendtransaction)
- [signTransaction](index.md#signtransaction)
- [stripTealComments](index.md#striptealcomments)
- [transactionFees](index.md#transactionfees)
- [transactionSignerAccount](index.md#transactionsigneraccount)
- [transferAlgos](index.md#transferalgos)
- [transferAsset](index.md#transferasset)
- [updateApp](index.md#updateapp)
- [waitForConfirmation](index.md#waitforconfirmation)

## Variables

### Config

• `Const` **Config**: [`UpdatableConfig`](../classes/types_config.UpdatableConfig.md)

The AlgoKit config. To update it use the configure method.

#### Defined in

[src/index.ts:17](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/index.ts#L17)

## Functions

### algos

▸ **algos**(`algos`): [`AlgoAmount`](../classes/types_amount.AlgoAmount.md)

Returns an amount of Algos using AlgoAmount

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `algos` | `number` | The amount in Algos |

#### Returns

[`AlgoAmount`](../classes/types_amount.AlgoAmount.md)

#### Defined in

[src/amount.ts:22](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/amount.ts#L22)

___

### callApp

▸ **callApp**(`call`, `algod`): `Promise`<[`AppCallTransactionResult`](../interfaces/types_app.AppCallTransactionResult.md)\>

Issues a call to a given app.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `call` | [`AppCallParams`](../interfaces/types_app.AppCallParams.md) | The call details. |
| `algod` | `default` | An algod client |

#### Returns

`Promise`<[`AppCallTransactionResult`](../interfaces/types_app.AppCallTransactionResult.md)\>

The result of the call

#### Defined in

[src/app.ts:305](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/app.ts#L305)

___

### capTransactionFee

▸ **capTransactionFee**(`transaction`, `maxAcceptableFee`): `void`

Limit the acceptable fee to a defined amount of µALGOs.
This also sets the transaction to be flatFee to ensure the transaction only succeeds at
the estimated rate.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `transaction` | `SuggestedParams` \| `Transaction` | The transaction to cap or suggested params object about to be used to create a transaction |
| `maxAcceptableFee` | [`AlgoAmount`](../classes/types_amount.AlgoAmount.md) | The maximum acceptable fee to pay |

#### Returns

`void`

#### Defined in

[src/transaction.ts:426](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/transaction.ts#L426)

___

### compileTeal

▸ **compileTeal**(`tealCode`, `algod`): `Promise`<[`CompiledTeal`](../interfaces/types_app.CompiledTeal.md)\>

Compiles the given TEAL using algod and returns the result, including source map.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `tealCode` | `string` | The TEAL code |
| `algod` | `default` | An algod client |

#### Returns

`Promise`<[`CompiledTeal`](../interfaces/types_app.CompiledTeal.md)\>

The information about the compiled file

#### Defined in

[src/app.ts:672](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/app.ts#L672)

___

### controlFees

▸ **controlFees**<`T`\>(`transaction`, `feeControl`): `T`

Allows for control of fees on a `Transaction` or `SuggestedParams` object

#### Type parameters

| Name | Type |
| :------ | :------ |
| `T` | extends `SuggestedParams` \| `Transaction` |

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `transaction` | `T` | The transaction or suggested params |
| `feeControl` | `Object` | The fee control parameters |
| `feeControl.fee?` | [`AlgoAmount`](../classes/types_amount.AlgoAmount.md) | - |
| `feeControl.maxFee?` | [`AlgoAmount`](../classes/types_amount.AlgoAmount.md) | - |

#### Returns

`T`

#### Defined in

[src/transaction.ts:449](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/transaction.ts#L449)

___

### createApp

▸ **createApp**(`create`, `algod`): `Promise`<`Partial`<[`AppCompilationResult`](../interfaces/types_app.AppCompilationResult.md)\> & [`AppCallTransactionResult`](../interfaces/types_app.AppCallTransactionResult.md) & [`AppReference`](../interfaces/types_app.AppReference.md)\>

Creates a smart contract app, returns the details of the created app.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `create` | [`CreateAppParams`](../interfaces/types_app.CreateAppParams.md) | The parameters to create the app with |
| `algod` | `default` | An algod client |

#### Returns

`Promise`<`Partial`<[`AppCompilationResult`](../interfaces/types_app.AppCompilationResult.md)\> & [`AppCallTransactionResult`](../interfaces/types_app.AppCallTransactionResult.md) & [`AppReference`](../interfaces/types_app.AppReference.md)\>

The details of the created app, or the transaction to create it if `skipSending` and the compilation result

#### Defined in

[src/app.ts:56](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/app.ts#L56)

___

### decodeAppState

▸ **decodeAppState**(`state`): [`AppState`](../interfaces/types_app.AppState.md)

Converts an array of global/local state values from the algod api to a more friendly
generic object keyed by the UTF-8 value of the key.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `state` | { `key`: `string` ; `value`: `TealValue` \| `EvalDelta`  }[] | A `global-state`, `local-state`, `global-state-deltas` or `local-state-deltas` |

#### Returns

[`AppState`](../interfaces/types_app.AppState.md)

An object keyeed by the UTF-8 representation of the key with various parsings of the values

#### Defined in

[src/app.ts:522](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/app.ts#L522)

___

### deployApp

▸ **deployApp**(`deployment`, `algod`, `indexer?`): `Promise`<`Partial`<[`AppCompilationResult`](../interfaces/types_app.AppCompilationResult.md)\> & [`ConfirmedTransactionResults`](../interfaces/types_transaction.ConfirmedTransactionResults.md) & [`AppMetadata`](../interfaces/types_app.AppMetadata.md) & { `operationPerformed`: ``"create"`` \| ``"update"`` ; `return?`: [`ABIReturn`](types_app.md#abireturn)  } \| [`ConfirmedTransactionResults`](../interfaces/types_transaction.ConfirmedTransactionResults.md) & [`AppMetadata`](../interfaces/types_app.AppMetadata.md) & { `deleteResult`: [`ConfirmedTransactionResult`](../interfaces/types_transaction.ConfirmedTransactionResult.md) ; `deleteReturn?`: [`ABIReturn`](types_app.md#abireturn) ; `operationPerformed`: ``"replace"`` ; `return?`: [`ABIReturn`](types_app.md#abireturn)  } \| [`AppMetadata`](../interfaces/types_app.AppMetadata.md) & { `operationPerformed`: ``"nothing"``  }\>

Idempotently deploy (create, update/delete if changed) an app against the given name via the given creator account, including deploy-time template placeholder substitutions.

To understand the architecture decisions behind this functionality please see https://github.com/algorandfoundation/algokit-cli/blob/main/docs/architecture-decisions/2023-01-12_smart-contract-deployment.md

**Note:** When using the return from this function be sure to check `operationPerformed` to get access to various return properties like `transaction`, `confirmation` and `deleteResult`.

**Note:** if there is a breaking state schema change to an existing app (and `onSchemaBreak` is set to `'replace'`) the existing app will be deleted and re-created.

**Note:** if there is an update (different TEAL code) to an existing app (and `onUpdate` is set to `'replace'`) the existing app will be deleted and re-created.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `deployment` | [`AppDeploymentParams`](../interfaces/types_app.AppDeploymentParams.md) | The arguments to control the app deployment |
| `algod` | `default` | An algod client |
| `indexer?` | `default` | An indexer client, needed if `existingDeployments` not passed in |

#### Returns

`Promise`<`Partial`<[`AppCompilationResult`](../interfaces/types_app.AppCompilationResult.md)\> & [`ConfirmedTransactionResults`](../interfaces/types_transaction.ConfirmedTransactionResults.md) & [`AppMetadata`](../interfaces/types_app.AppMetadata.md) & { `operationPerformed`: ``"create"`` \| ``"update"`` ; `return?`: [`ABIReturn`](types_app.md#abireturn)  } \| [`ConfirmedTransactionResults`](../interfaces/types_transaction.ConfirmedTransactionResults.md) & [`AppMetadata`](../interfaces/types_app.AppMetadata.md) & { `deleteResult`: [`ConfirmedTransactionResult`](../interfaces/types_transaction.ConfirmedTransactionResult.md) ; `deleteReturn?`: [`ABIReturn`](types_app.md#abireturn) ; `operationPerformed`: ``"replace"`` ; `return?`: [`ABIReturn`](types_app.md#abireturn)  } \| [`AppMetadata`](../interfaces/types_app.AppMetadata.md) & { `operationPerformed`: ``"nothing"``  }\>

The app reference of the new/existing app

#### Defined in

[src/app-deploy.ts:38](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/app-deploy.ts#L38)

___

### encodeTransactionNote

▸ **encodeTransactionNote**(`note?`): `Uint8Array` \| `undefined`

Encodes a transaction note into a byte array ready to be included in an Algorand transaction.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `note?` | [`TransactionNote`](types_transaction.md#transactionnote) | The transaction note |

#### Returns

`Uint8Array` \| `undefined`

the transaction note ready for inclusion in a transaction

 Case on the value of `data` this either be:
  * `null` | `undefined`: `undefined`
  * `string`: The string value
  * Uint8Array: passthrough
  * Arc2TransactionNote object: ARC-0002 compatible transaction note
  * Else: The object/value converted into a JSON string representation

#### Defined in

[src/transaction.ts:39](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/transaction.ts#L39)

___

### ensureFunded

▸ **ensureFunded**<`T`\>(`funding`, `algod`, `kmd?`): `Promise`<[`EnsureFundedReturnType`](../interfaces/types_transfer.EnsureFundedReturnType.md) \| `undefined`\>

Funds a given account using a funding source such that it has a certain amount of algos free to spend (accounting for ALGOs locked in minimum balance requirement).

https://developer.algorand.org/docs/get-details/accounts/#minimum-balance

#### Type parameters

| Name | Type |
| :------ | :------ |
| `T` | extends [`EnsureFundedParams`](../interfaces/types_transfer.EnsureFundedParams.md) |

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `funding` | `T` | The funding configuration of type `EnsureFundedParams`, including the account to fund, minimum spending balance, and optional parameters. If you set `useDispenserApi` to true, you must also set `ALGOKIT_DISPENSER_ACCESS_TOKEN` in your environment variables. |
| `algod` | `default` | An instance of the Algodv2 client. |
| `kmd?` | `default` | An optional instance of the Kmd client. |

#### Returns

`Promise`<[`EnsureFundedReturnType`](../interfaces/types_transfer.EnsureFundedReturnType.md) \| `undefined`\>

- `EnsureFundedReturnType` if funds were transferred.
- `undefined` if no funds were needed.

#### Defined in

[src/transfer.ts:114](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/transfer.ts#L114)

___

### executePaginatedRequest

▸ **executePaginatedRequest**<`TResult`, `TRequest`\>(`extractItems`, `buildRequest`): `Promise`<`TResult`[]\>

#### Type parameters

| Name | Type |
| :------ | :------ |
| `TResult` | `TResult` |
| `TRequest` | extends `Object` |

#### Parameters

| Name | Type |
| :------ | :------ |
| `extractItems` | (`response`: `any`) => `TResult`[] |
| `buildRequest` | (`nextToken?`: `string`) => `TRequest` |

#### Returns

`Promise`<`TResult`[]\>

#### Defined in

[src/indexer-lookup.ts:108](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/indexer-lookup.ts#L108)

___

### getABIMethodSignature

▸ **getABIMethodSignature**(`method`): `string`

Returns the encoded ABI spec for a given ABI Method

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `method` | `ABIMethodParams` \| `ABIMethod` | The method to return a signature for |

#### Returns

`string`

The encoded ABI method spec e.g. `method_name(uint64,string)string`

#### Defined in

[src/app.ts:688](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/app.ts#L688)

___

### getABIReturn

▸ **getABIReturn**(`args?`, `confirmation?`): [`ABIReturn`](types_app.md#abireturn) \| `undefined`

Returns any ABI return values for the given app call arguments and transaction confirmation.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `args?` | [`AppCallArgs`](types_app.md#appcallargs) | The arguments that were used for the call |
| `confirmation?` | `PendingTransactionResponse` | The transaction confirmation from algod |

#### Returns

[`ABIReturn`](types_app.md#abireturn) \| `undefined`

The return value for the method call

#### Defined in

[src/app.ts:386](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/app.ts#L386)

___

### getAccount

▸ **getAccount**(`account`, `algod`, `kmdClient?`): `Promise`<`Account` \| [`SigningAccount`](../classes/types_account.SigningAccount.md)\>

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `account` | `string` \| { `fundWith?`: [`AlgoAmount`](../classes/types_amount.AlgoAmount.md) ; `name`: `string`  } | The details of the account to get, either the name identifier (string) or an object with: * `name`: The name identifier of the account * `fundWith`: The amount to fund the account with when it gets created (when targeting LocalNet), if not specified then 1000 Algos will be funded from the dispenser account |
| `algod` | `default` | An algod client |
| `kmdClient?` | `default` | An optional KMD client to use to create an account (when targeting LocalNet), if not specified then a default KMD client will be loaded from environment variables |

#### Returns

`Promise`<`Account` \| [`SigningAccount`](../classes/types_account.SigningAccount.md)\>

The requested account with private key loaded from the environment variables or when targeting LocalNet from KMD (idempotently creating and funding the account)

**`Deprecated`**

use mnemonicAccountFromEnvironment instead

Returns an Algorand account with private key loaded by convention based on the given name identifier.

Note: This function expects to run in a Node.js environment.

## Convention:
* **Non-LocalNet:** will load process.env['{NAME}_MNEMONIC'] as a mnemonic secret; **Note: Be careful how the mnemonic is handled**,
 never commit it into source control and ideally load it via a secret storage service rather than the file system.
  If process.env['{NAME}_SENDER'] is defined then it will use that for the sender address (i.e. to support rekeyed accounts)
* **LocalNet:** will load the account from a KMD wallet called {NAME} and if that wallet doesn't exist it will create it and fund the account for you

This allows you to write code that will work seamlessly in production and local development (LocalNet) without manual config locally (including when you reset the LocalNet).

**`Example`**

Default

If you have a mnemonic secret loaded into `process.env.ACCOUNT_MNEMONIC` then you can call the following to get that private key loaded into an account object:
```typescript
const account = await getAccount('ACCOUNT', algod)
```

If that code runs against LocalNet then a wallet called `ACCOUNT` will automatically be created with an account that is automatically funded with 1000 (default) ALGOs from the default LocalNet dispenser.

#### Defined in

[src/account.ts:92](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/account.ts#L92)

▸ **getAccount**(`account`, `algod`, `kmdClient?`): `Promise`<`Account` \| [`SigningAccount`](../classes/types_account.SigningAccount.md)\>

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `account` | `Object` | The details of the account to get, an object with: * `config`: Account configuration. To get from environment use getAccountConfigFromEnvironment(accountName) * `fundWith`: The amount to fund the account with when it gets created (when targeting LocalNet), if not specified then 1000 Algos will be funded from the dispenser account |
| `account.config` | [`AccountConfig`](../interfaces/types_account.AccountConfig.md) | - |
| `account.fundWith?` | [`AlgoAmount`](../classes/types_amount.AlgoAmount.md) | - |
| `algod` | `default` | An algod client |
| `kmdClient?` | `default` | An optional KMD client to use to create an account (when targeting LocalNet), if not specified then a default KMD client will be loaded from environment variables |

#### Returns

`Promise`<`Account` \| [`SigningAccount`](../classes/types_account.SigningAccount.md)\>

The requested account with private key loaded from the environment variables or when targeting LocalNet from KMD (idempotently creating and funding the account)

**`Deprecated`**

use mnemonicAccountFromEnvironment instead
Returns an Algorand account with private key loaded by convention based on the given name identifier.

Note: This function expects to run in a Node.js environment.

**`Example`**

Default

If you have a mnemonic secret loaded into `process.env.ACCOUNT_MNEMONIC` then you can call the following to get that private key loaded into an account object:
```typescript
const account = await getAccount({config: getAccountConfigFromEnvironment('ACCOUNT')}, algod)
```

If that code runs against LocalNet then a wallet called `ACCOUNT` will automatically be created with an account that is automatically funded with 1000 (default) ALGOs from the default LocalNet dispenser.

#### Defined in

[src/account.ts:119](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/account.ts#L119)

___

### getAccountAddressAsString

▸ **getAccountAddressAsString**(`addressEncodedInB64`): `string`

Returns the string address of an Algorand account from a base64 encoded version of the underlying byte array of the address public key

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `addressEncodedInB64` | `string` | The base64 encoded version of the underlying byte array of the address public key |

#### Returns

`string`

#### Defined in

[src/account.ts:270](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/account.ts#L270)

___

### getAccountAddressAsUint8Array

▸ **getAccountAddressAsUint8Array**(`account`): `Uint8Array`

Returns an account's address as a byte array

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `account` | `string` \| [`SendTransactionFrom`](types_transaction.md#sendtransactionfrom) | Either an account (with private key loaded) or the string address of an account |

#### Returns

`Uint8Array`

#### Defined in

[src/account.ts:262](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/account.ts#L262)

___

### getAccountConfigFromEnvironment

▸ **getAccountConfigFromEnvironment**(`accountName`): [`AccountConfig`](../interfaces/types_account.AccountConfig.md)

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `accountName` | `string` | account name |

#### Returns

[`AccountConfig`](../interfaces/types_account.AccountConfig.md)

**`Deprecated`**

Use algokit.mnemonicAccountFromEnvironment, which doesn't need this function
Returns the Account configuration from environment variables

**`Example`**

```ts
environment variables
{accountName}_MNEMONIC
{accountName}_SENDER
```

#### Defined in

[src/account.ts:298](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/account.ts#L298)

___

### getAlgoClient

▸ **getAlgoClient**(`config?`): `Algodv2`

Returns an algod SDK client that automatically retries on idempotent calls

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `config?` | [`AlgoClientConfig`](../interfaces/types_network_client.AlgoClientConfig.md) | The config if you want to override the default (getting config from process.env) |

#### Returns

`Algodv2`

**`Example`**

Default (load from environment variables)

 ```typescript
 // Uses process.env.ALGOD_SERVER, process.env.ALGOD_PORT and process.env.ALGOD_TOKEN
 // Automatically detects if you are using PureStake to switch in the right header name for ALGOD_TOKEN
 const algod = getAlgoClient()
 await algod.healthCheck().do()
 ```

**`Example`**

AlgoNode (testnet)
```typescript
 const algod = getAlgoClient(getAlgoNodeConfig('testnet', 'algod'))
 await algod.healthCheck().do()
```

**`Example`**

AlgoNode (mainnet)
```typescript
 const algod = getAlgoClient(getAlgoNodeConfig('mainnet', 'algod'))
 await algod.healthCheck().do()
```

**`Example`**

Custom (e.g. default LocalNet, although we recommend loading this into a .env and using the Default option instead)
```typescript
 const algod = getAlgoClient({server: 'http://localhost', port: '4001', token: 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa'})
 await algod.healthCheck().do()
```

#### Defined in

[src/network-client.ts:126](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/network-client.ts#L126)

___

### getAlgoIndexerClient

▸ **getAlgoIndexerClient**(`config?`): `Indexer`

Returns an indexer SDK client that automatically retries on idempotent calls

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `config?` | [`AlgoClientConfig`](../interfaces/types_network_client.AlgoClientConfig.md) | The config if you want to override the default (getting config from process.env) |

#### Returns

`Indexer`

**`Example`**

Default (load from environment variables)

 ```typescript
 // Uses process.env.INDEXER_SERVER, process.env.INDEXER_PORT and process.env.INDEXER_TOKEN
 // Automatically detects if you are using PureStake to switch in the right header name for INDEXER_TOKEN
 const indexer = getAlgoIndexerClient()
 await indexer.makeHealthCheck().do()
 ```

**`Example`**

AlgoNode (testnet)
```typescript
 const indexer = getAlgoIndexerClient(getAlgoNodeConfig('testnet', 'indexer'))
 await indexer.makeHealthCheck().do()
```

**`Example`**

AlgoNode (mainnet)
```typescript
 const indexer = getAlgoIndexerClient(getAlgoNodeConfig('mainnet', 'indexer'))
 await indexer.makeHealthCheck().do()
```

**`Example`**

Custom (e.g. default LocalNet, although we recommend loading this into a .env and using the Default option instead)
```typescript
 const indexer = getAlgoIndexerClient({server: 'http://localhost', port: '8980', token: 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa'})
 await indexer.makeHealthCheck().do()
```

#### Defined in

[src/network-client.ts:159](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/network-client.ts#L159)

___

### getAlgoKmdClient

▸ **getAlgoKmdClient**(`config?`): `Kmd`

Returns a KMD SDK client that automatically retries on idempotent calls

KMD client allows you to export private keys, which is useful to get the default account in a LocalNet network.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `config?` | [`AlgoClientConfig`](../interfaces/types_network_client.AlgoClientConfig.md) | The config if you want to override the default (getting config from process.env) |

#### Returns

`Kmd`

**`Example`**

Default (load from environment variables)

 ```typescript
 // Uses process.env.ALGOD_SERVER, process.env.KMD_PORT (or if not specified: port 4002) and process.env.ALGOD_TOKEN
 const kmd = getAlgoKmdClient()
 ```

**`Example`**

Custom (e.g. default LocalNet, although we recommend loading this into a .env and using the Default option instead)
```typescript
 const kmd = getAlgoKmdClient({server: 'http://localhost', port: '4002', token: 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa'})
```

#### Defined in

[src/network-client.ts:182](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/network-client.ts#L182)

___

### getAlgoNodeConfig

▸ **getAlgoNodeConfig**(`network`, `config`): [`AlgoClientConfig`](../interfaces/types_network_client.AlgoClientConfig.md)

Returns the Algorand configuration to point to the AlgoNode service

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `network` | ``"testnet"`` \| ``"mainnet"`` | Which network to connect to - TestNet or MainNet |
| `config` | ``"algod"`` \| ``"indexer"`` | Which algod config to return - Algod or Indexer |

#### Returns

[`AlgoClientConfig`](../interfaces/types_network_client.AlgoClientConfig.md)

#### Defined in

[src/network-client.ts:72](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/network-client.ts#L72)

___

### getAlgodConfigFromEnvironment

▸ **getAlgodConfigFromEnvironment**(): [`AlgoClientConfig`](../interfaces/types_network_client.AlgoClientConfig.md)

Retrieve the algod configuration from environment variables (expects to be called from a Node.js environment not algod-side)

#### Returns

[`AlgoClientConfig`](../interfaces/types_network_client.AlgoClientConfig.md)

#### Defined in

[src/network-client.ts:34](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/network-client.ts#L34)

___

### getAppArgsForABICall

▸ **getAppArgsForABICall**(`args`, `from`): `Promise`<{ `appAccounts`: `undefined` \| `string`[] ; `appForeignApps`: `undefined` \| `number`[] = args.apps; `appForeignAssets`: `undefined` \| `number`[] = args.assets; `boxes`: `undefined` \| `BoxReference`[] ; `lease`: `undefined` \| `Uint8Array` ; `method`: `ABIMethod` ; `methodArgs`: (`string` \| `number` \| `bigint` \| `boolean` \| `Uint8Array` \| `ABIValue`[] \| `TransactionWithSigner`)[] = methodArgs; `rekeyTo`: `undefined` = undefined; `sender`: `string` ; `signer`: `TransactionSigner` = signer }\>

Returns the app args ready to load onto an ABI method call in `AtomicTransactionComposer`

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `args` | [`ABIAppCallArgs`](types_app.md#abiappcallargs) | The ABI app call args |
| `from` | [`SendTransactionFrom`](types_transaction.md#sendtransactionfrom) | The transaction signer |

#### Returns

`Promise`<{ `appAccounts`: `undefined` \| `string`[] ; `appForeignApps`: `undefined` \| `number`[] = args.apps; `appForeignAssets`: `undefined` \| `number`[] = args.assets; `boxes`: `undefined` \| `BoxReference`[] ; `lease`: `undefined` \| `Uint8Array` ; `method`: `ABIMethod` ; `methodArgs`: (`string` \| `number` \| `bigint` \| `boolean` \| `Uint8Array` \| `ABIValue`[] \| `TransactionWithSigner`)[] = methodArgs; `rekeyTo`: `undefined` = undefined; `sender`: `string` ; `signer`: `TransactionSigner` = signer }\>

The parameters ready to pass into `addMethodCall` within AtomicTransactionComposer

#### Defined in

[src/app.ts:589](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/app.ts#L589)

___

### getAppArgsForTransaction

▸ **getAppArgsForTransaction**(`args?`): `undefined` \| { `accounts`: `undefined` \| `string`[] ; `appArgs`: `undefined` \| `Uint8Array`[] ; `boxes`: `undefined` \| `BoxReference`[] ; `foreignApps`: `undefined` \| `number`[] = args.apps; `foreignAssets`: `undefined` \| `number`[] = args.assets; `lease`: `undefined` \| `Uint8Array`  }

Returns the app args ready to load onto an app `Transaction` object

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `args?` | [`RawAppCallArgs`](../interfaces/types_app.RawAppCallArgs.md) | The app call args |

#### Returns

`undefined` \| { `accounts`: `undefined` \| `string`[] ; `appArgs`: `undefined` \| `Uint8Array`[] ; `boxes`: `undefined` \| `BoxReference`[] ; `foreignApps`: `undefined` \| `number`[] = args.apps; `foreignAssets`: `undefined` \| `number`[] = args.assets; `lease`: `undefined` \| `Uint8Array`  }

The args ready to load into a `Transaction`

#### Defined in

[src/app.ts:569](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/app.ts#L569)

___

### getAppBoxNames

▸ **getAppBoxNames**(`appId`, `algod`): `Promise`<[`BoxName`](../interfaces/types_app.BoxName.md)[]\>

Returns the names of the boxes for the given app.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `appId` | `number` \| `bigint` | The ID of the app return box names for |
| `algod` | `default` | An algod client instance |

#### Returns

`Promise`<[`BoxName`](../interfaces/types_app.BoxName.md)[]\>

The current box names

#### Defined in

[src/app.ts:458](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/app.ts#L458)

___

### getAppBoxValue

▸ **getAppBoxValue**(`appId`, `boxName`, `algod`): `Promise`<`Uint8Array`\>

Returns the value of the given box name for the given app.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `appId` | `number` \| `bigint` | The ID of the app return box names for |
| `boxName` | `string` \| `Uint8Array` \| [`BoxName`](../interfaces/types_app.BoxName.md) | The name of the box to return either as a string, binary array or `BoxName` |
| `algod` | `default` | An algod client instance |

#### Returns

`Promise`<`Uint8Array`\>

The current box value as a byte array

#### Defined in

[src/app.ts:476](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/app.ts#L476)

___

### getAppBoxValueFromABIType

▸ **getAppBoxValueFromABIType**(`request`, `algod`): `Promise`<`ABIValue`\>

Returns the value of the given box name for the given app decoded based on the given ABI type.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `request` | [`BoxValueRequestParams`](../interfaces/types_app.BoxValueRequestParams.md) | The parameters for the box value request |
| `algod` | `default` | An algod client instance |

#### Returns

`Promise`<`ABIValue`\>

The current box value as an ABI value

#### Defined in

[src/app.ts:499](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/app.ts#L499)

___

### getAppBoxValues

▸ **getAppBoxValues**(`appId`, `boxNames`, `algod`): `Promise`<`Uint8Array`[]\>

Returns the value of the given box names for the given app.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `appId` | `number` | The ID of the app return box names for |
| `boxNames` | (`string` \| `Uint8Array` \| [`BoxName`](../interfaces/types_app.BoxName.md))[] | The names of the boxes to return either as a string, binary array or `BoxName` |
| `algod` | `default` | An algod client instance |

#### Returns

`Promise`<`Uint8Array`[]\>

The current box values as a byte array in the same order as the passed in box names

#### Defined in

[src/app.ts:489](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/app.ts#L489)

___

### getAppBoxValuesFromABIType

▸ **getAppBoxValuesFromABIType**(`request`, `algod`): `Promise`<`ABIValue`[]\>

Returns the value of the given box names for the given app decoded based on the given ABI type.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `request` | [`BoxValuesRequestParams`](../interfaces/types_app.BoxValuesRequestParams.md) | The parameters for the box value request |
| `algod` | `default` | An algod client instance |

#### Returns

`Promise`<`ABIValue`[]\>

The current box values as an ABI value in the same order as the passed in box names

#### Defined in

[src/app.ts:511](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/app.ts#L511)

___

### getAppById

▸ **getAppById**(`appId`, `algod`): `Promise`<`Application`\>

Gets the current data for the given app from algod.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `appId` | `number` \| `bigint` | The id of the app |
| `algod` | `default` | An algod client |

#### Returns

`Promise`<`Application`\>

The data about the app

#### Defined in

[src/app.ts:661](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/app.ts#L661)

___

### getAppClient

▸ **getAppClient**(`appDetails`, `algod`): [`ApplicationClient`](../classes/types_app_client.ApplicationClient.md)

Create a new ApplicationClient instance

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `appDetails` | [`AppSpecAppDetails`](types_app_client.md#appspecappdetails) | The details of the app |
| `algod` | `default` | An algod instance |

#### Returns

[`ApplicationClient`](../classes/types_app_client.ApplicationClient.md)

The application client

**`Example`**

```ts
Resolve by creator and name
const client = algokit.getAppClient(
    {
      resolveBy: 'creatorAndName',
      app: {appSpec},
      sender: {account},
      creatorAddress: {creator},
      findExistingUsing: indexerClient,
    },
    algodClient,
  )
```

**`Example`**

```ts
Resolve by id:
const client = algokit.getAppClient(
    {
      resolveBy: 'id',
      app: {appSpec},
      sender: {account},
      id: {id},
    },
   algodClient,
)
```

#### Defined in

[src/app-client.ts:34](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/app-client.ts#L34)

___

### getAppClientByCreatorAndName

▸ **getAppClientByCreatorAndName**(`appDetails`, `algod`): [`ApplicationClient`](../classes/types_app_client.ApplicationClient.md)

Create a new ApplicationClient instance by creator and name

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `appDetails` | [`AppSpecAppDetailsByCreatorAndName`](types_app_client.md#appspecappdetailsbycreatorandname) | The details of the app by creator and name |
| `algod` | `default` | An algod instance |

#### Returns

[`ApplicationClient`](../classes/types_app_client.ApplicationClient.md)

The application client

**`Example`**

```ts
const client = algokit.getAppClientByCreatorAndName(
    {
      app: {appSpec},
      sender: {account},
      creatorAddress: {account.addr},
      findExistingUsing: {indexerClient},
    },
    algodClient,
  )
```

#### Defined in

[src/app-client.ts:77](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/app-client.ts#L77)

___

### getAppClientById

▸ **getAppClientById**(`appDetails`, `algod`): [`ApplicationClient`](../classes/types_app_client.ApplicationClient.md)

Create a new ApplicationClient instance by id

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `appDetails` | [`AppSpecAppDetailsById`](types_app_client.md#appspecappdetailsbyid) | The details of the app |
| `algod` | `default` | An algod instance |

#### Returns

[`ApplicationClient`](../classes/types_app_client.ApplicationClient.md)

The application client

**`Example`**

```ts
const client = algokit.getAppClientById(
    {
      app: {appSpec},
      sender: {account},
      id: {id},
    },
    algodClient,
  )
```

#### Defined in

[src/app-client.ts:55](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/app-client.ts#L55)

___

### getAppDeploymentTransactionNote

▸ **getAppDeploymentTransactionNote**(`metadata`): [`Arc2TransactionNote`](types_transaction.md#arc2transactionnote)

Return the transaction note for an app deployment.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `metadata` | [`AppDeployMetadata`](../interfaces/types_app.AppDeployMetadata.md) | The metadata of the deployment |

#### Returns

[`Arc2TransactionNote`](types_transaction.md#arc2transactionnote)

The transaction note as a utf-8 string

#### Defined in

[src/app-deploy.ts:528](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/app-deploy.ts#L528)

___

### getAppGlobalState

▸ **getAppGlobalState**(`appId`, `algod`): `Promise`<[`AppState`](../interfaces/types_app.AppState.md)\>

Returns the current global state values for the given app ID

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `appId` | `number` \| `bigint` | The ID of the app return global state for |
| `algod` | `default` | An algod client instance |

#### Returns

`Promise`<[`AppState`](../interfaces/types_app.AppState.md)\>

The current global state

#### Defined in

[src/app.ts:422](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/app.ts#L422)

___

### getAppLocalState

▸ **getAppLocalState**(`appId`, `account`, `algod`): `Promise`<[`AppState`](../interfaces/types_app.AppState.md)\>

Returns the current global state values for the given app ID and account

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `appId` | `number` \| `bigint` | The ID of the app return global state for |
| `account` | `string` \| [`SendTransactionFrom`](types_transaction.md#sendtransactionfrom) | Either the string address of an account or an account object for the account to get local state for the given app |
| `algod` | `default` | An algod client instance |

#### Returns

`Promise`<[`AppState`](../interfaces/types_app.AppState.md)\>

The current local state for the given (app, account) combination

#### Defined in

[src/app.ts:439](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/app.ts#L439)

___

### getAppOnCompleteAction

▸ **getAppOnCompleteAction**(`onCompletionAction?`): `OnApplicationComplete`

Returns an `algosdk.OnApplicationComplete` for the given onCompleteAction.

If given `undefined` will return `OnApplicationComplete.NoOpOC`.

If given an `AppCallType` will convert the string enum to the correct underlying `algosdk.OnApplicationComplete`.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `onCompletionAction?` | [`AppCallType`](types_app.md#appcalltype) \| `OnApplicationComplete` | The on completion action |

#### Returns

`OnApplicationComplete`

The `algosdk.OnApplicationComplete`

#### Defined in

[src/app.ts:275](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/app.ts#L275)

___

### getAtomicTransactionComposerTransactions

▸ **getAtomicTransactionComposerTransactions**(`atc`): `TransactionWithSigner`[]

Returns the array of transactions currently present in the given `AtomicTransactionComposer`

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `atc` | `AtomicTransactionComposer` | The atomic transaction composer |

#### Returns

`TransactionWithSigner`[]

The array of transactions with signers

#### Defined in

[src/transaction.ts:481](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/transaction.ts#L481)

___

### getBoxReference

▸ **getBoxReference**(`box`): `algosdk.BoxReference`

Returns a `algosdk.BoxReference` given a `BoxIdentifier` or `BoxReference`.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `box` | [`BoxReference`](../interfaces/types_app.BoxReference.md) \| [`BoxIdentifier`](types_app.md#boxidentifier) \| `BoxReference` | The box to return a reference for |

#### Returns

`algosdk.BoxReference`

The box reference ready to pass into a `Transaction`

#### Defined in

[src/app.ts:631](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/app.ts#L631)

___

### getConfigFromEnvOrDefaults

▸ **getConfigFromEnvOrDefaults**(): [`AlgoConfig`](../interfaces/types_network_client.AlgoConfig.md)

Retrieve configurations from environment variables when defined or get defaults (expects to be called from a Node.js environment not algod-side)

#### Returns

[`AlgoConfig`](../interfaces/types_network_client.AlgoConfig.md)

#### Defined in

[src/network-client.ts:7](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/network-client.ts#L7)

___

### getCreatorAppsByName

▸ **getCreatorAppsByName**(`creatorAccount`, `indexer`): `Promise`<[`AppLookup`](../interfaces/types_app.AppLookup.md)\>

Returns a lookup of name => app metadata (id, address, ...metadata) for all apps created by the given account that have an `AppDeployNote` in the transaction note of the creation transaction.

**Note:** It's recommended this is only called once and then stored since it's a somewhat expensive operation (multiple indexer calls).

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `creatorAccount` | `string` \| [`SendTransactionFrom`](types_transaction.md#sendtransactionfrom) | The account (with private key loaded) or string address of an account that is the creator of the apps you want to search for |
| `indexer` | `default` | An indexer client |

#### Returns

`Promise`<[`AppLookup`](../interfaces/types_app.AppLookup.md)\>

A name-based lookup of the app information (id, address)

#### Defined in

[src/app-deploy.ts:421](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/app-deploy.ts#L421)

___

### getDefaultLocalNetConfig

▸ **getDefaultLocalNetConfig**(`configOrPort`): [`AlgoClientConfig`](../interfaces/types_network_client.AlgoClientConfig.md)

Returns the Algorand configuration to point to the default LocalNet

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `configOrPort` | `number` \| ``"algod"`` \| ``"indexer"`` \| ``"kmd"`` | Which algod config to return - algod, kmd, or indexer OR a port number |

#### Returns

[`AlgoClientConfig`](../interfaces/types_network_client.AlgoClientConfig.md)

#### Defined in

[src/network-client.ts:83](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/network-client.ts#L83)

___

### getDispenserAccount

▸ **getDispenserAccount**(`algod`, `kmd?`): `Promise`<`default` \| [`SigningAccount`](../classes/types_account.SigningAccount.md)\>

Returns an account (with private key loaded) that can act as a dispenser

If running on LocalNet then it will return the default dispenser account automatically,
 otherwise it will load the account mnemonic stored in process.env.DISPENSER_MNEMONIC

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `algod` | `default` | An algod client |
| `kmd?` | `default` | A KMD client, if not specified then a default KMD client will be loaded from environment variables |

#### Returns

`Promise`<`default` \| [`SigningAccount`](../classes/types_account.SigningAccount.md)\>

#### Defined in

[src/account.ts:282](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/account.ts#L282)

___

### getIndexerConfigFromEnvironment

▸ **getIndexerConfigFromEnvironment**(): [`AlgoClientConfig`](../interfaces/types_network_client.AlgoClientConfig.md)

Retrieve the indexer configuration from environment variables (expects to be called from a Node.js environment not algod-side)

#### Returns

[`AlgoClientConfig`](../interfaces/types_network_client.AlgoClientConfig.md)

#### Defined in

[src/network-client.ts:51](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/network-client.ts#L51)

___

### getKmdWalletAccount

▸ **getKmdWalletAccount**(`walletAccount`, `algod`, `kmdClient?`): `Promise`<`Account` \| `undefined`\>

Returns an Algorand account with private key loaded from the given KMD wallet (identified by name).

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `walletAccount` | `Object` | The details of the wallet, with: * `name`: The name of the wallet to retrieve an account from * `predicate`: An optional filter to use to find the account (otherwise it will return a random account from the wallet) |
| `walletAccount.name` | `string` | - |
| `walletAccount.predicate?` | (`account`: `Record`<`string`, `any`\>) => `boolean` | - |
| `algod` | `default` | An algod client |
| `kmdClient?` | `default` | A KMD client, if not specified then a default KMD client will be loaded from environment variables |

#### Returns

`Promise`<`Account` \| `undefined`\>

**`Example`**

Get default funded account in a LocalNet

```typescript
const defaultDispenserAccount = await getKmdWalletAccount(algod,
  'unencrypted-default-wallet',
  a => a.status !== 'Offline' && a.amount > 1_000_000_000
)
```

#### Defined in

[src/localnet.ts:90](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/localnet.ts#L90)

___

### getLocalNetDispenserAccount

▸ **getLocalNetDispenserAccount**(`algod`, `kmd?`): `Promise`<`Account`\>

Returns an Algorand account with private key loaded for the default LocalNet dispenser account (that can be used to fund other accounts)

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `algod` | `default` | An algod client |
| `kmd?` | `default` | A KMD client, if not specified then a default KMD client will be loaded from environment variables |

#### Returns

`Promise`<`Account`\>

#### Defined in

[src/localnet.ts:141](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/localnet.ts#L141)

___

### getOrCreateKmdWalletAccount

▸ **getOrCreateKmdWalletAccount**(`walletAccount`, `algod`, `kmdClient?`): `Promise`<`Account`\>

Gets an account with private key loaded from a KMD wallet of the given name, or alternatively creates one with funds in it via a KMD wallet of the given name.

This is useful to get idempotent accounts from LocalNet without having to specify the private key (which will change when resetting the LocalNet).

This significantly speeds up local dev time and improves experience since you can write code that *just works* first go without manual config in a fresh LocalNet.

If this is used via `mnemonicAccountFromEnvironment`, then you can even use the same code that runs on production without changes for local development!

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `walletAccount` | `Object` | The wallet details with: * `name`: The name of the wallet to retrieve / create * `fundWith`: The number of Algos to fund the account with when it gets created, if not specified then 1000 Algos will be funded from the dispenser account |
| `walletAccount.fundWith?` | [`AlgoAmount`](../classes/types_amount.AlgoAmount.md) | - |
| `walletAccount.name` | `string` | - |
| `algod` | `default` | An algod client |
| `kmdClient?` | `default` | A KMD client, if not specified then a default KMD client will be loaded from environment variables |

#### Returns

`Promise`<`Account`\>

An Algorand account with private key loaded - either one that already existed in the given KMD wallet, or a new one that is funded for you

#### Defined in

[src/localnet.ts:32](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/localnet.ts#L32)

___

### getSenderAddress

▸ **getSenderAddress**(`sender`): `string`

Returns the public address of the given transaction sender.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `sender` | [`SendTransactionFrom`](types_transaction.md#sendtransactionfrom) | A transaction sender |

#### Returns

`string`

The public address

#### Defined in

[src/transaction.ts:60](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/transaction.ts#L60)

___

### getSenderTransactionSigner

▸ **getSenderTransactionSigner**(`val`): `TransactionSigner`

Returns a `TransactionSigner` for the given transaction sender.
This function has memoization, so will return the same transaction signer for a given sender.

#### Parameters

| Name | Type |
| :------ | :------ |
| `val` | [`SendTransactionFrom`](types_transaction.md#sendtransactionfrom) |

#### Returns

`TransactionSigner`

A transaction signer

#### Defined in

[src/transaction.ts:70](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/transaction.ts#L70)

___

### getTestNetDispenserApiClient

▸ **getTestNetDispenserApiClient**(`params?`): [`TestNetDispenserApiClient`](../classes/types_dispenser_client.TestNetDispenserApiClient.md)

Create a new TestNetDispenserApiClient instance.
Refer to [docs](https://github.com/algorandfoundation/algokit/blob/main/docs/testnet_api.md) on guidance to obtain an access token.

#### Parameters

| Name | Type | Default value | Description |
| :------ | :------ | :------ | :------ |
| `params` | ``null`` \| [`TestNetDispenserApiClientParams`](../interfaces/types_dispenser_client.TestNetDispenserApiClientParams.md) | `null` | An object containing parameters for the TestNetDispenserApiClient class. Or null if you want the client to load the access token from the environment variable `ALGOKIT_DISPENSER_ACCESS_TOKEN`. |

#### Returns

[`TestNetDispenserApiClient`](../classes/types_dispenser_client.TestNetDispenserApiClient.md)

An instance of the TestNetDispenserApiClient class.

**`Example`**

```ts
const client = algokit.getTestNetDispenserApiClient(
    {
      authToken: 'your_auth_token',
      requestTimeout: 15,
    }
)
```

#### Defined in

[src/dispenser-client.ts:19](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/dispenser-client.ts#L19)

___

### getTransactionParams

▸ **getTransactionParams**(`params`, `algod`): `Promise`<`SuggestedParamsWithMinFee` \| { `fee`: `number` ; `firstRound`: `number` ; `flatFee?`: `boolean` ; `genesisHash`: `string` ; `genesisID`: `string` ; `lastRound`: `number`  }\>

Returns suggested transaction parameters from algod unless some are already provided.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `params` | `undefined` \| `SuggestedParams` | Optionally provide parameters to use |
| `algod` | `default` | Algod algod |

#### Returns

`Promise`<`SuggestedParamsWithMinFee` \| { `fee`: `number` ; `firstRound`: `number` ; `flatFee?`: `boolean` ; `genesisHash`: `string` ; `genesisID`: `string` ; `lastRound`: `number`  }\>

The suggested transaction parameters

#### Defined in

[src/transaction.ts:472](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/transaction.ts#L472)

___

### getTransactionWithSigner

▸ **getTransactionWithSigner**(`transaction`, `defaultSender?`): `Promise`<`TransactionWithSigner`\>

Given a transaction in a variety of supported formats, returns a TransactionWithSigner object ready to be passed to an
AtomicTransactionComposer's addTransaction method.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `transaction` | `Transaction` \| [`TransactionToSign`](../interfaces/types_transaction.TransactionToSign.md) \| `Promise`<[`SendTransactionResult`](../interfaces/types_transaction.SendTransactionResult.md)\> \| `TransactionWithSigner` | One of: A TransactionWithSigner object (returned as is), a TransactionToSign object (signer is obtained from the signer property), a Transaction object (signer is extracted from the defaultSender parameter), an async SendTransactionResult returned by one of algokit utils' helpers (signer is obtained from the defaultSender parameter) |
| `defaultSender?` | [`SendTransactionFrom`](types_transaction.md#sendtransactionfrom) | The default sender to be used to obtain a signer where the object provided to the transaction parameter does not include a signer. |

#### Returns

`Promise`<`TransactionWithSigner`\>

A TransactionWithSigner object.

#### Defined in

[src/transaction.ts:83](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/transaction.ts#L83)

___

### isLocalNet

▸ **isLocalNet**(`algod`): `Promise`<`boolean`\>

Returns true if the algod client is pointing to a LocalNet Algorand network

#### Parameters

| Name | Type |
| :------ | :------ |
| `algod` | `default` |

#### Returns

`Promise`<`boolean`\>

#### Defined in

[src/localnet.ts:9](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/localnet.ts#L9)

___

### isMainNet

▸ **isMainNet**(`algod`): `Promise`<`boolean`\>

#### Parameters

| Name | Type |
| :------ | :------ |
| `algod` | `default` |

#### Returns

`Promise`<`boolean`\>

#### Defined in

[src/network-client.ts:194](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/network-client.ts#L194)

___

### isSchemaIsBroken

▸ **isSchemaIsBroken**(`before`, `after`): `boolean`

Returns true is there is a breaking change in the application state schema from before to after.
 i.e. if the schema becomes larger, since applications can't ask for more schema after creation.
 Otherwise, there is no error, the app just doesn't store data in the extra schema :(

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `before` | `ApplicationStateSchema` | The existing schema |
| `after` | `ApplicationStateSchema` | The new schema |

#### Returns

`boolean`

Whether or not there is a breaking change

#### Defined in

[src/app-deploy.ts:408](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/app-deploy.ts#L408)

___

### isTestNet

▸ **isTestNet**(`algod`): `Promise`<`boolean`\>

#### Parameters

| Name | Type |
| :------ | :------ |
| `algod` | `default` |

#### Returns

`Promise`<`boolean`\>

#### Defined in

[src/network-client.ts:189](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/network-client.ts#L189)

___

### lookupAccountByAddress

▸ **lookupAccountByAddress**(`accountAddress`, `indexer`): `Promise`<[`AccountLookupResult`](../interfaces/types_indexer.AccountLookupResult.md)\>

Looks up an account by address using Indexer.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `accountAddress` | `string` | - |
| `indexer` | `default` | An indexer client |

#### Returns

`Promise`<[`AccountLookupResult`](../interfaces/types_indexer.AccountLookupResult.md)\>

The result of the look-up

#### Defined in

[src/indexer-lookup.ts:29](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/indexer-lookup.ts#L29)

___

### lookupAccountCreatedApplicationByAddress

▸ **lookupAccountCreatedApplicationByAddress**(`indexer`, `address`, `getAll?`, `paginationLimit?`): `Promise`<[`ApplicationResult`](../interfaces/types_indexer.ApplicationResult.md)[]\>

Looks up applications that were created by the given address.

#### Parameters

| Name | Type | Default value | Description |
| :------ | :------ | :------ | :------ |
| `indexer` | `default` | `undefined` | An indexer instance |
| `address` | `string` | `undefined` | The address of the creator to look up |
| `getAll` | `undefined` \| `boolean` | `undefined` | Whether or not to include deleted applications |
| `paginationLimit?` | `number` | `undefined` | The number of records to return per paginated request, default 1000 |

#### Returns

`Promise`<[`ApplicationResult`](../interfaces/types_indexer.ApplicationResult.md)[]\>

The list of application results

#### Defined in

[src/indexer-lookup.ts:41](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/indexer-lookup.ts#L41)

___

### lookupTransactionById

▸ **lookupTransactionById**(`transactionId`, `indexer`): `Promise`<[`TransactionLookupResult`](../interfaces/types_indexer.TransactionLookupResult.md)\>

Looks up a transaction by ID using Indexer.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `transactionId` | `string` | The ID of the transaction to look up |
| `indexer` | `default` | An indexer client |

#### Returns

`Promise`<[`TransactionLookupResult`](../interfaces/types_indexer.TransactionLookupResult.md)\>

The result of the look-up

#### Defined in

[src/indexer-lookup.ts:19](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/indexer-lookup.ts#L19)

___

### microAlgos

▸ **microAlgos**(`microAlgos`): [`AlgoAmount`](../classes/types_amount.AlgoAmount.md)

Returns an amount of µAlgos using AlgoAmount

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `microAlgos` | `number` | The amount in µAlgos |

#### Returns

[`AlgoAmount`](../classes/types_amount.AlgoAmount.md)

#### Defined in

[src/amount.ts:29](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/amount.ts#L29)

___

### mnemonicAccount

▸ **mnemonicAccount**(`mnemonicSecret`): `Account`

Returns an Algorand account with secret key loaded (i.e. that can sign transactions) by taking the mnemonic secret.

This is a wrapper around algosdk.mnemonicToSecretKey to provide a more friendly/obvious name.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `mnemonicSecret` | `string` | The mnemonic secret representing the private key of an account; **Note: Be careful how the mnemonic is handled**, never commit it into source control and ideally load it from the environment (ideally via a secret storage service) rather than the file system. |

#### Returns

`Account`

#### Defined in

[src/account.ts:47](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/account.ts#L47)

___

### mnemonicAccountFromEnvironment

▸ **mnemonicAccountFromEnvironment**(`account`, `algod`, `kmdClient?`): `Promise`<`Account` \| [`SigningAccount`](../classes/types_account.SigningAccount.md)\>

Returns an Algorand account with private key loaded by convention from environment variables based on the given name identifier.

Note: This function expects to run in a Node.js environment.

## Convention:
* **Non-LocalNet:** will load process.env['{NAME}_MNEMONIC'] as a mnemonic secret; **Note: Be careful how the mnemonic is handled**,
 never commit it into source control and ideally load it via a secret storage service rather than the file system.
  If process.env['{NAME}_SENDER'] is defined then it will use that for the sender address (i.e. to support rekeyed accounts)
* **LocalNet:** will load the account from a KMD wallet called {NAME} and if that wallet doesn't exist it will create it and fund the account for you

This allows you to write code that will work seamlessly in production and local development (LocalNet) without manual config locally (including when you reset the LocalNet).

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `account` | `string` \| { `fundWith?`: [`AlgoAmount`](../classes/types_amount.AlgoAmount.md) ; `name`: `string`  } | The details of the account to get, either the name identifier (string) or an object with: * `name`: string: The name identifier of the account * `fundWith`: The amount to fund the account with when it gets created (when targeting LocalNet), if not specified then 1000 Algos will be funded from the dispenser account |
| `algod` | `default` | An algod client |
| `kmdClient?` | `default` | An optional KMD client to use to create an account (when targeting LocalNet), if not specified then a default KMD client will be loaded from environment variables |

#### Returns

`Promise`<`Account` \| [`SigningAccount`](../classes/types_account.SigningAccount.md)\>

The requested account with private key loaded from the environment variables or when targeting LocalNet from KMD (idempotently creating and funding the account)

**`Example`**

Default

If you have a mnemonic secret loaded into `process.env.MY_ACCOUNT_MNEMONIC` then you can call the following to get that private key loaded into an account object:
```typescript
const account = await mnemonicAccountFromEnvironment('MY_ACCOUNT', algod)
```

If that code runs against LocalNet then a wallet called `MY_ACCOUNT` will automatically be created with an account that is automatically funded with 1000 (default) ALGOs from the default LocalNet dispenser.
If not running against LocalNet then it will use proces.env.MY_ACCOUNT_MNEMONIC as the private key and (if present) process.env.MY_ACCOUNT_SENDER as the sender address.

#### Defined in

[src/account.ts:230](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/account.ts#L230)

___

### multisigAccount

▸ **multisigAccount**(`multisigParams`, `signingAccounts`): [`MultisigAccount`](../classes/types_account.MultisigAccount.md)

Returns an account wrapper that supports partial or full multisig signing.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `multisigParams` | `MultisigMetadata` | The parameters that define the multisig account |
| `signingAccounts` | (`default` \| [`SigningAccount`](../classes/types_account.SigningAccount.md))[] | The signers that are currently present |

#### Returns

[`MultisigAccount`](../classes/types_account.MultisigAccount.md)

A multisig account wrapper

#### Defined in

[src/account.ts:16](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/account.ts#L16)

___

### optIn

▸ **optIn**(`client`, `account`, `assetIds`): `Promise`<`Record`<`number`, `string`\>\>

Opt in to a list of assets on the Algorand blockchain.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `client` | `default` | An instance of the Algodv2 class from the `algosdk` library. |
| `account` | `default` | An instance of the Account class from the `algosdk` library representing the account that wants to opt in to the assets. |
| `assetIds` | `number`[] | An array of asset IDs that the account wants to opt in to. |

#### Returns

`Promise`<`Record`<`number`, `string`\>\>

A record object where the keys are the asset IDs and the values are the corresponding transaction IDs for successful opt-ins.

**`Throws`**

If there is an error during the opt-in process.

#### Defined in

[src/asset.ts:77](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/asset.ts#L77)

___

### optOut

▸ **optOut**(`client`, `account`, `assetIds`): `Promise`<`Record`<`number`, `string`\>\>

Opt out of multiple assets in Algorand blockchain.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `client` | `default` | An instance of the Algodv2 client used to interact with the Algorand blockchain. |
| `account` | `default` | The Algorand account that wants to opt out of the assets. |
| `assetIds` | `number`[] | An array of asset IDs representing the assets to opt out of. |

#### Returns

`Promise`<`Record`<`number`, `string`\>\>

- A record object containing asset IDs as keys and their corresponding transaction IDs as values.

#### Defined in

[src/asset.ts:130](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/asset.ts#L130)

___

### performAtomicTransactionComposerDryrun

▸ **performAtomicTransactionComposerDryrun**(`atc`, `algod`): `Promise`<`DryrunResult`\>

Performs a dry run of the transactions loaded into the given AtomicTransactionComposer`
@param atc The AtomicTransactionComposer` with transaction(s) loaded

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `atc` | `AtomicTransactionComposer` | - |
| `algod` | `default` | An Algod client |

#### Returns

`Promise`<`DryrunResult`\>

The dryrun result

#### Defined in

[src/transaction.ts:278](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/transaction.ts#L278)

___

### performAtomicTransactionComposerSimulate

▸ **performAtomicTransactionComposerSimulate**(`atc`, `algod`): `Promise`<`SimulateResponse`\>

Performs a simulation of the transactions loaded into the given AtomicTransactionComposer.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `atc` | `AtomicTransactionComposer` | The AtomicTransactionComposer with transaction(s) loaded. |
| `algod` | `default` | An Algod client to perform the simulation. |

#### Returns

`Promise`<`SimulateResponse`\>

The simulation result, which includes various details about how the transactions would be processed.

#### Defined in

[src/transaction.ts:293](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/transaction.ts#L293)

___

### performTemplateSubstitution

▸ **performTemplateSubstitution**(`tealCode`, `templateParams?`): `string`

Performs template substitution of a teal file.

Looks for `TMPL_{parameter}` for template replacements.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `tealCode` | `string` | The TEAL logic to compile |
| `templateParams?` | [`TealTemplateParams`](../interfaces/types_app.TealTemplateParams.md) | Any parameters to replace in the .teal file before compiling |

#### Returns

`string`

The TEAL code with replacements

#### Defined in

[src/app-deploy.ts:580](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/app-deploy.ts#L580)

___

### performTemplateSubstitutionAndCompile

▸ **performTemplateSubstitutionAndCompile**(`tealCode`, `algod`, `templateParams?`, `deploymentMetadata?`): `Promise`<[`CompiledTeal`](../interfaces/types_app.CompiledTeal.md)\>

Performs template substitution of a teal file and compiles it, returning the compiled result.

Looks for `TMPL_{parameter}` for template replacements.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `tealCode` | `string` | The TEAL logic to compile |
| `algod` | `default` | An algod client |
| `templateParams?` | [`TealTemplateParams`](../interfaces/types_app.TealTemplateParams.md) | Any parameters to replace in the .teal file before compiling |
| `deploymentMetadata?` | [`AppDeployMetadata`](../interfaces/types_app.AppDeployMetadata.md) | The deployment metadata the app will be deployed with |

#### Returns

`Promise`<[`CompiledTeal`](../interfaces/types_app.CompiledTeal.md)\>

The information about the compiled code

#### Defined in

[src/app-deploy.ts:610](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/app-deploy.ts#L610)

___

### randomAccount

▸ **randomAccount**(): `Account`

Returns a new, random Algorand account with secret key loaded.

This is a wrapper around algosdk.generateAccount to provide a more friendly/obvious name.

#### Returns

`Account`

#### Defined in

[src/account.ts:57](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/account.ts#L57)

___

### rekeyedAccount

▸ **rekeyedAccount**(`signer`, `sender`): [`SigningAccount`](../classes/types_account.SigningAccount.md)

Returns an account wrapper that supports a rekeyed account.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `signer` | `default` | The account, with private key loaded, that is signing |
| `sender` | `string` | The address of the rekeyed account that will act as a sender |

#### Returns

[`SigningAccount`](../classes/types_account.SigningAccount.md)

The SigningAccount wrapper

#### Defined in

[src/account.ts:26](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/account.ts#L26)

___

### replaceDeployTimeControlParams

▸ **replaceDeployTimeControlParams**(`tealCode`, `params`): `string`

Replaces deploy-time deployment control parameters within the given teal code.

* `TMPL_UPDATABLE` for updatability / immutability control
* `TMPL_DELETABLE` for deletability / permanence control

Note: If these values are not undefined, but the corresponding `TMPL_*` value
 isn't in the teal code it will throw an exception.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `tealCode` | `string` | The TEAL code to substitute |
| `params` | `Object` | The deploy-time deployment control parameter value to replace |
| `params.deletable?` | `boolean` | - |
| `params.updatable?` | `boolean` | - |

#### Returns

`string`

The replaced TEAL code

#### Defined in

[src/app-deploy.ts:549](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/app-deploy.ts#L549)

___

### searchTransactions

▸ **searchTransactions**(`indexer`, `searchCriteria`, `paginationLimit?`): `Promise`<[`TransactionSearchResults`](../interfaces/types_indexer.TransactionSearchResults.md)\>

Allows transactions to be searched for the given criteria.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `indexer` | `default` | An indexer client |
| `searchCriteria` | (`s`: `default`) => `default` | The criteria to search for |
| `paginationLimit?` | `number` | The number of records to return per paginated request, default 1000 |

#### Returns

`Promise`<[`TransactionSearchResults`](../interfaces/types_indexer.TransactionSearchResults.md)\>

The search results

#### Defined in

[src/indexer-lookup.ts:74](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/indexer-lookup.ts#L74)

___

### sendAtomicTransactionComposer

▸ **sendAtomicTransactionComposer**(`atcSend`, `algod`): `Promise`<[`SendAtomicTransactionComposerResults`](../interfaces/types_transaction.SendAtomicTransactionComposerResults.md)\>

Signs and sends transactions that have been collected by an `AtomicTransactionComposer`.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `atcSend` | [`AtomicTransactionComposerToSend`](../interfaces/types_transaction.AtomicTransactionComposerToSend.md) | The parameters controlling the send, including: * `atc` The `AtomicTransactionComposer` * `sendParams` The parameters to control the send behaviour |
| `algod` | `default` | An algod client |

#### Returns

`Promise`<[`SendAtomicTransactionComposerResults`](../interfaces/types_transaction.SendAtomicTransactionComposerResults.md)\>

An object with transaction IDs, transactions, group transaction ID (`groupTransactionId`) if more than 1 transaction sent, and (if `skipWaiting` is `false` or unset) confirmation (`confirmation`)

#### Defined in

[src/transaction.ts:190](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/transaction.ts#L190)

___

### sendGroupOfTransactions

▸ **sendGroupOfTransactions**(`groupSend`, `algod`): `Promise`<`Omit`<[`SendAtomicTransactionComposerResults`](../interfaces/types_transaction.SendAtomicTransactionComposerResults.md), ``"returns"``\>\>

Signs and sends a group of [up to 16](https://developer.algorand.org/docs/get-details/atomic_transfers/#create-transactions) transactions to the chain

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `groupSend` | [`TransactionGroupToSend`](../interfaces/types_transaction.TransactionGroupToSend.md) | The group details to send, with: * `transactions`: The array of transactions to send along with their signing account * `sendParams`: The parameters to dictate how the group is sent |
| `algod` | `default` | An algod client |

#### Returns

`Promise`<`Omit`<[`SendAtomicTransactionComposerResults`](../interfaces/types_transaction.SendAtomicTransactionComposerResults.md), ``"returns"``\>\>

An object with transaction IDs, transactions, group transaction ID (`groupTransactionId`) if more than 1 transaction sent, and (if `skipWaiting` is `false` or unset) confirmation (`confirmation`)

#### Defined in

[src/transaction.ts:324](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/transaction.ts#L324)

___

### sendTransaction

▸ **sendTransaction**(`send`, `algod`): `Promise`<[`SendTransactionResult`](../interfaces/types_transaction.SendTransactionResult.md)\>

Prepares a transaction for sending and then (if instructed) signs and sends the given transaction to the chain.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `send` | `Object` | The details for the transaction to prepare/send, including: * `transaction`: The unsigned transaction * `from`: The account to sign the transaction with: either an account with private key loaded or a logic signature account * `config`: The sending configuration for this transaction |
| `send.from` | [`SendTransactionFrom`](types_transaction.md#sendtransactionfrom) | - |
| `send.sendParams?` | [`SendTransactionParams`](../interfaces/types_transaction.SendTransactionParams.md) | - |
| `send.transaction` | `Transaction` | - |
| `algod` | `default` | An algod client |

#### Returns

`Promise`<[`SendTransactionResult`](../interfaces/types_transaction.SendTransactionResult.md)\>

An object with transaction (`transaction`) and (if `skipWaiting` is `false` or `undefined`) confirmation (`confirmation`)

#### Defined in

[src/transaction.ts:146](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/transaction.ts#L146)

___

### signTransaction

▸ **signTransaction**(`transaction`, `signer`): `Promise`<`Uint8Array`\>

Signs a single transaction by the given signer.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `transaction` | `Transaction` | The transaction to sign |
| `signer` | [`SendTransactionFrom`](types_transaction.md#sendtransactionfrom) | The signer to sign |

#### Returns

`Promise`<`Uint8Array`\>

The signed transaction as a `Uint8Array`

#### Defined in

[src/transaction.ts:126](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/transaction.ts#L126)

___

### stripTealComments

▸ **stripTealComments**(`tealCode`): `string`

Remove comments from TEAL Code

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `tealCode` | `string` | The TEAL logic to compile |

#### Returns

`string`

The TEAL without comments

#### Defined in

[src/app-deploy.ts:633](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/app-deploy.ts#L633)

___

### transactionFees

▸ **transactionFees**(`numberOfTransactions`): [`AlgoAmount`](../classes/types_amount.AlgoAmount.md)

Returns an amount of µAlgos to cover standard fees for the given number of transactions using AlgoAmount

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `numberOfTransactions` | `number` | The of standard transaction fees to return the amount of ALGOs |

#### Returns

[`AlgoAmount`](../classes/types_amount.AlgoAmount.md)

#### Defined in

[src/amount.ts:36](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/amount.ts#L36)

___

### transactionSignerAccount

▸ **transactionSignerAccount**(`signer`, `sender`): [`TransactionSignerAccount`](../interfaces/types_account.TransactionSignerAccount.md)

Returns an account wrapper that supports a transaction signer with associated sender address.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `signer` | `TransactionSigner` | The transaction signer |
| `sender` | `string` | The address of sender account |

#### Returns

[`TransactionSignerAccount`](../interfaces/types_account.TransactionSignerAccount.md)

The SigningAccount wrapper

#### Defined in

[src/account.ts:36](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/account.ts#L36)

___

### transferAlgos

▸ **transferAlgos**(`transfer`, `algod`): `Promise`<[`SendTransactionResult`](../interfaces/types_transaction.SendTransactionResult.md)\>

Transfer ALGOs between two accounts.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `transfer` | [`AlgoTransferParams`](../interfaces/types_transfer.AlgoTransferParams.md) | The transfer definition |
| `algod` | `default` | An algod client |

#### Returns

`Promise`<[`SendTransactionResult`](../interfaces/types_transaction.SendTransactionResult.md)\>

The transaction object and optionally the confirmation if it was sent to the chain (`skipSending` is `false` or unset)

**`Example`**

Usage example
```typescript
await algokit.transferAlgos({ from, to, amount: algokit.algos(1) }, algod)
```

#### Defined in

[src/transfer.ts:82](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/transfer.ts#L82)

___

### transferAsset

▸ **transferAsset**(`transfer`, `algod`): `Promise`<[`SendTransactionResult`](../interfaces/types_transaction.SendTransactionResult.md)\>

Transfer asset between two accounts.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `transfer` | [`TransferAssetParams`](../interfaces/types_transfer.TransferAssetParams.md) | The transfer definition |
| `algod` | `default` | An algod client |

#### Returns

`Promise`<[`SendTransactionResult`](../interfaces/types_transaction.SendTransactionResult.md)\>

The transaction object and optionally the confirmation if it was sent to the chain (`skipSending` is `false` or unset)

**`Example`**

Usage example
```typescript
await algokit.transferAsset({ from, to, assetId, amount }, algod)
```

#### Defined in

[src/transfer.ts:165](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/transfer.ts#L165)

___

### updateApp

▸ **updateApp**(`update`, `algod`): `Promise`<`Partial`<[`AppCompilationResult`](../interfaces/types_app.AppCompilationResult.md)\> & [`AppCallTransactionResult`](../interfaces/types_app.AppCallTransactionResult.md)\>

Updates a smart contract app.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `update` | [`UpdateAppParams`](../interfaces/types_app.UpdateAppParams.md) | The parameters to update the app with |
| `algod` | `default` | An algod client |

#### Returns

`Promise`<`Partial`<[`AppCompilationResult`](../interfaces/types_app.AppCompilationResult.md)\> & [`AppCallTransactionResult`](../interfaces/types_app.AppCallTransactionResult.md)\>

The transaction send result and the compilation result

#### Defined in

[src/app.ts:188](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/app.ts#L188)

___

### waitForConfirmation

▸ **waitForConfirmation**(`transactionId`, `maxRoundsToWait`, `algod`): `Promise`<`PendingTransactionResponse`\>

Wait until the transaction is confirmed or rejected, or until `timeout`
number of rounds have passed.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `transactionId` | `string` | The transaction ID to wait for |
| `maxRoundsToWait` | `number` \| `bigint` | Maximum number of rounds to wait |
| `algod` | `default` | An algod client |

#### Returns

`Promise`<`PendingTransactionResponse`\>

Pending transaction information

**`Throws`**

Throws an error if the transaction is not confirmed or rejected in the next `timeout` rounds

#### Defined in

[src/transaction.ts:369](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/transaction.ts#L369)
