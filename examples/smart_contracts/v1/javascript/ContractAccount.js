const algosdk = require('algosdk');

// const token = "algod-token"<PLACEHOLDER>;
// const server = "algod-address"<PLACEHOLDER>;
// const port = algod - port<PLACEHOLDER>;

const token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
const server = "http://localhost";
const port = 4001;

// Function used to wait for a tx confirmation
const waitForConfirmation = async function (algodclient, txId) {
    let lastround = (await algodclient.status()).lastRound;
    while (true) {
        const pendingInfo = await algodclient.pendingTransactionInformation(txId);
        if (pendingInfo.round !== null && pendingInfo.round > 0) {
            //Got the completed Transaction
            console.log("Transaction " + pendingInfo.tx + " confirmed in round " + pendingInfo.round);
            break;
        }
        lastround++;
        await algodclient.statusAfterBlock(lastround);
    }
};
// create an algod client
let algodclient = new algosdk.Algod(token, server, port);

(async () => {

    // before every transaction in this example
    let params = await algodclient.getTransactionParams();
    console.log(params);
    let endRound = params.lastRound + parseInt(1000);
    let fee = await algodclient.suggestedFee();

    // create logic sig

    // b64 example "ASABACI=" is `int 0`
    // see more info here: https://developer.algorand.org/docs/features/asc1/sdks/#accessing-teal-program-from-sdks
    // let program = new Uint8Array(Buffer.from("base64-encoded-program" < PLACEHOLDER >, "base64"));
    let program = new Uint8Array(Buffer.from("ASABACI=" , "base64"));
    let lsig = algosdk.makeLogicSig(program);


    // create a transaction
    let txn = {
        "from": lsig.address(),
        // "to": "receiver-address" < PLACEHOLDER >,
        "to": "SOEI4UA72A7ZL5P25GNISSVWW724YABSGZ7GHW5ERV4QKK2XSXLXGXPG5Y",      
        "fee": params.fee,
        // "amount": amount < PLACEHOLDER >,
        "amount": 10000,
        "firstRound": params.lastRound,
        "lastRound": endRound,
        "genesisID": params.genesisID,
        "genesisHash": params.genesishashb64
    };
    // Create the LogicSigTransaction with contract account LogicSig 
    let rawSignedTxn = algosdk.signLogicSigTransaction(txn, lsig);

    // send raw LogicSigTransaction to network
    console.log("This is expected to fail as the program is int 0 and will return false : ");     
    let tx = (await algodclient.sendRawTransaction(rawSignedTxn.blob));
    console.log("Transaction : " + tx.txId);    
    await waitForConfirmation(algodclient, tx.txId);

})().catch(e => {
    console.log(e);
});