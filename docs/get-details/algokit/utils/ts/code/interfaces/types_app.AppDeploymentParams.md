[@algorandfoundation/algokit-utils](../index.md) / [types/app](../modules/types_app.md) / AppDeploymentParams

# Interface: AppDeploymentParams

[types/app](../modules/types_app.md).AppDeploymentParams

The parameters to deploy an app

## Hierarchy

- `Omit`<[`CreateAppParams`](types_app.CreateAppParams.md), ``"onCompleteAction"`` \| ``"args"`` \| ``"note"`` \| ``"skipSending"`` \| ``"skipWaiting"`` \| ``"atc"``\>

  ↳ **`AppDeploymentParams`**

## Table of contents

### Properties

- [approvalProgram](types_app.AppDeploymentParams.md#approvalprogram)
- [clearStateProgram](types_app.AppDeploymentParams.md#clearstateprogram)
- [createArgs](types_app.AppDeploymentParams.md#createargs)
- [createOnCompleteAction](types_app.AppDeploymentParams.md#createoncompleteaction)
- [deleteArgs](types_app.AppDeploymentParams.md#deleteargs)
- [deployTimeParams](types_app.AppDeploymentParams.md#deploytimeparams)
- [existingDeployments](types_app.AppDeploymentParams.md#existingdeployments)
- [fee](types_app.AppDeploymentParams.md#fee)
- [from](types_app.AppDeploymentParams.md#from)
- [maxFee](types_app.AppDeploymentParams.md#maxfee)
- [maxRoundsToWaitForConfirmation](types_app.AppDeploymentParams.md#maxroundstowaitforconfirmation)
- [metadata](types_app.AppDeploymentParams.md#metadata)
- [onSchemaBreak](types_app.AppDeploymentParams.md#onschemabreak)
- [onUpdate](types_app.AppDeploymentParams.md#onupdate)
- [schema](types_app.AppDeploymentParams.md#schema)
- [suppressLog](types_app.AppDeploymentParams.md#suppresslog)
- [transactionParams](types_app.AppDeploymentParams.md#transactionparams)
- [updateArgs](types_app.AppDeploymentParams.md#updateargs)

## Properties

### approvalProgram

• **approvalProgram**: `string` \| `Uint8Array`

The approval program as raw teal (string) or compiled teal, base 64 encoded as a byte array (Uint8Array)

#### Inherited from

Omit.approvalProgram

#### Defined in

[src/types/app.ts:121](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L121)

___

### clearStateProgram

• **clearStateProgram**: `string` \| `Uint8Array`

The clear state program as raw teal (string) or compiled teal, base 64 encoded as a byte array (Uint8Array)

#### Inherited from

Omit.clearStateProgram

#### Defined in

[src/types/app.ts:123](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L123)

___

### createArgs

• `Optional` **createArgs**: [`AppCallArgs`](../modules/types_app.md#appcallargs)

Any args to pass to any create transaction that is issued as part of deployment

#### Defined in

[src/types/app.ts:300](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L300)

___

### createOnCompleteAction

• `Optional` **createOnCompleteAction**: ``"no_op"`` \| ``"opt_in"`` \| ``"close_out"`` \| ``"update_application"`` \| ``"delete_application"`` \| `NoOpOC` \| `OptInOC` \| `CloseOutOC` \| `UpdateApplicationOC` \| `DeleteApplicationOC`

Override the on-completion action for the create call; defaults to NoOp

#### Defined in

[src/types/app.ts:302](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L302)

___

### deleteArgs

• `Optional` **deleteArgs**: [`AppCallArgs`](../modules/types_app.md#appcallargs)

Any args to pass to any delete transaction that is issued as part of deployment

#### Defined in

[src/types/app.ts:306](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L306)

___

### deployTimeParams

• `Optional` **deployTimeParams**: [`TealTemplateParams`](types_app.TealTemplateParams.md)

Any deploy-time parameters to replace in the TEAL code

#### Defined in

[src/types/app.ts:292](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L292)

___

### existingDeployments

• `Optional` **existingDeployments**: [`AppLookup`](types_app.AppLookup.md)

Optional cached value of the existing apps for the given creator

#### Defined in

[src/types/app.ts:298](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L298)

___

### fee

• `Optional` **fee**: [`AlgoAmount`](../classes/types_amount.AlgoAmount.md)

The flat fee you want to pay, useful for covering extra fees in a transaction group or app call

#### Inherited from

Omit.fee

#### Defined in

[src/types/transaction.ts:34](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L34)

___

### from

• **from**: [`SendTransactionFrom`](../modules/types_transaction.md#sendtransactionfrom)

The account (with private key loaded) that will send the transaction

#### Inherited from

Omit.from

#### Defined in

[src/types/app.ts:119](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L119)

___

### maxFee

• `Optional` **maxFee**: [`AlgoAmount`](../classes/types_amount.AlgoAmount.md)

The maximum fee that you are happy to pay (default: unbounded) - if this is set it's possible the transaction could get rejected during network congestion

#### Inherited from

Omit.maxFee

#### Defined in

[src/types/transaction.ts:36](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L36)

___

### maxRoundsToWaitForConfirmation

• `Optional` **maxRoundsToWaitForConfirmation**: `number`

The maximum number of rounds to wait for confirmation, only applies if `skipWaiting` is `undefined` or `false`, default: wait up to 5 rounds

#### Inherited from

Omit.maxRoundsToWaitForConfirmation

#### Defined in

[src/types/transaction.ts:38](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L38)

___

### metadata

• **metadata**: [`AppDeployMetadata`](types_app.AppDeployMetadata.md)

The deployment metadata

#### Defined in

[src/types/app.ts:290](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L290)

___

### onSchemaBreak

• `Optional` **onSchemaBreak**: [`OnSchemaBreak`](../enums/types_app.OnSchemaBreak.md) \| ``"replace"`` \| ``"fail"`` \| ``"append"``

What action to perform if a schema break is detected

#### Defined in

[src/types/app.ts:294](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L294)

___

### onUpdate

• `Optional` **onUpdate**: ``"replace"`` \| ``"fail"`` \| ``"append"`` \| [`OnUpdate`](../enums/types_app.OnUpdate.md) \| ``"update"``

What action to perform if a TEAL update is detected

#### Defined in

[src/types/app.ts:296](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L296)

___

### schema

• **schema**: [`AppStorageSchema`](types_app.AppStorageSchema.md)

The storage schema to request for the created app

#### Inherited from

Omit.schema

#### Defined in

[src/types/app.ts:135](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L135)

___

### suppressLog

• `Optional` **suppressLog**: `boolean`

Whether to suppress log messages from transaction send, default: do not suppress

#### Inherited from

Omit.suppressLog

#### Defined in

[src/types/transaction.ts:32](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L32)

___

### transactionParams

• `Optional` **transactionParams**: `SuggestedParams`

Optional transaction parameters

#### Inherited from

Omit.transactionParams

#### Defined in

[src/types/app.ts:125](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L125)

___

### updateArgs

• `Optional` **updateArgs**: [`AppCallArgs`](../modules/types_app.md#appcallargs)

Any args to pass to any update transaction that is issued as part of deployment

#### Defined in

[src/types/app.ts:304](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L304)
