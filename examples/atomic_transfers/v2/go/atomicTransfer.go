package main

import (
	"context"
	"crypto/ed25519"
	json "encoding/json"
	"fmt"

	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/mnemonic"
	"github.com/algorand/go-algorand-sdk/transaction"
	"github.com/algorand/go-algorand-sdk/types"
	"github.com/algorand/go-algorand-sdk/future"
)

// This atomic transfer example code requires three (3) acounts:
//  - account1 requires a user-defined mnemonic and be funded with 101000 microAlgos
//  - account2 requires a user-defined mnemonic and be funded with 201000 microAlgos
//  - account3 auto-generated within the code, 100000 microAlgos will be transfered here
// For account1 and account2, replcace the string "Your 25-word mnemonic goes here" in the code below.
// For account3, ensure you note the mnemonic generated for future.
// Faucents available for funding accounts:
//  - TestNet: https://developer.algorand.org/docs/reference/algorand-networks/testnet/#faucet
//  - BetaNet: https://developer.algorand.org/docs/reference/algorand-networks/betanet/#faucet
// Replace the algodToken and algodAddress parameters below to connect to your API host.

// user declared account mnemonics for account1 and account2
// never us mnemonics in code. This is for demo purposes only.
const mnemonic1 = "Your 25-word mnemonic goes here"
const mnemonic2 = "Your 25-word mnemonic goes here"


// user declared algod connection parameters
const algodAddress = "http://localhost:4001"
const algodToken = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"


// const algodToken = "ef920e2e7e002953f4b29a8af720efe8e4ecc75ff102b165e0472834b25832c1"
// const algodAddress = "http://hackathon.algodev.network"

// prettyPrint prints Go structs
func prettyPrint(data interface{}) {
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}


// utility function to get address string
func getAddress(mn string) string {
	sk, err := mnemonic.ToPrivateKey(mn)
	if err != nil {
		fmt.Printf("error recovering account: %s\n", err)
		return ""
	}
	pk := sk.Public()
	var a types.Address
	cpk := pk.(ed25519.PublicKey)
	copy(a[:], cpk[:])
	fmt.Printf("Address: %s\n", a.String())
	address := a.String()
	return address
}

// utility function to generate new account
func generateNewAccount() string {
	account := crypto.GenerateAccount()
	passphrase, err := mnemonic.FromPrivateKey(account.PrivateKey)
	if err != nil {
		fmt.Printf("Error creating new account: %s\n", err)
	} else {
		fmt.Printf("Created new account: %s\n", account.Address)
		fmt.Printf("Generated mnemonic: \"%s\"\n", passphrase)
	}
	return account.Address.String()
}

// utility function to display account balance
func displayAccountAlgoBalance(account string, client *algod.Client) {
	accountInfo, err := client.AccountInformation(account).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to get account info: %v\n", err)
		return
	}
	fmt.Printf("%s: %v microAlgos\n", accountInfo.Address, accountInfo.Amount)
}

func main() {
	// Initialize an algodClient
	algodClient, err := algod.MakeClient(algodAddress, algodToken)
	if err != nil {
		fmt.Printf("failed to make algod client: %v\n", err)
		return
	}

	// declared account1 and account2 based on user supplied mnemonics
	fmt.Println("Loading two existing accounts...")
	account1 := getAddress(mnemonic1)
	account2 := getAddress(mnemonic2)

	// convert mnemonic1 and mnemonic2 using the mnemonic.ToPrivateKey() helper function
	sk1, err := mnemonic.ToPrivateKey(mnemonic1)
	sk2, err := mnemonic.ToPrivateKey(mnemonic2)

	// generate account3, display mnemonic, wait
	fmt.Println("Generating new account...")
	account3 := generateNewAccount()
	fmt.Println("!! NOTICE !! Please retain the above generated \"25-word mnemonic passphrase\" for future use.")

	// display account balances
	fmt.Println("Initial balances:")
	displayAccountAlgoBalance(account1, algodClient)
	displayAccountAlgoBalance(account2, algodClient)
	displayAccountAlgoBalance(account3, algodClient)

	// get node suggested parameters
	txParams, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("error getting suggested tx params: %s\n", err)
		return
	}
	txParams.FlatFee = true
	minFee := uint64(txParams.MinFee)
	genID := txParams.GenesisID
	genHash := txParams.GenesisHash
	firstValidRound := uint64(txParams.FirstRoundValid)
	lastValidRound := uint64(txParams.LastRoundValid)

	// make transactions
	fmt.Println("Creating transactions...")
	// from account 1 to account 3
	fromAddr := account1
	toAddr := account3
	var amount uint64 = 100000

	tx1, err := transaction.MakePaymentTxnWithFlatFee(fromAddr, toAddr, minFee, amount, firstValidRound, lastValidRound, nil, "", genID, genHash)
	if err != nil {
		fmt.Printf("Error creating transaction: %s\n", err)
		return
	}
	fmt.Printf("...tx1: from %s to %s for %v microAlgos\n", fromAddr, toAddr, amount)

	// from account 2 to account 1
	fromAddr = account2
	toAddr = account1
	amount = 200000
	tx2, err := transaction.MakePaymentTxnWithFlatFee(fromAddr, toAddr, minFee, amount, firstValidRound, lastValidRound, nil, "", genID, genHash)
	if err != nil {
		fmt.Printf("Error creating transaction: %s\n", err)
		return
	}
	fmt.Printf("...tx2: from %s to %s for %v microAlgos\n", fromAddr, toAddr, amount)

	// combine transations
	fmt.Println("Combining transactions...")
	// the SDK does this implicitly within grouping below

	fmt.Println("Grouping transactions...")
	// compute group id and put it into each transaction
	gid, err := crypto.ComputeGroupID([]types.Transaction{tx1, tx2})
	tx1.Group = gid
	tx2.Group = gid
	fmt.Println("...computed groupId: ", gid)

	// split transaction group
	fmt.Println("...splitting unsigned transaction group")

	// sign transactions
	fmt.Println("Signing transactions...")
	sTxID1, stx1, err := crypto.SignTransaction(sk1, tx1)
	if err != nil {
		fmt.Printf("Failed to sign transaction: %s\n", err)
		return
	}
	fmt.Println("...account1 signed tx1: ", sTxID1)
	sTxID2, stx2, err := crypto.SignTransaction(sk2, tx2)
	if err != nil {
		fmt.Printf("Failed to sign transaction: %s\n", err)
		return
	}
	fmt.Println("...account2 signed tx2: ", sTxID2)

	// assemble transaction group
	fmt.Println("Assembling transaction group...")
	var signedGroup []byte
	signedGroup = append(signedGroup, stx1...)
	signedGroup = append(signedGroup, stx2...)

	// send transactions
	fmt.Println("Sending transaction group...")
	pendingTxID, err := algodClient.SendRawTransaction(signedGroup).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
    fmt.Printf("Submitted transaction %s\n", pendingTxID)

    confirmedTxn, err := future.WaitForConfirmation(algodClient, pendingTxID, 4, context.Background())
    if err != nil {
		fmt.Printf("Error waiting for confirmation on txID: %s\n", pendingTxID)
        return
    }
	fmt.Printf("Confirmed Transaction: %s in Round %d\n", pendingTxID ,confirmedTxn.ConfirmedRound)


	txnJSON, err := json.MarshalIndent(confirmedTxn.Transaction.Txn, "", "\t")
	if err != nil {
		fmt.Printf("Can not marshall txn data: %s\n", err)
	}
	fmt.Printf("Transaction information: %s\n", txnJSON)

	// display account balances
	fmt.Println("Final balances:")
	displayAccountAlgoBalance(account1, algodClient)
	displayAccountAlgoBalance(account2, algodClient)
	displayAccountAlgoBalance(account3, algodClient)

	//display confirmed transaction group
	// tx1
	confirmedTx, _, err := algodClient.PendingTransactionInformation(sTxID1).Do(context.Background())
	if err != nil {
		fmt.Printf("error getting pending transaction: %s\n", err)
		return
	}
	fmt.Printf("Transaction information tx1: \n")
    prettyPrint(confirmedTx.Transaction.Txn)

	// tx2
	confirmedTx, _, err = algodClient.PendingTransactionInformation(sTxID2).Do(context.Background())
	if err != nil {
		fmt.Printf("error getting pending transaction: %s\n", err)
		return
	}
	fmt.Printf("Transaction information tx2: \n")
    prettyPrint(confirmedTx.Transaction.Txn)

}