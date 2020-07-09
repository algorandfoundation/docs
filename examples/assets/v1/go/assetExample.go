package main

import (
	"crypto/ed25519"
	b64 "encoding/base64"
	json "encoding/json"
	"fmt"

	"github.com/algorand/go-algorand-sdk/client/algod"
	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/mnemonic"
	"github.com/algorand/go-algorand-sdk/transaction"
	"github.com/algorand/go-algorand-sdk/types"
)
// 332920
// UPDATE THESE VALUES

// const algodAddress = "Your ADDRESS"
// const algodToken = "Your TOKEN"

// hackathon
// const algodAddress = "http://hackathon.algodev.network:9100"
// const algodToken = "ef920e2e7e002953f4b29a8af720efe8e4ecc75ff102b165e0472834b25832c1"

// sandbox
const algodAddress = "http://localhost:4001"
const algodToken = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

var txHeaders = append([]*algod.Header{}, &algod.Header{"Content-Type", "application/json"})

// Accounts to be used through examples
func loadAccounts() (map[int][]byte, map[int]string) {
	// Shown for demonstration purposes. NEVER reveal secret mnemonics in practice.
	// Change these values to use the accounts created previously.

	// Paste in mnemonic phrases for all three accounts
	// mnemonic1 := "PASTE your phrase for account 1"
	// mnemonic2 := "PASTE your phrase for account 2"
	// mnemonic3 := "PASTE your phrase for account 3"

	mnemonic1 := "portion never forward pill lunch organ biology weird catch curve isolate plug innocent skin grunt bounce clown mercy hole eagle soul chunk type absorb trim"
	mnemonic2 := "place blouse sad pigeon wing warrior wild script problem team blouse camp soldier breeze twist mother vanish public glass code arrow execute convince ability there"
	mnemonic3 := "image travel claw climb bottom spot path roast century also task cherry address curious save item clean theme amateur loyal apart hybrid steak about blanket"

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

// Main function to demonstrate ASA examples
func main() {
	// Get pre-defined set of keys for example
	sks, pks := loadAccounts()

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
	// Print asset info for newly created asset.
	PrettyPrint(txParams)
	PrettyPrint(sks)
	PrettyPrint(pks)
	// note: you would not normally show secret keys for security reasons,
	// they are shown here for tutorial clarity

	// Debug console should look similar to this...

	// Loaded Key 1: THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM
	// Loaded Key 2: AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU
	// Loaded Key 3: 3ZQ3SHCYIKSGK7MTZ7PE7S6EDOFWLKDQ6RYYVMT7OHNQ4UJ774LE52AQCU
	// {
	// 	"fee": 1,
	// 	"genesisID": "testnet-v1.0",
	// 	"genesishashb64": "SGO1GKSzyE7IEPItTxCByw9x8FmnrCDexi9/cOUJOiI=",
	// 	"lastRound": 4268229,
	// 	"consensusVersion": "https://github.com/algorandfoundation/specs/tree/4a9db6a25595c6fd097cf9cc137cc83027787eaa"
	// }
	// {
	// 	"1": "QkWlt0yawnHOIvkgkQ3tbEo6KudsGmDRYtlQ1OeieN2Z4HMPhyEk58kpxgu+MspyO/PcN+Xj7ZHhUKz3JN1VHg==",
	// 	"2": "Lg1Ge0vafd1jv8FbrXcwDEJnbnA9kIpH68XQUoY88SUCWtLBvxyj+BMf96v0jNvrns7vMml+KmvcBZzqGlkLFg==",
	// 	"3": "iuM5VLAiDUsfFLsr0QG8d7KB1/jXdlIBeA9IKAXAoXreYbkcWEKkZX2Tz95Py8Qbi2WocPRxirJ/cdsOUT//Fg=="
	// }
	// {
	// 	"1": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM",
	// 	"2": "AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU",
	// 	"3": "3ZQ3SHCYIKSGK7MTZ7PE7S6EDOFWLKDQ6RYYVMT7OHNQ4UJ774LE52AQCU"
	// }

	// Create an Asset
	// uncomment these imports at the top
	// 	-	b64 "encoding/base64"
	//  - 	"github.com/algorand/go-algorand-sdk/transaction"
	// 	-	"github.com/algorand/go-algorand-sdk/crypto"
	fee := txParams.Fee
	firstRound := txParams.LastRound
	lastRound := txParams.LastRound + 1000
	genHash := b64.StdEncoding.EncodeToString(txParams.GenesisHash)
	genID := txParams.GenesisID

	// Create an asset
	// Set parameters for asset creation transaction
	creator := pks[1]
	assetName := "latinum"
	unitName := "latinum"
	assetURL := "https://path/to/my/asset/details"
	assetMetadataHash := "thisIsSomeLength32HashCommitment"
	defaultFrozen := false
	decimals := uint32(0)
	totalIssuance := uint64(1000)
	manager := pks[2]
	reserve := pks[2]
	freeze := pks[2]
	clawback := pks[2]
	note := []byte(nil)
	txn, err := transaction.MakeAssetCreateTxn(creator, fee, firstRound, lastRound, note,
		genID, genHash, totalIssuance, decimals, defaultFrozen, manager, reserve, freeze, clawback,
		unitName, assetName, assetURL, assetMetadataHash)
	if err != nil {
		fmt.Printf("Failed to make asset: %s\n", err)
		return
	}
	fmt.Printf("Asset created AssetName: %s\n", txn.AssetConfigTxnFields.AssetParams.AssetName)

	txid, stx, err := crypto.SignTransaction(sks[1], txn)
	if err != nil {
		fmt.Printf("Failed to sign transaction: %s\n", err)
		return
	}
	fmt.Printf("Transaction ID: %s\n", txid)
	// Broadcast the transaction to the network
	sendResponse, err := algodClient.SendRawTransaction(stx)
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}

	// Wait for transaction to be confirmed
	waitForConfirmation(algodClient, sendResponse.TxID)

	// Retrieve asset ID by grabbing the max asset ID
	// from the creator account's holdings.
	act, err := algodClient.AccountInformation(pks[1], txHeaders...)
	if err != nil {
		fmt.Printf("failed to get account information: %s\n", err)
		return
	}
	assetID := uint64(0)
	for i, _ := range act.AssetParams {
		if i > assetID {
			assetID = i
		}
	}
	fmt.Printf("Asset ID from AssetParams: %d\n", assetID)

	// Retrieve asset info.
	assetInfo, err := algodClient.AssetInformation(assetID, txHeaders...)

	// Print asset info for newly created asset.
	PrettyPrint(assetInfo)
	// terminal output should look similar to this

	// terminal output should look similar to this
	// Asset created AssetName: latinum
	// Transaction ID: 4P4ACUIZTWYGFPSRZ6BPD4P64XZCYN2ZOHO33N4V7TYE2KWWDA4Q
	// Transaction 4P4ACUIZTWYGFPSRZ6BPD4P64XZCYN2ZOHO33N4V7TYE2KWWDA4Q confirmed in round 4308303
	// Asset ID from AssetParams: 151771
	// {
	// 	"creator": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM",
	// 	"total": 1000,
	// 	"decimals": 0,
	// 	"defaultfrozen": false,
	// 	"unitname": "latinum",
	// 	"assetname": "latinum",
	// 	"url": "https://path/to/my/asset/details",
	// 	"metadatahash": "dGhpc0lzU29tZUxlbmd0aDMySGFzaENvbW1pdG1lbnQ=",
	// 	"managerkey": "AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU",
	// 	"reserveaddr": "AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU",
	// 	"freezeaddr": "AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU",
	// 	"clawbackaddr": "AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU"
	// }

	// Change Asset Manager from Account 2 to Account 1
	manager = pks[1]
	oldmanager := pks[2]
	strictEmptyAddressChecking := true
	txn, err = transaction.MakeAssetConfigTxn(oldmanager, fee, firstRound, lastRound, note, genID, genHash, assetID, manager, reserve, freeze, clawback, strictEmptyAddressChecking)
	if err != nil {
		fmt.Printf("Failed to send txn: %s\n", err)
		return
	}

	txid, stx, err = crypto.SignTransaction(sks[2], txn)
	if err != nil {
		fmt.Printf("Failed to sign transaction: %s\n", err)
		return
	}
	fmt.Printf("Transaction ID: %s\n", txid)
	// Broadcast the transaction to the network
	sendResponse, err = algodClient.SendRawTransaction(stx)
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Transaction ID raw: %s\n", sendResponse.TxID)

	// Wait for transaction to be confirmed
	waitForConfirmation(algodClient, sendResponse.TxID)
	// Retrieve asset info.
	assetInfo, err = algodClient.AssetInformation(assetID, txHeaders...)
	// Print asset info showing updated manager address.
	PrettyPrint(assetInfo)

	// Your terminal output should appear similar to this...

	// 	Transaction HC6XR26GVMUZW26HBFHWU45Z4TI3XK4E7WR6IJ3BABPXVTLXI3UA confirmed in round 4309244

	// {
	// 	"creator": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM",
	// 	"total": 1000,
	// 	"decimals": 0,
	// 	"defaultfrozen": false,
	// 	"unitname": "latinum",
	// 	"assetname": "latinum",
	// 	"url": "https://path/to/my/asset/details",
	// 	"metadatahash": "dGhpc0lzU29tZUxlbmd0aDMySGFzaENvbW1pdG1lbnQ=",
	// 	"managerkey": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM",
	// 	"reserveaddr": "AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU",
	// 	"freezeaddr": "AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU",
	// 	"clawbackaddr": "AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU"
	// }

	// Opt-in
	// Account 3 opts in to receive latinum
	// Use previously set transaction parameters and update sending address to account 3
	txn, err = transaction.MakeAssetAcceptanceTxn(pks[3], fee, firstRound, lastRound, note, genID, genHash, assetID)
	if err != nil {
		fmt.Printf("Failed to send transaction MakeAssetAcceptanceTxn: %s\n", err)
		return
	}
	txid, stx, err = crypto.SignTransaction(sks[3], txn)
	if err != nil {
		fmt.Printf("Failed to sign transaction: %s\n", err)
		return
	}

	fmt.Printf("Transaction ID: %s\n", txid)
	// Broadcast the transaction to the network
	sendResponse, err = algodClient.SendRawTransaction(stx)
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Transaction ID raw: %s\n", sendResponse.TxID)

	// Wait for transaction to be confirmed
	waitForConfirmation(algodClient, sendResponse.TxID)

	act, err = algodClient.AccountInformation(pks[3], txHeaders...)
	if err != nil {
		fmt.Printf("failed to get account information: %s\n", err)
		return
	}
	PrettyPrint(act.Assets[assetID])

	// your terminal output should be similar to this...

	// 	Transaction CYF2K2KLIQB7XE4KZY4ZET3KGRULPNNQWWMMK77WLPAOUN5M5VFA confirmed in round 4309336
	// {
	// 	"creator": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM",
	// 	"amount": 0,
	// 	"frozen": false
	// }

	// Transfer an Asset
	// Send  10 latinum from Account 1 to Account 3

	sender := pks[1]
	recipient := pks[3]
	amount := uint64(10)
	closeRemainderTo := ""
	txn, err = transaction.MakeAssetTransferTxn(sender, recipient, closeRemainderTo, amount, fee, firstRound, lastRound, note,
		genID, genHash, assetID)
	if err != nil {
		fmt.Printf("Failed to send transaction MakeAssetTransfer Txn: %s\n", err)
		return
	}
	txid, stx, err = crypto.SignTransaction(sks[1], txn)
	if err != nil {
		fmt.Printf("Failed to sign transaction: %s\n", err)
		return
	}
	fmt.Printf("Transaction ID: %s\n", txid)
	// Broadcast the transaction to the network
	sendResponse, err = algodClient.SendRawTransaction(stx)
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Transaction ID raw: %s\n", sendResponse.TxID)

	// Wait for transaction to be confirmed
	waitForConfirmation(algodClient, sendResponse.TxID)

	act, err = algodClient.AccountInformation(pks[3], txHeaders...)
	if err != nil {
		fmt.Printf("failed to get account information: %s\n", err)
		return
	}
	PrettyPrint(act.Assets[assetID])

	// Your terminal output should look similar to this
	// Transaction QAHANEYNUGYLKBDRM2BMLG63Y6PET2JNZGLJGH4Y2ROYTLEWZWGQ confirmed in round 4309416
	// {
	// 	"creator": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM",
	// 	"amount": 10,
	// 	"frozen": false
	// }

	// Freeze an Asset
	// The freeze address (Account 2) Freeze's asset for Account 3.
	newFreezeSetting := true
	target := pks[3]
	txn, err = transaction.MakeAssetFreezeTxn(freeze, fee, firstRound, lastRound, note, genID, genHash, assetID, target, newFreezeSetting)
	if err != nil {
		fmt.Printf("Failed to send txn: %s\n", err)
		return
	}
	txid, stx, err = crypto.SignTransaction(sks[2], txn)
	if err != nil {
		fmt.Printf("Failed to sign transaction: %s\n", err)
		return
	}
	fmt.Printf("Transaction ID: %s\n", txid)
	// Broadcast the transaction to the network
	sendResponse, err = algodClient.SendRawTransaction(stx)
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Transaction ID raw: %s\n", sendResponse.TxID)
	// Wait for transaction to be confirmed
	waitForConfirmation(algodClient, sendResponse.TxID)

	act, err = algodClient.AccountInformation(pks[3], txHeaders...)
	if err != nil {
		fmt.Printf("failed to get account information: %s\n", err)
		return
	}
	PrettyPrint(act.Assets[assetID])

	// Your terminal output should look similar to this:

	// Transaction RFFNJOQ3QSVJNTOXW2YYBD2RKJGCWCB6DY2JJZ4RSN5PJBFAKIPA confirmed in round 4309523
	// {
	// 	"creator": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM",
	// 	"amount": 10,
	// 	"frozen": true
	// }

	// Revoke an Asset
	// The clawback address (Account 2) revokes 10 latinum from Account 3 (target)
	// and places it back with Account 1 (creator).

	target = pks[3]
	txn, err = transaction.MakeAssetRevocationTxn(clawback, target, creator, amount, fee, firstRound, lastRound, note,
		genID, genHash, assetID)
	if err != nil {
		fmt.Printf("Failed to send txn: %s\n", err)
		return
	}
	txid, stx, err = crypto.SignTransaction(sks[2], txn)
	if err != nil {
		fmt.Printf("Failed to sign transaction: %s\n", err)
		return
	}
	fmt.Printf("Transaction ID: %s\n", txid)
	// Broadcast the transaction to the network
	sendResponse, err = algodClient.SendRawTransaction(stx)
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Transaction ID raw: %s\n", sendResponse.TxID)
	// Wait for transaction to be confirmed
	waitForConfirmation(algodClient, sendResponse.TxID)

	act, err = algodClient.AccountInformation(pks[3], txHeaders...)
	if err != nil {
		fmt.Printf("failed to get account information: %s\n", err)
		return
	}
	fmt.Printf("target")
	PrettyPrint(act.Assets[assetID])

	act, err = algodClient.AccountInformation(pks[1], txHeaders...)
	if err != nil {
		fmt.Printf("failed to get account information: %s\n", err)
		return
	}
	fmt.Printf("recipient")
	PrettyPrint(act.Assets[assetID])

	// Your terminal output should look similar to this...

	// Transaction PJFMUJPOSPQG6ZAKHZ4RW3WZT7BIU7B2QRR6OZA3TG7OSDK7U5ZA confirmed in round 4309690
	// target{
	// 	"creator": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM",
	// 	"amount": 0,
	// 	"frozen": true
	// }
	// recipient{
	// 	"creator": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM",
	// 	"amount": 1000,
	// 	"frozen": false
	// }

	// Destroy the asset
	// Make sure all funds are back in the creator's account. Then use the
	// Manager account (Account 1) to destroy the asset.
	txn, err = transaction.MakeAssetDestroyTxn(manager, fee, firstRound, lastRound, note, genID, genHash, assetID)
	if err != nil {
		fmt.Printf("Failed to send txn: %s\n", err)
		return
	}
	txid, stx, err = crypto.SignTransaction(sks[1], txn)
	if err != nil {
		fmt.Printf("Failed to sign transaction: %s\n", err)
		return
	}
	fmt.Printf("Transaction ID: %s\n", txid)
	// Broadcast the transaction to the network
	sendResponse, err = algodClient.SendRawTransaction(stx)
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Transaction ID raw: %s\n", sendResponse.TxID)
	// Wait for transaction to be confirmed
	waitForConfirmation(algodClient, sendResponse.TxID)
	// Retrieve asset info. This should now throw an error.
	assetInfo, err = algodClient.AssetInformation(assetID, txHeaders...)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	// Your terminal output should look similar to this...

	// Transaction ID: WBJTR4TWEVVWMNTUHXUY6CHXKRK3NVF76WQ74VF5HT2GFA75OI4A
	// Transaction ID raw: WBJTR4TWEVVWMNTUHXUY6CHXKRK3NVF76WQ74VF5HT2GFA75OI4A
	// Transaction WBJTR4TWEVVWMNTUHXUY6CHXKRK3NVF76WQ74VF5HT2GFA75OI4A confirmed in round 4309776
	// HTTP 404 Not Found: failed to retrieve asset creator from the ledger
}
