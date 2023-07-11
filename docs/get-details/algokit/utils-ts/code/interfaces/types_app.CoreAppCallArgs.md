[@algorandfoundation/algokit-utils](../README.md) / [types/app](../modules/types_app.md) / CoreAppCallArgs

# Interface: CoreAppCallArgs

[types/app](../modules/types_app.md).CoreAppCallArgs

Common app call arguments for ABI and non-ABI (raw) calls

## Hierarchy

- **`CoreAppCallArgs`**

  ↳ [`RawAppCallArgs`](types_app.RawAppCallArgs.md)

## Table of contents

### Properties

- [boxes](types_app.CoreAppCallArgs.md#boxes)
- [lease](types_app.CoreAppCallArgs.md#lease)

## Properties

### boxes

• `Optional` **boxes**: ([`BoxReference`](types_app.BoxReference.md) \| [`BoxIdentifier`](../modules/types_app.md#boxidentifier) \| `BoxReference`)[]

Any box references to load

#### Defined in

[src/types/app.ts:72](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L72)

___

### lease

• `Optional` **lease**: `string` \| `Uint8Array`

The optional lease for the transaction

#### Defined in

[src/types/app.ts:70](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L70)
