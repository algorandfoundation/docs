package main

import (
    "fmt"

    "golang.org/x/crypto/ed25519"

    "github.com/algorand/go-algorand-sdk/client/algod"
    "github.com/algorand/go-algorand-sdk/crypto"
    "github.com/algorand/go-algorand-sdk/mnemonic"
    "github.com/algorand/go-algorand-sdk/transaction"
    "github.com/algorand/go-algorand-sdk/types"
)

const algodToken = <algod-token>
const algodAddress = <algod-address>

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
// utility funcitn to get address string
func getAddress(mn string )(string) {
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
func main() {
    // Initialize an algodClient
    algodClient, err := algod.MakeClient(algodAddress, algodToken)
    if err != nil {
        fmt.Printf("failed to make algod client: %v\n", err)
        return
    }

    // get node suggested parameters
    txParams, err := algodClient.SuggestedParams()
    if err != nil {
        fmt.Printf("error getting suggested tx params: %s\n", err)
        return
    }

    // declare account mnemonics for later consumption for private key conversion
    const mnemonic1 = <25-word-passphrase>;
    const mnemonic2 = <25-word-passphrase>;
    const mnemonic3 = <25-word-passphrase>;

    // convert mnemonic1 and mnemonic2 using the mnemonic.ToPrivateKey() helper function
    sk1, err := mnemonic.ToPrivateKey(mnemonic1)
    sk2, err := mnemonic.ToPrivateKey(mnemonic2)
    // declare accounts
    account1 := getAddress(mnemonic1)
    account2 := getAddress(mnemonic2)
    account3 := getAddress(mnemonic3)

    // make transactions
    tx1, err := transaction.MakePaymentTxn(account1, account3, 1, 100000,
        txParams.LastRound, txParams.LastRound+100, nil, "", 
        txParams.GenesisID, txParams.GenesisHash)
    if err != nil {
        fmt.Printf("Error creating transaction: %s\n", err)
        return
    }

    tx2, err := transaction.MakePaymentTxn(account2, account1, 1, 100000,
        txParams.LastRound, txParams.LastRound+100, nil, "", 
        txParams.GenesisID, txParams.GenesisHash)
    if err != nil {
        fmt.Printf("Error creating transaction: %s\n", err)
        return
    }

    // compute group id and put it into each transaction
    gid, err := crypto.ComputeGroupID([]types.Transaction{tx1, tx2})
    tx1.Group = gid
    tx2.Group = gid

    // sign transactions
    _, stx1, err := crypto.SignTransaction(sk1, tx1)
    if err != nil {
        fmt.Printf("Failed to sign transaction: %s\n", err)
        return
    }
    _, stx2, err := crypto.SignTransaction(sk2, tx2)
    if err != nil {
        fmt.Printf("Failed to sign transaction: %s\n", err)
        return
    }

    // send transactions
    var signedGroup []byte
    signedGroup = append(signedGroup, stx1...)
    signedGroup = append(signedGroup, stx2...)
    signed, err := algodClient.SendRawTransaction(signedGroup)
    if err != nil {
        fmt.Printf("Failed to create payment transaction: %v\n", err)
        return
    }
    fmt.Printf("Transaction ID: %s\n", signed.TxID)
    waitForConfirmation(algodClient, signed.TxID)
}