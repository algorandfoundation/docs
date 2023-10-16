[@algorandfoundation/algokit-utils](../README.md) / [types/indexer](../modules/types_indexer.md) / TransactionLookupResult

# Interface: TransactionLookupResult

[types/indexer](../modules/types_indexer.md).TransactionLookupResult

Indexer result for a transaction lookup, https://developer.algorand.org/docs/rest-apis/indexer/#get-v2transactionstxid

## Table of contents

### Properties

- [current-round](types_indexer.TransactionLookupResult.md#current-round)
- [transaction](types_indexer.TransactionLookupResult.md#transaction)

## Properties

### current-round

• **current-round**: `number`

Round at which the results were computed.

#### Defined in

[src/types/indexer.ts:62](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L62)

___

### transaction

• **transaction**: [`TransactionResult`](types_indexer.TransactionResult.md)

The returned transaction

#### Defined in

[src/types/indexer.ts:64](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L64)
