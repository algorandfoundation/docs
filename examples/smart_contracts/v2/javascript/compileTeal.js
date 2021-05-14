const fs = require('fs');
const path = require('path');
const algosdk = require('algosdk');
const { waitForConfirmation } = require('./utils/waitForConfirmation');

// We assume that testing is done off of sandbox, hence the settings below
const token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
const server = "http://localhost";
const port = 4001;

const algodClient = new algosdk.Algodv2(token, server, port);

const main = async () => {
    const filePath = path.join(__dirname, '/teal/sample.teal');
    const data = fs.readFileSync(filePath);

    const results = await algodClient.compile(data).do();
    return results;
}

main().then((results) => {
    console.log("Hash = " + results.hash);
    console.log("Result = " + results.result);
    // output would be similar to this... 
    // Hash = KI4DJG2OOFJGUERJGSWCYGFZWDNEU2KWTU56VRJHITP62PLJ5VYMBFDBFE
    // Result = ASABACI=
}).catch(e => {
    const error = e.body && e.body.message ? e.body.message : e;
    console.log(error);
});
