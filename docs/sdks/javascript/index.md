title: Your First Transaction

This section is a quick start guide for interacting with the Algorand network using JavaScript. This guide will help to install the ***[Algorand sandbox](https://github.com/algorand/sandbox){:target="_blank"}***, which provides a node for testing and development. This guide will also help to install the JavaScript SDK, create an account and submit your first transaction using different JavaScript Runtimes.
​
# Alternative Guide
 
If you are a visual learner, try our [live demo](https://replit.com/@Algorand/Getting-Started-with-JavaScript){:target="_blank"} of submitting your first transaction or watch a [full video](https://youtu.be/WuhaGp2yrak){:target="_blank"} that explains the following steps.

 
# Sandbox Install
!!! Prerequisites
    - Docker Compose ([install guide](https://docs.docker.com/compose/install/))
    - Git ([install guide](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)) 
 
Algorand provides a Docker instance for setting up a node or private network, which can be used to get started developing. You can find more information about setting up a development environment in [this section](https://developer.algorand.org/docs/get-started/devenv/). To install and use this instance, follow these instructions.
​
```bash
git clone https://github.com/algorand/sandbox.git
cd sandbox
./sandbox up testnet
```

[More Information about the sandbox](https://developer.algorand.org/articles/introducing-sandbox-20/) and [how to use](https://developer.algorand.org/tutorials/exploring-the-algorand-sandbox/) it.
​

This will install a sandbox node connected to the Algorand TestNet. To read more about Algorand networks see [Algorand Networks](https://developer.algorand.org/docs/reference/algorand-networks/).
​

!!! Info 
    The Indexer allows quick searching of the entire  blockchain for transactions, assets, applications and accounts in a timely manner. To learn more about this capability, see [Searching the Blockchain](https://developer.algorand.org/docs/features/indexer/). When running Algorand Sandbox for TestNet, BetaNet or MainNet, you will not have access to the Sandbox Algorand Indexer. When running a private network with `./sandbox up`, you'll have access to the indexer configured for your private network.
 
!!! Warning
    The sandbox installation may take a few minutes to startup in order to catch up to the current block round. To learn more about fast catchup, see [Sync Node Network using Fast Catchup](https://developer.algorand.org/docs/run-a-node/setup/install/#sync-node-network-using-fast-catchup)
    .
# Install SDK For Runtime
Algorand provides an [SDK for JavaScript](https://github.com/algorand/js-algorand-sdk). The instructions for installing the SDK will depend on what runtime you plan on using. 

!!! Prerequisites
   - Install [Node.js](https://nodejs.org/download)
​
```bash 
# initialize project
npm init
# install Algorand sdk
npm install algosdk
# list the version
npm list algosdk

# This package provides TypeScript types, but you will need TypeScript version 4.2 or higher to use them properly.
```

[`Watch Video`](https://youtu.be/WuhaGp2yrak?t=164 ){:target="_blank"}

[`More Information`](https://github.com/algorand/js-algorand-sdk)
 
The SDK is installed with the specific runtime and can now interact with the Sandbox created earlier.

​
!!! Info
    Using a Web Runtime requires the AlgoSigner extension or other web-based private key management software. For more information see [community wallets](https://developer.algorand.org/docs/community/#wallets). 
​
# Create an Account on Algorand
In order to interact with the Algorand blockchain, you must have a funded account. To quickly create a test account use the following code. The account object contains an address (`addr`) and private key (`sk`). You can also export the mnemonic so you can later import the account in key management software like AlgoSigner.
​
```javascript
const algosdk = require('algosdk');
const createAccount = function() {
    try {  
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

[Watch Video](https://youtu.be/WuhaGp2yrak?t=212){:target="_blank"}

[More Information](https://developer.algorand.org/docs/features/accounts/create/#standalone)
 
!!! Tip
    Make sure to save your account's address and passphrase at a separate place, as they will be used later on.    ​

!!! Warning 
    Never share mnemonic private keys. Production environments require stringent private key management. For more information on key management in community wallets, click [here](https://developer.algorand.org/docs/community/#wallets). For the [Algorand open source wallet](https://developer.algorand.org/articles/algorand-wallet-now-open-source/), click [here](https://github.com/algorand/algorand-wallet).

​
# Fund the Account
The code below prompts to fund the newly created account. Before sending transactions to the Algorand network, the account must be funded to cover the minimal transaction fees that exist on Algorand. To fund the account use the [Algorand TestNet faucet](https://dispenser.testnet.aws.algodev.network/).

​
!!! Info
    All Algorand accounts require a minimum balance to be registered in the ledger. To read more about Algorand minimum balance see [Account Overview](https://developer.algorand.org/docs/features/accounts/#minimum-balance)
 
​
[Watch Video](https://youtu.be/WuhaGp2yrak?t=307){:target="_blank"}

# Viewing the Transaction
To view the transaction, click on the transaction link in the dispenser or copy the transaction ID and paste to the search bar on the [Algorand Blockchain Explorer](https://testnet.algoexplorer.io/) or [Goal Seeker](https://goalseeker.purestake.io/algorand/testnet).
​

[Watch Video](https://youtu.be/WuhaGp2yrak?t=326){:target="_blank"}
​​
# Connect Your Client
Client must be instantiated prior to making calls to the API endpoints. You must provide values for `<algod-address>` and `<algod-token>`. The CLI tools implement the client natively. By default, the `algodToken` for each [sandbox](https://github.com/algorand/sandbox) is set to its `aaa...` value (64 "a"s) with server address `http://localhost` and port `4001`.
 
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


```

 
# Build First Transaction
To interact with the Algorand blockchain, you can send different types of transactions. The following code shows how to create a payment transaction to transfer Algo tokens to a different address. To construct a transaction, you need to retrieve the parameters about the Algorand network first. You can choose to set a fee yourself, however, by default the fee is set to 1000 microAlgos (0.001 Algo). Optionally, you can add a message to the transaction using the `note` field (up to 1000 bytes). You can find more information about transaction fields in the [documentation](https://developer.algorand.org/docs/get-details/transactions/transactions/#common-fields-header-and-type).
​
```javascript 
        // Construct the transaction
        let params = await algodClient.getTransactionParams().do();
        // comment out the next two lines to use suggested fee
        params.fee = algosdk.ALGORAND_MIN_TX_FEE;
        params.flatFee = true;

        const receiver = "HZ57J3K46JIJXILONBBZOHX6BKPXEM2VVXNRFSUED6DKFD5ZD24PMJ3MVA";
        const enc = new TextEncoder();
        const note = enc.encode("Hello World");
        let amount = 1000000; // equals 1 ALGO
        let sender = myAccount.addr;

        let txn = algosdk.makePaymentTxnWithSuggestedParamsFromObject({
            from: sender, 
            to: receiver, 
            amount: amount, 
            note: note, 
            suggestedParams: params
        });
```
[`Watch Video`](https://youtu.be/WuhaGp2yrak?t=386){:target="_blank"}


​
!!! Info
    Algorand supports many transaction types. To see what types are supported see [Transactions](https://developer.algorand.org/docs/get-details/transactions/).
​
# Sign First Transaction
Before the transaction is considered valid, it must be signed by a private key. Use the following code to sign the transaction. Now, you can extract the transaction ID. Actually, you can even extract the transaction ID before signing the transaction. You'll use the `txId` to look up the status of the transaction in the following sections of this guide.
​
```javascript 
        // Sign the transaction
        let signedTxn = txn.signTxn(myAccount.sk);
        let txId = txn.txID().toString();
        console.log("Signed transaction with txID: %s", txId);
```
​
[`Watch Video`](https://youtu.be/WuhaGp2yrak?t=500){:target="_blank"}

!!! Info
    Algorand provides many ways to sign transactions. To see other ways see [Authorization](https://developer.algorand.org/docs/features/transactions/signatures/#single-signatures).
 

 
# Submit the Transaction
The signed transaction can now be submitted to the network.`waitForConfirmation` is called after the transaction is submitted to wait until the transaction is broadcast to the Algorand blockchain and is confirmed. The below snippet also shows how you can decode the data in the node field again to make it readable.
 
 
​
```javascript
        // Submit the transaction
        await algodClient.sendRawTransaction(signedTxn).do();

        // Wait for confirmation
        let confirmedTxn = await algosdk.waitForConfirmation(algodClient, txId, 4);
        //Get the completed Transaction
        console.log("Transaction " + txId + " confirmed in round " + confirmedTxn["confirmed-round"]);
        // let mytxinfo = JSON.stringify(confirmedTxn.txn.txn, undefined, 2);
        // console.log("Transaction information: %o", mytxinfo);
        let string = new TextDecoder().decode(confirmedTxn.txn.txn.note);
        console.log("Note field: ", string);
        accountInfo = await algodClient.accountInformation(myAccount.addr).do();
        console.log("Transaction Amount: %d microAlgos", confirmedTxn.txn.txn.amt);        
        console.log("Transaction Fee: %d microAlgos", confirmedTxn.txn.txn.fee);
        console.log("Account balance: %d microAlgos", accountInfo.amount);

```

[`Watch Video`](https://youtu.be/WuhaGp2yrak?t=508){:target="_blank"}
 
# Complete Example
 
​The complete example below illustrates how to quickly submit your first transaction. If you want to learn more about other transaction types, you can read the documentation and try out some of the [examples listed on GitHub](https://github.com/algorand/js-algorand-sdk/tree/develop/examples) to quickly learn more.
 
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

        //Check your balance
        let accountInfo = await algodClient.accountInformation(myAccount.addr).do();
        console.log("Account balance: %d microAlgos", accountInfo.amount);

        // Construct the transaction
        let params = await algodClient.getTransactionParams().do();
        // comment out the next two lines to use suggested fee
        params.fee = algosdk.ALGORAND_MIN_TX_FEE;
        params.flatFee = true;

        // receiver defined as TestNet faucet address 
        const receiver = "HZ57J3K46JIJXILONBBZOHX6BKPXEM2VVXNRFSUED6DKFD5ZD24PMJ3MVA";
        const enc = new TextEncoder();
        const note = enc.encode("Hello World");
        let amount = 1000000;
        let sender = myAccount.addr;
        let txn = algosdk.makePaymentTxnWithSuggestedParamsFromObject({
            from: sender, 
            to: receiver, 
            amount: amount, 
            note: note, 
            suggestedParams: params
        });


        // Sign the transaction
        let signedTxn = txn.signTxn(myAccount.sk);
        let txId = txn.txID().toString();
        console.log("Signed transaction with txID: %s", txId);

        // Submit the transaction
        await algodClient.sendRawTransaction(signedTxn).do();

        // Wait for confirmation
        let confirmedTxn = await algosdk.waitForConfirmation(algodClient, txId, 4);
        //Get the completed Transaction
        console.log("Transaction " + txId + " confirmed in round " + confirmedTxn["confirmed-round"]);
        let string = new TextDecoder().decode(confirmedTxn.txn.txn.note);
        console.log("Note field: ", string);
        accountInfo = await algodClient.accountInformation(myAccount.addr).do();
        console.log("Transaction Amount: %d microAlgos", confirmedTxn.txn.txn.amt);        
        console.log("Transaction Fee: %d microAlgos", confirmedTxn.txn.txn.fee);
     
        console.log("Account balance: %d microAlgos", accountInfo.amount);
    }
    catch (err) {
        console.log("err", err);
    }
    process.exit();
};

firstTransaction();
```

[Run Code](https://replit.com/@Algorand/Getting-Started-with-JavaScript){:target="_blank"}

[Watch Video](https://youtu.be/WuhaGp2yrak){:target="_blank"}


​
!!! Warning
    In order for this transaction to be successful, the generated account must be [funded](https://dispenser.testnet.aws.algodev.network/).
​

# Setting Up Your Editor/Framework
The Algorand community provides many editors, frameworks, and plugins that can be used to work with the Algorand Network. Tutorials have been created for configuring each of these for use with Algorand. Select your Editor preference below.
​

* [Setting Up VSCode](https://developer.algorand.org/tutorials/vs-code-javascript/)

* [Algorand Studio](https://developer.algorand.org/articles/intro-algorand-studio-algorand-vs-code-extension/)

* [Algorand Studio VSCode Extension](https://developer.algorand.org/articles/intro-algorand-studio-algorand-vs-code-extension/)

* [AlgoDEA IntelliJ Plugin](https://developer.algorand.org/articles/making-development-easier-algodea-intellij-plugin/)

* [Algorand Builder Framework](https://developer.algorand.org/articles/introducing-algorand-builder/) and [Algo Builder Tutorial series](https://developer.algorand.org/tutorials/algorand-builder-tutorial-part1-creating-local-network-and-deploying-asa/)

