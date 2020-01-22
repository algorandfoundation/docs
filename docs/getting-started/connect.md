title: Connect to the Network

After you have you an algod address and access token, instantiate an **algod** client with your preferred SDK.

```JavaScript tab=
const algosdk = require('algosdk');

async function gettingStartedExample() {

	const server = algod-address<PLACEHOLDER>;
	const port = port-number<PLACEHOLDER>;
	const token = algod-token<PLACEHOLDER>;

	let algodClient = new algosdk.Algod(token, server, port);
	...
}
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

!!! info	
	If you are using a third-party service, use the API key header instead when instantiating an algod client.

	```JavaScript tab=
	const algosdk = require("algosdk");

	async function gettingStartedExample() {

		const server = algod-address<PLACEHOLDER>;
		const port = "";
		const token = {
			'X-API-Key': service-api-key<PLACEHOLDER>
		};

		let algodClient = new algosdk.Algod(token, server, port);
		...
	}

	```

	```Python tab=
	from algodsdk import algod

	algod_address = "algod-address"<PLACEHOLDER>
	algod_token = ""
	headers = {
   		"X-API-Key": service-api-key<PLACEHOLDER>,
	}

	algod_client = algod.AlgodClient(algod_token, algod_address, headers)
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

Call the _status_ and _version_ methods fron the algod client to check the details of your connection. This information is also available through equivalent REST API calls and `goal` commands.

```javascript tab="JavaScript"
...
	let status = await algodClient.status();
	console.log("Algorand network status: %o", status);
...
```

```python tab="Python"
...
	try:
		status = algod_client.status()
		print(json.dumps(status, indent=4))
	except Exception as e:
		print(e)
...
```

```java tab="Java"
	...
        try {
            NodeStatus status = algodApiInstance.getStatus();
            System.out.println("Algorand network status: " + status);
        } catch (ApiException e) {
            System.err.println("Exception when calling algod#getStatus");
            e.printStackTrace();
        }
	...
```

```go tab="Go"
...
	status, err := algodClient.Status()
	if err != nil {
		fmt.Printf("Error getting status: %s\n", err)
		return
	}
	statusJSON, err := json.MarshalIndent(status, "", "\t")
	if err != nil {
		fmt.Printf("Can not marshall status data: %s\n", err)
	}
	fmt.Printf("%s\n", statusJSON)
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

The _status_ methods return information about the status of the node, such as the latest round<LINK TO GLOSSARY>, referred to as `lastRound`, from the perspective of the node you are connected to. Each of the SDKs may differ slightly in which information they return for each call. Shown below is the response from the REST API call.

```json tab="Response"
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

Check if the node is caught up by validating against others running nodes, like a [public block explorer](../community.md#block-explorers). As a secondary check, see if your `catchupTime` is 0 and your rounds are progressing at a rate of less than 5 seconds on average. This is the time it takes to confirm a block on Algorand. Note that the `timeSinceLastRound` is represented in nanoseconds.

!!! warning
	If your node is out-of-sync with the rest of the network you cannot send transactions and account balances will be out-of-date. 

The _version_ methods return information about the identity of the network and the current software build. 


```javascript tab="JavaScript"
...
	let version = await algodClient.versions();
	console.log("Algorand protocol version: %o", version)
...
```

```python tab="Python"
...
	try:
		versions = algod_client.versions()
		print(json.dumps(versions, indent=4))
	except Exception as e:
		print(e)
...
```

```java tab="Java"
	...
        try {
            Version version = algodApiInstance.getVersion();
            System.out.println("Algorand network version: " + version);
        } catch (ApiException e) {
            System.err.println("Exception when calling algod#getVersion");
            e.printStackTrace();
        }
	...
```

```go tab="Go"
...
	version, err := algodClient.Versions()
	if err != nil {
		fmt.Printf("Error getting versions: %s\n", err)
		return
	}
	versionJSON, err := json.MarshalIndent(version, "", "\t")
	if err != nil {
		fmt.Printf("Can not marshall version data: %s\n", err)
	}
	fmt.Printf("%s\n", versionJSON)
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

```json hl_lines="5 6" tab="Response"
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

??? example "Complete Example - Connect to the Network"

	```python tab="Python"
	import json
	from algosdk import algod

	def main():
		algod_address = algod-address<PLACEHOLDER>
		algod_token = algod-token<PLACEHOLDER>
		algod_client = algod.AlgodClient(algod_token, algod_address)

		try:
			status = algod_client.status()
			versions = algod_client.versions()
			print(json.dumps(status, indent=4))
			print(json.dumps(versions, indent=4))
		except Exception as e:
			print(e)
	main()
	```

??? example "Complete Example - Connect to the Network with API Service"

	```python tab="Python"
	import json
	from algosdk import algod

	def main():
		algod_address = algod-address<PLACEHOLDER>
		algod_token = ""
		headers = {
			'X-API-Key': service-api-key<PLACEHOLDER>
		}
		algod_client = algod.AlgodClient(algod_token, algod_address, headers)

		try:
			status = algod_client.status()
			versions = algod_client.versions()
			print(json.dumps(status, indent=4))
			print(json.dumps(versions, indent=4))
		except Exception as e:
			print(e)
	main()
	```


