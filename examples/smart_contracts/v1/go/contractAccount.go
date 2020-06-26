    package main

    import (
        "fmt"
        "encoding/base64"
		"crypto/ed25519"

        "github.com/algorand/go-algorand-sdk/client/algod"
        "github.com/algorand/go-algorand-sdk/crypto"
        "github.com/algorand/go-algorand-sdk/transaction"
    )

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
		

        program, err :=  base64.StdEncoding.DecodeString("ASABACI=")
        var args [][]byte
        lsig, err := crypto.MakeLogicSig(program, args, sk, ma)
        addr := crypto.LogicSigAddress(lsig).String()
        fmt.Printf("Escrow Address: %s\n" , addr )
        

        // Create an algod client
        algodClient, err := algod.MakeClient(algodAddress, algodToken)
        if err != nil {
            fmt.Printf("failed to make algod client: %s\n", err)
            return
        }	
        // Get suggested params for the transaction
        txParams, err := algodClient.SuggestedParams()
        if err != nil {
                fmt.Printf("error getting suggested tx params: %s\n", err)
                return
        }

        // Make transaction
        // const receiver = "transaction-receiver"<PLACEHOLDER>
        // const fee = fee<PLACEHOLDER>
		// const amount = amount<PLACEHOLDER>
        const receiver = "QUDVUXBX4Q3Y2H5K2AG3QWEOMY374WO62YNJFFGUTMOJ7FB74CMBKY6LPQ"
        const fee = 1000
        const amount = 100000		

        var note []byte	
        genID := txParams.GenesisID
        genHash := txParams.GenesisHash	
        tx, err := transaction.MakePaymentTxnWithFlatFee(
            addr, receiver, fee, amount, txParams.LastRound, (txParams.LastRound + 1000),
            note, "", genID, genHash )

        txid, stx, err := crypto.SignLogicsigTransaction(lsig, tx)
        if err != nil {
            fmt.Printf("Signing failed with %v", err)
            return
        }
        fmt.Printf("Signed tx: %v\n", txid)

		// Submit the raw transaction to network
		fmt.Printf("expected to fail with int 0 program, always false %s\n", err)
        transactionID, err := algodClient.SendRawTransaction(stx)
        if err != nil {
            fmt.Printf("Sending failed with %v\n", err)
        }
        fmt.Printf("Transaction ID: %v\n", transactionID)

    }