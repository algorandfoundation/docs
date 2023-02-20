title: Indexer

This guide explains how to search the Algorand Blockchain using the Indexer. This Indexer not part of the Algorand node and requires a separate binary download. See the [Indexer Install Guide](../run-a-node/setup/indexer.md) for instructions on installing and configuring the new Indexer.

![Block Proposal](../imgs/indexerv2.png)
<center>*Algorand V2 Indexer*</center>

The primary purpose of this Indexer is to provide a REST API interface of API calls to support searching the Algorand Blockchain. The Indexer REST APIs retrieve the blockchain data from a [PostgreSQL](https://www.postgresql.org/) compatible database that must be populated. This database is populated using the same indexer instance or a separate instance of the indexer which must connect to the algod process of a running Algorand node to read block data. This node must also be an Archival node to make searching the entire blockchain possible. 

The Indexer provides a set of REST API calls for searching blockchain Transactions, Accounts, Assets and Blocks. Each of these calls also provides several filter parameters to support refining searches. The latest Algorand native SDKs (Python, JavaScript, Go, and Java) provide similar functionality. These REST calls are based on the Open API specification and are described in the REST SDK reference documentation. 

See the full description of endpoints available in the [indexer docs](../rest-apis/indexer.md).

# SDK client instantiations

=== "JavaScript"
    <!-- ===JSSDK_CREATE_INDEXER_CLIENT=== -->
	```javascript
    const algosdk = require('algosdk');

    const indexer_token = "";
    const indexer_server = "http://localhost";
    const indexer_port = 8980;

    const indexerClient = new algosdk.Indexer(indexer_token, indexer_server, indexer_port);
    ```
    <!-- ===JSSDK_CREATE_INDEXER_CLIENT=== -->

=== "Python"
    <!-- ===PYSDK_CREATE_INDEXER_CLIENT=== -->
```python
# instantiate indexer client
indexer_host = "http://localhost:8980"
indexer_token = "a" * 64
myindexer = indexer.IndexerClient(
    indexer_token=indexer_token, indexer_address=indexer_host
)
```
    <!-- ===PYSDK_CREATE_INDEXER_CLIENT=== -->

=== "Java"
    <!-- ===JAVASDK_CREATE_INDEXER_CLIENT=== -->
	```java
    package com.algorand.javatest.indexer;

    import com.algorand.algosdk.v2.client.common.IndexerClient;
    import com.algorand.algosdk.v2.client.common.Client;

    public class InstantiateIndexer {
        public Client indexerInstance = null;
        // utility function to connect to a node
        private Client connectToNetwork(){
            final String INDEXER_API_ADDR = "localhost";
            final int INDEXER_API_PORT = 8980;       
            IndexerClient indexerClient = new IndexerClient(INDEXER_API_ADDR, INDEXER_API_PORT); 
            return indexerClient;
        }
        public static void main(String args[]) throws Exception {
            InstantiateIndexer ex = new InstantiateIndexer();
            IndexerClient indexerClientInstance = (IndexerClient)ex.connectToNetwork();
            System.out.println("IndexerClient Instantiated : " + indexerClientInstance); // pretty print json
        }
    }
    ```
    <!-- ===JAVASDK_CREATE_INDEXER_CLIENT=== -->

=== "Go"
<!-- ===GOSDK_CREATE_INDEXER_CLIENT=== -->
	```go
    package main

    import (
        "context"
        "encoding/json"
        "fmt"

        "github.com/algorand/go-algorand-sdk/v2/client/v2/indexer"
    )

    const indexerAddress = "http://localhost:8980"
    const indexerToken = ""

    func main() {
        // instantiate indexer client
        indexerClient, err := indexer.MakeClient(indexerAddress, indexerToken)
    }
    ```
<!-- ===GOSDK_CREATE_INDEXER_CLIENT=== -->

!!! info 
    When using cURL be aware that the parameters may need to be URL encoded. The SDKs handle the encoding of parameter data. 

The indexer provides two primary ways to access information:

1) [Lookup](#lookup) a single item (a single account, a single transaction, a single block)
2) Search for items that match a query (transactions for > 100A and between round N and M)

# Lookup

When performing a lookup of a single item ...

To get the details of a specific asset the indexer provides the `/assets/{asset-id}` REST call.  This call takes no parameters as the asset id is passed in the URL. This call returns the details of the asset and the round the results were calculated in. 

=== "JavaScript"
    <!-- ===JSSDK_INDEXER_LOOKUP_ASSET=== -->
	```javascript
    (async () => {
        let assetIndex = 2044572;
        let assetInfo = await indexerClient.searchForAssets()
            .index(assetIndex).do();
        console.log("Information for Asset: " + JSON.stringify(assetInfo, undefined, 2));
    })().catch(e => {
        console.log(e);
        console.trace();
    });
    ```
    <!-- ===JSSDK_INDEXER_LOOKUP_ASSET=== -->

=== "Python"
    <!-- ===PYSDK_INDEXER_LOOKUP_ASSET=== -->
```python
# lookup a single asset
asset_id = 2044572
# by passing include_all, we specify that we want to see deleted assets as well
response = myindexer.asset_info(asset_id, include_all=True)
print(f"Asset Info: {json.dumps(response, indent=2,)}")
```
    <!-- ===PYSDK_INDEXER_LOOKUP_ASSET=== -->

=== "Java"
    <!-- ===JAVASDK_INDEXER_LOOKUP_ASSET=== -->
	```java
    public static void main(String args[]) throws Exception {
        SearchAssets ex = new SearchAssets();
        IndexerClient indexerClientInstance = (IndexerClient)ex.connectToNetwork();
        Long asset_id = Long.valueOf(12215366);        
        Response<AssetsResponse> response = indexerClientInstance
            .searchForAssets()
            .assetId(asset_id)
            .execute();
        if (!response.isSuccessful()) {
            throw new Exception(response.message());
        } 

        JSONObject jsonObj = new JSONObject(response.body().toString());
        System.out.println("Asset Info: " + jsonObj.toString(2)); // pretty print json            
    }
    ```
    <!-- ===JAVASDK_INDEXER_LOOKUP_ASSET=== -->

=== "Go"
    <!-- ===GOSDK_INDEXER_LOOKUP_ASSET=== -->
	```go
    // query parameters
    var assetID uint64 = 2044572

    // Query
    _, result, err := indexerClient.LookupAssetByID(assetID).Do(context.Background())

    // Print the results
    JSON, err := json.MarshalIndent(result, "", "\t")
    fmt.Printf(string(JSON) + "\n")
    ```
    <!-- ===GOSDK_INDEXER_LOOKUP_ASSET=== -->

=== "cURL"
    <!-- ===CURL_INDEXER_LOOKUP_ASSET=== -->
	``` bash
    $ curl "localhost:8980/v2/assets/2044572" 
    ```
    <!-- ===CURL_INDEXER_LOOKUP_ASSET=== -->



# Search

... show search with chained filters

=== "JavaScript"
    <!-- ===JSSDK_INDEXER_SEARCH_MIN_AMOUNT=== -->
	```javascript
    (async () => {
        let currencyGreater = 10;
        let transactionInfo = await indexerClient.searchForTransactions()
            .currencyGreaterThan(currencyGreater).do();
        console.log("Information for Transaction search: " + JSON.stringify(transactionInfo, undefined, 2));
    })().catch(e => {
        console.log(e);
        console.trace();
    });
    ```
    <!-- ===JSSDK_INDEXER_SEARCH_MIN_AMOUNT=== -->

=== "Python"
    <!-- ===PYSDK_INDEXER_SEARCH_MIN_AMOUNT=== -->
```python
response = myindexer.search_transactions(
    min_amount=10, min_round=1000, max_round=1500
)
print(f"Transaction results: {json.dumps(response, indent=2)}")
```
    <!-- ===PYSDK_INDEXER_SEARCH_MIN_AMOUNT=== -->

=== "Java"
    <!-- ===JAVASDK_INDEXER_SEARCH_MIN_AMOUNT=== -->
	```java
    public static void main(String args[]) throws Exception {
        SearchTransactionsMinAmount ex = new SearchTransactionsMinAmount();
        IndexerClient indexerClientInstance = (IndexerClient)ex.connectToNetwork();
        Long min_amount = Long.valueOf(10);     
        Response<TransactionsResponse> response = indexerClientInstance
                .searchForTransactions()
                .currencyGreaterThan(min_amount)
                .execute();
        if (!response.isSuccessful()) {
            throw new Exception(response.message());
        }                

        JSONObject jsonObj = new JSONObject(response.body().toString());
        System.out.println("Transaction Info: " + jsonObj.toString(2)); // pretty print json        
    }
    ```
    <!-- ===JAVASDK_INDEXER_SEARCH_MIN_AMOUNT=== -->

=== "Go"
    <!-- ===GOSDK_INDEXER_SEARCH_MIN_AMOUNT=== -->
	```go
    // query parameters
    var minAmount uint64 = 10

    // Query
    result, err := indexerClient.SearchForTransactions().CurrencyGreaterThan(minAmount).Do

    // Print results
    result, err := json.MarshalIndent(transactions, "", "\t")
    fmt.Printf(string(JSON) + "\n")
    ```
    <!-- ===GOSDK_INDEXER_SEARCH_MIN_AMOUNT=== -->

=== "cURL"
    <!-- ===CURL_INDEXER_SEARCH_MIN_AMOUNT=== -->
	```bash
    $ curl "localhost:8980/v2/transactions?currency-greater-than=10"
    ```
    <!-- ===CURL_INDEXER_SEARCH_MIN_AMOUNT=== -->

# Pagination

When searching large amounts of blockchain data often the results may be too large to process in one given operation. In fact, the indexer imposes hard limits on the number of results returned for specific searches.  The default limits for these searches are summarized in the table below.

| Search Type  | Maximum number of results per search  |
| ------------- | ------------- |
| API Resources Per Account | 1,000 |
| Transactions Search | 1,000 |
| Accounts Search | 100 |
| Assets Search | 100 |
| Balances Search | 1,000 |
| Applications Search | 100 |

When trying to find specific transactions, the Indexer supplies a pagination method that allows separating the results into several REST calls to return larger result sets. When used with the limit parameter the results for large data sets can be returned in expected result counts.

For example, adding a limit parameter of 5 to the previous call will cause only 5 results to be returned in each page. To get the next 5 transactions simply add the next-token as a parameter to the next REST call. The parameter is named `next` and this token is only good for the next 5 results.

=== "JavaScript"
    <!-- ===JSSDK_INDEXER_PAGINATE_RESULTS=== -->
	```javascript
    let nexttoken = "";
    let numtx = 1;
    // loop until there are no more transactions in the response
    // for the limit(max limit is 1000  per request)    
    (async () => {
        let min_amount = 100000000000000;
        let limit = 5;
        while (numtx > 0) {
            // execute code as long as condition is true
            let next_page = nexttoken;
            let response = await indexerClient.searchForTransactions()
                .limit(limit)
                .currencyGreaterThan(min_amount)
                .nextToken(next_page).do();
            let transactions = response['transactions'];
            numtx = transactions.length;
            if (numtx > 0)
            {
                nexttoken = response['next-token']; 
                console.log("Transaction Information: " + JSON.stringify(response, undefined, 2));           
            }
        }
    })().catch(e => {
        console.log(e);
        console.trace();
    });
    ```
    <!-- ===JSSDK_INDEXER_PAGINATE_RESULTS=== -->

=== "Python"
    <!-- ===PYSDK_INDEXER_PAGINATE_RESULTS=== -->
```python

nexttoken = ""
has_results = True
page = 0

# loop using next_page to paginate until there are 
# no more transactions in the response
while has_results:
    response = myindexer.search_transactions(
        min_amount=10, min_round=1000, max_round=1500
    )

    has_results = len(response['transactions'])>0

    if has_results:
        nexttoken = response['next-token']
        print(f"Tranastion on page {page}: " + json.dumps(response, indent=2))

    page += 1
```
    <!-- ===PYSDK_INDEXER_PAGINATE_RESULTS=== -->

=== "Java"
    <!-- ===JAVASDK_INDEXER_PAGINATE_RESULTS=== -->
	```java
    public static void main(String args[]) throws Exception {
        SearchTransactionsPaging ex = new SearchTransactionsPaging();
        IndexerClient indexerClientInstance = (IndexerClient) ex.connectToNetwork();
        String nexttoken = "";
        Integer numtx = 1;  
        Long maxround=Long.valueOf(30000);           
        // loop until there are no more transactions in the response
        // for the limit (max limit is 1000 per request)
        while (numtx > 0) {
            Long min_amount = Long.valueOf(500000000000L);
            Long limit = Long.valueOf(4);
            String next_page = nexttoken;
            Response<TransactionsResponse> response = indexerClientInstance
                .searchForTransactions()
                .next(next_page)
                .currencyGreaterThan(min_amount)
                .maxRound(maxround)
                .limit(limit)
                .execute();
            if (!response.isSuccessful()) {
                throw new Exception(response.message());
            }        
            JSONObject jsonObj = new JSONObject(response.body().toString());

            JSONArray jsonArray = (JSONArray) jsonObj.get("transactions");
            numtx = jsonArray.length();
            if (numtx > 0) {
                nexttoken = jsonObj.get("next-token").toString();
                JSONObject jsonObjAll = new JSONObject(response.body().toString());
                System.out.println("Transaction Info: " + jsonObjAll.toString(2)); // pretty print json
            }
        }
    }
    ```
    <!-- ===JAVASDK_INDEXER_PAGINATE_RESULTS=== -->

=== "Go"
    <!-- ===GOSDK_INDEXER_PAGINATE_RESULTS=== -->
	```go
    var nextToken = ""
    var numTx = 1
    var numPages = 1
    var minAmount uint64 = 100000000000000
    var limit uint64 = 2

    for numTx > 0 {
    // Query
    result, err := indexerClient.SearchForTransactions().CurrencyGreaterThan(minAmount).Limit(limit).NextToken(nextToken).Do(context.Background())
    if err != nil {
        return
    }
    transactions := result.Transactions
    numTx = len(transactions)
    nextToken = result.NextToken

    if numTx > 0 {
        // Print results
        JSON, err := json.MarshalIndent(transactions, "", "\t")
        if err != nil {
            return
        }
        fmt.Printf(string(JSON) + "\n")
        fmt.Println("End of page : ", numPages)
        fmt.Println("Transaction printed : ", len(transactions))
        fmt.Println("Next Token : ", nextToken)
        numPages++
    }
    }
    ```
    <!-- ===GOSDK_INDEXER_PAGINATE_RESULTS=== -->

=== "cURL"
    <!-- ===CURL_INDEXER_PAGINATE_RESULTS=== -->
	```bash
    $ curl "localhost:8980/v2/transactions?currency-greater-than=10&limit=5"
    # note the "next-token" field in the most resent results and supply the value to the "next" parameter
    $ curl "localhost:8979/v2/transactions?currency-greater-than=10&limit=5&next=cAoBAAAAAAAAAAAA"
    ```
    <!-- ===CURL_INDEXER_PAGINATE_RESULTS=== -->

Results showing "next-token"
```json
{
   "next-token" : "cAoBAAAAAAAAAAAA",
   "current-round" : 7050272,
   "transactions" : []
}
```

A new next token will be returned to get the next five. This token acts as a marker in the current result set and allows the next call to pick up where the last search ended. We note that the 'limit' parameter can also be used to specify a larger (rather than smaller) results set than the defaults above. These limits are shown in the following table and are per call, not the total result set.

| Search Type  | Search Limit with a limit parameter |
| ------------- | ------------- |
| Transaction Search | 10000 |
| Account Search | 1000 |
| Asset Search | 1000 |
| Asset Balance Search | 10000 |

The following REST calls support paginated results.

* `/accounts` - Search for specific accounts.
* `/accounts/{account-id}/transactions` - Search for transactions for a specific account.
* `/accounts/{account-id}/created-apps` - Search for created applications for a specific account.
* `/accounts/{account-id}/created-assets` - Search for created assets for a specific account.
* `/accounts/{account-id}/assets` - Search for assets on a specific account.
* `/accounts/{account-id}/apps-local-state` - Search for application local state for a specific account.
* `/assets` - Search Assets.
* `/assets/{asset-id}/balances` - Search Asset balances.
* `/assets/{asset-id}/transactions` - Search for Transactions with a specific Asset.
* `/transactions` - Search all transactions

# Note field searching
Every transaction has the ability to add up to a 1kb note in the note field. Several of the REST APIs provide the ability to search for a prefix that is present in the note field, meaning that the note starts with a specific string. This can be a very powerful way to quickly locate transactions that are specific to an application. The REST calls that support prefix searching are the following.

* `/accounts/{account-id}/transactions` - Search for a prefix for a specific accounts transactions.
* `/assets/{asset-id}/transactions` - Search for a prefix for a specific Asset Id.
* `/transactions` - Search all Transactions for a specific transaction note field prefix.

To search for a specific prefix use the `note-prefix` parameter. For the Javascript and direct REST API, the value needs to be base64 encoded to return results. (The other SDKs perform the base64 encoding for you.) For example, if the contents of the note field started with the string “showing prefix searches”, encoding the beginning of that sentence using python like the following:

``` bash
$ python3 -c "import base64;print(base64.b64encode('showing prefix'.encode()))"
```

This will return an encoded value of `c2hvd2luZyBwcmVmaXg=`.  This value can then be passed to the search. To search all transactions use the following command.

=== "JavaScript"
    <!-- ===JSSDK_INDEXER_PREFIX_SEARCH=== -->
	```javascript
    (async () => {
        const enc = new TextEncoder();
        let note = enc.encode("Hello");  
        let s = Buffer.from(note).toString("base64");
        let transactionInfo = await indexerClient.searchForTransactions()
            .notePrefix(s).do();
        console.log("Information for Transaction search: " + JSON.stringify(transactionInfo, undefined, 2));
    })().catch(e => {
        console.log(e);
        console.trace();
    });
    ```
    <!-- ===JSSDK_INDEXER_PREFIX_SEARCH=== -->

=== "Python"
    <!-- ===PYSDK_INDEXER_PREFIX_SEARCH=== -->
```python
note_prefix = "showing prefix".encode()
response = myindexer.search_transactions(note_prefix=note_prefix)
print(f"result: {json.dumps(response, indent=2)}")
```
    <!-- ===PYSDK_INDEXER_PREFIX_SEARCH=== -->

=== "Java"
    <!-- ===JAVASDK_INDEXER_PREFIX_SEARCH=== -->
	```java
    public static void main(String args[]) throws Exception {
        IndexerClient indexerClientInstance = connectToNetwork();
        Response<TransactionsResponse> resp = indexerClientInstance.searchForTransactions()
                .notePrefix("showing prefix".getBytes())
                .minRound(10894697L)
                .maxRound(10994697L).execute();
        if (!resp.isSuccessful()) {
            throw new Exception(resp.message());
        }

        // pretty print json
        JSONObject jsonObj = new JSONObject(resp.body().toString());
        System.out.println("Transaction Info: " + jsonObj.toString(2));

        int i = 0;
        for (Transaction tx : resp.body().transactions) {
            i++;
            System.out.println("Transaction " + i);
            System.out.println("  Note Info: " + new String(tx.note));
        }
    }
    ```
    <!-- ===JAVASDK_INDEXER_PREFIX_SEARCH=== -->

=== "Go"
    <!-- ===GOSDK_INDEXER_PREFIX_SEARCH=== -->
	```go
    // Parameters
    var notePrefix = "showing prefix"

    // Query
    result, err := indexerClient.SearchForTransactions().NotePrefix([]byte(notePrefix)).Do(context.Background())

    // Print results
    JSON, err := json.MarshalIndent(result, "", "\t")
    fmt.Printf(string(JSON) + "\n")

    ```
    <!-- ===GOSDK_INDEXER_PREFIX_SEARCH=== -->

=== "cURL"
    <!-- ===CURL_INDEXER_PREFIX_SEARCH=== -->
	```bash
    python3 -c "import base64;print(base64.b64encode('showing prefix'.encode()))"
    $ curl "localhost:8980/v2/transactions?note-prefix=c2hvd2luZyBwcmVmaXg=" | json_pp
    ```
    <!-- ===CURL_INDEXER_PREFIX_SEARCH=== -->