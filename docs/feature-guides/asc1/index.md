title: Algorand Smart Contracts

**Algorand Smart Contracts (ASC1)** are small programs written in an assembly-like language that can be used as a replacement for signatures within a transaction. The language of Algorand Smart Contracts is named **Transaction Execution Approval Langauge** or **TEAL**. 

TEAL programs have one primary function and that is to determine whether or not a transaction is approved by analyzing it against its own logic and returning either `true` or `false` - approved or not approved, respectively. Algorand transactions can be authorized by [a signature from a single account](../signatures.md#single-signatures) or a [multisignature account](../signatures.md#multisignatures). Smart Contracts allow for a third type of signature using a TEAL program, called a **logic signature (LogicSig)**. Algorand Smart Contracts provide two modes for TEAL logic to operate as a LogicSig, which are discussed in [Usage Modes](modes.md).

Before getting started, it is important to understand that there are several ways to approach learning about Algorand Smart Contracts. At a high-level, learning about Smart Contracts can be separated into *learning how to write smart contracts in TEAL* and *learning how to use smart contracts from either the command line, with `goal`, or by calling them from the SDKs*.

To learn more about the TEAL language specification see the [TEAL Specification Reference](../../reference-docs/teal/specification.md). To get a simplified understanding of how TEAL is processed see the [TEAL Overview guide](teal_overview.md).

To compile and use TEAL programs using `goal`, see [A Contract Walkthrough](goal_teal_walkthrough.md). 

To start working with TEAL with the SDKs, visit the [Using the SDKs](asc1_sdk_usage.md) section.




