title: indexer 🆕
---

<a name="paths"></a>
## Paths

<a name="makehealthcheck"></a>
### GET /health
Returns 200 if healthy.
```
GET /health
```


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|(empty)|[HealthCheck](#healthcheck)|
|**default**|Unknown Error|No Content|


**Produces**

* `application/json`


**Tags**

* common


<a name="searchforaccounts"></a>
### GET /v2/accounts

**Description**
Search for accounts.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Query**|**asset-id**  <br>*optional*|Asset ID|integer|
|**Query**|**auth-addr**  <br>*optional*|Include accounts configured to use this spending key.|string|
|**Query**|**currency-greater-than**  <br>*optional*|Results should have an amount greater than this value. MicroAlgos are the default currency unless an asset-id is provided, in which case the asset will be used.|integer|
|**Query**|**currency-less-than**  <br>*optional*|Results should have an amount less than this value. MicroAlgos are the default currency unless an asset-id is provided, in which case the asset will be used.|integer|
|**Query**|**limit**  <br>*optional*|Maximum number of results to return.|integer|
|**Query**|**next**  <br>*optional*|The next page of results. Use the next token provided by the previous results.|string|
|**Query**|**round**  <br>*optional*|Include results for the specified round. For performance reasons, this parameter may be disabled on some configurations.|integer|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|(empty)|[Response 200](#searchforaccounts-response-200)|

<a name="searchforaccounts-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**accounts**  <br>*required*||< [Account](#account) > array|
|**current-round**  <br>*required*|Round at which the results were computed.|integer|
|**next-token**  <br>*optional*|Used for pagination, when making another request provide this token with the next parameter.|string|


**Consumes**

* `application/json`


**Produces**

* `application/json`


**Tags**

* search


<a name="lookupaccountbyid"></a>
### GET /v2/accounts/{account-id}

**Description**
Lookup account information.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**account-id**  <br>*required*|account string|string|
|**Query**|**round**  <br>*optional*|Include results for the specified round.|integer|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|(empty)|[Response 200](#lookupaccountbyid-response-200)|

<a name="lookupaccountbyid-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**account**  <br>*required*||[Account](#account)|
|**current-round**  <br>*required*|Round at which the results were computed.|integer|


**Consumes**

* `application/json`


**Produces**

* `application/json`


**Tags**

* lookup


<a name="lookupaccounttransactions"></a>
### GET /v2/accounts/{account-id}/transactions

**Description**
Lookup account transactions.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**account-id**  <br>*required*|account string|string|
|**Query**|**after-time**  <br>*optional*|Include results after the given time. Must be an RFC 3339 formatted string.|string (date-time)|
|**Query**|**asset-id**  <br>*optional*|Asset ID|integer|
|**Query**|**before-time**  <br>*optional*|Include results before the given time. Must be an RFC 3339 formatted string.|string (date-time)|
|**Query**|**currency-greater-than**  <br>*optional*|Results should have an amount greater than this value. MicroAlgos are the default currency unless an asset-id is provided, in which case the asset will be used.|integer|
|**Query**|**currency-less-than**  <br>*optional*|Results should have an amount less than this value. MicroAlgos are the default currency unless an asset-id is provided, in which case the asset will be used.|integer|
|**Query**|**limit**  <br>*optional*|Maximum number of results to return.|integer|
|**Query**|**max-round**  <br>*optional*|Include results at or before the specified max-round.|integer|
|**Query**|**min-round**  <br>*optional*|Include results at or after the specified min-round.|integer|
|**Query**|**next**  <br>*optional*|The next page of results. Use the next token provided by the previous results.|string|
|**Query**|**note-prefix**  <br>*optional*|Specifies a prefix which must be contained in the note field.|string|
|**Query**|**rekey-to**  <br>*optional*|Include results which include the rekey-to field.|boolean|
|**Query**|**round**  <br>*optional*|Include results for the specified round.|integer|
|**Query**|**sig-type**  <br>*optional*|SigType filters just results using the specified type of signature:<br>* sig - Standard<br>* msig - MultiSig<br>* lsig - LogicSig|enum (sig, msig, lsig)|
|**Query**|**tx-type**  <br>*optional*||enum (pay, keyreg, acfg, axfer, afrz)|
|**Query**|**txid**  <br>*optional*|Lookup the specific transaction by ID.|string|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|(empty)|[Response 200](#lookupaccounttransactions-response-200)|

<a name="lookupaccounttransactions-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**current-round**  <br>*required*|Round at which the results were computed.|integer|
|**next-token**  <br>*optional*|Used for pagination, when making another request provide this token with the next parameter.|string|
|**transactions**  <br>*required*||< [Transaction](#transaction) > array|


**Consumes**

* `application/json`


**Produces**

* `application/json`


**Tags**

* lookup


<a name="searchforassets"></a>
### GET /v2/assets

**Description**
Search for assets.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Query**|**asset-id**  <br>*optional*|Asset ID|integer|
|**Query**|**creator**  <br>*optional*|Filter just assets with the given creator address.|string|
|**Query**|**limit**  <br>*optional*|Maximum number of results to return.|integer|
|**Query**|**name**  <br>*optional*|Filter just assets with the given name.|string|
|**Query**|**next**  <br>*optional*|The next page of results. Use the next token provided by the previous results.|string|
|**Query**|**unit**  <br>*optional*|Filter just assets with the given unit.|string|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|(empty)|[Response 200](#searchforassets-response-200)|

<a name="searchforassets-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**assets**  <br>*required*||< [Asset](#asset) > array|
|**current-round**  <br>*required*|Round at which the results were computed.|integer|
|**next-token**  <br>*optional*|Used for pagination, when making another request provide this token with the next parameter.|string|


**Consumes**

* `application/json`


**Produces**

* `application/json`


**Tags**

* search


<a name="lookupassetbyid"></a>
### GET /v2/assets/{asset-id}

**Description**
Lookup asset information.


**Parameters**

|Type|Name|Schema|
|---|---|---|
|**Path**|**asset-id**  <br>*required*|integer|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|(empty)|[Response 200](#lookupassetbyid-response-200)|

<a name="lookupassetbyid-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**asset**  <br>*required*||[Asset](#asset)|
|**current-round**  <br>*required*|Round at which the results were computed.|integer|


**Consumes**

* `application/json`


**Produces**

* `application/json`


**Tags**

* lookup


<a name="lookupassetbalances"></a>
### GET /v2/assets/{asset-id}/balances

**Description**
Lookup the list of accounts who hold this asset


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**asset-id**  <br>*required*||integer|
|**Query**|**currency-greater-than**  <br>*optional*|Results should have an amount greater than this value. MicroAlgos are the default currency unless an asset-id is provided, in which case the asset will be used.|integer|
|**Query**|**currency-less-than**  <br>*optional*|Results should have an amount less than this value. MicroAlgos are the default currency unless an asset-id is provided, in which case the asset will be used.|integer|
|**Query**|**limit**  <br>*optional*|Maximum number of results to return.|integer|
|**Query**|**next**  <br>*optional*|The next page of results. Use the next token provided by the previous results.|string|
|**Query**|**round**  <br>*optional*|Include results for the specified round.|integer|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|(empty)|[Response 200](#lookupassetbalances-response-200)|

<a name="lookupassetbalances-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**balances**  <br>*required*||< [MiniAssetHolding](#miniassetholding) > array|
|**current-round**  <br>*required*|Round at which the results were computed.|integer|
|**next-token**  <br>*optional*|Used for pagination, when making another request provide this token with the next parameter.|string|


**Consumes**

* `application/json`


**Produces**

* `application/json`


**Tags**

* lookup


<a name="lookupassettransactions"></a>
### GET /v2/assets/{asset-id}/transactions

**Description**
Lookup transactions for an asset.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**asset-id**  <br>*required*||integer|
|**Query**|**address**  <br>*optional*|Only include transactions with this address in one of the transaction fields.|string|
|**Query**|**address-role**  <br>*optional*|Combine with the address parameter to define what type of address to search for.|enum (sender, receiver, freeze-target)|
|**Query**|**after-time**  <br>*optional*|Include results after the given time. Must be an RFC 3339 formatted string.|string (date-time)|
|**Query**|**before-time**  <br>*optional*|Include results before the given time. Must be an RFC 3339 formatted string.|string (date-time)|
|**Query**|**currency-greater-than**  <br>*optional*|Results should have an amount greater than this value. MicroAlgos are the default currency unless an asset-id is provided, in which case the asset will be used.|integer|
|**Query**|**currency-less-than**  <br>*optional*|Results should have an amount less than this value. MicroAlgos are the default currency unless an asset-id is provided, in which case the asset will be used.|integer|
|**Query**|**exclude-close-to**  <br>*optional*|Combine with address and address-role parameters to define what type of address to search for. The close to fields are normally treated as a receiver, if you would like to exclude them set this parameter to true.|boolean|
|**Query**|**limit**  <br>*optional*|Maximum number of results to return.|integer|
|**Query**|**max-round**  <br>*optional*|Include results at or before the specified max-round.|integer|
|**Query**|**min-round**  <br>*optional*|Include results at or after the specified min-round.|integer|
|**Query**|**next**  <br>*optional*|The next page of results. Use the next token provided by the previous results.|string|
|**Query**|**note-prefix**  <br>*optional*|Specifies a prefix which must be contained in the note field.|string|
|**Query**|**rekey-to**  <br>*optional*|Include results which include the rekey-to field.|boolean|
|**Query**|**round**  <br>*optional*|Include results for the specified round.|integer|
|**Query**|**sig-type**  <br>*optional*|SigType filters just results using the specified type of signature:<br>* sig - Standard<br>* msig - MultiSig<br>* lsig - LogicSig|enum (sig, msig, lsig)|
|**Query**|**tx-type**  <br>*optional*||enum (pay, keyreg, acfg, axfer, afrz)|
|**Query**|**txid**  <br>*optional*|Lookup the specific transaction by ID.|string|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|(empty)|[Response 200](#lookupassettransactions-response-200)|

<a name="lookupassettransactions-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**current-round**  <br>*required*|Round at which the results were computed.|integer|
|**next-token**  <br>*optional*|Used for pagination, when making another request provide this token with the next parameter.|string|
|**transactions**  <br>*required*||< [Transaction](#transaction) > array|


**Consumes**

* `application/json`


**Produces**

* `application/json`


**Tags**

* lookup


<a name="lookupblock"></a>
### GET /v2/blocks/{round-number}

**Description**
Lookup block.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**round-number**  <br>*required*|Round number|integer|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|(empty)|[Block](#block)|


**Consumes**

* `application/json`


**Produces**

* `application/json`


**Tags**

* lookup


<a name="searchfortransactions"></a>
### GET /v2/transactions

**Description**
Search for transactions.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Query**|**address**  <br>*optional*|Only include transactions with this address in one of the transaction fields.|string|
|**Query**|**address-role**  <br>*optional*|Combine with the address parameter to define what type of address to search for.|enum (sender, receiver, freeze-target)|
|**Query**|**after-time**  <br>*optional*|Include results after the given time. Must be an RFC 3339 formatted string.|string (date-time)|
|**Query**|**asset-id**  <br>*optional*|Asset ID|integer|
|**Query**|**before-time**  <br>*optional*|Include results before the given time. Must be an RFC 3339 formatted string.|string (date-time)|
|**Query**|**currency-greater-than**  <br>*optional*|Results should have an amount greater than this value. MicroAlgos are the default currency unless an asset-id is provided, in which case the asset will be used.|integer|
|**Query**|**currency-less-than**  <br>*optional*|Results should have an amount less than this value. MicroAlgos are the default currency unless an asset-id is provided, in which case the asset will be used.|integer|
|**Query**|**exclude-close-to**  <br>*optional*|Combine with address and address-role parameters to define what type of address to search for. The close to fields are normally treated as a receiver, if you would like to exclude them set this parameter to true.|boolean|
|**Query**|**limit**  <br>*optional*|Maximum number of results to return.|integer|
|**Query**|**max-round**  <br>*optional*|Include results at or before the specified max-round.|integer|
|**Query**|**min-round**  <br>*optional*|Include results at or after the specified min-round.|integer|
|**Query**|**next**  <br>*optional*|The next page of results. Use the next token provided by the previous results.|string|
|**Query**|**note-prefix**  <br>*optional*|Specifies a prefix which must be contained in the note field.|string|
|**Query**|**rekey-to**  <br>*optional*|Include results which include the rekey-to field.|boolean|
|**Query**|**round**  <br>*optional*|Include results for the specified round.|integer|
|**Query**|**sig-type**  <br>*optional*|SigType filters just results using the specified type of signature:<br>* sig - Standard<br>* msig - MultiSig<br>* lsig - LogicSig|enum (sig, msig, lsig)|
|**Query**|**tx-type**  <br>*optional*||enum (pay, keyreg, acfg, axfer, afrz)|
|**Query**|**txid**  <br>*optional*|Lookup the specific transaction by ID.|string|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|(empty)|[Response 200](#searchfortransactions-response-200)|
|**500**|error|No Content|

<a name="searchfortransactions-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**current-round**  <br>*required*|Round at which the results were computed.|integer|
|**next-token**  <br>*optional*|Used for pagination, when making another request provide this token with the next parameter.|string|
|**transactions**  <br>*required*||< [Transaction](#transaction) > array|


**Consumes**

* `application/json`


**Produces**

* `application/json`


**Tags**

* search




<a name="definitions"></a>
## Definitions

<a name="account"></a>
### Account
Account information at a given round.

Definition:
data/basics/userBalance.go : AccountData


|Name|Description|Schema|
|---|---|---|
|**address**  <br>*required*|the account public key|string|
|**amount**  <br>*required*|\[algo\] total number of MicroAlgos in the account|integer|
|**amount-without-pending-rewards**  <br>*required*|specifies the amount of MicroAlgos in the account, without the pending rewards.|integer|
|**assets**  <br>*optional*|\[asset\] assets held by this account.<br><br>Note the raw object uses `map[int] -> AssetHolding` for this type.|< [AssetHolding](#assetholding) > array|
|**auth-addr**  <br>*optional*|\[spend\] the address against which signing should be checked. If empty, the address of the current account is used. This field can be updated in any transaction by setting the RekeyTo field.|string|
|**created-assets**  <br>*optional*|\[apar\] parameters of assets created by this account.<br><br>Note: the raw account uses `map[int] -> Asset` for this type.|< [Asset](#asset) > array|
|**participation**  <br>*optional*||[AccountParticipation](#accountparticipation)|
|**pending-rewards**  <br>*required*|amount of MicroAlgos of pending rewards in this account.|integer|
|**reward-base**  <br>*optional*|\[ebase\] used as part of the rewards computation. Only applicable to accounts which are participating.|integer|
|**rewards**  <br>*required*|\[ern\] total rewards of MicroAlgos the account has received, including pending rewards.|integer|
|**round**  <br>*required*|The round for which this information is relevant.|integer|
|**sig-type**  <br>*optional*|Indicates what type of signature is used by this account, must be one of:<br>* sig<br>* msig<br>* lsig|enum (sig, msig, lsig)|
|**status**  <br>*required*|\[onl\] delegation status of the account's MicroAlgos<br>* Offline - indicates that the associated account is delegated.<br>*  Online  - indicates that the associated account used as part of the delegation pool.<br>*   NotParticipating - indicates that the associated account is neither a delegator nor a delegate.|string|


<a name="accountparticipation"></a>
### AccountParticipation
AccountParticipation describes the parameters used by this account in consensus protocol.


|Name|Description|Schema|
|---|---|---|
|**selection-participation-key**  <br>*required*|\[sel\] Selection public key (if any) currently registered for this round.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**vote-first-valid**  <br>*required*|\[voteFst\] First round for which this participation is valid.|integer|
|**vote-key-dilution**  <br>*required*|\[voteKD\] Number of subkeys in each batch of participation keys.|integer|
|**vote-last-valid**  <br>*required*|\[voteLst\] Last round for which this participation is valid.|integer|
|**vote-participation-key**  <br>*required*|\[vote\] root participation public key (if any) currently registered for this round.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|


<a name="asset"></a>
### Asset
Specifies both the unique identifier and the parameters for an asset


|Name|Description|Schema|
|---|---|---|
|**index**  <br>*required*|unique asset identifier|integer|
|**params**  <br>*required*||[AssetParams](#assetparams)|


<a name="assetholding"></a>
### AssetHolding
Describes an asset held by an account.

Definition:
data/basics/userBalance.go : AssetHolding


|Name|Description|Schema|
|---|---|---|
|**amount**  <br>*required*|\[a\] number of units held.|integer|
|**asset-id**  <br>*required*|Asset ID of the holding.|integer|
|**creator**  <br>*required*|Address that created this asset. This is the address where the parameters for this asset can be found, and also the address where unwanted asset units can be sent in the worst case.|string|
|**is-frozen**  <br>*required*|\[f\] whether or not the holding is frozen.|boolean|


<a name="assetparams"></a>
### AssetParams
AssetParams specifies the parameters for an asset.

\[apar\] when part of an AssetConfig transaction.

Definition:
data/transactions/asset.go : AssetParams


|Name|Description|Schema|
|---|---|---|
|**clawback**  <br>*optional*|\[c\] Address of account used to clawback holdings of this asset.  If empty, clawback is not permitted.|string|
|**creator**  <br>*required*|The address that created this asset. This is the address where the parameters for this asset can be found, and also the address where unwanted asset units can be sent in the worst case.|string|
|**decimals**  <br>*required*|\[dc\] The number of digits to use after the decimal point when displaying this asset. If 0, the asset is not divisible. If 1, the base unit of the asset is in tenths. If 2, the base unit of the asset is in hundredths, and so on. This value must be between 0 and 19 (inclusive).  <br>**Minimum value** : `0`  <br>**Maximum value** : `19`|integer|
|**default-frozen**  <br>*optional*|\[df\] Whether holdings of this asset are frozen by default.|boolean|
|**freeze**  <br>*optional*|\[f\] Address of account used to freeze holdings of this asset.  If empty, freezing is not permitted.|string|
|**manager**  <br>*optional*|\[m\] Address of account used to manage the keys of this asset and to destroy it.|string|
|**metadata-hash**  <br>*optional*|\[am\] A commitment to some unspecified asset metadata. The format of this metadata is up to the application.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**name**  <br>*optional*|\[an\] Name of this asset, as supplied by the creator.|string|
|**reserve**  <br>*optional*|\[r\] Address of account holding reserve (non-minted) units of this asset.|string|
|**total**  <br>*required*|\[t\] The total number of units of this asset.|integer|
|**unit-name**  <br>*optional*|\[un\] Name of a unit of this asset, as supplied by the creator.|string|
|**url**  <br>*optional*|\[au\] URL where more information about the asset can be retrieved.|string|


<a name="block"></a>
### Block
Block information.

Definition:
data/bookkeeping/block.go : Block


|Name|Description|Schema|
|---|---|---|
|**genesis-hash**  <br>*required*|\[gh\] hash to which this block belongs.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**genesis-id**  <br>*required*|\[gen\] ID to which this block belongs.|string|
|**previous-block-hash**  <br>*required*|\[prev\] Previous block hash.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**rewards**  <br>*optional*||[BlockRewards](#blockrewards)|
|**round**  <br>*required*|\[rnd\] Current round on which this block was appended to the chain.|integer|
|**seed**  <br>*required*|\[seed\] Sortition seed.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**timestamp**  <br>*required*|\[ts\] Block creation timestamp in seconds since eposh|integer|
|**transactions**  <br>*optional*|\[txns\] list of transactions corresponding to a given round.|< [Transaction](#transaction) > array|
|**transactions-root**  <br>*required*|\[txn\] TransactionsRoot authenticates the set of transactions appearing in the block. More specifically, it's the root of a merkle tree whose leaves are the block's Txids, in lexicographic order. For the empty block, it's 0. Note that the TxnRoot does not authenticate the signatures on the transactions, only the transactions themselves. Two blocks with the same transactions but in a different order and with different signatures will have the same TxnRoot.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**txn-counter**  <br>*optional*|\[tc\] TxnCounter counts the number of transactions committed in the ledger, from the time at which support for this feature was introduced.<br><br>Specifically, TxnCounter is the number of the next transaction that will be committed after this block.  It is 0 when no transactions have ever been committed (since TxnCounter started being supported).|integer|
|**upgrade-state**  <br>*optional*||[BlockUpgradeState](#blockupgradestate)|
|**upgrade-vote**  <br>*optional*||[BlockUpgradeVote](#blockupgradevote)|


<a name="blockrewards"></a>
### BlockRewards
Fields relating to rewards,


|Name|Description|Schema|
|---|---|---|
|**fee-sink**  <br>*required*|\[fees\] accepts transaction fees, it can only spend to the incentive pool.|string|
|**rewards-calculation-round**  <br>*required*|\[rwcalr\] number of leftover MicroAlgos after the distribution of rewards-rate MicroAlgos for every reward unit in the next round.|integer|
|**rewards-level**  <br>*required*|\[earn\] How many rewards, in MicroAlgos, have been distributed to each RewardUnit of MicroAlgos since genesis.|integer|
|**rewards-pool**  <br>*required*|\[rwd\] accepts periodic injections from the fee-sink and continually redistributes them as rewards.|string|
|**rewards-rate**  <br>*required*|\[rate\] Number of new MicroAlgos added to the participation stake from rewards at the next round.|integer|
|**rewards-residue**  <br>*required*|\[frac\] Number of leftover MicroAlgos after the distribution of RewardsRate/rewardUnits MicroAlgos for every reward unit in the next round.|integer|


<a name="blockupgradestate"></a>
### BlockUpgradeState
Fields relating to a protocol upgrade.


|Name|Description|Schema|
|---|---|---|
|**current-protocol**  <br>*required*|\[proto\] The current protocol version.|string|
|**next-protocol**  <br>*optional*|\[nextproto\] The next proposed protocol version.|string|
|**next-protocol-approvals**  <br>*optional*|\[nextyes\] Number of blocks which approved the protocol upgrade.|integer|
|**next-protocol-switch-on**  <br>*optional*|\[nextswitch\] Round on which the protocol upgrade will take effect.|integer|
|**next-protocol-vote-before**  <br>*optional*|\[nextbefore\] Deadline round for this protocol upgrade (No votes will be consider after this round).|integer|


<a name="blockupgradevote"></a>
### BlockUpgradeVote
Fields relating to voting for a protocol upgrade.


|Name|Description|Schema|
|---|---|---|
|**upgrade-approve**  <br>*optional*|\[upgradeyes\] Indicates a yes vote for the current proposal.|boolean|
|**upgrade-delay**  <br>*optional*|\[upgradedelay\] Indicates the time between acceptance and execution.|integer|
|**upgrade-propose**  <br>*optional*|\[upgradeprop\] Indicates a proposed upgrade.|string|


<a name="errorresponse"></a>
### ErrorResponse
An error response with optional data field.


|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|


<a name="healthcheck"></a>
### HealthCheck
A health check response.


|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|


<a name="miniassetholding"></a>
### MiniAssetHolding
A simplified version of AssetHolding


|Name|Schema|
|---|---|
|**address**  <br>*required*|string|
|**amount**  <br>*required*|integer|
|**is-frozen**  <br>*required*|boolean|


<a name="transaction"></a>
### Transaction
Contains all fields common to all transactions and serves as an envelope to all transactions type.

Definition:
data/transactions/signedtxn.go : SignedTxn
data/transactions/transaction.go : Transaction


|Name|Description|Schema|
|---|---|---|
|**asset-config-transaction**  <br>*optional*||[TransactionAssetConfig](#transactionassetconfig)|
|**asset-freeze-transaction**  <br>*optional*||[TransactionAssetFreeze](#transactionassetfreeze)|
|**asset-transfer-transaction**  <br>*optional*||[TransactionAssetTransfer](#transactionassettransfer)|
|**auth-addr**  <br>*optional*|\[sgnr\] The address used to sign the transaction. This is used for rekeyed accounts to indicate that the sender address did not sign the transaction.|string|
|**close-rewards**  <br>*optional*|\[rc\] rewards applied to close-remainder-to account.|integer|
|**closing-amount**  <br>*optional*|\[ca\] closing amount for transaction.|integer|
|**confirmed-round**  <br>*optional*|Round when the transaction was confirmed.|integer|
|**created-asset-index**  <br>*optional*|Specifies an asset index (ID) if an asset was created with this transaction.|integer|
|**fee**  <br>*required*|\[fee\] Transaction fee.|integer|
|**first-valid**  <br>*required*|\[fv\] First valid round for this transaction.|integer|
|**genesis-hash**  <br>*optional*|\[gh\] Hash of genesis block.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**genesis-id**  <br>*optional*|\[gen\] genesis block ID.|string|
|**group**  <br>*optional*|\[grp\] Base64 encoded byte array of a sha512/256 digest. When present indicates that this transaction is part of a transaction group and the value is the sha512/256 hash of the transactions in that group.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**id**  <br>*required*|Transaction ID|string|
|**intra-round-offset**  <br>*optional*|Offset into the round where this transaction was confirmed.|integer|
|**keyreg-transaction**  <br>*optional*||[TransactionKeyreg](#transactionkeyreg)|
|**last-valid**  <br>*required*|\[lv\] Last valid round for this transaction.|integer|
|**lease**  <br>*optional*|\[lx\] Base64 encoded 32-byte array. Lease enforces mutual exclusion of transactions.  If this field is nonzero, then once the transaction is confirmed, it acquires the lease identified by the (Sender, Lease) pair of the transaction until the LastValid round passes.  While this transaction possesses the lease, no other transaction specifying this lease can be confirmed.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**note**  <br>*optional*|\[note\] Free form data.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**payment-transaction**  <br>*optional*||[TransactionPayment](#transactionpayment)|
|**receiver-rewards**  <br>*optional*|\[rr\] rewards applied to receiver account.|integer|
|**rekey-to**  <br>*optional*|\[rekey\] when included in a valid transaction, the accounts auth addr will be updated with this value and future signatures must be signed with the key represented by this address.|string|
|**round-time**  <br>*optional*|Time when the block this transaction is in was confirmed.|integer|
|**sender**  <br>*required*|\[snd\] Sender's address.|string|
|**sender-rewards**  <br>*optional*|\[rs\] rewards applied to sender account.|integer|
|**signature**  <br>*required*||[TransactionSignature](#transactionsignature)|
|**tx-type**  <br>*required*|\[type\] Indicates what type of transaction this is. Different types have different fields.<br><br>Valid types, and where their fields are stored:<br>* \[pay\] payment-transaction<br>* \[keyreg\] keyreg-transaction<br>* \[acfg\] asset-config-transaction<br>* \[axfer\] asset-transfer-transaction<br>* \[afrz\] asset-freeze-transaction|enum (pay, keyreg, acfg, axfer, afrz)|


<a name="transactionassetconfig"></a>
### TransactionAssetConfig
Fields for asset allocation, re-configuration, and destruction.


A zero value for asset-id indicates asset creation.
A zero value for the params indicates asset destruction.

Definition:
data/transactions/asset.go : AssetConfigTxnFields


|Name|Description|Schema|
|---|---|---|
|**asset-id**  <br>*optional*|\[xaid\] ID of the asset being configured or empty if creating.|integer|
|**params**  <br>*optional*||[AssetParams](#assetparams)|


<a name="transactionassetfreeze"></a>
### TransactionAssetFreeze
Fields for an asset freeze transaction.

Definition:
data/transactions/asset.go : AssetFreezeTxnFields


|Name|Description|Schema|
|---|---|---|
|**address**  <br>*required*|\[fadd\] Address of the account whose asset is being frozen or thawed.|string|
|**asset-id**  <br>*required*|\[faid\] ID of the asset being frozen or thawed.|integer|
|**new-freeze-status**  <br>*required*|\[afrz\] The new freeze status.|boolean|


<a name="transactionassettransfer"></a>
### TransactionAssetTransfer
Fields for an asset transfer transaction.

Definition:
data/transactions/asset.go : AssetTransferTxnFields


|Name|Description|Schema|
|---|---|---|
|**amount**  <br>*required*|\[aamt\] Amount of asset to transfer. A zero amount transferred to self allocates that asset in the account's Assets map.|integer|
|**asset-id**  <br>*required*|\[xaid\] ID of the asset being transferred.|integer|
|**close-amount**  <br>*optional*|Number of assets transfered to the close-to account as part of the transaction.|integer|
|**close-to**  <br>*optional*|\[aclose\] Indicates that the asset should be removed from the account's Assets map, and specifies where the remaining asset holdings should be transferred.  It's always valid to transfer remaining asset holdings to the creator account.|string|
|**receiver**  <br>*required*|\[arcv\] Recipient address of the transfer.|string|
|**sender**  <br>*optional*|\[asnd\] The effective sender during a clawback transactions. If this is not a zero value, the real transaction sender must be the Clawback address from the AssetParams.|string|


<a name="transactionkeyreg"></a>
### TransactionKeyreg
Fields for a keyreg transaction.

Definition:
data/transactions/keyreg.go : KeyregTxnFields


|Name|Description|Schema|
|---|---|---|
|**non-participation**  <br>*optional*|\[nonpart\] Mark the account as participating or non-participating.|boolean|
|**selection-participation-key**  <br>*optional*|\[selkey\] Public key used with the Verified Random Function (VRF) result during committee selection.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**vote-first-valid**  <br>*optional*|\[votefst\] First round this participation key is valid.|integer|
|**vote-key-dilution**  <br>*optional*|\[votekd\] Number of subkeys in each batch of participation keys.|integer|
|**vote-last-valid**  <br>*optional*|\[votelst\] Last round this participation key is valid.|integer|
|**vote-participation-key**  <br>*optional*|\[votekey\] Participation public key used in key registration transactions.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|


<a name="transactionpayment"></a>
### TransactionPayment
Fields for a payment transaction.

Definition:
data/transactions/payment.go : PaymentTxnFields


|Name|Description|Schema|
|---|---|---|
|**amount**  <br>*required*|\[amt\] number of MicroAlgos intended to be transferred.|integer|
|**close-amount**  <br>*optional*|Number of MicroAlgos that were sent to the close-remainder-to address when closing the sender account.|integer|
|**close-remainder-to**  <br>*optional*|\[close\] when set, indicates that the sending account should be closed and all remaining funds be transferred to this address.|string|
|**receiver**  <br>*required*|\[rcv\] receiver's address.|string|


<a name="transactionsignature"></a>
### TransactionSignature
Validation signature associated with some data. Only one of the signatures should be provided.


|Name|Description|Schema|
|---|---|---|
|**logicsig**  <br>*optional*||[TransactionSignatureLogicsig](#transactionsignaturelogicsig)|
|**multisig**  <br>*optional*||[TransactionSignatureMultisig](#transactionsignaturemultisig)|
|**sig**  <br>*optional*|\[sig\] Standard ed25519 signature.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|


<a name="transactionsignaturelogicsig"></a>
### TransactionSignatureLogicsig
\[lsig\] Programatic transaction signature.

Definition:
data/transactions/logicsig.go


|Name|Description|Schema|
|---|---|---|
|**args**  <br>*optional*|\[arg\] Logic arguments, base64 encoded.|< string > array|
|**logic**  <br>*required*|\[l\] Program signed by a signature or multi signature, or hashed to be the address of ana ccount. Base64 encoded TEAL program.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**multisig-signature**  <br>*optional*||[TransactionSignatureMultisig](#transactionsignaturemultisig)|
|**signature**  <br>*optional*|\[sig\] ed25519 signature.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|


<a name="transactionsignaturemultisig"></a>
### TransactionSignatureMultisig
\[msig\] structure holding multiple subsignatures.

Definition:
crypto/multisig.go : MultisigSig


|Name|Description|Schema|
|---|---|---|
|**subsignature**  <br>*optional*|\[subsig\] holds pairs of public key and signatures.|< [TransactionSignatureMultisigSubsignature](#transactionsignaturemultisigsubsignature) > array|
|**threshold**  <br>*optional*|\[thr\]|integer|
|**version**  <br>*optional*|\[v\]|integer|


<a name="transactionsignaturemultisigsubsignature"></a>
### TransactionSignatureMultisigSubsignature

|Name|Description|Schema|
|---|---|---|
|**public-key**  <br>*optional*|\[pk\]  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**signature**  <br>*optional*|\[s\]  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|



