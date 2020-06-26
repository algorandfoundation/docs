package main

import (
	"context"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/transaction"
)

// Function that waits for a given txId to be confirmed by the network
func waitForConfirmation(txID string, client *algod.Client) {
	status, err := client.Status().Do(context.Background())
	if err != nil {
		fmt.Printf("error getting algod status: %s\n", err)
		return
	}
	lastRound := status.LastRound
	for {
		pt, _, err := client.PendingTransactionInformation(txID).Do(context.Background())
		if err != nil {
			fmt.Printf("error getting pending transaction: %s\n", err)
			return
		}
		if pt.ConfirmedRound > 0 {
			fmt.Printf("Transaction "+txID+" confirmed in round %d\n", pt.ConfirmedRound)
			break
		}
		fmt.Printf("waiting for confirmation\n")
		lastRound++
		status, err = client.StatusAfterBlock(lastRound).Do(context.Background())
	}
}
func main() {

	// const algodToken = "algod-token<PLACEHOLDER>"
	// const algodAddress = "algod-address<PLACEHOLDER>"

	// sandbox
	const algodAddress = "http://localhost:4001"
	const algodToken = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	// Create logic signature
	// example base64 encoded program "ASABACI=" int 0
	var sk ed25519.PrivateKey
	var ma crypto.MultisigAccount
	// program, err :=  base64.StdEncoding.DecodeString("base64-encoded-program<PLACEHOLDER>")

	program, err := base64.StdEncoding.DecodeString("ASABACI=")
	var args [][]byte
	lsig, err := crypto.MakeLogicSig(program, args, sk, ma)
	addr := crypto.LogicSigAddress(lsig).String()
	fmt.Printf("Escrow Address: %s\n", addr)

	// Create an algod client
	algodClient, err := algod.MakeClient(algodAddress, algodToken)
	if err != nil {
		fmt.Printf("failed to make algod client: %s\n", err)
		return
	}
	// Get suggested params for the transaction
	txParams, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	txParams.FlatFee = true
	txParams.Fee = 1000

	// Make transaction
	// const receiver = "transaction-receiver"<PLACEHOLDER>
	// const fee = fee<PLACEHOLDER>
	// const amount = amount<PLACEHOLDER>
	const receiver = "QUDVUXBX4Q3Y2H5K2AG3QWEOMY374WO62YNJFFGUTMOJ7FB74CMBKY6LPQ"
	const fee = 1000
	const amount = 100000
	var minFee uint64 = 1000
	note := []byte("Hello World")
	genID := txParams.GenesisID
	genHash := txParams.GenesisHash
	firstValidRound := uint64(txParams.FirstRoundValid)
	lastValidRound := uint64(txParams.LastRoundValid)
	tx, err := transaction.MakePaymentTxnWithFlatFee(
		addr, receiver, minFee, amount, firstValidRound, lastValidRound, note, "", genID, genHash)

	txID, stx, err := crypto.SignLogicsigTransaction(lsig, tx)
	if err != nil {
		fmt.Printf("Signing failed with %v", err)
		return
	}
	fmt.Printf("Signed tx: %v\n", txID)

	// Submit the raw transaction to network
	fmt.Printf("expected to fail with int 0 program, always false %s\n", err)
	transactionID, err := algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		fmt.Printf("Sending failed with %v\n", err)
	}
	fmt.Printf("Transaction ID: %v\n", transactionID)

	// Read transaction
	confirmedTxn, stxn, err := algodClient.PendingTransactionInformation(txID).Do(context.Background())
	if err != nil {
		fmt.Printf("Error retrieving transaction %s\n", txID)
		return
	}
	txnJSON, err := json.MarshalIndent(confirmedTxn.Transaction.Txn, "", "\t")
	if err != nil {
		fmt.Printf("Can not marshall txn data: %s\n", err)
	}
	fmt.Printf("Transaction information: %s\n", txnJSON)
	fmt.Printf("Decoded note: %s\n", string(stxn.Txn.Note))
}
