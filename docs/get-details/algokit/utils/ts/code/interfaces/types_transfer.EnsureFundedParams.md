[@algorandfoundation/algokit-utils](../index.md) / [types/transfer](../modules/types_transfer.md) / EnsureFundedParams

# Interface: EnsureFundedParams

[types/transfer](../modules/types_transfer.md).EnsureFundedParams

Parameters for `ensureFunded` call.

## Hierarchy

- [`SendTransactionParams`](types_transaction.SendTransactionParams.md)

  ↳ **`EnsureFundedParams`**

## Table of contents

### Properties

- [accountToFund](types_transfer.EnsureFundedParams.md#accounttofund)
- [atc](types_transfer.EnsureFundedParams.md#atc)
- [fee](types_transfer.EnsureFundedParams.md#fee)
- [fundingSource](types_transfer.EnsureFundedParams.md#fundingsource)
- [lease](types_transfer.EnsureFundedParams.md#lease)
- [maxFee](types_transfer.EnsureFundedParams.md#maxfee)
- [maxRoundsToWaitForConfirmation](types_transfer.EnsureFundedParams.md#maxroundstowaitforconfirmation)
- [minFundingIncrement](types_transfer.EnsureFundedParams.md#minfundingincrement)
- [minSpendingBalance](types_transfer.EnsureFundedParams.md#minspendingbalance)
- [note](types_transfer.EnsureFundedParams.md#note)
- [skipSending](types_transfer.EnsureFundedParams.md#skipsending)
- [skipWaiting](types_transfer.EnsureFundedParams.md#skipwaiting)
- [suppressLog](types_transfer.EnsureFundedParams.md#suppresslog)
- [transactionParams](types_transfer.EnsureFundedParams.md#transactionparams)

## Properties

### accountToFund

• **accountToFund**: `string` \| [`SendTransactionFrom`](../modules/types_transaction.md#sendtransactionfrom)

The account to fund

#### Defined in

[src/types/transfer.ts:40](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transfer.ts#L40)

___

### atc

• `Optional` **atc**: `AtomicTransactionComposer`

An optional `AtomicTransactionComposer` to add the transaction to, if specified then `skipSending: undefined` has the same effect as `skipSending: true`

#### Inherited from

[SendTransactionParams](types_transaction.SendTransactionParams.md).[atc](types_transaction.SendTransactionParams.md#atc)

#### Defined in

[src/types/transaction.ts:35](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L35)

___

### fee

• `Optional` **fee**: [`AlgoAmount`](../classes/types_amount.AlgoAmount.md)

The flat fee you want to pay, useful for covering extra fees in a transaction group or app call

#### Inherited from

[SendTransactionParams](types_transaction.SendTransactionParams.md).[fee](types_transaction.SendTransactionParams.md#fee)

#### Defined in

[src/types/transaction.ts:39](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L39)

___

### fundingSource

• `Optional` **fundingSource**: [`SendTransactionFrom`](../modules/types_transaction.md#sendtransactionfrom) \| [`TestNetDispenserApiClient`](../classes/types_dispenser_client.TestNetDispenserApiClient.md)

The account to use as a funding source, will default to using the dispenser account returned by `algokit.getDispenserAccount`

#### Defined in

[src/types/transfer.ts:42](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transfer.ts#L42)

___

### lease

• `Optional` **lease**: `string` \| `Uint8Array`

An (optional) [transaction lease](https://developer.algorand.org/articles/leased-transactions-securing-advanced-smart-contract-design/) to apply

#### Defined in

[src/types/transfer.ts:52](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transfer.ts#L52)

___

### maxFee

• `Optional` **maxFee**: [`AlgoAmount`](../classes/types_amount.AlgoAmount.md)

The maximum fee that you are happy to pay (default: unbounded) - if this is set it's possible the transaction could get rejected during network congestion

#### Inherited from

[SendTransactionParams](types_transaction.SendTransactionParams.md).[maxFee](types_transaction.SendTransactionParams.md#maxfee)

#### Defined in

[src/types/transaction.ts:41](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L41)

___

### maxRoundsToWaitForConfirmation

• `Optional` **maxRoundsToWaitForConfirmation**: `number`

The maximum number of rounds to wait for confirmation, only applies if `skipWaiting` is `undefined` or `false`, default: wait up to 5 rounds

#### Inherited from

[SendTransactionParams](types_transaction.SendTransactionParams.md).[maxRoundsToWaitForConfirmation](types_transaction.SendTransactionParams.md#maxroundstowaitforconfirmation)

#### Defined in

[src/types/transaction.ts:43](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L43)

___

### minFundingIncrement

• `Optional` **minFundingIncrement**: [`AlgoAmount`](../classes/types_amount.AlgoAmount.md)

When issuing a funding amount, the minimum amount to transfer (avoids many small transfers if this gets called often on an active account)

#### Defined in

[src/types/transfer.ts:46](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transfer.ts#L46)

___

### minSpendingBalance

• **minSpendingBalance**: [`AlgoAmount`](../classes/types_amount.AlgoAmount.md)

The minimum balance of ALGOs that the account should have available to spend (i.e. on top of minimum balance requirement)

#### Defined in

[src/types/transfer.ts:44](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transfer.ts#L44)

___

### note

• `Optional` **note**: [`TransactionNote`](../modules/types_transaction.md#transactionnote)

The (optional) transaction note, default: "Funding account to meet minimum requirement"

#### Defined in

[src/types/transfer.ts:50](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transfer.ts#L50)

___

### skipSending

• `Optional` **skipSending**: `boolean`

Whether to skip signing and sending the transaction to the chain (default: transaction signed and sent to chain, unless `atc` specified)
and instead just return the raw transaction, e.g. so you can add it to a group of transactions

#### Inherited from

[SendTransactionParams](types_transaction.SendTransactionParams.md).[skipSending](types_transaction.SendTransactionParams.md#skipsending)

#### Defined in

[src/types/transaction.ts:31](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L31)

___

### skipWaiting

• `Optional` **skipWaiting**: `boolean`

Whether to skip waiting for the submitted transaction (only relevant if `skipSending` is `false` or unset)

#### Inherited from

[SendTransactionParams](types_transaction.SendTransactionParams.md).[skipWaiting](types_transaction.SendTransactionParams.md#skipwaiting)

#### Defined in

[src/types/transaction.ts:33](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L33)

___

### suppressLog

• `Optional` **suppressLog**: `boolean`

Whether to suppress log messages from transaction send, default: do not suppress

#### Inherited from

[SendTransactionParams](types_transaction.SendTransactionParams.md).[suppressLog](types_transaction.SendTransactionParams.md#suppresslog)

#### Defined in

[src/types/transaction.ts:37](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transaction.ts#L37)

___

### transactionParams

• `Optional` **transactionParams**: `SuggestedParams`

Optional transaction parameters

#### Defined in

[src/types/transfer.ts:48](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/transfer.ts#L48)
