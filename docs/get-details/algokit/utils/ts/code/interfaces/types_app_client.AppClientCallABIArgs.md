[@algorandfoundation/algokit-utils](../index.md) / [types/app-client](../modules/types_app_client.md) / AppClientCallABIArgs

# Interface: AppClientCallABIArgs

[types/app-client](../modules/types_app_client.md).AppClientCallABIArgs

## Hierarchy

- `Omit`\<[`ABIAppCallArgs`](../modules/types_app.md#abiappcallargs), ``"method"``\>

  ↳ **`AppClientCallABIArgs`**

## Table of contents

### Properties

- [accounts](types_app_client.AppClientCallABIArgs.md#accounts)
- [apps](types_app_client.AppClientCallABIArgs.md#apps)
- [assets](types_app_client.AppClientCallABIArgs.md#assets)
- [boxes](types_app_client.AppClientCallABIArgs.md#boxes)
- [lease](types_app_client.AppClientCallABIArgs.md#lease)
- [method](types_app_client.AppClientCallABIArgs.md#method)
- [methodArgs](types_app_client.AppClientCallABIArgs.md#methodargs)
- [rekeyTo](types_app_client.AppClientCallABIArgs.md#rekeyto)

## Properties

### accounts

• `Optional` **accounts**: (`string` \| `Address`)[]

The address of any accounts to load in

#### Inherited from

Omit.accounts

#### Defined in

[src/types/app.ts:73](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L73)

___

### apps

• `Optional` **apps**: `number`[]

IDs of any apps to load into the foreignApps array

#### Inherited from

Omit.apps

#### Defined in

[src/types/app.ts:75](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L75)

___

### assets

• `Optional` **assets**: `number`[]

IDs of any assets to load into the foreignAssets array

#### Inherited from

Omit.assets

#### Defined in

[src/types/app.ts:77](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L77)

___

### boxes

• `Optional` **boxes**: ([`BoxReference`](types_app.BoxReference.md) \| [`BoxIdentifier`](../modules/types_app.md#boxidentifier) \| `BoxReference`)[]

Any box references to load

#### Inherited from

Omit.boxes

#### Defined in

[src/types/app.ts:71](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L71)

___

### lease

• `Optional` **lease**: `string` \| `Uint8Array`

The optional lease for the transaction

#### Inherited from

Omit.lease

#### Defined in

[src/types/app.ts:69](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L69)

___

### method

• **method**: `string`

If calling an ABI method then either the name of the method, or the ABI signature

#### Defined in

[src/types/app-client.ts:166](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L166)

___

### methodArgs

• **methodArgs**: [`ABIAppCallArg`](../modules/types_app.md#abiappcallarg)[]

The ABI method args to pass in

#### Inherited from

Omit.methodArgs

#### Defined in

[src/types/app.ts:111](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L111)

___

### rekeyTo

• `Optional` **rekeyTo**: `string` \| [`SendTransactionFrom`](../modules/types_transaction.md#sendtransactionfrom)

Optional account / account address that should be authorised to transact on behalf of the from account the app call is sent from after this transaction.

**Note:** Use with extreme caution and review the [official rekey guidance](https://developer.algorand.org/docs/get-details/accounts/rekey/) first.

#### Inherited from

Omit.rekeyTo

#### Defined in

[src/types/app.ts:82](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app.ts#L82)
