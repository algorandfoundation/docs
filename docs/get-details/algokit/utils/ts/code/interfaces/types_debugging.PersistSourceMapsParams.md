[@algorandfoundation/algokit-utils](../index.md) / [types/debugging](../modules/types_debugging.md) / PersistSourceMapsParams

# Interface: PersistSourceMapsParams

[types/debugging](../modules/types_debugging.md).PersistSourceMapsParams

Parameters to a call that persists source maps

## Table of contents

### Properties

- [client](types_debugging.PersistSourceMapsParams.md#client)
- [projectRoot](types_debugging.PersistSourceMapsParams.md#projectroot)
- [sources](types_debugging.PersistSourceMapsParams.md#sources)
- [withSources](types_debugging.PersistSourceMapsParams.md#withsources)

## Properties

### client

• **client**: `default`

An Algodv2 client to perform the compilation.

#### Defined in

[src/types/debugging.ts:148](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/debugging.ts#L148)

___

### projectRoot

• **projectRoot**: `string`

The root directory of the project.

#### Defined in

[src/types/debugging.ts:146](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/debugging.ts#L146)

___

### sources

• **sources**: [`PersistSourceMapInput`](../classes/types_debugging.PersistSourceMapInput.md)[]

An array of PersistSourceMapInput objects. Each object can either contain rawTeal, in which case the function will execute a compile to obtain byte code, or it can accept an object of type CompiledTeal provided by algokit, which is used for source codes that have already been compiled and contain the traces.

#### Defined in

[src/types/debugging.ts:144](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/debugging.ts#L144)

___

### withSources

• `Optional` **withSources**: `boolean`

A boolean indicating whether to include the source files in the output.

#### Defined in

[src/types/debugging.ts:150](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/debugging.ts#L150)
