title: Set up your development environment

To start building on Algorand, you need to prepare your development environment. A development environment requires getting access to an Algorand node. You need access to a node to submit new transactions, read blockchain data, and manage wallets.

Let's take a look at the different options you have.

## 1. Sandbox

The most-used option is setting up the **[Algorand sandbox](https://github.com/algorand/sandbox)**. The sandbox allows developers to create local, private networks. Moreover, you can quickly remove a network, reset its state, or spin up a new network. The only requirement for using the Algorand sandbox is installing Docker and Docker-Compose. Additionally, you can choose to connect to a real network and the sandbox will take care of catching up to the latest round.

## 2. Third-party API services

Use **third-party API services** to access native Algorand REST APIs for the mainnet, testnet, and betanet. This is an excellent choice if you don't want to set up a local network using Docker, and just want to experiment with Algorand development initially. Existing service providers are [Purestake](https://developer.purestake.io/), [Bloq](https://www.bloq.com/products/platform/bloq-connect/), [Rand Labs](https://randlabs.io/products?product=api), and [AlgoNode.io](https://algonode.io/). 
A list of third-party API services is provided on the [Ecosystem Tools & Projects page](https://developer.algorand.org/ecosystem-projects/?tags=api-services).

Bear in mind that the free tiers for both service providers come with certain limitations, like the number of requests per second (more information on both websites). You can access these services via an API key and Algorand node address.

## 3. Run your Algorand node

You can decide to **[run your own Algorand node](../../run-a-node/setup/install.md)**, which contains the full implementation of the Algorand software. This solution is more complex to set up and is also less flexible. Unlike the sandbox, you can't throw away a network and set up a new one when you want to. Setting up an Algorand node takes much more time than setting up a local, private network using the sandbox.

## Recommendation?

If you only want to try out Algorand code snippets, you can get started quickly with a third-party API service. 

However, to continue learning about Algorand and its smart signatures/smart contracts, [installing the Algorand sandbox](sandbox.md) is recommended. Also, in local private networks, the sandbox comes with the [indexer software](https://developer.algorand.org/docs/rest-apis/indexer/#create-publication-overlay), which runs locally so you can quickly lookup applications, assets, transactions, and accounts.

!!! info
    Even if you do not want to run your own node, it is recommended to [install the Algorand software](../../run-a-node/setup/install.md), as it provides useful command line tools such as [`goal`](../../../clis/goal/goal), [`algokey`](../../../clis/algokey/algokey), and  `msgpacktool`.