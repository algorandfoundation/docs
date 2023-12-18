[@algorandfoundation/algokit-utils](../index.md) / [types/indexer](../modules/types_indexer.md) / AssetConfigTransactionResult

# Interface: AssetConfigTransactionResult

[types/indexer](../modules/types_indexer.md).AssetConfigTransactionResult

Fields for asset allocation, re-configuration, and destruction.
https://developer.algorand.org/docs/rest-apis/indexer/#transactionassetconfig

A zero value for asset-id indicates asset creation. A zero value for the params indicates asset destruction.

## Table of contents

### Properties

- [asset-id](types_indexer.AssetConfigTransactionResult.md#asset-id)
- [params](types_indexer.AssetConfigTransactionResult.md#params)

## Properties

### asset-id

• **asset-id**: `number`

[xaid] ID of the asset being configured or empty if creating.

#### Defined in

[src/types/indexer.ts:289](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L289)

___

### params

• **params**: [`AssetParams`](types_indexer.AssetParams.md)

[apar] the parameters for the asset.

#### Defined in

[src/types/indexer.ts:291](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L291)
