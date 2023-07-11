[@algorandfoundation/algokit-utils](../README.md) / [types/app](../modules/types_app.md) / RawAppCallArgs

# Interface: RawAppCallArgs

[types/app](../modules/types_app.md).RawAppCallArgs

App call args with non-ABI (raw) values (minus some processing like encoding strings as binary)

## Hierarchy

- [`CoreAppCallArgs`](types_app.CoreAppCallArgs.md)

  ↳ **`RawAppCallArgs`**

  ↳↳ [`AppClientCallRawArgs`](types_app_client.AppClientCallRawArgs.md)

## Table of contents

### Properties

- [accounts](types_app.RawAppCallArgs.md#accounts)
- [appArgs](types_app.RawAppCallArgs.md#appargs)
- [apps](types_app.RawAppCallArgs.md#apps)
- [assets](types_app.RawAppCallArgs.md#assets)
- [boxes](types_app.RawAppCallArgs.md#boxes)
- [lease](types_app.RawAppCallArgs.md#lease)
- [method](types_app.RawAppCallArgs.md#method)

## Properties

### accounts

• `Optional` **accounts**: (`string` \| `Address`)[]

The address of any accounts to load in

#### Defined in

[src/types/app.ts:82](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L82)

___

### appArgs

• `Optional` **appArgs**: (`string` \| `Uint8Array`)[]

Any application arguments to pass through

#### Defined in

[src/types/app.ts:80](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L80)

___

### apps

• `Optional` **apps**: `number`[]

IDs of any apps to load into the foreignApps array

#### Defined in

[src/types/app.ts:84](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L84)

___

### assets

• `Optional` **assets**: `number`[]

IDs of any assets to load into the foreignAssets array

#### Defined in

[src/types/app.ts:86](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L86)

___

### boxes

• `Optional` **boxes**: ([`BoxReference`](types_app.BoxReference.md) \| [`BoxIdentifier`](../modules/types_app.md#boxidentifier) \| `BoxReference`)[]

Any box references to load

#### Inherited from

[CoreAppCallArgs](types_app.CoreAppCallArgs.md).[boxes](types_app.CoreAppCallArgs.md#boxes)

#### Defined in

[src/types/app.ts:72](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L72)

___

### lease

• `Optional` **lease**: `string` \| `Uint8Array`

The optional lease for the transaction

#### Inherited from

[CoreAppCallArgs](types_app.CoreAppCallArgs.md).[lease](types_app.CoreAppCallArgs.md#lease)

#### Defined in

[src/types/app.ts:70](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L70)

___

### method

• `Optional` **method**: `undefined`

Property to aid intellisense

#### Defined in

[src/types/app.ts:88](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L88)
