[@algorandfoundation/algokit-utils](../index.md) / [types/indexer](../modules/types_indexer.md) / ApplicationParams

# Interface: ApplicationParams

[types/indexer](../modules/types_indexer.md).ApplicationParams

Stores the global information associated with an application https://developer.algorand.org/docs/rest-apis/indexer/#applicationparams

## Table of contents

### Properties

- [approval-program](types_indexer.ApplicationParams.md#approval-program)
- [clear-state-program](types_indexer.ApplicationParams.md#clear-state-program)
- [creator](types_indexer.ApplicationParams.md#creator)
- [extra-program-pages](types_indexer.ApplicationParams.md#extra-program-pages)
- [global-state](types_indexer.ApplicationParams.md#global-state)
- [global-state-schema](types_indexer.ApplicationParams.md#global-state-schema)
- [local-state-schema](types_indexer.ApplicationParams.md#local-state-schema)

## Properties

### approval-program

• **approval-program**: `string`

[apap]/[approv] Logic executed for every application transaction, except when on-completion is set to "clear".

It can read and write global state for the application, as well as account-specific local state.

Approval programs may reject the transaction.

*Pattern:* `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`

#### Defined in

[src/types/indexer.ts:637](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L637)

___

### clear-state-program

• **clear-state-program**: `string`

[apsu]/[clearp] Logic executed for application transactions with on-completion set to "clear".

It can read and write global state for the application, as well as account-specific local state.

Clear state programs cannot reject the transaction.

*Pattern:* `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`

#### Defined in

[src/types/indexer.ts:647](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L647)

___

### creator

• **creator**: `string`

The address that created this application. This is the address where the parameters and global state for this application can be found.

#### Defined in

[src/types/indexer.ts:627](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L627)

___

### extra-program-pages

• `Optional` **extra-program-pages**: `number`

[epp] the amount of extra program pages available to this app.

#### Defined in

[src/types/indexer.ts:649](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L649)

___

### global-state

• **global-state**: `TealKeyValue`[]

[gs] global schema

#### Defined in

[src/types/indexer.ts:651](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L651)

___

### global-state-schema

• `Optional` **global-state-schema**: [`StateSchema`](types_indexer.StateSchema.md)

[gsch] global schema

#### Defined in

[src/types/indexer.ts:653](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L653)

___

### local-state-schema

• `Optional` **local-state-schema**: [`StateSchema`](types_indexer.StateSchema.md)

[lsch] local schema

#### Defined in

[src/types/indexer.ts:655](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L655)
