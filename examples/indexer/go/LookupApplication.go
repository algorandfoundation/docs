package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/algorand/go-algorand-sdk/client/v2/indexer"
)

const indexerAddress = "http://localhost:59998"
const indexerToken = ""

func main() {

	// Create an indexer client
	indexerClient, err := indexer.MakeClient(indexerAddress, indexerToken)
	if err != nil {
		return
	}
    var applicationID uint64  = 22
	// Lookup application
	result, err := indexerClient.LookupApplicationByID(applicationID).Do(context.Background())

	// Print the results
	JSON, err := json.MarshalIndent(result, "", "\t")
	fmt.Printf(string(JSON) + "\n")

}

// results should look similar to this...
// {
// 	"application": {
// 		"id": 22,
// 		"params": {
// 			"creator": "GHFRLVOMKJNTJ4HY3P74ZR4CNE2PB7CYAUAJ6HVAVVDX7ZKEMLJX6AAF4M",
// 			"global-state-schema": {},
// 			"local-state-schema": {}
// 		}
// 	},
// 	"current-round": 377
// }

