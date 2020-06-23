const algosdk = require('algosdk');
const fs = require('fs');

var client = null;
async function setupClient() {
    if( client == null){
        const ALGOD_API_ADDR = "http://localhost:4001";
        const ALGOD_API_TOKEN = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
        const port = 4001;
        let algodClient = new algosdk.Algod(token, server, port);
        client = algodClient;
    } else {
        return client;
    }
    return client;
}
// recover first account
function recoverAccount1(){
    const passphrase = "Your 25-word mnemonic goes here";
    let myAccount = algosdk.mnemonicToSecretKey(passphrase);
    return myAccount;
}
// recover second account
function recoverAccount2(){
    const passphrase = "Your 25-word mnemonic goes here";
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
        const receiver = "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A";
        // sample show account A to C
        // B to A 
        // grouped
        let algodClient = await setupClient();

        // Creat transaction A to C and write to a file
        await writeUnsignedTransactionToFile();

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