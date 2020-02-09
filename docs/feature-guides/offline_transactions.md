title: Authorizing Transactions Offline

This section explains how to authorize transactions with private keys that are kept **offline**. In particular, this guide shows how to create and save transactions to a file that can then be transferred to an offline device for signing. To learn how to construct and authorize transactions in general visit the [Constructing Transactions](./transactions.md) and [Authorizing Transactions](./signatures.md) guides, respectively.

The same methodology described here can also be used to work with [LogicSignatures](asc1/modes.md#logic-signatures) and [Multisignatures](signatures.md#multisignatures). All objects in the following examples use msgpack to store the transaction object ensuring interoperability with the SDKs and `goal`.

!!! info
    Storing keys _offline_ is also referred to as placing them in **cold storage**. An _online_ device that stores private keys is often referred to as a **hot wallet**.  

# Unsigned Transaction File Operations
Algorand SDK's and `goal` support writing and reading both signed and unsigned transactions to a file. Examples of these scenarios are shown in the following code snippets.

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
	// Save transaction to file
    fs.writeFileSync('./unsigned.txn', algosdk.encodeObj( txn ));
    
	// read transaction from file and sign it
    let txn = algosdk.decodeObj(fs.readFileSync('./unsigned.txn')); 
	let signedTxn = algosdk.signTransaction(txn, myAccount.sk);
	let txId = signedTxn.txID;
	
	// send signed transaction to node
	await algodClient.sendRawTransaction(signedTxn.blob);         
```

``` python tab="Python"
	# create transaction
	receiver = <transaction-receiver>
	data = {
		"sender": my_address,
		"receiver": receiver,
		"fee": params.get('minFee'),
		"flat_fee": True,
		"amt": <amount>,
		"first": params.get('lastRound'),
		"last": params.get('lastRound') + 1000,
		"gen": params.get('genesisID'),
		"gh": params.get('genesishashb64')
	}
	txn = transaction.PaymentTxn(**data)

	# write to file
	dir_path = os.path.dirname(os.path.realpath(__file__))
	transaction.write_to_file([txn], dir_path + "/unsigned.txn")

	# read from file
	txns = transaction.retrieve_from_file("./unsigned.txn")

	# sign and submit transaction
	txn = txns[0]
	signed_txn = txn.sign(private_key)
	txid = signed_txn.transaction.get_txid()
	algod_client.send_transaction(signed_txn)    
```

``` java tab="Java"
    BigInteger amount = BigInteger.valueOf(200000);
    BigInteger lastRound = firstRound.add(BigInteger.valueOf(1000));  
    Transaction tx = new Transaction(new Address(SRC_ADDR),  
            BigInteger.valueOf(1000), firstRound, lastRound, 
            null, amount, new Address(DEST_ADDR), genId, genesisHash);
    // save as signed even though it has not been
    SignedTransaction stx = new SignedTransaction();
    stx.tx = tx;  
    // Save transaction to a file 
    Files.write(Paths.get("./unsigned.txn"), Encoder.encodeToMsgPack(stx));

    // read transaction from file
    SignedTransaction decodedTransaction = Encoder.decodeFromMsgPack(
        Files.readAllBytes(Paths.get("./unsigned.txn")), 
        SignedTransaction.class);            
    Transaction tx = decodedTransaction.tx;          

    // recover account    
    String SRC_ACCOUNT = <25-word-passphrase>;
    Account src = new Account(SRC_ACCOUNT);

    // sign transaction
    SignedTransaction signedTx = src.signTransaction(tx);
    byte[] encodedTxBytes = Encoder.encodeToMsgPack(signedTx);
            
    // submit the encoded transaction to the network
    TransactionID id = algodApiInstance.rawTransaction(encodedTxBytes);
```

``` go tab="Go"
	tx, err := transaction.MakePaymentTxn(addr, toAddr, 1, 100000,
		 txParams.LastRound, txParams.LastRound+100, nil, "", 
		 genID, txParams.GenesisHash)
	if err != nil {
		fmt.Printf("Error creating transaction: %s\n", err)
		return
    }
    // save as signed tx object without sig
    unsignedTx := types.SignedTxn{
		Txn:  tx,
	 }

	// save unsigned Transaction to file
	err = ioutil.WriteFile("./unsigned.txn", msgpack.Encode(unsignedTx), 0644)
	if err == nil {
		fmt.Printf("Saved unsigned transaction to file\n")
		return
    }

    // read unsigned transaction from file
	dat, err := ioutil.ReadFile("./unsigned.txn")
	if err != nil {
		fmt.Printf("Error reading transaction from file: %s\n", err)
		return
	}
	var unsignedTxRaw types.SignedTxn 
	var unsignedTxn types.Transaction

	msgpack.Decode(dat, &unsignedTxRaw)

    unsignedTxn = unsignedTxRaw.Txn

	// recover account and sign transaction
	addr, sk := recoverAccount();
	fmt.Printf("Address is: %s\n", addr)
	txid, stx, err := crypto.SignTransaction(sk, unsignedTxn)
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
```


``` goal tab="goal"
$ goal clerk send --from=<my-account> --to=GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A --fee=1000 --amount=1000000 --out="unsigned.txn"

$ goal clerk sign --infile unsigned.txn --outfile signed.txn

$ goal clerk rawsend --filename signed.txn

```
# Signed Transaction File Operations 
Signed Transactions are similar, but require an account to sign the transaction before writing it to a file.

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

	// sign transaction and write to file
	let signedTxn = algosdk.signTransaction(txn, myAccount.sk);
    fs.writeFileSync('./signed.stxn', algosdk.encodeObj( signedTxn ));
    
	// read signed transaction from file
	let stx = algosdk.decodeObj(fs.readFileSync("./signed.stxn"));
		
	// send signed transaction to node
	let tx = await algodClient.sendRawTransaction(stx.blob);    
```

``` python tab="Python"
	# create transaction
    receiver = <transaction-receiver>
    data = {
        "sender": my_address,
        "receiver": receiver,
        "fee": params.get('minFee'),
        "flat_fee": True,
        "amt": <amount>,
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

	# read signed transaction from file
	txns = transaction.retrieve_from_file("./signed.txn")
	signed_txn = txns[0]
	txid = signed_txn.transaction.get_txid()
	print("Signed transaction with txID: {}".format(txid))
	
	# send transaction to network
	algod_client.send_transaction(signed_txn)    
```

``` java tab="Java"
    // create transaction 
    BigInteger amount = BigInteger.valueOf(200000);
    BigInteger lastRound = firstRound.add(BigInteger.valueOf(1000));  
    Transaction tx = new Transaction(new Address(SRC_ADDR),  
            BigInteger.valueOf(1000), firstRound, lastRound, 
            null, amount, new Address(DEST_ADDR), genId, genesisHash);

    // recover account    
    String SRC_ACCOUNT = <25-word-passphrase>;                    
    Account src = new Account(SRC_ACCOUNT);

    // sign transaction
    SignedTransaction signedTx = src.signTransaction(tx);                    

    // save signed transaction to  a file 
    Files.write(Paths.get("./signed.txn"), Encoder.encodeToMsgPack(signedTx));

    //Read the transaction from a file 
    SignedTransaction decodedSignedTransaction = Encoder.decodeFromMsgPack(
        Files.readAllBytes(Paths.get("./signed.txn")), SignedTransaction.class);   
    System.out.println("Signed transaction with txid: " + decodedSignedTransaction.transactionID);           

    // Msgpack encode the signed transaction
    byte[] encodedTxBytes = Encoder.encodeToMsgPack(decodedSignedTransaction);

    //submit the encoded transaction to the network
    TransactionID id = algodApiInstance.rawTransaction(encodedTxBytes);    
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
	fmt.Printf("Made signed transaction with TxID %s: %x\n", txid, stx)

	//Save the signed transaction to file
	err = ioutil.WriteFile("./signed.stxn", msgpack.Encode(stx), 0644)
	if err == nil {
		fmt.Printf("Saved signed transaction to file\n")
		return
    }
    
	// read unsigned transaction from file
	dat, err := ioutil.ReadFile("./signed.stxn")
	if err != nil {
		fmt.Printf("Error reading signed transaction from file: %s\n", err)
		return
	}
	var signedTx []byte 
	msgpack.Decode(dat, &signedTx)
	
	// send the transaction to the network
	sendResponse, err := algodClient.SendRawTransaction(signedTx)
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}

```

``` goal tab="goal"
$ goal clerk rawsend --filename signed.txn
```

??? example "Complete Example = Saving Signed and Unsigned Transactions to a File"
    
    ```javascript tab="JavaScript"
    const algosdk = require('algosdk');
    const fs = require('fs');
    var client = null;
    // make connection to node
    async function setupClient() {
        if( client == null){
            const server = <algod-address>;
            const token = <algod-token>;
            const port = <port-number>;
            let algodClient = new algosdk.Algod(token, server, port);
            client = algodClient;
        } else {
            return client;
        }
        return client;
    }
    // recover account for example
    function recoverAccount(){
        const passphrase =<your-25-word-mnemonic>;
        let myAccount = algosdk.mnemonicToSecretKey(passphrase);
        return myAccount;
    }
    // Function used to wait for a tx confirmation
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

        try{
            const receiver = <transaction-receiver>;

            // setup accounts and make node connection
            let algodClient = await setupClient();

            // recover account
            let myAccount = await recoverAccount();
            console.log("My address: %s", myAccount.addr)

            // get network suggested parameters
            let params = await algodClient.getTransactionParams();
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
            // Save transaction to file
            fs.writeFileSync('./unsigned.txn', algosdk.encodeObj( txn ));	
        }catch( e ){
            console.log( e );
        }
    }; 
    async function readUnsignedTransactionFromFile() {

        try{
            // setup connection to node
            let algodClient = await setupClient();

            // recover account
            let myAccount = await recoverAccount(); 
            console.log("My address: %s", myAccount.addr)

            // read transaction from file and sign it
            let txn = algosdk.decodeObj(fs.readFileSync('./unsigned.txn'));  
            let signedTxn = algosdk.signTransaction(txn, myAccount.sk);
            let txId = signedTxn.txID;
            console.log("Signed transaction with txID: %s", txId);

            // send signed transaction to node
            await algodClient.sendRawTransaction(signedTxn.blob);

            // Wait for transaction to be confirmed
            await waitForConfirmation(algodClient, txId);
        } catch ( e ){
            console.log( e );
        }	
    }; 
    async function writeSignedTransactionToFile() {

        try{
            const receiver = <transaction-receiver>;

            // setup connection to node
            let algodClient = await setupClient();
            let myAccount = await recoverAccount();
            console.log("My address: %s", myAccount.addr)

            // get network suggested parameters
            let params = await algodClient.getTransactionParams();

            // setup a transaction
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

            // sign transaction and write to file
            let signedTxn = algosdk.signTransaction(txn, myAccount.sk);
            fs.writeFileSync('./signed.stxn', algosdk.encodeObj( signedTxn ));	
        } catch( e ) {
            console.log(e);
        }
    }; 
    async function readSignedTransactionFromFile() {

        try{
            // setup connection to node
            let algodClient = await setupClient();
            
            // read signed transaction from file
            let stx = algosdk.decodeObj(fs.readFileSync("./signed.stxn"));
            
            // send signed transaction to node
            let tx = await algodClient.sendRawTransaction(stx.blob);
            console.log("Signed transaction with txID: %s", tx.txId);

            // Wait for transaction to be confirmed
            await waitForConfirmation(algodClient, tx.txId);
        } catch( e ) {
            console.log(e);
        }	
    }; 

    async function testUnsigned(){
        await writeUnsignedTransactionToFile();
        await readUnsignedTransactionFromFile();
    }
    async function testSigned(){
        await writeSignedTransactionToFile();
        await readSignedTransactionFromFile();
    }
    //testUnsigned();
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
        algod_address = <algod-address>
        algod_token = <algod-token>
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
        passphrase = <25-word-passphrase>
        private_key = mnemonic.to_private_key(passphrase)
        my_address = mnemonic.to_public_key(passphrase)
        print("My address: {}".format(my_address))

        # get suggested parameters
        params = algod_client.suggested_params()

        # create transaction
        receiver = <transaction-receiver>
        data = {
            "sender": my_address,
            "receiver": receiver,
            "fee": params.get('minFee'),
            "flat_fee": True,
            "amt": <amount>,
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
        passphrase = <25-word-passphrase>
        private_key = mnemonic.to_private_key(passphrase)
        my_address = mnemonic.to_public_key(passphrase)
        print("My address: {}".format(my_address))
        
        # get node suggested parameters
        params = algod_client.suggested_params()
        
        # create transaction
        receiver = <transaction-receiver>
        data = {
            "sender": my_address,
            "receiver": receiver,
            "fee": params.get('minFee'),
            "flat_fee": True,
            "amt": <amount>,
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

    import java.math.BigInteger;
    import java.nio.file.Files;
    import java.nio.file.Paths;

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
            final String ALGOD_API_ADDR = <algod-address>;
            final String ALGOD_API_TOKEN = <algod-token>;

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

            // connect to node
            if( algodApiInstance == null ) connectToNetwork();

            final String DEST_ADDR = <transaction-reciever>;
            final String SRC_ADDR = <transaction-sender>;

            try { 
                // Get suggested parameters from the node
                TransactionParams params = algodApiInstance.transactionParams();                     
                BigInteger firstRound = params.getLastRound();
                String genId = params.getGenesisID();
                Digest genesisHash = new Digest(params.getGenesishashb64());

                // create transaction
                BigInteger amount = BigInteger.valueOf(200000);
                BigInteger lastRound = firstRound.add(BigInteger.valueOf(1000));  
                Transaction tx = new Transaction(new Address(SRC_ADDR),  
                        BigInteger.valueOf(1000), firstRound, lastRound, 
                        null, amount, new Address(DEST_ADDR), genId, genesisHash);
                // save as signed even though it has not been
                SignedTransaction stx = new SignedTransaction();
                stx.tx = tx;  
                // Save transaction to a file 
                Files.write(Paths.get("./unsigned.txn"), Encoder.encodeToMsgPack(stx));
                System.out.println("Transaction written to a file");
            } catch (Exception e) { 
                System.out.println("Save Exception: " + e); 
            }

        }
        public void readUnsignedTransaction(){

            try {
                // connect to node
                if( algodApiInstance == null ) connectToNetwork();

                // read transaction from file
                SignedTransaction decodedTransaction = Encoder.decodeFromMsgPack(
                    Files.readAllBytes(Paths.get("./unsigned.txn")), SignedTransaction.class);            
                Transaction tx = decodedTransaction.tx;           

                // recover account    
                String SRC_ACCOUNT = <25-word-passphrase>;
                Account src = new Account(SRC_ACCOUNT);

                // sign transaction
                SignedTransaction signedTx = src.signTransaction(tx);
                byte[] encodedTxBytes = Encoder.encodeToMsgPack(signedTx);
                
                // submit the encoded transaction to the network
                TransactionID id = algodApiInstance.rawTransaction(encodedTxBytes);
                System.out.println("Successfully sent tx with id: " + id);
                waitForConfirmation(id.getTxId());

            } catch (Exception e) {
                System.out.println("Submit Exception: " + e); 
            }


        }
        public void writeSignedTransaction(){

            // connect to node
            if( algodApiInstance == null ) connectToNetwork();

            final String DEST_ADDR = <transaction-reciever>;
            final String SRC_ADDR = <transaction-sender>;;

            try { 

                // Get suggested parameters from the node
                TransactionParams params = algodApiInstance.transactionParams();
                BigInteger firstRound = params.getLastRound();
                String genId = params.getGenesisID();
                Digest genesisHash = new Digest(params.getGenesishashb64());

                // create transaction 
                BigInteger amount = BigInteger.valueOf(200000);
                BigInteger lastRound = firstRound.add(BigInteger.valueOf(1000));  
                Transaction tx = new Transaction(new Address(SRC_ADDR),  
                        BigInteger.valueOf(1000), firstRound, lastRound, 
                        null, amount, new Address(DEST_ADDR), genId, genesisHash);

                // recover account    
                String SRC_ACCOUNT = <25-word-passphrase>;                    
                Account src = new Account(SRC_ACCOUNT);

                // sign transaction
                SignedTransaction signedTx = src.signTransaction(tx);                    

                // save signed transaction to  a file 
                Files.write(Paths.get("./signed.txn"), Encoder.encodeToMsgPack(signedTx));
            } catch (Exception e) { 
                System.out.println("Save Exception: " + e); 
            }

        }

        public void readSignedTransaction(){

            try {
                // connect to a node
                if( algodApiInstance == null ) connectToNetwork();

                //Read the transaction from a file 
                SignedTransaction decodedSignedTransaction = Encoder.decodeFromMsgPack(
                    Files.readAllBytes(Paths.get("./signed.txn")), SignedTransaction.class);   
                System.out.println("Signed transaction with txid: " + decodedSignedTransaction.transactionID);           

                // Msgpack encode the signed transaction
                byte[] encodedTxBytes = Encoder.encodeToMsgPack(decodedSignedTransaction);

                //submit the encoded transaction to the network
                TransactionID id = algodApiInstance.rawTransaction(encodedTxBytes);
                System.out.println("Successfully sent tx with id: " + id); 
                waitForConfirmation(id.getTxId());

            } catch (Exception e) {
                System.out.println("Submit Exception: " + e); 
            }


        }
        public static void main(String args[]) throws Exception {
            SaveTransactionOffline mn = new SaveTransactionOffline();
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
        "fmt"
        "io/ioutil"

        "golang.org/x/crypto/ed25519"

        "github.com/algorand/go-algorand-sdk/client/algod"
        "github.com/algorand/go-algorand-sdk/crypto"
        "github.com/algorand/go-algorand-sdk/mnemonic"
        "github.com/algorand/go-algorand-sdk/encoding/msgpack"
        "github.com/algorand/go-algorand-sdk/transaction"
        "github.com/algorand/go-algorand-sdk/types"
    )

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
    // utility function to recover account and return sk and address
    func recoverAccount()(string, ed25519.PrivateKey) {
        const passphrase = <your-25-word-mnemonic>

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
        const algodToken = <algod-token>
        const algodAddress = <algod-address>
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
        toAddr := <transaction-receiver>
        genID := txParams.GenesisID
        tx, err := transaction.MakePaymentTxn(addr, toAddr, 1, 100000,
            txParams.LastRound, txParams.LastRound+100, nil, "", 
            genID, txParams.GenesisHash)
        if err != nil {
            fmt.Printf("Error creating transaction: %s\n", err)
            return
        }
	    unsignedTx := types.SignedTxn{
		    Txn:  tx,
	    }        

        // save unsigned Transaction to file
        err = ioutil.WriteFile("./unsigned.txn", msgpack.Encode(unsignedTx), 0644)
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
        dat, err := ioutil.ReadFile("./unsigned.txn")
        if err != nil {
            fmt.Printf("Error reading transaction from file: %s\n", err)
            return
        }
	    var unsignedTxRaw types.SignedTxn 
	    var unsignedTxn types.Transaction

	    msgpack.Decode(dat, &unsignedTxRaw)

	    unsignedTxn = unsignedTxRaw.Txn

        // recover account and sign transaction
        addr, sk := recoverAccount();
        fmt.Printf("Address is: %s\n", addr)
        txid, stx, err := crypto.SignTransaction(sk, unsignedTxn)
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
        toAddr := <transaction-receiver>
        genID := txParams.GenesisID
        tx, err := transaction.MakePaymentTxn(addr, toAddr, 1, 100000,
            txParams.LastRound, txParams.LastRound+100, nil, "", 
            genID, txParams.GenesisHash)
        if err != nil {
            fmt.Printf("Error creating transaction: %s\n", err)
            return
        }

        // sign the Transaction, msgpack encoding happens in sign
        txid, stx, err := crypto.SignTransaction(sk, tx)
        if err != nil {
            fmt.Printf("Failed to sign transaction: %s\n", err)
            return
        }
        fmt.Printf("Made signed transaction with TxID %s: %x\n", txid, stx)

        //Save the signed transaction to file
        err = ioutil.WriteFile("./signed.stxn", stx, 0644)
        if err == nil {
            fmt.Printf("Saved signed transaction to file\n")
            return
        }
        fmt.Printf("Failed in saving trx to file, error %s\n", err)

    }
    func readSignedTransaction(){

        // setup connection
        algodClient := setupConnection()

        // read unsigned transaction from file
        dat, err := ioutil.ReadFile("./signed.stxn")
        if err != nil {
            fmt.Printf("Error reading signed transaction from file: %s\n", err)
            return
        }
        
        // send the transaction to the network
        sendResponse, err := algodClient.SendRawTransaction(dat)
        if err != nil {
            fmt.Printf("failed to send transaction: %s\n", err)
            return
        }

        fmt.Printf("Transaction ID: %s\n", sendResponse.TxID)
        waitForConfirmation(algodClient, sendResponse.TxID)
    }
    func main() {
        //saveUnsignedTransaction()
        //readUnsignedTransaction()

        saveSignedTransaction()
        readSignedTransaction()

    }    
    ```
