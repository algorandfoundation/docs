package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/algorand/go-algorand-sdk/client/v2/indexer"
)

// indexer host
const indexerAddress = "http://localhost:8980"
const indexerToken = ""

// query parameters
var minAmount uint64 = 10
var data = "showing prefix"
var encodedNote = base64.StdEncoding.EncodeToString([]byte(data))

func main() {

	// Create an indexer client
	indexerClient, err := indexer.MakeClient(indexerAddress, indexerToken)
	if err != nil {
		return
	}

	// Query
	result, err := indexerClient.SearchForTransactions().NotePrefix([]byte(data)).Do(context.Background())

	// Print the results
	JSON, err := json.MarshalIndent(result, "", "\t")
	fmt.Printf(string(JSON) + "\n")
}
