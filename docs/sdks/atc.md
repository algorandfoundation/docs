title: Atomic Transaction Composer


Constructing [atomic transactions](../get-details/atomic_transfers.md) requires building up the set of transactions, assigning a group id, and then signing the transactions. While this is not a difficult thing to do on its own, it can be made more difficult if one of the transactions is an Application Call to an [ABI](../get-details/dapps/smart-contracts/ABI/index.md) compliant application. This is because the arguments passed to the application call must be properly encoded and may include transactions or accounts that are part of the application call or group. 

The Atomic Transaction Composer is a convenient way to build out an atomic group of transactions that handles encoding and decoding of ABI arguments and return values. 

!!! Note 
    The following code examples are snippits to demonstrate usage. The full code for the below snippits is available [here](https://github.com/algorand-devrel/demo-abi).


## Create ATC

To use the Atomic Transaction Composer, first initialize the composer: 

=== Python

    ```py
    from algosdk.atomic_transaction_composer import AtomicTransactionComposer

    comp = AtomicTransactionComposer()
    ```

=== JavaScript

    ```js
    import algosdk from 'algosdk'

    const comp = new algosdk.AtomicTransactionComposer()
    ```
===

=== Go

    ```go
	    import "github.com/algorand/go-algorand-sdk/future"
        //...
    	var atc = future.AtomicTransactionComposer{}
    ```
===

=== Java

    ```java
        AtomicTransactionComposer comp = new AtomicTransactionComposer();
    ```
===


## Add individual transactions

Individual transactions being passed to the composer must be wrapped in a `TransactionWithSigner`. This allows some conveniences we'll see later on.

Constructing a Transaction with Signer and adding it to the transaction composer can be done as follows:

=== Python

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
    comp.add_transaction(tws)

    ```

=== JavaScript

    ```js

    //...

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
    comp.addTransaction(tws)

    ```

=== Go 

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

    comp.AddTransaction(tws)

    ```

=== Java

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
    comp.addTransaction(tws);

    ```

The call to add a transaction may may be performed multiple times, each time adding a new transaction to the atomic group. Recall that a maximum of 16 transactions may be included in a single group.


## Calling ABI Methods

=== Python

    ```py

    signer = AccountTransactionSigner(sk)

    comp.add_method_call(app_id, get_method("add"), addr, sp, signer, method_args=[1,1])

    txn = TransactionWithSigner(PaymentTxn(addr, sp, addr, 10000), signer)

    comp.add_method_call(app_id, get_method("txntest"), addr, sp, signer, method_args=[10000, txn, 1000])

    ```

=== JavaScript

    ```js

    // Read in the local contract.json file
    const buff = fs.readFileSync("../contract.json")

    // Parse the json file into an object, pass it to create an ABIContract object
    const contract = new algosdk.ABIContract(JSON.parse(buff.toString()))

    // Utility function to return an ABIMethod by its name
    function getMethodByName(name: string): algosdk.ABIMethod  {
        const m = contract.methods.find((mt: algosdk.ABIMethod)=>{ return mt.name==name })
        if(m === undefined)
            throw Error("Method undefined: "+name)
        return m
    }

    const sp = await client.getTransactionParams().do()
    const commonParams = {
        appID:contract.networks["default"].appID,
        sender:acct.addr,
        suggestedParams:sp,
        signer: algosdk.makeBasicAccountTransactionSigner(acct)
    }

    const comp = new algosdk.AtomicTransactionComposer()

    // Simple ABI Calls with standard arguments, return type
    comp.addMethodCall({
        method: getMethodByName("add"), methodArgs: [1,1], ...commonParams
    })

    // Transaction being passed as an argument, this removes the transaction from the 
    // args list, but includes it in the atomic grouped transaction
    comp.addMethodCall({
        method: getMethodByName("txntest"), 
        methodArgs: [
            10000,
            {
                txn: new Transaction({
                    from: acct.addr,
                    to: acct.addr,
                    amount: 10000,
                    ...sp
                }),
                signer: algosdk.makeBasicAccountTransactionSigner(acct)
            },
            1000
        ], 
        ...commonParams
    })
    ```

=== Go

    ```go

	f, err := os.Open("../contract.json")
	if err != nil {
		log.Fatalf("Failed to open contract file: %+v", err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("Failed to read file: %+v", err)
	}

	contract := &abi.Contract{}
	if err := json.Unmarshal(b, contract); err != nil {
		log.Fatalf("Failed to marshal contract: %+v", err)
	}

	mcp := future.AddMethodCallParams{
		AppID:           contract.Networks["default"].AppID,
		Sender:          acct.Address,
		SuggestedParams: sp,
		OnComplete:      types.NoOpOC,
		Signer:          signer,
	}

	atc.AddMethodCall(combine(mcp, getMethod(contract, "add"), []interface{}{1, 1}))
	// Txn arg, uint return
	txn, _ := future.MakePaymentTxn(acct.Address.String(), acct.Address.String(), 10000, nil, "", sp)
	stxn := future.TransactionWithSigner{Txn: txn, Signer: signer}
	atc.AddMethodCall(combine(mcp, getMethod(contract, "txntest"), []interface{}{10000, stxn, 1000}))

    func getMethod(c *abi.Contract, name string) (m abi.Method) {
        for _, m = range c.Methods {
            if m.Name == name {
                return
            }
        }
        log.Fatalf("No method named: %s", name)
        return
    }

    func combine(mcp future.AddMethodCallParams, m abi.Method, a []interface{}) future.AddMethodCallParams {
        mcp.Method = m
        mcp.MethodArgs = a
        return mcp
    }
    ```

=== Java

    ```java
    ```

## Execution 

Once all the transactions are added to the atomic group the Atomic Transaction Composer allows several ways to perform the transactions. 

    - Build Group will construct the group of transactions and taking care of assigning the group id, returning an array of unsigned TransactionWithSigner objects.
    - Submit will call build group first, then gather the signatures associated to the transactions, then submit the group without blocking. It will return the full list of transaction ids that can be passed to a wait for confirmation function.
    - Execute will perform submit then wait for confirmation given a number of rounds. It will return the resulting confirmed round, list of transaction ids and any parsed ABI return values if relevant.  


=== Python

    ```py
    resp = comp.execute(client, 2)

    for result in dryrun.abi_results:
        print(result.return_value)
    ```

=== JavaScript

    ```js
    const result = await comp.execute(client, 2)
    for(const idx in result.methodResults){
        console.log(result.methodResults[idx])
    }
    ```

=== Go 

    ```go
	ret, err := atc.Execute(client, context.Background(), 2)
	if err != nil {
		log.Fatalf("Failed to execute call: %+v", err)
	}

	for _, r := range ret.MethodResults {
		log.Printf("%s returned %+v", r.TxID, r.ReturnValue)
	}
    ```

=== Java

    ```java
    ```