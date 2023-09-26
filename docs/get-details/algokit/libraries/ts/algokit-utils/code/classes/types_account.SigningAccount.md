[@algorandfoundation/algokit-utils](../README.md) / [types/account](../modules/types_account.md) / SigningAccount

# Class: SigningAccount

[types/account](../modules/types_account.md).SigningAccount

Account wrapper that supports a rekeyed account

## Implements

- `default`

## Table of contents

### Constructors

- [constructor](types_account.SigningAccount.md#constructor)

### Properties

- [\_account](types_account.SigningAccount.md#_account)
- [\_sender](types_account.SigningAccount.md#_sender)
- [\_signer](types_account.SigningAccount.md#_signer)

### Accessors

- [addr](types_account.SigningAccount.md#addr)
- [sender](types_account.SigningAccount.md#sender)
- [signer](types_account.SigningAccount.md#signer)
- [sk](types_account.SigningAccount.md#sk)

## Constructors

### constructor

• **new SigningAccount**(`account`, `sender`)

#### Parameters

| Name | Type |
| :------ | :------ |
| `account` | `default` |
| `sender` | `undefined` \| `string` |

#### Defined in

[src/types/account.ts:100](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/account.ts#L100)

## Properties

### \_account

• `Private` **\_account**: `default`

#### Defined in

[src/types/account.ts:65](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/account.ts#L65)

___

### \_sender

• `Private` **\_sender**: `string`

#### Defined in

[src/types/account.ts:67](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/account.ts#L67)

___

### \_signer

• `Private` **\_signer**: `TransactionSigner`

#### Defined in

[src/types/account.ts:66](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/account.ts#L66)

## Accessors

### addr

• `get` **addr**(): `string`

Algorand address of the sender

#### Returns

`string`

#### Implementation of

Account.addr

#### Defined in

[src/types/account.ts:72](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/account.ts#L72)

___

### sender

• `get` **sender**(): `default`

Algorand account of the sender address and signer private key

#### Returns

`default`

#### Defined in

[src/types/account.ts:93](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/account.ts#L93)

___

### signer

• `get` **signer**(): `TransactionSigner`

Transaction signer for the underlying signing account

#### Returns

`TransactionSigner`

#### Defined in

[src/types/account.ts:86](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/account.ts#L86)

___

### sk

• `get` **sk**(): `Readonly`<`Uint8Array`\>

Secret key belonging to the signer

#### Returns

`Readonly`<`Uint8Array`\>

#### Implementation of

Account.sk

#### Defined in

[src/types/account.ts:79](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/account.ts#L79)
