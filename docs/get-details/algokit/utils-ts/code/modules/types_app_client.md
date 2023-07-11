[@algorandfoundation/algokit-utils](../README.md) / types/app-client

# Module: types/app-client

## Table of contents

### Classes

- [ApplicationClient](../classes/types_app_client.ApplicationClient.md)

### Interfaces

- [AppClientCallABIArgs](../interfaces/types_app_client.AppClientCallABIArgs.md)
- [AppClientCallCoreParams](../interfaces/types_app_client.AppClientCallCoreParams.md)
- [AppClientCallRawArgs](../interfaces/types_app_client.AppClientCallRawArgs.md)
- [AppClientCompilationParams](../interfaces/types_app_client.AppClientCompilationParams.md)
- [AppClientDeployCallInterfaceParams](../interfaces/types_app_client.AppClientDeployCallInterfaceParams.md)
- [AppClientDeployCoreParams](../interfaces/types_app_client.AppClientDeployCoreParams.md)
- [AppClientDeployParams](../interfaces/types_app_client.AppClientDeployParams.md)
- [AppSourceMaps](../interfaces/types_app_client.AppSourceMaps.md)
- [FundAppAccountParams](../interfaces/types_app_client.FundAppAccountParams.md)
- [ResolveAppById](../interfaces/types_app_client.ResolveAppById.md)
- [SourceMapExport](../interfaces/types_app_client.SourceMapExport.md)

### Type Aliases

- [AppClientCallArgs](types_app_client.md#appclientcallargs)
- [AppClientCallParams](types_app_client.md#appclientcallparams)
- [AppClientClearStateParams](types_app_client.md#appclientclearstateparams)
- [AppClientCreateOnComplete](types_app_client.md#appclientcreateoncomplete)
- [AppClientCreateParams](types_app_client.md#appclientcreateparams)
- [AppClientUpdateParams](types_app_client.md#appclientupdateparams)
- [AppDetails](types_app_client.md#appdetails)
- [AppSpecAppDetails](types_app_client.md#appspecappdetails)
- [ResolveAppByCreatorAndName](types_app_client.md#resolveappbycreatorandname)

## Type Aliases

### AppClientCallArgs

Ƭ **AppClientCallArgs**: [`AppClientCallRawArgs`](../interfaces/types_app_client.AppClientCallRawArgs.md) \| [`AppClientCallABIArgs`](../interfaces/types_app_client.AppClientCallABIArgs.md)

The arguments to pass to an Application Client smart contract call

#### Defined in

[src/types/app-client.ts:150](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L150)

___

### AppClientCallParams

Ƭ **AppClientCallParams**: [`AppClientCallArgs`](types_app_client.md#appclientcallargs) & [`AppClientCallCoreParams`](../interfaces/types_app_client.AppClientCallCoreParams.md)

Parameters to construct a ApplicationClient contract call

#### Defined in

[src/types/app-client.ts:163](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L163)

___

### AppClientClearStateParams

Ƭ **AppClientClearStateParams**: [`AppClientCallRawArgs`](../interfaces/types_app_client.AppClientCallRawArgs.md) & [`AppClientCallCoreParams`](../interfaces/types_app_client.AppClientCallCoreParams.md)

Parameters to construct a ApplicationClient clear state contract call

#### Defined in

[src/types/app-client.ts:166](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L166)

___

### AppClientCreateOnComplete

Ƭ **AppClientCreateOnComplete**: `Object`

On-complete action parameter for creating a contract using ApplicationClient

#### Type declaration

| Name | Type | Description |
| :------ | :------ | :------ |
| `onCompleteAction?` | `Exclude`<[`AppCallType`](types_app.md#appcalltype), ``"clear_state"``\> \| `Exclude`<`OnApplicationComplete`, `OnApplicationComplete.ClearStateOC`\> | Override the on-completion action for the create call; defaults to NoOp |

#### Defined in

[src/types/app-client.ts:178](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L178)

___

### AppClientCreateParams

Ƭ **AppClientCreateParams**: [`AppClientCallParams`](types_app_client.md#appclientcallparams) & [`AppClientCompilationParams`](../interfaces/types_app_client.AppClientCompilationParams.md) & [`AppClientCreateOnComplete`](types_app_client.md#appclientcreateoncomplete)

Parameters for creating a contract using ApplicationClient

#### Defined in

[src/types/app-client.ts:184](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L184)

___

### AppClientUpdateParams

Ƭ **AppClientUpdateParams**: [`AppClientCallParams`](types_app_client.md#appclientcallparams) & [`AppClientCompilationParams`](../interfaces/types_app_client.AppClientCompilationParams.md)

Parameters for updating a contract using ApplicationClient

#### Defined in

[src/types/app-client.ts:187](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L187)

___

### AppDetails

Ƭ **AppDetails**: { `deployTimeParams?`: [`TealTemplateParams`](../interfaces/types_app.TealTemplateParams.md) ; `params?`: `SuggestedParams` ; `sender?`: [`SendTransactionFrom`](types_transaction.md#sendtransactionfrom)  } & [`ResolveAppById`](../interfaces/types_app_client.ResolveAppById.md) \| [`ResolveAppByCreatorAndName`](types_app_client.md#resolveappbycreatorandname)

The details of an AlgoKit Utils deployed app

#### Defined in

[src/types/app-client.ts:82](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L82)

___

### AppSpecAppDetails

Ƭ **AppSpecAppDetails**: [`AppDetails`](types_app_client.md#appdetails) & { `app`: [`AppSpec`](../interfaces/types_app_spec.AppSpec.md) \| `string`  }

The details of an ARC-0032 app spec specified, AlgoKit Utils deployed app

#### Defined in

[src/types/app-client.ts:94](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L94)

___

### ResolveAppByCreatorAndName

Ƭ **ResolveAppByCreatorAndName**: `Object`

Configuration to resolve app by creator and name `getCreatorAppsByName`

#### Type declaration

| Name | Type | Description |
| :------ | :------ | :------ |
| `creatorAddress` | `string` | The address of the app creator account to resolve the app by |
| `findExistingUsing` | `Indexer` \| [`AppLookup`](../interfaces/types_app.AppLookup.md) | The mechanism to find an existing app instance metadata for the given creator and name; either: * An indexer instance to search the creator account apps; or * The cached value of the existing apps for the given creator from `getCreatorAppsByName` |
| `name?` | `string` | The optional name override to resolve the app by within the creator account (default: uses the name in the ABI contract) |
| `resolveBy` | ``"creatorAndName"`` | How the app ID is resolved, either by `'id'` or `'creatorAndName'`; must be `'creatorAndName'` if you want to use `deploy` |

#### Defined in

[src/types/app-client.ts:57](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L57)
