// SearchTxAddressAsset.js
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
    let address = "SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU";
    let asset_id = 2044572;
    let min_amount = 50;

    let response = await indexerClient.searchForTransactions()
            .address(address)
            .currencyGreaterThan(min_amount)
            .assetID(asset_id).do();
    console.log("Information for Transaction search: " + JSON.stringify(response, undefined, 2));
    }   
)().catch(e => {
    console.log(e);
    console.trace();
});