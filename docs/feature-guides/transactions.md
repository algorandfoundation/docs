title: Transactions

In the [Getting Started Guide](../getting-started/tutorial.md), you created your first transaction on Algorand and were introduced to the pattern of constructing a transaction, authorizing it, sending it to the network, and validating its inclusion in a block. This section looks at the construction portion of that pattern. 

# Overview
There are five transaction types in the Algorand Protocol:

1. [Payment](#payment)
2. [Key Registration](#keyreg)
3. [Asset Configuration](#asset-configuration)
4. [Asset Freeze](#asset-freeze)
5. [Asset Transfers](#asset-freeze)

This guide is meant to show you how these transactions differ underneath (i.e. bottom up), but it will link to associated guides that approach these same transactions from the user-facing (i.e. top down) perspective.

As an example, a transaction to create an asset and destroy an asset use the same underlying transaction type (`AssetConfigTx`) and are distinguishable only by the fields they specify. This guide will explain what their underlying differences are. Fortunately, the SDKs abstract away from the transaction types where possible so that as a developer, you interface instead with a method that allows you to create and destroy a transaction, respectively. That perspective is explained in the [Algorand Standard Assets guide](./asa.md). It is the hope that the combination of these guides will aid in a comprehensive understanding of how transactions work on Algorand. 

Note that all of the transactions shown in this guide are not yet authorized and would fail if submitted to the network. The following section, [Signatures](./signatures.md), will guide you on how to authorize transactions before sending them to the network.

# Fields common to all transactions
There are several parameters that are common to all transactions on Algorand.

## Sender
*required*

The address of the account that will pay the `amount` and `"fee"` and authorize the transaction. 

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
Here is an example transaction that sends 5 Algos from one account to another on TestNet. The output below was generated with `goal clerk inspect`.

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
From top to bottom, this transaction specifies an [amount (`"amt"`)](#amount) of 5 million microAlgos (i.e. 5 Algos). The [`"fee"`](#fee) is 1000 microAlgos. `"fv"` refers to the [FirstValidRound](#first-valid-round), set for round 6 million. If this transaction is sent prior to round 6 million, it will be rejected. The `"gen"` field is the optional [Genesis ID](#genesis-id). `"gh"` is the [Genesis Hash](#genesis-hash) in base64 encoding. This transaction is only valid on [TestNet](../algorand-networks/testnet.md) which can be verified by matching the hashes. `"lv"` refers to the [Last Valid Round](#last-valid-round) which is set to 1,000 rounds after the first valid round, i.e. the maximum range. The optional `"note"` field is a base64 representation of the string "Hello World". The base64 encoding is the representation returned by `goal clerk inspect`. The [Receiver](#receiver) (`"rcv"`) address is the TestNet faucet address. The [Sender](#sender) (`"snd"`) address is a random generated account. Finally, the `"type"` field is filled in by the developer tools. The value in this transaction is `pay` which maps to the `PaymentTxn` [type](https://github.com/algorand/go-algorand/blob/master/protocol/txntype.go) in the Algorand protocol. 

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
The purpose of a `KeyRegistrationTx` is to register an account `online` or `offline` to participate in Algorand Consensus. The context of key registration transactions is described in detail in the [Participate in Consensus](../network-participation/participate-in-consensus/overview.md) section. 

### Important Fields

#### Sender
The Sender pays the fee and its `status` will change to `online` once the transaction is confirmed.

!!! warning
	An account that is marked `online` does not necessarily mean it is participating in consensus. The process of registering an account online involves first generating a participation key *prior* to issuing a KeyReg transaction. Follow the steps in the [Participate in Consensus section](../network-participation/participate-in-consensus/overview.md) to ensure that you are following good network behavior.


#### Vote Participation Key
*required for online*

A base64 encoded string corresponding to the root participation public key. See [Generate a Participation Key](../network-participation/participate-in-consensus/generate_keys.md) to learn more.

#### Vote Selection Key 
*required for online*

This is a base64-encoded string corresponding to the VRF public key.

#### Vote First Round
*required for online*

The first round that the *participation key* is valid. This should not be confused with the [First Valid Round](#first-valid-round) of the keyreg transaction.

#### Vote Last Round
*required for online*

The last round that the *participation key* is valid. This should not be confused with the [Last Valid Round](#last-valid-round) of the keyreg transaction.

#### Vote Key Dilution
*required for online*

This is the dilution for the 2-level participation key. 

!!! info
	The participation key-related transaction values come from the participation key generated on the participating node. 

### Example - Register account online 
This is an example of an online registration transaction, written to file, and then viewed with `goal clerk inspect`.

```json
{
  "txn": {
    "fee": 1000,
    "fv": 6000000,
    "gh": "SGO1GKSzyE7IEPItTxCByw9x8FmnrCDexi9/cOUJOiI=",
    "lv": 6001000,
    "selkey": "X84ReKTmp+yfgmMCbbokVqeFFFrKQeFZKEXG89SXwm4=",
    "snd": "EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4",
    "type": "keyreg",
    "votefst": 6000000,
    "votekd": 1730,
    "votekey": "eXq34wzh2UIxCZaI1leALKyAvSz/+XOe0wqdHagM+bw=",
    "votelst": 9000000
  }
}
```
First, take a look at the  fields that are common across all transactions. This transaction specifies a [`"fee"`](#fee) of 1000 microAlgos and is valid between rounds 6,000,000 and 6,001,000 as shown by the values of `"fv"` and `"lv"` respectively. The transaction is only valid on [TestNet](../algorand-networks/testnet.md#genesis-hash) as referenced by the [Genesis Hash](#genesis-hash) (`"gh"`) and the corresponding [Genesis ID](#genesis-id). The `"type"` of transaction is a `keyreg` transaction which maps to the [`KeyRegistrationTx` type](https://github.com/algorand/go-algorand/blob/master/protocol/txntype.go) in the Algorand protocol.

The [Sender](#sender) is the account whose state will change to `online` after this transaction is confirmed by the network. The rest of the fields correspond to data about the participation key, namely the [Vote Selection Key](#vote-selection-key) (`"votekey"`), the [Vote Participation Key](#vote-participation-key) (`"votekey"`), the [Vote First Round](#vote-first-round) (`"votefst"`), the [Vote Last Round](#vote-last-round) (`"votelst"`) and the [Vote Key Dilution](#vote-key-dilution). These values can be obtained on the node where the participation key is hosted using the process described in the [Generate a Participation Key Section](../network-participation/participate-in-consensus/generate_keys.md#view-participation-key-info)

Use the SDKs to create the transaction above.

```javascript tab="JavaScript"
```

```python tab="Python"
```

```java tab="Java"
```

```go tab="Go"
```

```zsh tab="goal"
$ goal account changeonlinestatus -a EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4 --fee=1000 --firstvalid=6000000 --lastvalid=6001000 --online=true --txfile=<online-tx-filename>
```

### Example - Register account offline

Here is an example of a transaction that takes an account offline.

```json
{
  "txn": {
    "fee": 1000,
    "fv": 7000000,
    "gh": "SGO1GKSzyE7IEPItTxCByw9x8FmnrCDexi9/cOUJOiI=",
    "lv": 7001000,
    "snd": "EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4",
    "type": "keyreg"
  }
}
```
No data about the participation key is required for an offline transaction since the account will no longer need a participation key if the transaction is confirmed. This lack of participation key data is what signals to the network that this is the `offline` version of a `KeyRegistrationTx`. 

This transaction is valid between rounds 7,000,000 (`"fv"`) and 7,001,000 (`"lv"`) on [TestNet](../algorand-networks/testnet.md#genesis-hash) as per the [Genesis Hash](#genesis-hash) (`"gh"`). The `"fee"` paid by the Sender is 1000 microAlgos.

!!! warning
	The latest keyreg transaction that is confirmed by the network will always overtake any previous for that account.

Generate this exact transaction with the SDKs and `goal`:

```javascript tab="JavaScript"
```

```python tab="Python"
```

```java tab="Java"
```

```go tab="Go"
```

```zsh tab="goal"
$ goal account changeonlinestatus -a EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4 --online=false --firstvalid=7000000 --lastvalid=7001000 --fee=1000 --txfile=<offline-txn-filename>
```

### See also

- The full guide to [Participate in Consensus](../network-participation/participate-in-consensus/overview.md)

## Asset Configuration
An `AssetConfigTx` is used to create an asset, modify certain parameters of an asset, or destroy an asset. 

### Important Fields

#### Creator
*required*

The **creator** is the address of the account that creates the asset and holds all units of the asset upon creation. 

#### Total
*required*

The total number of asset units to be created. This number cannot be changed.

#### Decimals
*required*

The number of digits to use after the decimal point when displaying this asset. If set to 0, the asset is not divisible beyond its base unit. If set to 1, the base asset unit is tenths. If 2, the base asset unit is hundredths, and so on.

#### Asset  ID
*required, except when creating an asset*

This is the integer identifier for an asset that is unique across the entire blockchain. The **asset id** increments according to the number of transactions on the network. [FACT CHECK THIS]

#### Freeze Address
*optional*

This address, if specified, can freeze and unfreeze specific accounts that hold this asset. If an account is frozen, it cannot send or receive the asset.

Once this value is set to `null` it cannot be changed. 

#### Manager Address
*optional*

This address manages the configuration of the asset and therefore has the authority to change any of the [manager](#manager-address), [freeze](#freeze-address), [clawback](#clawback-address), or [reserve](#reserve-address) addresses. See [Reconfigure an Asset](#example---reconfigure-an-asset)

Never set this value to `null`, as it cannot be changed and you will lose access to management of your asset.

#### Clawback Address
*optional*

This address, if specified, has the authority to transfer units of this asset out of any account and into another. See an example of this type of transaction in the [example below](#example---revoke-an-asset).

Once this value is set to `null` it cannot be changed. 

#### Reserve Address
*optional*

This address can act as a holdings account for an asset after creation however. It has know specific authority in the protocol itself. This may be used in the case where you want to signal to holders of your asset that the unissued units of the asset live in account that is different from the **creator** account.

Once this value is set to `null` it cannot be changed. 

#### Asset Name
*optional but recommended*

This is the human-readable name for your asset. The name cannot exceed 32 bytes. 

#### Asset UnitName
*optional but recommended*

This a human-readable short-form name for a unit of the asset. This could be used as a ticker reference, like USDT. Unitnames cannot exceed 8 bytes. 

#### Asset URL
*optional*

Specify a URL that provides more details about the asset. This URL could also include information about what is represented in the [Asset Metadata Hash](#asset-metadata-hash). The URL size cannot exceed 32 bytes.

#### Default Frozen
*optional, defaults to false*

Set this value to `true`, if you want all units of the asset to be frozen by default. This means that no one can trade the asset, without explicit authorization from the [freeze address](#freeze-address). 

Set this value to `false`, if you want to allow anyone who owns this asset to trade it freely by default. 

#### Asset Metadata Hash
*optional*

This field is intended to represent a base64-encoded 32-byte hash of some metadata that is relevant to your asset and/or asset holders. This field can _only_ be specified upon creation. An example might be the hash of some certificate that acknowledges the digitized asset as the official representation of a particular real-world asset. 


### Example - Create an Asset

Here is an example asset creation transaction:

```json
{
  "txn": {
    "apar": {
      "am": "gXHjtDdtVpY7IKwJYsJWdCSrnUyRsX4jr3ihzQ2U9CQ=",
      "an": "My New Coin",
      "au": "developer.algorand.org",
      "c": "EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4",
      "dc": 2,
      "f": "EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4",
      "m": "EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4",
      "r": "EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4",
      "t": 50000000,
      "un": "MNC"
    },
    "fee": 1000,
    "fv": 6000000,
    "gh": "SGO1GKSzyE7IEPItTxCByw9x8FmnrCDexi9/cOUJOiI=",
    "lv": 6001000,
    "snd": "EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4",
    "type": "acfg"
  }
}
```
In an asset creation transaction, the [sender](#sender) (`"snd"`) _is_ the [creator](#creator). A missing [**asset id**](#asset-id) plus the `"type"` equal to `"acfg"`, the [`AssetConfigTx`](https://github.com/algorand/go-algorand/blob/master/protocol/txntype.go) and the existence of the `"apar"` object is a clear indication that this transaction is for *creating* a new asset.

Within the asset parameters (`"apar"`) object is the initial configuration of the asset. The asset is [named](#asset-name) (`an`) "My New Coin". the [unitname](#asset-unitname) (`"un"`) is "MNC". There are 50,000,000 [total](#total) base units of this asset. Combine this with the [decimals](#decimals) (`"dc"`) value set to 2, means that there are 500,000.00 of this asset. There is an [asset URL](#asset-url) (`"au"`) specified which points to [developer.algorand.org](https://developer.algorand.org/) and a base64-encoded [metadata hash](#asset-metadata-hash) (`"am"`). This specific value corresponds to the SHA512/256 hash of the string "My New Coin Certificate of Value".

The [manager](#manager-address) (`"m"`), [freeze](#freeze-address) (`"f"`), [clawback](#clawback-address) (`"c"`), and [reserve](#reserve-address) (`"r"`) are all set to the match the [creator](#creator) address.

The remaining parameters are the `fee` (1000 microAlgos), the first (`"fv"`) and last valid round (`"lv"`)(6,000,000 to 6,001,000), and the [genesis hash](#genesis-hash) for [TestNet](../algorand-networks/testnet.md#genesis-hash).

```zsh tab="goal"
goal asset create --creator EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4 --total 5000 --unitname="MNC" --name="My New Coin" --decimals=2 --defaultfrozen=false --fee=1000 --firstvalid 6000000 --lastvalid 6001000 --asseturl="developer.algorand.org" --assetmetadatab64="gXHjtDdtVpY7IKwJYsJWdCSrnUyRsX4jr3ihzQ2U9CQ=" --out asset.txn --note=""
```


### Example - Reconfigure an Asset
A **Reconfiguration Transaction** is a form of an `AssetConfigTx` that includes the **asset id** (since it now exists) and asset parameters that specify the [Clawback](#clawback-address), [Freeze](#freeze-address), [Manager](#manager-address), and [Reserve](#reserve-address) address.

Here is what an example reconfiguration transaction that changes the manager address for the asset id `168103` that was [created above](#example---create-an-asset).  

```json
{
  "txn": {
    "apar": {
      "c": "EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4",
      "f": "EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4",
      "m": "QC7XT7QU7X6IHNRJZBR67RBMKCAPH67PCSX4LYH4QKVSQ7DQZ32PG5HSVQ",
      "r": "EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4"
    },
    "caid": 168103,
    "fee": 1000,
    "fv": 6002000,
    "gh": "SGO1GKSzyE7IEPItTxCByw9x8FmnrCDexi9/cOUJOiI=",
    "lv": 6003000,
    "snd": "EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4",
    "type": "acfg"
  }
}
```
This transaction is valid on [TestNet](../algorand-networks/testnet.md#genesis-hash) between rounds 6,002,000 annd 6,003,000 as per the values of `"gh"`, `"fv"`, and `"lv"`, respectively. A `"fee"` of 1,000 microAlgos will be paid by the sender if confirmed. The transaction `"type"` is `"acfg"`.

The [sender](#sender) (`"snd"`) is the current [manager](#manager-address) of the asset, specified by the [asset ID](#asset-id) (`"caid"`) `168103`. The [manager address](#manager-address) (`"m"`) is set to `QC7XT7QU7X6IHNRJZBR67RBMKCAPH67PCSX4LYH4QKVSQ7DQZ32PG5HSVQ` and the [clawback](#clawback-address) (`"c"`), [freeze](#freeze-address) (`"f"`), and [reserve](#reserve-address) (`"r"`) address stay the same and therefore _must be_ re-specified in any reconfiguration transaction. 

!!! warning
	The protocol interprets unspecified addresses in an `AssetConfigTx` as an explicit action to set those values to null for the asset. Once set to `null`, this action cannot be undone.


### Example - Destroy an Asset

A **Destroy Transaction** is a form of the `AssetConfigTx` that *must* include the **asset ID** of the asset to be destroyed and _must **not**_ include any asset parameters.

To destroy an existing asset on Algorand, the original `creator` must be in possession of all units of the asset and the `manager` must send and therefore authorize the transaction. 

Here is what an example transaction looks like with `goal clerk inspect`:

```json
{
  "txn": {
    "caid": 168103,
    "fee": 1000,
    "fv": 7000000,
    "gh": "SGO1GKSzyE7IEPItTxCByw9x8FmnrCDexi9/cOUJOiI=",
    "lv": 7001000,
    "snd": "EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4",
    "type": "acfg"
  }
}
```

This transaction differentiates itself from an **Asset Creation** transaction in that it contains an **asset ID** (`caid`) pointing to the asset to be destroyed. It differentiates itself from an **Asset Reconfiguration** transaction by the *lack* of any asset parameters. 

```zsh tab="goal"
goal asset destroy --assetid=168103 --firstvalid=7000000 --lastvalid=7001000 --fee=1000 --creator=EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4 --note="" -o destroy.txn
```
### See also
- [Creating an Asset](./asa.md#creating-an-asset)
- [Modifying an Asset](./asa.md#modifying-an-asset)
- [Destroying an Asset](./asa.md#destroying-an-asset)

## Asset Transfer

### Example - Opt-in to an Asset

### Example - Transfer an Asset

### Example - Revoke an Asset

### See also
- [Receiving an Asset](./asa.md#receiving-an-asset)
- [Transferring an Asset](./asa.md#transferring-an-asset)
- [Revoking an Asset](./asa.md#revoking-an-asset)

## Asset Freeze

### Example - Freeze an Asset

### See also
- [Freezing an Asset](./asa.md#freezing-an-asset)

# Useful How-Tos

## Sending a Transaction in the Future
[this was copied and pasted from another doc - needs to be fixed and formatted]
Algorand transactions are valid for a specific round range. If you plan to submit the transaction within 1000 blocks (the maximum range allowed), specifying this round range is trivial. However, when the transaction requires offline signing and/or you plan to make frequent transactions from that account, it may be beneficial to specify a future round range or ranges that are more convenient. You can sign these transactions in a single secure session, and then submit them to the network when the valid round range is reached.

Calculating the round range requires you to know the current round, the average block time, and the target submission time. 

Current round: Run goal node status - the current round is the 
“Last committed block”. 
Average block time: Currently, average block time is 4.26 seconds
Target submission time:  Variable; up to you.

You can calculate as follows:
Calculate the delta between the target submission time and the current time in seconds.
Divide that time by the average seconds per block to get the number of blocks spanning that time period.
Add that number to the current round to get the first valid round for your transaction.
Add 1000 to the first valid round to get the last valid round.

Keep in mind that these block times are estimations and it is not possible to be exactly precise for a given target time. Also, the longer out you project a round range, the wider the potential drift given the variability in time that a block is created (i.e. 4.26 is just the average now but may vary during certain time periods).
Example Scenarios
Here are three example scenarios and how the round range could be calculated for each. We will assume for each that the current round is 162,159 and the current time is June 20, 2019 09:58 EST.
Scenario 1: I want to transfer funds on June 21, 2019 at 20:00 EST.

The delta between June 20, 2019 09:58 EST and June 21, 2019 20:00 EST:
1 day + 10 hours + 2 minutes = 1*24 hours*3600 + 10 hours*3600 + 2*60 = 122520 seconds
122520 seconds/4.25 seconds per block: about 28828 blocks
First Valid Round = 162159 + 28828 = 190987
Last Valid Round = 190987 + 1000 = 191987
Scenario 2: I want to transfer funds in about 24 hours, but I need the flexibility to submit up to 3 hours after that.

Given that I need some flexibility on when to submit, I can create three duplicate transactions with different consecutive valid round ranges starting at roughly 24 hours from now. I will assume a max round range of 1000 which will give me between 3 and 4 hours.

24 hours = 24*60 minutes*60 seconds = 86400 seconds
86400 seconds/4.25 seconds per block: about 20329 blocks
First Valid Round = 162159 + 20329 = 182488
Last Valid Round = 182488 + 1000 = 183488
Duplicate transaction 1:
First Valid Round = 183488
Last Valid Round = 183488 + 1000 = 184488
Duplicate transaction 2:
First Valid Round = 184488
Last Valid Round = 184488 + 1000 = 185488

Note that in any scenario where you create duplicate transactions for convenience, you need to be careful not to submit more than one. You may want to destroy the duplicate transaction files, to be sure not to send twice the amount of algos by mistake. See What is the security risk to having pre-signed transactions online?

??? Example - I want to transfer funds daily to a predetermined account.

	In this situation, I may create two transactions per day (every 12 hours) and submit one each day, with a backup in case the first one cannot be submitted for some reason. In this situation, I will assume a longer average block time of 4.5 seconds to make it more likely that I will have at least 2 valid transactions in one day. I will schedule the first one for 12 hours from the current time. Transaction A and B below would be projected out for the number of days that this type of transaction should be in effect.

	Transaction A:

	12 hours = 12*60 minutes*60 seconds = 43200 seconds
	43200 seconds/4.5 seconds per block: about 9600 blocks
	First Valid Round = 162159 + 9600 = 171759
	Last Valid Round = 171759 + 1000 = 172759

	Transaction B:

	12 hours = 12*60 minutes*60 seconds = 43200 seconds
	43200 seconds/4.5 seconds per block: about 9600 blocks
	First Valid Round = 171759 + 9600 = 181359
	Last Valid Round = 181359 + 1000 = 182359

## Determining the Optimal Fee
The minimum fee to send a transaction on Algorand is 1000 microAlgos. If blocks are not full, this fee is generally sufficient for the transaction to be prioritized into a block. The SDKs provide `suggestedFee` methods to help determine an optimal fee. 

[Need more explanation around how suggested fee works and how to use in SDKs.]

## Closing an Account

Closing an account means removing it from the Algorand ledger. Since there is a [minimum balance requirement](../feature-guides/accounts.md#minimum-balance) for every account on Algorand, the only way to completely remove it is to use the [Close Remainder To](#close-remainder-to) field. 

If you have asset holdings, you must first close out those asset holdings before you can close out the Algorand account completely. Close out your asset holdings by specifying a Close Remainder To address within an [Asset Transfer](#asset-transfer) transaction of the relevant asset.

## Setting First and Last Valid

Unless you have specific security concerns or logical constraints embedded within a specific Algorand Smart Contract, it is generally recommended that you set your default range to the maximum, currently 1000. This will give you an ample window of validity time to submit your transaction. 
