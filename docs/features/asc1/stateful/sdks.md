title: Using the SDKs

This guide covers using stateful smart contracts with the Algorand SDKs. Stateful smart contracts form the basis for applications written in [Transaction Execution Approval Language (TEAL)](../teal/index.md) or with Python using the [PyTeal](../teal/pyteal.md) library.

Each SDK's install process is discussed in the [SDK Reference](../../../reference/sdks/index.md) documentation.

!!! info
    The example code snippets are provided throughout this page and are abbreviated for conciseness and clarity. Full running code examples for each SDK are available within the GitHub repo at [/examples/smart_contracts](https://github.com/algorand/docs/tree/master/examples/smart_contracts) and for [download](https://github.com/algorand/docs/blob/master/examples/smart_contracts/smart_contracts.zip?raw=true) (.zip).

# Application Lifecycle

This guide follows an application throughout its lifecycle from initial creation, to usage, to modification and finally deletion. The application is a variation of the Hello World counter application. Initially, the application stores the number of times called within its _global state_ and also stores the number of times each user account calls the application within their _local state_. Midway through the lifecycle the application is upgraded to add an additional key:value pair to the user's _local storage_ for storing the call timestamp.

# Environment Setup

This guide requires two accounts:

```python tab="Python"
# user declared account mnemonics
creator_mnemonic = "Your first 25-word mnemonic goes here"
user_mnemonic = "A second distinct 25-word mnemonic goes here"
```
```javascript tab="JavaScript"
```
```Java tab="Java"
// user declared account mnemonics
String creatorMnemonic = "Your 25-word mnemonic goes here";
String userMnemonic = "A second distinct 25-word mnemonic goes here";
```
```go tab="Go"
```

An `algod` client connection is also required. The following connects using Sandbox:

```python tab="Python"
# user declared algod connection parameters
algod_address = "http://localhost:4001"
algod_token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
algod_client = algod.AlgodClient(algod_token, algod_address)
```
```javascript tab="JavaScript"
```
```Java tab="Java"
// user declared account mnemonics
String ALGOD_API_ADDR = "localhost";
Integer ALGOD_PORT = 4001;
String ALGOD_API_TOKEN = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
```
```go tab="Go"
```

!!! Info
    Ensure the `algod` node has the _"EnableDeveloperAPI"_ parameter set to **true** within the `config.json` file.

# Declarations

All stateful applications are comprised of state storage, an approval program and a clear program. Details of each are found within the stateful smart contract guide.

## State Storage
Begin by defining the application's _global_schema_ and _local_schema_ storage requirements. These values are immutable once the application is created, so they must accommodate the maximum required upon upgrade. 

This application will ultimately hold one each of `bytes` and `ints` value within the _local storage_ of the user account, as well as a single `ints` value within _global storage_ of the application:

```python tab="Python"
# declare application state storage (immutable)
local_ints = 1
local_bytes = 1
global_ints = 1
global_bytes = 0
global_schema = transaction.StateSchema(global_ints, global_bytes)
local_schema = transaction.StateSchema(local_ints, local_bytes)
```
```javascript tab="JavaScript"
```
```Java tab="Java"
// declare application state storage (immutable)
int localInts = 1;
int localBytes = 1;
int globalInts = 1;
int globalBytes = 0;
```
```go tab="Go"
```

## Approval Program

The approval program handles the main logic of the application. A detailed walk through of this code is provided in the appendix of this guide.

```python tab="Python"
# user declared approval program (initial)
approval_program_source_initial = b"""#pragma version 2
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

"""
```
```javascript tab="JavaScript"
```
```Java tab="Java"
// user declared approval program (initial)
String approvalProgramSourceInitial = "#pragma version 2\n" +
"///// Handle each possible OnCompletion type. We don't have to worry about\n" +
"//// handling ClearState, because the ClearStateProgram will execute in that\n" +
"//// case, not the ApprovalProgram.\n" +

"txn OnCompletion\n" +
"int NoOp\n" +
"==\n" +
"bnz handle_noop\n" +

"txn OnCompletion\n" +
"int OptIn\n" +
"==\n" +
"bnz handle_optin\n" +

"txn OnCompletion\n" +
"int CloseOut\n" +
"==\n" +
"bnz handle_closeout\n" +

"txn OnCompletion\n" +
"int UpdateApplication\n" +
"==\n" +
"bnz handle_updateapp\n" +

"txn OnCompletion\n" +
"int DeleteApplication\n" +
"==\n" +
"bnz handle_deleteapp\n" +

"//// Unexpected OnCompletion value. Should be unreachable.\n" +
"err\n" +

"handle_noop:\n" +
"//// Handle NoOp\n" +
"//// Check for creator\n" +
"addr 5XWY6RBNYHCSY2HK5HCTO62DUJJ4PT3G4L77FQEBUKE6ZYRGQAFTLZSQQ4\n" +
"txn Sender\n" +
"==\n" +
"bnz handle_optin\n" +

"//// read global state\n" +
"byte \"counter\"\n" +
"dup\n" +
"app_global_get\n" +

"//// increment the value\n" +
"int 1\n" +
"+\n" +

"//// store to scratch space\n" +
"dup\n" +
"store 0\n" +

"//// update global state\n" +
"app_global_put\n" +

"//// read local state for sender\n" +
"int 0\n" +
"byte \"counter\"\n" +
"app_local_get\n" +

"//// increment the value\n" +
"int 1\n" +
"+\n" +
"store 1\n" +

"//// update local state for sender\n" +
"int 0\n" +
"byte \"counter\"\n" +
"load 1\n" +
"app_local_put\n" +

"//// load return value as approval\n" +
"load 0\n" +
"return\n" +

"handle_optin:\n" +
"//// Handle OptIn\n" +
"//// approval\n" +
"int 1\n" +
"return\n" +

"handle_closeout:\n" +
"//// Handle CloseOut\n" +
"////approval\n" +
"int 1\n" +
"return\n" +

"handle_deleteapp:\n" +
"//// Check for creator\n" +
"addr 5XWY6RBNYHCSY2HK5HCTO62DUJJ4PT3G4L77FQEBUKE6ZYRGQAFTLZSQQ4\n" +
"txn Sender\n" +
"==\n" +
"return\n" +

"handle_updateapp:\n" +
"//// Check for creator\n" +
"addr 5XWY6RBNYHCSY2HK5HCTO62DUJJ4PT3G4L77FQEBUKE6ZYRGQAFTLZSQQ4\n" +
"txn Sender\n" +
"==\n" +
"return\n";
```
```go tab="Go"
```

## Clear Program

This is the most basic clear program available:

```python tab="Python"
# declare clear state program source
clear_program_source = b"""#pragma version 2
int 1
"""
```
```javascript tab="JavaScript"
```
```Java tab="Java"
// declare clear state program source
String clearProgramSource = "#pragma version 2\n" +
"int 1\n";
```
```go tab="Go"
```

# Application Methods

## Create

The creator will deploy the application using the create app method. It requires 7 parameters:

- sender: address, representing the creator of the app
- sp: suggested parameters obtained from the network
- on_complete: enum value, representing NoOp
- approval_program: compiled program
- clear program: compiled program
- local_schema: maximum _local storage_ allocation, immutable
- global_schema: maximum _global storage_ allocation, immutable

Use the creator_mnemonic to define sender:

```python tab="Python"
# declare creator
private_key = mnemonic.to_private_key(creator_mnemonic)
creator = account.address_from_private_key(private_key)
```
```javascript tab="JavaScript"
```
```Java tab="Java"
// get accounts from mnemonic
Account creatorAccount = new Account(creatorMnemonic);
Address sender = creatorAccount.getAddress();
```
```go tab="Go"
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
```
```Java tab="Java"
// get node suggested parameters
TransactionParametersResponse params = client.TransactionParams().execute().body();
```
```go tab="Go"
```


Set the on_complete parameter to NoOp: 

```python tab="Python"
# declare on_complete as NoOp
on_complete = transaction.OnComplete.NoOpOC.real
```
```javascript tab="JavaScript"
```
```Java tab="Java"
// not required
```
```go tab="Go"
```


Compile the programs using the `compile` endpoint:

```python tab="Python"
# helper function to compile program source
def compile_program(client, source_code) :
    compile_response = client.compile(source_code.decode('utf-8'))
    return base64.b64decode(compile_response['result'])
```
```javascript tab="JavaScript"
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
```

Construct the transaction with defined values:

```python tab="Python"
# create unsigned transaction
txn = transaction.ApplicationCreateTxn(sender, params, on_complete, \
                                        approval_program, clear_program, \
                                        global_schema, local_schema)
```
```javascript tab="JavaScript"
```
```Java tab="Java"
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
```


!!! Notice
    Note down the app-id from the confirmed transaction response. Place this value into the `index` parameter within all remaining code samples. 

## Opt-In

The user must optin to use the application. This method requires 3 parameters:

- sender: address, representing the user intending to optin to using the app
- sp: suggested parameters obtained from the network
- index: the app-id as defined by the create method result

Use the user_mnemonic to define sender:

```python tab="Python"
# declare sender
private_key = mnemonic.to_private_key(user_mnemonic)
sender = account.address_from_private_key(private_key)
```
```javascript tab="JavaScript"
```
```Java tab="Java"
Account userAccount = new Account(userMnemonic);
Address sender = userAccount.getAddress();
```
```go tab="Go"
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
```
```Java tab="Java"
// get node suggested parameters
TransactionParametersResponse params = client.TransactionParams().execute().body();
```
```go tab="Go"
```


Construct the transaction with defined values:
```python tab="Python"
txn = transaction.ApplicationOptInTxn(sender, params, index)
```
```javascript tab="JavaScript"
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
print("OptIn to app-id: ",transaction_response['txn']['txn']['apid'])  
```  
```javascript tab="JavaScript"
```
```Java tab="Java"
// sign transaction
SignedTransaction signedTxn = account.signTransaction(txn);

// send to network
byte[] encodedTxBytes = Encoder.encodeToMsgPack(signedTxn);
String id = client.RawTransaction().rawtxn(encodedTxBytes).execute().body().txId;

// await confirmation
waitForConfirmation(id);

// display results
PendingTransactionResponse pTrx = client.PendingTransactionInformation(id).execute().body();
System.out.println("OptIn to app-id: " + pTrx.txn.tx.applicationId);       
```
```go tab="Go"
```


## Call (NoOp)

The user may now call the application. This method requires 3 parameters:

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
```

## Read State

Anyone may read the global state of any application or the local state of an application within a given user account using the REST API account_info endpoint. 

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
```


## Update

The creator may update the approval program using the update method (if the current approval program allows it). The following refactored approval program source code adds a key/value pair to the user's local storage indicating the timestamp when the application was called. The original clear program will be reused.

```python tab="Python"
# user declared approval program (refactored)
approval_program_source_refactored = b"""#pragma version 2
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
"""
```
```javascript tab="JavaScript"
```
```Java tab="Java"
// user declared approval program (refactored)
String approvalProgramSourceRefactored = "#pragma version 2\n" +
"//// Handle each possible OnCompletion type. We don't have to worry about\n" +
"//// handling ClearState, because the ClearStateProgram will execute in that\n" +
"//// case, not the ApprovalProgram.\n" +

"txn OnCompletion\n" +
"int NoOp\n" +
"==\n" +
"bnz handle_noop\n" +

"txn OnCompletion\n" +
"int OptIn\n" +
"==\n" +
"bnz handle_optin\n" +

"txn OnCompletion\n" +
"int CloseOut\n" +
"==\n" +
"bnz handle_closeout\n" +

"txn OnCompletion\n" +
"int UpdateApplication\n" +
"==\n" +
"bnz handle_updateapp\n" +

"txn OnCompletion\n" +
"int DeleteApplication\n" +
"==\n" +
"bnz handle_deleteapp\n" +

"//// Unexpected OnCompletion value. Should be unreachable.\n" +
"err\n" +

"handle_noop:\n" +
"//// Handle NoOp\n" +
"//// Check for creator\n" +
"addr 5XWY6RBNYHCSY2HK5HCTO62DUJJ4PT3G4L77FQEBUKE6ZYRGQAFTLZSQQ4\n" +
"txn Sender\n" +
"==\n" +
"bnz handle_optin\n" +

"//// read global state\n" +
"byte \"counter\"\n" +
"dup\n" +
"app_global_get\n" +

"//// increment the value\n" +
"int 1\n" +
"+\n" +

"//// store to scratch space\n" +
"dup\n" +
"store 0\n" +

"//// update global state\n" +
"app_global_put\n" +

"//// read local state for sender\n" +
"int 0\n" +
"byte \"counter\"\n" +
"app_local_get\n" +

"//// increment the value\n" +
"int 1\n" +
"+\n" +
"store 1\n" +

"//// update local state for sender\n" +
"//// update \"counter\"\n" +
"int 0\n" +
"byte \"counter\"\n" +
"load 1\n" +
"app_local_put\n" +

"//// update \"timestamp\"\n" +
"int 0\n" +
"byte \"timestamp\"\n" +
"txn ApplicationArgs 0\n" +
"app_local_put\n" +

"//// load return value as approval\n" +
"load 0\n" +
"return\n" +

"handle_optin:\n" +
"//// Handle OptIn\n" +
"//// approval\n" +
"int 1\n" +
"return\n" +

"handle_closeout:\n" +
"//// Handle CloseOut\n" +
"////approval\n" +
"int 1\n" +
"return\n" +

"handle_deleteapp:\n" +
"//// Check for creator\n" +
"addr 5XWY6RBNYHCSY2HK5HCTO62DUJJ4PT3G4L77FQEBUKE6ZYRGQAFTLZSQQ4\n" +
"txn Sender\n" +
"==\n" +
"return\n" +

"handle_updateapp:\n" +
"//// Check for creator\n" +
"addr 5XWY6RBNYHCSY2HK5HCTO62DUJJ4PT3G4L77FQEBUKE6ZYRGQAFTLZSQQ4\n" +
"txn Sender\n" +
"==\n" +
"return\n";
```
```go tab="Go"
```

The update method method requires 6 parameters:

- sender: address, representing the user intending to optin to using the app
- sp: suggested parameters obtained from the network
- index: the app-id as defined by the create method result
- approval_program: compiled program
- clear program: compiled program
- app_args: application arguments used by approval program

The refactored application expects a timestamp be supplied with the application call, including this update transaction. Therefore, the creator will supply the initial value as an application argument.

```python tab="Python"
# define initial value for key "timestamp"
app_args = [b'initial value']
```
```javascript tab="JavaScript"
```
```Java tab="Java"
// not required
```
```go tab="Go"
```

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
```

## Call with Arguments

A program may process arguments passed at run-time. The NoOp call method has an optional app_args parameter where the timestamp may be supplied:

```python tab="Python"
# call application with arguments
now = datetime.datetime.now().strftime("%H:%M:%S")
app_args = [now.encode("utf-8")]
call_app(algod_client, user_private_key, app_id, app_args)

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
```

## Close Out

The user may discontinue use of the application by sending a close out transaction. This will remove the local state for this application from the user's account. This method requires 3 parameters:

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
```

## Delete

The creator is the only account able to delete the application. This removes the global state, but doe not impact any user's local state. This method uses the same 3 parameter.

```python tab="Python"
# create unsigned transaction
txn = transaction.ApplicationDeleteTxn(sender, params, index)

# sign, send, await

# display results
transaction_response = client.pending_transaction_info(tx_id)
print("Deleted app-id: ",transaction_response['txn']['txn']['apid'])    
```
```javascript tab="JavaScript"
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
```

## Clear State

The user may clear the local state for an application at any time, even if the application was deleted by the creator. This method uses the same 3 parameter.

```python tab="Python"
# create unsigned transaction
txn = transaction.ApplicationClearStateTxn(sender, params, index)

# sign, send, await 

# display results
transaction_response = client.pending_transaction_info(tx_id)
print("Cleared app-id: ",transaction_response['txn']['txn']['apid']) 
```
```javascript tab="JavaScript"
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
```

# Appendix
## Approval Program Walkthrough
## Clear Program Walkthrough
