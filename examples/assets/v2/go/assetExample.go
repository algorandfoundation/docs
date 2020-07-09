package main

import (
	"context"
	"crypto/ed25519"
	json "encoding/json"
	"fmt"

	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/mnemonic"
	"github.com/algorand/go-algorand-sdk/types"
)
import transaction "github.com/algorand/go-algorand-sdk/future"


// UPDATE THESE VALUES
// const algodAddress = "Your ADDRESS"
// const algodToken = "Your TOKEN"

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

// printAssetHolding utility to print asset holding for account
func printAssetHolding(assetID uint64, account string, client *algod.Client) {

	act, err := client.AccountInformation(account).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to get account information: %s\n", err)
		return
	}
	for _, assetholding := range act.Assets {
		if assetID == assetholding.AssetId {
			prettyPrint(assetholding)
			break
		}
	}
}

// printCreatedAsset utility to print created assert for account
func printCreatedAsset(assetID uint64, account string, client *algod.Client) {

	act, err := client.AccountInformation(account).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to get account information: %s\n", err)
		return
	}
	for _, asset := range act.CreatedAssets {
		if assetID == asset.Index {
			prettyPrint(asset)
			break
		}
	}
}

// Main function to demonstrate ASA examples
func main() {

	// Initialize an algodClient
	algodClient, err := algod.MakeClient(algodAddress, algodToken)
	if err != nil {
		return
	}

	// Get network-related transaction parameters and assign
	txParams, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	txParams.FlatFee = true
	txParams.Fee = 1000

	// Get pre-defined set of keys for example
	sks, pks := loadAccounts()

	// Print asset info for newly created asset.
	prettyPrint(txParams)
	prettyPrint(sks)
	prettyPrint(pks)
	// note: you would not normally show secret keys for security reasons,
	// they are shown here for tutorial clarity

	// Debug console should look similar to this...

	// Loaded Key 1: THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM
	// Loaded Key 2: AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU
	// Loaded Key 3: 3ZQ3SHCYIKSGK7MTZ7PE7S6EDOFWLKDQ6RYYVMT7OHNQ4UJ774LE52AQCU
	// {
	// 	"Fee": 1000,
	// 	"GenesisID": "betanet-v1.0",
	// 	"GenesisHash": "mFgazF+2uRS1tMiL9dsj01hJGySEmPN28B/TjjvpVW0=",
	// 	"FirstRoundValid": 4072061,
	// 	"LastRoundValid": 4073061,
	// 	"ConsensusVersion": "https://github.com/algorandfoundation/specs/tree/e5f565421d720c6f75cdd186f7098495caf9101f",
	// 	"FlatFee": true
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


	// CREATE ASSET

	// Construct the transaction
	// Set parameters for asset creation 
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
	txn, err := transaction.MakeAssetCreateTxn(creator,
		note,
		txParams, totalIssuance, decimals,
		defaultFrozen, manager, reserve, freeze, clawback,
		unitName, assetName, assetURL, assetMetadataHash)

	if err != nil {
		fmt.Printf("Failed to make asset: %s\n", err)
		return
	}
	fmt.Printf("Asset created AssetName: %s\n", txn.AssetConfigTxnFields.AssetParams.AssetName)
	// sign the transaction
	txid, stx, err := crypto.SignTransaction(sks[1], txn)
	if err != nil {
		fmt.Printf("Failed to sign transaction: %s\n", err)
		return
	}
	fmt.Printf("Transaction ID: %s\n", txid)
	// Broadcast the transaction to the network
	sendResponse, err := algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Submitted transaction %s\n", sendResponse)
	// Wait for transaction to be confirmed
	waitForConfirmation(txid, algodClient)
	//    response := algodClient.PendingTransactionInformation(txid)
	//    prettyPrint(response)
	// Retrieve asset ID by grabbing the max asset ID
	// from the creator account's holdings.
	act, err := algodClient.AccountInformation(pks[1]).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to get account information: %s\n", err)
		return
	}

	assetID := uint64(0)
	//	find newest (highest) asset for this account
	for _, asset := range act.CreatedAssets {
		if asset.Index > assetID {
			assetID = asset.Index
		}
	}

	// print created asset and asset holding info for this asset
	fmt.Printf("Asset ID: %d\n", assetID)
	printCreatedAsset(assetID, pks[1], algodClient)
	printAssetHolding(assetID, pks[1], algodClient)
	// Your output should look similar to this...
	
	// Asset created AssetName: latinum
	// Transaction ID: BEBUEATOOWSYDKN7W56Y2DHRED2Q45Z3M6ENGU4OWWMETC6CFW7Q
	// Submitted transaction BEBUEATOOWSYDKN7W56Y2DHRED2Q45Z3M6ENGU4OWWMETC6CFW7Q
	// waiting for confirmation
	// Transaction BEBUEATOOWSYDKN7W56Y2DHRED2Q45Z3M6ENGU4OWWMETC6CFW7Q confirmed in round 4086072
	// Asset ID: 2654040
	// {
	// 	"index": 2654040,
	// 	"params": {
	// 		"clawback": "AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU",
	// 		"creator": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM",
	// 		"decimals": 0,
	// 		"freeze": "AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU",
	// 		"manager": "AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU",
	// 		"metadata-hash": "dGhpc0lzU29tZUxlbmd0aDMySGFzaENvbW1pdG1lbnQ=",
	// 		"name": "latinum",
	// 		"reserve": "AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU",
	// 		"total": 1000,
	// 		"unit-name": "latinum",
	// 		"url": "https://path/to/my/asset/details"
	// 	}
	// } 
	// {
	// 	"amount": 1000,
	// 	"asset-id": 2654040,
	// 	"creator": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM"
	// } 

    // CHANGE MANAGER
	// Change Asset Manager from Account 2 to Account 1
	// assetID := uint64(332920)
	// Get network-related transaction parameters and assign
	txParams, err = algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	txParams.FlatFee = true
	txParams.Fee = 1000

	manager = pks[1]
	oldmanager := pks[2]
	strictEmptyAddressChecking := true
	txn, err = transaction.MakeAssetConfigTxn(oldmanager, note, txParams, assetID, manager, reserve, freeze, clawback, strictEmptyAddressChecking)
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
	sendResponse, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Transaction ID raw: %s\n", txid)

	// Wait for transaction to be confirmed
	waitForConfirmation(txid,algodClient )
	// print created assetinfo for this asset
	fmt.Printf("Asset ID: %d\n", assetID)
	printCreatedAsset(assetID, pks[1], algodClient)


	// Your terminal output should appear similar to this...

	// Transaction ID: 2XR5UANBPZ74MA5FCK2TB7KGOEEX3C3PFQEBBFYQ4UOBONR764VA
	// Transaction ID raw: 2XR5UANBPZ74MA5FCK2TB7KGOEEX3C3PFQEBBFYQ4UOBONR764VA
	// waiting for confirmation
	// Transaction 2XR5UANBPZ74MA5FCK2TB7KGOEEX3C3PFQEBBFYQ4UOBONR764VA confirmed in round 4086076
	// Asset ID: 2654040
	// {
	// 	"index": 2654040,
	// 	"params": {
	// 		"clawback": "AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU",
	// 		"creator": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM",
	// 		"decimals": 0,
	// 		"freeze": "AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU",
	// 		"manager": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM",
	// 		"metadata-hash": "dGhpc0lzU29tZUxlbmd0aDMySGFzaENvbW1pdG1lbnQ=",
	// 		"name": "latinum",
	// 		"reserve": "AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU",
	// 		"total": 1000,
	// 		"unit-name": "latinum",
	// 		"url": "https://path/to/my/asset/details"
	// 	}
	// } 

	// OPT-IN

	// Account 3 opts in to receive latinum
	// Use previously set transaction parameters and update sending address to account 3
	// assetID := uint64(332920)
	// Get network-related transaction parameters and assign
	txParams, err = algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	txParams.FlatFee = true
	txParams.Fee = 1000

	txn, err = transaction.MakeAssetAcceptanceTxn(pks[3], note, txParams, assetID)
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
	sendResponse, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Transaction ID raw: %s\n", txid)

	// Wait for transaction to be confirmed
	waitForConfirmation(txid, algodClient)

	// print created assetholding for this asset and Account 3, showing 0 balance
	fmt.Printf("Asset ID: %d\n", assetID)
	fmt.Printf("Account 3: %s\n", pks[3])
	printAssetHolding(assetID, pks[3], algodClient)

	// your terminal output should be similar to this...

	// Transaction ID: JYVJEB25YMAVNSAFDTZECWMJTKZHSFJGICGGXF64TH5RTXDICIUA
	// Transaction ID raw: JYVJEB25YMAVNSAFDTZECWMJTKZHSFJGICGGXF64TH5RTXDICIUA
	// waiting for confirmation
	// Transaction JYVJEB25YMAVNSAFDTZECWMJTKZHSFJGICGGXF64TH5RTXDICIUA confirmed in round 4086079
	// Asset ID: 2654040
	// Account 3: 3ZQ3SHCYIKSGK7MTZ7PE7S6EDOFWLKDQ6RYYVMT7OHNQ4UJ774LE52AQCU
	// {
	// 	"amount": 0,
	// 	"asset-id": 2654040,
	// 	"creator": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM"
	// } 

	// TRANSFER ASSET
	
	// Send  10 latinum from Account 1 to Account 3
	// assetID := uint64(332920)
	// Get network-related transaction parameters and assign
	txParams, err = algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	txParams.FlatFee = true
	txParams.Fee = 1000

	sender := pks[1]
	recipient := pks[3]
	amount := uint64(10)
	closeRemainderTo := ""
	txn, err = transaction.MakeAssetTransferTxn(sender, recipient, amount, note, txParams, closeRemainderTo, 
		assetID)
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
	sendResponse, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Transaction ID raw: %s\n", txid)

	// Wait for transaction to be confirmed
	waitForConfirmation(txid,algodClient)

	// print created assetholding for this asset and Account 3 and Account 1
	// You should see amount of 10 in Account 3, and 990 in Account 1
	fmt.Printf("Asset ID: %d\n", assetID)
	fmt.Printf("Account 3: %s\n", pks[3])
	printAssetHolding(assetID, pks[3], algodClient)
	fmt.Printf("Account 1: %s\n", pks[1])
	printAssetHolding(assetID, pks[1], algodClient)
	// Your terminal output should look similar to this
	// Transaction ID: 7GPXSVF6YYHHHIGHDCGGR2AS2XXLMDXTUR6GUTSZU4GMIOK2V7TQ
	// Transaction ID raw: 7GPXSVF6YYHHHIGHDCGGR2AS2XXLMDXTUR6GUTSZU4GMIOK2V7TQ
	// waiting for confirmation
	// Transaction 7GPXSVF6YYHHHIGHDCGGR2AS2XXLMDXTUR6GUTSZU4GMIOK2V7TQ confirmed in round 4086081
	// Asset ID: 2654040
	// Account 3: 3ZQ3SHCYIKSGK7MTZ7PE7S6EDOFWLKDQ6RYYVMT7OHNQ4UJ774LE52AQCU
	// {
	// 	"amount": 10,
	// 	"asset-id": 2654040,
	// 	"creator": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM"
	// } 
	// Account 1: THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM
	// {
	// 	"amount": 990,
	// 	"asset-id": 2654040,
	// 	"creator": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM"
	// } 

	// FREEZE ASSET
	// The freeze address (Account 2) Freeze's asset for Account 3.
	// assetID := uint64(332920)
	// Get network-related transaction parameters and assign
	txParams, err = algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	txParams.FlatFee = true
	txParams.Fee = 1000
	newFreezeSetting := true
	target := pks[3]
	txn, err = transaction.MakeAssetFreezeTxn(freeze, note, txParams, assetID, target, newFreezeSetting)
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
	sendResponse, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Transaction ID raw: %s\n", txid)
	// Wait for transaction to be confirmed
	waitForConfirmation(txid,algodClient)
    // You should now see is-frozen value of true
	fmt.Printf("Asset ID: %d\n", assetID)
	fmt.Printf("Account 3: %s\n", pks[3])
	printAssetHolding(assetID, pks[3], algodClient)

	

	// Your terminal output should look similar to this:

	// Transaction ID: FHFLUVKQ5Q4S2RRLOA6EJ6NVQDZEVU6TDKNOVJK5ZNKCDYUZFNXQ
	// Transaction ID raw: FHFLUVKQ5Q4S2RRLOA6EJ6NVQDZEVU6TDKNOVJK5ZNKCDYUZFNXQ
	// waiting for confirmation
	// Transaction FHFLUVKQ5Q4S2RRLOA6EJ6NVQDZEVU6TDKNOVJK5ZNKCDYUZFNXQ confirmed in round 4086084
	// Asset ID: 2654040
	// Account 3: 3ZQ3SHCYIKSGK7MTZ7PE7S6EDOFWLKDQ6RYYVMT7OHNQ4UJ774LE52AQCU
	// {
	// 	"amount": 10,
	// 	"asset-id": 2654040,
	// 	"creator": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM",
	// 	"is-frozen": true
	// }
	
	// REVOKE ASSET
	// Revoke an Asset
	// The clawback address (Account 2) revokes 10 latinum from Account 3 (target)
	// and places it back with Account 1 (creator).
	// assetID := uint64(332920)
	// Get network-related transaction parameters and assign
	txParams, err = algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	txParams.FlatFee = true
	txParams.Fee = 1000
	target = pks[3]
	txn, err = transaction.MakeAssetRevocationTxn(clawback, target, amount, creator, note,
		txParams, assetID)
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
	sendResponse, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Transaction ID raw: %s\n", txid)
	// Wait for transaction to be confirmed
	waitForConfirmation( txid, algodClient)
	// print created assetholding for this asset and Account 3 and Account 1
	// You should see amount of 0 in Account 3, and 1000 in Account 1
	fmt.Printf("Asset ID: %d\n", assetID)
	fmt.Printf("recipient")
	fmt.Printf("Account 3: %s\n", pks[3])
	printAssetHolding(assetID, pks[3], algodClient)
	fmt.Printf("target")
	fmt.Printf("Account 1: %s\n", pks[1])
	printAssetHolding(assetID, pks[1], algodClient)

	// Your terminal output should look similar to this...

	// Transaction XH32YUIX2VTEH3QPJECVNVXHVHU2LBQVGIHMPPSRE4XGLFNUG63Q confirmed in round 4086090
	// Asset ID: 2654040
	// recipientAccount 3: 3ZQ3SHCYIKSGK7MTZ7PE7S6EDOFWLKDQ6RYYVMT7OHNQ4UJ774LE52AQCU
	// {
	// 	"amount": 0,
	// 	"asset-id": 2654040,
	// 	"creator": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM",
	// 	"is-frozen": true
	// } 
	// targetAccount 1: THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM
	// {
	// 	"amount": 1000,
	// 	"asset-id": 2654040,
	// 	"creator": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM"
	// }

	// DESTROY ASSET
	// Destroy the asset
	// Make sure all funds are back in the creator's account. Then use the
	// Manager account (Account 1) to destroy the asset.

	// assetID := uint64(332920)
	// Get network-related transaction parameters and assign
	txParams, err = algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	txParams.FlatFee = true
	txParams.Fee = 1000

	txn, err = transaction.MakeAssetDestroyTxn(manager, note, txParams, assetID)
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
	sendResponse, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Transaction ID raw: %s\n", txid)
	// Wait for transaction to be confirmed
	waitForConfirmation(txid,algodClient)
	fmt.Printf("Asset ID: %d\n", assetID)	
	fmt.Printf("Account 3 must do a transaction for an amount of 0, \n" )
    fmt.Printf("with a closeRemainderTo to the creator account, to clear it from its accountholdings. \n")
    fmt.Printf("For Account 1, nothing should print after this as the asset is destroyed on the creator account \n")

	// print created asset and asset holding info for this asset (should not print anything)

	printCreatedAsset(assetID, pks[1], algodClient)
	printAssetHolding(assetID, pks[1], algodClient)

	// Your terminal output should look similar to this...

	// Transaction PI4U7DJZYDKEZS2PKTNGB6DFNVCCEYN5FNLZBBWNONTWMA7RH6AA confirmed in round 4086093
	// Asset ID: 2654040
	// Account 3 must do a transaction for an amount of 0, 
	// with a closeRemainderTo to the creator account, to clear it from its accountholdings.
	// For Account 1, nothing should print after this as the asset is destroyed on the creator account
}
