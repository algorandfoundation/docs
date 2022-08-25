title: Your First Transaction

This section is a quick start guide for sending your first transaction on the Algorand TestNet network using the Go programming language. This guide installs the Go SDK, creates an account and submits a payment transaction. This guide also installs Algorand Sandbox, which provides required infrastructure for development and testing. 

!!! Info
    If you are a visual learner, try our [live demo](https://replit.com/@Algorand/Getting-Started-with-Go) or watch a [video walkthrough](https://youtu.be/rFG7Zo2JvIY?t=) explaining all the code in the steps below.
 
# Install Algorand Sandbox

Algorand Sandbox is developer-focused tool for quickly spinning up the Algorand infrastructure portion of your development environment. It uses Docker to provide an `algod` instance for connecting to the network of your choosing and an `indexer` instance for querying blockchain data. APIs are exposed by both instances for client access provided within the SDK. Read more about [Algorand networks](../../get-details/algorand-networks/index.md), their capabilities and intended use.

!!! Prerequisites
    - Docker Compose ([install guide](https://docs.docker.com/compose/install/))
    - Git ([install guide](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git))

From a terminal window, install Algorand Sandbox connected to TestNet:

```bash
git clone https://github.com/algorand/sandbox.git
cd sandbox
./sandbox up testnet
```

!!! Warning
    The Algorand Sandbox installation may take a few minutes to complete in order to catch up to the current round on TestNet. To learn more about fast catchup, see [Sync Node Network using Fast Catchup](https://developer.algorand.org/docs/run-a-node/setup/install/#sync-node-network-using-fast-catchup).

!!! Info
    The `indexer` is enabled **only** for _private networks_. Therefore, all blockchain queries in this guide will use the `algod` API.

- [Watch Video](https://youtu.be/rFG7Zo2JvIY?t=18)
- [More Information](https://developer.algorand.org/articles/introducing-sandbox-20/)

# Install Go SDK

Algorand provides an SDK for Go. 

!!! Prerequisites
    - Go programming language ([install guide](https://golang.org/doc/install))

From a terminal window, install the Go SDK:

```bash
go get -u github.com/algorand/go-algorand-sdk/...
```

- [`Watch Video`](https://youtu.be/rFG7Zo2JvIY?t=88)
- [`More Information`](https://github.com/algorand/go-algorand-sdk)
 
The SDK is installed and can now interact with the running Algorand Sandbox environment, as configured above.

# Create account

In order to interact with the Algorand blockchain, you must have a funded account on the network. To quickly create an account on Algorand TestNet create a new file **yourFirstTransaction.go** and insert the following code:

```go linenums="1"
package main

import (
    "context"
    json "encoding/json"
    "fmt"
	"github.com/algorand/go-algorand-sdk/future"
    "github.com/algorand/go-algorand-sdk/client/v2/algod"
    "github.com/algorand/go-algorand-sdk/crypto"
    "github.com/algorand/go-algorand-sdk/mnemonic"
    "github.com/algorand/go-algorand-sdk/transaction"
)

// TODO: insert additional utility functions here

func main() {
    // Create account
    account := crypto.GenerateAccount()
    passphrase, err := mnemonic.FromPrivateKey(account.PrivateKey)
    myAddress := account.Address.String()

    if err != nil {
        fmt.Printf("Error creating transaction: %s\n", err)
    } else {
        fmt.Printf("My address: %s\n", myAddress)
        fmt.Printf("My passphrase: %s\n", passphrase)
        fmt.Println("--> Copy down your address and passphrase for future use.")
        fmt.Println("--> Once secured, press ENTER key to continue...")
        fmt.Scanln()
    }

    // TODO: insert additional codeblocks here
}
```

!!! Note 
    Lines 17 and 35 contain TODO: comments about inserting additional code. As you proceed with this guide, ensure the line numbers remain in sync.

!!! Tip
    Make sure to save the generated address and passphrase in a secure location, as they will be used later on.

!!! Warning 
    Never share your mnemonic passphrase or private keys. Production environments require stringent private key management. For more information on key management in community Wallets, click [here](https://developer.algorand.org/docs/community/#wallets). For the open source [Algorand Wallet](https://developer.algorand.org/articles/algorand-wallet-now-open-source/), click [here](https://github.com/algorand/algorand-wallet).

- [Watch Video](https://youtu.be/rFG7Zo2JvIY?t=97)
- [More Information](https://developer.algorand.org/docs/features/accounts/create/#standalone)
 
# Fund account

The code below prompts to fund the newly generated account. Before sending transactions to the Algorand network, the account must be funded to cover the minimal transaction fees that exist on Algorand. To fund the account use the [Algorand TestNet faucet](https://dispenser.testnet.aws.algodev.network/). 

```go linenums="35"
// Fund account
fmt.Println("Fund the created account using the Algorand TestNet faucet:\n--> https://dispenser.testnet.aws.algodev.network?account=" + myAddress)
fmt.Println("--> Once funded, press ENTER key to continue...")
fmt.Scanln()
```

!!! Info
    All Algorand accounts require a minimum balance to be registered in the ledger. To read more about Algorand minimum balance see [Account Overview](https://developer.algorand.org/docs/features/accounts/#minimum-balance)

- [Watch Video](https://youtu.be/rFG7Zo2JvIY?t=138)

# Instantiate client

You must instantiate a client prior to making calls to the API endpoints. The Go SDK implements the client natively using the following code:

```go  linenums="40"
// instantiate algod client to Algorand Sandbox
const algodAddress = "http://localhost:4001"
const algodToken = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

algodClient, err := algod.MakeClient(algodAddress, algodToken)
if err != nil {
    fmt.Printf("Issue with creating algod client: %s\n", err)
    return
}
```
 
!!! Info
    This guide provides values for `algodAddress` and `algodToken` as specified by Algorand Sandbox. If you want to connect to a third-party service provider, see [Purestake](https://developer.purestake.io/code-samples) or [AlgoExplorer Developer API](https://algoexplorer.io/api-dev/v2) and adjust these values accordingly.
 
- [Watch Video](https://youtu.be/rFG7Zo2JvIY?t=149)

# Check account balance

Before moving on to the next step, make sure your account has been funded by the faucet.
 
```go  linenums="50"
//Check account balance
fmt.Printf("My address: %s\n", myAddress)

accountInfo, err := algodClient.AccountInformation(myAddress).Do(context.Background())
if err != nil {
    fmt.Printf("Error getting account info: %s\n", err)
    return
}

fmt.Printf("Account balance: %d microAlgos\n", accountInfo.Amount)
fmt.Println("--> Ensure balance greater than 0, press ENTER key to continue...")
fmt.Scanln()
```

- [Watch Video](https://youtu.be/rFG7Zo2JvIY?t=161)

# Build transaction

Communication with the Algorand network is performed using transactions. Create a payment transaction sending 1 ALGO from your account to the TestNet faucet address:

```go linenums="62"
// Construct the transaction
txParams, err := algodClient.SuggestedParams().Do(context.Background())
if err != nil {
    fmt.Printf("Error getting suggested tx params: %s\n", err)
    return
}
fromAddr := myAddress
toAddr := "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A"


var amount uint64 = 1000000
var minFee uint64 = transaction.MinTxnFee 
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
```

!!! Info
    Algorand supports many transaction types. To see what types are supported see [Transactions](https://developer.algorand.org/docs/features/transactions/).

[`Watch Video`](https://youtu.be/rFG7Zo2JvIY?t=178)

# Sign transaction

Before the transaction is considered valid, it must be signed by a private key. Use the following code to sign the transaction.

```go linenums="83"
// Sign the transaction
txID, signedTxn, err := crypto.SignTransaction(account.PrivateKey, txn)
if err != nil {
    fmt.Printf("Failed to sign transaction: %s\n", err)
    return
}
fmt.Printf("Signed txid: %s\n", txID)
```

!!! Info
    Algorand provides many ways to sign transactions. To see other ways see [Authorization](https://developer.algorand.org/docs/features/transactions/signatures/#single-signatures).

[`Watch Video`](https://youtu.be/rFG7Zo2JvIY?t=204)

# Submit transaction

The signed transaction can now be broadcast to the network for validation and inclusion in a future block. The `waitForConfirmation` SDK method polls the `algod` node for the transaction ID to ensure it succeeded.

```go linenums="91"
// Submit the transaction
sendResponse, err := algodClient.SendRawTransaction(signedTxn).Do(context.Background())
if err != nil {
    fmt.Printf("failed to send transaction: %s\n", err)
    return
}
fmt.Printf("Submitted transaction %s\n", sendResponse)

// Wait for confirmation
confirmedTxn, err := future.WaitForConfirmation(algodClient, txID, 4)
if err != nil {
    fmt.Printf("Error waiting for confirmation on txID: %s\n", txID)
    return
}
```

- [Watch Video](https://youtu.be/rFG7Zo2JvIY?t=216)

# Display completed transaction

Finally, we can query the blockchain for the committed transaction data and display in on the command line. 

```go linenums="106"
// Display completed transaction
txnJSON, err := json.MarshalIndent(confirmedTxn.Transaction.Txn, "", "\t")
if err != nil {
    fmt.Printf("Can not marshall txn data: %s\n", err)
}
fmt.Printf("Transaction information: %s\n", txnJSON)
fmt.Printf("Decoded note: %s\n", string(confirmedTxn.Transaction.Txn.Note))
fmt.Printf("Amount sent: %d microAlgos\n", confirmedTxn.Transaction.Txn.Amount)
fmt.Printf("Fee: %d microAlgos\n", confirmedTxn.Transaction.Txn.Fee)	
		
fmt.Printf("Decoded note: %s\n", string(confirmedTxn.Transaction.Txn.Note))

```

- [Watch Video](https://youtu.be/rFG7Zo2JvIY?t=232)
 
 
# Run the program
 
Save your file and execute the program:

```bash
go run yourFirstTransaction.go
```

!!! Warning
    In order for your transaction to be successful, you must fund the generated account during runtime.

!!! Info
	View the confirmed transaction in your web browser by clicking the link to these third-party block explorers and inserting the transactionID within their search bar:
	
	- [AlgoExplorer](https://testnet.algoexplorer.io/)
	- [Goal Seeker](https://goalseeker.purestake.io/algorand/testnet)

- [Watch Video](https://youtu.be/rFG7Zo2JvIY?t=232)
 
# Complete example

If you have any trouble compiling or running your program, please check the complete example below which details how to quickly submit your first transaction.
 
[Run Code](https://replit.com/@Algorand/Getting-Started-with-Go)

[Watch Video](https://youtu.be/rFG7Zo2JvIY?t=)

# Setting up your editor/framework

The Algorand community provides many editors, frameworks, and plugins that can be used to work with the Algorand Network. Tutorials have been created for configuring each of these for use with Algorand. Select your Editor preference below.

* [Setting Up VSCode](https://developer.algorand.org/tutorials/vs-code-go/)
* [AlgoDEA IntelliJ Plugin](https://developer.algorand.org/articles/making-development-easier-algodea-intellij-plugin/)
* [Algo Builder Framework](https://developer.algorand.org/articles/introducing-algorand-builder/)
