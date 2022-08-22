const algosdk = require('algosdk');

const token = "a".repeat(64);
const server = "http://localhost";
const port = 4001;

// Import the filesystem module 
const fs = require('fs');

let algodclient = new algosdk.Algodv2(token, server, port);

(async () => {
    // Read file for Teal code - int 1
    var fs = require('fs'),
        path = require('path'),
        filePath = path.join(__dirname, 'samplearg.teal');
    let data = fs.readFileSync(filePath);
    let results = await algodclient.compile(data).do();
    console.log("Hash = " + results.hash);
    console.log("Result = " + results.result);

})().catch(e => {
    console.log(e.body.message);
    console.log(e);
});

// results should be the following:
// Response Result =  Bi0XgXsS
// Response Hash =  PBS3JWRY5HGL46SDYLDVMEJVJXKMSCFOAY7SKIW2RQ7OTJVWNNGW4QXQ5A