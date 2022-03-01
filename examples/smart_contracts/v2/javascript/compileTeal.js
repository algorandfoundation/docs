const algosdk = require('algosdk');

// const token = "<algod-token>";
// const server = "<algod-address>";
// const port = <algod-port>;

const token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
const server = "http://localhost";
const port = 4001;



// Import the filesystem module 
const fs = require('fs');
let algodclient = new algosdk.Algodv2(token, server, port);

(async () => {
    // Read file for Teal code - int 0
    var fs = require('fs'),
        path = require('path'),
        filePath = path.join(__dirname, 'sample.teal');
    //  filePath = path.join(__dirname, '<filename>');    
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