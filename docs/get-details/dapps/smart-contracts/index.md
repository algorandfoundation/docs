title: Introduction

Algorand Smart Contracts (ASC1) are small programs that serve various functions on the blockchain and operate on layer-1. Smart contracts are separated into two main categories, smart contracts, and smart signatures. These types are also referred to as stateful and stateless contracts respectively. The type of contract that is written will determine when and how the logic of the program is evaluated. See the following sections to understand how each type of contract is used on the Algorand blockchain. Both types of contracts are written in the [Transaction Execution Approval Language (TEAL)](../avm/teal), which is an assembly-like language that is interpreted by the [Algorand Virtual Machine (AVM)](../avm) running within an Algorand node. TEAL programs can be written by hand or by using the Python language with the PyTEAL compiler. 


!!! warning
    When writing smart contracts, make sure to follow [TEAL guidelines](../avm/teal/guidelines). This is very important in order to prevent smart contracts from being compromised.

# Quick start videos

If you prefer videos, take a look at this 6 minute guide to learn about Smart Contracts Overview.

<iframe width="100%" style="aspect-ratio:16/9" src="https://www.youtube-nocookie.com/embed/_8vLEeJDjlc" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>

# Smart contracts
Smart contracts are contracts that, once deployed, are remotely callable from any node in the Algorand blockchain. Once deployed, the on-chain instantiation of the contract is referred to as an Application and assigned an Application Id. These applications are triggered by a specific type of transaction called an Application Call transaction.  These on-chain applications handle the primary decentralized logic of a dApp. 

- Applications can [modify state](./apps/#modifying-state-in-smart-contract) associated with the application (global state) or a per application+account (local state) basis. 
- Applications can [access](./apps/#using-assets-in-smart-contracts) on-chain values, such as account balances, asset configuration parameters, or the latest block time.
- Applications can [execute transactions](./apps/#inner-transactions) as part of the execution of the logic. One type of transaction they can perform, as of AVM 1.1, is an Application Call transaction which allows one application to call another.  This ability to call other applications enables composability between applications. 
- Applications have an associated [Application Account](./apps/#using-a-smart-contract-as-an-escrow) that can hold Algos or ASAs balances and can be used as on-chain escrow accounts. 

To provide a standard method for exposing an API and encoding/decoding data types from application call transactions, the [ABI](/docs/get-details/dapps/smart-contracts/ABI/) should be used. 

For more information on smart contracts, see the [smart contract documentation](./apps). For more information on building smart contracts in PyTeal see the [build with python documentation](/docs/get-details/dapps/writing-contracts/pyteal).
For more information on using smart contracts with the SDKs see the [Interacting with smart contracts documentation](./frontend/smartsigs.md).
For more information on debugging a smart contract, see the [debugging](./debugging.md) page.

# Smart signatures
Smart signatures contain logic that is used to sign transactions, primarily for signature delegation. The logic of the smart signature is submitted with a transaction. While the logic in the smart signature is stored on the chain as part of resolving the transaction, the logic is not remotely callable. Any new transaction that relies on the same smart signature would resubmit the logic. When the logic is submitted to a node the AVM evaluates the logic, where it either fails or succeeds. If a smart signature’s logic fails when executed by the AVM, the associated transaction will not be executed. 

Smart signatures can be used in two different modes. When compiled smart signatures produce an Algorand account that functions similar to any other account on the blockchain. These accounts can hold Algos or assets. These funds are only allowed to leave the account if a transaction occurs from the account that successfully executes the logic within the smart signature. This is similar in functionality to a smart contract escrow, but the logic must be submitted for every transaction from the account. Smart signatures can also also be used to delegate some portion of authority to another account. In this case, an account can sign the smart signature which can then be used at a later time to sign a transaction from the original signer’s account. This is referred to as account delegation. See the [modes of use documentation](./smartsigs/modes.md) for more details on these two types of smart signatures.   

Once a transaction that is signed with a smart signature, is submitted it is evaluated by an Algorand node using the Algorand Virtual Machine. These contracts only have access to a few global variables, some temporary scratch space, and the properties of the transaction(s) they are submitted with. 

For more information on smart signatures, see the [smart signature documentation](./smartsigs). For more information on building contracts in PyTeal see the [build with Python documentation](/docs/get-details/dapps/writing-contracts/pyteal).
For more information on using smart signatures with the SDKs see the [Interacting with smart signature documentation](./frontend/smartsigs).
For more information on debugging a smart contract, see the [debugging](./debugging.md) page.


For more information on the [AVM](../avm) or the [TEAL language](../avm/teal) see the developer documentation.
