Below you will find examples for creating group transactions and sending them to the network in each of the available SDKs. The example code is separated into snippets categorized by these core functions.
# Create transactions
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
