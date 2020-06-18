package main

import (

    "encoding/json"
    "fmt"

    "golang.org/x/crypto/ed25519"

    "github.com/algorand/go-algorand-sdk/client/algod"
    "github.com/algorand/go-algorand-sdk/crypto"
    "github.com/algorand/go-algorand-sdk/mnemonic"
    "github.com/algorand/go-algorand-sdk/transaction"
    "github.com/algorand/go-algorand-sdk/types"
)

const algodAddress = <algod-address>
const algodToken = <algod-tokenn>
// Function that waits for a given txId to be confirmed by the network
func waitForConfirmation(algodClient algod.Client, txID string) {
    for {
        pt, err := algodClient.PendingTransactionInformation(txID)
        if err != nil {
            fmt.Printf("waiting for confirmation... (pool error, if any): %s\n", err)
            continue
        }
        if pt.ConfirmedRound > 0 {
            fmt.Printf("Transaction "+pt.TxID+" confirmed in round %d\n", pt.ConfirmedRound)
            break
        }
        nodeStatus, err := algodClient.Status()
        if err != nil {
            fmt.Printf("error getting algod status: %s\n", err)
            return
        }
        algodClient.StatusAfterBlock( nodeStatus.LastRound + 1)
    }
}


func main() {

    algodClient, err := algod.MakeClient(algodAddress, algodToken)
    if err != nil {
        return
    }

    passphrase := <25-word-mnemonic>

    privateKey, err := mnemonic.ToPrivateKey(passphrase)
    if err != nil {
        fmt.Printf("Issue with mnemonic conversion: %s\n", err)
    }

    var myAddress types.Address
    publicKey := privateKey.Public()
    cpk := publicKey.(ed25519.PublicKey)
    copy(myAddress[:], cpk[:])
    fmt.Printf("My address: %s\n", myAddress.String())

    // Check account balance
    accountInfo, err := algodClient.AccountInformation(myAddress.String())
    if err != nil {
        fmt.Printf("Error getting account info: %s\n", err)
    }
    fmt.Printf("Account balance: %d microAlgos\n", accountInfo.Amount)

    // Construct the transaction
    txParams, err := algodClient.SuggestedParams()
    if err != nil {
        fmt.Printf("Error getting suggested tx params: %s\n", err)
        return
    }

    fromAddr := myAddress
    toAddr := "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A"
    var amount uint64 = 10000000
    var minFee uint64 = 1000
    note := []byte("Hello World")
    genID := txParams.GenesisID
    genHash := txParams.GenesisHash
    firstValidRound := txParams.LastRound
    lastValidRound := firstValidRound + 1000

    txn, err := transaction.MakePaymentTxnWithFlatFee(fromAddr.String(), toAddr, minFee, amount, firstValidRound, lastValidRound, note, "", genID, genHash)
    if err != nil {
        fmt.Printf("Error creating transaction: %s\n", err)
        return
    }

    // Sign the transaction
    txId, bytes, err := crypto.SignTransaction(privateKey, txn)
    if err != nil {
        fmt.Printf("Failed to sign transaction: %s\n", err)
        return
    }
    fmt.Printf("Signed txid: %s\n", txId)

    // Submit the transaction

    sendResponse, err := algodClient.SendRawTransaction(bytes)
    if err != nil {
        fmt.Printf("failed to send transaction: %s\n", err)
        return
    }
    fmt.Printf("Submitted transaction %s\n", sendResponse.TxID)

    // Wait for confirmation
    waitForConfirmation(algodClient, sendResponse.TxID)

    // Read confirmed transaction from block
    confirmedTxn, err := algodClient.TransactionInformation(myAddress.String(), txId)
    if err != nil {
        fmt.Printf("Error retrieving transaction %s\n", txId)
        return
    }
    txnJSON, err := json.MarshalIndent(confirmedTxn, "", "\t")
    if err != nil {
        fmt.Printf("Can not marshall txn data: %s\n", err)
    }
    fmt.Printf("Transaction information: %s\n", txnJSON)
    fmt.Printf("Decoded note: %s\n", string(confirmedTxn.Note))
}