[@algorandfoundation/algokit-utils](../index.md) / [types/app](../modules/types_app.md) / CoreAppCallArgs

# Interface: CoreAppCallArgs

[types/app](../modules/types_app.md).CoreAppCallArgs

Common app call arguments for ABI and non-ABI (raw) calls

## Hierarchy

- **`CoreAppCallArgs`**

  ↳ [`RawAppCallArgs`](types_app.RawAppCallArgs.md)

## Table of contents

### Properties

- [accounts](types_app.CoreAppCallArgs.md#accounts)
- [apps](types_app.CoreAppCallArgs.md#apps)
- [assets](types_app.CoreAppCallArgs.md#assets)
- [boxes](types_app.CoreAppCallArgs.md#boxes)
- [lease](types_app.CoreAppCallArgs.md#lease)
- [rekeyTo](types_app.CoreAppCallArgs.md#rekeyto)

## Properties

### accounts

• `Optional` **accounts**: (`string` \| `Address`)[]

The address of any accounts to load in

#### Defined in

[src/types/app.ts:73](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L73)

___

### apps

• `Optional` **apps**: `number`[]

IDs of any apps to load into the foreignApps array

#### Defined in

[src/types/app.ts:75](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L75)

___

### assets

• `Optional` **assets**: `number`[]

IDs of any assets to load into the foreignAssets array

#### Defined in

[src/types/app.ts:77](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L77)

___

### boxes

• `Optional` **boxes**: ([`BoxReference`](types_app.BoxReference.md) \| [`BoxIdentifier`](../modules/types_app.md#boxidentifier) \| `BoxReference`)[]

Any box references to load

#### Defined in

[src/types/app.ts:71](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L71)

___

### lease

• `Optional` **lease**: `string` \| `Uint8Array`

The optional lease for the transaction

#### Defined in

[src/types/app.ts:69](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L69)

___

### rekeyTo

• `Optional` **rekeyTo**: `string` \| [`SendTransactionFrom`](../modules/types_transaction.md#sendtransactionfrom)

Optional account / account address that should be authorised to transact on behalf of the from account the app call is sent from after this transaction.

**Note:** Use with extreme caution and review the [official rekey guidance](https://developer.algorand.org/docs/get-details/accounts/rekey/) first.

#### Defined in

[src/types/app.ts:82](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L82)
