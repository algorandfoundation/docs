title: 3. Your First Transaction

After you successfully connect to **algod** using your preferred SDK, explore the methods available to read from and write to the blockchain. Remember that writing to the Algorand blockchain is simply sending a transaction to the network that is later confirmed within a block. 

Follow the guide below to send your first transaction on Algorand and familiarize yourself with some of the core functions of the SDKs. Examples of `goal` commands and REST API calls are included when they are the same or similar, allowing you to cross-verify and gain fluency across all available tools and platforms. 

Code snippets are abbreviated for conciseness and clarity. See the full code example for each SDK at the bottom of this guide.

!!! info
    The examples in this section have been updated to the v2 API, which was launched to MainNet on June 16, 2020. Visit the [v2 Migration Guide](../reference/sdks/migration.md) for information on how to migrate your code from v1. 

    Full running code examples for each SDK and both API versions are available within the GitHub repo at [/examples/start_building](https://github.com/algorand/docs/tree/master/examples/start_building) and for [download](https://github.com/algorand/docs/blob/master/examples/start_building/start_building.zip?raw=true) (.zip).

# Create an account
In order to send a transaction, you first need an [account](../features/accounts/index.md#accounts) on Algorand. Create an account by generating an Algorand public/private key pair and then funding the public address with Algos on your chosen network. 

!!! info
	The terms **account**, **public key**, and **address** are used interchangeably in certain contexts, but they have slightly different meanings. Read more about these differences in the [Accounts Overview](../features/accounts/index.md).


## Generate a public/private key pair

```javascript tab="JavaScript"
const algosdk = require('algosdk');

var account = algosdk.generateAccount();
var passphrase = algosdk.secretKeyToMnemonic(account.sk);
console.log( "My address: " + account.addr );
console.log( "My passphrase: " + passphrase );
```

```python tab="Python"
from algosdk import account, mnemonic

def generate_algorand_keypair():
	private_key, address = account.generate_account()
	print("My address: {}".format(address))
	print("My passphrase: {}".format(mnemonic.from_private_key(private_key)))
```

```java tab="Java"
import com.algorand.algosdk.account.Account;	

Account myAccount = new Account();
System.out.println("My Address: " + myAccount.getAddress());
System.out.println("My Passphrase: " + myAccount.toMnemonic());
```

```go tab="Go"
import (
	"fmt"

	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/mnemonic"
)

func main() {
	account := crypto.GenerateAccount()
	passphrase, err := mnemonic.FromPrivateKey(account.PrivateKey)

	if err != nil {
		fmt.Printf("Error creating transaction: %s\n", err)
	} else {
		fmt.Printf("My address: %s\n", account.Address)
		fmt.Printf("My passphrase: %s\n", passphrase)
	}
}
```

```text tab="goal"
$ goal account new
Created new account with address [ADDRESS]

$ goal account export -a <address>
Exported key for account [ADDRESS]: [PASSPHRASE]
```

```bash tab="algokey"
$ algokey generate
Private key mnemonic: [PASSPHRASE]
Public key: [ADDRESS]
```

_Learn more about [Creating Accounts on Algorand](../features/accounts/create.md)._

## Add funds
For [TestNet](../../reference/algorand-networks/testnet/#faucet) and [BetaNet](../../reference/algorand-networks/betanet/#faucet), copy and paste the public portion of your key pair in the corresponding faucet prompt and click "Submit". A `200` response means the transaction went through and your balance increased by 100,000,000 microAlgos (i.e. 100 Algos).

!!! info
	Amounts are returned in microAlgos - the base unit for Algos. Micro denotes a unit x 10^-6. Therefore, 1 Algo equals 1,000,000 microAlgos.

## Connect your client

Each SDK provides a client which must instantiate prior to making calls to the API endpoints. You must provide values for `<algod-address>`, `<port>` and `<algod-token>`. The CLI tools implement the client natively. 

_Learn more about [Connecting to a Node](connect.md)._

```JavaScript tab=
const algodToken = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
const algodServer = "http://localhost";
const algodPort = 4001;

let algodClient = new algosdk.Algodv2(algodToken, algodServer, algodPort);
```

```Python tab=
from algosdk.v2client import algod

algod_address = "http://localhost:4001"
algod_token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
algod_client = algod.AlgodClient(algod_token, algod_address)
```

```Java tab=
import com.algorand.algosdk.v2.client.common.AlgodClient;
import com.algorand.algosdk.v2.client.common.Client;

final String ALGOD_API_ADDR = "localhost";
final Integer ALGOD_PORT = 4001;
final String ALGOD_API_TOKEN = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";

AlgodClient client = (AlgodClient) new AlgodClient(ALGOD_API_ADDR, ALGOD_PORT, ALGOD_API_TOKEN);
```

```Go tab=
package main

import (
	"github.com/algorand/go-algorand-sdk/client/v2/algod" 
)

const algodAddress = "http://localhost:4001"
const algodToken = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

func main() {
	algodClient, err := algod.MakeClient(algodAddress, algodToken)
	if err != nil {
        fmt.Printf("Issue with creating algod client: %s\n", err)
        return
	}
}
```

## Check your balance

Check your balance to confirm the added funds.

```javascript tab="JavaScript"

	const passphrase = "Your 25-word mnemonic generated and displayed above";

	let myAccount = algosdk.mnemonicToSecretKey(passphrase)
	console.log("My address: %s", myAccount.addr)

    let accountInfo = await algodClient.accountInformation(myAccount.addr).do();
    console.log("Account balance: %d microAlgos", accountInfo.amount);
```

```python tab="Python"
passphrase = "Your 25-word mnemonic generated and displayed above"

private_key = mnemonic.to_private_key(passphrase)
my_address = mnemonic.to_public_key(passphrase)
print("My address: {}".format(my_address))

account_info = algod_client.account_info(my_address)
print("Account balance: {} microAlgos".format(account_info.get('amount')))
```

```java tab="Java"
final String PASSPHRASE = "Your 25-word mnemonic generated and displayed above";
String myAddress = myAccount.getAddress().toString();

com.algorand.algosdk.v2.client.model.Account accountInfo = client.AccountInformation(myAccount.getAddress()).execute().body();

System.out.println(String.format("Account Balance: %d microAlgos", accountInfo.amount));
```

```go tab="Go"
passphrase := "Your 25-word mnemonic generated and displayed above"
privateKey, err := mnemonic.ToPrivateKey(passphrase)
if err != nil {
    fmt.Printf("Issue with mnemonic conversion: %s\n", err)
    return
}

var myAddress types.Address
publicKey := privateKey.Public()
cpk := publicKey.(ed25519.PublicKey)
copy(myAddress[:], cpk[:])
fmt.Printf("My address: %s\n", myAddress.String())

accountInfo, err := algodClient.AccountInformation(myAddress.String()).Do(context.Background())
if err != nil {
    fmt.Printf("Error getting account info: %s\n", err)
    return
}
fmt.Printf("Account balance: %d microAlgos\n", accountInfo.Amount)
```

```bash tab="cURL"
curl -i -X GET \
   -H "X-Algo-API-Token:<algod-token> \
 'http://<algod-address>:<algod-port>/v1/account/<address>'
```

```bash tab="goal"
$ goal account balance -a <my-address>
[AMOUNT] microAlgos
```

# Construct the transaction

Create a transaction to send 1 Algo from your account to the TestNet faucet address (`GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A`) with the note "Hello World".

Transactions require a certain minimum set of parameters to be valid. Mandatory fields include the **round validity range**, the **fee**, and the **genesis hash** for the network the transaction is valid for. Read all about Transaction types, fields, and configurations in the [Transactions Feature Guide](../features/transactions/index.md). For now, construct a payment transaction as follows. Use the _suggested parameters_ methods to initialize network-related 
fields. 

```javascript tab="JavaScript"
let params = await algodClient.getTransactionParams().do();
// comment out the next two lines to use suggested fee
params.fee = 1000;
params.flatFee = true;
const receiver = "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A";
let note = algosdk.encodeObj("Hello World");

let txn = algosdk.makePaymentTxnWithSuggestedParams(myAccount.addr, receiver, 1000000, undefined, note, params);        
```

```python tab="Python"
params = algod_client.suggested_params()
# comment out the next two (2) lines to use suggested fees
params.flat_fee = True
params.fee = 1000
receiver = "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A"
note = "Hello World".encode()

unsigned_txn = PaymentTxn(my_address, params, receiver, 1000000, None, note)
```

```java tab="Java"
// Construct the transaction
final String RECEIVER = "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A";
String note = "Hello World";
TransactionParametersResponse params = client.TransactionParams().execute().body();
Transaction txn = Transaction.PaymentTransactionBuilder()
.sender(myAddress)
.note(note.getBytes())
.amount(100000)
.receiver(new Address(RECEIVER))
.suggestedParams(params)
.build();
```

```go tab="Go"
txParams, err := algodClient.SuggestedParams().Do(context.Background())
if err != nil {
    fmt.Printf("Error getting suggested tx params: %s\n", err)
    return
}
// comment out the next two (2) lines to use suggested fees
txParams.FlatFee = true
txParams.Fee = 1000

fromAddr := myAddress.String()
toAddr := "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A"
var amount uint64 = 1000000
var minFee uint64 = 1000
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

```bash tab="goal"
$ goal clerk send --from=<my-account> --to=GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A --fee=1000 --amount=1000000 --note="Hello World" --out="hello-world.txn"
```

!!! info
    Some of the SDKs provide wrapper functions for creating certain types of transactions, like `makePaymentTxn` in Go. 

_Learn more about the [Structure of Transactions on Algorand](../features/transactions/index.md)._

# Sign the transaction
Sign the transaction with your private key. This creates a new signed transaction object in the SDKs. Retrieve the transaction ID of the signed transaction.

```javascript tab="JavaScript"
let signedTxn = txn.signTxn(myAccount.sk);
let txId = txn.txID().toString();
console.log("Signed transaction with txID: %s", txId);
```

```python tab="Python"
signed_txn = unsigned_txn.sign(mnemonic.to_private_key(passphrase))
```

```java tab="Java"
SignedTransaction signedTxn = myAccount.signTransaction(txn);
System.out.println("Signed transaction with txid: " + signedTxn.transactionID);
```

```go tab="Go"
txID, signedTxn, err := crypto.SignTransaction(privateKey, txn)
if err != nil {
    fmt.Printf("Failed to sign transaction: %s\n", err)
    return
}
fmt.Printf("Signed txid: %s\n", txID)
```

```bash tab="goal"
$ goal clerk sign --infile="hello-world.txn" --outfile="hello-world.stxn"
```

```bash tab="algokey"
```

_Learn more about [Authorizing Transactions on Algorand](../features/transactions/signatures.md)._

# Submit the transaction
Send the signed transaction to the network with your algod client. 

```javascript tab="JavaScript"
await algodClient.sendRawTransaction(signedTxn).do();
```

```python tab="Python"
txid = algod_client.send_transaction(signed_txn)
print("Successfully sent transaction with txID: {}".format(txid)
```

```java tab="Java"
byte[] encodedTxBytes = Encoder.encodeToMsgPack(signedTxn);
String id = client.RawTransaction().rawtxn(encodedTxBytes).execute().body().txId;
System.out.println("Successfully sent tx with ID: " + id);
```

```go tab="Go"
sendResponse, err := algodClient.SendRawTransaction(signedTxn).Do(context.Background())
if err != nil {
    fmt.Printf("failed to send transaction: %s\n", err)
    return
}
fmt.Printf("Submitted transaction %s\n", sendResponse)
```

```bash tab="cURL"
curl -i -X POST \
   -H "X-Algo-API-Token:<algod-token> \
   -H "Content-Type:application/x-binary" \
   -T "hello-world.stxn" \
 'http://<algod-address>:<algod-port>/v1/transactions'
```

```bash tab="goal"
$ goal clerk rawsend --filename="hello-world.stxn"
Sent 1000000 MicroAlgos from account [ADDRESS] to address GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A, transaction ID: [TXID]. Fee set to 1000
Transaction [TXID] still pending as of round [LAST_ROUND]
Transaction [TXID] committed in round [COMMITTED_ROUND]

# Or construct, sign, and submit in one line
$ goal clerk send --from=<my-account> --to=GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A --fee=1000 --amount=1000000 --note="Hello World"
Sent 1000000 MicroAlgos from account [ADDRESS] to address GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A, transaction ID: [TXID]. Fee set to 1000
Transaction [TXID] still pending as of round [LAST_ROUND]
Transaction [TXID] committed in round [COMMITTED_ROUND]
```


# Wait for confirmation

Successfully submitting your transaction to the network does not necessarily mean the network confirmed it. Always check that the network confirmed your transaction within a block before proceeding. 

!!! info
    On Algorand, transactions are final as soon as they are incorporated into a block and blocks are produced, on average, every 5 seconds. This means that transactions are confirmed, on average, in **5 seconds**! Read more about the [Algorand's Consensus Protocol](../algorand_consensus.md) and how it achieves such high confirmation speeds and immediate transaction finality.

```javascript tab="JavaScript"
const waitForConfirmation = async function (algodclient, txId) {
    let status = (await algodclient.status().do());
    let lastRound = status["last-round"];
    while (true) {
        const pendingInfo = await algodclient.pendingTransactionInformation(txId).do();
        if (pendingInfo["confirmed-round"] !== null && pendingInfo["confirmed-round"] > 0) {
            //Got the completed Transaction
            console.log("Transaction " + txId + " confirmed in round " + pendingInfo["confirmed-round"]);
            break;
        }
        lastRound++;
        await algodclient.statusAfterBlock(lastRound).do();
    }
};
```

```python tab="Python"
def wait_for_confirmation(client, txid):
	"""
	Utility function to wait until the transaction is
	confirmed before proceeding.
	"""
	last_round = client.status().get('last-round')
	txinfo = client.pending_transaction_info(txid)
	while not (txinfo.get('confirmed-round') and txinfo.get('confirmed-round') > 0):
		print("Waiting for confirmation")
		last_round += 1
		client.status_after_block(last_round)
		txinfo = client.pending_transaction_info(txid)
	print("Transaction {} confirmed in round {}.".format(txid, txinfo.get('confirmed-round')))
	return txinfo
```

```java tab="Java"
// utility function to wait on a transaction to be confirmed    
public void waitForConfirmation( String txID ) throws Exception{
    if( client == null ) this.client = connectToNetwork();
    Long lastRound = client.GetStatus().execute().body().lastRound;
    while(true) {
        try {
            //Check the pending tranactions
            Response<PendingTransactionResponse> pendingInfo = client.PendingTransactionInformation(txID).execute();
            if (pendingInfo.body().confirmedRound != null && pendingInfo.body().confirmedRound > 0) {
                //Got the completed Transaction
                System.out.println("Transaction " + txID + " confirmed in round " + pendingInfo.body().confirmedRound);
                break;
            } 
            lastRound++;
            client.WaitForBlock(lastRound).execute();
        } catch (Exception e) {
            throw( e );
        }
    }
}
```

```go tab="Go"
// Function that waits for a given txId to be confirmed by the network
func waitForConfirmation(txID string, client *algod.Client) {
	status, err := client.Status().Do(context.Background())
	if err != nil {
		fmt.Printf("error getting algod status: %s\n", err)
		return
	}
	lastRound := status.LastRound
	for {
		pt, _, err := client.PendingTransactionInformation(txID).Do(context.Background())
		if err != nil {
			fmt.Printf("error getting pending transaction: %s\n", err)
			return
		}
		if pt.ConfirmedRound > 0 {
			fmt.Printf("Transaction "+txID+" confirmed in round %d\n", pt.ConfirmedRound)
			break
		}
		fmt.Printf("waiting for confirmation\n")
		lastRound++
		status, err = client.StatusAfterBlock(lastRound).Do(context.Background())
	}
}
```

```bash tab="cURL"
curl -i -X GET \
   -H "X-Algo-API-Token:<algod-token> \
 'http://<algod-address>:<algod-port>/v1/transactions/pending/<txid>'
```

```bash tab="goal"
$ goal clerk rawsend --filename="hello-world.stxn"
Sent 1000000 MicroAlgos from account [ADDRESS] to address GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A, transaction ID: [TXID]. Fee set to 1000
Transaction [TXID] still pending as of round [LAST_ROUND]
Transaction [TXID] committed in round [COMMITTED_ROUND]

# Or construct, sign, and submit in one line
$ goal clerk send --from=<my-account> --to=GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A --fee=1000 --amount=1000000 --note="Hello World"
Sent 1000000 MicroAlgos from account [ADDRESS] to address GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A, transaction ID: [TXID]. Fee set to 1000
Transaction [TXID] still pending as of round [LAST_ROUND]
Transaction [TXID] committed in round [COMMITTED_ROUND]
```

# Read the transaction from the blockchain

Read your transaction back from the blockchain. 

!!! info
    Although you can read any transaction on the blockchain, only archival nodes store the whole history. By default, most nodes store only the last 1000 rounds and the APIs return errors when calling for information from earlier rounds. If you need to access data further back, make sure your algod client is connected to an archival, indexer node. Read more about node configurations in the Network Participation Guide or reach out to your service provider to understand how their node is configured. 

```javascript tab="JavaScript"
let confirmedTxn = await algodClient.pendingTransactionInformation(txId).do();
console.log("Transaction information: %o", confirmedTxn.txn.txn);
console.log("Decoded note: %s", algosdk.decodeObj(confirmedTxn.txn.txn.note));
```

```python tab="Python"
confirmed_txn = algod_client.pending_transaction_info(txid)
print("Transaction information: {}".format(json.dumps(confirmed_txn, indent=4)))
print("Decoded note: {}".format(base64.b64decode(confirmed_txn["txn"]["txn"]["note"]).decode()))
```

```java tab="Java"
PendingTransactionResponse pTrx = client.PendingTransactionInformation(id).execute().body();
System.out.println("Transaction information (with notes): " + pTrx.toString());
System.out.println("Decoded note: " + new String(pTrx.txn.tx.note));
```

```go tab="Go"
confirmedTxn, stxn, err := algodClient.PendingTransactionInformation(txID).Do(context.Background())
if err != nil {
    fmt.Printf("Error retrieving transaction %s\n", txID)
    return
}
txnJSON, err := json.MarshalIndent(confirmedTxn.Transaction.Txn, "", "\t")
if err != nil {
    fmt.Printf("Can not marshall txn data: %s\n", err)
}
fmt.Printf("Transaction information: %s\n", txnJSON)
fmt.Printf("Decoded note: %s\n", string(stxn.Txn.Note))
```

```bash tab="cURL"
curl -i -X GET \
   -H "X-Algo-API-Token:<algod-token> \
 'http://<algod-address>:<port>/v1/account/<my-address>/transaction/<txid>'
```

Notice above the pattern of constructing a transaction, authorizing it, submitting it to the network, and confirming its inclusion in a block. This is a framework to familiarize yourself with as it appears often in blockchain-related development.

!!! info
    Example code snippets were provided throughout this page. Full running code examples for each SDK are available within the GitHub repo at [/examples/start_building](https://github.com/algorand/docs/tree/master/examples/start_building) and for [download](https://github.com/algorand/docs/blob/master/examples/start_building/start_building.zip?raw=true) (.zip).

