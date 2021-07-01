title: Getting Started With JavaScript
​
This section is a quick start guide for interacting with the Algorand network using JavaScript. This guide will help to install ***sandbox***, which provides a node for testing and development. This guide will also help to install the JavaScript SDK, create an account and submit your first transaction using different JavaScript Runtimes.
​
# Alternative Guide
 
If you are a visual learner, try our [live demo](https://replit.com/@Algorand/gettingStartedPython#main.py) of submitting your first transaction or watch a [full video]() that explains the following steps.
 
 
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
// Must Have X installed
// initialize project
npm init
// install Algorand sdk
npm install algosdk
 
```
[`Watch Video`](https://this_is_url/)
[`More Information`](https://this_is_url/)
 
``` javascript tab="React"
//Must have X installed
npx create-react-app my-app
cd my-app
npm install algosdk
yarn start
​
//Or clone the Algorand Starter Project
git clone need repo
```
[`Watch Video`](https://this_is_url/)
[`More Information`](https://this_is_url/)
  
The SDK is installed with the specific runtime and can now interact with the Sandbox created earlier.
​
!!! Info
   Using a Web Runtime requires the AlgoSigner or plugin or other web-based private key management software. For more information see xxxxxxx
​
# Create an Account on Algorand
In order to interact with the Algorand blockchain, you must have a funded account. To quickly create a test account use the following code.
​
```javascript
var account = algosdk.generateAccount();
var passphrase = algosdk.secretKeyToMnemonic(account.sk);
console.log( "My address: " + account.addr );
console.log( "My passphrase: " + passphrase );
```
[Watch Video](https://this_is_url/)
[More Information](https://developer.algorand.org/docs/features/accounts/create/#standalone)
 
!!! Tip
   Make sure to save your account's address and passphrase at a seperate place, as they will be used later on.
​
!!! Warning 
    Never share Mnemonic private keys. Production environments require stringent private key management. For more information on key management in community Wallets, click [here](https://developer.algorand.org/docs/community/#wallets). For the [Algorand open source wallet](https://developer.algorand.org/articles/algorand-wallet-now-open-source/), click [here](https://github.com/algorand/algorand-wallet).
 
​
# Fund the Account
The code above prompts to fund the newly created account. Before sending transactions to the Algorand network, the account must be funded to cover the minimal transaction fees that exist on Algorand. To fund the account use the [Algorand faucet](https://dispenser.testnet.aws.algodev.network/).
​
<button name="button" onclick="http://www.google.com">Watch Video</button>
​
!!! Info
   All Algorand accounts require a minimum balance to be registered in the ledger. To read more about Algorand minimum balance see [Account Overview](https://developer.algorand.org/docs/features/accounts/#minimum-balance)
 
​
# Connect Your Client
Client must be instantiated prior to making calls to the API endpoints. You must provide values for `<algod-address>` and `<algod-token>`. The CLI tools implement the client natively.
 
```python
from algosdk.v2client import algod
 
algod_address = "http://localhost:4001"
algod_token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
algod_client = algod.AlgodClient(algod_token, algod_address)
```
 
!!! Info
The example code connects to the sandbox Algod client. If you want to connect to a Purestake client, see [Purestake](https://developer.purestake.io/code-samples).
 
# Check Your Balance
Before moving on to the next step, make sure your account has been funded by the faucet.
 
```python
   account_info = algod_client.account_info(my_address)
   print("Account balance: {} microAlgos".format(account_info.get('amount')) + "\n")
```
[`Watch Video`](https://this_is_url/)
 
# Build First Transaction
Communication with the Algorand network is performed using transactions. To create a payment transaction use the following code.
​
```javascript tab="Node"
// Construct the transaction
// replace myAddress with address created earlier
let params = await algodClient.getTransactionParams().do();
const receiver = "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A";
let note = algosdk.encodeObj("Hello World");
​
let txn = algosdk.makePaymentTxnWithSuggestedParams(myAccount.addr, receiver, 1000000, undefined, note, params);
[`Watch Video`](https://this_is_url/)
```
​
```javascript tab="React"
// Construct the transaction
// replace myAddress with address created earlier
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
// Construct the transaction
// replace myAddress with address created earlier
let myAddress = myAccount.getAddress();
final String RECEIVER = "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A";
String note = "Hello World";
TransactionParametersResponse params = client.TransactionParams().execute().body();
Transaction txn = Transaction.PaymentTransactionBuilder()
.sender(myAddress)
.amount(100000)
.receiver(new Address(RECEIVER))
.suggestedParams(params)
.build();
[`Watch Video`](https://this_is_url/)
```
​
```javascript tab="React"
need react code here
[`Watch Video`](https://this_is_url/)
```
​
​!!! Info
   Algorand provides many ways to sign transactions. To see other ways see [Authorization](https://developer.algorand.org/docs/features/transactions/signatures/#single-signatures).
 
# Submit the Transaction
The signed transaction can now be submitted to the network.`waitForConfirmation` is called after the transaction is submitted to wait until the transaction is broadcast to the Algorand blockchain and is confirmed. For more information, see [Wait for Confirmation](https://developer.algorand.org/docs/build-apps/hello_world/#wait-for-confirmation)
 
 
​
```javascript tab="Node"
byte[] encodedTxBytes = Encoder.encodeToMsgPack(signedTxn);
String id = client.RawTransaction().rawtxn(encodedTxBytes).execute().body().txId;
System.out.println("Successfully sent tx with ID: " + id);
[`Run Code`](https://this_is_url/)
[`Watch Video`](https://this_is_url/)
```
​
```javascript tab="React"
need react code here
[`Run Code`](https://this_is_url/)
[`Watch Video`](https://this_is_url/)
```
 
# Complete Example
 
​The complete example below illustrates how to quickly submit your first transaction.
 
```javascript tab="Node"
const algosdk = require('algosdk');
const getInput = require('cli-interact');
​
​
async function yourFirstTransaction() {
​
   try {
       var myAccount = algosdk.generateAccount();
       var passphrase = algosdk.secretKeyToMnemonic(myAccount.sk);
       console.log("My address: " + myAccount.addr);
       console.log("My passphrase: " + passphrase);
​
       let query = getInput.getYesNo;
       var answer = query('Has the Account Been Funded');
       console.log('you answered:', answer);
​
       // Add funds
       // TestNet Faucet: https://bank.testnet.algorand.network/
       // BetaNet Faucet: https://bank.betanet.algodev.network/
​
       // Connect your client
       //const algodToken = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
       //const algodServer = "http://localhost";
       //const algodPort = 4001;
​
       const algodServer = "https://testnet-algorand.api.purestake.io/ps2";
       const algodPort = '';
       const algodToken = {
           'X-API-Key': 'B3SU4KcVKi94Jap2VXkK83xx38bsv95K5UZm2lab'
       }
​
​
       let algodClient = new algosdk.Algodv2(algodToken, algodServer, algodPort);
​
       //Check your balance
       let accountInfo = await algodClient.accountInformation(myAccount.addr).do();
       console.log("Account balance: %d microAlgos", accountInfo.amount);
​
       // Construct the transaction
       let params = await algodClient.getTransactionParams().do();
​
       // receiver defined as TestNet faucet address
       const receiver = "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A";
       const enc = new TextEncoder();
       const note = enc.encode("Hello World");
​
       let txn = algosdk.makePaymentTxnWithSuggestedParams(myAccount.addr, receiver, 1000000, undefined, note, params);
​
       // Sign the transaction
       let signedTxn = txn.signTxn(myAccount.sk);
       let txId = txn.txID().toString();
       console.log("Signed transaction with txID: %s", txId);
​
       // Submit the transaction
       await algodClient.sendRawTransaction(signedTxn).do();
​
   } catch (err) {
       console.log("err", err);
   }
};
yourFirstTransaction();
​```
[Run Code](https://repl.it/@JasonWeathersby/AlgorandFirstTransaction#index.js))
[Watch Video](https://this_is_url/)
​[`Go to Github`](https://github.com/algorand/docs/blob/staging/examples/start_building/v2/python/your_first_transaction.py)
​
```javascript tab="React"
//need react code here
​
```
​
!!! Warning
   In order for this transaction to be successful, the account must be funded.
​
# Viewing the Transaction
To view the transaction, open the [Algorand Blockchain Explorer](https://testnet.algoexplorer.io/) or [Goal Seeker](https://goalseeker.purestake.io/algorand/testnet) and paste the transaction ID into the search bar.
​
[Watch Video](https://this_is_url/)
​​
# Setting Up Your Editor/Framework
The Algorand community provides many editors, frameworks, and plugins that can be used to work with the Algorand Network. Tutorials have been created for configuring each of these for use with Algorand. Select your Editor preference below.
​
* [Setting Up VSCode](https://developer.algorand.org/tutorials/vs-code-javascript/)
* [Algorand Studio](https://developer.algorand.org/articles/intro-algorand-studio-algorand-vs-code-extension/)
* [Algorand Studio VSCode Extension](https://developer.algorand.org/articles/intro-algorand-studio-algorand-vs-code-extension/)
* [AlgoDEA InteliJ Plugin](https://developer.algorand.org/articles/making-development-easier-algodea-intellij-plugin/)
* [Algorand Builder Framework](https://developer.algorand.org/articles/introducing-algorand-builder/)
