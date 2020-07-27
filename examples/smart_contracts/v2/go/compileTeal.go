package main

import (
	//	"bytes"
	"context"
	"io/ioutil"
	"log"

	//	"crypto/ed25519"
	//	"encoding/base64"
	//  "encoding/gob"
	//	"encoding/json"

	// "encoding/binary"
	"fmt"
	"os"

	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	//	"github.com/algorand/go-algorand-sdk/crypto"
	//	"github.com/algorand/go-algorand-sdk/transaction"
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
		fmt.Printf("failed to make client: %s\n", err)
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
		fmt.Printf("failed to read file: %s\n", err)
		return}


	response, err := algodClient.TealCompile(tealFile).Do(context.Background())
	fmt.Printf("Hash = %s\n", response.Hash)
	fmt.Printf("Result = %s\n", response.Result)

}
// results should look similar to
// Hash = KI4DJG2OOFJGUERJGSWCYGFZWDNEU2KWTU56VRJHITP62PLJ5VYMBFDBFE
// Result = ASABACI=