[@algorandfoundation/algokit-utils](../index.md) / [types/config](../modules/types_config.md) / UpdatableConfig

# Class: UpdatableConfig

[types/config](../modules/types_config.md).UpdatableConfig

Updatable AlgoKit config

## Implements

- `Readonly`<[`Config`](../interfaces/types_config.Config.md)\>

## Table of contents

### Constructors

- [constructor](types_config.UpdatableConfig.md#constructor)

### Properties

- [config](types_config.UpdatableConfig.md#config)

### Accessors

- [debug](types_config.UpdatableConfig.md#debug)
- [logger](types_config.UpdatableConfig.md#logger)

### Methods

- [configure](types_config.UpdatableConfig.md#configure)
- [getLogger](types_config.UpdatableConfig.md#getlogger)
- [withDebug](types_config.UpdatableConfig.md#withdebug)

## Constructors

### constructor

• **new UpdatableConfig**()

#### Defined in

[src/types/config.ts:50](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/config.ts#L50)

## Properties

### config

• `Private` **config**: [`Config`](../interfaces/types_config.Config.md)

#### Defined in

[src/types/config.ts:13](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/config.ts#L13)

## Accessors

### debug

• `get` **debug**(): `boolean`

#### Returns

`boolean`

#### Implementation of

Readonly.debug

#### Defined in

[src/types/config.ts:19](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/config.ts#L19)

___

### logger

• `get` **logger**(): [`Logger`](../modules/types_logging.md#logger)

#### Returns

[`Logger`](../modules/types_logging.md#logger)

#### Implementation of

Readonly.logger

#### Defined in

[src/types/config.ts:15](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/config.ts#L15)

## Methods

### configure

▸ **configure**(`newConfig`): `void`

Update the AlgoKit configuration with your own configuration settings

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `newConfig` | `Partial`<[`Config`](../interfaces/types_config.Config.md)\> | Partial or complete config to replace |

#### Returns

`void`

#### Defined in

[src/types/config.ts:61](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/config.ts#L61)

___

### getLogger

▸ **getLogger**(`returnNullLogger?`): [`Logger`](../modules/types_logging.md#logger)

Returns the current logger, or the null logger if true is passed in to `returnNullLogger`

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `returnNullLogger?` | `boolean` | Whether or not to return the null logger |

#### Returns

[`Logger`](../modules/types_logging.md#logger)

The requested logger

#### Defined in

[src/types/config.ts:28](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/config.ts#L28)

___

### withDebug

▸ **withDebug**(`lambda`): `void`

Temporarily run with debug set to true.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `lambda` | () => `unknown` | A lambda expression with code to run with debug config set to true |

#### Returns

`void`

#### Defined in

[src/types/config.ts:40](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/config.ts#L40)
