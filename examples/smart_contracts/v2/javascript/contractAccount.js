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
    // comment out the next two lines to use suggested fee
    params.fee = 1000;
    params.flatFee = true;

    const filePath = path.join(__dirname, '/teal/samplearg.teal');
      
    const data = fs.readFileSync(filePath);
    const  results = await algodClient.compile(data).do();
    console.log("Hash = " + results.hash);
    console.log("Result = " + results.result);

    const program = new Uint8Array(Buffer.from(results.result, "base64"));
    // Use this if no args
    // const lsig = algosdk.makeLogicSig(program);

    // Initialize arguments array
    const args = [];

    // String parameter
    // args.push([...Buffer.from("my string")]);

    // Integer parameter
    args.push(algosdk.encodeUint64(123));

    const lsig = algosdk.makeLogicSig(program, args);
    console.log("lsig : " + lsig.address());   

    // Create a transaction
    const sender = lsig.address();
    // This is a sandbox address and will be different in your environment
    const receiver = "HZ57J3K46JIJXILONBBZOHX6BKPXEM2VVXNRFSUED6DKFD5ZD24PMJ3MVA";
    const amount = 10000;
    const closeToRemaninder = undefined;
    const note = undefined;
    const txn = algosdk.makePaymentTxnWithSuggestedParams(sender, receiver, amount, closeToRemaninder, note, params)

    const rawSignedTxn = algosdk.signLogicSigTransactionObject(txn, lsig);

    // send raw LogicSigTransaction to network
    const tx = await algodClient.sendRawTransaction(rawSignedTxn.blob).do();
    console.log("Transaction : " + tx.txId);   
    await waitForConfirmation(algodClient, tx.txId);
};

main().then().catch(e => {
    const error = e.body && e.body.message ? e.body.message : e;
    console.log(error);
});