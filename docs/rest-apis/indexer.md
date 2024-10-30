title: indexer
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
|**Query**|**application-id**  <br>*optional*|Application ID|integer|
|**Query**|**asset-id**  <br>*optional*|Asset ID|integer|
|**Query**|**auth-addr**  <br>*optional*|Include accounts configured to use this spending key.|string|
|**Query**|**currency-greater-than**  <br>*optional*|Results should have an amount greater than this value. MicroAlgos are the default currency unless an asset-id is provided, in which case the asset will be used.|integer|
|**Query**|**currency-less-than**  <br>*optional*|Results should have an amount less than this value. MicroAlgos are the default currency unless an asset-id is provided, in which case the asset will be used.|integer|
|**Query**|**exclude**  <br>*optional*|Exclude additional items such as asset holdings, application local data stored for this account, asset parameters created by this account, and application parameters created by this account.|< enum (all, assets, created-assets, apps-local-state, created-apps, none) > array|
|**Query**|**include-all**  <br>*optional*|Include all items including closed accounts, deleted applications, destroyed assets, opted-out asset holdings, and closed-out application localstates.|boolean|
|**Query**|**limit**  <br>*optional*|Maximum number of results to return. There could be additional pages even if the limit is not reached.|integer|
|**Query**|**next**  <br>*optional*|The next page of results. Use the next token provided by the previous results.|string|
|**Query**|**round**  <br>*optional*|Deprecated and disallowed. This parameter used to include results for a specified round. Requests with this parameter set are now rejected.|integer|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|(empty)|[Response 200](#searchforaccounts-response-200)|
|**400**|Response for errors|[Response 400](#searchforaccounts-response-400)|
|**500**|Response for errors|[Response 500](#searchforaccounts-response-500)|

<a name="searchforaccounts-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**accounts**  <br>*required*||< [Account](#account) > array|
|**current-round**  <br>*required*|Round at which the results were computed.|integer|
|**next-token**  <br>*optional*|Used for pagination, when making another request provide this token with the next parameter.|string|

<a name="searchforaccounts-response-400"></a>
**Response 400**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|

<a name="searchforaccounts-response-500"></a>
**Response 500**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|


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
|**Query**|**exclude**  <br>*optional*|Exclude additional items such as asset holdings, application local data stored for this account, asset parameters created by this account, and application parameters created by this account.|< enum (all, assets, created-assets, apps-local-state, created-apps, none) > array|
|**Query**|**include-all**  <br>*optional*|Include all items including closed accounts, deleted applications, destroyed assets, opted-out asset holdings, and closed-out application localstates.|boolean|
|**Query**|**round**  <br>*optional*|Deprecated and disallowed. This parameter used to include results for a specified round. Requests with this parameter set are now rejected.|integer|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|(empty)|[Response 200](#lookupaccountbyid-response-200)|
|**400**|Response for errors|[Response 400](#lookupaccountbyid-response-400)|
|**404**|Response for errors|[Response 404](#lookupaccountbyid-response-404)|
|**500**|Response for errors|[Response 500](#lookupaccountbyid-response-500)|

<a name="lookupaccountbyid-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**account**  <br>*required*||[Account](#account)|
|**current-round**  <br>*required*|Round at which the results were computed.|integer|

<a name="lookupaccountbyid-response-400"></a>
**Response 400**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|

<a name="lookupaccountbyid-response-404"></a>
**Response 404**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|

<a name="lookupaccountbyid-response-500"></a>
**Response 500**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|


**Consumes**

* `application/json`


**Produces**

* `application/json`


**Tags**

* lookup


<a name="lookupaccountapplocalstates"></a>
### GET /v2/accounts/{account-id}/apps-local-state

**Description**
Lookup an account's asset holdings, optionally for a specific ID.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**account-id**  <br>*required*|account string|string|
|**Query**|**application-id**  <br>*optional*|Application ID|integer|
|**Query**|**include-all**  <br>*optional*|Include all items including closed accounts, deleted applications, destroyed assets, opted-out asset holdings, and closed-out application localstates.|boolean|
|**Query**|**limit**  <br>*optional*|Maximum number of results to return. There could be additional pages even if the limit is not reached.|integer|
|**Query**|**next**  <br>*optional*|The next page of results. Use the next token provided by the previous results.|string|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|(empty)|[Response 200](#lookupaccountapplocalstates-response-200)|
|**400**|Response for errors|[Response 400](#lookupaccountapplocalstates-response-400)|
|**404**|Response for errors|[Response 404](#lookupaccountapplocalstates-response-404)|
|**500**|Response for errors|[Response 500](#lookupaccountapplocalstates-response-500)|

<a name="lookupaccountapplocalstates-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**apps-local-states**  <br>*required*||< [ApplicationLocalState](#applicationlocalstate) > array|
|**current-round**  <br>*required*|Round at which the results were computed.|integer|
|**next-token**  <br>*optional*|Used for pagination, when making another request provide this token with the next parameter.|string|

<a name="lookupaccountapplocalstates-response-400"></a>
**Response 400**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|

<a name="lookupaccountapplocalstates-response-404"></a>
**Response 404**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|

<a name="lookupaccountapplocalstates-response-500"></a>
**Response 500**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|


**Consumes**

* `application/json`


**Produces**

* `application/json`


**Tags**

* lookup


<a name="lookupaccountassets"></a>
### GET /v2/accounts/{account-id}/assets

**Description**
Lookup an account's asset holdings, optionally for a specific ID.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**account-id**  <br>*required*|account string|string|
|**Query**|**asset-id**  <br>*optional*|Asset ID|integer|
|**Query**|**include-all**  <br>*optional*|Include all items including closed accounts, deleted applications, destroyed assets, opted-out asset holdings, and closed-out application localstates.|boolean|
|**Query**|**limit**  <br>*optional*|Maximum number of results to return. There could be additional pages even if the limit is not reached.|integer|
|**Query**|**next**  <br>*optional*|The next page of results. Use the next token provided by the previous results.|string|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|(empty)|[Response 200](#lookupaccountassets-response-200)|
|**400**|Response for errors|[Response 400](#lookupaccountassets-response-400)|
|**404**|Response for errors|[Response 404](#lookupaccountassets-response-404)|
|**500**|Response for errors|[Response 500](#lookupaccountassets-response-500)|

<a name="lookupaccountassets-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**assets**  <br>*required*||< [AssetHolding](#assetholding) > array|
|**current-round**  <br>*required*|Round at which the results were computed.|integer|
|**next-token**  <br>*optional*|Used for pagination, when making another request provide this token with the next parameter.|string|

<a name="lookupaccountassets-response-400"></a>
**Response 400**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|

<a name="lookupaccountassets-response-404"></a>
**Response 404**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|

<a name="lookupaccountassets-response-500"></a>
**Response 500**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|


**Consumes**

* `application/json`


**Produces**

* `application/json`


**Tags**

* lookup


<a name="lookupaccountcreatedapplications"></a>
### GET /v2/accounts/{account-id}/created-applications

**Description**
Lookup an account's created application parameters, optionally for a specific ID.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**account-id**  <br>*required*|account string|string|
|**Query**|**application-id**  <br>*optional*|Application ID|integer|
|**Query**|**include-all**  <br>*optional*|Include all items including closed accounts, deleted applications, destroyed assets, opted-out asset holdings, and closed-out application localstates.|boolean|
|**Query**|**limit**  <br>*optional*|Maximum number of results to return. There could be additional pages even if the limit is not reached.|integer|
|**Query**|**next**  <br>*optional*|The next page of results. Use the next token provided by the previous results.|string|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|(empty)|[Response 200](#lookupaccountcreatedapplications-response-200)|
|**400**|Response for errors|[Response 400](#lookupaccountcreatedapplications-response-400)|
|**404**|Response for errors|[Response 404](#lookupaccountcreatedapplications-response-404)|
|**500**|Response for errors|[Response 500](#lookupaccountcreatedapplications-response-500)|

<a name="lookupaccountcreatedapplications-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**applications**  <br>*required*||< [Application](#application) > array|
|**current-round**  <br>*required*|Round at which the results were computed.|integer|
|**next-token**  <br>*optional*|Used for pagination, when making another request provide this token with the next parameter.|string|

<a name="lookupaccountcreatedapplications-response-400"></a>
**Response 400**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|

<a name="lookupaccountcreatedapplications-response-404"></a>
**Response 404**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|

<a name="lookupaccountcreatedapplications-response-500"></a>
**Response 500**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|


**Consumes**

* `application/json`


**Produces**

* `application/json`


**Tags**

* lookup


<a name="lookupaccountcreatedassets"></a>
### GET /v2/accounts/{account-id}/created-assets

**Description**
Lookup an account's created asset parameters, optionally for a specific ID.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**account-id**  <br>*required*|account string|string|
|**Query**|**asset-id**  <br>*optional*|Asset ID|integer|
|**Query**|**include-all**  <br>*optional*|Include all items including closed accounts, deleted applications, destroyed assets, opted-out asset holdings, and closed-out application localstates.|boolean|
|**Query**|**limit**  <br>*optional*|Maximum number of results to return. There could be additional pages even if the limit is not reached.|integer|
|**Query**|**next**  <br>*optional*|The next page of results. Use the next token provided by the previous results.|string|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|(empty)|[Response 200](#lookupaccountcreatedassets-response-200)|
|**400**|Response for errors|[Response 400](#lookupaccountcreatedassets-response-400)|
|**404**|Response for errors|[Response 404](#lookupaccountcreatedassets-response-404)|
|**500**|Response for errors|[Response 500](#lookupaccountcreatedassets-response-500)|

<a name="lookupaccountcreatedassets-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**assets**  <br>*required*||< [Asset](#asset) > array|
|**current-round**  <br>*required*|Round at which the results were computed.|integer|
|**next-token**  <br>*optional*|Used for pagination, when making another request provide this token with the next parameter.|string|

<a name="lookupaccountcreatedassets-response-400"></a>
**Response 400**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|

<a name="lookupaccountcreatedassets-response-404"></a>
**Response 404**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|

<a name="lookupaccountcreatedassets-response-500"></a>
**Response 500**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|


**Consumes**

* `application/json`


**Produces**

* `application/json`


**Tags**

* lookup


<a name="lookupaccounttransactions"></a>
### GET /v2/accounts/{account-id}/transactions

**Description**
Lookup account transactions. Transactions are returned newest to oldest.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**account-id**  <br>*required*|account string|string|
|**Query**|**after-time**  <br>*optional*|Include results after the given time. Must be an RFC 3339 formatted string.|string (date-time)|
|**Query**|**asset-id**  <br>*optional*|Asset ID|integer|
|**Query**|**before-time**  <br>*optional*|Include results before the given time. Must be an RFC 3339 formatted string.|string (date-time)|
|**Query**|**currency-greater-than**  <br>*optional*|Results should have an amount greater than this value. MicroAlgos are the default currency unless an asset-id is provided, in which case the asset will be used.|integer|
|**Query**|**currency-less-than**  <br>*optional*|Results should have an amount less than this value. MicroAlgos are the default currency unless an asset-id is provided, in which case the asset will be used.|integer|
|**Query**|**limit**  <br>*optional*|Maximum number of results to return. There could be additional pages even if the limit is not reached.|integer|
|**Query**|**max-round**  <br>*optional*|Include results at or before the specified max-round.|integer|
|**Query**|**min-round**  <br>*optional*|Include results at or after the specified min-round.|integer|
|**Query**|**next**  <br>*optional*|The next page of results. Use the next token provided by the previous results.|string|
|**Query**|**note-prefix**  <br>*optional*|Specifies a prefix which must be contained in the note field.|string|
|**Query**|**rekey-to**  <br>*optional*|Include results which include the rekey-to field.|boolean|
|**Query**|**round**  <br>*optional*|Include results for the specified round.|integer|
|**Query**|**sig-type**  <br>*optional*|SigType filters just results using the specified type of signature:<br>* sig - Standard<br>* msig - MultiSig<br>* lsig - LogicSig|enum (sig, msig, lsig)|
|**Query**|**tx-type**  <br>*optional*||enum (pay, keyreg, acfg, axfer, afrz, appl, stpf)|
|**Query**|**txid**  <br>*optional*|Lookup the specific transaction by ID.|string|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|(empty)|[Response 200](#lookupaccounttransactions-response-200)|
|**400**|Response for errors|[Response 400](#lookupaccounttransactions-response-400)|
|**500**|Response for errors|[Response 500](#lookupaccounttransactions-response-500)|

<a name="lookupaccounttransactions-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**current-round**  <br>*required*|Round at which the results were computed.|integer|
|**next-token**  <br>*optional*|Used for pagination, when making another request provide this token with the next parameter.|string|
|**transactions**  <br>*required*||< [Transaction](#transaction) > array|

<a name="lookupaccounttransactions-response-400"></a>
**Response 400**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|

<a name="lookupaccounttransactions-response-500"></a>
**Response 500**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|


**Consumes**

* `application/json`


**Produces**

* `application/json`


**Tags**

* lookup


<a name="searchforapplications"></a>
### GET /v2/applications

**Description**
Search for applications


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Query**|**application-id**  <br>*optional*|Application ID|integer|
|**Query**|**creator**  <br>*optional*|Filter just applications with the given creator address.|string|
|**Query**|**include-all**  <br>*optional*|Include all items including closed accounts, deleted applications, destroyed assets, opted-out asset holdings, and closed-out application localstates.|boolean|
|**Query**|**limit**  <br>*optional*|Maximum number of results to return. There could be additional pages even if the limit is not reached.|integer|
|**Query**|**next**  <br>*optional*|The next page of results. Use the next token provided by the previous results.|string|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|(empty)|[Response 200](#searchforapplications-response-200)|
|**500**|Response for errors|[Response 500](#searchforapplications-response-500)|

<a name="searchforapplications-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**applications**  <br>*required*||< [Application](#application) > array|
|**current-round**  <br>*required*|Round at which the results were computed.|integer|
|**next-token**  <br>*optional*|Used for pagination, when making another request provide this token with the next parameter.|string|

<a name="searchforapplications-response-500"></a>
**Response 500**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|


**Consumes**

* `application/json`


**Produces**

* `application/json`


**Tags**

* search


<a name="lookupapplicationbyid"></a>
### GET /v2/applications/{application-id}

**Description**
Lookup application.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**application-id**  <br>*required*||integer|
|**Query**|**include-all**  <br>*optional*|Include all items including closed accounts, deleted applications, destroyed assets, opted-out asset holdings, and closed-out application localstates.|boolean|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|(empty)|[Response 200](#lookupapplicationbyid-response-200)|
|**404**|Response for errors|[Response 404](#lookupapplicationbyid-response-404)|
|**500**|Response for errors|[Response 500](#lookupapplicationbyid-response-500)|

<a name="lookupapplicationbyid-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**application**  <br>*optional*||[Application](#application)|
|**current-round**  <br>*required*|Round at which the results were computed.|integer|

<a name="lookupapplicationbyid-response-404"></a>
**Response 404**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|

<a name="lookupapplicationbyid-response-500"></a>
**Response 500**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|


**Consumes**

* `application/json`


**Produces**

* `application/json`


**Tags**

* lookup


<a name="lookupapplicationboxbyidandname"></a>
### GET /v2/applications/{application-id}/box
Get box information for a given application.
```
GET /v2/applications/{application-id}/box
```


**Description**
Given an application ID and box name, returns base64 encoded box name and value. Box names must be in the goal app call arg form 'encoding:value'. For ints, use the form 'int:1234'. For raw bytes, encode base 64 and use 'b64' prefix as in 'b64:A=='. For printable strings, use the form 'str:hello'. For addresses, use the form 'addr:XYZ...'.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**application-id**  <br>*required*||integer|
|**Query**|**name**  <br>*required*|A box name in goal-arg form 'encoding:value'. For ints, use the form 'int:1234'. For raw bytes, use the form 'b64:A=='. For printable strings, use the form 'str:hello'. For addresses, use the form 'addr:XYZ...'.|string|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Box information|[Box](#box)|
|**400**|Response for errors|[Response 400](#lookupapplicationboxbyidandname-response-400)|
|**404**|Response for errors|[Response 404](#lookupapplicationboxbyidandname-response-404)|
|**500**|Response for errors|[Response 500](#lookupapplicationboxbyidandname-response-500)|

<a name="lookupapplicationboxbyidandname-response-400"></a>
**Response 400**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|

<a name="lookupapplicationboxbyidandname-response-404"></a>
**Response 404**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|

<a name="lookupapplicationboxbyidandname-response-500"></a>
**Response 500**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|


**Consumes**

* `application/json`


**Produces**

* `application/json`


**Tags**

* lookup


<a name="searchforapplicationboxes"></a>
### GET /v2/applications/{application-id}/boxes
Get box names for a given application.
```
GET /v2/applications/{application-id}/boxes
```


**Description**
Given an application ID, returns the box names of that application sorted lexicographically.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**application-id**  <br>*required*||integer|
|**Query**|**limit**  <br>*optional*|Maximum number of results to return. There could be additional pages even if the limit is not reached.|integer|
|**Query**|**next**  <br>*optional*|The next page of results. Use the next token provided by the previous results.|string|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Box names of an application|[Response 200](#searchforapplicationboxes-response-200)|
|**400**|Response for errors|[Response 400](#searchforapplicationboxes-response-400)|
|**404**|Response for errors|[Response 404](#searchforapplicationboxes-response-404)|
|**500**|Response for errors|[Response 500](#searchforapplicationboxes-response-500)|

<a name="searchforapplicationboxes-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**application-id**  <br>*required*|\[appidx\] application index.|integer|
|**boxes**  <br>*required*||< [BoxDescriptor](#boxdescriptor) > array|
|**next-token**  <br>*optional*|Used for pagination, when making another request provide this token with the next parameter.|string|

<a name="searchforapplicationboxes-response-400"></a>
**Response 400**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|

<a name="searchforapplicationboxes-response-404"></a>
**Response 404**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|

<a name="searchforapplicationboxes-response-500"></a>
**Response 500**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|


**Consumes**

* `application/json`


**Produces**

* `application/json`


**Tags**

* search


<a name="lookupapplicationlogsbyid"></a>
### GET /v2/applications/{application-id}/logs

**Description**
Lookup application logs.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**application-id**  <br>*required*||integer|
|**Query**|**limit**  <br>*optional*|Maximum number of results to return. There could be additional pages even if the limit is not reached.|integer|
|**Query**|**max-round**  <br>*optional*|Include results at or before the specified max-round.|integer|
|**Query**|**min-round**  <br>*optional*|Include results at or after the specified min-round.|integer|
|**Query**|**next**  <br>*optional*|The next page of results. Use the next token provided by the previous results.|string|
|**Query**|**sender-address**  <br>*optional*|Only include transactions with this sender address.|string|
|**Query**|**txid**  <br>*optional*|Lookup the specific transaction by ID.|string|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|(empty)|[Response 200](#lookupapplicationlogsbyid-response-200)|

<a name="lookupapplicationlogsbyid-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**application-id**  <br>*required*|\[appidx\] application index.|integer|
|**current-round**  <br>*required*|Round at which the results were computed.|integer|
|**log-data**  <br>*optional*||< [ApplicationLogData](#applicationlogdata) > array|
|**next-token**  <br>*optional*|Used for pagination, when making another request provide this token with the next parameter.|string|


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
|**Query**|**include-all**  <br>*optional*|Include all items including closed accounts, deleted applications, destroyed assets, opted-out asset holdings, and closed-out application localstates.|boolean|
|**Query**|**limit**  <br>*optional*|Maximum number of results to return. There could be additional pages even if the limit is not reached.|integer|
|**Query**|**name**  <br>*optional*|Filter just assets with the given name.|string|
|**Query**|**next**  <br>*optional*|The next page of results. Use the next token provided by the previous results.|string|
|**Query**|**unit**  <br>*optional*|Filter just assets with the given unit.|string|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|(empty)|[Response 200](#searchforassets-response-200)|
|**400**|Response for errors|[Response 400](#searchforassets-response-400)|
|**500**|Response for errors|[Response 500](#searchforassets-response-500)|

<a name="searchforassets-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**assets**  <br>*required*||< [Asset](#asset) > array|
|**current-round**  <br>*required*|Round at which the results were computed.|integer|
|**next-token**  <br>*optional*|Used for pagination, when making another request provide this token with the next parameter.|string|

<a name="searchforassets-response-400"></a>
**Response 400**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|

<a name="searchforassets-response-500"></a>
**Response 500**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|


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

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**asset-id**  <br>*required*||integer|
|**Query**|**include-all**  <br>*optional*|Include all items including closed accounts, deleted applications, destroyed assets, opted-out asset holdings, and closed-out application localstates.|boolean|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|(empty)|[Response 200](#lookupassetbyid-response-200)|
|**400**|Response for errors|[Response 400](#lookupassetbyid-response-400)|
|**404**|Response for errors|[Response 404](#lookupassetbyid-response-404)|
|**500**|Response for errors|[Response 500](#lookupassetbyid-response-500)|

<a name="lookupassetbyid-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**asset**  <br>*required*||[Asset](#asset)|
|**current-round**  <br>*required*|Round at which the results were computed.|integer|

<a name="lookupassetbyid-response-400"></a>
**Response 400**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|

<a name="lookupassetbyid-response-404"></a>
**Response 404**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|

<a name="lookupassetbyid-response-500"></a>
**Response 500**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|


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
|**Query**|**include-all**  <br>*optional*|Include all items including closed accounts, deleted applications, destroyed assets, opted-out asset holdings, and closed-out application localstates.|boolean|
|**Query**|**limit**  <br>*optional*|Maximum number of results to return. There could be additional pages even if the limit is not reached.|integer|
|**Query**|**next**  <br>*optional*|The next page of results. Use the next token provided by the previous results.|string|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|(empty)|[Response 200](#lookupassetbalances-response-200)|
|**400**|Response for errors|[Response 400](#lookupassetbalances-response-400)|
|**500**|Response for errors|[Response 500](#lookupassetbalances-response-500)|

<a name="lookupassetbalances-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**balances**  <br>*required*||< [MiniAssetHolding](#miniassetholding) > array|
|**current-round**  <br>*required*|Round at which the results were computed.|integer|
|**next-token**  <br>*optional*|Used for pagination, when making another request provide this token with the next parameter.|string|

<a name="lookupassetbalances-response-400"></a>
**Response 400**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|

<a name="lookupassetbalances-response-500"></a>
**Response 500**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|


**Consumes**

* `application/json`


**Produces**

* `application/json`


**Tags**

* lookup


<a name="lookupassettransactions"></a>
### GET /v2/assets/{asset-id}/transactions

**Description**
Lookup transactions for an asset. Transactions are returned oldest to newest.


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
|**Query**|**limit**  <br>*optional*|Maximum number of results to return. There could be additional pages even if the limit is not reached.|integer|
|**Query**|**max-round**  <br>*optional*|Include results at or before the specified max-round.|integer|
|**Query**|**min-round**  <br>*optional*|Include results at or after the specified min-round.|integer|
|**Query**|**next**  <br>*optional*|The next page of results. Use the next token provided by the previous results.|string|
|**Query**|**note-prefix**  <br>*optional*|Specifies a prefix which must be contained in the note field.|string|
|**Query**|**rekey-to**  <br>*optional*|Include results which include the rekey-to field.|boolean|
|**Query**|**round**  <br>*optional*|Include results for the specified round.|integer|
|**Query**|**sig-type**  <br>*optional*|SigType filters just results using the specified type of signature:<br>* sig - Standard<br>* msig - MultiSig<br>* lsig - LogicSig|enum (sig, msig, lsig)|
|**Query**|**tx-type**  <br>*optional*||enum (pay, keyreg, acfg, axfer, afrz, appl, stpf)|
|**Query**|**txid**  <br>*optional*|Lookup the specific transaction by ID.|string|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|(empty)|[Response 200](#lookupassettransactions-response-200)|
|**400**|Response for errors|[Response 400](#lookupassettransactions-response-400)|
|**500**|Response for errors|[Response 500](#lookupassettransactions-response-500)|

<a name="lookupassettransactions-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**current-round**  <br>*required*|Round at which the results were computed.|integer|
|**next-token**  <br>*optional*|Used for pagination, when making another request provide this token with the next parameter.|string|
|**transactions**  <br>*required*||< [Transaction](#transaction) > array|

<a name="lookupassettransactions-response-400"></a>
**Response 400**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|

<a name="lookupassettransactions-response-500"></a>
**Response 500**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|


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
|**Query**|**header-only**  <br>*optional*|Header only flag. When this is set to true, returned block does not contain the transactions|boolean|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|(empty)|[Block](#block)|
|**404**|Response for errors|[Response 404](#lookupblock-response-404)|
|**500**|Response for errors|[Response 500](#lookupblock-response-500)|

<a name="lookupblock-response-404"></a>
**Response 404**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|

<a name="lookupblock-response-500"></a>
**Response 500**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|


**Consumes**

* `application/json`


**Produces**

* `application/json`


**Tags**

* lookup


<a name="searchfortransactions"></a>
### GET /v2/transactions

**Description**
Search for transactions. Transactions are returned oldest to newest unless the address parameter is used, in which case results are returned newest to oldest.


**Parameters**

|Type|Name|Description|Schema|
|---|---|---|---|
|**Query**|**address**  <br>*optional*|Only include transactions with this address in one of the transaction fields.|string|
|**Query**|**address-role**  <br>*optional*|Combine with the address parameter to define what type of address to search for.|enum (sender, receiver, freeze-target)|
|**Query**|**after-time**  <br>*optional*|Include results after the given time. Must be an RFC 3339 formatted string.|string (date-time)|
|**Query**|**application-id**  <br>*optional*|Application ID|integer|
|**Query**|**asset-id**  <br>*optional*|Asset ID|integer|
|**Query**|**before-time**  <br>*optional*|Include results before the given time. Must be an RFC 3339 formatted string.|string (date-time)|
|**Query**|**currency-greater-than**  <br>*optional*|Results should have an amount greater than this value. MicroAlgos are the default currency unless an asset-id is provided, in which case the asset will be used.|integer|
|**Query**|**currency-less-than**  <br>*optional*|Results should have an amount less than this value. MicroAlgos are the default currency unless an asset-id is provided, in which case the asset will be used.|integer|
|**Query**|**exclude-close-to**  <br>*optional*|Combine with address and address-role parameters to define what type of address to search for. The close to fields are normally treated as a receiver, if you would like to exclude them set this parameter to true.|boolean|
|**Query**|**limit**  <br>*optional*|Maximum number of results to return. There could be additional pages even if the limit is not reached.|integer|
|**Query**|**max-round**  <br>*optional*|Include results at or before the specified max-round.|integer|
|**Query**|**min-round**  <br>*optional*|Include results at or after the specified min-round.|integer|
|**Query**|**next**  <br>*optional*|The next page of results. Use the next token provided by the previous results.|string|
|**Query**|**note-prefix**  <br>*optional*|Specifies a prefix which must be contained in the note field.|string|
|**Query**|**rekey-to**  <br>*optional*|Include results which include the rekey-to field.|boolean|
|**Query**|**round**  <br>*optional*|Include results for the specified round.|integer|
|**Query**|**sig-type**  <br>*optional*|SigType filters just results using the specified type of signature:<br>* sig - Standard<br>* msig - MultiSig<br>* lsig - LogicSig|enum (sig, msig, lsig)|
|**Query**|**tx-type**  <br>*optional*||enum (pay, keyreg, acfg, axfer, afrz, appl, stpf)|
|**Query**|**txid**  <br>*optional*|Lookup the specific transaction by ID.|string|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|(empty)|[Response 200](#searchfortransactions-response-200)|
|**400**|Response for errors|[Response 400](#searchfortransactions-response-400)|
|**500**|Response for errors|[Response 500](#searchfortransactions-response-500)|

<a name="searchfortransactions-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**current-round**  <br>*required*|Round at which the results were computed.|integer|
|**next-token**  <br>*optional*|Used for pagination, when making another request provide this token with the next parameter.|string|
|**transactions**  <br>*required*||< [Transaction](#transaction) > array|

<a name="searchfortransactions-response-400"></a>
**Response 400**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|

<a name="searchfortransactions-response-500"></a>
**Response 500**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|


**Consumes**

* `application/json`


**Produces**

* `application/json`


**Tags**

* search


<a name="lookuptransaction"></a>
### GET /v2/transactions/{txid}

**Description**
Lookup a single transaction.


**Parameters**

|Type|Name|Schema|
|---|---|---|
|**Path**|**txid**  <br>*required*|string|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|(empty)|[Response 200](#lookuptransaction-response-200)|
|**400**|Response for errors|[Response 400](#lookuptransaction-response-400)|
|**404**|Response for errors|[Response 404](#lookuptransaction-response-404)|
|**500**|Response for errors|[Response 500](#lookuptransaction-response-500)|

<a name="lookuptransaction-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**current-round**  <br>*required*|Round at which the results were computed.|integer|
|**transaction**  <br>*required*||[Transaction](#transaction)|

<a name="lookuptransaction-response-400"></a>
**Response 400**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|

<a name="lookuptransaction-response-404"></a>
**Response 404**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|

<a name="lookuptransaction-response-500"></a>
**Response 500**

|Name|Schema|
|---|---|
|**data**  <br>*optional*|object|
|**message**  <br>*required*|string|


**Consumes**

* `application/json`


**Produces**

* `application/json`


**Tags**

* lookup




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
|**amount**  <br>*required*|total number of MicroAlgos in the account|integer|
|**amount-without-pending-rewards**  <br>*required*|specifies the amount of MicroAlgos in the account, without the pending rewards.|integer|
|**apps-local-state**  <br>*optional*|application local data stored in this account.<br><br>Note the raw object uses `map[int] -> AppLocalState` for this type.|< [ApplicationLocalState](#applicationlocalstate) > array|
|**apps-total-extra-pages**  <br>*optional*|the sum of all extra application program pages for this account.|integer|
|**apps-total-schema**  <br>*optional*|the sum of all of the local schemas and global schemas in this account.<br><br>Note: the raw account uses `StateSchema` for this type.|[ApplicationStateSchema](#applicationstateschema)|
|**assets**  <br>*optional*|assets held by this account.<br><br>Note the raw object uses `map[int] -> AssetHolding` for this type.|< [AssetHolding](#assetholding) > array|
|**auth-addr**  <br>*optional*|The address against which signing should be checked. If empty, the address of the current account is used. This field can be updated in any transaction by setting the RekeyTo field.|string|
|**closed-at-round**  <br>*optional*|Round during which this account was most recently closed.|integer|
|**created-apps**  <br>*optional*|parameters of applications created by this account including app global data.<br><br>Note: the raw account uses `map[int] -> AppParams` for this type.|< [Application](#application) > array|
|**created-assets**  <br>*optional*|parameters of assets created by this account.<br><br>Note: the raw account uses `map[int] -> Asset` for this type.|< [Asset](#asset) > array|
|**created-at-round**  <br>*optional*|Round during which this account first appeared in a transaction.|integer|
|**deleted**  <br>*optional*|Whether or not this account is currently closed.|boolean|
|**incentive-eligible**  <br>*optional*|can the account receive block incentives if its balance is in range at proposal time.|boolean|
|**last-heartbeat**  <br>*optional*|The round in which this account last went online, or explicitly renewed their online status.|integer|
|**last-proposed**  <br>*optional*|The round in which this account last proposed the block.|integer|
|**min-balance**  <br>*required*|MicroAlgo balance required by the account.<br><br>The requirement grows based on asset and application usage.|integer|
|**participation**  <br>*optional*||[AccountParticipation](#accountparticipation)|
|**pending-rewards**  <br>*required*|amount of MicroAlgos of pending rewards in this account.|integer|
|**reward-base**  <br>*optional*|used as part of the rewards computation. Only applicable to accounts which are participating.|integer|
|**rewards**  <br>*required*|total rewards of MicroAlgos the account has received, including pending rewards.|integer|
|**round**  <br>*required*|The round for which this information is relevant.|integer|
|**sig-type**  <br>*optional*|the type of signature used by this account, must be one of:<br>* sig<br>* msig<br>* lsig<br>* or null if unknown|enum (sig, msig, lsig)|
|**status**  <br>*required*|voting status of the account's MicroAlgos<br>* Offline - indicates that the associated account is delegated.<br>*  Online  - indicates that the associated account used as part of the delegation pool.<br>*   NotParticipating - indicates that the associated account is neither a delegator nor a delegate.|string|
|**total-apps-opted-in**  <br>*required*|The count of all applications that have been opted in, equivalent to the count of application local data (AppLocalState objects) stored in this account.|integer|
|**total-assets-opted-in**  <br>*required*|The count of all assets that have been opted in, equivalent to the count of AssetHolding objects held by this account.|integer|
|**total-box-bytes**  <br>*required*|For app-accounts only. The total number of bytes allocated for the keys and values of boxes which belong to the associated application.|integer|
|**total-boxes**  <br>*required*|For app-accounts only. The total number of boxes which belong to the associated application.|integer|
|**total-created-apps**  <br>*required*|The count of all apps (AppParams objects) created by this account.|integer|
|**total-created-assets**  <br>*required*|The count of all assets (AssetParams objects) created by this account.|integer|


<a name="accountparticipation"></a>
### AccountParticipation
AccountParticipation describes the parameters used by this account in consensus protocol.


|Name|Description|Schema|
|---|---|---|
|**selection-participation-key**  <br>*required*|Selection public key (if any) currently registered for this round.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**state-proof-key**  <br>*optional*|Root of the state proof key (if any)  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**vote-first-valid**  <br>*required*|First round for which this participation is valid.|integer|
|**vote-key-dilution**  <br>*required*|Number of subkeys in each batch of participation keys.|integer|
|**vote-last-valid**  <br>*required*|Last round for which this participation is valid.|integer|
|**vote-participation-key**  <br>*required*|root participation public key (if any) currently registered for this round.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|


<a name="accountstatedelta"></a>
### AccountStateDelta
Application state delta.


|Name|Schema|
|---|---|
|**address**  <br>*required*|string|
|**delta**  <br>*required*|[StateDelta](#statedelta)|


<a name="application"></a>
### Application
Application index and its parameters


|Name|Description|Schema|
|---|---|---|
|**created-at-round**  <br>*optional*|Round when this application was created.|integer|
|**deleted**  <br>*optional*|Whether or not this application is currently deleted.|boolean|
|**deleted-at-round**  <br>*optional*|Round when this application was deleted.|integer|
|**id**  <br>*required*|application index.|integer|
|**params**  <br>*required*|application parameters.|[ApplicationParams](#applicationparams)|


<a name="applicationlocalstate"></a>
### ApplicationLocalState
Stores local state associated with an application.


|Name|Description|Schema|
|---|---|---|
|**closed-out-at-round**  <br>*optional*|Round when account closed out of the application.|integer|
|**deleted**  <br>*optional*|Whether or not the application local state is currently deleted from its account.|boolean|
|**id**  <br>*required*|The application which this local state is for.|integer|
|**key-value**  <br>*optional*|storage.|[TealKeyValueStore](#tealkeyvaluestore)|
|**opted-in-at-round**  <br>*optional*|Round when the account opted into the application.|integer|
|**schema**  <br>*required*|schema.|[ApplicationStateSchema](#applicationstateschema)|


<a name="applicationlogdata"></a>
### ApplicationLogData
Stores the global information associated with an application.


|Name|Description|Schema|
|---|---|---|
|**logs**  <br>*required*|Logs for the application being executed by the transaction.|< string (byte) > array|
|**txid**  <br>*required*|Transaction ID|string|


<a name="applicationparams"></a>
### ApplicationParams
Stores the global information associated with an application.


|Name|Description|Schema|
|---|---|---|
|**approval-program**  <br>*required*|approval program.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**clear-state-program**  <br>*required*|clear state program.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**creator**  <br>*optional*|The address that created this application. This is the address where the parameters and global state for this application can be found.|string|
|**extra-program-pages**  <br>*optional*|the number of extra program pages available to this app.|integer|
|**global-state**  <br>*optional*|global state|[TealKeyValueStore](#tealkeyvaluestore)|
|**global-state-schema**  <br>*optional*|global schema|[ApplicationStateSchema](#applicationstateschema)|
|**local-state-schema**  <br>*optional*|local schema|[ApplicationStateSchema](#applicationstateschema)|


<a name="applicationstateschema"></a>
### ApplicationStateSchema
Specifies maximums on the number of each type that may be stored.


|Name|Description|Schema|
|---|---|---|
|**num-byte-slice**  <br>*required*|number of byte slices.|integer|
|**num-uint**  <br>*required*|number of uints.|integer|


<a name="asset"></a>
### Asset
Specifies both the unique identifier and the parameters for an asset


|Name|Description|Schema|
|---|---|---|
|**created-at-round**  <br>*optional*|Round during which this asset was created.|integer|
|**deleted**  <br>*optional*|Whether or not this asset is currently deleted.|boolean|
|**destroyed-at-round**  <br>*optional*|Round during which this asset was destroyed.|integer|
|**index**  <br>*required*|unique asset identifier|integer|
|**params**  <br>*required*||[AssetParams](#assetparams)|


<a name="assetholding"></a>
### AssetHolding
Describes an asset held by an account.

Definition:
data/basics/userBalance.go : AssetHolding


|Name|Description|Schema|
|---|---|---|
|**amount**  <br>*required*|number of units held.|integer|
|**asset-id**  <br>*required*|Asset ID of the holding.|integer|
|**deleted**  <br>*optional*|Whether or not the asset holding is currently deleted from its account.|boolean|
|**is-frozen**  <br>*required*|whether or not the holding is frozen.|boolean|
|**opted-in-at-round**  <br>*optional*|Round during which the account opted into this asset holding.|integer|
|**opted-out-at-round**  <br>*optional*|Round during which the account opted out of this asset holding.|integer|


<a name="assetparams"></a>
### AssetParams
AssetParams specifies the parameters for an asset.

\[apar\] when part of an AssetConfig transaction.

Definition:
data/transactions/asset.go : AssetParams


|Name|Description|Schema|
|---|---|---|
|**clawback**  <br>*optional*|Address of account used to clawback holdings of this asset.  If empty, clawback is not permitted.|string|
|**creator**  <br>*required*|The address that created this asset. This is the address where the parameters for this asset can be found, and also the address where unwanted asset units can be sent in the worst case.|string|
|**decimals**  <br>*required*|The number of digits to use after the decimal point when displaying this asset. If 0, the asset is not divisible. If 1, the base unit of the asset is in tenths. If 2, the base unit of the asset is in hundredths, and so on. This value must be between 0 and 19 (inclusive).  <br>**Minimum value** : `0`  <br>**Maximum value** : `19`|integer|
|**default-frozen**  <br>*optional*|Whether holdings of this asset are frozen by default.|boolean|
|**freeze**  <br>*optional*|Address of account used to freeze holdings of this asset.  If empty, freezing is not permitted.|string|
|**manager**  <br>*optional*|Address of account used to manage the keys of this asset and to destroy it.|string|
|**metadata-hash**  <br>*optional*|A commitment to some unspecified asset metadata. The format of this metadata is up to the application.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**name**  <br>*optional*|Name of this asset, as supplied by the creator. Included only when the asset name is composed of printable utf-8 characters.|string|
|**name-b64**  <br>*optional*|Base64 encoded name of this asset, as supplied by the creator.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**reserve**  <br>*optional*|Address of account holding reserve (non-minted) units of this asset.|string|
|**total**  <br>*required*|The total number of units of this asset.|integer|
|**unit-name**  <br>*optional*|Name of a unit of this asset, as supplied by the creator. Included only when the name of a unit of this asset is composed of printable utf-8 characters.|string|
|**unit-name-b64**  <br>*optional*|Base64 encoded name of a unit of this asset, as supplied by the creator.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**url**  <br>*optional*|URL where more information about the asset can be retrieved. Included only when the URL is composed of printable utf-8 characters.|string|
|**url-b64**  <br>*optional*|Base64 encoded URL where more information about the asset can be retrieved.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|


<a name="block"></a>
### Block
Block information.

Definition:
data/bookkeeping/block.go : Block


|Name|Description|Schema|
|---|---|---|
|**bonus**  <br>*optional*|the potential bonus payout for this block.|integer|
|**fees-collected**  <br>*optional*|the sum of all fees paid by transactions in this block.|integer|
|**genesis-hash**  <br>*required*|\[gh\] hash to which this block belongs.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**genesis-id**  <br>*required*|\[gen\] ID to which this block belongs.|string|
|**participation-updates**  <br>*optional*||[ParticipationUpdates](#participationupdates)|
|**previous-block-hash**  <br>*required*|\[prev\] Previous block hash.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**proposer**  <br>*optional*|the proposer of this block.|string|
|**proposer-payout**  <br>*optional*|the actual amount transferred to the proposer from the fee sink.|integer|
|**rewards**  <br>*optional*||[BlockRewards](#blockrewards)|
|**round**  <br>*required*|\[rnd\] Current round on which this block was appended to the chain.|integer|
|**seed**  <br>*required*|\[seed\] Sortition seed.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**state-proof-tracking**  <br>*optional*|Tracks the status of state proofs.|< [StateProofTracking](#stateprooftracking) > array|
|**timestamp**  <br>*required*|\[ts\] Block creation timestamp in seconds since eposh|integer|
|**transactions**  <br>*optional*|\[txns\] list of transactions corresponding to a given round.|< [Transaction](#transaction) > array|
|**transactions-root**  <br>*required*|\[txn\] TransactionsRoot authenticates the set of transactions appearing in the block. More specifically, it's the root of a merkle tree whose leaves are the block's Txids, in lexicographic order. For the empty block, it's 0. Note that the TxnRoot does not authenticate the signatures on the transactions, only the transactions themselves. Two blocks with the same transactions but in a different order and with different signatures will have the same TxnRoot.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**transactions-root-sha256**  <br>*required*|\[txn256\] TransactionsRootSHA256 is an auxiliary TransactionRoot, built using a vector commitment instead of a merkle tree, and SHA256 hash function instead of the default SHA512_256. This commitment can be used on environments where only the SHA256 function exists.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
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
Box descriptor describes an app box without a value.


|Name|Description|Schema|
|---|---|---|
|**name**  <br>*required*|Base64 encoded box name  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|


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


<a name="hashfactory"></a>
### HashFactory

|Name|Description|Schema|
|---|---|---|
|**hash-type**  <br>*optional*|\[t\]|integer|


<a name="hashtype"></a>
### Hashtype
The type of hash function used to create the proof, must be one of: 
* sha512_256 
* sha256

*Type* : enum (sha512_256, sha256)


<a name="healthcheck"></a>
### HealthCheck
A health check response.


|Name|Description|Schema|
|---|---|---|
|**data**  <br>*optional*||object|
|**db-available**  <br>*required*||boolean|
|**errors**  <br>*optional*||< string > array|
|**is-migrating**  <br>*required*||boolean|
|**message**  <br>*required*||string|
|**round**  <br>*required*||integer|
|**version**  <br>*required*|Current version.|string|


<a name="indexerstateproofmessage"></a>
### IndexerStateProofMessage

|Name|Description|Schema|
|---|---|---|
|**block-headers-commitment**  <br>*optional*|\[b\]  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**first-attested-round**  <br>*optional*|\[f\]|integer|
|**latest-attested-round**  <br>*optional*|\[l\]|integer|
|**ln-proven-weight**  <br>*optional*|\[P\]|integer|
|**voters-commitment**  <br>*optional*|\[v\]  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|


<a name="merklearrayproof"></a>
### MerkleArrayProof

|Name|Description|Schema|
|---|---|---|
|**hash-factory**  <br>*optional*||[HashFactory](#hashfactory)|
|**path**  <br>*optional*|\[pth\]|< string (byte) > array|
|**tree-depth**  <br>*optional*|\[td\]|integer|


<a name="miniassetholding"></a>
### MiniAssetHolding
A simplified version of AssetHolding


|Name|Description|Schema|
|---|---|---|
|**address**  <br>*required*||string|
|**amount**  <br>*required*||integer|
|**deleted**  <br>*optional*|Whether or not this asset holding is currently deleted from its account.|boolean|
|**is-frozen**  <br>*required*||boolean|
|**opted-in-at-round**  <br>*optional*|Round during which the account opted into the asset.|integer|
|**opted-out-at-round**  <br>*optional*|Round during which the account opted out of the asset.|integer|


<a name="oncompletion"></a>
### OnCompletion
\[apan\] defines the what additional actions occur with the transaction.

Valid types:
* noop
* optin
* closeout
* clear
* update
* update
* delete

*Type* : enum (noop, optin, closeout, clear, update, delete)


<a name="participationupdates"></a>
### ParticipationUpdates
Participation account data that needs to be checked/acted on by the network.


|Name|Description|Schema|
|---|---|---|
|**absent-participation-accounts**  <br>*optional*|\[partupabs\] a list of online accounts that need to be suspended.|< string > array|
|**expired-participation-accounts**  <br>*optional*|\[partupdrmv\] a list of online accounts that needs to be converted to offline since their participation key expired.|< string > array|


<a name="statedelta"></a>
### StateDelta
Application state delta.

*Type* : < [EvalDeltaKeyValue](#evaldeltakeyvalue) > array


<a name="stateprooffields"></a>
### StateProofFields
\[sp\] represents a state proof.

Definition:
crypto/stateproof/structs.go : StateProof


|Name|Description|Schema|
|---|---|---|
|**part-proofs**  <br>*optional*|\[P\]|[MerkleArrayProof](#merklearrayproof)|
|**positions-to-reveal**  <br>*optional*|\[pr\] Sequence of reveal positions.|< integer > array|
|**reveals**  <br>*optional*|\[r\] Note that this is actually stored as a map[uint64] - Reveal in the actual msgp|< [StateProofReveal](#stateproofreveal) > array|
|**salt-version**  <br>*optional*|\[v\] Salt version of the merkle signature.|integer|
|**sig-commit**  <br>*optional*|\[c\]  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**sig-proofs**  <br>*optional*|\[S\]|[MerkleArrayProof](#merklearrayproof)|
|**signed-weight**  <br>*optional*|\[w\]|integer|


<a name="stateproofparticipant"></a>
### StateProofParticipant

|Name|Description|Schema|
|---|---|---|
|**verifier**  <br>*optional*|\[p\]|[StateProofVerifier](#stateproofverifier)|
|**weight**  <br>*optional*|\[w\]|integer|


<a name="stateproofreveal"></a>
### StateProofReveal

|Name|Description|Schema|
|---|---|---|
|**participant**  <br>*optional*|\[p\]|[StateProofParticipant](#stateproofparticipant)|
|**position**  <br>*optional*|The position in the signature and participants arrays corresponding to this entry.|integer|
|**sig-slot**  <br>*optional*|\[s\]|[StateProofSigSlot](#stateproofsigslot)|


<a name="stateproofsigslot"></a>
### StateProofSigSlot

|Name|Description|Schema|
|---|---|---|
|**lower-sig-weight**  <br>*optional*|\[l\] The total weight of signatures in the lower-numbered slots.|integer|
|**signature**  <br>*optional*||[StateProofSignature](#stateproofsignature)|


<a name="stateproofsignature"></a>
### StateProofSignature

|Name|Description|Schema|
|---|---|---|
|**falcon-signature**  <br>*optional*|**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**merkle-array-index**  <br>*optional*||integer|
|**proof**  <br>*optional*||[MerkleArrayProof](#merklearrayproof)|
|**verifying-key**  <br>*optional*|\[vkey\]  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|


<a name="stateprooftracking"></a>
### StateProofTracking

|Name|Description|Schema|
|---|---|---|
|**next-round**  <br>*optional*|\[n\] Next round for which we will accept a state proof transaction.|integer|
|**online-total-weight**  <br>*optional*|\[t\] The total number of microalgos held by the online accounts during the StateProof round.|integer|
|**type**  <br>*optional*|State Proof Type. Note the raw object uses map with this as key.|integer|
|**voters-commitment**  <br>*optional*|\[v\] Root of a vector commitment containing online accounts that will help sign the proof.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|


<a name="stateproofverifier"></a>
### StateProofVerifier

|Name|Description|Schema|
|---|---|---|
|**commitment**  <br>*optional*|\[cmt\] Represents the root of the vector commitment tree.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**key-lifetime**  <br>*optional*|\[lf\] Key lifetime.|integer|


<a name="stateschema"></a>
### StateSchema
Represents a \[apls\] local-state or \[apgs\] global-state schema. These schemas determine how much storage may be used in a local-state or global-state for an application. The more space used, the larger minimum balance must be maintained in the account holding the data.


|Name|Description|Schema|
|---|---|---|
|**num-byte-slice**  <br>*required*|Maximum number of TEAL byte slices that may be stored in the key/value store.|integer|
|**num-uint**  <br>*required*|Maximum number of TEAL uints that may be stored in the key/value store.|integer|


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
|**bytes**  <br>*required*|bytes value.|string|
|**type**  <br>*required*|type of the value. Value `1` refers to **bytes**, value `2` refers to **uint**|integer|
|**uint**  <br>*required*|uint value.|integer|


<a name="transaction"></a>
### Transaction
Contains all fields common to all transactions and serves as an envelope to all transactions type. Represents both regular and inner transactions.

Definition:
data/transactions/signedtxn.go : SignedTxn
data/transactions/transaction.go : Transaction


|Name|Description|Schema|
|---|---|---|
|**application-transaction**  <br>*optional*||[TransactionApplication](#transactionapplication)|
|**asset-config-transaction**  <br>*optional*||[TransactionAssetConfig](#transactionassetconfig)|
|**asset-freeze-transaction**  <br>*optional*||[TransactionAssetFreeze](#transactionassetfreeze)|
|**asset-transfer-transaction**  <br>*optional*||[TransactionAssetTransfer](#transactionassettransfer)|
|**auth-addr**  <br>*optional*|\[sgnr\] this is included with signed transactions when the signing address does not equal the sender. The backend can use this to ensure that auth addr is equal to the accounts auth addr.|string|
|**close-rewards**  <br>*optional*|\[rc\] rewards applied to close-remainder-to account.|integer|
|**closing-amount**  <br>*optional*|\[ca\] closing amount for transaction.|integer|
|**confirmed-round**  <br>*optional*|Round when the transaction was confirmed.|integer|
|**created-application-index**  <br>*optional*|Specifies an application index (ID) if an application was created with this transaction.|integer|
|**created-asset-index**  <br>*optional*|Specifies an asset index (ID) if an asset was created with this transaction.|integer|
|**fee**  <br>*required*|\[fee\] Transaction fee.|integer|
|**first-valid**  <br>*required*|\[fv\] First valid round for this transaction.|integer|
|**genesis-hash**  <br>*optional*|\[gh\] Hash of genesis block.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**genesis-id**  <br>*optional*|\[gen\] genesis block ID.|string|
|**global-state-delta**  <br>*optional*|\[gd\] Global state key/value changes for the application being executed by this transaction.|[StateDelta](#statedelta)|
|**group**  <br>*optional*|\[grp\] Base64 encoded byte array of a sha512/256 digest. When present indicates that this transaction is part of a transaction group and the value is the sha512/256 hash of the transactions in that group.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**id**  <br>*optional*|Transaction ID|string|
|**inner-txns**  <br>*optional*|Inner transactions produced by application execution.|< [Transaction](#transaction) > array|
|**intra-round-offset**  <br>*optional*|Offset into the round where this transaction was confirmed.|integer|
|**keyreg-transaction**  <br>*optional*||[TransactionKeyreg](#transactionkeyreg)|
|**last-valid**  <br>*required*|\[lv\] Last valid round for this transaction.|integer|
|**lease**  <br>*optional*|\[lx\] Base64 encoded 32-byte array. Lease enforces mutual exclusion of transactions.  If this field is nonzero, then once the transaction is confirmed, it acquires the lease identified by the (Sender, Lease) pair of the transaction until the LastValid round passes.  While this transaction possesses the lease, no other transaction specifying this lease can be confirmed.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**local-state-delta**  <br>*optional*|\[ld\] Local state key/value changes for the application being executed by this transaction.|< [AccountStateDelta](#accountstatedelta) > array|
|**logs**  <br>*optional*|\[lg\] Logs for the application being executed by this transaction.|< string (byte) > array|
|**note**  <br>*optional*|\[note\] Free form data.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**payment-transaction**  <br>*optional*||[TransactionPayment](#transactionpayment)|
|**receiver-rewards**  <br>*optional*|\[rr\] rewards applied to receiver account.|integer|
|**rekey-to**  <br>*optional*|\[rekey\] when included in a valid transaction, the accounts auth addr will be updated with this value and future signatures must be signed with the key represented by this address.|string|
|**round-time**  <br>*optional*|Time when the block this transaction is in was confirmed.|integer|
|**sender**  <br>*required*|\[snd\] Sender's address.|string|
|**sender-rewards**  <br>*optional*|\[rs\] rewards applied to sender account.|integer|
|**signature**  <br>*optional*||[TransactionSignature](#transactionsignature)|
|**state-proof-transaction**  <br>*optional*||[TransactionStateProof](#transactionstateproof)|
|**tx-type**  <br>*required*|\[type\] Indicates what type of transaction this is. Different types have different fields.<br><br>Valid types, and where their fields are stored:<br>* \[pay\] payment-transaction<br>* \[keyreg\] keyreg-transaction<br>* \[acfg\] asset-config-transaction<br>* \[axfer\] asset-transfer-transaction<br>* \[afrz\] asset-freeze-transaction<br>* \[appl\] application-transaction<br>* \[stpf\] state-proof-transaction|enum (pay, keyreg, acfg, axfer, afrz, appl, stpf)|


<a name="transactionapplication"></a>
### TransactionApplication
Fields for application transactions.

Definition:
data/transactions/application.go : ApplicationCallTxnFields


|Name|Description|Schema|
|---|---|---|
|**accounts**  <br>*optional*|\[apat\] List of accounts in addition to the sender that may be accessed from the application's approval-program and clear-state-program.|< string > array|
|**application-args**  <br>*optional*|\[apaa\] transaction specific arguments accessed from the application's approval-program and clear-state-program.|< string > array|
|**application-id**  <br>*required*|\[apid\] ID of the application being configured or empty if creating.|integer|
|**approval-program**  <br>*optional*|\[apap\] Logic executed for every application transaction, except when on-completion is set to "clear". It can read and write global state for the application, as well as account-specific local state. Approval programs may reject the transaction.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**clear-state-program**  <br>*optional*|\[apsu\] Logic executed for application transactions with on-completion set to "clear". It can read and write global state for the application, as well as account-specific local state. Clear state programs cannot reject the transaction.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**extra-program-pages**  <br>*optional*|\[epp\] specifies the additional app program len requested in pages.|integer|
|**foreign-apps**  <br>*optional*|\[apfa\] Lists the applications in addition to the application-id whose global states may be accessed by this application's approval-program and clear-state-program. The access is read-only.|< integer > array|
|**foreign-assets**  <br>*optional*|\[apas\] lists the assets whose parameters may be accessed by this application's ApprovalProgram and ClearStateProgram. The access is read-only.|< integer > array|
|**global-state-schema**  <br>*optional*||[StateSchema](#stateschema)|
|**local-state-schema**  <br>*optional*||[StateSchema](#stateschema)|
|**on-completion**  <br>*required*||[OnCompletion](#oncompletion)|


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
|**close-amount**  <br>*optional*|Number of assets transferred to the close-to account as part of the transaction.|integer|
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
|**state-proof-key**  <br>*optional*|\[sprfkey\] State proof key used in key registration transactions.  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
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


<a name="transactionstateproof"></a>
### TransactionStateProof
Fields for a state proof transaction. 

Definition:
data/transactions/stateproof.go : StateProofTxnFields


|Name|Description|Schema|
|---|---|---|
|**message**  <br>*optional*|\[spmsg\]|[IndexerStateProofMessage](#indexerstateproofmessage)|
|**state-proof**  <br>*optional*||[StateProofFields](#stateprooffields)|
|**state-proof-type**  <br>*optional*|\[sptype\] Type of the state proof. Integer representing an entry defined in protocol/stateproof.go|integer|



