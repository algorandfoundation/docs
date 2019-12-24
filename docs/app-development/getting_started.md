This page is a guide for developers who want to get started building on the Algorand blockchain. 

# What does it mean to build on Algorand?
Building an application on Algorand means that your application reads from or writes to the Algorand blockchain. Writing to the Algorand blockchain is synonymous with issuing a transaction that will later be confirmed within a block.

A program connects to the Algorand blockchain through an `algod` client. The `algod` client requires a valid `address` and access `token` of an Algorand node that is connected to the network you plan to interact with. 

Algorand officially supports four SDKs for developing applications on Algorand: Javascript, Python, Java, and Go. 

# Choosing a network
There are three public Algorand Networks: 

 || MainNet | TestNet | BetaNet |
:-- |:-------------:| :-------------: | :-------------: |
**Protocol Version** | Current | Current| Future
**Genesis Distribution** | Unique | Unique | Unique
**Reliability**         | Most Stable  | Very Stable (restart possible; but very infrequent) | Experimental; frequent restarts


- **MainNet**: The primary Algorand Network with real-value assets including the native Algo.  
- **TestNet**: Has a different genesis block than MainNet, but the protocol version and features are the same. 
- **BetaNet**: Where future protocols are available for public testing. As its name suggests, BetaNet is meant to be a beta-like environment. Therefore, quality and features may not be final, and protocol upgrades and network restarts may be common.

## Best Practices
1. If your application depends on features currently available on MainNet, use TestNet as your public testing network. 
2. If your application depends on features only available on BetaNet, use BetaNet as your public testing network.
3. Use private networks as needed for greater control and isolation of your development environment.

!!! tip
	Create private networks with the `goal network` commands using protocol versions from any of the available networks or if you modified source. This is highly recommended for stability if your application uses BetaNet-only features.



# How do I obtain an `address` and `token`?
There are 3 recommended ways to obtain an `address` and `token`. 

## 1. Use a third-party service

!!! info
	- Recommended if you do not need or want to manage a node.
	- Fastest way to get connected.
	- Compatible with SDKs (check with service provider) and `algod` RESTful interfaces.
	- **No** access to `goal`, `algokey`, or `kmd`.
	- **No** access to private network functions.

Sign up with a service and they will provide you with the `address` and `token` for the desired Algorand network. Use this information to instantiate an `algod` client.

Known available services: [Purestake API](https://developer.purestake.io/)

## 2. Bootstrap from S3

!!! info
	- Recommended if you eventually want to run your own node, but you also want to start developing right away.
	- Access to `goal`, `algokey`, and `kmd`.
	- Compatible with all SDKs and RESTful interfaces for `algod` and `kmd`.
	- Access to private networks.
	- Great for experimentation, but not suitable for use in production. See warning below.
  
!!! warning
	Bootstrapping from a snapshot bypasses the normal node catchup procedure that cryptographically verifies the whole history of the blockchain - a procedure that is imperative to maintaining a healthy network. Therefore, this method is *only* recommended in the context of early stage application development to avoid catchup wait times and get started quickly. It should *never* be used to run a node in production or participate in consensus. Make sure that you migrate your application to a node that has undergone full catchup prior to launching your application in production.

Follow the directions here: https://github.com/algorand/sandbox

## 3. Run your own node

!!! info
	- Gives you full control of node configuration.
	- Access to `goal`, `algokey`, and `kmd`.
	- Compatible with all SDKs and RESTful interfaces for `algod` and `kmd`.
	- Access to private networks.
	- Requires longest time to setup, as the node validates all blocks since the beginning.
	- Suitable for production.
  
This method gives you full control of your node and its configuration. Read the docs to setup and run a node [here](../node/types.md). [Configure your node](../node/config.md) according to your specific needs.

Your `algod` net `address` is located here:

```bash
$ cat $ALGORAND_DATA/algod.net
```

Your algod access `token` is located here:

```bash
$ cat $ALGORAND_DATA/algod.token
```

## Side-by-Side Comparison
 || Use a third-party service | Bootstrap with s3 | Run your own node |
:-- |:-------------:| :-------------: | :-------------: |
**Time**         | **Seconds** - Just signup| **Minutes** - Same as running a node with no catchup	| **Days** - need to wait for node to catchup
**Trust**         | 1 party 🧐       | 1 party 🧐	| Yourself 🙂
**Cost**         | Usually free for development; pay based on rate limits in production | Variable (with free option) - see [node types](../node/types.md)	| Variable (with free option) - see [node types](../node/types.md)	
**Private Networks**| ❌ | ✅ | ✅
**`goal`, `algokey`, `kmd`**| ❌ | ✅ | ✅
**Platform**|Varied|Linux|MacOS; Linux|
**Production Ready**| ✅ | ❌ | ✅

# Install SDKs
Install your desired SDK by visiting the corresponding SDK documentation: Python, JavaScript, [Go](https://github.com/algorand/go-algorand-sdk), Java. 

# Connect to `algod`

After you have your address and access token. You can instantiate an `algod` client with any of the SDKs as follows:

```Python tab=
from algod import AlgodClient

algod_address = "http://127.0.0.1:8080"
algod_token = "ef3120d25723fc3fc22c61f9ab4aa4b989f27ef0855d4f860248cddcb25977ab"

algod_client = AlgodClient(algod_token, algod_address)
```

```JavaScript tab=
const TOKEN = "ef3120d25723fc3fc22c61f9ab4aa4b989f27ef0855d4f860248cddcb25977ab";
const SERVER = "http://127.0.0.1";
const PORT = 8080;

let algodClient = new algosdk.Algod(TOKEN, SERVER, PORT);
```

```Go tab=
package main

import (
	"github.com/algorand/go-algorand-sdk/client/algod"
)

const algodAddress = "http://127.0.0.1:8080"
const algodToken = "ef3120d25723fc3fc22c61f9ab4aa4b989f27ef0855d4f860248cddcb25977ab"

func main() {
	algodClient, err := algod.MakeClient(algodAddress, algodToken)
	if err != nil {
		return
	}
}
```

```Java tab=
package main

import (
	"github.com/algorand/go-algorand-sdk/client/algod"
)

const algodAddress = "http://127.0.0.1:8080"
const algodToken = "ef3120d25723fc3fc22c61f9ab4aa4b989f27ef0855d4f860248cddcb25977ab"

func main() {
	algodClient, err := algod.MakeClient(algodAddress, algodToken)
	if err != nil {
		return
	}
}
```