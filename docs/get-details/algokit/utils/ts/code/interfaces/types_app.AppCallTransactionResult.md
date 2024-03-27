[@algorandfoundation/algokit-utils](../index.md) / [types/app](../modules/types_app.md) / AppCallTransactionResult

# Interface: AppCallTransactionResult

[types/app](../modules/types_app.md).AppCallTransactionResult

Result from calling an app

## Hierarchy

- [`AppCallTransactionResultOfType`](types_app.AppCallTransactionResultOfType.md)\<[`ABIReturn`](../modules/types_app.md#abireturn)\>

  ↳ **`AppCallTransactionResult`**

## Table of contents

### Properties

- [confirmation](types_app.AppCallTransactionResult.md#confirmation)
- [confirmations](types_app.AppCallTransactionResult.md#confirmations)
- [return](types_app.AppCallTransactionResult.md#return)
- [transaction](types_app.AppCallTransactionResult.md#transaction)
- [transactions](types_app.AppCallTransactionResult.md#transactions)

## Properties

### confirmation

• `Optional` **confirmation**: `PendingTransactionResponse`

The response if the transaction was sent and waited for

#### Inherited from

[AppCallTransactionResultOfType](types_app.AppCallTransactionResultOfType.md).[confirmation](types_app.AppCallTransactionResultOfType.md#confirmation)

#### Defined in

[src/types/transaction.ts:53](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L53)

___

### confirmations

• `Optional` **confirmations**: `PendingTransactionResponse`[]

The responses if the transactions were sent and waited for,
the index of the confirmation will match the index of the underlying transaction

#### Inherited from

[AppCallTransactionResultOfType](types_app.AppCallTransactionResultOfType.md).[confirmations](types_app.AppCallTransactionResultOfType.md#confirmations)

#### Defined in

[src/types/transaction.ts:63](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L63)

___

### return

• `Optional` **return**: [`ABIReturn`](../modules/types_app.md#abireturn)

If an ABI method was called the processed return value

#### Inherited from

[AppCallTransactionResultOfType](types_app.AppCallTransactionResultOfType.md).[return](types_app.AppCallTransactionResultOfType.md#return)

#### Defined in

[src/types/app.ts:209](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L209)

___

### transaction

• **transaction**: `Transaction`

The transaction

#### Inherited from

[AppCallTransactionResultOfType](types_app.AppCallTransactionResultOfType.md).[transaction](types_app.AppCallTransactionResultOfType.md#transaction)

#### Defined in

[src/types/transaction.ts:51](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L51)

___

### transactions

• **transactions**: `Transaction`[]

The transactions that have been prepared and/or sent

#### Inherited from

[AppCallTransactionResultOfType](types_app.AppCallTransactionResultOfType.md).[transactions](types_app.AppCallTransactionResultOfType.md#transactions)

#### Defined in

[src/types/transaction.ts:59](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L59)
