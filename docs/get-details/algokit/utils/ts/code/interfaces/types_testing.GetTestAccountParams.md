[@algorandfoundation/algokit-utils](../index.md) / [types/testing](../modules/types_testing.md) / GetTestAccountParams

# Interface: GetTestAccountParams

[types/testing](../modules/types_testing.md).GetTestAccountParams

Parameters for the `getTestAccount` function.

## Table of contents

### Properties

- [initialFunds](types_testing.GetTestAccountParams.md#initialfunds)
- [suppressLog](types_testing.GetTestAccountParams.md#suppresslog)

## Properties

### initialFunds

• **initialFunds**: [`AlgoAmount`](../classes/types_amount.AlgoAmount.md)

Initial funds to ensure the account has

#### Defined in

[src/types/testing.ts:40](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/testing.ts#L40)

___

### suppressLog

• `Optional` **suppressLog**: `boolean`

Whether to suppress the log (which includes a mnemonic) or not (default: do not suppress the log)

#### Defined in

[src/types/testing.ts:42](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/testing.ts#L42)
