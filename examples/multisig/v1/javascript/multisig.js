const algosdk = require('algosdk');


// function used to wait for a tx confirmation
var waitForConfirmation = async function (algodclient, txId) {
    while (true) {
        let lastround = (await algodclient.status()).lastRound;
        let pendingInfo = await algodclient.pendingTransactionInformation(txId);
        if (pendingInfo.round != null && pendingInfo.round > 0) {
            //Got the completed Transaction
            console.log("Transaction " + pendingInfo.tx + " confirmed in round " + pendingInfo.round);
            break;
        }
        await algodclient.statusAfterBlock(lastround + 1);
    }
};
// const token = <algod-token>;
// const server = <algod-address>;
// const port = <algod-port>;
// sandbox
const token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
const server = "http://localhost";
const port = 4001;

const keypress = async() => {
    process.stdin.setRawMode(true)
    return new Promise(resolve => process.stdin.once('data', () => {
                process.stdin.setRawMode(false)
        resolve()
    }))
}

(async() => {
    // restore an account

    // var account1_mnemonic = "PASTE your phrase for account 1";
    // var account2_mnemonic = "PASTE your phrase for account 2";
    // var account3_mnemonic = "PASTE your phrase for account 3"

    var account1_mnemonic = "predict mandate aware dizzy limit match hazard fantasy victory auto fortune hello public dragon ostrich happy blue spray parrot island odor actress only ability hurry";
    var account2_mnemonic = "moon grid random garlic effort faculty fence gym write skin they joke govern home huge there claw skin way bid fit bean damp able only";
    var account3_mnemonic = "mirror zone together remind rural impose balcony position minimum quick manage climb quit draft lion device pluck rug siege robust spirit fine luggage ability actual"


    var account1 = algosdk.mnemonicToSecretKey(account1_mnemonic);
    var account2 = algosdk.mnemonicToSecretKey(account2_mnemonic);
    var account3 = algosdk.mnemonicToSecretKey(account3_mnemonic);
    console.log(account1.addr);
    console.log(account2.addr);
    console.log(account3.addr);

    //Setup the parameters for the multisig account
    const mparams = {
                version: 1,
        threshold: 2,
        addrs: [
            account1.addr,
            account2.addr,
            account3.addr,
        ],
    };

    var multsigaddr = algosdk.multisigAddress(mparams);
    console.log("Multisig Address: " + multsigaddr);
    //Pause execution to allow using the dispenser on testnet to put tokens in account
    console.log('Make sure address above has tokens using the dispenser');
   // await keypress();
    try {
    //    let algodclient = new algosdk.Algod(token, server, port);
        let algodClient = new algosdk.Algodv2(token, server, port);
        // Construct the transaction
        let params = await algodClient.getTransactionParams().do();
        // comment out the next two lines to use suggested fee
        params.fee = 1000;
        params.flatFee = true;
        //create a transaction
        // let txn = {
        //     "from": multsigaddr,
        //     "to": account3.addr,
        //     "fee": params.fee,
        //     "amount": 200000,
        //     "firstRound": firstRoundint,
        //     "lastRound": lastRoundint,
        //     "genesisID": params.genesisID,
        //     "genesisHash": params.genesisHash,
        //     "note": new Uint8Array(0)
        // };

        // let txn = {
        //     "from": multsigaddr,
        //     "to": account3.addr,
        //     "amount": 200000,
        //     "suggestedParams": params,
        //     "note": new Uint8Array(0)
        // };

        // receiver 
        const receiver = account3.addr;
        let note = algosdk.encodeObj("Hello World");
        amount = 200000;
        //  image BUG - 
        //  let txn = algosdk.makePaymentTxnWithSuggestedParams(multsigaddr, receiver, amount, undefined, note, params); 
        //Sign wiith first signature
        let rawSignedTxn = algosdk.signMultisigTransaction(txn, mparams, account1.sk).blob;
        //sign with second account
        let twosigs = algosdk.appendSignMultisigTransaction(rawSignedTxn, mparams, account2.sk).blob;
        //submit the transaction
        let tx = (await algodClient.sendRawTransaction(twosigs));
        // Wait for confirmation
        await waitForConfirmation(algodClient, tx.txId);
        console.log("Transaction : " + tx.txId);

    } catch (err) {
                console.log(err);
    }
})().then(process.exit)
