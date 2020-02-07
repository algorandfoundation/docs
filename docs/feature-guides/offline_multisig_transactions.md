title: Authorizing Multisig Transactions Offline

This section explains how to authorize multisig transactions with private keys that are kept **offline**. In particular, this guide shows how to create and save transactions to a file that can then be transferred to an offline device for signing. To learn how to construct and authorize transactions in general visit the [Constructing Transactions](./transactions.md) and [Authorizing Transactions](./signatures.md) guides, respectively.

The same methodology described here can also be used to work with [LogicSignatures](asc1/modes.md#logic-signatures) and [Single Signatures](signatures.md#single-signatures). All objects in the following examples use msgpack to store the transaction object ensuring interoperability with the SDKs and `goal`.

!!! info
    Storing keys _offline_ is also referred to as placing them in **cold storage**. An _online_ device that stores private keys is often referred to as a **hot wallet**.  


# Unsigned Multisig Transaction File Operations
Algorand SDK's and `goal` support writing and reading both signed and unsigned transactions to a file. Examples of these scenarios are shown in the following code snippets.

Unsigned transactions require the transaction object to be created before writing to a file. Also the multisig params are written out as well.


``` javascript tab="JavaScript"
    const params = {
        version: 1,
        threshold: 2,
        addrs: [
            myAccount1.addr,
            myAccount2.addr,
            myAccount3.addr,
        ],
    };

    let txparams = await algodClient.getTransactionParams();
    let msa = algosdk.multisigAddress(params);
    console.log("My multisig address is: %s", msa)

    // BREAK HERE here and add algos from the testnet dispenser to the multisig account if 0 balance
    // CPZ3OU3QDH2565PNDN7HY5UF7IIC7EUYSHDPVAZ3MDE3WKGOVWE7HQNVL4
    // send from msa to Account 3
    let txn = {
        "from": msa,
        "to": myAccount3.addr,
        "fee": txparams.minFee,
        "amount": 1000000,
        "firstRound": txparams.lastRound,
        "lastRound": txparams.lastRound + 1000,
        "genesisID": txparams.genesisID,
        "genesisHash": txparams.genesishashb64
    };
    fs.writeFileSync('./unsignedmultisig.txn', algosdk.encodeObj(txn));
    console.log("The txn file was saved!");
    fs.writeFileSync('./unsignedmultisig.params', algosdk.encodeObj(params));
    console.log("The params file was saved!");    

    txn = algosdk.decodeObj(fs.readFileSync('./unsignedmultisig.txn'));
    const params = algosdk.decodeObj(fs.readFileSync('./unsignedmultisig.params'));
    let msa = algosdk.multisigAddress(params);
    console.log("My multisig address is: %s", msa)

    let stxn = algosdk.signMultisigTransaction(txn, params, myAccount1.sk);
    let partialTxn1 = new Uint8Array(stxn.blob);
    let partialTxn2 = algosdk.appendSignMultisigTransaction(partialTxn1, params, myAccount2.sk).blob;
    let mergedTsigTxn = algosdk.mergeMultisigTransactions([partialTxn1, partialTxn2]);
    // now two of the three accounts in the multisig have signed the transaction, so the threshold of 2 has been met
    // submit the transaction     
    let signedTxn = (await algodClient.sendRawTransaction(mergedTsigTxn));
    console.log("Transaction : " + signedTxn.txId);    
```

``` python tab="Python"
accounts = {}
counter = 1
for m in [mnemonic1, mnemonic2, mnemonic3]:
    accounts[counter] = {}
    accounts[counter]['pk'] = mnemonic.to_public_key(m)
    accounts[counter]['sk'] = mnemonic.to_private_key(m)
    counter += 1

print("Account 1 address: {}".format(accounts[1]['pk']))
print("Account 2 address: {}".format(accounts[2]['pk']))
print("Account 3 address: {}".format(accounts[3]['pk']))

# create a multisig account
version = 1  # multisig version
threshold = 2  # how many signatures are necessary
multisigparams = {
    "version": version,
    "threshold": threshold,
    "addresses": [
        accounts[1]['pk'],
        accounts[2]['pk']
    ]}
msig = transaction.Multisig(**multisigparams)
print(msig.address())
input("Please go to: https://bank.testnet.algorand.network/ to fund your multisig account." +
    '\n' + "Press Enter to continue...")
# get suggested parameters
params = algod_client.suggested_params()
gen = params["genesisID"]
gh = params["genesishashb64"]
last_round = params["lastRound"]
fee = params["fee"]
# create a transaction
sender = msig.address()
amount = 1000000
txn = transaction.PaymentTxn(
    sender, fee, last_round, last_round + 1000, gh, accounts[3]['pk'], amount)
print(msig.address())
# create a MultisigTransaction object
mtx = transaction.MultisigTransaction(txn, msig)
# write to file
dir_path = os.path.dirname(os.path.realpath(__file__))
transaction.write_to_file(
    [mtx], dir_path + "/unsigned.mtx") 

# For ease of reference, add account public and private keys to
# an accounts dict.
accounts = {}
counter = 1
for m in [mnemonic1, mnemonic2, mnemonic3]:
    accounts[counter] = {}
    accounts[counter]['pk'] = mnemonic.to_public_key(m)
    accounts[counter]['sk'] = mnemonic.to_private_key(m)
    counter += 1

print("Account 1 address: {}".format(accounts[1]['pk']))
print("Account 2 address: {}".format(accounts[2]['pk']))
print("Account 3 address: {}".format(accounts[3]['pk']))
# read from file
dir_path = os.path.dirname(os.path.realpath(__file__))
msigs = transaction.retrieve_from_file(dir_path + "/unsigned.mtx")
mtx = msigs[0]
# sign the transaction
mtx.sign(accounts[1]['sk'])
mtx.sign(accounts[2]['sk'])
# print encoded transaction
print(encoding.msgpack_encode(mtx))
# send the transaction
transaction_id = algod_client.send_raw_transaction(encoding.msgpack_encode(mtx))
print("\nTransaction was sent!")
print("Transaction ID: " + transaction_id + "\n")
# wait for confirmation
wait_for_confirmation(algod_client, transaction_id)  
```

``` java tab="Java"
    List<Ed25519PublicKey> publicKeys = new ArrayList<>();
    publicKeys.add(acct1.getEd25519PublicKey());
    publicKeys.add(acct2.getEd25519PublicKey());
    publicKeys.add(acct3.getEd25519PublicKey());
    // Please go to: https://bank.testnet.algorand.network/ to fund your accounts.
    MultisigAddress msig = new MultisigAddress(1, 2, publicKeys);
    System.out.println("Multisig Address: " + msig.toString());

    // Save multisig to a file
    Files.write(Paths.get("./unsigned.msig"), Encoder.encodeToMsgPack(msig));

    MultisigAddress decodedMultisig = Encoder.decodeFromMsgPack(Files.readAllBytes(Paths.get("./unsigned.msig")), MultisigAddress.class);
    System.out.println("Multisig Address: " + decodedMultisig.toString());

    msig = decodedMultisig;

    // final String DEST_ADDR = "transaction-reciever<PLACEHOLDER>";
    // final String SRC_ADDR = "transaction-sender<PLACEHOLDER>";

    final String DEST_ADDR = "transaction-reciever<PLACEHOLDER>";
    final Address SRC_ADDR = msig.toAddress();

    // Get suggested parameters from the node
    TransactionParams params = algodApiInstance.transactionParams();
    BigInteger firstRound = params.getLastRound();
    String genId = params.getGenesisID();
    Digest genesisHash = new Digest(params.getGenesishashb64());

    // create transaction
    BigInteger lastRound = firstRound.add(BigInteger.valueOf(1000));
    BigInteger fee = BigInteger.valueOf(1000);
    fee = params.getFee();
    BigInteger amount = BigInteger.valueOf(1000000);
    Transaction tx = new Transaction(SRC_ADDR, fee, firstRound, lastRound, null, amount, new Address(DEST_ADDR),
            genId, genesisHash);

    System.out.println("Unsigned transaction with txid: " + tx.txID());

    // NOTE: save as signed even though it has not been
    SignedTransaction stx = new SignedTransaction();
    stx.tx = tx;
    // Save transaction to a file
    Files.write(Paths.get("./unsigned.txn"), Encoder.encodeToMsgPack(stx));
    System.out.println("Transaction written to a file");
    

    // read transaction from file
    SignedTransaction decodedTransaction = Encoder.decodeFromMsgPack(Files.readAllBytes(Paths.get("./unsigned.txn")), SignedTransaction.class);
    System.out.println("Signed transaction with txid: " + decodedTransaction.tx.txID());

    MultisigAddress decodedMultisig = Encoder.decodeFromMsgPack(Files.readAllBytes(Paths.get("./unsigned.msig")), MultisigAddress.class);
    System.out.println("Multisig Address: " + decodedMultisig.toString());

    msig = decodedMultisig;
    System.out.println("Multisig Address: " + msig.toString());

    Transaction tx1 = decodedTransaction.tx;

    SignedTransaction signedTransaction = acct1.signMultisigTransaction(msig, tx1);

    SignedTransaction signedTrx2 = acct2.appendMultisigTransaction(msig, signedTransaction);

    System.err.println(
            "Please go to: https://bank.testnet.algorand.network/ to fund your multisig account. Press enter when ready \n"
                    + msig.toAddress());
    System.in.read();

    try {
        // sign transaction
        byte[] encodedTxBytes = Encoder.encodeToMsgPack(signedTrx2);
        // submit the encoded transaction to the network
        TransactionID id = algodApiInstance.rawTransaction(encodedTxBytes);
        System.out.println("Successfully sent multisig: " + id);
        waitForConfirmation(id.getTxId());
    } catch (ApiException e) {
        // This may be generally expected, but should give us an informative error
        // message.
        System.err.println("Exception when calling algod#rawTransaction: " + e.getResponseBody());
    }

```

``` go tab="Go"
    // setup connection
    algodClient := setupConnection()

    // recover account for example
    sks, pks := recoverAccounts()
    PrettyPrint(sks)
    // get network suggested parameters
    txParams, err := algodClient.SuggestedParams()
    if err != nil {
        fmt.Printf("error getting suggested tx params: %s\n", err)
        return
    }
    addr1, _ := types.DecodeAddress(pks[1])
    addr2, _ := types.DecodeAddress(pks[2])
    addr3, _ := types.DecodeAddress(pks[3])
    ma, err := crypto.MultisigAccountWithParams(1, 2, []types.Address{
        addr1,
        addr2,
        addr3,
    })
    if err != nil {
        panic("invalid multisig parameters")
    }
    fromAddr, _ := ma.Address()
    fmt.Printf("Here is your multisig address : %s \n", fromAddr.String())
    fmt.Println("Please go to: https://bank.testnet.algorand.network/ to fund your multisig account.")
    // wait for Enter Key
    // this does not return if debugging in go in vscode. known issue
    // comment out this line after adding funds to the account, uncomment if loading for the fist time
    // or just set a break point here and add funds
    // fmt.Scanln()
    // send from multisig to account 3
    txn, err := transaction.MakePaymentTxn(
        fromAddr.String(),
        addr3.String(),
        txParams.Fee,           // fee per byte
        1000000,                // amount
        txParams.LastRound,     // first valid round
        txParams.LastRound+100, // last valid round
        nil,                    // note
        "",                     // closeRemainderTo
        txParams.GenesisID,     // genesisID
        txParams.GenesisHash,   // genesisHash (Cannot be empty in practice)
    )
    unsignedTx := types.SignedTxn{
        Txn: txn,
    }
    // save unsigned transction to file
    err = ioutil.WriteFile("./unsigned.txn", msgpack.Encode(unsignedTx), 0644)
    if err != nil {
        fmt.Printf("Failed in saving trx to file, error %s\n", err)
        return
    }
    fmt.Printf("Saved unsigned transaction to file\n")

    err = ioutil.WriteFile("./ma.txn", msgpack.Encode(ma), 0644)
    if err != nil {
        fmt.Printf("Failed in saving ma to file, error %s\n", err)
        return
    }

    fmt.Printf("Saved ma to file\n")

  // read unsigned transaction from file
    dat, err := ioutil.ReadFile("./unsigned.txn")
    if err != nil {
        fmt.Printf("Error reading transaction from file: %s\n", err)
        return
    }
    var unsignedTxRaw types.SignedTxn
    var unsignedTxn types.Transaction

    msgpack.Decode(dat, &unsignedTxRaw)

    unsignedTxn = unsignedTxRaw.Txn

    // recover account and sign transaction
    sks, pks := recoverAccounts()
    addr := pks[1]
    fmt.Printf("Address is: %s\n", addr)

    msgpack.Decode(dat, &unsignedTxRaw)
    unsignedTxn = unsignedTxRaw.Txn

    // read ma from file
    datma, err := ioutil.ReadFile("./ma.txn")
    if err != nil {
        fmt.Printf("Error reading ma from file: %s\n", err)
        return
    }

    var ma crypto.MultisigAccount
    msgpack.Decode(datma, &ma)

    txid, txBytes, err := crypto.SignMultisigTransaction(sks[1], ma, unsignedTxn)
    if err != nil {
        panic("could not sign multisig transaction")
    }
    fmt.Printf("Made partially-signed multisig transaction with TxID %s: %x\n", txid, txBytes)

    // append our signature to readTxBytes
    txid, twoOfThreeTxBytes, err := crypto.AppendMultisigTransaction(sks[2], ma, txBytes)
    if err != nil {
        panic("could not append signature to multisig transaction")
    }
    fmt.Printf("Made 2-out-of-3 multisig transaction with TxID %s: %x\n", txid, twoOfThreeTxBytes)

    // Send transaction to the network
    sendResponse, err := algodClient.SendRawTransaction(twoOfThreeTxBytes)
    if err != nil {
        fmt.Printf("Failed to create payment transaction: %v\n", err)
        return
    }

    fmt.Printf("Transaction ID: %s\n", sendResponse.TxID)
    waitForConfirmation(algodClient, sendResponse.TxID)    
```


``` goal tab="goal"
Create a multisig account by listing all of the accounts in the multisig and specifying the threshold number of accounts to sign with the -T flag
$ goal account multisig new my-account1<PLACEHOLDER> my-account2<PLACEHOLDER> my-account3<PLACEHOLDER> etc… -T 2    

Create an unsigned transaction and write to file
$ goal clerk send --from my-multisig-account<PLACEHOLDER>  --to AZLR2XP4O2WFHLX6TX7AZVY23HLVLG3K5K3FRIKIYDOYN6ISIF54SA4RNY --fee=1000 --amount=1000000 --out="unsigned.txn"

Sign by the required number of accounts to meet the threshold. 
$ goal clerk multisig sign -a F my-account1<PLACEHOLDER> -t=unsigned.txn
goal clerk multisig sign -a F my-account2<PLACEHOLDER> -t=unsigned.txn

Merge signings 
$ goal clerk multisig merge --out signed.txn unsigned.txn

Broadcast 
$ goal clerk rawsend --filename signed.txn
```
# Signed Multisig Transaction File Operations 
Signed Transactions are similar, but require an account to sign the transaction before writing it to a file. Also for multisig transactions, the threshold of signings must be met for the transaction to complete.

``` javascript tab="JavaScript"
    const params = {
        version: 1,
        threshold: 2,
        addrs: [
            myAccount1.addr,
            myAccount2.addr,
            myAccount3.addr,
        ],
    };
    let txparams = await algodClient.getTransactionParams();
    let msa = algosdk.multisigAddress(params);
    console.log("My multisig address is: %s", msa)
    // BREAK HERE here and add algos from the testnet dispenser to the multisig account if 0 balance
    // CPZ3OU3QDH2565PNDN7HY5UF7IIC7EUYSHDPVAZ3MDE3WKGOVWE7HQNVL4
    // send from msa to Account 3
    let txn = {
        "from": msa,
        "to": myAccount3.addr,
        "fee": txparams.minFee,
        "amount": 1000000,
        "firstRound": txparams.lastRound,
        "lastRound": txparams.lastRound + 1000,
        "genesisID": txparams.genesisID,
        "genesisHash": txparams.genesishashb64
    };
    fs.writeFileSync('./unsignedmultisig.txn', algosdk.encodeObj(txn));
    console.log("The txn file was saved!");
    fs.writeFileSync('./unsignedmultisig.params', algosdk.encodeObj(params));
    console.log("The params file was saved!");

    txn = algosdk.decodeObj(fs.readFileSync('./unsignedmultisig.txn'));
    const params = algosdk.decodeObj(fs.readFileSync('./unsignedmultisig.params'));
    let msa = algosdk.multisigAddress(params);
    console.log("My multisig address is: %s", msa)

    let stxn = algosdk.signMultisigTransaction(txn, params, myAccount1.sk);
    let partialTxn1 = new Uint8Array(stxn.blob);
    let partialTxn2 = algosdk.appendSignMultisigTransaction(partialTxn1, params, myAccount2.sk).blob;
    let mergedTsigTxn = algosdk.mergeMultisigTransactions([partialTxn1, partialTxn2]);
    // now two of the three accounts in the multisig have signed the transaction, so the threshold of 2 has been met
    // submit the transaction     
    let signedTxn = (await algodClient.sendRawTransaction(mergedTsigTxn));
    console.log("Transaction : " + signedTxn.txId);

    stxn = algosdk.decodeObj(fs.readFileSync('./signedmultisig.stxn'));
    const params = algosdk.decodeObj(fs.readFileSync('./signedmultisig.params'));

    let msa = algosdk.multisigAddress(params);
    console.log("My multisig address is: %s", msa)     
    let partialTxn1 = new Uint8Array(stxn.blob);
    let partialTxn2 = algosdk.appendSignMultisigTransaction(partialTxn1, params, myAccount2.sk).blob;
    let mergedTsigTxn = algosdk.mergeMultisigTransactions([partialTxn1, partialTxn2]);
    // now two of the three accounts in the multisig have signed the transaction, so the threshold of 2 has been met
    // submit the transaction     
    let signedTxn = (await algodClient.sendRawTransaction(mergedTsigTxn));
    console.log("Transaction : " + signedTxn.txId);
```

``` python tab="Python"

# set up connection to node
algod_client = connect_to_network()
    
# For ease of reference, add account public and private keys to
# an accounts dict.
accounts = {}
counter = 1
for m in [mnemonic1, mnemonic2, mnemonic3]:
    accounts[counter] = {}
    accounts[counter]['pk'] = mnemonic.to_public_key(m)
    accounts[counter]['sk'] = mnemonic.to_private_key(m)
    counter += 1

print("Account 1 address: {}".format(accounts[1]['pk']))
print("Account 2 address: {}".format(accounts[2]['pk']))
print("Account 3 address: {}".format(accounts[3]['pk']))

# create a multisig account
version = 1  # multisig version
threshold = 2  # how many signatures are necessary
multisigparams = {
    "version": version,
    "threshold": threshold,
    "addresses": [
        accounts[1]['pk'],
        accounts[2]['pk']
    ]}

msig = transaction.Multisig(**multisigparams)
print(msig.address())
input("Please go to: https://bank.testnet.algorand.network/ to fund your multisig account." +
    '\n' + "Press Enter to continue...")

# get suggested parameters
params = algod_client.suggested_params()
gen = params["genesisID"]
gh = params["genesishashb64"]
last_round = params["lastRound"]
fee = params["fee"]
# create a transaction
sender = msig.address()
amount = 1000000
txn = transaction.PaymentTxn(
    sender, fee, last_round, last_round + 1000, gh, accounts[3]['pk'], amount)
print(msig.address())
# create a MultisigTransaction object
mtx = transaction.MultisigTransaction(txn, msig)
# sign the transaction
mtx.sign(accounts[1]['sk'])
mtx.sign(accounts[2]['sk'])
# print encoded transaction
print(encoding.msgpack_encode(mtx))
# write to file
dir_path = os.path.dirname(os.path.realpath(__file__))
transaction.write_to_file(
    [mtx], dir_path + "/signed.mtx")
print("Signed mtx file saved!")   
# read from file
dir_path = os.path.dirname(os.path.realpath(__file__))
msigs = transaction.retrieve_from_file(dir_path + "/signed.mtx")
mtx = msigs[0]
# send the transaction
transaction_id = algod_client.send_raw_transaction(encoding.msgpack_encode(mtx))
print("\nTransaction was sent!")
print("Transaction ID: " + transaction_id + "\n")
wait_for_confirmation(algod_client, transaction_id)

```

``` java tab="Java"
    List<Ed25519PublicKey> publicKeys = new ArrayList<>();
    publicKeys.add(acct1.getEd25519PublicKey());
    publicKeys.add(acct2.getEd25519PublicKey());
    publicKeys.add(acct3.getEd25519PublicKey());
    // Please go to: https://bank.testnet.algorand.network/ to fund your accounts.
    MultisigAddress msig = new MultisigAddress(1, 2, publicKeys);
    System.out.println("Multisig Address: " + msig.toString());

    // Save multisig to a file
    Files.write(Paths.get("./unsigned.msig"), Encoder.encodeToMsgPack(msig));

    MultisigAddress decodedMultisig = Encoder.decodeFromMsgPack(Files.readAllBytes(Paths.get("./unsigned.msig")), MultisigAddress.class);
    System.out.println("Multisig Address: " + decodedMultisig.toString());

    msig = decodedMultisig;

    // final String DEST_ADDR = "transaction-reciever<PLACEHOLDER>";
    // final String SRC_ADDR = "transaction-sender<PLACEHOLDER>";

    final String DEST_ADDR = "transaction-reciever<PLACEHOLDER>";
    final Address SRC_ADDR = msig.toAddress();

    // Get suggested parameters from the node
    TransactionParams params = algodApiInstance.transactionParams();
    BigInteger firstRound = params.getLastRound();
    String genId = params.getGenesisID();
    Digest genesisHash = new Digest(params.getGenesishashb64());

    // create transaction
    BigInteger lastRound = firstRound.add(BigInteger.valueOf(1000));
    BigInteger fee = BigInteger.valueOf(1000);
    fee = params.getFee();
    BigInteger amount = BigInteger.valueOf(1000000);
    Transaction tx = new Transaction(SRC_ADDR, fee, firstRound, lastRound, null, amount, new Address(DEST_ADDR),
            genId, genesisHash);

    System.out.println("Unsigned transaction with txid: " + tx.txID());

    SignedTransaction signedTransaction = acct1.signMultisigTransaction(msig, tx);

    SignedTransaction signedTrx2 = acct2.appendMultisigTransaction(msig, signedTransaction);

    // save signed transaction to a file
    Files.write(Paths.get("./signed.txn"), Encoder.encodeToMsgPack(signedTrx2));

    // Read the transaction from a file
    SignedTransaction decodedSignedTransaction = Encoder
            .decodeFromMsgPack(Files.readAllBytes(Paths.get("./signed.txn")), SignedTransaction.class);
    System.out.println("Signed transaction with txid: " + decodedSignedTransaction.transactionID);

    // Msgpack encode the signed transaction
    byte[] encodedTxBytes = Encoder.encodeToMsgPack(decodedSignedTransaction);

    // submit the encoded transaction to the network
    TransactionID id = algodApiInstance.rawTransaction(encodedTxBytes);
    System.out.println("Successfully sent tx with id: " + id);
    waitForConfirmation(id.getTxId());         
```

``` go tab="Go"
  // setup connection
    algodClient := setupConnection()

    // recover account for example
    //addr, _ := recoverAccounts();
    sks, pks := recoverAccounts()
    PrettyPrint(sks)
    // get network suggested parameters
    txParams, err := algodClient.SuggestedParams()
    if err != nil {
        fmt.Printf("error getting suggested tx params: %s\n", err)
        return
    }
    addr1, _ := types.DecodeAddress(pks[1])
    addr2, _ := types.DecodeAddress(pks[2])
    addr3, _ := types.DecodeAddress(pks[3])
    ma, err := crypto.MultisigAccountWithParams(1, 2, []types.Address{
        addr1,
        addr2,
        addr3,
    })
    if err != nil {
        panic("invalid multisig parameters")
    }
    fromAddr, _ := ma.Address()
    fmt.Printf("Here is your multisig address : %s \n", fromAddr.String())
    fmt.Println("Please go to: https://bank.testnet.algorand.network/ to fund your multisig account.")
    // wait for Enter Key
    // this does not return if debugging in go in vscode. known issue
    // comment out this line after adding funds to the account, uncomment if loading for the fist time
    // or just set a break point here and add funds
    // fmt.Scanln()
    // send from multisig to account 3
    txn, err := transaction.MakePaymentTxn(
        fromAddr.String(),
        addr3.String(),
        txParams.Fee,           // fee per byte
        1000000,                // amount
        txParams.LastRound,     // first valid round
        txParams.LastRound+100, // last valid round
        nil,                    // note
        "",                     // closeRemainderTo
        txParams.GenesisID,     // genesisID
        txParams.GenesisHash,   // genesisHash (Cannot be empty in practice)
    )

    txid, txBytes, err := crypto.SignMultisigTransaction(sks[1], ma, txn)
    if err != nil {
        panic("could not sign multisig transaction")
    }
    fmt.Printf("Made partially-signed multisig transaction with TxID %s: %x\n", txid, txBytes)

    // append our signature to readTxBytes
    txid, twoOfThreeTxBytes, err := crypto.AppendMultisigTransaction(sks[2], ma, txBytes)
    if err != nil {
        panic("could not append signature to multisig transaction")
    }
    fmt.Printf("Made 2-out-of-3 multisig transaction with TxID %s: %x\n", txid, twoOfThreeTxBytes)

    // otherTxBytes := ... // generate another raw multisig transaction somehow
    // txid, mergedTxBytes, err := crypto.MergeMultisigTransactions(twoOfThreeTxBytes, otherTxBytes)

    //Save the signed transaction to file
    err = ioutil.WriteFile("./signed.stxn", twoOfThreeTxBytes, 0644)
    if err != nil {

        fmt.Printf("Failed in saving stxn to file, error %s\n", err)
        return
    }
    fmt.Printf("Saved signed transaction to file\n")
    // read unsigned transaction from file
    dat, err := ioutil.ReadFile("./signed.stxn")
    if err != nil {
        fmt.Printf("Error reading signed transaction from file: %s\n", err)
        return
    }

    // send the transaction to the network
    sendResponse, err := algodClient.SendRawTransaction(dat)
    if err != nil {
        fmt.Printf("failed to send transaction: %s\n", err)
        return
    }
    fmt.Printf("Transaction ID: %s\n", sendResponse.TxID)
    waitForConfirmation(algodClient, sendResponse.TxID)    
```

``` goal tab="goal"
$ goal clerk rawsend --filename signed.txn
```

??? example "Complete Example = Saving Signed and Unsigned Multisig Transactions to a File"
    
    ```javascript tab="JavaScript"
    const algosdk = require('algosdk');
    const fs = require('fs');
    var client = null;
    // make connection to node
    async function setupClient() {
        if (client == null) {
            const ALGOD_API_ADDR = "algod-address<PLACEHOLDER>";
            const ALGOD_API_TOKEN = "algod-token<PLACEHOLDER>";
            const port = port-number<PLACEHOLDER>;
            let algodClient = new algosdk.Algod(token, server, port);
            client = algodClient;
        } else {
            return client;
        }
        return client;
    }
    // recover acccounts for example

    function recoverAccount1() {
        const passphrase = "your-25-word-mnemonic<PLACEHOLDER>";
        let myAccount = algosdk.mnemonicToSecretKey(passphrase);   
        return myAccount;
    }
    function recoverAccount2() {
        const passphrase = "your-25-word-mnemonic<PLACEHOLDER>";
        let myAccount = algosdk.mnemonicToSecretKey(passphrase);
        return myAccount;
    }
    function recoverAccount3() {
        const passphrase = "your-25-word-mnemonic<PLACEHOLDER>";
        let myAccount = algosdk.mnemonicToSecretKey(passphrase);
        return myAccount;
    }

    // Function used to wait for a tx confirmation
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

    async function writeSignedMultisigTransctionToFile() {

        try {
            let algodClient = await setupClient();
            let myAccount1 = await recoverAccount1();
            console.log("My address1: %s", myAccount1.addr)
            let myAccount2 = await recoverAccount2();
            console.log("My address2: %s", myAccount2.addr)        
            let myAccount3 = await recoverAccount3();
            console.log("My address3: %s", myAccount3.addr)   
            const params = {
                version: 1,
                threshold: 2,
                addrs: [
                    myAccount1.addr,
                    myAccount2.addr,
                    myAccount3.addr,
                ],
            };  
            let txparams = await algodClient.getTransactionParams();
            let msa = algosdk.multisigAddress(params);
            console.log("My multisig address is: %s", msa)  

            // BREAK HERE here and add algos from the testnet dispenser to the multisig account if 0 balance
            // CPZ3OU3QDH2565PNDN7HY5UF7IIC7EUYSHDPVAZ3MDE3WKGOVWE7HQNVL4
            // send from msa to Account 3
            let txn = {
                "from" : msa,
                "to": myAccount3.addr,
                "fee": txparams.minFee,
                "amount": 1000000,
                "firstRound": txparams.lastRound,
                "lastRound": txparams.lastRound + 1000,
                "genesisID": txparams.genesisID,
                "genesisHash": txparams.genesishashb64
            };
            let stxn = algosdk.signMultisigTransaction(txn, params, myAccount1.sk);
            fs.writeFileSync('./signedmultisig.stxn', algosdk.encodeObj(stxn));
            console.log("The stxn file was saved!");  
            fs.writeFileSync('./signedmultisig.params', algosdk.encodeObj(params));
            console.log("The params file was saved!");       
        } catch (e) {
            console.log(e);
        }
    };
    async function readSignedMultisigTransctionFromFile() {
        try {
            let algodClient = await setupClient();
            let myAccount1 = await recoverAccount1();
            console.log("My address1: %s", myAccount1.addr)
            let myAccount2 = await recoverAccount2();
            console.log("My address2: %s", myAccount2.addr)
            let myAccount3 = await recoverAccount3();
            console.log("My address3: %s", myAccount3.addr)
            stxn = algosdk.decodeObj(fs.readFileSync('./signedmultisig.stxn'));
            const params = algosdk.decodeObj(fs.readFileSync('./signedmultisig.params'));
            let msa = algosdk.multisigAddress(params);
            console.log("My multisig address is: %s", msa)     
            let partialTxn1 = new Uint8Array(stxn.blob);
            let partialTxn2 = algosdk.appendSignMultisigTransaction(partialTxn1, params, myAccount2.sk).blob;
            let mergedTsigTxn = algosdk.mergeMultisigTransactions([partialTxn1, partialTxn2]);
            // now two of the three accounts in the multisig have signed the transaction, so the threshold of 2 has been met
            // submit the transaction     
            let signedTxn = (await algodClient.sendRawTransaction(mergedTsigTxn));
            console.log("Transaction : " + signedTxn.txId);
        } catch (e) {
            console.log(e);
        }
    };

    async function writeUnsignedMultisigTransctionToFile() {
        try {
            let algodClient = await setupClient();
            let myAccount1 = await recoverAccount1();
            console.log("My address1: %s", myAccount1.addr)
            let myAccount2 = await recoverAccount2();
            console.log("My address2: %s", myAccount2.addr)
            let myAccount3 = await recoverAccount3();
            console.log("My address3: %s", myAccount3.addr)
            const params = {
                version: 1,
                threshold: 2,
                addrs: [
                    myAccount1.addr,
                    myAccount2.addr,
                    myAccount3.addr,
                ],
            };
            let txparams = await algodClient.getTransactionParams();
            let msa = algosdk.multisigAddress(params);
            console.log("My multisig address is: %s", msa)
            // BREAK HERE here and add algos from the testnet dispenser to the multisig account if 0 balance
            // CPZ3OU3QDH2565PNDN7HY5UF7IIC7EUYSHDPVAZ3MDE3WKGOVWE7HQNVL4
            // send from msa to Account 3
            let txn = {
                "from": msa,
                "to": myAccount3.addr,
                "fee": txparams.minFee,
                "amount": 1000000,
                "firstRound": txparams.lastRound,
                "lastRound": txparams.lastRound + 1000,
                "genesisID": txparams.genesisID,
                "genesisHash": txparams.genesishashb64
            };
            fs.writeFileSync('./unsignedmultisig.txn', algosdk.encodeObj(txn));
            console.log("The txn file was saved!");
            fs.writeFileSync('./unsignedmultisig.params', algosdk.encodeObj(params));
            console.log("The params file was saved!");
        } catch (e) {
            console.log(e);
        }
    };
    async function readUnsignedMultisigTransctionFromFile() {
        try {
            let algodClient = await setupClient();
            let myAccount1 = await recoverAccount1();
            console.log("My address1: %s", myAccount1.addr)
            let myAccount2 = await recoverAccount2();
            console.log("My address2: %s", myAccount2.addr)
            let myAccount3 = await recoverAccount3();
            console.log("My address3: %s", myAccount3.addr) 
            // BREAK HERE here and add algos from the testnet dispenser to the multisig account if 0 balance
            // CPZ3OU3QDH2565PNDN7HY5UF7IIC7EUYSHDPVAZ3MDE3WKGOVWE7HQNVL4
            // send from msa to Account 3
            txn = algosdk.decodeObj(fs.readFileSync('./unsignedmultisig.txn'));
            const params = algosdk.decodeObj(fs.readFileSync('./unsignedmultisig.params'));
            let msa = algosdk.multisigAddress(params);
            console.log("My multisig address is: %s", msa)
            let stxn = algosdk.signMultisigTransaction(txn, params, myAccount1.sk);
            let partialTxn1 = new Uint8Array(stxn.blob);
            let partialTxn2 = algosdk.appendSignMultisigTransaction(partialTxn1, params, myAccount2.sk).blob;
            let mergedTsigTxn = algosdk.mergeMultisigTransactions([partialTxn1, partialTxn2]);
            // now two of the three accounts in the multisig have signed the transaction, so the threshold of 2 has been met
            // submit the transaction     
            let signedTxn = (await algodClient.sendRawTransaction(mergedTsigTxn));
            console.log("Transaction : " + signedTxn.txId);
        } catch (e) {
            console.log(e);
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

    testMultisigSigned();
    testMultisigUnsigned();
    ```

    ```python tab="Python"
    import params
    import os
    import json

    from algosdk import account, transaction, algod, encoding, mnemonic

    def connect_to_network():
        algod_address = "algod-address<PLACEHOLDER>"
        algod_token = "algod-token<PLACEHOLDER>"
        algod_client = algod.AlgodClient(algod_token, algod_address)
        return algod_client

    def wait_for_confirmation(algod_client, txid):
        while True:
            txinfo = algod_client.pending_transaction_info(txid)
            if txinfo.get('round') and txinfo.get('round') > 0:
                print("Transaction {} confirmed in round {}.".format(
                    txid, txinfo.get('round')))
                break
            else:
                print("Waiting for confirmation...")
                algod_client.status_after_block(
                    algod_client.status().get('lastRound') + 1)

    def write_unsigned():
        # setup none connection
        algod_client = connect_to_network()
        # generate three new accounts if desired
        # private_key_1, account_1 = account.generate_account()
        # private_key_2, account_2 = account.generate_account()
        # private_key_3, account_3 = account.generate_account()
        # print("Account 1:", account_1)
        # mnemonic1 = mnemonic.from_private_key(private_key_1)
        # print("Mnemonic 1 = ", mnemonic1)
        # print("Account 2:", account_2)
        # mnemonic2 = mnemonic.from_private_key(private_key_2)
        # print("Mnemonic 2 = ", mnemonic2)
        # print("Account 3:", account_3)
        # mnemonic3 = mnemonic.from_private_key(private_key_3)
        # print("Mnemonic 3 = ", mnemonic3)
        # input("Please go to: https://bank.testnet.algorand.network/ to fund your new accounts." +
        #       '\n' + "Press Enter to continue...")

        # or use existing accounts
        mnemonic1 = "25-word-passphrase<PLACEHOLDER>"
        mnemonic2 = "25-word-passphrase<PLACEHOLDER>"
        mnemonic3 = "25-word-passphrase<PLACEHOLDER>"
        # For ease of reference, add account public and private keys to
        # an accounts dict.
        accounts = {}
        counter = 1
        for m in [mnemonic1, mnemonic2, mnemonic3]:
            accounts[counter] = {}
            accounts[counter]['pk'] = mnemonic.to_public_key(m)
            accounts[counter]['sk'] = mnemonic.to_private_key(m)
            counter += 1

        print("Account 1 address: {}".format(accounts[1]['pk']))
        print("Account 2 address: {}".format(accounts[2]['pk']))
        print("Account 3 address: {}".format(accounts[3]['pk']))

        # create a multisig account
        version = 1  # multisig version
        threshold = 2  # how many signatures are necessary
        multisigparams = {
            "version": version,
            "threshold": threshold,
            "addresses": [
                accounts[1]['pk'],
                accounts[2]['pk']
            ]}
        msig = transaction.Multisig(**multisigparams)
        print(msig.address())
        input("Please go to: https://bank.testnet.algorand.network/ to fund your multisig account." +
            '\n' + "Press Enter to continue...")
        # get suggested parameters
        params = algod_client.suggested_params()
        gen = params["genesisID"]
        gh = params["genesishashb64"]
        last_round = params["lastRound"]
        fee = params["fee"]
        # create a transaction
        sender = msig.address()
        amount = 1000000
        txn = transaction.PaymentTxn(
            sender, fee, last_round, last_round + 1000, gh, accounts[3]['pk'], amount)
        print(msig.address())
        # create a MultisigTransaction object
        mtx = transaction.MultisigTransaction(txn, msig)
        # write to file
        dir_path = os.path.dirname(os.path.realpath(__file__))
        transaction.write_to_file(
            [mtx], dir_path + "/unsigned.mtx")

    def read_unsigned():
        # setup node connection
        algod_client = connect_to_network()
        # generate three new accounts if desired
        # private_key_1, account_1 = account.generate_account()
        # private_key_2, account_2 = account.generate_account()
        # private_key_3, account_3 = account.generate_account()
        # print("Account 1:", account_1)
        # mnemonic1 = mnemonic.from_private_key(private_key_1)
        # print("Mnemonic 1 = ", mnemonic1)
        # print("Account 2:", account_2)
        # mnemonic2 = mnemonic.from_private_key(private_key_2)
        # print("Mnemonic 2 = ", mnemonic2)
        # print("Account 3:", account_3)
        # mnemonic3 = mnemonic.from_private_key(private_key_3)
        # print("Mnemonic 3 = ", mnemonic3)
        # input("Please go to: https://bank.testnet.algorand.network/ to fund your new accounts." +
        #       '\n' + "Press Enter to continue...")

        # or use existing accounts
        mnemonic1 = "25-word-passphrase<PLACEHOLDER>"
        mnemonic2 = "25-word-passphrase<PLACEHOLDER>"
        mnemonic3 = "25-word-passphrase<PLACEHOLDER>"

        # For ease of reference, add account public and private keys to
        # an accounts dict.
        accounts = {}
        counter = 1
        for m in [mnemonic1, mnemonic2, mnemonic3]:
            accounts[counter] = {}
            accounts[counter]['pk'] = mnemonic.to_public_key(m)
            accounts[counter]['sk'] = mnemonic.to_private_key(m)
            counter += 1

        print("Account 1 address: {}".format(accounts[1]['pk']))
        print("Account 2 address: {}".format(accounts[2]['pk']))
        print("Account 3 address: {}".format(accounts[3]['pk']))
        # read from file
        dir_path = os.path.dirname(os.path.realpath(__file__))
        msigs = transaction.retrieve_from_file(dir_path + "/unsigned.mtx")
        mtx = msigs[0]
        # sign the transaction
        mtx.sign(accounts[1]['sk'])
        mtx.sign(accounts[2]['sk'])
        # print encoded transaction
        print(encoding.msgpack_encode(mtx))
        # send the transaction
        transaction_id = algod_client.send_raw_transaction(encoding.msgpack_encode(mtx))
        print("\nTransaction was sent!")
        print("Transaction ID: " + transaction_id + "\n")
        # wait for confirmation
        wait_for_confirmation(algod_client, transaction_id)

    def write_signed():
        # setup connection to node
        algod_client = connect_to_network()
        # generate three new accounts if desired
        # private_key_1, account_1 = account.generate_account()
        # private_key_2, account_2 = account.generate_account()
        # private_key_3, account_3 = account.generate_account()
        # print("Account 1:", account_1)
        # mnemonic1 = mnemonic.from_private_key(private_key_1)
        # print("Mnemonic 1 = ", mnemonic1)
        # print("Account 2:", account_2)
        # mnemonic2 = mnemonic.from_private_key(private_key_2)
        # print("Mnemonic 2 = ", mnemonic2)
        # print("Account 3:", account_3)
        # mnemonic3 = mnemonic.from_private_key(private_key_3)
        # print("Mnemonic 3 = ", mnemonic3)
        # input("Please go to: https://bank.testnet.algorand.network/ to fund your new accounts." +
        #       '\n' + "Press Enter to continue...")

        # or use existing accounts
        mnemonic1 = "25-word-passphrase<PLACEHOLDER>"
        mnemonic2 = "25-word-passphrase<PLACEHOLDER>"
        mnemonic3 = "25-word-passphrase<PLACEHOLDER>"

        # For ease of reference, add account public and private keys to
        # an accounts dict.
        accounts = {}
        counter = 1
        for m in [mnemonic1, mnemonic2, mnemonic3]:
            accounts[counter] = {}
            accounts[counter]['pk'] = mnemonic.to_public_key(m)
            accounts[counter]['sk'] = mnemonic.to_private_key(m)
            counter += 1

        print("Account 1 address: {}".format(accounts[1]['pk']))
        print("Account 2 address: {}".format(accounts[2]['pk']))
        print("Account 3 address: {}".format(accounts[3]['pk']))

        # create a multisig account
        version = 1  # multisig version
        threshold = 2  # how many signatures are necessary
        multisigparams = {
            "version": version,
            "threshold": threshold,
            "addresses": [
                accounts[1]['pk'],
                accounts[2]['pk']
            ]}

        msig = transaction.Multisig(**multisigparams)
        print(msig.address())
        input("Please go to: https://bank.testnet.algorand.network/ to fund your multisig account." +
            '\n' + "Press Enter to continue...")

        # get suggested parameters
        params = algod_client.suggested_params()
        gen = params["genesisID"]
        gh = params["genesishashb64"]
        last_round = params["lastRound"]
        fee = params["fee"]
        # create a transaction
        sender = msig.address()
        amount = 1000000
        txn = transaction.PaymentTxn(
            sender, fee, last_round, last_round + 1000, gh, accounts[3]['pk'], amount)
        print(msig.address())
        # create a MultisigTransaction object
        mtx = transaction.MultisigTransaction(txn, msig)
        # sign the transaction
        mtx.sign(accounts[1]['sk'])
        mtx.sign(accounts[2]['sk'])
        # print encoded transaction
        print(encoding.msgpack_encode(mtx))
        # write to file
        dir_path = os.path.dirname(os.path.realpath(__file__))
        transaction.write_to_file(
            [mtx], dir_path + "/signed.mtx")
        print("Signed mtx file saved!")

    def read_signed():
        # set up connection to node
        algod_client = connect_to_network()
        # read from file
        dir_path = os.path.dirname(os.path.realpath(__file__))
        msigs = transaction.retrieve_from_file(dir_path + "/signed.mtx")
        mtx = msigs[0]
        # send the transaction
        transaction_id = algod_client.send_raw_transaction(encoding.msgpack_encode(mtx))
        print("\nTransaction was sent!")
        print("Transaction ID: " + transaction_id + "\n")
        wait_for_confirmation(algod_client, transaction_id)

    # Test Runs
    #write_unsigned()
    #read_unsigned()
    write_signed()
    read_signed()
    ```

    ```java tab="Java"
    package com.algorand.javatest;

    import com.algorand.algosdk.algod.client.AlgodClient;
    import com.algorand.algosdk.algod.client.ApiException;
    import com.algorand.algosdk.algod.client.api.AlgodApi;
    import com.algorand.algosdk.algod.client.auth.ApiKeyAuth;
    import com.algorand.algosdk.account.Account;
    import com.algorand.algosdk.crypto.Address;
    import com.algorand.algosdk.algod.client.model.TransactionParams;
    import com.algorand.algosdk.crypto.Digest;
    import com.algorand.algosdk.crypto.Ed25519PublicKey;
    import com.algorand.algosdk.algod.client.model.TransactionID;
    import com.algorand.algosdk.transaction.Transaction;
    import com.algorand.algosdk.util.Encoder;

    import com.algorand.algosdk.crypto.MultisigAddress;
    import com.algorand.algosdk.transaction.SignedTransaction;

    import java.math.BigInteger;
    import java.util.ArrayList;
    import java.util.List;
    import java.nio.file.Files;
    import java.nio.file.Paths;

    public class SaveMultisigTransactionOffline {
        public AlgodApi algodApiInstance = null;

        // utility function to connect to a node
        private AlgodApi connectToNetwork() {

            // Initialize an algod client
            final String ALGOD_API_ADDR = "http://hackathon.algodev.network:9100";
            final String ALGOD_API_TOKEN = "ef920e2e7e002953f4b29a8af720efe8e4ecc75ff102b165e0472834b25832c1";

            // final String ALGOD_API_ADDR = "algod-address<PLACEHOLDER>";
            // final String ALGOD_API_TOKEN = "algod-token<PLACEHOLDER>";

            AlgodClient client = (AlgodClient) new AlgodClient().setBasePath(ALGOD_API_ADDR);
            ApiKeyAuth api_key = (ApiKeyAuth) client.getAuthentication("api_key");
            api_key.setApiKey(ALGOD_API_TOKEN);
            algodApiInstance = new AlgodApi(client);
            return algodApiInstance;
        }

        // utility function to wait on a transaction to be confirmed
        public void waitForConfirmation(String txID) throws Exception {
            if (algodApiInstance == null)
                connectToNetwork();
            while (true) {
                try {
                    // Check the pending tranactions
                    com.algorand.algosdk.algod.client.model.Transaction pendingInfo = algodApiInstance
                            .pendingTransactionInformation(txID);
                    if (pendingInfo.getRound() != null && pendingInfo.getRound().longValue() > 0) {
                        // Got the completed Transaction
                        System.out.println("Transaction " + pendingInfo.getTx() + " confirmed in round "
                                + pendingInfo.getRound().longValue());
                        break;
                    }
                    algodApiInstance
                            .waitForBlock(BigInteger.valueOf(algodApiInstance.getStatus().getLastRound().longValue() + 1));
                } catch (Exception e) {
                    throw (e);
                }
            }

        }

        public void writeMultisigUnsignedTransaction() {

            // connect to node
            if (algodApiInstance == null)
                connectToNetwork();
            try {
                // create new accounts or...
                // Account acct1 = new Account();
                // Account acct2 = new Account();
                // Account acct3 = new Account();
                // System.out.println("Account 1 Address: " + acct1.getAddress());
                // System.out.println("Account 2 Address: " + acct2.getAddress());
                // System.out.println("Account 3 Address: " + acct3.getAddress());
                // System.out.println("Account 1 Mnemonic: " + acct1.toMnemonic());
                // System.out.println("Account 2 Mnemonic: " + acct2.toMnemonic());
                // System.out.println("Account 3 Mnemonic: " + acct3.toMnemonic());
                // Please go to: https://bank.testnet.algorand.network/ to fund your accounts.

                // ... or recover accounts
                final String account1_mnemonic = "your-25-word-mnemonic<PLACEHOLDER>";
                final String account2_mnemonic = "your-25-word-mnemonic<PLACEHOLDER>";
                final String account3_mnemonic = "your-25-word-mnemonic<PLACEHOLDER>";

                final Account acct1 = new Account(account1_mnemonic);
                final Account acct2 = new Account(account2_mnemonic);
                final Account acct3 = new Account(account3_mnemonic);

                List<Ed25519PublicKey> publicKeys = new ArrayList<>();
                publicKeys.add(acct1.getEd25519PublicKey());
                publicKeys.add(acct2.getEd25519PublicKey());
                publicKeys.add(acct3.getEd25519PublicKey());
                // Please go to: https://bank.testnet.algorand.network/ to fund your accounts.
                MultisigAddress msig = new MultisigAddress(1, 2, publicKeys);
                System.out.println("Multisig Address: " + msig.toString());

                // Save multisig to a file
                Files.write(Paths.get("./unsigned.msig"), Encoder.encodeToMsgPack(msig));

                MultisigAddress decodedMultisig = Encoder.decodeFromMsgPack(Files.readAllBytes(Paths.get("./unsigned.msig")), MultisigAddress.class);
                System.out.println("Multisig Address: " + decodedMultisig.toString());

                msig = decodedMultisig;

                // final String DEST_ADDR = "transaction-reciever<PLACEHOLDER>";
                // final String SRC_ADDR = "transaction-sender<PLACEHOLDER>";

                final String DEST_ADDR "transaction-reciever<PLACEHOLDER>";
                final Address SRC_ADDR = msig.toAddress();

                // Get suggested parameters from the node
                TransactionParams params = algodApiInstance.transactionParams();
                BigInteger firstRound = params.getLastRound();
                String genId = params.getGenesisID();
                Digest genesisHash = new Digest(params.getGenesishashb64());

                // create transaction
                BigInteger lastRound = firstRound.add(BigInteger.valueOf(1000));
                BigInteger fee = BigInteger.valueOf(1000);
                fee = params.getFee();
                BigInteger amount = BigInteger.valueOf(1000000);
                Transaction tx = new Transaction(SRC_ADDR, fee, firstRound, lastRound, null, amount, new Address(DEST_ADDR),
                        genId, genesisHash);

                System.out.println("Unsigned transaction with txid: " + tx.txID());

                // NOTE: save as signed even though it has not been
                SignedTransaction stx = new SignedTransaction();
                stx.tx = tx;
                // Save transaction to a file
                Files.write(Paths.get("./unsigned.txn"), Encoder.encodeToMsgPack(stx));
                System.out.println("Transaction written to a file");

            } catch (Exception e) {
                System.out.println("Save Exception: " + e);
            }

        }

        public void readMultisigUnsignedTransaction() {
            try {
                // recover accounts
                final String account1_mnemonic = "your-25-word-mnemonic<PLACEHOLDER>";
                final String account2_mnemonic = "your-25-word-mnemonic<PLACEHOLDER>";
                final String account3_mnemonic = "your-25-word-mnemonic<PLACEHOLDER>";

                final Account acct1 = new Account(account1_mnemonic);
                final Account acct2 = new Account(account2_mnemonic);
                final Account acct3 = new Account(account3_mnemonic);

                // read transaction from file

                SignedTransaction decodedTransaction = Encoder
                        .decodeFromMsgPack(Files.readAllBytes(Paths.get("./unsigned.txn")), SignedTransaction.class);
                System.out.println("Signed transaction with txid: " + decodedTransaction.tx.txID());

                MultisigAddress decodedMultisig = Encoder
                .decodeFromMsgPack(Files.readAllBytes(Paths.get("./unsigned.msig")),
                MultisigAddress.class);
                System.out.println("Multisig Address: " + decodedMultisig.toString());

                msig = decodedMultisig;

                System.out.println("Multisig Address: " + msig.toString());

                Transaction tx1 = decodedTransaction.tx;

                SignedTransaction signedTransaction = acct1.signMultisigTransaction(msig, tx1);

                SignedTransaction signedTrx2 = acct2.appendMultisigTransaction(msig, signedTransaction);

                // save signed transaction to a file
                Files.write(Paths.get("./signed.txn"), Encoder.encodeToMsgPack(signedTrx2));

                // Read the transaction from a file
                signedTrx2 = Encoder.decodeFromMsgPack(Files.readAllBytes(Paths.get("./signed.txn")),
                        SignedTransaction.class);

                System.err.println(
                        "Please go to: https://bank.testnet.algorand.network/ to fund your multisig account. Press enter when ready \n"
                                + msig.toAddress());
                // System.in.read();

                try {
                    // sign transaction
                    byte[] encodedTxBytes = Encoder.encodeToMsgPack(signedTrx2);
                    // submit the encoded transaction to the network
                    TransactionID id = algodApiInstance.rawTransaction(encodedTxBytes);
                    System.out.println("Successfully sent multisig: " + id);
                    waitForConfirmation(id.getTxId());
                } catch (ApiException e) {
                    // This may be generally expected, but should give us an informative error
                    // message.
                    System.err.println("Exception when calling algod#rawTransaction: " + e.getResponseBody());
                }
            } catch (Exception e) {
                System.out.println("Submit Exception: " + e);
            }
        }

        public void writeMultisigSignedTransaction() {

            // connect to node
            if (algodApiInstance == null)
                connectToNetwork();
            try {

                // Account acct1 = new Account();
                // Account acct2 = new Account();
                // Account acct3 = new Account();
                // System.out.println("Account 1 Address: " + acct1.getAddress());
                // System.out.println("Account 2 Address: " + acct2.getAddress());
                // System.out.println("Account 3 Address: " + acct3.getAddress());
                // System.out.println("Account 1 Mnemonic: " + acct1.toMnemonic());
                // System.out.println("Account 2 Mnemonic: " + acct2.toMnemonic());
                // System.out.println("Account 3 Mnemonic: " + acct3.toMnemonic());
                // Please go to: https://bank.testnet.algorand.network/ to fund your accounts.

                final String account1_mnemonic = "your-25-word-mnemonic<PLACEHOLDER>";
                final String account2_mnemonic = "your-25-word-mnemonic<PLACEHOLDER>";
                final String account3_mnemonic = "your-25-word-mnemonic<PLACEHOLDER>";

                final Account acct1 = new Account(account1_mnemonic);
                final Account acct2 = new Account(account2_mnemonic);
                final Account acct3 = new Account(account3_mnemonic);

                List<Ed25519PublicKey> publicKeys = new ArrayList<>();
                publicKeys.add(acct1.getEd25519PublicKey());
                publicKeys.add(acct2.getEd25519PublicKey());
                publicKeys.add(acct3.getEd25519PublicKey());
                // Please go to: https://bank.testnet.algorand.network/ to fund your accounts.
                MultisigAddress msig = new MultisigAddress(1, 2, publicKeys);
                System.out.println("Multisig Address: " + msig.toString());

                // Save multisig to a file
                Files.write(Paths.get("./unsigned.msig"), Encoder.encodeToMsgPack(msig));

                MultisigAddress decodedMultisig = Encoder.decodeFromMsgPack(Files.readAllBytes(Paths.get("./unsigned.msig")), MultisigAddress.class);
                System.out.println("Multisig Address: " + decodedMultisig.toString());

                msig = decodedMultisig;

                // final String DEST_ADDR = 
                final String DEST_ADDR = "transaction-reciever<PLACEHOLDER>";
                final Address SRC_ADDR = msig.toAddress();

                // Get suggested parameters from the node
                TransactionParams params = algodApiInstance.transactionParams();
                BigInteger firstRound = params.getLastRound();
                String genId = params.getGenesisID();
                Digest genesisHash = new Digest(params.getGenesishashb64());

                // create transaction
                BigInteger lastRound = firstRound.add(BigInteger.valueOf(1000));
                BigInteger fee = BigInteger.valueOf(1000);
                fee = params.getFee();
                BigInteger amount = BigInteger.valueOf(1000000);
                Transaction tx = new Transaction(SRC_ADDR, fee, firstRound, lastRound, null, amount, new Address(DEST_ADDR),
                        genId, genesisHash);

                System.out.println("Unsigned transaction with txid: " + tx.txID());

                SignedTransaction signedTransaction = acct1.signMultisigTransaction(msig, tx);

                SignedTransaction signedTrx2 = acct2.appendMultisigTransaction(msig, signedTransaction);

                // save signed transaction to a file
                Files.write(Paths.get("./signed.txn"), Encoder.encodeToMsgPack(signedTrx2));

            } catch (Exception e) {
                System.out.println("Save Exception: " + e);
            }

        }

        public void readMultisigSignedTransaction() {

            try {
                // connect to a node
                if (algodApiInstance == null)
                    connectToNetwork();

                // Read the transaction from a file
                SignedTransaction decodedSignedTransaction = Encoder
                        .decodeFromMsgPack(Files.readAllBytes(Paths.get("./signed.txn")), SignedTransaction.class);
                System.out.println("Signed transaction with txid: " + decodedSignedTransaction.transactionID);

                // Msgpack encode the signed transaction
                byte[] encodedTxBytes = Encoder.encodeToMsgPack(decodedSignedTransaction);

                // submit the encoded transaction to the network
                TransactionID id = algodApiInstance.rawTransaction(encodedTxBytes);
                System.out.println("Successfully sent tx with id: " + id);
                waitForConfirmation(id.getTxId());

            } catch (Exception e) {
                System.out.println("Submit Exception: " + e);
            }

        }

        public static void main(String args[]) throws Exception {
            SaveMultisigTransactionOffline mn = new SaveMultisigTransactionOffline();
            // mn.writeMultisigUnsignedTransaction();
            // mn.readMultisigUnsignedTransaction();

            mn.writeMultisigSignedTransaction();
            mn.readMultisigSignedTransaction();

        }
    } 
    ```

    ```go tab="Go"
    package main

    import (
        json "encoding/json"
        "fmt"
        "io/ioutil"

        "github.com/algorand/go-algorand-sdk/client/algod"
        "github.com/algorand/go-algorand-sdk/crypto"
        "github.com/algorand/go-algorand-sdk/encoding/msgpack"
        "github.com/algorand/go-algorand-sdk/mnemonic"
        "github.com/algorand/go-algorand-sdk/transaction"
        "github.com/algorand/go-algorand-sdk/types"
    )

    // Function that waits for a given txId to be confirmed by the network
    func waitForConfirmation(algodClient algod.Client, txID string) {
        for {
            pt, err := algodClient.PendingTransactionInformation(txID)
            if err != nil {
                fmt.Printf("waiting for confirmation... (pool error, if any): %s\n", err)
                continue
            }
            if pt.ConfirmedRound > 0 {
                fmt.Printf("Transaction "+pt.TxID+" confirmed in round %d\n", pt.ConfirmedRound)
                break
            }
            nodeStatus, err := algodClient.Status()
            if err != nil {
                fmt.Printf("error getting algod status: %s\n", err)
                return
            }
            algodClient.StatusAfterBlock(nodeStatus.LastRound + 1)
        }
    }

    // createAccounts
    func createAccounts() {
        account1 := crypto.GenerateAccount()
        account2 := crypto.GenerateAccount()
        account3 := crypto.GenerateAccount()
        address1 := account1.Address.String()
        address2 := account2.Address.String()
        address3 := account3.Address.String()

        mnemonic1, err := mnemonic.FromPrivateKey(account1.PrivateKey)
        if err != nil {
            return
        }
        mnemonic2, err := mnemonic.FromPrivateKey(account2.PrivateKey)
        if err != nil {
            return
        }
        mnemonic3, err := mnemonic.FromPrivateKey(account3.PrivateKey)
        if err != nil {
            return
        }
        fmt.Printf("1 : \"%s\"\n", address1)
        fmt.Printf("2 : \"%s\"\n", address2)
        fmt.Printf("3 : \"%s\"\n", address3)
        fmt.Printf("")
        fmt.Printf("Copy off accounts above and add TestNet Algo funds using the TestNet Dispenser at https://bank.testnet.algorand.network/\n")
        fmt.Printf("Copy off the following mnemonic code for use in Step 1B\n")
        fmt.Printf("\n")
        fmt.Printf("mnemonic1 := \"%s\"\n", mnemonic1)
        fmt.Printf("mnemonic2 := \"%s\"\n", mnemonic2)
        fmt.Printf("mnemonic3 := \"%s\"\n", mnemonic3)

    }

    // Accounts to be used through examples
    func recoverAccounts() (map[int][]byte, map[int]string) {
        // Shown for demonstration purposes. NEVER reveal secret mnemonics in practice.
        // Change these values to use previously generated and funded accounts

        var pks = map[int]string{
        	1: "Account Address 1 ",
        	2: "Account Address 2 ",
        	3: "Account Address 3 ",
        }
        mnemonic1 := "your-25-word-mnemonic<PLACEHOLDER>"
        mnemonic2 := "your-25-word-mnemonic<PLACEHOLDER>"
        mnemonic3 := "your-25-word-mnemonic<PLACEHOLDER>"

        mnemonics := []string{mnemonic1, mnemonic2, mnemonic3}
        var sks = make(map[int][]byte)
        for i, m := range mnemonics {
            var err error
            sks[i+1], err = mnemonic.ToPrivateKey(m)
            if err != nil {
                fmt.Printf("Issue with account %d private key conversion.", i+1)
            } else {
                fmt.Printf("Loaded Key %d: %s\n", i+1, pks[i+1])
            }
        }
        return sks, pks
    }

    // PrettyPrint prints Go structs
    func PrettyPrint(data interface{}) {
        var p []byte
        //    var err := error
        p, err := json.MarshalIndent(data, "", "\t")
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Printf("%s \n", p)
    }

    // utility funciton to setup connection to node
    func setupConnection() algod.Client {
        // const algodToken = "algod-token<PLACEHOLDER>"
        // const algodAddress = "algod-address<PLACEHOLDER>"
        algodClient, err := algod.MakeClient(algodAddress, algodToken)
        if err != nil {
            fmt.Printf("failed to make algod client: %s\n", err)
        }
        return algodClient
    }

    func saveUnsignedMultisigTransaction() {

        // setup connection
        algodClient := setupConnection()

        // recover account for example
        //addr, _ := recoverAccounts();
        sks, pks := recoverAccounts()
        PrettyPrint(sks)
        // get network suggested parameters
        txParams, err := algodClient.SuggestedParams()
        if err != nil {
            fmt.Printf("error getting suggested tx params: %s\n", err)
            return
        }
        addr1, _ := types.DecodeAddress(pks[1])
        addr2, _ := types.DecodeAddress(pks[2])
        addr3, _ := types.DecodeAddress(pks[3])
        ma, err := crypto.MultisigAccountWithParams(1, 2, []types.Address{
            addr1,
            addr2,
            addr3,
        })
        if err != nil {
            panic("invalid multisig parameters")
        }
        fromAddr, _ := ma.Address()
        fmt.Printf("Here is your multisig address : %s \n", fromAddr.String())
        fmt.Println("Please go to: https://bank.testnet.algorand.network/ to fund your multisig account.")
        // wait for Enter Key
        // this does not return if debugging in go in vscode. known issue
        // comment out this line after adding funds to the account, uncomment if loading for the fist time
        // or just set a break point here and add funds
        // fmt.Scanln()
        // send from multisig to account 3
        txn, err := transaction.MakePaymentTxn(
            fromAddr.String(),
            addr3.String(),
            txParams.Fee,           // fee per byte
            1000000,                // amount
            txParams.LastRound,     // first valid round
            txParams.LastRound+100, // last valid round
            nil,                    // note
            "",                     // closeRemainderTo
            txParams.GenesisID,     // genesisID
            txParams.GenesisHash,   // genesisHash (Cannot be empty in practice)
        )
        unsignedTx := types.SignedTxn{
            Txn: txn,
        }
        // save unsigned transction to file
        err = ioutil.WriteFile("./unsigned.txn", msgpack.Encode(unsignedTx), 0644)
        if err != nil {
            fmt.Printf("Failed in saving trx to file, error %s\n", err)
            return
        }
        fmt.Printf("Saved unsigned transaction to file\n")

        err = ioutil.WriteFile("./ma.txn", msgpack.Encode(ma), 0644)
        if err != nil {
            fmt.Printf("Failed in saving ma to file, error %s\n", err)
            return
        }

        fmt.Printf("Saved ma to file\n")

        return

    }
    func readUnsignedMultisigTransaction() {

        // setup connection
        algodClient := setupConnection()

        // read unsigned transaction from file
        dat, err := ioutil.ReadFile("./unsigned.txn")
        if err != nil {
            fmt.Printf("Error reading transaction from file: %s\n", err)
            return
        }
        var unsignedTxRaw types.SignedTxn
        var unsignedTxn types.Transaction

        msgpack.Decode(dat, &unsignedTxRaw)

        unsignedTxn = unsignedTxRaw.Txn

        // recover account and sign transaction
        sks, pks := recoverAccounts()
        addr := pks[1]
        fmt.Printf("Address is: %s\n", addr)

        msgpack.Decode(dat, &unsignedTxRaw)
        unsignedTxn = unsignedTxRaw.Txn

        // read ma from file
        datma, err := ioutil.ReadFile("./ma.txn")
        if err != nil {
            fmt.Printf("Error reading ma from file: %s\n", err)
            return
        }

        var ma crypto.MultisigAccount
        msgpack.Decode(datma, &ma)

        txid, txBytes, err := crypto.SignMultisigTransaction(sks[1], ma, unsignedTxn)
        if err != nil {
            panic("could not sign multisig transaction")
        }
        fmt.Printf("Made partially-signed multisig transaction with TxID %s: %x\n", txid, txBytes)

        // append our signature to readTxBytes
        txid, twoOfThreeTxBytes, err := crypto.AppendMultisigTransaction(sks[2], ma, txBytes)
        if err != nil {
            panic("could not append signature to multisig transaction")
        }
        fmt.Printf("Made 2-out-of-3 multisig transaction with TxID %s: %x\n", txid, twoOfThreeTxBytes)

        // otherTxBytes := ... // generate another raw multisig transaction somehow
        // txid, mergedTxBytes, err := crypto.MergeMultisigTransactions(twoOfThreeTxBytes, otherTxBytes)

        // Send transaction to the network
        sendResponse, err := algodClient.SendRawTransaction(twoOfThreeTxBytes)
        if err != nil {
            fmt.Printf("Failed to create payment transaction: %v\n", err)
            return
        }

        fmt.Printf("Transaction ID: %s\n", sendResponse.TxID)
        waitForConfirmation(algodClient, sendResponse.TxID)
        return

    }

    func saveSignedMultisigTransaction() {

        // setup connection
        algodClient := setupConnection()

        // recover account for example
        //addr, _ := recoverAccounts();
        sks, pks := recoverAccounts()
        PrettyPrint(sks)
        // get network suggested parameters
        txParams, err := algodClient.SuggestedParams()
        if err != nil {
            fmt.Printf("error getting suggested tx params: %s\n", err)
            return
        }
        addr1, _ := types.DecodeAddress(pks[1])
        addr2, _ := types.DecodeAddress(pks[2])
        addr3, _ := types.DecodeAddress(pks[3])
        ma, err := crypto.MultisigAccountWithParams(1, 2, []types.Address{
            addr1,
            addr2,
            addr3,
        })
        if err != nil {
            panic("invalid multisig parameters")
        }
        fromAddr, _ := ma.Address()
        fmt.Printf("Here is your multisig address : %s \n", fromAddr.String())
        fmt.Println("Please go to: https://bank.testnet.algorand.network/ to fund your multisig account.")
        // wait for Enter Key
        // this does not return if debugging in go in vscode. known issue
        // comment out this line after adding funds to the account, uncomment if loading for the fist time
        // or just set a break point here and add funds
        // fmt.Scanln()
        // send from multisig to account 3
        txn, err := transaction.MakePaymentTxn(
            fromAddr.String(),
            addr3.String(),
            txParams.Fee,           // fee per byte
            1000000,                // amount
            txParams.LastRound,     // first valid round
            txParams.LastRound+100, // last valid round
            nil,                    // note
            "",                     // closeRemainderTo
            txParams.GenesisID,     // genesisID
            txParams.GenesisHash,   // genesisHash (Cannot be empty in practice)
        )

        txid, txBytes, err := crypto.SignMultisigTransaction(sks[1], ma, txn)
        if err != nil {
            panic("could not sign multisig transaction")
        }
        fmt.Printf("Made partially-signed multisig transaction with TxID %s: %x\n", txid, txBytes)

        // append our signature to readTxBytes
        txid, twoOfThreeTxBytes, err := crypto.AppendMultisigTransaction(sks[2], ma, txBytes)
        if err != nil {
            panic("could not append signature to multisig transaction")
        }
        fmt.Printf("Made 2-out-of-3 multisig transaction with TxID %s: %x\n", txid, twoOfThreeTxBytes)

        // otherTxBytes := ... // generate another raw multisig transaction somehow
        // txid, mergedTxBytes, err := crypto.MergeMultisigTransactions(twoOfThreeTxBytes, otherTxBytes)

        //Save the signed transaction to file
        err = ioutil.WriteFile("./signed.stxn", twoOfThreeTxBytes, 0644)
        if err != nil {

            fmt.Printf("Failed in saving stxn to file, error %s\n", err)
            return
        }
        fmt.Printf("Saved signed transaction to file\n")

        return

    }
    func readSignedMultisigTransaction() {

        // setup connection
        algodClient := setupConnection()

        // read unsigned transaction from file
        dat, err := ioutil.ReadFile("./signed.stxn")
        if err != nil {
            fmt.Printf("Error reading signed transaction from file: %s\n", err)
            return
        }

        // send the transaction to the network
        sendResponse, err := algodClient.SendRawTransaction(dat)
        if err != nil {
            fmt.Printf("failed to send transaction: %s\n", err)
            return
        }

        fmt.Printf("Transaction ID: %s\n", sendResponse.TxID)
        waitForConfirmation(algodClient, sendResponse.TxID)
    }

    func main() {

        createAccounts()
        //Copy off accounts above and add TestNet Algo funds using the
        //TestNet Dispenser at https://bank.testnet.algorand.network/

        saveUnsignedMultisigTransaction()
        readUnsignedMultisigTransaction()

        saveSignedMultisigTransaction()
        readSignedMultisigTransaction()

    }   
    ```

    ```java tab="goal"
    Create a multisig account by listing all of the accounts in the multisig and specifying the threshold number of accounts to sign with the -T flag
    $ goal account multisig new my-account1<PLACEHOLDER> my-account2<PLACEHOLDER> my-account3<PLACEHOLDER> etc… -T 2    

    Create an unsigned transaction and write to file
    $ goal clerk send --from my-multisig-account<PLACEHOLDER>  --to AZLR2XP4O2WFHLX6TX7AZVY23HLVLG3K5K3FRIKIYDOYN6ISIF54SA4RNY --fee=1000 --amount=1000000 --out="unsigned.txn"

    Sign by the required number of accounts to meet the threshold. 
    $ goal clerk multisig sign -a F my-account1<PLACEHOLDER> -t=unsigned.txn
    goal clerk multisig sign -a F my-account2<PLACEHOLDER> -t=unsigned.txn

    Merge signings 
    $ goal clerk multisig merge --out signed.txn unsigned.txn

    Broadcast 
    $ goal clerk rawsend --filename signed.txn
    ```
