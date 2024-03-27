[@algorandfoundation/algokit-utils](../index.md) / [types/indexer](../modules/types_indexer.md) / AppLocalState

# Interface: AppLocalState

[types/indexer](../modules/types_indexer.md).AppLocalState

Stores local state associated with an application. https://developer.algorand.org/docs/rest-apis/indexer/#applicationlocalstate

## Table of contents

### Properties

- [closed-out-at-round](types_indexer.AppLocalState.md#closed-out-at-round)
- [deleted](types_indexer.AppLocalState.md#deleted)
- [id](types_indexer.AppLocalState.md#id)
- [key-value](types_indexer.AppLocalState.md#key-value)
- [opted-in-at-round](types_indexer.AppLocalState.md#opted-in-at-round)
- [schema](types_indexer.AppLocalState.md#schema)

## Properties

### closed-out-at-round

• `Optional` **closed-out-at-round**: `number`

Round when account closed out of the application.

#### Defined in

[src/types/indexer.ts:805](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L805)

___

### deleted

• `Optional` **deleted**: `boolean`

Whether or not the application local state is currently deleted from its account.

#### Defined in

[src/types/indexer.ts:807](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L807)

___

### id

• **id**: `number`

The application which this local state is for.

#### Defined in

[src/types/indexer.ts:809](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L809)

___

### key-value

• `Optional` **key-value**: `TealKeyValue`[]

[tkv] storage.

#### Defined in

[src/types/indexer.ts:811](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L811)

___

### opted-in-at-round

• `Optional` **opted-in-at-round**: `number`

Round when the account opted into the application.

#### Defined in

[src/types/indexer.ts:813](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L813)

___

### schema

• **schema**: [`StateSchema`](types_indexer.StateSchema.md)

[hsch] schema.

#### Defined in

[src/types/indexer.ts:815](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L815)
