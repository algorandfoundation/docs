title: 2. Connect to Node

The **algod IP address** and access **token** you obtained in the [Workspace Setup](./setup.md) section gives your application the credentials to interface with the Algorand blockchain. Interfacing with the Algorand blockchain using the SDKs is accomplished through an **algod client**. If using the algod REST API directly, these credentials can be supplied with each request.

_Read more about the [node's algod process](../reference/node/artifacts.md#algod)._

!!! info
    The examples in this section have been updated to the v2 API, which was launched to MainNet on June 16, 2020. Visit the [v2 Migration Guide](../reference/sdks/migration.md) for information on how to migrate your code from v1. Access archived v1 examples [here](https://github.com/algorand/docs/tree/master/examples/start-building). 

# Create an algod client
Instantiate an **algod** client with your preferred SDK. 

```JavaScript tab=
const algosdk = require('algosdk');

const algodToken = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
const algodServer = "http://localhost";
const algodPort = 4001;

let algodClient = new algosdk.Algodv2(algodToken, algodServer, algodPort);
```

```Python tab=
from algosdk.v2client import algod

algod_address = "http://localhost:4001"
algod_token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
algod_client = algod.AlgodClient(algod_token, algod_address)
```

```Java tab=
import com.algorand.algosdk.v2.client.common.AlgodClient;
import com.algorand.algosdk.v2.client.common.Client;

final String ALGOD_API_ADDR = "localhost";
final Integer ALGOD_PORT = 4001;
final String ALGOD_API_TOKEN = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";

AlgodClient client = (AlgodClient) new AlgodClient(ALGOD_API_ADDR, ALGOD_PORT, ALGOD_API_TOKEN);
```

```Go tab=
package main

import (
    "github.com/algorand/go-algorand-sdk/client/v2/algod" 
)

const algodAddress = "http://localhost:4001"
const algodToken = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

func main() {
    algodClient, err := algod.MakeClient(algodAddress, algodToken)
    if err != nil {
        fmt.Printf("Issue with creating algod client: %s\n", err)
        return
    }
}
```

If you are using a third-party service, use the API key header instead when instantiating an algod client.

```JavaScript tab=
from algosdk.v2client import algod

const algodServer = "https://api.host.com";
const port = "";
const token = {
	'X-API-Key': "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
};
let algodClient = new algosdk.Algodv2(algodToken, algodServer, algodPort);
```

```Python tab=
from algosdk.v2client import algod

algod_address = "https://api.host.com"
algod_token = ""
headers = {
   	"X-API-Key": "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
}

algod_client = algod.AlgodClient(algod_token, algod_address, headers)
```

```Java tab=
import com.algorand.algosdk.v2.client.common.AlgodClient;
import com.algorand.algosdk.v2.client.common.Client;

final String ALGOD_API_ADDR = "https://api.host.com";
final String ALGOD_API_KEY = "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb";

AlgodClient client = new AlgodClient();
client.setBasePath(ALGOD_API_ADDR);
client.addDefaultHeader("X-API-Key", ALGOD_API_KEY);
AlgodApi algodApiInstance = new AlgodApi(client);
```

```Go tab=
package main

import (
    "github.com/algorand/go-algorand-sdk/client/v2/algod" 
)

const algodAddress = "https://api.host.com"
const apiKey = "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"

func main() {
	var headers []*algod.Header
	headers = append(headers, &algod.Header{"X-API-Key", apiKey})

    algodClient, err := algod.MakeClientWithHeaders(algodAddress, "", headers)
    if err != nil {
        fmt.Printf("Issue with creating algod client: %s\n", err)
        return
    }
}
```

# Check node status

Call the _status_ method from the algod client to check the details of your connection. This information is also available through equivalent REST API calls and `goal` commands.

```javascript tab="JavaScript"
let status = (await algodclient.status().do());
console.log("Algorand network status: %o", status);
```

```python tab="Python"
status = algod_client.status()
print(json.dumps(status, indent=4))
```

```java tab="Java"
try {
	NodeStatus status = algodApiInstance.getStatus();
	System.out.println("Algorand network status: " + status);
} catch (ApiException e) {
	System.err.println("Exception when calling algod#getStatus");
	e.printStackTrace();
}
```

```go tab="Go"
	status, err := algodClient.Status().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting status: %s\n", err)
		return
	}
	statusJSON, err := json.MarshalIndent(status, "", "\t")
	if err != nil {
		fmt.Printf("Can not marshall status data: %s\n", err)
	}
	fmt.Printf("%s\n", statusJSON)
```

```bash tab="cURL"
curl -i -X GET \
   -H "X-Algo-API-Token:<algod-token>" \
 'http://<algod-address>:<algod-port>/v2/status'
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
Last Catchpoint: []
Genesis ID: [GENESIS_ID]
Genesis hash: [GENESIS_HASH]
```

The _status_ methods returns information about the status of the node, such as the latest round<LINK TO GLOSSARY>, referred to as `last-round`, from the perspective of the node you are connected to. Each of the SDKs may differ slightly in which information they return for each call. Shown below is the response from the REST API call.

```json tab="Response"
{
	"catchpoint": "",
    "catchpoint-acquired-blocks": 0,
    "catchpoint-processed-accounts": 0,
    "catchpoint-total-accounts": 0,
    "catchpoint-total-blocks": 0,
    "catchup-time": 0,
    "last-catchpoint": "",
    "last-round": 4243027,
    "last-version": "https://github.com/algorandfoundation/specs/tree/4a9db6a25595c6fd097cf9cc137cc83027787eaa",
    "next-version": "https://github.com/algorandfoundation/specs/tree/4a9db6a25595c6fd097cf9cc137cc83027787eaa",
    "next-version-round": 4243028,
    "next-version-supported": true,
    "stopped-at-unsupported-round": false,
    "time-since-last-round": 4261519666,
}
```

Check if the node is caught up by validating against others running nodes, like a [public block explorer](../community.md#block-explorers). As a secondary check, see if your `catchup-time` is 0 and your rounds are progressing at a rate of less than 5 seconds on average. This is the time it takes to confirm a block on Algorand. Note that the `time-since-last-round` is represented in nanoseconds.

!!! warning
	If your node is out-of-sync with the rest of the network you cannot send transactions and account balances will be out-of-date. 

# Check suggested transaction parameters

The _/v2/transactions/params_ endpoint returns information about the identity of the network and parameters for constructing a new transaction. 


```javascript tab="JavaScript"
...
	let params = await algodClient.getTransactionParams().do();
	console.log("Algorand suggested parameters: %o", params)
...
```

```python tab="Python"
...
	try:
		params = algod_client.suggested_params()
		print(json.dumps(params, indent=4))
	except Exception as e:
		print(e)
...
```

```java tab="Java"
	...
        try {
            TransactionParametersResponse params = client.TransactionParams().execute().body();
            System.out.println("Algorand suggested parameters: " + TransactionParametersResponse);
        } catch (ApiException e) {
            System.err.println("Exception when calling algod#TransactionParams");
            e.printStackTrace();
        }
	...
```

```go tab="Go"
...
	txParams, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error Algorand suggested parameters: %s\n", err)
		return
	}
	JSON, err := json.MarshalIndent(txParams, "", "\t")
	if err != nil {
		fmt.Printf("Can not marshall suggested parameters data: %s\n", err)
	}
	fmt.Printf("%s\n", JSON)
...
```

```bash tab="cURL"
curl -i -X GET \
   -H "X-Algo-API-Token:<algod-token>" \
 'http://<algod-address>:<algod-port>/v2/transactions/params'
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
Last Catchpoint: []
Genesis ID: [GENESIS_ID]
Genesis hash: [GENESIS_HASH]
```

Check the `genesis-id` and the `genesis-hash`, as shown in the REST response below. Ensure both match your chosen network before proceeding.

```json hl_lines="4 5" tab="Response"
{
    "consensus-version": "https://github.com/algorandfoundation/specs/tree/e5f565421d720c6f75cdd186f7098495caf9101f",
    "fee": 1,
    "genesis-hash": "SGO1GKSzyE7IEPItTxCByw9x8FmnrCDexi9/cOUJOiI=",
    "genesis-id": "testnet-v1.0",
    "last-round": 7430522,
    "min-fee": 1000
}
```
