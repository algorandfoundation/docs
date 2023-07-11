[@algorandfoundation/algokit-utils](../README.md) / [types/app-client](../modules/types_app_client.md) / AppClientCallCoreParams

# Interface: AppClientCallCoreParams

[types/app-client](../modules/types_app_client.md).AppClientCallCoreParams

Common (core) parameters to construct a ApplicationClient contract call

## Table of contents

### Properties

- [note](types_app_client.AppClientCallCoreParams.md#note)
- [sendParams](types_app_client.AppClientCallCoreParams.md#sendparams)
- [sender](types_app_client.AppClientCallCoreParams.md#sender)

## Properties

### note

• `Optional` **note**: [`TransactionNote`](../modules/types_transaction.md#transactionnote)

The transaction note for the smart contract call

#### Defined in

[src/types/app-client.ts:157](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L157)

___

### sendParams

• `Optional` **sendParams**: [`SendTransactionParams`](types_transaction.SendTransactionParams.md)

Parameters to control transaction sending

#### Defined in

[src/types/app-client.ts:159](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L159)

___

### sender

• `Optional` **sender**: [`SendTransactionFrom`](../modules/types_transaction.md#sendtransactionfrom)

The optional sender to send the transaction from, will use the application client's default sender by default if specified

#### Defined in

[src/types/app-client.ts:155](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L155)
