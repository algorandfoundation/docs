title: Section Index

The **Explore Features** section contains explanations and how-tos for all Algorand features. Code examples are embedded throughout in each of the [SDKs](../reference/index.md#sdks) and [command line interface tools](../reference/index.md#command-line-interface-tools-cli-tools).

# Accounts

[Overview](accounts) - Learn all about Algorand accounts, keys, wallets, and addresses.

[Creation Methods](accounts/create) - How to create standalone accounts, wallet-derived (kmd) accounts, or multisignature accounts.

[Rekeying](accounts/rekey) - Learn how to change the private spending key of an Algorand address.

# Transactions

[Structure](transactions) - A look at the underlying structure of the different types of transactions.

[Signatures](transactions/signatures) - How to authorize transactions with Single Signatures, Multisignatures, and Logic Signatures.

[Offline signatures](transactions/offline_transactions) - How to authorize transactions _without_ requiring an internet connection.

# Assets

[Algorand Standard Assets Overview](asa) - Learn all about assets and how they are implemented on Algorand at layer-1.

[Asset Parameters](asa#asset-parameters) - The mutable and immutable properties of an asset.

[Creating an Asset](asa#creating-an-asset) - [Modifying an Asset](asa#modifying-an-asset) - [Receiving an Asset](asa#receiving-an-asset) - [Transferring an Asset](asa#transferring-an-asset) - [Freezing an Asset](asa#freezing-an-asset) - [Revoking an Asset](asa#revoking-an-asset) - [Destroying an Asset](asa#destroying-an-asset)

[Retrieving Asset Information](asa#retrieve-asset-information) - How to retrieve information about an asset.

# Atomic Transfers

[Overview](atomic_transfers) - What atomic transfers are and how they are they implemented on Algorand.

[Step-by-Step Guide](atomic_transfers#step-by-step-guide) - Walkthrough an atomic transfer in the SDKs and goal.

# Smart Contracts

[Smart Contract Overview](dapps/pyteal/smart-contracts) - An overview of Smart Contracts on Algorand, including Stateful and Stateless versions.

[Stateful Smart Contracts](dapps/pyteal/smart-contracts/stateful) 🔷 - An overview of Stateful Smart Contracts.


[Stateless Smart Contracts](dapps/pyteal/smart-contracts/stateless) - An overview of Stateless Smart Contracts. 

[Logic Signatures](dapps/pyteal/smart-contracts/stateless/modes#logic-signatures) - The authorization method for transactions that use Smart Contracts.

[Modes of Use - Contract Account](dapps/pyteal/smart-contracts/stateless/modes#contract-account) - How to use a Smart Contract as its own Algorand account.

[Modes of Use - Delegated Approval](dapps/pyteal/smart-contracts/stateless/modes#delegated-approval) - How to use a Smart Contract to authorize spending from an existing account.

[Using the SDKS](dapps/pyteal/smart-contracts/frontend/stateless-sdks) - Stateless Smart Contract functionality from within the SDKs.

[The Smart Contract Language](dapps/avm/teal) - Learn all about TEAL, the language of Smart Contracts.

[Contract Walkthrough](dapps/pyteal/smart-contracts/stateless/walkthrough) - A simple command-line walkthrough of creating stateless smart contracts.

[Debugging TEAL](dapps/pyteal/smart-contracts/test-and-deploy/debugging) - Tools and methods for debugging TEAL code.




# Searching the Blockchain (Indexer V2)

[Overview](indexer) - How to use Indexer v2 functionality to efficiently query data on the blockchain.