title: SmartSig details

**Algorand Smart Contracts (ASC1)**, provides both smart contracts and smart signatures. The intended use case for the contract will determine the appropriate type to use. Smart signatures are primarily used to delegate signature authority. Smart signatures can also be used as escrow or contract accounts, but in most cases it is preferable to use a [smart contract](../apps/) when an escrow is required.

Most Algorand transactions are authorized by a [signature from a single account](../../../transactions/signatures.md#single-signatures) or a [multisignature account](../../../transactions/signatures#multisignatures). Algorand’s smart signatures allow for a third type of signature using a **Transaction Execution Approval Language (TEAL)** program, called a **logic signature (LogicSig)**. Smart signatures provide two modes for TEAL logic to operate as a LogicSig, to create a contract account that functions similar to an escrow or to delegate signature authority to another account. These two modes are discussed in detail in [Usage Modes](./modes/).

These smart signatures can be written in [TEAL](/docs/get-details/dapps/avm/teal/) or in Python using the [PyTeal library](/docs/get-details/dapps/writing-contracts/pyteal). If written in Python the library will return the automatically generated TEAL to be used by either the SDKs or the `goal` command-line tool.  

To learn more about the TEAL language specification see the [TEAL Specification Reference](../../avm/teal/specification.md). To get a high-level understanding of how TEAL is processed see the [TEAL Overview guide](../../avm/teal/).

To compile and use a TEAL program as a smart signature with `goal`, see [CLI smart signatures](walkthrough).

To start working with smart signatures with the SDKs, visit the [Interact with smart signatures](../frontend/smartsigs/) section.


