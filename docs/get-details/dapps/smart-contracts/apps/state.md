title: Contract storage

# Smart contract storage 
Smart Contracts have three different types of storage: [local storage](#local-storage), [global storage](#global-storage), and [box storage](#box-storage). 

Global state and boxes are associated with the app itself, whereas local state is associated with each account that opts into the application. Global and local storage are Key/Value pairs that are limited to 128 bytes per pair. Boxes are keyed storage segments up to 32kb of data per box.

Each storage option’s properties are described below. 

## Global Storage

* Allocation: 
    * Can include between 0 and 64 key/value pairs for a total of 8K of memory to share among them.
    * The amount of global storage is allocated in k/v units, and determined at contract creation. This schema is immutable after creation. 
    * The contract creator address is responsible for funding the global storage (by an increase to their minimum balance requirement, see below). 
* Reading: 
    * Can be read by any app call that has specified app a’s ID in its foreign apps array.
    * Can be read on-chain using the k/v pairs defined (from off-chain, can be read using goal or APIs + SDKs). 
* Writing: 
    * Can only be written by app a. 
* Deletion: 
    * Is deleted when app a is deleted. Cannot otherwise be deallocated (though of course the contents can be cleared by app a, but this does not change the minimum balance requirement). 

## Local Storage

* Allocation: 
    * Is allocated when account x opts in to app a (submits a transaction to opt-in to app a).
    * Can include between 0 and 16 key/value pairs for a total of 2KB of memory to share among them. 
    * The amount of local storage is allocated in k/v units, and determined at contract creation. This cannot be edited later. 
    * The opted-in user address is responsible for funding the local storage (by an increase to their minimum balance). 
* Reading: 
    * Can be read by any app call that has app a in its foreign apps array and account x in its foreign accounts array. 
    * Can be read on-chain using the k/v pairs defined (from off-chain, can be read using goal and the SDKs). 
* Writing: 
    * Is editable only by app a, but is delete-able by app a or the user x (using a ClearState call, see below). 
* Deletion: 
    * Deleting an app does not affect its local storage. Accounts must clear out of app to recover minimum balance.
    * _Clear state_. Every Smart Contract on Algorand has two programs: the _approval_ and the _clear state_ program. An account holder can clear their local state for an app at any time (deleting their data and freeing up their locked minimum balance). The purpose of the clear state program is to allow the app to handle the clearing of that local state gracefully. 
    * Account x can request to clear its local state using a [close out transaction](https://developer.algorand.org/docs/get-details/transactions/#application-close-out-transaction). 
    * Account x can clear its local state for app a using a [clear state transaction](https://developer.algorand.org/docs/get-details/transactions/#application-clear-state-transaction), which will always succeed, even after app a is deleted. 


## Box Storage 

* Allocation: 
    * App a can allocate as many boxes as it needs, when it needs them.
    * App a allocates a box using the `box_create` opcode in its TEAL program, specifying the name and the size of the box being allocated. 
        * Boxes can be any size from 0 to 32K bytes. 
        * Box names must be at least 1 byte, at most 64 bytes, and must be unique within app a. 
    * The app account(the smart contract) is responsible for funding the box storage (with an increase to its minimum balance requirement, see below for details). 
    * A box name and app id must be referenced in the boxes array of the app call to be allocated. 
* Reading: 
    * App a is the only app that can read the contents of its boxes on-chain. This on-chain privacy is unique to box storage. Recall that everything can be read by anybody from off-chain using the algod or indexer APIs. 
    * To read box b from app a, the app call must include b in its boxes array. 
    * Read budget: Each box reference in the boxes array allows an app call to access 1K bytes of box state - 1K of “box read budget”. To read a box larger than 1K, multiple box references must be put in the boxes arrays. 
        * The box read budget is shared across the transaction group. 
        * The total box read budget must be larger than the sum of the sizes of all the individual boxes referenced (it is not possible to use this read budget for a part of a box - the whole box is read in).
    * Box data is unstructured. This is unique to box storage. 
    * A box is referenced by including its app ID and box name. 
* Writing: 
    * App a is the only app that can write the contents of its boxes.
    * As with reading, each box ref in the boxes array allows an app call to write 1kb of box state - 1kb of “box write budget”. 
* Deletion: 
    * App a is the only app that can delete its boxes. 
    * If an app is deleted, its boxes are not deleted. The boxes will not be modifiable but still can be queried using the SDKs. The minimum balance will also be locked. (the correct cleanup design is to look up the boxes from off-chain and call the app to delete all its boxes before deleting the app itself). 


# Manipulate global or local state in smart contract
Smart contracts can create, update, and delete values in global or local state. The number of values that can be written is limited based on how the contract was first created. See [Creating the Smart Contract](#creating-the-smart-contract) for details on configuring the initial global and local storage. State is represented with key-value pairs. The key is limited to 64 bytes. The key plus the value is limited to 128 bytes total. Using smaller keys to have more storage available for the value is possible. The keys are stored as byte slices (byte-array value) and the values are stored as either byte slices (byte-array value) or uint64s. The TEAL language provides several opcodes for facilitating reading and writing to state.

## Reading local state from other accounts
Local storage values are stored in the account's balance record. Any account that sends a transaction to the smart contract can have its local storage modified by the smart contract as long as the account has opted into the smart contract. In addition, any call to the smart contract can also reference up to four additional accounts which can also have their local storage manipulated for the current smart contract as long as the account has opted into the contract. These five accounts can also have their storage values for any smart contract on Algorand read by specifying the application id of the smart contract, if the additional contract is in the applications array for the transaction. This is a read-only operation and does not allow one smart contract to modify the local state of another smart contract. The additionally referenced accounts can be changed per smart contract call (transaction). The process for reading local state from another account is described in the following sections.

## Reading global state from other smart contracts
Global storage for the current contract can also be modified by the smart contract code. In addition, the global storage of any contract in the applications array can be read. This is a read-only operation. The global state can not be changed for other smart contracts. The external smart contracts can be changed per smart contract call (transaction). The process for reading global state from another smart contract is described in the following sections.

## Write to state
To write to either local or global state, the opcodes `app_global_put` and `app_local_put` should be used. These calls are similar but with local storage, you provide an additional account parameter. This determines what account should have its local storage modified. In addition to the sender of the transaction, any call to the smart contract can reference up to four additional accounts. Below is an example of doing a global write. See [Reference arrays](index.md#reference-arrays) for more details.

=== "PyTeal"
	<!-- ===PYTEAL_WRITE_GLOBAL_STATE=== -->
	```python
	    program = App.globalPut(Bytes("Mykey"), Int(50))
	    print(compileTeal(program, Mode.Application))
	```
	[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/kv_state.py#L5-L7)
	<!-- ===PYTEAL_WRITE_GLOBAL_STATE=== -->

=== "TEAL"
	<!-- ===TEAL_WRITE_GLOBAL_STATE=== -->
	```teal
	byte "GlobalKey"
	int 42
	app_global_put
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/state_manipulation/approval.teal#L8-L11)
	<!-- ===TEAL_WRITE_GLOBAL_STATE=== -->

To store a value in local storage, the following contract code can be used.

=== "PyTeal"
	```python
	    program = App.localPut(Txn.sender(), Bytes("MyLocalKey"), Int(50))
	    print(compileTeal(program, Mode.Application))
	```


=== "TEAL"
	```teal
	txn Sender
	byte "OwnLocalKey"
	int 1337
	app_local_put
	```


In this example, the `txn Sender` represents the sender of the transaction. Any account in the accounts array can be specified. See [Reference arrays](index.md#reference-arrays) for more details.


=== "PyTeal"
	```python
	    program = App.localPut(Addr("GHZ..."), Bytes("MyLocalKey"), Int(50))
	    print(compileTeal(program, Mode.Application))
	```


=== "TEAL"
	```teal
	addr GHZ....
	byte "OtherLocalKey"
	int 200
	app_local_put
	```


The account specified must be in the accounts array. See [Reference arrays](index.md#reference-arrays) for more details.

!!! info
    Local storage writes are only allowed if the account has opted into the smart contract.

## Read from state
TEAL provides calls to read global and local state values for the current smart contract.  To read from local or global state TEAL provides the `app_local_get`, `app_global_get`, `app_local_get_ex` , and `app_global_get_ex` opcodes. The following contract code reads a value from global state for the current smart contract.


=== "PyTeal"
	<!-- ===PYTEAL_READ_GLOBAL_STATE=== -->
	```python
	    program = App.globalGet(Bytes("MyGlobalKey"))
	    print(compileTeal(program, Mode.Application))
	```
	[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/kv_state.py#L26-L28)
	<!-- ===PYTEAL_READ_GLOBAL_STATE=== -->

=== "TEAL"
	```teal
	byte "GlobalKey"
	app_global_get		
	```


The following contract code reads the local state of the sender account.

=== "PyTeal"
	```python
	    program = App.localGet(Txn.sender(), Bytes("MyLocalKey"))
	    print(compileTeal(program, Mode.Application))
	```
	

=== "TEAL"
	```teal
	txn Sender
	byte "OwnLocalState"
	app_local_get
	```

The `_ex` opcodes return two values to the stack. The first value is a 0 or a 1 indicating the value was returned successfully or not, and the second value on the stack contains the actual value. These calls allow local and global states to be read from other accounts and applications (smart contracts) as long as the account and the contract are in the accounts and applications arrays. To read a local storage value with the `app_local_get_ex` opcode the following contract code should be used.

=== "PyTeal"
	```python
	    program = App.localGetEx(Txn.sender(), Txn.application_id(), Bytes("MyAmountGiven"))
	    print(compileTeal(program, Mode.Application))
	```

=== "TEAL"
	```teal
	txn Sender
	txn ApplicationID
	byte "MyAmountGiven"
	app_local_get_ex
	```

!!! note
    The PyTeal code snippet preemptively stores the return values from `localGetEx` in scratch space for later reference. 

The `txn Sender` represents the sender of the transaction. The `txn ApplicationID` line refers to the current application, but could be any application that exists on Algorand as long as the contract's ID is in the applications array. See [Reference arrays](index.md#reference-arrays) for more details.  The top value on the stack will either return 0 or 1 depending on if the variable was found.  Most likely branching logic will be used after a call to the `_ex` opcode. The following example illustrates this concept.

=== "PyTeal"
	```python
	    get_amount_given = App.localGetEx(
	        Txn.sender(), Txn.application_id(), Bytes("MyAmountGiven")
	    )
	
	    # Change these to appropriate logic for new and previous givers.
	    new_giver_logic = Seq(Return(Int(1)))
	
	    previous_giver_logic = Seq(Return(Int(1)))
	
	    program = Seq(
	        get_amount_given,
	        If(get_amount_given.hasValue(), previous_giver_logic, new_giver_logic),
	    )
	
	    print(compileTeal(program, Mode.Application))
	```


=== "TEAL"
	```teal
	txn Sender
	txn ApplicationID
	byte "deposited"
	app_local_get_ex
	bz new_deposit
	// Account has deposited before
	
	new_deposit:
	// Account is making their first deposit
	```

The `app_global_get_ex` is used to read not only the global state of the current contract but any contract that is in the applications array. To access these foreign apps, they must be passed in with the application call. See [Reference arrays](index.md#reference-arrays) for more details. 

```bash
$ goal app call --foreign-app APP1ID --foreign-app APP2ID
```

To read from the global state with the `app_global_get_ex` opcode, use the following TEAL.

=== "PyTeal"
	```python
	    get_global_key = App.globalGetEx(Global.current_application_id(), Bytes("MyGlobalKey"))
	
	    # Update with appropriate logic for use case
	    increment_existing = Seq(Return(Int(1)))
	
	    program = Seq(
	        get_global_key,
	        If(get_global_key.hasValue(), increment_existing, Return(Int(1))),
	    )
	
	    print(compileTeal(program, Mode.Application))
	```


=== "TEAL"
	```teal
	global CurrentApplicationID
	byte "GlobalKey"
	app_global_get_ex
	```

The specified contract's ID must be in the applications array. See [Reference arrays](index.md#reference-arrays) for more details. Similar to the `app_local_get_ex` opcode, generally, there will be branching logic testing whether the value was found or not. 

## Summary of global and Local state operations

| Context            | Write            | Read                | Delete           | Check If Exists     |
| ---                | ---              | ---                 | ---              | ---                 |
| Current App Global | `app_global_put` | `app_global_get`    | `app_global_del` | `app_global_get_ex` |
| Current App Local  | `app_local_put`  | `app_local_get`     | `app_local_del`  | `app_local_get_ex`  |
| Other App Global   |                  | `app_global_get_ex` |                  | `app_global_get_ex` |
| Other App Local    |                  | `app_local_get_ex`  |                  | `app_local_get_ex`  |

# Box Details
Boxes are useful in many scenarios:

* Applications that need larger or unbound contract storage.
* Applications that want to store data per user, but do not wish to require users to opt-in to the contract or need the account data to persist even after the user closes or clears out of the application.
* Applications that have dynamic storage requirements.
* Applications that require larger storage blocks that can not fit in the existing global state key-value pairs.
* Applications that require storing arbitrary maps or hash tables. 
  
The following sections cover the details of manipulating boxes within a smart contract. 

## Box Array 
The box array is an array of pairs: the first element of each pair is an integer specifying the index into the foreign application array, and the second element is the key name of the box to be accessed.

Each entry in the box array allows access to only 1kb of data. For example, if a box is sized to 4kb, the transaction must use four entries in this array. To claim an allotted entry a corresponding app Id and box name need to be added to the box ref array. If you need more than the 1kb associated with that specific box name, you can either specify the box ref entry more than once or, preferably, add “empty” box refs `[0,””]` into the array. If you specify 0 as the app Id the box ref is for the application being called. 

For example, suppose the contract needs to read “BoxA” which is 1.5kb, and “Box B” which is 2.5kb, this would require four entries in the box ref array and would look something like:

```py
boxes=[[0, "BoxA"],[0,"BoxB"], [0,""],[0,""]]
``` 

The required box I/O budget is based on the sizes of the boxes accessed, not the amount of data read or written. For example, if a contract accesses “Box A” with a size of 2kb and “Box B” with a size of 10 bytes, this requires both boxes be in the box reference array and one additional reference ( ceil((2kb + 10b) / 1kb), which can be an “empty” box reference. 

Access budgets are summed across multiple application calls in the same transaction group. For example in a group of two smart contract calls, there is room for 16 array entries (8 per app call), allowing access to 16kb of data. If an application needs to access a 16kb box named “Box A”, it will need to be grouped with one additional application call and the box reference array for each transaction in the group should look similar to this:

Transaction 0: [0,”Box A”],[0,””],[0,””],[0,””],[0,””],[0,””],[0,””],[0,””]
Transaction 1: [0,””],[0,””],[0,””],[0,””],[0,””],[0,””],[0,””],[0,””]

Box refs can be added to the boxes array using `goal` or any of the SDKs.

=== "Goal"
    ```goal
    goal app method --app-id=53 --method="add_member2()void" --box="53,str:BoxA" --from=CONP4XZSXVZYA7PGYH7426OCAROGQPBTWBUD2334KPEAZIHY7ZRR653AFY
    ```

=== "Python"
    <!-- ===PYSDK_ATC_BOX_REF=== -->
	```python
	atc = AtomicTransactionComposer()
	atc.add_method_call(
	    app_id,
	    my_method,
	    addr,
	    sp,
	    signer,
	    boxes=[[app_id, b"key"]],
	)
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/atc.py#L78-L87)
    <!-- ===PYSDK_ATC_BOX_REF=== -->

=== "JavaScript"
    <!-- ===JSSDK_ATC_BOX_REF=== -->
	```javascript
	const boxATC = new algosdk.AtomicTransactionComposer();
	const boxKey = new Uint8Array(Buffer.from('key'));
	boxATC.addMethodCall({
	  appID: appIndex,
	  method: boxAccessorMethod,
	  methodArgs: [],
	  boxes: [
	    {
	      appIndex: 0,
	      name: boxKey,
	    },
	  ],
	  sender: sender.addr,
	  signer: sender.signer,
	  suggestedParams,
	});
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/atc.ts#L101-L117)
    <!-- ===JSSDK_ATC_BOX_REF=== -->

=== "Go"
    <!-- ===GOSDK_ATC_BOX_REF=== -->
	```go
	boxName := "coolBoxName"
	mcp = transaction.AddMethodCallParams{
		AppID:           appID,
		Sender:          acct1.Address,
		SuggestedParams: sp,
		OnComplete:      types.NoOpOC,
		Signer:          signer,
		Method:          addMethod,
		MethodArgs:      []interface{}{1, 1},
		// Here we're passing a box reference so our app
		// can reference it during evaluation
		BoxReferences: []types.AppBoxReference{
			{AppID: appID, Name: []byte(boxName)},
		},
	}
	// ...
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/atc/main.go#L95-L111)
    <!-- ===GOSDK_ATC_BOX_REF=== -->

=== "Java"
    <!-- ===JAVASDK_ATC_BOX_REF=== -->
	```java
	MethodCallTransactionBuilder<?> mctBuilder = MethodCallTransactionBuilder.Builder();
	
	List<AppBoxReference> boxRefs = new ArrayList<>();
	boxRefs.add(new AppBoxReference(appId.intValue(), "cool-box".getBytes()));
	MethodCallParams boxRefMcp = mctBuilder
	                .suggestedParams(sp)
	                .applicationId(appId)
	                .sender(acct.getAddress())
	                .method(contract.getMethodByName("add"))
	                .methodArguments(methodArgs)
	                .signer(acct.getTransactionSigner())
	                .onComplete(Transaction.OnCompletion.NoOpOC)
	                // Include reference to a box so the app logic may
	                // use it during evaluation
	                .boxReferences(boxRefs)
	                .build();
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/ATC.java#L98-L114)
    <!-- ===JAVASDK_ATC_BOX_REF=== -->

=== "Beaker"
    ```py
    #Beaker framework
    result = app_client.call(
        Myapp.my_method,
        boxes=[[app_client.app_id, "key"]],
    )
    ```
## Creating a Box
The AVM supports two opcodes `box_create` and `box_put` that can be used to create a box. 
The `box_create` opcode takes two parameters, the name and the size in bytes for the created box. The `box_put` opcode takes two parameters as well. The first parameter is the name and the second is a byte array to write. Because the AVM limits any element on the stack to 4kb, `box_put` can only be used for boxes with length <= 4kb.

=== "TEAL"
	<!-- ===TEAL_BOX_CREATE=== -->
	```teal
	// 100 byte box created with box_create
	byte "Mykey"
	int 100
	box_create
	// ... OR ...
	// create with a box_put
	byte "Mykey"
	byte "My data values"
	box_put
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/boxes/approval.teal#L6-L15)
	<!-- ===TEAL_BOX_CREATE=== -->

=== "PyTeal"
	<!-- ===PYTEAL_BOX_CREATE=== -->
	```python
	        # ...
	        # box created with box_create, size 100 bytes
	        App.box_create(Bytes("MyKey"), Int(100)),
	        # OR box created with box_put, size is implicitly the
	        # length of bytes written
	        App.box_put(Bytes("MyKey"), Bytes("My data values"))
	        # ...
	```
	[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/box.py#L6-L13)
	<!-- ===PYTEAL_BOX_CREATE=== -->

Box names must be unique within an application. If using `box_create`, and an existing box name is passed with a different size, the creation will fail. If an existing box name is used with the existing size, the call will return a 0 without modifying the box contents. When creating a new box the call will return a 1. When using `box_put` with an existing key name, the put will fail if the size of the second argument (data array) is different from the original box size. 

!!!info
    When creating a box, the key name to be created must be in the box ref array.

## Writing to a Box  
The AVM provides two opcodes, `box_put` and `box_replace`,  to write data to a box. The `box_put` opcode is described in the previous section. The `box_replace` opcode takes three parameters, the key name, the starting location and replacement bytes.

=== "TEAL"
	<!-- ===TEAL_BOX_WRITE=== -->
	```teal
	byte "MyKey"
	int 10
	byte "best"
	box_replace
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/boxes/approval.teal#L18-L22)
	<!-- ===TEAL_BOX_WRITE=== -->

=== "PyTeal"
	<!-- ===BEAKER_BOX_WRITE=== -->
    ```py
    #Beaker Framework
        @external
        def replace_string(self, ky: abi.String, start: abi.Uint64, replacement: abi.String, *, output: abi.String):
            return Seq(
                App.box_replace(ky.get(), start.get(), replacement.get()),
                boxstr :=  App.box_get(ky.get()),
                Assert( boxstr.hasValue()),
                output.set(boxstr.value()),
            ) 

    ```
	<!-- ===BEAKER_BOX_WRITE=== -->

When using `box_replace`, the box size can not increase. This means if the replacement bytes, when added to the start byte location, exceed the upper bounds of the box, the call will fail. 

## Reading from a Box
The AVM provides two opcodes for reading the contents of a box, `box_get` and `box_extract`. The `box_get` opcode takes one parameter which is the key name for the box. It reads the entire contents of a box. The `box_get` opcode returns two values. The top-of-stack is an integer that has the value of 1 or 0. A value of 1 means that the box was found and read. A value of 0 means that the box was not found. The next stack element contains the bytes read if the box exists, else it contains an empty byte array. `box_get` fails if the box length exceeds 4kb.

=== "TEAL"
	<!-- ===TEAL_BOX_GET=== -->
	```teal
	byte "MyKey"
	box_get
	assert //verify that the read occurred and we have a value
	//box contents at the top of the stack
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/boxes/approval.teal#L25-L29)
	<!-- ===TEAL_BOX_GET=== -->

=== "PyTeal"
	<!-- ===PYTEAL_BOX_GET=== -->
	```python
	        boxval := App.box_get(Bytes("MyKey")),
	        Assert(boxval.hasValue()),
	        # do something with boxval.value()
	        # ...
	```
	[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/box.py#L21-L25)
	<!-- ===PYTEAL_BOX_GET=== -->

Note that when using either opcode to read the contents of a box, the AVM is limited to reading no more than 4kb at a time. This is because the stack is limited to 4kb entries. For larger boxes, the `box_extract` opcode should be used to perform multiple reads to retrieve the entire contents.

The `box_extract` opcode requires three parameters: the box key name, the starting location, and the length to read. If the box is not found or if the read exceeds the boundaries of the box the opcode will fail. 


=== "TEAL"
	<!-- ===TEAL_BOX_EXTRACT=== -->
	```teal
	byte "BoxA"
	byte "this is a test of a very very very very long string"
	box_put
	
	byte "BoxA"
	int 5
	int 9
	box_extract
	
	byte "is a test"
	==
	assert
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/boxes/approval.teal#L32-L44)
	<!-- ===TEAL_BOX_EXTRACT=== -->

=== "PyTeal"
	<!-- ===PYTEAL_BOX_EXTRACT=== -->
	```python
	        # ...
	        App.box_put(
	            Bytes("BoxA"), Bytes("this is a test of a very very very very long string")
	        ),
	        scratchVar.store(App.box_extract(Bytes("BoxA"), Int(5), Int(9))),
	        Assert(scratchVar.load() == Bytes("is a test"))
	        # ...
	```
	[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/box.py#L33-L40)
	<!-- ===PYTEAL_BOX_EXTRACT=== -->

## Getting a Box Length
The AVM offers the `box_len` opcode to retrieve the length of a box. This opcode can also be used to verify the existence of a particular box. The opcode takes the box key name and returns two unsigned integers (uint64). The top-of-stack is either a 0 or 1, where 1 indicates the existence of the box and 0 indicates the box does not exist. The next is the length of the box if it exists, else it is 0.

=== "TEAL"
	<!-- ===TEAL_BOX_LEN=== -->
	```teal
	byte "BoxA"
	byte "this is a test of a very very very very long string"
	box_put
	
	byte "BoxA"
	box_len
	assert
	
	int 51
	==
	assert
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/boxes/approval.teal#L47-L58)
	<!-- ===TEAL_BOX_LEN=== -->

=== "PyTeal"
	<!-- ===PYTEAL_BOX_LEN=== -->
	```python
	        App.box_put(
	            Bytes("BoxA"), Bytes("this is a test of a very very very very long string")
	        ),
	        # box length is equal to the size of the box created
	        # not a measure of how many bytes have been _written_
	        # by the smart contract
	        bt := App.box_length(Bytes("BoxA")),
	        Assert(bt.hasValue()),
	        Assert(bt.value() == 51),
	```
	[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/box.py#L47-L56)
	<!-- ===PYTEAL_BOX_LEN=== -->

## Deleting a Box
The AVM offers the `box_del` opcode to delete a box. This opcode takes the box key name. The opcode returns one unsigned integer (uint64) with a value of 0 or 1. A value of 1 indicates the box existed and was deleted. A value of 0 indicates the box did not exist.


=== "TEAL"
	<!-- ===TEAL_BOX_DELETE=== -->
	```teal
	byte "BoxA"
	byte "this is a test of a very very very very long string"
	box_put
	
	byte "BoxA"
	box_del
	bnz existed
	```
	[Snippet Source](https://github.com/nullun/algorand-teal-examples/blob/main/_examples/boxes/approval.teal#L61-L68)
	<!-- ===TEAL_BOX_DELETE=== -->

=== "PyTeal"
	<!-- ===PYTEAL_BOX_DELETE=== -->
	```python
	        App.box_put(
	            Bytes("BoxA"), Bytes("this is a test of a very very very very long string")
	        ),
	        # Box delete returns a 1/0 on the stack
	        # depending on if it was successful
	        Assert(App.box_delete(Bytes("BoxA"))),
	```
	[Snippet Source](https://github.com/barnjamin/pyteal/blob/examples-for-docs/_examples/box.py#L63-L69)
	<!-- ===PYTEAL_BOX_DELETE=== -->

!!!warning
    You must delete all boxes before deleting a contract. If this is not done, the minimum balance for that box is not recoverable.

## Example: Storing Named Tuples in a Box
If your contract is using the ABI and authored in PyTeaI, you might want to store a named tuple in a Box. It is preferable that the tuple only contain static data types, as that will allow easy indexing into the box. The following example creates a box for every address that calls the contract’s add_member method. This is an effective way of storing data for every user of the contract without having to have the user’s account opt-in to the contract.

=== "PyTeal"
    ```py
    # This example uses the Beaker framework

    from algosdk import *
    from pyteal import *
    from beaker import *


    class NamedTupleBox(Application):

        class MembershipRecord(abi.NamedTuple):
            role: abi.Field[abi.Uint8]
            voted: abi.Field[abi.Bool]

        
        @external
        def add_member(self, role: abi.Uint8, voted: abi.Bool,*, output: MembershipRecord):
            return Seq(
                output.set(role, voted),
                App.box_put(Txn.sender(), output.encode()),
            )

        @external
        def del_member(self,*, output: abi.Uint64):
            return Seq(
                output.set(App.box_delete(Txn.sender())),
            )     
            
    if __name__ == "__main__":
        accts = sandbox.get_accounts()
        acct = accts.pop()


        app_client = client.ApplicationClient(
            sandbox.get_algod_client(), NamedTupleBox(), signer=acct.signer
        )

        app_client.create()
        app_client.fund(100 * consts.algo)
        print("APP ID")
        print(app_client.app_id)
        print(acct.address)
        ls = acct.address.encode()

        result = app_client.call(
            NamedTupleBox.add_member,
            role=2,
            voted=False,
            boxes=[[app_client.app_id, encoding.decode_address(acct.address)]],
        )
        result = app_client.call(
            NamedTupleBox.del_member,
            boxes=[[app_client.app_id, encoding.decode_address(acct.address)]],
        )    

        print(result.return_value)
        NamedTupleBox().dump('./artifacts')
    ```

# Reading a smart contracts state
In addition to being able to read the state of a smart contract using TEAL, these global and local values can be read externally with the SDKs and `goal`. These reads are not transactions and just query the current state of the contract. 

```bash
$ goal app read --app-id 1 --guess-format --global --from [ADDRESS]
```

In the above example, the global state of the smart contract with the application ID of 1 is returned. The `--guess-format` opt in the above example tries programmatically to display the properly formatted values of the state variables. To get the local state, replace `--global` with `--local` and note that this call will only return the local state of the `--from` account.

Here is an example output with 3 keys/values:

```json
{
  "Creator": {
    "tb": "FRYCPGH25DHCYQGXEB54NJ6LHQG6I2TWMUV2P3UWUU7RWP7BQ2BMBBDPD4",
    "tt": 1
  },
  "MyBytesKey": {
    "tb": "hello",
    "tt": 1
  },
  "MyUintKey": {
    "tt": 2,
    "ui": 50
  }
}
```

Interpretation:

* the keys are `Creator`, `MyBytesKey`, `MyUintKey`.
* the field `tt` is the type of the value: 1 for byte slices (byte-array value), 2 for uint.
* when `tt=1`, the value is in the field `tb`. Note that because of `--guess-format`, the value for `Creator` is automatically converted to an Algorand address with checksum (as opposed to a 32-byte public key.
* when `tt=2`, the value is in the field `ui`.
    