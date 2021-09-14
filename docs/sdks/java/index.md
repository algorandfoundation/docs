
title: Your First Transaction

This section is a quick start guide for interacting with Algorand network using Java. This guide will help to install **_sandbox_**, which provides a node for testing and development. This guide will also help to install the Java SDK, create an account and submit your first transaction.

# Sandbox Install

!!! Prerequisites

- Docker Compose ([install guide](https://docs.docker.com/compose/install/))
- Git ([install guide](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git))

Algorand provides a docker instance for setting up a node, which can be used to get started developing. To install and use this instance, follow these instructions.

```bash
git clone https://github.com/algorand/sandbox.git
cd sandbox
./sandbox up testnet
```

[Watch Video](https://youtu.be/_EEMXHrbMzU?t=21)

[More Information](https://developer.algorand.org/articles/introducing-sandbox-20/)

This will install a Sandbox node connected to the Algorand TestNet. To read more about Algorand networks see [Algorand Networks](https://developer.algorand.org/docs/reference/algorand-networks/).


!!! Warning
The sandbox installation may take a few minutes to startup in order to catch up to the current block round. To learn more about fast catchup, see [Sync Node Network using Fast Catchup](https://developer.algorand.org/docs/run-a-node/setup/install/#sync-node-network-using-fast-catchup)
. 

# Install SDK For Runtime

Algorand provides an SDK for Java. The instructions for installing the SDK are as follows. The Java SDK is available in the MVN repository and can be used in your Maven project by including the following dependency.

Requirements: Java SDK requires Java 7+ and Android minSdkVersion 16+. Check for the latest version of the Java SDK [here](https://github.com/algorand/java-algorand-sdk#installation).

```java
<dependency>
    <groupId>com.algorand</groupId>
    <artifactId>algosdk</artifactId>
    <version>1.6.0</version>
</dependency>
```


[`Watch Video`](https://youtu.be/_EEMXHrbMzU?t=98)

[`More Information`](https://github.com/algorand/java-algorand-sdk#installation)

The GitHub repository contains additional documentation and examples.

See the [Java SDK reference documentation](https://algorand.github.io/java-algorand-sdk/) for more information on packages and methods.

The SDK is installed and can now interact with the Sandbox created earlier.

!!! Info
Web based solutions require the AlgoSigner Chrome plugin or other web-based private key management software. For more information see [community wallets](https://developer.algorand.org/docs/community/#wallets).

# Create an Account on Algorand

In order to interact with the Algorand blockchain, you must have a funded account. To quickly create a test account use the following code.

```java
import com.algorand.algosdk.account.Account;
class GettingStarted{

    // Create Account
    static Scanner scan = new Scanner(System.in);
    public Account createAccount()  throws Exception {
        try {
            Account myAccount1 = new Account();
            System.out.println("My Address: " + myAccount1.getAddress());
            System.out.println("My Passphrase: " + myAccount1.toMnemonic());
            System.out.println("Navigate to this link:  https://dispenser.testnet.aws.algodev.network/");
            System.out.println("Copy TestNet Account Address to Dispense Funds to: ");
            System.out.println(myAccount1.getAddress().toString());
            System.out.println("PRESS ENTER KEY TO CONTINUE...");
            scan.nextLine();
            return myAccount1;
            // Copy off account and mnemonic
            // Dispense TestNet Algos to account:
            // https://dispenser.testnet.aws.algodev.network/
            // resource:
            // https://developer.algorand.org/docs/features/accounts/create/#standalone

        } catch (Exception e) {
            e.printStackTrace();
            throw new Exception("Account creation error " + e.getMessage() );
        }

        public static void main(String args[]) throws Exception {
            GettingStarted t = new GettingStarted();
            Account myAccount1 = t.createAccount();
            t.gettingStartedExample(myAccount1); // to be added below
        }

    }
}

```

[Watch Video](https://youtu.be/_EEMXHrbMzU?t=139)

[More Information](https://developer.algorand.org/docs/features/accounts/create/#standalone)

!!! Warning
Never share Mnemonic private keys. Production environments require stringent private key management. For more information on key management in community Wallets, click [here](https://developer.algorand.org/docs/community/#wallets). For the [Algorand open source wallet](https://developer.algorand.org/articles/algorand-wallet-now-open-source/), click [here](https://github.com/algorand/algorand-wallet).

# Fund the Account

The code above prompts to fund the newly created account. Before sending transactions to the Algorand network, the account must be funded to cover the minimal transaction fees that exist on Algorand. To fund the account use the [Algorand TestNet faucet](https://dispenser.testnet.aws.algodev.network/). The 

!!! Info
All Algorand accounts require a minimum balance to be registered in the ledger. To read more about Algorand minimums see this [link](https://developer.algorand.org/docs/features/accounts/#minimum-balance). 

[Watch Video](https://youtu.be/_EEMXHrbMzU?t=178)

# Viewing the Transaction

To view the transaction, open the [Algorand Blockchain Explorer](https://testnet.algoexplorer.io/) or [Goal Seeker](https://goalseeker.purestake.io/algorand/testnet) and paste the transaction ID into the search bar or simply click on the funded transaction link on the dispenser page.

[Watch Video](https://youtu.be/_EEMXHrbMzU?t=209)

# Build First Transaction

Communication with the Algorand network is performed using transactions. To create a payment transaction use the following code, which also includes some utility functions, `connectToNetwork` , `waitForConfirmation` and `PrintBalance`. Add this code to the GettingStarted class above.

```java
package com.algorand.javatest.firsttransaction;
import com.algorand.algosdk.account.Account;
import com.algorand.algosdk.crypto.Address;
import com.algorand.algosdk.transaction.SignedTransaction;
import com.algorand.algosdk.transaction.Transaction;
import com.algorand.algosdk.util.Encoder;
import com.algorand.algosdk.v2.client.common.AlgodClient;
import com.algorand.algosdk.v2.client.common.Response;
import com.algorand.algosdk.v2.client.model.NodeStatusResponse;
import com.algorand.algosdk.v2.client.model.PendingTransactionResponse;
import com.algorand.algosdk.v2.client.model.PostTransactionsResponse;
import com.algorand.algosdk.v2.client.model.TransactionParametersResponse;
import org.json.JSONObject;

private AlgodClient client = null;
    // utility function to connect to a node
    private AlgodClient connectToNetwork() {
    final String ALGOD_API_ADDR = "localhost";
    final String ALGOD_API_TOKEN = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
    final Integer ALGOD_PORT = 4001;
    AlgodClient client = new AlgodClient(ALGOD_API_ADDR,
        ALGOD_PORT, ALGOD_API_TOKEN);
    return client;
}

/**
    * utility function to wait on a transaction to be confirmed
    * the timeout parameter indicates how many rounds do you wish to check pending transactions for
    */
private PendingTransactionResponse waitForConfirmation(AlgodClient myclient, String txID, Integer timeout)
throws Exception {
    if (myclient == null || txID == null || timeout < 0) {
        throw new IllegalArgumentException("Bad arguments for waitForConfirmation.");
    }
    Response < NodeStatusResponse > resp = myclient.GetStatus().execute();
    if (!resp.isSuccessful()) {
        throw new Exception(resp.message());
    }
    NodeStatusResponse nodeStatusResponse = resp.body();
    Long startRound = nodeStatusResponse.lastRound + 1;
    Long currentRound = startRound;
    while (currentRound < (startRound + timeout)) {
        // Check the pending transactions
        Response < PendingTransactionResponse > resp2 = myclient.PendingTransactionInformation(txID).execute();
        if (resp2.isSuccessful()) {
            PendingTransactionResponse pendingInfo = resp2.body();
            if (pendingInfo != null) {
                if (pendingInfo.confirmedRound != null && pendingInfo.confirmedRound > 0) {
                    // Got the completed Transaction
                    return pendingInfo;
                }
                if (pendingInfo.poolError != null && pendingInfo.poolError.length() > 0) {
                    // If there was a pool error, then the transaction has been rejected!
                    throw new Exception("The transaction has been rejected with a pool error: " + pendingInfo.poolError);
                }
            }
        }
        resp = myclient.WaitForBlock(currentRound).execute();
        if (!resp.isSuccessful()) {
            throw new Exception(resp.message());
        }
        currentRound++;
    }
    throw new Exception("Transaction not confirmed after " + timeout + " rounds!");
}

private String printBalance(com.algorand.algosdk.account.Account myAccount) throws Exception {
        String myAddress = myAccount.getAddress().toString();
        Response < com.algorand.algosdk.v2.client.model.Account > respAcct = client.AccountInformation(myAccount.getAddress()).execute();
        if (!respAcct.isSuccessful()) {
            throw new Exception(respAcct.message());
        }
        com.algorand.algosdk.v2.client.model.Account accountInfo = respAcct.body();
        System.out.println(String.format("Account Balance: %d microAlgos", accountInfo.amount));
        return myAddress;
}

public void gettingStartedExample(Account myAccount) throws Exception {
    if (client == null)
        this.client = connectToNetwork();
    printBalance(myAccount);
    try {
        // Construct the transaction
        final String RECEIVER = "HZ57J3K46JIJXILONBBZOHX6BKPXEM2VVXNRFSUED6DKFD5ZD24PMJ3MVA";
        String note = "Hello World";
        Response < TransactionParametersResponse > resp = client.TransactionParams().execute();
        if (!resp.isSuccessful()) {
            throw new Exception(resp.message());
        }
        TransactionParametersResponse params = resp.body();
        if (params == null) {
            throw new Exception("Params retrieval error");
        }
        JSONObject jsonObj = new JSONObject(params.toString());
        System.out.println("Algorand suggested parameters: " + jsonObj.toString(2));
        Transaction txn = Transaction.PaymentTransactionBuilder()
            .sender(myAccount.getAddress().toString())
            .note(note.getBytes())
            .amount(1000000) // 1 algo = 1000000 microalgos
            .receiver(new Address(RECEIVER))
            .suggestedParams(params)
            .closeRemainderTo(RECEIVER) // WARNING! all remaining funds in the sender account will be sent to the closeRemainderTo Account, omit RECEIVER account when in use otherwise all funds from the sender account will be sent to that account.
            .build();
        // CloseRemainder can be used to reset sender account to 0.
        // Normally this would be omitted. For more info see:
        // https://developer.algorand.org/docs/reference/transactions/#payment-transaction

        // more code below
```

[`Watch Video`](https://youtu.be/_EEMXHrbMzU?t=288)

!!! Info
Algorand supports many transaction types. To see what types are supported see [Transactions](https://developer.algorand.org/docs/features/transactions/).

# Sign First Transaction

Before the transaction is considered valid, it must be signed by a private key. Use the following code to sign the transaction.

```java
        // Sign the transaction
        SignedTransaction signedTxn = myAccount.signTransaction(txn);
        System.out.println("Signed transaction with txid: " + signedTxn.transactionID);
```

[`Watch Video`](https://youtu.be/_EEMXHrbMzU?t=401)

!!! Info
Algorand provides additional ways for transactions to be signed, other than by a standalone account. For more information see [Authorization](https://developer.algorand.org/docs/features/transactions/signatures).

# Submit the Transaction

The signed transaction can now be submitted to the network. `waitForConfirmation` is called after the transaction is submitted to wait until the transaction is broadcast to the Algorand blockchain and is confirmed. For more information, see [Wait for Confirmation](https://developer.algorand.org/docs/build-apps/hello_world/#wait-for-confirmation)

```java
        // Submit the transaction to the network
        byte[] encodedTxBytes = Encoder.encodeToMsgPack(signedTxn);
        Response < PostTransactionsResponse > rawtxresponse = client.RawTransaction().rawtxn(encodedTxBytes).execute();

        if (!rawtxresponse.isSuccessful()) {
            throw new Exception(rawtxresponse.message());
        }
        String id = rawtxresponse.body().txId;

        // Wait for transaction confirmation
        PendingTransactionResponse pTrx = waitForConfirmation(client, id, 4);

        System.out.println("Transaction " + id + " confirmed in round " + pTrx.confirmedRound);
        // Read the transaction
        JSONObject jsonObj2 = new JSONObject(pTrx.toString());
        System.out.println("Transaction information (with notes): " + jsonObj2.toString(2));
        System.out.println("Decoded note: " + new String(pTrx.txn.tx.note));
        System.out.println("Transaction information (with notes): " + jsonObj2.toString(2));
        System.out.println("Decoded note: " + new String(pTrx.txn.tx.note));
        System.out.println("Amount: " + new String(pTrx.txn.tx.amount.toString()));
        System.out.println("Fee: " + new String(pTrx.txn.tx.fee.toString()));
        if (pTrx.closingAmount != null){
            System.out.println("Closing Amount: " + new String(pTrx.closingAmount.toString()));
        }
        printBalance(myAccount);
    } catch (Exception e) {
        System.err.println("Exception when calling algod#transactionInformation: " + e.getMessage());
    }
}
```

[`Run Code`](https://replit.com/@Algorand/GettingStarted-with-Java#GettingStarted.java)

[`Watch Video`](https://youtu.be/_EEMXHrbMzU?t=410)

# Complete Example

The example below illustrates creating an account and using it to submit your first transaction.

```java
package com.algorand.javatest.firsttransaction;
import com.algorand.algosdk.account.Account;
import java.util.Scanner;
import com.algorand.algosdk.crypto.Address;
import com.algorand.algosdk.transaction.SignedTransaction;
import com.algorand.algosdk.transaction.Transaction;
import com.algorand.algosdk.util.Encoder;
import com.algorand.algosdk.v2.client.common.AlgodClient;
import com.algorand.algosdk.v2.client.common.Response;
import com.algorand.algosdk.v2.client.model.NodeStatusResponse;
import com.algorand.algosdk.v2.client.model.PendingTransactionResponse;
import com.algorand.algosdk.v2.client.model.PostTransactionsResponse;
import com.algorand.algosdk.v2.client.model.TransactionParametersResponse;
import org.json.JSONObject;

class GettingStarted{
        // Crete Account
    static Scanner scan = new Scanner(System.in);

    public Account createAccount()  throws Exception {
        try {
            Account myAccount1 = new Account();
            System.out.println("My Address: " + myAccount1.getAddress());
            System.out.println("My Passphrase: " + myAccount1.toMnemonic());

            System.out.println("Navigate to this link:  https://dispenser.testnet.aws.algodev.network/);
            System.out.println("Copy TestNet Account Address to Dispense Funds to: ");
            System.out.println(myAccount1.getAddress().toString());
            System.out.println("PRESS ENTER KEY TO CONTINUE...");
            scan.nextLine();
            return myAccount1;
            // Copy off account and mnemonic
            // Dispense TestNet Algos to account:
            // https://dispenser.testnet.aws.algodev.network/
            // resource:
            // https://developer.algorand.org/docs/features/accounts/create/#standalone
        } catch (Exception e) {
            e.printStackTrace();
            throw new Exception("Account creation error " + e.getMessage() );
        }
    }

    private AlgodClient client = null;
      // utility function to connect to a node
      private AlgodClient connectToNetwork() {
    final String ALGOD_API_ADDR = "localhost";
    final String ALGOD_API_TOKEN = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
    final Integer ALGOD_PORT = 4001;
      AlgodClient client = new AlgodClient(ALGOD_API_ADDR,
          ALGOD_PORT, ALGOD_API_TOKEN);
      return client;
    }
    /**
      * utility function to wait on a transaction to be confirmed
      * the timeout parameter indicates how many rounds do you wish to check pending transactions for
      */
    private PendingTransactionResponse waitForConfirmation(AlgodClient myclient, String txID, Integer timeout)
    throws Exception {
        if (myclient == null || txID == null || timeout < 0) {
            throw new IllegalArgumentException("Bad arguments for waitForConfirmation.");
        }
        Response < NodeStatusResponse > resp = myclient.GetStatus().execute();
        if (!resp.isSuccessful()) {
            throw new Exception(resp.message());
        }
        NodeStatusResponse nodeStatusResponse = resp.body();
        Long startRound = nodeStatusResponse.lastRound + 1;
        Long currentRound = startRound;
        while (currentRound < (startRound + timeout)) {
            // Check the pending transactions
            Response < PendingTransactionResponse > resp2 = myclient.PendingTransactionInformation(txID).execute();
            if (resp2.isSuccessful()) {
                PendingTransactionResponse pendingInfo = resp2.body();
                if (pendingInfo != null) {
                    if (pendingInfo.confirmedRound != null && pendingInfo.confirmedRound > 0) {
                        // Got the completed Transaction
                        return pendingInfo;
                    }
                    if (pendingInfo.poolError != null && pendingInfo.poolError.length() > 0) {
                        // If there was a pool error, then the transaction has been rejected!
                        throw new Exception("The transaction has been rejected with a pool error: " + pendingInfo.poolError);
                    }
                }
            }
            resp = myclient.WaitForBlock(currentRound).execute();
            if (!resp.isSuccessful()) {
                throw new Exception(resp.message());
            }
            currentRound++;
        }
        throw new Exception("Transaction not confirmed after " + timeout + " rounds!");
    }

    private String printBalance(com.algorand.algosdk.account.Account myAccount) throws Exception {
        String myAddress = myAccount.getAddress().toString();
        Response < com.algorand.algosdk.v2.client.model.Account > respAcct = client.AccountInformation(myAccount.getAddress()).execute();
        if (!respAcct.isSuccessful()) {
            throw new Exception(respAcct.message());
        }
        com.algorand.algosdk.v2.client.model.Account accountInfo = respAcct.body();
        System.out.println(String.format("Account Balance: %d microAlgos", accountInfo.amount));
        return myAddress;
    }

    public void gettingStartedExample(Account myAccount) throws Exception {

        if (client == null)
            this.client = connectToNetwork();

        printBalance(myAccount);

        try {
            // Construct the transaction
            final String RECEIVER = "HZ57J3K46JIJXILONBBZOHX6BKPXEM2VVXNRFSUED6DKFD5ZD24PMJ3MVA";
            String note = "Hello World";
            Response < TransactionParametersResponse > resp = client.TransactionParams().execute();
            if (!resp.isSuccessful()) {
                throw new Exception(resp.message());
            }
            TransactionParametersResponse params = resp.body();
            if (params == null) {
                throw new Exception("Params retrieval error");
            }
            JSONObject jsonObj = new JSONObject(params.toString());
            System.out.println("Algorand suggested parameters: " + jsonObj.toString(2));
            Transaction txn = Transaction.PaymentTransactionBuilder()
                .sender(myAccount.getAddress().toString())
                .note(note.getBytes())
                .amount(1000000) // 1 algo = 1000000 microalgos
                .receiver(new Address(RECEIVER))
                .suggestedParams(params)
                .build();

            // Sign the transaction
            SignedTransaction signedTxn = myAccount.signTransaction(txn);
            System.out.println("Signed transaction with txid: " + signedTxn.transactionID);
            // Submit the transaction to the network
            byte[] encodedTxBytes = Encoder.encodeToMsgPack(signedTxn);
            Response < PostTransactionsResponse > rawtxresponse = client.RawTransaction().rawtxn(encodedTxBytes).execute();

            if (!rawtxresponse.isSuccessful()) {
                throw new Exception(rawtxresponse.message());
            }
            String id = rawtxresponse.body().txId;
            // Wait for transaction confirmation
            PendingTransactionResponse pTrx = waitForConfirmation(client, id, 4);
            System.out.println("Transaction " + id + " confirmed in round " + pTrx.confirmedRound);
            // Read the transaction
            JSONObject jsonObj2 = new JSONObject(pTrx.toString());
            System.out.println("Transaction information (with notes): " + jsonObj2.toString(2));
            System.out.println("Decoded note: " + new String(pTrx.txn.tx.note));
            System.out.println("Amount: " + new String(pTrx.txn.tx.amount.toString()));
            System.out.println("Fee: " + new String(pTrx.txn.tx.fee.toString()));
            if (pTrx.closingAmount != null){
             System.out.println("Closing Amount: " + new String(pTrx.closingAmount.toString()));
            }
            printBalance(myAccount);


        } catch (Exception e) {
            System.err.println("Exception when calling algod#transactionInformation: " + e.getMessage());
        }
    }

    public static void main(String args[]) throws Exception {
        GettingStarted t = new GettingStarted();
        Account myAccount1 = t.createAccount();
        t.gettingStartedExample(myAccount1);
    }
  }

```

[Run Code](https://replit.com/@Algorand/GettingStarted-with-Java#GettingStarted.java)

[Watch Video](https://youtu.be/_EEMXHrbMzU)

!!! Warning
In order for this transaction to be successful, the account must be funded.


# Setting Up Your Editor/Framework

The Algorand community provides many editors, frameworks, and plugins that can be used to work with the Algorand Network. Tutorials have been created for configuring each of these for use with Algorand. Select your Editor preference below.

- [Setting Up VSCode](https://developer.algorand.org/tutorials/vs-code-java/)
- [Algorand Studio](https://developer.algorand.org/articles/intro-algorand-studio-algorand-vs-code-extension/)
- [Algorand Studio VSCode Extension](https://developer.algorand.org/articles/intro-algorand-studio-algorand-vs-code-extension/)
- [AlgoDEA InteliJ Plugin](https://developer.algorand.org/articles/making-development-easier-algodea-intellij-plugin/)
- [Algorand Builder Framework](https://developer.algorand.org/articles/introducing-algorand-builder/)

