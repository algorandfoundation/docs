Atomic Transfers are irreducible batch operations that allow a group of [transactions](../features/transactions.md) to be submitted as a unit. All transactions in the batch either pass or fail. The batch operation will only be successful if each transaction within the group is successful.

Transactions can contain Alogs or Algorand Assets and may also be governed by Algorand Smart Contracts. 

Individual transactions are first created and then grouped into a data structure or a file. The grouped transactions are then singed individually by the originators within the group. This data structure or file is then submitted to the network, where the node verifies the transactions and submits them all at once. This eliminates the need for hashed time-locks for atomic swaps in other blockchains and also avoids the delays inherent with hashed timed-locks.
Atomic Transfers enable applications such as: 

* Circular trades: Alice pays Bob if and only if Bob pays Clare if and only if Clare pays Alice
* Group payments: Group funding where everyone pays or none pay.
* Decentralized exchanges: Atomic multi-party transfers require no trusted intermediaries.
* Distributed payments: Payments to multiple recipients


Atomic Transfers are Created with the following steps:

* Create unsigned transactions 
* Combine transactions into one data structure or file
* Group transactions
* Sign the grouped transaction with private keys
* Submit signed grouped transaction data structure or file to the network 

<center>![Atomic Transfer Flow](../../imgs/atomic_transfers-1.png)</center>
<center>*Atomic Transfer Flow*</center>


Below you will find examples for creating group transactions and sending them to the network in each of the available SDKs. The same capability is also shown using the `goal` command-line tool. The example code is separated into snippets categorized by these core functions.
# Create transactions
Transaction creation functions are described in the [Transactions](../features/transactions.md) documentation. Atomic Transfers are created by constructing two or more transactions that are not signed and are written to a file or stored in a data structure within an application. By not signing the transaction, this allows atomic transfers to be created by one or more parties, possibly at different times. For example, an asset exchange application can create the entire atomic transfer and allow individual parties to sign from their location. If an application is using files to store the transactions, see the [Offline Transactions](../features/offline_transactions.md#saving-unsigned-transactions-to-file) documentation which explains how to create and save individual **unsigned** transactions to a file. 

# Combine transactions into one data structure or file
Individual transactions can be combined in various ways. Each SDK language provides specific methodologies for combining transactions. Transactions can be stored as a local varible within a service application or possibly read from a file. How the transactions are combined will be application dependent. If using the `goal` command-line tool all transaction files will be combined using an OS-level command such as `cat`. If using one of the SDKs, the application may store all the transactions individually or in an array. From the SDK it is also possible to read a transaction from a file created at an earlier time, which is described in the [Offline Transactions](../features/offline_transactions.md) documentation. 



The SDKs or `goal` can be used to create unsigned transactions. The following snippets assume one transaction is read from a file and another is created using the SDK. The example illustrates Account A sending a transaction to Account C and Account B sending a transaction to Account A.

``` javascript tab="JavaScript"
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
```

``` python tab="Python"
```

``` java tab="Java"
```

``` go tab="Go"
```

``` goal tab="goal"
```


# Group transactions
All transactions must be grouped. The follow code illustrates how this is done.

``` javascript tab="JavaScript"
	// Group both transactions
	let txgroup = algosdk.assignGroupID(txns);
```

``` python tab="Python"
import tensorflow as tf
# get group id and assign it to transactions
gid = transaction.calculate_group_id([txn1, txn2])
txn1.group = gid
txn2.group = gid

```

``` Go tab="Go"
// compute group id and put it into each transaction
gid, err := crypto.ComputeGroupID([]types.Transaction{tx1, tx2})
tx1.Group = gid
tx2.Group = gid

```

``` Java tab=
// compute group id and put it into each transaction
gid, err := crypto.ComputeGroupID([]types.Transaction{tx1, tx2})
tx1.Group = gid
tx2.Group = gid

```


# Sign transactions
After the unsigned transactions are grouped they must be signed individually by every account or multisig account that is sending funds within the group.

``` javascript tab="JavaScript"
	// Sign each transaction in the group with
	// correct key and store in array
	let signed = []
	signed.push( transaction1.signTxn( myAccountA.sk ) )
	signed.push( transaction2.signTxn( myAccountB.sk ) )
```



# Send transactions
Grouped signed transactions are then sent to a node to process. 

``` javascript tab="JavaScript"
	let tx = (await algodClient.sendRawTransactions(signed));
	console.log("Transaction : " + tx.txId);

	// Wait for transaction to be confirmed
	await waitForConfirmation(algodClient, tx.txId)
```

??? example "Complete Example = Atomic Transfer"
    
    ```javascript tab="JavaScript"
    const algosdk = require('algosdk');
    const fs = require('fs');

    var client = null;
    async function setupClient() {
        if( client == null){
            const ALGOD_API_ADDR = "algod-address<PLACEHOLDER>";
            const ALGOD_API_TOKEN = "algod-token<PLACEHOLDER>";
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
        const passphrase ="your-25-word-mnemonic<PLACEHOLDER>";
        let myAccount = algosdk.mnemonicToSecretKey(passphrase);
        return myAccount;
    }
    // recover second account
    function recoverAccount2(){
        const passphrase ="your-25-word-mnemonic<PLACEHOLDER>";
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
            const receiver = "transaction-receiver-address<PLACEHOLDER>"
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