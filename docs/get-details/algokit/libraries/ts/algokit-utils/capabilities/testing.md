# Automated testing

Automated testing is a higher-order use case capability provided by AlgoKit Utils that builds on top of the core capabilities. It allows you to use terse, robust automated testing primitives that work across any testing framework (including jest and vitest) to facilitate fixture management, quickly generating isolated and funded test accounts, transaction logging, indexer wait management and log capture.

To see some usage examples check out the all of the [automated tests](../../src/) and the various \*.spec.ts files (AlgoKit Utils [dogfoods](https://en.wikipedia.org/wiki/Eating_your_own_dog_food) it's own testing library). Alternatively, you can see an example of using this library to test a smart contract with [the tests](https://github.com/algorandfoundation/nft_voting_tool/blob/main/src/algorand/smart_contracts/tests/voting.spec.ts) for the [on-chain voting tool](https://github.com/algorandfoundation/nft_voting_tool#readme).

## Module import

The testing capability is not exposed from the [root algokit module](../code/modules/index.md) so there is a clear separation between testing functionality and non-testing functionality.

To access all of the functionality in the testing capability individually, you can import the [testing module](../code/modules/testing.md):

```typescript
import * as algotesting from '@algorandfoundation/algokit-utils/testing'
```

## Algorand fixture

In general, the only entrypoint you will need to use the testing capability is just by importing the `algorandFixture` since it exposes the rest of the functionality in a manner that is easy to integrate with an underlying test framework like Jest or vitest:

```typescript
import { algorandFixture } from '@algorandfoundation/algokit-utils/testing'
```

### Using with Jest

To integrate with [Jest](https://jestjs.io/) you need to pass the `fixture.beforeEach` method into Jest's `beforeEach` method and then within each test you can access `fixture.context` to access per-test isolated fixture values.

```typescript
import { describe, test, beforeEach } from '@jest/globals'
import { algorandFixture } from './testing'

describe('MY MODULE', () => {
  const fixture = algorandFixture()
  beforeEach(fixture.beforeEach, 10_000)

  test('MY TEST', async () => {
    const { algod, testAccount /* ... */ } = fixture.context

    // Test stuff!
  })
})
```

Occasionally there may be a delay when first running the fixture setup so we add a 10s timeout to avoid intermittent test failures (`10_000`).

### Using with vitest

To integrate with [vitest](https://vitest.dev/) you need to pass the `fixture.beforeEach` method into vitest's `beforeEach` method and then within each test you can access `fixture.context` to access per-test isolated fixture values.

```typescript
import { describe, test, beforeEach } from 'vitest'
import { algorandFixture } from './testing'

describe('MY MODULE', () => {
  const fixture = algorandFixture()
  beforeEach(fixture.beforeEach, 10_000)

  test('MY TEST', async () => {
    const { algod, testAccount /* ... */ } = fixture.context

    // Test stuff!
  })
})
```

### Fixture configuration

When calling `algorandFixture()` you can optionally pass in some fixture configuration, with any of these properties (all optional):

- `algod?: Algodv2` - An optional algod client, if not specified then it will create one against environment variables defined network (if present) or default LocalNet
- `indexer?: Indexer` - An optional indexer client, if not specified then it will create one against environment variables defined network (if present) or default LocalNet
- `kmd?: Kmd` - An optional kmd client, if not specified then it will create one against environment variables defined network (if present) or default LocalNet
- `testAccountFunding?: AlgoAmount` - The [amount](./amount.md) of funds to allocate to the default testing account, if not specified then it will get `10` ALGOs

### Using the fixture context

The `fixture.context` property is of type [`AlgorandTestAutomationContext`](../code/interfaces/types_testing.AlgorandTestAutomationContext.md) exposes the following properties from which you can pick which ones you want in a given test using an object [destructuring assignment](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Destructuring_assignment):

- `algod: Algodv2` - Proxy Algod client instance that will log sent transactions in `transactionLogger`
- `indexer: Indexer` - Indexer client instance
- `kmd: Kmd` - KMD client instance
- `transactionLogger: TransactionLogger` - Transaction logger that will log transaction IDs for all transactions issued by `algod`
- `testAccount: Account` - Funded test account that is ephemerally created for each test
- `generateAccount: (params: GetTestAccountParams) => Promise<Account>` - Generate and fund an additional ephemerally created account
- `waitForIndexer: () => Promise<void>` - Wait for the indexer to catch up with all transactions logged by transactionLogger
- `waitForIndexerTransaction: (transactionId: string) => Promise<TransactionLookupResult>` - Wait for the indexer to catch up with the given transaction ID

## Log capture fixture

If you want to capture log messages from AlgoKit that are issued within your test so that you can assert on them or parse them for debugging information etc. then you can use the log capture fixture.

```typescript
import { algoKitLogCaptureFixture } from '@algorandfoundation/algokit-utils/testing'
```

The log capture fixture works by setting the logger within the AlgoKit configuration to be a [`TestLogger`](../code/classes/testing.TestLogger.md) during the test run.

### Using with Jest

To integrate with [Jest](https://jestjs.io/) you need to pass the `fixture.beforeEach` method into Jest's `beforeEach` method and then within each test you can access `fixture.context` to access per-test isolated fixture values.

```typescript
import { describe, test, beforeEach, afterEach } from '@jest/globals'
import { algoKitLogCaptureFixture } from './testing'

describe('MY MODULE', () => {
  const logs = algoKitLogCaptureFixture()
  beforeEach(logs.beforeEach)
  afterEach(logs.afterEach)

  test('MY TEST', async () => {
    // Test stuff!

    const capturedLogs = logs.testLogger.capturedLogs
    // do stuff with the logs
  })
})
```

### Using with vitest

To integrate with [vitest](https://vitest.dev/) you need to pass the `fixture.beforeEach` method into vitest's `beforeEach` method and then within each test you can access `fixture.context` to access per-test isolated fixture values.

```typescript
import { describe, test, beforeEach, afterEach } from 'vitest'
import { algoKitLogCaptureFixture } from './testing'

describe('MY MODULE', () => {
  const logs = algoKitLogCaptureFixture()
  beforeEach(logs.beforeEach)
  afterEach(logs.afterEach)

  test('MY TEST', async () => {
    // Test stuff!

    const capturedLogs = logs.testLogger.capturedLogs
    // do stuff with the logs
  })
})
```

### Snapshot testing the logs

If you want to quickly pin some behaviour of what logic you have does in terms of invoking AlgoKit methods you can do a [snapshot test](https://jestjs.io/docs/snapshot-testing) / [approval test](https://approvaltests.com/) of the captured log output. The only problem is this output will contain identifiers that will change for every test run and thus will constantly break the snapshot. In order to work around this you can use the `getLogSnapshot` method on the `TestLogger`, which will replace those changing values with predictable strings to keep the snapshot integrity intact.

This might look something like this:

```typescript
const { algod, indexer, testAccount } = fixture.context
const result = await algokit.deployApp(getAppDeploymentParams(), algod, indexer)
expect(
  logging.testLogger.getLogSnapshot({
    accounts: [testAccount],
    transactions: [result.transaction],
    apps: [result.appId],
  }),
).toMatchSnapshot()
```

## Waiting for indexer

Often there will be things that you do in your test that you may want to assert in using data that is exclusively in indexer such as transaction notes. The problem is indexer asynchronously indexes the data in algod, even when devmode is turned on and algod instantly confirms transactions.

This means it's easy to create tests that are flaky and have intermittent test failures (sometimes indexer is up to date and other times it hasn't caught up yet).

The testing capability provides mechanisms for waiting for indexer to catch up, namely:

- `algotesting.runWhenIndexerCaughtUp(run: () => Promise<T>)` - Executes the given action every 200ms up to 20 times until there is no longer an error with a `status` property with `404` and then returns the result of the action; this will work for any call that calls indexer APIs expecting to return a single record
- `algorandFixture.waitForIndexer()` - Waits for indexer to catch up with all transactions that have been captured by the `transactionLogger` in the Algorand fixture
- `algorandFixture.waitForIndexerTransaction(transactionId)` - Waits for indexer to catch up with the single transaction of the given ID

## Logging transactions

When testing, it can be useful to capture all of the transactions that have been issued with a given test run. They can then be asserted on, or used for [waiting for indexer](#waiting-for-indexer), etc.

The testing capability provides the ability to capture transactions via the [`TransactionLogger`](../code/classes/testing.TransactionLogger.md) class.

The `TransactionLogger` has the following methods:

- `logRawTransaction(signedTransactions: Uint8Array | Uint8Array[])` - Logs the IDs of the given signed transaction(s)
- `capture(algod)` - Returns a proxy `algosdk.Algodv2` instance that wraps the given `algod` client that will call `logRawTransaction` for every call to `sendRawTransaction` on that algod instance
- `sentTransactionIds` - Returns the currently captured list of transaction IDs that have been logged
- `clear()` - Clears the current list of transaction IDs
- `waitForIndexer(indexer)` - [Waits for the given indexer instance to catch up](#waiting-for-indexer) with the currently logged transaction IDs

The easiest way to use this functionality is via the [Algorand fixture](#algorand-fixture), which automatically provides a `transactionLogger` and a proxy `algod` connected to that `transactionLogger`.

## Getting a test account

When testing, it's often useful to ephemerally generate random accounts, fund them with some number of ALGOs and then use that account to perform transactions. By creating an ephemeral, random account you naturally get isolation between tests and test runs and don't need to start from a specific blockchain network state. This makes test less flakey, and also means the same test can be run against LocalNet and (say) TestNet.

The key when generating a test account is getting hold of a [dispenser](./transfer.md#dispenser) and then [ensuring the test account is funded](./transfer.md#ensurefunded).

To make it easier to quickly get a test account the testing capability provides the following mechanisms:

- [`algotesting.getTestAccount(testAccountParams, algod, kmd?)`](../code/modules/testing.md#gettestaccount) - Generates a random new account, logs the mnemonic of the account (unless suppressed), funds it from the [dispenser](./transfer.md#dispenser)
- `algorandFixture.testAccount` - A test account that is always generated for every test (log output suppressed to reduce noise, but worth noting that means the mnemonic isn't logged for this account), by default it is given 10 ALGOs unless overridden in the [fixture config](#fixture-configuration)
- [`algorandFixture.generateAccount(testAccountParams)`](../code/interfaces/types_testing.AlgorandTestAutomationContext.md#generateaccount) - Allows you to quickly generate a test account with the `algod` and `kmd` instances that are part of the given fixture

The parameters object that controls test account generation, [`GetTestAccountParams`](../code/interfaces/types_testing.GetTestAccountParams.md), has the following properties:

- `initialFunds: AlgoAmount` - Initial funds to ensure the account has
- `suppressLog?: boolean` - Whether to suppress the log (which includes a mnemonic) or not (default: do not suppress the log)
