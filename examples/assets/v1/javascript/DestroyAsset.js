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
var cp = {
    fee: 0,
    firstRound: 0,
    lastRound: 0,
    genID: "",
    genHash: ""
}
// Utility function to update params from blockchain
var getChangingParms = async function (algodclient) {
    let params = await algodclient.getTransactionParams();
    cp.firstRound = params.lastRound;
    cp.lastRound = cp.firstRound + parseInt(1000);
    let sfee = await algodclient.suggestedFee();
    cp.fee = sfee.fee;
    cp.genID = params.genesisID;
    cp.genHash = params.genesishashb64;
}

// Function used to wait for a tx confirmation

const waitForConfirmation = async function (algodclient, txId) {
    let lastround = (await algodclient.status()).lastRound;
    while (true) {
        const pendingInfo = await algodclient.pendingTransactionInformation(txId);
        if (pendingInfo.round !== null && pendingInfo.round > 0) {
            //Got the completed Transaction
            console.log("Transaction " + pendingInfo.tx + " confirmed in round " + pendingInfo.round);
            break;
        }
        lastround++;
        await algodclient.statusAfterBlock(lastround);
    }
};


// Recover accounts
// paste in mnemonic phrases here for each account

// var account1_mnemonic = "PASTE your phrase for account 1";
// var account2_mnemonic = "PASTE your phrase for account 2";
// var account3_mnemonic = "PASTE your phrase for account 3"

var account1_mnemonic = "portion never forward pill lunch organ biology" +
    " weird catch curve isolate plug innocent skin grunt" +
    " bounce clown mercy hole eagle soul chunk type absorb trim";
var account2_mnemonic = "place blouse sad pigeon wing warrior wild script" +
    " problem team blouse camp soldier breeze twist mother" +
    " vanish public glass code arrow execute convince ability" +
    " there";
var account3_mnemonic = "image travel claw climb bottom spot path roast" +
    " century also task cherry address curious save item" +
    " clean theme amateur loyal apart hybrid steak about blanket"


var recoveredAccount1 = algosdk.mnemonicToSecretKey(account1_mnemonic);
var recoveredAccount2 = algosdk.mnemonicToSecretKey(account2_mnemonic);
var recoveredAccount3 = algosdk.mnemonicToSecretKey(account3_mnemonic);
console.log(recoveredAccount1.addr);
console.log(recoveredAccount2.addr);
console.log(recoveredAccount3.addr);

// Instantiate the algod wrapper
let algodclient = new algosdk.Algod(token, server, port);

// Debug Console should look similar to this

// THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM  
// AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU   
// 3ZQ3SHCYIKSGK7MTZ7PE7S6EDOFWLKDQ6RYYVMT7OHNQ4UJ774LE52AQCU   


(async () => {
    // Destroy and Asset:
    // paste in the asset id from the create asset tutorial
    let assetID = 327472;
    // All of the created assets should now be back in the creators
    // Account so we can delete the asset.
    // If this is not the case the asset deletion will fail
    // The address for the from field must be the manager account
    // Which is currently the creator addr1
    addr = recoveredAccount1.addr;
    note = undefined;

    // First update changing transaction parameters
    // We will account for changing transaction parameters
    // before every transaction in this example
    await getChangingParms(algodclient);

    // if all assets are held by the asset creator,
    // the asset creator can sign and issue "txn" to remove the asset from the ledger. 
    let dtxn = algosdk.makeAssetDestroyTxn(addr, cp.fee, cp.firstRound, cp.lastRound, note, cp.genHash, cp.genID, assetID);
    // The transaction must be signed by the manager which 
    // is currently set to account1
    rawSignedTxn = dtxn.signTxn(recoveredAccount1.sk)
    let dtx = (await algodclient.sendRawTransaction(rawSignedTxn));
    console.log("Asset Destroy Transaction : " + dtx.txId);
    // wait for transaction to be confirmed
    await waitForConfirmation(algodclient, dtx.txId);

    // The account3 and account1 should no longer contain the asset as it has been destroyed
    console.log("Asset ID: " + assetID);
    act = await algodclient.accountInformation(recoveredAccount3.addr);
    console.log("Account Information for Account 3: " + JSON.stringify(act.assets[assetID]));
    act = await algodclient.accountInformation(recoveredAccount1.addr);
    console.log("Account Information for Account 1: " + JSON.stringify(act.assets[assetID]));

    // Notice that although the asset was destroyed, the asset id and associated 
    // metadata still exists in account holdings for Account 3. 
    // When you destroy an asset, the global parameters associated with that asset
    // (manager addresses, name, etc.) are deleted from the creator's balance record (Account 1).
    // However, holdings are not deleted automatically -- users still need to close out of the deleted asset.
    // This is necessary for technical reasons because we currently can't have a single transaction touch potentially 
    // thousands of accounts (all the holdings that would need to be deleted).


    // Asset Destroy Transaction: 2M72QUASX3MLKSSPLAXOE2KZC6BQWBDEOGP4FANLNUDCPT3U4EQQ
    // Transaction 2M72QUASX3MLKSSPLAXOE2KZC6BQWBDEOGP4FANLNUDCPT3U4EQQ confirmed in round 6026503
    // Asset ID: 327472
    // Account Information for Account 3: { "creator": "", "amount": 0, "frozen": true }
    // Account Information for Account 1: undefined



})().catch(e => {
    console.log(e);
    console.trace();
});


