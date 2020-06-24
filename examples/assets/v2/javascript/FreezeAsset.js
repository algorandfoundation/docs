const algosdk = require('algosdk');
// Retrieve the token, server and port values for your installation in the 
// algod.net and algod.token files within the data directory

// UPDATE THESE VALUES
// const token = "TOKEN";
// const server = "SERVER";
// const port = PORT;

//hackathon
// const token = "ef920e2e7e002953f4b29a8af720efe8e4ecc75ff102b165e0472834b25832c1";
// const server = "http://hackathon.algodev.network";
// const port = 9100;

// sandbox
const token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
const server = "http://localhost";
const port = 4001;

// Structure for changing blockchain params
var cp = {
    fee: 0,
    firstRound: 0,
    lastRound: 0,
    genID: "",
    genHash: ""
}
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
// let algodclient = new algosdk.Algod(token, server, port);
let algodclient = new algosdk.Algodv2(token, server, port);

// Debug Console should look similar to this

// THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM  
// AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU   
// 3ZQ3SHCYIKSGK7MTZ7PE7S6EDOFWLKDQ6RYYVMT7OHNQ4UJ774LE52AQCU   


(async () => {
    // Freeze Asset:
    // replace with your  assetid
    let assetID = 2653801;

    // The asset was created and configured to allow freezing an account
    // If the freeze address is set "", it will no longer be possible to do this.
    // In this example we will now freeze account3 from transacting with the 
    // The newly created asset. 
    // The freeze transaction is sent from the freeze acount
    // Which in this example is account2 
    from = recoveredAccount2.addr;
    freezeTarget = recoveredAccount3.addr;
    freezeState = true;
    note = undefined;

    // First update changing transaction parameters
    // We will account for changing transaction parameters
    // before every transaction in this example
    let params = await algodclient.getTransactionParams().do();
    //comment out the next two lines to use suggested fee
    params.fee = 1000;
    params.flatFee = true;
    console.log(params);


    // The freeze transaction needs to be signed by the freeze account
    let ftxn = algosdk.makeAssetFreezeTxnWithSuggestedParams(from, note,
        assetID, freezeTarget, freezeState, params);

    // Must be signed by the freeze account   
    rawSignedTxn = ftxn.signTxn(recoveredAccount2.sk)
    let ftx = (await algodclient.sendRawTransaction(rawSignedTxn).do());
    console.log("Asset Freeze Transaction id: " + ftx.txId);
    // wait for transaction to be confirmed
    await waitForConfirmation(algodclient, ftx.txId);

    // You should now see the asset is frozen listed in the account information
    console.log("Account 3 = " + recoveredAccount3.addr);
    await printAssetHolding(algodclient, recoveredAccount3.addr, assetID);

    //you should see console/terminal output similar to this witht he frozen vales set to true
    //    Asset Freeze Transaction id: CZ7VYA7UFGRVJ4EHQOLKCPRESDJC4ULHUTWM3QCAPZR3K4KVT2IQ
    //    Transaction CZ7VYA7UFGRVJ4EHQOLKCPRESDJC4ULHUTWM3QCAPZR3K4KVT2IQ confirmed in round 4274065
    //    Account Information for: { "creator": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM", "amount": 10, "frozen": true }

    // Attempt to spend Frozen Asset (this should fail)
    sender = recoveredAccount3.addr;
    recipient = recoveredAccount2.addr;
    revocationTarget = undefined;
    closeRemainderTo = undefined;
    // Amount of the asset attempting to transfer
    amount = 3;
    let ftxnattempt = algosdk.makeAssetTransferTxn(recoveredAccount3.addr, recoveredAccount2.addr, closeRemainderTo, revocationTarget, cp.fee, amount, cp.firstRound, cp.lastRound, note, cp.genHash, cp.genID, assetID);
    rawSignedTxn = ftxnattempt.signTxn(recoveredAccount3.sk);
    // this should fail as we are attempting to send from a frozen account  
    let ftxna = (await algodclient.sendRawTransaction(rawSignedTxn));
    console.log("Asset Transaction id : " + ftxna.txId);
})().catch(e => {
    console.log(e);
    console.trace();
});


