[@algorandfoundation/algokit-utils](../README.md) / [types/app-client](../modules/types_app_client.md) / AppClientCallABIArgs

# Interface: AppClientCallABIArgs

[types/app-client](../modules/types_app_client.md).AppClientCallABIArgs

## Hierarchy

- `Omit`<[`ABIAppCallArgs`](../modules/types_app.md#abiappcallargs), ``"method"``\>

  ↳ **`AppClientCallABIArgs`**

## Table of contents

### Properties

- [boxes](types_app_client.AppClientCallABIArgs.md#boxes)
- [lease](types_app_client.AppClientCallABIArgs.md#lease)
- [method](types_app_client.AppClientCallABIArgs.md#method)
- [methodArgs](types_app_client.AppClientCallABIArgs.md#methodargs)

## Properties

### boxes

• `Optional` **boxes**: ([`BoxReference`](types_app.BoxReference.md) \| [`BoxIdentifier`](../modules/types_app.md#boxidentifier) \| `BoxReference`)[]

Any box references to load

#### Inherited from

Omit.boxes

#### Defined in

[src/types/app.ts:72](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L72)

___

### lease

• `Optional` **lease**: `string` \| `Uint8Array`

The optional lease for the transaction

#### Inherited from

Omit.lease

#### Defined in

[src/types/app.ts:70](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L70)

___

### method

• **method**: `string`

If calling an ABI method then either the name of the method, or the ABI signature

#### Defined in

[src/types/app-client.ts:146](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L146)

___

### methodArgs

• **methodArgs**: [`ABIAppCallArg`](../modules/types_app.md#abiappcallarg)[]

The ABI method args to pass in

#### Inherited from

Omit.methodArgs

#### Defined in

[src/types/app.ts:107](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L107)
