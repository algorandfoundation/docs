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
var numTx = 1
var numPages = 1
var minAmount uint64 = 100000000000000
var limit uint64 = 2

func main() {

	// Create an indexer client
	indexerClient, err := indexer.MakeClient(indexerAddress, indexerToken)
	if err != nil {
		return
	}

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
}
