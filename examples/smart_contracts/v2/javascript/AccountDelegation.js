const algosdk = require('algosdk');
const fs = require('fs');

// Create an algod client, using default sandbox parameters here
const token = "a" * 64;
const server = "http://localhost";
const port = 4001;
let algodclient = new algosdk.Algodv2(token, server, port);

// Get account info
// mnemonics should not be used in production code, for demo purposes only
let PASSPHRASE = "< Your 25 word mnemonic >"
let myAccount = algosdk.mnemonicToSecretKey(PASSPHRASE);
console.log("My Address: " + myAccount.addr);

(async () => {
    // Get suggested parameters
    let params = await algodclient.getTransactionParams().do();
    console.log(params);

    // Read a file
    // see more info here: https://developer.algorand.org/docs/features/asc1/sdks/#accessing-teal-program-from-sdks
    let fs = require('fs'),
        path = require('path'),
        filePath = path.join(__dirname, 'samplearg.teal');
    let data = fs.readFileSync(filePath);

    //Compile the program against the algod
    let results = await algodclient.compile(data).do();
    console.log("Hash = " + results.hash);
    console.log("Result = " + results.result);
    // let program = new Uint8Array(Buffer.from(<"base64-encoded-program">, "base64"));
    let program = new Uint8Array(Buffer.from(results.result, "base64"));
    // Use this if no args
    // let lsig = algosdk.makeLogicSig(program);

    // String parameter
    // let args = ["<my string>"];
    // let lsig = algosdk.makeLogicSig(program, args);
    // Integer parameter
    // Create the lsig account, passing arg
    let args = getUint8Int(12345);
    let lsig = new algosdk.LogicSigAccount(program, args);

    // sign the logic signature with an account sk
    lsig.sign(myAccount.sk);

    // Setup a transaction
    let sender = myAccount.addr;
    let receiver = "SOEI4UA72A7ZL5P25GNISSVWW724YABSGZ7GHW5ERV4QKK2XSXLXGXPG5Y";
    // let receiver = "<receiver-address>"";
    let amount = 10000;
    let closeToRemaninder = undefined;
    let note = undefined;
    let txn = algosdk.makePaymentTxnWithSuggestedParams(sender,
        receiver, amount, closeToRemaninder, note, params)

    // Create the LogicSigTransaction with contract account LogicSig
    let rawSignedTxn = algosdk.signLogicSigTransactionObject(txn, lsig);

    // fs.writeFileSync("simple.stxn", rawSignedTxn.blob);
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
