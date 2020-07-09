package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/algorand/go-algorand-sdk/client/v2/indexer"
)

// indexer host
const indexerAddress = "http://localhost:8980"

var indexerToken = ""

// query parameters
var assetId uint64 = 2044572
var minBalance uint64 = 200

func main() {

	// Create an indexer client
	indexerClient, err := indexer.MakeClient(indexerAddress, indexerToken)
	if err != nil {
		return
	}

	// Query
	result, err := indexerClient.LookupAssetBalances(assetId).CurrencyGreaterThan(minBalance).Do(context.Background())

	// Print the results
	JSON, err := json.MarshalIndent(result, "", "\t")
	fmt.Printf(string(JSON) + "\n")
}
