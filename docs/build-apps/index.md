title: Section Overview

This section is a getting started guide for developers looking to build applications on the Algorand blockchain. Start here to learn what it means to build an application on Algorand and how to get started quickly.

# What does it mean to build on Algorand?
Building an application on Algorand means that your application, directly or indirectly, reads from or writes to the Algorand blockchain. Writing to the Algorand blockchain is synonymous with issuing a transaction that will later be confirmed within a block. Reading from the blockchain means reading back transactions that have been confirmed within prior blocks.

A program connects to the Algorand blockchain through an **algod** client. The algod client requires a valid **algod REST endpoint IP address** and **algod token** from an Algorand node that is connected to the network you plan to interact with. 

The rest of this section is organized as follows.

1. [Workspace Setup](setup.md) - Learn how to setup up your development environment, which Algorand network to use, and how to obtain credentials to connect to a blockchain node.
2. [Connect to Node](connect.md) - Connect to a node through your preferred SDK.
3. [Your First Transaction](hello_world.md) - Step-by-step guide to send your first transaction on TestNet.