title: 2. Connect to Node

The **algod IP address** and access **token** you obtained in the [Workspace Setup](./setup.md) section gives your application the credentials to interface with the Algorand blockchain. Interfacing with the Algorand blockchain using the SDKs is accomplished through an **algod client**. If using the algod REST API directly, these credentials can be supplied with each request.

_Read more about the [node's algod process](../reference-docs/node/artifacts.md#algod)._

# Create an algod client
Instantiate an **algod** client with your preferred SDK. 

```JavaScript tab=
const algosdk = require('algosdk');

async function gettingStartedExample() {

	const server = <algod-address>;
	const port = <port-number>;
	const token = <algod-token>;

	let algodClient = new algosdk.Algod(token, server, port);
	...
}
```

```Python tab=
from algosdk import algod

algod_address = <algod-address>
algod_token = <algod-token>

algod_client = algod.AlgodClient(algod_token, algod_address)
```

```Java tab=
import com.algorand.algosdk.algod.client.AlgodClient;
import com.algorand.algosdk.algod.client.api.AlgodApi;
import com.algorand.algosdk.algod.client.auth.ApiKeyAuth;

public class ConnectToNetwork {
    public static void main(String args[]) throws Exception {
        
        final String ALGOD_API_ADDR = <algod-address>;
        final String ALGOD_API_TOKEN = <algod-token>;

        //Create an instance of the algod API client
        AlgodClient client = (AlgodClient) new AlgodClient()
		client.setBasePath(ALGOD_API_ADDR);
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

const algodAddress = <algod-address>
const algodToken = <algod-token>

func main() {
	algodClient, err := algod.MakeClient(algodAddress, algodToken)
	if err != nil {
		return
	}
}
```

If you are using a third-party service, use the API key header instead when instantiating an algod client.

```JavaScript tab=
const algosdk = require("algosdk");

async function gettingStartedExample() {

	const server = <algod-address>;
	const port = "";
	const token = {
		'X-API-Key': <service-api-key>
	};

	let algodClient = new algosdk.Algod(token, server, port);
	...
}

```

```Python tab=
from algodsdk import algod

algod_address = <algod-address>
algod_token = ""
headers = {
   	"X-API-Key": <service-api-key>,
}

algod_client = algod.AlgodClient(algod_token, algod_address, headers)
```

```Java tab=
import com.algorand.algosdk.algod.client.AlgodClient;
import com.algorand.algosdk.algod.client.api.AlgodApi;
import com.algorand.algosdk.algod.client.auth.ApiKeyAuth;

public class ConnectToNetwork {
	public static void main(String[] args) {

		final String ALGOD_API_ADDR = <algod-address>;
		final String ALGOD_API_KEY = <service-api-key>;

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
const algodAddress = <algod-address>
const apiKey = <your-api-key>

func main() {
	var headers []*algod.Header
	headers = append(headers, &algod.Header{"X-API-Key", apiKey})
	algodClient, err := algod.MakeClientWithHeaders(algodAddress, "", headers)
	...
}
```



# Check node status and network version

Call the _status_ and _version_ methods from the algod client to check the details of your connection. This information is also available through equivalent REST API calls and `goal` commands.

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
   -H "X-Algo-API-Token:<algod-token>" \
 'http://<algod-address>:<algod-port>/v1/status'
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
   -H "X-Algo-API-Token:<algod-token>" \
 'http://<algod-address>:<algod-port>/versions'
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
		algod_address = <algod-address>
		algod_token = <algod-token>
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

	```java tab="Java"
	import com.algorand.algosdk.algod.client.AlgodClient;
	import com.algorand.algosdk.algod.client.api.AlgodApi;
	import com.algorand.algosdk.algod.client.auth.ApiKeyAuth;
	import com.algorand.algosdk.algod.client.model.*;

	public class ConnectToNetwork { 
		public static void main(String args[]) throws Exception {
			
			final String ALGOD_API_ADDR = <algod-address>;
			final String ALGOD_API_TOKEN = <algod-token>;

			//Create an instance of the algod API client
			AlgodClient client = (AlgodClient) new AlgodClient()
			client.setBasePath(ALGOD_API_ADDR);
			ApiKeyAuth api_key = (ApiKeyAuth) client.getAuthentication("api_key");
			api_key.setApiKey(ALGOD_API_TOKEN);
			AlgodApi algodApiInstance = new AlgodApi(client); 
			try {
				NodeStatus status = algodApiInstance.getStatus();
				Version version = algodApiInstance.getVersion();
				System.out.println("Algorand network status: " + status);
				System.out.println("Algorand network version: " + version);
			} catch (ApiException e) {
				System.err.println("Exception when calling algod#getStatus or algod#getVersion");
				e.printStackTrace();
			}
		}
	}
	```

??? example "Complete Example - Connect to the Network with API Service"

	```python tab="Python"
	import json
	from algosdk import algod

	def main():
		algod_address = <algod-address>
		algod_token = ""
		headers = {
			'X-API-Key': <service-api-key>
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

	```java tab="Java"
	package example;

	import java.math.BigInteger;

	import java.util.concurrent.TimeUnit;

	import com.algorand.algosdk.account.Account;
	import com.algorand.algosdk.algod.client.AlgodClient;
	import com.algorand.algosdk.algod.client.ApiException;
	import com.algorand.algosdk.algod.client.api.AlgodApi;
	import com.algorand.algosdk.algod.client.auth.ApiKeyAuth;
	import com.algorand.algosdk.algod.client.model.*;

	public class ConnectToNetwork {

		public static void main(String args[]) throws Exception {

			final String ALGOD_API_ADDR = <algod-address>;
			final String ALGOD_API_KEY = <service-api-key>;

			AlgodClient client = new AlgodClient();
			client.setBasePath(ALGOD_API_ADDR);
			client.addDefaultHeader("X-API-Key", ALGOD_API_KEY);
			AlgodApi algodApiInstance = new AlgodApi(client);
			try {
				NodeStatus status = algodApiInstance.getStatus();
				Version version = algodApiInstance.getVersion();
				System.out.println("Algorand network status: " + status);
				System.out.println("Algorand network version: " + version);
			} catch (ApiException e) {
				System.err.println("Exception when calling algod#getStatus or algod#getVersion");
				e.printStackTrace();
			}
		}
	}
	```


