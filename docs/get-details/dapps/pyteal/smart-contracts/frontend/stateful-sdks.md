title: Interact with stateful contracts

This guide covers using stateful smart contracts with the Algorand SDKs. Stateful smart contracts form the basis for applications written in [Transaction Execution Approval Language (TEAL)](../../../../avm/teal/) or with Python using the [PyTeal](../../../../pyteal/) library.

Each SDK's install process is discussed in the [SDK Reference](../../../reference/sdks/index.md) documentation.

!!! info
    The example code snippets are provided throughout this page and are abbreviated for conciseness and clarity. Full running code examples for each SDK are available within the GitHub repo at [/examples/smart_contracts](https://github.com/algorand/docs/tree/master/examples/smart_contracts) and for [download](https://github.com/algorand/docs/blob/master/examples/smart_contracts/smart_contracts.zip?raw=true) (.zip).

# Application Lifecycle

This guide follows an application throughout its [lifecycle](../../stateful/#the-lifecycle-of-a-stateful-smart-contract) from initial creation, to usage, to modification and finally deletion. The application is a variation of the [Hello World](hello_world.md) counter application. Initially, the application stores the number of times called within its _global state_ and also stores the number of times each user account calls the application within their _local state_. Midway through the lifecycle, the application is upgraded to add an additional key:value pair to the user's _local storage_ for storing the call timestamp. 

# Environment Setup

This guide requires two accounts:

```python tab="Python"
# user declared account mnemonics
creator_mnemonic = "Your first 25-word mnemonic goes here"
user_mnemonic = "A second distinct 25-word mnemonic goes here"
```

```javascript tab="JavaScript"
// user declared account mnemonics
creatorMnemonic = "Your 25-word mnemonic goes here";
userMnemonic = "A second distinct 25-word mnemonic goes here";
```

```Java tab="Java"
// user declared account mnemonics
String creatorMnemonic = "Your 25-word mnemonic goes here";
String userMnemonic = "A second distinct 25-word mnemonic goes here";
```

```go tab="Go"
// user defined mnemonics
const creatorMnemonic = "Your 25-word mnemonic goes here"
const userMnemonic = "A second distinct 25-word mnemonic goes here"
```

An `algod` client connection is also required. The following connects using [Sandbox](https://developer.algorand.org/articles/introducing-sandbox-quick-way-get-started-algorand/):

```python tab="Python"
# user declared algod connection parameters
algod_address = "http://localhost:4001"
algod_token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
algod_client = algod.AlgodClient(algod_token, algod_address)
```

```javascript tab="JavaScript"
// user declared algod connection parameters
algodAddress = "http://localhost:4001";
algodToken = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
let algodClient = new algosdk.Algodv2(algodToken, algodServer, algodPort);
```

```Java tab="Java"
// user declared account mnemonics
String ALGOD_API_ADDR = "localhost";
Integer ALGOD_PORT = 4001;
String ALGOD_API_TOKEN = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
AlgodClient client = (AlgodClient) new AlgodClient(ALGOD_API_ADDR, ALGOD_PORT, ALGOD_API_TOKEN);
```

```go tab="Go"
// user defined algod client settings
const algodAddress = "http://localhost:8080"
const algodToken = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
algodClient, err := algod.MakeClient(algodAddress, algodToken)
```

!!! Info
    Ensure the `algod` node has the _"EnableDeveloperAPI"_ parameter set to **true** within the `config.json` file. This is required to enable the SDK access to the _compile_ and _dryrun_ endpoints.

# Declarations

All stateful applications are comprised of state storage, an approval program and a clear program. Details of each are found within the [stateful smart contract guide](../../stateful).

## State Storage
Begin by defining the application's _global_schema_ and _local_schema_ storage requirements. These values are immutable once the application is created, so they must specify the maximum number required by the initial application and any future updates. 

The example application defined below may hold up to one each of `bytes` and `ints` value within the _local storage_ of the user account, as well as a single `ints` value within _global storage_ of the application:

```python tab="Python"
# declare application state storage (immutable)
local_ints = 1
local_bytes = 1
global_ints = 1
global_bytes = 0

# define schema
global_schema = transaction.StateSchema(global_ints, global_bytes)
local_schema = transaction.StateSchema(local_ints, local_bytes)
```

```javascript tab="JavaScript"
// declare application state storage (immutable)
localInts = 1;
localBytes = 1;
globalInts = 1;
globalBytes = 0;
```

```Java tab="Java"
// declare application state storage (immutable)
int localInts = 1;
int localBytes = 1;
int globalInts = 1;
int globalBytes = 0;
```

```go tab="Go"
// declare application state storage (immutable)
const localInts = 1
const localBytes = 1
const globalInts = 1
const globalBytes = 0

// define schema
globalSchema := types.StateSchema{NumUint: uint64(globalInts), NumByteSlice: uint64(globalBytes)}
localSchema := types.StateSchema{NumUint: uint64(localInts), NumByteSlice: uint64(localBytes)}
```

!!! Info
    The example application is not allowed to hold any `bytes` value within global storage.


## Approval Program

The [approval program](../../stateful/#the-lifecycle-of-a-stateful-smart-contract) handles the main logic of the application. A detailed walkthrough of this code is provided in the [appendix](#appendix) of this guide.

## Clear Program

This is the most basic [clear program](../../stateful/#the-lifecycle-of-a-stateful-smart-contract) and returns _true_ when an account clears its participation in a smart contract:

```python tab="Python"
# declare clear state program source
clear_program_source = b"""#pragma version 4
int 1
"""
```

```javascript tab="JavaScript"
// declare clear state program source
clearProgramSource = `#pragma version 4
int 1
`;
```

```Java tab="Java"
// declare clear state program source
String clearProgramSource = "#pragma version 4\n" +
"int 1\n";
```

```go tab="Go"
// declare clear state program source
const clearProgramSource = `#pragma version 4
int 1
`
```

# Application Methods

## Create

The creator will deploy the application using the [create app](../../stateful#creating-the-smart-contract) method. It requires 7 parameters:

- sender: address, representing the creator of the app
- sp: suggested parameters obtained from the network
- on_complete: enum value, representing NoOp
- approval_program: compiled program
- clear program: compiled program
- local_schema: maximum _local storage_ allocation, immutable
- global_schema: maximum _global storage_ allocation, immutable

Use the creator_mnemonic to define sender:

```python tab="Python"
# get account from mnemonic
private_key = mnemonic.to_private_key(creator_mnemonic)
sender = account.address_from_private_key(private_key)
```

```javascript tab="JavaScript"
// get account from mnemonic
let creatorAccount = algosdk.mnemonicToSecretKey(creatorMnemonic);
let sender = creatorAccount.addr;
```

```Java tab="Java"
// get account from mnemonic
Account creatorAccount = new Account(creatorMnemonic);
Address sender = creatorAccount.getAddress();
```

```go tab="Go"
// get account from mnemonic
creatorAccount := recoverAccount(creatorMnemonic)
sender := creatorAccount.Address
```

Use the `suggested_params` endpoint:

```python tab="Python"
# get node suggested parameters
params = client.suggested_params()
# comment out the next two (2) lines to use suggested fees
params.flat_fee = True
params.fee = 1000
```

```javascript tab="JavaScript"
// get node suggested parameters
let params = await client.getTransactionParams().do();
// comment out the next two lines to use suggested fee
params.fee = 1000;
params.flatFee = true;
```

```Java tab="Java"
// get node suggested parameters
TransactionParametersResponse params = client.TransactionParams().execute().body();
```

```go tab="Go"
// get transaction suggested parameters
params, _ := client.SuggestedParams().Do(context.Background())
// comment out the next two (2) lines to use suggested fees
params.FlatFee = true
params.Fee = 1000
```

Set the [on_complete](../../teal/specification/#oncomplete) parameter to NoOp: 

```python tab="Python"
# declare on_complete as NoOp
on_complete = transaction.OnComplete.NoOpOC.real
```

```javascript tab="JavaScript"
// declare onComplete as NoOp
onComplete = algosdk.OnApplicationComplete.NoOpOC;
```

```Java tab="Java"
// not required
```

```go tab="Go"
// not required
```

Compile the programs using the `compile` endpoint:

```python tab="Python"
# helper function to compile program source
def compile_program(client, source_code) :
    compile_response = client.compile(source_code.decode('utf-8'))
    return base64.b64decode(compile_response['result'])
```

```javascript tab="JavaScript"
// helper function to compile program source  
async function compileProgram(client, programSource) {
    let encoder = new TextEncoder();
    let programBytes = encoder.encode(programSource);
    let compileResponse = await client.compile(programBytes).do();
    let compiledBytes = new Uint8Array(Buffer.from(compileResponse.result, "base64"));
    return compiledBytes;
}
```

```Java tab="Java"
// helper function to compile program source
public String compileProgram(AlgodClient client, byte[] programSource) {
    Response<CompileResponse> compileResponse = null;
    try {
        compileResponse = client.TealCompile().source(programSource).execute();
    } catch (Exception e) {
        e.printStackTrace();
    }
    System.out.println(compileResponse.body().result);
    return compileResponse.body().result;
}
```

```go tab="Go"
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

Construct the transaction with defined values:

```python tab="Python"
# create unsigned transaction
txn = transaction.ApplicationCreateTxn(sender, params, on_complete, \
                                        approval_program, clear_program, \
                                        global_schema, local_schema)
```

```javascript tab="JavaScript"
// create unsigned transaction
let txn = algosdk.makeApplicationCreateTxn(sender, params, onComplete, 
                                        approvalProgram, clearProgram, 
                                        localInts, localBytes, globalInts, globalBytes,);
let txId = txn.txID().toString();
```

```Java tab="Java"
// create unsigned transaction
Transaction txn = Transaction.ApplicationCreateTransactionBuilder()
                    .sender(sender)
                    .suggestedParams(params)
                    .approvalProgram(approvalProgramSource)
                    .clearStateProgram(clearProgramSource)
                    .globalStateSchema(new StateSchema(globalInts, globalBytes))
                    .localStateSchema(new StateSchema(localInts, localBytes))
                    .build();
```

```go tab="Go"
// create unsigned transaction
txn, _ := future.MakeApplicationCreateTx(false, approvalProgram, clearProgram, globalSchema, localSchema, 
                    nil, nil, nil, nil, params, creatorAccount.Address, nil, 
                    types.Digest{}, [32]byte{}, types.Address{}, uint32(0))
```

Sign, send, await confirmation and display the results:

```python tab="Python"
# sign transaction
signed_txn = txn.sign(private_key)
tx_id = signed_txn.transaction.get_txid()

# send transaction
client.send_transactions([signed_txn])

# await confirmation
wait_for_confirmation(client, tx_id)

# display results
transaction_response = client.pending_transaction_info(tx_id)
app_id = transaction_response['application-index']
print("Created new app-id: ",app_id)
```

```javascript tab="JavaScript"
// Sign the transaction
let signedTxn = txn.signTxn(creatorAccount.sk);
console.log("Signed transaction with txID: %s", txId);

// Submit the transaction
await client.sendRawTransaction(signedTxn).do();

// Wait for confirmation
await waitForConfirmation(client, txId);

// display results
let transactionResponse = await client.pendingTransactionInformation(txId).do();
let appId = transactionResponse['application-index'];
console.log("Created new app-id: ",appId);
```

```Java tab="Java"
// sign transaction
SignedTransaction signedTxn = creator.signTransaction(txn);
System.out.println("Signed transaction with txid: " + signedTxn.transactionID);

// send to network
byte[] encodedTxBytes = Encoder.encodeToMsgPack(signedTxn);
String id = client.RawTransaction().rawtxn(encodedTxBytes).execute().body().txId;
System.out.println("Successfully sent tx with ID: " + id);

// await confirmation
waitForConfirmation(id);

// display results
PendingTransactionResponse pTrx = client.PendingTransactionInformation(id).execute().body();
Long appId = pTrx.applicationIndex;
System.out.println("Created new app-id: " + appId);    
```

```go tab="Go"
// Sign the transaction
txID, signedTxn, _ := crypto.SignTransaction(creatorAccount.PrivateKey, txn)
fmt.Printf("Signed txid: %s\n", txID)

// Submit the transaction
sendResponse, _ := client.SendRawTransaction(signedTxn).Do(context.Background())
fmt.Printf("Submitted transaction %s\n", sendResponse)

// Wait for confirmation
waitForConfirmation(txID, client)

// display results
confirmedTxn, _, _ := client.PendingTransactionInformation(txID).Do(context.Background())
appId = confirmedTxn.ApplicationIndex
fmt.Printf("Created new app-id: %d\n", appId)
```

!!! Notice
    Note down the app-id from the confirmed transaction response. Place this value into the `index` parameter within all remaining code samples. 

## Opt-In

The user must [opt-in](../../stateful/#opt-in-to-the-smart-contract) to use the application. This method requires 3 parameters:

- sender: address, representing the user intending to opt-in to using the app
- sp: suggested parameters obtained from the network
- index: the app-id as defined by the create method result

Use the user_mnemonic to define sender:

```python tab="Python"
# declare sender
private_key = mnemonic.to_private_key(user_mnemonic)
sender = account.address_from_private_key(private_key)
```

```javascript tab="JavaScript"
// get accounts from mnemonic
let userAccount = algosdk.mnemonicToSecretKey(userMnemonic);
let sender = userAccount.addr;
```

```Java tab="Java"
// declare sender
Account userAccount = new Account(userMnemonic);
Address sender = userAccount.getAddress();
```

```go tab="Go"
// declare sender
userAccount := recoverAccount(userMnemonic)
sender := userAccount.Address
```

Construct the transaction with defined values:

```python tab="Python"
txn = transaction.ApplicationOptInTxn(sender, params, index)
```

```javascript tab="JavaScript"
// create unsigned transaction
let txn = algosdk.makeApplicationOptInTxn(sender, params, index);
```

```Java tab="Java"
// create unsigned transaction
Transaction txn = Transaction.ApplicationOptInTransactionBuilder()
                        .sender(sender)
                        .suggestedParams(params)
                        .applicationId(appId)
                        .build();
```

```go tab="Go"
// create unsigned transaction
txn, _ := future.MakeApplicationOptInTx(index, nil, nil, nil, nil, params,
                        sender, nil, types.Digest{}, [32]byte{}, types.Address{})
```

Sign, send, await confirmation and display the results:

```python tab="Python"
# sign, send, await

# display results
transaction_response = client.pending_transaction_info(tx_id)
print("OptIn to app-id: ",transaction_response['txn']['txn']['apid'])  
```  

```javascript tab="JavaScript"
// sign, send, await

// display results
let transactionResponse = await client.pendingTransactionInformation(txId).do();
console.log("Opted-in to app-id:",transactionResponse['txn']['txn']['apid'])
```

```Java tab="Java"
// sign, send, await

// display results
PendingTransactionResponse pTrx = client.PendingTransactionInformation(id).execute().body();
System.out.println("OptIn to app-id: " + pTrx.txn.tx.applicationId);       
```

```go tab="Go"
// sign, send, await

// display results
confirmedTxn, _, _ := client.PendingTransactionInformation(txID).Do(context.Background())
fmt.Printf("Oped-in to app-id: %d\n", confirmedTxn.Transaction.Txn.ApplicationID)
```

## Call (NoOp)

The user may now [call](../../stateful/#call-the-stateful-smart-contract) the application. This method requires 3 parameters:

- sender: address, representing the user intending to optin to using the app
- sp: suggested parameters obtained from the network
- index: the app-id as defined by the create method result

```python tab="Python"
# create unsigned transaction
txn = transaction.ApplicationNoOpTxn(sender, params, index)

# sign, send, await 

# display results
transaction_response = client.pending_transaction_info(tx_id)
print("Called app-id: ",transaction_response['txn']['txn']['apid'])
if "global-state-delta" in transaction_response :
    print("Global State updated :\n",transaction_response['global-state-delta'])
if "local-state-delta" in transaction_response :
    print("Local State updated :\n",transaction_response['local-state-delta'])
```

```javascript tab="JavaScript"
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

```Java tab="Java"
// create unsigned transaction
Transaction txn = Transaction.ApplicationCallTransactionBuilder()
                        .sender(sender)
                        .suggestedParams(params)
                        .applicationId(appId)
                        .args(appArgs)
                        .build();

// sign, send, await 

// display results
PendingTransactionResponse pTrx = client.PendingTransactionInformation(id).execute().body();
System.out.println("Called app-id: " + pTrx.txn.tx.applicationId);
if (pTrx.globalStateDelta != null) {
    System.out.println("    Global state: " + pTrx.globalStateDelta.toString());
}
if (pTrx.localStateDelta != null) {
    System.out.println("    Local state: " + pTrx.localStateDelta.toString());
}
```

```go tab="Go"
// create unsigned transaction
txn, _:= future.MakeApplicationNoOpTx(index, appArgs, nil, nil, nil, params, sender, 
                        nil, types.Digest{}, [32]byte{}, types.Address{})

// sign, send, await 

// display results
confirmedTxn, _, err := client.PendingTransactionInformation(txID).Do(context.Background())
fmt.Printf("Called app-id: %d\n", confirmedTxn.Transaction.Txn.ApplicationID)
```

## Read State

Anyone may read the [global state](../../stateful/#reading-global-state-from-other-smart-contracts) of any application or the [local state](../../stateful/#reading-local-state-from-other-accounts) of an application within a given user account using the REST API account_info endpoint. 

```python tab="Python"
# read user local state
def read_local_state(client, addr, app_id) :   
    results = client.account_info(addr)
    local_state = results['apps-local-state'][0]
    for index in local_state :
        if local_state[index] == app_id :
            print(f"local_state of account {addr} for app_id {app_id}: ", local_state['key-value'])

# read app global state
def read_global_state(client, addr, app_id) :   
    results = client.account_info(addr)
    apps_created = results['created-apps']
    for app in apps_created :
        if app['id'] == app_id :
            print(f"global_state for app_id {app_id}: ", app['params']['global-state'])
```

```javascript tab="JavaScript"
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
async function readGlobalState(client, account, index){
    let accountInfoResponse = await client.accountInformation(account.addr).do();
    for (let i = 0; i < accountInfoResponse['created-apps'].length; i++) { 
        if (accountInfoResponse['created-apps'][i].id == index) {
            console.log("Application's global state:");
            for (let n = 0; n < accountInfoResponse['created-apps'][i]['params']['global-state'].length; n++) {
                console.log(accountInfoResponse['created-apps'][i]['params']['global-state'][n]);
            }
        }
    }
}
```

```Java tab="Java"
public void readLocalState(AlgodClient client, Account account, Long appId) throws Exception {
    Response<com.algorand.algosdk.v2.client.model.Account> acctResponse = client.AccountInformation(account.getAddress()).execute();
    List<ApplicationLocalState> applicationLocalState = acctResponse.body().appsLocalState;
    for (int i = 0; i < applicationLocalState.size(); i++) { 
        if (applicationLocalState.get(i).id.equals(appId)) {
            System.out.println("User's application local state: " + applicationLocalState.get(i).keyValue.toString());
        }
    }
}

public void readGlobalState(AlgodClient client, Account account, Long appId) throws Exception {
    Response<com.algorand.algosdk.v2.client.model.Account> acctResponse = client.AccountInformation(account.getAddress()).execute();
    List<Application> createdApplications = acctResponse.body().createdApps;
    for (int i = 0; i < createdApplications.size(); i++) {
        if (createdApplications.get(i).id.equals(appId)) {
            System.out.println("Application global state: " + createdApplications.get(i).params.globalState.toString());
        }
    }
}
```

```go tab="Go"
func readLocalState(client *algod.Client, account crypto.Account, index uint64) {
	accountInfo, _ := client.AccountInformation(account.Address.String()).Do(context.Background())
	for _, ap := range accountInfo.AppsLocalState {
		if ap.Id == index {
			fmt.Printf("Local state for app-id %d (account %s):\n", ap.Id, account.Address.String())
			fmt.Println(ap.KeyValue)
		}
	}
}

func readGlobalState(client *algod.Client, account crypto.Account, index uint64) {
	accountInfo, _ := client.AccountInformation(account.Address.String()).Do(context.Background())
	for _, ap := range accountInfo.CreatedApps {
		if ap.Id == index {
			fmt.Printf("Global state for app-id %d:\n", ap.Id)
			fmt.Println(ap.Params.GlobalState)
		}
	}
}
```

## Update

The creator may [update the approval program](../../stateful/#update-stateful-smart-contract) using the update method (if the current approval program allows it). The refactored approval program source code adds a key/value pair to the user's local storage indicating the timestamp when the application was called. Refer to the [appendix](#refactored-approval-program) for details. The original clear program will be reused.

The update method method requires 6 parameters:

- sender: address, representing the user intending to opt-in to using the app
- sp: suggested parameters obtained from the network
- index: the app-id as defined by the create method result
- approval_program: compiled program
- clear program: compiled program
- app_args: application arguments used by approval program

Construct the update transaction and await the response:

```python tab="Python"
# create unsigned transaction
txn = transaction.ApplicationUpdateTxn(sender, params, app_id, \
                                        approval_program, clear_program, app_args)

# sign, send, await 

# display results
transaction_response = client.pending_transaction_info(tx_id)
app_id = transaction_response['txn']['txn']['apid']
print("Updated existing app-id: ",app_id)
```

```javascript tab="JavaScript"
// create unsigned transaction
let txn = algosdk.makeApplicationUpdateTxn(sender, params, index, approvalProgram, clearProgram);

// sign, send, await

// display results
let transactionResponse = await client.pendingTransactionInformation(txId).do();
let appId = transactionResponse['txn']['txn'].apid;
console.log("Updated app-id: ",appId);
```

```Java tab="Java"
// create unsigned transaction
Transaction txn = Transaction.ApplicationUpdateTransactionBuilder()
                        .sender(sender)
                        .suggestedParams(params)
                        .applicationId(appId)
                        .approvalProgram(approvalProgramSource)
                        .clearStateProgram(clearProgramSource)
                        .build();

// sign, send, await

// display results
PendingTransactionResponse pTrx = client.PendingTransactionInformation(id).execute().body();
System.out.println("Updated new app-id: " + appId);    
```

```go tab="Go"
	// create unsigned transaction
    txn, _ := future.MakeApplicationUpdateTx(index, nil, nil, nil, nil, 
                        approvalProgram, clearProgram, params, creatorAccount.Address, 
                        nil, types.Digest{}, [32]byte{}, types.Address{})

    // sign, send, await
    
	// display results
	confirmedTxn, _, _ := client.PendingTransactionInformation(txID).Do(context.Background())
	fmt.Printf("Updated app-id: %d\n", confirmedTxn.Transaction.Txn.ApplicationID)
}
```

## Call with Arguments

A program may [process arguments passed](../../stateful/#passing-arguments-to-stateful-smart-contracts) at run-time. The NoOp call method has an optional app_args parameter where the timestamp may be supplied:

The refactored application expects a timestamp be supplied with the application call.

```python tab="Python"
# call application with arguments
now = datetime.datetime.now().strftime("%H:%M:%S")
app_args = [now.encode("utf-8")]

# create unsigned transaction
txn = transaction.ApplicationNoOpTxn(sender, params, index, app_args)

# sign, send, await 

# display results
transaction_response = client.pending_transaction_info(tx_id)
print("Called app-id: ",transaction_response['txn']['txn']['apid'])
if "global-state-delta" in transaction_response :
    print("Global State updated :\n",transaction_response['global-state-delta'])
if "local-state-delta" in transaction_response :
    print("Local State updated :\n",transaction_response['local-state-delta'])
```

```javascript tab="JavaScript"
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

```Java tab="Java"
// call application with arguments
SimpleDateFormat formatter= new SimpleDateFormat("yyyy-MM-dd 'at' HH:mm:ss");
Date date = new Date(System.currentTimeMillis());
System.out.println(formatter.format(date));
List<byte[]> appArgs = new ArrayList<byte[]>();
appArgs.add(formatter.format(date).toString().getBytes("UTF8"));  

// create unsigned transaction
Transaction txn = Transaction.ApplicationCallTransactionBuilder()
                        .sender(sender)
                        .suggestedParams(params)
                        .applicationId(appId)
                        .args(appArgs)
                        .build();

// sign, send, await

// display results
PendingTransactionResponse pTrx = client.PendingTransactionInformation(id).execute().body();
System.out.println("Called app-id: " + pTrx.txn.tx.applicationId);
if (pTrx.globalStateDelta != null) {
    System.out.println("    Global state: " + pTrx.globalStateDelta.toString());
}
if (pTrx.localStateDelta != null) {
    System.out.println("    Local state: " + pTrx.localStateDelta.toString());
}
```

```go tab="Go"
// call application with arguments
now := time.Now().Format("Mon Jan _2 15:04:05 2006")
appArgs := make([][]byte, 1)
appArgs[0] = []byte(now)

// create unsigned transaction
txn, _ := future.MakeApplicationNoOpTx(index, appArgs, nil, nil, nil, params, sender, 
                            nil, types.Digest{}, [32]byte{}, types.Address{})

// sign, send, await

// display results
confirmedTxn, _, _ := client.PendingTransactionInformation(txID).Do(context.Background())
fmt.Printf("Called app-id: %d\n", confirmedTxn.Transaction.Txn.ApplicationID)
```

## Close Out

The user may discontinue use of the application by sending a [close out](../../stateful/#the-lifecycle-of-a-stateful-smart-contract) transaction. This will remove the local state for this application from the user's account. This method requires 3 parameters:

- sender: address, representing the user intending to optin to using the app
- sp: suggested parameters obtained from the network
- index: the app-id as defined by the create method result

```python tab="Python"
# create unsigned transaction
txn = transaction.ApplicationCloseOutTxn(sender, params, index)

# sign, send, await 

# display results
transaction_response = client.pending_transaction_info(tx_id)
print("Closed out from app-id: ",transaction_response['txn']['txn']['apid'])
```

```javascript tab="JavaScript"
// create unsigned transaction
let txn = algosdk.makeApplicationCloseOutTxn(sender, params, index)

// sign, send, await

// display results
let transactionResponse = await client.pendingTransactionInformation(txId).do();
console.log("Closed out from app-id:",transactionResponse['txn']['txn']['apid'])
```

```Java tab="Java"
// create unsigned transaction
Transaction txn = Transaction.ApplicationCloseTransactionBuilder()
                        .sender(sender)
                        .suggestedParams(params)
                        .applicationId(appId)
                        .build();

// sign, send, await 

// display results
PendingTransactionResponse pTrx = client.PendingTransactionInformation(id).execute().body();
System.out.println("Closed out from app-id: " + appId);  
```

```go tab="Go"
// create unsigned transaction
txn, _ := future.MakeApplicationCloseOutTx(index, nil, nil, nil, nil, params, account.Address, nil, types.Digest{}, [32]byte{}, types.Address{})

// sign, send, await

// display results
confirmedTxn, _, _ := client.PendingTransactionInformation(txID).Do(context.Background())
fmt.Printf("Closed out from app-id: %d\n", confirmedTxn.Transaction.Txn.ApplicationID)
```

## Delete

The approval program defines the creator as the only account able to [delete the application](../../stateful/#delete-stateful-smart-contract). This removes the global state, but does not impact any user's local state. This method uses the same 3 parameters.

```python tab="Python"
# create unsigned transaction
txn = transaction.ApplicationDeleteTxn(sender, params, index)

# sign, send, await

# display results
transaction_response = client.pending_transaction_info(tx_id)
print("Deleted app-id: ",transaction_response['txn']['txn']['apid'])    
```

```javascript tab="JavaScript"
// create unsigned transaction
let txn = algosdk.makeApplicationDeleteTxn(sender, params, index);

// sign, send, await

// display results
let transactionResponse = await client.pendingTransactionInformation(txId).do();
let appId = transactionResponse['txn']['txn'].apid;
console.log("Deleted app-id: ",appId);
```

```Java tab="Java"
// create unsigned transaction
Transaction txn = Transaction.ApplicationDeleteTransactionBuilder()
                        .sender(sender)
                        .suggestedParams(params)
                        .applicationId(appId)
                        .build();

// sign, send, await

// display results
PendingTransactionResponse pTrx = client.PendingTransactionInformation(id).execute().body();
System.out.println("Deleted app-id: " + appId);
```

```go tab="Go"
// create unsigned transaction
txn, _ := future.MakeApplicationDeleteTx(index, nil, nil, nil, nil, params, sender, 
                        nil, types.Digest{}, [32]byte{}, types.Address{})

// sign, send, await

// display results
confirmedTxn, _, _ := client.PendingTransactionInformation(txID).Do(context.Background())

fmt.Printf("Deleted app-id: %d\n", confirmedTxn.Transaction.Txn.ApplicationID)
```

## Clear State

The user may [clear the local state](../../stateful/#the-lifecycle-of-a-stateful-smart-contract) for an application at any time, even if the application was deleted by the creator. This method uses the same 3 parameter.

```python tab="Python"
# create unsigned transaction
txn = transaction.ApplicationClearStateTxn(sender, params, index)

# sign, send, await 

# display results
transaction_response = client.pending_transaction_info(tx_id)
print("Cleared app-id: ",transaction_response['txn']['txn']['apid']) 
```

```javascript tab="JavaScript"
// create unsigned transaction
let txn = algosdk.makeApplicationClearStateTxn(sender, params, index);

// sign, send, await

// display results
let transactionResponse = await client.pendingTransactionInformation(txId).do();
let appId = transactionResponse['txn']['txn'].apid;
console.log("Cleared local state for app-id: ",appId);
```

```Java tab="Java"
// create unsigned transaction
Transaction txn = Transaction.ApplicationClearTransactionBuilder()
                        .sender(sender)
                        .suggestedParams(params)
                        .applicationId(appId)
                        .build();

// sign, send, await

// display results
PendingTransactionResponse pTrx = client.PendingTransactionInformation(id).execute().body();
System.out.println("Cleared local state for app-id: " + appId);   
```

```go tab="Go"
// create unsigned transaction
txn, _ := future.MakeApplicationClearStateTx(index, nil, nil, nil, nil, params, sender, 
                        nil, types.Digest{}, [32]byte{}, types.Address{})

// sign, send, await

// display results
confirmedTxn, _, _ := client.PendingTransactionInformation(txID).Do(context.Background())
fmt.Printf("Cleared local state for app-id: %d\n", confirmedTxn.Transaction.Txn.ApplicationID)
```

# Appendix

## Approval Program Walkthrough

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

## Clear Program Walkthrough

## Refactored Approval Program

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
