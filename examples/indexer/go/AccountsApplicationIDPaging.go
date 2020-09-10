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

// query parameters
var applicationID uint64 = 70
var nextToken = ""
var numAccounts = 1
var numPages = 1
var limit uint64 = 2
func main() {
	// Create an indexer client
	indexerClient, err := indexer.MakeClient(indexerAddress, indexerToken)
	if err != nil {
		return
	}
	for numAccounts > 0 {
	// Query
	result, err := indexerClient.SearchAccounts().ApplicationId(applicationID).Limit(limit).NextToken(nextToken).Do(context.Background())
	if err != nil {
		return
	}
	accounts := result.Accounts
	numAccounts = len(accounts)
	nextToken = result.NextToken
	if numAccounts > 0 {
		// Print results
		JSON, err := json.MarshalIndent(accounts, "", "\t")
		if err != nil {
			return
		}
		fmt.Println("Account Info for Application ID : ", string(JSON) )
		fmt.Println("End of page : ", numPages)
		fmt.Println("Accounts printed on this page : ", len(accounts))
		fmt.Println("Next Token : ", nextToken)
		numPages++
	}
	}
}
