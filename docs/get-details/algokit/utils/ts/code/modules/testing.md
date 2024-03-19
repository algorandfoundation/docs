[@algorandfoundation/algokit-utils](../index.md) / testing

# Module: testing

## Table of contents

### Classes

- [TestLogger](../classes/testing.TestLogger.md)
- [TransactionLogger](../classes/testing.TransactionLogger.md)

### Functions

- [algoKitLogCaptureFixture](testing.md#algokitlogcapturefixture)
- [algorandFixture](testing.md#algorandfixture)
- [getTestAccount](testing.md#gettestaccount)
- [runWhenIndexerCaughtUp](testing.md#runwhenindexercaughtup)

## Functions

### algoKitLogCaptureFixture

▸ **algoKitLogCaptureFixture**(): [`AlgoKitLogCaptureFixture`](../interfaces/types_testing.AlgoKitLogCaptureFixture.md)

Creates a test fixture for capturing AlgoKit logs.

#### Returns

[`AlgoKitLogCaptureFixture`](../interfaces/types_testing.AlgoKitLogCaptureFixture.md)

The fixture

**`Example`**

```typescript
const logs = algoKitLogCaptureFixture()

beforeEach(logs.beforeEach)
afterEach(logs.afterEach)

test('My test', () => {
    const capturedLogs = logs.testLogger.capturedLogs
})
```

#### Defined in

[src/testing/fixtures/algokit-log-capture-fixture.ts:22](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/testing/fixtures/algokit-log-capture-fixture.ts#L22)

___

### algorandFixture

▸ **algorandFixture**(`fixtureConfig?`): [`AlgorandFixture`](../interfaces/types_testing.AlgorandFixture.md)

Creates a test fixture for automated testing against Algorand.
By default it tests against an environment variable specified client
 if the standard environment variables are specified, otherwise against
 a default LocalNet instance, but you can pass in an algod, indexer
 and/or kmd if you want to test against an explicitly defined network.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `fixtureConfig?` | [`AlgorandFixtureConfig`](../interfaces/types_testing.AlgorandFixtureConfig.md) | The fixture configuration |

#### Returns

[`AlgorandFixture`](../interfaces/types_testing.AlgorandFixture.md)

The fixture

**`Example`**

```typescript
const algorand = algorandFixture()

beforeEach(algorand.beforeEach, 10_000)

test('My test', async () => {
    const {algod, indexer, testAccount, ...} = algorand.context
    // test things...
})
```

#### Defined in

[src/testing/fixtures/algorand-fixture.ts:38](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/testing/fixtures/algorand-fixture.ts#L38)

▸ **algorandFixture**(`fixtureConfig`, `config`): [`AlgorandFixture`](../interfaces/types_testing.AlgorandFixture.md)

Creates a test fixture for automated testing against Algorand.
By default it tests against an environment variable specified client
 if the standard environment variables are specified, otherwise against
 a default LocalNet instance, but you can pass in an algod, indexer
 and/or kmd if you want to test against an explicitly defined network.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `fixtureConfig` | `undefined` \| [`AlgorandFixtureConfig`](../interfaces/types_testing.AlgorandFixtureConfig.md) | The fixture configuration |
| `config` | [`AlgoConfig`](../interfaces/types_network_client.AlgoConfig.md) | The algo configuration |

#### Returns

[`AlgorandFixture`](../interfaces/types_testing.AlgorandFixture.md)

The fixture

**`Example`**

```typescript
const algorand = algorandFixture(undefined, getConfigFromEnvOrDefaults())

beforeEach(algorand.beforeEach, 10_000)

test('My test', async () => {
    const {algod, indexer, testAccount, ...} = algorand.context
    // test things...
})
```

#### Defined in

[src/testing/fixtures/algorand-fixture.ts:63](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/testing/fixtures/algorand-fixture.ts#L63)

___

### getTestAccount

▸ **getTestAccount**(`param0`, `algod`, `kmd?`): `Promise`\<`Account`\>

Creates an ephemeral Algorand account for the purposes of testing.
Returns a newly created random test account that is funded from the dispenser
DO NOT USE THIS TO CREATE A MAINNET ACCOUNT!
Note: By default this will log the mnemonic of the account.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `param0` | [`GetTestAccountParams`](../interfaces/types_testing.GetTestAccountParams.md) | The config for the test account to generate |
| `algod` | `default` | An algod client |
| `kmd?` | `default` | A KMD client, if not specified then a default KMD client will be loaded from environment variables |

#### Returns

`Promise`\<`Account`\>

The account, with private key loaded

#### Defined in

[src/testing/account.ts:19](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/testing/account.ts#L19)

___

### runWhenIndexerCaughtUp

▸ **runWhenIndexerCaughtUp**\<`T`\>(`run`): `Promise`\<`T`\>

Runs the given indexer call until a 404 error is no longer returned.
Tried every 200ms up to 100 times.
Very rudimentary implementation designed for automated testing.

#### Type parameters

| Name |
| :------ |
| `T` |

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `run` | () => `Promise`\<`T`\> | The code to run |

#### Returns

`Promise`\<`T`\>

The result (as a promise), or throws if the indexer didn't catch up in time

**`Example`**

```typescript
const transaction = await runWhenIndexerCaughtUp(() => indexer.lookupTransactionByID(txnId).do())
```

#### Defined in

[src/testing/indexer.ts:12](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/testing/indexer.ts#L12)
