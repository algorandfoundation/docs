const algosdk = require('algosdk');

// const token = "algod-token"<PLACEHOLDER>;
// const server = "algod-address"<PLACEHOLDER>;
// const port = algod - port<PLACEHOLDER>;

const token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
const server = "http://localhost";
const port = 4001;


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

// Import the filesystem module 
const fs = require('fs');
let algodclient = new algosdk.Algodv2(token, server, port);

(async () => {
    // Read file for Teal code - int 0
    var fs = require('fs'),
        path = require('path'),
        filePath = path.join(__dirname, 'sample.teal');
    let data = fs.readFileSync(filePath);
    //  algodclient.compile(data.toString())
    let results = await algodclient.compile(data).do();
    console.log("Hash = " + results.hash);
    console.log("Result = " + results.result);
    
})().catch(e => {
    console.log(e.body.message);
    console.log(e);
});
// output would be similar to this... 
// Hash = KI4DJG2OOFJGUERJGSWCYGFZWDNEU2KWTU56VRJHITP62PLJ5VYMBFDBFE
// Result = ASABACI=