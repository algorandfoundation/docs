title: Interact with smart contracts

This guide covers using smart contracts with the Algorand SDKs. Smart contracts form the basis for applications written in [Transaction Execution Approval Language (TEAL)](../../avm/teal/index.md) or with Python using the [PyTeal](../../pyteal/index.md) library.


!!! info
    The example code snippets are provided throughout this page and are abbreviated for conciseness and clarity. Full running code examples for each SDK are available within the GitHub repo at [/examples/smart_contracts](https://github.com/algorand/docs/tree/master/examples/smart_contracts) and for [download](https://github.com/algorand/docs/blob/master/examples/smart_contracts/smart_contracts.zip?raw=true) (.zip).

# Application lifecycle

This guide follows an application throughout its [lifecycle](../apps/index.md#the-lifecycle-of-a-smart-contract) from initial creation, to usage, to modification and finally deletion. The application stores the number of times it is called within its _global state_ and also stores the number of times each user account calls the application within their _local state_. Midway through the lifecycle, the application is upgraded to add an additional key:value pair to the user's _local storage_ for storing the call timestamp. 

# Environment setup

This guide requires two accounts:

=== "Python"
	```python
    # user declared account mnemonics
    creator_mnemonic = "Your first 25-word mnemonic goes here"
    user_mnemonic = "A second distinct 25-word mnemonic goes here"
    ```

=== "JavaScript"
	```javascript
    // user declared account mnemonics
    creatorMnemonic = "Your 25-word mnemonic goes here";
    userMnemonic = "A second distinct 25-word mnemonic goes here";
    ```

=== "Java"
	```Java
    // user declared account mnemonics
    String creatorMnemonic = "Your 25-word mnemonic goes here";
    String userMnemonic = "A second distinct 25-word mnemonic goes here";
    ```

=== "Go"
	```go
    // user defined mnemonics
    const creatorMnemonic = "Your 25-word mnemonic goes here"
    const userMnemonic = "A second distinct 25-word mnemonic goes here"
    ```

An `algod` client connection is also required. The following connects using Sandbox:

=== "Python"
    <!-- ===PYSDK_ALGOD_CREATE_CLIENT=== -->
```python
# Create a new algod client, configured to connect to our local sandbox
algod_address = "http://localhost:4001"
algod_token = "a" * 64
algod_client = algod.AlgodClient(algod_token, algod_address)

# Or, if necessary, pass alternate headers

# Create a new client with an alternate api key header
special_algod_client = algod.AlgodClient(
    "", algod_address, headers={"X-API-Key": algod_token}
)
```
[Snippet Source](https://github.com/barnjamin/py-algorand-sdk/blob/doc-examples/_examples/overview.py#L10-L21)
    <!-- ===PYSDK_ALGOD_CREATE_CLIENT=== -->

=== "JavaScript"
    <!-- ===JSSDK_ALGOD_CREATE_CLIENT=== -->
	```javascript
    // user declared algod connection parameters
    algodAddress = "http://localhost:4001";
    algodToken = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
    let algodClient = new algosdk.Algodv2(algodToken, algodServer);
    ```
    <!-- ===JSSDK_ALGOD_CREATE_CLIENT=== -->

=== "Java"
    <!-- ===JAVASDK_ALGOD_CREATE_CLIENT=== -->
```java
        String algodHost = "http://localhost";
        int algodPort = 4001;
        String algodToken = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
        AlgodClient algodClient = new AlgodClient(algodHost, algodPort, algodToken);

        // OR if the API provider requires a specific header key for the token
        String tokenHeader = "X-API-Key";
        AlgodClient otherAlgodClient = new AlgodClient(algodHost, algodPort, algodToken, tokenHeader);
```
[Snippet Source](https://github.com/barnjamin/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/Overview.java#L91-L99)
    <!-- ===JAVASDK_ALGOD_CREATE_CLIENT=== -->

=== "Go"
    <!-- ===GOSDK_ALGOD_CREATE_CLIENT=== -->
```go
	// Create a new algod client, configured to connect to out local sandbox
	var algodAddress = "http://localhost:4001"
	var algodToken = strings.Repeat("a", 64)
	algodClient, err := algod.MakeClient(
		algodAddress,
		algodToken,
	)

	// Or, if necessary, pass alternate headers

	var algodHeader common.Header
	algodHeader.Key = "X-API-Key"
	algodHeader.Value = algodToken
	algodClientWithHeaders, err := algod.MakeClientWithHeaders(
		algodAddress,
		algodToken,
		[]*common.Header{&algodHeader},
	)
```
[Snippet Source](https://github.com/barnjamin/go-algorand-sdk/blob/examples/_examples/overview.go#L13-L31)
    <!-- ===GOSDK_ALGOD_CREATE_CLIENT=== -->

!!! Info
    Ensure the `algod` node has the _"EnableDeveloperAPI"_ parameter set to **true** within the `config.json` file. This is required to enable the SDK access to the _compile_ and _dryrun_ endpoints.

# Declarations

All smart contracts are comprised of state storage, an approval program and a clear program. Details of each are found within the [smart contract guide](../apps/index.md).

## State storage
Begin by defining the application's _global_schema_ and _local_schema_ storage requirements. These values are immutable once the application is created, so they must specify the maximum number required by the initial application and any future updates. 

The example application defined below may hold up to one each of `bytes` and `ints` value within the _local storage_ of the user account, as well as a single `ints` value within _global storage_ of the application:

=== "Python"
    <!-- ===PYSDK_APP_SCHEMA=== -->
```python
# create schema for both global and local state, specifying
# how many keys of each type we need to have available
local_schema = transaction.StateSchema(num_uints=1, num_byte_slices=1)
global_schema = transaction.StateSchema(num_uints=1, num_byte_slices=1)
```
[Snippet Source](https://github.com/barnjamin/py-algorand-sdk/blob/doc-examples/_examples/apps.py#L11-L15)
    <!-- ===PYSDK_APP_SCHEMA=== -->

=== "JavaScript"
    <!-- ===JSSDK_APP_SCHEMA=== -->
	```javascript
    // declare application state storage (immutable)
    localInts = 1;
    localBytes = 1;
    globalInts = 1;
    globalBytes = 0;
    ```
    <!-- ===JSSDK_APP_SCHEMA=== -->

=== "Java"
    <!-- ===JAVASDK_APP_SCHEMA=== -->
```java
        int localInts = 1;
        int localBytes = 1;
        int globalInts = 1;
        int globalBytes = 0;
        StateSchema localSchema = new StateSchema(localInts, localBytes);
        StateSchema globalSchema = new StateSchema(globalInts, globalBytes);
```
[Snippet Source](https://github.com/barnjamin/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/AppExamples.java#L40-L46)
    <!-- ===JAVASDK_APP_SCHEMA=== -->

=== "Go"
    <!-- ===GOSDK_APP_SCHEMA=== -->
	```go
    // declare application state storage (immutable)
    const localInts = 1
    const localBytes = 1
    const globalInts = 1
    const globalBytes = 0

    // define schema
    globalSchema := types.StateSchema{NumUint: uint64(globalInts), NumByteSlice: uint64(globalBytes)}
    localSchema := types.StateSchema{NumUint: uint64(localInts), NumByteSlice: uint64(localBytes)}
    ```
    <!-- ===GOSDK_APP_SCHEMA=== -->

!!! Info
    The example application is not allowed to hold any `bytes` value within global storage.


## Approval program

The [approval program](../apps/#the-lifecycle-of-a-stateful-smart-contract) handles the main logic of the application.

## Clear program

This is the most basic [clear program](../apps/#the-lifecycle-of-a-stateful-smart-contract) and returns _true_ when an account clears its participation in a smart contract:

=== "Python"
    <!-- ===PYSDK_APP_SOURCE=== -->
```python
# read the `.teal` source files from disk
with open("application/approval.teal", "r") as f:
    approval_program = f.read()

with open("application/clear.teal", "r") as f:
    clear_program = f.read()
```
[Snippet Source](https://github.com/barnjamin/py-algorand-sdk/blob/doc-examples/_examples/apps.py#L18-L24)
    <!-- ===PYSDK_APP_SOURCE=== -->

=== "JavaScript"
    <!-- ===JSSDK_APP_SOURCE=== -->
	```javascript
    // declare clear state program source
    clearProgramSource = `#pragma version 4
    int 1
    `;
    ```
    <!-- ===JSSDK_APP_SOURCE=== -->

=== "Java"
    <!-- ===JAVASDK_APP_SOURCE=== -->
```java
        // Read in the `teal` source files as a string
        String approvalSource = Files.readString(Paths.get("application/approval.teal"));
        String clearSource = Files.readString(Paths.get("application/clear.teal"));
```
[Snippet Source](https://github.com/barnjamin/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/AppExamples.java#L49-L52)
    <!-- ===JAVASDK_APP_SOURCE=== -->

=== "Go"
    <!-- ===GOSDK_APP_SOURCE=== -->
	```go
    // declare clear state program source
    const clearProgramSource = `#pragma version 4
    int 1
    `
    ```
    <!-- ===GOSDK_APP_SOURCE=== -->

# Application methods

## Create

The creator will deploy the application using the [create app](../apps#creating-the-smart-contract) method. It requires 7 parameters:

- sender: address, representing the creator of the app
- sp: suggested parameters obtained from the network
- on_complete: enum value, representing NoOp
- approval_program: compiled program
- clear program: compiled program
- local_schema: maximum _local storage_ allocation, immutable
- global_schema: maximum _global storage_ allocation, immutable

Use the creator_mnemonic to define sender:

=== "Python"
    <!-- ===PYSDK_ACCOUNT_RECOVER_MNEMONIC=== -->
```python
mn = "cost piano sample enough south bar diet garden nasty mystery mesh sadness convince bacon best patch surround protect drum actress entire vacuum begin abandon hair"
pk = mnemonic.to_private_key(mn)
print(f"Base64 encoded private key: {pk}")
addr = account.address_from_private_key(pk)
print(f"Address: {addr}")
```
[Snippet Source](https://github.com/barnjamin/py-algorand-sdk/blob/doc-examples/_examples/account.py#L12-L17)
    <!-- ===PYSDK_ACCOUNT_RECOVER_MNEMONIC=== -->

=== "JavaScript"
    <!-- ===JSSDK_ACCOUNT_RECOVER_MNEMONIC=== -->
	```javascript
    // get account from mnemonic
    let creatorAccount = algosdk.mnemonicToSecretKey(creatorMnemonic);
    let sender = creatorAccount.addr;
    ```
    <!-- ===JSSDK_ACCOUNT_RECOVER_MNEMONIC=== -->

=== "Java"
    <!-- ===JAVASDK_ACCOUNT_RECOVER_MNEMONIC=== -->
```java
                // Space delimited 25 word mnemonic
                String mn = "cost piano sample enough south bar diet garden nasty mystery mesh sadness convince bacon best patch surround protect drum actress entire vacuum begin abandon hair";
                // We can get the private key
                byte[] key = Mnemonic.toKey(mn);
                // Or just init the account directly from the mnemonic
                Account acct = new Account(mn);
```
[Snippet Source](https://github.com/barnjamin/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/AccountExamples.java#L63-L69)
    <!-- ===JAVASDK_ACCOUNT_RECOVER_MNEMONIC=== -->

=== "Go"
    <!-- ===GOSDK_ACCOUNT_RECOVER_MNEMONIC=== -->
```go
	k, err := mnemonic.ToPrivateKey(mn)
	if err != nil {
		log.Fatalf("failed to parse mnemonic: %s", err)
	}

	recovered, err := crypto.AccountFromPrivateKey(k)
	if err != nil {
		log.Fatalf("failed to recover account from key: %s", err)
	}

	log.Printf("%+v", recovered)
```
[Snippet Source](https://github.com/barnjamin/go-algorand-sdk/blob/examples/_examples/account.go#L23-L34)
    <!-- ===GOSDK_ACCOUNT_RECOVER_MNEMONIC=== -->

Compile the programs using the `compile` endpoint:

=== "Python"
    <!-- ===PYSDK_APP_COMPILE=== -->
```python
# pass the `.teal` files to the compile endpoint
# and b64 decode the result to bytes
approval_result = algod_client.compile(approval_program)
approval_binary = base64.b64decode(approval_result["result"])

clear_result = algod_client.compile(clear_program)
clear_binary = base64.b64decode(clear_result["result"])
```
[Snippet Source](https://github.com/barnjamin/py-algorand-sdk/blob/doc-examples/_examples/apps.py#L29-L36)
    <!-- ===PYSDK_APP_COMPILE=== -->

=== "JavaScript"
    <!-- ===JSSDK_APP_COMPILE=== -->
	```javascript
    // helper function to compile program source  
    async function compileProgram(client, programSource) {
        let encoder = new TextEncoder();
        let programBytes = encoder.encode(programSource);
        let compileResponse = await client.compile(programBytes).do();
        let compiledBytes = new Uint8Array(Buffer.from(compileResponse.result, "base64"));
        return compiledBytes;
    }
    ```
    <!-- ===JSSDK_APP_COMPILE=== -->

=== "Java"
    <!-- ===JAVASDK_APP_COMPILE=== -->
```java
        CompileResponse approvalResponse = algodClient.TealCompile().source(approvalSource.getBytes()).execute()
                .body();
        CompileResponse clearResponse = algodClient.TealCompile().source(clearSource.getBytes()).execute()
                .body();

        TEALProgram approvalProg = new TEALProgram(approvalResponse.result);
        TEALProgram clearProg = new TEALProgram(clearResponse.result);
```
[Snippet Source](https://github.com/barnjamin/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/AppExamples.java#L55-L62)
    <!-- ===JAVASDK_APP_COMPILE=== -->

=== "Go"
    <!-- ===GOSDK_APP_COMPILE=== -->
	```go
    func compileProgram(client *algod.Client, programSource string) (compiledProgram []byte) {
        compileResponse, err := client.TealCompile([]byte(programSource)).Do(context.Background())
        if err != nil {
            fmt.Printf("Issue with compile: %s\n", err)
            return
        }
        compiledProgram, _ = base64.StdEncoding.DecodeString(compileResponse.Result)
        return compiledProgram
    }
    ```
    <!-- ===GOSDK_APP_COMPILE=== -->

Construct the transaction with defined values then sign, send, and await confirmation

=== "Python"
    <!-- ===PYSDK_APP_CREATE=== -->
```python
sp = algod_client.suggested_params()
# create the app create transaction, passing compiled programs and schema
app_create_txn = transaction.ApplicationCreateTxn(
    creator.address,
    sp,
    transaction.OnComplete.NoOpOC,
    approval_program=approval_binary,
    clear_program=clear_binary,
    global_schema=global_schema,
    local_schema=local_schema,
)
# sign transaction
signed_create_txn = app_create_txn.sign(creator.private_key)
txid = algod_client.send_transaction(signed_create_txn)
result = transaction.wait_for_confirmation(algod_client, txid, 4)
app_id = result["application-index"]
print(f"Created app with id: {app_id}")
```
[Snippet Source](https://github.com/barnjamin/py-algorand-sdk/blob/doc-examples/_examples/apps.py#L39-L56)
    <!-- ===PYSDK_APP_CREATE=== -->

=== "JavaScript"
    <!-- ===JSSDK_APP_CREATE=== -->
	```javascript
    // create unsigned transaction
    let txn = algosdk.makeApplicationCreateTxn(sender, params, onComplete, 
                                            approvalProgram, clearProgram, 
                                            localInts, localBytes, globalInts, globalBytes,);
    let txId = txn.txID().toString();

    // Sign the transaction
    let signedTxn = txn.signTxn(creatorAccount.sk);
    console.log("Signed transaction with txID: %s", txId);

    // Submit the transaction
    await client.sendRawTransaction(signedTxn).do();

    // Wait for transaction to be confirmed
    confirmedTxn = await algosdk.waitForConfirmation(client, txId, 4);
    //Get the completed Transaction
    console.log("Transaction " + txId + " confirmed in round " + confirmedTxn["confirmed-round"]);

    // display results
    let transactionResponse = await client.pendingTransactionInformation(txId).do();
    let appId = transactionResponse['application-index'];
    console.log("Created new app-id: ",appId);
    ```
    <!-- ===JSSDK_APP_CREATE=== -->

=== "Java"
    <!-- ===JAVASDK_APP_CREATE=== -->
```java
        Response<TransactionParametersResponse> rsp = algodClient.TransactionParams().execute();
        TransactionParametersResponse sp = rsp.body();

        Transaction appCreate = ApplicationCreateTransactionBuilder.Builder()
                .sender(creator.getAddress())
                .suggestedParams(sp)
                .approvalProgram(approvalProg)
                .clearStateProgram(clearProg)
                .localStateSchema(localSchema)
                .globalStateSchema(globalSchema)
                .build();

        SignedTransaction signedAppCreate = creator.signTransaction(appCreate);
        Response<PostTransactionsResponse> createResponse = algodClient.RawTransaction()
                .rawtxn(Encoder.encodeToMsgPack(signedAppCreate)).execute();

        PendingTransactionResponse result = Utils.waitForConfirmation(algodClient, createResponse.body().txId, 4);
        Long appId = result.applicationIndex;
        System.out.printf("Created application with id: %d\n", appId);
```
[Snippet Source](https://github.com/barnjamin/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/AppExamples.java#L65-L84)
    <!-- ===JAVASDK_APP_CREATE=== -->

=== "Go"
    <!-- ===GOSDK_APP_CREATE=== -->
	```go
    // create unsigned transaction
    txn, _ := transaction.MakeApplicationCreateTx(false, approvalProgram, clearProgram, globalSchema, localSchema, 
                        nil, nil, nil, nil, params, creatorAccount.Address, nil, 
                        types.Digest{}, [32]byte{}, types.Address{}, uint32(0))

    // Sign the transaction
    txID, signedTxn, _ := crypto.SignTransaction(creatorAccount.PrivateKey, txn)
    fmt.Printf("Signed txid: %s\n", txID)

    // Submit the transaction
    sendResponse, _ := client.SendRawTransaction(signedTxn).Do(context.Background())
    fmt.Printf("Submitted transaction %s\n", sendResponse)

    // Wait for confirmation
    confirmedTxn, err := transaction.WaitForConfirmation(client, txID, 4, context.Background())
    if err != nil {
		fmt.Printf("Error waiting for confirmation on txID: %s\n", txID)
        return
    }
	fmt.Printf("Confirmed Transaction: %s in Round %d\n", txID ,confirmedTxn.ConfirmedRound)


    // display results
    confirmedTxn, _, _ := client.PendingTransactionInformation(txID).Do(context.Background())
    appId = confirmedTxn.ApplicationIndex
    fmt.Printf("Created new app-id: %d\n", appId)
    ```
    <!-- ===GOSDK_APP_CREATE=== -->


!!! Notice
    Note down the app-id from the confirmed transaction response. Place this value into the `index` parameter within all remaining code samples. 

## Opt-in

The user must [opt-in](../apps/#opt-in-to-the-smart-contract) to use the application. This method requires 3 parameters:

- sender: address, representing the user intending to opt-in to using the app
- sp: suggested parameters obtained from the network
- index: the app-id as defined by the create method result

Use the user_mnemonic to define sender:

=== "Python"
<!-- ===PYSDK_ACCOUNT_RECOVER_MNEMONIC=== -->
```python
mn = "cost piano sample enough south bar diet garden nasty mystery mesh sadness convince bacon best patch surround protect drum actress entire vacuum begin abandon hair"
pk = mnemonic.to_private_key(mn)
print(f"Base64 encoded private key: {pk}")
addr = account.address_from_private_key(pk)
print(f"Address: {addr}")
```
[Snippet Source](https://github.com/barnjamin/py-algorand-sdk/blob/doc-examples/_examples/account.py#L12-L17)
<!-- ===PYSDK_ACCOUNT_RECOVER_MNEMONIC=== -->

=== "JavaScript"
<!-- ===JSSDK_ACCOUNT_RECOVER_MNEMONIC=== -->
	```javascript
    // get accounts from mnemonic
    let userAccount = algosdk.mnemonicToSecretKey(userMnemonic);
    let sender = userAccount.addr;
    ```
<!-- ===JSSDK_ACCOUNT_RECOVER_MNEMONIC=== -->

=== "Java"
<!-- ===JAVASDK_ACCOUNT_RECOVER_MNEMONIC=== -->
```java
                // Space delimited 25 word mnemonic
                String mn = "cost piano sample enough south bar diet garden nasty mystery mesh sadness convince bacon best patch surround protect drum actress entire vacuum begin abandon hair";
                // We can get the private key
                byte[] key = Mnemonic.toKey(mn);
                // Or just init the account directly from the mnemonic
                Account acct = new Account(mn);
```
[Snippet Source](https://github.com/barnjamin/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/AccountExamples.java#L63-L69)
<!-- ===JAVASDK_ACCOUNT_RECOVER_MNEMONIC=== -->

=== "Go"
<!-- ===GOSDK_ACCOUNT_RECOVER_MNEMONIC=== -->
```go
	k, err := mnemonic.ToPrivateKey(mn)
	if err != nil {
		log.Fatalf("failed to parse mnemonic: %s", err)
	}

	recovered, err := crypto.AccountFromPrivateKey(k)
	if err != nil {
		log.Fatalf("failed to recover account from key: %s", err)
	}

	log.Printf("%+v", recovered)
```
[Snippet Source](https://github.com/barnjamin/go-algorand-sdk/blob/examples/_examples/account.go#L23-L34)
<!-- ===GOSDK_ACCOUNT_RECOVER_MNEMONIC=== -->

Construct the transaction with defined values then sign, send, and await confirmation:

=== "Python"
    <!-- ===PYSDK_APP_OPTIN=== -->
```python
opt_in_txn = transaction.ApplicationOptInTxn(user.address, sp, app_id)
signed_opt_in = opt_in_txn.sign(user.private_key)
txid = algod_client.send_transaction(signed_opt_in)
optin_result = transaction.wait_for_confirmation(algod_client, txid, 4)
assert optin_result["confirmed-round"] > 0
```
[Snippet Source](https://github.com/barnjamin/py-algorand-sdk/blob/doc-examples/_examples/apps.py#L59-L64)
    <!-- ===PYSDK_APP_OPTIN=== -->

=== "JavaScript"
    <!-- ===JSSDK_APP_OPTIN=== -->
	```javascript
    // create unsigned transaction
    let txn = algosdk.makeApplicationOptInTxn(sender, params, app_id);

    // ... sign send wait

    let transactionResponse = await client.pendingTransactionInformation(txId).do();
    console.log("Opted-in to app-id:",transactionResponse['txn']['txn']['apid'])
    ```
    <!-- ===JSSDK_APP_OPTIN=== -->

=== "Java"
    <!-- ===JAVASDK_APP_OPTIN=== -->
```java
        Transaction optInTxn = ApplicationOptInTransactionBuilder.Builder()
                .sender(user.getAddress())
                .suggestedParams(sp)
                .applicationId(appId)
                .build();

        SignedTransaction signedOptIn = user.signTransaction(optInTxn);
        Response<PostTransactionsResponse> optInResponse = algodClient.RawTransaction()
                .rawtxn(Encoder.encodeToMsgPack(signedOptIn)).execute();

        PendingTransactionResponse optInResult = Utils.waitForConfirmation(algodClient, optInResponse.body().txId, 4);
        assert optInResult.confirmedRound > 0;
```
[Snippet Source](https://github.com/barnjamin/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/AppExamples.java#L87-L99)
    <!-- ===JAVASDK_APP_OPTIN=== -->

=== "Go"
    <!-- ===GOSDK_APP_OPTIN=== -->
	```go
    // create unsigned transaction
    txn, _ := transaction.MakeApplicationOptInTx(index, nil, nil, nil, nil, params,
                            sender, nil, types.Digest{}, [32]byte{}, types.Address{})

    // ... sign, send, await

    // display results
    confirmedTxn, _, _ := client.PendingTransactionInformation(txID).Do(context.Background())
    fmt.Printf("Oped-in to app-id: %d\n", confirmedTxn.Transaction.Txn.ApplicationID)
    ```
    <!-- ===GOSDK_APP_OPTIN=== -->


## Call (NoOp)

The user may now [call](../apps/#call-the-stateful-smart-contract) the application. This method requires 3 parameters:

- sender: address, representing the user intending to optin to using the app
- sp: suggested parameters obtained from the network
- index: the app-id as defined by the create method result

=== "Python"
    <!-- ===PYSDK_APP_NOOP=== -->
```python
noop_txn = transaction.ApplicationNoOpTxn(user.address, sp, app_id)
signed_noop = noop_txn.sign(user.private_key)
txid = algod_client.send_transaction(signed_noop)
noop_result = transaction.wait_for_confirmation(algod_client, txid, 4)
assert noop_result["confirmed-round"] > 0
```
[Snippet Source](https://github.com/barnjamin/py-algorand-sdk/blob/doc-examples/_examples/apps.py#L67-L72)
    <!-- ===PYSDK_APP_NOOP=== -->

=== "JavaScript"
    <!-- ===JSSDK_APP_NOOP=== -->
	```javascript
    // create unsigned transaction
    let txn = algosdk.makeApplicationNoOpTxn(sender, params, index, appArgs)

    // ... sign, send, wait

    // display results
    let transactionResponse = await client.pendingTransactionInformation(txId).do();
    console.log("Called app-id:",transactionResponse['txn']['txn']['apid'])
    if (transactionResponse['global-state-delta'] !== undefined ) {
        console.log("Global State updated:",transactionResponse['global-state-delta']);
    }
    if (transactionResponse['local-state-delta'] !== undefined ) {
        console.log("Local State updated:",transactionResponse['local-state-delta']);
    }
    ```
    <!-- ===JSSDK_APP_NOOP=== -->

=== "Java"
    <!-- ===JAVASDK_APP_NOOP=== -->
```java
        Transaction noopTxn = ApplicationCallTransactionBuilder.Builder()
                .sender(user.getAddress())
                .suggestedParams(sp)
                .applicationId(appId)
                .build();

        SignedTransaction signedNoop = user.signTransaction(noopTxn);
        Response<PostTransactionsResponse> noopResponse = algodClient.RawTransaction()
                .rawtxn(Encoder.encodeToMsgPack(signedNoop)).execute();

        PendingTransactionResponse noopResult = Utils.waitForConfirmation(algodClient, noopResponse.body().txId, 4);
        assert noopResult.confirmedRound > 0;
```
[Snippet Source](https://github.com/barnjamin/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/AppExamples.java#L102-L114)
    <!-- ===JAVASDK_APP_NOOP=== -->

=== "Go"
    <!-- ===GOSDK_APP_NOOP=== -->
	```go
    // create unsigned transaction
    txn, _:= transaction.MakeApplicationNoOpTx(index, appArgs, nil, nil, nil, params, sender, 
                            nil, types.Digest{}, [32]byte{}, types.Address{})

    // sign, send, await 

    // display results
    confirmedTxn, _, err := client.PendingTransactionInformation(txID).Do(context.Background())
    fmt.Printf("Called app-id: %d\n", confirmedTxn.Transaction.Txn.ApplicationID)
    ```
    <!-- ===GOSDK_APP_NOOP=== -->

## Read state

Anyone may read the [global state](../apps/#reading-global-state-from-other-smart-contracts) of any application or the [local state](../apps/#reading-local-state-from-other-accounts) of an application within a given user account using the REST API account_info endpoint. 

=== "Python"
    <!-- ===PYSDK_APP_READ_STATE=== -->
```python
acct_info = algod_client.account_application_info(user.address, app_id)
# base64 encoded keys and values
print(acct_info["app-local-state"]["key-value"])
```
[Snippet Source](https://github.com/barnjamin/py-algorand-sdk/blob/doc-examples/_examples/apps.py#L75-L78)
    <!-- ===PYSDK_APP_READ_STATE=== -->

=== "JavaScript"
    <!-- ===JSSDK_APP_READ_STATE=== -->
	```javascript
    // read local state of application from user account
    async function readLocalState(client, account, index){
        let accountInfoResponse = await client.accountInformation(account.addr).do();
        for (let i = 0; i < accountInfoResponse['apps-local-state'].length; i++) { 
            if (accountInfoResponse['apps-local-state'][i].id == index) {
                console.log("User's local state:");
                for (let n = 0; n < accountInfoResponse['apps-local-state'][i][`key-value`].length; n++) {
                    console.log(accountInfoResponse['apps-local-state'][i][`key-value`][n]);
                }
            }
        }
    }

    // read global state of application
    async function readGlobalState(client, index){
        let applicationInfoResponse = await client.applicationInfo(index).do();
        let globalState = []
        if(applicationInfoResponse['params'].includes('global-state')) {
            globalState = applicationInfoResponse['params']['global-state']
        }
        for (let n = 0; n < globalState.length; n++) {
            console.log(applicationInfoResponse['params']['global-state'][n]);
        }
    }
    ```
    <!-- ===JSSDK_APP_READ_STATE=== -->

=== "Java"
    <!-- ===JAVASDK_APP_READ_STATE=== -->
```java
```
[Snippet Source](https://github.com/barnjamin/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/AppExamples.java#L117-L117)
    <!-- ===JAVASDK_APP_READ_STATE=== -->

=== "Go"
    <!-- ===GOSDK_APP_READ_STATE=== -->
	```go
    func readLocalState(client *algod.Client, account crypto.Account, index uint64) {
        accountInfo, _ := client.AccountInformation(account.Address.String()).Do(context.Background())
        for _, ap := range accountInfo.AppsLocalState {
            if ap.Id == index {
                fmt.Printf("Local state for app-id %d (account %s):\n", ap.Id, account.Address.String())
                fmt.Println(ap.KeyValue)
            }
        }
    }

    func readGlobalState(client *algod.Client, index uint64) {
        applicationInfo, _ := client.GetApplicationByID(index)
        globalState, _ := applicationInfo.Params.GlobalState
        if globalState != nil {
            fmt.Println(globalState)
        }
    }
    ```
    <!-- ===GOSDK_APP_READ_STATE=== -->

## Update

The creator may [update the approval program](../apps/#update-stateful-smart-contract) using the update method (if the current approval program allows it). The refactored approval program source code adds a key/value pair to the user's local storage indicating the timestamp when the application was called. Refer to the [appendix](#refactored-approval-program) for details. The original clear program will be reused.

The update method method requires 6 parameters:

- sender: address, representing the user intending to opt-in to using the app
- sp: suggested parameters obtained from the network
- index: the app-id as defined by the create method result
- approval_program: compiled program
- clear program: compiled program
- app_args: application arguments used by approval program

Construct the update transaction and await the response:

=== "Python"
    <!-- ===PYSDK_APP_UPDATE=== -->
```python
with open("application/approval_refactored.teal", "r") as f:
    approval_program = f.read()

approval_result = algod_client.compile(approval_program)
approval_binary = base64.b64decode(approval_result["result"])


sp = algod_client.suggested_params()
# create the app update transaction, passing compiled programs and schema
# note that schema is immutable, we cant change it after create
app_update_txn = transaction.ApplicationUpdateTxn(
    creator.address,
    sp,
    app_id,
    approval_program=approval_binary,
    clear_program=clear_binary,
)
signed_update = app_update_txn.sign(creator.private_key)
txid = algod_client.send_transaction(signed_update)
update_result = transaction.wait_for_confirmation(algod_client, txid, 4)
assert update_result["confirmed-round"] > 0
```
[Snippet Source](https://github.com/barnjamin/py-algorand-sdk/blob/doc-examples/_examples/apps.py#L81-L102)
    <!-- ===PYSDK_APP_UPDATE=== -->

=== "JavaScript"
    <!-- ===JSSDK_APP_UPDATE=== -->
	```javascript
    // create unsigned transaction
    let txn = algosdk.makeApplicationUpdateTxn(sender, params, index, approvalProgram, clearProgram);

    // sign, send, await

    // display results
    let transactionResponse = await client.pendingTransactionInformation(txId).do();
    let appId = transactionResponse['txn']['txn'].apid;
    console.log("Updated app-id: ",appId);
    ```
    <!-- ===JSSDK_APP_UPDATE=== -->

=== "Java"
    <!-- ===JAVASDK_APP_UPDATE=== -->
```java
        String approvalSourceUpdated = Files.readString(Paths.get("application/approval_refactored.teal"));
        CompileResponse approvalUpdatedResponse = algodClient.TealCompile().source(approvalSourceUpdated.getBytes())
                .execute()
                .body();
        TEALProgram approvalProgUpdated = new TEALProgram(approvalUpdatedResponse.result);

        Transaction appUpdate = ApplicationUpdateTransactionBuilder.Builder()
                .sender(creator.getAddress())
                .suggestedParams(sp)
                .applicationId(appId)
                .approvalProgram(approvalProgUpdated)
                .clearStateProgram(clearProg)
                .build();

        SignedTransaction signedAppUpdate = creator.signTransaction(appUpdate);
        Response<PostTransactionsResponse> updateResponse = algodClient.RawTransaction()
                .rawtxn(Encoder.encodeToMsgPack(signedAppUpdate)).execute();
        PendingTransactionResponse updateResult = Utils.waitForConfirmation(algodClient, updateResponse.body().txId, 4);
        assert updateResult.confirmedRound > 0;
```
[Snippet Source](https://github.com/barnjamin/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/AppExamples.java#L120-L139)
    <!-- ===JAVASDK_APP_UPDATE=== -->

=== "Go"
    <!-- ===GOSDK_APP_UPDATE=== -->
	```go
        // create unsigned transaction
        txn, _ := transaction.MakeApplicationUpdateTx(index, nil, nil, nil, nil, 
                            approvalProgram, clearProgram, params, creatorAccount.Address, 
                            nil, types.Digest{}, [32]byte{}, types.Address{})

        // sign, send, await
        
        // display results
        confirmedTxn, _, _ := client.PendingTransactionInformation(txID).Do(context.Background())
        fmt.Printf("Updated app-id: %d\n", confirmedTxn.Transaction.Txn.ApplicationID)
    }
    ```
    <!-- ===GOSDK_APP_UPDATE=== -->

## Call with arguments

A program may [process arguments passed](../apps/#passing-arguments-to-stateful-smart-contracts) at run-time. The NoOp call method has an optional app_args parameter where the timestamp may be supplied:

The refactored application expects a timestamp be supplied with the application call.

=== "Python"
    <!-- ===PYSDK_APP_CALL=== -->
```python
now = datetime.datetime.now().strftime("%H:%M:%S")
app_args = [now.encode("utf-8")]
call_txn = transaction.ApplicationNoOpTxn(user.address, sp, app_id, app_args)

signed_call = call_txn.sign(user.private_key)
txid = algod_client.send_transaction(signed_call)
call_result = transaction.wait_for_confirmation(algod_client, txid, 4)
assert call_result["confirmed-round"] > 0

# display results
print("Called app-id: ", call_result["txn"]["txn"]["apid"])
if "global-state-delta" in call_result:
    print("Global State updated :\n", call_result["global-state-delta"])
if "local-state-delta" in call_result:
    print("Local State updated :\n", call_result["local-state-delta"])
```
[Snippet Source](https://github.com/barnjamin/py-algorand-sdk/blob/doc-examples/_examples/apps.py#L105-L120)
    <!-- ===PYSDK_APP_CALL=== -->

=== "JavaScript"
    <!-- ===JSSDK_APP_CALL=== -->
	```javascript
    // call application with arguments
    let ts = new Date(new Date().toUTCString());
    console.log(ts)
    let appArgs = [];
    appArgs.push(new Uint8Array(Buffer.from(ts)));

    // create unsigned transaction
    let txn = algosdk.makeApplicationNoOpTxn(sender, params, index, appArgs)

    // sign, send, await

    // display results
    let transactionResponse = await client.pendingTransactionInformation(txId).do();
    console.log("Called app-id:",transactionResponse['txn']['txn']['apid'])
    if (transactionResponse['global-state-delta'] !== undefined ) {
        console.log("Global State updated:",transactionResponse['global-state-delta']);
    }
    if (transactionResponse['local-state-delta'] !== undefined ) {
        console.log("Local State updated:",transactionResponse['local-state-delta']);
    }
    ```
    <!-- ===JSSDK_APP_CALL=== -->

=== "Java"
    <!-- ===JAVASDK_APP_CALL=== -->
```java
        SimpleDateFormat formatter = new SimpleDateFormat("yyyy-MM-dd 'at' HH:mm:ss");
        Date date = new Date(System.currentTimeMillis());

        List<byte[]> appArgs = new ArrayList<byte[]>();
        appArgs.add(formatter.format(date).toString().getBytes());

        // create unsigned transaction
        Transaction callTransaction = ApplicationCallTransactionBuilder.Builder()
                .sender(user.getAddress())
                .suggestedParams(sp)
                .applicationId(appId)
                .args(appArgs)
                .build();

        SignedTransaction signedCallTransaction = user.signTransaction(callTransaction);
        Response<PostTransactionsResponse> callResponse = algodClient.RawTransaction()
                .rawtxn(Encoder.encodeToMsgPack(signedCallTransaction)).execute();

        PendingTransactionResponse callResult = Utils.waitForConfirmation(algodClient, callResponse.body().txId, 4);
        assert callResult.confirmedRound > 0;
        // display results
        if (callResult.globalStateDelta != null) {
            System.out.printf("\tGlobal state: %s\n", callResult.globalStateDelta);
        }

        if (callResult.localStateDelta != null) {
            System.out.printf("\tLocal state: %s\n", callResult.localStateDelta);
        }
```
[Snippet Source](https://github.com/barnjamin/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/AppExamples.java#L142-L170)
    <!-- ===JAVASDK_APP_CALL=== -->

=== "Go"
    <!-- ===GOSDK_APP_CALL=== -->
	```go
    // call application with arguments
    now := time.Now().Format("Mon Jan _2 15:04:05 2006")
    appArgs := make([][]byte, 1)
    appArgs[0] = []byte(now)

    // create unsigned transaction
    txn, _ := transaction.MakeApplicationNoOpTx(index, appArgs, nil, nil, nil, params, sender, 
                                nil, types.Digest{}, [32]byte{}, types.Address{})

    // sign, send, await

    // display results
    confirmedTxn, _, _ := client.PendingTransactionInformation(txID).Do(context.Background())
    fmt.Printf("Called app-id: %d\n", confirmedTxn.Transaction.Txn.ApplicationID)
    ```
    <!-- ===GOSDK_APP_CALL=== -->

## Close out

The user may discontinue use of the application by sending a [close out](../apps/#the-lifecycle-of-a-stateful-smart-contract) transaction. This will remove the local state for this application from the user's account. This method requires 3 parameters:

- sender: address, representing the user intending to optin to using the app
- sp: suggested parameters obtained from the network
- index: the app-id as defined by the create method result

=== "Python"
    <!-- ===PYSDK_APP_CLOSEOUT=== -->
```python
close_txn = transaction.ApplicationCloseOutTxn(user.address, sp, app_id)
signed_close = close_txn.sign(user.private_key)
txid = algod_client.send_transaction(signed_close)
optin_result = transaction.wait_for_confirmation(algod_client, txid, 4)
assert optin_result["confirmed-round"] > 0
```
[Snippet Source](https://github.com/barnjamin/py-algorand-sdk/blob/doc-examples/_examples/apps.py#L123-L128)
    <!-- ===PYSDK_APP_CLOSEOUT=== -->

=== "JavaScript"
    <!-- ===JSSDK_APP_CLOSEOUT=== -->
	```javascript
    // create unsigned transaction
    let txn = algosdk.makeApplicationCloseOutTxn(sender, params, index)

    // sign, send, await

    // display results
    let transactionResponse = await client.pendingTransactionInformation(txId).do();
    console.log("Closed out from app-id:",transactionResponse['txn']['txn']['apid'])
    ```
    <!-- ===JSSDK_APP_CLOSEOUT=== -->

=== "Java"
    <!-- ===JAVASDK_APP_CLOSEOUT=== -->
```java
        Transaction closeOutTxn = ApplicationCloseTransactionBuilder.Builder()
                .sender(user.getAddress())
                .suggestedParams(sp)
                .applicationId(appId)
                .build();

        SignedTransaction signedCloseOut = user.signTransaction(closeOutTxn);
        Response<PostTransactionsResponse> closeOutResponse = algodClient.RawTransaction()
                .rawtxn(Encoder.encodeToMsgPack(signedCloseOut)).execute();

        PendingTransactionResponse closeOutResult = Utils.waitForConfirmation(algodClient, closeOutResponse.body().txId,
                4);
        assert closeOutResult.confirmedRound > 0;
```
[Snippet Source](https://github.com/barnjamin/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/AppExamples.java#L173-L186)
    <!-- ===JAVASDK_APP_CLOSEOUT=== -->

=== "Go"
    <!-- ===GOSDK_APP_CLOSEOUT=== -->
	```go
    // create unsigned transaction
    txn, _ := transaction.MakeApplicationCloseOutTx(index, nil, nil, nil, nil, params, account.Address, nil, types.Digest{}, [32]byte{}, types.Address{})

    // sign, send, await

    // display results
    confirmedTxn, _, _ := client.PendingTransactionInformation(txID).Do(context.Background())
    fmt.Printf("Closed out from app-id: %d\n", confirmedTxn.Transaction.Txn.ApplicationID)
    ```
    <!-- ===GOSDK_APP_CLOSEOUT=== -->

## Delete

The approval program defines the creator as the only account able to [delete the application](../apps/#delete-stateful-smart-contract). This removes the global state, but does not impact any user's local state. This method uses the same 3 parameters.

=== "Python"
    <!-- ===PYSDK_APP_DELETE=== -->
```python
delete_txn = transaction.ApplicationDeleteTxn(creator.address, sp, app_id)
signed_delete = delete_txn.sign(creator.private_key)
txid = algod_client.send_transaction(signed_delete)
optin_result = transaction.wait_for_confirmation(algod_client, txid, 4)
assert optin_result["confirmed-round"] > 0
```
[Snippet Source](https://github.com/barnjamin/py-algorand-sdk/blob/doc-examples/_examples/apps.py#L131-L136)
    <!-- ===PYSDK_APP_DELETE=== -->

=== "JavaScript"
    <!-- ===JSSDK_APP_DELETE=== -->
	```javascript
    // create unsigned transaction
    let txn = algosdk.makeApplicationDeleteTxn(sender, params, index);

    // sign, send, await

    // display results
    let transactionResponse = await client.pendingTransactionInformation(txId).do();
    let appId = transactionResponse['txn']['txn'].apid;
    console.log("Deleted app-id: ",appId);
    ```
    <!-- ===JSSDK_APP_DELETE=== -->

=== "Java"
    <!-- ===JAVASDK_APP_DELETE=== -->
```java
        Transaction appDelete = ApplicationDeleteTransactionBuilder.Builder()
                .sender(creator.getAddress())
                .suggestedParams(sp)
                .applicationId(appId)
                .build();

        SignedTransaction signedAppDelete = creator.signTransaction(appDelete);
        Response<PostTransactionsResponse> deleteResponse = algodClient.RawTransaction()
                .rawtxn(Encoder.encodeToMsgPack(signedAppDelete)).execute();
        PendingTransactionResponse deleteResult = Utils.waitForConfirmation(algodClient, deleteResponse.body().txId, 4);
        assert deleteResult.confirmedRound > 0;
```
[Snippet Source](https://github.com/barnjamin/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/AppExamples.java#L189-L200)
    <!-- ===JAVASDK_APP_DELETE=== -->

=== "Go"
    <!-- ===GOSDK_APP_DELETE=== -->
	```go
    // create unsigned transaction
    txn, _ := transaction.MakeApplicationDeleteTx(index, nil, nil, nil, nil, params, sender, 
                            nil, types.Digest{}, [32]byte{}, types.Address{})

    // sign, send, await

    // display results
    confirmedTxn, _, _ := client.PendingTransactionInformation(txID).Do(context.Background())

    fmt.Printf("Deleted app-id: %d\n", confirmedTxn.Transaction.Txn.ApplicationID)
    ```
    <!-- ===GOSDK_APP_DELETE=== -->

## Clear state

The user may [clear the local state](../apps/#the-lifecycle-of-a-stateful-smart-contract) for an application at any time, even if the application was deleted by the creator. This method uses the same 3 parameter.

=== "Python"
    <!-- ===PYSDK_APP_CLEAR=== -->
```python
clear_txn = transaction.ApplicationClearStateTxn(user.address, sp, app_id)
# .. sign, send, wait
```
[Snippet Source](https://github.com/barnjamin/py-algorand-sdk/blob/doc-examples/_examples/apps.py#L139-L141)
    <!-- ===PYSDK_APP_CLEAR=== -->

=== "JavaScript"
    <!-- ===JSSDK_APP_CLEAR=== -->
	```javascript
    // create unsigned transaction
    let txn = algosdk.makeApplicationClearStateTxn(sender, params, index);

    // sign, send, await

    // display results
    let transactionResponse = await client.pendingTransactionInformation(txId).do();
    let appId = transactionResponse['txn']['txn'].apid;
    console.log("Cleared local state for app-id: ",appId);
    ```
    <!-- ===JSSDK_APP_CLEAR=== -->

=== "Java"
    <!-- ===JAVASDK_APP_CLEAR=== -->
```java
        Transaction clearTxn = ApplicationClearTransactionBuilder.Builder()
                .sender(user.getAddress())
                .suggestedParams(sp)
                .applicationId(appId)
                .build();

        SignedTransaction signedClear = user.signTransaction(clearTxn);
        // ... sign, send, wait
```
[Snippet Source](https://github.com/barnjamin/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/AppExamples.java#L203-L211)
    <!-- ===JAVASDK_APP_CLEAR=== -->

=== "Go"
    <!-- ===GOSDK_APP_CLEAR=== -->
	```go
    // create unsigned transaction
    txn, _ := transaction.MakeApplicationClearStateTx(index, nil, nil, nil, nil, params, sender, 
                            nil, types.Digest{}, [32]byte{}, types.Address{})

    // sign, send, await

    // display results
    confirmedTxn, _, _ := client.PendingTransactionInformation(txID).Do(context.Background())
    fmt.Printf("Cleared local state for app-id: %d\n", confirmedTxn.Transaction.Txn.ApplicationID)
    ```
    <!-- ===GOSDK_APP_CLEAR=== -->

# Appendix

## Approval program walkthrough

```teal
#pragma version 4
// Handle each possible OnCompletion type. We don't have to worry about
// handling ClearState, because the ClearStateProgram will execute in that
// case, not the ApprovalProgram.

txn OnCompletion
int NoOp
==
bnz handle_noop

txn OnCompletion
int OptIn
==
bnz handle_optin

txn OnCompletion
int CloseOut
==
bnz handle_closeout

txn OnCompletion
int UpdateApplication
==
bnz handle_updateapp

txn OnCompletion
int DeleteApplication
==
bnz handle_deleteapp

// Unexpected OnCompletion value. Should be unreachable.
err

handle_noop:
// Handle NoOp
// Check for creator
addr 5XWY6RBNYHCSY2HK5HCTO62DUJJ4PT3G4L77FQEBUKE6ZYRGQAFTLZSQQ4
txn Sender
==
bnz handle_optin

// read global state
byte "counter"
dup
app_global_get

// increment the value
int 1
+

// store to scratch space
dup
store 0

// update global state
app_global_put

// read local state for sender
int 0
byte "counter"
app_local_get

// increment the value
int 1
+
store 1

// update local state for sender
int 0
byte "counter"
load 1
app_local_put

// load return value as approval
load 0
return

handle_optin:
// Handle OptIn
// approval
int 1
return

handle_closeout:
// Handle CloseOut
//approval
int 1
return

handle_deleteapp:
// Check for creator
addr 5XWY6RBNYHCSY2HK5HCTO62DUJJ4PT3G4L77FQEBUKE6ZYRGQAFTLZSQQ4
txn Sender
==
return

handle_updateapp:
// Check for creator
addr 5XWY6RBNYHCSY2HK5HCTO62DUJJ4PT3G4L77FQEBUKE6ZYRGQAFTLZSQQ4
txn Sender
==
return
```

## Clear program walkthrough

## Refactored approval program

```teal
#pragma version 4
// Handle each possible OnCompletion type. We don't have to worry about
// handling ClearState, because the ClearStateProgram will execute in that
// case, not the ApprovalProgram.

txn OnCompletion
int NoOp
==
bnz handle_noop

txn OnCompletion
int OptIn
==
bnz handle_optin

txn OnCompletion
int CloseOut
==
bnz handle_closeout

txn OnCompletion
int UpdateApplication
==
bnz handle_updateapp

txn OnCompletion
int DeleteApplication
==
bnz handle_deleteapp

// Unexpected OnCompletion value. Should be unreachable.
err

handle_noop:
// Handle NoOp
// Check for creator
addr 5XWY6RBNYHCSY2HK5HCTO62DUJJ4PT3G4L77FQEBUKE6ZYRGQAFTLZSQQ4
txn Sender
==
bnz handle_optin

// read global state
byte "counter"
dup
app_global_get

// increment the value
int 1
+

// store to scratch space
dup
store 0

// update global state
app_global_put

// read local state for sender
int 0
byte "counter"
app_local_get

// increment the value
int 1
+
store 1

// update local state for sender
// update "counter"
int 0
byte "counter"
load 1
app_local_put

// update "timestamp"
int 0
byte "timestamp"
txn ApplicationArgs 0
app_local_put

// load return value as approval
load 0
return

handle_optin:
// Handle OptIn
// approval
int 1
return

handle_closeout:
// Handle CloseOut
//approval
int 1
return

handle_deleteapp:
// Check for creator
addr 5XWY6RBNYHCSY2HK5HCTO62DUJJ4PT3G4L77FQEBUKE6ZYRGQAFTLZSQQ4
txn Sender
==
return

handle_updateapp:
// Check for creator
addr 5XWY6RBNYHCSY2HK5HCTO62DUJJ4PT3G4L77FQEBUKE6ZYRGQAFTLZSQQ4
txn Sender
==
return
```