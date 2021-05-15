const algosdk = require('algosdk');
const { waitForConfirmation } = require('./utils/waitForConfirmation');

const main = async () => {
    // Generate a public/private key pair
    const passphrase = "price clap dilemma swim genius fame lucky crack torch hunt maid palace ladder unlock symptom rubber scale load acoustic drop oval cabbage review abstract embark";
    const myAccount = algosdk.mnemonicToSecretKey(passphrase);
    console.log("My address: %s", myAccount.addr);
    console.log("My passphrase: " + passphrase);

    // Faucets to add funds
    // TestNet Faucet: https://bank.testnet.algorand.network/
    // BetaNet Faucet: https://bank.betanet.algodev.network/

    // Connect your client
    const algodToken = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
    const algodServer = "http://localhost";
    const algodPort = 4001;

    const algodClient = new algosdk.Algodv2(algodToken, algodServer, algodPort);

    //Check your balance
    const accountInfo = await algodClient.accountInformation(myAccount.addr).do();
    console.log("Account balance: %d microAlgos", accountInfo.amount);

    // Construct the transaction
    let params = await algodClient.getTransactionParams().do();
    // comment out the next two lines to use suggested fee
    params.fee = 1000;
    params.flatFee = true;

    // receiver defined as TestNet faucet address 
    const receiver = "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A";
    const enc = new TextEncoder();
    const note = enc.encode("Hello World");

    const txn = algosdk.makePaymentTxnWithSuggestedParams(myAccount.addr, receiver, 1000000, undefined, note, params);

    // Sign the transaction
    const signedTxn = txn.signTxn(myAccount.sk);
    const txId = txn.txID().toString();
    console.log("Signed transaction with txID: %s", txId);

    // Submit the transaction
    await algodClient.sendRawTransaction(signedTxn).do();

    const confirmedTxn = await waitForConfirmation(algodClient, txId, 4);
    console.log("Transaction " + txId + " confirmed in round " + confirmedTxn["confirmed-round"]);

    const mytxinfo = JSON.stringify(confirmedTxn.txn.txn, undefined, 2);
    console.log("Transaction information: %o", mytxinfo);

    const noteString = new TextDecoder().decode(confirmedTxn.txn.txn.note);
    console.log("Note field: ", noteString);
};

main().then().catch(e => {
    const error = e.body && e.body.message ? e.body.message : e;
    console.log(error);
});