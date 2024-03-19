[@algorandfoundation/algokit-utils](../index.md) / [types/indexer](../modules/types_indexer.md) / TransactionResult

# Interface: TransactionResult

[types/indexer](../modules/types_indexer.md).TransactionResult

Indexer result for a transaction, https://developer.algorand.org/docs/rest-apis/indexer/#transaction

## Hierarchy

- `Record`\<`string`, `any`\>

  ↳ **`TransactionResult`**

## Table of contents

### Properties

- [application-transaction](types_indexer.TransactionResult.md#application-transaction)
- [asset-config-transaction](types_indexer.TransactionResult.md#asset-config-transaction)
- [asset-freeze-transaction](types_indexer.TransactionResult.md#asset-freeze-transaction)
- [asset-transfer-transaction](types_indexer.TransactionResult.md#asset-transfer-transaction)
- [auth-addr](types_indexer.TransactionResult.md#auth-addr)
- [close-rewards](types_indexer.TransactionResult.md#close-rewards)
- [closing-amount](types_indexer.TransactionResult.md#closing-amount)
- [confirmed-round](types_indexer.TransactionResult.md#confirmed-round)
- [created-application-index](types_indexer.TransactionResult.md#created-application-index)
- [created-asset-index](types_indexer.TransactionResult.md#created-asset-index)
- [fee](types_indexer.TransactionResult.md#fee)
- [first-valid](types_indexer.TransactionResult.md#first-valid)
- [genesis-hash](types_indexer.TransactionResult.md#genesis-hash)
- [genesis-id](types_indexer.TransactionResult.md#genesis-id)
- [global-state-delta](types_indexer.TransactionResult.md#global-state-delta)
- [group](types_indexer.TransactionResult.md#group)
- [id](types_indexer.TransactionResult.md#id)
- [inner-txns](types_indexer.TransactionResult.md#inner-txns)
- [intra-round-offset](types_indexer.TransactionResult.md#intra-round-offset)
- [keyreg-transaction](types_indexer.TransactionResult.md#keyreg-transaction)
- [last-valid](types_indexer.TransactionResult.md#last-valid)
- [lease](types_indexer.TransactionResult.md#lease)
- [local-state-delta](types_indexer.TransactionResult.md#local-state-delta)
- [logs](types_indexer.TransactionResult.md#logs)
- [note](types_indexer.TransactionResult.md#note)
- [payment-transaction](types_indexer.TransactionResult.md#payment-transaction)
- [receiver-rewards](types_indexer.TransactionResult.md#receiver-rewards)
- [rekey-to](types_indexer.TransactionResult.md#rekey-to)
- [round-time](types_indexer.TransactionResult.md#round-time)
- [sender](types_indexer.TransactionResult.md#sender)
- [sender-rewards](types_indexer.TransactionResult.md#sender-rewards)
- [signature](types_indexer.TransactionResult.md#signature)
- [state-proof-transaction](types_indexer.TransactionResult.md#state-proof-transaction)
- [tx-type](types_indexer.TransactionResult.md#tx-type)

## Properties

### application-transaction

• `Optional` **application-transaction**: [`ApplicationTransactionResult`](types_indexer.ApplicationTransactionResult.md)

If the transaction is an `appl` transaction this will be populated see `tx-type`

#### Defined in

[src/types/indexer.ts:126](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L126)

___

### asset-config-transaction

• `Optional` **asset-config-transaction**: [`AssetConfigTransactionResult`](types_indexer.AssetConfigTransactionResult.md)

If the transaction is an `acfg` transaction this will be populated see `tx-type`

#### Defined in

[src/types/indexer.ts:132](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L132)

___

### asset-freeze-transaction

• `Optional` **asset-freeze-transaction**: [`AssetFreezeTransactionResult`](types_indexer.AssetFreezeTransactionResult.md)

If the transaction is an `afrz` transaction this will be populated see `tx-type`

#### Defined in

[src/types/indexer.ts:138](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L138)

___

### asset-transfer-transaction

• `Optional` **asset-transfer-transaction**: [`AssetTransferTransactionResult`](types_indexer.AssetTransferTransactionResult.md)

If the transaction is an `axfer` transaction this will be populated see `tx-type`

#### Defined in

[src/types/indexer.ts:140](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L140)

___

### auth-addr

• `Optional` **auth-addr**: `string`

[sgnr] this is included with signed transactions when the signing address does not equal the sender.
The backend can use this to ensure that auth addr is equal to the accounts auth addr.

#### Defined in

[src/types/indexer.ts:150](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L150)

___

### close-rewards

• `Optional` **close-rewards**: `number`

[rc] rewards applied to close-remainder-to account.

#### Defined in

[src/types/indexer.ts:185](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L185)

___

### closing-amount

• `Optional` **closing-amount**: `number`

[ca] closing amount for transaction.

#### Defined in

[src/types/indexer.ts:152](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L152)

___

### confirmed-round

• `Optional` **confirmed-round**: `number`

Round when the transaction was confirmed.

#### Defined in

[src/types/indexer.ts:102](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L102)

___

### created-application-index

• `Optional` **created-application-index**: `number`

If the transaction is an `appl` transaction that resulted in an application creation then this
specifies the application index (ID) of that application.

#### Defined in

[src/types/indexer.ts:130](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L130)

___

### created-asset-index

• `Optional` **created-asset-index**: `number`

If the transaction is an `acfg` transaction that resulted in an asset creation then this
specifies the asset index (ID) of that asset.

#### Defined in

[src/types/indexer.ts:136](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L136)

___

### fee

• **fee**: `number`

[fee] Transaction fee.

#### Defined in

[src/types/indexer.ts:94](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L94)

___

### first-valid

• **first-valid**: `number`

[fv] First valid round for this transaction.

#### Defined in

[src/types/indexer.ts:98](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L98)

___

### genesis-hash

• `Optional` **genesis-hash**: `string`

[gh] Hash of genesis block.

*Pattern:* `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`

#### Defined in

[src/types/indexer.ts:157](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L157)

___

### genesis-id

• `Optional` **genesis-id**: `string`

[gen] genesis block ID.

#### Defined in

[src/types/indexer.ts:159](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L159)

___

### global-state-delta

• `Optional` **global-state-delta**: `Record`\<`string`, [`EvalDelta`](types_indexer.EvalDelta.md)\>[]

[gd] Global state key/value changes for the application being executed by this transaction.

#### Defined in

[src/types/indexer.ts:179](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L179)

___

### group

• `Optional` **group**: `string`

[grp] Base64 encoded byte array of a sha512/256 digest.

When present indicates that this transaction is part of a transaction group
 and the value is the sha512/256 hash of the transactions in that group.

*Pattern:* `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`

#### Defined in

[src/types/indexer.ts:110](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L110)

___

### id

• **id**: `string`

Transaction ID

#### Defined in

[src/types/indexer.ts:81](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L81)

___

### inner-txns

• `Optional` **inner-txns**: [`TransactionResult`](types_indexer.TransactionResult.md)[]

Inner transactions produced by application execution.

#### Defined in

[src/types/indexer.ts:161](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L161)

___

### intra-round-offset

• `Optional` **intra-round-offset**: `number`

Offset into the round where this transaction was confirmed.

#### Defined in

[src/types/indexer.ts:122](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L122)

___

### keyreg-transaction

• `Optional` **keyreg-transaction**: [`KeyRegistrationTransactionResult`](types_indexer.KeyRegistrationTransactionResult.md)

If the transaction is a `keyreg` transaction this will be populated see `tx-type`

#### Defined in

[src/types/indexer.ts:142](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L142)

___

### last-valid

• **last-valid**: `number`

[lv] Last valid round for this transaction.

#### Defined in

[src/types/indexer.ts:100](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L100)

___

### lease

• `Optional` **lease**: `string`

[lx] Base64 encoded 32-byte array. Lease enforces mutual exclusion of transactions.

If this field is nonzero, then once the transaction is confirmed, it acquires the lease
identified by the (Sender, Lease) pair of the transaction until the LastValid round passes.

While this transaction possesses the lease, no other transaction specifying this lease can be confirmed.

*Pattern:* `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`

#### Defined in

[src/types/indexer.ts:175](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L175)

___

### local-state-delta

• `Optional` **local-state-delta**: `Record`\<`string`, [`EvalDelta`](types_indexer.EvalDelta.md)\>[]

[ld] Local state key/value changes for the application being executed by this transaction.

#### Defined in

[src/types/indexer.ts:177](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L177)

___

### logs

• `Optional` **logs**: `string`[]

[lg] Logs for the application being executed by this transaction.

#### Defined in

[src/types/indexer.ts:118](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L118)

___

### note

• `Optional` **note**: `string`

[note] Free form data.

*Pattern:* `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`

#### Defined in

[src/types/indexer.ts:116](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L116)

___

### payment-transaction

• `Optional` **payment-transaction**: [`PaymentTransactionResult`](types_indexer.PaymentTransactionResult.md)

If the transaction is a `pay` transaction this will be populated see `tx-type`

#### Defined in

[src/types/indexer.ts:144](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L144)

___

### receiver-rewards

• `Optional` **receiver-rewards**: `number`

[rr] rewards applied to receiver account.

#### Defined in

[src/types/indexer.ts:181](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L181)

___

### rekey-to

• `Optional` **rekey-to**: `string`

[rekey] when included in a valid transaction, the accounts auth addr will be updated with
this value and future signatures must be signed with the key represented by this address.

#### Defined in

[src/types/indexer.ts:165](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L165)

___

### round-time

• `Optional` **round-time**: `number`

Time when the block this transaction is in was confirmed.

#### Defined in

[src/types/indexer.ts:120](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L120)

___

### sender

• **sender**: `string`

[snd] Sender's address.

#### Defined in

[src/types/indexer.ts:96](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L96)

___

### sender-rewards

• `Optional` **sender-rewards**: `number`

[rs] rewards applied to sender account.

#### Defined in

[src/types/indexer.ts:183](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L183)

___

### signature

• `Optional` **signature**: [`TransactionSignature`](types_indexer.TransactionSignature.md)

Signature of the transaction

#### Defined in

[src/types/indexer.ts:124](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L124)

___

### state-proof-transaction

• `Optional` **state-proof-transaction**: [`StateProofTransactionResult`](types_indexer.StateProofTransactionResult.md)

If the transaction is a `stpf` transaction this will be populated see `tx-type`

#### Defined in

[src/types/indexer.ts:146](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L146)

___

### tx-type

• **tx-type**: `TransactionType`

[type] Indicates what type of transaction this is. Different types have different fields.
Valid types, and where their fields are stored:
 * [pay] payment-transaction
 * [keyreg] keyreg-transaction
 * [acfg] asset-config-transaction
 * [axfer] asset-transfer-transaction
 * [afrz] asset-freeze-transaction
 * [appl] application-transaction
 * [stpf] state-proof-transaction

#### Defined in

[src/types/indexer.ts:92](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L92)
