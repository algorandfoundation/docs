title: Transaction reference

Each table below specifies the **field name**, whether it is optional or required, its **type** within the protocol code (note that SDKs input types for these fields may differ), the **codec**, which is the name of the field when viewed within a transaction, and a **description** of the field. 

# Common Fields (`Header` and `Type`)
These fields are common to all transactions.


|Field|Required|Type|codec| Description|
|---|---|---|---|---|
|<a name="fee">Fee</a>| _required_| uint64|`"fee"`|Paid by the sender to the FeeSink to prevent denial-of-service. The minimum fee on Algorand is currently 1000 microAlgos.|
|<a name="firstvalid">FirstValid</a>| _required_ | uint64 | `"fv"`|The first round for when the transaction is valid. If the transaction is sent prior to this round it will be rejected by the network.|
|<a name="genesishash">GenesisHash</a>|_required_|[32]byte|`"gh"`|The hash of the genesis block of the network for which the transaction is valid. See the genesis hash for [MainNet](../../algorand-networks/mainnet#genesis-hash), [TestNet](../../algorand-networks/testnet#genesis-hash), and [BetaNet](../../algorand-networks/betanet#genesis-hash).
|<a name="lastvalid">LastValid</a>| _required_ | uint64 | `"lv"`|The ending round for which the transaction is valid. After this round, the transaction will be rejected by the network.|
|<a name="sender">Sender</a>| _required_ |Address|`"snd"`|The address of the account that pays the fee and amount.|
|<a name="type">TxType</a>|_required_|string|`"type"`| Specifies the type of transaction. This value is automatically generated using any of the developer tools.|
|<a name="genesisid">GenesisID</a>|_optional_|string|`"gen"`| The human-readable string that identifies the network for the transaction. The genesis ID is found in the genesis block. See the genesis ID for [MainNet](../../algorand-networks/mainnet#genesis-id), [TestNet](../../algorand-networks/testnet#genesis_id), and [BetaNet](../../algorand-networks/betanet#genesis-id). |
<a name="group">Group</a>|_optional_|[32]byte|`"grp"`|The group specifies that the transaction is part of a group and, if so, specifies the hash of the transaction group. Assign a group ID to a transaction through the workflow described in the [Atomic Transfers Guide](../../atomic_transfers).|
<a name="lease">Lease</a>|_optional_|[32]byte|`"lx"`|A lease enforces mutual exclusion of transactions. If this field is nonzero, then once the transaction is confirmed, it acquires the lease identified by the (Sender, Lease) pair of the transaction until the LastValid round passes. While this transaction possesses the lease, no other transaction specifying this lease can be confirmed.  A lease is often used in the context of Algorand Smart Contracts to prevent replay attacks. Read more about [Algorand Smart Contracts](../../dapps/smart-contracts/). Leases can also be used to safeguard against unintended duplicate spends. For example, if I send a transaction to the network and later realize my fee was too low, I could send another transaction with a higher fee, but the same lease value. This would ensure that only one of those transactions ends up getting confirmed during the validity period. |
|<a name="note">Note</a>|_optional_|[]byte|`"note"`| Any data up to 1000 bytes. |
|<a name="rekeyto">RekeyTo</a>|_optional_|Address|`"rekey"`| Specifies the authorized address. This address will be used to authorize all future transactions. Learn more about [Rekeying](../../accounts/rekey) accounts. |

# Payment Transaction
Transaction Object Type: `PaymentTx`

Includes all fields in [Header](#common-fields-header-and-type) and `"type"` is `"pay"`.

|Field|Required|Type|codec| Description|
|---|---|---|---|---|
|<a name="receiver">Receiver</a>| _required_ |Address|`"rcv"`|The address of the account that receives the [amount](#amount).|
|<a name="amount">Amount</a>|_required_|uint64|`"amt"`| The total amount to be sent in microAlgos.|
|<a name="closeremainderto">CloseRemainderTo</a>|_optional_|Address|`"close"`|When set, it indicates that the transaction is requesting that the [Sender](#sender) account should be closed, and all remaining funds, after the [fee](#fee) and [amount](#amount) are paid, be transferred to this address.|

# Key Registration Transaction
Transaction Object Type: `KeyRegistrationTx`

Includes all fields in [Header](#common-fields-header-and-type) and `"type"` is `"keyreg"`.

|Field|Required|Type|codec| Description|
|---|---|---|---|---|
|<a name="votepk">VotePk</a>| _required for online_ |ed25519PublicKey|`"votekey"`|The root participation public key. See [Generate a Participation Key](../../run-a-node/participate/generate_keys.md) to learn more.|
|<a name="selectionpk">SelectionPK</a>|_required for online_|VrfPubkey|`"selkey"`| The VRF public key.|
|<a name="stateproofpk">StateProofPk</a>|_required for online_|MerkleSignature Verifier (64 bytes)|`"sprfkey"`| The 64 byte state proof public key commitment.|
|<a name="votefirst">VoteFirst</a>|_required for online_|uint64|`"votefst"`|The first round that the *participation key* is valid. Not to be confused with the [FirstValid](#firstvalid) round of the keyreg transaction.|
|<a name="votelast">VoteLast</a>|_required for online_|uint64|`"votelst"`|The last round that the *participation key* is valid. Not to be confused with the [LastValid](#lastvalid) round of the keyreg transaction.|
|<a name="votekeydilution">VoteKeyDilution</a>|_required for online_|uint64|`"votekd"`|This is the dilution for the 2-level participation key. It determines the interval (number of rounds) for generating new ephemeral keys. |
|<a name="nonparticipation">Nonparticipation</a>|_optional_|bool|`"nonpart"`| All new Algorand accounts are participating by default. This means that they earn rewards. Mark an account nonparticipating by setting this value to `true` and this account will no longer earn rewards. It is unlikely that you will ever need to do this and exists mainly for economic-related functions on the network.|

# Asset Configuration Transaction
Transaction Object Type: `AssetConfigTx`

Includes all fields in [Header](#common-fields-header-and-type) and `"type"` is `"acfg"`.

This is used to create, configure and destroy an asset depending on which fields are set.

|Field|Required|Type|codec| Description|
|---|---|---|---|---|
|<a name="configasset">ConfigAsset</a>| _required, except on create_ |uint64|`"caid"`|For re-configure or destroy transactions, this is the unique asset ID. On asset creation, the ID is set to zero.|
|[AssetParams](#asset-parameters) |_required, except on destroy_|[AssetParams](#asset-parameters)|`"apar"`| See AssetParams table for all available fields.|

## Asset Parameters
Object Name: `AssetParams`

|Field|Required|Type|codec| Description|
|---|---|---|---|---|
|<a name="total">Total</a>|_required on creation_|uint64|`"t"`| The total number of base units of the asset to create. This number cannot be changed.|
|<a name="decimals">Decimals</a>|_required on creation_|uint32|`"dc"`| The number of digits to use after the decimal point when displaying the asset. If 0, the asset is not divisible. If 1, the base unit of the asset is in tenths. If 2, the base unit of the asset is in hundredths, if 3, the base unit of the asset is in thousandths, and so on up to 19 decimal places |
|<a name="defaultfrozen">DefaultFrozen</a>|_required on creation_|bool|`"df"`| True to freeze holdings for this asset by default. |
|<a name="unitname">UnitName</a>|_optional_|string|`"un"`| The name of a unit of this asset. Supplied on creation. Max size is 8 bytes. Example: USDT |
|<a name="assetname">AssetName</a>|_optional_|string|`"an"`| The name of the asset. Supplied on creation. Max size is 32 bytes. Example: Tether|
|<a name="url">URL</a>|_optional_|string|`"au"`| Specifies a URL where more information about the asset can be retrieved. Max size is 96 bytes. |
|<a name="metadatahash">MetaDataHash</a>|_optional_|[]byte|`"am"`| This field is intended to be a 32-byte hash of some metadata that is relevant to your asset and/or asset holders. The format of this metadata is up to the application. This field can _only_ be specified upon creation. An example might be the hash of some certificate that acknowledges the digitized asset as the official representation of a particular real-world asset.  |
|<a name="manageraddr">ManagerAddr</a>|_optional_|Address|`"m"`| The address of the account that can manage the configuration of the asset and destroy it. |
|<a name="reserveaddr">ReserveAddr</a>|_optional_|Address|`"r"`| The address of the account that holds the reserve (non-minted) units of the asset. This address has no specific authority in the protocol itself. It is used in the case where you want to signal to holders of your asset that the non-minted units of the asset reside in an account that is different from the default creator account (the sender). |
|<a name="freezeaddr">FreezeAddr</a>|_optional_|Address|`"f"`| The address of the account used to freeze holdings of this asset. If empty, freezing is not permitted. |
|<a name="clawbackaddr">ClawbackAddr</a>|_optional_|Address|`"c"`| The address of the account that can clawback holdings of  this asset. If empty, clawback is not permitted. |

# Asset Transfer Transaction 
Transaction Object Type: `AssetTransferTx`

Includes all fields in [Header](#common-fields-header-and-type) and `"type"` is `"axfer"`.

|Field|Required|Type|codec| Description|
|---|---|---|---|---|
|<a name="xferasset">XferAsset</a>| _required_ |uint64|`"xaid"`|The unique ID of the asset to be transferred.|
|<a name="assetamount">AssetAmount</a>|_required_|uint64|`"aamt"`| The amount of the asset to be transferred. A zero amount transferred to self allocates that asset in the account's Asset map.|
|<a name="assetsender">AssetSender</a>|_required_|Address|`"asnd"`|The sender of the transfer. The regular [sender](#sender) field should be used and this one set to the zero value for regular transfers between accounts. If this value is nonzero, it indicates a clawback transaction where the [sender](#sender) is the asset's clawback address and the asset sender is the address from which the funds will be withdrawn.|
|<a name="assetreceiver">AssetReceiver</a>|_required_|Address|`"arcv"`| The recipient of the asset transfer.|
|<a name="assetcloseto">AssetCloseTo</a>|_optional_|Address|`"aclose"`|Specify this field to remove the asset holding from the [sender](#sender) account and reduce the account's minimum balance (i.e. opt-out of the asset). |

# Asset OptIn Transaction 
Transaction Object Type: `AssetTransferTx`

Includes all fields in [Header](#common-fields-header-and-type) and `"type"` is `"axfer"`.

This is a special form of an Asset Transfer Transaction.

|Field|Required|Type|codec| Description|
|---|---|---|---|---|
|<a name="xferasset">XferAsset</a>| _required_ |uint64|`"xaid"`|The unique ID of the asset to opt-in to.|
|<a name="sender">Sender</a>|_required_|Address|`"snd"`| The account which is allocating the asset to their account's Asset map.|
|<a name="assetreceiver">AssetReceiver</a>|_required_|Address|`"arcv"`| The account which is allocating the asset to their account's Asset map.|

# Asset Clawback Transaction 
Transaction Object Type: `AssetTransferTx`

Includes all fields in [Header](#common-fields-header-and-type) and `"type"` is `"axfer"`.

This is a special form of an Asset Transfer Transaction.

|Field|Required|Type|codec| Description|
|---|---|---|---|---|
|<a name="sender">Sender</a>| _required_ |Address|`"snd"`|The sender of this transaction must be the clawback account specified in the asset configuration.|
|<a name="xferasset">XferAsset</a>| _required_ |uint64|`"xaid"`|The unique ID of the asset to be transferred.|
|<a name="assetamount">AssetAmount</a>|_required_|uint64|`"aamt"`| The amount of the asset to be transferred.|
|<a name="assetsender">AssetSender</a>|_required_|Address|`"asnd"`| The address from which the funds will be withdrawn.|
|<a name="assetreceiver">AssetReceiver</a>|_required_|Address|`"arcv"`| The recipient of the asset transfer.|

# Asset Freeze Transaction
Transaction Object Type: `AssetFreezeTx`

Includes all fields in [Header](#common-fields-header-and-type) and `"type"` is `"afrz"`.

|Field|Required|Type|codec| Description|
|---|---|---|---|---|
|<a name="freezeaccount">FreezeAccount</a>| _required_ |Address|`"fadd"`|The address of the account whose asset is being frozen or unfrozen.|
|<a name="freezeasset">FreezeAsset</a>|_required_|uint64|`"faid"`| The asset ID being frozen or unfrozen.|
|<a name="assetfrozen">AssetFrozen</a>|_required_|bool|`"afrz"`| True to freeze the asset.|

# Application Call Transaction
Transaction Object Type: `ApplicationCallTx`

Includes all fields in [Header](#common-fields-header-and-type) and `"type"` is `"appl"`.

|Field|Required|Type|codec| Description|
|---|---|---|---|---|
| <a name="">Application ID</a>| _required_| uint64| `"apid"`| ID of the application being configured or empty if creating.|
| <a name="">OnComplete</a>| _required_| uint64| `"apan"`| Defines what additional actions occur with the transaction. See the [OnComplete](../../dapps/avm/teal/specification#oncomplete) section of the TEAL spec for details.|
| <a name="">Accounts</a>| _optional_| []Address| `"apat"`| List of accounts in addition to the sender that may be accessed from the application's approval-program and clear-state-program.|
| <a name="">Approval Program</a>| _optional_| []byte | `"apap"`| Logic executed for every application transaction, except when on-completion is set to "clear". It can read and write global state for the application, as well as account-specific local state. Approval programs may reject the transaction.|
| <a name="">App Arguments</a>| _optional_| [][]byte | `"apaa"`| Transaction specific arguments accessed from the application's approval-program and clear-state-program.|
| <a name="">Clear State Program</a>| _optional_| []byte | `"apsu"`| Logic executed for application transactions with on-completion set to "clear". It can read and write global state for the application, as well as account-specific local state. Clear state programs cannot reject the transaction.|
| <a name="">Foreign Apps</a>| _optional_| []uint64 | `"apfa"`| Lists the applications in addition to the application-id whose global states may be accessed by this application's approval-program and clear-state-program. The access is read-only.|
| <a name="">Foreign Assets</a>| _optional_| []uint64 | `"apas"`| Lists the assets whose AssetParams may be accessed by this application's approval-program and clear-state-program. The access is read-only.|
| <a name="">GlobalStateSchema</a>| _optional_| <a href=#storage-state-schema>StateSchema</a>| `"apgs"`| Holds the maximum number of global state values defined within a <a href=#storage-state-schema>StateSchema</a> object.|
| <a name="">LocalStateSchema</a>| _optional_| <a href=#storage-state-schema>StateSchema</a>| `"apls"`| Holds the maximum number of local state values defined within a <a href=#storage-state-schema>StateSchema</a> object.|
| <a name="">ExtraProgramPages</a>| _optional_| uint64 | `"apep"`| Number of additional pages allocated to the application's approval and clear state programs. Each `ExtraProgramPages` is 2048 bytes. The sum of `ApprovalProgram` and `ClearStateProgram` may not exceed 2048*(1+`ExtraProgramPages`) bytes. |
| <a name="">Boxes</a>| _optional_| []BoxRef | `"apbx"`| The boxes that should be made available for the runtime of the program. |

## Storage State Schema
Object Name: `StateSchema`

The `StateSchema` object is only required for the create application call transaction. The `StateSchema` object must be fully populated for both the `GlobalStateSchema` and `LocalStateSchema` objects.  

|Field|Required|Type|codec| Description|
|---|---|---|---|---|
| <a name="">Number Ints</a>| _required_| uint64| `"nui"`| Maximum number of integer values that may be stored in the [global \|\| local] application key/value store. Immutable.|
| <a name="">Number ByteSlices</a>| _required_| uint64| `"nbs"`| Maximum number of byte slices values that may be stored in the [global \|\| local] application key/value store. Immutable.|

# Signed Transaction
Transaction Object Type: `SignedTxn`

|Field|Required|Type|codec| Description|
|---|---|---|---|---|
|<a name="sig">Sig</a>| _required, if no other sig specified_ |crypto.Signature|`"sig"`||
|<a name="msig">Msig</a>|_required, if no other sig specified_|crypto.MultisigSig|`"msig"`||
|<a name="lsig">LogicSig</a>|_required, if no other sig specified_|LogicSig|`"lsig"`| |
|<a name="txn">Transaction</a>|_required_|Transaction|`"txn"`| [`PaymentTx`](#payment-transaction), [`KeyRegistrationTx`](#key-registration-transaction), [`AssetConfigTx`](#asset-configuration-transaction), [`AssetTransferTx`](#asset-transfer-transaction), [`AssetFreezeTx`](#asset-freeze-transaction) or [`ApplicationCallTx`](#application-call-transaction)
