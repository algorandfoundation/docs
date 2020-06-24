const algosdk = require('algosdk');


// Retrieve the token, server and port values for your installation in the 
// algod.net and algod.token files within the data directory

// UPDATE THESE VALUES
// const token = "TOKEN";
// const server = "SERVER";
// const port = PORT;

// hackathon
// const token = "ef920e2e7e002953f4b29a8af720efe8e4ecc75ff102b165e0472834b25832c1";
// const server = "http://hackathon.algodev.network";
// const port = 9100;

// sandbox
const token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
const server = "http://localhost";
const port = 4001;

// Structure for changing blockchain params
// var cp = {
//     fee: 0,
//     firstRound: 0,
//     lastRound: 0,
//     genID: "",
//     genHash: ""
// }
// Utility function to update params from blockchain
// var getChangingParms = async function (algodclient) {
//     let params = await algodclient.getTransactionParams();
//     cp.firstRound = params.lastRound;
//     cp.lastRound = cp.firstRound + parseInt(1000);
//     let sfee = await algodclient.suggestedFee();
//     cp.fee = sfee.fee;
//     cp.genID = params.genesisID;
//     cp.genHash = params.genesishashb64;
// }

// Function used to wait for a tx confirmation

const waitForConfirmation = async function (algodclient, txId) {
    let response = await algodclient.status().do();
    let lastround = response["last-round"];
    while (true) {
        const pendingInfo = await algodclient.pendingTransactionInformation(txId).do();
        if (pendingInfo["confirmed-round"] !== null && pendingInfo["confirmed-round"] > 0) {
            //Got the completed Transaction
            console.log("Transaction " + txId + " confirmed in round " + pendingInfo["confirmed-round"]);
            break;
        }
        lastround++;
        await algodclient.statusAfterBlock(lastround).do();
    }
};

// Function used to print created asset for account and assetid
const printCreatedAsset = async function (algodclient, account, assetid) {
    // note: if you have an indexer instance available it is easier to just use this
    //     let accountInfo = await indexerClient.searchAccounts()
    //    .assetID(assetIndex).do();
    // and in the loop below use this to extract the asset for a particular account
    // accountInfo['accounts'][idx][account]);
    let accountInfo = await algodclient.accountInformation(account).do();
    for (idx = 0; idx < accountInfo['created-assets'].length; idx++) {
        let scrutinizedAsset = accountInfo['created-assets'][idx];
        if (scrutinizedAsset['index'] == assetid) {
            console.log("AssetID = " + scrutinizedAsset['index']);
            let myparms = JSON.stringify(scrutinizedAsset['params'], undefined, 2);
            console.log("parms = " + myparms);
            break;
        }
    }
};
// Function used to print asset holding for account and assetid
const printAssetHolding = async function (algodclient, account, assetid) {
    // note: if you have an indexer instance available it is easier to just use this
    //     let accountInfo = await indexerClient.searchAccounts()
    //    .assetID(assetIndex).do();
    // and in the loop below use this to extract the asset for a particular account
    // accountInfo['accounts'][idx][account]);
    let accountInfo = await algodclient.accountInformation(account).do();
    for (idx = 0; idx < accountInfo['assets'].length; idx++) {
        let scrutinizedAsset = accountInfo['assets'][idx];
        if (scrutinizedAsset['asset-id'] == assetid) {
            let myassetholding = JSON.stringify(scrutinizedAsset, undefined, 2);
            console.log("assetholdinginfo = " + myassetholding);
            break;
        }
    }
};


// Recover accounts
// paste in mnemonic phrases here for each account

// var account1_mnemonic = "PASTE your phrase for account 1";
// var account2_mnemonic = "PASTE your phrase for account 2";
// var account3_mnemonic = "PASTE your phrase for account 3"

// var account1_mnemonic = "portion never forward pill lunch organ biology" +
//     " weird catch curve isolate plug innocent skin grunt" +
//     " bounce clown mercy hole eagle soul chunk type absorb trim";
// var account2_mnemonic = "place blouse sad pigeon wing warrior wild script" +
//     " problem team blouse camp soldier breeze twist mother" +
//     " vanish public glass code arrow execute convince ability" +
//     " there";
// var account3_mnemonic = "image travel claw climb bottom spot path roast" +
//     " century also task cherry address curious save item" +
//     " clean theme amateur loyal apart hybrid steak about blanket"
var account1_mnemonic = "canal enact luggage spring similar zoo couple stomach shoe laptop middle wonder eager monitor weather number heavy skirt siren purity spell maze warfare ability ten";
var account2_mnemonic = "beauty nurse season autumn curve slice cry strategy frozen spy panic hobby strong goose employ review love fee pride enlist friend enroll clip ability runway";
var account3_mnemonic = "picnic bright know ticket purity pluck stumble destroy ugly tuna luggage quote frame loan wealth edge carpet drift cinnamon resemble shrimp grain dynamic absorb edge";



var recoveredAccount1 = algosdk.mnemonicToSecretKey(account1_mnemonic);
var recoveredAccount2 = algosdk.mnemonicToSecretKey(account2_mnemonic);
var recoveredAccount3 = algosdk.mnemonicToSecretKey(account3_mnemonic);
console.log(recoveredAccount1.addr);
console.log(recoveredAccount2.addr);
console.log(recoveredAccount3.addr);

// Instantiate the algod wrapper
//let algodclient = new algosdk.Algod(token, server, port);
let algodclient = new algosdk.Algodv2(token, server, port);

// Debug Console should look similar to this

// THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM  
// AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU   
// 3ZQ3SHCYIKSGK7MTZ7PE7S6EDOFWLKDQ6RYYVMT7OHNQ4UJ774LE52AQCU   


(async () => {
    // Destroy and Asset:
    // paste in the asset id from the create asset tutorial
    let assetID = 2653801;
    // All of the created assets should now be back in the creators
    // Account so we can delete the asset.
    // If this is not the case the asset deletion will fail
    // The address for the from field must be the manager account
    // Which is currently the creator addr1

    // First update changing transaction parameters
    // We will account for changing transaction parameters
    // before every transaction in this example
    let params = await algodclient.getTransactionParams().do();
    //comment out the next two lines to use suggested fee
    params.fee = 1000;
    params.flatFee = true;
    console.log(params); 

    addr = recoveredAccount1.addr;
    note = undefined;

    // if all assets are held by the asset creator,
    // the asset creator can sign and issue "txn" to remove the asset from the ledger. 
    let dtxn = algosdk.makeAssetDestroyTxnWithSuggestedParams(addr, note, assetID, params);
    // The transaction must be signed by the manager which 
    // is currently set to account1
    rawSignedTxn = dtxn.signTxn(recoveredAccount1.sk);
    let dtx = (await algodclient.sendRawTransaction(rawSignedTxn).do());
    console.log("Asset Destroy Transaction : " + dtx.txId);
    // wait for transaction to be confirmed
    await waitForConfirmation(algodclient, dtx.txId);

     // The account3 and account1 should no longer contain the asset as it has been destroyed
    console.log("Asset ID: " + assetID);
    console.log("Account 1 = " + recoveredAccount1.addr);
    await printCreatedAsset(algodclient, recoveredAccount1.addr, assetID);
    await printAssetHolding(algodclient, recoveredAccount1.addr, assetID);
    console.log("Account 3 = " + recoveredAccount3.addr);
    await printAssetHolding(algodclient, recoveredAccount3.addr, assetID);  
  

    // Notice that although the asset was destroyed, the asset id and associated 
    // metadata still exists in account holdings for Account 3. 
    // When you destroy an asset, the global parameters associated with that asset
    // (manager addresses, name, etc.) are deleted from the creator's balance record (Account 1).
    // However, holdings are not deleted automatically -- users still need to close out of the deleted asset.
    // This is necessary for technical reasons because we currently can't have a single transaction touch potentially 
    // thousands of accounts (all the holdings that would need to be deleted).

    // your console/terminal output should look similar to this: 
    // Transaction: KMSVPDQMVGIASU2ZLKRBSWOQVXR3EOJ4Z4WKS5GTB7MKE4D57L6Q
    // Transaction KMSVPDQMVGIASU2ZLKRBSWOQVXR3EOJ4Z4WKS5GTB7MKE4D57L6Q confirmed in round 3961877
    // Asset ID: 2653785
    // Account 1 = ATTR6RUEHHBHXKUHT4GUOYWNBVDV2GJ5FHUWCSFZLHD55EVKZWOWSM7ABQ
    // Account 3 = IWR4CLLCN2TIVX2QPVVKVR5ER5OZGMWAV5QB2UIPYMPKBPLJZX4C37C4AA
    // AssetExample.js: 480
    // assetholdinginfo = {
    //     "amount": 0,
    //     "asset-id": 2653785,
    //     "creator": "",
    //     "is-frozen": true
    // }



})().catch(e => {
    console.log(e);
    console.trace();
});


