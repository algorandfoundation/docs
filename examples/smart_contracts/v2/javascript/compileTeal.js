const algosdk = require('algosdk');

// const token = "algod-token"<PLACEHOLDER>;
// const server = "algod-address"<PLACEHOLDER>;
// const port = algod - port<PLACEHOLDER>;

const token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
const server = "http://localhost";
const port = 4001;


// Import the filesystem module 
const fs = require('fs');
// create v2 client
let algodclient = new algosdk.Algodv2(token, server, port);

(async () => {
    // Read file for Teal code - int 0
    let data = fs.readFileSync('sample.teal');
    // Compile teal
    let results = await algodclient.compile(data).do();
    // Print results
    console.log("Hash = " + results['hash']);
    console.log("Result = " + results['result']);   
})().catch(e => {
    console.log(e.body.message);
    console.log(e);
});