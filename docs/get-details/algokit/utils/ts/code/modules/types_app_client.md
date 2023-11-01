[@algorandfoundation/algokit-utils](../index.md) / types/app-client

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
- [ResolveAppByIdBase](../interfaces/types_app_client.ResolveAppByIdBase.md)
- [SourceMapExport](../interfaces/types_app_client.SourceMapExport.md)

### Type Aliases

- [AppClientCallArgs](types_app_client.md#appclientcallargs)
- [AppClientCallParams](types_app_client.md#appclientcallparams)
- [AppClientClearStateParams](types_app_client.md#appclientclearstateparams)
- [AppClientCreateOnComplete](types_app_client.md#appclientcreateoncomplete)
- [AppClientCreateParams](types_app_client.md#appclientcreateparams)
- [AppClientUpdateParams](types_app_client.md#appclientupdateparams)
- [AppDetails](types_app_client.md#appdetails)
- [AppDetailsBase](types_app_client.md#appdetailsbase)
- [AppSpecAppDetails](types_app_client.md#appspecappdetails)
- [AppSpecAppDetailsBase](types_app_client.md#appspecappdetailsbase)
- [AppSpecAppDetailsByCreatorAndName](types_app_client.md#appspecappdetailsbycreatorandname)
- [AppSpecAppDetailsById](types_app_client.md#appspecappdetailsbyid)
- [ResolveAppByCreatorAndName](types_app_client.md#resolveappbycreatorandname)
- [ResolveAppByCreatorAndNameBase](types_app_client.md#resolveappbycreatorandnamebase)

## Type Aliases

### AppClientCallArgs

Ƭ **AppClientCallArgs**: [`AppClientCallRawArgs`](../interfaces/types_app_client.AppClientCallRawArgs.md) \| [`AppClientCallABIArgs`](../interfaces/types_app_client.AppClientCallABIArgs.md)

The arguments to pass to an Application Client smart contract call

#### Defined in

[src/types/app-client.ts:168](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L168)

___

### AppClientCallParams

Ƭ **AppClientCallParams**: [`AppClientCallArgs`](types_app_client.md#appclientcallargs) & [`AppClientCallCoreParams`](../interfaces/types_app_client.AppClientCallCoreParams.md)

Parameters to construct a ApplicationClient contract call

#### Defined in

[src/types/app-client.ts:181](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L181)

___

### AppClientClearStateParams

Ƭ **AppClientClearStateParams**: [`AppClientCallRawArgs`](../interfaces/types_app_client.AppClientCallRawArgs.md) & [`AppClientCallCoreParams`](../interfaces/types_app_client.AppClientCallCoreParams.md)

Parameters to construct a ApplicationClient clear state contract call

#### Defined in

[src/types/app-client.ts:184](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L184)

___

### AppClientCreateOnComplete

Ƭ **AppClientCreateOnComplete**: `Object`

On-complete action parameter for creating a contract using ApplicationClient

#### Type declaration

| Name | Type | Description |
| :------ | :------ | :------ |
| `onCompleteAction?` | `Exclude`<[`AppCallType`](types_app.md#appcalltype), ``"clear_state"``\> \| `Exclude`<`OnApplicationComplete`, `OnApplicationComplete.ClearStateOC`\> | Override the on-completion action for the create call; defaults to NoOp |

#### Defined in

[src/types/app-client.ts:196](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L196)

___

### AppClientCreateParams

Ƭ **AppClientCreateParams**: [`AppClientCallParams`](types_app_client.md#appclientcallparams) & [`AppClientCompilationParams`](../interfaces/types_app_client.AppClientCompilationParams.md) & [`AppClientCreateOnComplete`](types_app_client.md#appclientcreateoncomplete)

Parameters for creating a contract using ApplicationClient

#### Defined in

[src/types/app-client.ts:202](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L202)

___

### AppClientUpdateParams

Ƭ **AppClientUpdateParams**: [`AppClientCallParams`](types_app_client.md#appclientcallparams) & [`AppClientCompilationParams`](../interfaces/types_app_client.AppClientCompilationParams.md)

Parameters for updating a contract using ApplicationClient

#### Defined in

[src/types/app-client.ts:205](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L205)

___

### AppDetails

Ƭ **AppDetails**: [`AppDetailsBase`](types_app_client.md#appdetailsbase) & [`ResolveAppById`](../interfaces/types_app_client.ResolveAppById.md) \| [`ResolveAppByCreatorAndName`](types_app_client.md#resolveappbycreatorandname)

The details of an AlgoKit Utils deployed app

#### Defined in

[src/types/app-client.ts:100](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L100)

___

### AppDetailsBase

Ƭ **AppDetailsBase**: `Object`

The details of an AlgoKit Utils deployed app

#### Type declaration

| Name | Type | Description |
| :------ | :------ | :------ |
| `deployTimeParams?` | [`TealTemplateParams`](../interfaces/types_app.TealTemplateParams.md) | Optionally provide any deploy-time parameters to replace in the TEAL code; if specified here will get used in calls to `deploy`, `create` and `update` unless overridden in those calls |
| `params?` | `SuggestedParams` | Default suggested params object to use |
| `sender?` | [`SendTransactionFrom`](types_transaction.md#sendtransactionfrom) | Default sender to use for transactions issued by this application client |

#### Defined in

[src/types/app-client.ts:88](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L88)

___

### AppSpecAppDetails

Ƭ **AppSpecAppDetails**: [`AppSpecAppDetailsBase`](types_app_client.md#appspecappdetailsbase) & [`AppDetails`](types_app_client.md#appdetails)

The details of an ARC-0032 app spec specified, AlgoKit Utils deployed app

#### Defined in

[src/types/app-client.ts:118](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L118)

___

### AppSpecAppDetailsBase

Ƭ **AppSpecAppDetailsBase**: `Object`

The details of an ARC-0032 app spec specified, AlgoKit Utils deployed app

#### Type declaration

| Name | Type | Description |
| :------ | :------ | :------ |
| `app` | [`AppSpec`](../interfaces/types_app_spec.AppSpec.md) \| `string` | The ARC-0032 application spec as either: * Parsed JSON `AppSpec` * Raw JSON string |

#### Defined in

[src/types/app-client.ts:103](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L103)

___

### AppSpecAppDetailsByCreatorAndName

Ƭ **AppSpecAppDetailsByCreatorAndName**: [`AppSpecAppDetailsBase`](types_app_client.md#appspecappdetailsbase) & [`AppDetailsBase`](types_app_client.md#appdetailsbase) & [`ResolveAppByCreatorAndNameBase`](types_app_client.md#resolveappbycreatorandnamebase)

The details of an ARC-0032 app spec specified, AlgoKit Utils deployed app by creator and name

#### Defined in

[src/types/app-client.ts:115](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L115)

___

### AppSpecAppDetailsById

Ƭ **AppSpecAppDetailsById**: [`AppSpecAppDetailsBase`](types_app_client.md#appspecappdetailsbase) & [`AppDetailsBase`](types_app_client.md#appdetailsbase) & [`ResolveAppByIdBase`](../interfaces/types_app_client.ResolveAppByIdBase.md)

The details of an ARC-0032 app spec specified, AlgoKit Utils deployed app by id

#### Defined in

[src/types/app-client.ts:112](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L112)

___

### ResolveAppByCreatorAndName

Ƭ **ResolveAppByCreatorAndName**: [`ResolveAppByCreatorAndNameBase`](types_app_client.md#resolveappbycreatorandnamebase) & { `resolveBy`: ``"creatorAndName"``  }

Configuration to resolve app by creator and name `getCreatorAppsByName`

#### Defined in

[src/types/app-client.ts:69](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L69)

___

### ResolveAppByCreatorAndNameBase

Ƭ **ResolveAppByCreatorAndNameBase**: `Object`

Configuration to resolve app by creator and name `getCreatorAppsByName`

#### Type declaration

| Name | Type | Description |
| :------ | :------ | :------ |
| `creatorAddress` | `string` | The address of the app creator account to resolve the app by |
| `findExistingUsing` | `Indexer` \| [`AppLookup`](../interfaces/types_app.AppLookup.md) | The mechanism to find an existing app instance metadata for the given creator and name; either: * An indexer instance to search the creator account apps; or * The cached value of the existing apps for the given creator from `getCreatorAppsByName` |
| `name?` | `string` | The optional name override to resolve the app by within the creator account (default: uses the name in the ABI contract) |

#### Defined in

[src/types/app-client.ts:56](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/app-client.ts#L56)
