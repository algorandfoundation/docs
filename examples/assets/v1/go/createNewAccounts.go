package main

import (
	"fmt"
	json "encoding/json"
	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/mnemonic"
	"github.com/algorand/go-algorand-sdk/client/algod"
)

// const algodAddress = "Your Address"
// const algodToken = "Your Token"

const algodAddress = "http://localhost:4001"
const algodToken = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
var txHeaders = append([]*algod.Header{}, &algod.Header{"Content-Type", "application/json"})

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
	account1 := crypto.GenerateAccount()
	account2 := crypto.GenerateAccount()
	account3 := crypto.GenerateAccount()
	address1 := account1.Address.String()
	address2 := account2.Address.String()
	address3 := account3.Address.String()

	mnemonic1, err := mnemonic.FromPrivateKey(account1.PrivateKey)
	if err != nil {
		return
	}
	mnemonic2, err := mnemonic.FromPrivateKey(account2.PrivateKey)
	if err != nil {
		return
	}
	mnemonic3, err := mnemonic.FromPrivateKey(account3.PrivateKey)
	if err != nil {
		return
	}
	fmt.Printf("1 : \"%s\"\n", address1)
	fmt.Printf("2 : \"%s\"\n", address2)
	fmt.Printf("3 : \"%s\"\n", address3)
	fmt.Printf("")
	fmt.Printf("Copy off accounts above and add TestNet Algo funds using the TestNet Dispenser at https://bank.testnet.algorand.network/\n")
	fmt.Printf("Copy off the following mnemonic code for future tutorial use\n")
	fmt.Printf("\n")
	fmt.Printf("mnemonic1 := \"%s\"\n", mnemonic1)
	fmt.Printf("mnemonic2 := \"%s\"\n", mnemonic2)
	fmt.Printf("mnemonic3 := \"%s\"\n", mnemonic3)

	// Initialize an algodClient
	algodClient, err := algod.MakeClient(algodAddress, algodToken)
	if err != nil {
		return
	}
	act, err := algodClient.AccountInformation(account1.Address.String(), txHeaders...)
	if err != nil {
		fmt.Printf("failed to get account information: %s\n", err)
		return
	}
	fmt.Print("Account 1: ")
	PrettyPrint(act)
	act, err = algodClient.AccountInformation(account2.Address.String(), txHeaders...)
	if err != nil {
		fmt.Printf("failed to get account information: %s\n", err)
		return
	}	
    fmt.Print("Account 2: ")
	PrettyPrint(act)
	act, err = algodClient.AccountInformation(account3.Address.String(), txHeaders...)
	if err != nil {
		fmt.Printf("failed to get account information: %s\n", err)
		return
	}
    fmt.Print("Account 3: ")
	PrettyPrint(act)
}

// Terminal output should look similar to this...

// 1 : "6RYS2DWSSAYVW5PSOB7WO35ETWVVP4BYKK2NLW4KHGUTHIA63AG3Z7GZIM"
// 2 : "TVIGQH4QAUQR6FBYFLKVPYUSAVYAKWF7TUL5FM3XVX47755ZHNRMCKISRI"
// 3 : "R6VFJ6E4XXSE5VX4FBMREU67AQETJHKH34OV3HXL4VBIXSUX4UHEHA65QQ"
// Copy off accounts above and add TestNet Algo funds using the TestNet Dispenser at https://bank.testnet.algorand.network/
// Copy off the following mnemonic code for use in Step 1B

// mnemonic1 := "scale damage enter shoot enrich traffic sunny evidence recycle budget hello sorry abstract acoustic foot enlist awful blame arch illness cash victory divide absorb ozone"
// mnemonic2 := "culture person industry summer asthma wood police truth design dinner virtual unfold dismiss flavor gospel emerge artwork theory group slide welcome immense direct abstract helmet"
// mnemonic3 := "method trumpet novel turkey super chat orchard resist unit evidence grow drift faint apology position easily frog armor cover anger list angle arrange absent project"
// Account 1: {
// 	"round": 5985502,
// 	"address": "6RYS2DWSSAYVW5PSOB7WO35ETWVVP4BYKK2NLW4KHGUTHIA63AG3Z7GZIM",
// 	"amount": 100000000,
// 	"pendingrewards": 0,
// 	"amountwithoutpendingrewards": 100000000,
// 	"rewards": 0,
// 	"status": "Offline"
// } 
// Account 2: {
// 	"round": 5985502,
// 	"address": "TVIGQH4QAUQR6FBYFLKVPYUSAVYAKWF7TUL5FM3XVX47755ZHNRMCKISRI",
// 	"amount": 100000000,
// 	"pendingrewards": 0,
// 	"amountwithoutpendingrewards": 100000000,
// 	"rewards": 0,
// 	"status": "Offline"
// } 
// Account 3: {
// 	"round": 5985502,
// 	"address": "R6VFJ6E4XXSE5VX4FBMREU67AQETJHKH34OV3HXL4VBIXSUX4UHEHA65QQ",
// 	"amount": 100000000,
// 	"pendingrewards": 0,
// 	"amountwithoutpendingrewards": 100000000,
// 	"rewards": 0,
// 	"status": "Offline"
// } 