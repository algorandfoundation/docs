const algosdk = require('algosdk');

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

async function yourFirstTransaction() {

    try{
        const algodToken = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
        const algodServer = "http://localhost";
        const algodPort = 4001;

        let algodClient = new algosdk.Algod(algodToken, algodServer, algodPort);

        const receiver = "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A";    
        const passphrase = <your-25-word-mnemonic>;

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
        await waitForConfirmation(algodClient, txId);
    }catch (err){
        console.log("err", err);  
    }
};
yourFirstTransaction();