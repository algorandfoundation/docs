[@algorandfoundation/algokit-utils](../index.md) / [types/debugging](../modules/types_debugging.md) / AVMDebuggerSourceMapEntry

# Class: AVMDebuggerSourceMapEntry

[types/debugging](../modules/types_debugging.md).AVMDebuggerSourceMapEntry

AVM debugger source map entry class.

## Table of contents

### Constructors

- [constructor](types_debugging.AVMDebuggerSourceMapEntry.md#constructor)

### Properties

- [location](types_debugging.AVMDebuggerSourceMapEntry.md#location)
- [programHash](types_debugging.AVMDebuggerSourceMapEntry.md#programhash)

### Methods

- [equals](types_debugging.AVMDebuggerSourceMapEntry.md#equals)
- [toString](types_debugging.AVMDebuggerSourceMapEntry.md#tostring)

## Constructors

### constructor

• **new AVMDebuggerSourceMapEntry**(`location`, `programHash`): [`AVMDebuggerSourceMapEntry`](types_debugging.AVMDebuggerSourceMapEntry.md)

Create an AVM debugger source map entry.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `location` | `string` | The location of the file the source map is for. |
| `programHash` | `string` | The hash of the TEAL binary. |

#### Returns

[`AVMDebuggerSourceMapEntry`](types_debugging.AVMDebuggerSourceMapEntry.md)

#### Defined in

[src/types/debugging.ts:23](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/debugging.ts#L23)

## Properties

### location

• **location**: `string`

The location of the file the source map is for.

#### Defined in

[src/types/debugging.ts:24](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/debugging.ts#L24)

___

### programHash

• **programHash**: `string`

The hash of the TEAL binary.

#### Defined in

[src/types/debugging.ts:25](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/debugging.ts#L25)

## Methods

### equals

▸ **equals**(`other`): `boolean`

#### Parameters

| Name | Type |
| :------ | :------ |
| `other` | [`AVMDebuggerSourceMapEntry`](types_debugging.AVMDebuggerSourceMapEntry.md) |

#### Returns

`boolean`

#### Defined in

[src/types/debugging.ts:28](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/debugging.ts#L28)

___

### toString

▸ **toString**(): `string`

#### Returns

`string`

#### Defined in

[src/types/debugging.ts:32](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/debugging.ts#L32)
