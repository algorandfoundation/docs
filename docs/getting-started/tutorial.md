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
```

```python tab="Python"
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
```

```bash tab="goal"
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
	Amounts are returned in microAlgos - the base unit for Algos. 1 Algo equals 1,000,000 microAlgos. 

## Check your balance

Check your balance to confirm the added funds.

```javascript tab="JavaScript"
```

```python tab="Python"
```

```java tab="Java"
import com.algorand.algosdk.algod.client.model.Account;	
...
String myAddress = myAccount.getAddress().toString();
Account accountInfo = algodApiInstance.accountInformation(myAddress);
System.out.println("Account Balance: " + accountInfo.getAmount());
...
```

```bash tab="cURL"
curl -i -X GET \
   -H "X-Algo-API-Token:algod-token<PLACEHOLDER>" \
 'http://algod-address<PLACEHOLDER>:algod-port<PLACEHOLDER>/v1/account/address<placeholder>'
```

```bash tab="goal"
$ goal account balance -a my-address<PLACEHOLDER>
```

_Read all about Accounts, Keys, and Wallets in the Accounts Feature Guide_
# Construct the transaction

Create a transaction to send 1 Algo from your account to the TestNet faucet address (`GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A`) with the note "Hello World".

## Initialize transaction fields
Transactions require a minimum set of parameters to be valid. Mandatory fields include the **round validity range**, the **fee**, and the **genesis hash** for the network the transaction is valid for. Read all about Transaction types, fields, and configurations in the Transactions Feature Guide<LINK>. For now, construct a payment transaction as follows. Use the _suggested parameters_ methods to initialize network-related 
fields. 

```javascript tab="JavaScript"
```

```python tab="Python"
```

```java tab="Java"
...
final RECEIVER = "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A";
...
BigInteger fee;
String genesisID;
Digest genesisHash;
long firstValidRound;

try {
    TransactionParams params = algodApiInstance.transactionParams();
    genesisHash = new Digest(params.getGenesishashb64());
    genesisID = params.getGenesisID();
    System.out.println("Minimum fee: " + fee);
    firstValidRound = params.getLastRound().longValue();
    System.out.println("Current Round: " + firstValidRound);
} catch (ApiException e) {
    throw new RuntimeException("Could not get params", e);
}

fee = BigInteger.valueOf(1000);
long amount = 1000000; // microAlgos
long lastValidRound = firstValidRound + 1000; // 1000 is the max tx window
String note = "Hello World";

Transaction txn = new Transaction(myAccount.getAddress(), fee, BigInteger.valueOf(firstValidRound),
    	BigInteger.valueOf(lastValidRound), note.getBytes(), BigInteger.valueOf(amount), new Address(RECEIVER),
        genesisID, genesisHash);
...
```

```bash tab="goal"
$ goal clerk send --from=my-account<PLACEHOLDER> --to=GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A --fee=1000 --amount=1000000 --note="Hello World" --out="hello-world.txn"
```

# Sign the transaction

```javascript tab="JavaScript"
```

```python tab="Python"
```

```java tab="Java"
...
Transaction txn = new Transaction(myAccount.getAddress(), fee, BigInteger.valueOf(firstValidRound),
    	BigInteger.valueOf(lastValidRound), note.getBytes(), BigInteger.valueOf(amount), new Address(RECEIVER),
        genesisID, genesisHash);
SignedTransaction signedTxn = myAccount.signTransaction(txn);
System.out.println("Signed transaction with txid: " + signedTxn.transactionID);

try {
    byte[] encodedTxnBytes = Encoder.encodeToMsgPack(signedTxn);
    TransactionID id = algodApiInstance.rawTransaction(encodedTxnBytes);
    System.out.println("Successfully sent tx with ID: " + id);
} catch (ApiException e) {
    System.err.println("Exception when calling algod#rawTransaction: " + e.getResponseBody());
}
...
```

```bash tab="goal"
$ goal clerk sign --infile="hello-world.txn" --outfile="hello-world.stxn"
```

```bash tab="algokey"
```

# Submit the transaction

```javascript tab="JavaScript"
```

```python tab="Python"
```

```java tab="Java"
...
SignedTransaction signedTx = myAccount.signTransaction(txn);
System.out.println("Signed transaction with txid: " + signedTx.transactionID);
...
```

```bash tab="cURL"
curl -i -X POST \
   -H "X-Algo-API-Token:algod-token<PLACEHOLDER>" \
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

# Locate the transaction

Notice above the pattern of constructing a transaction, authenticating it, submitting it to the network, and confirming its inclusion in a block. This is a framework to familiarize yourself with as it appears often in blockchain-related development.