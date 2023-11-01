[@algorandfoundation/algokit-utils](../index.md) / [types/indexer](../modules/types_indexer.md) / AssetFreezeTransactionResult

# Interface: AssetFreezeTransactionResult

[types/indexer](../modules/types_indexer.md).AssetFreezeTransactionResult

Fields for an asset freeze transaction. https://developer.algorand.org/docs/rest-apis/indexer/#transactionassetfreeze

## Table of contents

### Properties

- [address](types_indexer.AssetFreezeTransactionResult.md#address)
- [asset-id](types_indexer.AssetFreezeTransactionResult.md#asset-id)
- [new-freeze-status](types_indexer.AssetFreezeTransactionResult.md#new-freeze-status)

## Properties

### address

• **address**: `string`

[fadd] Address of the account whose asset is being frozen or thawed.

#### Defined in

[src/types/indexer.ts:295](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L295)

___

### asset-id

• **asset-id**: `number`

[faid] ID of the asset being frozen or thawed.

#### Defined in

[src/types/indexer.ts:297](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L297)

___

### new-freeze-status

• **new-freeze-status**: `boolean`

[afrz] The new freeze status.

#### Defined in

[src/types/indexer.ts:299](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L299)
