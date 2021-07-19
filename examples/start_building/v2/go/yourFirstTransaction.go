package main

import (
	"context"
	json "encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/client/v2/common/models"
	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/mnemonic"
	"github.com/algorand/go-algorand-sdk/transaction"
)

// Utility function that waits for a given txId to be confirmed by the network
func waitForConfirmation(txID string, client *algod.Client, timeout uint64) (models.PendingTransactionInfoResponse, error) {
	pt := new(models.PendingTransactionInfoResponse)
	if client == nil || txID == "" || timeout < 0 {
		fmt.Printf("Bad arguments for waitForConfirmation")
		var msg = errors.New("Bad arguments for waitForConfirmation")
		return *pt, msg

	}

	status, err := client.Status().Do(context.Background())
	if err != nil {
		fmt.Printf("error getting algod status: %s\n", err)
		var msg = errors.New(strings.Join([]string{"error getting algod status: "}, err.Error()))
		return *pt, msg
	}
	startRound := status.LastRound + 1
	currentRound := startRound

	for currentRound < (startRound + timeout) {

		*pt, _, err = client.PendingTransactionInformation(txID).Do(context.Background())
		if err != nil {
			fmt.Printf("error getting pending transaction: %s\n", err)
			var msg = errors.New(strings.Join([]string{"error getting pending transaction: "}, err.Error()))
			return *pt, msg
		}
		if pt.ConfirmedRound > 0 {
			fmt.Printf("Transaction "+txID+" confirmed in round %d\n", pt.ConfirmedRound)
			return *pt, nil
		}
		if pt.PoolError != "" {
			fmt.Printf("There was a pool error, then the transaction has been rejected!")
			var msg = errors.New("There was a pool error, then the transaction has been rejected")
			return *pt, msg
		}
		fmt.Printf("waiting for confirmation\n")
		status, err = client.StatusAfterBlock(currentRound).Do(context.Background())
		currentRound++
	}
	msg := errors.New("Tx not found in round range")
	return *pt, msg
}

// TODO: insert aditional utility functions here

func main() {
	// Create account
	account := crypto.GenerateAccount()
	passphrase, err := mnemonic.FromPrivateKey(account.PrivateKey)
	myAddress := account.Address.String()

	if err != nil {
		fmt.Printf("Error creating transaction: %s\n", err)
	} else {
		fmt.Printf("My address: %s\n", myAddress)
		fmt.Printf("My passphrase: %s\n", passphrase)
		fmt.Println("--> Copy down your address and passpharse for future use.")
		fmt.Println("--> Once secured, press ENTER key to continue...")
		fmt.Scanln()
	}

	// Fund account
	fmt.Println("Fund the created account using testnet faucet:\n--> https://dispenser.testnet.aws.algodev.network?account=" + myAddress)
	fmt.Println("--> Once funded, press ENTER key to continue...")
	fmt.Scanln()

	// instantiate algod client to Algorand Sandbox
	const algodAddress = "http://localhost:4001"
	const algodToken = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	algodClient, err := algod.MakeClient(algodAddress, algodToken)
	if err != nil {
		fmt.Printf("Issue with creating algod client: %s\n", err)
		return
	}

	//Check account balance
	fmt.Printf("My address: %s\n", myAddress)

	accountInfo, err := algodClient.AccountInformation(myAddress).Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting account info: %s\n", err)
		return
	}
	fmt.Printf("Account balance: %d microAlgos\n", accountInfo.Amount)
	fmt.Println("--> Ensure balance greater than 0, press ENTER key to continue...")
	fmt.Scanln()

	// Construct the transaction
	txParams, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	fromAddr := myAddress
	toAddr := "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A"
	var amount uint64 = 9000000
	var minFee uint64 = 1000
	note := []byte("Hello World")
	genID := txParams.GenesisID
	genHash := txParams.GenesisHash
	firstValidRound := uint64(txParams.FirstRoundValid)
	lastValidRound := uint64(txParams.LastRoundValid)
	txn, err := transaction.MakePaymentTxnWithFlatFee(fromAddr, toAddr, minFee, amount, firstValidRound, lastValidRound, note, "", genID, genHash)
	if err != nil {
		fmt.Printf("Error creating transaction: %s\n", err)
		return
	}

	// Sign the transaction
	txID, signedTxn, err := crypto.SignTransaction(account.PrivateKey, txn)
	if err != nil {
		fmt.Printf("Failed to sign transaction: %s\n", err)
		return
	}
	fmt.Printf("Signed txid: %s\n", txID)

	// Submit the transaction
	sendResponse, err := algodClient.SendRawTransaction(signedTxn).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Submitted transaction %s\n", sendResponse)

	// Wait for confirmation
	confirmedTxn, err := waitForConfirmation(txID, algodClient, 4)
	if err != nil {
		fmt.Printf("Error wating for confirmation on txID: %s\n", txID)
		return
	}

	// Get the completed Transaction
	txnJSON, err := json.MarshalIndent(confirmedTxn.Transaction.Txn, "", "\t")
	if err != nil {
		fmt.Printf("Can not marshall txn data: %s\n", err)
	}
	fmt.Printf("Transaction information: %s\n", txnJSON)
	fmt.Printf("Decoded note: %s\n", string(confirmedTxn.Transaction.Txn.Note))

	// TODO: insert additional codeblocks here
}
