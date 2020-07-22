package main

import (

	"context"
	"io/ioutil"
	"log"
	"fmt"
	"os"
	"github.com/algorand/go-algorand-sdk/client/v2/algod"
)

func main() {

	// const algodToken = "algod-token<PLACEHOLDER>"
	// const algodAddress = "algod-address<PLACEHOLDER>"
	// sandbox
	const algodAddress = "http://localhost:4001"
	const algodToken = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	// Create an algod client
	algodClient, err := algod.MakeClient(algodAddress, algodToken)
	if err != nil {
		fmt.Printf("failed to read teal file: %s\n", err)
		return
	}
	// int 0 in sample.teal
	file, err := os.Open("./sample.teal")
    if err != nil {
        log.Fatal(err)
    }    
	defer file.Close()
	tealFile, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("failed to teal file: %s\n", err)
		return}
    // compile teal program
	response, err := algodClient.TealCompile(tealFile).Do(context.Background())
    // print response	
	fmt.Printf("Hash = %s\n", response.Hash)
	fmt.Printf("Result = %s\n", response.Result)
}
// results should look similar to
// Hash = KI4DJG2OOFJGUERJGSWCYGFZWDNEU2KWTU56VRJHITP62PLJ5VYMBFDBFE
// Result = ASABACI=