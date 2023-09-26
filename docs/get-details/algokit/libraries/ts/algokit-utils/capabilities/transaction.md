# Transaction management

Transaction management is one of the core capabilities provided by AlgoKit Utils. It allows you to send single, grouped or Atomic Transaction Composer transactions with consistent and highly configurable semantics, including configurable control of transaction notes (including ARC-0002), logging, fees, multiple sender account types, and sending behaviour.

## `SendTransactionParams`

Any AlgoKit Utils function that needs to sign/send a transaction will generally take all or part of [`SendTransactionParams`](../code/interfaces/types_transaction.SendTransactionParams.md) interface, which represents a standard set of configurations that can be applied to a given transaction or transactions that are to be sent to an Algorand network.

The following parameters are able to be provided, with all of them being optional:

- **skipSending: boolean** - Whether to skip signing and sending the transaction to the chain (default: transaction signed and sent to chain, unless `atc` is specified) and instead just return the raw transaction, e.g. so you can add it to a group of transactions
- **skipWaiting: boolean** - Whether to skip waiting for the submitted transaction (only relevant if `skipSending` is `false` or unset) and turn this transaction send into an async operation
- **atc: AtomicTransactionComposer** - An optional `AtomicTransactionComposer` to add the transaction to, if specified then `skipSending: undefined` has the same effect as `skipSending: true`
- **suppressLog: boolean** - Whether to suppress log messages from transaction send, default: do not suppress
- **fee: AlgoAmount** - The flat fee you want to pay, useful for covering extra fees in a transaction group or app call
- **maxFee: AlgoAmount** - The maximum fee that you are happy to pay (default: unbounded) - if this is set it's possible the transaction could get rejected during network congestion
- **maxRoundsToWaitForConfirmation: number** - The maximum number of rounds to wait for confirmation, only applies if `skipWaiting` is `undefined` or `false`, default: wait up to 5 rounds

## `SendTransactionResult`

All AlgoKit Utils functions that prepare and/or send a transaction will generally return a [`SendTransactionResult` interface](../code/interfaces/types_transaction.SendTransactionResult.md) or some superset of that. This provides a consistent mechanism to interpret the results of a transaction send.

It consists of two properties:

- `transaction`: An `algosdk.Transaction` object that is either ready to send or represents the transaction that was sent
- `confirmation` (optional): An `algosdk.modelsv2.PendingTransactionResponse` object, which is a type-safe wrapper of the return from the algod pending transaction API noting that it will only be returned if the transaction was able to be confirmed (so won't represent a "pending" transaction)

A useful pattern to use to access these properties is destructuring, e.g.:

```typescript
const { transaction, confirmation } = await algokit.sendTransaction(...)
```

There are various variations of the `SendTransactionResult` that are exposed by various functions within AlgoKit Utils, including:

- [`ConfirmedTransactionResult`](../code/interfaces/types_transaction.ConfirmedTransactionResult.md) - Where it's guaranteed that a confirmation will be returned (i.e. `skipSending` is not made available)
- [`ConfirmedTransactionResults`](../code/interfaces/types_transaction.ConfirmedTransactionResults.md) - Where it's both guaranteed that a confirmation will be returned, there is a primary driving transaction, but multiple transactions may be sent (e.g. when making an ABI app call which has dependant transactions)
- [`SendTransactionResults`](../code/interfaces/types_transaction.SendTransactionResults.md) - Where multiple transactions are being sent (`transactions` and `confirmations` are arrays that replace the singular `transaction` and `confirmation`)
- [`SendAtomicTransactionComposerResults`](../code/interfaces/types_transaction.SendAtomicTransactionComposerResults.md) - The result from sending the transactions within an `AtomicTransactionComposer`, it extends `SendTransactionResults` and adds a few other useful properties
- [`AppCallTransactionResult`](../code/interfaces/types_app.AppCallTransactionResult.md) - Result from calling a single app call (which potentially may result in multiple other transaction calls if it was an ABI method with dependant transactions)

## Sending a transaction

AlgoKit Utils provides three core helper methods that allow you to use the `SendTransactionParams` configuration when sending transactions. It also has many other methods that delegate to these underlying core methods such as `algokit.transferAlgos(...)`, etc..

- [`sendTransaction`](../code/modules/index.md#sendtransaction) - Signs and sends (if instructed to) a single transaction using the `SendTransactionParams` configuration
- [`sendGroupOfTransactions`](../code/modules/index.md#sendgroupoftransactions) - Takes an array of transactions to sign with associated `SendTransactionFrom` accounts and signs and sends (if instructed to) them as an atomic transaction group using the `SendTransactionParams` configuration
- [`sendAtomicTransactionComposer`](../code/modules/index.md#sendatomictransactioncomposer) - Signs and sends (if instructed to) the transactions and/or method calls loaded into an `AtomicTransactionComposer` object using the `SendTransactionParams` configuration

All of these methods take `algosdk.Transaction` and/or `algosdk.AtomicTransactionComposer` objects so you can use them with transaction generation mechanisms outside of AlgoKit Utils per the [modularity principle](../README.md#core-principles).

## Helpers

The functionality provided by the transaction capability includes a set of lower level helpers that might be useful in their own right per the [modularity principle](../README.md#core-principles).

### Signing

If you want to sign a transaction there are the [`algokit.signTransaction(transaction, signer)`](../code/modules/index.md#signtransaction) method and [`algokit.getSenderTransactionSigner(sender)`](../code/modules/index.md#getsendertransactionsigner) methods that both work with `SendTransactionFrom` as described in the [Account capability](./account.md).

There are also some methods that take a [`TransactionToSign`](../code/interfaces/types_transaction.TransactionToSign.md), which is the AlgoKit Utils equivalent of `algosdk.TransactionWithSigner`, but has a `SendTransactionFrom` as the signer.

### Waiting

There is a [`algokit.waitForConfirmation(transactionId, maxRoundsToWait, algod)`](../code/modules/index.md#waitforconfirmation) method which helps you wait until a given `algosdk.Transaction` has been confirmed by the network.

### Fees

If you want to control the fees of a transaction before sending then you can use:

- [`algokit.capTransactionFee(transaction, maxAcceptableFee)`](../code/modules/index.md#captransactionfee) - Limit the acceptable maximum fee of a `algosdk.Transaction` or `algosdk.SuggestedParams` to a defined amount of algos.
- [`algokit.controlFees(transaction, feeControl)`](../code/modules/index.md#controlfees) - Allows you to control fees on a `algosdk.Transaction` or `algosdk.SuggestedParams` object either applying a flat fee or a max fee

### Transaction notes

If you want to create an encoded transaction note for adding to a transaction you can use the [`algokit.encodeTransactionNote(note)`](../code/modules/index.md#encodetransactionnote) function. This takes a `TransactionNote` type, which is a union of:

- `null` or `undefined` if there is no note
- `Uint8Array` which is passed straight through
- Data that us turned into JSON, any one of:
  - `string`
  - `number`
  - `any[]`
  - `Record<string, any>`
- A `Arc2TransactionNote` object, which creates an [ARC-0002 compliant transaction note](https://github.com/algorandfoundation/ARCs/blob/main/ARCs/arc-0002.md) and has the following properties:
  - `dAppName` - The name of the app that is generating the note
  - `format`:
    - `m` - Message Pack format
    - `b` - Byte string
    - `u` - UTF-8 string
    - `j` - JSON data
  - `data` either a string or an object that is encoded to JSON (if `format` is `j`)

### Transaction params

If you want to specify transaction params to add to a transaction you can use the [`algokit.getTransactionParams(params, algod)`](../code/modules/index.md#gettransactionparams) method. This let's you pass in an existing params object if one exists or if that's `undefined` then it will retrieve a new params object from the Algod client.
