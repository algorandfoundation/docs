[@algorandfoundation/algokit-utils](../index.md) / types/app-spec

# Module: types/app-spec

## Table of contents

### Interfaces

- [AppSources](../interfaces/types_app_spec.AppSources.md)
- [AppSpec](../interfaces/types_app_spec.AppSpec.md)
- [CallConfig](../interfaces/types_app_spec.CallConfig.md)
- [DeclaredSchemaValueSpec](../interfaces/types_app_spec.DeclaredSchemaValueSpec.md)
- [Hint](../interfaces/types_app_spec.Hint.md)
- [ReservedSchemaValueSpec](../interfaces/types_app_spec.ReservedSchemaValueSpec.md)
- [Schema](../interfaces/types_app_spec.Schema.md)
- [SchemaSpec](../interfaces/types_app_spec.SchemaSpec.md)
- [StateSchemaSpec](../interfaces/types_app_spec.StateSchemaSpec.md)
- [Struct](../interfaces/types_app_spec.Struct.md)

### Type Aliases

- [ABIType](types_app_spec.md#abitype)
- [AVMType](types_app_spec.md#avmtype)
- [CallConfigValue](types_app_spec.md#callconfigvalue)
- [DefaultArgument](types_app_spec.md#defaultargument)
- [FieldName](types_app_spec.md#fieldname)
- [HintSpec](types_app_spec.md#hintspec)
- [StateSchema](types_app_spec.md#stateschema)
- [StructElement](types_app_spec.md#structelement)

## Type Aliases

### ABIType

Ƭ **ABIType**: `string`

The string name of an ABI type

#### Defined in

[src/types/app-spec.ts:65](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-spec.ts#L65)

___

### AVMType

Ƭ **AVMType**: ``"uint64"`` \| ``"bytes"``

AVM data type

#### Defined in

[src/types/app-spec.ts:121](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-spec.ts#L121)

___

### CallConfigValue

Ƭ **CallConfigValue**: ``"NEVER"`` \| ``"CALL"`` \| ``"CREATE"`` \| ``"ALL"``

The various call configs:
 * `NEVER`: Will not be called
 * `CALL`: Can be called during a non-create call i.e. app id != 0
 * `CREATE`: Can be called during a create call i.e. app id = 0
 * `ALL`: Can be during a create OR non-create call

#### Defined in

[src/types/app-spec.ts:36](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-spec.ts#L36)

___

### DefaultArgument

Ƭ **DefaultArgument**: { `data`: `ABIMethodParams` ; `source`: ``"abi-method"``  } \| { `data`: `string` ; `source`: ``"global-state"``  } \| { `data`: `string` ; `source`: ``"local-state"``  } \| { `data`: `string` \| `number` ; `source`: ``"constant"``  }

Defines a strategy for obtaining a default value for a given ABI arg.

#### Defined in

[src/types/app-spec.ts:81](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-spec.ts#L81)

___

### FieldName

Ƭ **FieldName**: `string`

The name of a field

#### Defined in

[src/types/app-spec.ts:62](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-spec.ts#L62)

___

### HintSpec

Ƭ **HintSpec**: `Record`<`string`, [`Hint`](../interfaces/types_app_spec.Hint.md)\>

A lookup of encoded method call spec to hint

#### Defined in

[src/types/app-spec.ts:20](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-spec.ts#L20)

___

### StateSchema

Ƭ **StateSchema**: `Object`

Schema spec summary for global or local storage

#### Type declaration

| Name | Type | Description |
| :------ | :------ | :------ |
| `num_byte_slices` | `number` | Number of byte slots |
| `num_uints` | `number` | Number of uint slots |

#### Defined in

[src/types/app-spec.ts:170](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-spec.ts#L170)

___

### StructElement

Ƭ **StructElement**: [[`FieldName`](types_app_spec.md#fieldname), [`ABIType`](types_app_spec.md#abitype)]

The elements of the struct/tuple: `FieldName`, `ABIType`

#### Defined in

[src/types/app-spec.ts:68](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-spec.ts#L68)
