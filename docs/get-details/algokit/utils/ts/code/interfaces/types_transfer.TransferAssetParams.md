[@algorandfoundation/algokit-utils](../index.md) / [types/transfer](../modules/types_transfer.md) / TransferAssetParams

# Interface: TransferAssetParams

[types/transfer](../modules/types_transfer.md).TransferAssetParams

Parameters for `transferAsset` call.

## Hierarchy

- [`SendTransactionParams`](types_transaction.SendTransactionParams.md)

  ↳ **`TransferAssetParams`**

## Table of contents

### Properties

- [amount](types_transfer.TransferAssetParams.md#amount)
- [assetId](types_transfer.TransferAssetParams.md#assetid)
- [atc](types_transfer.TransferAssetParams.md#atc)
- [clawbackFrom](types_transfer.TransferAssetParams.md#clawbackfrom)
- [fee](types_transfer.TransferAssetParams.md#fee)
- [from](types_transfer.TransferAssetParams.md#from)
- [lease](types_transfer.TransferAssetParams.md#lease)
- [maxFee](types_transfer.TransferAssetParams.md#maxfee)
- [maxRoundsToWaitForConfirmation](types_transfer.TransferAssetParams.md#maxroundstowaitforconfirmation)
- [note](types_transfer.TransferAssetParams.md#note)
- [populateAppCallResources](types_transfer.TransferAssetParams.md#populateappcallresources)
- [skipSending](types_transfer.TransferAssetParams.md#skipsending)
- [skipWaiting](types_transfer.TransferAssetParams.md#skipwaiting)
- [suppressLog](types_transfer.TransferAssetParams.md#suppresslog)
- [to](types_transfer.TransferAssetParams.md#to)
- [transactionParams](types_transfer.TransferAssetParams.md#transactionparams)

## Properties

### amount

• **amount**: `number` \| `bigint`

The amount to send as the smallest divisible unit value

#### Defined in

[src/types/transfer.ts:64](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transfer.ts#L64)

___

### assetId

• **assetId**: `number`

The asset id that will be transfered

#### Defined in

[src/types/transfer.ts:62](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transfer.ts#L62)

___

### atc

• `Optional` **atc**: `AtomicTransactionComposer`

An optional `AtomicTransactionComposer` to add the transaction to, if specified then `skipSending: undefined` has the same effect as `skipSending: true`

#### Inherited from

[SendTransactionParams](types_transaction.SendTransactionParams.md).[atc](types_transaction.SendTransactionParams.md#atc)

#### Defined in

[src/types/transaction.ts:35](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L35)

___

### clawbackFrom

• `Optional` **clawbackFrom**: `string` \| [`SendTransactionFrom`](../modules/types_transaction.md#sendtransactionfrom)

An address of a target account from which to perform a clawback operation. Please note, in such cases senderAccount must be equal to clawback field on ASA metadata.

#### Defined in

[src/types/transfer.ts:68](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transfer.ts#L68)

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

The account that will send the asset

#### Defined in

[src/types/transfer.ts:58](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transfer.ts#L58)

___

### lease

• `Optional` **lease**: `string` \| `Uint8Array`

An (optional) [transaction lease](https://developer.algorand.org/articles/leased-transactions-securing-advanced-smart-contract-design/) to apply

#### Defined in

[src/types/transfer.ts:72](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transfer.ts#L72)

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

[src/types/transfer.ts:70](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transfer.ts#L70)

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

### to

• **to**: `string` \| [`SendTransactionFrom`](../modules/types_transaction.md#sendtransactionfrom)

The account / account address that will receive the asset

#### Defined in

[src/types/transfer.ts:60](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transfer.ts#L60)

___

### transactionParams

• `Optional` **transactionParams**: `SuggestedParams`

Optional transaction parameters

#### Defined in

[src/types/transfer.ts:66](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transfer.ts#L66)
