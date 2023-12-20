[@algorandfoundation/algokit-utils](../index.md) / [types/debugging](../modules/types_debugging.md) / PersistSourceMapInput

# Class: PersistSourceMapInput

[types/debugging](../modules/types_debugging.md).PersistSourceMapInput

Class representing a debugger source maps input for persistence.

Note: rawTeal and compiledTeal are mutually exclusive. Only one of them should be provided.

## Table of contents

### Constructors

- [constructor](types_debugging.PersistSourceMapInput.md#constructor)

### Properties

- [\_fileName](types_debugging.PersistSourceMapInput.md#_filename)
- [\_rawTeal](types_debugging.PersistSourceMapInput.md#_rawteal)
- [appName](types_debugging.PersistSourceMapInput.md#appname)
- [compiledTeal](types_debugging.PersistSourceMapInput.md#compiledteal)

### Accessors

- [fileName](types_debugging.PersistSourceMapInput.md#filename)
- [rawTeal](types_debugging.PersistSourceMapInput.md#rawteal)

### Methods

- [stripTealExtension](types_debugging.PersistSourceMapInput.md#striptealextension)
- [fromCompiledTeal](types_debugging.PersistSourceMapInput.md#fromcompiledteal)
- [fromRawTeal](types_debugging.PersistSourceMapInput.md#fromrawteal)

## Constructors

### constructor

• **new PersistSourceMapInput**(`appName`, `fileName`, `rawTeal?`, `compiledTeal?`): [`PersistSourceMapInput`](types_debugging.PersistSourceMapInput.md)

#### Parameters

| Name | Type |
| :------ | :------ |
| `appName` | `string` |
| `fileName` | `string` |
| `rawTeal?` | `string` |
| `compiledTeal?` | [`CompiledTeal`](../interfaces/types_app.CompiledTeal.md) |

#### Returns

[`PersistSourceMapInput`](types_debugging.PersistSourceMapInput.md)

#### Defined in

[src/types/debugging.ts:55](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/debugging.ts#L55)

## Properties

### \_fileName

• `Private` **\_fileName**: `string`

#### Defined in

[src/types/debugging.ts:52](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/debugging.ts#L52)

___

### \_rawTeal

• `Private` `Optional` **\_rawTeal**: `string`

#### Defined in

[src/types/debugging.ts:53](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/debugging.ts#L53)

___

### appName

• **appName**: `string`

#### Defined in

[src/types/debugging.ts:50](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/debugging.ts#L50)

___

### compiledTeal

• `Optional` **compiledTeal**: [`CompiledTeal`](../interfaces/types_app.CompiledTeal.md)

#### Defined in

[src/types/debugging.ts:51](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/debugging.ts#L51)

## Accessors

### fileName

• `get` **fileName**(): `string`

#### Returns

`string`

#### Defined in

[src/types/debugging.ts:80](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/debugging.ts#L80)

___

### rawTeal

• `get` **rawTeal**(): `string`

#### Returns

`string`

#### Defined in

[src/types/debugging.ts:70](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/debugging.ts#L70)

## Methods

### stripTealExtension

▸ **stripTealExtension**(`fileName`): `string`

Strips the '.teal' extension from a filename, if present.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `fileName` | `string` | The filename to strip the extension from. |

#### Returns

`string`

The filename without the '.teal' extension.

#### Defined in

[src/types/debugging.ts:90](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/debugging.ts#L90)

___

### fromCompiledTeal

▸ **fromCompiledTeal**(`compiledTeal`, `appName`, `fileName`): [`PersistSourceMapInput`](types_debugging.PersistSourceMapInput.md)

#### Parameters

| Name | Type |
| :------ | :------ |
| `compiledTeal` | [`CompiledTeal`](../interfaces/types_app.CompiledTeal.md) |
| `appName` | `string` |
| `fileName` | `string` |

#### Returns

[`PersistSourceMapInput`](types_debugging.PersistSourceMapInput.md)

#### Defined in

[src/types/debugging.ts:66](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/debugging.ts#L66)

___

### fromRawTeal

▸ **fromRawTeal**(`rawTeal`, `appName`, `fileName`): [`PersistSourceMapInput`](types_debugging.PersistSourceMapInput.md)

#### Parameters

| Name | Type |
| :------ | :------ |
| `rawTeal` | `string` |
| `appName` | `string` |
| `fileName` | `string` |

#### Returns

[`PersistSourceMapInput`](types_debugging.PersistSourceMapInput.md)

#### Defined in

[src/types/debugging.ts:62](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/debugging.ts#L62)
