const algosdk = require('algosdk');
// In order to do this ASA tutorial, we will need to generate 3 accounts
// once created copy off the values which we will past into the TutorialASA code
// once created sucessfully, you will need to add funds to all three
// The Algorand TestNet Dispenser is located here: 
// https://bank.testnet.algorand.network/
var acct = null;

acct = algosdk.generateAccount();

account1 = acct.addr;
console.log("Account 1 = " + account1);
var account1_mnemonic = algosdk.secretKeyToMnemonic(acct.sk);
console.log("Account Mnemonic 1 = "+ account1_mnemonic);
var recoveredAccount1 = algosdk.mnemonicToSecretKey(account1_mnemonic);
var isValid = algosdk.isValidAddress(recoveredAccount1.addr);
console.log("Is this a valid address: " + isValid);
console.log("Account created. Save off Mnemonic and address");


acct = algosdk.generateAccount();

account2 = acct.addr;
console.log("Account 2 = " + account2);
var account2_mnemonic = algosdk.secretKeyToMnemonic(acct.sk);
console.log("Account Mnemonic 2 = " +account2_mnemonic);
var recoveredAccount2 = algosdk.mnemonicToSecretKey(account2_mnemonic);
var isValid = algosdk.isValidAddress(recoveredAccount2.addr);
console.log("Is this a valid address: " + isValid);
console.log("Account created. Save off Mnemonic and address");

acct = algosdk.generateAccount();

account3 = acct.addr;
console.log("Account 3 = " + account3);
var account3_mnemonic = algosdk.secretKeyToMnemonic(acct.sk);
console.log("Account Mnemonic 3 = " +account3_mnemonic);
var recoveredAccount3 = algosdk.mnemonicToSecretKey(account3_mnemonic);
var isValid = algosdk.isValidAddress(recoveredAccount3.addr);
console.log("Is this a valid address: " + isValid);
console.log("Account created. Save off Mnemonic and address");
console.log("");
console.log("Add funds to all of these accounts using the TestNet Dispenser at https://bank.testnet.algorand.network/ ");
console.log("");
console.log("Copy off these 3 lines of code and they will be pasted in the subsequent Tutorial code");
console.log("");
console.log("var account1_mnemonic = \"" + account1_mnemonic + "\"");
console.log("var account2_mnemonic = \"" + account2_mnemonic + "\"");
console.log("var account3_mnemonic = \"" + account3_mnemonic + "\"");

// sandbox
const token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
const server = "http://localhost";
const port = 4001;
// Instantiate the algod wrapper
let algodclient = new algosdk.Algod(token, server, port);

(async () => {
    let account_info = (await algodclient.accountInformation(recoveredAccount1.addr));
    let acct_string = JSON.stringify(account_info);
    console.log("Account 1 Info: " + acct_string);
    account_info = (await algodclient.accountInformation(recoveredAccount2.addr));
    acct_string = JSON.stringify(account_info);
    console.log("Account 2 Info: " + acct_string);
    account_info = (await algodclient.accountInformation(recoveredAccount3.addr));
    acct_string = JSON.stringify(account_info);
    console.log("Account 3 Info: " + acct_string);
})().catch(e => {
    console.log(e);
});

// Account 1 = RNFJFZRDOKY3ZTDXDZY7JXZF6PXRJX3Z6OKJREJCATXKAHE27PEN6S3WSI
// Account Mnemonic 1 = actor float tired slice holiday craft prefer shell enough fog girl assume edge employ piece address antenna kidney square chuckle example congress tell able ketchup
// Is this a valid address: true
// Account created.Save off Mnemonic and address

// Account 2 = NONFSLZNME4AKQMEPV5FTOKZEQPF4UB6GN5ERFZ5UGWIZB3IUBZ6MET5AI
// Account Mnemonic 2 = crumble foil love below clog way cluster first castle energy rich coin thing tribe skull sentence awful destroy main buyer cable warm welcome abstract excit
// Is this a valid address: true
// Account created.Save off Mnemonic and address

// Account 3 = SYYUGUEKECUK7ORTRH3MM2TPSG6ZCTB4ORQGUN7DKNJ7R26B36NIVMZLIY
// Account Mnemonic 3 = green inside final anchor antenna radio vintage rubber coil leaf anger insane round room moment industry basket entire lazy quiz enlist dad dilemma about program
// Is this a valid address: true
// Account created.Save off Mnemonic and address

// Add funds to all of these accounts using the TestNet Dispenser at https://bank.testnet.algorand.network/

// Copy off these 3 lines of code and they will be pasted in the subsequent Tutorial code
// var account1_mnemonic = "actor float tired slice holiday craft prefer shell enough fog girl assume edge employ piece address antenna kidney square chuckle example congress tell able ketchup"
// var account2_mnemonic = "crumble foil love below clog way cluster first castle energy rich coin thing tribe skull sentence awful destroy main buyer cable warm welcome abstract excite"
// var account3_mnemonic = "green inside final anchor antenna radio vintage rubber coil leaf anger insane round room moment industry basket entire lazy quiz enlist dad dilemma about program"

// Account 1 Info: { "round": 5983626, "address": "RNFJFZRDOKY3ZTDXDZY7JXZF6PXRJX3Z6OKJREJCATXKAHE27PEN6S3WSI", "amount": 100000000, "pendingrewards": 0, "amountwithoutpendingrewards": 100000000, "rewards": 0, "status": "Offline" }
// Account 2 Info: { "round": 5983626, "address": "NONFSLZNME4AKQMEPV5FTOKZEQPF4UB6GN5ERFZ5UGWIZB3IUBZ6MET5AI", "amount": 100000000, "pendingrewards": 0, "amountwithoutpendingrewards": 100000000, "rewards": 0, "status": "Offline" }
// Account 3 Info: { "round": 5983626, "address": "SYYUGUEKECUK7ORTRH3MM2TPSG6ZCTB4ORQGUN7DKNJ7R26B36NIVMZLIY", "amount": 100000000, "pendingrewards": 0, "amountwithoutpendingrewards": 100000000, "rewards": 0, "status": "Offline" }

