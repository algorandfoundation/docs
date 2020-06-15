// SearchTxAddressBlockRange.js
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
    let address = "XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4";
    let min_round = 7048876;
    let max_round = 7048878;    
    let response = await indexerClient.searchForTransactions()
        .address(address).maxRound(max_round)
        .minRound(min_round).do();
    console.log("Information for Transaction search: " + JSON.stringify(response, undefined, 2));
    }  
)().catch(e => {
    console.log(e);
    console.trace();
});