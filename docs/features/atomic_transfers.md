Atomic Transfers are irreducible batch transactions that allow groups of transactions to be submitted in one operation. If any of the transactions fail, then all the transactions will fail. Transactions can contain Alogs, Algorand Assets and Algorand Smart Contracts. 

The implementation allows individual transactions to be created and then grouped into a file. Each transfer is signed separately by the originator of the transfer. This file is then submitted to the network, where the node verifies the transactions and submits them all at once. This eliminates the need for hashed time-locks for atomic swaps in other blockchains and also avoids the delays inherent with hashed timed-locks.
Atomic Transfers enables applications such as: 

* Circular trades: Alice pays Bob if and only if Bob pays Clare if and only if Clare pays Alice
* Group payments: Either everyone pays or no one does.
* Decentralized exchanges: Atomic multi-party transfers enable trades without trusted intermediaries.
* Internal units of accounting: Several companies can settle their accounts with each other with a single transaction.
* Payments to multiple recipients


Atomic Transfers are Created with the following steps:

* Create unsigned transactions and save to file
* Combine transactions into file
* Group transactions
* Sign the grouped transaction with private keys
* Submit signed grouped transaction file to network 

<center>![Atomic Transfer Flow](../../imgs/atomic_transfers-1.png)</center>
<center>*Atomic Transfer Flow*</center>


Below you will find examples for creating group transactions and sending them to the network in each of the available SDKs. The same capability is also shown using the `goal` command line tool. The example code is separated into snippets categorized by these core functions.
# Create transactions
Transaction creation functions are described in the Transaciton documentation<LINK>. Atomic Transfers are created by constructing two or more transactions that are not signed and are written to a file. By not signing the transaction, this allows atomic transfers to be created by one or more parties, possibly at different times. For example, an asset exchange application can create the entire atomic transfer and allow individual parties to sign from their location.

The SDKs or `goal` can be used to create unsigned transactions. 

``` javascript tab="JavaScript"
```

``` python tab="Python"
```

``` java tab="Java"
```

``` go tab="Go"
```

``` goal tab="goal"
```


# Group transactions

Compute the group id of the transactions and assign it the each transaction using tx.Group = gid

``` python tab="Python"
import tensorflow as tf
# get group id and assign it to transactions
gid = transaction.calculate_group_id([txn1, txn2])
txn1.group = gid
txn2.group = gid

```

``` Go tab="Go"
// compute group id and put it into each transaction
gid, err := crypto.ComputeGroupID([]types.Transaction{tx1, tx2})
tx1.Group = gid
tx2.Group = gid

```

``` Java tab=
// compute group id and put it into each transaction
gid, err := crypto.ComputeGroupID([]types.Transaction{tx1, tx2})
tx1.Group = gid
tx2.Group = gid

```

``` Javascript tab=
// compute group id and put it into each transaction
gid, err := crypto.ComputeGroupID([]types.Transaction{tx1, tx2})
tx1.Group = gid
tx2.Group = gid

```

# Sign transactions
# Send transactions
