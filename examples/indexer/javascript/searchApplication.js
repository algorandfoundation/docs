// AccountInfo.js
// requires algosdk@1.6.2 or higher 
// verify installed version
// npm list algosdk
const algosdk = require('algosdk');
const indexer_token = "";
const indexer_server = "http://localhost";
const indexer_port = 59998;

// Instantiate the indexer client wrapper
let indexerClient = new algosdk.Indexer(indexer_token, indexer_server, indexer_port);

(async () => {
    let response = await indexerClient.searchForApplications().do();
    console.log("Response: " + JSON.stringify(response, undefined, 2));
})().catch(e => {
    console.log(e.message);
    console.trace();
});

// Response should look similar to this...
//
// Response: {
//     "applications": [
//         {
//             "id": 20,
//             "params": {
//                 "approval-program": "ASABASI=",
//                 "clear-state-program": "ASABASI=",
//                 "creator": "DQ5PMCTEBZLM4UJEDSGZLKAV6ZGXRK2C5WYAFC63RSHI54ASQSJHDMMTUM",
//                 "global-state-schema": {
//                     "num-byte-slice": 0,
//                     "num-uint": 0
//                 },
//                 "local-state-schema": {
//                     "num-byte-slice": 0,
//                     "num-uint": 0
//                 }
//             }
//         },
//         {
//             "id": 22,
//             "params": {
//                 "approval-program": null,
//                 "clear-state-program": null,
//                 "creator": "GHFRLVOMKJNTJ4HY3P74ZR4CNE2PB7CYAUAJ6HVAVVDX7ZKEMLJX6AAF4M",
//                 "global-state-schema": {
//                     "num-byte-slice": 0,
//                     "num-uint": 0
//                 },
//                 "local-state-schema": {
//                     "num-byte-slice": 0,
//                     "num-uint": 0
//                 }
//             }
//         },
//         ...

