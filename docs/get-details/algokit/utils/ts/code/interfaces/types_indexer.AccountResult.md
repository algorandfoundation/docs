[@algorandfoundation/algokit-utils](../index.md) / [types/indexer](../modules/types_indexer.md) / AccountResult

# Interface: AccountResult

[types/indexer](../modules/types_indexer.md).AccountResult

Account information at a given round https://developer.algorand.org/docs/rest-apis/indexer/#account

## Hierarchy

- `Record`\<`string`, `any`\>

  ↳ **`AccountResult`**

## Table of contents

### Properties

- [address](types_indexer.AccountResult.md#address)
- [amount](types_indexer.AccountResult.md#amount)
- [amount-without-pending-rewards](types_indexer.AccountResult.md#amount-without-pending-rewards)
- [apps-local-state](types_indexer.AccountResult.md#apps-local-state)
- [apps-total-extra-pages](types_indexer.AccountResult.md#apps-total-extra-pages)
- [apps-total-schema](types_indexer.AccountResult.md#apps-total-schema)
- [assets](types_indexer.AccountResult.md#assets)
- [auth-addr](types_indexer.AccountResult.md#auth-addr)
- [closed-at-round](types_indexer.AccountResult.md#closed-at-round)
- [created-apps](types_indexer.AccountResult.md#created-apps)
- [created-assets](types_indexer.AccountResult.md#created-assets)
- [created-at-round](types_indexer.AccountResult.md#created-at-round)
- [deleted](types_indexer.AccountResult.md#deleted)
- [participation](types_indexer.AccountResult.md#participation)
- [pending-rewards](types_indexer.AccountResult.md#pending-rewards)
- [reward-base](types_indexer.AccountResult.md#reward-base)
- [rewards](types_indexer.AccountResult.md#rewards)
- [round](types_indexer.AccountResult.md#round)
- [sig-type](types_indexer.AccountResult.md#sig-type)
- [status](types_indexer.AccountResult.md#status)
- [total-apps-opted-in](types_indexer.AccountResult.md#total-apps-opted-in)
- [total-assets-opted-in](types_indexer.AccountResult.md#total-assets-opted-in)
- [total-box-bytes](types_indexer.AccountResult.md#total-box-bytes)
- [total-boxes](types_indexer.AccountResult.md#total-boxes)
- [total-created-apps](types_indexer.AccountResult.md#total-created-apps)
- [total-created-assets](types_indexer.AccountResult.md#total-created-assets)

## Properties

### address

• **address**: `string`

the account public key

#### Defined in

[src/types/indexer.ts:192](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L192)

___

### amount

• **amount**: `number`

[algo] total number of MicroAlgos in the account

#### Defined in

[src/types/indexer.ts:194](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L194)

___

### amount-without-pending-rewards

• **amount-without-pending-rewards**: `number`

specifies the amount of MicroAlgos in the account, without the pending rewards.

#### Defined in

[src/types/indexer.ts:196](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L196)

___

### apps-local-state

• `Optional` **apps-local-state**: [`AppLocalState`](types_indexer.AppLocalState.md)[]

[appl] applications local data stored in this account.

Note the raw object uses map[int] -> AppLocalState for this type.

#### Defined in

[src/types/indexer.ts:201](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L201)

___

### apps-total-extra-pages

• `Optional` **apps-total-extra-pages**: `number`

[teap] the sum of all extra application program pages for this account.

#### Defined in

[src/types/indexer.ts:203](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L203)

___

### apps-total-schema

• `Optional` **apps-total-schema**: [`StateSchema`](types_indexer.StateSchema.md)

[tsch] stores the sum of all of the local schemas and global schemas in this account.

Note: the raw account uses StateSchema for this type.

#### Defined in

[src/types/indexer.ts:208](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L208)

___

### assets

• `Optional` **assets**: [`AssetHolding`](types_indexer.AssetHolding.md)[]

[asset] assets held by this account.

Note the raw object uses map[int] -> AssetHolding for this type.

#### Defined in

[src/types/indexer.ts:213](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L213)

___

### auth-addr

• `Optional` **auth-addr**: `string`

[spend] the address against which signing should be checked.

If empty, the address of the current account is used.

This field can be updated in any transaction by setting the RekeyTo field.

#### Defined in

[src/types/indexer.ts:220](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L220)

___

### closed-at-round

• `Optional` **closed-at-round**: `number`

Round during which this account was most recently closed.

#### Defined in

[src/types/indexer.ts:222](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L222)

___

### created-apps

• `Optional` **created-apps**: [`ApplicationResult`](types_indexer.ApplicationResult.md)[]

[appp] parameters of applications created by this account including app global data.

Note: the raw account uses map[int] -> AppParams for this type.

#### Defined in

[src/types/indexer.ts:227](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L227)

___

### created-assets

• `Optional` **created-assets**: [`AssetResult`](types_indexer.AssetResult.md)[]

[apar] parameters of assets created by this account.

Note: the raw account uses map[int] -> Asset for this type.

#### Defined in

[src/types/indexer.ts:232](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L232)

___

### created-at-round

• `Optional` **created-at-round**: `number`

Round during which this account first appeared in a transaction.

#### Defined in

[src/types/indexer.ts:234](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L234)

___

### deleted

• `Optional` **deleted**: `boolean`

Whether or not this account is currently closed.

#### Defined in

[src/types/indexer.ts:236](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L236)

___

### participation

• `Optional` **participation**: [`AccountParticipation`](types_indexer.AccountParticipation.md)

If participating in consensus, the parameters used by this account in the consensus protocol.

#### Defined in

[src/types/indexer.ts:238](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L238)

___

### pending-rewards

• **pending-rewards**: `number`

amount of MicroAlgos of pending rewards in this account.

#### Defined in

[src/types/indexer.ts:240](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L240)

___

### reward-base

• `Optional` **reward-base**: `number`

[ebase] used as part of the rewards computation. Only applicable to accounts which are participating.

#### Defined in

[src/types/indexer.ts:242](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L242)

___

### rewards

• **rewards**: `number`

[ern] total rewards of MicroAlgos the account has received, including pending rewards.

#### Defined in

[src/types/indexer.ts:244](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L244)

___

### round

• **round**: `number`

The round for which this information is relevant.

#### Defined in

[src/types/indexer.ts:246](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L246)

___

### sig-type

• **sig-type**: [`SignatureType`](../enums/types_indexer.SignatureType.md)

Indicates what type of signature is used by this account

#### Defined in

[src/types/indexer.ts:248](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L248)

___

### status

• **status**: [`AccountStatus`](../enums/types_indexer.AccountStatus.md)

[onl] delegation status of the account's MicroAlgos

#### Defined in

[src/types/indexer.ts:250](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L250)

___

### total-apps-opted-in

• **total-apps-opted-in**: `number`

The count of all applications that have been opted in, equivalent to the count of application local data (AppLocalState objects) stored in this account.

#### Defined in

[src/types/indexer.ts:252](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L252)

___

### total-assets-opted-in

• **total-assets-opted-in**: `number`

The count of all assets that have been opted in, equivalent to the count of AssetHolding objects held by this account.

#### Defined in

[src/types/indexer.ts:254](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L254)

___

### total-box-bytes

• **total-box-bytes**: `number`

For app-accounts only. The total number of bytes allocated for the keys and values of boxes which belong to the associated application.

#### Defined in

[src/types/indexer.ts:256](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L256)

___

### total-boxes

• **total-boxes**: `number`

For app-accounts only. The total number of boxes which belong to the associated application.

#### Defined in

[src/types/indexer.ts:258](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L258)

___

### total-created-apps

• **total-created-apps**: `number`

The count of all apps (AppParams objects) created by this account.

#### Defined in

[src/types/indexer.ts:260](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L260)

___

### total-created-assets

• **total-created-assets**: `number`

The count of all assets (AssetParams objects) created by this account.

#### Defined in

[src/types/indexer.ts:262](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L262)
