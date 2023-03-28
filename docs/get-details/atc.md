title: Atomic Transaction Composer


Constructing [atomic transactions](../get-details/atomic_transfers.md) requires building up the set of transactions, assigning a group id, and then signing the transactions. While this is not a difficult thing to do on its own, it can be made more difficult if one of the transactions is an Application Call to an [ABI](/docs/get-details/dapps/smart-contracts/ABI/) compliant application. This is because the arguments passed to the application call must be properly encoded and may include transactions or accounts that are part of the application call or group. 

The Atomic Transaction Composer is a convenient way to build out an atomic group of transactions that handles encoding and decoding of ABI arguments and return values. 

!!! Note 
    The following code examples are snippits to demonstrate usage. Additional examples are available [here](https://github.com/algorand-devrel/demo-abi).


## Create Atomic Transaction Composer 

To use the Atomic Transaction Composer, first initialize the composer: 

=== "JavaScript"
    <!-- ===JSSDK_ATC_CREATE=== -->
	```javascript
	const atc = new algosdk.AtomicTransactionComposer();
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/atc.ts#L51-L52)
    <!-- ===JSSDK_ATC_CREATE=== -->

=== "Python"
    <!-- ===PYSDK_ATC_CREATE=== -->
	```python
	atc = AtomicTransactionComposer()
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/atc.py#L12-L13)
    <!-- ===PYSDK_ATC_CREATE=== -->

=== "Go"
    <!-- ===GOSDK_ATC_CREATE=== -->
	```go
	// Create the atc we'll use to compose our transaction group
	var atc = transaction.AtomicTransactionComposer{}
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/atc/main.go#L39-L41)
    <!-- ===GOSDK_ATC_CREATE=== -->

=== "Java"
    <!-- ===JAVASDK_ATC_CREATE=== -->
	```java
	AtomicTransactionComposer atc = new AtomicTransactionComposer();
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/ATC.java#L47-L48)
    <!-- ===JAVASDK_ATC_CREATE=== -->

## Add individual transactions

Individual transactions being passed to the composer must be wrapped in a `TransactionWithSigner`. This allows some conveniences we'll see later on.

Constructing a Transaction with Signer and adding it to the transaction composer can be done as follows:

=== "Python"
    <!-- ===PYSDK_ATC_ADD_TRANSACTION=== -->
	```python
	addr, sk = acct.address, acct.private_key
	
	# Create signer object
	signer = AccountTransactionSigner(sk)
	
	# Get suggested params from the client
	sp = algod_client.suggested_params()
	
	# Create a transaction
	ptxn = transaction.PaymentTxn(addr, sp, addr, 10000)
	
	# Construct TransactionWithSigner
	tws = TransactionWithSigner(ptxn, signer)
	
	# Pass TransactionWithSigner to ATC
	atc.add_transaction(tws)
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/atc.py#L21-L37)
    <!-- ===PYSDK_ATC_ADD_TRANSACTION=== -->

=== "JavaScript"
    <!-- ===JSSDK_ATC_ADD_TRANSACTION=== -->
	```javascript
	// construct a transaction
	const paymentTxn = algosdk.makePaymentTxnWithSuggestedParamsFromObject({
	  from: sender.addr,
	  suggestedParams,
	  to: sender.addr,
	  amount: 1000,
	});
	
	// add the transaction to the ATC with a signer
	atc.addTransaction({ txn: paymentTxn, signer: sender.signer });
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/atc.ts#L62-L72)
    <!-- ===JSSDK_ATC_ADD_TRANSACTION=== -->

=== "Go"
    <!-- ===GOSDK_ATC_ADD_TRANSACTION=== -->
	```go
	// Get suggested params and make a transaction as usual
	sp, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		log.Fatalf("error getting suggested tx params: %s", err)
	}
	
	txn, err := transaction.MakePaymentTxn(acct1.Address.String(), acct1.Address.String(), 10000, nil, "", sp)
	if err != nil {
		log.Fatalf("failed to make transaction: %s", err)
	}
	
	// Construct a TransactionWithSigner and pass it to the atc
	signer := transaction.BasicAccountTransactionSigner{Account: acct1}
	atc.AddTransaction(transaction.TransactionWithSigner{Txn: txn, Signer: signer})
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/atc/main.go#L44-L58)
    <!-- ===GOSDK_ATC_ADD_TRANSACTION=== -->

=== "Java"
    <!-- ===JAVASDK_ATC_ADD_TRANSACTION=== -->
	```java
	// Create a transaction
	Transaction ptxn = PaymentTransactionBuilder.Builder().amount(10000).suggestedParams(sp)
	                .sender(acct.getAddress()).receiver(acct.getAddress()).build();
	
	// Construct TransactionWithSigner
	TransactionWithSigner tws = new TransactionWithSigner(ptxn,
	                acct.getTransactionSigner());
	
	// Pass TransactionWithSigner to atc
	atc.addTransaction(tws);
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/ATC.java#L51-L61)
    <!-- ===JAVASDK_ATC_ADD_TRANSACTION=== -->

The call to add a transaction may be performed multiple times, each time adding a new transaction to the atomic group. Recall that a maximum of 16 transactions may be included in a single group.


## Calling ABI Methods

When calling an [ABI](/docs/get-details/dapps/smart-contracts/ABI/) compliant application, the Atomic Transaction Composer will handle encoding and decoding of the arguments passed and the return value.  It will also make sure that any [reference types](/docs/get-details/dapps/smart-contracts/ABI/#reference-types) are packed into the transaction group appropriately.  Additionally, since it knows the method signature and types required, it will do some type checking to make sure the arguments passed are valid for the method call. 

In order to call the methods, a Contract or Interface is constructed. Typically this will be done using a [json file](/docs/get-details/dapps/smart-contracts/ABI/#api) that describes the api for the application.   

Once the Contract object is constructed, it can be used to look up and pass method objects into the Atomic Transaction Composers `add_method_call`

=== "Python"
    <!-- ===PYSDK_ATC_CONTRACT_INIT=== -->
	```python
	with open("calculator/contract.json") as f:
	    js = f.read()
	contract = abi.Contract.from_json(js)
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/atc.py#L43-L46)
    <!-- ===PYSDK_ATC_CONTRACT_INIT=== -->
    <!-- ===PYSDK_ATC_ADD_METHOD_CALL=== -->
	```python
	
	# Simple call to the `add` method, method_args can be any type but _must_
	# match those in the method signature of the contract
	atc.add_method_call(
	    app_id,
	    contract.get_method_by_name("add"),
	    addr,
	    sp,
	    signer,
	    method_args=[1, 1],
	)
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/atc.py#L49-L60)
    <!-- ===PYSDK_ATC_ADD_METHOD_CALL=== -->
    <!-- ===PYSDK_ATC_RESULTS=== -->
	```python
	# Other options:
	# txngroup = atc.build_group()
	# txids = atc.submit(client)
	result = atc.execute(algod_client, 4)
	for res in result.abi_results:
	    print(res.return_value)
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/atc.py#L64-L70)
    <!-- ===PYSDK_ATC_RESULTS=== -->

=== "JavaScript"
    <!-- ===JSSDK_ATC_CONTRACT_INIT=== -->
	```javascript
	const abi = JSON.parse(
	  fs.readFileSync(path.join(__dirname, '/calculator/contract.json'), 'utf8')
	);
	const contract = new algosdk.ABIContract(abi);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/atc.ts#L55-L59)
    <!-- ===JSSDK_ATC_CONTRACT_INIT=== -->
    <!-- ===JSSDK_ATC_ADD_METHOD_CALL=== -->
	```javascript
	atc.addMethodCall({
	  appID: appIndex,
	  method: contract.getMethodByName('add'),
	  methodArgs: [1, 2],
	  sender: sender.addr,
	  signer: sender.signer,
	  suggestedParams,
	});
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/atc.ts#L75-L83)
    <!-- ===JSSDK_ATC_ADD_METHOD_CALL=== -->
    <!-- ===JSSDK_ATC_RESULTS=== -->
	```javascript
	const result = await atc.execute(client, 4);
	for (const mr of result.methodResults) {
	  console.log(`${mr.returnValue}`);
	}
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/atc.ts#L86-L90)
    <!-- ===JSSDK_ATC_RESULTS=== -->

=== "Go"
    <!-- ===GOSDK_ATC_CONTRACT_INIT=== -->
	```go
	b, err := ioutil.ReadFile("calculator/contract.json")
	if err != nil {
		log.Fatalf("failed to read contract file: %s", err)
	}
	
	contract := &abi.Contract{}
	if err := json.Unmarshal(b, contract); err != nil {
		log.Fatalf("failed to unmarshal contract: %s", err)
	}
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/atc/main.go#L27-L36)
    <!-- ===GOSDK_ATC_CONTRACT_INIT=== -->
	<!-- ===GOSDK_ATC_ADD_METHOD_CALL=== -->
	```go
	// Grab the method from out contract object
	addMethod, err := contract.GetMethodByName("add")
	if err != nil {
		log.Fatalf("failed to get add method: %s", err)
	}
	
	// Set up method call params
	mcp := transaction.AddMethodCallParams{
		AppID:           appID,
		Sender:          acct1.Address,
		SuggestedParams: sp,
		OnComplete:      types.NoOpOC,
		Signer:          signer,
		Method:          addMethod,
		MethodArgs:      []interface{}{1, 1},
	}
	if err := atc.AddMethodCall(mcp); err != nil {
		log.Fatalf("failed to add method call: %s", err)
	}
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/atc/main.go#L61-L80)
	<!-- ===GOSDK_ATC_ADD_METHOD_CALL=== -->
    <!-- ===GOSDK_ATC_RESULTS=== -->
	```go
	result, err := atc.Execute(algodClient, context.Background(), 4)
	if err != nil {
		log.Fatalf("failed to get add method: %s", err)
	}
	
	for _, r := range result.MethodResults {
		log.Printf("%s => %v", r.Method.Name, r.ReturnValue)
	}
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/atc/main.go#L83-L91)
    <!-- ===GOSDK_ATC_RESULTS=== -->

=== "Java"
    <!-- ===JAVASDK_ATC_CONTRACT_INIT=== -->
	```java
	// Read the json from disk
	String jsonContract = Files.readString(Paths.get("calculator/contract.json"));
	// Create Contract from Json
	Contract contract = Encoder.decodeFromJson(jsonContract, Contract.class);
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/ATC.java#L64-L68)
    <!-- ===JAVASDK_ATC_CONTRACT_INIT=== -->
	<!-- ===JAVASDK_ATC_ADD_METHOD_CALL=== -->
	```java
	// create methodCallParams by builder (or create by constructor) for add method
	List<Object> methodArgs = new ArrayList<Object>();
	methodArgs.add(1);
	methodArgs.add(1);
	
	MethodCallTransactionBuilder<?> mctb = MethodCallTransactionBuilder.Builder();
	
	MethodCallParams mcp = mctb.applicationId(appId).signer(acct.getTransactionSigner())
	                .sender(acct.getAddress())
	                .method(contract.getMethodByName("add")).methodArguments(methodArgs)
	                .onComplete(Transaction.OnCompletion.NoOpOC).suggestedParams(sp).build();
	
	atc.addMethodCall(mcp);
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/ATC.java#L71-L84)
	<!-- ===JAVASDK_ATC_ADD_METHOD_CALL=== -->
    <!-- ===JAVASDK_ATC_RESULTS=== -->
	```java
	ExecuteResult res = atc.execute(algodClient, 2);
	System.out.printf("App call (%s) confirmed in round %d\n", res.txIDs, res.confirmedRound);
	res.methodResults.forEach(methodResult -> {
	        System.out.printf("Result from calling '%s' method: %s\n", methodResult.method.name,
	                        methodResult.value);
	});
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/ATC.java#L87-L93)
    <!-- ===JAVASDK_ATC_RESULTS=== -->