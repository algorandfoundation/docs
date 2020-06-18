// SearchAssetTransactionsRole.js
// requires algosdk@1.6.1 or higher 
// verify installed version
// npm list algosdk

const algosdk = require('algosdk');
const indexer_token = "";
const indexer_server = "http://localhost";
const indexer_port = 8980;

// Instantiate the indexer client wrapper
let indexerClient = new algosdk.Indexer(indexer_token, indexer_server, indexer_port);

(async () => {
    let asset_id = 2044572;
    let address_role = "receiver";
    let address = "UF7ATOM6PBLWMQMPUQ5QLA5DZ5E35PXQ2IENWGZQLEJJAAPAPGEGC3ZYNI";
    let tracsactionInfo = await indexerClient.lookupAssetTransactions(asset_id)
        .addressRole(address_role)
        .address(address).do();
    console.log("Information for Transaction for Asset: " + JSON.stringify(tracsactionInfo, undefined, 2));
})().catch(e => {
    console.log(e);
    console.trace();
});