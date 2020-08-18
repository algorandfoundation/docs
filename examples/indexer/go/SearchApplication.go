package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/algorand/go-algorand-sdk/client/v2/indexer"
)

// indexer host
const indexerAddress = "http://localhost:59998"
const indexerToken = ""

func main() {

	// Create an indexer client
	indexerClient, err := indexer.MakeClient(indexerAddress, indexerToken)
	if err != nil {
		return
	}

	// Lookup account
	result, err := indexerClient.SearchForApplications().Do(context.Background())

	// Print the results
	JSON, err := json.MarshalIndent(result, "", "\t")
	fmt.Printf(string(JSON) + "\n")

}

// results should look similar to this...
// {
// 	"applications": [
// 		{
// 			"id": 20,
// 			"params": {
// 				"approval-program": "ASABASI=",
// 				"clear-state-program": "ASABASI=",
// 				"creator": "DQ5PMCTEBZLM4UJEDSGZLKAV6ZGXRK2C5WYAFC63RSHI54ASQSJHDMMTUM",
// 				"global-state-schema": {},
// 				"local-state-schema": {}
// 			}
// 		},
// 		{
// 			"id": 22,
// 			"params": {
// 				"creator": "GHFRLVOMKJNTJ4HY3P74ZR4CNE2PB7CYAUAJ6HVAVVDX7ZKEMLJX6AAF4M",
// 				"global-state-schema": {},
// 				"local-state-schema": {}
// 			}
// 		},

