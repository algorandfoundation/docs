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
// var data = "showing prefix"
// var encodedNote = base64.StdEncoding.EncodeToString([]byte(data))
var notePrefix = "showing prefix"



func main() {
	address, _ := types.DecodeAddress("IAMIRIFW3ERXIMR5LWNYHNK7KRTESUGS4QHOPKF2GL3CLHWWGW32XWB7OI")

	// Create an indexer client
	indexerClient, err := indexer.MakeClient(indexerAddress, indexerToken)
	if err != nil {
		return
	}

	// Query
	result, err := indexerClient.SearchForTransactions().Address(address).NotePrefix([]byte(notePrefix)).Do(context.Background())

	// Print the results
	JSON, err := json.MarshalIndent(result, "", "\t")
	fmt.Printf(string(JSON) + "\n")
}
