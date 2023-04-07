title: Signatures

In the [Transactions Section](../../transactions), you learned how transactions are composed. In this section you will learn how to authorize them. 

Before a transaction is sent to the network, it must first be authorized by the [sender](../../transactions/#transaction-walkthroughs). Authorization occurs through the addition of a **signature** to the transaction object. Specifically, a transaction object, when signed, is wrapped in a [`SignedTxn`](../../transactions/transactions#signed-transaction) object that includes the [transaction](../../transactions/transactions#txn) and a type of [signature](../../transactions/transactions#sig). 

There are three types of signatures:

- [Single Signatures](#single-signatures)
- [Multisignatures](#multisignatures)
- [Logic Signatures](#logic-signatures)



# Single Signatures
A single signature corresponds to a signature from the private key of an [Algorand public/private key pair](../../accounts#keys-and-addresses).

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
This transaction sends 10 Algo from `"EW64GC..."` to `"QC7XT7..."` on TestNet. The transaction was signed with the private key that corresponds to the `"snd"` address of `"EW64GC..."`. The base64 encoded signature is shown as the value of the [`"sig"`](./transactions.md#sig) field.


# Multisignatures

When the [sender](../../transactions/transactions#sender) of a transaction is the address of a [multisignature account](../../accounts/create#multisignature) then authorization requires a subset of signatures, _equal to or greater than the threshold value_, from the associated private keys of the addresses that multisignature account is composed of. See [Multisignature Accounts](../../accounts/create#multisignature) for details on how to configure a multisignature account.

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
The difference between this transaction and the one above is the form of its signature component. For multisignature accounts, an [`"msig"`](../../transactions/transactions#msig) struct is added which contains the 3 public addresses (`"pk"`), the threshold value (`"thr"`) and the multisig version `"v"`. This transaction is still unsigned but the addition of the correct `"msig"` struct is confirmation that the transaction is "aware" of the fact that the sender is multisig and will have no trouble accepting sub-signatures from single keys even if the signing agent does not contain information about its multisignature properties.

!!! tip
	Adding the `"msig"` template to make the transaction "aware" of its multisig sender is highly recommended, particularly in cases where the transaction is signed by multiple parties or offline. Without it, the signing agent would need to have its own knowledge of the multisignature account. For example, `goal` can sign a multisig transaction that does not contain an `"msig"` template _if_ the multisig address was created within its wallet. On signing, it will add the `"msig"` template. 


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

Extend the example from the [Multisignature Account](../../accounts/create#multisignature) section by creating, signing, and sending a transaction from a multisig account on TestNet.

=== "JavaScript"
    <!-- ===JSSDK_MULTISIG_CREATE=== -->
	```javascript
	const signerAccounts: algosdk.Account[] = [];
	signerAccounts.push(algosdk.generateAccount());
	signerAccounts.push(algosdk.generateAccount());
	signerAccounts.push(algosdk.generateAccount());
	
	// multiSigParams is used when creating the address and when signing transactions
	const multiSigParams = {
	  version: 1,
	  threshold: 2,
	  addrs: signerAccounts.map((a) => a.addr),
	};
	const multisigAddr = algosdk.multisigAddress(multiSigParams);
	
	console.log('Created MultiSig Address: ', multisigAddr);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/accounts.ts#L27-L41)
    <!-- ===JSSDK_MULTISIG_CREATE=== -->
    <!-- ===JSSDK_MULTISIG_SIGN=== -->
	```javascript
	const msigTxn = algosdk.makePaymentTxnWithSuggestedParamsFromObject({
	  from: multisigAddr,
	  to: funder.addr,
	  amount: 100,
	  suggestedParams,
	});
	
	// First signature uses signMultisigTransaction
	const msigWithFirstSig = algosdk.signMultisigTransaction(
	  msigTxn,
	  multiSigParams,
	  signerAccounts[0].sk
	).blob;
	
	// Subsequent signatures use appendSignMultisigTransaction
	const msigWithSecondSig = algosdk.appendSignMultisigTransaction(
	  msigWithFirstSig,
	  multiSigParams,
	  signerAccounts[1].sk
	).blob;
	
	await client.sendRawTransaction(msigWithSecondSig).do();
	await algosdk.waitForConfirmation(client, msigTxn.txID().toString(), 3);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/accounts.ts#L54-L77)
    <!-- ===JSSDK_MULTISIG_SIGN=== -->

=== "Python"
    <!-- ===PYSDK_MULTISIG_CREATE=== -->
	```python
	version = 1  # multisig version
	threshold = 2  # how many signatures are necessary
	# create a Multisig given the set of participants and threshold
	msig = transaction.Multisig(
	    version,
	    threshold,
	    [account_1.address, account_2.address, account_3.address],
	)
	print("Multisig Address: ", msig.address())
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/account.py#L25-L34)
    <!-- ===PYSDK_MULTISIG_CREATE=== -->
    <!-- ===PYSDK_MULTISIG_SIGN=== -->
	```python
	msig_pay = transaction.PaymentTxn(
	    msig.address(),
	    sp,
	    account_1.address,
	    0,
	    close_remainder_to=account_1.address,
	)
	msig_txn = transaction.MultisigTransaction(msig_pay, msig)
	msig_txn.sign(account_2.private_key)
	msig_txn.sign(account_3.private_key)
	txid = algod_client.send_transaction(msig_txn)
	result = transaction.wait_for_confirmation(algod_client, txid, 4)
	print(
	    f"Payment made from msig account confirmed in round {result['confirmed-round']}"
	)
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/account.py#L46-L61)
    <!-- ===PYSDK_MULTISIG_SIGN=== -->

=== "Java"
    <!-- ===JAVASDK_MULTISIG_CREATE=== -->
	```java
	int version = 1; // no other versions at the time of writing
	int threshold = 2; // we're making a 2/3 msig
	
	// Populate a list of Ed25519 pubkeys
	List<Ed25519PublicKey> accts = new ArrayList<>();
	accts.add(addr1.getEd25519PublicKey());
	accts.add(addr2.getEd25519PublicKey());
	accts.add(addr3.getEd25519PublicKey());
	// create the MultisigAddress object
	MultisigAddress msig = new MultisigAddress(version, threshold, accts);
	System.out.printf("msig address: %s\n", msig.toAddress().toString());
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/AccountExamples.java#L77-L88)
    <!-- ===JAVASDK_MULTISIG_CREATE=== -->
    <!-- ===JAVASDK_MULTISIG_SIGN=== -->
	```java
	// Construct transaction with sender as address of msig
	Transaction msigPayTxn = Transaction.PaymentTransactionBuilder()
	                .sender(msig.toAddress())
	                .amount(1000)
	                .receiver(acct1.getAddress())
	                .suggestedParams(sp)
	                .build();
	
	// For each subsig, sign or append to the existing partially signed transaction
	SignedTransaction signedMsigPayTxn = acct1.signMultisigTransaction(msig, msigPayTxn);
	signedMsigPayTxn = acct2.appendMultisigTransaction(msig, signedMsigPayTxn);
	Response<PostTransactionsResponse> msigSubResponse = algodClient.RawTransaction()
	                .rawtxn(Encoder.encodeToMsgPack(signedMsigPayTxn)).execute();
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/AccountExamples.java#L42-L55)
    <!-- ===JAVASDK_MULTISIG_SIGN=== -->

=== "Go"
    <!-- ===GOSDK_MULTISIG_CREATE=== -->
	```go
	// Get pre-defined set of keys for example
	_, pks := loadAccounts()
	addr1, _ := types.DecodeAddress(pks[1])
	addr2, _ := types.DecodeAddress(pks[2])
	addr3, _ := types.DecodeAddress(pks[3])
	
	ma, err := crypto.MultisigAccountWithParams(1, 2, []types.Address{
		addr1,
		addr2,
		addr3,
	})
	
	if err != nil {
		panic("invalid multisig parameters")
	}
	fromAddr, _ := ma.Address()
	// Print multisig account
	fmt.Printf("Multisig address : %s \n", fromAddr)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/kmd/main.go#L178-L196)
    <!-- ===GOSDK_MULTISIG_CREATE=== -->
    <!-- ===GOSDK_MULTISIG_SIGN=== -->
    ```go
    package main

    import (
        "context"
        "crypto/ed25519"
        "fmt"
        json "encoding/json"
        "github.com/algorand/go-algorand-sdk/client/v2/algod"	
        "github.com/algorand/go-algorand-sdk/crypto"
        "github.com/algorand/go-algorand-sdk/mnemonic"
        "github.com/algorand/go-algorand-sdk/transaction"
        "github.com/algorand/go-algorand-sdk/types"
    )

    // UPDATE THESE VALUES
    // const algodAddress = "Your ADDRESS"
    // const algodToken = "Your TOKEN"

    // sandbox
    const algodAddress = "http://localhost:4001"
    const algodToken = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

    // Accounts to be used through examples
    func loadAccounts() (map[int][]byte, map[int]string) {
        // Shown for demonstration purposes. NEVER reveal secret mnemonics in practice.
        // Change these values to use the accounts created previously.

        // Paste in mnemonic phrases for all three accounts
        mnemonic1 := "PASTE phrase for account 1"
        mnemonic2 := "PASTE phrase for account 2"
        mnemonic3 := "PASTE phrase for account 3"
        // never use mnemonics in production code, replace for demo purposes only


        mnemonics := []string{mnemonic1, mnemonic2, mnemonic3}
        pks := map[int]string{1: "", 2: "", 3: ""}
        var sks = make(map[int][]byte)

        for i, m := range mnemonics {
            var err error
            sk, err := mnemonic.ToPrivateKey(m)
            sks[i+1] = sk
            if err != nil {
                fmt.Printf("Issue with account %d private key conversion.", i+1)
            }
            // derive public address from Secret Key.
            pk := sk.Public()
            var a types.Address
            cpk := pk.(ed25519.PublicKey)
            copy(a[:], cpk[:])
            pks[i+1] = a.String()
            fmt.Printf("Loaded Key %d: %s\n", i+1, pks[i+1])
        }
        return sks, pks
    }


    // PrettyPrint prints Go structs
    func PrettyPrint(data interface{}) {
        var p []byte
        //    var err := error
        p, err := json.MarshalIndent(data, "", "\t")
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Printf("%s \n", p)
    }

    func main() {

        // Initialize an algodClient
        algodClient, err := algod.MakeClient(algodAddress, algodToken)
        if err != nil {
            return
        }
        // Get network-related transaction parameters and assign
        txParams, err := algodClient.SuggestedParams().Do(context.Background())
        if err != nil {
            fmt.Printf("error getting suggested tx params: %s\n", err)
            return
        }
        // comment out the next two (2) lines to use suggested fees
        // txParams.FlatFee = true
        // txParams.Fee = 1000
        // Get pre-defined set of keys for example
        sks, pks := loadAccounts()


        addr1, _ := types.DecodeAddress(pks[1])
        addr2, _ := types.DecodeAddress(pks[2])
        addr3, _ := types.DecodeAddress(pks[3])		
        
        ma, err := crypto.MultisigAccountWithParams(1, 2, []types.Address{
            addr1,
            addr2,
            addr3,
        })

        if err != nil {
            panic("invalid multisig parameters")
        }

        fromAddr, _ := ma.Address()
        // Fund account
        fmt.Println("Fund multisig account using testnet faucet:\n--> https://dispenser.testnet.aws.algodev.network?account=" + fromAddr.String())
        fmt.Println("--> Once funded, press ENTER key to continue...")


        //	fmt.Scanln() // wait for Enter Key

        toAddr := addr3.String()
        var amount uint64 = 10000
        note := []byte("Hello World")
        genID := txParams.GenesisID
        genHash := txParams.GenesisHash
        firstValidRound := uint64(txParams.FirstRoundValid)
        lastValidRound := uint64(txParams.LastRoundValid)
        var minFee uint64 = 1000
        txn, err := transaction.MakePaymentTxn(
            fromAddr.String(),
            toAddr,
            minFee,     // fee per byte
            amount,  // amount
            firstValidRound, // first valid round
            lastValidRound, // last valid round
            note,    // note
            "",     // closeRemainderTo
            genID,     // genesisHash
            genHash,     // genesisHash
        )

        txid, txBytes, err := crypto.SignMultisigTransaction(sks[1], ma, txn)
        if err != nil {
            println(err.Error)
            panic("could not sign multisig transaction")
        }
        fmt.Printf("Made partially-signed multisig transaction with TxID %s: %x\n", txid, txBytes)

        txid, twoOfThreeTxBytes, err := crypto.AppendMultisigTransaction(sks[2], ma, txBytes)

        if err != nil {
            panic("could not append signature to multisig transaction")
        }
        fmt.Printf("Appended bytes %x\n", twoOfThreeTxBytes)

        fmt.Printf("Made 2-out-of-3 multisig transaction with TxID %s: %x\n", txid, twoOfThreeTxBytes)


        // We can also merge raw, partially-signed multisig transactions:
        // otherTxBytes := ... // generate another raw multisig transaction 
        // txid, mergedTxBytes, err := crypto.MergeMultisigTransactions(twoOfThreeTxBytes, otherTxBytes)

        // Broadcast the transaction to the network
        txid, err = algodClient.SendRawTransaction(twoOfThreeTxBytes).Do(context.Background())


        // Wait for confirmation
        confirmedTxn, err := transaction.WaitForConfirmation(algodClient,txid,  4, context.Background())
        if err != nil {
            fmt.Printf("Error waiting for confirmation on txID: %s\n", txid)
            return
        }
        fmt.Printf("Confirmed Transaction: %s in Round %d\n", txid ,confirmedTxn.ConfirmedRound)

        txnJSON, err := json.MarshalIndent(confirmedTxn.Transaction.Txn, "", "\t")
        if err != nil {
            fmt.Printf("Can not marshall txn data: %s\n", err)
        }
        fmt.Printf("Transaction information: %s\n", txnJSON)

        fmt.Printf("Decoded note: %s\n", string(confirmedTxn.Transaction.Txn.Note))

    }

    ```
    <!-- ===GOSDK_MULTISIG_SIGN=== -->

=== "goal"
    ```zsh
    # Sign cumulatively
    $ goal clerk multisig sign -t multisig.txn -a $ADDRESS1
    $ goal clerk multisig sign -t multisig.txn -a $ADDRESS2

    # Or sign two separate files and merge
    $ goal clerk multisig sign -t multisig1.txn -a $ADDRESS1
    $ goal clerk multisig sign -t multisig2.txn -a $ADDRESS2
    $ goal clerk multisig merge multisig1.txn multisig2.txn --out=merged.stxn
    ```

=== "algokey"
    ```zsh
    # algokey takes account-level mnemonics at the time of signing
    # requires the transaction to include the msig struct before signing
    $ algokey multisig --txfile=multisig1.txn --outfile=multisig1.stxn -m <25-word-mnemonic>
    $ algokey multisig --txfile=multisig2.txn --outfile=multisig2.stxn -m <25-word-mnemonic>

    # Use goal to merge the the *.stxn files.
    ```

# Logic Signatures

Logic signatures authorize transactions associated with an Algorand Smart Signature. Logic signatures are added to transactions to authorize spends from a [Contract Account](../../dapps/smart-contracts/smartsigs/modes#contract-account) or from a [Delegated Account](../../dapps/smart-contracts/smartsigs/modes#delegated-approval).

A full explanation of Logic Signatures can be found in the [Algorand Smart Contract Usage Modes Guide](../../dapps/smart-contracts/smartsigs/modes).

**Related How-To**

- [Use LogicSigs with the SDKs](../../dapps/smart-contracts/frontend/smartsigs)
- [Attach a LogicSig with `goal`](../../dapps/smart-contracts/smartsigs/walkthrough)