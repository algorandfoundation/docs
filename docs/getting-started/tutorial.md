title: 3. Your First Transaction

PAGE STATUS: IN PROGRESS

After you successfully connect to **algod** using your preferred SDK, explore the methods available to read from and write to the blockchain. Remember that writing to the Algorand blockchain is simply sending a transaction to the network that is later confirmed within a block. 

Follow the guide below to send your first transaction on Algorand and familiarize yourself with some of the core functions of the SDKs. Examples of `goal` commands and REST API calls are included when they are the same or similar, allowing you to cross-verify and gain fluency across all available tools and platforms. 

Code snippets are abbreviated for conciseness and clarity. See the full code example for each SDK at the bottom of this guide.

# Create an account
In order to send a transaction, you first need an account<LINK TO GLOSSARY> on Algorand. Create an account by generating an Algorand public/private key pair and then funding the public address with Algos on your chosen network. 

!!! info
	The terms **account**, **public key**, and **address** are used interchangeably in certain contexts, but they have slightly different meanings. Read more about these differences on the Account Feature Guide<LINK> and within their respective definitions on the Glossary<LINK> page.


## Generate a public/private key pair

```javascript tab="JavaScript"
const algosdk = require('algosdk');

function generateAlgorandKeyPair() {
	var account = algosdk.generateAccount();
	var passphrase = algosdk.secretKeyToMnemonic(account.sk);
	console.log( "My address: " + account.addr );
	console.log( "My passphrase: " + passphrase );
}
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

public class GenerateAlgorandKeyPair {
	public static void main(String args[]) {
		Account myAccount = new Account();
        System.out.println("My Address: " + myAccount.getAddress());
		System.out.println("My Passphrase: " + myAccount.toMnemonic());
	}
}
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

$ goal account export -a address<PLACEHOLDER>
Exported key for account [ADDRESS]: [PASSPHRASE]
```

```bash tab="algokey"
$ algokey generate
Private key mnemonic: [PASSPHRASE]
Public key: [ADDRESS]
```

## Add funds
For TestNet and BetaNet, copy and paste the public portion of your key pair in the corresponding faucet prompt and click "Submit". A `200` response means the transaction went through and your balance increased by 100,000,000 microAlgos (i.e. 100 Algos).

!!! info
	Amounts are returned in microAlgos - the base unit for Algos. 1 Algo equals 1,000,000 microAlgos. <add micro explanation>

## Check your balance

Check your balance to confirm the added funds.

```javascript tab="JavaScript"
...
	const passphrase = "your-25-word-mnemonic<PLACEHOLDER>";
	let myAccount = algosdk.mnemonicToSecretKey(passphrase)
	console.log("My address: %s", myAccount.addr)

	let accountInfo = await algodClient.accountInformation(myAccount.addr);
    console.log("Account balance: %d microAlgos", accountInfo.amount)
...
```

```python tab="Python"
...
	passphrase = 25-word-mnemonic<PLACEHOLDER>
	private_key = mnemonic.to_private_key(passphrase)
	my_address = mnemonic.to_public_key(passphrase)
	print("My address: {}".format(my_address))

	account_info = algod_client.account_info(my_address)
	print("Account balance: {} microAlgos".format(account_info.get('amount')))
...
```

```java tab="Java"
...
    final String PASSPHRASE = "25-word-mnemonic<PLACEHOLDER>";
    com.algorand.algosdk.account.Account myAccount = new Account(PASSPHRASE);
    String myAddress = myAccount.getAddress().toString();
    System.out.println("My Address: " + myAddress);

    Account accountInfo = algodApiInstance.accountInformation(myAddress);
    System.out.println(String.format("Account Balance: %d microAlgos", accountInfo.getAmount()));
...
```

```go tab="Go"
...
	passphrase := 25-word-mnemonic<PLACEHOLDER>
	privateKey, err := mnemonic.ToPrivateKey(passphrase)
	if err != nil {
		fmt.Printf("Issue with mnemonic conversion: %s\n", err)
	}

	var myAddress types.Address
	publicKey := privateKey.Public()
	cpk := publicKey.(ed25519.PublicKey)
	copy(myAddress[:], cpk[:])
	fmt.Printf("My address: %s\n", myAddress.String())

	accountInfo, err := algodClient.AccountInformation(myAddress.String())
	if err != nil {
		fmt.Printf("Error getting account info: %s\n", err)
	}
    fmt.Printf("Account balance: %d microAlgos\n", accountInfo.Amount)
...
```

```bash tab="cURL"
curl -i -X GET \
   -H "X-Algo-API-Token:algod-token<PLACEHOLDER>" \
 'http://algod-address<PLACEHOLDER>:algod-port<PLACEHOLDER>/v1/account/address<placeholder>'
```

```bash tab="goal"
$ goal account balance -a my-address<PLACEHOLDER>
[AMOUNT] microAlgos
```

_Read all about Accounts, Keys, and Wallets in the Accounts Feature Guide_
# Construct the transaction

Create a transaction to send 1 Algo from your account to the TestNet faucet address (`GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A`) with the note "Hello World".

Transactions require a certain minimum set of parameters to be valid. Mandatory fields include the **round validity range**, the **fee**, and the **genesis hash** for the network the transaction is valid for. Read all about Transaction types, fields, and configurations in the Transactions Feature Guide<LINK>. For now, construct a payment transaction as follows. Use the _suggested parameters_ methods to initialize network-related 
fields. 

```javascript tab="JavaScript"
...
	let params = await algodClient.getTransactionParams();
	let note = algosdk.encodeObj("Hello World");
	let txn = {
		"from": myAccount.addr,
		"to": receiver,
		"fee": params.minFee,
		"amount": 1000000,
		"firstRound": params.lastRound,
		"lastRound": params.lastRound + 1000,
		"note": note,
		"genesisID": params.genesisID,
		"genesisHash": params.genesishashb64
	};
...
```

```python tab="Python"
...
	params = algod_client.suggested_params()
	note = "Hello World".encode()
	receiver = "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A"

	data = {
		"sender": my_address,
		"receiver": receiver,
		"fee": params.get('minFee'),
		"flat_fee": True,
		"amt": 1000000,
		"first": params.get('lastRound'),
		"last": params.get('lastRound') + 1000,
		"note": note,
		"gen": params.get('genesisID'),
		"gh": params.get('genesishashb64')
	}
	
	txn = transaction.PaymentTxn(**data)
...
```

```java tab="Java"
...
final RECEIVER = "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A";
...
BigInteger fee;
String genesisID;
Digest genesisHash;
BigInteger firstValidRound;
fee = BigInteger.valueOf(1000);
try {
    TransactionParams params = algodApiInstance.transactionParams();
    genesisHash = new Digest(params.getGenesishashb64());
    genesisID = params.getGenesisID();
    System.out.println("Minimum Fee: " + fee);
    firstValidRound = params.getLastRound();
    System.out.println("Current Round: " + firstValidRound);
} catch (ApiException e) {
    throw new RuntimeException("Could not get params", e);
}
BigInteger amount = BigInteger.valueOf(1000000); // microAlgos
BigInteger lastValidRound = firstValidRound.add(BigInteger.valueOf(1000)); // 1000 is the max tx window
String note = "Hello World";
Transaction txn = new Transaction(myAccount.getAddress(), fee, firstValidRound,
        lastValidRound, note.getBytes(), amount, new Address(RECEIVER),
        genesisID, genesisHash);
...
```

```go tab="Go"
...
	txParams, err := algodClient.SuggestedParams()
	if err != nil {
		fmt.Printf("Error getting suggested tx params: %s\n", err)
		return
    }
    
	fromAddr := myAddress
	toAddr := "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A"
	var amount uint64 = 10000000
	var minFee uint64 = 1000
	note := []byte("Hello World")
	genID := txParams.GenesisID
	genHash := txParams.GenesisHash
	firstValidRound := txParams.LastRound
	lastValidRound := firstValidRound + 1000

	txn, err := transaction.MakePaymentTxnWithFlatFee(fromAddr.String(), toAddr, minFee, amount, firstValidRound, lastValidRound, note, "", genID, genHash)
	if err != nil {
		fmt.Printf("Error creating transaction: %s\n", err)
		return
    }
...
```

```bash tab="goal"
$ goal clerk send --from=my-account<PLACEHOLDER> --to=GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A --fee=1000 --amount=1000000 --note="Hello World" --out="hello-world.txn"
```

!!! info
    Some of the SDKs provide wrapper functions for creating certain types of transactions, like `makePaymentTxn` in Go. 

# Sign the transaction
Sign the transaction with your private key. This creates a new signed transaction object in the SDKs. Retrieve the transaction ID of the signed transaction.

```javascript tab="JavaScript"
...
	let signedTxn = algosdk.signTransaction(txn, myAccount.sk);
	let txId = signedTxn.txID;
    console.log("Signed transaction with txId: %s", txId);
...
```

```python tab="Python"
...
	signed_txn = txn.sign(private_key)
	txid = signed_txn.transaction.get_txid()
	print("Signed transaction with txID: {}".format(txid))
...
```

```java tab="Java"
...
    SignedTransaction signedTx = myAccount.signTransaction(txn);
    System.out.println("Signed transaction with txId: " + signedTx.transactionID);
...
```

```go tab="Go"
...
	txId, bytes, err := crypto.SignTransaction(privateKey, txn)
	if err != nil {
		fmt.Printf("Failed to sign transaction: %s\n", err)
		return
	}
    fmt.Printf("Signed txid: %s\n", txId)
...
```

```bash tab="goal"
$ goal clerk sign --infile="hello-world.txn" --outfile="hello-world.stxn"
```

```bash tab="algokey"
```

# Submit the transaction
Send the signed transaction to the network with your algod client. 

```javascript tab="JavaScript"
...
    await postAlgodClient.sendRawTransaction(signedTxn.blob)
...
```

```python tab="Python"
...
	algod_client.send_transaction(signed_txn)
...
```

```java tab="Java"
...
try {
    byte[] encodedTxBytes = Encoder.encodeToMsgPack(signedTx);
    TransactionID id = algodApiInstance.rawTransaction(encodedTxBytes);
    System.out.println("Successfully sent tx with ID: " + id);
} catch (ApiException e) {
    System.err.println("Exception when calling algod#rawTransaction: " + e.getResponseBody());
}
...
```

```go tab="Go"
...
	sendResponse, err := algodClient.SendRawTransaction(bytes)
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
    fmt.Printf("Submitted transaction %s\n", sendResponse.TxID)
...
```

```bash tab="cURL"
curl -i -X POST \
   -H "X-Algo-API-Token:algod-token<PLACEHOLDER>" \
   -H "Content-Type:application/x-binary" \
   -T "hello-world.stxn" \
 'http://algod-address<PLACEHOLDER>:algod-port<PLACEHOLDER>/v1/transactions'
```

```bash tab="goal"
$ goal clerk rawsend --filename="hello-world.stxn"
Sent 1000000 MicroAlgos from account [ADDRESS] to address GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A, transaction ID: [TXID]. Fee set to 1000
Transaction [TXID] still pending as of round [LAST_ROUND]
Transaction [TXID] committed in round [COMMITTED_ROUND]

# Or construct, sign, and submit in one line
$ goal clerk send --from=my-account<PLACEHOLDER> --to=GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A --fee=1000 --amount=1000000 --note="Hello World"
Sent 1000000 MicroAlgos from account [ADDRESS] to address GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A, transaction ID: [TXID]. Fee set to 1000
Transaction [TXID] still pending as of round [LAST_ROUND]
Transaction [TXID] committed in round [COMMITTED_ROUND]
```

!!! Info
    If you are using a third-party service, it may require you to specify a `Content-Type` header when you send transactions to the network. Set the `Content-Type` to `application/x-binary` and add it as a header to the algod client or the specific request that sends the transaction.
    
    ```JavaScript tab=
	const algosdk = require("algosdk");

	async function gettingStartedExample() {

		const server = algod-address<PLACEHOLDER>;
		const port = "";
		const token = {
			'X-API-Key': service-api-key<PLACEHOLDER>,
            'Content-Type': 'appliation/x-binary'
		};

		let postAlgodClient = new algosdk.Algod(token, server, port);
		...
	}

	```

	```Python tab=
    ...
	algod_client.send_transaction(signed_txn, headers={'content-type': 'application/x-binary'})
    ...
	```

	```Go tab=
    ...
        txHeaders := append([]*algod.Header{}, &algod.Header{"Content-Type", "application/x-binary"})
        sendResponse, err := algodClient.SendRawTransaction(bytes, txHeaders...)
    ...
	```

# Wait for confirmation

Successfully submitting your transaction to the network does not necessarily mean the network confirmed it. Always check that the network confirmed your transaction within a block before proceeding. 

!!! info
    On Algorand, transactions are final as soon as they are incorporated into a block and blocks are produced, on average, every 5 seconds. This means that transactions are confirmed, on average, in **5 seconds**! Read more about the [Algorand's Consensus Protocol](../learn/algorand_consensus.md) and how it achieves such high confirmation speeds and immediate transaction finality in the [Learn](../learn/algorand_consensus.md) section.

```javascript tab="JavaScript"
...
let timerId = 
    setInterval(async function() {
        let txInfo = await algodClient.pendingTransactionInformation(txId);
        if (txInfo.round != null && txInfo.round > 0) {
            console.log("Transaction %s confirmed in round %d", txId, txInfo.round);
            ...
        } else {
                console.log("Waiting for confirmation...");
        }
    }, 1000)
...
```

```python tab="Python"
	while True:
		txinfo = algod_client.pending_transaction_info(txid)
		if txinfo.get('round') and txinfo.get('round') > 0:
			print("Transaction {} confirmed in round {}.".format(txid, txinfo.get('round')))
			break
		else:
			print("Waiting for confirmation...")
		time.sleep(2)
```

```java tab="Java"
...
 while(true) {
    try {
        //Check the pending transactions
        com.algorand.algosdk.algod.client.model.Transaction txInfo =
                algodApiInstance.pendingTransactionInformation(signedTx.transactionID);
        if (txInfo.getRound() != null && txInfo.getRound().longValue() > 0) {
                    System.out.println("Transaction " + txInfo.getTx()
                            + " confirmed in round " + txInfo.getRound().longValue());
                    break;
                } else {
                    System.out.println("Waiting for confirmation... (pool error, if any:)" + txInfo.getPoolerror());
                }
            } catch (ApiException e) {
                System.err.println("Exception when calling algod#pendingTxInformation: " + e.getMessage());
                break;
            }
        }

...
```

```go tab="Go"
...
	for {
		pendingTxnInfo, err := algodClient.PendingTransactionInformation(txId)
		if err != nil {
			fmt.Printf("Error with pending txn: %s\n", err)
			return
		}
		if pendingTxnInfo.ConfirmedRound > 0 {
			fmt.Printf("Transaction "+pendingTxnInfo.TxID+" confirmed in round %d\n", pendingTxnInfo.ConfirmedRound)
			break
		}
		fmt.Printf("waiting for confirmation...\n")
		time.Sleep(2 * time.Second)
    }
...
```

```bash tab="cURL"
curl -i -X GET \
   -H "X-Algo-API-Token:algod-token<PLACEHOLDER>" \
 'http://algod-address<PLACEHOLDER>:algod-port<PLACEHOLDER>/v1/transactions/pending/txid<PLACEHOLDER>'
```

```bash tab="goal"
$ goal clerk rawsend --filename="hello-world.stxn"
Sent 1000000 MicroAlgos from account [ADDRESS] to address GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A, transaction ID: [TXID]. Fee set to 1000
Transaction [TXID] still pending as of round [LAST_ROUND]
Transaction [TXID] committed in round [COMMITTED_ROUND]

# Or construct, sign, and submit in one line
$ goal clerk send --from=my-account<PLACEHOLDER> --to=GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A --fee=1000 --amount=1000000 --note="Hello World"
Sent 1000000 MicroAlgos from account [ADDRESS] to address GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A, transaction ID: [TXID]. Fee set to 1000
Transaction [TXID] still pending as of round [LAST_ROUND]
Transaction [TXID] committed in round [COMMITTED_ROUND]
```

# Read the transaction from the blockchain

Read your transaction back from the blockchain. 

!!! info
    Although you can read any transaction on the blockchain, only archival nodes store the whole history. By default, most nodes store only the last 1000 rounds and the APIs return errors when calling for information from earlier rounds. If you need to access data further back, make sure your algod client is connected to an archival, indexer node. Read more about node configurations in the Network Participation Guide or reach out to your service provider to understand how their node is configured. 

```javascript tab="JavaScript"
...
try {
	let confirmedTxn = await algodClient.transactionInformation(myAccount.addr, txId);
	console.log("Transaction information: %o", confirmedTxn);
	console.log("Decoded note: %s", algosdk.decodeObj(confirmedTxn.note));
} catch(e) {
	console.log(e.response.text);
}
...
```

```python tab="Python"
...
	confirmed_txn = algod_client.transaction_info(my_address, txid)
	print("Transaction information: {}".format(json.dumps(confirmed_txn, indent=4)))
	print("Decoded note: {}".format(base64.b64decode(confirmed_txn.get('noteb64')).decode()))
...
```

```java tab="Java"
...
    try {
        com.algorand.algosdk.algod.client.model.Transaction confirmedTxn =
                algodApiInstance.transactionInformation(RECEIVER, signedTxn.transactionID);
        System.out.println("Transaction information (with notes): " + confirmedTxn.toString());
        System.out.println("Decoded note: " + new String(confirmedTxn.getNoteb64()));
    } catch (ApiException e) {
        System.err.println("Exception when calling algod#transactionInformation: " + e.getCode());
    }
...
```

```go tab="Go"
...
	confirmedTxn, err := algodClient.TransactionInformation(myAddress.String(), txId)
	if err != nil {
		fmt.Printf("Error retrieving transaction %s\n", txId)
		return
	}
	txnJSON, err := json.MarshalIndent(confirmedTxn, "", "\t")
	if err != nil {
		fmt.Printf("Can not marshall txn data: %s\n", err)
	}
	fmt.Printf("Transaction information: %s\n", txnJSON)
    fmt.Printf("Decoded note: %s\n", string(confirmedTxn.Note))
...
```

```bash tab="cURL"
curl -i -X GET \
   -H "X-Algo-API-Token:algod-token<PLACEHOLDER>" \
 'http://algod-address<PLACEHOLDER>:port<PLACEHOLDER>/v1/account/my-address<PLACEHOLDER>/transaction/txid<PLACEHOLDER>'
```

Notice above the pattern of constructing a transaction, authenticating it, submitting it to the network, and confirming its inclusion in a block. This is a framework to familiarize yourself with as it appears often in blockchain-related development.

??? example "Complete Example - Sending a Payment Transaction"
    
    ```javascript tab="JavaScript"
    const algosdk = require('algosdk');

    async function gettingStartedExample() {

        const token = algod-token<PLACEHOLDER>;
        const server = algod-address<PLACEHOLDER>;
        const port = port<PLACEHOLDER>;

        let algodClient = new algosdk.Algod(token, server, port);

        const receiver = "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A";
        const passphrase = "your-25-word-mnemonic<PLACEHOLDER>";
        
        let myAccount = algosdk.mnemonicToSecretKey(passphrase)
        console.log("My address: %s", myAccount.addr)

        let accountInfo = await algodClient.accountInformation(myAccount.addr);
        console.log("Account balance: %d microAlgos", accountInfo.amount)

        let params = await algodClient.getTransactionParams();
        console.log(params);

        let note = algosdk.encodeObj("Hello World");

        let txn = {
            "from": myAccount.addr,
            "to": receiver,
            "fee": params.minFee,
            "flatFee": true,
            "amount": 1000000,
            "firstRound": params.lastRound,
            "lastRound": params.lastRound + 1000,
            "note": note,
            "genesisID": params.genesisID,
            "genesisHash": params.genesishashb64
        };

        let signedTxn = algosdk.signTransaction(txn, myAccount.sk);
        let txId = signedTxn.txID;
        console.log("Signed transaction with txID: %s", txId);

        await algodClient.sendRawTransaction(signedTxn.blob)

        // Wait for confirmation
        let timerId = 
            setInterval(async function() {
                let txInfo = await algodClient.pendingTransactionInformation(txId);
                if (txInfo.round != null && txInfo.round > 0) {
                    console.log("Transaction %s confirmed in round %d", txId, txInfo.round);
                    try {
                        // Read the confirmed transaction
                        let confirmedTxn = await algodClient.transactionInformation(myAccount.addr, txId);
                        console.log("Transaction information: %o", confirmedTxn);
                        console.log("Decoded note: %s", algosdk.decodeObj(confirmedTxn.note));
                    } catch(e) {
                        console.log(e.response.text);
                    }
                    clearInterval(timerId);
                } else {
                        console.log("Waiting for confirmation...");
                }
            }, 1000)
    };
    ```

    ```python tab="Python"
    import json
    import time
    import base64
    from algosdk import algod
    from algosdk import mnemonic
    from algosdk import transaction

    def main():
        algod_address = algod-address<PLACEHOLDER>
        algod_token = algod-token<PLACEHOLDER>
        algod_client = algod.AlgodClient(algod_token, algod_address)
        
        passphrase = 25-word-mnemonic<PLACEHOLDER>
        private_key = mnemonic.to_private_key(passphrase)
        my_address = mnemonic.to_public_key(passphrase)
        print("My address: {}".format(my_address))

        account_info = algod_client.account_info(my_address)
        print("Account balance: {} microAlgos".format(account_info.get('amount')))

        params = algod_client.suggested_params()
        note = "Hello World".encode()
        receiver = "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A"

        data = {
            "sender": my_address,
            "receiver": receiver,
            "fee": params.get('minFee'),
            "flat_fee": True,
            "amt": 1000000,
            "first": params.get('lastRound'),
            "last": params.get('lastRound') + 1000,
            "note": note,
            "gen": params.get('genesisID'),
            "gh": params.get('genesishashb64')
        }
        
        txn = transaction.PaymentTxn(**data)
        signed_txn = txn.sign(private_key)
        txid = signed_txn.transaction.get_txid()
        print("Signed transaction with txID: {}".format(txid))

        algod_client.send_transaction(signed_txn)
        
        while True:
            txinfo = algod_client.pending_transaction_info(txid)
            if txinfo.get('round') and txinfo.get('round') > 0:
                print("Transaction {} confirmed in round {}.".format(txid, txinfo.get('round')))
                break
            else:
                print("Waiting for confirmation...")
            time.sleep(2)
        
        # Read the transction
	    try:
            confirmed_txn = algod_client.transaction_info(my_address, txdfid)
        except Exception as err:
            print(err)
        print("Transaction information: {}".format(json.dumps(confirmed_txn, indent=4)))
        print("Decoded note: {}".format(base64.b64decode(confirmed_txn.get('noteb64')).decode()))
    ```

    ```java tab="Java"
    package example;

    import java.math.BigInteger;

    import java.util.concurrent.TimeUnit;

    import com.algorand.algosdk.account.Account;
    import com.algorand.algosdk.algod.client.AlgodClient;
    import com.algorand.algosdk.algod.client.ApiException;
    import com.algorand.algosdk.algod.client.api.AlgodApi;
    import com.algorand.algosdk.algod.client.model.*;
    import com.algorand.algosdk.crypto.Address;
    import com.algorand.algosdk.crypto.Digest;
    import com.algorand.algosdk.transaction.Transaction;
    import com.algorand.algosdk.transaction.SignedTransaction;
    import com.algorand.algosdk.util.Encoder;

    public class MyApp {

        public static void main(String args[]) throws Exception {

            // Initialize an algod client
            final String ALGOD_API_ADDR = "algod-address"<PLACEHOLDER>;
            final String ALGOD_API_TOKEN = "algod-token"<PLACEHOLDER>;

            //Create an instance of the algod API client
            AlgodClient client = (AlgodClient) new AlgodClient().setBasePath(ALGOD_API_ADDR);
            ApiKeyAuth api_key = (ApiKeyAuth) client.getAuthentication("api_key");
            api_key.setApiKey(ALGOD_API_TOKEN);
            AlgodApi algodApiInstance = new AlgodApi(client);

            // Import your private key mnemonic and address
            final String PASSPHRASE = "25-word-mnemonic<PLACEHOLDER>";
            com.algorand.algosdk.account.Account myAccount = new Account(PASSPHRASE);
            System.out.println("My Address: " + myAccount.getAddress());

            String myAddress = myAccount.getAddress().toString();
            com.algorand.algosdk.algod.client.model.Account accountInfo = 
                algodApiInstance.accountInformation(myAddress);
            System.out.println(String.format("Account Balance: %d microAlgos", accountInfo.getAmount()));

            // Construct the transaction
            final String RECEIVER = "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A";
            BigInteger fee;
            String genesisID;
            Digest genesisHash;
            BigInteger firstValidRound;
            fee = BigInteger.valueOf(1000);
            try {
                TransactionParams params = algodApiInstance.transactionParams();
                genesisHash = new Digest(params.getGenesishashb64());
                genesisID = params.getGenesisID();
                System.out.println("Minimum Fee: " + fee);
                firstValidRound = params.getLastRound();
                System.out.println("Current Round: " + firstValidRound);
            } catch (ApiException e) {
                throw new RuntimeException("Could not get params", e);
            }
            BigInteger amount = BigInteger.valueOf(1000000); // microAlgos
            BigInteger lastValidRound = firstValidRound.add(BigInteger.valueOf(1000)); // 1000 is the max tx window
            String note = "Hello World";


            Transaction txn = new Transaction(myAccount.getAddress(), fee, firstValidRound,
                    lastValidRound, note.getBytes(), amount, new Address(RECEIVER),
                    genesisID, genesisHash);

            // Sign the transaction
            SignedTransaction signedTxn = myAccount.signTransaction(txn);
            System.out.println("Signed transaction with txid: " + signedTxn.transactionID);

            // Submit the transaction to the network
            try {
                byte[] encodedTxBytes = Encoder.encodeToMsgPack(signedTxn);
                TransactionID id = algodApiInstance.rawTransaction(encodedTxBytes);
                System.out.println("Successfully sent tx with ID: " + id);
            } catch (ApiException e) {
                System.err.println("Exception when calling algod#rawTransaction: " + e.getResponseBody());
            }

            // Wait for transaction confirmation
            while(true) {
                try {
                    //Check the pending transactions
                    com.algorand.algosdk.algod.client.model.Transaction txInfo =
                            algodApiInstance.pendingTransactionInformation(signedTxn.transactionID);
                    if (txInfo.getRound() != null && txInfo.getRound().longValue() > 0) {
                        System.out.println("Transaction " + txInfo.getTx()
                                + " confirmed in round " + txInfo.getRound().longValue());
                        break;
                    } else {
                        System.out.println("Waiting for confirmation... (pool error, if any:)" + txInfo.getPoolerror());
                    }
                } catch (ApiException e) {
                    System.err.println("Exception when calling algod#pendingTxInformation: " + e.getMessage());
                    break;
                }
                TimeUnit.SECONDS.sleep(1);
            }

            //Read the transaction from the blockchain
            try {
                com.algorand.algosdk.algod.client.model.Transaction confirmedTxn =
                        algodApiInstance.transactionInformation(RECEIVER, signedTxn.transactionID);
                System.out.println("Transaction information (with notes): " + confirmedTxn.toString());
                System.out.println("Decoded note: " + new String(confirmedTxn.getNoteb64()));
            } catch (ApiException e) {
                System.err.println("Exception when calling algod#transactionInformation: " + e.getCode());
            }
        }
    }
    ```

    ```go tab="Go"
    package main

    import (
        "crypto/ed25519"
        "encoding/json"
        "fmt"
        "time"

        "github.com/algorand/go-algorand-sdk/client/algod"
        "github.com/algorand/go-algorand-sdk/crypto"
        "github.com/algorand/go-algorand-sdk/mnemonic"
        "github.com/algorand/go-algorand-sdk/transaction"
        "github.com/algorand/go-algorand-sdk/types"
    )

    const algodAddress = algod-address<PLACEHOLDER>
    const algodToken = algod-tokenn<PLACEHOLDER>

    func main() {

        algodClient, err := algod.MakeClient(algodAddress, algodToken)
        if err != nil {
            return
        }

        passphrase := 25-word-mnemonic<PLACEHOLDER>
        privateKey, err := mnemonic.ToPrivateKey(passphrase)
        if err != nil {
            fmt.Printf("Issue with mnemonic conversion: %s\n", err)
        }

        var myAddress types.Address
        publicKey := privateKey.Public()
        cpk := publicKey.(ed25519.PublicKey)
        copy(myAddress[:], cpk[:])
        fmt.Printf("My address: %s\n", myAddress.String())

        // Check account balance
        accountInfo, err := algodClient.AccountInformation(myAddress.String())
        if err != nil {
            fmt.Printf("Error getting account info: %s\n", err)
        }
        fmt.Printf("Account balance: %d microAlgos\n", accountInfo.Amount)

        // Construct the transaction
        txParams, err := algodClient.SuggestedParams()
        if err != nil {
            fmt.Printf("Error getting suggested tx params: %s\n", err)
            return
        }

        fromAddr := myAddress
        toAddr := "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A"
        var amount uint64 = 10000000
        var minFee uint64 = 1000
        note := []byte("Hello World")
        genID := txParams.GenesisID
        genHash := txParams.GenesisHash
        firstValidRound := txParams.LastRound
        lastValidRound := firstValidRound + 1000

        txn, err := transaction.MakePaymentTxnWithFlatFee(fromAddr.String(), toAddr, minFee, amount, firstValidRound, lastValidRound, note, "", genID, genHash)
        if err != nil {
            fmt.Printf("Error creating transaction: %s\n", err)
            return
        }

        // Sign the transaction
        txId, bytes, err := crypto.SignTransaction(privateKey, txn)
        if err != nil {
            fmt.Printf("Failed to sign transaction: %s\n", err)
            return
        }
        fmt.Printf("Signed txid: %s\n", txId)

        // Submit the transaction

        sendResponse, err := algodClient.SendRawTransaction(bytes)
        if err != nil {
            fmt.Printf("failed to send transaction: %s\n", err)
            return
        }
        fmt.Printf("Submitted transaction %s\n", sendResponse.TxID)

        // Wait for confirmation
        for {
            pendingTxnInfo, err := algodClient.PendingTransactionInformation(txId)
            if err != nil {
                fmt.Printf("Error with pending txn: %s\n", err)
                return
            }
            if pendingTxnInfo.ConfirmedRound > 0 {
                fmt.Printf("Transaction "+pendingTxnInfo.TxID+" confirmed in round %d\n", pendingTxnInfo.ConfirmedRound)
                break
            }
            fmt.Printf("waiting for confirmation...\n")
            time.Sleep(2 * time.Second)
        }

        // Read confirmed transaction from block
        confirmedTxn, err := algodClient.TransactionInformation(myAddress.String(), txId)
        if err != nil {
            fmt.Printf("Error retrieving transaction %s\n", txId)
            return
        }
        txnJSON, err := json.MarshalIndent(confirmedTxn, "", "\t")
        if err != nil {
            fmt.Printf("Can not marshall txn data: %s\n", err)
        }
        fmt.Printf("Transaction information: %s\n", txnJSON)
        fmt.Printf("Decoded note: %s\n", string(confirmedTxn.Note))
    }
    ```

    ```bash tab="goal"
    #!/usr/bin/bash

    $ goal account balance -a my-address<PLACEHOLDER>
    [AMOUNT] microAlgos

    $ goal clerk send --from=my-account<PLACEHOLDER> --to=GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A --fee=1000 --amount=1000000 --note="Hello World" --out="hello-world.txn"

    $ goal clerk sign --infile="hello-world.txn" --outfile="hello-world.stxn"

    $ goal clerk rawsend --filename="hello-world.stxn"
    Sent 1000000 MicroAlgos from account [ADDRESS] to address GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A, transaction ID: [TXID]. Fee set to 1000
    Transaction [TXID] still pending as of round [LAST_ROUND]
    Transaction [TXID] committed in round [COMMITTED_ROUND]

    # Or construct, sign, and submit in one line
     $ goal clerk send --from=my-account<PLACEHOLDER> --to=GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A --fee=1000 --amount=1000000 --note="Hello World"
     Sent 1000000 MicroAlgos from account [ADDRESS] to address GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A, transaction ID: [TXID]. Fee set to 1000
     Transaction [TXID] still pending as of round [LAST_ROUND]
     Transaction [TXID] committed in round [COMMITTED_ROUND]
    ``` 