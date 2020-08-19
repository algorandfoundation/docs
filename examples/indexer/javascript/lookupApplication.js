// AccountInfo.js
// requires algosdk@1.6.1 or higher 
// verify installed version
// npm list algosdk
const algosdk = require('algosdk');

const indexer_token = "";
const indexer_server = "http://localhost";
const indexer_port = 59998;

// Instantiate the indexer client wrapper
let indexerClient = new algosdk.Indexer(indexer_token, indexer_server, indexer_port);

(async () => {
    let response = await indexerClient.lookupApplications(59).do();
    console.log("Response: " + JSON.stringify(response, undefined, 2));
})().catch(e => {
    console.log(e.message);
    console.trace();
});

// Repsonse should look similar to this...

// Response: {
// "application": {
//     "id": 59,
//         "params": {
//         "approval-program": "AiACAQAmAgNmb28DYmFyIihlQQAIKRJBAANCAAIjQyJD",
//             "clear-state-program": "ASABASI=",
//                 "creator": "7IROB3J2FTR7LYQA3QOUYSTKCQTSVJK4FTTC77KWSE5NMRATEZXP6TARPA",
//                     "global-state-schema": {
//             "num-byte-slice": 0,
//                 "num-uint": 0
//         },
//         "local-state-schema": {
//             "num-byte-slice": 0,
//                 "num-uint": 0
//         }
//     }
// },
// "current-round": 377
// }
