package main

import (
    "context"
    "crypto/ed25519"
    "fmt"

    "github.com/algorand/go-algorand-sdk/client/v2/algod"
    "github.com/algorand/go-algorand-sdk/crypto"
    "github.com/algorand/go-algorand-sdk/mnemonic"
    "github.com/algorand/go-algorand-sdk/transaction"
    "github.com/algorand/go-algorand-sdk/types"
)

const algodAddress = "http://localhost:4001"
const algodToken = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

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
    algodClient, err := algod.MakeClient(algodAddress, algodToken)
    if err != nil {
        fmt.Printf("Issue with creating algod client: %s\n", err)
        return
    }

    passphrase := "Your 25-word mnemonic generated and displayed above"

    privateKey, err := mnemonic.ToPrivateKey(passphrase)
    if err != nil {
        fmt.Printf("Issue with mnemonic conversion: %s\n", err)
        return
    }

    var myAddress types.Address
    publicKey := privateKey.Public()
    cpk := publicKey.(ed25519.PublicKey)
    copy(myAddress[:], cpk[:])
    fmt.Printf("My address: %s\n", myAddress.String())

    // Check account balance
    accountInfo, err := algodClient.AccountInformation(myAddress.String()).Do(context.Background())
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
    txParams.FlatFee = true
    txParams.Fee = 1000

    fromAddr := myAddress.String()
    toAddr := "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A"
    var amount uint64 = 1000000
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
    txID, signedTxn, err := crypto.SignTransaction(privateKey, txn)
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
    waitForConfirmation(txID, algodClient)

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