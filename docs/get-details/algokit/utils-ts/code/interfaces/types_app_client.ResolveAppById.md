[@algorandfoundation/algokit-utils](../README.md) / [types/app-client](../modules/types_app_client.md) / ResolveAppById

# Interface: ResolveAppById

[types/app-client](../modules/types_app_client.md).ResolveAppById

Configuration to resolve app by ID

## Table of contents

### Properties

- [id](types_app_client.ResolveAppById.md#id)
- [name](types_app_client.ResolveAppById.md#name)
- [resolveBy](types_app_client.ResolveAppById.md#resolveby)

## Properties

### id

• **id**: `number` \| `bigint`

The id of an existing app to call using this client, or 0 if the app hasn't been created yet

#### Defined in

[src/types/app-client.ts:76](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L76)

___

### name

• `Optional` **name**: `string`

The optional name to use to mark the app when deploying `ApplicationClient.deploy` (default: uses the name in the ABI contract)

#### Defined in

[src/types/app-client.ts:78](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L78)

___

### resolveBy

• **resolveBy**: ``"id"``

How the app ID is resolved, either by `'id'` or `'creatorAndName'`; must be `'creatorAndName'` if you want to use `deploy`

#### Defined in

[src/types/app-client.ts:74](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L74)
