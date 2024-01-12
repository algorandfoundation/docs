[@algorandfoundation/algokit-utils](../index.md) / [types/app](../modules/types_app.md) / BoxValueRequestParams

# Interface: BoxValueRequestParams

[types/app](../modules/types_app.md).BoxValueRequestParams

Parameters to get and decode a box value as an ABI type.

## Table of contents

### Properties

- [appId](types_app.BoxValueRequestParams.md#appid)
- [boxName](types_app.BoxValueRequestParams.md#boxname)
- [type](types_app.BoxValueRequestParams.md#type)

## Properties

### appId

• **appId**: `number` \| `bigint`

The ID of the app return box names for

#### Defined in

[src/types/app.ts:353](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L353)

___

### boxName

• **boxName**: `string` \| `Uint8Array` \| [`BoxName`](types_app.BoxName.md)

The name of the box to return either as a string, binary array or `BoxName`

#### Defined in

[src/types/app.ts:355](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L355)

___

### type

• **type**: `ABIType`

The ABI type to decode the value using

#### Defined in

[src/types/app.ts:357](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L357)
