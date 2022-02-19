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
	"github.com/algorand/go-algorand-sdk/mnemonic"
	"github.com/algorand/go-algorand-sdk/transaction"
	"github.com/algorand/go-algorand-sdk/types"
)



func main() {
    // sandbox
    const algodAddress = "http://localhost:4001"
    const algodToken = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
    // const algodToken = "<algod-token>"
    // const algodAddress = "<algod-address>"
    // Create an algod client
    algodClient, err := algod.MakeClient(algodAddress, algodToken)
    if err != nil {
        fmt.Printf("failed to make algod client: %s\n", err)
        return
    }	
    // Get private key for sender address
    // Do not use mnemonics in production code. For demo purposes only
    PASSPHRASE := "<25-word-mnemonic>"
    sk, err := mnemonic.ToPrivateKey(PASSPHRASE)	
    pk := sk.Public()
    var a types.Address
    cpk := pk.(ed25519.PublicKey)
    copy(a[:], cpk[:])
    fmt.Printf("Address: %s\n", a.String())	
    sender := a.String()

    // Create logic signature
    var ma crypto.MultisigAccount

    // samplearg.teal  ... This code is meant for learning purposes only
    // It should not be used in production
    // arg_0
    // btoi
    // int 123
    // == 
    // file, err := os.Open("<filename>")
    file, err := os.Open("./samplearg.teal")
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
    // lsig, err := crypto.MakeLogicSig(program, args, sk, m a)

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
    
    // Construct the transaction
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
    const receiver = "QUDVUXBX4Q3Y2H5K2AG3QWEOMY374WO62YNJFFGUTMOJ7FB74CMBKY6LPQ"	
    // const fee = <fee>
    // const amount = <amount>
    const fee = 1000
    const amount = 200000
    note := []byte("Hello World")
    genID := txParams.GenesisID
    genHash := txParams.GenesisHash
    firstValidRound := uint64(txParams.FirstRoundValid)
    lastValidRound := uint64(txParams.LastRoundValid)
    tx, err := transaction.MakePaymentTxnWithFlatFee(
        sender, receiver, fee, amount, firstValidRound, lastValidRound,
        note, "", genID, genHash )

    txID, stx, err := crypto.SignLogicsigTransaction(lsig, tx)
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

    // Submit the raw transaction as normal 

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