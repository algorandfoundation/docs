[@algorandfoundation/algokit-utils](../index.md) / [testing](../modules/testing.md) / TestLogger

# Class: TestLogger

[testing](../modules/testing.md).TestLogger

Exposes an AlgoKit logger which captures log messages, while wrapping an original logger.
This is useful for automated testing.

## Implements

- [`Logger`](../modules/types_logging.md#logger)

## Table of contents

### Constructors

- [constructor](testing.TestLogger.md#constructor)

### Properties

- [logs](testing.TestLogger.md#logs)
- [originalLogger](testing.TestLogger.md#originallogger)

### Accessors

- [capturedLogs](testing.TestLogger.md#capturedlogs)

### Methods

- [clear](testing.TestLogger.md#clear)
- [debug](testing.TestLogger.md#debug)
- [error](testing.TestLogger.md#error)
- [getLogSnapshot](testing.TestLogger.md#getlogsnapshot)
- [info](testing.TestLogger.md#info)
- [verbose](testing.TestLogger.md#verbose)
- [warn](testing.TestLogger.md#warn)

## Constructors

### constructor

• **new TestLogger**(`originalLogger?`)

Create a new test logger that wraps the given logger if provided.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `originalLogger?` | [`Logger`](../modules/types_logging.md#logger) | The optional original logger to wrap. |

#### Defined in

[src/testing/test-logger.ts:16](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/testing/test-logger.ts#L16)

## Properties

### logs

• `Private` **logs**: `string`[]

#### Defined in

[src/testing/test-logger.ts:10](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/testing/test-logger.ts#L10)

___

### originalLogger

• `Private` **originalLogger**: `undefined` \| [`Logger`](../modules/types_logging.md#logger)

#### Defined in

[src/testing/test-logger.ts:9](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/testing/test-logger.ts#L9)

## Accessors

### capturedLogs

• `get` **capturedLogs**(): `string`[]

Returns all logs captured thus far.

#### Returns

`string`[]

#### Defined in

[src/testing/test-logger.ts:22](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/testing/test-logger.ts#L22)

## Methods

### clear

▸ **clear**(): `void`

Clears all logs captured so far.

#### Returns

`void`

#### Defined in

[src/testing/test-logger.ts:27](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/testing/test-logger.ts#L27)

___

### debug

▸ **debug**(`message`, `...optionalParams`): `void`

#### Parameters

| Name | Type |
| :------ | :------ |
| `message` | `string` |
| `...optionalParams` | `unknown`[] |

#### Returns

`void`

#### Implementation of

Logger.debug

#### Defined in

[src/testing/test-logger.ts:77](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/testing/test-logger.ts#L77)

___

### error

▸ **error**(`message`, `...optionalParams`): `void`

#### Parameters

| Name | Type |
| :------ | :------ |
| `message` | `string` |
| `...optionalParams` | `unknown`[] |

#### Returns

`void`

#### Implementation of

Logger.error

#### Defined in

[src/testing/test-logger.ts:61](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/testing/test-logger.ts#L61)

___

### getLogSnapshot

▸ **getLogSnapshot**(`config?`): `string`

Returns a captured log snapshot.
This helps ensure that the provided configuration items won't appear
 with random values in the log snapshot, but rather will get substituted with predictable ids.

https://jestjs.io/docs/snapshot-testing#2-tests-should-be-deterministic

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `config?` | [`LogSnapshotConfig`](../interfaces/types_testing.LogSnapshotConfig.md) | The snapshot configuration |

#### Returns

`string`

The snapshotted logs.

**`Example`**

Jest Example
```typescript
const logger = new TestLogger()
...
expect(logger.getLogSnapshot()).toMatchSnapshot()
```

#### Defined in

[src/testing/test-logger.ts:47](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/testing/test-logger.ts#L47)

___

### info

▸ **info**(`message`, `...optionalParams`): `void`

#### Parameters

| Name | Type |
| :------ | :------ |
| `message` | `string` |
| `...optionalParams` | `unknown`[] |

#### Returns

`void`

#### Implementation of

Logger.info

#### Defined in

[src/testing/test-logger.ts:69](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/testing/test-logger.ts#L69)

___

### verbose

▸ **verbose**(`message`, `...optionalParams`): `void`

#### Parameters

| Name | Type |
| :------ | :------ |
| `message` | `string` |
| `...optionalParams` | `unknown`[] |

#### Returns

`void`

#### Implementation of

Logger.verbose

#### Defined in

[src/testing/test-logger.ts:73](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/testing/test-logger.ts#L73)

___

### warn

▸ **warn**(`message`, `...optionalParams`): `void`

#### Parameters

| Name | Type |
| :------ | :------ |
| `message` | `string` |
| `...optionalParams` | `unknown`[] |

#### Returns

`void`

#### Implementation of

Logger.warn

#### Defined in

[src/testing/test-logger.ts:65](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/testing/test-logger.ts#L65)
