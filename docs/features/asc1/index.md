title: Section Overview

**Algorand Smart Contracts (ASC1)** are small programs written in an assembly-like language that can be used as a replacement for signatures within a transaction. The language of Algorand Smart Contracts is named **Transaction Execution Approval Language** or **TEAL**. 

TEAL programs have one primary function and that is to determine whether or not a transaction is approved by analyzing it against its own logic and returning either `true` or `false` - approved or not approved, respectively. Algorand transactions can be authorized by [a signature from a single account](../transactions/signatures.md#single-signatures) or a [multisignature account](../transactions/signatures.md#multisignatures). Smart Contracts allow for a third type of signature using a TEAL program, called a **logic signature (LogicSig)**. Algorand Smart Contracts provide two modes for TEAL logic to operate as a LogicSig, which are discussed in [Usage Modes](modes.md).

Before getting started, it is important to understand that there are several ways to approach learning about Algorand Smart Contracts. At a high-level, learning about Smart Contracts can be separated into learning how to write smart contracts in TEAL and learning how to use smart contracts from either the command line, with `goal`, or by calling them from the SDKs.


# How to use and deploy Smart Contracts
- [Modes of Use](./modes.md) - How to issue transactions with Smart Contracts. 

- [A Contract Walkthrough](goal_teal_walkthrough.md) - A walkthrough of a sample contract, including how to compile and use with `goal`
- [Smart Contracts in the SDKs](sdks.md) - Learn how you can use Smart Contracts within your preferred SDK.

# All about TEAL, the language of Smart Contracts
- [TEAL, the Language of Smart Contracts](teal_overview.md) - A detailed overview of TEAL, the Smart Contract language, and how it is processed.
- [TEAL Specification Reference](../../reference/teal/specification.md) - The full technical spec for TEAL the programming language.
- [TEAL Opcodes](../../reference/teal/opcodes.md) - The full list of TEAL opcodes.
- [Guidelines for writing with TEAL](../../reference/teal/guidelines.md) - Tips and best practices when writing your own Smart Contracts.





