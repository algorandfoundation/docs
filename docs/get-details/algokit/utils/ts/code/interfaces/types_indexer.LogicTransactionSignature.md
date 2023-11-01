[@algorandfoundation/algokit-utils](../index.md) / [types/indexer](../modules/types_indexer.md) / LogicTransactionSignature

# Interface: LogicTransactionSignature

[types/indexer](../modules/types_indexer.md).LogicTransactionSignature

[lsig] Programatic transaction signature.

https://developer.algorand.org/docs/rest-apis/indexer/#transactionsignaturelogicsig

https://developer.algorand.org/docs/get-details/transactions/signatures/#logic-signatures

## Table of contents

### Properties

- [args](types_indexer.LogicTransactionSignature.md#args)
- [logic](types_indexer.LogicTransactionSignature.md#logic)
- [multisig-signature](types_indexer.LogicTransactionSignature.md#multisig-signature)
- [signature](types_indexer.LogicTransactionSignature.md#signature)

## Properties

### args

• `Optional` **args**: `string`[]

[arg] Logic arguments, base64 encoded.

#### Defined in

[src/types/indexer.ts:391](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L391)

___

### logic

• **logic**: `string`

[l] Program signed by a signature or multi signature, or hashed to be the address of ana ccount.

Base64 encoded TEAL program.

*Pattern:* `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`

#### Defined in

[src/types/indexer.ts:398](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L398)

___

### multisig-signature

• `Optional` **multisig-signature**: [`MultisigTransactionSignature`](types_indexer.MultisigTransactionSignature.md)

The signature of the multisig the logic signature delegating the logicsig. https://developer.algorand.org/docs/get-details/dapps/smart-contracts/smartsigs/modes/#delegated-approval

#### Defined in

[src/types/indexer.ts:400](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L400)

___

### signature

• `Optional` **signature**: `string`

[sig] Standard ed25519 signature delegating the logicsig. https://developer.algorand.org/docs/get-details/dapps/smart-contracts/smartsigs/modes/#delegated-approval

*Pattern:* `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`

#### Defined in

[src/types/indexer.ts:405](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L405)
