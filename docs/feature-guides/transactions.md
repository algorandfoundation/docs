title: Transactions

In the getting started section<LINK>, you created your first transaction on Algorand and were introduced to the pattern of constructing a transaction, authorizing it, sending it to the network, and validating its inclusion in a block. This section deals in depth with the construction portion of that pattern. 

It begins with a deeper look into the anatomy of a transaction on Algorand and then looks at the five different types of Algorand transactions and how they differ. 

# Fields common to all transactions
There are several parameters that are common to all transactions on Algorand.

## Sender
*required*

The address of the account that will pay the `amount` and `fee` and authorize the transaction. 

## Fee
*required*

An amount paid by the sender to the FeeSink to prevent denial-of-service. The minimum fee on Algorand is currently 1000 microAlgos. Read more about [how to set the optimal fee for your transaction](#determining-the-optimal-fee).

## First Valid Round
*required*

The starting round for when the transaction is valid. If the transaction is sent prior to this round it will be rejected by the network.

## Last Valid Round
*required*

The ending round for which the transaction is valid. After this round, the transaction will be rejected by the network.

!!!	info
  	The maximum validity range is 1000, i.e. it must be the case that **lastValidRound** - **firstValidRound** <= 1000.

## Genesis Hash
*required*

The SHA512/256 hash of the genesis block of the network for which the transaction is valid. See the genesis hash for [MainNet](../algorand-networks/mainnet.md), [TestNet](../algorand-networks/testnet.md), and [BetaNet](../algorand-networks/betanet.md).

## Genesis ID
*optional*

The human-readable string that identifies the network for the transaction. The network ID is specified in the genesis block. See the genesis ID for [MainNet](../algorand-networks/mainnet.md), [TestNet](../algorand-networks/testnet.md), and [BetaNet](../algorand-networks/betanet.md).

## Note
*optional*

All transactions have an optional note field that allows up to 1000 bytes of data.

## Lease
*optional*

A lease is generally used in the context of Algorand Smart Contracts to prevent replay attacks. Read more about [Algorand Smart Contracts](../feature-guides/asc1/index.md) and see the Delegate [Key Registration TEAL template](../reference-docs/teal/templates/delegate_keyreg.md) for an example implementation of leases.

## Group ID
*optional*

A group ID is assigned to a transaction through the workflow described in the [Atomic Transfers Guide](../feature-guides/atomic_transfers.md). The existence of a group ID, signals that the transaction is only valid when executed alongside the other transactions that contain the same group ID. 



# Transaction Types
The following sections describe the 5 types of transactions on Algorand. Create these transactions with the SDKs or with `goal`. Transactions are encoded in msgpack before they are sent to the network but this is generally masked by the developer tools. Write transactions to a file and use `goal clerk inspect <var>filename</var>` to view the transaction in a human-readable way.

## Payment
A `PaymentTxn` is the most basic type of transaction on Algorand. With it, you can send Algos (the Algorand blockchain's native currency) from one account to another.

### Important Fields
#### Amount 
*required*

This is the total number of microAlgos to be sent from the sending account. 

#### Receiver

*required*

The address of the account that will receive the [amount](#amount).

#### Close Remainder To

*optional*

The address of the account where the remaining balance of the [sender](#sender) account will be sent. Use this field to *close* the account that is sending the Algos. Closing an account means that the remaining balance, after the [amount](#amount) and [fee](#fee) are paid, will be transferred to the specified **Close Remainder To** address and the **Sender** account will be removed from the ledger.

### Example - Send 5 Algos
Here is an example transaction that sends 5 Algos from one account to another on TestNet followed by an explanation of what each of the fields mean. The output below was generated with `goal clerk inspect`.

```json
{
  "txn": {
    "amt": 5000000,
    "fee": 1000,
    "fv": 6000000,
    "gen": "testnet-v1.0",
    "gh": "SGO1GKSzyE7IEPItTxCByw9x8FmnrCDexi9/cOUJOiI=",
    "lv": 6001000,
    "note": "SGVsbG8gV29ybGQ=",
    "rcv": "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A",
    "snd": "EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4",
    "type": "pay"
  }
}
```
From top to bottom, this transaction specifies an [amount (`amt`)](#amount) of 5 million microAlgos (i.e. 5 Algos). The [`fee`](#fee) is 1000 microAlgos. `fv` refers to the [FirstValidRound](#first-valid-round), set for round 6 million. If this transaction is sent prior to round 6 million, it will be rejected. The `gen` field is the optional [Genesis ID](#genesis-id). `gh` is the [Genesis Hash](#genesis-hash) in base64 encoding. This transaction is only valid on [TestNet](../algorand-networks/testnet.md) which can be verified by matching the hashes. `lv` refers to the [Last Valid Round](#last-valid-round) which is set to 1,000 rounds after the first valid round, i.e. the maximum range. The optional `note` field is a base64 representation of the string "Hello World". The base64 encoding is the representation returned by `goal clerk inspect`. The [Receiver](#receiver) (`rcv`) address is the TestNet faucet address. The [Sender](#sender) (`snd`) address is a random generated account. Finally, the `type` field is filled in by the developer tools. The value in this transaction is `pay` which maps to the `PaymentTxn` [type](https://github.com/algorand/go-algorand/blob/master/protocol/txntype.go) in the Algorand protocol. 

Generate this transaction with the SDKs and `goal` as follows:

```javascript tab="JavaScript"
```

```python tab="Python"
```

```java tab="Java"
```

```go tab="Go"
```

```zsh tab="goal"
goal clerk send --from=EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4 --to=GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A --amount=5000000 --note="Hello World" --firstvalid=6000000 --lastvalid=6001000 -o <txn-filename>
```

### See also
- See a step-by-step guide to creating a Payment Transaction in the [Getting Started Section](../getting-started/tutorial.md). 
  

## KeyReg
The purpose of a `KeyRegistrationTx` is to register an account `online` or `offline` to participate in Algorand Consensus. The context of key registration transactions is described in detail in the [Participate in Consensus](../network-participation/participate-in-consensus/generate_keys.md) section. 

### Important Fields

#### Sender
The Sender pays the fee and its `status` will change to `online` once the transaction is confirmed.

!!! warning
	An account that is marked `online` does not necessarily mean it is participating in consensus. The process of registering an account online involves first generating a participation key *prior* to issuing a KeyReg transaction. Follow the steps in the Participate in Consensus section to ensure good network behavior.


#### Vote Participation Key
*required*

A base64 encoded string corresponding to the root participation public key. See [Generate a Participation Key](../network-participation/participate-in-consensus/generate_keys.md) to learn more.

#### Selection Key 
*required*

This is a base64-encoded string corresponding to the VRF public key.

#### Vote First Round
*required*
The first round that the *participation key* is valid. This should not be confused with the [First Valid Round](#first-valid-round) of the keyreg transaction.

#### Vote Last Round
*required*

The last round that the *participation key* is valid. This should not be confused with the [Last Valid Round](#last-valid-round) of the keyreg transaction.

#### Vote Key Dilution
*required*

This is the dilution for the 2-level participation key. 

!!! info
	The participation key-related transaction values come from the participation key generated on the participating node. 


```javascript tab="JavaScript"
```

```python tab="Python"
```

```java tab="Java"
```

```go tab="Go"
```

```text tab="goal"
```

## Asset Configuration
## Asset Transfer
## Asset Freeze

# Useful How-Tos

## Sending a Transaction in the Future

## Determining the Optimal Fee

## Closing an Account

Closing an account means removing it from the Algorand ledger. Since there is a [minimum balance requirement](../feature-guides/accounts.md#minimum-balance) for every account on Algorand, the only way to completely remove it is to use the [Close Remainder To](#close-remainder-to) field. 

If you have asset holdings, you must first close out those asset holdings before you can close out the Algorand account completely. Close out your asset holdings by specifying a Close Remainder To address within an [Asset Transfer](#asset-transfer) transaction of the relevant asset.

## Setting First and Last Valid

Unless you have specific security concerns or logical constraints embedded within a specific Algorand Smart Contract, it is generally recommended that you set your default range to the maximum, currently 1000. This will give you an ample window of validity time to submit your transaction. 



[maybe include matrix of all different fields and transaction types]