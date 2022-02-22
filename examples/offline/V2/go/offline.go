package main

import (
	"context"
	"crypto/ed25519"
	"fmt"
	"io/ioutil"
	"github.com/algorand/go-algorand-sdk/future"
	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/encoding/msgpack"
	"github.com/algorand/go-algorand-sdk/mnemonic"
	"github.com/algorand/go-algorand-sdk/transaction"
	"github.com/algorand/go-algorand-sdk/types"
)

const algodAddress = "http://localhost:4001"
const algodToken = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"



// utility function to recover account and return sk and address
// Shown for demonstration purposes. NEVER reveal secret mnemonics in practice.
func recoverAccount() (string, ed25519.PrivateKey) {
	const passphrase = "PASTE phrase for account"

	sk, err := mnemonic.ToPrivateKey(passphrase)
	if err != nil {
		fmt.Printf("error recovering account: %s\n", err)
		return "", nil
	}
	pk := sk.Public()
	var a types.Address
	cpk := pk.(ed25519.PublicKey)
	copy(a[:], cpk[:])
	fmt.Printf("Address: %s\n", a.String())
	address := a.String()
	return address, sk
}

func saveUnsignedTransaction() {

	algodClient, err := algod.MakeClient(algodAddress, algodToken)
	if err != nil {
		fmt.Printf("Issue with creating algod client: %s\n", err)
		return
	}
	// recover account
	myAddress, _ := recoverAccount()

	// Check account balance
	accountInfo, err := algodClient.AccountInformation(myAddress).Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting account info: %s\n", err)
		return
	}
	fmt.Printf("Account balance: %d microAlgos\n", accountInfo.Amount)

	// Construct the transaction
	txParams, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	// txParams.FlatFee = true
	// txParams.Fee = 1000

	fromAddr := myAddress
	toAddr := "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A"
	var amount uint64 = 100000
	var minFee uint64 = uint64(txParams.MinFee)
	note := []byte("Hello World")
	genID := txParams.GenesisID
	genHash := txParams.GenesisHash
	firstValidRound := uint64(txParams.FirstRoundValid)
	lastValidRound := uint64(txParams.LastRoundValid)

	tx, err := transaction.MakePaymentTxnWithFlatFee(fromAddr, toAddr, minFee, amount, firstValidRound, lastValidRound, note, "", genID, genHash)
	if err != nil {
		fmt.Printf("Error creating transaction: %s\n", err)
		return
	}
	unsignedTx := types.SignedTxn{
		Txn: tx,
	}

	// save unsigned Transaction to file
	err = ioutil.WriteFile("./unsigned.txn", msgpack.Encode(unsignedTx), 0644)
	if err == nil {
		fmt.Printf("Saved unsigned transaction to file\n")
		return
	}
	fmt.Printf("Failed in saving trx to file, error %s\n", err)

}

func readUnsignedTransaction() {
	algodClient, err := algod.MakeClient(algodAddress, algodToken)
	if err != nil {
		fmt.Printf("Issue with creating algod client: %s\n", err)
		return
	}

	// read unsigned transaction from file
	dat, err := ioutil.ReadFile("./unsigned.txn")
	if err != nil {
		fmt.Printf("Error reading transaction from file: %s\n", err)
		return
	}
	var unsignedTxRaw types.SignedTxn
	var unsignedTxn types.Transaction
	msgpack.Decode(dat, &unsignedTxRaw)
	unsignedTxn = unsignedTxRaw.Txn

	// recover account
	myAddress, privateKey := recoverAccount()
	fmt.Printf("Address is: %s\n", myAddress)
	// Check account balance
	accountInfo, err := algodClient.AccountInformation(myAddress).Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting account info: %s\n", err)
		return
	}
	fmt.Printf("Account balance: %d microAlgos\n", accountInfo.Amount)

	// Sign the transaction
	txID, signedTxn, err := crypto.SignTransaction(privateKey, unsignedTxn)
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
	confirmedTxn, err := future.WaitForConfirmation(algodClient,txID,  4, context.Background())
	if err != nil {
		fmt.Printf("Error waiting for confirmation on txID: %s\n", txID)
		return
	}
	fmt.Printf("Confirmed Transaction: %s in Round %d\n", txID ,confirmedTxn.ConfirmedRound)

	fmt.Printf("Decoded note: %s\n", string(confirmedTxn.Transaction.Txn.Note))
}

func saveSignedTransaction() {
	algodClient, err := algod.MakeClient(algodAddress, algodToken)
	if err != nil {
		fmt.Printf("Issue with creating algod client: %s\n", err)
		return
	}
	// recover account
	myAddress, privateKey := recoverAccount()

	// Check account balance
	accountInfo, err := algodClient.AccountInformation(myAddress).Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting account info: %s\n", err)
		return
	}
	fmt.Printf("Account balance: %d microAlgos\n", accountInfo.Amount)

	// Construct the transaction
	txParams, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	// txParams.FlatFee = true
	// txParams.Fee = 1000

	fromAddr := myAddress
	toAddr := "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A"
	var amount uint64 = 100000
	var minFee uint64 = uint64(txParams.MinFee)
	note := []byte("Hello World")
	genID := txParams.GenesisID
	genHash := txParams.GenesisHash
	firstValidRound := uint64(txParams.FirstRoundValid)
	lastValidRound := uint64(txParams.LastRoundValid)

	tx, err := transaction.MakePaymentTxnWithFlatFee(fromAddr, toAddr, minFee, amount, firstValidRound, lastValidRound, note, "", genID, genHash)
	if err != nil {
		fmt.Printf("Error creating transaction: %s\n", err)
		return
	}
	// Sign the transaction
	txID, signedTxn, err := crypto.SignTransaction(privateKey, tx)
	if err != nil {
		fmt.Printf("Failed to sign transaction: %s\n", err)
		return
	}
	fmt.Printf("Signed txid: %s\n", txID)

	//Save the signed transaction to file
	err = ioutil.WriteFile("./signed.stxn", signedTxn, 0644)
	if err == nil {
		fmt.Printf("Saved signed transaction to file\n")
		return
	}
	fmt.Printf("Failed in saving trx to file, error %s\n", err)

}

func readSignedTransaction() {
	algodClient, err := algod.MakeClient(algodAddress, algodToken)
	if err != nil {
		fmt.Printf("Issue with creating algod client: %s\n", err)
		return
	}
	// read signed transaction from file
	dat, err := ioutil.ReadFile("./signed.stxn")
	if err != nil {
		fmt.Printf("Error reading signed transaction from file: %s\n", err)
		return
	}

	// Submit the transaction
	sendResponse, err := algodClient.SendRawTransaction(dat).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Submitted transaction %s\n", sendResponse)


	// Wait for confirmation
	confirmedTxn, err := future.WaitForConfirmation(algodClient,sendResponse,  4, context.Background())
	if err != nil {
		fmt.Printf("Error waiting for confirmation on txID: %s\n", sendResponse)
		return
	}
	fmt.Printf("Confirmed Transaction: %s in Round %d\n", sendResponse ,confirmedTxn.ConfirmedRound)

	fmt.Printf("Decoded note: %s\n", string(confirmedTxn.Transaction.Txn.Note))
}
func main() {
	saveUnsignedTransaction()
	readUnsignedTransaction()

	// saveSignedTransaction()
	// readSignedTransaction()

}
