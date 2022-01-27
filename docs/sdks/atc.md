title: Atomic Transaction Composer


Constructing [atomic transactions](../get-details/atomic_transfers.md) requires building up the set of transactions, assigning a group id, and then signing the transactions. While this is not a difficult thing to do on its own, it can be made more difficult if one of the transactions is an Application Call to an [ABI](../get-details/dapps/smart-contracts/ABI/index.md) compliant application. This is because the arguments passed to the application call must be properly encoded and may include transactions or accounts that are part of the application call or group. 

The Atomic Transaction Composer is a convenient way to build out an atomic group of transactions that handles encoding and decoding of ABI arguments and return values. 

!!! Note 
    The following code examples are snippits to demonstrate usage. The full code for the below snippits is available [here](https://github.com/algorand-devrel/demo-abi).


## Create Atomic Transaction Composer 

To use the Atomic Transaction Composer, first initialize the composer: 

=== "Python"
    ```py
    from algosdk.atomic_transaction_composer import AtomicTransactionComposer

    atc = AtomicTransactionComposer()
    ```

=== "JavaScript"
    ```js
    import algosdk from 'algosdk'

    const atc = new algosdk.AtomicTransactionComposer()
    ```

=== "Go"
    ```go
	    import "github.com/algorand/go-algorand-sdk/future"
        //...
    	var atc = future.AtomicTransactionComposer{}
    ```

=== "Java"
    ```java
        AtomicTransactionComposer atc = new AtomicTransactionComposer();
    ```

## Add individual transactions

Individual transactions being passed to the composer must be wrapped in a `TransactionWithSigner`. This allows some conveniences we'll see later on.

Constructing a Transaction with Signer and adding it to the transaction composer can be done as follows:

=== "Python"
    ```py

    #...

    addr, sk = get_account()

    # Create signer object
    signer = AccountTransactionSigner(sk)

    # Get suggested params from the client
    sp = client.suggested_params()

    # Create a transaction
    ptxn = PaymentTxn(addr, sp, addr, 10000)

    # Construct TransactionWithSigner
    tws = TransactionWithSigner(ptxn, signer)

    # Pass TransactionWithSigner to ATC
    atc.add_transaction(tws)
    ```

=== "JavaScript"
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

=== "Go"
    ```go

    // ...

    // Returns Account 
    acct, _ := GetAccount()

    // Create signer object
	signer := future.BasicAccountTransactionSigner{Account: acct}

    // Get suggested params from client
	sp, _ := client.SuggestedParams().Do(context.Background())

    // Create a transaction
	ptxn, _ := future.MakePaymentTxn(acct.Address.String(), acct.Address.String(), 10000, nil, "", sp)

    // Construct TransactionWithSigner
	tws := future.TransactionWithSigner{Txn: txn, Signer: signer}

    // Pass TransactionWithSigner to atc

    atc.AddTransaction(tws)

    ```

=== "Java"
    ```java

    // ...

    // Returns Account 
    Account acct = getAccount();

    // Create signer object
	BasicAccountTransactionSigner signer = new BasicAccountTransactionSigner(acct);

    // Get suggested params from client
    Response<TransactionParametersResponse> rsp = client.TransactionParams().execute();
    TransactionParametersResponse sp = rsp.body();

    // Create a transaction
    Transaction ptxn = PaymentTransactionBuilder.Builder().amount(10000).suggestedParams(sp)
            .sender(acct.getAddress()).receiver(acct.getAddress()).build();

    // Construct TransactionWithSigner
	TransactionWithSigner tws = new TransactionWithSigner(ptxn, signer);

    // Pass TransactionWithSigner to atc
    atc.addTransaction(tws);

    ```

The call to add a transaction may may be performed multiple times, each time adding a new transaction to the atomic group. Recall that a maximum of 16 transactions may be included in a single group.


## Calling ABI Methods

When calling an [ABI](../get-details/dapps/smart-contracts/ABI/index.md) compliant application, the Atomic Transaction Composer will handle encoding and decoding of the arguments passed and the return value.  It will also make sure that any [reference types](../get-details/dapps/smart-contracts/ABI/index.md#reference-types) are packed into the transaction group appropriately.  Additionally, since it knows the method signature and types required, it will do some type checking to make sure the arguments passed are valid for the method call. 

In order to call the methods, a Contract or Interface is constructed. Typically this will be done using a [json file](../get-details/dapps/smart-contracts/ABI/index.md#api) that describes the api for the application.   

Once the Contract object is constructed, it can be used to look up and pass method objects into the Atomic Transaction Composers `add_method_call`

=== "Python"
    ```py
    from algosdk.abi import Contract

    with open("path/to/contract.json") as f:
        js = f.read()
    c = Contract.from_json(js)

    # Using the app id from the "sandnet" network, which is hardcoded in the json file
    app_id = c.networks["sandnet"].app_id

    # Utility function to get the Method object for a given method name
    def get_method(name: str) -> Method:
        for m in c.methods:
            if m.name == name:
                return m
        raise Exception("No method with the name {}".format(name))



    # Simple call to the `add` method, method_args can be any type but _must_ 
    # match those in the method signature of the contract
    atc.add_method_call(app_id, get_method("add"), addr, sp, signer, method_args=[1,1])

    # This method requires a `transaction` as its second argument. Construct the transaction and pass it in as an argument.
    # The ATC will handle adding it to the group transaction and setting the reference in the application arguments.
    txn = TransactionWithSigner(PaymentTxn(addr, sp, addr, 10000), signer)
    atc.add_method_call(app_id, get_method("txntest"), addr, sp, signer, method_args=[10000, txn, 1000])

    ```

=== "JavaScript"
    ```js

    // Read in the local contract.json file
    const buff = fs.readFileSync("path/to/contract.json")

    // Parse the json file into an object, pass it to create an ABIContract object
    const contract = new algosdk.ABIContract(JSON.parse(buff.toString()))

    // Utility function to return an ABIMethod by its name
    function getMethodByName(name: string): algosdk.ABIMethod  {
        const m = contract.methods.find((mt: algosdk.ABIMethod)=>{ return mt.name==name })
        if(m === undefined)
            throw Error("Method undefined: "+name)
        return m
    }

    const commonParams = {
        appID:contract.networks["sandnet"].appID,
        sender:acct.addr,
        suggestedParams:sp,
        signer: algosdk.makeBasicAccountTransactionSigner(acct)
    }


    // Simple call to the `add` method, method_args can be any type but _must_ 
    // match those in the method signature of the contract
    atc.addMethodCall({
        method: getMethodByName("add"), methodArgs: [1,1], ...commonParams
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
    ```go

    // Read in contract json file and marshal into Contract
	f, _ := os.Open("path/to/contract.json")
	b, _ := ioutil.ReadAll(f)
	contract := &abi.Contract{}
	_ = json.Unmarshal(b, contract)


    // Utility function to get a Method given the name 
    func getMethod(c *abi.Contract, name string) (abi.Method, error) {
        for _, m = range c.Methods {
            if m.Name == name {
                return m, nil
            }
        }
        return abi.Method{}, fmt.Errorf("No method named: %s", name)
    }

    func combine(mcp future.AddMethodCallParams, m abi.Method, a []interface{}) future.AddMethodCallParams {
        mcp.Method = m
        mcp.MethodArgs = a
        return mcp
    }

	mcp := future.AddMethodCallParams{
		AppID:           contract.Networks["sandnet"].AppID,
		Sender:          acct.Address,
		SuggestedParams: sp,
		OnComplete:      types.NoOpOC,
		Signer:          signer,
	}

    // Simple call to the `add` method, method_args can be any type but _must_ 
    // match those in the method signature of the contract
	atc.AddMethodCall(combine(mcp, getMethod(contract, "add"), []interface{}{1, 1}))


    // This method requires a `transaction` as its second argument. Construct the transaction and pass it in as an argument.
    // The ATC will handle adding it to the group transaction and setting the reference in the application arguments.
	txn, _ := future.MakePaymentTxn(acct.Address.String(), acct.Address.String(), 10000, nil, "", sp)
	stxn := future.TransactionWithSigner{Txn: txn, Signer: signer}
	atc.AddMethodCall(combine(mcp, getMethod(contract, "txntest"), []interface{}{10000, stxn, 1000}))

    ```

=== "Java"
    ```java
    ```

## Execution 

Once all the transactions are added to the atomic group the Atomic Transaction Composer allows several ways to perform the transactions. 

    - Build Group will construct the group of transactions and taking care of assigning the group id, returning an array of unsigned TransactionWithSigner objects.
    - Submit will call build group first, then gather the signatures associated to the transactions, then submit the group without blocking. It will return the full list of transaction ids that can be passed to a wait for confirmation function.
    - Execute will perform submit then wait for confirmation given a number of rounds. It will return the resulting confirmed round, list of transaction ids and any parsed ABI return values if relevant.  


=== "Python"
    ```py
    # Other options:
    # txngroup = atc.build_group()
    # txids = atc.submit(client)

    result = atc.execute(client, 2)
    for res in result.abi_results:
        print(res.return_value)
    ```

=== "JavaScript"
    ```js
    // Other options:
    // const txgroup = atc.buildGroup()
    // const txids = atc.submit(client)

    const result = await atc.execute(client, 2)
    for(const idx in result.methodResults){
        console.log(result.methodResults[idx])
    }
    ```

=== "Go"
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

=== "Java"
    ```java

    ```