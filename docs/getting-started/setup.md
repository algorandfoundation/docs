PAGE STATUS: IN PROGRESS

This section is a getting started guide for developers who want to build applications on the Algorand blockchain. 

# What does it mean to build on Algorand?
Building an application on Algorand means that your application reads from or writes to the Algorand blockchain. Writing to the Algorand blockchain is synonymous with issuing a transaction that will later be confirmed within a block.

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
2. [Bootstrap from s3](#2-bootstrap-from-s3)
3. [Run your own node](#3-run-your-own-node)

## 1. Use a third-party service

!!! summary
	:thumbsup:
	Recommended if you plan to use _just_ the SDKs or algod RESTful interfaces, and want to get connected as fast as possible.

A third-party service runs a node and provides access to that node through their own API keys. On signup, the service provides you with an algod address and an API key which will replace your algod token.

Known available services: [Purestake API](https://developer.purestake.io/)

## 2. Bootstrap from S3

!!! summary
	:thumbsup:
	Recommended if you want access to all developer tools including `goal`, `kmd`, and `algokey`, but can't wait days for your node to catchup.


Follow the directions here: https://github.com/algorand/sandbox

!!! warning
	Bootstrapping from a snapshot bypasses the normal node catchup procedure that cryptographically verifies the whole history of the blockchain - a procedure that is imperative to maintaining a healthy network. Therefore, this method is *only* recommended in the context of early stage application development to avoid catchup wait times and get started quickly. It should *never* be used to run a node in production or participate in consensus. Make sure that you migrate your application to a node that has undergone full catchup prior to launching your application in production.

## 3. Run your own node

!!! summary
	:thumbsup:
	Recommended if you want access to all developer tools including `goal`, `kmd`, and `algokey`, and want to setup a production-ready environment. This should be follow-on to option 2 prior to launching an application on MainNet.

  
This method gives you full control of your node and its configuration. Read the docs to setup and run a node [here](../node/types.md). [Configure your node](../node/config.md) according to your specific needs.

After setup, find your **algod address** here:

```bash
$ cat $ALGORAND_DATA/algod.net
```

and your **algod token** here:

```bash
$ cat $ALGORAND_DATA/algod.token
```

<center>*Side-by-Side Comparison*</center>

 || Use a third-party service | Bootstrap with s3 | Run your own node |
:-- |:-------------:| :-------------: | :-------------: |
**Time**         | **Seconds** - Just signup| **Minutes** - Same as running a node with no catchup	| **Days** - need to wait for node to catchup
**Trust**         | 1 party       | 1 party	| Yourself 
**Cost**         | Usually free for development; pay based on rate limits in production | Variable (with free option) - see [node types](../node/types.md)	| Variable (with free option) - see [node types](../node/types.md)	
**Private Networks**| ❌ | ✅ | ✅
**`goal`, `algokey`, `kmd`**| ❌ | ✅ | ✅
**Platform**|Varied|MacOS; Linux|MacOS; Linux|
**Production Ready**| ✅ | ❌ | ✅

# Install your preferred SDK
Install your preferred SDK by following the setup instructions on the corresponding SDK documentation page.

[Python](../reference/sdks/python.md)

[JavaScript](../reference/sdks/javascript.md) 

[Go](../reference/sdks/go.md)

[Java](../reference/sdks/java.md)

# Other Setup Tips

If you use `goal`, it is recommended that you set the `ALGORAND_DATA` environment variable to avoid the need to specify it for each `goal` command. It is also recommended that you place `goal`, `kmd`, and `algokey` within your executable path. The examples in these docs will assume you have done both of these.

So you will see:

```bash
$ goal node status 
```

instead of:

```bash

$ goal node status -d your-node-directory <PLACEHOLDER>
```


# Connect to algod

After you have your address and access token. Instantiate an **algod** client with any of the SDKs as follows:

```JavaScript tab=
const TOKEN = "algod-address";
const SERVER = "algod-address";
const PORT = algod-port // <PLACEHOLDER>;

let algodClient = new algosdk.Algod(TOKEN, SERVER, PORT);
```

```Python tab=
from algosdk import algod

algod_address = "algod-address" <PLACEHOLDER>
algod_token = "algod-token" <PLACEHOLDER>

algod_client = algod.AlgodClient(algod_token, algod_address)
```

```Java tab=
public class GettingStartedExample 
{
    public static void main(String args[]) throws Exception {
        
        final String ALGOD_API_ADDR = "algod-address"<PLACEHOLDER>;
        final String ALGOD_API_TOKEN = "algod-token"<PLACEHOLDER>;

        //Create an instance of the algod API client
        AlgodClient client = (AlgodClient) new AlgodClient().setBasePath(ALGOD_API_ADDR);
        ApiKeyAuth api_key = (ApiKeyAuth) client.getAuthentication("api_key");
        api_key.setApiKey(ALGOD_API_TOKEN);
        AlgodApi algodApiInstance = new AlgodApi(client); 
		...
	}
}
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

!!! note	
	If you are using a third-party service, use the API key header to instantiate a client instead.

	```JavaScript tab=
	const TOKEN = "ef3120d25723fc3fc22c61f9ab4aa4b989f27ef0855d4f860248cddcb25977ab";
	const SERVER = "http://127.0.0.1";
	const PORT = 8080;

	let algodClient = new algosdk.Algod(TOKEN, SERVER, PORT);
	```

	```Python tab=
	from algodsdk import algod

	algod_address = "https://testnet-algorand.api.purestake.io/ps1"
	algod_token = ""
	headers = {
   		"X-API-Key": API-key,
	}
	```

	```Java tab=
	import com.algorand.algosdk.algod.client.AlgodClient;
	import com.algorand.algosdk.algod.client.api.AlgodApi;

	public class GettingStartedExample {
		public static void main(String[] args) {

			final String ALGOD_API_ADDR = "algod-address"<PLACEHOLDER>;
			final String ALGOD_API_KEY = "service-api-key"<PLACEHOLDER>;

			AlgodClient client = new AlgodClient();
			client.setBasePath(ALGOD_API_ADDR);
			client.addDefaultHeader("X-API-Key", ALGOD_API_KEY);
			AlgodApi algodApiInstance = new AlgodApi(client);
			...
		}
	}
	```

	```Go tab=
	import (
		"github.com/algorand/go-algorand-sdk/client/algod"
	)
	const algodAddress = "algod-address"<PLACEHOLDER>
	const apiKey = "your-api-key"<PLACEHOLDER>

	func main() {
		var headers []*algod.Header
		headers = append(headers, &algod.Header{"X-API-Key", apiKey})
		algodClient, err := algod.MakeClientWithHeaders(algodAddress, "", headers)
		...
	}
	```



# Check node status and network version

Call the status and version methods on your algod client to validate the details of your connection.

```javascript tab="JavaScript"
```

```python tab="Python"
```

```java tab="Java"
	...
        try {
            NodeStatus status = algodApiInstance.getStatus();
            System.out.println("Algorand network status: " + status);
            Version version = algodApiInstance.getVersion();
            System.out.println("Algorand network version: " + version);
        } catch (ApiException e) {
            System.err.println("Exception when calling algod#getStatus or #getVersion");
            e.printStackTrace();
        }
	...
```

```bash tab="cURL"
curl -i -X GET \
   -H "X-Algo-API-Token:algod-token<PLACEHOLDER>" \
 'http://algod-address<PLACEHOLDER>:algod-port<PLACEHOLDER>/v1/status'
```

```bash tab="goal" hl_lines="2 3 4 5 6 7 8 9"
$ goal node status
Last committed block: [LATEST_ROUND]
Time since last block: [TIME_IN_SECONDS]
Sync Time: [TIME_IN_SECONDS]
Last consensus protocol: [LINK_TO_CURRENT_PROTOCOL_SPEC]
Next consensus protocol: [LINK_TO_FUTURE_PROTOCOL_SPEC]
Round for next consensus protocol: [ROUND_FOR_FUTURE_PROTOCOL]
Next consensus protocol supported: [true|false]
Has Synced Since Startup: [true|false]
Genesis ID: [GENESIS_ID]
Genesis hash: [BASE64_GENESIS_HASH]
```

The _status_ methods return information such as the latest round<LINK TO GLOSSARY>, referred to as `lastRound` from the perspective of the node you are connected to. Each of the SDKs may differ slightly in which information they return for each call so below is the full REST response with the latest round highlighted.

```json hl_lines="2"
{
    "lastRound": 4243027,
    "lastConsensusVersion": "https://github.com/algorandfoundation/specs/tree/4a9db6a25595c6fd097cf9cc137cc83027787eaa",
    "nextConsensusVersion": "https://github.com/algorandfoundation/specs/tree/4a9db6a25595c6fd097cf9cc137cc83027787eaa",
    "nextConsensusVersionRound": 4243028,
    "nextConsensusVersionSupported": true,
    "timeSinceLastRound": 4261519666,
    "catchupTime": 0,
    "hasSyncedSinceStartup": false
}

```

Check if the node is caught up by validating against others running nodes, like a public block explorer<LINK TO COMMUNITY PROJECTS>. As a secondary check, see if your `catchupTime` is 0 and your rounds are progressing at a rate of less than 5 seconds on average. This is the time it takes to confirm a block on Algorand. Note that the `timeSinceLastRound` is represented in nanoseconds.

!!! warning
	If your node is out-of-sync with the rest of the network you cannot send transactions and account balances will be out of date. 

The _version_ methods return information about the identity of the network and the current software build. 


```javascript tab="JavaScript"
```

```python tab="Python"
```

```java tab="Java"
	...
        try {
            NodeStatus status = algodApiInstance.getStatus();
            System.out.println("Algorand network status: " + status);
            Version version = algodApiInstance.getVersion();
            System.out.println("Algorand network version: " + version);
        } catch (ApiException e) {
            System.err.println("Exception when calling algod#getStatus or #getVersion");
            e.printStackTrace();
        }
	...
```

```bash tab="cURL"
curl -i -X GET \
   -H "X-Algo-API-Token:algod-token<PLACEHOLDER>" \
 'http://algod-address<PLACEHOLDER>:algod-port<PLACEHOLDER>/versions'
```

```bash tab="goal" hl_lines="10 11"
$ goal node status
Last committed block: [LATEST_ROUND]
Time since last block: [TIME_IN_SECONDS]
Sync Time: [TIME_IN_SECONDS]
Last consensus protocol: [LINK_TO_CURRENT_PROTOCOL_SPEC]
Next consensus protocol: [LINK_TO_FUTURE_PROTOCOL_SPEC]
Round for next consensus protocol: [ROUND_FOR_FUTURE_PROTOCOL]
Next consensus protocol supported: [true|false]
Has Synced Since Startup: [true|false]
Genesis ID: [GENESIS_ID]
Genesis hash: [BASE64_GENESIS_HASH]
```

Check that the `genesis_id` and the `genesis_hash_b64`, as shown in the REST response below, match your chosen network before proceeding.

```json hl_lines="5 6"
{
    "versions": [
        "v1"
    ],
    "genesis_id": "testnet-v1.0",
    "genesis_hash_b64": "SGO1GKSzyE7IEPItTxCByw9x8FmnrCDexi9/cOUJOiI=",
    "build": {
        "major": 2,
        "minor": 0,
        "build_number": 2,
        "commit_hash": "1bc1d4c9",
        "branch": "rel/stable",
        "channel": "stable"
    }
}

```




