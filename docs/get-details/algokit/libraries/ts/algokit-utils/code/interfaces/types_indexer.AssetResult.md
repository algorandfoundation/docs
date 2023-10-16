[@algorandfoundation/algokit-utils](../README.md) / [types/indexer](../modules/types_indexer.md) / AssetResult

# Interface: AssetResult

[types/indexer](../modules/types_indexer.md).AssetResult

Specifies both the unique identifier and the parameters for an asset. https://developer.algorand.org/docs/rest-apis/indexer/#asset

## Table of contents

### Properties

- [created-at-round](types_indexer.AssetResult.md#created-at-round)
- [deleted](types_indexer.AssetResult.md#deleted)
- [destroyed-at-round](types_indexer.AssetResult.md#destroyed-at-round)
- [index](types_indexer.AssetResult.md#index)
- [params](types_indexer.AssetResult.md#params)

## Properties

### created-at-round

• `Optional` **created-at-round**: `number`

Round during which this asset was created.

#### Defined in

[src/types/indexer.ts:352](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L352)

___

### deleted

• `Optional` **deleted**: `boolean`

Whether or not this asset is currently deleted.

#### Defined in

[src/types/indexer.ts:350](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L350)

___

### destroyed-at-round

• `Optional` **destroyed-at-round**: `number`

Round during which this asset was destroyed.

#### Defined in

[src/types/indexer.ts:354](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L354)

___

### index

• **index**: `number`

Unique asset identifier.

#### Defined in

[src/types/indexer.ts:348](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L348)

___

### params

• **params**: [`AssetParams`](types_indexer.AssetParams.md)

The parameters for the asset

#### Defined in

[src/types/indexer.ts:356](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L356)
