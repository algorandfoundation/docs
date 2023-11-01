[@algorandfoundation/algokit-utils](../index.md) / [types/account](../modules/types_account.md) / MultisigAccount

# Class: MultisigAccount

[types/account](../modules/types_account.md).MultisigAccount

Account wrapper that supports partial or full multisig signing.

## Table of contents

### Constructors

- [constructor](types_account.MultisigAccount.md#constructor)

### Properties

- [\_addr](types_account.MultisigAccount.md#_addr)
- [\_params](types_account.MultisigAccount.md#_params)
- [\_signer](types_account.MultisigAccount.md#_signer)
- [\_signingAccounts](types_account.MultisigAccount.md#_signingaccounts)

### Accessors

- [addr](types_account.MultisigAccount.md#addr)
- [params](types_account.MultisigAccount.md#params)
- [signer](types_account.MultisigAccount.md#signer)
- [signingAccounts](types_account.MultisigAccount.md#signingaccounts)

### Methods

- [sign](types_account.MultisigAccount.md#sign)

## Constructors

### constructor

• **new MultisigAccount**(`multisigParams`, `signingAccounts`)

#### Parameters

| Name | Type |
| :------ | :------ |
| `multisigParams` | `MultisigMetadata` |
| `signingAccounts` | (`default` \| [`SigningAccount`](types_account.SigningAccount.md))[] |

#### Defined in

[src/types/account.ts:34](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/account.ts#L34)

## Properties

### \_addr

• **\_addr**: `string`

#### Defined in

[src/types/account.ts:12](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/account.ts#L12)

___

### \_params

• **\_params**: `MultisigMetadata`

#### Defined in

[src/types/account.ts:10](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/account.ts#L10)

___

### \_signer

• **\_signer**: `TransactionSigner`

#### Defined in

[src/types/account.ts:13](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/account.ts#L13)

___

### \_signingAccounts

• **\_signingAccounts**: (`default` \| [`SigningAccount`](types_account.SigningAccount.md))[]

#### Defined in

[src/types/account.ts:11](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/account.ts#L11)

## Accessors

### addr

• `get` **addr**(): `string`

The address of the multisig account

#### Returns

`string`

#### Defined in

[src/types/account.ts:26](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/account.ts#L26)

___

### params

• `get` **params**(): `Readonly`<`MultisigMetadata`\>

The parameters for the multisig account

#### Returns

`Readonly`<`MultisigMetadata`\>

#### Defined in

[src/types/account.ts:16](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/account.ts#L16)

___

### signer

• `get` **signer**(): `TransactionSigner`

#### Returns

`TransactionSigner`

#### Defined in

[src/types/account.ts:30](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/account.ts#L30)

___

### signingAccounts

• `get` **signingAccounts**(): readonly (`default` \| [`SigningAccount`](types_account.SigningAccount.md))[]

The list of accounts that are present to sign

#### Returns

readonly (`default` \| [`SigningAccount`](types_account.SigningAccount.md))[]

#### Defined in

[src/types/account.ts:21](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/account.ts#L21)

## Methods

### sign

▸ **sign**(`transaction`): `Uint8Array`

Sign the given transaction

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `transaction` | `Uint8Array` \| `Transaction` | Either a transaction object or a raw, partially signed transaction |

#### Returns

`Uint8Array`

The transaction signed by the present signers

#### Defined in

[src/types/account.ts:49](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/account.ts#L49)
