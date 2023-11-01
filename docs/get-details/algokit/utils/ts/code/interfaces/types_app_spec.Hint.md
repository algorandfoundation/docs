[@algorandfoundation/algokit-utils](../index.md) / [types/app-spec](../modules/types_app_spec.md) / Hint

# Interface: Hint

[types/app-spec](../modules/types_app_spec.md).Hint

Hint information for a given method call to allow client generation

## Table of contents

### Properties

- [call\_config](types_app_spec.Hint.md#call_config)
- [default\_arguments](types_app_spec.Hint.md#default_arguments)
- [read\_only](types_app_spec.Hint.md#read_only)
- [structs](types_app_spec.Hint.md#structs)

## Properties

### call\_config

• **call\_config**: [`CallConfig`](types_app_spec.CallConfig.md)

#### Defined in

[src/types/app-spec.ts:58](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-spec.ts#L58)

___

### default\_arguments

• `Optional` **default\_arguments**: `Record`<`string`, [`DefaultArgument`](../modules/types_app_spec.md#defaultargument)\>

#### Defined in

[src/types/app-spec.ts:57](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-spec.ts#L57)

___

### read\_only

• `Optional` **read\_only**: `boolean`

#### Defined in

[src/types/app-spec.ts:56](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-spec.ts#L56)

___

### structs

• `Optional` **structs**: `Record`<`string`, [`Struct`](types_app_spec.Struct.md)\>

Any user-defined struct/tuple types used in the method call, keyed by parameter name or `output` for return type

#### Defined in

[src/types/app-spec.ts:55](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-spec.ts#L55)
