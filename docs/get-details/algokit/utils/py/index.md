title: Overview

A set of core Algorand utilities written in Python and released via PyPi that make it easier to build solutions on Algorand.
This project is part of [AlgoKit](../../index.md).

The goal of this library is to provide intuitive, productive utility functions that make it easier, quicker and safer to build applications on Algorand.
Largely these functions wrap the underlying Algorand SDK, but provide a higher level interface with sensible defaults and capabilities for common tasks.

!!! Note
	If you prefer TypeScript there's an equivalent [TypeScript utility library](../ts/index.md).

[Core principles](#core-principles) | [Installation](#installation) | [Usage](#usage) | [Capabilities](#capabilities)

# Core principles

This library is designed with the following principles:

- **Modularity** - This library is a thin wrapper of modular building blocks over the Algorand SDK; the primitives from the underlying Algorand SDK are
  exposed and used wherever possible so you can opt-in to which parts of this library you want to use without having to use an all or nothing approach.
- **Type-safety** - This library provides strong TypeScript support with effort put into creating types that provide good type safety and intellisense.
- **Productivity** - This library is built to make solution developers highly productive; it has a number of mechanisms to make common code easier and terser to write


# Installation

This library can be installed from PyPi using pip or poetry, e.g.:

```
pip install algokit-utils
poetry add algokit-utils
```


# Usage

To use this library simply include the following at the top of your file:

```python
import algokit_utils
```

Then you can use intellisense to auto-complete the various functions and types that are available by typing `algokit_utils.` in your favourite Integrated Development Environment (IDE),
or you can refer to the [reference documentation](apidocs.md).

## Types

The library contains extensive type hinting combined with a tool like MyPy this can help identify issues where incorrect types have been used, or used incorrectly.


# Capabilities

The library helps you with the following capabilities:

- Core capabilities
  - [**Client management**](capabilities/client.md) - Creation of algod, indexer and kmd clients against various networks resolved from environment or specified configuration
  - [**Account management**](capabilities/account.md) - Creation and use of accounts including mnemonic, multisig, transaction signer, idempotent KMD accounts and environment variable injected
- Higher-order use cases
  - [**ARC-0032 Application Spec client**](capabilities/app-client.md) - Builds on top of the App management and App deployment capabilities to provide a high productivity application client that works with ARC-0032 application spec defined smart contracts (e.g. via Beaker)
  - [**App deployment**](capabilities/app-deploy.md) - Idempotent (safely retryable) deployment of an app, including deploy-time immutability and permanence control and TEAL template substitution
  - [**Algo transfers**](capabilities/transfer.md) - Ability to easily initiate algo transfers between accounts, including dispenser management and idempotent account funding


