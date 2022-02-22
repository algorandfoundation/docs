const algosdk = require('algosdk');
const fs = require('fs');
let  client = null;
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
// Shown for demonstration purposes. NEVER reveal secret mnemonics in practice.
function recoverAccount(){
    const passphrase = "<your-25-word-mnemonic>";
    let myAccount = algosdk.mnemonicToSecretKey(passphrase);
    return myAccount;
}


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
        let txn = algosdk.makePaymentTxnWithSuggestedParams(myAccount.addr, receiver, 100000, undefined, note, params);        
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
        const confirmedTxn = await algosdk.waitForConfirmation(algodClient, txId, 4);
        //Get the completed Transaction
        console.log("Transaction " + txId + " confirmed in round " + confirmedTxn["confirmed-round"]);
        let mytxinfo = JSON.stringify(confirmedTxn.txn.txn, undefined, 2);
        console.log("Transaction information: %o", mytxinfo);
        let  string = new TextDecoder().decode(confirmedTxn.txn.txn.note);
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
        let txn = algosdk.makePaymentTxnWithSuggestedParams(myAccount.addr, receiver, 100000, undefined, note, params);        
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
        const confirmedTxn = await algosdk.waitForConfirmation(algodClient, tx.txId, 4);
        //Get the completed Transaction
        console.log("Transaction " + tx.txId + " confirmed in round " + confirmedTxn["confirmed-round"]);
        let mytxinfo = JSON.stringify(confirmedTxn.txn.txn, undefined, 2);
        console.log("Transaction information: %o", mytxinfo);
        let  string = new TextDecoder().decode(confirmedTxn.txn.txn.note);
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
// testUnsigned();
testSigned();