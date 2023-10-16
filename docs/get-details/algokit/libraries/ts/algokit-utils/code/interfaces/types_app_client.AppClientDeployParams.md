[@algorandfoundation/algokit-utils](../README.md) / [types/app-client](../modules/types_app_client.md) / AppClientDeployParams

# Interface: AppClientDeployParams

[types/app-client](../modules/types_app_client.md).AppClientDeployParams

Parameters to pass into ApplicationClient.deploy

## Hierarchy

- [`AppClientDeployCoreParams`](types_app_client.AppClientDeployCoreParams.md)

- [`AppClientDeployCallInterfaceParams`](types_app_client.AppClientDeployCallInterfaceParams.md)

  ↳ **`AppClientDeployParams`**

## Table of contents

### Properties

- [allowDelete](types_app_client.AppClientDeployParams.md#allowdelete)
- [allowUpdate](types_app_client.AppClientDeployParams.md#allowupdate)
- [createArgs](types_app_client.AppClientDeployParams.md#createargs)
- [createOnCompleteAction](types_app_client.AppClientDeployParams.md#createoncompleteaction)
- [deleteArgs](types_app_client.AppClientDeployParams.md#deleteargs)
- [deployTimeParams](types_app_client.AppClientDeployParams.md#deploytimeparams)
- [onSchemaBreak](types_app_client.AppClientDeployParams.md#onschemabreak)
- [onUpdate](types_app_client.AppClientDeployParams.md#onupdate)
- [sendParams](types_app_client.AppClientDeployParams.md#sendparams)
- [sender](types_app_client.AppClientDeployParams.md#sender)
- [updateArgs](types_app_client.AppClientDeployParams.md#updateargs)
- [version](types_app_client.AppClientDeployParams.md#version)

## Properties

### allowDelete

• `Optional` **allowDelete**: `boolean`

Whether or not to allow deletes in the contract using the deploy-time deletability control if present in your contract.
If this is not specified then it will automatically be determined based on the AppSpec definition

#### Inherited from

[AppClientDeployCoreParams](types_app_client.AppClientDeployCoreParams.md).[allowDelete](types_app_client.AppClientDeployCoreParams.md#allowdelete)

#### Defined in

[src/types/app-client.ts:135](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L135)

___

### allowUpdate

• `Optional` **allowUpdate**: `boolean`

Whether or not to allow updates in the contract using the deploy-time updatability control if present in your contract.
If this is not specified then it will automatically be determined based on the AppSpec definition

#### Inherited from

[AppClientDeployCoreParams](types_app_client.AppClientDeployCoreParams.md).[allowUpdate](types_app_client.AppClientDeployCoreParams.md#allowupdate)

#### Defined in

[src/types/app-client.ts:131](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L131)

___

### createArgs

• `Optional` **createArgs**: [`AppClientCallArgs`](../modules/types_app_client.md#appclientcallargs)

Any args to pass to any create transaction that is issued as part of deployment

#### Inherited from

[AppClientDeployCallInterfaceParams](types_app_client.AppClientDeployCallInterfaceParams.md).[createArgs](types_app_client.AppClientDeployCallInterfaceParams.md#createargs)

#### Defined in

[src/types/app-client.ts:147](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L147)

___

### createOnCompleteAction

• `Optional` **createOnCompleteAction**: ``"no_op"`` \| ``"opt_in"`` \| ``"close_out"`` \| ``"update_application"`` \| ``"delete_application"`` \| `NoOpOC` \| `OptInOC` \| `CloseOutOC` \| `UpdateApplicationOC` \| `DeleteApplicationOC`

Override the on-completion action for the create call; defaults to NoOp

#### Inherited from

[AppClientDeployCallInterfaceParams](types_app_client.AppClientDeployCallInterfaceParams.md).[createOnCompleteAction](types_app_client.AppClientDeployCallInterfaceParams.md#createoncompleteaction)

#### Defined in

[src/types/app-client.ts:149](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L149)

___

### deleteArgs

• `Optional` **deleteArgs**: [`AppClientCallArgs`](../modules/types_app_client.md#appclientcallargs)

Any args to pass to any delete transaction that is issued as part of deployment

#### Inherited from

[AppClientDeployCallInterfaceParams](types_app_client.AppClientDeployCallInterfaceParams.md).[deleteArgs](types_app_client.AppClientDeployCallInterfaceParams.md#deleteargs)

#### Defined in

[src/types/app-client.ts:153](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L153)

___

### deployTimeParams

• `Optional` **deployTimeParams**: [`TealTemplateParams`](types_app.TealTemplateParams.md)

Any deploy-time parameters to replace in the TEAL code

#### Inherited from

[AppClientDeployCallInterfaceParams](types_app_client.AppClientDeployCallInterfaceParams.md).[deployTimeParams](types_app_client.AppClientDeployCallInterfaceParams.md#deploytimeparams)

#### Defined in

[src/types/app-client.ts:145](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L145)

___

### onSchemaBreak

• `Optional` **onSchemaBreak**: [`OnSchemaBreak`](../enums/types_app.OnSchemaBreak.md) \| ``"replace"`` \| ``"fail"`` \| ``"append"``

What action to perform if a schema break is detected

#### Inherited from

[AppClientDeployCoreParams](types_app_client.AppClientDeployCoreParams.md).[onSchemaBreak](types_app_client.AppClientDeployCoreParams.md#onschemabreak)

#### Defined in

[src/types/app-client.ts:137](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L137)

___

### onUpdate

• `Optional` **onUpdate**: ``"replace"`` \| ``"fail"`` \| ``"append"`` \| [`OnUpdate`](../enums/types_app.OnUpdate.md) \| ``"update"``

What action to perform if a TEAL update is detected

#### Inherited from

[AppClientDeployCoreParams](types_app_client.AppClientDeployCoreParams.md).[onUpdate](types_app_client.AppClientDeployCoreParams.md#onupdate)

#### Defined in

[src/types/app-client.ts:139](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L139)

___

### sendParams

• `Optional` **sendParams**: `Omit`<[`SendTransactionParams`](types_transaction.SendTransactionParams.md), ``"skipSending"`` \| ``"skipWaiting"``\>

Parameters to control transaction sending

#### Inherited from

[AppClientDeployCoreParams](types_app_client.AppClientDeployCoreParams.md).[sendParams](types_app_client.AppClientDeployCoreParams.md#sendparams)

#### Defined in

[src/types/app-client.ts:127](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L127)

___

### sender

• `Optional` **sender**: [`SendTransactionFrom`](../modules/types_transaction.md#sendtransactionfrom)

The optional sender to send the transaction from, will use the application client's default sender by default if specified

#### Inherited from

[AppClientDeployCoreParams](types_app_client.AppClientDeployCoreParams.md).[sender](types_app_client.AppClientDeployCoreParams.md#sender)

#### Defined in

[src/types/app-client.ts:125](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L125)

___

### updateArgs

• `Optional` **updateArgs**: [`AppClientCallArgs`](../modules/types_app_client.md#appclientcallargs)

Any args to pass to any update transaction that is issued as part of deployment

#### Inherited from

[AppClientDeployCallInterfaceParams](types_app_client.AppClientDeployCallInterfaceParams.md).[updateArgs](types_app_client.AppClientDeployCallInterfaceParams.md#updateargs)

#### Defined in

[src/types/app-client.ts:151](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L151)

___

### version

• `Optional` **version**: `string`

The version of the contract, uses "1.0" by default

#### Inherited from

[AppClientDeployCoreParams](types_app_client.AppClientDeployCoreParams.md).[version](types_app_client.AppClientDeployCoreParams.md#version)

#### Defined in

[src/types/app-client.ts:123](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L123)
