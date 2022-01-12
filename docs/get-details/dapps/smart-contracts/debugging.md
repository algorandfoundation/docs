title: Smart contract debugging

Smart contracts can be debugged using two different methods. The first is an interactive debugger that uses the `tealdbg` command-line tool to launch a debug session where the smart contract can be examined as the contract is being evaluated. The second method uses the `goal clerk dryrun-remote` command which outputs a line by line result of the evaluation of the smart contract.  These two methods are described below.

# Using the TEAL debugger
Algorand provides the `tealdbg` command-line tool to launch an interactive session to debug smart contracts. This tool is explained in the project’s [README](https://github.com/algorand/go-algorand/blob/master/cmd/tealdbg/README.md).

This debugger can debug local smart contracts or connect remotely to an on-chain smart contract. The examples below illustrate using the debugger locally, which will be the predominant use case for developers when they are developing smart contracts. For more information on the `tealdbg` utility and how you can use it remotely, see the project’s [README](https://github.com/algorand/go-algorand/blob/master/cmd/tealdbg/README.md) which covers in detail the different debugging options. The `tealdbg` command can also be called through a [sandbox](https://github.com/algorand/sandbox) with `./sandbox tealdbg {OPTIONS}`.

The debugger process supports both Chrome Developer Tools and a simple Web Frontend.

To launch the debugger locally, for use with the CDT, execute the following command from a terminal.

```
$ tealdbg debug program.teal
.
.
2020/08/25 14:05:38 CDT debugger listening on: ws://localhost:9392/091d04a69152223ae84c8b40271c3d62f8490ea9b795aae95868932163f89735
2020/08/25 14:05:38 Or open in Chrome:
2020/08/25 14:05:38 devtools://devtools/bundled/js_app.html?experiments=true&v8only=false&ws=localhost:9392/091d04a69152223ae84c8b40271c3d62f8490ea9b795aae95868932163f89735
```

This will launch the debugger process and return an endpoint that is listening for connections. This process can be connected to directly with the Chrome Developer Tools (CDT). The simplest way to do this is to enter `chrome://inspect/` into the address bar of the browser, click “Configure” to add “localhost:9392”, and select the Algorand TEAL Debugger in the Remote Target Section (click on the inspect link).

!!! note "A note about ports"
    You may have to pass a specific port to the tealdbg command line tool to prevent it from trying to use a port that is already being used. For example if you have the sandbox running and you'd like to run the tealdbg on the host machine.
    ```
     tealdbg debug program.teal --remote-debugging-port 9393
    ```

![CDT Remote Connection](../../../imgs/tealdbg-1.png)
<center>*Configure CDT Remote Connection*</center>

This will launch the debugger and allow the smart contract to be inspected. The debugger provides standard debugger controls, like single-stepping, pausing, breakpoints etc.

!!! note "A note on viewing TEAL source"
    The TEAL source file may not open immediately but you can bring up a menu with the source files by pressing CTRL+P or CMD+P on mac 

![Teal Debugger](../../../imgs/tealdbg-2.png)
<center>*TEAL Debugger*</center>

The Scope pane contains the current stack and is useful for determining what values the current line of code is processing. When a smart contract returns, if anything other than one positive value is left on the stack or the `return` opcode is used with a nonpositive value on the top of the stack, the program will fail. The Scope pane also displays the current transaction with all its properties, the current scratch space, global variables, and any state associated with the contract. If this transaction is part of an atomic transfer, all transactions will also be available in the Scope pane.

The debugging session runs in a mock version of the ledger so the context including; balance records of the accounts, application parameters, and assets used during program evaluation must be supplied.  Exactly what context is needed for a debug session will depend on the type of contract that is being debugged and what opcodes are used in the program. 

Most frequently, a `Dryrun Dump` file is used to pass all this context information in a single payload. 

This file may be msgpack or json and can be created using goal or the SDKs


=== "Python"
    ```py

    app_txn = transaction.ApplicationCallTxn(addr, sp, app_id, transaction.OnComplete.NoOpOC, accounts=[other_addr], foreign_assets=[asset_id])
    s_app_txn = app_txn.sign(sk)

    drr = transaction.create_dryrun(client, [s_app_txn])

    filename = "dryrun.msgp"
    with open(filename, "wb") as f:
        import base64
        f.write(base64.b64decode(encoding.msgpack_encode(drr)))

    ```

=== "JavaScript"
    ```js

    const app_txn = algosdk.makeApplicationNoOpTxn(addr, sp, app_id, undefined, [other_addr], undefined, undefined)
    const s_app_txn = algosdk.signTransaction(app_txn, sk)

    const drr = await algosdk.createDryrun({
        client: client, 
        txns: [
            algosdk.decodeSignedTransaction(s_app_txn['blob']),
        ]
    })

    const filename = 'dryrun.msgp'
    fs.writeFileSync(filename, algosdk.encodeObj(drr.get_obj_for_encoding(true)))

    ```

=== "Go"
    ```go

	app_txn, err := future.MakeApplicationNoOpTx(app_id, nil, []string{other_addr}, nil, []uint64{asset_id}, sp, addr, nil, types.Digest{}, [32]byte{}, types.Address{})
	if err != nil {
		log.Fatalf("Failed to create app call txn: %+v", err)
	}

    _, s_app_bytes, err := crypto.SignTransaction(sk, app_txn)
	if err != nil {
		log.Fatalf("Failed to sign app txn: %+v", err)
	}
	s_app_txn := types.SignedTxn{}
	msgpack.Decode(s_app_bytes, &s_app)

    drr, err := future.CreateDryrun(client, []types.SignedTxn{s_app_txn}, nil, context.Background())
	if err != nil {
		log.Fatalf("Failed to create dryrun: %+v", err)
	}

	filename := "dryrun.msgp"
	os.WriteFile(filename, msgpack.Encode(drr), 0666)

    ```

=== "Java"
    ```java

    Transaction app_txn = ApplicationCallTransactionBuilder.Builder().sender(addr)
        .suggestedParams(sp).applicationId(app_id).foreignAssets(fassets).accounts(addrs).build();


    List<SignedTransaction> stxns = new ArrayList<>();
    stxns.add(acct.signTransaction(app_txn));

    DryrunRequest drr = Utils.createDryrun(client, stxns, "", 0L, 0L);

    String fname = "dryrun.msgp";
    FileOutputStream outfile = new FileOutputStream(fname);
    outfile.write(Encoder.encodeToMsgPack(drr));
    outfile.close();

    ```

=== "goal"
    ```sh

    $ goal app call --app-id {appid} --from {ACCOUNT} --out=dryrun.json --dryrun-dump

    ```



This can be supplied directly by calling `tealdbg` with arguments including: 


    - Transaction(s) - Supply one or more transactions to the debug session.
    - Balance Records - Account balance records of accounts used in the contract. Needed when debugging smart contracts that make use of state. 
    - Latest Timestamp - Set the latest timestamp for debugging smart contracts that make use of time values.
    - Protocol - Set the protocol of consensus that the debug session uses when evaluating the contract.
    - Round - Set the current round for the debug session.
    Group Index - If using more than one transaction, the specific index that is being processed can be set.
    - Mode - The debugger can debug both smart signatures (Signature) and smart contracts (Application). The mode can be set to determine which type is being processed. By default, the debugger scans the program to determine the type of contract.
    - Application ID - Manually set the application ID of the current smart contract.
    - Program - Pass the program to debug

To learn more about setting individual context items, see the `tealdbg` [documentation](https://github.com/algorand/go-algorand/blob/master/cmd/tealdbg/README.md).

The context can also be read directly from an instance of the Indexer. See the `tealdbg` [documentation](https://github.com/algorand/go-algorand/blob/master/cmd/tealdbg/README.md) for more information.


## Calling the debugger with context

```
$ tealdbg debug -d dryrun.msgp
```

!!! note "a note on supplying the Teal code"
    The `tealdbg` command does not require the Teal program above as the decompiled version is available in the dryrun-req dump file. In this example it is supplied for easier readability while in the debugger. Supplying the program will allow debugging the original source and not the decompiled version.

The scope panel will now have the proper context data for the debugging session.

![Teal Debugger Scope](../../../imgs/tealdbg-3.png)
<center>*TEAL Debugger Scope*</center>

One or more transactions that are stored in a file can be debugged by dumping the context data with the `goal clerk dryrun` command. For example, two smart contracts are grouped below and the context data is generated using the `dryrun-dump` option. The `tealdbg` command is then used to start the debugger.

```
#generate calls
$ goal app call --app-id 1  --app-arg "str:test1" --from {ACCOUNT} --out=unsginedtransaction1.tx  
$ goal app call --app-id 1  --app-arg "str:test2" --from {ACCOUNT} --out=unsginedtransaction2.tx

# atomically group them
$ cat unsginedtransaction1.tx unsginedtransaction2.tx  > combinedtransactions.tx
$ goal clerk group -i combinedtransactions.tx -o groupedtransactions.tx 
$ goal clerk split -i groupedtransactions.tx -o split.tx 

# sign individual transactions
$ goal clerk sign -i split-0.tx -o signout-0.tx
$ goal clerk sign -i split-1.tx -o signout-1.tx 

# concatenate the two signed transactions
$ cat signout-0.tx signout-1.tx  > signout.tx

# generate context debug file
$ goal clerk dryrun -t signout.tx --dryrun-dump  -o dryrun.json

# debug first transaction. Change index to 1 to debug second transaction
$ tealdbg debug program.teal -d dryrun.json --group-index 0
```

A similar flow may be implemented in any of the sdks, passing the list of transactions to the `create_dryrun` function will produce the DryrunDump object and either write the object out to a file as shown above or send it directly to the [/v2/teal/dryrun](/rest-apis/algod/v2/#post-v2tealdryrun) endpoint

Debugging a smart signature functions identically to the process described above except the state is not required. For example, a smart signature may act as an escrow account. The following call exports the transaction for debugging purposes. This call will not execute on the blockchain as it is not submitted to the network but is written to the output file.

```
$ goal clerk send -a=0 -c={ACCOUNT_1} --to={ACCOUNT_2} --from-program=stateless.teal --out=statelesstx.tx --dryrun-dump
```

This contract can then be debugged by using the following command.

```
$ tealdbg debug stateless.teal -d statelesstx.tx
```

The `tealdbg` utility has many more options for setting specific context items. For information on these capabilities, see the project’s [README](https://github.com/algorand/go-algorand/blob/master/cmd/tealdbg/README.md).


# Using dryrun for debugging a TEAL program

The SDKs and the `goal` command-line tool provide the functionality to do a test run of a TEAL smart contract.


!!! important "enabling dryrun-remote"
    The dryrun REST API is only available on a node if it has been enabled in the node's configuration. This can be done using the following commands.

    ```
    $ algocfg set -p EnableDeveloperAPI -v true
    $ goal node restart
    ```