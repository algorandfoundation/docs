[@algorandfoundation/algokit-utils](../index.md) / [types/debugging](../modules/types_debugging.md) / AVMDebuggerSourceMap

# Class: AVMDebuggerSourceMap

[types/debugging](../modules/types_debugging.md).AVMDebuggerSourceMap

## Table of contents

### Constructors

- [constructor](types_debugging.AVMDebuggerSourceMap.md#constructor)

### Properties

- [txnGroupSources](types_debugging.AVMDebuggerSourceMap.md#txngroupsources)

### Methods

- [toDict](types_debugging.AVMDebuggerSourceMap.md#todict)
- [fromDict](types_debugging.AVMDebuggerSourceMap.md#fromdict)

## Constructors

### constructor

• **new AVMDebuggerSourceMap**(`txnGroupSources`): [`AVMDebuggerSourceMap`](types_debugging.AVMDebuggerSourceMap.md)

#### Parameters

| Name | Type |
| :------ | :------ |
| `txnGroupSources` | [`AVMDebuggerSourceMapEntry`](types_debugging.AVMDebuggerSourceMapEntry.md)[] |

#### Returns

[`AVMDebuggerSourceMap`](types_debugging.AVMDebuggerSourceMap.md)

#### Defined in

[src/types/debugging.ts:29](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/debugging.ts#L29)

## Properties

### txnGroupSources

• **txnGroupSources**: [`AVMDebuggerSourceMapEntry`](types_debugging.AVMDebuggerSourceMapEntry.md)[]

#### Defined in

[src/types/debugging.ts:27](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/debugging.ts#L27)

## Methods

### toDict

▸ **toDict**(): [`AVMDebuggerSourceMapDict`](../interfaces/types_debugging.AVMDebuggerSourceMapDict.md)

#### Returns

[`AVMDebuggerSourceMapDict`](../interfaces/types_debugging.AVMDebuggerSourceMapDict.md)

#### Defined in

[src/types/debugging.ts:39](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/debugging.ts#L39)

___

### fromDict

▸ **fromDict**(`data`): [`AVMDebuggerSourceMap`](types_debugging.AVMDebuggerSourceMap.md)

#### Parameters

| Name | Type |
| :------ | :------ |
| `data` | [`AVMDebuggerSourceMapDict`](../interfaces/types_debugging.AVMDebuggerSourceMapDict.md) |

#### Returns

[`AVMDebuggerSourceMap`](types_debugging.AVMDebuggerSourceMap.md)

#### Defined in

[src/types/debugging.ts:33](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/debugging.ts#L33)
