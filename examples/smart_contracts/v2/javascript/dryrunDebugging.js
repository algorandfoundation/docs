const algosdk = require('algosdk');

// const token = "<algod-token>";
// const server = "<algod-address>";
// const port = <algod-port>;
// sandbox
const token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
const server = "http://localhost";
const port = 4001;

// Import the filesystem module 
const fs = require('fs'); 
// import your private key mnemonic
// mnemonics should not be used in prodcution code, for demo purposes only
let PASSPHRASE = "<25-word-mnemonic>";

let  myAccount = algosdk.mnemonicToSecretKey(PASSPHRASE);
console.log("My Address: " + myAccount.addr);
// create an algod v2 client
let algodclient = new algosdk.Algodv2(token, server, port);

(async () => {

    // get suggested parameter
    let params = await algodclient.getTransactionParams().do();
    // comment out the next two lines to use suggested fee 
    // params.fee = 1000;
    // params.flatFee = true;
    console.log(params);
    // create logic sig
    // samplearg.teal
    // This code is meant for learning purposes only
    // It should not be used in production
    // arg_0
    // btoi
    // int 12345
    // ==
    let  fs = require('fs'),
        path = require('path'),
        filePath = path.join(__dirname, 'samplearg.teal');
    // filePath = path.join(__dirname, <'fileName'>);
    let data = fs.readFileSync(filePath);
    let results = await algodclient.compile(data).do();
    console.log("Hash = " + results.hash);
    console.log("Result = " + results.result);
    // let program = new Uint8Array(Buffer.from(<"base64-encoded-program">, "base64"));
    let program = new Uint8Array(Buffer.from(results.result, "base64"));
    // Use this if no args
    // let lsig = new algosdk.LogicSigAccount(program);

    // String parameter
    // let args = ["<my string>"];
    // let lsig = new algosdk.LogicSigAccount(program, args);
    // Integer parameter
    // let args = getUint8Int(12345);

    let args = getUint8Int(12345);
    let lsig = new algosdk.LogicSigAccount(program, args);

    // sign the logic signature with an account sk
    lsig.sign(myAccount.sk);
    
    // Setup a transaction
    let sender = myAccount.addr;
    let receiver = "SOEI4UA72A7ZL5P25GNISSVWW724YABSGZ7GHW5ERV4QKK2XSXLXGXPG5Y";
    // let receiver = "<receiver-address>"";
    let amount = 10000;
    let closeToRemaninder = undefined;
    let note = undefined;
    let txn = algosdk.makePaymentTxnWithSuggestedParams(sender, receiver, amount, closeToRemaninder, note, params)

    let rawSignedTxn = algosdk.signLogicSigTransactionObject(txn, lsig);
    // fs.writeFileSync("simple.stxn", rawSignedTxn.blob);


    //compile debugging
    const dryrunResponse = await algosdk.createDryrun({
        client: algodclient, 
        txns: [
            algosdk.decodeSignedTransaction(rawSignedTxn['blob'])
        ]
    })
    let  textedJson = JSON.stringify(dryrunResponse, undefined, 4);
    console.log("compile Response ");  
    console.log(textedJson);

    //source debugging
    const dryrunResponseSource = await algosdk.createDryrun({
        client: algodclient, 
        txns: [
            algosdk.decodeSignedTransaction(rawSignedTxn['blob'])
        ],
        sources: data
        
    })
    textedJson = JSON.stringify(dryrunResponseSource, undefined, 4);
    console.log("source Response ");  
    console.log(textedJson);
    
    // send raw LogicSigTransaction to network 
    let tx = (await algodclient.sendRawTransaction(rawSignedTxn.blob).do());
    console.log("Transaction : " + tx.txId);    
    const confirmedTxn = await algosdk.waitForConfirmation(algodclient, tx.txId, 4);
    //Get the completed Transaction
    console.log("Transaction " + tx.txId + " confirmed in round " + confirmedTxn["confirmed-round"]);

})().catch(e => {
    console.log(e.message);
    console.log(e);
});
async function dryrunDebugging(lsig, txn, data) {
    if (data == null)
    {
        //compile
        txns = [{
            lsig: lsig,
            txn: txn,
        }];        
    }
    else
    {
        // source
        txns = [{
            txn: txn,
        }];
        sources = [new algosdk.modelsv2.DryrunSource("lsig", data.toString("utf8"), 0)];
    }
    const dr = new algosdk.modelsv2.DryrunRequest({
        txns: txns,
        sources: sources,
    });
    dryrunResponse = await algodclient.dryrun(dr).do();
    return dryrunResponse;
}
async function dryrunDebuggingdr(lsig, txn, data) {
    if (data == null) {
        //compile
        txns = [{
            lsig: lsig,
            txn: txn,
        }];
    }
    else {
        // source
        txns = [{
            txn: txn,
        }];
        sources = [new algosdk.modelsv2.DryrunSource("lsig", data.toString("utf8"), 0)];
    }
    const dr = new algosdk.modelsv2.DryrunRequest({
        txns: txns,
        sources: sources,
    });

    return dr;
}

function getUint8Int(number) {
    const buffer = Buffer.alloc(8);
    const bigIntValue = BigInt(number);
    buffer.writeBigUInt64BE(bigIntValue);
    return  [Uint8Array.from(buffer)];
}
// output should look like this
// compile Response 
// {
//     "accounts": [],
//     "apps": [],
//     "txns": [
//         {
//             "lsig": {
//                 "arg": [
//                     {
//                         "0": 0,
//                         "1": 0,
//                         "2": 0,
//                         "3": 0,
//                         "4": 0,
//                         "5": 0,
//                         "6": 48,
//                         "7": 57
//                     }
//                 ],
//                 "l": {
//                     "0": 5,
//                     "1": 45,
//                     "2": 23,
//                     "3": 129,
//                     "4": 185,
//                     "5": 96,
//                     "6": 18
//                 },
//                 "sig": {
//                     "0": 121,
//                     "1": 119,
//                     "2": 184,
//                     "3": 224,
//                     "4": 25,
//                     "5": 132,
//                     "6": 2,
//                     "7": 28,
//                     "8": 251,
//                     "9": 65,
//                     "10": 237,
//                     "11": 243,
//                     "12": 107,
//                     "13": 144,
//                     "14": 204,
//                     "15": 191,
//                     "16": 224,
//                     "17": 143,
//                     "18": 127,
//                     "19": 67,
//                     "20": 211,
//                     "21": 70,
//                     "22": 225,
//                     "23": 113,
//                     "24": 59,
//                     "25": 182,
//                     "26": 44,
//                     "27": 38,
//                     "28": 49,
//                     "29": 223,
//                     "30": 88,
//                     "31": 197,
//                     "32": 23,
//                     "33": 44,
//                     "34": 84,
//                     "35": 94,
//                     "36": 218,
//                     "37": 170,
//                     "38": 87,
//                     "39": 228,
//                     "40": 153,
//                     "41": 189,
//                     "42": 36,
//                     "43": 84,
//                     "44": 239,
//                     "45": 151,
//                     "46": 112,
//                     "47": 207,
//                     "48": 176,
//                     "49": 90,
//                     "50": 113,
//                     "51": 176,
//                     "52": 70,
//                     "53": 248,
//                     "54": 108,
//                     "55": 134,
//                     "56": 79,
//                     "57": 188,
//                     "58": 174,
//                     "59": 37,
//                     "60": 89,
//                     "61": 56,
//                     "62": 147,
//                     "63": 13
//                 }
//             },
//             "txn": {
//                 "amt": 10000,
//                 "fee": 1000,
//                 "fv": 19648790,
//                 "lv": 19649790,
//                 "snd": {
//                     "type": "Buffer",
//                     "data": [
//                         28,
//                         237,
//                         98,
//                         5,
//                         224,
//                         62,
//                         69,
//                         120,
//                         208,
//                         212,
//                         106,
//                         121,
//                         86,
//                         13,
//                         189,
//                         96,
//                         255,
//                         10,
//                         84,
//                         3,
//                         58,
//                         247,
//                         98,
//                         31,
//                         212,
//                         224,
//                         214,
//                         48,
//                         227,
//                         69,
//                         181,
//                         97
//                     ]
//                 },
//                 "type": "pay",
//                 "gen": "testnet-v1.0",
//                 "gh": {
//                     "type": "Buffer",
//                     "data": [
//                         72,
//                         99,
//                         181,
//                         24,
//                         164,
//                         179,
//                         200,
//                         78,
//                         200,
//                         16,
//                         242,
//                         45,
//                         79,
//                         16,
//                         129,
//                         203,
//                         15,
//                         113,
//                         240,
//                         89,
//                         167,
//                         172,
//                         32,
//                         222,
//                         198,
//                         47,
//                         127,
//                         112,
//                         229,
//                         9,
//                         58,
//                         34
//                     ]
//                 },
//                 "rcv": {
//                     "type": "Buffer",
//                     "data": [
//                         147,
//                         136,
//                         142,
//                         80,
//                         31,
//                         208,
//                         63,
//                         149,
//                         245,
//                         250,
//                         233,
//                         154,
//                         137,
//                         74,
//                         182,
//                         183,
//                         245,
//                         204,
//                         0,
//                         50,
//                         54,
//                         126,
//                         99,
//                         219,
//                         164,
//                         141,
//                         121,
//                         5,
//                         43,
//                         87,
//                         149,
//                         215
//                     ]
//                 }
//             }
//         }
//     ],
//     "attribute_map": {
//         "accounts": "accounts",
//         "apps": "apps",
//         "latestTimestamp": "latest-timestamp",
//         "protocolVersion": "protocol-version",
//         "round": "round",
//         "sources": "sources",
//         "txns": "txns"
//     }
// }

// source Response 
// {
//     "accounts": [],
//     "apps": [],
//     "sources": {
//         "type": "Buffer",
//         "data": [
//             35,
//             112,
//             114,
//             97,
//             103,
//             109,
//             97,
//             32,
//             118,
//             101,
//             114,
//             115,
//             105,
//             111,
//             110,
//             32,
//             53,
//             10,
//             47,
//             47,
//             32,
//             84,
//             104,
//             105,
//             115,
//             32,
//             99,
//             111,
//             100,
//             101,
//             32,
//             105,
//             115,
//             32,
//             109,
//             101,
//             97,
//             110,
//             116,
//             32,
//             102,
//             111,
//             114,
//             32,
//             108,
//             101,
//             97,
//             114,
//             110,
//             105,
//             110,
//             103,
//             32,
//             112,
//             117,
//             114,
//             112,
//             111,
//             115,
//             101,
//             115,
//             32,
//             111,
//             110,
//             108,
//             121,
//             10,
//             47,
//             47,
//             32,
//             73,
//             116,
//             32,
//             115,
//             104,
//             111,
//             117,
//             108,
//             100,
//             32,
//             110,
//             111,
//             116,
//             32,
//             98,
//             101,
//             32,
//             117,
//             115,
//             101,
//             100,
//             32,
//             105,
//             110,
//             32,
//             112,
//             114,
//             111,
//             100,
//             117,
//             99,
//             116,
//             105,
//             111,
//             110,
//             10,
//             97,
//             114,
//             103,
//             95,
//             48,
//             10,
//             98,
//             116,
//             111,
//             105,
//             10,
//             105,
//             110,
//             116,
//             32,
//             49,
//             50,
//             51,
//             52,
//             53,
//             10,
//             61,
//             61
//         ]
//     },
//     "txns": [
//         {
//             "lsig": {
//                 "arg": [
//                     {
//                         "0": 0,
//                         "1": 0,
//                         "2": 0,
//                         "3": 0,
//                         "4": 0,
//                         "5": 0,
//                         "6": 48,
//                         "7": 57
//                     }
//                 ],
//                 "l": {
//                     "0": 5,
//                     "1": 45,
//                     "2": 23,
//                     "3": 129,
//                     "4": 185,
//                     "5": 96,
//                     "6": 18
//                 },
//                 "sig": {
//                     "0": 121,
//                     "1": 119,
//                     "2": 184,
//                     "3": 224,
//                     "4": 25,
//                     "5": 132,
//                     "6": 2,
//                     "7": 28,
//                     "8": 251,
//                     "9": 65,
//                     "10": 237,
//                     "11": 243,
//                     "12": 107,
//                     "13": 144,
//                     "14": 204,
//                     "15": 191,
//                     "16": 224,
//                     "17": 143,
//                     "18": 127,
//                     "19": 67,
//                     "20": 211,
//                     "21": 70,
//                     "22": 225,
//                     "23": 113,
//                     "24": 59,
//                     "25": 182,
//                     "26": 44,
//                     "27": 38,
//                     "28": 49,
//                     "29": 223,
//                     "30": 88,
//                     "31": 197,
//                     "32": 23,
//                     "33": 44,
//                     "34": 84,
//                     "35": 94,
//                     "36": 218,
//                     "37": 170,
//                     "38": 87,
//                     "39": 228,
//                     "40": 153,
//                     "41": 189,
//                     "42": 36,
//                     "43": 84,
//                     "44": 239,
//                     "45": 151,
//                     "46": 112,
//                     "47": 207,
//                     "48": 176,
//                     "49": 90,
//                     "50": 113,
//                     "51": 176,
//                     "52": 70,
//                     "53": 248,
//                     "54": 108,
//                     "55": 134,
//                     "56": 79,
//                     "57": 188,
//                     "58": 174,
//                     "59": 37,
//                     "60": 89,
//                     "61": 56,
//                     "62": 147,
//                     "63": 13
//                 }
//             },
//             "txn": {
//                 "amt": 10000,
//                 "fee": 1000,
//                 "fv": 19648790,
//                 "lv": 19649790,
//                 "snd": {
//                     "type": "Buffer",
//                     "data": [
//                         28,
//                         237,
//                         98,
//                         5,
//                         224,
//                         62,
//                         69,
//                         120,
//                         208,
//                         212,
//                         106,
//                         121,
//                         86,
//                         13,
//                         189,
//                         96,
//                         255,
//                         10,
//                         84,
//                         3,
//                         58,
//                         247,
//                         98,
//                         31,
//                         212,
//                         224,
//                         214,
//                         48,
//                         227,
//                         69,
//                         181,
//                         97
//                     ]
//                 },
//                 "type": "pay",
//                 "gen": "testnet-v1.0",
//                 "gh": {
//                     "type": "Buffer",
//                     "data": [
//                         72,
//                         99,
//                         181,
//                         24,
//                         164,
//                         179,
//                         200,
//                         78,
//                         200,
//                         16,
//                         242,
//                         45,
//                         79,
//                         16,
//                         129,
//                         203,
//                         15,
//                         113,
//                         240,
//                         89,
//                         167,
//                         172,
//                         32,
//                         222,
//                         198,
//                         47,
//                         127,
//                         112,
//                         229,
//                         9,
//                         58,
//                         34
//                     ]
//                 },
//                 "rcv": {
//                     "type": "Buffer",
//                     "data": [
//                         147,
//                         136,
//                         142,
//                         80,
//                         31,
//                         208,
//                         63,
//                         149,
//                         245,
//                         250,
//                         233,
//                         154,
//                         137,
//                         74,
//                         182,
//                         183,
//                         245,
//                         204,
//                         0,
//                         50,
//                         54,
//                         126,
//                         99,
//                         219,
//                         164,
//                         141,
//                         121,
//                         5,
//                         43,
//                         87,
//                         149,
//                         215
//                     ]
//                 }
//             }
//         }
//     ],
//     "attribute_map": {
//         "accounts": "accounts",
//         "apps": "apps",
//         "latestTimestamp": "latest-timestamp",
//         "protocolVersion": "protocol-version",
//         "round": "round",
//         "sources": "sources",
//         "txns": "txns"
//     }
// }
