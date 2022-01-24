package main

import (
	"context"
	json "encoding/json"
	// "errors"
	"fmt"
	// "strings"
	"github.com/algorand/go-algorand-sdk/future"
	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	// "github.com/algorand/go-algorand-sdk/client/v2/common/models"
	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/mnemonic"
	"github.com/algorand/go-algorand-sdk/transaction"
)



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
	var startingAmount uint64 = accountInfo.Amount
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
	// close to dispenser
	closeToAddr := "HZ57J3K46JIJXILONBBZOHX6BKPXEM2VVXNRFSUED6DKFD5ZD24PMJ3MVA"
	var amount uint64 = 1000000
	var minFee uint64 = 1000
	note := []byte("DevPortal - My First Transaction with Go SDK")
	genID := txParams.GenesisID
	genHash := txParams.GenesisHash
	firstValidRound := uint64(txParams.FirstRoundValid)
	lastValidRound := uint64(txParams.LastRoundValid)
	txn, err := transaction.MakePaymentTxnWithFlatFee(fromAddr, toAddr, minFee, amount, firstValidRound, lastValidRound, note, closeToAddr, genID, genHash)
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
	confirmedTxn, err := future.WaitForConfirmation(algodClient, txID, 4)
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
	fmt.Printf("Amount sent: %d microAlgos\n", confirmedTxn.Transaction.Txn.Amount)
	fmt.Printf("Fee: %d microAlgos\n", confirmedTxn.Transaction.Txn.Fee)	
	amountAndFee := uint64 (confirmedTxn.Transaction.Txn.Amount + confirmedTxn.Transaction.Txn.Fee)
	fmt.Printf("Close to Amount: %d microAlgos\n", startingAmount - amountAndFee)		
	fmt.Printf("Decoded note: %s\n", string(confirmedTxn.Transaction.Txn.Note))

	// TODO: insert additional codeblocks here
}
