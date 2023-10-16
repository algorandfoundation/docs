[@algorandfoundation/algokit-utils](../README.md) / [types/indexer](../modules/types_indexer.md) / AccountParticipation

# Interface: AccountParticipation

[types/indexer](../modules/types_indexer.md).AccountParticipation

AccountParticipation describes the parameters used by this account in consensus protocol. https://developer.algorand.org/docs/rest-apis/indexer/#accountparticipation

## Table of contents

### Properties

- [selection-participation-key](types_indexer.AccountParticipation.md#selection-participation-key)
- [state-proof-key](types_indexer.AccountParticipation.md#state-proof-key)
- [vote-first-valid](types_indexer.AccountParticipation.md#vote-first-valid)
- [vote-key-dilution](types_indexer.AccountParticipation.md#vote-key-dilution)
- [vote-last-valid](types_indexer.AccountParticipation.md#vote-last-valid)
- [vote-participation-key](types_indexer.AccountParticipation.md#vote-participation-key)

## Properties

### selection-participation-key

• **selection-participation-key**: `string`

[sel] Selection public key (if any) currently registered for this round.

*Pattern:* `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`

#### Defined in

[src/types/indexer.ts:601](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L601)

___

### state-proof-key

• `Optional` **state-proof-key**: `string`

[stprf] Root of the state proof key (if any).

*Pattern:* `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`

#### Defined in

[src/types/indexer.ts:606](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L606)

___

### vote-first-valid

• **vote-first-valid**: `number`

[voteFst] First round for which this participation is valid.

#### Defined in

[src/types/indexer.ts:608](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L608)

___

### vote-key-dilution

• **vote-key-dilution**: `number`

[voteKD] Number of subkeys in each batch of participation keys.

#### Defined in

[src/types/indexer.ts:610](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L610)

___

### vote-last-valid

• **vote-last-valid**: `number`

[voteLst] Last round for which this participation is valid.

#### Defined in

[src/types/indexer.ts:612](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L612)

___

### vote-participation-key

• **vote-participation-key**: `string`

[vote] root participation public key (if any) currently registered for this round.

*Pattern:* `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`

#### Defined in

[src/types/indexer.ts:617](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts#L617)
