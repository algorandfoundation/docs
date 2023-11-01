[@algorandfoundation/algokit-utils](../index.md) / [types/indexer](../modules/types_indexer.md) / ApplicationLookupResult

# Interface: ApplicationLookupResult

[types/indexer](../modules/types_indexer.md).ApplicationLookupResult

Indexer result for an application lookup, https://developer.algorand.org/docs/rest-apis/indexer/#get-v2applicationsapplication-id

## Table of contents

### Properties

- [application](types_indexer.ApplicationLookupResult.md#application)
- [current-round](types_indexer.ApplicationLookupResult.md#current-round)

## Properties

### application

• **application**: [`ApplicationResult`](types_indexer.ApplicationResult.md)

The returned application

#### Defined in

[src/types/indexer.ts:72](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L72)

___

### current-round

• **current-round**: `number`

Round at which the results were computed.

#### Defined in

[src/types/indexer.ts:70](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L70)
