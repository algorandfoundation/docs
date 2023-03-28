title: Your First Transaction

This section is a quick start guide for interacting with the Algorand network using JavaScript. This guide will help to install the ***[Algorand sandbox](https://github.com/algorand/sandbox){:target="_blank"}***, which provides a node for testing and development. This guide will also help to install the JavaScript SDK, create an account and submit your first transaction using different JavaScript Runtimes.
​
# Install Sandbox

!!! info
    This step is only required if you are not using AlgoKit. If you are using AlgoKit, you can spin up a sandbox using the LocalNet, see [AlgoKit getting started guide](/docs/get-started/algokit.md#start-a-localnet) for more information. 

!!! Prerequisites
    - Docker Compose ([install guide](https://docs.docker.com/compose/install/))
    - Git ([install guide](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)) 
 
Algorand provides a Docker instance for setting up a node or private network, which can be used to get started developing. You can find more information about setting up a development environment in [this section](/docs/get-started/algokit). To install and use this instance, follow these instructions.
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

<!-- ===JSSDK_ACCOUNT_GENERATE=== -->
```javascript
const generatedAccount = algosdk.generateAccount();
const passphrase = algosdk.secretKeyToMnemonic(generatedAccount.sk);
console.log(`My address: ${generatedAccount.addr}`);
console.log(`My passphrase: ${passphrase}`);
```
[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/accounts.ts#L80-L84)
<!-- ===JSSDK_ACCOUNT_GENERATE=== -->


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
​
# Connect Your Client
Client must be instantiated prior to making calls to the API endpoints. You must provide values for `<algod-address>` and `<algod-token>`. The CLI tools implement the client natively. By default, the `algodToken` for each [sandbox](https://github.com/algorand/sandbox) is set to its `aaa...` value (64 "a"s) with server address `http://localhost` and port `4001`.

<!-- ===JSSDK_ALGOD_CREATE_CLIENT=== -->
```javascript
const algodToken = 'a'.repeat(64);
const algodServer = 'http://localhost';
const algodPort = 4001;

const algodClient = new algosdk.Algodv2(algodToken, algodServer, algodPort);
```
[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/overview.ts#L7-L12)
<!-- ===JSSDK_ALGOD_CREATE_CLIENT=== -->
 
!!! Info
    The example code connects to the sandbox Algod client. If you want to connect to a other clients, see [Purestake](https://developer.purestake.io/code-samples) or [AlgoExplorer Developer API](https://algoexplorer.io/api-dev/v2).
 
# Check Your Balance
Before moving on to the next step, make sure your account has been funded by the faucet.
 
<!-- ===JSSDK_ALGOD_FETCH_ACCOUNT_INFO=== -->
```javascript
const acctInfo = await algodClient.accountInformation(acct.addr).do();
console.log(`Account balance: ${acctInfo.amount} microAlgos`);
```
[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/overview.ts#L46-L48)
<!-- ===JSSDK_ALGOD_FETCH_ACCOUNT_INFO=== -->

 
# Build First Transaction
To interact with the Algorand blockchain, you can send different types of transactions. The following code shows how to create a payment transaction to transfer Algo tokens to a different address. To construct a transaction, you need to retrieve the parameters about the Algorand network first. You can choose to set a fee yourself, however, by default the fee is set to 1000 microAlgos (0.001 Algo). Optionally, you can add a message to the transaction using the `note` field (up to 1000 bytes). You can find more information about transaction fields in the [documentation](https://developer.algorand.org/docs/get-details/transactions/transactions/#common-fields-header-and-type).
​
<!-- ===JSSDK_TRANSACTION_PAYMENT_CREATE=== -->
```javascript
const suggestedParams = await algodClient.getTransactionParams().do();
const ptxn = algosdk.makePaymentTxnWithSuggestedParamsFromObject({
  from: acct.addr,
  suggestedParams,
  to: acct2.addr,
  amount: 10000,
  note: new Uint8Array(Buffer.from('hello world')),
});
```
[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/overview.ts#L23-L31)
<!-- ===JSSDK_TRANSACTION_PAYMENT_CREATE=== -->

​
!!! Info
    Algorand supports many transaction types. To see what types are supported see [Transactions](https://developer.algorand.org/docs/get-details/transactions/).
​
# Sign First Transaction
Before the transaction is considered valid, it must be signed by a private key. Use the following code to sign the transaction. Now, you can extract the transaction ID. Actually, you can even extract the transaction ID before signing the transaction. You'll use the `txId` to look up the status of the transaction in the following sections of this guide.
​
<!-- ===JSSDK_TRANSACTION_PAYMENT_SIGN=== -->
```javascript
const signedTxn = ptxn.signTxn(acct.privateKey);
```
[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/overview.ts#L34-L35)
<!-- ===JSSDK_TRANSACTION_PAYMENT_SIGN=== -->

!!! Info
    Algorand provides many ways to sign transactions. To see other ways see [Authorization](https://developer.algorand.org/docs/features/transactions/signatures/#single-signatures).
 

 
# Submit the Transaction
The signed transaction can now be submitted to the network.`waitForConfirmation` is called after the transaction is submitted to wait until the transaction is broadcast to the Algorand blockchain and is confirmed. The below snippet also shows how you can decode the data in the node field again to make it readable.
 
 
 <!-- ===JSSDK_TRANSACTION_PAYMENT_SUBMIT=== -->
```javascript
const { txId } = await algodClient.sendRawTransaction(signedTxn).do();
const result = await algosdk.waitForConfirmation(algodClient, txId, 4);
console.log(result);
console.log(`Transaction Information: ${result.txn}`);
console.log(`Decoded Note: ${Buffer.from(result.txn.txn.note).toString()}`);
```
[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/overview.ts#L38-L43)
 <!-- ===JSSDK_TRANSACTION_PAYMENT_SUBMIT=== -->
 
​
!!! Warning
    In order for this transaction to be successful, the generated account must be [funded](https://dispenser.testnet.aws.algodev.network/).
​
# View the Transaction
To view the transaction, open [AlgoExplorer](https://testnet.algoexplorer.io/){target=_blank} or [Goal Seeker](https://goalseeker.purestake.io/algorand/testnet){target=_blank} and paste the transaction ID into the search bar.  