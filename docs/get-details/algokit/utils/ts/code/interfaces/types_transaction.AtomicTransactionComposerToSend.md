[@algorandfoundation/algokit-utils](../index.md) / [types/transaction](../modules/types_transaction.md) / AtomicTransactionComposerToSend

# Interface: AtomicTransactionComposerToSend

[types/transaction](../modules/types_transaction.md).AtomicTransactionComposerToSend

An `AtomicTransactionComposer` with transactions to send.

## Table of contents

### Properties

- [atc](types_transaction.AtomicTransactionComposerToSend.md#atc)
- [sendParams](types_transaction.AtomicTransactionComposerToSend.md#sendparams)

## Properties

### atc

• **atc**: `AtomicTransactionComposer`

The `AtomicTransactionComposer` with transactions loaded to send

#### Defined in

[src/types/transaction.ts:126](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L126)

___

### sendParams

• `Optional` **sendParams**: `Omit`\<[`SendTransactionParams`](types_transaction.SendTransactionParams.md), ``"fee"`` \| ``"maxFee"`` \| ``"skipSending"`` \| ``"atc"``\>

Any parameters to control the semantics of the send to the network

#### Defined in

[src/types/transaction.ts:128](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L128)
