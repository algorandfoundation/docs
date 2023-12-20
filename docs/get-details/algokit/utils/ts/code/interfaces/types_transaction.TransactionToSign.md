[@algorandfoundation/algokit-utils](../index.md) / [types/transaction](../modules/types_transaction.md) / TransactionToSign

# Interface: TransactionToSign

[types/transaction](../modules/types_transaction.md).TransactionToSign

Defines an unsigned transaction that will appear in a group of transactions along with its signing information

## Table of contents

### Properties

- [signer](types_transaction.TransactionToSign.md#signer)
- [transaction](types_transaction.TransactionToSign.md#transaction)

## Properties

### signer

• **signer**: [`SendTransactionFrom`](../modules/types_transaction.md#sendtransactionfrom)

The account to use to sign the transaction, either an account (with private key loaded) or a logic signature account

#### Defined in

[src/types/transaction.ts:105](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L105)

___

### transaction

• **transaction**: `Transaction`

The unsigned transaction to sign and send

#### Defined in

[src/types/transaction.ts:103](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L103)
