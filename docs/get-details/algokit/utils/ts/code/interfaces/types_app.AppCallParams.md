[@algorandfoundation/algokit-utils](../index.md) / [types/app](../modules/types_app.md) / AppCallParams

# Interface: AppCallParams

[types/app](../modules/types_app.md).AppCallParams

Parameters representing a call to an app.

## Hierarchy

- [`SendTransactionParams`](types_transaction.SendTransactionParams.md)

  ↳ **`AppCallParams`**

## Table of contents

### Properties

- [appId](types_app.AppCallParams.md#appid)
- [args](types_app.AppCallParams.md#args)
- [atc](types_app.AppCallParams.md#atc)
- [callType](types_app.AppCallParams.md#calltype)
- [fee](types_app.AppCallParams.md#fee)
- [from](types_app.AppCallParams.md#from)
- [maxFee](types_app.AppCallParams.md#maxfee)
- [maxRoundsToWaitForConfirmation](types_app.AppCallParams.md#maxroundstowaitforconfirmation)
- [note](types_app.AppCallParams.md#note)
- [skipSending](types_app.AppCallParams.md#skipsending)
- [skipWaiting](types_app.AppCallParams.md#skipwaiting)
- [suppressLog](types_app.AppCallParams.md#suppresslog)
- [transactionParams](types_app.AppCallParams.md#transactionparams)

## Properties

### appId

• **appId**: `number` \| `bigint`

The id of the app to call

#### Defined in

[src/types/app.ts:162](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L162)

___

### args

• `Optional` **args**: [`AppCallArgs`](../modules/types_app.md#appcallargs)

The arguments passed in to the app call

#### Defined in

[src/types/app.ts:172](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L172)

___

### atc

• `Optional` **atc**: `AtomicTransactionComposer`

An optional `AtomicTransactionComposer` to add the transaction to, if specified then `skipSending: undefined` has the same effect as `skipSending: true`

#### Inherited from

[SendTransactionParams](types_transaction.SendTransactionParams.md).[atc](types_transaction.SendTransactionParams.md#atc)

#### Defined in

[src/types/transaction.ts:30](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L30)

___

### callType

• **callType**: ``"no_op"`` \| ``"opt_in"`` \| ``"close_out"`` \| ``"clear_state"`` \| ``"delete_application"`` \| `NoOpOC` \| `OptInOC` \| `CloseOutOC` \| `ClearStateOC` \| `DeleteApplicationOC`

The type of call, everything except create (see `createApp`) and update (see `updateApp`)

#### Defined in

[src/types/app.ts:164](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L164)

___

### fee

• `Optional` **fee**: [`AlgoAmount`](../classes/types_amount.AlgoAmount.md)

The flat fee you want to pay, useful for covering extra fees in a transaction group or app call

#### Inherited from

[SendTransactionParams](types_transaction.SendTransactionParams.md).[fee](types_transaction.SendTransactionParams.md#fee)

#### Defined in

[src/types/transaction.ts:34](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L34)

___

### from

• **from**: [`SendTransactionFrom`](../modules/types_transaction.md#sendtransactionfrom)

The account to make the call from

#### Defined in

[src/types/app.ts:166](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L166)

___

### maxFee

• `Optional` **maxFee**: [`AlgoAmount`](../classes/types_amount.AlgoAmount.md)

The maximum fee that you are happy to pay (default: unbounded) - if this is set it's possible the transaction could get rejected during network congestion

#### Inherited from

[SendTransactionParams](types_transaction.SendTransactionParams.md).[maxFee](types_transaction.SendTransactionParams.md#maxfee)

#### Defined in

[src/types/transaction.ts:36](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L36)

___

### maxRoundsToWaitForConfirmation

• `Optional` **maxRoundsToWaitForConfirmation**: `number`

The maximum number of rounds to wait for confirmation, only applies if `skipWaiting` is `undefined` or `false`, default: wait up to 5 rounds

#### Inherited from

[SendTransactionParams](types_transaction.SendTransactionParams.md).[maxRoundsToWaitForConfirmation](types_transaction.SendTransactionParams.md#maxroundstowaitforconfirmation)

#### Defined in

[src/types/transaction.ts:38](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L38)

___

### note

• `Optional` **note**: [`TransactionNote`](../modules/types_transaction.md#transactionnote)

The (optional) transaction note

#### Defined in

[src/types/app.ts:170](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L170)

___

### skipSending

• `Optional` **skipSending**: `boolean`

Whether to skip signing and sending the transaction to the chain (default: transaction signed and sent to chain, unless `atc` specified)
and instead just return the raw transaction, e.g. so you can add it to a group of transactions

#### Inherited from

[SendTransactionParams](types_transaction.SendTransactionParams.md).[skipSending](types_transaction.SendTransactionParams.md#skipsending)

#### Defined in

[src/types/transaction.ts:26](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L26)

___

### skipWaiting

• `Optional` **skipWaiting**: `boolean`

Whether to skip waiting for the submitted transaction (only relevant if `skipSending` is `false` or unset)

#### Inherited from

[SendTransactionParams](types_transaction.SendTransactionParams.md).[skipWaiting](types_transaction.SendTransactionParams.md#skipwaiting)

#### Defined in

[src/types/transaction.ts:28](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L28)

___

### suppressLog

• `Optional` **suppressLog**: `boolean`

Whether to suppress log messages from transaction send, default: do not suppress

#### Inherited from

[SendTransactionParams](types_transaction.SendTransactionParams.md).[suppressLog](types_transaction.SendTransactionParams.md#suppresslog)

#### Defined in

[src/types/transaction.ts:32](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L32)

___

### transactionParams

• `Optional` **transactionParams**: `SuggestedParams`

Optional transaction parameters

#### Defined in

[src/types/app.ts:168](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L168)
