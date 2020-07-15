const algosdk = require('algosdk');


// function used to wait for a tx confirmation
const waitForConfirmation = async function (algodclient, txId) {
    let status = (await algodclient.status().do());
    let lastRound = status["last-round"];
    while (true) {
        const pendingInfo = await algodclient.pendingTransactionInformation(txId).do();
        if (pendingInfo["confirmed-round"] !== null && pendingInfo["confirmed-round"] > 0) {
            //Got the completed Transaction
            console.log("Transaction " + txId + " confirmed in round " + pendingInfo["confirmed-round"]);
            break;
        }
        lastRound++;
        await algodclient.statusAfterBlock(lastRound).do();
    }
};
// enter token, server, and port
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
    // recover accounts
    // paste in mnemonic phrases here for each account

    // var account1_mnemonic = "PASTE phrase for account 1";
    // var account2_mnemonic = "PASTE phrase for account 2";
    // var account3_mnemonic = "PASTE phrase for account 3"

    var account1_mnemonic = "predict mandate aware dizzy limit match hazard fantasy victory auto fortune hello public dragon ostrich happy blue spray parrot island odor actress only ability hurry";
    var account2_mnemonic = "moon grid random garlic effort faculty fence gym write skin they joke govern home huge there claw skin way bid fit bean damp able only";
    var account3_mnemonic = "mirror zone together remind rural impose balcony position minimum quick manage climb quit draft lion device pluck rug siege robust spirit fine luggage ability actual"


    var account1 = algosdk.mnemonicToSecretKey(account1_mnemonic);
    var account2 = algosdk.mnemonicToSecretKey(account2_mnemonic);
    var account3 = algosdk.mnemonicToSecretKey(account3_mnemonic);
    console.log(account1.addr);
    console.log(account2.addr);
    console.log(account3.addr);

    // Setup the parameters for the multisig account
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
    console.log('Dispense funds to this account on TestNet https://bank.testnet.algorand.network/');
    // await keypress();
    try {
        let algodclient = new algosdk.Algodv2(token, server, port);

        // Get the relevant params from the algod
        let params = await algodclient.getTransactionParams().do();
        // comment out the next two lines to use suggested fee
        params.fee = 1000;
        params.flatFee = true;

        //create a transaction
        // let txn = {
        //     "from": multsigaddr,
        //     "to": account3.addr,
        //     "fee": params.fee,
        //     "amount": 200000,
        //     "firstRound": params.firstRound,
        //     "lastRound": params.lastRound,
        //     "genesisID": params.genesisID,
        //     "genesisHash": params.genesisHash,
        //     "note": new Uint8Array(0)
        // };
        const receiver = account3.addr;
        let note = algosdk.encodeObj("Hello World");
        
        let txn = algosdk.makePaymentTxnWithSuggestedParams(multsigaddr, receiver, 1000000, undefined, note, params);       
        let txId = txn.txID().toString();
        // Sign with first signature
        // At the time of writing this, there is a V2 js sdk bug - signMultisigTransaction - 
        // Error: The transaction sender address and multisig preimage do not match
        // Workaround use V1
        // https://github.com/algorand/js-algorand-sdk/issues/188
        let rawSignedTxn = algosdk.signMultisigTransaction(txn, mparams, account1.sk).blob;
        //sign with second account
        let twosigs = algosdk.appendSignMultisigTransaction(rawSignedTxn, mparams, account2.sk).blob;
        //submit the transaction
        await algodclient.sendRawTransaction(twosigs).do();
        // Wait for confirmation
        await waitForConfirmation(algodclient, txId);

        // Read the transaction from the blockchain
        let confirmedTxn = await algodClient.pendingTransactionInformation(txId).do();
        console.log("Transaction information: %o", confirmedTxn.txn.txn);      
        console.log("Decoded note: %s", algosdk.decodeObj(confirmedTxn.txn.txn.note));

    } catch (err) {
                console.log(err.message);
    }
})().then(process.exit)
