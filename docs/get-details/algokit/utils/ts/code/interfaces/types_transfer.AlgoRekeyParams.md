[@algorandfoundation/algokit-utils](../index.md) / [types/transfer](../modules/types_transfer.md) / AlgoRekeyParams

# Interface: AlgoRekeyParams

[types/transfer](../modules/types_transfer.md).AlgoRekeyParams

Parameters for `rekeyAccount` call.

## Hierarchy

- [`SendTransactionParams`](types_transaction.SendTransactionParams.md)

  ↳ **`AlgoRekeyParams`**

## Table of contents

### Properties

- [atc](types_transfer.AlgoRekeyParams.md#atc)
- [fee](types_transfer.AlgoRekeyParams.md#fee)
- [from](types_transfer.AlgoRekeyParams.md#from)
- [lease](types_transfer.AlgoRekeyParams.md#lease)
- [maxFee](types_transfer.AlgoRekeyParams.md#maxfee)
- [maxRoundsToWaitForConfirmation](types_transfer.AlgoRekeyParams.md#maxroundstowaitforconfirmation)
- [note](types_transfer.AlgoRekeyParams.md#note)
- [rekeyTo](types_transfer.AlgoRekeyParams.md#rekeyto)
- [skipSending](types_transfer.AlgoRekeyParams.md#skipsending)
- [skipWaiting](types_transfer.AlgoRekeyParams.md#skipwaiting)
- [suppressLog](types_transfer.AlgoRekeyParams.md#suppresslog)
- [transactionParams](types_transfer.AlgoRekeyParams.md#transactionparams)

## Properties

### atc

• `Optional` **atc**: `AtomicTransactionComposer`

An optional `AtomicTransactionComposer` to add the transaction to, if specified then `skipSending: undefined` has the same effect as `skipSending: true`

#### Inherited from

[SendTransactionParams](types_transaction.SendTransactionParams.md).[atc](types_transaction.SendTransactionParams.md#atc)

#### Defined in

[src/types/transaction.ts:35](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L35)

___

### fee

• `Optional` **fee**: [`AlgoAmount`](../classes/types_amount.AlgoAmount.md)

The flat fee you want to pay, useful for covering extra fees in a transaction group or app call

#### Inherited from

[SendTransactionParams](types_transaction.SendTransactionParams.md).[fee](types_transaction.SendTransactionParams.md#fee)

#### Defined in

[src/types/transaction.ts:39](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L39)

___

### from

• **from**: [`SendTransactionFrom`](../modules/types_transaction.md#sendtransactionfrom)

The account that will be rekeyed

#### Defined in

[src/types/transfer.ts:26](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transfer.ts#L26)

___

### lease

• `Optional` **lease**: `string` \| `Uint8Array`

An (optional) [transaction lease](https://developer.algorand.org/articles/leased-transactions-securing-advanced-smart-contract-design/) to apply

#### Defined in

[src/types/transfer.ts:34](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transfer.ts#L34)

___

### maxFee

• `Optional` **maxFee**: [`AlgoAmount`](../classes/types_amount.AlgoAmount.md)

The maximum fee that you are happy to pay (default: unbounded) - if this is set it's possible the transaction could get rejected during network congestion

#### Inherited from

[SendTransactionParams](types_transaction.SendTransactionParams.md).[maxFee](types_transaction.SendTransactionParams.md#maxfee)

#### Defined in

[src/types/transaction.ts:41](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L41)

___

### maxRoundsToWaitForConfirmation

• `Optional` **maxRoundsToWaitForConfirmation**: `number`

The maximum number of rounds to wait for confirmation, only applies if `skipWaiting` is `undefined` or `false`, default: wait up to 5 rounds

#### Inherited from

[SendTransactionParams](types_transaction.SendTransactionParams.md).[maxRoundsToWaitForConfirmation](types_transaction.SendTransactionParams.md#maxroundstowaitforconfirmation)

#### Defined in

[src/types/transaction.ts:43](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L43)

___

### note

• `Optional` **note**: [`TransactionNote`](../modules/types_transaction.md#transactionnote)

The (optional) transaction note

#### Defined in

[src/types/transfer.ts:32](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transfer.ts#L32)

___

### rekeyTo

• **rekeyTo**: `string` \| [`SendTransactionFrom`](../modules/types_transaction.md#sendtransactionfrom)

The account / account address that will have the private key that is authorised to transact on behalf of the from account from now on

#### Defined in

[src/types/transfer.ts:28](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transfer.ts#L28)

___

### skipSending

• `Optional` **skipSending**: `boolean`

Whether to skip signing and sending the transaction to the chain (default: transaction signed and sent to chain, unless `atc` specified)
and instead just return the raw transaction, e.g. so you can add it to a group of transactions

#### Inherited from

[SendTransactionParams](types_transaction.SendTransactionParams.md).[skipSending](types_transaction.SendTransactionParams.md#skipsending)

#### Defined in

[src/types/transaction.ts:31](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L31)

___

### skipWaiting

• `Optional` **skipWaiting**: `boolean`

Whether to skip waiting for the submitted transaction (only relevant if `skipSending` is `false` or unset)

#### Inherited from

[SendTransactionParams](types_transaction.SendTransactionParams.md).[skipWaiting](types_transaction.SendTransactionParams.md#skipwaiting)

#### Defined in

[src/types/transaction.ts:33](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L33)

___

### suppressLog

• `Optional` **suppressLog**: `boolean`

Whether to suppress log messages from transaction send, default: do not suppress

#### Inherited from

[SendTransactionParams](types_transaction.SendTransactionParams.md).[suppressLog](types_transaction.SendTransactionParams.md#suppresslog)

#### Defined in

[src/types/transaction.ts:37](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L37)

___

### transactionParams

• `Optional` **transactionParams**: `SuggestedParams`

Optional transaction parameters

#### Defined in

[src/types/transfer.ts:30](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transfer.ts#L30)
