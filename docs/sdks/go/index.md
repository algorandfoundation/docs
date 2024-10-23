title: Go SDK: Your First Transaction

This section is a quick start guide for interacting with the Algorand network using Go. This guide will help to install [Algorand sandbox](https://github.com/algorand/sandbox){target=blank}, which provides a node for testing and development. This guide will also help to install the Go SDK, create an account and submit your first transaction on Algorand.  
 
# Install Sandbox

!!! info
    This step is only required if you are not using AlgoKit. If you are using AlgoKit, you can spin up a sandbox using the LocalNet, see [AlgoKit getting started guide](/docs/get-started/algokit/#start-a-localnet) for more information. 
	
!!! Prerequisites
    - Docker Compose ([install guide](https://docs.docker.com/compose/install/))
    - Git ([install guide](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git))

Algorand provides a docker instance for setting up a node, which can be used to get started developing quickly. To install and use this instance, follow these instructions.

```bash
git clone https://github.com/algorand/sandbox.git
cd sandbox
./sandbox up dev 
```

This will install and start private network. To read more about Algorand networks see [Algorand Networks](../../get-details/algorand-networks/index.md){target=_blank}. 

[More Information about the sandbox](https://developer.algorand.org/articles/introducing-sandbox-20/) and [how to use](https://developer.algorand.org/tutorials/exploring-the-algorand-sandbox/) it.

 

# Install Go SDK
Algorand provides an SDK for Go. 

!!! Prerequisites
    - Go programming language ([install guide](https://golang.org/doc/install))

From a terminal window, install the Go SDK:

```bash
go get -u github.com/algorand/go-algorand-sdk/v2
```

The [GitHub repository](https://github.com/algorand/go-algorand-sdk){target=_blank} contains additional documentation and examples.

See the JavaScript SDK [reference documentation](https://pkg.go.dev/github.com/algorand/go-algorand-sdk/v2){target=_blank} for more information on methods.  

The SDK is installed and can now interact with the running Algorand Sandbox environment, as configured above.

# Create an account
In order to interact with the Algorand blockchain, you must have a funded account on the network. To quickly create a test account use the following code.

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
[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/account/main.go#L16-L25)
<!-- ===GOSDK_ACCOUNT_GENERATE=== -->

[`More Information`](../../get-details/accounts/create.md#standalone){target=_blank}  

!!! Warning 
    Never share your mnemonic passphrase or private keys. Production environments require stringent private key management. For more information on key management in community Wallets, click [here](https://developer.algorand.org/docs/community/#wallets). For the open source [Algorand Wallet](https://developer.algorand.org/articles/algorand-wallet-now-open-source/), click [here](https://github.com/algorand/algorand-wallet).

 
# Fund the account
Before sending transactions to the Algorand network, the account must be funded to cover the minimal transaction fees that exist on Algorand. In this example, we'll be using prefunded accounts available in the Sandbox. To fund an account on Testnet account use the [Algorand faucet](https://dispenser.testnet.aws.algodev.network/){target=_blank}. 

!!! Info
    All Algorand accounts require a minimum balance to be registered in the ledger. To read more about Algorand minimum balance see [Account Overview](https://developer.algorand.org/docs/features/accounts/#minimum-balance)


# Connect Your Client
An Algod client must be instantiated prior to making calls to the API endpoints. You must provide values for `<algod-address>` and `<algod-token>`. The CLI tools implement the client natively. By default, the `algodToken` for each [sandbox](https://github.com/algorand/sandbox) is set to its `aaa...` value and the `algodAddress` corresponds to `http://localhost:4001`.


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
[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/overview/main.go#L18-L36)
<!-- ===GOSDK_ALGOD_CREATE_CLIENT=== -->
 
!!! Info
    The example code connects to the sandbox Algod client. If you want to connect to a public API client, change the host, port, and token parameters to match the API service. See some service available [here](https://developer.algorand.org/ecosystem-projects/?tags=api-services)

!!! Info
    If you are connecting to the Testnet, a dispenser is available [here](https://dispenser.testnet.aws.algodev.network/){target=_blank}

# Check Your Balance
Before moving on to the next step, make sure your account has been funded.
 
 <!-- ===GOSDK_ALGOD_FETCH_ACCOUNT_INFO=== -->
```go
acctInfo, err := algodClient.AccountInformation(acct.Address.String()).Do(context.Background())
if err != nil {
	log.Fatalf("failed to fetch account info: %s", err)
}
log.Printf("Account balance: %d microAlgos", acctInfo.Amount)
```
[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/overview/main.go#L51-L56)
 <!-- ===GOSDK_ALGOD_FETCH_ACCOUNT_INFO=== -->


# Build First Transaction
Transactions are used to interact with the Algorand network. To create a payment transaction use the following code.

<!-- ===GOSDK_TRANSACTION_PAYMENT_CREATE=== -->
```go
sp, err := algodClient.SuggestedParams().Do(context.Background())
if err != nil {
	log.Fatalf("failed to get suggested params: %s", err)
}
// payment from account to itself
ptxn, err := transaction.MakePaymentTxn(acct.Address.String(), acct.Address.String(), 100000, nil, "", sp)
if err != nil {
	log.Fatalf("failed creating transaction: %s", err)
}
```
[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/overview/main.go#L59-L68)
<!-- ===GOSDK_TRANSACTION_PAYMENT_CREATE=== -->

!!! Info
    Algorand supports many transaction types. To see what types are supported see [Transactions](https://developer.algorand.org/docs/features/transactions/).


# Sign First transaction
Before the transaction is considered valid, it must be signed by a private key. Use the following code to sign the transaction.

<!-- ===GOSDK_TRANSACTION_PAYMENT_SIGN=== -->
```go
_, sptxn, err := crypto.SignTransaction(acct.PrivateKey, ptxn)
if err != nil {
	fmt.Printf("Failed to sign transaction: %s\n", err)
	return
}
```
[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/overview/main.go#L71-L76)
<!-- ===GOSDK_TRANSACTION_PAYMENT_SIGN=== -->

!!! Info
    Algorand provides many ways to sign transactions. To see other ways see [Authorization](https://developer.algorand.org/docs/features/transactions/signatures/#single-signatures).


# Submit transaction
The signed transaction can now be broadcast to the network for validation and inclusion in a future block. The `waitForConfirmation` SDK method polls the `algod` node for the transaction ID to ensure it succeeded.

<!-- ===GOSDK_TRANSACTION_PAYMENT_SUBMIT=== -->
```go
pendingTxID, err := algodClient.SendRawTransaction(sptxn).Do(context.Background())
if err != nil {
	fmt.Printf("failed to send transaction: %s\n", err)
	return
}
confirmedTxn, err := transaction.WaitForConfirmation(algodClient, pendingTxID, 4, context.Background())
if err != nil {
	fmt.Printf("Error waiting for confirmation on txID: %s\n", pendingTxID)
	return
}
fmt.Printf("Confirmed Transaction: %s in Round %d\n", pendingTxID, confirmedTxn.ConfirmedRound)
```
[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/overview/main.go#L79-L90)
<!-- ===GOSDK_TRANSACTION_PAYMENT_SUBMIT=== -->
 
# Viewing the Transaction
To view the transaction we submitted to the sandbox Algod, open [Lora](https://lora.algokit.io/localnet){target=_blank} and choose `LocalNet` configuration option, then search for the transaction ID. 

To view a transaction submitted to public network like testnet, open [Lora](https://lora.algokit.io/testnet){target=_blank} or [Pera Explorer](https://testnet.explorer.perawallet.app/){target=blank} and paste the transaction ID into the search bar.
