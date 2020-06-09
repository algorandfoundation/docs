title: Structure

This section looks at how transactions are constructed, and in particular, how to _read and understand_ the underlying transaction composition after it has been created. To learn how to _create_ those same transactions visit the corresponding feature guide that is linked in each of the examples below. It is the hope that the combination of these guides will aid in developing a comprehensive understanding of how transactions work on Algorand. 

At the end of this section are several useful transaction-related how-tos.

!!! tip
	When you are given a transaction to sign, understanding its underlying representation will help you verify that the details of the transaction are correct.

# Transaction Types
There are [five transaction types](https://github.com/algorand/go-algorand/blob/master/protocol/txntype.go) in the Algorand Protocol: 1) [Payment](#payment-transaction), 2) [Key Registration](#key-registration-transaction), 3) [Asset Configuration](#asset-configuration-transaction), 4) [Asset Freeze](#asset-freeze-transaction), 5) [Asset Transfer](#asset-transfer-transaction).

These five transaction types can be specified in particular ways that result in more granular perceived transaction types. As an example, a transaction to [create an asset](../atomic_transfers.md#creating-an-asset) and [destroy an asset](../atomic_transfers.md#destroying-an-asset) use the same underlying `AssetConfigTx` type. Distinguishing these two transactions requires knowing which combination of `AssetConfigTx` fields and values result in one versus the other. This guide will help explain those differences.  Fortunately, the SDKs provide intuitive methods to create these more granular transaction types without having to necessarily worry about the underlying structure. However, if you are signing a pre-made transaction, correctly interpreting the underlying structure is critical. 

Note that all of the transactions shown in this guide are not yet authorized and would fail if submitted to the network. The next section will explain how to [authorize transactions](../transactions/signatures.md) before sending them to the network.

# Transaction Walkthroughs
The following sections describe the five types of Algorand transactions through example transactions that represent common use cases. Each transaction is displayed using the `goal clerk inspect` command which takes a signed or unsigned transaction file (msgpack-encoded) as input and outputs a human-readable json object. 

## Payment Transaction

A `PaymentTxn` sends Algos (the Algorand blockchain's native currency) from one account to another.

[_Payment Transaction Fields Reference_](../../reference/transactions.md#payment-transaction)


### Send 5 Algos
Here is an example transaction that sends 5 Algos from one account to another on MainNet. 

```json
{
  "txn": {
    "amt": 5000000,
    "fee": 1000,
    "fv": 6000000,
    "gen": "mainnet-v1.0",
    "gh": "wGHE2Pwdvd7S12BL5FaOP20EGYesN73ktiC1qzkkit8=",
    "lv": 6001000,
    "note": "SGVsbG8gV29ybGQ=",
    "rcv": "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A",
    "snd": "EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4",
    "type": "pay"
  }
}
```
The `"type": "pay"` signals that this is a payment transaction. 

This transaction transfers 5 Algos (shown as 5000000 microAlgos) from the account represented by the address starting with `"EW64GC..."` to the account with the address starting with `"GD64YI..."`. The sender address (`"EW64GC..."`) will pay a fee of `1000` microAlgos, which is also the minimum fee. An optional note is included in this transaction, which corresponds to the base64-encoded bytes for `"Hello World"`. Note that the base64 representation is a by product of the output of the `goal clerk inspect` command. 

This transaction is valid on MainNet, as per the genesis hash value which corresponds to [MainNet's genesis hash](../../reference/algorand-networks/mainnet.md#genesis-hash). The genesis ID is also provided for human-readability and also matches [MainNet](../../reference/algorand-networks/mainnet.md#genesis-id). Be sure to validate against the genesis hash value since it is unique to the specific network. The genesis ID is not; anyone could spin up a private network and call it `"mainnet-v1.0"` if desired. This transaction is valid if submitted between rounds 6000000 and 6001000.

**Related How-To**

- [Create a Payment Transaction](../../build-apps/hello_world.md). 

### Close an Account

Closing an account means removing it from the Algorand ledger. Since there is a minimum balance requirement for every account on Algorand, the only way to completely remove it is to use the [Close Remainder To](../../reference/transactions.md#closeremainderto) field as in the transaction below.

```json
{
  "txn": {
    "close": "EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4",
    "fee": 1000,
    "fv": 4695599,
    "gen": "testnet-v1.0",
    "gh": "SGO1GKSzyE7IEPItTxCByw9x8FmnrCDexi9/cOUJOiI=",
    "lv": 4696599,
    "rcv": "EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4",
    "snd": "SYGHTA2DR5DYFWJE6D4T34P4AWGCG7JTNMY4VI6EDUVRMX7NG4KTA2WMDA",
    "type": "pay"
  }
}
```
In this transaction, after the fee and the transaction `"amt"` are paid to the [receiver](../../reference/transactions.md#receiver) from the [sender](../../reference/transactions.md#closeremainderto) account (`"SYGHTA..."`), the remaining balance is transferred to the [closeto](../../reference/transactions.md#closeremainderto) account (`"SYGHTA..."`). Note that there is an implicit `"amt"` of 0 Algos when none is specified.

!!! info
    If you have asset holdings, you must first close out those asset holdings before you can close out the Algorand account completely. Close out your asset holdings by specifying an [Asset Close Remainder To](../../reference/transactions.md#closeassetto) address within an Asset Transfer transaction.

## Key Registration Transaction
The purpose of a `KeyRegistrationTx` is to register an account either `online` or `offline` to participate (i.e. vote) in Algorand Consensus. 

An account that is marked `online` does not necessarily mean it is participating in consensus. The process of registering an account online involves first generating a participation key *prior* to issuing a KeyReg transaction. It is important to follow the steps in the [Participate in Consensus section](../../run-a-node/participate/index.md) for a full overview participation and to ensure that you follow good network behavior. 

[_Key Registration Transaction Fields Reference_](../../reference/transactions.md#key-registration-transaction)

### Register account online 
This is an example of an **online** key registration transaction. 

```json
{
  "txn": {
    "fee": 2000,
    "fv": 6002000,
    "gh": "SGO1GKSzyE7IEPItTxCByw9x8FmnrCDexi9/cOUJOiI=",
    "lv": 6003000,
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
What distinguishes this as a key registration transaction is `"type": "keyreg"` and what distinguishes it as an _online_ key registration is the existence of the participation key-related fields, namely `"votekey"`, `"selkey"`, `"votekd"`, `"votefst"`, and `"votelst"`. The values for these fields are obtained by dumping the participation key info on the node where the participation key lives. The [sender](../../reference/transactions.md#sender) (`"EW64GC..."`) will pay a fee of `2000` microAlgos and its account state will change to `online` after this transaction is confirmed by the network. The transaction is valid between rounds 6002000 annd 6003000 on [TestNet](../../reference/algorand-networks/testnet.md).

**Related How-To**

- [Generate a Participation Key](../../run-a-node/participate/generate_keys.md)
- [Register an Account Online](../../run-a-node/participate/online.md) 

### Register account offline

Here is an example of an **offline** key registration transaction.

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
What distinguishes this from an _online_ transaction is that it does _not_ contain any participation key-related fields, since the account will no longer need a participation key if the transaction is confirmed. The [sender](../../reference/transactions.md#sender) (`"EW64GC..."`) will pay a fee of `2000` microAlgos and its account state will change to `offline` after this transaction is confirmed by the network. This transaction is valid between rounds 7,000,000 (`"fv"`) and 7,001,000 (`"lv"`) on [TestNet](../../reference/algorand-networks/testnet.md#genesis-hash) as per the [Genesis Hash](#genesis-hash) (`"gh"`) value.

**Related How-To**

- [Register an Account Offline](../../run-a-node/participate/offline.md) 

## Asset Configuration Transaction
An `AssetConfigTx` is used to create an asset, modify certain parameters of an asset, or destroy an asset. 

[_Asset Configuration Transaction Fields Reference_](../../reference/transactions.md#asset-configuration-transaction)

### Create an Asset

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
The `"type": "acfg"` distinguishes this as an Asset Configuration transaction. What makes this uniquely an **asset creation** transaction is that _no_ [asset ID (`"caid"`)](../../reference/transactions.md#configasset) is specified and there exists an [asset parameters](../../reference/transactions.md#asset-parameters) struct that includes all the initial configurations for the asset. The asset is [named](../../reference/transactions.md#assetname) (`an`) "My New Coin". the [unitname](../../reference/transactions.md#unitname) (`"un"`) is "MNC". There are 50,000,000 [total](../../reference/transactions.md#total) base units of this asset. Combine this with the [decimals](../../reference/transactions.md#decimals) (`"dc"`) value set to 2, means that there are 500,000.00 of this asset. There is an [asset URL](../../reference/transactions.md#asseturl) (`"au"`) specified which points to [developer.algorand.org](https://developer.algorand.org/) and a base64-encoded [metadata hash](../../reference/transactions.md#metadatahash) (`"am"`). This specific value corresponds to the SHA512/256 hash of the string "My New Coin Certificate of Value". The [manager](../../reference/transactions.md#manageraddr) (`"m"`), [freeze](../../reference/transactions.md#freezeaddr) (`"f"`), [clawback](../../reference/transactions.md#clawbackaddr) (`"c"`), and [reserve](../../reference/transactions.md#reserveaddr) (`"r"`) are the same as the sender. The [sender](../../reference/transactions.md#sender) is also the [creator](../../reference/transactions.md#creator).

This transaction is valid between rounds 6000000 (`"fv"`) and 6001000 (`"lv"`) on [TestNet](../../reference/algorand-networks/testnet.md#genesis-hash) as per the [Genesis Hash](../../reference/transactions.md#genesishash) (`"gh"`) value.

**Related How-To**

- [Create an Asset](../asa.md#creating-an-asset)

### Reconfigure an Asset
A **Reconfiguration Transaction** is issued by the asset manager to change the configuration of an already created asset.

Here is what an example reconfiguration transaction that changes the manager address for the asset with the Id `168103` that was [created above](#create-an-asset).  

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
What distinguishes this from an asset creation transaction is the inclusion of the **asset id** to be changed. The only fields that can be reconfigured are the [manager](../../reference/transactions.md#manageraddr), [freeze](../../reference/transactions.md#freezeaddr), [clawback](../../reference/transactions.md#clawbackaddr), and [reserve](../../reference/transactions.md#reserveaddr) addresses. All of them must be specified even if they do not change. 

!!! warning
	The protocol interprets unspecified addresses in an `AssetConfigTx` as an explicit action to set those values to null for the asset. Once set to `null`, this action cannot be undone.

Upon confirmation, this transaction will change the manager of the asset from `"EW64GC..."` to `"QC7XT7..."`.
This transaction is valid on [TestNet](../../reference/algorand-networks/testnet.md#genesis-hash) between rounds 6002000 and 6003000. A fee of `1000` microAlgos will be paid by the sender if confirmed. 

**Related How-To**

- [Modifying an Asset](../asa.md#modifying-an-asset)

### Destroy an Asset

A **Destroy Transaction** is issued to remove an asset from the Algorand ledger. To destroy an existing asset on Algorand, the original `creator` must be in possession of all units of the asset and the `manager` must send and therefore authorize the transaction. 

Here is what an example transaction destroy transaction looks like:

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

**Related How-To**

- [Destroying an Asset](../asa.md#destroying-an-asset)

## Asset Transfer Transaction
An Asset Transfer Transaction is used to opt-in to receive a specific type of Algorand Standard Asset, transfer an Algorand Standard asset, or revoke an Algorand Standard Asset from a specific account.

[_Asset Transfer Transaction Fields Reference_](../../reference/transactions.md#asset-transfer-transaction)

### Opt-in to an Asset
Here is an example of an opt-in transaction:

```json
{
  "txn": {
    "arcv": "QC7XT7QU7X6IHNRJZBR67RBMKCAPH67PCSX4LYH4QKVSQ7DQZ32PG5HSVQ",
    "fee": 1000,
    "fv": 6631154,
    "gh": "SGO1GKSzyE7IEPItTxCByw9x8FmnrCDexi9/cOUJOiI=",
    "lv": 6632154,
    "snd": "QC7XT7QU7X6IHNRJZBR67RBMKCAPH67PCSX4LYH4QKVSQ7DQZ32PG5HSVQ",
    "type": "axfer",
    "xaid": 168103
  }
}
```
The `"type": "axfer"` distinguishes this as an asset transfer transaction. The fields used in the transaction are the same as any other asset transfer. What distinguishes it as an opt-in transaction is in how those fields are specified and the sender account's asset holdings state prior to sending the transaction. In particular, the address `"QC7XT7...` is both the [sender](../../reference/transactions.md#sender) and [asset receiver](../../reference/transactions.md#assetreceiver) and it is assumed that the sender does not yet possess any of the desired asset identified with the [asset ID](../../reference/transactions.md#xferasset) `168103`. The asset amount is not specified in this example, which is equivalent to adding an [asset amount](../../reference/transactions.md#assetreceiver) equal to 0 (`"aamt": 0`). This transaction is valid on TestNet between rounds 6631154 and 6632154.

**Related How-To**

- [Receiving an Asset](../asa.md#receiving-an-asset)
  
### Transfer an Asset

Here is an example of an asset transfer transaction. 
```json
{
  "txn": {
    "aamt": 1000000,
    "arcv": "QC7XT7QU7X6IHNRJZBR67RBMKCAPH67PCSX4LYH4QKVSQ7DQZ32PG5HSVQ",
    "fee": 3000,
    "fv": 7631196,
    "gh": "SGO1GKSzyE7IEPItTxCByw9x8FmnrCDexi9/cOUJOiI=",
    "lv": 7632196,
    "snd": "EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4",
    "type": "axfer",
    "xaid": 168103
  }
}
```
An asset transfer transaction assumes that the asset receiver has already [opted-in](#opt-in-to-an-asset). The account represented by address `"EW64GC6..."` sends 1 million base units (or 10,000.00 units) of asset `168103` between rounds 7631196 annd 7632196 on TestNet. `"EW64GC6..."` pays a fee of 3000 microAlgos.

!!! tip
	If you are displaying asset amounts to users, be sure to include the asset's `"decimal"` configuration for easier readability. 

**Related How-To**

- [Transferring an Asset](../asa.md#transferring-an-asset)

### Revoke an Asset

Here is an example of the clawback account revoking assets from another account.

```json
{
  "txn": {
    "aamt": 500000,
    "arcv": "EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4",
    "asnd": "QC7XT7QU7X6IHNRJZBR67RBMKCAPH67PCSX4LYH4QKVSQ7DQZ32PG5HSVQ",
    "fee": 1000,
    "fv": 7687457,
    "gh": "SGO1GKSzyE7IEPItTxCByw9x8FmnrCDexi9/cOUJOiI=",
    "lv": 7688457,
    "snd": "EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4",
    "type": "axfer",
    "xaid": 168103
  }
}
```
The existence of an [asset sender](../../reference/transactions.md#assetsender) tells us that this transaction is utilizing the clawback workflow. During a clawback, the clawback address (`"EW64GC..."`) sends the transactions and therefore authorizes it and pays the `1000` microAlgo fee. The [asset sender](../../reference/transactions.md#assetsender) (`"QC7XT7..."`) is the address of the account from which the assets will be revoked. In this case, 5 million base units (5,000.00 units) of asset `168103` will be revoked from `"QC7XT7..."` and transferred to `"EW64GC..."`.

**Related How-To**

- [Revoking an Asset](../asa.md#revoking-an-asset)

## Asset Freeze Transaction
An Asset Freeze Transaction is issued by the Freeze Address and results in the asset receiver address losing or being granted the ability to send or receive the frozen asset.

### Freeze an Asset

```json
{
  "txn": {
    "afrz": true,
    "fadd": "QC7XT7QU7X6IHNRJZBR67RBMKCAPH67PCSX4LYH4QKVSQ7DQZ32PG5HSVQ",
    "faid": 168103,
    "fee": 1000,
    "fv": 7687793,
    "gh": "SGO1GKSzyE7IEPItTxCByw9x8FmnrCDexi9/cOUJOiI=",
    "lv": 7688793,
    "snd": "EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4",
    "type": "afrz"
  }
}
```
An asset freeze transaction is identified by `"type": "afrz"`. In this example, the [freeze manager](../../reference/transactions.md#freezeaddr) `"EW64GC..."` (i.e. the sender) freezes the asset `168103` for the account represented by address `"QC7XT7..."`. To unfreeze the asset, the [`"afrz"`](../../reference/transactions.md#assetfrozen) field is set to `true`.

### See also
- [Freezing an Asset](../asa.md#freezing-an-asset)

# Sending a Transaction in the Future

Algorand transactions are valid for a specific round range and the range maximum is 1000 rounds. If you plan to submit the transaction right away, specifying this round range is trivial. However, when the transaction requires offline signing or you plan to make frequent transactions from that account, it may be beneficial to specify a future round range or ranges that are more convenient. You can sign these transactions in a single secure session, and then submit them to the network when the valid round range is reached.

!!! tip
	For recurring transactions, Algorand Smart Contracts can be a more secure option. Read the corresponding [guide](../asc1/index.md) to learn more.

Calculating the round range requires you to know the **current round**, the **average block time**, and the **target submission time**. 

## Current Round

To retrieve the **current round** check the latest round passed for the network where you plan to submit the transaction/s. Check existing [block explorers](../../community.md#block-explorers) or get this info from your node's REST endpoint or `goal`. See [Check Node Status and Version](../../build-apps/connect.md#check-node-status-and-network-version).

## Average Block Time

This refers to the number of seconds it takes, on average, for a block to be committed on the Algorand blockchain. This number is not dynamically available through the Algorand developer tools, but at the time of writing this, blocks are confirmed in less than 5 seconds on Algorand so you can use a rough estimate of 4.5 seconds if precision is not critical. It is highly recommended that you validate this number against your own analytics or check our [Community Projects](../../community.md) for other projects that may provide this information since the average has shown above may be out of date at the time of reading this.

## Target Submission Time

This is the clock time at which you are targeting to send the transaction.

## Calculation
Calculate the delta between the target submission time and the current time in seconds. Divide that time by the average seconds per block to get the number of blocks spanning that time period. Add that number to the current round to get the first valid round for your transaction. Add 1000 to the first valid round to get the last valid round.

Keep in mind that these block times are estimations and it is not possible to be exactly precise for a given target time. Also, the longer out you project a round range, the wider the potential drift of round against clock time given natural variability in block times (i.e. 4.5 is just the average now but may vary during certain time periods).

## Example Scenarios

Here are three example scenarios and how the round range may be calculated for each. 

??? Example "Example - Today is Jan. 31, 2020 and I want to transfer funds on Feb. 2, 2020 at 20:00 UTC."

	Calculate the delta in seconds between current time (January 31, 2020 09:58 UTC) and February 2, 2020 20:00 UTC:
	
	```
	2 days + 10 hours + 2 minutes = 
		(2 days * 24 hours/day * 3600 seconds/hour) + 
		(10 hours * 3600 seconds/hour) + 
		(2 minutes * 60/seconds per minute) = 208,920 seconds
	```
		
	Calculate the average number of blocks produced in that time period, i.e. divide the total number of seconds by the average number of seconds it takes to produce a block. Assume for this example, that the average block time is 4.5 seconds.
		
	```
	208920 seconds/4.25 seconds per block ~ 46,427 blocks
	```

	Calculate the first valid round for the transaction by addng the total number of blocks to the current block. Assume the current block (i.e. round) is round 5,000,000. Then:
	
	```
	firstValidRound = 5000000 + 46427 = 5,046,427
	```

	Add 1000 rounds to get the last valid round for the transaction.
	
	```
	lastValidRound = 5046427 + 1000 = 5,047,427
	```

??? Example "Example I want to transfer funds in about 24 hours, but I need the flexibility to submit up to 3 hours after that."

	Given that I need some flexibility on when to submit, I create three [roughly] duplicate transactions with consecutive valid round ranges starting at roughly 24 hours from now. I will assume a max round range of 1000 which will give me between 3 and 4 hours to submit given an average block time of 4.5 secons. Assume the current time is January 31, 2020 09:58 UTC.

	Convert 24 hours to seconds (to determine the delta from now to target time):

   	```
	24 hours * 60 minutes/hour * 60 seconds/minute = 86,400 seconds
	```
	Calculate the number of blocks produced on average during that time period:
   
    ```
	86400 seconds/4.5 seconds per block: about 19,200 blocks
	```
	Determine first valid round and last valid round for first transaction. Assume current network round is 6,000,000:

    ```
	First Valid Round =  6,000,000 + 19,200 = 6,019,200
	Last Valid Round = 6,019,200 + 1000 = 6,020,200
	```
	Calculate the first and last valid rounds for the next 2 duplicate transactions:

	```
	Duplicate transaction 1:
	First Valid Round = 6,020,200
	Last Valid Round = 6,020,200 + 1000 = 6,021,200
	Duplicate transaction 2:
	First Valid Round = 6,021,200
	Last Valid Round = 6,021,200 + 1000 = 6,022,200
	```



??? Example "Example - I want to transfer funds daily to a predetermined account."

	In this situation, I may create two transactions per day (every 12 hours) and submit one each day, with a backup in case the first one cannot be submitted for some reason. I assume a longer average block time of 4.7 seconds to make it more likely that I will have at least 2 valid transactions in one day. I schedule the first one for 12 hours from the current time. Transaction A and B below would be projected out for the number of days that this type of transaction should be in effect.

	Transaction A:

	```
	12 hours * 60 minutes/hour * 60 seconds/minute = 43200 seconds
	43200 seconds/4.7 seconds per block ~ 9191 blocks
	First Valid Round = 6000000 + 9191 = 6009191
	Last Valid Round = 6009191 + 1000 = 6010191
	```
	Transaction B:

	```
	First Valid Round = 6009191 + 9191 = 6018382
	Last Valid Round = 6018382 + 1000 = 6019382
	```

# Fees

There are two primary ways to set the fee for a transaction. 

## Suggested Fee

The SDK provides a method to get the [**suggested fee** per byte (`fee`)](../../reference/rest-apis/algod/v1.md#transactionfee) which can be used to set the total fee for a transaction. This value is multiplied by the estimated size of the transaction in bytes to determine the total transaction fee. If the result is less than the minimum fee, the minimum fee is used instead. 

For larger transactions (> 1 KB in size), the resulting total fee will be greater than the network minimum, which in certain network conditions, may be more than you need to pay to get the transaction processed into the blockchain quickly. In particular, when blocks have enough room for all transactions, the minimum transaction fee will generally suffice. In this network scenario, set the fee (per byte) to 0 if you want to ensure that the minimum fee is chosen instead. 

In the future as more transactions are added to the network (and blocks are full), the minimum transaction fee may not guarantee that your transaction is processed as quickly as other transactions with higher fees set. In this case, using the returned suggested fee, which is based on the current transaction load, is the preferred method.

## Flat Fee
You can also manually set a **flat fee**. If you choose this method, make sure that your fee covers at least the [minimum transaction fee (`minFee`)](../../reference/rest-apis/algod/#transactionparams), which can be obtained from the suggested parameters method call in each of the SDKs.  Flat fees may be useful for applications that want to guarantee showing a specific rounded fee to users or for a transaction that is meant to be sent in the future where the network traffic conditions are unknown.

# Setting First and Last Valid

Unless you have specific security concerns or logical constraints embedded within a specific Algorand Smart Contract, it is generally recommended that you set your default range to the maximum, currently 1000. This will give you an ample window of validity time to submit your transaction. 
