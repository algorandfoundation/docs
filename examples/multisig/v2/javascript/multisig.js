const algosdk = require('algosdk');

/**
 * utility function to wait on a transaction to be confirmed
 * the timeout parameter indicates how many rounds do you wish to check pending transactions for
 */
const waitForConfirmation = async function (algodclient, txId, timeout) {
    // Wait until the transaction is confirmed or rejected, or until 'timeout'
    // number of rounds have passed.
    //     Args:
    // txId(str): the transaction to wait for
    // timeout(int): maximum number of rounds to wait
    // Returns:
    // pending transaction information, or throws an error if the transaction
    // is not confirmed or rejected in the next timeout rounds
    if (algodclient == null || txId == null || timeout < 0) {
        throw "Bad arguments.";
    }
    let status = (await algodclient.status().do());
    if (status == undefined) throw new Error("Unable to get node status");
    let startround = status["last-round"] + 1;
    let currentround = startround;

    while (currentround < (startround + timeout)) {
        let pendingInfo = await algodclient.pendingTransactionInformation(txId).do();
        if (pendingInfo != undefined) {
            if (pendingInfo["confirmed-round"] !== null && pendingInfo["confirmed-round"] > 0) {
                //Got the completed Transaction
                return pendingInfo;
            }
            else {
                if (pendingInfo["pool-error"] != null && pendingInfo["pool-error"].length > 0) {
                    // If there was a pool error, then the transaction has been rejected!
                    throw new Error("Transaction Rejected" + " pool error" + pendingInfo["pool-error"]);
                }
            }
        }
        await algodclient.statusAfterBlock(currentround).do();
        currentround++;
    }
    throw new Error("Transaction not confirmed after " + timeout + " rounds!");
};
// enter token, server, and port
// const token = <algod-token>;
// const server = <algod-address>;
// const port = <algod-port>;
// sandbox
const token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
const server = "http://localhost";
const port = 4001;

const keypress = async () => {
    process.stdin.setRawMode(true)
    return new Promise(resolve => process.stdin.once('data', () => {
        process.stdin.setRawMode(false)
        resolve()
    }))
}

(async () => {
    // recover accounts
    // paste in mnemonic phrases here for each account

    // let account1_mnemonic = "PASTE phrase for account 1";
    // let account2_mnemonic = "PASTE phrase for account 2";
    // let account3_mnemonic = "PASTE phrase for account 3"

    let account1_mnemonic = "patrol target joy dial ethics flip usual fatigue bulb security prosper brand coast arch casino burger inch cricket scissors shoe evolve eternal calm absorb school";
    let account2_mnemonic = "genius inside turtle lock alone blame parent civil depend dinosaur tag fiction fun skill chief use damp daughter expose pioneer today weasel box about silly";
    let account3_mnemonic = "off canyon mystery cable pluck emotion manual legal journey grit lunch include friend social monkey approve lava steel school mango auto cactus huge ability basket"

    let account1 = algosdk.mnemonicToSecretKey(account1_mnemonic);
    let account2 = algosdk.mnemonicToSecretKey(account2_mnemonic);
    let account3 = algosdk.mnemonicToSecretKey(account3_mnemonic);
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

    let multsigaddr = algosdk.multisigAddress(mparams);
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

        const receiver = account3.addr;
        let names = '{"firstName":"John", "lastName":"Doe"}';
        const enc = new TextEncoder();
        const note = enc.encode(names);


        let txn = algosdk.makePaymentTxnWithSuggestedParams(multsigaddr, receiver, 1000000, undefined, note, params);
        let txId = txn.txID().toString();
        // Sign with first signature

        let rawSignedTxn = algosdk.signMultisigTransaction(txn, mparams, account1.sk).blob;
        //sign with second account
        let twosigs = algosdk.appendSignMultisigTransaction(rawSignedTxn, mparams, account2.sk).blob;
        //submit the transaction
        await algodclient.sendRawTransaction(twosigs).do();

        // Wait for confirmation
        let confirmedTxn = await waitForConfirmation(algodclient, txId, 4);
        //Get the completed Transaction
        console.log("Transaction " + txId + " confirmed in round " + confirmedTxn["confirmed-round"]);
        let mytxinfo = JSON.stringify(confirmedTxn.txn.txn, undefined, 2);
        console.log("Transaction information: %o", mytxinfo);
        let string = new TextDecoder().decode(confirmedTxn.txn.txn.note);
        console.log("Note field: ", string);
        const obj = JSON.parse(string);
        console.log("Note first name: %s", obj.firstName);


    } catch (err) {
        console.log(err.message);
    }
})().then(process.exit)
