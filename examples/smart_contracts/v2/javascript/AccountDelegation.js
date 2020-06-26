const algosdk = require('algosdk');

// sandbox
const token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
const server = "http://localhost";
const port = 4001;
    
// const token = "algod-token"<PLACEHOLDER>;
// const server = "algod-address"<PLACEHOLDER>;
// const port = algod - port<PLACEHOLDER>;

// import your private key mnemonic and address
// let PASSPHRASE = "25-word-mnemonic<PLACEHOLDER>";
let PASSPHRASE = "awake used crawl list cruel harvest useful flag essay speed glad salmon camp sudden ride symptom test kind version together project inquiry diet abandon budget";

var myAccount = algosdk.mnemonicToSecretKey(PASSPHRASE);
console.log("My Address: " + myAccount.addr);
// Function used to wait for a tx confirmation
const waitForConfirmation = async function (algodclient, txId) {
    let response = await algodclient.status().do();
    let lastround = response["last-round"];
    while (true) {
        const pendingInfo = await algodclient.pendingTransactionInformation(txId).do();
        if (pendingInfo["confirmed-round"] !== null && pendingInfo["confirmed-round"] > 0) {
            //Got the completed Transaction
            console.log("Transaction " + txId + " confirmed in round " + pendingInfo["confirmed-round"]);
            break;
        }
        lastround++;
        await algodclient.statusAfterBlock(lastround).do();
    }
};
// create an algod v2 client
let algodclient = new algosdk.Algodv2(token, server, port);

(async () => {

    // // get suggested parameters
    // let params = await algodclient.getTransactionParams().do();
    // let endRound = params.lastRound + parseInt(1000);
    // let fee = await algodclient.suggestedFee();
    // get suggested parameters
    let params = await algodclient.getTransactionParams().do();
    // comment out the next two lines to use suggested fee 
    params.fee = 1000;
    params.flatFee = true;
    console.log(params);
    // create logic sig

    // b64 example "ASABACI=" is `int 0`
    // see more info here: https://developer.algorand.org/docs/features/asc1/sdks/#accessing-teal-program-from-sdks
    // let program = new Uint8Array(Buffer.from("base64-encoded-program" < PLACEHOLDER >, "base64"));

    let program = new Uint8Array(Buffer.from("ASABACI=" , "base64"));
    let lsig = algosdk.makeLogicSig(program);

    // sign the logic signature with an account sk
    lsig.sign(myAccount.sk);

    // create a transaction
    let txn = {
        "from": myAccount.addr,
        // "to": "receiver-address" < PLACEHOLDER >,
        "to": "SOEI4UA72A7ZL5P25GNISSVWW724YABSGZ7GHW5ERV4QKK2XSXLXGXPG5Y",
        "fee": params.fee,
        "flatFee": params.flatFee,
        "amount": 10000,
        "firstRound": params.firstRound,
        "lastRound": params.lastRound,
        "genesisID": params.genesisID,
        "genesisHash": params.genesisHash
    };

    // Create the LogicSigTransaction with contract account LogicSig 
    let rawSignedTxn = algosdk.signLogicSigTransaction(txn, lsig);

    // send raw LogicSigTransaction to network    
    let tx = (await algodclient.sendRawTransaction(rawSignedTxn.blob).do());
    console.log("Transaction : " + tx.txId);    
    await waitForConfirmation(algodclient, tx.txId);

})().catch(e => {
    console.log(e);
});
