title: Getting Started With JavaScript
​
This section is a quick start guide for interacting with the Algorand network using JavaScript. This guide will help to install ***sandbox***, which provides a node for testing and development. This guide will also help to install the JavaScript SDK, create an account and submit your first transaction using different JavaScript Runtimes.
​
# Alternative Guide
 
If you are a visual learner, try our [live demo](https://replit.com/@Algorand/Getting-Started-with-JavaScript#index.js) of submitting your first transaction or watch a [full video]() that explains the following steps.
 
 
# Sandbox Install
!!! Prerequisites
   - Docker Compose ([install guide](https://docs.docker.com/compose/install/))
   - Git ([install guide](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)) 
 
Algorand provides a docker instance for setting up a node, which can be used to get started developing. To install and use this instance, follow these instructions.
​
```bash
git clone https://github.com/algorand/sandbox.git
cd sandbox
./sandbox up testnet
```
 
​[Watch Video](https://youtu.be/uOZ95YzU9hU)
 
[More Information](https://developer.algorand.org/articles/introducing-sandbox-20/)
​
This will install a Sandbox node connected to the Algorand TestNet. To read more about Algorand networks see [Algorand Networks](https://developer.algorand.org/docs/reference/algorand-networks/).
​
​To use Indexer in the sandbox, start it to the default private network as follows. 
 
```bash
./sandbox up
```
 
!!! Info 
    The Indexer allows quick searching of the entire  blockchain for transactions, assets, applications and accounts in a timely manner. To learn more about this capability, see [Searching the Blockchain](https://developer.algorand.org/docs/features/indexer/). When running Algorand Sandbox for TestNet, BetaNet or MainNet, you will not have access to the Sandbox Algorand Indexer. 
 
!!! Warning
    The sandbox installation may take a few minutes to startup in order to catch up to the current block round. To learn more about fast catchup, see [Sync Node Network using Fast Catchup](https://developer.algorand.org/docs/run-a-node/setup/install/#sync-node-network-using-fast-catchup)
    .
# Install SDK For Runtime
Algorand provides an SDK for JavaScript. The instructions for installing the SDK will depend on what runtime you plan on using.
​
``` javascript tab="Node"

// Install Node.js https://nodejs.org/download
// initialize project
npm init
// install Algorand sdk
npm install algosdk
// list the version
npm list algosdk

// This package provides TypeScript types, but you will need TypeScript version 4.2 or higher to use them properly.
 
```
[`Watch Video`](https://this_is_url/)
[`More Information`](https://this_is_url/)
 
``` javascript tab="React"
//Get started with React https://reactjs.org/
npx create-react-app my-app
cd my-app
npm install algosdk
yarn start
​
// Or clone the Algorand Starter Project

//git clone need https://github.com/rekpero/algorand-sdk-react-component

[`More Information`](https://github.com/rekpero/algorand-sdk-react-component)
```

The SDK is installed with the specific runtime and can now interact with the Sandbox created earlier.
​
!!! Info
   Using a Web Runtime requires the AlgoSigner or plugin or other web-based private key management software. For more information see [community wallets](https://developer.algorand.org/docs/community/#wallets). 
​
# Create an Account on Algorand
In order to interact with the Algorand blockchain, you must have a funded account. To quickly create a test account use the following code.
​
```javascript
const algosdk = require('algosdk');
const createAccount =  function (){
    try{  
        const myaccount = algosdk.generateAccount();
        console.log("Account Address = " + myaccount.addr);
        let account_mnemonic = algosdk.secretKeyToMnemonic(myaccount.sk);
        console.log("Account Mnemonic = "+ account_mnemonic);
        console.log("Account created. Save off Mnemonic and address");
        console.log("Add funds to account using the TestNet Dispenser: ");
        console.log("https://dispenser.testnet.aws.algodev.network/ ");
        return myaccount;
    }
    catch (err) {
        console.log("err", err);
    }
};
```
[Watch Video](https://this_is_url/)

[More Information](https://developer.algorand.org/docs/features/accounts/create/#standalone)
 
!!! Tip
    Make sure to save your account's address and passphrase at a separate place, as they will be used later on.
    ​

!!! Warning 
    Never share Mnemonic private keys. Production environments require stringent private key management. For more information on key management in community Wallets, click [here](https://developer.algorand.org/docs/community/#wallets). For the [Algorand open source wallet](https://developer.algorand.org/articles/algorand-wallet-now-open-source/), click [here](https://github.com/algorand/algorand-wallet).

[Watch Video](https://this_is_url/) 
​
# Fund the Account
The code above prompts to fund the newly created account. Before sending transactions to the Algorand network, the account must be funded to cover the minimal transaction fees that exist on Algorand. To fund the account use the [Algorand faucet](https://dispenser.testnet.aws.algodev.network/).

​
!!! Info
   All Algorand accounts require a minimum balance to be registered in the ledger. To read more about Algorand minimum balance see [Account Overview](https://developer.algorand.org/docs/features/accounts/#minimum-balance)
 
​[Watch Video](https://this_is_url/)

# Viewing the Transaction
To view the transaction, open the [Algorand Blockchain Explorer](https://testnet.algoexplorer.io/) or [Goal Seeker](https://goalseeker.purestake.io/algorand/testnet) and paste the transaction ID into the search bar.
​
[Watch Video](https://this_is_url/)
​​
# Connect Your Client
Client must be instantiated prior to making calls to the API endpoints. You must provide values for `<algod-address>` and `<algod-token>`. The CLI tools implement the client natively.
 
```javascript
async function firstTransaction() {

    try {
        let myAccount = createAccount();
        console.log("Press any key when the account is funded");
        await keypress();
        // Connect your client
        const algodToken = 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa';
        const algodServer = 'http://localhost';
        const algodPort = 4001;
        let algodClient = new algosdk.Algodv2(algodToken, algodServer, algodPort);

```
 
!!! Info
The example code connects to the sandbox Algod client. If you want to connect to a other clients, see [Purestake](https://developer.purestake.io/code-samples) or [AlgoExplorer Developer API](https://algoexplorer.io/api-dev/v2).
 
# Check Your Balance
Before moving on to the next step, make sure your account has been funded by the faucet.
 
```javascript
        //Check your balance
        let accountInfo = await algodClient.accountInformation(myAccount.addr).do();
        console.log("Account balance: %d microAlgos", accountInfo.amount);
        let startingAmount = accountInfo.amount;

```
[`Watch Video`](https://this_is_url/)
 
# Build First Transaction
Communication with the Algorand network is performed using transactions. To create a payment transaction use the following code.
​
```javascript 
        // Construct the transaction
        let params = await algodClient.getTransactionParams().do();
        // comment out the next two lines to use suggested fee
        params.fee = 1000;
        params.flatFee = true;

        const receiver = "HZ57J3K46JIJXILONBBZOHX6BKPXEM2VVXNRFSUED6DKFD5ZD24PMJ3MVA";
        const enc = new TextEncoder();
        const note = enc.encode("Hello World");
        let amount = 1000000;
        let sender = myAccount.addr;
        let txn = algosdk.makePaymentTxnWithSuggestedParams(sender, receiver, amount, undefined, note, params);


[`Watch Video`](https://this_is_url/)
```

​
!!! Info
   Algorand supports many transaction types. To see what types are supported see [Transactions](https://developer.algorand.org/docs/features/transactions/).
​
# Sign First Transaction
Before the transaction is considered valid, it must be signed by a private key. Use the following code to sign the transaction.
​
```javascript tab="Node"
       // Sign the transaction
        let signedTxn = txn.signTxn(myAccount.sk);
        let txId = txn.txID().toString();
        console.log("Signed transaction with txID: %s", txId);
```
​[`Watch Video`](https://this_is_url/)


​!!! Info
   Algorand provides many ways to sign transactions. To see other ways see [Authorization](https://developer.algorand.org/docs/features/transactions/signatures/#single-signatures).
 
# Submit the Transaction
The signed transaction can now be submitted to the network.`waitForConfirmation` is called after the transaction is submitted to wait until the transaction is broadcast to the Algorand blockchain and is confirmed. For more information, see [Wait for Confirmation](https://developer.algorand.org/docs/build-apps/hello_world/#wait-for-confirmation)
 
 
​
```javascript
        // Submit the transaction
        await algodClient.sendRawTransaction(signedTxn).do();

        // Wait for confirmation
        let confirmedTxn = await waitForConfirmation(algodClient, txId, 4);
        //Get the completed Transaction
        console.log("Transaction " + txId + " confirmed in round " + confirmedTxn["confirmed-round"]);
        // let mytxinfo = JSON.stringify(confirmedTxn.txn.txn, undefined, 2);
        // console.log("Transaction information: %o", mytxinfo);
        var string = new TextDecoder().decode(confirmedTxn.txn.txn.note);
        console.log("Note field: ", string);
        accountInfo = await algodClient.accountInformation(myAccount.addr).do();
        console.log("Transaction Amount: %d microAlgos", confirmedTxn.txn.txn.amt);        
        console.log("Transaction Fee: %d microAlgos", confirmedTxn.txn.txn.fee);
        console.log("Account balance: %d microAlgos", accountInfo.amount);

```
[`Watch Video`](https://this_is_url/)​

 
# Complete Example
 
​The complete example below illustrates how to quickly submit your first transaction.
 
```javascript 
const algosdk = require('algosdk');
const keypress = async () => {
    process.stdin.setRawMode(true)
    return new Promise(resolve => process.stdin.once('data', () => {
        process.stdin.setRawMode(false)
        resolve()
    })) 
}
// Create an account and add funds to it. Copy the address off
// The Algorand TestNet Dispenser is located here: 
// https://dispenser.testnet.aws.algodev.network/

const createAccount =  function (){
    try{  
        const myaccount = algosdk.generateAccount();
        console.log("Account Address = " + myaccount.addr);
        let account_mnemonic = algosdk.secretKeyToMnemonic(myaccount.sk);
        console.log("Account Mnemonic = "+ account_mnemonic);
        console.log("Account created. Save off Mnemonic and address");
        console.log("Add funds to account using the TestNet Dispenser: ");
        console.log("https://dispenser.testnet.aws.algodev.network/ ");

        return myaccount;
    }
    catch (err) {
        console.log("err", err);
    }
};

/**
 * Wait until the transaction is confirmed or rejected, or until 'timeout'
 * number of rounds have passed.
 * @param {algosdk.Algodv2} algodClient the Algod V2 client
 * @param {string} txId the transaction ID to wait for
 * @param {number} timeout maximum number of rounds to wait
 * @return {Promise<*>} pending transaction information
 * @throws Throws an error if the transaction is not confirmed or rejected in the next timeout rounds
 */
 const waitForConfirmation = async function (algodClient, txId, timeout) {
    if (algodClient == null || txId == null || timeout < 0) {
        throw new Error("Bad arguments");
    }

    const status = (await algodClient.status().do());
    if (status === undefined) {
        throw new Error("Unable to get node status");
    }

    const startround = status["last-round"] + 1;
    let currentround = startround;

    while (currentround < (startround + timeout)) {
        const pendingInfo = await algodClient.pendingTransactionInformation(txId).do();
        if (pendingInfo !== undefined) {
            if (pendingInfo["confirmed-round"] !== null && pendingInfo["confirmed-round"] > 0) {
                //Got the completed Transaction
                return pendingInfo;
            } else {
                if (pendingInfo["pool-error"] != null && pendingInfo["pool-error"].length > 0) {
                    // If there was a pool error, then the transaction has been rejected!
                    throw new Error("Transaction " + txId + " rejected - pool error: " + pendingInfo["pool-error"]);
                }
            }
        }
        await algodClient.statusAfterBlock(currentround).do();
        currentround++;
    }
    throw new Error("Transaction " + txId + " not confirmed after " + timeout + " rounds!");
};
async function firstTransaction() {

    try {
        let myAccount = createAccount();
        console.log("Press any key when the account is funded");
        await keypress();
        // Connect your client
        // const algodToken = '';
        // const algodServer = 'https://testnet.algoexplorerapi.io';
        // const algodPort = '';
        const algodToken = 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa';
        const algodServer = 'http://localhost';
        const algodPort = 4001;
        let algodClient = new algosdk.Algodv2(algodToken, algodServer, algodPort);

        //Check your balance
        let accountInfo = await algodClient.accountInformation(myAccount.addr).do();
        console.log("Account balance: %d microAlgos", accountInfo.amount);
        let startingAmount = accountInfo.amount;
        // Construct the transaction
        let params = await algodClient.getTransactionParams().do();
        // comment out the next two lines to use suggested fee
        params.fee = 1000;
        params.flatFee = true;

        // receiver defined as TestNet faucet address 
        const receiver = "HZ57J3K46JIJXILONBBZOHX6BKPXEM2VVXNRFSUED6DKFD5ZD24PMJ3MVA";
        const enc = new TextEncoder();
        const note = enc.encode("Hello World");
        let amount = 1000000;
        let closeout = receiver; //closeRemainderTo
        let sender = myAccount.addr;
        let txn = algosdk.makePaymentTxnWithSuggestedParams(sender, receiver, amount, undefined, note, params);
        // WARNING! all remaining funds in the sender account above will be sent to the closeRemainderTo Account 
        // In order to keep all remaning funds in the sender account after tx, set closeout parameter to undefined.
        // For more info see: 
        // https://developer.algorand.org/docs/reference/transactions/#payment-transaction

        // Sign the transaction
        let signedTxn = txn.signTxn(myAccount.sk);
        let txId = txn.txID().toString();
        console.log("Signed transaction with txID: %s", txId);

        // Submit the transaction
        await algodClient.sendRawTransaction(signedTxn).do();

        // Wait for confirmation
        let confirmedTxn = await waitForConfirmation(algodClient, txId, 4);
        //Get the completed Transaction
        console.log("Transaction " + txId + " confirmed in round " + confirmedTxn["confirmed-round"]);
        // let mytxinfo = JSON.stringify(confirmedTxn.txn.txn, undefined, 2);
        // console.log("Transaction information: %o", mytxinfo);
        var string = new TextDecoder().decode(confirmedTxn.txn.txn.note);
        console.log("Note field: ", string);
        accountInfo = await algodClient.accountInformation(myAccount.addr).do();
        console.log("Transaction Amount: %d microAlgos", confirmedTxn.txn.txn.amt);        
        console.log("Transaction Fee: %d microAlgos", confirmedTxn.txn.txn.fee);
        let closeoutamt = startingAmount - confirmedTxn.txn.txn.amt - confirmedTxn.txn.txn.fee;        
        console.log("Close To Amount: %d microAlgos", closeoutamt);
        console.log("Account balance: %d microAlgos", accountInfo.amount);
    }
    catch (err) {
        console.log("err", err);
    }
    process.exit();
};

firstTransaction();
```



[Run Code](https://replit.com/@Algorand/Getting-Started-with-JavaScript))
[Watch Video](https://this_is_url/)

​

​
!!! Warning
   In order for this transaction to be successful, the account must be funded.
​

# Setting Up Your Editor/Framework
The Algorand community provides many editors, frameworks, and plugins that can be used to work with the Algorand Network. Tutorials have been created for configuring each of these for use with Algorand. Select your Editor preference below.
​
* [Setting Up VSCode](https://developer.algorand.org/tutorials/vs-code-javascript/)
* [Algorand Studio](https://developer.algorand.org/articles/intro-algorand-studio-algorand-vs-code-extension/)
* [Algorand Studio VSCode Extension](https://developer.algorand.org/articles/intro-algorand-studio-algorand-vs-code-extension/)
* [AlgoDEA InteliJ Plugin](https://developer.algorand.org/articles/making-development-easier-algodea-intellij-plugin/)
* [Algorand Builder Framework](https://developer.algorand.org/articles/introducing-algorand-builder/)
