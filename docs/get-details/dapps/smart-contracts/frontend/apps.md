title: Interact with smart contracts

This guide covers using smart contracts with the Algorand SDKs. Smart contracts form the basis for applications written in [Transaction Execution Approval Language (TEAL)](/docs/get-details/dapps/avm/teal/) or with Python using the [PyTeal](/docs/get-details/dapps/writing-contracts/pyteal) library.


# Application lifecycle

This guide follows an application throughout its [lifecycle](/docs/get-details/dapps/smart-contracts/apps/#the-lifecycle-of-a-smart-contract) from initial creation, to usage, to modification and finally deletion. The application stores the number of times it is called within its _global state_ and also stores the number of times each user account calls the application within their _local state_. Midway through the lifecycle, the application is upgraded to add an additional key:value pair to the user's _local storage_ for storing the call timestamp. 

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
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/overview.py#L10-L21)
    <!-- ===PYSDK_ALGOD_CREATE_CLIENT=== -->

=== "JavaScript"
    <!-- ===JSSDK_ALGOD_CREATE_CLIENT=== -->
	```javascript
	const algodToken = 'a'.repeat(64);
	const algodServer = 'http://localhost';
	const algodPort = 4001;
	
	const algodClient = new algosdk.Algodv2(algodToken, algodServer, algodPort);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/overview.ts#L7-L12)
    <!-- ===JSSDK_ALGOD_CREATE_CLIENT=== -->

=== "Java"
    <!-- ===JAVASDK_ALGOD_CREATE_CLIENT=== -->
	```java
	String algodHost = "http://localhost";
	int algodPort = 4001;
	String algodToken = "a".repeat(64);
	AlgodClient algodClient = new AlgodClient(algodHost, algodPort, algodToken);
	
	// OR if the API provider requires a specific header key for the token
	String tokenHeader = "X-API-Key";
	AlgodClient otherAlgodClient = new AlgodClient(algodHost, algodPort, algodToken, tokenHeader);
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/Overview.java#L94-L102)
    <!-- ===JAVASDK_ALGOD_CREATE_CLIENT=== -->

=== "Go"
    <!-- ===GOSDK_ALGOD_CREATE_CLIENT=== -->
	```go
	// Create a new algod client, configured to connect to out local sandbox
	var algodAddress = "http://localhost:4001"
	var algodToken = strings.Repeat("a", 64)
	algodClient, _ := algod.MakeClient(
		algodAddress,
		algodToken,
	)
	
	// Or, if necessary, pass alternate headers
	
	var algodHeader common.Header
	algodHeader.Key = "X-API-Key"
	algodHeader.Value = algodToken
	algodClientWithHeaders, _ := algod.MakeClientWithHeaders(
		algodAddress,
		algodToken,
		[]*common.Header{&algodHeader},
	)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/overview/main.go#L18-L36)
    <!-- ===GOSDK_ALGOD_CREATE_CLIENT=== -->

!!! Info
    Ensure the `algod` node has the _"EnableDeveloperAPI"_ parameter set to **true** within the `config.json` file. This is required to enable the SDK access to the _compile_ and _dryrun_ endpoints.

# Declarations

All smart contracts are comprised of state storage, an approval program and a clear program. Details of each are found within the [smart contract guide](/docs/get-details/dapps/smart-contracts/apps/).

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
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/apps.py#L11-L15)
    <!-- ===PYSDK_APP_SCHEMA=== -->

=== "JavaScript"
    <!-- ===JSSDK_APP_SCHEMA=== -->
	```javascript
	// define uint64s and byteslices stored in global/local storage
	const numGlobalByteSlices = 1;
	const numGlobalInts = 1;
	const numLocalByteSlices = 1;
	const numLocalInts = 1;
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/app.ts#L47-L52)
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
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/AppExamples.java#L40-L46)
    <!-- ===JAVASDK_APP_SCHEMA=== -->

=== "Go"
    <!-- ===GOSDK_APP_SCHEMA=== -->
	```go
	// declare application state storage (immutable)
	var (
		localInts   uint64 = 1
		localBytes  uint64 = 1
		globalInts  uint64 = 1
		globalBytes uint64 = 0
	)
	
	// define schema
	globalSchema := types.StateSchema{NumUint: globalInts, NumByteSlice: globalBytes}
	localSchema := types.StateSchema{NumUint: localInts, NumByteSlice: localBytes}
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/apps/main.go#L55-L66)
    <!-- ===GOSDK_APP_SCHEMA=== -->

!!! Info
    The example application is not allowed to hold any `bytes` value within global storage.


## Approval program

The [approval program](../apps/#the-lifecycle-of-a-smart-contract) handles the main logic of the application.

## Clear program

This is the most basic [clear program](../apps/#the-lifecycle-of-a-smart-contract) and returns _true_ when an account clears its participation in a smart contract:

=== "Python"
    <!-- ===PYSDK_APP_SOURCE=== -->
	```python
	# read the `.teal` source files from disk
	with open("application/approval.teal", "r") as f:
	    approval_program = f.read()
	
	with open("application/clear.teal", "r") as f:
	    clear_program = f.read()
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/apps.py#L18-L24)
    <!-- ===PYSDK_APP_SOURCE=== -->

=== "JavaScript"
    <!-- ===JSSDK_APP_SOURCE=== -->
	```javascript
	// define TEAL source from string or from a file
	const approvalProgram = fs.readFileSync(
	  path.join(__dirname, '/application/approval.teal'),
	  'utf8'
	);
	const clearProgram = fs.readFileSync(
	  path.join(__dirname, '/application/clear.teal'),
	  'utf8'
	);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/app.ts#L17-L26)
    <!-- ===JSSDK_APP_SOURCE=== -->

=== "Java"
    <!-- ===JAVASDK_APP_SOURCE=== -->
	```java
	// Read in the `teal` source files as a string
	String approvalSource = Files.readString(Paths.get("application/approval.teal"));
	String clearSource = Files.readString(Paths.get("application/clear.teal"));
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/AppExamples.java#L49-L52)
    <!-- ===JAVASDK_APP_SOURCE=== -->

=== "Go"
    <!-- ===GOSDK_APP_SOURCE=== -->
	```go
	approvalTeal, err := ioutil.ReadFile("application/approval.teal")
	if err != nil {
		log.Fatalf("failed to read approval program: %s", err)
	}
	clearTeal, err := ioutil.ReadFile("application/clear.teal")
	if err != nil {
		log.Fatalf("failed to read clear program: %s", err)
	}
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/apps/main.go#L69-L77)
    <!-- ===GOSDK_APP_SOURCE=== -->

# Application methods

## Create

The creator will deploy the application using the [create app](../apps/#creating-the-smart-contract) method. It requires 7 parameters:

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
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/account.py#L12-L17)
    <!-- ===PYSDK_ACCOUNT_RECOVER_MNEMONIC=== -->

=== "JavaScript"
    <!-- ===JSSDK_ACCOUNT_RECOVER_MNEMONIC=== -->
	```javascript
	// restore 25-word mnemonic from a string
	// Note the mnemonic should _never_ appear in your source code
	const mnemonic =
	  'creek phrase island true then hope employ veteran rapid hurdle above liberty tissue connect alcohol timber idle ten frog bulb embody crunch taxi abstract month';
	const recoveredAccount = algosdk.mnemonicToSecretKey(mnemonic);
	console.log('Recovered mnemonic account: ', recoveredAccount.addr);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/accounts.ts#L16-L22)
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
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/AccountExamples.java#L64-L70)
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
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/account/main.go#L28-L39)
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
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/apps.py#L29-L36)
    <!-- ===PYSDK_APP_COMPILE=== -->

=== "JavaScript"
    <!-- ===JSSDK_APP_COMPILE=== -->
	```javascript
	const approvalCompileResp = await algodClient
	  .compile(Buffer.from(approvalProgram))
	  .do();
	
	const compiledApprovalProgram = new Uint8Array(
	  Buffer.from(approvalCompileResp.result, 'base64')
	);
	
	const clearCompileResp = await algodClient
	  .compile(Buffer.from(clearProgram))
	  .do();
	
	const compiledClearProgram = new Uint8Array(
	  Buffer.from(clearCompileResp.result, 'base64')
	);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/app.ts#L29-L44)
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
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/AppExamples.java#L55-L62)
    <!-- ===JAVASDK_APP_COMPILE=== -->

=== "Go"
    <!-- ===GOSDK_APP_COMPILE=== -->
	```go
	approvalResult, err := algodClient.TealCompile(approvalTeal).Do(context.Background())
	if err != nil {
		log.Fatalf("failed to compile program: %s", err)
	}
	
	approvalBinary, err := base64.StdEncoding.DecodeString(approvalResult.Result)
	if err != nil {
		log.Fatalf("failed to decode compiled program: %s", err)
	}
	
	clearResult, err := algodClient.TealCompile(clearTeal).Do(context.Background())
	if err != nil {
		log.Fatalf("failed to compile program: %s", err)
	}
	
	clearBinary, err := base64.StdEncoding.DecodeString(clearResult.Result)
	if err != nil {
		log.Fatalf("failed to decode compiled program: %s", err)
	}
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/apps/main.go#L80-L99)
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
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/apps.py#L39-L56)
    <!-- ===PYSDK_APP_CREATE=== -->

=== "JavaScript"
    <!-- ===JSSDK_APP_CREATE=== -->
	```javascript
	const appCreateTxn = algosdk.makeApplicationCreateTxnFromObject({
	  from: creator.addr,
	  approvalProgram: compiledApprovalProgram,
	  clearProgram: compiledClearProgram,
	  numGlobalByteSlices,
	  numGlobalInts,
	  numLocalByteSlices,
	  numLocalInts,
	  suggestedParams,
	  onComplete: algosdk.OnApplicationComplete.NoOpOC,
	});
	
	// Sign and send
	await algodClient
	  .sendRawTransaction(appCreateTxn.signTxn(creator.privateKey))
	  .do();
	const result = await algosdk.waitForConfirmation(
	  algodClient,
	  appCreateTxn.txID().toString(),
	  3
	);
	// Grab app id from confirmed transaction result
	const appId = result['application-index'];
	console.log(`Created app with index: ${appId}`);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/app.ts#L55-L79)
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
	
	PendingTransactionResponse result = Utils.waitForConfirmation(algodClient, createResponse.body().txId,
	                4);
	Long appId = result.applicationIndex;
	System.out.printf("Created application with id: %d\n", appId);
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/AppExamples.java#L65-L85)
    <!-- ===JAVASDK_APP_CREATE=== -->

=== "Go"
    <!-- ===GOSDK_APP_CREATE=== -->
	```go
	// Create application
	sp, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		log.Fatalf("error getting suggested tx params: %s", err)
	}
	
	txn, err := transaction.MakeApplicationCreateTx(
		false, approvalBinary, clearBinary, globalSchema, localSchema,
		nil, nil, nil, nil, sp, creator.Address, nil,
		types.Digest{}, [32]byte{}, types.ZeroAddress,
	)
	if err != nil {
		log.Fatalf("failed to make txn: %s", err)
	}
	
	txid, stx, err := crypto.SignTransaction(creator.PrivateKey, txn)
	if err != nil {
		log.Fatalf("failed to sign transaction: %s", err)
	}
	
	_, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		log.Fatalf("failed to send transaction: %s", err)
	}
	
	confirmedTxn, err := transaction.WaitForConfirmation(algodClient, txid, 4, context.Background())
	if err != nil {
		log.Fatalf("error waiting for confirmation:  %s", err)
	}
	appID := confirmedTxn.ApplicationIndex
	log.Printf("Created app with id: %d", appID)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/apps/main.go#L102-L133)
    <!-- ===GOSDK_APP_CREATE=== -->


!!! Notice
    Note down the app-id from the confirmed transaction response. Place this value into the `index` parameter within all remaining code samples. 

## Opt-in

The user must [opt-in](../apps/#opt-into-the-smart-contract) to call the application if some local state is used during evaluation of the call. This method requires 3 parameters:

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
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/account.py#L12-L17)
	<!-- ===PYSDK_ACCOUNT_RECOVER_MNEMONIC=== -->

=== "JavaScript"
	<!-- ===JSSDK_ACCOUNT_RECOVER_MNEMONIC=== -->
	```javascript
	// restore 25-word mnemonic from a string
	// Note the mnemonic should _never_ appear in your source code
	const mnemonic =
	  'creek phrase island true then hope employ veteran rapid hurdle above liberty tissue connect alcohol timber idle ten frog bulb embody crunch taxi abstract month';
	const recoveredAccount = algosdk.mnemonicToSecretKey(mnemonic);
	console.log('Recovered mnemonic account: ', recoveredAccount.addr);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/accounts.ts#L16-L22)
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
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/AccountExamples.java#L64-L70)
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
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/account/main.go#L28-L39)
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
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/apps.py#L59-L64)
    <!-- ===PYSDK_APP_OPTIN=== -->

=== "JavaScript"
    <!-- ===JSSDK_APP_OPTIN=== -->
	```javascript
	const appOptInTxn = algosdk.makeApplicationOptInTxnFromObject({
	  from: caller.addr,
	  appIndex: appId,
	  suggestedParams,
	});
	
	await algodClient
	  .sendRawTransaction(appOptInTxn.signTxn(caller.privateKey))
	  .do();
	await algosdk.waitForConfirmation(
	  algodClient,
	  appOptInTxn.txID().toString(),
	  3
	);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/app.ts#L84-L98)
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
	
	PendingTransactionResponse optInResult = Utils.waitForConfirmation(algodClient,
	                optInResponse.body().txId, 4);
	assert optInResult.confirmedRound > 0;
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/AppExamples.java#L88-L101)
    <!-- ===JAVASDK_APP_OPTIN=== -->

=== "Go"
    <!-- ===GOSDK_APP_OPTIN=== -->
	```go
	sp, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		log.Fatalf("error getting suggested tx params: %s", err)
	}
	
	// Create a new clawback transaction with the target of the user address and the recipient as the creator
	// address, being sent from the address marked as `clawback` on the asset, in this case the same as creator
	txn, err := transaction.MakeApplicationOptInTx(
		appID, nil, nil, nil, nil, sp,
		caller.Address, nil, types.Digest{}, [32]byte{}, types.ZeroAddress,
	)
	if err != nil {
		log.Fatalf("failed to make txn: %s", err)
	}
	// sign the transaction
	txid, stx, err := crypto.SignTransaction(caller.PrivateKey, txn)
	if err != nil {
		log.Fatalf("failed to sign transaction: %s", err)
	}
	
	// Broadcast the transaction to the network
	_, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		log.Fatalf("failed to send transaction: %s", err)
	}
	
	// Wait for confirmation
	confirmedTxn, err := transaction.WaitForConfirmation(algodClient, txid, 4, context.Background())
	if err != nil {
		log.Fatalf("error waiting for confirmation:  %s", err)
	}
	
	log.Printf("OptIn Transaction: %s confirmed in Round %d\n", txid, confirmedTxn.ConfirmedRound)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/apps/main.go#L139-L172)
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
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/apps.py#L67-L72)
    <!-- ===PYSDK_APP_NOOP=== -->

=== "JavaScript"
    <!-- ===JSSDK_APP_NOOP=== -->
	```javascript
	const appNoOpTxn = algosdk.makeApplicationNoOpTxnFromObject({
	  from: caller.addr,
	  appIndex: appId,
	  suggestedParams,
	});
	
	await algodClient
	  .sendRawTransaction(appNoOpTxn.signTxn(caller.privateKey))
	  .do();
	await algosdk.waitForConfirmation(
	  algodClient,
	  appNoOpTxn.txID().toString(),
	  3
	);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/app.ts#L101-L115)
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
	
	PendingTransactionResponse noopResult = Utils.waitForConfirmation(algodClient, noopResponse.body().txId,
	                4);
	assert noopResult.confirmedRound > 0;
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/AppExamples.java#L104-L117)
    <!-- ===JAVASDK_APP_NOOP=== -->

=== "Go"
    <!-- ===GOSDK_APP_NOOP=== -->
	```go
	sp, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		log.Fatalf("error getting suggested tx params: %s", err)
	}
	
	var (
		appArgs [][]byte
		accts   []string
		apps    []uint64
		assets  []uint64
	)
	
	// Add an arg to our app call
	appArgs = append(appArgs, []byte("arg0"))
	
	txn, err := transaction.MakeApplicationNoOpTx(
		appID, appArgs, accts, apps, assets, sp,
		caller.Address, nil, types.Digest{}, [32]byte{}, types.ZeroAddress,
	)
	if err != nil {
		log.Fatalf("failed to make txn: %s", err)
	}
	
	// sign the transaction
	txid, stx, err := crypto.SignTransaction(caller.PrivateKey, txn)
	if err != nil {
		log.Fatalf("failed to sign transaction: %s", err)
	}
	
	// Broadcast the transaction to the network
	_, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		log.Fatalf("failed to send transaction: %s", err)
	}
	
	// Wait for confirmation
	confirmedTxn, err := transaction.WaitForConfirmation(algodClient, txid, 4, context.Background())
	if err != nil {
		log.Fatalf("error waiting for confirmation:  %s", err)
	}
	
	log.Printf("NoOp Transaction: %s confirmed in Round %d\n", txid, confirmedTxn.ConfirmedRound)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/apps/main.go#L177-L219)
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
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/apps.py#L75-L78)
    <!-- ===PYSDK_APP_READ_STATE=== -->

=== "JavaScript"
    <!-- ===JSSDK_APP_READ_STATE=== -->
	```javascript
	const appInfo = await algodClient.getApplicationByID(appId).do();
	const globalState = appInfo.params['global-state'][0];
	console.log(`Raw global state - ${JSON.stringify(globalState)}`);
	
	// decode b64 string key with Buffer
	const globalKey = Buffer.from(globalState.key, 'base64').toString();
	// decode b64 address value with encodeAddress and Buffer
	const globalValue = algosdk.encodeAddress(
	  Buffer.from(globalState.value.bytes, 'base64')
	);
	
	console.log(`Decoded global state - ${globalKey}: ${globalValue}`);
	
	const accountAppInfo = await algodClient
	  .accountApplicationInformation(caller.addr, appId)
	  .do();
	
	const localState = accountAppInfo['app-local-state']['key-value'][0];
	console.log(`Raw local state - ${JSON.stringify(localState)}`);
	
	// decode b64 string key with Buffer
	const localKey = Buffer.from(localState.key, 'base64').toString();
	// get uint value directly
	const localValue = localState.value.uint;
	
	console.log(`Decoded local state - ${localKey}: ${localValue}`);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/app.ts#L154-L180)
    <!-- ===JSSDK_APP_READ_STATE=== -->

=== "Java"
    <!-- ===JAVASDK_APP_READ_STATE=== -->
	```java
	
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/AppExamples.java#L120-L121)
    <!-- ===JAVASDK_APP_READ_STATE=== -->

=== "Go"
    <!-- ===GOSDK_APP_READ_STATE=== -->
	```go
	// grab global state and config of application
	appInfo, err := algodClient.GetApplicationByID(appID).Do(context.Background())
	if err != nil {
		log.Fatalf("failed to get app info: %s", err)
	}
	log.Printf("app info: %+v", appInfo)
	
	// grab local state for an app id for a single account
	acctInfo, err := algodClient.AccountApplicationInformation(
		acct1.Address.String(), appID,
	).Do(context.Background())
	if err != nil {
		log.Fatalf("failed to get app info: %s", err)
	}
	log.Printf("app info: %+v", acctInfo)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/apps/main.go#L28-L43)
    <!-- ===GOSDK_APP_READ_STATE=== -->

## Update

The creator may [update the approval program](../apps/#update-smart-contract) using the update method (if the current approval program allows it). The refactored approval program source code adds a key/value pair to the user's local storage indicating the timestamp when the application was called. Refer to the [appendix](#refactored-approval-program) for details. The original clear program will be reused.

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
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/apps.py#L81-L102)
    <!-- ===PYSDK_APP_UPDATE=== -->

=== "JavaScript"
    <!-- ===JSSDK_APP_UPDATE=== -->
	```javascript
	const newProgram = fs.readFileSync(
	  path.join(__dirname, '/application/approval_refactored.teal'),
	  'utf8'
	);
	const compiledNewProgram = await compileProgram(algodClient, newProgram);
	
	const appUpdateTxn = algosdk.makeApplicationUpdateTxnFromObject({
	  from: creator.addr,
	  suggestedParams,
	  appIndex: appId,
	  // updates must define both approval and clear programs, even if unchanged
	  approvalProgram: compiledNewProgram,
	  clearProgram: compiledClearProgram,
	});
	
	await algodClient
	  .sendRawTransaction(appUpdateTxn.signTxn(creator.privateKey))
	  .do();
	await algosdk.waitForConfirmation(
	  algodClient,
	  appUpdateTxn.txID().toString(),
	  3
	);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/app.ts#L200-L223)
    <!-- ===JSSDK_APP_UPDATE=== -->

=== "Java"
    <!-- ===JAVASDK_APP_UPDATE=== -->
	```java
	String approvalSourceUpdated = Files.readString(Paths.get("application/approval_refactored.teal"));
	CompileResponse approvalUpdatedResponse = algodClient.TealCompile()
	                .source(approvalSourceUpdated.getBytes())
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
	PendingTransactionResponse updateResult = Utils.waitForConfirmation(algodClient,
	                updateResponse.body().txId, 4);
	assert updateResult.confirmedRound > 0;
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/AppExamples.java#L123-L144)
    <!-- ===JAVASDK_APP_UPDATE=== -->

=== "Go"
    <!-- ===GOSDK_APP_UPDATE=== -->
	```go
	sp, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		log.Fatalf("error getting suggested tx params: %s", err)
	}
	
	var (
		appArgs [][]byte
		accts   []string
		apps    []uint64
		assets  []uint64
	)
	
	txn, err := transaction.MakeApplicationUpdateTx(
		appID, appArgs, accts, apps, assets, approvalBinary, clearBinary,
		sp, caller.Address, nil, types.Digest{}, [32]byte{}, types.ZeroAddress,
	)
	if err != nil {
		log.Fatalf("failed to make txn: %s", err)
	}
	
	// sign the transaction
	txid, stx, err := crypto.SignTransaction(caller.PrivateKey, txn)
	if err != nil {
		log.Fatalf("failed to sign transaction: %s", err)
	}
	
	// Broadcast the transaction to the network
	_, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		log.Fatalf("failed to send transaction: %s", err)
	}
	
	// Wait for confirmation
	confirmedTxn, err := transaction.WaitForConfirmation(algodClient, txid, 4, context.Background())
	if err != nil {
		log.Fatalf("error waiting for confirmation:  %s", err)
	}
	
	log.Printf("Update Transaction: %s confirmed in Round %d\n", txid, confirmedTxn.ConfirmedRound)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/apps/main.go#L227-L266)
    <!-- ===GOSDK_APP_UPDATE=== -->

## Call with arguments

A program may [process arguments passed](../apps/##passing-arguments-to-smart-contracts) at run-time. The NoOp call method has an optional app_args parameter where smart contract parameters can be supplied:

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
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/apps.py#L105-L120)
    <!-- ===PYSDK_APP_CALL=== -->

=== "JavaScript"
    <!-- ===JSSDK_APP_CALL=== -->
	```javascript
	const now = new Date().toString();
	const simpleAddTxn = algosdk.makeApplicationNoOpTxnFromObject({
	  from: caller.addr,
	  suggestedParams,
	  appIndex: appId,
	  appArgs: [new Uint8Array(Buffer.from(now))],
	});
	
	await algodClient
	  .sendRawTransaction(simpleAddTxn.signTxn(caller.privateKey))
	  .do();
	await algosdk.waitForConfirmation(
	  algodClient,
	  simpleAddTxn.txID().toString(),
	  3
	);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/app.ts#L135-L151)
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
	
	PendingTransactionResponse callResult = Utils.waitForConfirmation(algodClient, callResponse.body().txId,
	                4);
	assert callResult.confirmedRound > 0;
	// display results
	if (callResult.globalStateDelta != null) {
	        System.out.printf("\tGlobal state: %s\n", callResult.globalStateDelta);
	}
	
	if (callResult.localStateDelta != null) {
	        System.out.printf("\tLocal state: %s\n", callResult.localStateDelta);
	}
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/AppExamples.java#L147-L176)
    <!-- ===JAVASDK_APP_CALL=== -->

=== "Go"
    <!-- ===GOSDK_APP_CALL=== -->
	```go
	sp, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		log.Fatalf("error getting suggested tx params: %s", err)
	}
	
	var (
		appArgs [][]byte
		accts   []string
		apps    []uint64
		assets  []uint64
	)
	
	datetime := time.Now().Format("2006-01-02 at 15:04:05")
	appArgs = append(appArgs, []byte(datetime))
	
	txn, err := transaction.MakeApplicationNoOpTx(
		appID, appArgs, accts, apps, assets, sp,
		caller.Address, nil, types.Digest{}, [32]byte{}, types.ZeroAddress,
	)
	if err != nil {
		log.Fatalf("failed to make txn: %s", err)
	}
	
	// sign the transaction
	txid, stx, err := crypto.SignTransaction(caller.PrivateKey, txn)
	if err != nil {
		log.Fatalf("failed to sign transaction: %s", err)
	}
	
	// Broadcast the transaction to the network
	_, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		log.Fatalf("failed to send transaction: %s", err)
	}
	
	// Wait for confirmation
	confirmedTxn, err := transaction.WaitForConfirmation(algodClient, txid, 4, context.Background())
	if err != nil {
		log.Fatalf("error waiting for confirmation:  %s", err)
	}
	
	log.Printf("NoOp Transaction: %s confirmed in Round %d\n", txid, confirmedTxn.ConfirmedRound)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/apps/main.go#L359-L401)
    <!-- ===GOSDK_APP_CALL=== -->

## Close out

The user may discontinue use of the application by sending a [close out](../apps/#the-lifecycle-of-a-smart-contract) transaction. This will remove the local state for this application from the user's account. This method requires 3 parameters:

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
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/apps.py#L123-L128)
    <!-- ===PYSDK_APP_CLOSEOUT=== -->

=== "JavaScript"
    <!-- ===JSSDK_APP_CLOSEOUT=== -->
	```javascript
	const appCloseOutTxn = algosdk.makeApplicationCloseOutTxnFromObject({
	  from: caller.addr,
	  appIndex: appId,
	  suggestedParams,
	});
	
	await algodClient
	  .sendRawTransaction(appCloseOutTxn.signTxn(caller.privateKey))
	  .do();
	await algosdk.waitForConfirmation(
	  algodClient,
	  appCloseOutTxn.txID().toString(),
	  3
	);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/app.ts#L183-L197)
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
	
	PendingTransactionResponse closeOutResult = Utils.waitForConfirmation(algodClient,
	                closeOutResponse.body().txId,
	                4);
	assert closeOutResult.confirmedRound > 0;
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/AppExamples.java#L179-L193)
    <!-- ===JAVASDK_APP_CLOSEOUT=== -->

=== "Go"
    <!-- ===GOSDK_APP_CLOSEOUT=== -->
	```go
	sp, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		log.Fatalf("error getting suggested tx params: %s", err)
	}
	
	var (
		appArgs [][]byte
		accts   []string
		apps    []uint64
		assets  []uint64
	)
	
	txn, err := transaction.MakeApplicationCloseOutTx(
		appID, appArgs, accts, apps, assets, sp,
		caller.Address, nil, types.Digest{}, [32]byte{}, types.ZeroAddress,
	)
	if err != nil {
		log.Fatalf("failed to make txn: %s", err)
	}
	
	// sign the transaction
	txid, stx, err := crypto.SignTransaction(caller.PrivateKey, txn)
	if err != nil {
		log.Fatalf("failed to sign transaction: %s", err)
	}
	
	// Broadcast the transaction to the network
	_, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		log.Fatalf("failed to send transaction: %s", err)
	}
	
	// Wait for confirmation
	confirmedTxn, err := transaction.WaitForConfirmation(algodClient, txid, 4, context.Background())
	if err != nil {
		log.Fatalf("error waiting for confirmation:  %s", err)
	}
	
	log.Printf("Closeout Transaction: %s confirmed in Round %d\n", txid, confirmedTxn.ConfirmedRound)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/apps/main.go#L271-L310)
    <!-- ===GOSDK_APP_CLOSEOUT=== -->

## Delete

The approval program defines the creator as the only account able to [delete the application](../apps/#delete-smart-contract). This removes the global state, but does not impact any user's local state. This method uses the same 3 parameters.

=== "Python"
    <!-- ===PYSDK_APP_DELETE=== -->
	```python
	delete_txn = transaction.ApplicationDeleteTxn(creator.address, sp, app_id)
	signed_delete = delete_txn.sign(creator.private_key)
	txid = algod_client.send_transaction(signed_delete)
	optin_result = transaction.wait_for_confirmation(algod_client, txid, 4)
	assert optin_result["confirmed-round"] > 0
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/apps.py#L131-L136)
    <!-- ===PYSDK_APP_DELETE=== -->

=== "JavaScript"
    <!-- ===JSSDK_APP_DELETE=== -->
	```javascript
	const appDeleteTxn = algosdk.makeApplicationDeleteTxnFromObject({
	  from: creator.addr,
	  suggestedParams,
	  appIndex: appId,
	});
	
	await algodClient
	  .sendRawTransaction(appDeleteTxn.signTxn(creator.privateKey))
	  .do();
	await algosdk.waitForConfirmation(
	  algodClient,
	  appDeleteTxn.txID().toString(),
	  3
	);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/app.ts#L243-L257)
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
	PendingTransactionResponse deleteResult = Utils.waitForConfirmation(algodClient,
	                deleteResponse.body().txId, 4);
	assert deleteResult.confirmedRound > 0;
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/AppExamples.java#L196-L208)
    <!-- ===JAVASDK_APP_DELETE=== -->

=== "Go"
    <!-- ===GOSDK_APP_DELETE=== -->
	```go
	sp, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		log.Fatalf("error getting suggested tx params: %s", err)
	}
	
	var (
		appArgs [][]byte
		accts   []string
		apps    []uint64
		assets  []uint64
	)
	
	txn, err := transaction.MakeApplicationDeleteTx(
		appID, appArgs, accts, apps, assets, sp,
		caller.Address, nil, types.Digest{}, [32]byte{}, types.ZeroAddress,
	)
	if err != nil {
		log.Fatalf("failed to make txn: %s", err)
	}
	
	// sign the transaction
	txid, stx, err := crypto.SignTransaction(caller.PrivateKey, txn)
	if err != nil {
		log.Fatalf("failed to sign transaction: %s", err)
	}
	
	// Broadcast the transaction to the network
	_, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		log.Fatalf("failed to send transaction: %s", err)
	}
	
	// Wait for confirmation
	confirmedTxn, err := transaction.WaitForConfirmation(algodClient, txid, 4, context.Background())
	if err != nil {
		log.Fatalf("error waiting for confirmation:  %s", err)
	}
	
	log.Printf("Delete Transaction: %s confirmed in Round %d\n", txid, confirmedTxn.ConfirmedRound)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/apps/main.go#L406-L445)
    <!-- ===GOSDK_APP_DELETE=== -->

## Clear state

The user may [clear the local state](../apps/#the-lifecycle-of-a-smart-contract) for an application at any time, even if the application was deleted by the creator. This method uses the same 3 parameters.

=== "Python"
    <!-- ===PYSDK_APP_CLEAR=== -->
	```python
	clear_txn = transaction.ApplicationClearStateTxn(user.address, sp, app_id)
	# .. sign, send, wait
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/apps.py#L139-L141)
    <!-- ===PYSDK_APP_CLEAR=== -->

=== "JavaScript"
    <!-- ===JSSDK_APP_CLEAR=== -->
	```javascript
	const appClearTxn = algosdk.makeApplicationClearStateTxnFromObject({
	  from: anotherCaller.addr,
	  suggestedParams,
	  appIndex: appId,
	});
	
	await algodClient
	  .sendRawTransaction(appClearTxn.signTxn(anotherCaller.privateKey))
	  .do();
	await algosdk.waitForConfirmation(
	  algodClient,
	  appClearTxn.txID().toString(),
	  3
	);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/app.ts#L226-L240)
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
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/AppExamples.java#L211-L219)
    <!-- ===JAVASDK_APP_CLEAR=== -->

=== "Go"
    <!-- ===GOSDK_APP_CLEAR=== -->
	```go
	sp, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		log.Fatalf("error getting suggested tx params: %s", err)
	}
	
	var (
		appArgs [][]byte
		accts   []string
		apps    []uint64
		assets  []uint64
	)
	
	txn, err := transaction.MakeApplicationClearStateTx(
		appID, appArgs, accts, apps, assets, sp,
		caller.Address, nil, types.Digest{}, [32]byte{}, types.ZeroAddress,
	)
	if err != nil {
		log.Fatalf("failed to make txn: %s", err)
	}
	
	// sign the transaction
	txid, stx, err := crypto.SignTransaction(caller.PrivateKey, txn)
	if err != nil {
		log.Fatalf("failed to sign transaction: %s", err)
	}
	
	// Broadcast the transaction to the network
	_, err = algodClient.SendRawTransaction(stx).Do(context.Background())
	if err != nil {
		log.Fatalf("failed to send transaction: %s", err)
	}
	
	// Wait for confirmation
	confirmedTxn, err := transaction.WaitForConfirmation(algodClient, txid, 4, context.Background())
	if err != nil {
		log.Fatalf("error waiting for confirmation:  %s", err)
	}
	
	log.Printf("ClearState Transaction: %s confirmed in Round %d\n", txid, confirmedTxn.ConfirmedRound)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/apps/main.go#L315-L354)
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