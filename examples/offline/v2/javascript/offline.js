const algosdk = require('algosdk');
const fs = require('fs');
var client = null;
// make connection to node
async function setupClient() {
    if( client == null){
        const token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
        const server = "http://localhost";
        const port = 4001;
        let algodClient = new algosdk.Algodv2(token, server, port);
        client = algodClient;
    } else {
        return client;
    }
    return client;
}
// recover account for example
function recoverAccount(){
    const passphrase = "liquid million govern habit nasty danger spoil air monitor lobster solar misery confirm problem tuna hollow ritual assume mean return enrich mistake seven abstract tent";
    let myAccount = algosdk.mnemonicToSecretKey(passphrase);
    return myAccount;
}

/**
 * utility function to wait on a transaction to be confirmed
 * the timeout parameter indicates how many rounds do you wish to check pending transactions for
 */
 const waitForConfirmation = async function (algodclient, txId, timeout) {
    // Wait until the transaction is confirmed or rejected, or until 'timeout'
    // number of rounds have passed.
    //     Args:
    // txId(str): the transaction to wait for
    // timeout(int): maximum number of rounds to wait
    // Returns:
    // pending transaction information, or throws an error if the transaction
    // is not confirmed or rejected in the next timeout rounds
    if (algodclient == null || txId == null || timeout < 0) {
        throw "Bad arguments.";
    }
    let status = (await algodclient.status().do());
    if (status == undefined) throw new Error("Unable to get node status");
    let startround = status["last-round"] + 1;   
    let currentround = startround;

    while (currentround < (startround + timeout)) {
        let pendingInfo = await algodclient.pendingTransactionInformation(txId).do();      
        if (pendingInfo != undefined) {
            if (pendingInfo["confirmed-round"] !== null && pendingInfo["confirmed-round"] > 0) {
                //Got the completed Transaction
                return pendingInfo;
            }
            else {
                if (pendingInfo["pool-error"] != null && pendingInfo["pool-error"].length > 0) {
                    // If there was a pool error, then the transaction has been rejected!
                    throw new Error("Transaction Rejected" + " pool error" + pendingInfo["pool-error"]);
                }
            }
        } 
        await algodclient.statusAfterBlock(currentround).do();
        currentround++;
    }
    throw new Error("Transaction not confirmed after " + timeout + " rounds!");
};

async function writeUnsignedTransactionToFile() {

    try{
        const receiver = "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A";

        // setup accounts and make node connection
        let algodClient = await setupClient();
        // recover account
        let myAccount = await recoverAccount();
        console.log("My address: %s", myAccount.addr)

        // get network suggested parameters
        let params = await algodClient.getTransactionParams().do();
        const enc = new TextEncoder();
        const note = enc.encode("Hello World");
        console.log(note);
        let txn = algosdk.makePaymentTxnWithSuggestedParams(myAccount.addr, receiver, 1000000, undefined, note, params);        
        // Save transaction to file
        fs.writeFileSync('./unsigned.txn', algosdk.encodeUnsignedTransaction( txn ));   
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
        let txn = algosdk.decodeUnsignedTransaction(fs.readFileSync('./unsigned.txn'));  
        let signedTxn = algosdk.signTransaction(txn, myAccount.sk);
        let txId = signedTxn.txID;
        console.log("Signed transaction with txID: %s", txId);
        // send signed transaction to node
        await algodClient.sendRawTransaction(signedTxn.blob).do();
        // Wait for transaction to be confirmed
        let confirmedTxn = await waitForConfirmation(algodClient, tx.txId, 4);
        //Get the completed Transaction
        console.log("Transaction " + tx.txId + " confirmed in round " + confirmedTxn["confirmed-round"]);
        let mytxinfo = JSON.stringify(confirmedTxn.txn.txn, undefined, 2);
        console.log("Transaction information: %o", mytxinfo);
        var string = new TextDecoder().decode(confirmedTxn.txn.txn.note);
        console.log("Note field: ", string);

    } catch ( e ){
        console.log( e );
    }   
}; 
async function writeSignedTransactionToFile() {

    try{
        const receiver = "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A";

        // setup connection to node
        let algodClient = await setupClient();
        let myAccount = await recoverAccount();
        console.log("My address: %s", myAccount.addr)

        // get network suggested parameters
        let params = await algodClient.getTransactionParams().do();
        // setup a transaction
        const enc = new TextEncoder();
        const note = enc.encode("Hello World");
        console.log(note);
        let txn = algosdk.makePaymentTxnWithSuggestedParams(myAccount.addr, receiver, 1000000, undefined, note, params);        
        // sign transaction and write to file
        let signedTxn = txn.signTxn(myAccount.sk);
        fs.writeFileSync('./signed.stxn', signedTxn );  

    } catch( e ) {
        console.log(e);
    }
}; 
async function readSignedTransactionFromFile() {

    try{
        // setup connection to node
        let algodClient = await setupClient();

        // read signed transaction from file
        let stx = fs.readFileSync("./signed.stxn");
        // send signed transaction to node
        let tx = await algodClient.sendRawTransaction(stx).do();
        console.log("Signed transaction with txID: %s", tx.txId);
        // Wait for confirmation
        let confirmedTxn = await waitForConfirmation(algodClient, tx.txId, 4);
        //Get the completed Transaction
        console.log("Transaction " + tx.txId + " confirmed in round " + confirmedTxn["confirmed-round"]);
        let mytxinfo = JSON.stringify(confirmedTxn.txn.txn, undefined, 2);
        console.log("Transaction information: %o", mytxinfo);
        var string = new TextDecoder().decode(confirmedTxn.txn.txn.note);
        console.log("Note field: ", string);


        
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