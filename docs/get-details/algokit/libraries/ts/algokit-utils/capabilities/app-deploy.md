# App deployment

Idempotent (safely retryable) deployment of an app, including deploy-time immutability and permanence control and TEAL template substitution

App deployment is a higher-order use case capability provided by AlgoKit Utils that builds on top of the core capabilities, particularly [App management](./app.md). It allows you to idempotently (with safe retryability) deploy an app, including deploy-time immutability and permanence control and TEAL template substitution.

To see some usage examples check out the [automated tests](../../src/app-deploy.spec.ts).

## Design

The architecture design behind app deployment is articulated in an [architecture decision record](https://github.com/algorandfoundation/algokit-cli/blob/main/docs/architecture-decisions/2023-01-12_smart-contract-deployment.md). While the implementation will naturally evolve over time and diverge from this record, the principles and design goals behind the design are comprehensively explained.

Namely, it described the concept of a smart contract development lifecycle:

1. Development
   1. **Write** smart contracts
   2. **Transpile** smart contracts with development-time parameters (code configuration) to TEAL Templates
   3. **Verify** the TEAL Templates maintain [output stability](https://github.com/algorandfoundation/algokit-cli/blob/main/docs/articles/output_stability.md) and any other static code quality checks
2. Deployment
   1. **Substitute** deploy-time parameters into TEAL Templates to create final TEAL code
   2. **Compile** the TEAL to create byte code using algod
   3. **Deploy** the byte code to one or more Algorand networks (e.g. LocalNet, TestNet, MainNet) to create Deployed Application(s)
3. Runtime
   1. **Validate** the deployed app via automated testing of the smart contracts to provide confidence in their correctness
   2. **Call** deployed smart contract with runtime parameters to utilise it

![App deployment lifecycle](../images/lifecycle.jpg)

The App deployment capability provided by AlgoKit Utils helps implement **#2 Deployment**.

Furthermore, the implementation contains the following implementation characteristics per the original architecture design:

- Deploy-time parameters can be provided and substituted into a TEAL Template by convention (by replacing `TMPL_{KEY}`)
- Contracts can be built by any smart contract framework that supports [ARC-0032](https://github.com/algorandfoundation/ARCs/pull/150) and [ARC-0004](https://github.com/algorandfoundation/ARCs/blob/main/ARCs/arc-0004.md) ([Beaker](https://beaker.algo.xyz/) or otherwise), which also means the deployment language can be different to the development language e.g. you can deploy a Python smart contract with TypeScript for instance
- There is explicit control of the immutability (updatability / upgradeability) and permanence (deletability) of the smart contract, which can be varied per environment to allow for easier development and testing in non-MainNet environments (by replacing `TMPL_UPDATABLE` and `TMPL_DELETABLE` at deploy-time by convention, if present)
- Contracts are resolvable by a string "name" for a given creator to allow automated determination of whether that contract had been deployed previously or not, but can also be resolved by ID instead

## `getCreatorAppsByName`

There is a method `algokit.getCreatorAppsByName(creatorAccount, indexer)`, which performs a series of indexer lookups that return all apps created by the given creator (as a [`SendTransactionFrom`](account.md#sendtransactionfrom), or the encoded string representation of the public key of the account) indexed by the name it was deployed under if the creation transaction contained the following payload in the transaction note field:

```
ALGOKIT_DEPLOYER:j{name:string, version:string, updatable?:boolean, deletable?:boolean}
```

Any creation transactions or update transactions are then retrieved and processed in chronological order to result in an `AppLookup` object which has the following shape (each app is keyed by name and has the [`AppMetadata`](../code/interfaces/types_app.AppMetadata.md) fields):

```json
{
  "creator": "<creator_address>",
  "apps": {
    "<app_name>": {
      /** The id of the app */
      "appId": 1,
      /** The Algorand address of the account associated with the app */
      "appAddress": "<app_account_address>",
      /** The unique name identifier of the app within the creator account */
      "name": "<app_name>",
      /** The version of app that is / will be deployed */
      "version": "1.0.0",
      /** Whether or not the app is deletable / permanent / unspecified */
      "deletable": true,
      /** Whether or not the app is updatable / immutable / unspecified */
      "updatable": true,
      /** The round the app was created */
      "createdRound": 1,
      /** The last round that the app was updated */
      "updatedRound": 1,
      /** Whether or not the app is deleted */
      "deleted": false,
      /** The metadata when the app was created */
      "createdMetadata": {
        /** The unique name identifier of the app within the creator account */
        "name": "<app_name>",
        /** The version of app that is / will be deployed */
        "version": "1.0.0",
        /** Whether or not the app is deletable / permanent / unspecified */
        "deletable": true,
        /** Whether or not the app is updatable / immutable / unspecified */
        "updatable": true
      }
    }
    //...
  }
}
```

Given there are a number of indexer calls to retrieve this data it's a non-trivial object to create and it's recommended that for the duration you are performing a single deployment you hold a value of it rather than recalculating it. Most AlgoKit Utils functions that need it will also take an optional value of it that will be used in preference to retrieving a fresh version.

## `deployApp`

The method that performs the deployment logic is `algokit.deployApp(deployment, algod, indexer?)`. It performs an idempotent (safely retryable) deployment. It will detect if the app already exists and if it doesn't it will create it. If the app does already exist then it will:

- Detect if the app has been updated (i.e. the logic has changed) and either fail or perform either an update or a replacement based on the deployment configuration.
- Detect if the app has a breaking schema change (i.e. more global or local storage is needed than was originally requested) and either fail or perform a replacement based on the deployment configuration.

It will automatically add metadata to the transaction note of the create or update calls that indicates the name, version, updatability and deletability of the contract. This metadata works in concert with [`getCreatorAppsByName`](#getcreatorappsbyname) to allow the app to be reliably retrieved against that creator in it's currently deployed state.

`deployApp` automatically executes [template substitution](#compilation-and-template-substitution) including deploy-time control of permanence and immutability.

### Input parameters

The first parameter `deployment` is an [`AppDeploymentParams`](../code/interfaces/types_app.AppDeploymentParams.md) object, that consists of the same parameters as [creating an app](app.md#createapp) with the following differences:

- `args` is replaced with three parameters:
  - `createArgs` - The [args](./app.md#appcallargs) to use if a create is needed
  - `updateArgs` - The [args](./app.md#appcallargs) to use if an update is needed
  - `deleteArgs` - The [args](./app.md#appcallargs) to use if a delete is needed
- `onCompleteAction` is replaced with `createOnCompleteAction` - Override the on-completion action for the create call; defaults to NoOp
- `note` is excluded since deployment must add a `ALGOKIT_DEPLOYER:j{...}` note for the [name lookup](#getcreatorappsbyname) to work
- `skipSending`, `skipWaiting` and `atc` are excluded because `deployApp` needs to control the deployment and actually send to the network to function correctly
- `metadata: AppDeployMetadata` is added to allow the [deployment metadata](../code/interfaces/types_app.AppDeployMetadata.md) to be provided (`name`, `version`, `updatable` and `deletable`)
- `deployTimeParams: TealTemplateParams` is added to allow automatic substitution of [deploy-time parameter values](#design)
  - [`TealTemplateParams`](../code/interfaces/types_app.TealTemplateParams.md) is a `key => value` object that will result in `TMPL_{key}` being replaced with `value` (where a string or `Uint8Array` will be appropriately encoded as bytes within the TEAL code)
- `onSchemaBreak?: 'replace' | 'fail' | OnSchemaBreak` determines what should happen if a breaking change to the schema is detected (e.g. if you need more global or local state that was previously requested when the contract was originally created)
- `onUpdate?: 'update' | 'replace' | 'fail' | OnUpdate` determines what should happen if an update to the smart contract is detected (e.g. the TEAL code has changed since last deployment)
- `existingDeployments?: AppLookup` optionally allows the [app lookup](#getcreatorappsbyname) to be skipped if it's already been retrieved

### Idempotency

`deployApp` is idempotent which means you can safely call it again multiple times and it will only apply any changes it detects. If you call it again straight after calling it then it will do nothing.

### Compilation and template substitution

When compiling TEAL template code, the capabilities described in the [design above](#design) are present, namely the ability to supply deploy-time parameters and the ability to control immutability and permanence of the smart contract at deploy-time.

In order for a smart contract to be able to use this functionality, it must have a TEAL Template that contains the following:

- `TMPL_{key}` - Which can be replaced with a number or a string / byte array which wil be automatically hexadecimal encoded
- `TMPL_UPDATABLE` - Which will be replaced with a `1` if an app should be updatable and `0` if it shouldn't (immutable)
- `TMPL_DELETABLE` - Which will be replaced with a `1` if an app should be deletable and `0` if it shouldn't (permanent)

If you are building a smart contract using the [beaker_production AlgoKit template](https://github.com/algorandfoundation/algokit-beaker-default-template) if provides a reference implementation out of the box for the deploy-time immutability and permanence control.

If you passed in a TEAL template for the approvalProgram or clearProgram (i.e. a `string` rather than a `Uint8Array`) then `deployApp` will return the [compilation result](../code/interfaces/types_app.CompiledTeal.md) of substituting then compiling the TEAL template(s) in the following properties of the return value:

- `compiledApproval?: CompiledTeal`
- `compiledClear?: CompiledTeal`

Template substitution is done by executing `algokit.performTemplateSubstitutionAndCompile(tealCode, algod, templateParams?, deploymentMetadata?)`, which in turn calls the following in order:

- `algokit.performTemplateSubstitution(tealCode, templateParams)` - Replaces the `templateParams` by looking for `TMPL_{key}`
- `algokit.replaceDeployTimeControlParams(tealCode, deploymentMetadata)` - If `deploymentMetadata` is provided, it allows for deploy-time immutability and permanence control by replacing `TMPL_UPDATABLE` with `deploymentMetadata.updatable` if it's not `undefined` and replacing `TMPL_DELETABLE` with `deploymentMetadata.deletable` if it's not `undefined`
- `algokit.compileTeal(tealCode, algod)` - Sends the final TEAL to algod for compilation and returns the result including the source map

### Return value

When `deployApp` executes it may do one of the following (which you can determine by looking at the `operationPerformed` field on the return value from the function):

- `create` - The smart contract app is created
- `update` - The smart contract app is updated
- `replace` - The smart contract app was deleted and created again (in an atomic transaction)
- `nothing` - Nothing was done since it was detected the existing smart contract app deployment was up to date

As well as the `operationPerformed` parameter and the [optional compilation result](#compilation-and-template-substitution), the return value will have the [`AppMetadata`](../code/interfaces/types_app.AppMetadata.md) [fields](#getcreatorappsbyname) present.

Based on the value of `operationPerformed` there will be other data available in the return value:

- If `create`, `update` or `replace` then it will have [`ConfirmedTransactionResults`](transaction.md#sendtransactionresult) and `{return?: ABIReturn}` (which will be populated if an ABI method was executed)
- If `replace` then it will also have `{deleteReturn?: ABIReturn, deleteResult: ConfirmedTransactionResult}` to capture the result of the deletion of the existing app
