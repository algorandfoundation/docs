title: Set up your development environment

To start building on Algorand, you need to prepare your development environment. A development environment consists of getting access to an Algorand node. You need access to a node to submit new transactions, read blockchain data, and manage wallets.

Let's take a look at the different options you have.

## 1. Sandbox

The most-used option is setting up the **[Algorand sandbox](https://github.com/algorand/sandbox)**. The sandbox allows developers to create local, private networks. Moreover, you can quickly remove a network, reset its state, or spin up a new network. The only requirement for using the Algorand sandbox is installing Docker and Docker-Compose. Additionally, you can choose to connect to a real network and catch up to the latest round.

## 2. Third-party API services

Use **third-party API services** to access native Algorand REST APIs for the mainnet, testnet, and betanet. This is an excellent choice if you don't want to set up a local network using Docker, and just want to experiment with Algorand development initially. Existing service providers are [Purestake](https://developer.purestake.io/) and [Rand Labs](https://randlabs.io/products?product=api). Bear in mind that the free tiers for both service providers come with certain limitations, like the number of requests per second (more information on both websites). You can access these services via an API key and Algorand node address.

## 3. Run your Algorand node

You can decide to **run your own [Algorand node](https://github.com/algorand/go-algorand)**, which contains the full implementation of the Algorand software. This solution is more complex to set up and is also less flexible. Unlike the Sandbox, you can't throw away a network and set up a new one when you want to. Setting up an Algorand node takes much more time than setting up a local, private network using the Sandbox.

## Recommendation?

If you only want to try out Algorand code snippets, you can get started quickly with a third-party API service. However, to continue learning about Algorand and its smart signatures/smart contracts, it's recommended to install the Algorand sandbox. Also, the sandbox comes with the [indexer software](https://developer.algorand.org/docs/rest-apis/indexer/#create-publication-overlay) which runs locally so you can quickly look up applications, assets, transactions, and accounts.

Here are steps to install the Algorand sandbox:

```sh
// clone the sandbox from GitHub
git clone https://github.com/algorand/sandbox.git

// enter the sandbox folder
cd sandbox

// run the sandbox executable to start a private network
./sandbox up
```

Starting the sandbox with the `up` command will take a couple of minutes if this is your first time running the sandbox. The script will make sure to pull all required Docker images before setting up your sandbox.

A successful node installation will print a list of prefunded accounts. Here's an example of the sandbox output.

```sh
# Available accounts
./sandbox goal account list

# [offline]	HCNMMIL3MKYILOLUO74NF6OPCJU4RU7IE5PX6JYBIT5YHAMQIVO5YADHMU	HCNMMIL3MKYILOLUO74NF6OPCJU4RU7IE5PX6JYBIT5YHAMQIVO5YADHMU	1000000000000000 microAlgos

# [offline]	3KHVQUNTXBFKPTWPPLRYZY3MZIW4EB6XYWRTTIA36O6ZSMRLSEWA2J2HTA	3KHVQUNTXBFKPTWPPLRYZY3MZIW4EB6XYWRTTIA36O6ZSMRLSEWA2J2HTA	4000000000000000 microAlgos

# [online]	5FRKKWRG3UAJQNB7QIOWBW2JICZS4YUF2WFAETHGN3CBM4R3N27NY3T2KQ	5FRKKWRG3UAJQNB7QIOWBW2JICZS4YUF2WFAETHGN3CBM4R3N27NY3T2KQ	4000000000000000 microAlgos
```

To connect to the sandbox using an SDK, you can use the below connection object.

```js
const algosdk = require('algosdk');

// create client object to connect to sandbox's algod client
const algodToken = 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa';
const algodServer = 'http://localhost';
const algodPort = 4001;
let algodClient = new algosdk.Algodv2(algodToken, algodServer, algodPort);
```
