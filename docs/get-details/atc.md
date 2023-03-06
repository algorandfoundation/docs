title: Atomic Transaction Composer


Constructing [atomic transactions](../get-details/atomic_transfers.md) requires building up the set of transactions, assigning a group id, and then signing the transactions. While this is not a difficult thing to do on its own, it can be made more difficult if one of the transactions is an Application Call to an [ABI](../get-details/dapps/smart-contracts/ABI/index.md) compliant application. This is because the arguments passed to the application call must be properly encoded and may include transactions or accounts that are part of the application call or group. 

The Atomic Transaction Composer is a convenient way to build out an atomic group of transactions that handles encoding and decoding of ABI arguments and return values. 

!!! Note 
    The following code examples are snippits to demonstrate usage. The full code for the below snippits is available [here](https://github.com/algorand-devrel/demo-abi).


## Create Atomic Transaction Composer 

To use the Atomic Transaction Composer, first initialize the composer: 

=== "JavaScript"
    <!-- ===JSSDK_ATC_CREATE=== -->
    ```js
    import algosdk from 'algosdk'

    const atc = new algosdk.AtomicTransactionComposer()
    ```
    <!-- ===JSSDK_ATC_CREATE=== -->

=== "Python"
    <!-- ===PYSDK_ATC_CREATE=== -->
```python
from algosdk.atomic_transaction_composer import (
    AtomicTransactionComposer,
    AccountTransactionSigner,
    TransactionWithSigner,
)

atc = AtomicTransactionComposer()
```
[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/doc-examples/_examples/atc.py#L5-L12)
    <!-- ===PYSDK_ATC_CREATE=== -->

=== "Go"
    <!-- ===GOSDK_ATC_CREATE=== -->
    ```go
	    import "github.com/algorand/go-algorand-sdk/transaction"
        //...
    	var atc = transaction.AtomicTransactionComposer{}
    ```
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
[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/doc-examples/_examples/atc.py#L22-L38)
    <!-- ===PYSDK_ATC_ADD_TRANSACTION=== -->

=== "JavaScript"
    <!-- ===JSSDK_ATC_ADD_TRANSACTION=== -->
    ```js
    // ...

    // Returns Account object
    const acct = get_account()

    // Create signer object
    const signer = algosdk.makeBasicAccountTransactionSigner(acct)

    // Get suggested params from the client
    const sp = await client.getTransactionParams().do()

    // Create a transaction
    const ptxn = new Transaction({
        from: acct.addr,
        to: acct.addr, 
        amount: 10000,
        ...sp
    })

    // Construct TransactionWithSigner
    const tws = {txn: ptxn, signer: signer}

    // Pass TransactionWithSigner to ATC
    atc.addTransaction(tws)

    ```
    <!-- ===JSSDK_ATC_ADD_TRANSACTION=== -->

=== "Go"
    <!-- ===GOSDK_ATC_ADD_TRANSACTION=== -->
    ```go

    // ...

    // Returns Account 
    acct, _ := GetAccount()

    // Create signer object
	signer := transaction.BasicAccountTransactionSigner{Account: acct}

    // Get suggested params from client
	sp, _ := client.SuggestedParams().Do(context.Background())

    // Create a transaction
	ptxn, _ := transaction.MakePaymentTxn(acct.Address.String(), acct.Address.String(), 10000, nil, "", sp)

    // Construct TransactionWithSigner
	tws := transaction.TransactionWithSigner{Txn: txn, Signer: signer}

    // Pass TransactionWithSigner to atc

    atc.AddTransaction(tws)

    ```
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

When calling an [ABI](../get-details/dapps/smart-contracts/ABI/index.md) compliant application, the Atomic Transaction Composer will handle encoding and decoding of the arguments passed and the return value.  It will also make sure that any [reference types](../get-details/dapps/smart-contracts/ABI/index.md#reference-types) are packed into the transaction group appropriately.  Additionally, since it knows the method signature and types required, it will do some type checking to make sure the arguments passed are valid for the method call. 

In order to call the methods, a Contract or Interface is constructed. Typically this will be done using a [json file](../get-details/dapps/smart-contracts/ABI/index.md#api) that describes the api for the application.   

Once the Contract object is constructed, it can be used to look up and pass method objects into the Atomic Transaction Composers `add_method_call`

=== "Python"
    <!-- ===PYSDK_ATC_CONTRACT_INIT=== -->
```python
with open("path/to/contract.json") as f:
    js = f.read()
contract = abi.Contract.from_json(js)
```
[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/doc-examples/_examples/atc.py#L41-L44)
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

# This method requires a `transaction` as its second argument.
# Construct the transaction and pass it in as an argument.
# The ATC will handle adding it to the group transaction and
# setting the reference in the application arguments.
ptxn = transaction.PaymentTxn(addr, sp, addr, 10000)
txn = TransactionWithSigner(ptxn, signer)
atc.add_method_call(
    app_id,
    contract.get_method_by_name("txntest"),
    addr,
    sp,
    signer,
    method_args=[10000, txn, 1000],
)
```
[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/doc-examples/_examples/atc.py#L51-L77)
    <!-- ===PYSDK_ATC_ADD_METHOD_CALL=== -->

=== "JavaScript"
    <!-- ===JSSDK_ATC_CONTRACT_INIT=== -->
    ```js

    // Read in the local contract.json file
    const buff = fs.readFileSync("path/to/contract.json")

    // Parse the json file into an object, pass it to create an ABIContract object
    const contract = new algosdk.ABIContract(JSON.parse(buff.toString()))
    ```
    <!-- ===JSSDK_ATC_CONTRACT_INIT=== -->

    <!-- ===JSSDK_ATC_ADD_METHOD_CALL=== -->
    ```js
    const commonParams = {
        appID:contract.networks[genesis_hash].appID,
        sender:acct.addr,
        suggestedParams:sp,
        signer: algosdk.makeBasicAccountTransactionSigner(acct)
    }


    // Simple call to the `add` method, method_args can be any type but _must_ 
    // match those in the method signature of the contract
    atc.addMethodCall({
        method: contract.getMethodByName("add"), methodArgs: [1,1], ...commonParams
    })

    // This method requires a `transaction` as its second argument. Construct the transaction and pass it in as an argument.
    // The ATC will handle adding it to the group transaction and setting the reference in the application arguments.
    txn = {
        txn: new Transaction({ from: acct.addr, to: acct.addr, amount: 10000, ...sp }),
        signer: algosdk.makeBasicAccountTransactionSigner(acct)
    }
    atc.addMethodCall({
        method: getMethodByName("txntest"), 
        methodArgs: [ 10000, txn, 1000 ], 
        ...commonParams
    })

    ```

=== "Go"
    <!-- ===GOSDK_ATC_CONTRACT_INIT=== --->
    ```go
    // Read in contract json file and marshal into Contract
	f, _ := os.Open("path/to/contract.json")
	b, _ := ioutil.ReadAll(f)
	contract := &abi.Contract{}
	_ = json.Unmarshal(b, contract)
    ```
    <!-- ===GOSDK_ATC_CONTRACT_INIT=== --->
    
    <!-- ===GOSDK_ATC_ADD_METHOD_CALL=== --->
    ```go


    func combine(mcp transaction.AddMethodCallParams, m abi.Method, a []interface{}) transaction.AddMethodCallParams {
        mcp.Method = m
        mcp.MethodArgs = a
        return mcp
    }

	mcp := transaction.AddMethodCallParams{
		AppID:           contract.Networks[genesis_hash].AppID,
		Sender:          acct.Address,
		SuggestedParams: sp,
		OnComplete:      types.NoOpOC,
		Signer:          signer,
	}

    // Simple call to the `add` method, method_args can be any type but _must_ 
    // match those in the method signature of the contract
	atc.AddMethodCall(combine(mcp, contract.GetMethodByName(contract, "add"), []interface{}{1, 1}))


    // This method requires a `transaction` as its second argument. Construct the transaction and pass it in as an argument.
    // The ATC will handle adding it to the group transaction and setting the reference in the application arguments.
	txn, _ := transaction.MakePaymentTxn(acct.Address.String(), acct.Address.String(), 10000, nil, "", sp)
	stxn := transaction.TransactionWithSigner{Txn: txn, Signer: signer}
	atc.AddMethodCall(combine(mcp, contract.GetMethodByName(contract, "txntest"), []interface{}{10000, stxn, 1000}))

    ```
    <!-- ===GOSDK_ATC_ADD_METHOD_CALL=== --->
    
=== "Java"
    <!-- ===JAVASDK_ATC_CONTRACT_INIT=== --->
```java
                // Read the json from disk
                String jsonContract = Files.readString(Paths.get("calculator/contract.json"));
                // Create Contract from Json
                Contract contract = Encoder.decodeFromJson(jsonContract, Contract.class);
```
[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/ATC.java#L64-L68)
    <!-- ===JAVASDK_ATC_CONTRACT_INIT=== --->

    <!-- ===JAVASDK_ATC_ADD_METHOD_CALL=== --->
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
    <!-- ===JAVASDK_ATC_ADD_METHOD_CALL=== --->
    
## Execution 

Once all the transactions are added to the atomic group the Atomic Transaction Composer allows several ways to perform the transactions. 

    - Build Group will construct the group of transactions and taking care of assigning the group id, returning an array of unsigned TransactionWithSigner objects.
    - Submit will call build group first, then gather the signatures associated to the transactions, then submit the group without blocking. It will return the full list of transaction ids that can be passed to a wait for confirmation function.
    - Execute will perform submit then wait for confirmation given a number of rounds. It will return the resulting confirmed round, list of transaction ids and any parsed ABI return values if relevant.  


=== "Python"
    <!-- ===PYSDK_ATC_RESULTS=== -->
```python
# Other options:
# txngroup = atc.build_group()
# txids = atc.submit(client)
result = atc.execute(algod_client, 4)
for res in result.abi_results:
    print(res.return_value)
```
[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/doc-examples/_examples/atc.py#L81-L87)
    <!-- ===PYSDK_ATC_RESULTS=== -->

=== "JavaScript"
    <!-- ===JSSDK_ATC_RESULTS=== -->
    ```js
    // Other options:
    // const txgroup = atc.buildGroup()
    // const txids = atc.submit(client)

    const result = await atc.execute(client, 2)
    for(const idx in result.methodResults){
        console.log(result.methodResults[idx])
    }
    ```
    <!-- ===JSSDK_ATC_RESULTS=== -->

=== "Go"
    <!-- ===GOSDK_ATC_RESULTS=== -->
    ```go
    // Other options:
    // txgroup := atc.BuildGroup()
    // txids := atc.Submit(client)

	ret, err := atc.Execute(client, context.Background(), 2)
	if err != nil {
		log.Fatalf("Failed to execute call: %+v", err)
	}

	for _, r := range ret.MethodResults {
		log.Printf("%s returned %+v", r.TxID, r.ReturnValue)
	}
    ```
    <!-- ===GOSDK_ATC_RESULTS=== -->

=== "Java"
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