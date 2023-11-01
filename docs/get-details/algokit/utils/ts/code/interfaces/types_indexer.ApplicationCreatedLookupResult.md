[@algorandfoundation/algokit-utils](../index.md) / [types/indexer](../modules/types_indexer.md) / ApplicationCreatedLookupResult

# Interface: ApplicationCreatedLookupResult

[types/indexer](../modules/types_indexer.md).ApplicationCreatedLookupResult

Indexer result for an account's created applications, https://developer.algorand.org/docs/rest-apis/indexer/#get-v2accountsaccount-idcreated-applications

## Table of contents

### Properties

- [applications](types_indexer.ApplicationCreatedLookupResult.md#applications)
- [current-round](types_indexer.ApplicationCreatedLookupResult.md#current-round)
- [next-token](types_indexer.ApplicationCreatedLookupResult.md#next-token)

## Properties

### applications

• **applications**: [`ApplicationResult`](types_indexer.ApplicationResult.md)[]

The returned applications

#### Defined in

[src/types/indexer.ts:48](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L48)

___

### current-round

• **current-round**: `number`

Round at which the results were computed.

#### Defined in

[src/types/indexer.ts:44](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L44)

___

### next-token

• **next-token**: `string`

Used for pagination, when making another request provide this token with the next parameter.

#### Defined in

[src/types/indexer.ts:46](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L46)
