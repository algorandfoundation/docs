[@algorandfoundation/algokit-utils](../index.md) / [types/debugging](../modules/types_debugging.md) / SimulateAndPersistResponseParams

# Interface: SimulateAndPersistResponseParams

[types/debugging](../modules/types_debugging.md).SimulateAndPersistResponseParams

Parameters to a call that simulates a transaction and persists the response.

## Table of contents

### Properties

- [algod](types_debugging.SimulateAndPersistResponseParams.md#algod)
- [atc](types_debugging.SimulateAndPersistResponseParams.md#atc)
- [bufferSizeMb](types_debugging.SimulateAndPersistResponseParams.md#buffersizemb)
- [projectRoot](types_debugging.SimulateAndPersistResponseParams.md#projectroot)

## Properties

### algod

• **algod**: `default`

algod An Algodv2 client to perform the simulation.

#### Defined in

[src/types/debugging.ts:158](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/debugging.ts#L158)

___

### atc

• **atc**: `AtomicTransactionComposer`

The AtomicTransactionComposer with transaction(s) loaded.

#### Defined in

[src/types/debugging.ts:160](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/debugging.ts#L160)

___

### bufferSizeMb

• **bufferSizeMb**: `number`

bufferSizeMb The buffer size in megabytes.

#### Defined in

[src/types/debugging.ts:164](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/debugging.ts#L164)

___

### projectRoot

• **projectRoot**: `string`

projectRoot The root directory of the project.

#### Defined in

[src/types/debugging.ts:162](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/debugging.ts#L162)
