const algosdk = require('algosdk');

// sandbox
const token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
const server = "http://localhost";
const port = 4001;

const main = async () => {
    // Generate 3 accounts
    const account1 = algosdk.generateAccount();
    const account2 = algosdk.generateAccount();
    const account3 = algosdk.generateAccount();

    // Account addresses
    console.log("Account 1: ", account1.addr);
    console.log("Account 2: ", account2.addr);
    console.log("Account 3: ", account3.addr);
    
    // Get the mnemonic for later usage
    const mn1 = algosdk.secretKeyToMnemonic(account1.sk);
    const mn2 = algosdk.secretKeyToMnemonic(account2.sk);
    const mn3 = algosdk.secretKeyToMnemonic(account3.sk);
    
    console.log("Account 1 Mnemonic: ", mn1);
    console.log("Account 2 Mnemonic: ", mn2);
    console.log("Account 3 Mnemonic: ", mn3);
}

main().then().catch(e => {
    const error = e.body && e.body.message ? e.body.message : e;
    console.log(error);
});
