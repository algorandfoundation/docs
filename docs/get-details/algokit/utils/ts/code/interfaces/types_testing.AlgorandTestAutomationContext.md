[@algorandfoundation/algokit-utils](../index.md) / [types/testing](../modules/types_testing.md) / AlgorandTestAutomationContext

# Interface: AlgorandTestAutomationContext

[types/testing](../modules/types_testing.md).AlgorandTestAutomationContext

Test automation context.

## Table of contents

### Properties

- [algod](types_testing.AlgorandTestAutomationContext.md#algod)
- [generateAccount](types_testing.AlgorandTestAutomationContext.md#generateaccount)
- [indexer](types_testing.AlgorandTestAutomationContext.md#indexer)
- [kmd](types_testing.AlgorandTestAutomationContext.md#kmd)
- [testAccount](types_testing.AlgorandTestAutomationContext.md#testaccount)
- [transactionLogger](types_testing.AlgorandTestAutomationContext.md#transactionlogger)
- [waitForIndexer](types_testing.AlgorandTestAutomationContext.md#waitforindexer)
- [waitForIndexerTransaction](types_testing.AlgorandTestAutomationContext.md#waitforindexertransaction)

## Properties

### algod

• **algod**: `default`

Algod client instance that will log transactions in `transactionLogger`

#### Defined in

[src/types/testing.ts:13](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/testing.ts#L13)

___

### generateAccount

• **generateAccount**: (`params`: [`GetTestAccountParams`](types_testing.GetTestAccountParams.md)) => `Promise`<`default`\>

#### Type declaration

▸ (`params`): `Promise`<`default`\>

Generate and fund an additional ephemerally created account

##### Parameters

| Name | Type |
| :------ | :------ |
| `params` | [`GetTestAccountParams`](types_testing.GetTestAccountParams.md) |

##### Returns

`Promise`<`default`\>

#### Defined in

[src/types/testing.ts:23](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/testing.ts#L23)

___

### indexer

• **indexer**: `default`

Indexer client instance

#### Defined in

[src/types/testing.ts:15](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/testing.ts#L15)

___

### kmd

• **kmd**: `default`

KMD client instance

#### Defined in

[src/types/testing.ts:17](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/testing.ts#L17)

___

### testAccount

• **testAccount**: `default`

Default, funded test account that is ephemerally created

#### Defined in

[src/types/testing.ts:21](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/testing.ts#L21)

___

### transactionLogger

• **transactionLogger**: [`TransactionLogger`](../classes/testing.TransactionLogger.md)

Transaction logger that will log transaction IDs for all transactions issued by `algod`

#### Defined in

[src/types/testing.ts:19](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/testing.ts#L19)

___

### waitForIndexer

• **waitForIndexer**: () => `Promise`<`void`\>

#### Type declaration

▸ (): `Promise`<`void`\>

Wait for the indexer to catch up with all transactions logged by `transactionLogger`

##### Returns

`Promise`<`void`\>

#### Defined in

[src/types/testing.ts:25](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/testing.ts#L25)

___

### waitForIndexerTransaction

• **waitForIndexerTransaction**: (`transactionId`: `string`) => `Promise`<[`TransactionLookupResult`](types_indexer.TransactionLookupResult.md)\>

#### Type declaration

▸ (`transactionId`): `Promise`<[`TransactionLookupResult`](types_indexer.TransactionLookupResult.md)\>

Wait for the indexer to catch up with the given transaction ID

##### Parameters

| Name | Type |
| :------ | :------ |
| `transactionId` | `string` |

##### Returns

`Promise`<[`TransactionLookupResult`](types_indexer.TransactionLookupResult.md)\>

#### Defined in

[src/types/testing.ts:27](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/testing.ts#L27)
