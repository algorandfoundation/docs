[@algorandfoundation/algokit-utils](../README.md) / [types/logic-error](../modules/types_logic_error.md) / LogicErrorDetails

# Interface: LogicErrorDetails

[types/logic-error](../modules/types_logic_error.md).LogicErrorDetails

Details about a smart contract logic error

## Table of contents

### Properties

- [desc](types_logic_error.LogicErrorDetails.md#desc)
- [msg](types_logic_error.LogicErrorDetails.md#msg)
- [pc](types_logic_error.LogicErrorDetails.md#pc)
- [traces](types_logic_error.LogicErrorDetails.md#traces)
- [txId](types_logic_error.LogicErrorDetails.md#txid)

## Properties

### desc

• **desc**: `string`

The full error description

#### Defined in

[src/types/logic-error.ts:16](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/logic-error.ts#L16)

___

### msg

• **msg**: `string`

The error message

#### Defined in

[src/types/logic-error.ts:14](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/logic-error.ts#L14)

___

### pc

• **pc**: `number`

The program counter where the error was

#### Defined in

[src/types/logic-error.ts:12](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/logic-error.ts#L12)

___

### traces

• **traces**: `Record`<`string`, `unknown`\>[]

Any trace information included in the error

#### Defined in

[src/types/logic-error.ts:18](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/logic-error.ts#L18)

___

### txId

• **txId**: `string`

The ID of the transaction with the logic error

#### Defined in

[src/types/logic-error.ts:10](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/logic-error.ts#L10)
