const algosdk = require('algosdk');
const fs = require('fs');
let client = null;

// make connection to node
async function setupClient() {
    if (client == null) {
        // const ALGOD_API_ADDR = <algod-address>;
        // const ALGOD_API_TOKEN = <algod-token>;
        // const port = <port-number>;
        const token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
        const server = "http://localhost";
        const port = 4001;
        let algodClient = new algosdk.Algodv2(token, server, port);
        client = algodClient;
    } else {
        return client;
    }
    return client;
}

// recover acccounts for example
// Shown for demonstration purposes. NEVER reveal secret mnemonics in practice.
function recoverAccount1() {
    const passphrase = "<your-25-word-mnemonic>";
    let myAccount = algosdk.mnemonicToSecretKey(passphrase);
    return myAccount;
}
function recoverAccount2() {
    const passphrase = "<your-25-word-mnemonic>";
    let myAccount = algosdk.mnemonicToSecretKey(passphrase);
    return myAccount;
}
function recoverAccount3() {
    const passphrase = "<your-25-word-mnemonic>";
    let myAccount = algosdk.mnemonicToSecretKey(passphrase);
    return myAccount;
}





const keypress = async () => {
    process.stdin.setRawMode(true)
    return new Promise(resolve => process.stdin.once('data', () => {
        process.stdin.setRawMode(false)
        resolve()
    }))
}

async function writeSignedMultisigTransctionToFile() {
  
    let algodClient = await setupClient();
    let account1 = await recoverAccount1();
    console.log("address1: %s", account1.addr)
    let account2 = await recoverAccount2();
    console.log("address2: %s", account2.addr)
    let account3 = await recoverAccount3();
    console.log("address3: %s", account3.addr)

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
    console.log('Dispense funds to this account on TestNet https://dispenser.testnet.aws.algodev.network/');
    // await keypress();
    try {
        // Get the relevant params from the algod
        let params = await algodClient.getTransactionParams().do();
        // comment out the next two lines to use suggested fee
        // params.fee = 1000;
        // params.flatFee = true;

        const receiver = account3.addr;
        let names = '{"firstName":"John", "lastName":"Doe"}';
        const enc = new TextEncoder();
        const note = enc.encode(names);

        let txn = algosdk.makePaymentTxnWithSuggestedParams(multsigaddr, receiver, 1000000, undefined, note, params);
        let txId = txn.txID().toString();
        // Sign with first signature

        let stxn = algosdk.signMultisigTransaction(txn, mparams, account1.sk).blob;
        //sign with second account - this can be signed before writing twosigs offline or when reading in, if second tx not signed here just writting stxn
        //let twosigs = algosdk.appendSignMultisigTransaction(stxn, mparams, account2.sk).blob;

        fs.writeFileSync('./signedmultisig.stxn', stxn ); 
        console.log("The stxn file was saved!");
        // write out mparams as json file
        let mparamsinfo = JSON.stringify(mparams, undefined, 2); 
        fs.writeFileSync('./multisigparams.json', mparamsinfo)        
        console.log("The multisig params file was saved!"); 

        // fs.writeFileSync('./multisig.mparams', algosdk.encodeObj(mparams))        
        // console.log("The multisig params file was saved!");        
       
    } catch (err) {
        console.log(err.message);
    }
};

async function readSignedMultisigTransctionFromFile() {
    let algodClient = await setupClient();
    let account1 = await recoverAccount1();
    console.log("address1: %s", account1.addr)
    let account2 = await recoverAccount2();
    console.log("address2: %s", account2.addr)
    let account3 = await recoverAccount3();
    console.log("address3: %s", account3.addr)
    // read transaction from file and sign it
    // read signed transaction from file
    let stxn = fs.readFileSync("./signedmultisig.stxn");
    // read mparams from json file

    const mparamstr = fs.readFileSync('./multisigparams.json');
    let string = new TextDecoder().decode(mparamstr);
    const mparams = JSON.parse(string);   
    // read multisig mparams
    let multsigaddr = algosdk.multisigAddress(mparams);
    console.log("Multisig Address: " + multsigaddr);
    //Pause execution to allow using the dispenser on testnet to put tokens in account
    console.log('Dispense funds to this account on TestNet https://bank.testnet.algorand.network/');
    // await keypress();
    try {


        // sign with second account
        let twosigs = algosdk.appendSignMultisigTransaction(stxn, mparams, account2.sk).blob;
        // submit the transaction     
        let txId = (await algodClient.sendRawTransaction(twosigs).do());        
        // Wait for confirmation
        const confirmedTxn = await algosdk.waitForConfirmation(algodClient, txId.txId, 4);
        //Get the completed Transaction
        console.log("Transaction " + txId.txId + " confirmed in round " + confirmedTxn["confirmed-round"]);
        // let mytxinfo = JSON.stringify(confirmedTxn.txn.txn, undefined, 2);
        // console.log("Transaction information: %o", mytxinfo);
        let string = new TextDecoder().decode(confirmedTxn.txn.txn.note);
        console.log("Note field: ", string);
        const obj = JSON.parse(string);
        console.log("Note first name: %s", obj.firstName);
    } catch (err) {
        console.log(err.message);
    }
};

async function writeUnsignedMultisigTransctionToFile() {
    let algodClient = await setupClient();
    let account1 = await recoverAccount1();
    console.log("address1: %s", account1)
    let account2 = await recoverAccount2();
    console.log("address2: %s", account2)
    let account3 = await recoverAccount3();
    console.log("address3: %s", account3)

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
    // Pause execution to allow using the dispenser on testnet to put tokens in account
    console.log('Dispense funds to this account on TestNet https://bank.testnet.algorand.network/');
    // await keypress();
    try {

        // Get the relevant params from the algod
        let params = await algodClient.getTransactionParams().do();
        // comment out the next two lines to use suggested fee
        params.fee = 1000;
        params.flatFee = true;

        const receiver = account3.addr;
        let names = '{"firstName":"John", "lastName":"Doe"}';
        const enc = new TextEncoder();
        const note = enc.encode(names);
        let txn = algosdk.makePaymentTxnWithSuggestedParams(multsigaddr, receiver, 1000000, undefined, note, params);
 
        // Save transaction to file
        fs.writeFileSync('./unsigned.txn', algosdk.encodeUnsignedTransaction( txn ));   
        console.log("Unsigned tx was saved!");
        // write out mparams as json file
        let mparamsinfo = JSON.stringify(mparams, undefined, 2); 
        fs.writeFileSync('./multisigparams.json', mparamsinfo)               
        console.log("The multisig params file was saved!"); 

    } catch (err) {
        console.log(err.message);
    }
};

async function readUnsignedMultisigTransctionFromFile() {
    let algodClient = await setupClient();
    let account1 = await recoverAccount1();
    console.log("address1: %s", account1.addr)
    let account2 = await recoverAccount2();
    console.log("address2: %s", account2.addr)
    let account3 = await recoverAccount3();
    console.log("address3: %s", account3.addr)
    
    // read transaction from file and sign it
    let txn = algosdk.decodeUnsignedTransaction(fs.readFileSync('./unsigned.txn'));  

    // read mparams from json file

    const mparamstr = fs.readFileSync('./multisigparams.json');
    let string = new TextDecoder().decode(mparamstr);
    const mparams = JSON.parse(string);   
    //const mparams = algosdk.decodeObj(fs.readFileSync('./multisig.mparams'));
    let multsigaddr = algosdk.multisigAddress(mparams);
    console.log("Multisig Address: " + multsigaddr);
    //Pause execution to allow using the dispenser on testnet to put tokens in account
    console.log('Dispense funds to this account on TestNet https://bank.testnet.algorand.network/');
    // await keypress();
    try {
        let txId = txn.txID().toString();
        // Sign with first signature

        let rawSignedTxn = algosdk.signMultisigTransaction(txn, mparams, account1.sk).blob;
        //sign with second account
        let twosigs = algosdk.appendSignMultisigTransaction(rawSignedTxn, mparams, account2.sk).blob;
        //submit the transaction
        await algodClient.sendRawTransaction(twosigs).do();
        // Wait for confirmation
        let confirmedTxn = await algosdk.waitForConfirmation(algodClient, txId, 4);
        //Get the completed Transaction
        console.log("Transaction " + txId + " confirmed in round " + confirmedTxn["confirmed-round"]);
        // let mytxinfo = JSON.stringify(confirmedTxn.txn.txn, undefined, 2);
        // console.log("Transaction information: %o", mytxinfo);
        let string = new TextDecoder().decode(confirmedTxn.txn.txn.note);
        console.log("Note field: ", string);
        const obj = JSON.parse(string);
        console.log("Note first name: %s", obj.firstName);
    } catch (err) {
        console.log(err.message);
    }
};
async function testMultisigSigned() {
    await writeSignedMultisigTransctionToFile();
    await readSignedMultisigTransctionFromFile();    
}
async function testMultisigUnsigned() {
    await writeUnsignedMultisigTransctionToFile();
    await readUnsignedMultisigTransctionFromFile();    
}

// testMultisigSigned();
testMultisigUnsigned();