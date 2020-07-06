    package main

    import (
        "fmt"
        "encoding/base64"
		"crypto/ed25519"

        "github.com/algorand/go-algorand-sdk/client/algod"
		"github.com/algorand/go-algorand-sdk/crypto"
        "github.com/algorand/go-algorand-sdk/mnemonic"
        "github.com/algorand/go-algorand-sdk/transaction"
        "github.com/algorand/go-algorand-sdk/types"
    )

    func main() {

        // const algodToken = "algod-token"<PLACEHOLDER>
        // const algodAddress = "algod-address"<PLACEHOLDER>
		// sandbox
		const algodAddress = "http://localhost:4001"
		const algodToken = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

        // Get private key for sender address
		// PASSPHRASE := "25-word-mnemonic<PLACEHOLDER>"
		PASSPHRASE := "buzz genre work meat fame favorite rookie stay tennis demand panic busy hedgehog snow morning acquire ball grain grape member blur armor foil ability seminar"		
        sk, err := mnemonic.ToPrivateKey(PASSPHRASE)	
        pk := sk.Public()
        var a types.Address
        cpk := pk.(ed25519.PublicKey)
        copy(a[:], cpk[:])
        fmt.Printf("Address: %s\n", a.String())	
        sender := a.String()

        // Create logic signature
        // example base64 encoded program "ASABACI=" int 0
        var ma crypto.MultisigAccount
		// program, err :=  base64.StdEncoding.DecodeString("base64-encoded-program"<PLACEHOLDER>)
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
		const receiver = "QUDVUXBX4Q3Y2H5K2AG3QWEOMY374WO62YNJFFGUTMOJ7FB74CMBKY6LPQ"	
        // const fee = fee<PLACEHOLDER>
		// const amount = amount<PLACEHOLDER>
		const fee = 1000
        const amount = 100000
        var note []byte	
        genID := txParams.GenesisID
        genHash := txParams.GenesisHash
        tx, err := transaction.MakePaymentTxnWithFlatFee(
            sender, receiver, fee, amount, txParams.LastRound, (txParams.LastRound + 1000),
            note, "", genID, genHash )

        txid, stx, err := crypto.SignLogicsigTransaction(lsig, tx)
        if err != nil {
            fmt.Printf("Signing failed with %v", err)
            return
        }
        fmt.Printf("Signed tx: %v\n", txid)

        // Submit the raw transaction as normal - expected to fail with int 0 program, always false
	    fmt.Printf("expected to fail with int 0 program, always false %s\n", err)
		transactionID, err := algodClient.SendRawTransaction(stx)
        if err != nil {
            fmt.Printf("Sending failed with %v\n", err)
        }
        fmt.Printf("Transaction ID: %v\n", transactionID)

    }