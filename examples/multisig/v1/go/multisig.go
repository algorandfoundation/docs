package main

import (
	"crypto/ed25519"
	"fmt"
	json "encoding/json"
	// "io/ioutil"
	"github.com/algorand/go-algorand-sdk/client/algod"
	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/mnemonic"
	"github.com/algorand/go-algorand-sdk/transaction"
	"github.com/algorand/go-algorand-sdk/types"
)

// sandbox
const algodAddress = "http://localhost:4001"
const algodToken = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

// Accounts to be used through examples
func loadAccounts() (map[int][]byte, map[int]string) {
	// Shown for demonstration purposes. NEVER reveal secret mnemonics in practice.
	// Change these values to use the accounts created previously.

	// Paste in mnemonic phrases for all three accounts
	// mnemonic1 := "PASTE your phrase for account 1"
	// mnemonic2 := "PASTE your phrase for account 2"
	// mnemonic3 := "PASTE your phrase for account 3"

	mnemonic1 := "predict mandate aware dizzy limit match hazard fantasy victory auto fortune hello public dragon ostrich happy blue spray parrot island odor actress only ability hurry"
	mnemonic2 := "moon grid random garlic effort faculty fence gym write skin they joke govern home huge there claw skin way bid fit bean damp able only"
	mnemonic3 := "mirror zone together remind rural impose balcony position minimum quick manage climb quit draft lion device pluck rug siege robust spirit fine luggage ability actual"

	mnemonics := []string{mnemonic1, mnemonic2, mnemonic3}
	pks := map[int]string{1: "", 2: "", 3: ""}
	var sks = make(map[int][]byte)

	for i, m := range mnemonics {
		var err error
		sk, err := mnemonic.ToPrivateKey(m)
		sks[i+1] = sk
		if err != nil {
			fmt.Printf("Issue with account %d private key conversion.", i+1)
		}
		// derive public address from Secret Key.
		pk := sk.Public()
		var a types.Address
		cpk := pk.(ed25519.PublicKey)
		copy(a[:], cpk[:])
		pks[i+1] = a.String()
		fmt.Printf("Loaded Key %d: %s\n", i+1, pks[i+1])
	}
	return sks, pks
}

// Function that waits for a given txId to be confirmed by the network
func waitForConfirmation(algodClient algod.Client, txID string) {
	nodeStatus, err := algodClient.Status()
	if err != nil {
		fmt.Printf("error getting algod status: %s\n", err)
		return
	}
	lastRound := nodeStatus.LastRound
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
		lastRound++
		algodClient.StatusAfterBlock(lastRound)
	}
}

// PrettyPrint prints Go structs
func PrettyPrint(data interface{}) {
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}

func main() {

    // Initialize an algodClient
	algodClient, err := algod.MakeClient(algodAddress, algodToken)
	if err != nil {
		return
	}
	// Get network-related transaction parameters and assign
	txParams, err := algodClient.SuggestedParams()
	if err != nil {
		fmt.Printf("error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	txParams.MinTxnFee = 1000
	txParams.Fee = 1000
	// Get pre-defined set of keys for example
	sks, pks := loadAccounts()


	addr1, _ := types.DecodeAddress(pks[1])
	addr2, _ := types.DecodeAddress(pks[2])
	addr3, _ := types.DecodeAddress(pks[3])		
	
	ma, err := crypto.MultisigAccountWithParams(1, 2, []types.Address{
		addr1,
		addr2,
		addr3,
	})

	if err != nil {
		panic("invalid multisig parameters")
	}

	
	fromAddr, _ := ma.Address()
	// Print multisig account
    fmt.Printf("Here is your multisig address : %s \n", fromAddr.String())
    fmt.Println("Please go to: https://bank.testnet.algorand.network/ to fund your multisig account.")
//	fmt.Scanln() // wait for Enter Key

    toAddr := addr3.String()
    var amount uint64 = 10000

    note := []byte("Hello World")
    genID := txParams.GenesisID
    genHash := txParams.GenesisHash
    firstValidRound := txParams.LastRound
	lastValidRound := firstValidRound + 1000
	
	txn, err := transaction.MakePaymentTxn(
		fromAddr.String(),
		toAddr,
		txParams.MinTxnFee,     // fee per byte
		amount,  // amount
		firstValidRound, // first valid round
		lastValidRound, // last valid round
		note,    // note
		"",     // closeRemainderTo
		genID,     // genesisHash
		genHash,     // genesisHash
	)

	txid, txBytes, err := crypto.SignMultisigTransaction(sks[1], ma, txn)
	if err != nil {
		println(err.Error)
		panic("could not sign multisig transaction")
	}
	fmt.Printf("Made partially-signed multisig transaction with TxID %s: %x\n", txid, txBytes)
	// ioutil.WriteFile("./arbitrary_file.tx", txBytes, 0644)
	// readTxBytes, _ := ioutil.ReadFile("./arbitrary_file.tx")
	// txid, twoOfThreeTxBytes, err := crypto.AppendMultisigTransaction(sks[2], ma, readTxBytes)
	txid, twoOfThreeTxBytes, err := crypto.AppendMultisigTransaction(sks[2], ma, txBytes)

	if err != nil {
		panic("could not append signature to multisig transaction")
	}
	fmt.Printf("Appended bytes %x\n", twoOfThreeTxBytes)

    fmt.Printf("Made 2-out-of-3 multisig transaction with TxID %s: %x\n", txid, twoOfThreeTxBytes)


	// We can also merge raw, partially-signed multisig transactions:

	// otherTxBytes := ... // generate another raw multisig transaction somehow
	// txid, mergedTxBytes, err := crypto.MergeMultisigTransactions(twoOfThreeTxBytes, otherTxBytes)

	// Broadcast the transaction to the network
	sendResponse, err := algodClient.SendRawTransaction(twoOfThreeTxBytes)
	waitForConfirmation(algodClient, sendResponse.TxID)
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Transaction ID: %s\n", sendResponse.TxID)
	

}

