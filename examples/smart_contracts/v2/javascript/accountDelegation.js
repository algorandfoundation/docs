const fs = require('fs');
const path = require('path');
const algosdk = require('algosdk');
const { waitForConfirmation } = require('./utils/waitForConfirmation');

// We assume that testing is done off of sandbox, hence the settings below
const token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
const server = "http://localhost";
const port = 4001;

// create v2 client
const algodClient = new algosdk.Algodv2(token, server, port);

const main = async () => {
    // get suggested parameters
    const params = await algodClient.getTransactionParams().do();
    params.fee = 1000;
    params.flatFee = true;

    const filePath = path.join(__dirname, '/teal/samplearg.teal');
    const data = fs.readFileSync(filePath);
    const  results = await algodClient.compile(data).do();
    const program = new Uint8Array(Buffer.from(results.result, "base64"));

    const args = [];
    args.push(algosdk.encodeUint64(123));
   
    const lsig = algosdk.makeLogicSig(program, args); 
    
    // *** Begin account delegation changes ***
    // import your private key mnemonic
    const PASSPHRASE = "< mnemonic >";
    const myAccount = algosdk.mnemonicToSecretKey(PASSPHRASE);

    // sign the logic signature with an account sk
    lsig.sign(myAccount.sk);
    // *** End account delegation changes ***

    // Setup a transaction
    const sender = myAccount.addr;
    const receiver = "GVVIH6IE3ZAFLIQF6UFBOH4TI5ZOPHFRMDQUTXQRKZ6XWSRTVRDBQ4GZLY";
    const amount = 10000;
    const closeToRemaninder = undefined;
    const note = undefined;
    const txn = algosdk.makePaymentTxnWithSuggestedParams(sender, receiver, amount, closeToRemaninder, note, params)
    const rawSignedTxn = algosdk.signLogicSigTransactionObject(txn, lsig);
   
    const tx = await algodClient.sendRawTransaction(rawSignedTxn.blob).do();
    console.log("Transaction : " + tx.txId);    
    await waitForConfirmation(algodClient, tx.txId);
};

main().then().catch(e => {
    const error = e.body && e.body.message ? e.body.message : e;
    console.log(error);
});