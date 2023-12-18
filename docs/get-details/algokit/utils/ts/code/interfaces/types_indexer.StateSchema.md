[@algorandfoundation/algokit-utils](../index.md) / [types/indexer](../modules/types_indexer.md) / StateSchema

# Interface: StateSchema

[types/indexer](../modules/types_indexer.md).StateSchema

Represents a [apls] local-state or [apgs] global-state schema.
https://developer.algorand.org/docs/rest-apis/indexer/#stateschema

These schemas determine how much storage may be used in a local-state or global-state for an application.

The more space used, the larger minimum balance must be maintained in the account holding the data.

## Table of contents

### Properties

- [num-byte-slice](types_indexer.StateSchema.md#num-byte-slice)
- [num-uint](types_indexer.StateSchema.md#num-uint)

## Properties

### num-byte-slice

• **num-byte-slice**: `number`

Maximum number of TEAL byte slices that may be stored in the key/value store.

#### Defined in

[src/types/indexer.ts:487](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L487)

___

### num-uint

• **num-uint**: `number`

Maximum number of TEAL uints that may be stored in the key/value store.

#### Defined in

[src/types/indexer.ts:489](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L489)
