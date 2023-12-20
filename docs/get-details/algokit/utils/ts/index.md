title: Overview

A set of core Algorand utilities written in TypeScript and released via npm that make it easier to build, test and deploy solutions on the Algorand Blockchain, including APIs, console apps and dApps. This project is part of [AlgoKit](../../index.md).

The goal of this library is to provide intuitive, productive utility functions that make it easier, quicker and safer to build applications on Algorand. Largely these functions provide a thin wrapper over the underlying Algorand SDK, but provide a higher level interface with sensible defaults and capabilities for common tasks that make development faster and easier.

!!! Note
	If you prefer Python there's an equivalent [Python utility library](../py).

[Core principles](#core-principles) | [Installation](#installation) | [Usage](#usage) | [Config and logging](#config-and-logging) | [Capabilities](#capabilities)

# Core principles

This library is designed with the following principles:

- **Modularity** - This library is a thin wrapper of modular building blocks over the Algorand SDK; the primitives from the underlying Algorand SDK are exposed and used wherever possible so you can opt-in to which parts of this library you want to use without having to use an all or nothing approach.
- **Type-safety** - This library provides strong TypeScript support with effort put into creating types that provide good type safety and intellisense.
- **Productivity** - This library is built to make solution developers highly productive; it has a number of mechanisms to make common code easier and terser to write

# Installation

This library can be installed from NPM using your favourite npm client, e.g.:

```
npm install @algorandfoundation/algokit-utils
```

## Peer Dependencies

This library requires `algosdk` as a peer dependency. Ensure you have it installed in your project.

# Usage

To use this library simply include the following at the top of your file:

```typescript
import * as algokit from '@algorandfoundation/algokit-utils'
```

Then you can use intellisense to auto-complete the various functions that are available by typing `algokit.` in your favourite Integrated Development Environment (IDE), or you can refer to the [reference documentation](code/modules/index.md).

## Testing

To use the automated testing functionality, you can import the testing module:

```typescript
import * as algotesting from '@algorandfoundation/algokit-utils/testing'
```

Or, you can generally get away with just importing the `algorandFixture` since it exposes the rest of the functionality in a manner that is easy to integrate with an underlying test framework like Jest or vitest:

```typescript
import { algorandFixture } from '@algorandfoundation/algokit-utils/testing'
```

To see what's available feel free to consult the [reference documentation](code/modules/testing.md) or consulting the [testing capability page](capabilities/testing.md).

## Types

If you want to extend or pass around any of the types the various functions take then they are all defined in isolated modules under the `types` namespace. This is to provide a better intellisense experience without overwhelming you with hundreds of types. If you determine a type to import then you can import it like so:

```typescript
import {<type>} from '@algorandfoundation/types/<module>'
```

Where `<type>` would be replaced with the type and `<module>` would be replaced with the module. You can use intellisense to discover the modules and types in your favourite IDE, or you can explore the [types modules in the reference documentation](code/index.md#modules).

# Config and logging

To configure the AlgoKit library you can make use of the `algokit.Config` object, which has a `configure` method that lets you configure some or all of the configuration options.

## Logging

AlgoKit has an in-built logging abstraction that allows the library to issue log messages without coupling the library to a particular logging library. This means you can access the AlgoKit Utils logs within your existing logging library if you have one.

To do this you need to create a logging translator that exposes the following interface ([`Logger`](./code/modules/types_logging.md#logger)):

```typescript
export type Logger = {
  error(message: string, ...optionalParams: unknown[]): void
  warn(message: string, ...optionalParams: unknown[]): void
  info(message: string, ...optionalParams: unknown[]): void
  verbose(message: string, ...optionalParams: unknown[]): void
  debug(message: string, ...optionalParams: unknown[]): void
}
```

!!! Note
	this interface type is directly compatible with [Winston](https://github.com/winstonjs/winston).

By default, the [`ConsoleLogger`](./code/modules/types_logging.md#consolelogger) is set as the logger, which will send log messages to the various `console.*` methods. There is also a [`NullLogger`](./code/modules/types_logging.md#nulllogger) if you want to disable logging.

If you want to override the logger you can use the following:

```typescript
algokit.configure({ logger: myLogger })
```

To retrieve the current debug state you can use [`algokit.Config.logger`](./code/interfaces/types_config.Config.md). To get a logger that is optionally set to the null logger based on a boolean flag you can use the [`algokit.Config.getLogger(useNullLogger)`](./code/classes/types_config.UpdatableConfig.md#getlogger) function.

## Debug mode

To turn on debug mode you can use the following:

```typescript
algokit.configure({ debug: true })
```

To retrieve the current debug state you can use [`algokit.Config.debug`](./code/interfaces/types_config.Config.md).

This will turn on things like automatic tracing and more verbose logging. It's likely this option will result in extra HTTP calls to algod so worth being careful when it's turned on.

If you want to temporarily turn it on you can use the [`withDebug`](./code/classes/types_config.UpdatableConfig.md#withdebug) function:

```typescript
algokit.Config.withDebug(() => {
  // Do stuff with algokit.Config.debug set to true
})
```

# Capabilities

The library helps you with the following capabilities:

- Core capabilities
  - [**Client management**](capabilities/client.md) - Creation of (auto-retry) algod, indexer and kmd clients against various networks resolved from environment or specified configuration
  - [**Account management**](capabilities/account.md) - Creation and use of accounts including mnemonic, rekeyed, multisig, transaction signer ([useWallet](https://github.com/TxnLab/use-wallet) for dApps and Atomic Transaction Composer compatible signers), idempotent KMD accounts and environment variable injected
  - [**Algo amount handling**](capabilities/amount.md) - Reliable and terse specification of microAlgo and Algo amounts and conversion between them
  - [**Transaction management**](capabilities/transaction.md) - Ability to send single, grouped or Atomic Transaction Composer transactions with consistent and highly configurable semantics, including configurable control of transaction notes (including ARC-0002), logging, fees, multiple sender account types, and sending behaviour
- Higher-order use cases
  - [**App management**](capabilities/app.md) - Creation, updating, deleting, calling (ABI and otherwise) smart contract apps and the metadata associated with them (including state and boxes)
  - [**App deployment**](capabilities/app-deploy.md) - Idempotent (safely retryable) deployment of an app, including deploy-time immutability and permanence control and TEAL template substitution
  - [**ARC-0032 Application Spec client**](capabilities/app-client.md) - Builds on top of the App management and App deployment capabilities to provide a high productivity application client that works with ARC-0032 application spec defined smart contracts (e.g. via Beaker)
  - [**Algo transfers**](capabilities/transfer.md) - Ability to easily initiate algo transfers between accounts, including dispenser management and idempotent account funding
  - [**Automated testing**](capabilities/testing.md) - Terse, robust automated testing primitives that work across any testing framework (including jest and vitest) to facilitate fixture management, quickly generating isolated and funded test accounts, transaction logging, indexer wait management and log capture
  - [**Indexer lookups / searching**](capabilities/indexer.md) - Type-safe indexer API wrappers (no more `Record<string, any>` pain), including automatic pagination control

