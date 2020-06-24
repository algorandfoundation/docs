title: Atomic Transfers

In traditional finance, trading assets generally requires a trusted intermediary, like a bank or an exchange, to make sure that both sides receive what they agreed to. On the Algorand blockchain, this type of trade is implemented within the protocol as an **Atomic Transfer**. This simply means that transactions that are part of the transfer either all succeed or all fail. Atomic transfers allow complete strangers to trade assets without the need for a trusted intermediary, all while guaranteeing that each party will receive what they agreed to. 

On Algorand, atomic transfers are implemented as irreducible batch operations, where a group of [transactions](./transactions/index.md) are submitted as a unit and all transactions in the batch either pass or fail. This also eliminates the need for more complex solutions like [hashed timelock contracts](https://en.bitcoinwiki.org/wiki/Hashed_Timelock_Contracts) that are implemented on other blockchains. An atomic transfer on Algorand is confirmed in less than 5 seconds, just like any other transaction. Transactions can contain Algos or Algorand Standard Assets and may also be governed by Algorand Smart Contracts. 

# Use Cases

Atomic transfers enable use cases such as:

**Circular trades** - Alice pays Bob if and only if Bob pays Claire if and only if Claire pays Alice.

**Group payments** - Everyone pays or no one pays.

**Decentralized exchanges** - Trade one asset for another without going through a centralized exchange. 

**Distributed payments** - Payments to multiple recipients. 

# Process Overview

To implement an atomic transfer, generate all of the transactions that will be involved in the transfer and then group those transactions together. The result of grouping is that each transaction is assigned the same group ID. Once all transactions contain this group ID, the transactions can be split up and sent to their respective senders to be authorized. A single party can then collect all the authorized transactions and submit them to the network together. 

!!! info
    An individual account involved in an atomic transfer, can verify that all the correct transfers are involved by creating the same set of [unsigned] transactions and grouping them in the same order. The group ID is a hash of the group of transactions and should match if the configuration is the same.

<center>![Atomic Transfer Flow](../imgs/atomic_transfers-1.png)</center>
<center>*Atomic Transfer Flow*</center>

Below you will find examples for creating and sending group transactions to the network in each of the available SDKs and with `goal`. 

# Step-by-Step Guide

## Create Transactions
Create two or more (up to 16 total) unsigned transactions of any type. Read about transaction types in the [Transactions Overview](./transactions/index.md) section. 

This could be done by a service or by each party involved in the transaction. For example, an asset exchange application can create the entire atomic transfer and allow individual parties to sign from their location.

The example below illustrates Account A sending a transaction to Account C and Account B sending a transaction to Account A.

!!! info
    The examples in this section have been updated to the v2 API, which was launched to MainNet on June 16, 2020. Visit the [v2 Migration Guide](../reference/sdks/migration.md) for information on how to migrate your code from v1. 

    Full running code examples for each SDK and both API versions are available within the GitHub repo at [/examples/atomic_transfers](https://github.com/algorand/docs/tree/master/examples/atomic_transfers) and for [download](https://github.com/algorand/docs/blob/master/examples/atomic_transfers/atomic_transfers.zip?raw=true) (.zip).

``` javascript tab="JavaScript"
// Transaction A to C 
let transaction1 = algosdk.makePaymentTxnWithSuggestedParams(myAccountA.addr, myAccountC.addr, 1000000, undefined, undefined, params);  
// Create transaction B to A
let transaction2 = algosdk.makePaymentTxnWithSuggestedParams(myAccountB.addr, myAccountA.addr, 2000000, undefined, undefined, params);  
    
```

``` python tab="Python"  
# from account 1 to account 3
txn_1 = transaction.PaymentTxn(account_1, params, account_3, 1000000)

# from account 2 to account 1
txn_2 = transaction.PaymentTxn(account_2, params, account_1, 2000000)

```

``` java tab="Java"
// Create the first transaction
Transaction tx1 = Transaction.PaymentTransactionBuilder()
.sender(acctA.getAddress())
.amount(1000000)
.receiver(acctC.getAddress())
.suggestedParams(params)
.build();

// Create the second transaction
Transaction tx2 = Transaction.PaymentTransactionBuilder()
.sender(acctB.getAddress())
.amount(2000000)
.receiver(acctA.getAddress())
.suggestedParams(params)
.build();

```

``` go tab="Go"
// from account 1 to account 3
tx1, err := transaction.MakePaymentTxnWithFlatFee(account1, account3, minFee, 1000000, firstValidRound, lastValidRound, nil, "", genID, genHash)
if err != nil {
    fmt.Printf("Error creating transaction: %s\n", err)
    return
}

// from account 2 to account 1
tx2, err := transaction.MakePaymentTxnWithFlatFee(account2, account1, minFee, 2000000, firstValidRound, lastValidRound, nil, "", genID, genHash)
if err != nil {
    fmt.Printf("Error creating transaction: %s\n", err)
    return
}

```

``` goal tab="goal"
$ goal clerk send --from=my-account-a<PLACEHOLDER> --to=my-account-c<PLACEHOLDER> --fee=1000 --amount=1000000 --out=unsginedtransaction1.txn"

$ goal clerk send --from=my-account-b<PLACEHOLDER> --to=my-account-a<PLACEHOLDER> --fee=1000 --amount=2000000 --out=unsginedtransaction2.txn"
```

At this point, these are just individual transactions. The next critical step is to combine them and then calculate the group ID.

See [Authorizing Transactions Offline](./transactions/offline_transactions.md#saving-unsigned-transactions-to-file) to learn how to create and save individual **unsigned** transactions to a file. This method can be used to distribute group transactions for signing.

## Combine Transactions 
Combining transactions just means concatenating them into a single file or ordering them in an array so that a group ID can then be assigned. 

If using `goal`, the transaction files can be combined using an OS-level command such as `cat`. If using one of the SDKs, the application may store all the transactions individually or in an array. From the SDK it is also possible to read a transaction from a file created at an earlier time, which is described in the [Offline Transactions](./transactions/offline_transactions.md) documentation. See the complete example at the bottom of this page that details how transactions are combined in the SDKs. To combine transactions in `goal` use a similar method to the one below.

``` javascript tab="JavaScript"
// Combine transactions
let txns = [transaction1, transaction2]
```

``` python tab="Python"
# the Python SDK performs combining implicitly within grouping below

```

``` java tab="Java"
// the Go SDK performs combining implicitly within grouping below

```

``` go tab="Go"
// the Go SDK performs combining implicitly within grouping below

```

``` goal tab="goal"
$ cat unsignedtransaction1.tx unsignedtransaction2.tx > combinedtransactions.tx
```

## Group Transactions

The result of this step is what ultimately guarantees that a particular transaction belongs to a group and is not valid if sent alone (even if properly signed). A group-id is calculated by hashing the concatenation of a set of related transactions. The resulting hash is assigned to the [Group](../reference/transactions.md#group) field within each transaction. This mechanism allows anyone to recreate all transactions and recalculate the group ID to verify that the contents are as agreed upon by all parties. Ordering of the transaction set must be maintained.

``` javascript tab="JavaScript"
// Group both transactions
let txgroup = algosdk.assignGroupID(txns);
```

``` python tab="Python"
# get group id and assign it to transactions
gid = transaction.calculate_group_id([txn1, txn2])
txn1.group = gid
txn2.group = gid
```

``` java tab="Java"
// group transactions an assign ids
Digest gid = TxGroup.computeGroupID(new Transaction[]{tx1, tx2});
tx1.assignGroupID(gid);
tx2.assignGroupID(gid);
```

``` go tab="Go"
// compute group id and put it into each transaction
gid, err := crypto.ComputeGroupID([]types.Transaction{tx1, tx2})
tx1.Group = gid
tx2.Group = gid
```

``` goal tab="goal"
$ goal clerk group -i yourwalletcombinedtransactions.tx -o groupedtransactions.tx -d data -w 
```

## Split Transactions

At this point the transaction set must be split to allow distributing each component transaction to the appropriate wallet for authorization. 

``` javascript tab="JavaScript"
// this example does not use files on disk, so splitting is implicit above
```

``` python tab="Python"
# this example does not use files on disk, so splitting is implicit above
```

``` java tab="Java"
// this example does not use files on disk, so splitting is implicit above
```

``` go tab="Go"
// this example does not use files on disk, so splitting is implicit above
```

``` goal tab="goal"
# keys in distinct wallets
$ goal clerk split -i groupedtransactions.tx -o splitfiles -d data -w yourwallet 

Wrote transaction 0 to splitfiles-0
Wrote transaction 1 to splitfiles-1

# distribute files for authorization
```

## Sign Transactions
With a group ID assigned, each transaction sender must authorize their respective transaction. 

``` javascript tab="JavaScript"
// Sign each transaction in the group 
signedTx1 = transaction1.signTxn( myAccountA.sk )
signedTx2 = transaction1.signTxn( myAccountB.sk )

```

``` python tab="Python"
# sign transactions
stxn_1 = txn_1.sign(sk_1)    
stxn_2 = txn_2.sign(sk_2)

```

``` java tab="Java"
// sign individual transactions
SignedTransaction signedTx1 = acctA.signTransaction(tx1);;
SignedTransaction signedTx2 = acctB.signTransaction(tx2);;

```        

``` go tab="Go"
sTxID1, stx1, err := crypto.SignTransaction(sk1, tx1)
if err != nil {
    fmt.Printf("Failed to sign transaction: %s\n", err)
    return
}
sTxID2, stx2, err := crypto.SignTransaction(sk2, tx2)
if err != nil {
    fmt.Printf("Failed to sign transaction: %s\n", err)
    return
}

```

``` goal tab="goal"
# sign from single wallet containing all keys
$ goal clerk sign -i groupedtransactions.tx -o signout.tx -d data -w yourwallet

# -- OR --

# sign from distinct wallets
$ goal clerk sign -i splitfiles-0 -o splitfiles-0.sig -d data -w my_wallet_1
$ goal clerk sign -i splitfiles-1 -o splitfiles-1.sig -d data -w my_wallet_2

```

## Assemble Transaction Group

All authorized transactions are now assembled into an array, maintaining the original transaction ordering, which represents the transaction group.

``` javascript tab="JavaScript"
let signed = []
signed.push( signedTx1 )
signed.push( signedTx2 )

```

``` python tab="Python"
# assemble transaction group
signedGroup =  []
signedGroup.append(stxn_1)
signedGroup.append(stxn_2)
```

``` java tab="Java"
// put both transaction in a byte array 
ByteArrayOutputStream byteOutputStream = new ByteArrayOutputStream( );
byte[] encodedTxBytes1 = Encoder.encodeToMsgPack(signedTx1);
byte[] encodedTxBytes2 = Encoder.encodeToMsgPack(signedTx2);
byteOutputStream.write(encodedTxBytes1);
byteOutputStream.write(encodedTxBytes2);
byte groupTransactionBytes[] = byteOutputStream.toByteArray();
```

``` go tab="Go"
var signedGroup []byte
signedGroup = append(signedGroup, stx1...)
signedGroup = append(signedGroup, stx2...)
```

``` goal tab="goal"
# combine signed transactions files
cat splitfiles-0.sig splitfiles-1.sig > signout.tx
```

## Send Transaction Group
The transaction group is now broadcast to the network. 

``` javascript tab="JavaScript"
let tx = (await algodClient.sendRawTransaction(signed).do());
console.log("Transaction : " + tx.txId);

// Wait for transaction to be confirmed
await waitForConfirmation(algodClient, tx.txId)
```

``` python tab="Python"
tx_id = algod_client.send_transactions(signedGroup)

# wait for confirmation
wait_for_confirmation(algod_client, tx_id) 
```

``` java tab="Java"
String id = client.RawTransaction().rawtxn(groupTransactionBytes).execute().body().txId;
System.out.println("Successfully sent tx with ID: " + id);
// write transaction to node
waitForConfirmation(id);
```

``` go tab="Go"
pendingTxID, err := algodClient.SendRawTransaction(signedGroup).Do(context.Background())
if err != nil {
    fmt.Printf("failed to send transaction: %s\n", err)
    return
}
waitForConfirmation(pendingTxID, algodClient)
```

``` goal tab="Goal"
goal clerk rawsend -f signout.tx -d data -w yourwallet
```

!!! info
    Full running code examples for each SDK and both API versions are available within the GitHub repo at [/examples/atomic_transfers](https://github.com/algorand/docs/tree/master/examples/atomic_transfers) and for [download](https://github.com/algorand/docs/blob/master/examples/atomic_transfers/atomic_transfers.zip?raw=true) (.zip).
