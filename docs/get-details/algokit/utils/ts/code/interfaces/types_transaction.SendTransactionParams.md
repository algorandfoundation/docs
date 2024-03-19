[@algorandfoundation/algokit-utils](../index.md) / [types/transaction](../modules/types_transaction.md) / SendTransactionParams

# Interface: SendTransactionParams

[types/transaction](../modules/types_transaction.md).SendTransactionParams

The sending configuration for a transaction

## Hierarchy

- **`SendTransactionParams`**

  ↳ [`AppCallParams`](types_app.AppCallParams.md)

  ↳ [`AssetOptInParams`](types_asset.AssetOptInParams.md)

  ↳ [`AlgoTransferParams`](types_transfer.AlgoTransferParams.md)

  ↳ [`AlgoRekeyParams`](types_transfer.AlgoRekeyParams.md)

  ↳ [`EnsureFundedParams`](types_transfer.EnsureFundedParams.md)

  ↳ [`TransferAssetParams`](types_transfer.TransferAssetParams.md)

## Table of contents

### Properties

- [atc](types_transaction.SendTransactionParams.md#atc)
- [fee](types_transaction.SendTransactionParams.md#fee)
- [maxFee](types_transaction.SendTransactionParams.md#maxfee)
- [maxRoundsToWaitForConfirmation](types_transaction.SendTransactionParams.md#maxroundstowaitforconfirmation)
- [populateAppCallResources](types_transaction.SendTransactionParams.md#populateappcallresources)
- [skipSending](types_transaction.SendTransactionParams.md#skipsending)
- [skipWaiting](types_transaction.SendTransactionParams.md#skipwaiting)
- [suppressLog](types_transaction.SendTransactionParams.md#suppresslog)

## Properties

### atc

• `Optional` **atc**: `AtomicTransactionComposer`

An optional `AtomicTransactionComposer` to add the transaction to, if specified then `skipSending: undefined` has the same effect as `skipSending: true`

#### Defined in

[src/types/transaction.ts:35](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L35)

___

### fee

• `Optional` **fee**: [`AlgoAmount`](../classes/types_amount.AlgoAmount.md)

The flat fee you want to pay, useful for covering extra fees in a transaction group or app call

#### Defined in

[src/types/transaction.ts:39](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L39)

___

### maxFee

• `Optional` **maxFee**: [`AlgoAmount`](../classes/types_amount.AlgoAmount.md)

The maximum fee that you are happy to pay (default: unbounded) - if this is set it's possible the transaction could get rejected during network congestion

#### Defined in

[src/types/transaction.ts:41](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L41)

___

### maxRoundsToWaitForConfirmation

• `Optional` **maxRoundsToWaitForConfirmation**: `number`

The maximum number of rounds to wait for confirmation, only applies if `skipWaiting` is `undefined` or `false`, default: wait up to 5 rounds

#### Defined in

[src/types/transaction.ts:43](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L43)

___

### populateAppCallResources

• `Optional` **populateAppCallResources**: `boolean`

**WARNING**: Not recommended for production use due to https://github.com/algorand/go-algorand/issues/5914. Whether to use simulate to automatically populate app call resources in the txn objects. Defaults to true when there are app calls in the group.

#### Defined in

[src/types/transaction.ts:45](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L45)

___

### skipSending

• `Optional` **skipSending**: `boolean`

Whether to skip signing and sending the transaction to the chain (default: transaction signed and sent to chain, unless `atc` specified)
and instead just return the raw transaction, e.g. so you can add it to a group of transactions

#### Defined in

[src/types/transaction.ts:31](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L31)

___

### skipWaiting

• `Optional` **skipWaiting**: `boolean`

Whether to skip waiting for the submitted transaction (only relevant if `skipSending` is `false` or unset)

#### Defined in

[src/types/transaction.ts:33](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L33)

___

### suppressLog

• `Optional` **suppressLog**: `boolean`

Whether to suppress log messages from transaction send, default: do not suppress

#### Defined in

[src/types/transaction.ts:37](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L37)
