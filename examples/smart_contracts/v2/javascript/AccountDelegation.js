const algosdk = require('algosdk');
const fs = require('fs');
const { getAccounts } = require('./sandbox.js');


// Create an algod client, using default sandbox parameters here
const token = "a".repeat(64);
const server = "http://localhost";
const port = 4001;
let algodclient = new algosdk.Algodv2(token, server, port);

(async () => {
    // Get receiver account info from sandbox
    const accounts = await getAccounts();
    const receiverAccount = accounts[0];
    console.log("Receiver Address: " + receiverAccount.addr);


    // Get sender account info from sandbox
    const senderAccount = accounts[0];
    console.log("Sender Address: " + senderAccount.addr);

    // Get suggested parameters
    let params = await algodclient.getTransactionParams().do();
    console.log(params);

    // Read TEAL file. See more info here: https://developer.algorand.org/docs/features/asc1/sdks/#accessing-teal-program-from-sdks
    let fs = require('fs'),
        path = require('path'),
        filePath = path.join(__dirname, 'samplearg.teal');
    let data = fs.readFileSync(filePath);

    //Compile the program against the algod
    let results = await algodclient.compile(data).do();
    console.log("Hash = " + results.hash);
    console.log("Result = " + results.result);
    let program = new Uint8Array(Buffer.from(results.result, "base64"));

    // Create the lsig account, passing arg
    let args = getUint8Int(123);
    let lsig = new algosdk.LogicSigAccount(program, args);
    console.log("lsig : " + lsig.address());

    // sign the logic signature with an account sk
    lsig.sign(senderAccount.sk);

    // Setup a transaction
    let sender = senderAccount.addr;
    let receiver = receiverAccount.addr;
    let amount = 10000;
    let closeToRemaninder = undefined;
    let note = undefined;
    let txn = algosdk.makePaymentTxnWithSuggestedParams(sender,
        receiver, amount, closeToRemaninder, note, params)

    // Create the LogicSigTransaction with contract account LogicSig
    let rawSignedTxn = algosdk.signLogicSigTransactionObject(txn, lsig.lsig);
    console.log("signedTxn: ", rawSignedTxn)

    // send raw LogicSigTransaction to network    
    let tx = (await algodclient.sendRawTransaction(rawSignedTxn.blob).do());
    console.log("Transaction : " + tx.txId);

    //Wait for confirmation
    const confirmedTxn = await algosdk.waitForConfirmation(algodclient, tx.txId, 4);

    //Get the completed Transaction
    console.log("Transaction " + tx.txId + " confirmed in round " + confirmedTxn["confirmed-round"]);


})().catch(e => {
    console.log(e.message);
    console.log(e);
});

function getUint8Int(number) {
    const buffer = Buffer.alloc(8);
    const bigIntValue = BigInt(number);
    buffer.writeBigUInt64BE(bigIntValue);
    return [Uint8Array.from(buffer)];
}
