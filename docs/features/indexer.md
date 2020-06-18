title: Searching the Blockchain (Indexer V2) 🆕

This guide explains how to search the Algorand Blockchain using the V2 Indexer. This Indexer is no longer part of the Algorand node and requires a separate binary download. See the [Indexer Install Guide](../run-a-node/setup/indexer.md) for instructions on installing and configuring the new Indexer.

![Block Proposal](../imgs/indexerv2.png)
<center>*Algorand V2 Indexer*</center>

The primary purpose of this Indexer is to provide a REST API interface of API calls to support searching the Algorand Blockchain. The Indexer REST APIs retrieve the blockchain data from a [PostgreSQL](https://www.postgresql.org/) compatible database that must be populated. This database is populated using the same indexer instance or a separate instance of the indexer which must connect to the algod process of a running Algorand node to read block data. This node must also be an Archival node to make searching the entire blockchain possible. 

!!! info
    The Indexer DB takes up a fraction of disk space compared to the actual blockchain data with archival mode on. For example, 100 GB of blockchain data takes about 1 GB of data in the Indexer DB.

The Indexer provides a set of REST API calls for searching blockchain Transactions, Accounts, Assets and Blocks. Each of these calls also provides several filter parameters to support refining searches. The latest Algorand native SDKs (Python, JavaScript, Go, and Java) provide similar functionality. Thes REST calls are based on the Open API specification and are described in the REST SDK reference documentation. 

!!! info
    Example code snippets are provided throughout this page. Full running code examples for each SDK are available within the GitHub repo at [/examples/indexer](https://github.com/algorand/docs/tree/master/examples/indexer) and for [download](https://github.com/algorand/docs/blob/master/examples/indexer/indexer.zip?raw=true) (.zip).

# SDK Client Instantiations

```javascript tab="JavaScript"
// requires algosdk@1.6.1 or higher 
// verify installed version
// npm list algosdk
const algosdk = require('algosdk');
const indexer_token = "";
const indexer_server = "http://localhost";
const indexer_port = 8980;
```

```python tab="Python"
import json
# requires Python SDK version 1.3 or higher
from algosdk.v2client import indexer

# instantiate indexer client
myindexer = indexer.IndexerClient(indexer_token="", indexer_address="http://localhost:8980")
```

```java tab="Java"
// /indexer/java/InstantiateIndexer.java
// requires java-algorand-sdk 1.4.1 or higher (see pom.xml)
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

```go tab="Go"
// requires Go SDK version 1.4 or higher

package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/algorand/go-algorand-sdk/client/v2/indexer"
)

const indexerAddress = "http://localhost:8980"
const indexerToken = ""

func main() {

	// instantiate indexer client
   indexerClient, err := indexer.MakeClient(indexerAddress, indexerToken)
}
```

!!! info 
    When using cURL be aware that the parameters may need to be URL encoded. The SDKs handle the encoding of parameter data. 

Before describing each of the available REST APIs, a few specific functions are described that are supported across many of the calls. This includes [Paginated Results](#paginated-results), [Historical Data searches](#historical-data-searches), and [Note Field Searching](#note-field-searching).

# Paginated Results
When searching large amounts of blockchain data often the results may be too large to process in one given operation. In fact, the indexer imposes hard limits on the number of results returned for specific searches.  The default limits for these searches are summarized in the table below.

| Search Type  | Maximum number of results per search  |
| ------------- | ------------- |
| Transaction Search | 1000 |
| Account Search | 100 |
| Asset Search | 100 |
| Asset Balance Search | 1000 |

These values represent the maximum number of results that will be returned when searching for specific results. For example, the following will return the last 1000 transactions that exceeded 10 microAlgos. 

```javascript tab="JavaScript"
// /indexer/javascript/SearchTransactionsMinAmount.js
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

```python tab="Python"
# /indexer/python/search_transactions_min_amount.py

response = myindexer.search_transactions(min_amount=10) 

# Pretty Printing JSON string
print(json.dumps(response, indent=2, sort_keys=True))
```

```java tab="Java"
// /indexer/java/SearchTransactionsMinAmount.java
    public static void main(String args[]) throws Exception {
        SearchTransactionsMinAmount ex = new SearchTransactionsMinAmount();
        IndexerClient indexerClientInstance = (IndexerClient)ex.connectToNetwork();
        Long min_amount = Long.valueOf(10);     
        String response = indexerClientInstance
                .searchForTransactions()
                .currencyGreaterThan(min_amount).execute().toString();
        JSONObject jsonObj = new JSONObject(response.toString());
        System.out.println("Transaction Info: " + jsonObj.toString(2)); // pretty print json
    }
};
```

```go tab="Go"
// query parameters
var minAmount uint64 = 10

// Query
result, err := indexerClient.SearchForTransactions().CurrencyGreaterThan(minAmount).Do

// Print results
result, err := json.MarshalIndent(transactions, "", "\t")
fmt.Printf(string(JSON) + "\n")
```

```bash tab="cURL"
 $ curl "localhost:8980/v2/transactions?currency-greater-than=10"
```

When trying to find specific transactions, the Indexer supplies a pagination method that allows separating the results into several REST calls to return larger result sets. When used with the limit parameter the results for large data sets can be returned in expected result counts.

For example, adding a limit parameter of 5 to the previous call

```javascript tab="JavaScript"
// /indexer/javascript/SearchTransactionsLimit.js
(async () => {
    let currencyGreater = 10;
    let limit = 5;
    let transactionInfo = await indexerClient.searchForTransactions()
        .currencyGreaterThan(currencyGreater)
        .limit(limit).do();
    console.log("Information for Transaction search: " + JSON.stringify(transactionInfo, undefined, 2));
})().catch(e => {
    console.log(e);
    console.trace();
});
```

```python tab="Python"
# /indexer/python/search_transactions_limit.py

response = myindexer.search_transactions(
   min_amount=10, limit=5) 

# Pretty Printing JSON string
print(json.dumps(response, indent=2, sort_keys=True))
```

```java tab="Java"
// /indexer/java/SearchTransactionsLimit.js
    public static void main(String args[]) throws Exception {
        SearchTransactionsLimit ex = new SearchTransactionsLimit();
        IndexerClient indexerClientInstance = (IndexerClient)ex.connectToNetwork();
        Long min_amount = Long.valueOf(10);
        Long limit = Long.valueOf(2);       
        String response = indexerClientInstance
                .searchForTransactions()
                .currencyGreaterThan(min_amount)
                .limit(limit)
        .execute().toString();
        JSONObject jsonObj = new JSONObject(response.toString());
        System.out.println("Transaction Info: " + jsonObj.toString(2)); // pretty print json
    }
```

```go tab="Go"
// query parameters
var minAmount uint64 = 10
var limit uint64 = 5
// Query
result, err := indexerClient.SearchForTransactions().CurrencyGreaterThan(minAmount).Limit(limit).Do

// Print results
result, err := json.MarshalIndent(transactions, "", "\t")
fmt.Printf(string(JSON) + "\n")
```

```bash tab="cURL"
 $ curl "localhost:8980/v2/transactions?currency-greater-than=10&limit=5"
```

Will return the last 5 transactions over 10 microAlgos. In addition, a token will be added to the result list that allows querying the next 5 transactions over 5 microAlgos.

Results showing "next-token"
```bash
{
   "next-token" : "cAoBAAAAAAAAAAAA",
   "current-round" : 7050272,
   "transactions" : [
	.
	.
```

To get the next 5 transactions simply add the next-token as a parameter to the next REST call. The parameter is named `next` and this token is only good for the next 5 results.

```javascript tab="JavaScript"
// /indexer/javascript/SearchTransactionsPaging.js
let nexttoken = "";
let numtx = 1;
// loop until there are no more transactions in the response
// for the limit(max limit is 1000  per request)    
(async () => {
    let min_amount = 100000000000000;
    let limit = 2;
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

```python tab="Python"
# /indexer/python/search_transactions_paging.py

nexttoken = ""
numtx = 1

# loop using next_page to paginate until there are no more transactions in the response
# for the limit (max is 1000  per request)

while (numtx > 0):

    response = myindexer.search_transactions(
        min_amount=100000000000000, limit=2, next_page=nexttoken) 
    transactions = response['transactions']
    numtx = len(transactions)
    if (numtx > 0):
        nexttoken = response['next-token']
        # Pretty Printing JSON string 
        print("Tranastion Info: " + json.dumps(response, indent=2, sort_keys=True))
```

```java tab="Java"
// /indexer/java/SearchTransactionsPaging.java
    public static void main(String args[]) throws Exception {
        SearchTransactionsPaging ex = new SearchTransactionsPaging();
        IndexerClient indexerClientInstance = (IndexerClient) ex.connectToNetwork();
        String nexttoken = "";
        Integer numtx = 1;               
        // loop until there are no more transactions in the response
        // for the limit (max limit is 1000 per request)
        while (numtx > 0) {
            Long min_amount = Long.valueOf(100000000000000L);
            Long limit = Long.valueOf(2);
            String next_page = nexttoken;
            String response = indexerClientInstance.searchForTransactions().next(next_page)
                    .currencyGreaterThan(min_amount).limit(limit).execute().toString();
            JSONObject jsonObj = new JSONObject(response.toString());

            JSONArray jsonArray = (JSONArray) jsonObj.get("transactions");
            numtx = jsonArray.length();
            if (numtx > 0) {

                nexttoken = jsonObj.get("next-token").toString();
                JSONObject jsonObjAll = new JSONObject(response.toString());
                System.out.println("Transaction Info: " + jsonObjAll.toString(2)); // pretty print json
            }
        }
```

```go tab="Go"
// query parameters
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

```bash tab="cURL"
# note the "next-token" field in the most resent results and supply the value to the "next" parameter
$ curl "localhost:8980/v2/transactions?currency-greater-than=10&limit=5&next=cAoBAAAAAAAAAAAA"
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
* `/assets` - Search Assets.
* `/assets/{asset-id}/balances` - Search Asset balances.
* `/assets/{asset-id}/transactions` - Search for Transactions with a specific Asset.
* `/transactions` - Search all transactions

# Historical Data Searches 
Many of the REST calls support getting values at specific rounds. This means that the Indexer will do calculations that determine what specific values were at a specific round. For example, if account A starts at round 50 with 200 ARCC tokens and spends 50 of those tokens in round 75, the following command would return a balance of 150

```javascript tab="JavaScript"
// /indexer/javascript/AccountInfo.js
(async () => {
    let acct = "7WENHRCKEAZHD37QMB5T7I2KWU7IZGMCC3EVAO7TQADV7V5APXOKUBILCI";
    let accountInfo = await indexerClient.lookupAccountByID(acct).do();
    console.log("Information for Account: " + JSON.stringify(accountInfo, undefined, 2));
})().catch(e => {
    console.log(e);
    console.trace();
});
```

```python tab="Python"
# /indexer/python/account_info.py

response = myindexer.account_info(
    address="7WENHRCKEAZHD37QMB5T7I2KWU7IZGMCC3EVAO7TQADV7V5APXOKUBILCI")
print("Account Info: " + json.dumps(response, indent=2, sort_keys=True))
```

```java tab="Java"
// /indexer/java/AccountInfo.java
    public static void main(String args[]) throws Exception {
        AccountInfo ex = new AccountInfo();
        IndexerClient indexerClientInstance = (IndexerClient)ex.connectToNetwork();
        Address account = new Address("7WENHRCKEAZHD37QMB5T7I2KWU7IZGMCC3EVAO7TQADV7V5APXOKUBILCI");
         String response = indexerClientInstance.lookupAccountByID(account).execute().toString();
        JSONObject jsonObj = new JSONObject(response.toString());
        System.out.println("Account Info: " + jsonObj.toString(2)); // pretty print json
    }
```

```go tab="Go"

```

```bash tab="cURL"
$ curl localhost:8980/v2/accounts/7WENHRCKEAZHD37QMB5T7I2KWU7IZGMCC3EVAO7TQADV7V5APXOKUBILCI
```

If the round parameter is used and set to 50 a balance of 200 would be returned.

```javascript tab="JavaScript"
// indexer/javascript/AccountInfoBlock.js
(async () => {
    let acct = "7WENHRCKEAZHD37QMB5T7I2KWU7IZGMCC3EVAO7TQADV7V5APXOKUBILCI";
    let round = 50;
    let accountInfo = await indexerClient.lookupAccountByID(acct)
        .round(round).do();
    console.log("Information for Account at block: " + JSON.stringify(accountInfo, undefined, 2));
})().catch(e => {
    console.log(e);
    console.trace();
});
```

```python tab="Python"
# /indexer/python/account_info_block.py

response = myindexer.account_info(
    address="7WENHRCKEAZHD37QMB5T7I2KWU7IZGMCC3EVAO7TQADV7V5APXOKUBILCI", block=50)
print("Account Info: " + json.dumps(response, indent=2, sort_keys=True))
```

```java tab="Java"
// /indexer/java/AccountInfoBlock.java
    public static void main(String args[]) throws Exception {
        AccountInfoBlock ex = new AccountInfoBlock();
        IndexerClient indexerClientInstance = (IndexerClient)ex.connectToNetwork();
        Address account = new Address("7WENHRCKEAZHD37QMB5T7I2KWU7IZGMCC3EVAO7TQADV7V5APXOKUBILCI");
        Long round = Long.valueOf(50);
        String response = indexerClientInstance.lookupAccountByID(account).round(round).execute().toString();
        JSONObject jsonObj = new JSONObject(response.toString());
        System.out.println("Account Info for block: " + jsonObj.toString(2)); // pretty print json
    }
```

```go tab="Go"
// Parameters 
var round uint64 = 50
var account = "7WENHRCKEAZHD37QMB5T7I2KWU7IZGMCC3EVAO7TQADV7V5APXOKUBILCI"

// Lookup block
_, result, err := indexerClient.LookupAccountByID(account).Round(round).Do(context.Background())

// Print results
JSON, err := json.MarshalIndent(result, "", "\t")
fmt.Printf(string(JSON) + "\n")
```

```bash tab="cURL"
$ curl localhost:8980/v2/accounts/7WENHRCKEAZHD37QMB5T7I2KWU7IZGMCC3EVAO7TQADV7V5APXOKUBILCI?round=50
```

The round parameter results are inclusive of all transactions within the given round. This parameter is available in many of the REST calls. This can be very computationally expensive so for performance reasons, it is by default disabled on the `/accounts` REST call but can be enabled by using the `--dev-mode` flag when starting the Indexer.

# Note Field Searching
Every transaction has the ability to add up to a 1kb note in the note field. Several of the REST APIs provide the ability to search for a prefix that is present in the note field, meaning that the note starts with a specific string. This can be a very powerful way to quickly locate transactions that are specific to an application. The REST calls that support prefix searching are the following.

* `/accounts/{account-id}/transactions` - Search for a prefix for a specific accounts transactions.
* `/assets/{asset-id}/transactions` - Search for a prefix for a specific Asset Id.
* `/transactions` - Search all Transactions for a specific transaction note field prefix.

To search for a specific prefix use the `note-prefix` parameter. The value needs to be base64 encoded to return results. For example, if the contents of the note field started with the string “showing prefix searches”, encoding the beginning of that sentence using python like the following:

``` bash
$ python3 -c "import base64;print(base64.b64encode('showing prefix'.encode()))"
```

This will return an encoded value of `c2hvd2luZyBwcmVmaXg=`.  This value can then be passed to the search. To search all transactions use the following command.

```javascript tab="JavaScript"
// /indexer/python/SearchTransactionsNote.js
(async () => {
    //let s = buffer.toString('base64');   
    let s = "c2hvd2luZyBwcmVmaXg=";
    let transactionInfo = await indexerClient.searchForTransactions()
        .notePrefix(s).do();
    console.log("Information for Transaction search: " + JSON.stringify(transactionInfo, undefined, 2));
})().catch(e => {
    console.log(e);
    console.trace();
});
```

```python tab="Python"
# /indexer/python/search_transactions_note.py

import base64
encodednote = base64.b64encode('showing prefix'.encode())
decodednote = base64.b64decode(encodednote)

response = myindexer.search_transactions(
    note_prefix=decodednote)

print("note_prefix = " +
      json.dumps(response, indent=2, sort_keys=True))
```

```java tab="Java"
// /indexer/java/SearchTransactionsNote.java
    public static void main(String args[]) throws Exception {
        SearchTransactionsNote ex = new SearchTransactionsNote();
        IndexerClient indexerClientInstance = (IndexerClient)ex.connectToNetwork();
        byte[] encodednote = Encoder.decodeFromBase64("c2hvd2luZyBwcmVmaXg="); // "showing prefix"
        String response = indexerClientInstance
                .searchForTransactions()
                .notePrefix(encodednote).execute().toString();
        JSONObject jsonObj = new JSONObject(response.toString());
        System.out.println("Transaction Info: " + jsonObj.toString(2)); // pretty print json
    }
```

```go tab="Go"
// Parameters
var minAmount uint64 = 10
var data = "showing prefix"
var encodedNote = base64.StdEncoding.EncodeToString([]byte(data))

// Query
result, err := indexerClient.SearchForTransactions().NotePrefix([]byte(data)).Do(context.Background())

// Print results
JSON, err := json.MarshalIndent(result, "", "\t")
fmt.Printf(string(JSON) + "\n")
}

```

```bash tab="cURL"
python3 -c "import base64;print(base64.b64encode('showing prefix'.encode()))"
$ curl "localhost:8980/v2/transactions?note-prefix=c2hvd2luZyBwcmVmaXg=" | json_pp
```

Results
``` bash
{
   "current-round" : 7087347,
   "next-token" : "qCFsAAAAAAAAAAAA",
   "transactions" : [
      {
         "note" : "c2hvd2luZyBwcmVmaXggc2VhcmNoZXM=",
         "round-time" : 1591374973,
         "intra-round-offset" : 0,
         "close-rewards" : 0,
         "id" : "KWLZ4USIJC7P2WFWFA7SKGUDPBF5FOYS7OZ3I6T5ZBBJATRMNWLQ",
         "tx-type" : "pay",
         "payment-transaction" : {
            "receiver" : "UF7ATOM6PBLWMQMPUQ5QLA5DZ5E35PXQ2IENWGZQLEJJAAPAPGEGC3ZYNI",
            "close-amount" : 0,
            "amount" : 1000000
         },
         "confirmed-round" : 7086504,
         "last-valid" : 7087499,
         "sender-rewards" : 14650,
         "fee" : 1000,
         "closing-amount" : 0,
         "sender" : "IAMIRIFW3ERXIMR5LWNYHNK7KRTESUGS4QHOPKF2GL3CLHWWGW32XWB7OI",
         "receiver-rewards" : 0,
         "first-valid" : 7086499,
         "signature" : {
            "sig" : "H8I6wnDdTO+q8N8t0H35GfphOkbaHoJ90gzz8YYX8SV1pUXXXbx0hvU0AT+r8MaJB09FafhQQfkl2HfoII0EBw=="
         }
      }
   ]
}
```

# Searching Accounts
The ‘/accounts’ call can be used to search for accounts on the Algorand blockchain. This query provides two main parameters for returning accounts with specific balances. These two calls are `currency-greater-than` and `currency-less-than` which returns all accounts with balances that match the criteria. By default, the currency these parameters look for is the Algo and the values are specified in microAlgos. This behavior can be changed by supplying the `asset-id` parameter which specifies the asset to search accounts for in the ledger. For example to search accounts that have the Tether UDSt token the following command would be used.

```javascript tab="JavaScript"
// /indexer/javascript/AccountsAssetID.js
(async () => {
    let assetIndex = 312769;
    let accountInfo = await indexerClient.searchAccounts()
        .assetID(assetIndex).do();
    console.log("Information for account info for Asset: " + JSON.stringify(accountInfo, undefined, 2));
})().catch(e => {
    console.log(e);
    console.trace();
});
```

```python tab="Python"
# /indexer/python/accounts_assetid.py

response = myindexer.accounts(
    asset_id=312769)
print("Account Info: " + json.dumps(response, indent=2, sort_keys=True))
```

```java tab="Java"
// /indexer/java/AccountsAssetID.java
    public static void main(String args[]) throws Exception {
        AccountsAssetID ex = new AccountsAssetID();
        IndexerClient indexerClientInstance = (IndexerClient)ex.connectToNetwork();
        Long asset_id = Long.valueOf(312769);
        String response = indexerClientInstance.searchForAccounts().assetId(asset_id).execute().toString();
        JSONObject jsonObj = new JSONObject(response.toString());
        System.out.println("Pretty Print of Account for Asset: " + jsonObj.toString(2)); // pretty print json
    }
```

```go tab="Go"
// query parameters
var assetId uint64 = 312769
var minBalance uint64 = 100

// Lookup accounts with minimum balance of asset
result, err := indexerClient.LookupAssetBalances(assetId).CurrencyGreaterThan(minBalance).Do(context.Background())

// Print the results
JSON, err := json.MarshalIndent(result, "", "\t")
fmt.Printf(string(JSON) + "\n")
```

```bash tab="cURL"
$ curl localhost:8980/v2/accounts?asset-id=312769
```

This search can be further refined to search for accounts that have a balance greater than 100 USDt by using the following query.

```javascript tab="JavaScript"
// /indexer/javascript/AccountsAssetIDMinBalance.js
(async () => {
    let assetIndex = 312769;
    let currencyGreater = 100;
    let accountInfo = await indexerClient.searchAccounts()
        .assetID(assetIndex)
        .currencyGreaterThan(currencyGreater).do();
    console.log("Information for Account Info for Asset: " + JSON.stringify(accountInfo, undefined, 2));
})().catch(e => {
    console.log(e);
    console.trace();
});
```

```python tab="Python"
# /indexer/python/accounts_assets_min_balance.py

# gets accounts with a min balance of 100 that have a particular AssetID
response = myindexer.accounts(
    asset_id=312769, min_balance=100)
print("Account Info: " + json.dumps(response, indent=2, sort_keys=True))
```

```java tab="Java"
// /indexer/java/AccountsAssetIDMinBalance.java
    public static void main(String args[]) throws Exception {
        AccountsAssetIDMinBalance ex = new AccountsAssetIDMinBalance();
        IndexerClient indexerClientInstance = (IndexerClient)ex.connectToNetwork();
        Long asset_id = Long.valueOf(312769);
        Long currencyGreaterThan = Long.valueOf(100);
        // searches for asset greater than currencyGreaterThan
        String response = indexerClientInstance.searchForAccounts()
                .assetId(asset_id)
                .currencyGreaterThan(currencyGreaterThan).execute().toString();
        JSONObject jsonObj = new JSONObject(response.toString());
        System.out.println("Account Info for Asset Min Balance: " + jsonObj.toString(2)); // pretty print json
    }
```

```go tab="Go"
// query parameters
var assetId uint64 = 312769
var minBalance uint64 = 100

// Lookup accounts with minimum balance of asset
result, err := indexerClient.LookupAssetBalances(assetId).CurrencyGreaterThan(minBalance).Do(context.Background())

// Print the results
JSON, err := json.MarshalIndent(result, "", "\t")
fmt.Printf(string(JSON) + "\n")
```

```bash tab="cURL"
$ curl localhost:8980/v2/accounts?asset-id=312769\&currency-greater-than=100
```

Results
``` bash
{
   "accounts" : [
      {
         "rewards" : 0,
         "reward-base" : 101609,
         "amount" : 3578058672,
         "status" : "Offline",
         "amount-without-pending-rewards" : 3496490000,
         "pending-rewards" : 81568672,
         "address" : "AAABLBTPDNMMT7VXHGTYLWY6WA5CXG7SSWSF5KXN6P7SRPJESNL6BTB5MY",
         "round" : 7069819
      },
.
.
   ],
   "current-round" : 7069819
}
```

This call also supports the pagination mechanism described in [Paginated Results](/#paginated-results) using the `next` and `limit` query parameters. This call also supports [Historical Data Searches](#historical-data-searches) if the Indexer is configured for the `/accounts` call.

This call returns a list of accounts with associated data, the round number the results were calculated for and optionally the `next-token` value if you are using pagination.

# Getting an Account
The `/accounts/{account-id}` can be used to look up ledger data for a specific account.
For example:

```javascript tab="JavaScript"
// /indexer/javascript/AccountInfo.js
(async () => {
    let acct = "7WENHRCKEAZHD37QMB5T7I2KWU7IZGMCC3EVAO7TQADV7V5APXOKUBILCI";
    let accountInfo = await indexerClient.lookupAccountByID(acct).do();
    console.log("Information for Account: " + JSON.stringify(accountInfo, undefined, 2));
})().catch(e => {
    console.log(e);
    console.trace();
});
```

```python tab="Python"
# /indexer/python/account_info.py

# gets account
response = myindexer.account_info(
    address="TDO7JWA77FH3T2HP5ZOZWFKUQDQEAPD25HDKDVEAAWQKBWTMNMYRXOOYGA")
print("Account Info: " + json.dumps(response, indent=2, sort_keys=True))
```

```java tab="Java"
// /indexer/java/AccountInfo.java
    public static void main(String args[]) throws Exception {
        AccountInfo ex = new AccountInfo();
        IndexerClient indexerClientInstance = (IndexerClient)ex.connectToNetwork();
        Address account = new Address("7WENHRCKEAZHD37QMB5T7I2KWU7IZGMCC3EVAO7TQADV7V5APXOKUBILCI");
         String response = indexerClientInstance.lookupAccountByID(account).execute().toString();
        JSONObject jsonObj = new JSONObject(response.toString());
        System.out.println("Account Info: " + jsonObj.toString(2)); // pretty print json
    }
```

```go tab="Go"
// query parameters
var accountID = "TDO7JWA77FH3T2HP5ZOZWFKUQDQEAPD25HDKDVEAAWQKBWTMNMYRXOOYGA"

// Lookup account
_, result, err := indexerClient.LookupAccountByID(accountID).Do(context.Background())

// Print the results
JSON, err := json.MarshalIndent(result, "", "\t")
fmt.Printf(string(JSON) + "\n")
```

``` bash tab="cURL"
$ $ curl localhost:8980/v2/accounts/TDO7JWA77FH3T2HP5ZOZWFKUQDQEAPD25HDKDVEAAWQKBWTMNMYRXOOYGA | json_pp
```

Results
``` bash
{
   "account" : {
      "round" : 7069909,
      "amount-without-pending-rewards" : 999000,
      "address" : "TDO7JWA77FH3T2HP5ZOZWFKUQDQEAPD25HDKDVEAAWQKBWTMNMYRXOOYGA",
      "pending-rewards" : 0,
      "rewards" : 0,
      "reward-base" : 123047,
      "amount" : 999000,
      "assets" : [
         {
            "creator" : "",
            "asset-id" : 163650,
            "is-frozen" : false,
            "amount" : 1000000
         }
      ],
      "sig-type" : "sig",
      "status" : "Offline"
   },
   "current-round" : 7069909
}
```

The ledger data will include both Algo balance and if the account has any assets with their associated balances. This call supports [Historical Data Searches](#historical-data-searches) with the `round` parameter which will return ledger data for the account at a specific round.

# Search Transactions for a Specific Account
The `/accounts/{account-id}/transactions` REST call provides a powerful mechanism for searching for specific transactions for a given account. 

## Date-Time
The range of transactions to be searched can be restricted based on the time by using the `before-time` and `after-time` parameters. These parameters must be [rfc3339](http://www.faqs.org/rfcs/rfc3339.html) formatted date-time strings. For example, the following query searches for all transactions that occurred after 10 am (Zulu/UTC) minus 5 hours (EST).

```javascript tab="JavaScript"
// /indexer/javascript/SearchTxAddressTime.js
(async () => {
    let address = "XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4";
    let start_time = "2020-06-03T10:00:00-05:00"; 
    let response = await indexerClient.searchForTransactions()
        .address(address)
        .afterTime(start_time).do();
    console.log("start_time: 06/03/2020 11:00:00 = " + JSON.stringify(response, undefined, 2));
    }   
)().catch(e => {
    console.log(e);
    console.trace();
});
```

```python tab="Python"
# /indexer/python/search_tx_address_time.py

# get transaction at specific time
response = myindexer.search_transactions_by_address(
    address="XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4", start_time="2020-06-03T10:00:00-05:00")

print("Transaction Start Time 2020-06-03T10:00:00-05:00 = " +
      json.dumps(response, indent=2, sort_keys=True))

```

```java tab="Java"
// /indexer/java/SearchTxAddressTime.java
    public static void main(String args[]) throws Exception {
        SearchTxAddressTime ex = new SearchTxAddressTime();
        IndexerClient indexerClientInstance = (IndexerClient) ex.connectToNetwork();
        Address account = new Address("XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4");
    	SimpleDateFormat sdf = new SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ssXXX");
        Date start_time = sdf.parse("2020-06-03T10:00:00-05:00");
        String response = indexerClientInstance
                .searchForTransactions()
                .address(account)
                .afterTime(start_time).execute().toString();
        JSONObject jsonObj = new JSONObject(response.toString());
        System.out.println("start_time: 06/03/2020 11:00:00-05:00 = " + jsonObj.toString(2)); // pretty print json
    }
```

```go tab="Go"
// query parameters
var startTime = "2020-06-03T10:00:00-05:00"
address, _ := types.DecodeAddress("XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4")

// Query
result, err := indexerClient.SearchForTransactions().Address(address).AfterTimeString(startTime).Do(context.Background())

// Print the results
JSON, err := json.MarshalIndent(result, "", "\t")
fmt.Printf(string(JSON) + "\n")
```

``` bash tab="cURL"
$ $ curl localhost:8980/v2/accounts/XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4/transactions?after-time=2020-06-03T10:00:00-05:00|json_pp
```

Results
``` bash
{
   "transactions" : [
      {
         "closing-amount" : 0,
         "intra-round-offset" : 0,
         "round-time" : 1591209213,
         "id" : "QZS3B2XBBS47S6X5CZGKKC2FC7HRP5VJ4UNS7LPGHP24DUECHAAA",
         "sender-rewards" : 0,
         "close-rewards" : 0,
         "signature" : {
            "sig" : "CUogpD2S7LErOsD8BLRMEUMzKsbrSohkCJYv/7+IeqQfo6T+oClTn0zU4hc8AKSlzm7k9d6qsPP2t53i55NfAw=="
         },
         "first-valid" : 7048875,
         "sender" : "RL6VDLXCN5G7N2GRTS7YLVDSFT4PVBBUOVTVS7T26OQ5MLXYQKRMI5ADXY",
         "fee" : 1000,
         "last-valid" : 7049875,
         "tx-type" : "axfer",
         "confirmed-round" : 7048877,
         "receiver-rewards" : 0,
         "asset-transfer-transaction" : {
            "amount" : 0,
            "receiver" : "XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4",
            "close-to" : "XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4",
            "close-amount" : 0,
            "asset-id" : 312769
         }
      }
   ],
   "current-round" : 7085397
} 
```

## Round Range
Transaction searches can also be restricted to round ranges using the `min-round` and `max-round` parameters. 

```javascript tab="JavaScript"
// /indexer/javascript/SearchTxAddressBlockRange.js
(async () => {
    let address = "XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4";
    let min_round = 7048876;
    let max_round = 7048878;    
    let response = await indexerClient.searchForTransactions()
        .address(address).maxRound(max_round)
        .minRound(min_round).do();
    console.log("Information for Transaction search: " + JSON.stringify(response, undefined, 2));
    }  
)().catch(e => {
    console.log(e);
    console.trace();
});
```

```python tab="Python"
# /indexer/python/search_tx_address_block_range.py

response = myindexer.search_transactions_by_address(
    address="XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4", min_round=7048876, max_round=7048878)

print("min-max rounds: 7048876-7048878 = " +
      json.dumps(response, indent=2, sort_keys=True))
```

```java tab="Java"
// /indexer/java/SearchTxAddressBlockRange.java
    public static void main(String args[]) throws Exception {
        SearchTxAddressBlockRange ex = new SearchTxAddressBlockRange();
        IndexerClient indexerClientInstance = (IndexerClient) ex.connectToNetwork();
        Address account = new Address("XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4");
        Long min_round = Long.valueOf(7048876);
        Long max_round = Long.valueOf(7048878);       
        String response = indexerClientInstance.searchForTransactions().address(account)
                    .minRound(min_round).maxRound(max_round).execute().toString();
        JSONObject jsonObj = new JSONObject(response.toString());
        System.out.println("Transaction Info: " + jsonObj.toString(2)); // pretty print json
    }
```

```go tab="Go"
// query parameters
var minRound uint64 = 7048876
var maxRound uint64 = 7048878
address, _ := types.DecodeAddress("XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4")

// query
result, err := indexerClient.SearchForTransactions().Address(address).MinRound(minRound).MaxRound(maxRound).Do(context.Background())

// Print the results
JSON, err := json.MarshalIndent(result, "", "\t")
fmt.Printf(string(JSON) + "\n")
```

``` bash tab="cURL"
$ $ curl "localhost:8980/v2/accounts/XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4/transactions?min-round=7048876&max-round=7048878" | json_pp
```

In addition, you can specify a specific round by using the round parameter.

```javascript tab="JavaScript"
// /indexer/javascript/SearchTxAddressBlock.js
(async () => {
    let address = "XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4";
    let round = 7048877;
    let response = await indexerClient.searchForTransactions()
        .address(address)
        .round(round).do();
    console.log("Information for Transaction search: " + JSON.stringify(response, undefined, 2));
    }   
)().catch(e => {
    console.log(e);
    console.trace();
});
```

```python tab="Python"
# /indexer/python/search_tx_address_block.py

response = myindexer.search_transactions_by_address(
    address="XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4", block=7048877)

print("block: 7048877 = " +
      json.dumps(response, indent=2, sort_keys=True))
```

```java tab="Java"
// /indexer/java/SearchTxAddressBlock.java
    public static void main(String args[]) throws Exception {
        SearchTxAddressBlock ex = new SearchTxAddressBlock();
        IndexerClient indexerClientInstance = (IndexerClient) ex.connectToNetwork();
        Address account = new Address("XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4");
        Long block = Long.valueOf(7048877);             
        String response = indexerClientInstance.searchForTransactions().address(account)
                    .round(block).execute().toString();
        JSONObject jsonObj = new JSONObject(response.toString());
        System.out.println("Transaction Info: " + jsonObj.toString(2)); // pretty print json
    }
```

```go tab="Go"
// query parameters
var round uint64 = 7048877
address, _ := types.DecodeAddress("XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4")

// query
result, err := indexerClient.SearchForTransactions().Address(address).Round(Round).Do(context.Background())

// Print the results
JSON, err := json.MarshalIndent(result, "", "\t")
fmt.Printf(string(JSON) + "\n")

```

``` bash tab="cURL"
$ $ curl "localhost:8980/v2/accounts/XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4/transactions?round=7048877"
```

## Transaction ID
Searching for a specific transaction can be achieved by supplying the transaction id using the `txid` parameter.

```javascript tab="JavaScript"
// /indexer/javascript/SearchTxAddressTxId.js
(async () => {
    let address = "XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4";
    let txid = "QZS3B2XBBS47S6X5CZGKKC2FC7HRP5VJ4UNS7LPGHP24DUECHAAA"; 
    let response = await indexerClient.searchForTransactions()
        .address(address)
        .txid(txid).do();
    console.log("txid: QZS3B2XBBS47S6X5CZGKKC2FC7HRP5VJ4UNS7LPGHP24DUECHAAA = " + JSON.stringify(response, undefined, 2));
    }  
)().catch(e => {
    console.log(e);
    console.trace();
});
```

```python tab="Python"
# /indexer/python/search_tx_address_txId.py

response = myindexer.search_transactions_by_address(
    address="XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4",
    txid="QZS3B2XBBS47S6X5CZGKKC2FC7HRP5VJ4UNS7LPGHP24DUECHAAA")

print("txid: QZS3B2XBBS47S6X5CZGKKC2FC7HRP5VJ4UNS7LPGHP24DUECHAAA = " +
      json.dumps(response, indent=2, sort_keys=True))
```

```java tab="Java"
// /indexer/java/SearchTxAddressTxId.java
    public static void main(String args[]) throws Exception {
        SearchTxAddressTxId ex = new SearchTxAddressTxId();
        IndexerClient indexerClientInstance = (IndexerClient) ex.connectToNetwork();
        Address account = new Address("XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4");
        String txid = "QZS3B2XBBS47S6X5CZGKKC2FC7HRP5VJ4UNS7LPGHP24DUECHAAA";
        String response = indexerClientInstance.searchForTransactions().address(account).txid(txid).execute().toString();
        JSONObject jsonObj = new JSONObject(response.toString());
        System.out.println("txid: QZS3B2XBBS47S6X5CZGKKC2FC7HRP5VJ4UNS7LPGHP24DUECHAAA = " + jsonObj.toString(2)); // pretty print json
    }
```

```go tab="Go"
// query parameters
var txID = "QZS3B2XBBS47S6X5CZGKKC2FC7HRP5VJ4UNS7LPGHP24DUECHAAA"
address, _ := types.DecodeAddress("XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4")

// Query
result, err := indexerClient.SearchForTransactions().Address(address).TXID(txID).Do(context.Background())

// Print the results
JSON, err := json.MarshalIndent(result, "", "\t")
fmt.Printf(string(JSON) + "\n")
```

``` bash tab="cURL"
$ $ curl "localhost:8980/v2/accounts/XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4/transactions?txid=QZS3B2XBBS47S6X5CZGKKC2FC7HRP5VJ4UNS7LPGHP24DUECHAAA" | json_pp
```

## Transaction Type
You can also search for specific transaction types that are described in the [transaction structure](https://developer.algorand.org/docs/features/transactions/#transaction-types) documentation. The Indexer supports looking for `pay`, `keyreg`, `acfg`, `axfer` and `afrz` transaction types. To search for a specific type of transaction use the `tx-type` parameter. The following example searches for the asset creation transaction for the DevDocsCoin.

```javascript tab="JavaScript"
// /indexer/javascript/SearchTxAddresstxntype.js
(async () => {
    let address = "SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU";
    let txn_type = "acfg"; 
    let response = await indexerClient.searchForTransactions()
        .address(address)
        .txType(txn_type).do();
    console.log("txn_type: acfg = " + JSON.stringify(response, undefined, 2));
    }  
)().catch(e => {
    console.log(e);
    console.trace();
});
```

```python tab="Python"
# /indexer/python/search_tx_address_txntype.py

response = myindexer.search_transactions_by_address(
    address="SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU",
    txn_type="acfg")

print("txn_type: acfg = " +
      json.dumps(response, indent=2, sort_keys=True))
```

```java tab="Java"
// /indexer/java/SearchTxAddresstxntype.java
    public static void main(String args[]) throws Exception {
        SearchTxAddresstxntype ex = new SearchTxAddresstxntype();
        IndexerClient indexerClientInstance = (IndexerClient) ex.connectToNetwork();
        Address account = new Address("SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU");
        TxType txType = TxType.ACFG;
        String response = indexerClientInstance.searchForTransactions().address(account).txType(txType).execute().toString();
        JSONObject jsonObj = new JSONObject(response.toString());
        System.out.println("txn_type: acfg = " + jsonObj.toString(2)); // pretty print json
    }
```

```go tab="Go"
// query parameters
var txType = "acfg"
address, _ := types.DecodeAddress("SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU")

// Query
result, err := indexerClient.SearchForTransactions().Address(address).TxType(txType).Do(context.Background())

// Print the results
JSON, err := json.MarshalIndent(result, "", "\t")
fmt.Printf(string(JSON) + "\n")
```

``` bash tab="cURL"
$ $ curl "localhost:8980/v2/accounts/SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU/transactions?tx-type=acfg" | json_pp
```

Results
``` bash
{
   "current-round" : 7086205,
   "transactions" : [
      {
         "last-valid" : 7087123,
         "close-rewards" : 0,
         "asset-config-transaction" : {
            "params" : {
               "unit-name" : "HiWorld",
               "reserve" : "SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU",
               "manager" : "SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU",
               "total" : 1000000,
               "url" : "https://bit.ly/374zKN2",
               "creator" : "SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU",
               "clawback" : "SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU",
               "default-frozen" : false,
               "name" : "DevDocsCoin",
               "decimals" : 0,
               "freeze" : "SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU"
            },
            "asset-id" : 0
         },
         "signature" : {
            "sig" : "uH2LhX3v5OFD+KQ61GZS2COlAgM55o2IKpfFox0ZEC079CyT3dxKfYEAhD2yTQsRV2bn83Ou7tNfFQDjWE+CAQ=="
         },
         "confirmed-round" : 7086124,
         "closing-amount" : 0,
         "id" : "L3W6Y6VUNEENZWUD4BSBFFWVQS6OEXMKKVJYE2MG7G27V647VILA",
         "note" : "lhEy6PuICYg=",
         "intra-round-offset" : 1,
         "sender-rewards" : 0,
         "tx-type" : "acfg",
         "created-asset-index" : 2044572,
         "first-valid" : 7086123,
         "receiver-rewards" : 0,
         "fee" : 1000,
         "round-time" : 1591373291,
         "sender" : "SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU"
      }
   ]
}
```

## Currency
Searching Transactions that are of certain values can be created by using the `currency-greater-than` and `currency-less-than` parameters. By default, these values are in microAlgos, but if this is used in conjunction with the `asset-id` parameter the currencies will refer to the specific asset. Searching for transactions greater than 50 DevDocsCoin (asset id=2044572) can be done with the following:

```javascript tab="JavaScript"
// /indexer/javascript/SearchTxAddressAsset.js
(async () => {
    let address = "SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU";
    let asset_id = 2044572;
    let min_amount = 50;

    let response = await indexerClient.searchForTransactions()
            .address(address)
            .currencyGreaterThan(min_amount)
            .assetID(asset_id).do();
    console.log("Information for Transaction search: " + JSON.stringify(response, undefined, 2));
    }   
)().catch(e => {
    console.log(e);
    console.trace();
});
```

```python tab="Python"
# /indexer/python/accounts_assets_min_balance.py

# gets assets with a min balance of 50 for AssetID
response = myindexer.asset_balances(
        address="SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU",asset_id=2044572, min_balance=50)

print("Asset Balances :" + json.dumps(response, indent=2, sort_keys=True))
```

```java tab="Java"
// /indexer/java/SearchTxAddressAsset.java
    public static void main(String args[]) throws Exception {
        SearchTxAddressAsset ex = new SearchTxAddressAsset();
        IndexerClient indexerClientInstance = (IndexerClient) ex.connectToNetwork();
        Address account = new Address("SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU");
        Long asset_id = Long.valueOf(2044572);   
        Long min_amount = Long.valueOf(50);              
        String response = indexerClientInstance.searchForTransactions().address(account)
                    .currencyGreaterThan(min_amount).assetId(asset_id).execute().toString();
        JSONObject jsonObj = new JSONObject(response.toString());
        System.out.println("Transaction Info: " + jsonObj.toString(2)); // pretty print json
    }
```

```go tab="Go"
// query parameters
var assetId uint64 = 2044572
var minBalance uint64 = 50
address, _ := types.DecodeAddress("SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU")

// Lookup accounts with minimum balance of asset
result, err := indexerClient.LookupAssetBalances(assetId).CurrencyGreaterThan(minBalance).Do(context.Background())

// Print the results
JSON, err := json.MarshalIndent(result, "", "\t")
fmt.Printf(string(JSON) + "\n")
```

``` bash tab="cURL"
$ $ curl "localhost:8980/v2/accounts/SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU/transactions?asset-id=2044572&currency-greater-than=50" | json_pp
```

Results
``` bash
{
   "transactions" : [
      {
         "sender" : "SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU",
         "asset-transfer-transaction" : {
            "amount" : 100,
            "asset-id" : 2044572,
            "close-amount" : 0,
            "receiver" : "UF7ATOM6PBLWMQMPUQ5QLA5DZ5E35PXQ2IENWGZQLEJJAAPAPGEGC3ZYNI"
         },
         "tx-type" : "axfer",
         "confirmed-round" : 7086187,
         "signature" : {
            "sig" : "ZTr7MGBt/TIgXBDMFylqwmrzsxTKGhcs/JOpbIdsmuMAVXZy9IDTiQR1pkA6jvqTWdt2kFbJJ+beJoFUyeM5CQ=="
         },
         "fee" : 1000,
         "last-valid" : 7087184,
         "round-time" : 1591373570,
         "close-rewards" : 0,
         "first-valid" : 7086184,
         "intra-round-offset" : 0,
         "receiver-rewards" : 0,
         "closing-amount" : 0,
         "sender-rewards" : 0,
         "id" : "GVM44H3RWUE557GT4GEEGNLPL5RQZ62F5FK3N5YDKCOTR26XV64A"
      }
   ],
   "current-round" : 7087524,
   "next-token" : "ayBsAAAAAAAAAAAA"
}
```

## Signature Type
Transaction searches can also look for specific [signature types](https://developer.algorand.org/docs/reference/transactions/#signed-transaction), which can be standard signatures, multi-signatures, or logic signatures. To search for transactions with a specific signature type use the `sig-type` parameter and specify either `sig`, `msig`, or `lsig` and the parameter values. These are signatures that specifically sign Algorand transactions. The following query will return all transactions that signed with a multi-signature for the specific account.

```javascript tab="JavaScript"
// /indexer/javascript/SearchTxAddresssigtype.js
(async () => {
    let address = "XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4";
    let sig_type = "msig"; 
    let response = await indexerClient.searchForTransactions()
        .address(address)
        .sigType(sig_type).do();
    console.log("Information for Transaction search: " + JSON.stringify(response, undefined, 2));
    }
   
)().catch(e => {
    console.log(e);
    console.trace();
});
```

```python tab="Python"
# /indexer/python/search_tx_address_sigtype.py

response = myindexer.search_transactions_by_address(
    address="XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4", sig_type="msig")

print("sig_type: msig = " +
      json.dumps(response, indent=2, sort_keys=True))
```

```java tab="Java"
// /indexer/java/SearchTxAddresssigtype.java
    public static void main(String args[]) throws Exception {
        SearchTxAddresssigtype ex = new SearchTxAddresssigtype();
        IndexerClient indexerClientInstance = (IndexerClient) ex.connectToNetwork();
        Address account = new Address("XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4");
        SigType sig_type = SigType.MSIG;    
        String response = indexerClientInstance.searchForTransactions().address(account).sigType(sig_type).execute().toString();
        JSONObject jsonObj = new JSONObject(response.toString());
        System.out.println("Transaction Info SigType msig: " + jsonObj.toString(2)); // pretty print json
    }
```

```go tab="Go"
// query parameters
var sigType = "msig"
address, _ := types.DecodeAddress("XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4")

// Query
result, err := indexerClient.SearchForTransactions().Address(address).SigType(sigType).Do(context.Background())

// Print the results
JSON, err := json.MarshalIndent(result, "", "\t")
fmt.Printf(string(JSON) + "\n")
```

``` bash tab="cURL"
$ $ curl "localhost:8980/v2/accounts/XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4/transactions?sig-type=msig" | json_pp
```

The call also supports [Paginated Results](#paginated-results) and [Note Field Searching](#note-field-searching).

This call returns a list of transactions, the round the results were calculated in and if using pagination the next token to retrieve more results.

# Search Assets 
The Indexer provides the `/assets` REST call to search the blockchain for specific assets. The call supports searching based on the asset id, the name of the asset, the unit name of the asset, and by the creator of the asset. For example, to search for DevDocsCoin the `name` parameter should be used. Note that this is a non-case sensitive search.

```javascript tab="JavaScript"
// /indexer/javascript/SearchAssetName.js
(async () => {
    let name = "DevDocsCoin";
    let assetInfo = await indexerClient.searchForAssets()
        .name(name).do();
    console.log("Information for Asset Name: " + JSON.stringify(assetInfo, undefined, 2));
})().catch(e => {
    console.log(e);
    console.trace();
});
```

```python tab="Python"
# /indexer/python/search_assets_name.py

response = myindexer.search_assets(
    name="DevDocsCoin")
print("Asset Name Info: " + json.dumps(response, indent=2, sort_keys=True))
```

```java tab="Java"
// /indexer/java/SearchAssetsName.java
    public static void main(String args[]) throws Exception {
        SearchAssetsName ex = new SearchAssetsName();
        IndexerClient indexerClientInstance = (IndexerClient)ex.connectToNetwork();
        String name = "DevDocsCoin";        
        String response = indexerClientInstance.searchForAssets()
                        .name(name).execute().toString();
        JSONObject jsonObj = new JSONObject(response.toString());
        System.out.println("Asset Info for Name: " + jsonObj.toString(2)); // pretty print json
    }
```

```go tab="Go"
// query parameters
var assetName = "DevDocsCoin"

// Lookup asset
_, result, err := indexerClient.SearchForAssets().Name(assetName).Do(context.Background())

// Search asset by name
JSON, err := json.MarshalIndent(result, "", "\t")
fmt.Printf(string(JSON) + "\n")
```

``` bash tab="cURL"
$ $ curl "localhost:8980/v2/assets?name=DevDocsCoin" | json_pp
```

Results
```bash 
{
   "current-round" : 7087643,
   "assets" : [
      {
         "params" : {
            "name" : "DevDocsCoin",
            "freeze" : "SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU",
            "decimals" : 0,
            "total" : 1000000,
            "default-frozen" : false,
            "creator" : "SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU",
            "manager" : "SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU",
            "unit-name" : "HiWorld",
            "reserve" : "SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU",
            "clawback" : "SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU",
            "url" : "https://bit.ly/374zKN2"
         },
         "index" : 2044572
      }
   ]
}
```

This call returns a list of assets and the round the results were calculated in. Additionally, the next token is also returned if using [Paginated Results](#paginated-results).

# Getting an Asset 
To get the details of a specific asset the indexer provides the `/assets/{asset-id}` REST call.  This call takes no parameters as the asset id is passed in the URL. This call returns the details of the asset and the round the results were calculated in. 

```javascript tab="JavaScript"
// /indexer/javascript/SearchAssets.js
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

```python tab="Python"
# /indexer/python/search_assets.py

response = myindexer.search_assets(
    asset_id=2044572)

print("Asset Info: " + json.dumps(response, indent=2, sort_keys=True))
```

```java tab="Java"
// /indexer/java/SearchAssets.java
    public static void main(String args[]) throws Exception {
        SearchAssets ex = new SearchAssets();
        IndexerClient indexerClientInstance = (IndexerClient)ex.connectToNetwork();
        Long asset_id = Long.valueOf(2044572);        
        String response = indexerClientInstance.searchForAssets()
                        .assetId(asset_id).execute().toString();
        JSONObject jsonObj = new JSONObject(response.toString());
        System.out.println("Asset Info: " + jsonObj.toString(2)); // pretty print json
    }
```

```go tab="Go"
// query parameters
var assetID uint64 = 2044572

// Query
_, result, err := indexerClient.LookupAssetByID(assetID).Do(context.Background())

// Print the results
JSON, err := json.MarshalIndent(result, "", "\t")
fmt.Printf(string(JSON) + "\n")
```

``` bash tab="cURL"
$ $ curl "localhost:8980/v2/assets/2044572" | json_pp
```

Results
``` bash
{
   "asset" : {
      "index" : 2044572,
      "params" : {
         "manager" : "SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU",
         "url" : "https://bit.ly/374zKN2",
         "reserve" : "SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU",
         "default-frozen" : false,
         "freeze" : "SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU",
         "total" : 1000000,
         "name" : "DevDocsCoin",
         "decimals" : 0,
         "clawback" : "SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU",
         "creator" : "SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU",
         "unit-name" : "HiWorld"
      }
   },
   "current-round" : 7088298
}
```

# Searching for Accounts based on Asset 
The Indexer provides the `/assets/{asset-id}/balances` REST API call to search for accounts that transact in a specific Asset. The call returns a list of balances, with their specific amount, the address holding the asset, and whether the account is frozen. Additionally, the round the results were calculated is returned.

```javascript tab="JavaScript"
// /indexer/javascript/AssetsBalances.js
(async () => {
    let assetIndex = 2044572;
    let assetInfo = await indexerClient.lookupAssetBalances(assetIndex).do();
    console.log("Information for Asset: " + JSON.stringify(assetInfo, undefined, 2));
})().catch(e => {
    console.log(e);
    console.trace();
});
```

```python tab="Python"
# /indexer/python/assets_balances.py

response = myindexer.asset_balances(
    asset_id=2044572)
print("Asset Balance: " + json.dumps(response, indent=2, sort_keys=True))
```

```java tab="Java"
// /indexer/java/AssetsBalances.java
    public static void main(String args[]) throws Exception {
        AssetsBalances ex = new AssetsBalances();
        IndexerClient indexerClientInstance = (IndexerClient)ex.connectToNetwork();
        Long asset_id = Long.valueOf(2044572);
        // searhes for asset greater than currencyGreaterThan
        String response = indexerClientInstance.lookupAssetBalances(asset_id).execute().toString();
        JSONObject jsonObj = new JSONObject(response.toString());
        System.out.println("Asset Info: " + jsonObj.toString(2)); // pretty print json
    }
```

```go tab="Go"
// query parameters
const assetId = 2044572

// Lookup balance of asset
result, err := indexerClient.LookupAssetBalances(assetId).Do(context.Background())

// Print the results
JSON, err := json.MarshalIndent(result, "", "\t")
fmt.Printf(string(JSON) + "\n")
```

```bash tab="cURL"
$ $ curl "localhost:8980/v2/assets/2044572/balances" | json_pp
```

Results
``` bash
{
   "balances" : [
      {
         "is-frozen" : false,
         "amount" : 999900,
         "address" : "SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU"
      },
      {
         "is-frozen" : false,
         "amount" : 100,
         "address" : "UF7ATOM6PBLWMQMPUQ5QLA5DZ5E35PXQ2IENWGZQLEJJAAPAPGEGC3ZYNI"
      }
   ],
   "current-round" : 7088478
}
```

This call can be refined by looking for addresses based on the current amount using the `currency-greater-than` and `currency-less-than` parameters.

```javascript tab="JavaScript"
// /indexer/javascript/AssetsBalancesMinBalance.js
(async () => {
    let assetIndex = 2044572;
    let currencyGreater = 200;
    let assetInfo = await indexerClient.lookupAssetBalances(assetIndex)
        .currencyGreaterThan(currencyGreater).do();
    console.log("Information for Asset: " + JSON.stringify(assetInfo, undefined, 2));
})().catch(e => {
    console.log(e);
    console.trace();
});
```

```python tab="Python"
# /indexer/python/assets_balances_min_balance.py

response = myindexer.asset_balances(
    asset_id=2044572, min_balance=200)

print("Asset Balances :" + json.dumps(response, indent=2, sort_keys=True))
```

```java tab="Java"
// /indexer/java/AccountsAssetIDMinBalance.java
    public static void main(String args[]) throws Exception {
        AccountsAssetIDMinBalance ex = new AccountsAssetIDMinBalance();
        IndexerClient indexerClientInstance = (IndexerClient)ex.connectToNetwork();
        Long asset_id = Long.valueOf(312769);
        Long currencyGreaterThan = Long.valueOf(100);
        // searches for asset greater than currencyGreaterThan
        String response = indexerClientInstance.searchForAccounts()
                .assetId(asset_id)
                .currencyGreaterThan(currencyGreaterThan).execute().toString();
        JSONObject jsonObj = new JSONObject(response.toString());
        System.out.println("Account Info for Asset Min Balance: " + jsonObj.toString(2)); // pretty print json
    }
```

```go tab="Go"
// query parameters
var assetId uint64 = 2044572
var minBalance uint64 = 200

// Query
result, err := indexerClient.LookupAssetBalances(assetId).CurrencyGreaterThan(minBalance).Do(context.Background())

// Print the results
JSON, err := json.MarshalIndent(result, "", "\t")
fmt.Printf(string(JSON) + "\n")
```

``` bash tab="cURL"
$ $ curl "localhost:8980/v2/assets/2044572/balances?currency-greater-than=200" | json_pp
```

Results
``` bash
{
   "balances" : [
      {
         "is-frozen" : false,
         "amount" : 999900,
         "address" : "SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU"
      }
   ],
   "current-round" : 7088514
}
```

This call also supports [Paginated Results](#paginated-results) and [Historical Data Searches](#historical-data-searches).

# Searching for Transactions based on Asset 
The `/assets/{asset-id}/transactions` call is provided by the Indexer to search for transactions involving a specific Asset. This call contains many parameters to refine the search for specific values. These parameters function very similarly to the parameters described in Search Transactions for a Specific Account. The following table list the common search parameters between this call and the `/accounts/{account-id}/transactions` call with a link to the description of this parameter.

| Parameter  | Description Link  |
| ------------- | ------------- |
| after-time | [Search Transactions for a Specific Account - Date-Time](#date-time) |
| before-time | [Search Transactions for a Specific Account - Date-Time](#date-time) |
| max-round | [Search Transactions for a Specific Account - Round Range](#round-range) |
| min-round | [Search Transactions for a Specific Account - Round Range](#round-range) |
| round | [Search Transactions for a Specific Account - Round Range](#round-range) |
| currency-greater-than | [Search Transactions for a Specific Account - Currency](/#currency) |
| currency-less-than | [Search Transactions for a Specific Account - Currency](/#currency) |
| tx-type | [Search Transactions for a Specific Account - Transaction Type](#transaction-type) |
| sig- type | [Search Transactions for a Specific Account - Signature Type](#signature-type) |
| tx-id | [Search Transactions for a Specific Account - Transaction ID](#signature-type) |


This call also supports [Paginated Results](#paginated-results) and [Note Field Searching](#note-field-searching).

# Address and Role
When searching for transactions that involve a specific Asset you can search for specific accounts as well as the address role. The roles that can be searched for are `sender`, `receiver`, and `freeze-target`. For example, a specific receiver with a given address for an Asset can be searched using the following:

```javascript tab="JavaScript"
// /indexer/javascript/SearchAssetTransactionsRole.js
(async () => {
    let asset_id = 2044572;
    let address_role = "receiver";
    let address = "UF7ATOM6PBLWMQMPUQ5QLA5DZ5E35PXQ2IENWGZQLEJJAAPAPGEGC3ZYNI";
    let tracsactionInfo = await indexerClient.lookupAssetTransactions(asset_id)
        .addressRole(address_role)
        .address(address).do();
    console.log("Information for Transaction for Asset: " + JSON.stringify(tracsactionInfo, undefined, 2));
})().catch(e => {
    console.log(e);
    console.trace();
})
```

```python tab="Python"
# /indexer/python/search_asset_transactions_role.py

response = myindexer.search_asset_transactions(
    asset_id=2044572, address_role="receiver", address="UF7ATOM6PBLWMQMPUQ5QLA5DZ5E35PXQ2IENWGZQLEJJAAPAPGEGC3ZYNI")

print("Asset Transaction Info: " + json.dumps(response, indent=2, sort_keys=True))
```

```java tab="Java"
// /indexer/java/SearchAssetsTransactionsRole.java
    public static void main(String args[]) throws Exception {
        SearchAssetsTransactionsRole ex = new SearchAssetsTransactionsRole();
        IndexerClient indexerClientInstance = (IndexerClient)ex.connectToNetwork();
        Long asset_id = Long.valueOf(2044572);
        AddressRole addressRole = AddressRole.RECEIVER;
        Address account = new Address("UF7ATOM6PBLWMQMPUQ5QLA5DZ5E35PXQ2IENWGZQLEJJAAPAPGEGC3ZYNI");           
        String response = indexerClientInstance.searchForTransactions().address(account).assetId(asset_id).addressRole(addressRole).execute().toString();
        JSONObject jsonObj = new JSONObject(response.toString());
        System.out.println("Asset Info for Name: " + jsonObj.toString(2)); // pretty print json
    }
```

```go tab="Go"
// query parameters
var assetID uint64 = 2044572
address, _ := types.DecodeAddress("UF7ATOM6PBLWMQMPUQ5QLA5DZ5E35PXQ2IENWGZQLEJJAAPAPGEGC3ZYNI")

// Lookup asset
result, err := indexerClient.SearchForTransactions().AssetID(assetID).AddressRole(role).Address(address).Do(context.Background())

// Search asset trnsaction by role
JSON, err := json.MarshalIndent(result, "", "\t")
fmt.Printf(string(JSON) + "\n")
```

``` bash tab="cURL"
$ $ curl "localhost:8980/v2/assets/2044572/transactions?address-role=receiver&address=UF7ATOM6PBLWMQMPUQ5QLA5DZ5E35PXQ2IENWGZQLEJJAAPAPGEGC3ZYNI" | json_pp
```

When searching for the receiver like in the above example, a `closeto` transaction property is also returned. You can remove this by supplying the `exclude-close-to` parameter and setting its value to true.

# Retrieving Blocks 
The Indexer provides the `/blocks/{round-number}` API call to retrieve specific blocks from the blockchain. This call does not take any parameters and returns data associated with the block.

```javascript tab="JavaScript"
// /indexer/javascript/BlockInfo.js
(async () => {
    let block = 50;
    let blockInfo = await indexerClient.lookupBlock(block).do();
    console.log("Information for Block: " + JSON.stringify(blockInfo, undefined, 2));
})().catch(e => {
    console.log(e);
    console.trace();
});
```

```python tab="Python"
# /indexer/python/block_info.py

response = myindexer.block_info(
    block=555)
print("Block Info: " + json.dumps(response, indent=2, sort_keys=True))
```

```java tab="Java"
// /indexer/java/BlockInfo.java
    public static void main(String args[]) throws Exception {
        BlockInfo ex = new BlockInfo();
        IndexerClient indexerClientInstance = (IndexerClient)ex.connectToNetwork();
        Long block = Long.valueOf(50);
        String response = indexerClientInstance.lookupBlock(block).execute().toString();
        JSONObject jsonObj = new JSONObject(response.toString());
        System.out.println("Block Info: " + jsonObj.toString(2)); // pretty print json
    }
```

```go tab="Go"
// query parameters
var round uint64 = 555

// Lookup block
result, err := indexerClient.LookupBlock(round).Do(context.Background())

// Print the results
JSON, err := json.MarshalIndent(result, "", "\t")
fmt.Printf(string(JSON) + "\n")
```

``` bash tab="cURL"
$ $ curl "localhost:8980/v2/blocks/555" | json_pp
```

Results
``` bash
{
   "transactions" : [],
   "genesis-hash" : "wGHE2Pwdvd7S12BL5FaOP20EGYesN73ktiC1qzkkit8=",
   "genesis-id" : "mainnet-v1.0",
   "upgrade-state" : {
      "next-protocol-approvals" : 0,
      "next-protocol-vote-before" : 0,
      "current-protocol" : "https://github.com/algorandfoundation/specs/tree/5615adc36bad610c7f165fa2967f4ecfa75125f0",
      "next-protocol-switch-on" : 0
   },
   "timestamp" : 1560225075,
   "seed" : "TikvdKIdKqotCFyHsVUqA8vY0obbwRdhgv1US1ZJs7Y=",
   "previous-block-hash" : "Fe6FewPRhbQo4xsFEU/QqBjxnoN00DjYovoXa76Q9Ss=",
   "rewards" : {
      "rewards-residue" : 320011132,
      "rewards-level" : 11,
      "fee-sink" : "",
      "rewards-rate" : 20000000,
      "rewards-pool" : "737777777777777777777777777777777777777777777777777UFEJ2CI",
      "rewards-calculation-round" : 500000
   },
   "transactions-root" : "tGWqr06HVvsHC2hFmu2EhwYBSiMkvIoo5nWxAkUxsLo=",
   "txn-counter" : 0,
   "round" : 555,
   "upgrade-vote" : {
      "upgrade-approve" : false,
      "upgrade-delay" : 0
   }
}
```

# Search Transactions
The `/transactions` REST method is provided to allow searching all transactions that have occurred on the blockchain.

This call contains many parameters to refine the search for specific values. These parameters function very similarly to the parameters described in Search Transactions for a Specific Account. The following table list the common search parameters between this call and the `/accounts/{account-id}/transactions` call with a link to the description of this parameter.


| Parameter  | Description Link  |
| ------------- | ------------- |
| after-time | [Search Transactions for a Specific Account - Date-Time](#date-time) |
| before-time | [Search Transactions for a Specific Account - Date-Time](#date-time) |
| max-round | [Search Transactions for a Specific Account - Round Range](#round-range) |
| min-round | [Search Transactions for a Specific Account - Round Range](#round-range) |
| round | [Search Transactions for a Specific Account - Round Range](#round-range) |
| currency-greater-than | [Search Transactions for a Specific Account - Currency](#currency) |
| currency-less-than | [Search Transactions for a Specific Account - Currency](#currency) |
| tx-type | [Search Transactions for a Specific Account - Transaction Type](#transaction-type) |
| sig- type | [Search Transactions for a Specific Account - Signature Type](#signature-type) |
| tx-id | [Search Transactions for a Specific Account - Transaction ID](#signature-type) |

Similar to asset searches this method supports searching based on address and role as well.

This call also supports [Paginated Results](#paginated-results) and [Note Field Searching](#note-field-searching).

# Complete Code Examples

!!! info
    Full running code examples for each SDK are available within the GitHub repo at [/examples/indexer](https://github.com/algorand/docs/tree/master/examples/indexer) and for [download](https://github.com/algorand/docs/blob/master/examples/indexer/indexer.zip?raw=true) (.zip).
