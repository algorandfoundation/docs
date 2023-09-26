[@algorandfoundation/algokit-utils](../README.md) / [types/indexer](../modules/types_indexer.md) / ApplicationTransactionResult

# Interface: ApplicationTransactionResult

[types/indexer](../modules/types_indexer.md).ApplicationTransactionResult

Fields for an application transaction https://developer.algorand.org/docs/rest-apis/indexer/#transactionapplication

## Hierarchy

- `Omit`<[`ApplicationParams`](types_indexer.ApplicationParams.md), ``"creator"`` \| ``"global-state"``\>

  ↳ **`ApplicationTransactionResult`**

## Table of contents

### Properties

- [accounts](types_indexer.ApplicationTransactionResult.md#accounts)
- [application-args](types_indexer.ApplicationTransactionResult.md#application-args)
- [application-id](types_indexer.ApplicationTransactionResult.md#application-id)
- [approval-program](types_indexer.ApplicationTransactionResult.md#approval-program)
- [clear-state-program](types_indexer.ApplicationTransactionResult.md#clear-state-program)
- [extra-program-pages](types_indexer.ApplicationTransactionResult.md#extra-program-pages)
- [foreign-apps](types_indexer.ApplicationTransactionResult.md#foreign-apps)
- [foreign-assets](types_indexer.ApplicationTransactionResult.md#foreign-assets)
- [global-state-schema](types_indexer.ApplicationTransactionResult.md#global-state-schema)
- [local-state-schema](types_indexer.ApplicationTransactionResult.md#local-state-schema)
- [on-completion](types_indexer.ApplicationTransactionResult.md#on-completion)

## Properties

### accounts

• `Optional` **accounts**: `string`[]

[apat] List of accounts in addition to the sender that may be accessed from the application's approval-program and clear-state-program.

#### Defined in

[src/types/indexer.ts:267](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L267)

___

### application-args

• `Optional` **application-args**: `string`[]

[apaa] transaction specific arguments accessed from the application's approval-program and clear-state-program.

#### Defined in

[src/types/indexer.ts:269](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L269)

___

### application-id

• **application-id**: `number`

[apid] ID of the application being configured or empty if creating.

#### Defined in

[src/types/indexer.ts:271](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L271)

___

### approval-program

• **approval-program**: `string`

[apap]/[approv] Logic executed for every application transaction, except when on-completion is set to "clear".

It can read and write global state for the application, as well as account-specific local state.

Approval programs may reject the transaction.

*Pattern:* `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`

#### Inherited from

Omit.approval-program

#### Defined in

[src/types/indexer.ts:455](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L455)

___

### clear-state-program

• **clear-state-program**: `string`

[apsu]/[clearp] Logic executed for application transactions with on-completion set to "clear".

It can read and write global state for the application, as well as account-specific local state.

Clear state programs cannot reject the transaction.

*Pattern:* `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`

#### Inherited from

Omit.clear-state-program

#### Defined in

[src/types/indexer.ts:465](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L465)

___

### extra-program-pages

• `Optional` **extra-program-pages**: `number`

[epp] the amount of extra program pages available to this app.

#### Inherited from

Omit.extra-program-pages

#### Defined in

[src/types/indexer.ts:467](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L467)

___

### foreign-apps

• `Optional` **foreign-apps**: `number`[]

[apfa] Lists the applications in addition to the application-id whose global states may be accessed by this application's approval-program and clear-state-program. The access is read-only.

#### Defined in

[src/types/indexer.ts:273](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L273)

___

### foreign-assets

• `Optional` **foreign-assets**: `number`[]

[apas] lists the assets whose parameters may be accessed by this application's ApprovalProgram and ClearStateProgram. The access is read-only.

#### Defined in

[src/types/indexer.ts:275](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L275)

___

### global-state-schema

• `Optional` **global-state-schema**: [`StateSchema`](types_indexer.StateSchema.md)

[gsch] global schema

#### Inherited from

Omit.global-state-schema

#### Defined in

[src/types/indexer.ts:471](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L471)

___

### local-state-schema

• `Optional` **local-state-schema**: [`StateSchema`](types_indexer.StateSchema.md)

[lsch] local schema

#### Inherited from

Omit.local-state-schema

#### Defined in

[src/types/indexer.ts:473](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L473)

___

### on-completion

• **on-completion**: [`ApplicationOnComplete`](../enums/types_indexer.ApplicationOnComplete.md)

[apan] defines the what additional actions occur with the transaction.

#### Defined in

[src/types/indexer.ts:277](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L277)
