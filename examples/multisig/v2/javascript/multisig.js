const algosdk = require('algosdk');
const { waitForConfirmation } = require('./utils/waitForConfirmation');

// sandbox
const token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
const server = "http://localhost";
const port = 4001;

const main = async () => {
    // mnemonic phrases here for each of your test accounts
    const account1_mnemonic = "< mnemonic 1 >";
    const account2_mnemonic = "< mnemonic 2 >";
    const account3_mnemonic = "< mnemonic 3 >";

    const account1 = algosdk.mnemonicToSecretKey(account1_mnemonic);
    const account2 = algosdk.mnemonicToSecretKey(account2_mnemonic);
    const account3 = algosdk.mnemonicToSecretKey(account3_mnemonic);

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

    const multsigaddr = algosdk.multisigAddress(mparams);
    console.log("Multisig Address: " + multsigaddr);

    const algodClient = new algosdk.Algodv2(token, server, port);

    // Get the relevant params from the algod
    let params = await algodClient.getTransactionParams().do();
    // comment out the next two lines to use suggested fee
    params.fee = 1000;
    params.flatFee = true;

    const receiver = account3.addr;
    const names = '{"firstName":"John", "lastName":"Doe"}';
    const enc = new TextEncoder();
    const note = enc.encode(names);

    const txn = algosdk.makePaymentTxnWithSuggestedParams(multsigaddr, receiver, 1000000, undefined, note, params);
    const txId = txn.txID().toString();
    
    // Sign with first signature
    const rawSignedTxn = algosdk.signMultisigTransaction(txn, mparams, account1.sk).blob;
    //sign with second account
    const twosigs = algosdk.appendSignMultisigTransaction(rawSignedTxn, mparams, account2.sk).blob;
    
    //submit the transaction
    await algodClient.sendRawTransaction(twosigs).do();

    // Wait for confirmation
    const confirmedTxn = await waitForConfirmation(algodClient, txId, 4);
    console.log("Transaction " + txId + " confirmed in round " + confirmedTxn["confirmed-round"]);

    const noteString = new TextDecoder().decode(confirmedTxn.txn.txn.note);
    console.log("Note field: ", noteString);

    const obj = JSON.parse(noteString);
    console.log("Note first name: %s", obj.firstName);
}

main().then().catch(e => {
    const error = e.body && e.body.message ? e.body.message : e;
    console.log(error);
});
