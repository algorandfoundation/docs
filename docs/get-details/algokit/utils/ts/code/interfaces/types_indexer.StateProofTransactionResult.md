[@algorandfoundation/algokit-utils](../index.md) / [types/indexer](../modules/types_indexer.md) / StateProofTransactionResult

# Interface: StateProofTransactionResult

[types/indexer](../modules/types_indexer.md).StateProofTransactionResult

Fields for a state proof transaction https://developer.algorand.org/docs/rest-apis/indexer/#transactionstateproof.

See also https://developer.algorand.org/docs/get-details/stateproofs/,
https://developer.algorand.org/docs/get-details/stateproofs/light_client/,
https://github.com/algorand/go-algorand/blob/master/data/transactions/stateproof.go,
https://github.com/algorand/go-algorand/blob/master/crypto/stateproof/structs.go,
https://github.com/algorand/go-algorand/blob/master/data/stateproofmsg/message.go, and
https://developer.algorand.org/docs/rest-apis/algod/#stateproof.

## Table of contents

### Properties

- [message](types_indexer.StateProofTransactionResult.md#message)
- [state-proof](types_indexer.StateProofTransactionResult.md#state-proof)
- [state-proof-type](types_indexer.StateProofTransactionResult.md#state-proof-type)

## Properties

### message

• **message**: `Object`

[spmsg] State proof message

Message represents the message that the state proofs are attesting to. This message can be
used by lightweight client and gives it the ability to verify proofs on the Algorand's state.

In addition to that proof, this message also contains fields that
are needed in order to verify the next state proofs (VotersCommitment and LnProvenWeight).

#### Type declaration

| Name | Type | Description |
| :------ | :------ | :------ |
| `block-headers-commitment` | `string` | [b] BlockHeadersCommitment contains a commitment on all light block headers within a state proof interval. |
| `first-attested-round` | `number` | [f] First round the message attests to |
| `latest-attested-round` | `number` | [l] Last round the message attests to |
| `ln-proven-weight` | `number` | [P] An integer value representing the natural log of the proven weight with 16 bits of precision. This value would be used to verify the next state proof. |
| `voters-commitment` | `string` | [v] The vector commitment root of the top N accounts to sign the next StateProof. Pattern : "^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\\|[A-Za-z0-9+/]{3}=)?$" |

#### Defined in

[src/types/indexer.ts:295](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L295)

___

### state-proof

• **state-proof**: `Object`

[sp] a proof on Algorand's state

#### Type declaration

| Name | Type | Description |
| :------ | :------ | :------ |
| `part-proofs` | [`MerkleArrayProof`](types_indexer.MerkleArrayProof.md) | [P] Part proofs that make up the overall proof |
| `positions-to-reveal` | `number`[] | [pr] The positions that are revealed |
| `reveals` | \{ `participant`: \{ `verifier`: \{ `commitment`: `string` ; `key-lifetime`: `number`  } ; `weight`: `number`  } ; `position`: `number` ; `sig-slot`: \{ `lower-sig-weight`: `number` ; `signature`: \{ `falcon-signature`: `string` ; `merkle-array-index`: `number` ; `proof`: [`MerkleArrayProof`](types_indexer.MerkleArrayProof.md) ; `verifying-key`: `string`  }  }  }[] | [r] Reveals is a sparse map from the position being revealed to the corresponding elements from the sigs and participants arrays. |
| `salt-version` | `number` | [v] Merkle signature salt version |
| `sig-commit` | `string` | [c] Digest of the signature commit |
| `sig-proofs` | [`MerkleArrayProof`](types_indexer.MerkleArrayProof.md) | [S] Proofs for the signature |
| `signed-weight` | `number` | [w] The combined weight of the signatures |

#### Defined in

[src/types/indexer.ts:310](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L310)

___

### state-proof-type

• **state-proof-type**: `number`

[sptype] State proof type, per https://github.com/algorand/go-algorand/blob/master/protocol/stateproof.go#L24

 * 0: StateProofBasic is our initial state proof setup. using falcon keys and subset-sum hash

#### Defined in

[src/types/indexer.ts:389](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L389)
