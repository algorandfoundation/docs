[@algorandfoundation/algokit-utils](../index.md) / [types/transaction](../modules/types_transaction.md) / ConfirmedTransactionResults

# Interface: ConfirmedTransactionResults

[types/transaction](../modules/types_transaction.md).ConfirmedTransactionResults

The result of sending and confirming one or more transactions, but where there is a primary transaction of interest

## Hierarchy

- [`SendTransactionResult`](types_transaction.SendTransactionResult.md)

- [`SendTransactionResults`](types_transaction.SendTransactionResults.md)

  ↳ **`ConfirmedTransactionResults`**

## Table of contents

### Properties

- [confirmation](types_transaction.ConfirmedTransactionResults.md#confirmation)
- [confirmations](types_transaction.ConfirmedTransactionResults.md#confirmations)
- [transaction](types_transaction.ConfirmedTransactionResults.md#transaction)
- [transactions](types_transaction.ConfirmedTransactionResults.md#transactions)

## Properties

### confirmation

• **confirmation**: `PendingTransactionResponse`

The response from sending and waiting for the primary transaction

#### Overrides

[SendTransactionResult](types_transaction.SendTransactionResult.md).[confirmation](types_transaction.SendTransactionResult.md#confirmation)

#### Defined in

[src/types/transaction.ts:78](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L78)

___

### confirmations

• **confirmations**: `PendingTransactionResponse`[]

The response from sending and waiting for the transactions

#### Overrides

[SendTransactionResults](types_transaction.SendTransactionResults.md).[confirmations](types_transaction.SendTransactionResults.md#confirmations)

#### Defined in

[src/types/transaction.ts:80](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L80)

___

### transaction

• **transaction**: `Transaction`

The transaction

#### Inherited from

[SendTransactionResult](types_transaction.SendTransactionResult.md).[transaction](types_transaction.SendTransactionResult.md#transaction)

#### Defined in

[src/types/transaction.ts:44](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L44)

___

### transactions

• **transactions**: `Transaction`[]

The transactions that have been prepared and/or sent

#### Inherited from

[SendTransactionResults](types_transaction.SendTransactionResults.md).[transactions](types_transaction.SendTransactionResults.md#transactions)

#### Defined in

[src/types/transaction.ts:52](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L52)
