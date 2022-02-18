package main

import (
	"context"
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/future"
	"github.com/algorand/go-algorand-sdk/mnemonic"
	"github.com/algorand/go-algorand-sdk/types"
)

// user defined algod client settings
const algodAddress = "http://localhost:4001"
const algodToken = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

// user defined mnemonics
// do not use mnemonics in production code - for demo purposes only
const creatorMnemonic = "your 25-word mnemonic goes right here"
const userMnemonic = "your second distinct 25-word mnemonic goes right here"


// declare application state storage (immutable)
const localInts = 1
const localBytes = 1
const globalInts = 1
const globalBytes = 0

// user declared approval program (initial)
const approvalProgramSourceInitial = `#pragma version 5
// Handle each possible OnCompletion type. We don't have to worry about
// handling ClearState, because the ClearStateProgram will execute in that
// case, not the ApprovalProgram.
txn OnCompletion
int NoOp
==
bnz handle_noop
txn OnCompletion
int OptIn
==
bnz handle_optin
txn OnCompletion
int CloseOut
==
bnz handle_closeout
txn OnCompletion
int UpdateApplication
==
bnz handle_updateapp
txn OnCompletion
int DeleteApplication
==
bnz handle_deleteapp
// Unexpected OnCompletion value. Should be unreachable.
err
handle_noop:
// Handle NoOp
// Check for creator
addr LD6R3YLIIIEQK5VEYXNXBR775EV4DEOLZB6R7WUZGOTCB2SMJEZTFRPFPY
txn Sender
==
bnz handle_optin
// read global state
byte "counter"
dup
app_global_get
// increment the value
int 1
+
// store to scratch space
dup
store 0
// update global state
app_global_put
// read local state for sender
int 0
byte "counter"
app_local_get
// increment the value
int 1
+
store 1
// update local state for sender
int 0
byte "counter"
load 1
app_local_put
// load return value as approval
load 0
return
handle_optin:
// Handle OptIn
// approval
int 1
return
handle_closeout:
// Handle CloseOut
//approval
int 1
return
handle_deleteapp:
// Check for creator
addr LD6R3YLIIIEQK5VEYXNXBR775EV4DEOLZB6R7WUZGOTCB2SMJEZTFRPFPY
txn Sender
==
return
handle_updateapp:
// Check for creator
addr LD6R3YLIIIEQK5VEYXNXBR775EV4DEOLZB6R7WUZGOTCB2SMJEZTFRPFPY
txn Sender
==
return
`

// user declared approval program (refactored)
const approvalProgramSourceRefactored = `#pragma version 5
// Handle each possible OnCompletion type. We don't have to worry about
// handling ClearState, because the ClearStateProgram will execute in that
// case, not the ApprovalProgram.
txn OnCompletion
int NoOp
==
bnz handle_noop
txn OnCompletion
int OptIn
==
bnz handle_optin
txn OnCompletion
int CloseOut
==
bnz handle_closeout
txn OnCompletion
int UpdateApplication
==
bnz handle_updateapp
txn OnCompletion
int DeleteApplication
==
bnz handle_deleteapp
// Unexpected OnCompletion value. Should be unreachable.
err
handle_noop:
// Handle NoOp
// Check for creator
addr 5XWY6RBNYHCSY2HK5HCTO62DUJJ4PT3G4L77FQEBUKE6ZYRGQAFTLZSQQ4
txn Sender
==
bnz handle_optin
// read global state
byte "counter"
dup
app_global_get
// increment the value
int 1
+
// store to scratch space
dup
store 0
// update global state
app_global_put
// read local state for sender
int 0
byte "counter"
app_local_get
// increment the value
int 1
+
store 1
// update local state for sender
// update "counter"
int 0
byte "counter"
load 1
app_local_put
// update "timestamp"
int 0
byte "timestamp"
txn ApplicationArgs 0
app_local_put
// load return value as approval
load 0
return
handle_optin:
// Handle OptIn
// approval
int 1
return
handle_closeout:
// Handle CloseOut
//approval
int 1
return
handle_deleteapp:
// Check for creator
addr LD6R3YLIIIEQK5VEYXNXBR775EV4DEOLZB6R7WUZGOTCB2SMJEZTFRPFPY
txn Sender
==
return
handle_updateapp:
// Check for creator
addr LD6R3YLIIIEQK5VEYXNXBR775EV4DEOLZB6R7WUZGOTCB2SMJEZTFRPFPY
txn Sender
==
return
`

// declare clear state program source
const clearProgramSource = `#pragma version 5
int 1
`



// utility function to recover account and return sk and address
func recoverAccount(passphrase string) crypto.Account {

	sk, err := mnemonic.ToPrivateKey(passphrase)
	if err != nil {
		fmt.Printf("error recovering account: %s\n", err)
	}
	pk := sk.Public()
	var addr types.Address
	cpk := pk.(ed25519.PublicKey)
	copy(addr[:], cpk[:])
	fmt.Printf("Address: %s\n", addr.String())
	account := crypto.GenerateAccount()
	account.Address = addr
	account.PrivateKey = sk
	return account
}

// helper function to compile program source
func compileProgram(client *algod.Client, programSource string) (compiledProgram []byte) {
	compileResponse, err := client.TealCompile([]byte(programSource)).Do(context.Background())
	if err != nil {
		fmt.Printf("Issue with compile: %s\n", err)
		return
	}
	compiledProgram, _ = base64.StdEncoding.DecodeString(compileResponse.Result)
	return compiledProgram
}

func createApp(client *algod.Client, creatorAccount crypto.Account, approvalProgram []byte, clearProgram []byte, localInts int, localBytes int, globalInts int, globalBytes int) (appId uint64) {
	// define schema
	globalSchema := types.StateSchema{NumUint: uint64(globalInts), NumByteSlice: uint64(globalBytes)}
	localSchema := types.StateSchema{NumUint: uint64(localInts), NumByteSlice: uint64(localBytes)}

	// get transaction suggested parameters
	params, err := client.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	// params.FlatFee = true
	// params.Fee = 1000

	// create unsigned transaction
	txn, err := future.MakeApplicationCreateTx(false, approvalProgram, clearProgram, globalSchema, localSchema, nil,
		nil, nil, nil, params, creatorAccount.Address, nil, types.Digest{}, [32]byte{}, types.Address{})
	if err != nil {
		fmt.Printf("Error creating transaction: %s\n", err)
		return
	}

	// Sign the transaction
	txID, signedTxn, err := crypto.SignTransaction(creatorAccount.PrivateKey, txn)
	if err != nil {
		fmt.Printf("Failed to sign transaction: %s\n", err)
		return
	}
	fmt.Printf("Signed txid: %s\n", txID)

	// Submit the transaction
	sendResponse, err := client.SendRawTransaction(signedTxn).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Submitted transaction %s\n", sendResponse)


	// Wait for confirmation
	confirmedTxn, err := future.WaitForConfirmation(client, txID,  4, context.Background())
	if err != nil {
		fmt.Printf("Error waiting for confirmation on txID: %s\n", txID)
		return
	}
	fmt.Printf("Confirmed Transaction: %s in Round %d\n", txID ,confirmedTxn.ConfirmedRound)

	// display results

	appId = confirmedTxn.ApplicationIndex
	fmt.Printf("Created new app-id: %d\n", appId)
	return appId
}

// opt-in to application
func optInApp(client *algod.Client, account crypto.Account, index uint64) {
	// get transaction suggested parameters
	params, err := client.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	// params.FlatFee = true
	// params.Fee = 1000

	// create unsigned transaction
	txn, err := future.MakeApplicationOptInTx(index, nil, nil, nil, nil, params, account.Address, nil, types.Digest{}, [32]byte{}, types.Address{})
	if err != nil {
		fmt.Printf("Error creating transaction: %s\n", err)
		return
	}

	// Sign the transaction
	txID, signedTxn, err := crypto.SignTransaction(account.PrivateKey, txn)
	if err != nil {
		fmt.Printf("Failed to sign transaction: %s\n", err)
		return
	}
	fmt.Printf("Signed txid: %s\n", txID)

	// Submit the transaction
	sendResponse, err := client.SendRawTransaction(signedTxn).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Submitted transaction %s\n", sendResponse)


	// Wait for confirmation
	confirmedTxn, err := future.WaitForConfirmation(client, txID,  4, context.Background())
	if err != nil {
		fmt.Printf("Error waiting for confirmation on txID: %s\n", txID)
		return
	}
	fmt.Printf("Confirmed Transaction: %s in Round %d\n", txID ,confirmedTxn.ConfirmedRound)

	// display results

	fmt.Printf("Oped-in to app-id: %d\n", confirmedTxn.Transaction.Txn.ApplicationID)
}

func callApp(client *algod.Client, account crypto.Account, index uint64, appArgs [][]byte) {
	// get transaction suggested parameters
	params, err := client.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	// params.FlatFee = true
	// params.Fee = 1000

	// create unsigned transaction
	txn, err := future.MakeApplicationNoOpTx(index, appArgs, nil, nil, nil, params, account.Address, nil, types.Digest{}, [32]byte{}, types.Address{})
	if err != nil {
		fmt.Printf("Error creating transaction: %s\n", err)
		return
	}

	// Sign the transaction
	txID, signedTxn, err := crypto.SignTransaction(account.PrivateKey, txn)
	if err != nil {
		fmt.Printf("Failed to sign transaction: %s\n", err)
		return
	}
	fmt.Printf("Signed txid: %s\n", txID)

	// Submit the transaction
	sendResponse, err := client.SendRawTransaction(signedTxn).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Submitted transaction %s\n", sendResponse)


	// Wait for confirmation
	confirmedTxn, err := future.WaitForConfirmation(client, txID,  4, context.Background())
	if err != nil {
		fmt.Printf("Error waiting for confirmation on txID: %s\n", txID)
		return
	}
	fmt.Printf("Confirmed Transaction: %s in Round %d\n", txID ,confirmedTxn.ConfirmedRound)

	// display results
	fmt.Printf("Called app-id: %d\n", confirmedTxn.Transaction.Txn.ApplicationID)

}

func readLocalState(client *algod.Client, account crypto.Account, index uint64) {
	accountInfo, err := client.AccountInformation(account.Address.String()).Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting account info: %s\n", err)
		return
	}
	for _, ap := range accountInfo.AppsLocalState {
		if ap.Id == index {
			fmt.Printf("Local state for app-id %d (account %s):\n", ap.Id, account.Address.String())
			fmt.Println(ap.KeyValue)
		}
	}
}

func readGlobalState(client *algod.Client, account crypto.Account, index uint64) {
	accountInfo, err := client.AccountInformation(account.Address.String()).Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting account info: %s\n", err)
		return
	}
	for _, ap := range accountInfo.CreatedApps {
		if ap.Id == index {
			fmt.Printf("Global state for app-id %d:\n", ap.Id)
			fmt.Println(ap.Params.GlobalState)
		}
	}
}

func updateApp(client *algod.Client, creatorAccount crypto.Account, index uint64, approvalProgram []byte, clearProgram []byte) {
	// get transaction suggested parameters
	params, err := client.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	// params.FlatFee = true
	// params.Fee = 1000

	// create unsigned transaction
	txn, err := future.MakeApplicationUpdateTx(index, nil, nil, nil, nil, approvalProgram, clearProgram, params, creatorAccount.Address, nil, types.Digest{}, [32]byte{}, types.Address{})
	if err != nil {
		fmt.Printf("Error creating transaction: %s\n", err)
		return
	}

	// Sign the transaction
	txID, signedTxn, err := crypto.SignTransaction(creatorAccount.PrivateKey, txn)
	if err != nil {
		fmt.Printf("Failed to sign transaction: %s\n", err)
		return
	}
	fmt.Printf("Signed txid: %s\n", txID)

	// Submit the transaction
	sendResponse, err := client.SendRawTransaction(signedTxn).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Submitted transaction %s\n", sendResponse)


	// Wait for confirmation
	confirmedTxn, err := future.WaitForConfirmation(client, txID,  4, context.Background())
	if err != nil {
		fmt.Printf("Error waiting for confirmation on txID: %s\n", txID)
		return
	}
	fmt.Printf("Confirmed Transaction: %s in Round %d\n", txID ,confirmedTxn.ConfirmedRound)

	// display results
	fmt.Printf("Updated app-id: %d\n", confirmedTxn.Transaction.Txn.ApplicationID)
}

func closeOutApp(client *algod.Client, account crypto.Account, index uint64) {
	// get transaction suggested parameters
	params, err := client.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	// params.FlatFee = true
	// params.Fee = 1000

	// create unsigned transaction
	txn, err := future.MakeApplicationCloseOutTx(index, nil, nil, nil, nil, params, account.Address, nil, types.Digest{}, [32]byte{}, types.Address{})
	if err != nil {
		fmt.Printf("Error creating transaction: %s\n", err)
		return
	}

	// Sign the transaction
	txID, signedTxn, err := crypto.SignTransaction(account.PrivateKey, txn)
	if err != nil {
		fmt.Printf("Failed to sign transaction: %s\n", err)
		return
	}
	fmt.Printf("Signed txid: %s\n", txID)

	// Submit the transaction
	sendResponse, err := client.SendRawTransaction(signedTxn).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Submitted transaction %s\n", sendResponse)


	// Wait for confirmation
	confirmedTxn, err := future.WaitForConfirmation(client, txID,  4, context.Background())
	if err != nil {
		fmt.Printf("Error waiting for confirmation on txID: %s\n", txID)
		return
	}
	fmt.Printf("Confirmed Transaction: %s in Round %d\n", txID ,confirmedTxn.ConfirmedRound)

	// display results

	fmt.Printf("Closed out from app-id: %d\n", confirmedTxn.Transaction.Txn.ApplicationID)
}

func deleteApp(client *algod.Client, account crypto.Account, index uint64) {
	// get transaction suggested parameters
	params, err := client.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	// params.FlatFee = true
	// params.Fee = 1000

	// create unsigned transaction
	txn, err := future.MakeApplicationDeleteTx(index, nil, nil, nil, nil, params, account.Address, nil, types.Digest{}, [32]byte{}, types.Address{})
	if err != nil {
		fmt.Printf("Error creating transaction: %s\n", err)
		return
	}

	// Sign the transaction
	txID, signedTxn, err := crypto.SignTransaction(account.PrivateKey, txn)
	if err != nil {
		fmt.Printf("Failed to sign transaction: %s\n", err)
		return
	}
	fmt.Printf("Signed txid: %s\n", txID)

	// Submit the transaction
	sendResponse, err := client.SendRawTransaction(signedTxn).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Submitted transaction %s\n", sendResponse)


	// Wait for confirmation
	confirmedTxn, err := future.WaitForConfirmation(client, txID,  4, context.Background())
	if err != nil {
		fmt.Printf("Error waiting for confirmation on txID: %s\n", txID)
		return
	}
	fmt.Printf("Confirmed Transaction: %s in Round %d\n", txID ,confirmedTxn.ConfirmedRound)

	// display results
	fmt.Printf("Deleted app-id: %d\n", confirmedTxn.Transaction.Txn.ApplicationID)
}

func clearApp(client *algod.Client, account crypto.Account, index uint64) {
	// get transaction suggested parameters
	params, err := client.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
	}
	// comment out the next two (2) lines to use suggested fees
	// params.FlatFee = true
	// params.Fee = 1000

	// create unsigned transaction
	txn, err := future.MakeApplicationClearStateTx(index, nil, nil, nil, nil, params, account.Address, nil, types.Digest{}, [32]byte{}, types.Address{})
	if err != nil {
		fmt.Printf("Error creating transaction: %s\n", err)
		return
	}

	// Sign the transaction
	txID, signedTxn, err := crypto.SignTransaction(account.PrivateKey, txn)
	if err != nil {
		fmt.Printf("Failed to sign transaction: %s\n", err)
		return
	}
	fmt.Printf("Signed txid: %s\n", txID)

	// Submit the transaction
	sendResponse, err := client.SendRawTransaction(signedTxn).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Submitted transaction %s\n", sendResponse)


	// Wait for confirmation
	confirmedTxn, err := future.WaitForConfirmation(client, txID,  4, context.Background())
	if err != nil {
		fmt.Printf("Error waiting for confirmation on txID: %s\n", txID)
		return
	}
	fmt.Printf("Confirmed Transaction: %s in Round %d\n", txID ,confirmedTxn.ConfirmedRound)

	// display results

	fmt.Printf("Cleared local state for app-id: %d\n", confirmedTxn.Transaction.Txn.ApplicationID)
}

func main() {
	// initialize an algodClient
	algodClient, err := algod.MakeClient(algodAddress, algodToken)
	if err != nil {
		fmt.Printf("Issue with creating algod client: %s\n", err)
		return
	}

	// get accounts from mnemonic
	creatorAccount := recoverAccount(creatorMnemonic)
	userAccount := recoverAccount(userMnemonic)

	// compile programs
	approvalProgram := compileProgram(algodClient, approvalProgramSourceInitial)
	clearProgram := compileProgram(algodClient, clearProgramSource)

	// create new application
	appId := createApp(algodClient, creatorAccount, approvalProgram, clearProgram, localInts, localBytes, globalInts, globalBytes)
	//var appId uint64 = 285
	//println(appId)

	// opt-in to application
	optInApp(algodClient, userAccount, appId)

	// call application without arguments
	callApp(algodClient, userAccount, appId, nil)

	// read local state of application from user account
	readLocalState(algodClient, userAccount, appId)

	// read global state of application
	readGlobalState(algodClient, creatorAccount, appId)

	// update application
	approvalProgram = compileProgram(algodClient, approvalProgramSourceRefactored)
	updateApp(algodClient, creatorAccount, appId, approvalProgram, clearProgram)

	// call application with arguments
	now := time.Now().Format("Mon Jan _2 15:04:05 2006")
	appArgs := make([][]byte, 1)
	appArgs[0] = []byte(now)
	callApp(algodClient, userAccount, appId, appArgs)

	// read local state of application from user account
	readLocalState(algodClient, userAccount, appId)

	// close-out from application
	closeOutApp(algodClient, userAccount, appId)

	// opt-in again to application
	optInApp(algodClient, userAccount, appId)

	// call application with arguments
	callApp(algodClient, userAccount, appId, appArgs)

	// read local state of application from user account
	readLocalState(algodClient, userAccount, appId)

	// delete application
	deleteApp(algodClient, creatorAccount, appId)

	// clear application from user account
	clearApp(algodClient, userAccount, appId)

}