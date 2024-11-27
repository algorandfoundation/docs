title: algod
---

<a name="paths"></a>
## Paths

<a name="getconfig"></a>
### GET /debug/settings/config
Gets the merged config file.
```
GET /debug/settings/config
```


**Description**
Returns the merged (defaults + overrides) config file in json.


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|The merged config file in json.|string|
|**default**|Unknown Error|No Content|


**Produces**

* `application/json`


**Tags**

* private


<a name="getdebugsettingsprof"></a>
### GET /debug/settings/pprof

**Description**
Retrieves the current settings for blocking and mutex profiles


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|DebugPprof is the response to the /debug/extra/pprof endpoint|[DebugSettingsProf](#debugsettingsprof)|


**Produces**

* `application/json`


**Tags**

* private


<a name="putdebugsettingsprof"></a>
### PUT /debug/settings/pprof

**Description**
Enables blocking and mutex profiles, and returns the old settings


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|DebugPprof is the response to the /debug/extra/pprof endpoint|[DebugSettingsProf](#debugsettingsprof)|


**Produces**

* `application/json`


**Tags**

* private


<a name="getgenesis"></a>
### GET /genesis
Gets the genesis information.
```
GET /genesis
```


**Description**
Returns the entire genesis file in json.


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|The genesis file in json.|string|
|**default**|Unknown Error|No Content|


**Produces**

* `application/json`


**Tags**

* common
* public


<a name="healthcheck"></a>
### GET /health
Returns OK if healthy.
```
GET /health
```


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|OK.|No Content|
|**default**|Unknown Error|No Content|


**Produces**

* `application/json`


**Tags**

* common
* public


<a name="metrics"></a>
### GET /metrics
Return metrics about algod functioning.
```
GET /metrics
```


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|text with \#-comments and key:value lines|No Content|
|**404**|metrics were compiled out|No Content|


**Produces**

* `text/plain`


**Tags**

* common
* public


<a name="getready"></a>
### GET /ready
Returns OK if healthy and fully caught up.
```
GET /ready
```


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|OK.|No Content|
|**500**|Internal Error|No Content|
|**503**|Node not ready yet|No Content|
|**default**|Unknown Error|No Content|


**Produces**

* `application/json`


**Tags**

* common
* public


<a name="swaggerjson"></a>
### GET /swagger.json
Gets the current swagger spec.
```
GET /swagger.json
```


**Description**
Returns the entire swagger spec in json.


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|The current swagger spec|string|
|**default**|Unknown Error|No Content|


**Produces**

* `application/json`


**Tags**

* common
* public


<a name="accountinformation"></a>
### GET /v2/accounts/{address}
Get account information.
```
GET /v2/accounts/{address}
```


**Description**
Given a specific account public key, this call returns the accounts status, balance and spendable amounts


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**address**  <br>*required*|An account public key|string|
|**Query**|**exclude**  <br>*optional*|When set to `all` will exclude asset holdings, application local state, created asset parameters, any created application parameters. Defaults to `none`.|enum (all, none)|
|**Query**|**format**  <br>*optional*|Configures whether the response object is JSON or MessagePack encoded. If not provided, defaults to JSON.|enum (json, msgpack)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|AccountResponse wraps the Account type in a response.|[Account](#account)|
|**400**|Bad request|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|


**Produces**

* `application/json`
* `application/msgpack`


**Tags**

* nonparticipating
* public


<a name="accountapplicationinformation"></a>
### GET /v2/accounts/{address}/applications/{application-id}
Get account information about a given app.
```
GET /v2/accounts/{address}/applications/{application-id}
```


**Description**
Given a specific account public key and application ID, this call returns the account's application local state and global state (AppLocalState and AppParams, if either exists). Global state will only be returned if the provided address is the application's creator.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**address**  <br>*required*|An account public key|string|
|**Path**|**application-id**  <br>*required*|An application identifier|integer|
|**Query**|**format**  <br>*optional*|Configures whether the response object is JSON or MessagePack encoded. If not provided, defaults to JSON.|enum (json, msgpack)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|AccountApplicationResponse describes the account's application local state and global state (AppLocalState and AppParams, if either exists) for a specific application ID. Global state will only be returned if the provided address is the application's creator.|[Response 200](#accountapplicationinformation-response-200)|
|**400**|Malformed address or application ID|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|

<a name="accountapplicationinformation-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**app-local-state**  <br>*optional*|\[appl\] the application local data stored in this account.<br><br>The raw account uses `AppLocalState` for this type.|[ApplicationLocalState](#applicationlocalstate)|
|**created-app**  <br>*optional*|\[appp\] parameters of the application created by this account including app global data.<br><br>The raw account uses `AppParams` for this type.|[ApplicationParams](#applicationparams)|
|**round**  <br>*required*|The round for which this information is relevant.|integer|


**Produces**

* `application/json`
* `application/msgpack`


**Tags**

* nonparticipating
* public


<a name="accountassetsinformation"></a>
### GET /v2/accounts/{address}/assets
Get a list of assets held by an account, inclusive of asset params.
```
GET /v2/accounts/{address}/assets
```


**Description**
Lookup an account's asset holdings.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**address**  <br>*required*|An account public key|string|
|**Query**|**limit**  <br>*optional*|Maximum number of results to return.|integer|
|**Query**|**next**  <br>*optional*|The next page of results. Use the next token provided by the previous results.|string|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|AccountAssetsInformationResponse contains a list of assets held by an account.|[Response 200](#accountassetsinformation-response-200)|
|**400**|Malformed address|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|

<a name="accountassetsinformation-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**asset-holdings**  <br>*optional*||< [AccountAssetHolding](#accountassetholding) > array|
|**next-token**  <br>*optional*|Used for pagination, when making another request provide this token with the next parameter.|string|
|**round**  <br>*required*|The round for which this information is relevant.|integer|


**Produces**

* `application/json`


**Tags**

* experimental
* public


<a name="accountassetinformation"></a>
### GET /v2/accounts/{address}/assets/{asset-id}
Get account information about a given asset.
```
GET /v2/accounts/{address}/assets/{asset-id}
```


**Description**
Given a specific account public key and asset ID, this call returns the account's asset holding and asset parameters (if either exist). Asset parameters will only be returned if the provided address is the asset's creator.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**address**  <br>*required*|An account public key|string|
|**Path**|**asset-id**  <br>*required*|An asset identifier|integer|
|**Query**|**format**  <br>*optional*|Configures whether the response object is JSON or MessagePack encoded. If not provided, defaults to JSON.|enum (json, msgpack)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|AccountAssetResponse describes the account's asset holding and asset parameters (if either exist) for a specific asset ID. Asset parameters will only be returned if the provided address is the asset's creator.|[Response 200](#accountassetinformation-response-200)|
|**400**|Malformed address or asset ID|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|

<a name="accountassetinformation-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**asset-holding**  <br>*optional*|\[asset\] Details about the asset held by this account.<br><br>The raw account uses `AssetHolding` for this type.|[AssetHolding](#assetholding)|
|**created-asset**  <br>*optional*|\[apar\] parameters of the asset created by this account.<br><br>The raw account uses `AssetParams` for this type.|[AssetParams](#assetparams)|
|**round**  <br>*required*|The round for which this information is relevant.|integer|


**Produces**

* `application/json`
* `application/msgpack`


**Tags**

* nonparticipating
* public


<a name="getpendingtransactionsbyaddress"></a>
### GET /v2/accounts/{address}/transactions/pending
Get a list of unconfirmed transactions currently in the transaction pool by address.
```
GET /v2/accounts/{address}/transactions/pending
```


**Description**
Get the list of pending transactions by address, sorted by priority, in decreasing order, truncated at the end at MAX. If MAX = 0, returns all pending transactions.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**address**  <br>*required*|An account public key|string|
|**Query**|**format**  <br>*optional*|Configures whether the response object is JSON or MessagePack encoded. If not provided, defaults to JSON.|enum (json, msgpack)|
|**Query**|**max**  <br>*optional*|Truncated number of transactions to display. If max=0, returns all pending txns.|integer|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|A potentially truncated list of transactions currently in the node's transaction pool. You can compute whether or not the list is truncated if the number of elements in the **top-transactions** array is fewer than **total-transactions**.|[Response 200](#getpendingtransactionsbyaddress-response-200)|
|**400**|Max must be a non-negative integer|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**503**|Service Temporarily Unavailable|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|

<a name="getpendingtransactionsbyaddress-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**top-transactions**  <br>*required*|An array of signed transaction objects.|< object > array|
|**total-transactions**  <br>*required*|Total number of transactions in the pool.|integer|


**Produces**

* `application/json`
* `application/msgpack`


**Tags**

* participating
* public


<a name="getapplicationbyid"></a>
### GET /v2/applications/{application-id}
Get application information.
```
GET /v2/applications/{application-id}
```


**Description**
Given a application ID, it returns application information including creator, approval and clear programs, global and local schemas, and global state.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**application-id**  <br>*required*|An application identifier|integer|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Application information|[Application](#application)|
|**400**|Bad Request|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**404**|Application Not Found|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|


**Produces**

* `application/json`


**Tags**

* nonparticipating
* public


<a name="getapplicationboxbyname"></a>
### GET /v2/applications/{application-id}/box
Get box information for a given application.
```
GET /v2/applications/{application-id}/box
```


**Description**
Given an application ID and box name, it returns the round, box name, and value (each base64 encoded). Box names must be in the goal app call arg encoding form 'encoding:value'. For ints, use the form 'int:1234'. For raw bytes, use the form 'b64:A=='. For printable strings, use the form 'str:hello'. For addresses, use the form 'addr:XYZ...'.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**application-id**  <br>*required*|An application identifier|integer|
|**Query**|**name**  <br>*required*|A box name, in the goal app call arg form 'encoding:value'. For ints, use the form 'int:1234'. For raw bytes, use the form 'b64:A=='. For printable strings, use the form 'str:hello'. For addresses, use the form 'addr:XYZ...'.|string|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Box information|[Box](#box)|
|**400**|Bad Request|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**404**|Box Not Found|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|


**Produces**

* `application/json`


**Tags**

* nonparticipating
* public


<a name="getapplicationboxes"></a>
### GET /v2/applications/{application-id}/boxes
Get all box names for a given application.
```
GET /v2/applications/{application-id}/boxes
```


**Description**
Given an application ID, return all Box names. No particular ordering is guaranteed. Request fails when client or server-side configured limits prevent returning all Box names.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**application-id**  <br>*required*|An application identifier|integer|
|**Query**|**max**  <br>*optional*|Max number of box names to return. If max is not set, or max == 0, returns all box-names.|integer|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Box names of an application|[Response 200](#getapplicationboxes-response-200)|
|**400**|Bad Request|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|

<a name="getapplicationboxes-response-200"></a>
**Response 200**

|Name|Schema|
|---|---|
|**boxes**  <br>*required*|< [BoxDescriptor](#boxdescriptor) > array|


**Produces**

* `application/json`


**Tags**

* nonparticipating
* public


<a name="getassetbyid"></a>
### GET /v2/assets/{asset-id}
Get asset information.
```
GET /v2/assets/{asset-id}
```


**Description**
Given a asset ID, it returns asset information including creator, name, total supply and special addresses.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**asset-id**  <br>*required*|An asset identifier|integer|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Asset information|[Asset](#asset)|
|**400**|Bad Request|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**404**|Application Not Found|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|


**Produces**

* `application/json`


**Tags**

* nonparticipating
* public


<a name="getblock"></a>
### GET /v2/blocks/{round}
Get the block for the given round.
```
GET /v2/blocks/{round}
```


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**round**  <br>*required*|The round from which to fetch block information.|integer|
|**Query**|**format**  <br>*optional*|Configures whether the response object is JSON or MessagePack encoded. If not provided, defaults to JSON.|enum (json, msgpack)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Encoded block object.|[Response 200](#getblock-response-200)|
|**400**|Bad Request - Non integer number|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**404**|None existing block|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|

<a name="getblock-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**block**  <br>*required*|Block header data.|object|
|**cert**  <br>*optional*|Optional certificate object. This is only included when the format is set to message pack.|object|


**Produces**

* `application/json`
* `application/msgpack`


**Tags**

* nonparticipating
* public


<a name="getblockhash"></a>
### GET /v2/blocks/{round}/hash
Get the block hash for the block on the given round.
```
GET /v2/blocks/{round}/hash
```


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**round**  <br>*required*|The round from which to fetch block hash information.|integer|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Hash of a block header.|[Response 200](#getblockhash-response-200)|
|**400**|Bad Request - Non integer number|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**404**|None existing block|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|

<a name="getblockhash-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**blockHash**  <br>*required*|Block header hash.|string|


**Produces**

* `application/json`


**Tags**

* nonparticipating
* public


<a name="getblockheader"></a>
### GET /v2/blocks/{round}/header
Get the block header for the block on the given round.
```
GET /v2/blocks/{round}/header
```


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**round**  <br>*required*|The round from which to fetch block header information.|integer|
|**Query**|**format**  <br>*optional*|Configures whether the response object is JSON or MessagePack encoded. If not provided, defaults to JSON.|enum (json, msgpack)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Block header.|[Response 200](#getblockheader-response-200)|
|**400**|Bad Request - Non integer number|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**404**|None existing block|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|

<a name="getblockheader-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**blockHeader**  <br>*required*|Block header data.|object|


**Produces**

* `application/json`
* `application/msgpack`


**Tags**

* nonparticipating
* public


<a name="getlightblockheaderproof"></a>
### GET /v2/blocks/{round}/lightheader/proof
Gets a proof for a given light block header inside a state proof commitment
```
GET /v2/blocks/{round}/lightheader/proof
```


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**round**  <br>*required*|The round to which the light block header belongs.|integer|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Proof of a light block header.|[LightBlockHeaderProof](#lightblockheaderproof)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**404**|Could not create proof since some data is missing|[ErrorResponse](#errorresponse)|
|**408**|timed out on request|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**503**|Service Temporarily Unavailable|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|


**Produces**

* `application/json`


**Tags**

* nonparticipating
* public


<a name="getblocklogs"></a>
### GET /v2/blocks/{round}/logs
Get all of the logs from outer and inner app calls in the given round
```
GET /v2/blocks/{round}/logs
```


**Description**
Get all of the logs from outer and inner app calls in the given round


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**round**  <br>*required*|The round from which to fetch block log information.|integer|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|All logs emitted in the given round. Each app call, whether top-level or inner, that contains logs results in a separate AppCallLogs object. Therefore there may be multiple AppCallLogs with the same application ID and outer transaction ID in the event of multiple inner app calls to the same app. App calls with no logs are not included in the response. AppCallLogs are returned in the same order that their corresponding app call appeared in the block (pre-order traversal of inner app calls)|[Response 200](#getblocklogs-response-200)|
|**400**|Bad Request - Non integer number|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**404**|Nonexistent block|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|

<a name="getblocklogs-response-200"></a>
**Response 200**

|Name|Schema|
|---|---|
|**logs**  <br>*required*|< [AppCallLogs](#appcalllogs) > array|


**Produces**

* `application/json`


**Tags**

* nonparticipating
* public


<a name="gettransactionproof"></a>
### GET /v2/blocks/{round}/transactions/{txid}/proof
Get a proof for a transaction in a block.
```
GET /v2/blocks/{round}/transactions/{txid}/proof
```


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**round**  <br>*required*|The round in which the transaction appears.|integer|
|**Path**|**txid**  <br>*required*|The transaction ID for which to generate a proof.|string|
|**Query**|**format**  <br>*optional*|Configures whether the response object is JSON or MessagePack encoded. If not provided, defaults to JSON.|enum (json, msgpack)|
|**Query**|**hashtype**  <br>*optional*|The type of hash function used to create the proof, must be one of: <br>* sha512_256 <br>* sha256|enum (sha512_256, sha256)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Proof of transaction in a block.|[Response 200](#gettransactionproof-response-200)|
|**400**|Malformed round number or transaction ID|[ErrorResponse](#errorresponse)|
|**401**|Invalid API token|[ErrorResponse](#errorresponse)|
|**404**|Non-existent block or transaction|[ErrorResponse](#errorresponse)|
|**500**|Internal error, including protocol not supporting proofs.|[ErrorResponse](#errorresponse)|
|**default**|Unknown error|No Content|

<a name="gettransactionproof-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**hashtype**  <br>*required*|The type of hash function used to create the proof, must be one of: <br>* sha512_256 <br>* sha256|enum (sha512_256, sha256)|
|**idx**  <br>*required*|Index of the transaction in the block's payset.|integer|
|**proof**  <br>*required*|Proof of transaction membership.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**stibhash**  <br>*required*|Hash of SignedTxnInBlock for verifying proof.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**treedepth**  <br>*required*|Represents the depth of the tree that is being proven, i.e. the number of edges from a leaf to the root.|integer|


**Produces**

* `application/json`


**Tags**

* nonparticipating
* public


<a name="getblocktxids"></a>
### GET /v2/blocks/{round}/txids
Get the top level transaction IDs for the block on the given round.
```
GET /v2/blocks/{round}/txids
```


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**round**  <br>*required*|The round from which to fetch block transaction IDs.|integer|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Top level transaction IDs in a block.|[Response 200](#getblocktxids-response-200)|
|**400**|Bad Request - Non integer number|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**404**|Non existing block|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|

<a name="getblocktxids-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**blockTxids**  <br>*required*|Block transaction IDs.|< string > array|


**Produces**

* `application/json`


**Tags**

* nonparticipating
* public


<a name="startcatchup"></a>
### POST /v2/catchup/{catchpoint}
Starts a catchpoint catchup.
```
POST /v2/catchup/{catchpoint}
```


**Description**
Given a catchpoint, it starts catching up to this catchpoint


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**catchpoint**  <br>*required*|A catch point|string (catchpoint)|
|**Query**|**min**  <br>*optional*|Specify the minimum number of blocks which the ledger must be advanced by in order to start the catchup. This is useful for simplifying tools which support fast catchup, they can run the catchup unconditionally and the node will skip the catchup if it is not needed.|integer|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**||[Response 200](#startcatchup-response-200)|
|**201**||[Response 201](#startcatchup-response-201)|
|**400**|Bad Request|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**408**|Request Timeout|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|

<a name="startcatchup-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**catchup-message**  <br>*required*|Catchup start response string|string|

<a name="startcatchup-response-201"></a>
**Response 201**

|Name|Description|Schema|
|---|---|---|
|**catchup-message**  <br>*required*|Catchup start response string|string|


**Produces**

* `application/json`


**Tags**

* nonparticipating
* private


<a name="abortcatchup"></a>
### DELETE /v2/catchup/{catchpoint}
Aborts a catchpoint catchup.
```
DELETE /v2/catchup/{catchpoint}
```


**Description**
Given a catchpoint, it aborts catching up to this catchpoint


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**catchpoint**  <br>*required*|A catch point|string (catchpoint)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**||[Response 200](#abortcatchup-response-200)|
|**400**|Bad Request|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|

<a name="abortcatchup-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**catchup-message**  <br>*required*|Catchup abort response string|string|


**Produces**

* `application/json`


**Tags**

* nonparticipating
* private


<a name="getledgerstatedeltafortransactiongroup"></a>
### GET /v2/deltas/txn/group/{id}
Get a LedgerStateDelta object for a given transaction group
```
GET /v2/deltas/txn/group/{id}
```


**Description**
Get a ledger delta for a given transaction group.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**id**  <br>*required*|A transaction ID, or transaction group ID|string|
|**Query**|**format**  <br>*optional*|Configures whether the response object is JSON or MessagePack encoded. If not provided, defaults to JSON.|enum (json, msgpack)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Response containing a ledger state delta for a single transaction group.|[LedgerStateDelta](#ledgerstatedelta)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**404**|Could not find a delta for transaction ID or group ID|[ErrorResponse](#errorresponse)|
|**408**|timed out on request|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**501**|Not Implemented|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|


**Produces**

* `application/json`
* `application/msgpack`


**Tags**

* nonparticipating
* public


<a name="getledgerstatedelta"></a>
### GET /v2/deltas/{round}
Get a LedgerStateDelta object for a given round
```
GET /v2/deltas/{round}
```


**Description**
Get ledger deltas for a round.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**round**  <br>*required*|The round for which the deltas are desired.|integer|
|**Query**|**format**  <br>*optional*|Configures whether the response object is JSON or MessagePack encoded. If not provided, defaults to JSON.|enum (json, msgpack)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Contains ledger deltas|[LedgerStateDelta](#ledgerstatedelta)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**404**|Could not find a delta for round|[ErrorResponse](#errorresponse)|
|**408**|timed out on request|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**503**|Service Temporarily Unavailable|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|


**Produces**

* `application/json`
* `application/msgpack`


**Tags**

* nonparticipating
* public


<a name="gettransactiongroupledgerstatedeltasforround"></a>
### GET /v2/deltas/{round}/txn/group
Get LedgerStateDelta objects for all transaction groups in a given round
```
GET /v2/deltas/{round}/txn/group
```


**Description**
Get ledger deltas for transaction groups in a given round.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**round**  <br>*required*|The round for which the deltas are desired.|integer|
|**Query**|**format**  <br>*optional*|Configures whether the response object is JSON or MessagePack encoded. If not provided, defaults to JSON.|enum (json, msgpack)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Response containing all ledger state deltas for transaction groups, with their associated Ids, in a single round.|[Response 200](#gettransactiongroupledgerstatedeltasforround-response-200)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**404**|Could not find deltas for round|[ErrorResponse](#errorresponse)|
|**408**|timed out on request|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**501**|Not Implemented|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|

<a name="gettransactiongroupledgerstatedeltasforround-response-200"></a>
**Response 200**

|Name|Schema|
|---|---|
|**Deltas**  <br>*required*|< [LedgerStateDeltaForTransactionGroup](#ledgerstatedeltafortransactiongroup) > array|


**Produces**

* `application/json`
* `application/msgpack`


**Tags**

* nonparticipating
* public


<a name="getblocktimestampoffset"></a>
### GET /v2/devmode/blocks/offset
Returns the timestamp offset. Timestamp offsets can only be set in dev mode.
```
GET /v2/devmode/blocks/offset
```


**Description**
Gets the current timestamp offset.


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Response containing the timestamp offset in seconds|[Response 200](#getblocktimestampoffset-response-200)|
|**400**|TimeStamp offset not set.|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|

<a name="getblocktimestampoffset-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**offset**  <br>*required*|Timestamp offset in seconds.|integer|


**Produces**

* `application/json`


**Tags**

* nonparticipating
* public


<a name="setblocktimestampoffset"></a>
### POST /v2/devmode/blocks/offset/{offset}
Given a timestamp offset in seconds, adds the offset to every subsequent block header's timestamp.
```
POST /v2/devmode/blocks/offset/{offset}
```


**Description**
Sets the timestamp offset (seconds) for blocks in dev mode. Providing an offset of 0 will unset this value and try to use the real clock for the timestamp.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**offset**  <br>*required*|The timestamp offset for blocks in dev mode.|integer|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|OK|No Content|
|**400**|Cannot set timestamp offset to a negative integer.|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|


**Tags**

* nonparticipating
* public


<a name="experimentalcheck"></a>
### GET /v2/experimental
Returns OK if experimental API is enabled.
```
GET /v2/experimental
```


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Experimental API enabled|No Content|
|**404**|Experimental API not enabled|No Content|
|**default**|Unknown Error|No Content|


**Produces**

* `application/json`


**Tags**

* experimental
* public


<a name="getsupply"></a>
### GET /v2/ledger/supply
Get the current supply reported by the ledger.
```
GET /v2/ledger/supply
```


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Supply represents the current supply of MicroAlgos in the system.|[Response 200](#getsupply-response-200)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|

<a name="getsupply-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**current_round**  <br>*required*|Round|integer|
|**online-money**  <br>*required*|OnlineMoney|integer|
|**total-money**  <br>*required*|TotalMoney|integer|


**Produces**

* `application/json`


**Tags**

* nonparticipating
* public


<a name="getsyncround"></a>
### GET /v2/ledger/sync
Returns the minimum sync round the ledger is keeping in cache.
```
GET /v2/ledger/sync
```


**Description**
Gets the minimum sync round for the ledger.


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Response containing the ledger's minimum sync round|[Response 200](#getsyncround-response-200)|
|**400**|Sync round not set.|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**503**|Service Temporarily Unavailable|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|

<a name="getsyncround-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**round**  <br>*required*|The minimum sync round for the ledger.|integer|


**Tags**

* data
* public


<a name="unsetsyncround"></a>
### DELETE /v2/ledger/sync
Removes minimum sync round restriction from the ledger.
```
DELETE /v2/ledger/sync
```


**Description**
Unset the ledger sync round.


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**||No Content|
|**400**|Sync round not set.|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**503**|Service Temporarily Unavailable|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|


**Tags**

* data
* public


<a name="setsyncround"></a>
### POST /v2/ledger/sync/{round}
Given a round, tells the ledger to keep that round in its cache.
```
POST /v2/ledger/sync/{round}
```


**Description**
Sets the minimum sync round on the ledger.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**round**  <br>*required*|The round for which the deltas are desired.|integer|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**||No Content|
|**400**|Can not set sync round to an earlier round than the current round.|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**503**|Service Temporarily Unavailable|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|


**Tags**

* data
* public


<a name="addparticipationkey"></a>
### POST /v2/participation
Add a participation key to the node
```
POST /v2/participation
```


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Body**|**participationkey**  <br>*required*|The participation key to add to the node|string (binary)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Participation ID of the submission|[Response 200](#addparticipationkey-response-200)|
|**400**|Bad Request|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**404**|Participation Key Not Found|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**503**|Service Temporarily Unavailable|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|

<a name="addparticipationkey-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**partId**  <br>*required*|encoding of the participation ID.|string|


**Consumes**

* `application/msgpack`


**Produces**

* `application/json`


**Tags**

* participating
* private


<a name="getparticipationkeys"></a>
### GET /v2/participation
Return a list of participation keys
```
GET /v2/participation
```


**Description**
Return a list of participation keys


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|A list of participation keys|< [ParticipationKey](#participationkey) > array|
|**400**|Bad Request|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**404**|Participation Key Not Found|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|


**Produces**

* `application/json`


**Tags**

* participating
* private


<a name="generateparticipationkeys"></a>
### POST /v2/participation/generate/{address}
Generate and install participation keys to the node.
```
POST /v2/participation/generate/{address}
```


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**address**  <br>*required*|An account public key|string|
|**Query**|**dilution**  <br>*optional*|Key dilution for two-level participation keys (defaults to sqrt of validity window).|integer|
|**Query**|**first**  <br>*required*|First round for participation key.|integer|
|**Query**|**last**  <br>*required*|Last round for participation key.|integer|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|An empty JSON object is returned if the generation process was started. Currently no status is available.|string|
|**400**|Bad Request|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**503**|Service Temporarily Unavailable|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|


**Produces**

* `application/json`


**Tags**

* participating
* private


<a name="appendkeys"></a>
### POST /v2/participation/{participation-id}
Append state proof keys to a participation key
```
POST /v2/participation/{participation-id}
```


**Description**
Given a participation ID, append state proof keys to a particular set of participation keys


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**participation-id**  <br>*required*||string|
|**Body**|**keymap**  <br>*required*|The state proof keys to add to an existing participation ID|string (binary)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|A detailed description of a participation ID|[ParticipationKey](#participationkey)|
|**400**|Bad Request|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**404**|Participation Key Not Found|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|


**Consumes**

* `application/msgpack`


**Produces**

* `application/json`


**Tags**

* participating
* private


<a name="getparticipationkeybyid"></a>
### GET /v2/participation/{participation-id}
Get participation key info given a participation ID
```
GET /v2/participation/{participation-id}
```


**Description**
Given a participation ID, return information about that participation key


**Parameters**

|Type|Name|Schema|
|---|---|---|
|**Path**|**participation-id**  <br>*required*|string|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|A detailed description of a participation ID|[ParticipationKey](#participationkey)|
|**400**|Bad Request|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**404**|Participation Key Not Found|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|


**Produces**

* `application/json`


**Tags**

* participating
* private


<a name="deleteparticipationkeybyid"></a>
### DELETE /v2/participation/{participation-id}
Delete a given participation key by ID
```
DELETE /v2/participation/{participation-id}
```


**Description**
Delete a given participation key by ID


**Parameters**

|Type|Name|Schema|
|---|---|---|
|**Path**|**participation-id**  <br>*required*|string|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Participation key got deleted by ID|No Content|
|**400**|Bad Request|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**404**|Participation Key Not Found|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|


**Produces**

* `application/json`


**Tags**

* participating
* private


<a name="shutdownnode"></a>
### POST /v2/shutdown

**Description**
Special management endpoint to shutdown the node. Optionally provide a timeout parameter to indicate that the node should begin shutting down after a number of seconds.


**Parameters**

|Type|Name|Schema|Default|
|---|---|---|---|
|**Query**|**timeout**  <br>*optional*|integer|`0`|


**Responses**

|HTTP Code|Schema|
|---|---|
|**200**|object|


**Tags**

* nonparticipating
* private


<a name="getstateproof"></a>
### GET /v2/stateproofs/{round}
Get a state proof that covers a given round
```
GET /v2/stateproofs/{round}
```


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**round**  <br>*required*|The round for which a state proof is desired.|integer|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|StateProofResponse wraps the StateProof type in a response.|[StateProof](#stateproof)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**404**|Could not find a state proof that covers a given round|[ErrorResponse](#errorresponse)|
|**408**|timed out on request|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**503**|Service Temporarily Unavailable|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|


**Produces**

* `application/json`


**Tags**

* nonparticipating
* public


<a name="getstatus"></a>
### GET /v2/status
Gets the current node status.
```
GET /v2/status
```


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**||[Response 200](#getstatus-response-200)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|string|
|**default**|Unknown Error|No Content|

<a name="getstatus-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**catchpoint**  <br>*optional*|The current catchpoint that is being caught up to|string|
|**catchpoint-acquired-blocks**  <br>*optional*|The number of blocks that have already been obtained by the node as part of the catchup|integer|
|**catchpoint-processed-accounts**  <br>*optional*|The number of accounts from the current catchpoint that have been processed so far as part of the catchup|integer|
|**catchpoint-processed-kvs**  <br>*optional*|The number of key-values (KVs) from the current catchpoint that have been processed so far as part of the catchup|integer|
|**catchpoint-total-accounts**  <br>*optional*|The total number of accounts included in the current catchpoint|integer|
|**catchpoint-total-blocks**  <br>*optional*|The total number of blocks that are required to complete the current catchpoint catchup|integer|
|**catchpoint-total-kvs**  <br>*optional*|The total number of key-values (KVs) included in the current catchpoint|integer|
|**catchpoint-verified-accounts**  <br>*optional*|The number of accounts from the current catchpoint that have been verified so far as part of the catchup|integer|
|**catchpoint-verified-kvs**  <br>*optional*|The number of key-values (KVs) from the current catchpoint that have been verified so far as part of the catchup|integer|
|**catchup-time**  <br>*required*|CatchupTime in nanoseconds|integer|
|**last-catchpoint**  <br>*optional*|The last catchpoint seen by the node|string|
|**last-round**  <br>*required*|LastRound indicates the last round seen|integer|
|**last-version**  <br>*required*|LastVersion indicates the last consensus version supported|string|
|**next-version**  <br>*required*|NextVersion of consensus protocol to use|string|
|**next-version-round**  <br>*required*|NextVersionRound is the round at which the next consensus version will apply|integer|
|**next-version-supported**  <br>*required*|NextVersionSupported indicates whether the next consensus version is supported by this node|boolean|
|**stopped-at-unsupported-round**  <br>*required*|StoppedAtUnsupportedRound indicates that the node does not support the new rounds and has stopped making progress|boolean|
|**time-since-last-round**  <br>*required*|TimeSinceLastRound in nanoseconds|integer|
|**upgrade-delay**  <br>*optional*|Upgrade delay|integer|
|**upgrade-next-protocol-vote-before**  <br>*optional*|Next protocol round|integer|
|**upgrade-no-votes**  <br>*optional*|No votes cast for consensus upgrade|integer|
|**upgrade-node-vote**  <br>*optional*|This node's upgrade vote|boolean|
|**upgrade-vote-rounds**  <br>*optional*|Total voting rounds for current upgrade|integer|
|**upgrade-votes**  <br>*optional*|Total votes cast for consensus upgrade|integer|
|**upgrade-votes-required**  <br>*optional*|Yes votes required for consensus upgrade|integer|
|**upgrade-yes-votes**  <br>*optional*|Yes votes cast for consensus upgrade|integer|


**Produces**

* `application/json`


**Tags**

* nonparticipating
* public


<a name="waitforblock"></a>
### GET /v2/status/wait-for-block-after/{round}
Gets the node status after waiting for a round after the given round.
```
GET /v2/status/wait-for-block-after/{round}
```


**Description**
Waits for a block to appear after round {round} and returns the node's status at the time. There is a 1 minute timeout, when reached the current status is returned regardless of whether or not it is the round after the given round.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**round**  <br>*required*|The round to wait until returning status|integer|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**||[Response 200](#waitforblock-response-200)|
|**400**|Bad Request -- number must be non-negative integer|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**503**|Service Temporarily Unavailable|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|

<a name="waitforblock-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**catchpoint**  <br>*optional*|The current catchpoint that is being caught up to|string|
|**catchpoint-acquired-blocks**  <br>*optional*|The number of blocks that have already been obtained by the node as part of the catchup|integer|
|**catchpoint-processed-accounts**  <br>*optional*|The number of accounts from the current catchpoint that have been processed so far as part of the catchup|integer|
|**catchpoint-processed-kvs**  <br>*optional*|The number of key-values (KVs) from the current catchpoint that have been processed so far as part of the catchup|integer|
|**catchpoint-total-accounts**  <br>*optional*|The total number of accounts included in the current catchpoint|integer|
|**catchpoint-total-blocks**  <br>*optional*|The total number of blocks that are required to complete the current catchpoint catchup|integer|
|**catchpoint-total-kvs**  <br>*optional*|The total number of key-values (KVs) included in the current catchpoint|integer|
|**catchpoint-verified-accounts**  <br>*optional*|The number of accounts from the current catchpoint that have been verified so far as part of the catchup|integer|
|**catchpoint-verified-kvs**  <br>*optional*|The number of key-values (KVs) from the current catchpoint that have been verified so far as part of the catchup|integer|
|**catchup-time**  <br>*required*|CatchupTime in nanoseconds|integer|
|**last-catchpoint**  <br>*optional*|The last catchpoint seen by the node|string|
|**last-round**  <br>*required*|LastRound indicates the last round seen|integer|
|**last-version**  <br>*required*|LastVersion indicates the last consensus version supported|string|
|**next-version**  <br>*required*|NextVersion of consensus protocol to use|string|
|**next-version-round**  <br>*required*|NextVersionRound is the round at which the next consensus version will apply|integer|
|**next-version-supported**  <br>*required*|NextVersionSupported indicates whether the next consensus version is supported by this node|boolean|
|**stopped-at-unsupported-round**  <br>*required*|StoppedAtUnsupportedRound indicates that the node does not support the new rounds and has stopped making progress|boolean|
|**time-since-last-round**  <br>*required*|TimeSinceLastRound in nanoseconds|integer|
|**upgrade-delay**  <br>*optional*|Upgrade delay|integer|
|**upgrade-next-protocol-vote-before**  <br>*optional*|Next protocol round|integer|
|**upgrade-no-votes**  <br>*optional*|No votes cast for consensus upgrade|integer|
|**upgrade-node-vote**  <br>*optional*|This node's upgrade vote|boolean|
|**upgrade-vote-rounds**  <br>*optional*|Total voting rounds for current upgrade|integer|
|**upgrade-votes**  <br>*optional*|Total votes cast for consensus upgrade|integer|
|**upgrade-votes-required**  <br>*optional*|Yes votes required for consensus upgrade|integer|
|**upgrade-yes-votes**  <br>*optional*|Yes votes cast for consensus upgrade|integer|


**Produces**

* `application/json`


**Tags**

* nonparticipating
* public


<a name="tealcompile"></a>
### POST /v2/teal/compile
Compile TEAL source code to binary, produce its hash
```
POST /v2/teal/compile
```


**Description**
Given TEAL source code in plain text, return base64 encoded program bytes and base32 SHA512_256 hash of program bytes (Address style). This endpoint is only enabled when a node's configuration file sets EnableDeveloperAPI to true.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Query**|**sourcemap**  <br>*optional*|When set to `true`, returns the source map of the program as a JSON. Defaults to `false`.|boolean|
|**Body**|**source**  <br>*required*|TEAL source code to be compiled|string (binary)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Teal compile Result|[Response 200](#tealcompile-response-200)|
|**400**|Bad Request - Teal Compile Error|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**404**|Developer API not enabled|No Content|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|

<a name="tealcompile-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**hash**  <br>*required*|base32 SHA512_256 of program bytes (Address style)|string|
|**result**  <br>*required*|base64 encoded program bytes|string|
|**sourcemap**  <br>*optional*|JSON of the source map|object|


**Consumes**

* `text/plain`


**Produces**

* `application/json`


**Tags**

* nonparticipating
* public


<a name="tealdisassemble"></a>
### POST /v2/teal/disassemble
Disassemble program bytes into the TEAL source code.
```
POST /v2/teal/disassemble
```


**Description**
Given the program bytes, return the TEAL source code in plain text. This endpoint is only enabled when a node's configuration file sets EnableDeveloperAPI to true.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Body**|**source**  <br>*required*|TEAL program binary to be disassembled|string (byte)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Teal disassembly Result|[Response 200](#tealdisassemble-response-200)|
|**400**|Bad Request - Teal Compile Error|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**404**|Developer API not enabled|No Content|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|

<a name="tealdisassemble-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**result**  <br>*required*|disassembled Teal code|string|


**Consumes**

* `application/x-binary`


**Produces**

* `application/json`


**Tags**

* nonparticipating
* public


<a name="tealdryrun"></a>
### POST /v2/teal/dryrun
Provide debugging information for a transaction (or group).
```
POST /v2/teal/dryrun
```


**Description**
Executes TEAL program(s) in context and returns debugging information about the execution. This endpoint is only enabled when a node's configuration file sets EnableDeveloperAPI to true.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Body**|**request**  <br>*optional*|Transaction (or group) and any accompanying state-simulation data.|[DryrunRequest](#dryrunrequest)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|DryrunResponse contains per-txn debug information from a dryrun.|[Response 200](#tealdryrun-response-200)|
|**400**|Bad Request|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**404**|Developer API not enabled|No Content|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|

<a name="tealdryrun-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**error**  <br>*required*||string|
|**protocol-version**  <br>*required*|Protocol version is the protocol version Dryrun was operated under.|string|
|**txns**  <br>*required*||< [DryrunTxnResult](#dryruntxnresult) > array|


**Consumes**

* `application/json`
* `application/msgpack`


**Produces**

* `application/json`


**Tags**

* nonparticipating
* public


<a name="rawtransaction"></a>
### POST /v2/transactions
Broadcasts a raw transaction or transaction group to the network.
```
POST /v2/transactions
```


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Body**|**rawtxn**  <br>*required*|The byte encoded signed transaction to broadcast to network|string (binary)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Transaction ID of the submission.|[Response 200](#rawtransaction-response-200)|
|**400**|Bad Request - Malformed Algorand transaction|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**503**|Service Temporarily Unavailable|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|

<a name="rawtransaction-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**txId**  <br>*required*|encoding of the transaction hash.|string|


**Consumes**

* `application/x-binary`


**Produces**

* `application/json`


**Tags**

* participating
* public


<a name="rawtransactionasync"></a>
### POST /v2/transactions/async
Fast track for broadcasting a raw transaction or transaction group to the network through the tx handler without performing most of the checks and reporting detailed errors. Should be only used for development and performance testing.
```
POST /v2/transactions/async
```


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Body**|**rawtxn**  <br>*required*|The byte encoded signed transaction to broadcast to network|string (binary)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**||No Content|
|**400**|Bad Request - Malformed Algorand transaction|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**404**|Developer or Experimental API not enabled|No Content|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**503**|Service Temporarily Unavailable|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|


**Consumes**

* `application/x-binary`


**Tags**

* experimental
* public


<a name="transactionparams"></a>
### GET /v2/transactions/params
Get parameters for constructing a new transaction
```
GET /v2/transactions/params
```


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|TransactionParams contains the parameters that help a client construct a new transaction.|[Response 200](#transactionparams-response-200)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**503**|Service Temporarily Unavailable|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|

<a name="transactionparams-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**consensus-version**  <br>*required*|ConsensusVersion indicates the consensus protocol version<br>as of LastRound.|string|
|**fee**  <br>*required*|Fee is the suggested transaction fee<br>Fee is in units of micro-Algos per byte.<br>Fee may fall to zero but transactions must still have a fee of<br>at least MinTxnFee for the current network protocol.|integer|
|**genesis-hash**  <br>*required*|GenesisHash is the hash of the genesis block.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**genesis-id**  <br>*required*|GenesisID is an ID listed in the genesis block.|string|
|**last-round**  <br>*required*|LastRound indicates the last round seen|integer|
|**min-fee**  <br>*required*|The minimum transaction fee (not per byte) required for the<br>txn to validate for the current network protocol.|integer|


**Produces**

* `application/json`


**Tags**

* nonparticipating
* public


<a name="getpendingtransactions"></a>
### GET /v2/transactions/pending
Get a list of unconfirmed transactions currently in the transaction pool.
```
GET /v2/transactions/pending
```


**Description**
Get the list of pending transactions, sorted by priority, in decreasing order, truncated at the end at MAX. If MAX = 0, returns all pending transactions.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Query**|**format**  <br>*optional*|Configures whether the response object is JSON or MessagePack encoded. If not provided, defaults to JSON.|enum (json, msgpack)|
|**Query**|**max**  <br>*optional*|Truncated number of transactions to display. If max=0, returns all pending txns.|integer|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|A potentially truncated list of transactions currently in the node's transaction pool. You can compute whether or not the list is truncated if the number of elements in the **top-transactions** array is fewer than **total-transactions**.|[Response 200](#getpendingtransactions-response-200)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**503**|Service Temporarily Unavailable|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|

<a name="getpendingtransactions-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**top-transactions**  <br>*required*|An array of signed transaction objects.|< object > array|
|**total-transactions**  <br>*required*|Total number of transactions in the pool.|integer|


**Produces**

* `application/json`
* `application/msgpack`


**Tags**

* participating
* public


<a name="pendingtransactioninformation"></a>
### GET /v2/transactions/pending/{txid}
Get a specific pending transaction.
```
GET /v2/transactions/pending/{txid}
```


**Description**
Given a transaction ID of a recently submitted transaction, it returns information about it.  There are several cases when this might succeed:
- transaction committed (committed round > 0)
- transaction still in the pool (committed round = 0, pool error = "")
- transaction removed from pool due to error (committed round = 0, pool error != "")
Or the transaction may have happened sufficiently long ago that the node no longer remembers it, and this will return an error.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**txid**  <br>*required*|A transaction ID|string|
|**Query**|**format**  <br>*optional*|Configures whether the response object is JSON or MessagePack encoded. If not provided, defaults to JSON.|enum (json, msgpack)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Given a transaction ID of a recently submitted transaction, it returns information about it.  There are several cases when this might succeed:<br>- transaction committed (committed round > 0)<br>- transaction still in the pool (committed round = 0, pool error = "")<br>- transaction removed from pool due to error (committed round = 0, pool error != "")<br><br>Or the transaction may have happened sufficiently long ago that the node no longer remembers it, and this will return an error.|[PendingTransactionResponse](#pendingtransactionresponse)|
|**400**|Bad Request|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**404**|Transaction Not Found|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|


**Produces**

* `application/json`
* `application/msgpack`


**Tags**

* participating
* public


<a name="simulatetransaction"></a>
### POST /v2/transactions/simulate
Simulates a raw transaction or transaction group as it would be evaluated on the network. The simulation will use blockchain state from the latest committed round.
```
POST /v2/transactions/simulate
```


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Query**|**format**  <br>*optional*|Configures whether the response object is JSON or MessagePack encoded. If not provided, defaults to JSON.|enum (json, msgpack)|
|**Body**|**request**  <br>*required*|The transactions to simulate, along with any other inputs.|[SimulateRequest](#simulaterequest)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Result of a transaction group simulation.|[Response 200](#simulatetransaction-response-200)|
|**400**|Bad Request|[ErrorResponse](#errorresponse)|
|**401**|Invalid API Token|[ErrorResponse](#errorresponse)|
|**500**|Internal Error|[ErrorResponse](#errorresponse)|
|**503**|Service Temporarily Unavailable|[ErrorResponse](#errorresponse)|
|**default**|Unknown Error|No Content|

<a name="simulatetransaction-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**eval-overrides**  <br>*optional*||[SimulationEvalOverrides](#simulationevaloverrides)|
|**exec-trace-config**  <br>*optional*||[SimulateTraceConfig](#simulatetraceconfig)|
|**initial-states**  <br>*optional*||[SimulateInitialStates](#simulateinitialstates)|
|**last-round**  <br>*required*|The round immediately preceding this simulation. State changes through this round were used to run this simulation.|integer|
|**txn-groups**  <br>*required*|A result object for each transaction group that was simulated.|< [SimulateTransactionGroupResult](#simulatetransactiongroupresult) > array|
|**version**  <br>*required*|The version of this response object.|integer|


**Consumes**

* `application/json`
* `application/msgpack`


**Produces**

* `application/json`
* `application/msgpack`


**Tags**

* nonparticipating
* public


<a name="getversion"></a>
### GET /versions

**Description**
Retrieves the supported API versions, binary build versions, and genesis information.


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|VersionsResponse is the response to 'GET /versions'|[Version](#version)|


**Produces**

* `application/json`


**Tags**

* common
* public




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
|**apps-local-state**  <br>*optional*|\[appl\] applications local data stored in this account.<br><br>Note the raw object uses `map[int] -> AppLocalState` for this type.|< [ApplicationLocalState](#applicationlocalstate) > array|
|**apps-total-extra-pages**  <br>*optional*|\[teap\] the sum of all extra application program pages for this account.|integer|
|**apps-total-schema**  <br>*optional*|\[tsch\] stores the sum of all of the local schemas and global schemas in this account.<br><br>Note: the raw account uses `StateSchema` for this type.|[ApplicationStateSchema](#applicationstateschema)|
|**assets**  <br>*optional*|\[asset\] assets held by this account.<br><br>Note the raw object uses `map[int] -> AssetHolding` for this type.|< [AssetHolding](#assetholding) > array|
|**auth-addr**  <br>*optional*|\[spend\] the address against which signing should be checked. If empty, the address of the current account is used. This field can be updated in any transaction by setting the RekeyTo field.|string|
|**created-apps**  <br>*optional*|\[appp\] parameters of applications created by this account including app global data.<br><br>Note: the raw account uses `map[int] -> AppParams` for this type.|< [Application](#application) > array|
|**created-assets**  <br>*optional*|\[apar\] parameters of assets created by this account.<br><br>Note: the raw account uses `map[int] -> Asset` for this type.|< [Asset](#asset) > array|
|**incentive-eligible**  <br>*optional*|Whether or not the account can receive block incentives if its balance is in range at proposal time.|boolean|
|**last-heartbeat**  <br>*optional*|The round in which this account last went online, or explicitly renewed their online status.|integer|
|**last-proposed**  <br>*optional*|The round in which this account last proposed the block.|integer|
|**min-balance**  <br>*required*|MicroAlgo balance required by the account.<br><br>The requirement grows based on asset and application usage.|integer|
|**participation**  <br>*optional*||[AccountParticipation](#accountparticipation)|
|**pending-rewards**  <br>*required*|amount of MicroAlgos of pending rewards in this account.|integer|
|**reward-base**  <br>*optional*|\[ebase\] used as part of the rewards computation. Only applicable to accounts which are participating.|integer|
|**rewards**  <br>*required*|\[ern\] total rewards of MicroAlgos the account has received, including pending rewards.|integer|
|**round**  <br>*required*|The round for which this information is relevant.|integer|
|**sig-type**  <br>*optional*|Indicates what type of signature is used by this account, must be one of:<br>* sig<br>* msig<br>* lsig|enum (sig, msig, lsig)|
|**status**  <br>*required*|\[onl\] delegation status of the account's MicroAlgos<br>* Offline - indicates that the associated account is delegated.<br>*  Online  - indicates that the associated account used as part of the delegation pool.<br>*   NotParticipating - indicates that the associated account is neither a delegator nor a delegate.|string|
|**total-apps-opted-in**  <br>*required*|The count of all applications that have been opted in, equivalent to the count of application local data (AppLocalState objects) stored in this account.|integer|
|**total-assets-opted-in**  <br>*required*|The count of all assets that have been opted in, equivalent to the count of AssetHolding objects held by this account.|integer|
|**total-box-bytes**  <br>*optional*|\[tbxb\] The total number of bytes used by this account's app's box keys and values.|integer|
|**total-boxes**  <br>*optional*|\[tbx\] The number of existing boxes created by this account's app.|integer|
|**total-created-apps**  <br>*required*|The count of all apps (AppParams objects) created by this account.|integer|
|**total-created-assets**  <br>*required*|The count of all assets (AssetParams objects) created by this account.|integer|


<a name="accountassetholding"></a>
### AccountAssetHolding
AccountAssetHolding describes the account's asset holding and asset parameters (if either exist) for a specific asset ID.


|Name|Description|Schema|
|---|---|---|
|**asset-holding**  <br>*required*|\[asset\] Details about the asset held by this account.<br><br>The raw account uses `AssetHolding` for this type.|[AssetHolding](#assetholding)|
|**asset-params**  <br>*optional*|\[apar\] parameters of the asset held by this account.<br><br>The raw account uses `AssetParams` for this type.|[AssetParams](#assetparams)|


<a name="accountparticipation"></a>
### AccountParticipation
AccountParticipation describes the parameters used by this account in consensus protocol.


|Name|Description|Schema|
|---|---|---|
|**selection-participation-key**  <br>*required*|\[sel\] Selection public key (if any) currently registered for this round.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**state-proof-key**  <br>*optional*|\[stprf\] Root of the state proof key (if any)  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**vote-first-valid**  <br>*required*|\[voteFst\] First round for which this participation is valid.|integer|
|**vote-key-dilution**  <br>*required*|\[voteKD\] Number of subkeys in each batch of participation keys.|integer|
|**vote-last-valid**  <br>*required*|\[voteLst\] Last round for which this participation is valid.|integer|
|**vote-participation-key**  <br>*required*|\[vote\] root participation public key (if any) currently registered for this round.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|


<a name="accountstatedelta"></a>
### AccountStateDelta
Application state delta.


|Name|Schema|
|---|---|
|**address**  <br>*required*|string|
|**delta**  <br>*required*|[StateDelta](#statedelta)|


<a name="appcalllogs"></a>
### AppCallLogs
The logged messages from an app call along with the app ID and outer transaction ID. Logs appear in the same order that they were emitted.


|Name|Description|Schema|
|---|---|---|
|**application-index**  <br>*required*|The application from which the logs were generated|integer|
|**logs**  <br>*required*|An array of logs|< string (byte) > array|
|**txId**  <br>*required*|The transaction ID of the outer app call that lead to these logs|string|


<a name="application"></a>
### Application
Application index and its parameters


|Name|Description|Schema|
|---|---|---|
|**id**  <br>*required*|\[appidx\] application index.|integer|
|**params**  <br>*required*|\[appparams\] application parameters.|[ApplicationParams](#applicationparams)|


<a name="applicationinitialstates"></a>
### ApplicationInitialStates
An application's initial global/local/box states that were accessed during simulation.


|Name|Description|Schema|
|---|---|---|
|**app-boxes**  <br>*optional*||[ApplicationKVStorage](#applicationkvstorage)|
|**app-globals**  <br>*optional*||[ApplicationKVStorage](#applicationkvstorage)|
|**app-locals**  <br>*optional*|An application's initial local states tied to different accounts.|< [ApplicationKVStorage](#applicationkvstorage) > array|
|**id**  <br>*required*|Application index.|integer|


<a name="applicationkvstorage"></a>
### ApplicationKVStorage
An application's global/local/box state.


|Name|Description|Schema|
|---|---|---|
|**account**  <br>*optional*|The address of the account associated with the local state.|string|
|**kvs**  <br>*required*|Key-Value pairs representing application states.|< [AvmKeyValue](#avmkeyvalue) > array|


<a name="applicationlocalreference"></a>
### ApplicationLocalReference
References an account's local state for an application.


|Name|Description|Schema|
|---|---|---|
|**account**  <br>*required*|Address of the account with the local state.|string|
|**app**  <br>*required*|Application ID of the local state application.|integer|


<a name="applicationlocalstate"></a>
### ApplicationLocalState
Stores local state associated with an application.


|Name|Description|Schema|
|---|---|---|
|**id**  <br>*required*|The application which this local state is for.|integer|
|**key-value**  <br>*optional*|\[tkv\] storage.|[TealKeyValueStore](#tealkeyvaluestore)|
|**schema**  <br>*required*|\[hsch\] schema.|[ApplicationStateSchema](#applicationstateschema)|


<a name="applicationparams"></a>
### ApplicationParams
Stores the global information associated with an application.


|Name|Description|Schema|
|---|---|---|
|**approval-program**  <br>*required*|\[approv\] approval program.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**clear-state-program**  <br>*required*|\[clearp\] approval program.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**creator**  <br>*required*|The address that created this application. This is the address where the parameters and global state for this application can be found.|string|
|**extra-program-pages**  <br>*optional*|\[epp\] the amount of extra program pages available to this app.|integer|
|**global-state**  <br>*optional*|\[gs\] global state|[TealKeyValueStore](#tealkeyvaluestore)|
|**global-state-schema**  <br>*optional*|\[gsch\] global schema|[ApplicationStateSchema](#applicationstateschema)|
|**local-state-schema**  <br>*optional*|\[lsch\] local schema|[ApplicationStateSchema](#applicationstateschema)|


<a name="applicationstateoperation"></a>
### ApplicationStateOperation
An operation against an application's global/local/box state.


|Name|Description|Schema|
|---|---|---|
|**account**  <br>*optional*|For local state changes, the address of the account associated with the local state.|string|
|**app-state-type**  <br>*required*|Type of application state. Value `g` is **global state**, `l` is **local state**, `b` is **boxes**.|string|
|**key**  <br>*required*|The key (name) of the global/local/box state.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**new-value**  <br>*optional*||[AvmValue](#avmvalue)|
|**operation**  <br>*required*|Operation type. Value `w` is **write**, `d` is **delete**.|string|


<a name="applicationstateschema"></a>
### ApplicationStateSchema
Specifies maximums on the number of each type that may be stored.


|Name|Description|Schema|
|---|---|---|
|**num-byte-slice**  <br>*required*|\[nbs\] num of byte slices.|integer|
|**num-uint**  <br>*required*|\[nui\] num of uints.|integer|


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
|**is-frozen**  <br>*required*|\[f\] whether or not the holding is frozen.|boolean|


<a name="assetholdingreference"></a>
### AssetHoldingReference
References an asset held by an account.


|Name|Description|Schema|
|---|---|---|
|**account**  <br>*required*|Address of the account holding the asset.|string|
|**asset**  <br>*required*|Asset ID of the holding.|integer|


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
|**name**  <br>*optional*|\[an\] Name of this asset, as supplied by the creator. Included only when the asset name is composed of printable utf-8 characters.|string|
|**name-b64**  <br>*optional*|Base64 encoded name of this asset, as supplied by the creator.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**reserve**  <br>*optional*|\[r\] Address of account holding reserve (non-minted) units of this asset.|string|
|**total**  <br>*required*|\[t\] The total number of units of this asset.|integer|
|**unit-name**  <br>*optional*|\[un\] Name of a unit of this asset, as supplied by the creator. Included only when the name of a unit of this asset is composed of printable utf-8 characters.|string|
|**unit-name-b64**  <br>*optional*|Base64 encoded name of a unit of this asset, as supplied by the creator.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**url**  <br>*optional*|\[au\] URL where more information about the asset can be retrieved. Included only when the URL is composed of printable utf-8 characters.|string|
|**url-b64**  <br>*optional*|Base64 encoded URL where more information about the asset can be retrieved.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|


<a name="avmkeyvalue"></a>
### AvmKeyValue
Represents an AVM key-value pair in an application store.


|Name|Description|Schema|
|---|---|---|
|**key**  <br>*required*|**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**value**  <br>*required*||[AvmValue](#avmvalue)|


<a name="avmvalue"></a>
### AvmValue
Represents an AVM value.


|Name|Description|Schema|
|---|---|---|
|**bytes**  <br>*optional*|bytes value.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**type**  <br>*required*|value type. Value `1` refers to **bytes**, value `2` refers to **uint64**|integer|
|**uint**  <br>*optional*|uint value.|integer|


<a name="box"></a>
### Box
Box name and its content.


|Name|Description|Schema|
|---|---|---|
|**name**  <br>*required*|\[name\] box name, base64 encoded  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**round**  <br>*required*|The round for which this information is relevant|integer|
|**value**  <br>*required*|\[value\] box value, base64 encoded.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|


<a name="boxdescriptor"></a>
### BoxDescriptor
Box descriptor describes a Box.


|Name|Description|Schema|
|---|---|---|
|**name**  <br>*required*|Base64 encoded box name  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|


<a name="boxreference"></a>
### BoxReference
References a box of an application.


|Name|Description|Schema|
|---|---|---|
|**app**  <br>*required*|Application ID which this box belongs to|integer|
|**name**  <br>*required*|Base64 encoded box name  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|


<a name="buildversion"></a>
### BuildVersion

|Name|Schema|
|---|---|
|**branch**  <br>*required*|string|
|**build_number**  <br>*required*|integer (int64)|
|**channel**  <br>*required*|string|
|**commit_hash**  <br>*required*|string|
|**major**  <br>*required*|integer (int64)|
|**minor**  <br>*required*|integer (int64)|


<a name="debugsettingsprof"></a>
### DebugSettingsProf
algod mutex and blocking profiling state.


|Name|Description|Schema|
|---|---|---|
|**block-rate**  <br>*optional*|The rate of blocking events. The profiler aims to sample an average of one blocking event per rate nanoseconds spent blocked. To turn off profiling entirely, pass rate 0.  <br>**Example** : `1000`|integer|
|**mutex-rate**  <br>*optional*|The rate of mutex events. On average 1/rate events are reported. To turn off profiling entirely, pass rate 0  <br>**Example** : `1000`|integer|


<a name="dryrunrequest"></a>
### DryrunRequest
Request data type for dryrun endpoint. Given the Transactions and simulated ledger state upload, run TEAL scripts and return debugging information.


|Name|Description|Schema|
|---|---|---|
|**accounts**  <br>*required*||< [Account](#account) > array|
|**apps**  <br>*required*||< [Application](#application) > array|
|**latest-timestamp**  <br>*required*|LatestTimestamp is available to some TEAL scripts. Defaults to the latest confirmed timestamp this algod is attached to.|integer (int64)|
|**protocol-version**  <br>*required*|ProtocolVersion specifies a specific version string to operate under, otherwise whatever the current protocol of the network this algod is running in.|string|
|**round**  <br>*required*|Round is available to some TEAL scripts. Defaults to the current round on the network this algod is attached to.|integer|
|**sources**  <br>*required*||< [DryrunSource](#dryrunsource) > array|
|**txns**  <br>*required*||< string (json) > array|


<a name="dryrunsource"></a>
### DryrunSource
DryrunSource is TEAL source text that gets uploaded, compiled, and inserted into transactions or application state.


|Name|Description|Schema|
|---|---|---|
|**app-index**  <br>*required*||integer|
|**field-name**  <br>*required*|FieldName is what kind of sources this is. If lsig then it goes into the transactions[this.TxnIndex].LogicSig. If approv or clearp it goes into the Approval Program or Clear State Program of application[this.AppIndex].|string|
|**source**  <br>*required*||string|
|**txn-index**  <br>*required*||integer|


<a name="dryrunstate"></a>
### DryrunState
Stores the TEAL eval step data


|Name|Description|Schema|
|---|---|---|
|**error**  <br>*optional*|Evaluation error if any|string|
|**line**  <br>*required*|Line number|integer|
|**pc**  <br>*required*|Program counter|integer|
|**scratch**  <br>*optional*||< [TealValue](#tealvalue) > array|
|**stack**  <br>*required*||< [TealValue](#tealvalue) > array|


<a name="dryruntxnresult"></a>
### DryrunTxnResult
DryrunTxnResult contains any LogicSig or ApplicationCall program debug information and state updates from a dryrun.


|Name|Description|Schema|
|---|---|---|
|**app-call-messages**  <br>*optional*||< string > array|
|**app-call-trace**  <br>*optional*||< [DryrunState](#dryrunstate) > array|
|**budget-added**  <br>*optional*|Budget added during execution of app call transaction.|integer|
|**budget-consumed**  <br>*optional*|Budget consumed during execution of app call transaction.|integer|
|**disassembly**  <br>*required*|Disassembled program line by line.|< string > array|
|**global-delta**  <br>*optional*||[StateDelta](#statedelta)|
|**local-deltas**  <br>*optional*||< [AccountStateDelta](#accountstatedelta) > array|
|**logic-sig-disassembly**  <br>*optional*|Disassembled lsig program line by line.|< string > array|
|**logic-sig-messages**  <br>*optional*||< string > array|
|**logic-sig-trace**  <br>*optional*||< [DryrunState](#dryrunstate) > array|
|**logs**  <br>*optional*||< string (byte) > array|


<a name="errorresponse"></a>
### ErrorResponse
An error response with optional data field.


|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|


<a name="evaldelta"></a>
### EvalDelta
Represents a TEAL value delta.


|Name|Description|Schema|
|---|---|---|
|**action**  <br>*required*|\[at\] delta action.|integer|
|**bytes**  <br>*optional*|\[bs\] bytes value.|string|
|**uint**  <br>*optional*|\[ui\] uint value.|integer|


<a name="evaldeltakeyvalue"></a>
### EvalDeltaKeyValue
Key-value pairs for StateDelta.


|Name|Schema|
|---|---|
|**key**  <br>*required*|string|
|**value**  <br>*required*|[EvalDelta](#evaldelta)|


<a name="kvdelta"></a>
### KvDelta
A single Delta containing the key, the previous value and the current value for a single round.


|Name|Description|Schema|
|---|---|---|
|**key**  <br>*optional*|The key, base64 encoded.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**value**  <br>*optional*|The new value of the KV store entry, base64 encoded.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|


<a name="ledgerstatedelta"></a>
### LedgerStateDelta
Ledger StateDelta object

*Type* : object


<a name="ledgerstatedeltafortransactiongroup"></a>
### LedgerStateDeltaForTransactionGroup
Contains a ledger delta for a single transaction group


|Name|Schema|
|---|---|
|**Delta**  <br>*required*|[LedgerStateDelta](#ledgerstatedelta)|
|**Ids**  <br>*required*|< string > array|


<a name="lightblockheaderproof"></a>
### LightBlockHeaderProof
Proof of membership and position of a light block header.


|Name|Description|Schema|
|---|---|---|
|**index**  <br>*required*|The index of the light block header in the vector commitment tree|integer|
|**proof**  <br>*required*|The encoded proof.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**treedepth**  <br>*required*|Represents the depth of the tree that is being proven, i.e. the number of edges from a leaf to the root.|integer|


<a name="participationkey"></a>
### ParticipationKey
Represents a participation key used by the node.


|Name|Description|Schema|
|---|---|---|
|**address**  <br>*required*|Address the key was generated for.|string|
|**effective-first-valid**  <br>*optional*|When registered, this is the first round it may be used.|integer|
|**effective-last-valid**  <br>*optional*|When registered, this is the last round it may be used.|integer|
|**id**  <br>*required*|The key's ParticipationID.|string|
|**key**  <br>*required*|Key information stored on the account.|[AccountParticipation](#accountparticipation)|
|**last-block-proposal**  <br>*optional*|Round when this key was last used to propose a block.|integer|
|**last-state-proof**  <br>*optional*|Round when this key was last used to generate a state proof.|integer|
|**last-vote**  <br>*optional*|Round when this key was last used to vote.|integer|


<a name="pendingtransactionresponse"></a>
### PendingTransactionResponse
Details about a pending transaction. If the transaction was recently confirmed, includes confirmation details like the round and reward details.


|Name|Description|Schema|
|---|---|---|
|**application-index**  <br>*optional*|The application index if the transaction was found and it created an application.|integer|
|**asset-closing-amount**  <br>*optional*|The number of the asset's unit that were transferred to the close-to address.|integer|
|**asset-index**  <br>*optional*|The asset index if the transaction was found and it created an asset.|integer|
|**close-rewards**  <br>*optional*|Rewards in microalgos applied to the close remainder to account.|integer|
|**closing-amount**  <br>*optional*|Closing amount for the transaction.|integer|
|**confirmed-round**  <br>*optional*|The round where this transaction was confirmed, if present.|integer|
|**global-state-delta**  <br>*optional*|Global state key/value changes for the application being executed by this transaction.|[StateDelta](#statedelta)|
|**inner-txns**  <br>*optional*|Inner transactions produced by application execution.|< [PendingTransactionResponse](#pendingtransactionresponse) > array|
|**local-state-delta**  <br>*optional*|Local state key/value changes for the application being executed by this transaction.|< [AccountStateDelta](#accountstatedelta) > array|
|**logs**  <br>*optional*|Logs for the application being executed by this transaction.|< string (byte) > array|
|**pool-error**  <br>*required*|Indicates that the transaction was kicked out of this node's transaction pool (and specifies why that happened).  An empty string indicates the transaction wasn't kicked out of this node's txpool due to an error.|string|
|**receiver-rewards**  <br>*optional*|Rewards in microalgos applied to the receiver account.|integer|
|**sender-rewards**  <br>*optional*|Rewards in microalgos applied to the sender account.|integer|
|**txn**  <br>*required*|The raw signed transaction.|object|


<a name="scratchchange"></a>
### ScratchChange
A write operation into a scratch slot.


|Name|Description|Schema|
|---|---|---|
|**new-value**  <br>*required*||[AvmValue](#avmvalue)|
|**slot**  <br>*required*|The scratch slot written.|integer|


<a name="simulateinitialstates"></a>
### SimulateInitialStates
Initial states of resources that were accessed during simulation.


|Name|Description|Schema|
|---|---|---|
|**app-initial-states**  <br>*optional*|The initial states of accessed application before simulation. The order of this array is arbitrary.|< [ApplicationInitialStates](#applicationinitialstates) > array|


<a name="simulaterequest"></a>
### SimulateRequest
Request type for simulation endpoint.


|Name|Description|Schema|
|---|---|---|
|**allow-empty-signatures**  <br>*optional*|Allows transactions without signatures to be simulated as if they had correct signatures.|boolean|
|**allow-more-logging**  <br>*optional*|Lifts limits on log opcode usage during simulation.|boolean|
|**allow-unnamed-resources**  <br>*optional*|Allows access to unnamed resources during simulation.|boolean|
|**exec-trace-config**  <br>*optional*||[SimulateTraceConfig](#simulatetraceconfig)|
|**extra-opcode-budget**  <br>*optional*|Applies extra opcode budget during simulation for each transaction group.|integer|
|**fix-signers**  <br>*optional*|If true, signers for transactions that are missing signatures will be fixed during evaluation.|boolean|
|**round**  <br>*optional*|If provided, specifies the round preceding the simulation. State changes through this round will be used to run this simulation. Usually only the 4 most recent rounds will be available (controlled by the node config value MaxAcctLookback). If not specified, defaults to the latest available round.|integer|
|**txn-groups**  <br>*required*|The transaction groups to simulate.|< [SimulateRequestTransactionGroup](#simulaterequesttransactiongroup) > array|


<a name="simulaterequesttransactiongroup"></a>
### SimulateRequestTransactionGroup
A transaction group to simulate.


|Name|Description|Schema|
|---|---|---|
|**txns**  <br>*required*|An atomic transaction group.|< string (json) > array|


<a name="simulatetraceconfig"></a>
### SimulateTraceConfig
An object that configures simulation execution trace.


|Name|Description|Schema|
|---|---|---|
|**enable**  <br>*optional*|A boolean option for opting in execution trace features simulation endpoint.|boolean|
|**scratch-change**  <br>*optional*|A boolean option enabling returning scratch slot changes together with execution trace during simulation.|boolean|
|**stack-change**  <br>*optional*|A boolean option enabling returning stack changes together with execution trace during simulation.|boolean|
|**state-change**  <br>*optional*|A boolean option enabling returning application state changes (global, local, and box changes) with the execution trace during simulation.|boolean|


<a name="simulatetransactiongroupresult"></a>
### SimulateTransactionGroupResult
Simulation result for an atomic transaction group


|Name|Description|Schema|
|---|---|---|
|**app-budget-added**  <br>*optional*|Total budget added during execution of app calls in the transaction group.|integer|
|**app-budget-consumed**  <br>*optional*|Total budget consumed during execution of app calls in the transaction group.|integer|
|**failed-at**  <br>*optional*|If present, indicates which transaction in this group caused the failure. This array represents the path to the failing transaction. Indexes are zero based, the first element indicates the top-level transaction, and successive elements indicate deeper inner transactions.|< integer > array|
|**failure-message**  <br>*optional*|If present, indicates that the transaction group failed and specifies why that happened|string|
|**txn-results**  <br>*required*|Simulation result for individual transactions|< [SimulateTransactionResult](#simulatetransactionresult) > array|
|**unnamed-resources-accessed**  <br>*optional*||[SimulateUnnamedResourcesAccessed](#simulateunnamedresourcesaccessed)|


<a name="simulatetransactionresult"></a>
### SimulateTransactionResult
Simulation result for an individual transaction


|Name|Description|Schema|
|---|---|---|
|**app-budget-consumed**  <br>*optional*|Budget used during execution of an app call transaction. This value includes budged used by inner app calls spawned by this transaction.|integer|
|**exec-trace**  <br>*optional*||[SimulationTransactionExecTrace](#simulationtransactionexectrace)|
|**fixed-signer**  <br>*optional*|The account that needed to sign this transaction when no signature was provided and the provided signer was incorrect.|string|
|**logic-sig-budget-consumed**  <br>*optional*|Budget used during execution of a logic sig transaction.|integer|
|**txn-result**  <br>*required*||[PendingTransactionResponse](#pendingtransactionresponse)|
|**unnamed-resources-accessed**  <br>*optional*||[SimulateUnnamedResourcesAccessed](#simulateunnamedresourcesaccessed)|


<a name="simulateunnamedresourcesaccessed"></a>
### SimulateUnnamedResourcesAccessed
These are resources that were accessed by this group that would normally have caused failure, but were allowed in simulation. Depending on where this object is in the response, the unnamed resources it contains may or may not qualify for group resource sharing. If this is a field in SimulateTransactionGroupResult, the resources do qualify, but if this is a field in SimulateTransactionResult, they do not qualify. In order to make this group valid for actual submission, resources that qualify for group sharing can be made available by any transaction of the group; otherwise, resources must be placed in the same transaction which accessed them.


|Name|Description|Schema|
|---|---|---|
|**accounts**  <br>*optional*|The unnamed accounts that were referenced. The order of this array is arbitrary.|< string > array|
|**app-locals**  <br>*optional*|The unnamed application local states that were referenced. The order of this array is arbitrary.|< [ApplicationLocalReference](#applicationlocalreference) > array|
|**apps**  <br>*optional*|The unnamed applications that were referenced. The order of this array is arbitrary.|< integer > array|
|**asset-holdings**  <br>*optional*|The unnamed asset holdings that were referenced. The order of this array is arbitrary.|< [AssetHoldingReference](#assetholdingreference) > array|
|**assets**  <br>*optional*|The unnamed assets that were referenced. The order of this array is arbitrary.|< integer > array|
|**boxes**  <br>*optional*|The unnamed boxes that were referenced. The order of this array is arbitrary.|< [BoxReference](#boxreference) > array|
|**extra-box-refs**  <br>*optional*|The number of extra box references used to increase the IO budget. This is in addition to the references defined in the input transaction group and any referenced to unnamed boxes.|integer|


<a name="simulationevaloverrides"></a>
### SimulationEvalOverrides
The set of parameters and limits override during simulation. If this set of parameters is present, then evaluation parameters may differ from standard evaluation in certain ways.


|Name|Description|Schema|
|---|---|---|
|**allow-empty-signatures**  <br>*optional*|If true, transactions without signatures are allowed and simulated as if they were properly signed.|boolean|
|**allow-unnamed-resources**  <br>*optional*|If true, allows access to unnamed resources during simulation.|boolean|
|**extra-opcode-budget**  <br>*optional*|The extra opcode budget added to each transaction group during simulation|integer|
|**fix-signers**  <br>*optional*|If true, signers for transactions that are missing signatures will be fixed during evaluation.|boolean|
|**max-log-calls**  <br>*optional*|The maximum log calls one can make during simulation|integer|
|**max-log-size**  <br>*optional*|The maximum byte number to log during simulation|integer|


<a name="simulationopcodetraceunit"></a>
### SimulationOpcodeTraceUnit
The set of trace information and effect from evaluating a single opcode.


|Name|Description|Schema|
|---|---|---|
|**pc**  <br>*required*|The program counter of the current opcode being evaluated.|integer|
|**scratch-changes**  <br>*optional*|The writes into scratch slots.|< [ScratchChange](#scratchchange) > array|
|**spawned-inners**  <br>*optional*|The indexes of the traces for inner transactions spawned by this opcode, if any.|< integer > array|
|**stack-additions**  <br>*optional*|The values added by this opcode to the stack.|< [AvmValue](#avmvalue) > array|
|**stack-pop-count**  <br>*optional*|The number of deleted stack values by this opcode.|integer|
|**state-changes**  <br>*optional*|The operations against the current application's states.|< [ApplicationStateOperation](#applicationstateoperation) > array|


<a name="simulationtransactionexectrace"></a>
### SimulationTransactionExecTrace
The execution trace of calling an app or a logic sig, containing the inner app call trace in a recursive way.


|Name|Description|Schema|
|---|---|---|
|**approval-program-hash**  <br>*optional*|SHA512_256 hash digest of the approval program executed in transaction.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**approval-program-trace**  <br>*optional*|Program trace that contains a trace of opcode effects in an approval program.|< [SimulationOpcodeTraceUnit](#simulationopcodetraceunit) > array|
|**clear-state-program-hash**  <br>*optional*|SHA512_256 hash digest of the clear state program executed in transaction.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**clear-state-program-trace**  <br>*optional*|Program trace that contains a trace of opcode effects in a clear state program.|< [SimulationOpcodeTraceUnit](#simulationopcodetraceunit) > array|
|**clear-state-rollback**  <br>*optional*|If true, indicates that the clear state program failed and any persistent state changes it produced should be reverted once the program exits.|boolean|
|**clear-state-rollback-error**  <br>*optional*|The error message explaining why the clear state program failed. This field will only be populated if clear-state-rollback is true and the failure was due to an execution error.|string|
|**inner-trace**  <br>*optional*|An array of SimulationTransactionExecTrace representing the execution trace of any inner transactions executed.|< [SimulationTransactionExecTrace](#simulationtransactionexectrace) > array|
|**logic-sig-hash**  <br>*optional*|SHA512_256 hash digest of the logic sig executed in transaction.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**logic-sig-trace**  <br>*optional*|Program trace that contains a trace of opcode effects in a logic sig.|< [SimulationOpcodeTraceUnit](#simulationopcodetraceunit) > array|


<a name="statedelta"></a>
### StateDelta
Application state delta.

*Type* : < [EvalDeltaKeyValue](#evaldeltakeyvalue) > array


<a name="stateproof"></a>
### StateProof
Represents a state proof and its corresponding message


|Name|Description|Schema|
|---|---|---|
|**Message**  <br>*required*||[StateProofMessage](#stateproofmessage)|
|**StateProof**  <br>*required*|The encoded StateProof for the message.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|


<a name="stateproofmessage"></a>
### StateProofMessage
Represents the message that the state proofs are attesting to.


|Name|Description|Schema|
|---|---|---|
|**BlockHeadersCommitment**  <br>*required*|The vector commitment root on all light block headers within a state proof interval.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**FirstAttestedRound**  <br>*required*|The first round the message attests to.|integer|
|**LastAttestedRound**  <br>*required*|The last round the message attests to.|integer|
|**LnProvenWeight**  <br>*required*|An integer value representing the natural log of the proven weight with 16 bits of precision. This value would be used to verify the next state proof.|integer|
|**VotersCommitment**  <br>*required*|The vector commitment root of the top N accounts to sign the next StateProof.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|


<a name="tealkeyvalue"></a>
### TealKeyValue
Represents a key-value pair in an application store.


|Name|Schema|
|---|---|
|**key**  <br>*required*|string|
|**value**  <br>*required*|[TealValue](#tealvalue)|


<a name="tealkeyvaluestore"></a>
### TealKeyValueStore
Represents a key-value store for use in an application.

*Type* : < [TealKeyValue](#tealkeyvalue) > array


<a name="tealvalue"></a>
### TealValue
Represents a TEAL value.


|Name|Description|Schema|
|---|---|---|
|**bytes**  <br>*required*|\[tb\] bytes value.|string|
|**type**  <br>*required*|\[tt\] value type. Value `1` refers to **bytes**, value `2` refers to **uint**|integer|
|**uint**  <br>*required*|\[ui\] uint value.|integer|


<a name="version"></a>
### Version
algod version information.


|Name|Description|Schema|
|---|---|---|
|**build**  <br>*required*||[BuildVersion](#buildversion)|
|**genesis_hash_b64**  <br>*required*|**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**genesis_id**  <br>*required*||string|
|**versions**  <br>*required*||< string > array|



