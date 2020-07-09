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
var minAmount uint64 = 10

func main() {

	// Create an indexer client
	indexerClient, err := indexer.MakeClient(indexerAddress, indexerToken)
	if err != nil {
		return
	}

	// Lookup asset
	result, err := indexerClient.SearchForTransactions().CurrencyGreaterThan(minAmount).Do(context.Background())

	// Query
	JSON, err := json.MarshalIndent(result, "", "\t")
	fmt.Printf(string(JSON) + "\n")
}
