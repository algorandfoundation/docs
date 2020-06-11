package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/algorand/go-algorand-sdk/client/v2/indexer"
)

// indexer host
const indexerAddress = "http://localhost:8980"
const indexerToken = ""

// query parameters
var nextToken = ""
var numTx uint64 = 1
var responseAll = ""

func main() {

	// Create an indexer client
	indexerClient, err := indexer.MakeClient(indexerAddress, indexerToken)
	if err != nil {
		return
	}

	// Loop until there are no more transactions in the reponse
	// for the limit (max is 1000  per request)
	// "min_amount": 100000000000000
	for numTx > 0 {
		// Lookup asset
		result, err := indexerClient.SearchForTransactions().Limit(numTx).Do(context.Background())
		transactions := result.Transactions
		numTx = len(transactions)
		if numTx > 0 {
			nextToken = transactions["next-token"]
			responseAll = responseAll + transactions
		}

	}

	// Print results
	JSON, err := json.MarshalIndent(result, "", "\t")
	fmt.Printf(string(JSON) + "\n")
}
