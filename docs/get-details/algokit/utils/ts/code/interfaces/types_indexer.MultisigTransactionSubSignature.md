[@algorandfoundation/algokit-utils](../index.md) / [types/indexer](../modules/types_indexer.md) / MultisigTransactionSubSignature

# Interface: MultisigTransactionSubSignature

[types/indexer](../modules/types_indexer.md).MultisigTransactionSubSignature

Sub-signature for a multisig signature https://developer.algorand.org/docs/rest-apis/indexer/#transactionsignaturemultisigsubsignature

## Table of contents

### Properties

- [public-key](types_indexer.MultisigTransactionSubSignature.md#public-key)
- [signature](types_indexer.MultisigTransactionSubSignature.md#signature)

## Properties

### public-key

• **public-key**: `string`

[pk] The public key of the account making the signature

*Pattern:* `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`

#### Defined in

[src/types/indexer.ts:606](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L606)

___

### signature

• `Optional` **signature**: `string`

[s] The signature

*Pattern:* `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`

#### Defined in

[src/types/indexer.ts:611](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L611)
