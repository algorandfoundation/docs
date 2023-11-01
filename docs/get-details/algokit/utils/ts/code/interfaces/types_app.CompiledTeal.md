[@algorandfoundation/algokit-utils](../index.md) / [types/app](../modules/types_app.md) / CompiledTeal

# Interface: CompiledTeal

[types/app](../modules/types_app.md).CompiledTeal

Information about a compiled teal program

## Table of contents

### Properties

- [compiled](types_app.CompiledTeal.md#compiled)
- [compiledBase64ToBytes](types_app.CompiledTeal.md#compiledbase64tobytes)
- [compiledHash](types_app.CompiledTeal.md#compiledhash)
- [sourceMap](types_app.CompiledTeal.md#sourcemap)
- [teal](types_app.CompiledTeal.md#teal)

## Properties

### compiled

• **compiled**: `string`

The compiled code

#### Defined in

[src/types/app.ts:194](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L194)

___

### compiledBase64ToBytes

• **compiledBase64ToBytes**: `Uint8Array`

The base64 encoded code as a byte array

#### Defined in

[src/types/app.ts:198](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L198)

___

### compiledHash

• **compiledHash**: `string`

The has returned by the compiler

#### Defined in

[src/types/app.ts:196](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L196)

___

### sourceMap

• **sourceMap**: `SourceMap`

Source map from the compilation

#### Defined in

[src/types/app.ts:200](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L200)

___

### teal

• **teal**: `string`

Original TEAL code

#### Defined in

[src/types/app.ts:192](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L192)
