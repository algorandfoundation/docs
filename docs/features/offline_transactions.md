Transactions are described in detail in the Transactions documumentation<LINK>. In many cases, transactions must be created for offline usage in an offline application. These transactions can be signed or unsigned depending on the usage case. This guide explains how transactions can be created and saved to a file. The samples do make use of a node connection to get network suggested parameters and to submit transactions. Suggested network parameters can be hardcoded and transactions can be transferred manually for a complete offline application.

# Saving Unsigned Transactions to File 
Algorand SDK's and `goal` support writing both signed and unsigned transactions to a file. Examples of these scenarios are shown in the following code snippets.

Unsigned transactions require the transaction object to be created before writting to a file.


``` javascript tab="JavaScript"
	let txn = {
		"from": myAccount.addr,
		"to": receiver,
		"fee": params.minFee,
		"flatFee": true,
		"amount": 1000000,
		"firstRound": params.lastRound,
		"lastRound": params.lastRound + 1000,
		"genesisID": params.genesisID,
		"genesisHash": params.genesishashb64
    };
	
	// write the unsigned transaction to a file
    fs.writeFileSync('./unsigned.txn', JSON.stringify(txn) , 'utf-8');
```

``` python tab="Python"
	# create transaction
	receiver = "transaction-receiver<PLACEHOLDER>"
	data = {
		"sender": my_address,
		"receiver": receiver,
		"fee": params.get('minFee'),
		"flat_fee": True,
		"amt": amount<PLACEHOLDER>,
		"first": params.get('lastRound'),
		"last": params.get('lastRound') + 1000,
		"gen": params.get('genesisID'),
		"gh": params.get('genesishashb64')
	}
	txn = transaction.PaymentTxn(**data)

	# write to file
	dir_path = os.path.dirname(os.path.realpath(__file__))
	transaction.write_to_file([txn], dir_path + "/unsigned.txn")
```

``` java tab="Java"
    try { 
            Transaction tx = new Transaction(new Address(SRC_ADDR),  
                    BigInteger.valueOf(1000), firstRound, lastRound, 
                    null, amount, new Address(DEST_ADDR), genId, genesisHash);
            FileOutputStream file = new FileOutputStream("./unsigned.txn"); 
            ObjectOutputStream out = new ObjectOutputStream(file); 
            out.writeObject(tx);
            out.close(); 
            file.close();
    } catch (Exception e) { 
        System.out.println("Exception: " + e); 
    }
```

``` go tab="Go"
	tx, err := transaction.MakePaymentTxn(addr, toAddr, 1, 100000,
		 txParams.LastRound, txParams.LastRound+100, nil, "", 
		 genID, txParams.GenesisHash)
	if err != nil {
		fmt.Printf("Error creating transaction: %s\n", err)
		return
	}

	// save unsigned transction to file
	file, err := os.Create("./unsigned.gob")
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(tx)
	}
	file.Close()
```


``` goal tab="goal"
$ goal clerk send --from=my-account<PLACEHOLDER> --to=GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A --fee=1000 --amount=1000000 --out="unsigned.txn"
```
# Saving Signed Transactions to File 
Signed Transactions are similar, but require an account to sign the transaction before writting it to a file.

``` javascript tab="JavaScript"
    let txn = {
        "from": myAccount.addr,
        "to": receiver,
        "fee": params.minFee,
        "flatFee": true,
        "amount": 1000000,
        "firstRound": params.lastRound,
        "lastRound": params.lastRound + 1000,
        "genesisID": params.genesisID,
        "genesisHash": params.genesishashb64
    };

    // sign the transaction and write to file 
    let signedTxn = algosdk.signTransaction(txn, myAccount.sk);
    fs.writeFileSync("./signed.stxn", signedTxn.blob);
```

``` python tab="Python"
	# create transaction
    receiver = "transaction-receiver<PLACEHOLDER>"
    data = {
        "sender": my_address,
        "receiver": receiver,
        "fee": params.get('minFee'),
        "flat_fee": True,
        "amt": amount<PLACEHOLDER>,
        "first": params.get('lastRound'),
        "last": params.get('lastRound') + 1000,
        "gen": params.get('genesisID'),
        "gh": params.get('genesishashb64')
    }
    txn = transaction.PaymentTxn(**data)

    # sign transaction
    signed_txn = txn.sign(private_key)

    # write to file
    dir_path = os.path.dirname(os.path.realpath(__file__))
    transaction.write_to_file([signed_txn], dir_path + "/signed.txn")
```

``` java tab="Java"
    try { 
        Transaction tx = new Transaction(new Address(SRC_ADDR),  
                BigInteger.valueOf(1000), firstRound, lastRound, 
                null, amount, new Address(DEST_ADDR), genId, genesisHash);

        // recover account and sign transaction
        String SRC_ACCOUNT = "25-word-passphrase<PLACEHOLDER>";
        Account src = new Account(SRC_ACCOUNT);
        SignedTransaction signedTx = src.signTransaction(tx);                    

        // save signed transaction to a file 
        FileOutputStream file = new FileOutputStream("./signed.stxn"); 
        ObjectOutputStream out = new ObjectOutputStream(file); 
        out.writeObject(signedTx);
        out.close(); 
        file.close();
    } catch (Exception e) { 
        System.out.println("Exception: " + e); 
    }
```

``` go tab="Go"
    tx, err := transaction.MakePaymentTxn(addr, toAddr, 1, 100000,
        txParams.LastRound, txParams.LastRound+100, nil, "", 
        genID, txParams.GenesisHash)
    if err != nil {
        fmt.Printf("Error creating transaction: %s\n", err)
        return
    }

    //Sign the Transaction
    txid, stx, err := crypto.SignTransaction(sk, tx)
    if err != nil {
        fmt.Printf("Failed to sign transaction: %s\n", err)
        return
    }
    //Save the signed transaction to file
    fmt.Printf("Made signed transaction with TxID %s: %x\n", txid, stx)
    file, err := os.Create("./signed.gob")
    if err == nil {
        encoder := gob.NewEncoder(file)
        encoder.Encode(stx)
    }
    file.Close()
```

``` goal tab="goal"
$ goal clerk send --from=my-account<PLACEHOLDER> --to=GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A --fee=1000 --amount=1000000 --out="signed.stxn" --sign
```

??? example "Complete Example = Saving Signed and Unsigned Transactions to a File"
    
    ```javascript tab="JavaScript"
    const algosdk = require('algosdk');
    const fs = require('fs');
    var client = null;
    // make connection to node
    async function setupClient() {
        
        if( client == null){
            const ALGOD_API_ADDR = "algod-address<PLACEHOLDER>";
            const ALGOD_API_TOKEN = "algod-token<PLACEHOLDER>";
            const port = port-number<PLACEHOLDER>;
            let algodClient = new algosdk.Algod(ALGOD_API_TOKEN, ALGOD_API_ADDR , port);
            client = algodClient;
        } else {
            return client;
        }
        return client;	

    }

    // recover acccount for example
    function recoverAccount(){
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
    async function writeUnsignedTransactionToFile() {

        // setup accounts and make node connection
        const receiver = "receiver-address<PLACEHOLDER>";
        let algodClient = await setupClient();
        let myAccount = await recoverAccount();
        console.log("My address: %s", myAccount.addr)

        // get the current params from the network
        let params = await algodClient.getTransactionParams();

        // setup transaction
        let txn = {
            "from": myAccount.addr,
            "to": receiver,
            "fee": params.minFee,
            "flatFee": true,
            "amount": 1000000,
            "firstRound": params.lastRound,
            "lastRound": params.lastRound + 1000,
            "genesisID": params.genesisID,
            "genesisHash": params.genesishashb64
        };
        
        // write the unsigned transaction to a file
        fs.writeFileSync('./unsigned.txn', JSON.stringify(txn) , 'utf-8');
    }; 
    async function readUnsignedTransactionFromFile() {

        // recover account and setup node connection
        let algodClient = await setupClient();
        let myAccount = await recoverAccount(); 
        console.log("My address: %s", myAccount.addr)
        
        // read transaction from file and sign
        let rawdata = fs.readFileSync('./unsigned.txn');
        let txn = JSON.parse(rawdata);    
        let signedTxn = algosdk.signTransaction(txn, myAccount.sk);
        
        let txId = signedTxn.txID;
        console.log("Signed transaction with txID: %s", txId);

        // send transaction to network
        await algodClient.sendRawTransaction(signedTxn.blob)

        // wait for transaction to be confirmed
        await waitForConfirmation(algodClient, txId)
    }; 
    async function writeSignedTransactionToFile() {

        // setup accounts and make node connection
        const receiver = "receiver-address<PLACEHOLDER>";
        let algodClient = await setupClient();
        let myAccount = await recoverAccount();
        console.log("My address: %s", myAccount.addr)

        // get the current params from the network
        let params = await algodClient.getTransactionParams();

        // setup transaction
        let txn = {
            "from": myAccount.addr,
            "to": receiver,
            "fee": params.minFee,
            "flatFee": true,
            "amount": 1000000,
            "firstRound": params.lastRound,
            "lastRound": params.lastRound + 1000,
            "genesisID": params.genesisID,
            "genesisHash": params.genesishashb64
        };

        // sign the transaction and write to file 
        let signedTxn = algosdk.signTransaction(txn, myAccount.sk);
        fs.writeFileSync("./signed.stxn", signedTxn.blob);

    }; 
    async function readSignedTransactionFromFile() {

        // setup node connection
        let algodClient = await setupClient();
        
        // read signed transaction from file 
        let stx = fs.readFileSync("./signed.stxn");
        const buf = Buffer.from(stx);
        let signedTxnBlob = Uint8Array.from(stx);

        // send signed transaction to the node
        let tx = await algodClient.sendRawTransaction(signedTxnBlob);
        console.log("Signed transaction with txID: %s", tx.txId);

        // wait for transaction to be confirmed
        await waitForConfirmation(algodClient, tx.txId)
    }; 
    // functions for testing
    async function testUnsigned(){
        await writeUnsignedTransactionToFile();
        await readUnsignedTransactionFromFile();
    }
    async function testSigned(){
        await writeSignedTransactionToFile();
        await readSignedTransactionFromFile();
    }
    // testUnsigned();
    testSigned();


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

    def connect_to_network():
        algod_address = "algod-address<PLACEHOLDER>"
        algod_token = "algod-token<PLACEHOLDER>"
        algod_client = algod.AlgodClient(algod_token, algod_address)
        return algod_client

    def wait_for_confirmation( algod_client, txid ):
        while True:
            txinfo = algod_client.pending_transaction_info(txid)
            if txinfo.get('round') and txinfo.get('round') > 0:
                print("Transaction {} confirmed in round {}.".format(txid, txinfo.get('round')))
                break
            else:
                print("Waiting for confirmation...")
                algod_client.status_after_block(algod_client.status().get('lastRound') +1)
            
    def write_unsigned():
        # setup none connection
        algod_client = connect_to_network()

        # recover account
        passphrase = "25-word-passphrase<PLACEHOLDER>"
        private_key = mnemonic.to_private_key(passphrase)
        my_address = mnemonic.to_public_key(passphrase)
        print("My address: {}".format(my_address))

        # get suggested parameters
        params = algod_client.suggested_params()

        # create transaction
        receiver = "transaction-receiver<PLACEHOLDER>"
        data = {
            "sender": my_address,
            "receiver": receiver,
            "fee": params.get('minFee'),
            "flat_fee": True,
            "amt": amount<PLACEHOLDER>,
            "first": params.get('lastRound'),
            "last": params.get('lastRound') + 1000,
            "gen": params.get('genesisID'),
            "gh": params.get('genesishashb64')
        }
        txn = transaction.PaymentTxn(**data)

        # write to file
        dir_path = os.path.dirname(os.path.realpath(__file__))
        transaction.write_to_file([txn], dir_path + "/unsigned.txn")


    def read_unsigned():
        
        # setup node connection
        algod_client = connect_to_network()
        
        # recover account
        passphrase = "25-word-passphrase"
        private_key = mnemonic.to_private_key(passphrase)
        my_address = mnemonic.to_public_key(passphrase)
        print("My address: {}".format(my_address))
        
        # read from file
        txns = transaction.retrieve_from_file("./unsigned.txn")

        # sign and submit transaction
        txn = txns[0]
        signed_txn = txn.sign(private_key)
        txid = signed_txn.transaction.get_txid()
        print("Signed transaction with txID: {}".format(txid))
        algod_client.send_transaction(signed_txn)

        # wait for confirmation
        wait_for_confirmation( algod_client, txid)

    def write_signed():
        
        # setup connection to node
        algod_client = connect_to_network()
        
        # recovere account
        passphrase = "25-word-passphrase<PLACEHOLDER>"
        private_key = mnemonic.to_private_key(passphrase)
        my_address = mnemonic.to_public_key(passphrase)
        print("My address: {}".format(my_address))
        
        # get node suggested parameters
        params = algod_client.suggested_params()
        
        # create transaction
        receiver = "transaction-receiver<PLACEHOLDER>"
        data = {
            "sender": my_address,
            "receiver": receiver,
            "fee": params.get('minFee'),
            "flat_fee": True,
            "amt": amount<PLACEHOLDER>,
            "first": params.get('lastRound'),
            "last": params.get('lastRound') + 1000,
            "gen": params.get('genesisID'),
            "gh": params.get('genesishashb64')
        }
        txn = transaction.PaymentTxn(**data)
        
        # sign transaction
        signed_txn = txn.sign(private_key)
        
        # write to file
        dir_path = os.path.dirname(os.path.realpath(__file__))
        transaction.write_to_file([signed_txn], dir_path + "/signed.txn")


    def read_signed():
        
        # set up connection to node
        algod_client = connect_to_network()
        
        # read signed transaction from file
        txns = transaction.retrieve_from_file("./signed.txn")
        signed_txn = txns[0]
        txid = signed_txn.transaction.get_txid()
        print("Signed transaction with txID: {}".format(txid))
        
        # send transaction to network
        algod_client.send_transaction(signed_txn)
    
        # wait for confirmation
        wait_for_confirmation( algod_client, txid)

    # Test Runs     
    #write_unsigned()
    #read_unsigned()
    write_signed()
    read_signed()    
    ```

    ```java tab="Java"
    package com.algorand.javatest;

    import java.io.FileInputStream;
    import java.io.FileOutputStream;
    import java.io.ObjectInputStream;
    import java.io.ObjectOutputStream;
    import java.math.BigInteger;

    import com.algorand.algosdk.account.Account;
    import com.algorand.algosdk.algod.client.AlgodClient;
    import com.algorand.algosdk.algod.client.api.AlgodApi;
    import com.algorand.algosdk.algod.client.auth.ApiKeyAuth;
    import com.algorand.algosdk.algod.client.model.TransactionID;
    import com.algorand.algosdk.algod.client.model.TransactionParams;
    import com.algorand.algosdk.crypto.Address;
    import com.algorand.algosdk.crypto.Digest;
    import com.algorand.algosdk.transaction.SignedTransaction;
    import com.algorand.algosdk.transaction.Transaction;
    import com.algorand.algosdk.util.Encoder;

    public class SaveTransactionOffline {
        public AlgodApi algodApiInstance = null;

        // utility function to connect to a node
        private AlgodApi connectToNetwork(){

            // Initialize an algod client
            final String ALGOD_API_ADDR = "algod-address<PLACEHOLDER>";
            final String ALGOD_API_TOKEN = "algod-token<PLACEHOLDER>";

            AlgodClient client = (AlgodClient) new AlgodClient().setBasePath(ALGOD_API_ADDR);
            ApiKeyAuth api_key = (ApiKeyAuth) client.getAuthentication("api_key");
            api_key.setApiKey(ALGOD_API_TOKEN);
            algodApiInstance = new AlgodApi(client);   
            return algodApiInstance;
        }
        // utility function to wait on a transaction to be confirmed    
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
        public void writeUnsignedTransaction(){

            if( algodApiInstance == null ) connectToNetwork();

            final String DEST_ADDR = "transaction-reciever<PLACEHOLDER>";
            final String SRC_ADDR = "transaction-sender<PLACEHOLDER>";

            //Save the Transaction to a file
            try { 
                
                // get last round and suggested tx fee
                BigInteger firstRound = BigInteger.valueOf(301);
                String genId = null;
                Digest genesisHash = null;

                // Get suggested parameters from the node
                TransactionParams params = algodApiInstance.transactionParams();

                //create transaction
                firstRound = params.getLastRound();
                genId = params.getGenesisID();
                genesisHash = new Digest(params.getGenesishashb64());
                BigInteger amount = BigInteger.valueOf(200000);
                BigInteger lastRound = firstRound.add(BigInteger.valueOf(1000)); 
                Transaction tx = new Transaction(new Address(SRC_ADDR),  
                        BigInteger.valueOf(1000), firstRound, lastRound, 
                        null, amount, new Address(DEST_ADDR), genId, genesisHash);

                //Save unsigned transaction to file 
                FileOutputStream file = new FileOutputStream("./unsigned.txn"); 
                ObjectOutputStream out = new ObjectOutputStream(file); 
                out.writeObject(tx);
                out.close(); 
                file.close();
                System.out.println("Transaction written to a file");
            } catch (Exception e) { 
                System.out.println("Save Exception: " + e); 
            }

        }
        public void readUnsignedTransaction(){

            Transaction tx = null;
            try {

                if( algodApiInstance == null ) connectToNetwork();

                //Reading  from a file 
                FileInputStream file = new FileInputStream("./unsigned.txn"); 
                ObjectInputStream in = new ObjectInputStream(file); 
                tx = (Transaction)in.readObject();
                in.close(); 
                file.close(); 

                // recover account and sign transaction
                String SRC_ACCOUNT = "25-word-passphrase<PLACEHOLDER>";
                Account src = new Account(SRC_ACCOUNT);
                SignedTransaction signedTx = src.signTransaction(tx);
                byte[] encodedTxBytes = Encoder.encodeToMsgPack(signedTx);

                //submit the encoded transaction to the network
                TransactionID id = algodApiInstance.rawTransaction(encodedTxBytes);
                System.out.println("Successfully sent tx with id: " + id);
                waitForConfirmation(id.getTxId());

            } catch (Exception e) {
                System.out.println("Submit Exception: " + e); 
            }


        }
        public void writeSignedTransaction(){

            if( algodApiInstance == null ) connectToNetwork();

            final String DEST_ADDR = "transaction-reciever<PLACEHOLDER>";
            final String SRC_ADDR = "transaction-sender<PLACEHOLDER>";

            //Save the Transaction to a file
            try { 
                
                // get last round and suggested tx fee
                BigInteger firstRound = BigInteger.valueOf(301);
                String genId = null;
                Digest genesisHash = null;

                // Get suggested parameters from the node
                TransactionParams params = algodApiInstance.transactionParams();

                // create transaction
                firstRound = params.getLastRound();
                genId = params.getGenesisID();
                genesisHash = new Digest(params.getGenesishashb64());
                BigInteger amount = BigInteger.valueOf(200000);
                BigInteger lastRound = firstRound.add(BigInteger.valueOf(1000));  
                Transaction tx = new Transaction(new Address(SRC_ADDR),  
                        BigInteger.valueOf(1000), firstRound, lastRound, 
                        null, amount, new Address(DEST_ADDR), genId, genesisHash);
    
                // recover account and sign transaction
                String SRC_ACCOUNT = "25-word-passphrase<PLACEHOLDER>";
                Account src = new Account(SRC_ACCOUNT);
                SignedTransaction signedTx = src.signTransaction(tx);                    
                
                // save signed transaction to a file 
                FileOutputStream file = new FileOutputStream("./signed.stxn"); 
                ObjectOutputStream out = new ObjectOutputStream(file); 
                out.writeObject(signedTx);
                out.close(); 
                file.close();
                System.out.println("Transaction Signed and written to a file");
            } catch (Exception e) { 
                System.out.println("Save Exception: " + e); 
            }

        }

        public void readSignedTransaction(){

            try {

                if( algodApiInstance == null ) connectToNetwork();

                //Read the transaction from a file 
                FileInputStream file = new FileInputStream("./signed.stxn"); 
                ObjectInputStream in = new ObjectInputStream(file); 
                SignedTransaction signedTx = (SignedTransaction)in.readObject();
                in.close(); 
                file.close(); 
                System.out.println("Signed transaction with txid: " + signedTx.transactionID);     
                    
                // Msgpack encode the signed transaction
                byte[] encodedTxBytes = Encoder.encodeToMsgPack(signedTx);
                
                //submit the encoded transaction to the network
                TransactionID id = algodApiInstance.rawTransaction(encodedTxBytes);
                System.out.println("Successfully sent tx with id: " + id); 
                waitForConfirmation(id.getTxId());

            } catch (Exception e) {
                System.out.println("Submit Exception: " + e); 
            }


        }
        public static void main(String args[]) throws Exception {
            SaveTxStxForOffline mn = new SaveTxStxForOffline();
            mn.writeUnsignedTransaction();
            mn.readUnsignedTransaction();

            //mn.writeSignedTransaction();
            //mn.readSignedTransaction();

        }

    }

    ```

    ```go tab="Go"
    package main

    import (
        "encoding/gob"
        "fmt"
        "os"

        "golang.org/x/crypto/ed25519"

        "github.com/algorand/go-algorand-sdk/client/algod"
        "github.com/algorand/go-algorand-sdk/crypto"
        "github.com/algorand/go-algorand-sdk/mnemonic"
        "github.com/algorand/go-algorand-sdk/transaction"
        "github.com/algorand/go-algorand-sdk/types"
    )
    // Function that waits for a given txId to be confirmed by the network
    func waitForConfirmation(algodClient algod.Client, txID string) {
        for {
            pendingInfo, err := algodClient.PendingTransactionInformation(txID)
            if err != nil {
                fmt.Printf("waiting for confirmation... (pool error, if any): %s\n", err)
                continue
            }
            if pendingInfo.ConfirmedRound > 0 {
                fmt.Printf("Transaction "+pendingInfo.TxID+" confirmed in round %d\n", pendingInfo.ConfirmedRound)
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
    // utility function to recover account and return sk and address
    func recoverAccount()(string, ed25519.PrivateKey) {
        const passphrase = "your-25-word-mnemonic<PLACEHOLDER>"

        sk, err := mnemonic.ToPrivateKey(passphrase)
        if err != nil {
            fmt.Printf("error recovering account: %s\n", err)
            return "", nil
        }
        pk := sk.Public()
        var a types.Address
        cpk := pk.(ed25519.PublicKey)
        copy(a[:], cpk[:])
        fmt.Printf("Address: %s\n", a.String())	
        address := a.String()
        return address, sk 
    }
    // utility funciton to setup connection to node
    func setupConnection()( algod.Client ){
        const algodToken = "algod-token<PLACEHOLDER>"
        const algodAddress = "algod-address<PLACEHOLDER>"
        algodClient, err := algod.MakeClient(algodAddress, algodToken)
        if err != nil {
            fmt.Printf("failed to make algod client: %s\n", err)
        }
        return algodClient
    }

    func saveUnsignedTransaction() {

        // setup connection
        algodClient := setupConnection()
        
        // recover account for example
        addr, _ := recoverAccount();

        // get network suggested parameters
        txParams, err := algodClient.SuggestedParams()
        if err != nil {
            fmt.Printf("error getting suggested tx params: %s\n", err)
            return
        }

        // create transaction
        toAddr := "transaction-receiver<PLACEHOLDER>"
        genID := txParams.GenesisID
        tx, err := transaction.MakePaymentTxn(addr, toAddr, 1, 100000,
            txParams.LastRound, txParams.LastRound+100, nil, "", 
            genID, txParams.GenesisHash)
        if err != nil {
            fmt.Printf("Error creating transaction: %s\n", err)
            return
        }

        // save unsigned transction to file
        file, err := os.Create("./unsigned.gob")
        if err == nil {
            encoder := gob.NewEncoder(file)
            encoder.Encode(tx)
        }
        file.Close()
        if err == nil {
            fmt.Printf("Saved unsigned transaction to file\n")
            return
        }
        fmt.Printf("Failed in saving trx to file, error %s\n", err)

    }
    func readUnsignedTransaction(){

        // setup connection
        algodClient := setupConnection()

        // read unsigned transaction from file
        var tx types.Transaction
        file, err := os.Open("./unsigned.gob")
        if err == nil {
            decoder := gob.NewDecoder(file)
            err := decoder.Decode(&tx)
            if err != nil {
                fmt.Printf("Error reading transaction from file: %s\n", err)
                return
            }
        }else{
            fmt.Printf("failed to open signed transaction: %s\n", err)
            return	
        }
        file.Close()

        // recover account and sign transaction
        addr, sk := recoverAccount();
        fmt.Printf("Address is: %s\n", addr)
        txid, stx, err := crypto.SignTransaction(sk, tx)
        if err != nil {
            fmt.Printf("Failed to sign transaction: %s\n", err)
            return
        }
        fmt.Printf("Transaction id: %s\n", txid)

        // send transaction to the network
        sendResponse, err := algodClient.SendRawTransaction(stx)
        if err != nil {
            fmt.Printf("failed to send transaction: %s\n", err)
            return
        }
        fmt.Printf("Transaction ID: %s\n", sendResponse.TxID)
        waitForConfirmation(algodClient, sendResponse.TxID)
    }


    func saveSignedTransaction() {

        // setup connection
        algodClient := setupConnection()

        // recover account
        addr, sk := recoverAccount();

        // get network suggested parameters
        txParams, err := algodClient.SuggestedParams()
        if err != nil {
            fmt.Printf("error getting suggested tx params: %s\n", err)
            return
        }

        // create transaction
        toAddr := "transaction-receiver<PLACEHOLDER>"
        genID := txParams.GenesisID
        tx, err := transaction.MakePaymentTxn(addr, toAddr, 1, 100000,
            txParams.LastRound, txParams.LastRound+100, nil, "", 
            genID, txParams.GenesisHash)
        if err != nil {
            fmt.Printf("Error creating transaction: %s\n", err)
            return
        }

        //Sign the Transaction
        txid, stx, err := crypto.SignTransaction(sk, tx)
        if err != nil {
            fmt.Printf("Failed to sign transaction: %s\n", err)
            return
        }
        //Save the signed transaction to file
        fmt.Printf("Made signed transaction with TxID %s: %x\n", txid, stx)
        file, err := os.Create("./signed.gob")
        if err == nil {
            encoder := gob.NewEncoder(file)
            encoder.Encode(stx)
        }
        file.Close()
        if err == nil {
            fmt.Printf("Saved signed transaction to file\n")
            return
        }
        fmt.Printf("Failed in saving trx to file, error %s\n", err)

    }
    func readSignedTransaction(){

        // setup connection
        algodClient := setupConnection()

        // read signed transaction from file
        var signedTransaction []byte
        file, err := os.Open("./signed.gob")
        if err == nil {
            decoder := gob.NewDecoder(file)
            err = decoder.Decode(&signedTransaction)
        }else{
            fmt.Printf("failed to open signed transaction: %s\n", err)
            return	
        }
        file.Close()

        // send the transaction to the network
        sendResponse, err := algodClient.SendRawTransaction(signedTransaction)
        if err != nil {
            fmt.Printf("failed to send transaction: %s\n", err)
            return
        }

        fmt.Printf("Transaction ID: %s\n", sendResponse.TxID)
        waitForConfirmation(algodClient, sendResponse.TxID)
    }
    func main() {
        saveUnsignedTransaction()
        readUnsignedTransaction()

        //saveSignedTransaction()
        //readSignedTransaction()

    }    
    ```
