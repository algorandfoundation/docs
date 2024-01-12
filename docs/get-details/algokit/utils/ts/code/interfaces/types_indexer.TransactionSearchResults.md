[@algorandfoundation/algokit-utils](../index.md) / [types/indexer](../modules/types_indexer.md) / TransactionSearchResults

# Interface: TransactionSearchResults

[types/indexer](../modules/types_indexer.md).TransactionSearchResults

Indexer result for a transaction search, https://developer.algorand.org/docs/rest-apis/indexer/#get-v2transactions

## Table of contents

### Properties

- [current-round](types_indexer.TransactionSearchResults.md#current-round)
- [next-token](types_indexer.TransactionSearchResults.md#next-token)
- [transactions](types_indexer.TransactionSearchResults.md#transactions)

## Properties

### current-round

• **current-round**: `number`

Round at which the results were computed.

#### Defined in

[src/types/indexer.ts:8](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L8)

___

### next-token

• **next-token**: `string`

Used for pagination, when making another request provide this token with the next parameter.

#### Defined in

[src/types/indexer.ts:10](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L10)

___

### transactions

• **transactions**: [`TransactionResult`](types_indexer.TransactionResult.md)[]

The returned transactions

#### Defined in

[src/types/indexer.ts:12](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L12)
