title: Authorization

In the [Transactions Section](../transactions/index.md), you learned how transactions are composed. In this section you will learn how to authorize them. 

Before a transaction is sent to the network, it must first be authorized by the [sender](../../reference/transactions.md#sender). Authorization occurs through the addition of a **signature** to the transaction object. Specifically, a transaction object, when signed, is wrapped in a [`SignedTxn`](../../reference/transactions.md#signed-transaction) object that includes the [transaction](../../reference/transactions.md#txn) and a type of [signature](../../reference/transactions.md#sig). 

There are three types of signatures:

- [Single Signatures](#single-signatures)
- [Multisignatures](#multisignatures)
- [Logic Signatures](#logic-signatures)



# Single Signatures
A single signature corresponds to a signature from the private key of an [Algorand public/private key pair](../accounts/index.md#keys-and-addresses).

This is an example of a transaction signed by an Algorand private key displayed with `goal clerk inspect` command:

```json
{
  "sig": "ynA5Hmq+qtMhRVx63pTO2RpDrYiY1wzF/9Rnnlms6NvEQ1ezJI/Ir9nPAT6+u+K8BQ32pplVrj5NTEMZQqy9Dw==",
  "txn": {
    "amt": 10000000,
    "fee": 1000,
    "fv": 4694301,
    "gen": "testnet-v1.0",
    "gh": "SGO1GKSzyE7IEPItTxCByw9x8FmnrCDexi9/cOUJOiI=",
    "lv": 4695301,
    "rcv": "QC7XT7QU7X6IHNRJZBR67RBMKCAPH67PCSX4LYH4QKVSQ7DQZ32PG5HSVQ",
    "snd": "EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4",
    "type": "pay"
  }
}
```
This transaction sends 1 Algo from `"EW64GC..."` to `"QC7XT7..."` on TestNet. The transaction was signed with the private key that corresponds to the `"snd"` address of `"EW64GC..."`. The base64 encoded signature is shown as the value of the [`"sig"`](../../reference/transactions.md#sig) field.

**Related How-To**

- [Sign a transaction with your private key](../../build-apps/hello_world.md#sign-the-transaction)

# Multisignatures

When the [sender](../../reference/transactions.md#sender) of a transaction is the address of a [multisignature account](../accounts/create.md#multisignature) then authorization requires a subset of signatures, _equal to or greater than the threshold value_, from the associated private keys of the addresses that multisignature account is composed of. See [Multisignature Accounts](../accounts/create.md#multisignature) for details on how to configure a multisignature account.

!!! important
	Upon signing, either the signing agent or the transaction needs to know the composition of the multisignature account, i.e. the ordered addresses, threshold, and version. 

Here is what the same transaction above would look like if sent from a 2/3 multisig account.

```json
{
  "msig": {
    "subsig": [
      {
        "pk": "SYGHTA2DR5DYFWJE6D4T34P4AWGCG7JTNMY4VI6EDUVRMX7NG4KTA2WMDA"
      },
      {
        "pk": "VBDMPQACQCH5M6SBXKQXRWQIL7QSR4FH2UI6EYI4RCJSB2T2ZYF2JDHZ2Q"
      },
      {
        "pk": "W3KONPXCGFNUGXGDCOCQYVD64KZOLUMHZ7BNM2ZBK5FSSARRDEXINLYHPI"
      }
    ],
    "thr": 2,
    "v": 1
  },
  "txn": {
    "amt": 10000000,
    "fee": 1000,
    "fv": 4694301,
    "gen": "testnet-v1.0",
    "gh": "SGO1GKSzyE7IEPItTxCByw9x8FmnrCDexi9/cOUJOiI=",
    "lv": 4695301,
    "rcv": "QC7XT7QU7X6IHNRJZBR67RBMKCAPH67PCSX4LYH4QKVSQ7DQZ32PG5HSVQ",
    "snd": "GQ3QPLJL4VKVGQCHPXT5UZTNZIJAGVJPXUHCJLRWQMFRVL4REVW7LJ3FGY",
    "type": "pay"
  }
}
```
The difference between this transaction and the one above is the form of its signature component. For multisignature accounts, an [`"msig"`](../../reference/transactions.md#msig) struct is added which contains the 3 public addresses (`"pk"`), the threshold value (`"thr"`) and the multisig version `"v"`. This transaction is still unsigned but the addition of the correct `"msig"` struct is confirmation that the transaction is "aware" of the fact that the sender is multisig and will have no trouble accepting sub-signatures from single keys even if the signing agent does not contain information about its multisignature properties.

!!! tip
	Adding the `"msig"` template to make the transaction "aware" of its multisig sender is highly recommended, particularly in cases where the transaction is signed by multiple parties or offline. Without it, the signing agent would need to have its own knowledge of the multisignature account. For example, `goal` can sign a multisig transaction that does not contain an `"msig"` templlate _if_ the multisig address was created within its wallet. On signing, it will add the `"msig"` template. 


Sub-signatures can be added to the transaction one at a time, cumulatively, or merged together from multiple transactions. Here is the same transaction above, fully authorized:

```json hl_lines="6 13"
{
  "msig": {
    "subsig": [
      {
        "pk": "SYGHTA2DR5DYFWJE6D4T34P4AWGCG7JTNMY4VI6EDUVRMX7NG4KTA2WMDA",
        "s": "xoQkPyyqCPEhodngmOTP2930Y2GgdmhU/YRQaxQXOwh775gyVSlb1NWn70KFRZvZU96cMtq6TXW+r4sK/lXBCQ=="
      },
      {
        "pk": "VBDMPQACQCH5M6SBXKQXRWQIL7QSR4FH2UI6EYI4RCJSB2T2ZYF2JDHZ2Q"
      },
      {
        "pk": "W3KONPXCGFNUGXGDCOCQYVD64KZOLUMHZ7BNM2ZBK5FSSARRDEXINLYHPI",
        "s": "p1ynP9+LZSOZCBcrFwt5JZB2F+zqw3qpLMY5vJBN83A+55cXDYp5uz/0b+vC0VKEKw+j+bL2TzKSL6aTESlDDw=="
      }
    ],
    "thr": 2,
    "v": 1
  },
  "txn": {
    "amt": 10000000,
    "fee": 1000,
    "fv": 4694301,
    "gen": "testnet-v1.0",
    "gh": "SGO1GKSzyE7IEPItTxCByw9x8FmnrCDexi9/cOUJOiI=",
    "lv": 4695301,
    "rcv": "QC7XT7QU7X6IHNRJZBR67RBMKCAPH67PCSX4LYH4QKVSQ7DQZ32PG5HSVQ",
    "snd": "GQ3QPLJL4VKVGQCHPXT5UZTNZIJAGVJPXUHCJLRWQMFRVL4REVW7LJ3FGY",
    "type": "pay"
  }
}
```

The two signatures are added underneath their respective addresses. Since 2 meets the threshold, this transaction is now fully authorized and can be sent to the network.

!!! info
	Adding more sub-signatures than the threshold requires is unnecessary but perfectly valid.

**How-To**

Extend the example from the [Multisignature Account](../accounts/create.md#multisignature) section by creating, signing, and sending a transaction from a multisig account on TestNet.

```javascript tab="JavaScript"
const algosdk = require('algosdk');

const token = <algod-token>;
const server = <algod-address>;
const port = <algod-port>;

const keypress = async() => {
    process.stdin.setRawMode(true)
    return new Promise(resolve => process.stdin.once('data', () => {
        process.stdin.setRawMode(false)
        resolve()
    }))
}

(async() => {
    //create an account
    var account1 = algosdk.generateAccount();
    console.log(account1.addr);
    //create an account
    var account2 = algosdk.generateAccount();
    console.log(account2.addr);
    //create an account
    var account3 = algosdk.generateAccount();
    console.log(account3.addr);

    //Setup teh parameters for the multisig account
    const mparams = {
        version: 1,
        threshold: 2,
        addrs: [
            account1.addr,
            account2.addr,
            account3.addr,
        ],
    };

    var multsigaddr = algosdk.multisigAddress(mparams);
    console.log("Multisig Address: " + multsigaddr);
    //Pause execution to allow using the dispenser on testnet to put tokens in account
    console.log('Make sure address above has tokens using the dispenser');
    await keypress();
    try {
        let algodclient = new algosdk.Algod(token, server, port);

        //Get the relevant params from the algod
        let params = await algodclient.getTransactionParams();
        let endRound = params.lastRound + parseInt(1000);
        //example of how to write an object into the notefield

        //create a transaction
        let txn = {
            "from": multsigaddr,
            "to": "7GBK5IJCWFPRWENNUEZI3K4CSE5KDIRSR55KWTSDDOBH3E3JJCKGCSFDGQ",
            "fee": params.fee,
            "amount": 200000,
            "firstRound": params.lastRound,
            "lastRound": endRound,
            "genesisID": params.genesisID,
            "genesisHash": params.genesishashb64,
            "note": new Uint8Array(0)
        };
        //Sign wiith first signature
        let rawSignedTxn = algosdk.signMultisigTransaction(txn, mparams, account1.sk).blob;
        //sign with second account
        let twosigs = algosdk.appendSignMultisigTransaction(rawSignedTxn, mparams, account2.sk).blob;
        //submit the transaction
        let tx = (await algodclient.sendRawTransaction(twosigs));
        console.log("Transaction : " + tx.txId);

    } catch (err) {
        console.log(err);
    }
})().then(process.exit)
```

```python tab="Python"
from algosdk import account, transaction, algod, encoding

algod_token = <algod_token>
algod_address = <algod_address>
acl = algod.AlgodClient(algod_token, algod_address)

# generate three accounts
private_key_1, account_1 = account.generate_account()
private_key_2, account_2 = account.generate_account()
private_key_3, account_3 = account.generate_account()
print("Account 1:", account_1)
print("Account 2", account_2)
print("Account 3:", account_3)

# create a multisig account
version = 1  # multisig version
threshold = 2  # how many signatures are necessary
msig = transaction.Multisig(version, threshold, [account_1, account_2])
print("Multisig Address: ", msig.address())
input("Please go to: https://bank.testnet.algorand.network/ to fund your multisig account." + '\n' + "Press Enter to continue...")

# get suggested parameters
params = acl.suggested_params()
gen = params["genesisID"]
gh = params["genesishashb64"]
last_round = params["lastRound"]
fee = params["fee"]

# create a transaction
sender = msig.address()
recipient = "4CXXFP3SJJW63HGEQD4OPSDPPIYDW7BXCVP4DQZG7T33Z3BXTOA4UMEDM4"
amount = 10000
txn = transaction.PaymentTxn(sender, fee, last_round, last_round+100, gh, recipient, amount)

# create a SignedTransaction object
mtx = transaction.MultisigTransaction(txn, msig)

# sign the transaction
mtx.sign(private_key_1)
mtx.sign(private_key_2)

# print encoded transaction
print(encoding.msgpack_encode(mtx))

# send the transaction
transaction_id = acl.send_raw_transaction(encoding.msgpack_encode(mtx))
print("\nTransaction was sent!")
print("Transaction ID: " + transaction_id + "\n")
```

```java tab="Java"
package com.algorand.algosdk.example;

import com.algorand.algosdk.algod.client.AlgodClient;
import com.algorand.algosdk.algod.client.ApiException;
import com.algorand.algosdk.algod.client.api.AlgodApi;
import com.algorand.algosdk.algod.client.auth.ApiKeyAuth;
import com.algorand.algosdk.account.Account;
import com.algorand.algosdk.crypto.Address;
import com.algorand.algosdk.algod.client.model.TransactionParams;
import com.algorand.algosdk.crypto.Digest;
import com.algorand.algosdk.crypto.Ed25519PublicKey;
import com.algorand.algosdk.algod.client.model.TransactionID;
import com.algorand.algosdk.transaction.Transaction;
import com.algorand.algosdk.util.Encoder;



import com.algorand.algosdk.crypto.MultisigAddress;
import com.algorand.algosdk.transaction.SignedTransaction;


import java.io.ByteArrayOutputStream;
import java.math.BigInteger;
import java.util.ArrayList;
import java.util.List;

public class MultisigAccount {

    // Inline class to handle changing block parameters
    // Throughout the example
    static class ChangingBlockParms {
        public BigInteger fee; 
        public BigInteger firstRound;  
        public BigInteger lastRound; 
        public String genID;
        public Digest genHash;
        public ChangingBlockParms() {
            this.fee = BigInteger.valueOf(0);
            this.firstRound = BigInteger.valueOf(0);
            this.lastRound = BigInteger.valueOf(0);
            this.genID = "";
            this.genHash = null;
        }
    };

    // Utility function to update changing block parameters 
    public static ChangingBlockParms getChangingParms(AlgodApi algodApiInstance) throws Exception{
        ChangingBlockParms cp = new MultisigAccount.ChangingBlockParms(); 
        try {
            TransactionParams params = algodApiInstance.transactionParams();
            cp.fee = params.getFee();
            cp.firstRound = params.getLastRound();
            cp.lastRound = cp.firstRound.add(BigInteger.valueOf(1000));
            cp.genID = params.getGenesisID();
            cp.genHash = new Digest(params.getGenesishashb64());

        } catch (ApiException e) {
           throw( e );
        }
        return( cp );
    }

    // Utility function for sending a raw signed transaction to the network
    public static TransactionID submitTransaction(AlgodApi algodApiInstance, SignedTransaction signedTx ) throws Exception{
        try {
            // Msgpack encode the signed transaction
            byte[] encodedTxBytes = Encoder.encodeToMsgPack(signedTx);
            TransactionID id = algodApiInstance.rawTransaction(encodedTxBytes);
            return( id );
        } catch (ApiException e) {
            throw( e );
        }
    }
    public static void main(String args[]) throws Exception {
        final String ALGOD_API_ADDR = "http://localhost:8080";
        final String ALGOD_API_TOKEN = "a967f42b017cd4c5c95a633e87b5ff14226ae60609e174bf5832722631946e13";

        AlgodClient client = new AlgodClient();
        client.setBasePath(ALGOD_API_ADDR);
        ApiKeyAuth api_key = (ApiKeyAuth) client.getAuthentication("api_key");
        api_key.setApiKey(ALGOD_API_TOKEN);

        AlgodApi algodApiInstance = new AlgodApi(client);

        Account acct1 = new Account();
        Account acct2 = new Account();
        Account acct3 = new Account();
        System.out.println("Account 1 Address: " + acct1.getAddress());
        System.out.println("Account 2 Address: " + acct2.getAddress());
        System.out.println("Account 3 Address: " + acct3.getAddress());

        ChangingBlockParms cp = null;
        try {
            cp = getChangingParms(algodApiInstance);
        } catch (ApiException e) {
            e.printStackTrace();
            return;
        }        	
        try {
            cp = getChangingParms(algodApiInstance);
        } catch (ApiException e) {
            e.printStackTrace();
            return;
        }
        
        List<Ed25519PublicKey> publicKeys = new ArrayList<>();
        publicKeys.add(acct1.getEd25519PublicKey());
        publicKeys.add(acct2.getEd25519PublicKey());
        publicKeys.add(acct3.getEd25519PublicKey());
         
        MultisigAddress msig = new MultisigAddress(1, 2, publicKeys);

        final String toAddr = "WICXIYCKG672UGFCCUPBAJ7UYZ2X7GZCNBLSAPBXW7M6DZJ5YY6SCXML4A";

        Transaction tx1 = new Transaction(msig.toAddress(), new Address(toAddr), 1000, cp.firstRound.intValue(), cp.lastRound.intValue(), cp.genID, cp.genHash);

        System.out.println("Multisig Address: " + msig.toString());
        
        SignedTransaction signedTransaction = acct1.signMultisigTransaction(msig, tx1);
       
        SignedTransaction signedTrx2 = acct2.appendMultisigTransaction(msig, signedTransaction);

        System.err.println("Please go to: https://bank.testnet.algorand.network/ to fund your multisig account. \n" + msig.toAddress());
        System.in.read();

        try {
            ByteArrayOutputStream byteOutputStream = new ByteArrayOutputStream( );
            byte[] encodedTxBytes = Encoder.encodeToMsgPack(signedTrx2);
            byteOutputStream.write(encodedTxBytes);

            byte stxBytes[] = byteOutputStream.toByteArray();
                        
            TransactionID id = algodApiInstance.rawTransaction(stxBytes);
            System.out.println("Successfully sent tx group with first tx id: " + id);
            } catch (ApiException e) {
                // This is generally expected, but should give us an informative error message.
                System.err.println("Exception when calling algod#rawTransaction: " + e.getResponseBody());
        }
    }
}
```

```go tab="Go"
package main

import (
	"fmt"

	"github.com/algorand/go-algorand-sdk/client/algod"
	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/transaction"
	"github.com/algorand/go-algorand-sdk/types"
)

// Change these values for your own algod.token and algod.net values
const algodAddress = "http://127.0.0.1:8080"
const algodToken = "a967f42b017cd4c5c95a633e87b5ff14226ae60609e174bf5832722631946e13"

func main() {

	// Initialize an algodClient
	algodClient, err := algod.MakeClient(algodAddress, algodToken)
	if err != nil {
		fmt.Printf("failed to make algod client: %v\n", err)
		return
	}

	txParams, err := algodClient.SuggestedParams()
	if err != nil {
		fmt.Printf("error getting suggested tx params: %s\n", err)
		return
	}

	// Generate Accounts
	acct1 := crypto.GenerateAccount()
	acct2 := crypto.GenerateAccount()
	acct3 := crypto.GenerateAccount()

	// Decode the account addresses
	addr1, _ := types.DecodeAddress(acct1.Address.String())
	addr2, _ := types.DecodeAddress(acct2.Address.String())
	addr3, _ := types.DecodeAddress(acct3.Address.String())

	ma, err := crypto.MultisigAccountWithParams(1, 2, []types.Address{
		addr1,
		addr2,
		addr3,
	})
	if err != nil {
		panic("invalid multisig parameters")
	}

	// declare txn parameters
	fee := txParams.Fee
	firstRound := txParams.LastRound
	lastRound := txParams.LastRound + 1000
	genesisID := txParams.GenesisID     // replace me
	genesisHash := txParams.GenesisHash // replace me
	const amount1 = 2000
	const amount2 = 1500
	var note []byte
	closeRemainderTo := ""

	fromAddr, _ := ma.Address()
	toAddr := "WICXIYCKG672UGFCCUPBAJ7UYZ2X7GZCNBLSAPBXW7M6DZJ5YY6SCXML4A"

	// Create the transaction
	txn, err := transaction.MakePaymentTxn(fromAddr.String(), toAddr, fee, amount1, firstRound, lastRound, note, closeRemainderTo, genesisID, genesisHash)

	// First signature on PST
	txid, preStxBytes, err := crypto.SignMultisigTransaction(acct1.PrivateKey, ma, txn)
	if err != nil {
		panic("could not sign multisig transaction")
	}
	fmt.Printf("Made partially-signed multisig transaction with TxID %s \n", txid)

	// Second signature on PST
	txid2, stxBytes, err := crypto.AppendMultisigTransaction(acct2.PrivateKey, ma, preStxBytes)
	if err != nil {
		panic("could not sign multisig transaction")
	}
	fmt.Printf("Made partially-signed multisig transaction with TxID %s \n", txid2)

	// Print multisig account
	fmt.Printf("Here is your multisig address : %s \n", fromAddr.String())

	fmt.Println("Please go to: https://bank.testnet.algorand.network/ to fund your multisig account.")
	fmt.Scanln() // wait for Enter Key

	// Send transaction to the network
	sendResponse, err := algodClient.SendRawTransaction(stxBytes)
	if err != nil {
		fmt.Printf("Failed to create payment transaction: %v\n", err)
		return
	}
	fmt.Printf("Transaction ID: %s\n", sendResponse.TxID)
}
```

```zsh tab="goal"
# Sign cumulatively
$ goal clerk multisig sign -t multisig.txn -a $ADDRESS1
$ goal clerk multisig sign -t multisig.txn -a $ADDRESS2

# Or sign two separate files and merge
$ goal clerk multisig sign -t multisig1.txn -a $ADDRESS1
$ goal clerk multisig sign -t multisig2.txn -a $ADDRESS2
$ goal clerk multisig merge multisig1.txn multisig2.txn --out=merged.stxn
```

```zsh tab="algokey"
# algokey takes account-level mnemonics at the time of signing
# requires the transaction to include the msig struct before signing
$ algokey multisig --txfile=multisig1.txn --outfile=multisig1.stxn -m <25-word-mnemonic>
$ algokey multisig --txfile=multisig2.txn --outfile=multisig2.stxn -m <25-word-mnemonic>

# Use goal to merge the the *.stxn files.
```

# Logic Signatures

Logic Signatures (or LogicSigs) authorize transactions associated with an Algorand Smart Contract. Logic signatures are added to transactions to authorize spends from a [Contract Account](../asc1/modes.md#contract-account) or from a [Delegated Account](../asc1/modes.md#delegated-account).

A full explanation of Logic Signatures can be found in the [Algorand Smart Contract Usage Modes Guide](.,/asc1/modes/#logic-signatures).

**Related How-To**

- [Use LogicSigs with the SDKs](../asc1/sdks.md)
- [Attach a LogicSig with `goal`](../asc1/goal_teal_walkthrough.md)
