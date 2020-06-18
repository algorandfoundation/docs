const algosdk = require('algosdk');

// function used to wait for a tx confirmation
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

async function yourFirstTransaction() {

    try{        
        // Generate a public/private key pair
        const passphrase = "Your 25-word mnemonic generated and displayed above";
        let myAccount = algosdk.mnemonicToSecretKey(passphrase);
        console.log("My address: %s", myAccount.addr);
        console.log( "My passphrase: " + passphrase );

        // Add funds
        // TestNet Faucet: https://bank.testnet.algorand.network/
        // BetaNet Faucet: https://bank.betanet.algodev.network/

        // Connect your client
        const algodToken = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
        const algodServer = "http://localhost";
        const algodPort = 4001;

        let algodClient = new algosdk.Algodv2(algodToken, algodServer, algodPort);
        
        //Check your balance
        let accountInfo = await algodClient.accountInformation(myAccount.addr).do();
        console.log("Account balance: %d microAlgos", accountInfo.amount);

        // Construct the transaction
        let params = await algodClient.getTransactionParams().do();
        // comment out the next two lines to use suggested fee
        params.fee = 1000;
        params.flatFee = true;

        // receiver defined as TestNet faucet address 
        const receiver = "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A";
        let note = algosdk.encodeObj("Hello World");

        let txn = algosdk.makePaymentTxnWithSuggestedParams(myAccount.addr, receiver, 1000000, undefined, note, params);       

        // Sign the transaction
        let signedTxn = txn.signTxn(myAccount.sk);
        let txId = txn.txID().toString();
        console.log("Signed transaction with txID: %s", txId);

        // Submit the transaction
        await algodClient.sendRawTransaction(signedTxn).do();

        // Wait for confirmation
        await waitForConfirmation(algodClient, txId);

        // Read the transaction from the blockchain
        let confirmedTxn = await algodClient.pendingTransactionInformation(txId).do();
        console.log("Transaction information: %o", confirmedTxn.txn.txn);
        console.log("Decoded note: %s", algosdk.decodeObj(confirmedTxn.txn.txn.note));
    }
    catch (err){
        console.log("err", err);  
    }
};
yourFirstTransaction();