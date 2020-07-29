title: Smart Contract Debugging

# Save Transaction Output for Debugging
The goal command-line tool provides functionality to do a test run of a TEAL program using the `goal clerk dryrun` command. This process is described in the [goal TEAL Walkthrough(goal_teal_walkthrough.md)] documentation. From the SDK a logic signature transaction can be written to a file to be used with the `goal clerk dryrun` command. The following code details how this is done. The goal tab illustrates run the `dryrun` on the generated file.

```javascript tab="JavaScript"
    let rawSignedTxn = algosdk.signLogicSigTransaction(txn, lsig);
    fs.writeFileSync("simple.stxn", rawSignedTxn.blob);
```

```python tab="Python"
    logicsig_txn = transaction.LogicSigTransaction(txn, lsig)
    transaction.write_to_file([logicsig_txn], "simple.stxn")
```

```java tab="Java"
    SignedTransaction stx = Account.signLogicsigTransaction(lsig, tx);
    byte[] outBytes = Encoder.encodeToMsgPack(stx);
    try {
        String FILEPATH = "./simple.stxn";
        File file = new File(FILEPATH);
        OutputStream os = new FileOutputStream(file);
        os.write(outBytes);
        os.close();
    }
    catch (Exception e) {
        System.out.println("Exception: " + e);
    }    
```

```go tab="Go"
	txid, stx, err := crypto.SignLogicsigTransaction(lsig, tx)
	if err != nil {
        ...
    }
    f, err := os.Create("simple.stxn")
    if err != nil {
        ...
    }
    defer f.Close()
    if _, err := f.Write(stx); err != nil {
        ...
    }
    if err := f.Sync(); err != nil {
        ...
    }    
```

```text tab="Goal"
$ goal clerk dryrun -t simple.stxn
tx[0] cost=2 trace:
  1 intcblock => <empty stack>
  4 intc_0 => 0 0x0

REJECT
```


# Debugging a TEAL Program
The `-o` option is used to write a signed transaction out to a file and the transaction will not be submitted to the network. This allows testing of the TEAL logic with the `goal clerk dryrun` command which shows how the TEAL is processed and approved or rejected.

```
$ goal clerk send -f C3MKH24QL3GHSD5CDQ47ZNQZMNZRX4MUTV6LVPAXMWAXMIISYSOWPGH674 -a 1000000 -t STF6TH6PKINM4CDIQHNSC7QEA4DM5OJKKSACAPWGTG776NWSQOMAYVGOQE -L mydelegatedsig.lsig -d ~/node/data -o out.stxn
$ goal clerk dryrun -t out.stxn
tx[0] cost=2 trace:
  1 intcblock => <empty stack>
  4 intc_0 => 0 0x0

REJECT
```