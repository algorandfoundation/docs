[@algorandfoundation/algokit-utils](../index.md) / [types/indexer](../modules/types_indexer.md) / PaymentTransactionResult

# Interface: PaymentTransactionResult

[types/indexer](../modules/types_indexer.md).PaymentTransactionResult

Fields for a payment transaction https://developer.algorand.org/docs/rest-apis/indexer/#transactionpayment

## Table of contents

### Properties

- [amount](types_indexer.PaymentTransactionResult.md#amount)
- [close-amount](types_indexer.PaymentTransactionResult.md#close-amount)
- [close-remainder-to](types_indexer.PaymentTransactionResult.md#close-remainder-to)
- [receiver](types_indexer.PaymentTransactionResult.md#receiver)

## Properties

### amount

• **amount**: `number`

[amt] number of MicroAlgos intended to be transferred.

#### Defined in

[src/types/indexer.ts:255](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L255)

___

### close-amount

• `Optional` **close-amount**: `number`

Number of MicroAlgos that were sent to the close-remainder-to address when closing the sender account.

#### Defined in

[src/types/indexer.ts:257](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L257)

___

### close-remainder-to

• `Optional` **close-remainder-to**: `string`

[close] when set, indicates that the sending account should be closed and all remaining funds be transferred to this address.

#### Defined in

[src/types/indexer.ts:259](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L259)

___

### receiver

• **receiver**: `string`

[rcv] receiver's address.

#### Defined in

[src/types/indexer.ts:261](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L261)
