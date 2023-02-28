title: Your First Transaction

This section is a quick start guide for interacting with Algorand network using Java. This guide will help to install **_sandbox_**, which provides a node for testing and development. This guide will also help to install the Java SDK, create an account and submit your first transaction.

# Sandbox Install

!!! Prerequisites
    - Docker Compose ([install guide](https://docs.docker.com/compose/install/){:target="_blank"})
    - Git ([install guide](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git){:target="_blank"})

Algorand provides a docker instance for setting up a node, which can be used to get started developing. To install and use this instance, follow these instructions.

```bash
git clone https://github.com/algorand/sandbox.git
cd sandbox
./sandbox up testnet
```


[More Information](https://developer.algorand.org/articles/introducing-sandbox-20/){:target="_blank"}

This will install a Sandbox node connected to the Algorand TestNet. To read more about Algorand networks see [Algorand Networks](../../get-details/algorand-networks/){:target="_blank"}.

!!! Warning
    The sandbox installation may take a few minutes to startup in order to catch up to the current block round. To learn more about fast catchup, see [Sync Node Network using Fast Catchup](../../run-a-node/setup/install/#sync-node-network-using-fast-catchup){:target="_blank"}
.

# Install SDK For Runtime

Algorand provides an SDK for Java. The instructions for installing the SDK are as follows. The Java SDK is available in the MVN repository and can be used in your Maven project by including the following dependency.

Requirements: Java SDK requires Java 7+ and Android minSdkVersion 16+. Check for the latest version of the Java SDK [here](https://github.com/algorand/java-algorand-sdk#installation){:target="_blank"}.

```java
<dependency>
    <groupId>com.algorand</groupId>
    <artifactId>algosdk</artifactId>
    <version>2.0.0</version>
</dependency>
```


[`More Information`](https://github.com/algorand/java-algorand-sdk#installation){:target="_blank"}

The GitHub repository contains additional documentation and examples.

See the [Java SDK reference documentation](https://algorand.github.io/java-algorand-sdk/){:target="_blank"} for more information on packages and methods.

The SDK is installed and can now interact with the Sandbox created earlier.

!!! Info
    Web based solutions require the AlgoSigner Chrome plugin or other web-based private key management software. For more information see [community wallets](../../../../ecosystem-projects/#wallets){:target="_blank"}.

# Create an Account on Algorand

In order to interact with the Algorand blockchain, you must have a funded account. To quickly create a test account use the following code.

<!-- ===JAVASDK_ACCOUNT_GENERATE=== -->
```java
        Account acct = new Account();
        System.out.println("Address: " + acct.getAddress());
        System.out.println("Passphrase: " + acct.toMnemonic());
```
<!-- ===JAVASDK_ACCOUNT_GENERATE=== -->

[More Information](../../get-details/accounts/create/#standalone){:target="_blank"}

!!! Warning
    Never share Mnemonic private keys. Production environments require stringent private key management. For more information on key management in community Wallets, click [here](../../../../ecosystem-projects/#wallets){:target="_blank"}. For the [Algorand open source wallet](https://developer.algorand.org/articles/algorand-wallet-now-open-source/){:target="_blank"}, click [here](https://github.com/algorand/algorand-wallet){:target="_blank"}.

# Fund the Account


The code above prompts to fund the newly created account. Before sending transactions to the Algorand network, the account must be funded to cover the minimal transaction fees that exist on Algorand. To fund the account use the [Algorand TestNet faucet](https://dispenser.testnet.aws.algodev.network/){:target="_blank"}.


!!! Info
    All Algorand accounts require a minimum balance to be registered in the ledger. To read more about Algorand minimums see this [link](../../get-details/accounts/#minimum-balance){:target="_blank"}.


# Connect Your Client
Client must be instantiated prior to making calls to the API endpoints. You must provide values for `<algod-address>` and `<algod-token>`. The CLI tools implement the client natively. By default, the `algodToken` for each [sandbox](https://github.com/algorand/sandbox) is set to its `aaa...` value and the `algodHost` corresponds to `http://localhost:4001`.


# Build First Transaction

Communication with the Algorand network is performed using transactions. To create a payment transaction use the following code, which also includes some utility functions, `connectToNetwork` ,  and `PrintBalance`. Add this code to the GettingStarted class above.

<!-- ===JAVASDK_ALGOD_CREATE_CLIENT=== -->
```java
        String algodHost = "http://localhost";
        int algodPort = 4001;
        String algodToken = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
        AlgodClient algodClient = new AlgodClient(algodHost, algodPort, algodToken);

        // OR if the API provider requires a specific header key for the token
        String tokenHeader = "X-API-Key";
        AlgodClient otherAlgodClient = new AlgodClient(algodHost, algodPort, algodToken, tokenHeader);
```
<!-- ===JAVASDK_ALGOD_CREATE_CLIENT=== -->


!!! Info
    Algorand supports many transaction types. To see what types are supported see [Transactions](../../get-details/transactions/){: target="_blank"}.

# Check Your Balance
Before moving on to the next step, make sure your account has been funded by the faucet.

<!-- ===JAVASDK_ALGOD_FETCH_ACCOUNT_INFO=== -->
```python
account_info: Dict[str, Any] = algod_client.account_info(address)
print(f"Account balance: {account_info.get('amount')} microAlgos")
```
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
<!-- ===JAVASDK_TRANSACTION_PAYMENT_SIGN=== -->

!!! Info
    Algorand provides additional ways for transactions to be signed, other than by a standalone account. For more information see [Authorization](../../get-details/transactions/signatures){:target="_blank"}.

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
<!-- ===JAVASDK_TRANSACTION_PAYMENT_SUBMIT=== -->

[`Run Code`](https://replit.com/@Algorand/GettingStarted-with-Java#GettingStarted.java){:target="_blank"}


# Complete Example
The complete example below illustrates how to quickly submit your first transaction.
​
# TODO::
[...](https://github.com/algorand/java-algorand-sdk/blob/master/examples/src/main/java/com/algorand/examples/Example.java)

!!! Warning 
    In order for this transaction to be successful, the account must be funded. 

# Viewing the Transaction

To view the transaction, open the [Algorand Blockchain Explorer](https://testnet.algoexplorer.io/){:target="_blank"} or [Goal Seeker](https://goalseeker.purestake.io/algorand/testnet){:target="_blank"} and paste the transaction ID into the search bar or simply click on the funded transaction link on the dispenser page.

# Setting Up Your Editor/Framework

The Algorand community provides many editors, frameworks, and plugins that can be used to work with the Algorand Network. Tutorials have been created for configuring each of these for use with Algorand. Select your Editor preference below.

- [Setting Up VSCode](https://developer.algorand.org/tutorials/vs-code-java/)
- [AlgoDEA IntelliJ Plugin](https://developer.algorand.org/articles/making-development-easier-algodea-intellij-plugin/)
- [Algorand Builder Framework](https://developer.algorand.org/articles/introducing-algorand-builder/)
  