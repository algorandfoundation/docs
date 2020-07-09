package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/algorand/go-algorand-sdk/client/v2/indexer"
	"github.com/algorand/go-algorand-sdk/types"
)

// indexer host
const indexerAddress = "http://localhost:8980"
const indexerToken = ""

// query parameters
var sigType = "msig"

func main() {
	address, _ := types.DecodeAddress("XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4")

	// Create an indexer client
	indexerClient, err := indexer.MakeClient(indexerAddress, indexerToken)
	if err != nil {
		return
	}

	// Query
	result, err := indexerClient.SearchForTransactions().Address(address).SigType(sigType).Do(context.Background())

	// Print the results
	JSON, err := json.MarshalIndent(result, "", "\t")
	fmt.Printf(string(JSON) + "\n")
}
