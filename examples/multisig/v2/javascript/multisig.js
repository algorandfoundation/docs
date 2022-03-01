const algosdk = require('algosdk');



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
    // never use mnemonics in production code, replace for demo purposes only
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
    console.log("Add funds to account using the TestNet Dispenser: ");
    console.log("https://dispenser.testnet.aws.algodev.network?account=" + multsigaddr);
    console.log("Once funded, press any key to continue");
    await keypress();
    try {
        let algodclient = new algosdk.Algodv2(token, server, port);

        // Get the relevant params from the algod
        let params = await algodclient.getTransactionParams().do();
        // comment out the next two lines to use suggested fee
        // params.fee = 1000;
        // params.flatFee = true;

        const receiver = account3.addr;
        let names = '{"firstName":"John", "lastName":"Doe"}';
        const enc = new TextEncoder();
        const note = enc.encode(names);


        let txn = algosdk.makePaymentTxnWithSuggestedParams(multsigaddr, receiver, 100000, undefined, note, params);
        let txId = txn.txID().toString();
        // Sign with first signature

        let rawSignedTxn = algosdk.signMultisigTransaction(txn, mparams, account1.sk).blob;
        //sign with second account
        let twosigs = algosdk.appendSignMultisigTransaction(rawSignedTxn, mparams, account2.sk).blob;
        //submit the transaction
        await algodclient.sendRawTransaction(twosigs).do();

        // Wait for confirmation
        const confirmedTxn = await algosdk.waitForConfirmation(algodclient, txId, 4);
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
