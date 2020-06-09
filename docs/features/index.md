title: Section Index

The **Explore Features** section contains explanations and how-tos for all Algorand features. Code examples are embedded throughout in each of the [SDKs](../reference/index.md#sdks) and [command line interface tools](../reference/index.md#command-line-interface-tools-cli-tools).

# Accounts

[Overview](./accounts/index.md) - Learn all about Algorand accounts, keys, wallets, and addresses.

[Creation Methods](./accounts/create.md) -How to create standalone accounts, wallet-derived (kmd) accounts, or multisignature accounts.

# Transactions

[Structure](./transactions/index.md) - A look at the underlying structure of the different types of transactions.

[Authorization](./transactions/signatures.md) - How to authorize transactions with Single Signatures, Multisignatures, and Logic Signatures.

[Offline Authorization](./transactions/offline_transactions.md) - How to authorize transactions _without_ requiring an internet connection.

# Assets

[Algorand Standard Assets Overview](./asa.md) - Learn all about assets and how they are implemented on Algorand at layer-1.

[Asset Parameters](./asa.md#asset-parameters) - The mutable and immutable properties of an asset.

[Creating an Asset](./asa.md#creating-an-asset) - [Modifying an Asset](./asa.md#modifying-an-asset) - [Receiving an Asset](./asa.md#receiving-an-asset) - [Transferring an Asset](./asa.md#transferring-an-asset) - [Freezing an Asset](./asa.md#freezing-an-asset) - [Revoking an Asset](./asa.md#revoking-an-asset) - [Destroying an Asset](./asa.md#destroying-an-asset)

[Retrieving Asset Information](./asa.md#retrieve-asset-information) - How to retrieve information about an asset.

# Atomic Transfers

[Overview](./atomic_transfers.md) - What atomic transfers are and how they are they implemented on Algorand.

[Step-by-Step Guide](./atomic_transfers.md#step-by-step-guide) - Walkthrough an atomic transfer in the SDKs and goal.

# Smart Contracts

[Introduction](./asc1/index.md) - An overview of stateless Algorand Smart Contracts (ASC1)

[Logic Signatures](./asc1/modes.md#logic-signatures) - The authorization method for transactions that use Smart Contracts.

[Modes of Use - Contract Account](./asc1/modes.md#contract-account) - How to use a Smart Contract as its own Algorand account.

[Modes of Use - Delegated Approval](./asc1/modes.md#delegated-approval) - How to use a Smart Contract to authorize spending from an existing account.

[The Smart Contract Language](./asc1/teal_overview.md) - Learn all about TEAL, the language of Smart Contracts.

[A Contract Walkthrough](./asc1/goal_teal_walkthrough.md) - A simple command-line walkthrough of creating a simple smart contract.

[Using the SDKS](./asc1/sdks.md) - Smart Contract functionality from within the SDKs.


# Searching the Blockchain 🆕

[Overview](./indexer.md) - How to use Indexer v2 functionality to efficiently query data on the blockchain.