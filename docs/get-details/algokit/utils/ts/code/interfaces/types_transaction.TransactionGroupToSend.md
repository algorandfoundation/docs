[@algorandfoundation/algokit-utils](../index.md) / [types/transaction](../modules/types_transaction.md) / TransactionGroupToSend

# Interface: TransactionGroupToSend

[types/transaction](../modules/types_transaction.md).TransactionGroupToSend

A group of transactions to send together as an atomic group
https://developer.algorand.org/docs/get-details/atomic_transfers/

## Table of contents

### Properties

- [sendParams](types_transaction.TransactionGroupToSend.md#sendparams)
- [signer](types_transaction.TransactionGroupToSend.md#signer)
- [transactions](types_transaction.TransactionGroupToSend.md#transactions)

## Properties

### sendParams

• `Optional` **sendParams**: `Omit`<[`SendTransactionParams`](types_transaction.SendTransactionParams.md), ``"fee"`` \| ``"maxFee"`` \| ``"skipSending"`` \| ``"atc"``\>

Any parameters to control the semantics of the send to the network

#### Defined in

[src/types/transaction.ts:108](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L108)

___

### signer

• `Optional` **signer**: [`SendTransactionFrom`](../modules/types_transaction.md#sendtransactionfrom)

Optional signer to pass in, required if at least one transaction provided is just the transaction, ignored otherwise

#### Defined in

[src/types/transaction.ts:115](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L115)

___

### transactions

• **transactions**: (`Transaction` \| [`TransactionToSign`](types_transaction.TransactionToSign.md) \| `Promise`<[`SendTransactionResult`](types_transaction.SendTransactionResult.md)\>)[]

The list of transactions to send, which can either be a raw transaction (in which case `signer` is required),
  the async result of an AlgoKit utils method that returns a `SendTransactionResult` (saves unwrapping the promise, be sure to pass `skipSending: true`, `signer` is also required)
  or the transaction with its signer (`signer` is ignored)

#### Defined in

[src/types/transaction.ts:113](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L113)
