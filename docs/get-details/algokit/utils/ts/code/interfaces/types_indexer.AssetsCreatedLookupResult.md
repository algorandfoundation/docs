[@algorandfoundation/algokit-utils](../index.md) / [types/indexer](../modules/types_indexer.md) / AssetsCreatedLookupResult

# Interface: AssetsCreatedLookupResult

[types/indexer](../modules/types_indexer.md).AssetsCreatedLookupResult

Indexer result for an account's created assets, https://developer.algorand.org/docs/rest-apis/indexer/#get-v2accountsaccount-idcreated-assets

## Table of contents

### Properties

- [assets](types_indexer.AssetsCreatedLookupResult.md#assets)
- [current-round](types_indexer.AssetsCreatedLookupResult.md#current-round)
- [next-token](types_indexer.AssetsCreatedLookupResult.md#next-token)

## Properties

### assets

• **assets**: [`AssetResult`](types_indexer.AssetResult.md)[]

The returned assets

#### Defined in

[src/types/indexer.ts:38](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L38)

___

### current-round

• **current-round**: `number`

Round at which the results were computed.

#### Defined in

[src/types/indexer.ts:34](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L34)

___

### next-token

• **next-token**: `string`

Used for pagination, when making another request provide this token with the next parameter.

#### Defined in

[src/types/indexer.ts:36](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L36)
