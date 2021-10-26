title: 2. Connect to Node

The **algod IP address** and access **token** you obtained in the [Workspace Setup](../setup) section gives your application the credentials to interface with the Algorand blockchain. Interfacing with the Algorand blockchain using the SDKs is accomplished through an **algod client**. If using the algod REST API directly, these credentials can be supplied with each request.

_Read more about the [node's algod process](../../../run-a-node/reference/artifacts#algod)._

!!! info
    The examples in this section have been updated to the v2 API, which was launched to MainNet on June 16, 2020. Visit the [v2 Migration Guide](../migration) for information on how to migrate your code from v1. Access v2 and archived v1 examples [here](https://github.com/algorand/docs/tree/master/examples/start_building/v2). 

# Create an algod client
Instantiate an **algod** client with your preferred SDK. 

=== "JavaScript"
    ```js
    const algosdk = require('algosdk');

    const algodToken = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
    const algodServer = "http://localhost";
    const algodPort = 4001;

    let algodClient = new algosdk.Algodv2(algodToken, algodServer, algodPort);
    ```

=== "Python"
    ```Python
    from algosdk.v2client import algod

    algod_address = "http://localhost:4001"
    algod_token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
    algod_client = algod.AlgodClient(algod_token, algod_address)
    ```

=== "Java"
    ```Java
    import com.algorand.algosdk.v2.client.common.AlgodClient;

    final String ALGOD_API_ADDR = "localhost";
    final Integer ALGOD_PORT = 4001;
    final String ALGOD_API_TOKEN = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";

    AlgodClient client = new AlgodClient(ALGOD_API_ADDR, ALGOD_PORT, ALGOD_API_TOKEN);
    ```
=== "Go"
    ```Go 
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

Some third-party services use a different API key header than the one used by default.
For example, if the API key header is `X-API-Key`, the client can be instantiated as follows:

=== "JavaScript"
    ```js
    from algosdk.v2client import algod

    const algodServer = "http://localhost";
    const port = "4001";
    const token = {
        'X-API-Key': "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
    };
    const algodClient = new algosdk.Algodv2(algodToken, algodServer, algodPort);
    ```
=== "Python"
    ```Python
    from algosdk.v2client import algod

    algod_address = "http://localhost:4001"
    algod_token = ""
    headers = {
        "X-API-Key": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
    }

    algod_client = algod.AlgodClient(algod_token, algod_address, headers)
    ```

=== "Java"
    ```Java 
    import com.algorand.algosdk.v2.client.common.AlgodClient;

    final String ALGOD_API_ADDR = "http://localhost";
    final Integer ALGOD_PORT = 4001;

    final String[] ALGOD_API_KEY_HEADERS = {"X-API-Key"};
    final String[] ALGOD_API_KEY_VALUES = {"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"};

    AlgodClient client = new AlgodClient(ALGOD_API_ADDR, ALGOD_PORT, "");

    // Contrary to the other SDK, the Java SDK does not allow to pass
    // the headers to all queries automatically.
    // Instead, the headers should be passed when executing each query
    // For example:
    //   client.GetStatus().execute(ALGOD_API_KEY_HEADERS, ALGOD_API_KEY_VALUES).body();
    ```
=== "Go"
    ```go 
    package main

    import (
        "github.com/algorand/go-algorand-sdk/client/v2/algod" 
    )

    const algodAddress = "http://localhost:4001"
    const apiKey = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

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

=== "JavaScript"
    ```javascript 
    let status = (await algodClient.status().do());
    console.log("Algorand network status: %o", status);
    ```
=== "Python"
    ```python 
    status = algod_client.status()
    print(json.dumps(status, indent=4))
    ```

=== "Java"
    ```java 
    try {
        Response < NodeStatusResponse > resp = myclient.GetStatus().execute();
        if (!resp.isSuccessful()) {
            throw new Exception(resp.message());
        }
        NodeStatusResponse status = resp.body();
        System.out.println("Algorand network status: " + status);
    } catch (ApiException e) {
        System.err.println("Exception when calling algod#getStatus");
        e.printStackTrace();
    }
    ```
=== "Go"
    ```go 
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
=== "cURL"
    ```bash 
    curl -i -X GET \
       -H "X-Algo-API-Token:<algod-token>" \
     'http://<algod-address>:<algod-port>/v2/status'
    ```

=== "goal"

    ```bash
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

=== "Response"
    ```json 
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

Check if the node is caught up by validating against others running nodes, like a [public block explorer](https://developer.algorand.org/ecosystem-projects/#block-explorers). As a secondary check, see if your `catchup-time` is 0 and your rounds are progressing at a rate of less than 5 seconds on average. This is the time it takes to confirm a block on Algorand. Note that the `time-since-last-round` is represented in nanoseconds.

!!! warning
	If your node is out-of-sync with the rest of the network you cannot send transactions and account balances will be out-of-date. 

# Check suggested transaction parameters

The _/v2/transactions/params_ endpoint returns information about the identity of the network and parameters for constructing a new transaction. 


=== "JavaScript"
    ```javascript 
    ...
        let params = await algodClient.getTransactionParams().do();
        console.log("Algorand suggested parameters: %o", params)
    ...
    ```

=== "Python"
    ```python 
    ...
        try:
            params = algod_client.suggested_params()
            print(json.dumps(vars(params), indent=4))
        except Exception as e:
            print(e)
    ...
    ```
=== "Java"
    ```java 
    ...
        try {
            Response < TransactionParametersResponse > resp = client.TransactionParams().execute();
            if (!resp.isSuccessful()) {
                throw new Exception(resp.message());
            }
            TransactionParametersResponse params = resp.body();
            if (params == null) {
                throw new Exception("Params retrieval error");
            }            
            System.out.println("Algorand suggested parameters: " + params);
        } catch (ApiException e) {
            System.err.println("Exception when calling algod#TransactionParams");
            e.printStackTrace();
        }
    ...
    ```

=== "Go"
    ```go 
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

=== "cURL"
    ```bash 
    curl -i -X GET \
       -H "X-Algo-API-Token:<algod-token>" \
     'http://<algod-address>:<algod-port>/v2/transactions/params'
    ```

=== "goal"
    ```bash
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

=== "Response"

    ```json 
    {
        "consensus-version": "https://github.com/algorandfoundation/specs/tree/e5f565421d720c6f75cdd186f7098495caf9101f",
        "fee": 1,
        "genesis-hash": "SGO1GKSzyE7IEPItTxCByw9x8FmnrCDexi9/cOUJOiI=",
        "genesis-id": "testnet-v1.0",
        "last-round": 7430522,
        "min-fee": 1000
    }
    ```
