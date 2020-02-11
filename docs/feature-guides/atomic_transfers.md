title: Swapping Assets (Atomic Transfers)

In traditional finance, trading assets generally requires a trusted intermediary, like a bank or an exchange, to make sure that both sides receive what they agreed to. On the Algorand blockchain, this type of trade is implemented within the protocol as an **Atomic Transfer**. This simply means that transactions that are part of the transfer either all succeed or all fail. Atomic transfers allow complete strangers to trade assets without the need for a trusted intermediary, all while guaranteeing that each party will receive what they agreed to. 

On Algorand, atomic transfers are implemented as irreducible batch operations, where a group of [transactions](../feature-guides/transactions.md) are submitted as a unit and all transactions in the batch either pass or fail. This also eliminates the need for more complex solutions like [hashed timelock contracts](https://en.bitcoinwiki.org/wiki/Hashed_Timelock_Contracts) that are implemented on other blockchains. An atomic transfer on Algorand is confirmed in less than 5 seconds, just like any other transaction. Transactions can contain Algos or Algorand Standard Assets and may also be governed by Algorand Smart Contracts. 

# Use Cases

Atomic transfers enable use cases such as:

**Circular trades** - Alice pays Bob if and only if Bob pays Claire if and only if Claire pays Alice.

**Group payments** - Everyone pays or no one pays.

**Decentralized exchanges** - Trade one asset for another without going through a centralized exchange. 

**Distributed payments** - Payments to multiple recipients. 

# Process Overview

To implement an atomic transfer, generate all of the transactions that will be involved in the transfer and then group those transactions together. The result of grouping is that each transaction is assigned the same group ID. Once the transactions contain the group ID, the transactions can be split up and sent to their respective senders to be authorized. A single party can then collect all the authorized transactions and submit them to the network together. 

!!! info
    An individual account involved in an atomic transfer, can verify that all the correct transfers are involved by creating the same set of unauthorized transactions and grouping them in the same order. The group ID is a hash of the group of transactions and should match if the configuration is the same.

<center>![Atomic Transfer Flow](../imgs/atomic_transfers-1.png)</center>
<center>*Atomic Transfer Flow*</center>

Below you will find examples for creating and sending group transactions to the network in each of the available SDKs and with `goal`. 

# Step-by-Step Guide

## Create transactions
Create two or more (up to 16 total) unsigned transactions of any type. Read about transaction types in the [Constructing Transactions](./transactions.md) guide. 

This could be done by a service or by each party involved in the transaction. For example, an asset exchange application can create the entire atomic transfer and allow individual parties to sign from their location.

The example below illustrates Account A sending a transaction to Account C and Account B sending a transaction to Account A.

``` javascript tab="JavaScript"
	// Transaction A to C 
	let transaction1 = algosdk.makePaymentTxn(myAccountA.addr, 
		receiver, params.minFee, 100000, undefined, 
		params.lastRound, params.lastRound + 1000, new Uint8Array(0), 
		params.genesishashb64, params.genesisID);

	// Create transaction B to A
	let transaction2 = algosdk.makePaymentTxn(myAccountB.addr, 
		myAccountA.addr, params.minFee, 200000, undefined, 
		params.lastRound, params.lastRound + 1000, new Uint8Array(0), 
		params.genesishashb64, params.genesisID);
			
```

``` python tab="Python"  
    # create transaction1
	txn1 = transaction.PaymentTxn(account_a, fee, last_round, last_round+100, gh, account_c, amount)

	# create transaction2
	txn2 = transaction.PaymentTxn(account_b, fee, last_round, last_round+100, gh, account_a, amount)
```

``` java tab="Java"
    // Create the first transaction
    Transaction tx1 = new Transaction(acctA.getAddress(), 
        acctC.getAddress(), 10000, cp.firstRound.intValue(), 
        cp.lastRound.intValue(), null, cp.genHash);
    tx1.fee = BigInteger.valueOf(1000);

    // Create the second transaction
    Transaction tx2 = new Transaction(acctB.getAddress(), 
        acctA.getAddress(), 20000, cp.firstRound.intValue(), 
        cp.lastRound.intValue(), null, cp.genHash);
    tx2.fee = BigInteger.valueOf(1000);
```

``` go tab="Go"
	tx1, err := transaction.MakePaymentTxn(account1, account3, 1, 100000,
		txParams.LastRound, txParams.LastRound+100, nil, "", 
		txParams.GenesisID, txParams.GenesisHash)
   	if err != nil {
	   	fmt.Printf("Error creating transaction: %s\n", err)
	   	return
	}
	   
	tx2, err := transaction.MakePaymentTxn(account2, account1, 1, 100000,
		txParams.LastRound, txParams.LastRound+100, nil, "", 
		txParams.GenesisID, txParams.GenesisHash)
   	if err != nil {
	   	fmt.Printf("Error creating transaction: %s\n", err)
	   	return
   	}
```

``` goal tab="goal"
$ goal clerk send --from=my-account-a<PLACEHOLDER> --to=my-account-c<PLACEHOLDER> --fee=1000 --amount=1000000 --out=unsginedtransaction1.txn"

$ goal clerk send --from=my-account-b<PLACEHOLDER> --to=my-account-a<PLACEHOLDER> --fee=1000 --amount=1000000 --out=unsginedtransaction2.txn"
```

At this point, these are just individual transactions. The next critical step is to combine them and then calculate the group ID.

## Combine transactions 
Combining transactions just means concatenating them into a single file or ordering them in an array so that a group ID can then be assigned. 

If using `goal`, the transaction files can be combined using an OS-level command such as `cat`. If using one of the SDKs, the application may store all the transactions individually or in an array. From the SDK it is also possible to read a transaction from a file created at an earlier time, which is described in the [Offline Transactions](../feature-guides/offline_transactions.md) documentation. See the complete example at the bottom of this page that details how transactions are combined in the SDKs. To combine transactions in `goal` use a similar method to the one below.

``` goal tab="goal"
cat unsignedtransaction1.tx unsignedtransaction2.tx > combinedtransactions.tx
```


## Group transactions

The result of this step is what ultimately guarantees that a particular transaction belongs to a group and is not valid if sent alone (even if authorized). A group ID is calculated by hashing the contents of the combined transaction and assigning the resulting hash as a [group ID](../reference-docs/transactions.md#group) to each transaction. This mechanism allows anyone to recreate all transactions and recalculate the group ID to verify that the contents are as agreed upon by all parties. 

``` javascript tab="JavaScript"
	// Group both transactions
	let txgroup = algosdk.assignGroupID(txns);
```

``` python tab="Python"
	# get group id and assign it to transactions
	gid = transaction.calculate_group_id([txn1, txn2])
	txn1.group = gid
	txn2.group = gid
```

``` java tab="Java"
    // group transactions an assign ids
    Digest gid = TxGroup.computeGroupID(new Transaction[]{tx1, tx2});
    tx1.assignGroupID(gid);
    tx2.assignGroupID(gid);
```

``` go tab="Go"
	// compute group id and put it into each transaction
	gid, err := crypto.ComputeGroupID([]types.Transaction{tx1, tx2})
	tx1.Group = gid
	tx2.Group = gid
```

``` goal tab="goal"
goal clerk group -i yourwalletcombinedtransactions.tx -o groupedtransactions.tx -d data -w 
```

At this point, transactions can be split and sent to individuals for authorization. See [Authorizing Transactions Offline](../feature-guides/offline_transactions.md#saving-unsigned-transactions-to-file) to learn how to create and save individual **unsigned** transactions to a file. This method can be used to distribute group transactions for signing.

## Sign transactions
With a group ID assigned, each transaction sender must authorize their respective transaction. All authorized transactions are then recombined before they are sent to the network.

``` javascript tab="JavaScript"
	// Sign each transaction in the group with
	// correct key and store in array
	let signed = []
	signed.push( transaction1.signTxn( myAccountA.sk ) )
	signed.push( transaction2.signTxn( myAccountB.sk ) )
```

``` python tab="Python"
	# sign transaction1
	stxn1 = txn1.sign(pk_account_a)

	# sign transaction2
	stxn2 = txn2.sign(pk_account_b)

	signedGroup =  []
	signedGroup.append(stxn1)
	signedGroup.append(stxn2)
```

``` java tab="Java"
    // sign individual transactions
    SignedTransaction signedTx1 = acctA.signTransaction(tx1);;
    SignedTransaction signedTx2 = acctB.signTransaction(tx2);;
    
    // put both transaction in a byte array 
    ByteArrayOutputStream byteOutputStream = new ByteArrayOutputStream( );
    byte[] encodedTxBytes1 = Encoder.encodeToMsgPack(signedTx1);
    byte[] encodedTxBytes2 = Encoder.encodeToMsgPack(signedTx2);
    byteOutputStream.write(encodedTxBytes1);
    byteOutputStream.write(encodedTxBytes2);
    byte groupTransactionBytes[] = byteOutputStream.toByteArray();    
```        

``` go tab="Go"
	_, stx1, err := crypto.SignTransaction(sk1, tx1)
	if err != nil {
		fmt.Printf("Failed to sign transaction: %s\n", err)
		return
	}
	_, stx2, err := crypto.SignTransaction(sk2, tx2)
	if err != nil {
		fmt.Printf("Failed to sign transaction: %s\n", err)
		return
	}

	var signedGroup []byte
	signedGroup = append(signedGroup, stx1...)
	signedGroup = append(signedGroup, stx2...)
```

``` goal tab="goal"
# keys on single machine
 $ goal clerk sign -i groupedtransactions.tx -o signout.tx -d data -w yourwallet

# keys on multiple machines
$ goal clerk split -i groupedtransactions.tx -o splitfiles -d data -w yourwallet 

Wrote transaction 0 to splitfiles-0
Wrote transaction 1 to splitfiles-1

# sign on individual machine
$ goal clerk sign -i splitfiles-0 -o splitfiles-0.sig -d data -w yourwallet
$ goal clerk sign -i splitfiles-1 -o splitfiles-1.sig -d data -w yourwallet

# combine signed transactions files
cat splitfiles-0.sig splitfiles-1.sig > signout.tx
```

## Send transactions
The signed group transactions are sent to the network together. 

``` javascript tab="JavaScript"
	let tx = (await algodClient.sendRawTransactions(signed));
	console.log("Transaction : " + tx.txId);

	// Wait for transaction to be confirmed
	await waitForConfirmation(algodClient, tx.txId)
```

``` python tab="Python"
	# send them over network
	sent = acl.send_transactions(signedGroup)
	# print txid
	print(sent)

	# wait for confirmation
	wait_for_confirmation( acl, sent) 
```

``` java tab="Java"
    // write transaction to node
    TransactionID id = algodApiInstance.rawTransaction(groupTransactionBytes);
    System.out.println("Successfully sent tx group with first tx id: " + id);
    waitForConfirmation(id.getTxId());
```

``` go tab="Go"
	signed, err := algodClient.SendRawTransaction(signedGroup)
	if err != nil {
		fmt.Printf("Failed to create payment transaction: %v\n", err)
		return
	}
	fmt.Printf("Transaction ID: %s\n", signed.TxID)
	waitForConfirmation(algodClient, signed.TxID)
```

``` goal tab="Goal"
goal clerk rawsend -f signout.tx -d data -w yourwallet
```

??? example "Complete Example = Atomic Transfer"
    
    ```javascript tab="JavaScript"
    const algosdk = require('algosdk');
    const fs = require('fs');

    var client = null;
    async function setupClient() {
        if( client == null){
            const ALGOD_API_ADDR = <algod-address>;
            const ALGOD_API_TOKEN = <algod-token>;
            const port = port-number<PLACEHOLDER>;
            let algodClient = new algosdk.Algod(token, server, port);
            client = algodClient;
        } else {
            return client;
        }
        return client;
    }
    // recover first account
    function recoverAccount1(){
        const passphrase = <25-word-passphrase>;
        let myAccount = algosdk.mnemonicToSecretKey(passphrase);
        return myAccount;
    }
    // recover second account
    function recoverAccount2(){
        const passphrase = <25-word-passphrase>;
        let myAccount = algosdk.mnemonicToSecretKey(passphrase);
        return myAccount;
    }
    // function used to wait for a tx confirmation
    var waitForConfirmation = async function(algodclient, txId) {
        while (true) {
            let lastround = (await algodclient.status()).lastRound;
            let pendingInfo = await algodclient.pendingTransactionInformation(txId);
            if (pendingInfo.round != null && pendingInfo.round > 0) {
                //Got the completed Transaction
                console.log("Transaction " + pendingInfo.tx + " confirmed in round " + pendingInfo.round);
                break;
            }
            await algodclient.statusAfterBlock(lastround + 1);
        }
    };
    async function submitGroupTransactions(){

        try{
            // receiver
            const receiver = <transaction-receiver-address>"
            // sample show account A to C
            // B to A 
            // grouped
            let algodClient = await setupClient();

            // Creat transactin A to C and write to a file
            await writeUnsignedTransctionToFile();

            // recover account
            // Account A
            let myAccountA = await recoverAccount1();
            console.log("My account A address: %s", myAccountA.addr)

            // recover an additional account
            // Account B
            let myAccountB = await recoverAccount2();
            console.log("My account B address: %s", myAccountB.addr)

            // get suggested params from the network
            let params = await algodClient.getTransactionParams();

            // Transaction A to C 
            let transaction1 = algosdk.makePaymentTxn(myAccountA.addr, 
                receiver, params.minFee, 100000, undefined, 
                params.lastRound, params.lastRound + 1000, new Uint8Array(0), 
                params.genesishashb64, params.genesisID);

            // Create transaction B to A
            let transaction2 = algosdk.makePaymentTxn(myAccountB.addr, 
                myAccountA.addr, params.minFee, 200000, undefined, 
                params.lastRound, params.lastRound + 1000, new Uint8Array(0), 
                params.genesishashb64, params.genesisID);
                
            // Store both transactions
            let txns = [transaction1, transaction2];

            // Group both transactions
            let txgroup = algosdk.assignGroupID(txns);

            // Sign each transaction in the group with
            // correct key
            let signed = []
            signed.push( transaction1.signTxn( myAccountA.sk ) )
            signed.push( transaction2.signTxn( myAccountB.sk ) )

            let tx = (await algodClient.sendRawTransactions(signed));
            console.log("Transaction : " + tx.txId);

            // Wait for transaction to be confirmed
            await waitForConfirmation(algodClient, tx.txId)
        } catch (err) {
            console.log("err", err);  
        }
    }
    submitGroupTransactions();
    ```

    ```python tab="Python"
    #/usr/bin/python3
    import json
    import time
    import base64
    import os
    from algosdk import algod
    from algosdk import mnemonic
    from algosdk import transaction
    from algosdk import encoding
    from algosdk import account

    # utility to connect to node
    def connect_to_network():
        algod_address = <algod-address>
        algod_token = <algod-token>
        algod_client = algod.AlgodClient(algod_token, algod_address)
        return algod_client

    # utility for waiting on a transaction confirmation
    def wait_for_confirmation( algod_client, txid ):
        while True:
            txinfo = algod_client.pending_transaction_info(txid)
            if txinfo.get('round') and txinfo.get('round') > 0:
                print("Transaction {} confirmed in round {}.".format(txid, txinfo.get('round')))
                break
            else:
                print("Waiting for confirmation...")
                algod_client.status_after_block(algod_client.status().get('lastRound') +1)

    # group transactions           
    def group_transactions() :
        
        # recover a account    
        passphrase1 = <25-word-passphrase>
        pk_account_a = mnemonic.to_private_key(passphrase1)
        account_a = account.address_from_private_key(pk_account_a)

        # recover b account
        passphrase2 = <25-word-passphrase>
        pk_account_b = mnemonic.to_private_key(passphrase2)
        account_b = account.address_from_private_key(pk_account_b)

        # recover c account
        passphrase3 = <25-word-passphrase>
        pk_account_c = mnemonic.to_private_key(passphrase3)
        account_c = account.address_from_private_key(pk_account_c)

        # connect to node
        acl = connect_to_network()

        # get suggested parameters
        params = acl.suggested_params()
        gen = params["genesisID"]
        gh = params["genesishashb64"]
        last_round = params["lastRound"]
        fee = params["fee"]
        amount = 1000
        
        # create transaction1
        txn1 = transaction.PaymentTxn(account_a, fee, last_round, last_round+100, gh, account_c, amount)

        # create transaction2
        txn2 = transaction.PaymentTxn(account_b, fee, last_round, last_round+100, gh, account_a, amount)

        # get group id and assign it to transactions
        gid = transaction.calculate_group_id([txn1, txn2])
        txn1.group = gid
        txn2.group = gid

        # sign transaction1
        stxn1 = txn1.sign(pk_account_a)

        # sign transaction2
        stxn2 = txn2.sign(pk_account_b)

        signedGroup =  []
        signedGroup.append(stxn1)
        signedGroup.append(stxn2)

        # send them over network
        sent = acl.send_transactions(signedGroup)
        # print txid
        print(sent)

        # wait for confirmation
        wait_for_confirmation( acl, sent) 

    # Test Runs     
    group_transactions()
    ```

    ```java tab="Java"
    package com.algorand.javatest;

    import java.io.ByteArrayOutputStream;
    import java.math.BigInteger;

    import com.algorand.algosdk.account.Account;
    import com.algorand.algosdk.algod.client.AlgodClient;
    import com.algorand.algosdk.algod.client.ApiException;
    import com.algorand.algosdk.algod.client.api.AlgodApi;
    import com.algorand.algosdk.algod.client.auth.ApiKeyAuth;
    import com.algorand.algosdk.algod.client.model.TransactionID;
    import com.algorand.algosdk.algod.client.model.TransactionParams;
    import com.algorand.algosdk.crypto.Digest;
    import com.algorand.algosdk.transaction.SignedTransaction;
    import com.algorand.algosdk.transaction.Transaction;
    import com.algorand.algosdk.transaction.TxGroup;
    import com.algorand.algosdk.util.Encoder;



    public class GroupedTransaction {   

        public AlgodApi algodApiInstance = null;
        
        // utility function to connect to a node
        private AlgodApi connectToNetwork(){

            final String ALGOD_API_ADDR = <algod-address>;
            final String ALGOD_API_TOKEN = <algod-token>;

            AlgodClient client = (AlgodClient) new AlgodClient().setBasePath(ALGOD_API_ADDR);
            ApiKeyAuth api_key = (ApiKeyAuth) client.getAuthentication("api_key");
            api_key.setApiKey(ALGOD_API_TOKEN);
            algodApiInstance = new AlgodApi(client);   
            return algodApiInstance;
        }    
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
            ChangingBlockParms cp = new GroupedTransaction.ChangingBlockParms(); 
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

        public void waitForConfirmation( String txID ) throws Exception{
            if( algodApiInstance == null ) connectToNetwork();
            while(true) {
                try {
                    //Check the pending tranactions
                    com.algorand.algosdk.algod.client.model.Transaction pendingInfo = algodApiInstance.pendingTransactionInformation(txID);
                    if (pendingInfo.getRound() != null && pendingInfo.getRound().longValue() > 0) {
                        //Got the completed Transaction
                        System.out.println("Transaction " + pendingInfo.getTx() + " confirmed in round " + pendingInfo.getRound().longValue());
                        break;
                    } 
                    algodApiInstance.waitForBlock(BigInteger.valueOf( algodApiInstance.getStatus().getLastRound().longValue() +1 ) );
                } catch (Exception e) {
                    throw( e );
                }
            }

        }

        public void AtomicTransfer() throws Exception {

            if( algodApiInstance == null ) connectToNetwork();;

            final String account1_mnemonic = <25-word-passphrase>;
            final String account2_mnemonic = <25-word-passphrase>;
            final String account3_mnemonic = <25-word-passphrase>;
            // recover account A, B, C
            Account acctA  = new Account(account1_mnemonic); 
            Account acctB  = new Account(account2_mnemonic);
            Account acctC  = new Account(account3_mnemonic); 
            
            // get node suggested parameters
            ChangingBlockParms cp = null;
            try {
                cp = getChangingParms(algodApiInstance);
            } catch (ApiException e) {
                e.printStackTrace();
                return;
            }        	

            // Create the first transaction
            Transaction tx1 = new Transaction(acctA.getAddress(), 
                acctC.getAddress(), 10000, cp.firstRound.intValue(), 
                cp.lastRound.intValue(), null, cp.genHash);
            tx1.fee = BigInteger.valueOf(1000);

            // Create the second transaction
            Transaction tx2 = new Transaction(acctB.getAddress(), 
                acctA.getAddress(), 20000, cp.firstRound.intValue(), 
                cp.lastRound.intValue(), null, cp.genHash);
            tx2.fee = BigInteger.valueOf(1000);

            // group transactions an assign ids
            Digest gid = TxGroup.computeGroupID(new Transaction[]{tx1, tx2});
            tx1.assignGroupID(gid);
            tx2.assignGroupID(gid);

            // sign individual transactions
            SignedTransaction signedTx1 = acctA.signTransaction(tx1);;
            SignedTransaction signedTx2 = acctB.signTransaction(tx2);;
                
            try {
                // put both transaction in a byte array 
                ByteArrayOutputStream byteOutputStream = new ByteArrayOutputStream( );
                byte[] encodedTxBytes1 = Encoder.encodeToMsgPack(signedTx1);
                byte[] encodedTxBytes2 = Encoder.encodeToMsgPack(signedTx2);
                byteOutputStream.write(encodedTxBytes1);
                byteOutputStream.write(encodedTxBytes2);
                byte groupTransactionBytes[] = byteOutputStream.toByteArray();
                
                // write transaction to node
                TransactionID id = algodApiInstance.rawTransaction(groupTransactionBytes);
                System.out.println("Successfully sent tx group with first tx id: " + id);
                waitForConfirmation(id.getTxId());

            } catch (Exception e) {
                System.out.println("Submit Exception: " + e); 
            }
        }
        public static void main(String args[]) throws Exception {
            GroupedTransaction mn = new GroupedTransaction();
            mn.AtomicTransfer();
        }
    }
    ```

    ```go tab="Go"
    package main

    import (
        "fmt"

        "golang.org/x/crypto/ed25519"

        "github.com/algorand/go-algorand-sdk/client/algod"
        "github.com/algorand/go-algorand-sdk/crypto"
        "github.com/algorand/go-algorand-sdk/mnemonic"
        "github.com/algorand/go-algorand-sdk/transaction"
        "github.com/algorand/go-algorand-sdk/types"
    )

    const algodToken = <algod-token>
    const algodAddress = <algod-address>

    // Function that waits for a given txId to be confirmed by the network
    func waitForConfirmation(algodClient algod.Client, txID string) {
        for {
            pt, err := algodClient.PendingTransactionInformation(txID)
            if err != nil {
                fmt.Printf("waiting for confirmation... (pool error, if any): %s\n", err)
                continue
            }
            if pt.ConfirmedRound > 0 {
                fmt.Printf("Transaction "+pt.TxID+" confirmed in round %d\n", pt.ConfirmedRound)
                break
            }
            nodeStatus, err := algodClient.Status()
            if err != nil {
                fmt.Printf("error getting algod status: %s\n", err)
                return
            }
            algodClient.StatusAfterBlock( nodeStatus.LastRound + 1)
        }
    }
    // utility funcitn to get address string
    func getAddress(mn string )(string) {
        sk, err := mnemonic.ToPrivateKey(mn)
        if err != nil {
            fmt.Printf("error recovering account: %s\n", err)
            return ""
        }
        pk := sk.Public()
        var a types.Address
        cpk := pk.(ed25519.PublicKey)
        copy(a[:], cpk[:])
        fmt.Printf("Address: %s\n", a.String())	
        address := a.String()
        return address
    }
    func main() {
        // Initialize an algodClient
        algodClient, err := algod.MakeClient(algodAddress, algodToken)
        if err != nil {
            fmt.Printf("failed to make algod client: %v\n", err)
            return
        }

        // get node suggested parameters
        txParams, err := algodClient.SuggestedParams()
        if err != nil {
            fmt.Printf("error getting suggested tx params: %s\n", err)
            return
        }

        // declare account mnemonics for later consumption for private key conversion
        const mnemonic1 = <25-word-passphrase>;
        const mnemonic2 = <25-word-passphrase>;
        const mnemonic3 = <25-word-passphrase>;

        // convert mnemonic1 and mnemonic2 using the mnemonic.ToPrivateKey() helper function
        sk1, err := mnemonic.ToPrivateKey(mnemonic1)
        sk2, err := mnemonic.ToPrivateKey(mnemonic2)
        // declare accounts
        account1 := getAddress(mnemonic1)
        account2 := getAddress(mnemonic2)
        account3 := getAddress(mnemonic3)

        // make transactions
        tx1, err := transaction.MakePaymentTxn(account1, account3, 1, 100000,
            txParams.LastRound, txParams.LastRound+100, nil, "", 
            txParams.GenesisID, txParams.GenesisHash)
        if err != nil {
            fmt.Printf("Error creating transaction: %s\n", err)
            return
        }
        
        tx2, err := transaction.MakePaymentTxn(account2, account1, 1, 100000,
            txParams.LastRound, txParams.LastRound+100, nil, "", 
            txParams.GenesisID, txParams.GenesisHash)
        if err != nil {
            fmt.Printf("Error creating transaction: %s\n", err)
            return
        }

        // compute group id and put it into each transaction
        gid, err := crypto.ComputeGroupID([]types.Transaction{tx1, tx2})
        tx1.Group = gid
        tx2.Group = gid

        // sign transactions
        _, stx1, err := crypto.SignTransaction(sk1, tx1)
        if err != nil {
            fmt.Printf("Failed to sign transaction: %s\n", err)
            return
        }
        _, stx2, err := crypto.SignTransaction(sk2, tx2)
        if err != nil {
            fmt.Printf("Failed to sign transaction: %s\n", err)
            return
        }

        // send transactions
        var signedGroup []byte
        signedGroup = append(signedGroup, stx1...)
        signedGroup = append(signedGroup, stx2...)
        signed, err := algodClient.SendRawTransaction(signedGroup)
        if err != nil {
            fmt.Printf("Failed to create payment transaction: %v\n", err)
            return
        }
        fmt.Printf("Transaction ID: %s\n", signed.TxID)
        waitForConfirmation(algodClient, signed.TxID)
    }

    ```