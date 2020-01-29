title: 1. Workspace Setup

This section is a getting started guide for developers, looking to build applications on the Algorand blockchain. Start here to learn what it means to build an application on Algorand and how to get started quickly.

# What does it mean to build on Algorand?
Building an application on Algorand means that your application, directly or indirectly, reads from or writes to the Algorand blockchain. Writing to the Algorand blockchain is synonymous with issuing a transaction that will later be confirmed within a block. Reading from the blockchain means reading back transactions that have been confirmed within prior blocks.

A program connects to the Algorand blockchain through an **algod** client. The algod client requires a valid algod **address** and access **token** of an Algorand node that is connected to the network you plan to interact with. 

# Available tools
Algorand officially supports four SDKs for developing applications on Algorand: Javascript<LINK>, Java, Python, and Go. 

There are also three command-line utilities packaged with Algorand node software: `goal`, `kmd`, and `algokey`.

`goal` is the primary tool for operating a node and it also contains functionality to manage keys, sign and send transactions, create assets, and perform many of the same or similar functions that are available in the SDKs. Although not required to build an application, developers who run nodes may find it useful to achieve some level of fluency in `goal` as a complementary tool during testing and validation. `goal` _is_ required to setup more advanced testing environments using private networks.

`kmd` is the CLI for the Algorand Key Management daemon and `algokey` is a standalone utility for generating Algorand accounts and for signing transactions. It is often used as a lightweight offline client for secure key signing. These two tools are not essential for getting started, so details of their use are described elsewhere.

There are also REST APIs available for both **algod** and **kmd** processes.

# Choosing a network
There are three **public** Algorand Networks paired with the functionality to create **private** networks using any protocol version. 

**MainNet** <LINK> is the primary Algorand Network with real-value assets, including Algorand's native currency - the Algo. **TestNet** <LINK> mirrors MainNet in terms of its protocol (i.e. software) version, but it has test Algos, available via a faucet, and a different genesis block, which means that the state of accounts and distribution of funds is different.  **BetaNet** is where new protocol-level features will be released for initial testing. Therefore, quality and features may not be final, and protocol upgrades and network restarts are common. <LINK>

 
## Recommended Use
If your application depends on features currently available on MainNet, use TestNet as your public testing network. If your application depends on features only available on BetaNet, use BetaNet as your public testing network. In all cases, use private networks, as needed, for greater control and isolation of your development environment.

If you are not sure which network to start with, TestNet is usually a good option as it allows you to develop against live features without risking real assets. Switching networks later will be trivial.

<center>*Side-by-side Network Comparison*</center>

 || **MainNet** | **TestNet** | **BetaNet** | 
:-- |:-------------:| :-------------: | :-------------: | 
**Protocol Version** | Current | Current| Future | Any |
**Genesis Distribution** | Unique | Unique | Unique | Any |
**Algo Accessibility** | For sale | Free from faucet | Free from faucet | 
**Network Reliability**         | Most Stable  | Very Stable, but restarts are possible | Experimental; frequent restarts | 

# How do I obtain an algod address and token?
There are three recommended ways to obtain an algod **address** and **token** that each have their respective pros and cons depending on your goals. Below is an explanation of each followed by a side-by-side comparison.

1. [Use a third-party service](#1-use-a-third-party-service)
2. [Use Docker Sandbox](#2-bootstrap-from-s3)
3. [Run your own node](#3-run-your-own-node)

## 1. Use a third-party service

This method is recommended if you plan to use _just_ the SDKs or algod RESTful interfaces, and want to get connected as fast as possible.

A third-party service runs a node and provides access to that node through their own API keys. On signup, the service provides you with an algod address and an API key which will replace your algod token.

Known available services: [Purestake API](https://developer.purestake.io/)

## 2. Use Docker Sandbox

This method is recommended if you want access to all developer tools including `goal`, `kmd`, and `algokey`, but can't wait days for your node to catchup.


Follow the directions here: https://github.com/algorand/sandbox

!!! warning
	Bootstrapping from a snapshot bypasses the normal node catchup procedure that cryptographically verifies the whole history of the blockchain - a procedure that is imperative to maintaining a healthy network. Therefore, this method is *only* recommended in the context of early stage application development to avoid catchup wait times and get started quickly. It should *never* be used to run a node in production or participate in consensus. Make sure that you migrate your application to a node that has undergone full catchup prior to launching your application in production.

## 3. Run your own node

This method is recommended if you want access to all developer tools including `goal`, `kmd`, and `algokey`, and want to setup a production-ready environment. This should be follow-on to option 2 prior to launching an application on MainNet.

  
This method gives you full control of your node and its configuration. Read the docs to setup and run a node [here](../network-participation/run-a-node/types.md). [Configure your node](../network-participation/run-a-node/config.md) according to your specific needs.

After setup, find your **REST endpoint's IP address** here:

```bash
$ cat $ALGORAND_DATA/algod.net
```

and your **algod token** (required by your application to authenticate against `algod`) here:

```bash
$ cat $ALGORAND_DATA/algod.token
```

<center>*Side-by-Side Comparison*</center>

 || Use a third-party service | Bootstrap with s3 | Run your own node |
:-- |:-------------:| :-------------: | :-------------: |
**Time**         | **Seconds** - Just signup| **Minutes** - Same as running a node with no catchup	| **Days** - need to wait for node to catchup
**Trust**         | 1 party       | 1 party	| Yourself 
**Cost**         | Usually free for development; pay based on rate limits in production | Variable (with free option) - see [node types](../network-participation/run-a-node/types.md)	| Variable (with free option) - see [node types](../network-participation/run-a-node/types.md)	
**Private Networks**| ❌ | ✅ | ✅
**`goal`, `algokey`, `kmd`**| ❌ | ✅ | ✅
**Platform**|Varied|MacOS; Linux|MacOS; Linux|
**Production Ready**| ✅ | ❌ | ✅

# Install your preferred SDK
Install your preferred SDK by following the setup instructions in the [SDK reference docs](../reference-docs/sdks.md).

# Other Setup Tips

If you use `goal`, it is recommended that you set the `ALGORAND_DATA` environment variable to avoid the need to specify it for each `goal` command. It is also recommended that you place `goal`, `kmd`, and `algokey` within your executable path. The examples in these docs will assume you have done both of these.

So you will use:

```bash
$ goal node status 
```

instead of:

```bash

$ goal node status -d your-node-directory <PLACEHOLDER>
```


