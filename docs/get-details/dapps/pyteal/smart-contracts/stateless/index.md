title: Overview

**Algorand Smart Contracts (ASC1)**, provides both stateful and stateless smart contracts. The intended use case for the contract will determine the appropriate type to use. Stateless smart contracts are used to approve or deny transactions and are evaluated at the time a transaction is submitted. In this sense, no long term state is associated with the contract. If your application requires long term state, see the stateful smart contract documentation. Stateless smart contracts are primarily intended to replace signature authority on a transaction.

Most Algorand transactions are authorized by [a signature from a single account](../../transactions/signatures.md#single-signatures) or a [multisignature account](../../transactions/signatures.md#multisignatures). Algorand’s stateful smart contracts allow for a third type of signature using a **Transaction Execution Approval Language (TEAL)** program, called a **logic signature (LogicSig)**. Stateless smart contracts provide two modes for TEAL logic to operate as a LogicSig, to create a contract account that functions similar to an escrow or to delegate signature authority to another account. These two modes are discussed in detail in [Usage Modes](modes.md).

These stateless smart contracts can be written in TEAL or in Python using the [PyTeal library](../teal/pyteal.md). If written in Python the library will return the automatically generated TEAL to be used by either the SDKs or the `goal` command-line tool.  

To learn more about the TEAL language specification see the [TEAL Specification Reference](../../../reference/teal/specification.md). To get a high-level understanding of how TEAL is processed see the [TEAL Overview guide](../teal/index.md).

To compile and use a TEAL program as a stateless smart contract, see [A Contract Walkthrough](walkthrough.md).

To start working with stateless smart contracts with the SDKs, visit the [Using the SDKs](sdks.md)section.


