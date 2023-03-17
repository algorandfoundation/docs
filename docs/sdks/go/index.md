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

# Install Go SDK

Algorand provides an SDK for Go. 

!!! Prerequisites
    - Go programming language ([install guide](https://golang.org/doc/install))

From a terminal window, install the Go SDK:

```bash
go get -u github.com/algorand/go-algorand-sdk/v2
```

- [SDK repository](https://github.com/algorand/go-algorand-sdk)
 
The SDK is installed and can now interact with the running Algorand Sandbox environment, as configured above.

# Create account

In order to interact with the Algorand blockchain, you must have a funded account on the network. To quickly create an account on Algorand TestNet create a new file **yourFirstTransaction.go** and insert the following code:

<!-- ===GOSDK_ACCOUNT_GENERATE=== -->
```go
account := crypto.GenerateAccount()
mn, err := mnemonic.FromPrivateKey(account.PrivateKey)

if err != nil {
	log.Fatalf("failed to generate account: %s", err)
}

log.Printf("Address: %s\n", account.Address)
log.Printf("Mnemonic: %s\n", mn)
```
[Snippet Source](https://github.com/barnjamin/go-algorand-sdk/blob/examples/_examples/account.go#L15-L24)
<!-- ===GOSDK_ACCOUNT_GENERATE=== -->

!!! Note 
    Lines 17 and 35 contain TODO: comments about inserting additional code. As you proceed with this guide, ensure the line numbers remain in sync.

!!! Tip
    Make sure to save the generated address and passphrase in a secure location, as they will be used later on.

!!! Warning 
    Never share your mnemonic passphrase or private keys. Production environments require stringent private key management. For more information on key management in community Wallets, click [here](https://developer.algorand.org/docs/community/#wallets). For the open source [Algorand Wallet](https://developer.algorand.org/articles/algorand-wallet-now-open-source/), click [here](https://github.com/algorand/algorand-wallet).

- [More Information](https://developer.algorand.org/docs/features/accounts/create/#standalone)
 
# Fund account

The code below prompts to fund the newly generated account. Before sending transactions to the Algorand network, the account must be funded to cover the minimal transaction fees that exist on Algorand. To fund the account use the [Algorand TestNet faucet](https://dispenser.testnet.aws.algodev.network/). 

!!! Info
    All Algorand accounts require a minimum balance to be registered in the ledger. To read more about Algorand minimum balance see [Account Overview](https://developer.algorand.org/docs/features/accounts/#minimum-balance)


# Connect Your client

Client must be instantiated prior to making calls to the API endpoints. You must provide values for `<algod-address>` and `<algod-token>`. The CLI tools implement the client natively. By default, the `algodToken` for each [sandbox](https://github.com/algorand/sandbox) is set to its `aaa...` value and the `algodAddress` corresponds to `http://localhost:4001`.


<!-- ===GOSDK_ALGOD_CREATE_CLIENT=== -->
```go
// Create a new algod client, configured to connect to out local sandbox
var algodAddress = "http://localhost:4001"
var algodToken = strings.Repeat("a", 64)
algodClient, _ := algod.MakeClient(
	algodAddress,
	algodToken,
)

// Or, if necessary, pass alternate headers

var algodHeader common.Header
algodHeader.Key = "X-API-Key"
algodHeader.Value = algodToken
algodClientWithHeaders, _ := algod.MakeClientWithHeaders(
	algodAddress,
	algodToken,
	[]*common.Header{&algodHeader},
)
```
[Snippet Source](https://github.com/barnjamin/go-algorand-sdk/blob/examples/_examples/overview.go#L47-L65)
<!-- ===GOSDK_ALGOD_CREATE_CLIENT=== -->
 
!!! Info
    This guide provides values for `algodAddress` and `algodToken` as specified by Algorand Sandbox. If you want to connect to a third-party service provider, see [Purestake](https://developer.purestake.io/code-samples) or [AlgoExplorer Developer API](https://algoexplorer.io/api-dev/v2) and adjust these values accordingly.
 

# Check account balance

Before moving on to the next step, make sure your account has been funded.
 
 <!-- ===GOSDK_ALGOD_FETCH_ACCOUNT_INFO=== -->
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
 <!-- ===GOSDK_ALGOD_FETCH_ACCOUNT_INFO=== -->


# Build transaction

Communication with the Algorand network is performed using transactions. Create a payment transaction sending 1 ALGO from your account to the TestNet faucet address:

<!-- ===GOSDK_TRANSACTION_PAYMENT_CREATE=== -->

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
<!-- ===GOSDK_TRANSACTION_PAYMENT_CREATE=== -->

!!! Info
    Algorand supports many transaction types. To see what types are supported see [Transactions](https://developer.algorand.org/docs/features/transactions/).


# Sign transaction

Before the transaction is considered valid, it must be signed by a private key. Use the following code to sign the transaction.

<!-- ===GOSDK_TRANSACTION_PAYMENT_SIGN=== -->
```go linenums="83"
// Sign the transaction
txID, signedTxn, err := crypto.SignTransaction(account.PrivateKey, txn)
if err != nil {
    fmt.Printf("Failed to sign transaction: %s\n", err)
    return
}
fmt.Printf("Signed txid: %s\n", txID)
```
<!-- ===GOSDK_TRANSACTION_PAYMENT_SIGN=== -->

!!! Info
    Algorand provides many ways to sign transactions. To see other ways see [Authorization](https://developer.algorand.org/docs/features/transactions/signatures/#single-signatures).


# Submit transaction

The signed transaction can now be broadcast to the network for validation and inclusion in a future block. The `waitForConfirmation` SDK method polls the `algod` node for the transaction ID to ensure it succeeded.

<!-- ===GOSDK_TRANSACTION_PAYMENT_SUBMIT=== -->
```go linenums="91"
// Submit the transaction
sendResponse, err := algodClient.SendRawTransaction(signedTxn).Do(context.Background())
if err != nil {
    fmt.Printf("failed to send transaction: %s\n", err)
    return
}
fmt.Printf("Submitted transaction %s\n", sendResponse)

// Wait for confirmation
confirmedTxn, err := transaction.WaitForConfirmation(algodClient, txID, 4, context.Background())
if err != nil {
    fmt.Printf("Error waiting for confirmation on txID: %s\n", txID)
    return
}
```
<!-- ===GOSDK_TRANSACTION_PAYMENT_SUBMIT=== -->
 
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

 
# Complete example

If you have any trouble compiling or running your program, please check the complete example below which details how to quickly submit your first transaction.

# Setting up your editor/framework

The Algorand community provides many editors, frameworks, and plugins that can be used to work with the Algorand Network. Tutorials have been created for configuring each of these for use with Algorand. Select your Editor preference below.

* [Setting Up VSCode](https://developer.algorand.org/tutorials/vs-code-go/)