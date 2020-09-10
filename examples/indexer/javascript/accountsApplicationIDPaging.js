// requires algosdk@1.7 or higher 
// verify installed version
// npm list algosdk

const algosdk = require('algosdk');

// const indexer_server = "https://testnet-algorand.api.purestake.io/idx2/";
// const indexer_port = "";
// const indexer_token = {
//     'X-API-key': 'B3SU4KcVKi94Jap2VXkK83xx38bsv95K5UZm2lab',
// }
const indexer_server = "http://localhost";
const indexer_port = 59998;
const indexer_token = "";
 
// Instantiate the indexer client wrapper
let indexerClient = new algosdk.Indexer(indexer_token, indexer_server, indexer_port);

let nexttoken = "";
let numAccounts = 1;
(async () => {

    let applicationID = 70;
    let limit = 2;
    while (numAccounts > 0) {
        // execute code as long as condition is true
        let next_page = nexttoken;
        let accountInfo = await indexerClient.searchAccounts()
            .applicationID(applicationID).limit(limit).nextToken(next_page).do();
        let accounts = accountInfo['accounts'];
        numAccounts = accounts.length;
        if (numAccounts > 0) {
            nexttoken = accountInfo['next-token'];
            console.log("Information for account info for Application ID: " + JSON.stringify(accountInfo, undefined, 2));
        }
    }

}) ().catch(e => {
    console.log(e);
    console.trace();
});