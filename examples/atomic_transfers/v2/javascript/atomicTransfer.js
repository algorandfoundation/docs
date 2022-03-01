const algosdk = require('algosdk');

let  client = null;
async function setupClient() {
    if (client == null) {
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
// never use mnemonics in code, for demo purposes only
function recoverAccount1() {
    const passphrase = "Your 25-word mnemonic goes here";
    let myAccount = algosdk.mnemonicToSecretKey(passphrase);
    return myAccount;
}
// recover second account
function recoverAccount2() {
    const passphrase = "Your 25-word mnemonic goes here";
    let myAccount = algosdk.mnemonicToSecretKey(passphrase);
    return myAccount;
}

async function submitAtomicTransfer() {

    try {
        // Account C
        const  myAccountC = "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A";
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
        //Check your balances
        let accountInfo = await algodClient.accountInformation(myAccountA.addr).do();
        console.log("Account A balance: %d microAlgos", accountInfo.amount);
        accountInfo = await algodClient.accountInformation(myAccountB.addr).do();
        console.log("Account B balance: %d microAlgos", accountInfo.amount);  
        accountInfo = await algodClient.accountInformation(myAccountC).do();
        console.log("Account C balance: %d microAlgos", accountInfo.amount);   

        // get suggested params from the network
        let params = await algodClient.getTransactionParams().do();

        // Transaction A to C 
        let transaction1 = algosdk.makePaymentTxnWithSuggestedParams(myAccountA.addr, myAccountC, 100000, undefined, undefined, params);
        // Create transaction B to A
        let transaction2 = algosdk.makePaymentTxnWithSuggestedParams(myAccountB.addr, myAccountA.addr, 200000, undefined, undefined, params);

        // Store both transactions
        let txns = [transaction1, transaction2];

        // Group both transactions
        let txgroup = algosdk.assignGroupID(txns);

        // Sign each transaction in the group 
        let signedTx1 = transaction1.signTxn(myAccountA.sk)
        let signedTx2 = transaction2.signTxn(myAccountB.sk)

        // Combine the signed transactions
        let signed = []
        signed.push(signedTx1)
        signed.push(signedTx2)

        let tx = (await algodClient.sendRawTransaction(signed).do());
        console.log("Transaction : " + tx.txId);

        // Wait for transaction to be confirmed
        confirmedTxn = await algosdk.waitForConfirmation(algodClient, tx.txId, 4);
        //Get the completed Transaction
        console.log("Transaction " + tx.txId + " confirmed in round " + confirmedTxn["confirmed-round"]);
        accountInfo = await algodClient.accountInformation(myAccountA.addr).do();
        console.log("Account A balance: %d microAlgos", accountInfo.amount);
        accountInfo = await algodClient.accountInformation(myAccountB.addr).do();
        console.log("Account B balance: %d microAlgos", accountInfo.amount);  
        accountInfo = await algodClient.accountInformation(myAccountC).do();
        console.log("Account C balance: %d microAlgos", accountInfo.amount);   

    } catch (err) {
        console.log("err", err);
    }
}
submitAtomicTransfer();