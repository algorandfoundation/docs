[@algorandfoundation/algokit-utils](../index.md) / [types/app-client](../modules/types_app_client.md) / AppClientCallRawArgs

# Interface: AppClientCallRawArgs

[types/app-client](../modules/types_app_client.md).AppClientCallRawArgs

App call args with non-ABI (raw) values (minus some processing like encoding strings as binary)

## Hierarchy

- [`RawAppCallArgs`](types_app.RawAppCallArgs.md)

  ↳ **`AppClientCallRawArgs`**

## Table of contents

### Properties

- [accounts](types_app_client.AppClientCallRawArgs.md#accounts)
- [appArgs](types_app_client.AppClientCallRawArgs.md#appargs)
- [apps](types_app_client.AppClientCallRawArgs.md#apps)
- [assets](types_app_client.AppClientCallRawArgs.md#assets)
- [boxes](types_app_client.AppClientCallRawArgs.md#boxes)
- [lease](types_app_client.AppClientCallRawArgs.md#lease)
- [method](types_app_client.AppClientCallRawArgs.md#method)
- [rekeyTo](types_app_client.AppClientCallRawArgs.md#rekeyto)

## Properties

### accounts

• `Optional` **accounts**: (`string` \| `Address`)[]

The address of any accounts to load in

#### Inherited from

[RawAppCallArgs](types_app.RawAppCallArgs.md).[accounts](types_app.RawAppCallArgs.md#accounts)

#### Defined in

[src/types/app.ts:73](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L73)

___

### appArgs

• `Optional` **appArgs**: (`string` \| `Uint8Array`)[]

Any application arguments to pass through

#### Inherited from

[RawAppCallArgs](types_app.RawAppCallArgs.md).[appArgs](types_app.RawAppCallArgs.md#appargs)

#### Defined in

[src/types/app.ts:90](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L90)

___

### apps

• `Optional` **apps**: `number`[]

IDs of any apps to load into the foreignApps array

#### Inherited from

[RawAppCallArgs](types_app.RawAppCallArgs.md).[apps](types_app.RawAppCallArgs.md#apps)

#### Defined in

[src/types/app.ts:75](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L75)

___

### assets

• `Optional` **assets**: `number`[]

IDs of any assets to load into the foreignAssets array

#### Inherited from

[RawAppCallArgs](types_app.RawAppCallArgs.md).[assets](types_app.RawAppCallArgs.md#assets)

#### Defined in

[src/types/app.ts:77](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L77)

___

### boxes

• `Optional` **boxes**: ([`BoxReference`](types_app.BoxReference.md) \| [`BoxIdentifier`](../modules/types_app.md#boxidentifier) \| `BoxReference`)[]

Any box references to load

#### Inherited from

[RawAppCallArgs](types_app.RawAppCallArgs.md).[boxes](types_app.RawAppCallArgs.md#boxes)

#### Defined in

[src/types/app.ts:71](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L71)

___

### lease

• `Optional` **lease**: `string` \| `Uint8Array`

The optional lease for the transaction

#### Inherited from

[RawAppCallArgs](types_app.RawAppCallArgs.md).[lease](types_app.RawAppCallArgs.md#lease)

#### Defined in

[src/types/app.ts:69](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L69)

___

### method

• `Optional` **method**: `undefined`

Property to aid intellisense

#### Inherited from

[RawAppCallArgs](types_app.RawAppCallArgs.md).[method](types_app.RawAppCallArgs.md#method)

#### Defined in

[src/types/app.ts:92](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L92)

___

### rekeyTo

• `Optional` **rekeyTo**: `string` \| [`SendTransactionFrom`](../modules/types_transaction.md#sendtransactionfrom)

Optional account / account address that should be authorised to transact on behalf of the from account the app call is sent from after this transaction.

**Note:** Use with extreme caution and review the [official rekey guidance](https://developer.algorand.org/docs/get-details/accounts/rekey/) first.

#### Inherited from

[RawAppCallArgs](types_app.RawAppCallArgs.md).[rekeyTo](types_app.RawAppCallArgs.md#rekeyto)

#### Defined in

[src/types/app.ts:82](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L82)
