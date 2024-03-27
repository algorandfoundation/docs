[@algorandfoundation/algokit-utils](../index.md) / [types/asset](../modules/types_asset.md) / AssetOptInParams

# Interface: AssetOptInParams

[types/asset](../modules/types_asset.md).AssetOptInParams

Parameters for `assetOptIn` call.

## Hierarchy

- [`SendTransactionParams`](types_transaction.SendTransactionParams.md)

  ↳ **`AssetOptInParams`**

  ↳↳ [`AssetOptOutParams`](types_asset.AssetOptOutParams.md)

## Table of contents

### Properties

- [account](types_asset.AssetOptInParams.md#account)
- [assetId](types_asset.AssetOptInParams.md#assetid)
- [atc](types_asset.AssetOptInParams.md#atc)
- [fee](types_asset.AssetOptInParams.md#fee)
- [lease](types_asset.AssetOptInParams.md#lease)
- [maxFee](types_asset.AssetOptInParams.md#maxfee)
- [maxRoundsToWaitForConfirmation](types_asset.AssetOptInParams.md#maxroundstowaitforconfirmation)
- [note](types_asset.AssetOptInParams.md#note)
- [populateAppCallResources](types_asset.AssetOptInParams.md#populateappcallresources)
- [skipSending](types_asset.AssetOptInParams.md#skipsending)
- [skipWaiting](types_asset.AssetOptInParams.md#skipwaiting)
- [suppressLog](types_asset.AssetOptInParams.md#suppresslog)
- [transactionParams](types_asset.AssetOptInParams.md#transactionparams)

## Properties

### account

• **account**: [`SendTransactionFrom`](../modules/types_transaction.md#sendtransactionfrom)

The account to opt in/out for

#### Defined in

[src/types/asset.ts:9](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/asset.ts#L9)

___

### assetId

• **assetId**: `number`

The ID of the assets to opt in for / out of

#### Defined in

[src/types/asset.ts:11](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/asset.ts#L11)

___

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

### lease

• `Optional` **lease**: `string` \| `Uint8Array`

An (optional) [transaction lease](https://developer.algorand.org/articles/leased-transactions-securing-advanced-smart-contract-design/) to apply

#### Defined in

[src/types/asset.ts:17](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/asset.ts#L17)

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

[src/types/asset.ts:15](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/asset.ts#L15)

___

### populateAppCallResources

• `Optional` **populateAppCallResources**: `boolean`

**WARNING**: Not recommended for production use due to https://github.com/algorand/go-algorand/issues/5914. Whether to use simulate to automatically populate app call resources in the txn objects. Defaults to true when there are app calls in the group.

#### Inherited from

[SendTransactionParams](types_transaction.SendTransactionParams.md).[populateAppCallResources](types_transaction.SendTransactionParams.md#populateappcallresources)

#### Defined in

[src/types/transaction.ts:45](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L45)

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

[src/types/asset.ts:13](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/asset.ts#L13)
