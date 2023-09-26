[@algorandfoundation/algokit-utils](../README.md) / [types/testing](../modules/types_testing.md) / GetTestAccountParams

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

[src/types/testing.ts:35](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/testing.ts#L35)

___

### suppressLog

• `Optional` **suppressLog**: `boolean`

Whether to suppress the log (which includes a mnemonic) or not (default: do not suppress the log)

#### Defined in

[src/types/testing.ts:37](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/testing.ts#L37)
