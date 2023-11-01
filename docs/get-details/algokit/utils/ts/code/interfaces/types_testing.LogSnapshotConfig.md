[@algorandfoundation/algokit-utils](../index.md) / [types/testing](../modules/types_testing.md) / LogSnapshotConfig

# Interface: LogSnapshotConfig

[types/testing](../modules/types_testing.md).LogSnapshotConfig

Configuration for preparing a captured log snapshot.
This helps ensure that the provided configuration items won't appear
 with random values in the log snapshot, but rather will get substituted with predictable ids.

## Table of contents

### Properties

- [accounts](types_testing.LogSnapshotConfig.md#accounts)
- [apps](types_testing.LogSnapshotConfig.md#apps)
- [transactions](types_testing.LogSnapshotConfig.md#transactions)

## Properties

### accounts

• `Optional` **accounts**: (`string` \| [`SendTransactionFrom`](../modules/types_transaction.md#sendtransactionfrom))[]

Any accounts/addresses to replace the address for predictably

#### Defined in

[src/types/testing.ts:79](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/testing.ts#L79)

___

### apps

• `Optional` **apps**: (`string` \| `number` \| `bigint`)[]

Any app IDs to replace predictably

#### Defined in

[src/types/testing.ts:81](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/testing.ts#L81)

___

### transactions

• `Optional` **transactions**: (`string` \| `Transaction`)[]

Any transaction IDs or transactions to replace the ID for predictably

#### Defined in

[src/types/testing.ts:77](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/testing.ts#L77)
