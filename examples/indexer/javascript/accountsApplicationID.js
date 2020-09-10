// AccountsAssetID.js
// requires algosdk@1.6.1 or higher 
// verify installed version
// npm list algosdk

const algosdk = require('algosdk');
const chalk = require('chalk');
// const indexer_token = "";
// const indexer_server = "http://localhost";
// const indexer_port = 8980;

// const indexer_server = "https://testnet-algorand.api.purestake.io/idx2/";
// const indexer_port = "";
// const indexer_token = {
//     'X-API-key': 'B3SU4KcVKi94Jap2VXkK83xx38bsv95K5UZm2lab',
// }
const indexer_server = "http://localhost";
const indexer_port = 59998;
const indexer_token = "";

// const indexer_server = "https://indexer-internal-testnet.aws.algodev.network";
// const indexer_port = 443;
// const indexer_token = "YddOUGbAjHLr1uPZtZwHOvMDmXvR1Zvw1f3Roj2PT1ufenXbNyIxIz0IeznrLbDsF"

 
// Instantiate the indexer client wrapper
let indexerClient = new algosdk.Indexer(indexer_token, indexer_server, indexer_port);

(async () => {
    let applicationID = 70;  
    let accountInfo = await indexerClient.searchAccounts()
        .applicationID(applicationID).do();
    console.log(chalk.black("Information for account info for Application ID: " + JSON.stringify(accountInfo, undefined, 2)));
})().catch(e => {
    console.log(e);
    console.trace();
});