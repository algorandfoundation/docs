title: Java SDK: Your First Transaction

This section is a quick start guide for interacting with Algorand network using Java. This guide will help to install [Algorand sandbox](https://github.com/algorand/sandbox){target=blank}, which provides a node for testing and development. This guide will also help to install the Java SDK, create an account and submit your first transaction.

# Install Sandbox

!!! info
    This step is only required if you are not using AlgoKit. If you are using AlgoKit, you can spin up a sandbox using the LocalNet, see [AlgoKit getting started guide](/docs/get-started/algokit/#start-a-localnet) for more information. 

!!! Prerequisites
    - Docker Compose ([install guide](https://docs.docker.com/compose/install/){target=blank})
    - Git ([install guide](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git){target=blank})

Algorand provides a docker instance for setting up a node, which can be used to get started developing. To install and use this instance, follow these instructions.

```bash
git clone https://github.com/algorand/sandbox.git
cd sandbox
./sandbox up dev 
```

This will install and start private network. To read more about Algorand networks see [Algorand Networks](../../get-details/algorand-networks/index.md){target=_blank}. 

[More Information about the sandbox](https://developer.algorand.org/articles/introducing-sandbox-20/) and [how to use](https://developer.algorand.org/tutorials/exploring-the-algorand-sandbox/) it.


# Install Java SDK 
Algorand provides an [SDK for Java](https://github.com/algorand/java-algorand-sdk). The instructions for installing the SDK are as follows. The Java SDK is available in the MVN repository and can be used in your Maven project by including the following dependency.

!!! Prerequisites
    Java SDK requires Java 7+ and Android minSdkVersion 16+. Check for the latest version of the Java SDK [here](https://github.com/algorand/java-algorand-sdk#installation){target=blank}.

```java
<dependency>
    <groupId>com.algorand</groupId>
    <artifactId>algosdk</artifactId>
    <version>2.0.0</version>
</dependency>
```

The [GitHub repository](https://github.com/algorand/js-algorand-sdk){target=_blank} contains additional documentation and examples.

See the Java SDK [reference documentation](https://algorand.github.io/java-algorand-sdk/){target=blank} for more information on packages and methods.

The SDK is installed and can now interact with the Sandbox created earlier.

# Create an Account
In order to interact with the Algorand blockchain, you must have a funded account. To quickly create a test account use the following code.

<!-- ===JAVASDK_ACCOUNT_GENERATE=== -->
```java
Account acct = new Account();
System.out.println("Address: " + acct.getAddress());
System.out.println("Passphrase: " + acct.toMnemonic());
```
[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/Overview.java#L76-L79)
<!-- ===JAVASDK_ACCOUNT_GENERATE=== -->

[More Information](../../get-details/accounts/create/#standalone){target=blank}

!!! Warning
    Never share Mnemonic private keys. Production environments require stringent private key management. For more information on key management in community Wallets, click [here](../../../../ecosystem-projects/#wallets){target=blank}. For the [Algorand open source wallet](https://developer.algorand.org/articles/algorand-wallet-now-open-source/){target=blank}, click [here](https://github.com/algorand/algorand-wallet){target=blank}.

# Fund the Account
Before sending transactions to the Algorand network, the account must be funded to cover the minimal transaction fees that exist on Algorand. In this example, we'll be using prefunded accounts available in the Sandbox. To fund an account on Testnet account use the [Algorand faucet](https://dispenser.testnet.aws.algodev.network/){target=_blank}. 


!!! Info
    All Algorand accounts require a minimum balance to be registered in the ledger. To read more about Algorand minimums see this [link](../../get-details/accounts/#minimum-balance){target=blank}.


# Connect Your Client
An Algod client must be instantiated prior to making calls to the API endpoints. You must provide values for `<algod-address>` and `<algod-token>`. The CLI tools implement the client natively. By default, the `algodToken` for each [sandbox](https://github.com/algorand/sandbox) is set to its `aaa...` value and the `algodAddress` corresponds to `http://localhost:4001`.


<!-- ===JAVASDK_ALGOD_CREATE_CLIENT=== -->
```java
String algodHost = "http://localhost";
int algodPort = 4001;
String algodToken = "a".repeat(64);
AlgodClient algodClient = new AlgodClient(algodHost, algodPort, algodToken);

// OR if the API provider requires a specific header key for the token
String tokenHeader = "X-API-Key";
AlgodClient otherAlgodClient = new AlgodClient(algodHost, algodPort, algodToken, tokenHeader);
```
[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/Overview.java#L94-L102)
<!-- ===JAVASDK_ALGOD_CREATE_CLIENT=== -->

!!! Info
    The example code connects to the sandbox Algod client. If you want to connect to a public API client, change the host, port, and token parameters to match the API service. See some service available [here](https://developer.algorand.org/ecosystem-projects/?tags=api-services)

!!! Info
    If you are connecting to the Testnet, a dispenser is available [here](https://dispenser.testnet.aws.algodev.network/){target=_blank}

# Check Your Balance
Before moving on to the next step, make sure your account has been funded by the faucet.

<!-- ===JAVASDK_ALGOD_FETCH_ACCOUNT_INFO=== -->
```java
Response<com.algorand.algosdk.v2.client.model.Account> acctInfoResp = algodClient
        .AccountInformation(acct.getAddress()).execute();
com.algorand.algosdk.v2.client.model.Account acctInfo = acctInfoResp.body();
// print one of the fields in the account info response
System.out.printf("Current balance: %d", acctInfo.amount);
```
[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/Overview.java#L84-L89)
<!-- ===JAVASDK_ALGOD_FETCH_ACCOUNT_INFO=== -->


# Build First Transaction
Transactions are used to interact with the Algorand network. To create a payment transaction use the following code.
​
<!-- ===JAVASDK_TRANSACTION_PAYMENT_CREATE=== -->
```java
Response<TransactionParametersResponse> suggestedParams = algodClient.TransactionParams().execute();
Integer amount = 1000000; // 1 Algo
Transaction ptxn = Transaction.PaymentTransactionBuilder()
        .sender(acct.getAddress())
        .amount(amount)
        .receiver(acct2.getAddress())
        .suggestedParams(suggestedParams.body()).build();
```
[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/Overview.java#L25-L32)
<!-- ===JAVASDK_TRANSACTION_PAYMENT_CREATE=== -->
​
!!! Info
    Algorand supports many transaction types. To see what types are supported see [Transactions](../../get-details/transactions/index.md#transaction-types){target=_blank}. 

# Sign First Transaction
Before the transaction is considered valid, it must be signed by a private key. Use the following code to sign the transaction.

<!-- ===JAVASDK_TRANSACTION_PAYMENT_SIGN=== -->
```java
SignedTransaction sptxn = acct.signTransaction(ptxn);
```
[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/Overview.java#L35-L36)
<!-- ===JAVASDK_TRANSACTION_PAYMENT_SIGN=== -->

!!! Info
    Algorand provides additional ways for transactions to be signed, other than by a standalone account. For more information see [Authorization](../../get-details/transactions/signatures){target=blank}.

# Submit the Transaction
The signed transaction can now be submitted to the network. The SDK `waitForConfirmation` utility function is called after the transaction is submitted to wait until the transaction is broadcast to the Algorand blockchain and is confirmed. 

<!-- ===JAVASDK_TRANSACTION_PAYMENT_SUBMIT=== -->
```java
// encode the transaction
byte[] encodedTxBytes = Encoder.encodeToMsgPack(sptxn);
// submit the transaction to the algod server
Response<PostTransactionsResponse> resp = algodClient.RawTransaction().rawtxn(encodedTxBytes).execute();
// wait for the transaction to be confirmed
String txid = resp.body().txId;
PendingTransactionResponse result = Utils.waitForConfirmation(algodClient, txid, 4);
System.out.printf("Transaction %s confirmed in round %d\n", txid, result.confirmedRound);
```
[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/Overview.java#L39-L47)
<!-- ===JAVASDK_TRANSACTION_PAYMENT_SUBMIT=== -->

# Viewing the Transaction

To view the transaction we submitted to the sandbox Algod, open [Lora](https://lora.algokit.io/localnet){target=_blank} and choose `LocalNet` configuration option, then search for the transaction ID. 

To view a transaction submitted to public network like testnet, open [Lora](https://lora.algokit.io/testnet){target=_blank} or [Pera Explorer](https://testnet.explorer.perawallet.app/){target=blank} and paste the transaction ID into the search bar.
