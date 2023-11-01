[@algorandfoundation/algokit-utils](../index.md) / [testing](../modules/testing.md) / TransactionLogger

# Class: TransactionLogger

[testing](../modules/testing.md).TransactionLogger

Allows you to keep track of Algorand transaction IDs by wrapping an `Algodv2` in a proxy.
Useful for automated tests.

## Table of contents

### Constructors

- [constructor](testing.TransactionLogger.md#constructor)

### Properties

- [\_sentTransactionIds](testing.TransactionLogger.md#_senttransactionids)

### Accessors

- [sentTransactionIds](testing.TransactionLogger.md#senttransactionids)

### Methods

- [capture](testing.TransactionLogger.md#capture)
- [clear](testing.TransactionLogger.md#clear)
- [logRawTransaction](testing.TransactionLogger.md#lograwtransaction)
- [waitForIndexer](testing.TransactionLogger.md#waitforindexer)

## Constructors

### constructor

• **new TransactionLogger**()

## Properties

### \_sentTransactionIds

• `Private` **\_sentTransactionIds**: `string`[] = `[]`

#### Defined in

[src/testing/transaction-logger.ts:9](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/testing/transaction-logger.ts#L9)

## Accessors

### sentTransactionIds

• `get` **sentTransactionIds**(): readonly `string`[]

The list of transaction IDs that has been logged thus far.

#### Returns

readonly `string`[]

#### Defined in

[src/testing/transaction-logger.ts:14](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/testing/transaction-logger.ts#L14)

## Methods

### capture

▸ **capture**(`algod`): `default`

Return a proxy that wraps the given Algodv2 with this transaction logger.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `algod` | `default` | The `Algodv2` to wrap |

#### Returns

`default`

The wrapped `Algodv2`, any transactions sent using this algod instance will be logged by this transaction logger

#### Defined in

[src/testing/transaction-logger.ts:45](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/testing/transaction-logger.ts#L45)

___

### clear

▸ **clear**(): `void`

Clear all logged IDs.

#### Returns

`void`

#### Defined in

[src/testing/transaction-logger.ts:21](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/testing/transaction-logger.ts#L21)

___

### logRawTransaction

▸ **logRawTransaction**(`signedTransactions`): `void`

The method that captures raw transactions and stores the transaction IDs.

#### Parameters

| Name | Type |
| :------ | :------ |
| `signedTransactions` | `Uint8Array` \| `Uint8Array`[] |

#### Returns

`void`

#### Defined in

[src/testing/transaction-logger.ts:28](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/testing/transaction-logger.ts#L28)

___

### waitForIndexer

▸ **waitForIndexer**(`indexer`): `Promise`<`void`\>

Wait until all logged transactions IDs appear in the given `Indexer`.

#### Parameters

| Name | Type |
| :------ | :------ |
| `indexer` | `default` |

#### Returns

`Promise`<`void`\>

#### Defined in

[src/testing/transaction-logger.ts:50](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/testing/transaction-logger.ts#L50)
