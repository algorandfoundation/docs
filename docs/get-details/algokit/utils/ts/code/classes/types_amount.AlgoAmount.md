[@algorandfoundation/algokit-utils](../index.md) / [types/amount](../modules/types_amount.md) / AlgoAmount

# Class: AlgoAmount

[types/amount](../modules/types_amount.md).AlgoAmount

Wrapper class to ensure safe, explicit conversion between µAlgos, Algos and numbers

## Table of contents

### Constructors

- [constructor](types_amount.AlgoAmount.md#constructor)

### Properties

- [amountInMicroAlgos](types_amount.AlgoAmount.md#amountinmicroalgos)

### Accessors

- [algos](types_amount.AlgoAmount.md#algos)
- [microAlgos](types_amount.AlgoAmount.md#microalgos)

### Methods

- [toString](types_amount.AlgoAmount.md#tostring)
- [valueOf](types_amount.AlgoAmount.md#valueof)
- [Algos](types_amount.AlgoAmount.md#algos-1)
- [MicroAlgos](types_amount.AlgoAmount.md#microalgos-1)

## Constructors

### constructor

• **new AlgoAmount**(`amount`)

#### Parameters

| Name | Type |
| :------ | :------ |
| `amount` | { `algos`: `number`  } \| { `microAlgos`: `number`  } |

#### Defined in

[src/types/amount.ts:17](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/amount.ts#L17)

## Properties

### amountInMicroAlgos

• `Private` **amountInMicroAlgos**: `number`

#### Defined in

[src/types/amount.ts:5](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/amount.ts#L5)

## Accessors

### algos

• `get` **algos**(): `number`

Return the amount as a number in Algos

#### Returns

`number`

#### Defined in

[src/types/amount.ts:13](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/amount.ts#L13)

___

### microAlgos

• `get` **microAlgos**(): `number`

Return the amount as a number in µAlgos

#### Returns

`number`

#### Defined in

[src/types/amount.ts:8](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/amount.ts#L8)

## Methods

### toString

▸ **toString**(): `string`

#### Returns

`string`

#### Defined in

[src/types/amount.ts:21](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/amount.ts#L21)

___

### valueOf

▸ **valueOf**(): `number`

valueOf allows you to use `AlgoAmount` in comparison operations such as `<` and `>=` etc.,
but it's not recommended to use this to convert to a number, it's much safer to explicitly call
the algos or microAlgos properties

#### Returns

`number`

#### Defined in

[src/types/amount.ts:29](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/amount.ts#L29)

___

### Algos

▸ `Static` **Algos**(`amount`): [`AlgoAmount`](types_amount.AlgoAmount.md)

Create a `AlgoAmount` object representing the given number of Algos

#### Parameters

| Name | Type |
| :------ | :------ |
| `amount` | `number` |

#### Returns

[`AlgoAmount`](types_amount.AlgoAmount.md)

#### Defined in

[src/types/amount.ts:34](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/amount.ts#L34)

___

### MicroAlgos

▸ `Static` **MicroAlgos**(`amount`): [`AlgoAmount`](types_amount.AlgoAmount.md)

Create a `AlgoAmount` object representing the given number of µAlgos

#### Parameters

| Name | Type |
| :------ | :------ |
| `amount` | `number` |

#### Returns

[`AlgoAmount`](types_amount.AlgoAmount.md)

#### Defined in

[src/types/amount.ts:39](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/amount.ts#L39)
