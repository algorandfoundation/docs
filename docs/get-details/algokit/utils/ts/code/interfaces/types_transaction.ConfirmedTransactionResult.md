[@algorandfoundation/algokit-utils](../index.md) / [types/transaction](../modules/types_transaction.md) / ConfirmedTransactionResult

# Interface: ConfirmedTransactionResult

[types/transaction](../modules/types_transaction.md).ConfirmedTransactionResult

The result of sending and confirming a transaction

## Hierarchy

- [`SendTransactionResult`](types_transaction.SendTransactionResult.md)

  ↳ **`ConfirmedTransactionResult`**

## Table of contents

### Properties

- [confirmation](types_transaction.ConfirmedTransactionResult.md#confirmation)
- [transaction](types_transaction.ConfirmedTransactionResult.md#transaction)

## Properties

### confirmation

• **confirmation**: `PendingTransactionResponse`

The response from sending and waiting for the transaction

#### Overrides

[SendTransactionResult](types_transaction.SendTransactionResult.md).[confirmation](types_transaction.SendTransactionResult.md#confirmation)

#### Defined in

[src/types/transaction.ts:79](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L79)

___

### transaction

• **transaction**: `Transaction`

The transaction

#### Inherited from

[SendTransactionResult](types_transaction.SendTransactionResult.md).[transaction](types_transaction.SendTransactionResult.md#transaction)

#### Defined in

[src/types/transaction.ts:51](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L51)
