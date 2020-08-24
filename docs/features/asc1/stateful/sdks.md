title: Using the SDKs

This guide covers using stateful smart contracts with the Algorand SDKs. Stateful smart contracts form the basis for applications written in [Transaction Execution Approval Language (TEAL)](../teal/index.md) or with Python using the [PyTeal](../teal/pyteal.md) library.

Each SDK's install process is discussed in the [SDK Reference](../../../reference/sdks/index.md) documentation.

!!! info
    The example code snippets are provided throughout this page and are abbreviated for conciseness and clarity. Full running code examples for each SDK are available within the GitHub repo at [/examples/smart_contracts](https://github.com/algorand/docs/tree/master/examples/smart_contracts) and for [download](https://github.com/algorand/docs/blob/master/examples/smart_contracts/smart_contracts.zip?raw=true) (.zip).

# Application Lifecycle

This guide follows an application throughout its lifecycle from initial creation, to usage, to modification and finally deletion. The application is a variation of the Hello World counter application. Initially, the application stores the number of times called within its _global state_ and also stores the number of times each user account calls the application within their _local state_. Midway through the lifecycle the application is upgraded to add an additional key:value pair to the user's _local storage_ for storing the call timestamp.

# Environment Setup

This guide requires two accounts:

```python
# user declared account mnemonics
creator_mnemonic = "Your first 25-word mnemonic goes here"
user_mnemonic = "A second distinct 25-word mnemonic goes here"
```

An `algod` client connection is also required. The following connects using Sandbox:

```python
# user declared algod connection parameters
algod_address = "http://localhost:4001"
algod_token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
algod_client = algod.AlgodClient(algod_token, algod_address)
```

!!! Info
    Ensure the `algod` node has the _"EnableDeveloperAPI"_ parameter set to **true** within the `config.json` file.

# Declarations

All stateful applications are comprised of state storage, an approval program and a clear program. Details of each are found within the stateful smart contract guide.

## State Storage
Begin by defining the application's _global_schema_ and _local_schema_ storage requirements. These values are immutable once the application is created, so they must accommodate the maximum required upon upgrade. 

This application will ultimately hold one each of `bytes` and `ints` value within the _local storage_ of the user account, as well as a single `ints` value within _global storage_ of the application:

```python
# declare application state storage (immutable)
local_ints = 1
local_bytes = 1
global_ints = 1
global_bytes = 0
global_schema = transaction.StateSchema(global_ints, global_bytes)
local_schema = transaction.StateSchema(local_ints, local_bytes)
```
## Approval Program

The approval program handles the main logic of the application. A detailed walk through of this code is provided in the appendix of this guide.

```python
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

## Clear Program

This is the most basic clear program available:

```python
# declare clear state program source
clear_program_source = b"""#pragma version 2
int 1
"""
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

```python
# declare creator
private_key = mnemonic.to_private_key(creator_mnemonic)
creator = account.address_from_private_key(private_key)
```

Use the `suggested_params` endpoint:

```python
# get node suggested parameters
params = client.suggested_params()
# comment out the next two (2) lines to use suggested fees
params.flat_fee = True
params.fee = 1000
```

Set the on_complete parameter to NoOp: 

```python
# declare on_complete as NoOp
on_complete = transaction.OnComplete.NoOpOC.real
```

Compile the programs using the `compile` endpoint:

```python
# helper function to compile program source
def compile_program(client, source_code) :
    compile_response = client.compile(source_code.decode('utf-8'))
    return base64.b64decode(compile_response['result'])
```

Construct the transaction with defined values:

```python
# create unsigned transaction
txn = transaction.ApplicationCreateTxn(sender, params, on_complete, \
                                        approval_program, clear_program, \
                                        global_schema, local_schema)
```

Sign, send, await confirmation and display the results:

```python
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

!!! Notice
    Note down the app-id from the confirmed transaction response. Place this value into the `index` parameter within all remaining code samples. 

## Opt-In

The user must optin to use the application. This method requires 3 parameters:

- sender: address, representing the user intending to optin to using the app
- sp: suggested parameters obtained from the network
- index: the app-id as defined by the create method result

Use the user_mnemonic to define sender:

```python
# declare sender
private_key = mnemonic.to_private_key(user_mnemonic)
sender = account.address_from_private_key(private_key)
```

Use the `suggested_params` endpoint:

```python
# get node suggested parameters
params = client.suggested_params()
# comment out the next two (2) lines to use suggested fees
params.flat_fee = True
params.fee = 1000
```

Construct the transaction with defined values:
```python
txn = transaction.ApplicationOptInTxn(sender, params, index)
```

Sign, send, await confirmation and display the results:

```python
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

## Call (NoOp)

The user may now call the application. This method requires 3 parameters:

- sender: address, representing the user intending to optin to using the app
- sp: suggested parameters obtained from the network
- index: the app-id as defined by the create method result

```python
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

## Read State

Anyone may read the global state of any application or the local state of an application within a given user account using the REST API account_info endpoint. 

```python
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

## Update

The creator may update the approval program using the update method (if the current approval program allows it). The following refactored approval program source code adds a key/value pair to the user's local storage indicating the timestamp when the application was called. The original clear program will be reused.

```python
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

The update method method requires 6 parameters:

- sender: address, representing the user intending to optin to using the app
- sp: suggested parameters obtained from the network
- index: the app-id as defined by the create method result
- approval_program: compiled program
- clear program: compiled program
- app_args: application arguments used by approval program

The refactored application expects a timestamp be supplied with the application call, including this update transaction. Therefore, the creator will supply the initial value as an application argument.

```python
# define initial value for key "timestamp"
app_args = [b'initial value']
```
Construct the update transaction and await the response:

```python
# create unsigned transaction
txn = transaction.ApplicationUpdateTxn(sender, params, app_id, \
                                        approval_program, clear_program, app_args)

# sign, send, await 

# display results
transaction_response = client.pending_transaction_info(tx_id)
app_id = transaction_response['txn']['txn']['apid']
print("Updated existing app-id: ",app_id)
```

## Call with Arguments

A program may process arguments passed at run-time. The NoOp call method has an optional app_args parameter where the timestamp may be supplied:

```python
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

## Close Out

The user may discontinue use of the application by sending a close out transaction. This will remove the local state for this application from the user's account. This method requires 3 parameters:

- sender: address, representing the user intending to optin to using the app
- sp: suggested parameters obtained from the network
- index: the app-id as defined by the create method result

```python
# create unsigned transaction
txn = transaction.ApplicationCloseOutTxn(sender, params, index)

# sign, send, await 

# display results
transaction_response = client.pending_transaction_info(tx_id)
print("Closed out from app-id: ",transaction_response['txn']['txn']['apid'])
```

## Delete

The creator is the only account able to delete the application. This removes the global state, but doe not impact any user's local state. This method uses the same 3 parameter.

```python
# create unsigned transaction
txn = transaction.ApplicationDeleteTxn(sender, params, index)

# sign, send, await

# display results
transaction_response = client.pending_transaction_info(tx_id)
print("Deleted app-id: ",transaction_response['txn']['txn']['apid'])    
```

## Clear State

The user may clear the local state for an application at any time, even if the application was deleted by the creator. This method uses the same 3 parameter.

```python
# create unsigned transaction
txn = transaction.ApplicationClearStateTxn(sender, params, index)

# sign, send, await 

# display results
transaction_response = client.pending_transaction_info(tx_id)
print("Cleared app-id: ",transaction_response['txn']['txn']['apid']) 
```

# Appendix
## Approval Program Walkthrough
## Clear Program Walkthrough

