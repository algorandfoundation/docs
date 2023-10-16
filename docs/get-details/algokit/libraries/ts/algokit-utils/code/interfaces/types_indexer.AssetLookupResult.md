[@algorandfoundation/algokit-utils](../README.md) / [types/indexer](../modules/types_indexer.md) / AssetLookupResult

# Interface: AssetLookupResult

[types/indexer](../modules/types_indexer.md).AssetLookupResult

Indexer result for an asset lookup, https://developer.algorand.org/docs/rest-apis/indexer/#get-v2assetsasset-id

## Table of contents

### Properties

- [asset](types_indexer.AssetLookupResult.md#asset)
- [current-round](types_indexer.AssetLookupResult.md#current-round)

## Properties

### asset

• **asset**: [`AssetResult`](types_indexer.AssetResult.md)

The returned asset

#### Defined in

[src/types/indexer.ts:56](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L56)

___

### current-round

• **current-round**: `number`

Round at which the results were computed.

#### Defined in

[src/types/indexer.ts:54](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L54)
