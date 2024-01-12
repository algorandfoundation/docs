[@algorandfoundation/algokit-utils](../index.md) / [types/asset](../modules/types_asset.md) / AssetBulkOptInOutParams

# Interface: AssetBulkOptInOutParams

[types/asset](../modules/types_asset.md).AssetBulkOptInOutParams

Parameters for `assetBulkOptIn` / `assetBulkOptOut` call.

## Table of contents

### Properties

- [account](types_asset.AssetBulkOptInOutParams.md#account)
- [assetIds](types_asset.AssetBulkOptInOutParams.md#assetids)
- [maxFee](types_asset.AssetBulkOptInOutParams.md#maxfee)
- [note](types_asset.AssetBulkOptInOutParams.md#note)
- [suppressLog](types_asset.AssetBulkOptInOutParams.md#suppresslog)
- [transactionParams](types_asset.AssetBulkOptInOutParams.md#transactionparams)
- [validateBalances](types_asset.AssetBulkOptInOutParams.md#validatebalances)

## Properties

### account

• **account**: [`SendTransactionFrom`](../modules/types_transaction.md#sendtransactionfrom)

The account to opt in/out for

#### Defined in

[src/types/asset.ts:31](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/asset.ts#L31)

___

### assetIds

• **assetIds**: `number`[]

The IDs of the assets to opt in for / out of

#### Defined in

[src/types/asset.ts:33](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/asset.ts#L33)

___

### maxFee

• `Optional` **maxFee**: [`AlgoAmount`](../classes/types_amount.AlgoAmount.md)

The maximum fee that you are happy to pay per transaction (default: unbounded) - if this is set it's possible the transaction could get rejected during network congestion

#### Defined in

[src/types/asset.ts:41](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/asset.ts#L41)

___

### note

• `Optional` **note**: [`TransactionNote`](../modules/types_transaction.md#transactionnote)

The (optional) transaction note

#### Defined in

[src/types/asset.ts:39](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/asset.ts#L39)

___

### suppressLog

• `Optional` **suppressLog**: `boolean`

Whether to suppress log messages from transaction send, default: do not suppress

#### Defined in

[src/types/asset.ts:43](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/asset.ts#L43)

___

### transactionParams

• `Optional` **transactionParams**: `SuggestedParams`

Optional transaction parameters

#### Defined in

[src/types/asset.ts:37](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/asset.ts#L37)

___

### validateBalances

• `Optional` **validateBalances**: `boolean`

Whether or not to validate the opt-in/out is valid before issuing transactions; default = true

#### Defined in

[src/types/asset.ts:35](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/asset.ts#L35)
