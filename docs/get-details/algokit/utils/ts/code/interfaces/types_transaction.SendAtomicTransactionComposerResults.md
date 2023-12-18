[@algorandfoundation/algokit-utils](../index.md) / [types/transaction](../modules/types_transaction.md) / SendAtomicTransactionComposerResults

# Interface: SendAtomicTransactionComposerResults

[types/transaction](../modules/types_transaction.md).SendAtomicTransactionComposerResults

The result of preparing and/or sending multiple transactions using an `AtomicTransactionComposer`

## Hierarchy

- [`SendTransactionResults`](types_transaction.SendTransactionResults.md)

  ↳ **`SendAtomicTransactionComposerResults`**

## Table of contents

### Properties

- [confirmations](types_transaction.SendAtomicTransactionComposerResults.md#confirmations)
- [groupId](types_transaction.SendAtomicTransactionComposerResults.md#groupid)
- [returns](types_transaction.SendAtomicTransactionComposerResults.md#returns)
- [transactions](types_transaction.SendAtomicTransactionComposerResults.md#transactions)
- [txIds](types_transaction.SendAtomicTransactionComposerResults.md#txids)

## Properties

### confirmations

• `Optional` **confirmations**: `PendingTransactionResponse`[]

The responses if the transactions were sent and waited for,
the index of the confirmation will match the index of the underlying transaction

#### Inherited from

[SendTransactionResults](types_transaction.SendTransactionResults.md).[confirmations](types_transaction.SendTransactionResults.md#confirmations)

#### Defined in

[src/types/transaction.ts:61](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L61)

___

### groupId

• **groupId**: `string`

base64 encoded representation of the group ID of the atomic group

#### Defined in

[src/types/transaction.ts:67](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L67)

___

### returns

• `Optional` **returns**: [`ABIReturn`](../modules/types_app.md#abireturn)[]

If ABI method(s) were called the processed return values

#### Defined in

[src/types/transaction.ts:71](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L71)

___

### transactions

• **transactions**: `Transaction`[]

The transactions that have been prepared and/or sent

#### Inherited from

[SendTransactionResults](types_transaction.SendTransactionResults.md).[transactions](types_transaction.SendTransactionResults.md#transactions)

#### Defined in

[src/types/transaction.ts:57](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L57)

___

### txIds

• **txIds**: `string`[]

The transaction IDs that have been prepared and/or sent

#### Defined in

[src/types/transaction.ts:69](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L69)
