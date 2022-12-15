package main

import (
	"context"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/binary"

	"io/ioutil"
	"log"
	"os"
	"fmt"
	"github.com/algorand/go-algorand-sdk/future"

	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/transaction"
)


func main() {

	// const algodToken = "<algod-token>"
	// const algodAddress = "<algod-address>"

	// sandbox
	const algodAddress = "http://localhost:4001"
	const algodToken = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	// Create an algod client
	algodClient, err := algod.MakeClient(algodAddress, algodToken)
	if err != nil {
		fmt.Printf("failed to make algod client: %s\n", err)
		return
	}
	// Create logic signature
	var sk ed25519.PrivateKey
    var ma crypto.MultisigAccount

    // samplearg.teal  ... This code is meant for learning purposes only
    // It should not be used in production
    // arg_0
    // btoi
    // int 123
    // == 
	file, err := os.Open("./samplearg.teal")
	// file, err := os.Open("<filename>")
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
    
    program, err :=  base64.StdEncoding.DecodeString(response.Result)	
    // if no args use these two lines
    // var args [][]byte
    // lsig, err := crypto.MakeLogicSig(program, args, sk, ma)

    // string parameter
    // args := make([][]byte, 1)
    // args[0] = []byte("<my string>")
    // lsig, err := crypto.MakeLogicSig(program, args, sk, ma)
    
    // integer args parameter
    args := make([][]byte, 1)
    var buf [8]byte
    binary.BigEndian.PutUint64(buf[:], 123)
    args[0] = buf[:]
    lsig, err := crypto.MakeLogicSig(program, args, sk, ma)

    addr := crypto.LogicSigAddress(lsig).String()


	// Get suggested params for the transaction
	txParams, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	// txParams.FlatFee = true
	// txParams.Fee = 1000

	// Make transaction
	// const receiver = "<receiver-address>"
	// const fee = <fee>
	// const amount = <amount>
	const receiver = "QUDVUXBX4Q3Y2H5K2AG3QWEOMY374WO62YNJFFGUTMOJ7FB74CMBKY6LPQ"
	const fee = 1000
	const amount = 100000
	var minFee uint64 = uint64(txParams.MinFee)
	note := []byte("Hello World")
	genID := txParams.GenesisID
	genHash := txParams.GenesisHash
	firstValidRound := uint64(txParams.FirstRoundValid)
	lastValidRound := uint64(txParams.LastRoundValid)
	tx, err := transaction.MakePaymentTxnWithFlatFee(
		addr, receiver, minFee, amount, firstValidRound, lastValidRound, note, "", genID, genHash)

	txID, stx, err := crypto.SignLogicSigTransaction(lsig, tx)
	if err != nil {
		fmt.Printf("Signing failed with %v", err)
		return
	}
	fmt.Printf("Signed tx: %v\n", txID)
	// logic signature transaction can be written to a file
	// f, err := os.Create("simple.stxn")

	// defer f.Close()
	// if _, err := f.Write(stx); err != nil {
	//     // handle
	// }
	// if err := f.Sync(); err != nil {
	//     // handle
	// }
		
	// Submit the raw transaction to network

	transactionID, err := algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		fmt.Printf("Sending failed with %v\n", err)
	}


	// Wait for confirmation
	confirmedTxn, err := future.WaitForConfirmation(algodClient,transactionID,  4, context.Background())
	if err != nil {
		fmt.Printf("Error waiting for confirmation on txID: %s\n", transactionID)
		return
	}
	fmt.Printf("Confirmed Transaction: %s in Round %d\n", transactionID ,confirmedTxn.ConfirmedRound)


}
