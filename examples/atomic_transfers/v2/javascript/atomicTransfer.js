const algosdk = require('algosdk');


var client = null;
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
async function submitGroupTransactions(){

    try{
        // receiver Account C
        const receiver = "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A";
        // sample show account A to C
        // B to A 
        // grouped
        let algodClient = await setupClient();

        // recover account
        // Account A
        let myAccountA = await recoverAccount1();
        console.log("My account A address: %s", myAccountA.addr)

        // recover an additional account
        // Account B
        let myAccountB = await recoverAccount2();
        console.log("My account B address: %s", myAccountB.addr)

        // get suggested params from the network
        let params = await algodClient.getTransactionParams().do();

        // Transaction A to C 
        let transaction1 = algosdk.makePaymentTxnWithSuggestedParams(myAccountA.addr, receiver, 1000000, undefined, undefined, params);  
        // Create transaction B to A
        let transaction2 = algosdk.makePaymentTxnWithSuggestedParams(myAccountB.addr, myAccountA.addr, 1000000, undefined, undefined, params);  

        // Store both transactions
        let txns = [transaction1, transaction2];

        // Group both transactions
        let txgroup = algosdk.assignGroupID(txns);

        // Sign each transaction in the group 
        signedTx1 = transaction1.signTxn( myAccountA.sk )
        signedTx2 = transaction1.signTxn( myAccountB.sk )
    
        // Combine the signed transactions
        let signed = []
        signed.push( signedTx1 )
        signed.push( signedTx2 )

        let tx = (await algodClient.sendRawTransaction(signed).do());
        console.log("Transaction : " + tx.txId);

        // Wait for transaction to be confirmed
        await waitForConfirmation(algodClient, tx.txId)
    } catch (err) {
        console.log("err", err);  
    }
}
submitGroupTransactions();