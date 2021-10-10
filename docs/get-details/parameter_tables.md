title: Algorand parameter tables
[comment]: <> (generated from: https://docs.google.com/spreadsheets/d/12E_B6AYBfhrOoylO8NZysqKns4rHKklV8I0fejiwMXE)

## Minimum balance
|Name|Current value|Developer doc|Consensus parameter name in (.go)|Note|
|-|-|-|-|-|
|Default|0.1 Algos|[reference](./accounts/index.md#minimum-balance)|MinBalance||
|Opt-in ASA|+ 0.1 Algos|[reference](./asa.md#assets-overview)|MinBalance||
|Created ASA|+ 0.1 Algos|[reference](./asa.md#assets-overview)|MinBalance|creator of ASA does not need to opt in|
|Opt-in Application|(see below)|||creator of application only needs to opt in if they use local state|
|Created Application|(see below)||||

## Minimum balance for smart contract
|Name|Current value|Developer doc|Consensus parameter name in (.go)|Note|
|-|-|-|-|-|
|Per page application creation fee|0.1 Algos|[reference](./dapps/smart-contracts/apps/index.md#minimum-balance-requirement-for-a-smart-contract)|AppFlatParamsMinBalance||
|Flat for application opt-in|0.1 Algos|[reference](./dapps/smart-contracts/apps/index.md#minimum-balance-requirement-for-a-smart-contract)|AppFlatOptInMinBalance||
|Per state entry|0.025 Algos|[reference](./dapps/smart-contracts/apps/index.md#minimum-balance-requirement-for-a-smart-contract)|SchemaMinBalancePerEntry||
|Addition per integer entry|0.0035 Algos|[reference](./dapps/smart-contracts/apps/index.md#minimum-balance-requirement-for-a-smart-contract)|SchemaUintMinBalance||
|Addition per byte slice entry|0.025 Algos|[reference](./dapps/smart-contracts/apps/index.md#minimum-balance-requirement-for-a-smart-contract)|SchemaBytesMinBalance||

## Transaction fee
|Name|Current value|Developer doc|Consensus parameter name in (.go)|Note|
|-|-|-|-|-|
|Minimum transaction fee, in all cases|0.001 Algos|[reference](./transactions/index.md#fees)|MinTxnFee||
|Additional minimum constraint if congestion|additional fee per byte|[reference](./transactions/index.md#fees)|-||

## Others
|Name|Current value|Developer doc|Consensus parameter name in (.go)|Note|
|-|-|-|-|-|
|Max number transactions in a group / atomic transfer|16|[reference](./atomic_transfers.md)|MaxTxGroupSize||
|Maximum size of a block|1000000 bytes|-|MaxTxnBytesPerBlock||
|Maximum size of note|1024 bytes|-|MaxTxnNoteBytes||
|Maximum transaction life|1000 rounds|[reference](./transactions/index.md#setting-first-and-last-valid)|MaxTxnLife||

## ASA constraints
|Name|Current value|Developer doc|Consensus parameter name in (.go)|Note|
|-|-|-|-|-|
|Max number of ASAs (create and opt-in)|1000|[reference](./asa.md)|MaxAssetsPerAccount||
|Max asset name size|32 bytes|[reference](./asa.md)|MaxAssetNameBytes||
|Max unit name size|8 bytes|[reference](./asa.md)|MaxAssetUnitNameBytes||
|Max URL size|96 bytes|[reference](./asa.md)|MaxAssetURLBytes||
|Metadata hash|32 bytes|[reference](./asa.md)||padded with zeros|

## Stateless smart contract constraints
|Name|Current value|Developer doc|Consensus parameter name in (.go)|Note|
|-|-|-|-|-|
|Max size of compiled TEAL code combined with arguments|1000 bytes|[reference](./dapps/avm/teal/specification.md#execution-environment-for-logicsigs)|LogicSigMaxSize||
|Max cost of TEAL code|20000|[reference](./dapps/avm/teal/specification.md#execution-environment-for-logicsigs)|LogicSigMaxCost||

## Stateful smart contract constraints
|Name|Current value|Developer doc|Consensus parameter name in (.go)|Note|
|-|-|-|-|-|
|Page size of compiled approval + clear TEAL code|2048 bytes|[reference](./dapps/avm/teal/index.md#dynamic-operational-cost-of-teal-opcodes)|MaxAppProgramLen|by default, each application has a single page|
|Max extra app pages |3|[reference](./dapps/smart-contracts/apps/index.md#creating-the-smart-contract)|MaxExtraAppProgramPages|an application can "pay" for additional pages via minimum balance|
|Max cost of approval TEAL code|700|[reference](./dapps/avm/teal/index.md#dynamic-operational-cost-of-teal-opcodes)|MaxAppProgramCost||
|Max cost of clear TEAL code|700|[reference](./dapps/avm/teal/index.md#dynamic-operational-cost-of-teal-opcodes)|MaxAppProgramCost||
|Max number of scratch variables|256][reference](./dapps/avm/teal/#storing-and-loading-from-scratchspace)|||
|Max depth of stack|1000][reference](./dapps/avm/teal/index.md)|MaxStackDepth||
|Max number of arguments|16|[reference](./dapps/smart-contracts/apps/index.md#passing-arguments-to-stateful-smart-contracts)|MaxAppArgs||
|Max combined size of arguments|2048 bytes|[reference](./dapps/smart-contracts/apps/index.md#passing-arguments-to-stateful-smart-contracts)|MaxAppTotalArgLen||
|Max number of global state keys|64|[reference](./dapps/smart-contracts/apps/index.md#creating-the-smart-contract)|MaxGlobalSchemaEntries||
|Max number of local state keys|16|[reference](./dapps/smart-contracts/apps/index.md#creating-the-smart-contract)|MaxLocalSchemaEntries||
|Max number of log messages|32|[reference](./dapps/smart-contracts/apps/index.md#TODO)|MaxLogCalls||
|Max size of log messages|1024|[reference](./dapps/smart-contracts/apps/index.md#TODO)|MaxLogSize||
|Max key size|64 bytes|[reference](./dapps/smart-contracts/apps/index.md#creating-the-smart-contract)|MaxAppKeyLen||
|Max []byte value size|128 bytes|[reference](./dapps/smart-contracts/apps/index.md#creating-the-smart-contract)|MaxAppBytesValueLen||
|Max key + value size|128 bytes|[reference](./dapps/smart-contracts/apps/index.md#creating-the-smart-contract)|MaxAppSumKeyValueLens||
|Max number of foreign accounts|4|[reference](./dapps/smart-contracts/apps/index.md#stateful-contract-arrays)|MaxAppTxnAccounts||
|Max number of foreign ASAs|8|[reference](./dapps/smart-contracts/apps/index.md#stateful-contract-arrays)|MaxAppTxnForeignAssets||
|Max number of foreign applications|8|[reference](./dapps/smart-contracts/apps/index.md#stateful-contract-arrays)|MaxAppTxnForeignApps||
|Max number of foreign accounts + ASAs + applications|8|[reference](./dapps/smart-contracts/apps/index.md#stateful-contract-arrays)|zMaxAppTotalTxnReferences||
|Max number of created applications|10|[reference](./dapps/smart-contracts/apps/index.md#opt-into-the-smart-contract)|MaxAppsCreated||
|Max number of opt-in applications|50|[reference](./dapps/smart-contracts/apps/index.md#opt-into-the-smart-contract)|MaxAppsOptedIn||
