title: kmd
---

<a name="paths"></a>
## Paths

<a name="swaggerhandler"></a>
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


<a name="generatekey"></a>
### POST /v1/key
Generate a key
```
POST /v1/key
```


**Description**
Generates the next key in the deterministic key sequence (as determined by the master derivation key) and adds it to the wallet, returning the public key.


**Parameters**

|Type|Name|Schema|
|---|---|---|
|**Body**|**Generate Key Request**  <br>*required*|[GenerateKeyRequest](#generatekeyrequest)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Response to `POST /v1/key`|[APIV1POSTKeyResponse](#apiv1postkeyresponse)|


**Produces**

* `application/json`


<a name="deletekey"></a>
### DELETE /v1/key
Delete a key
```
DELETE /v1/key
```


**Description**
Deletes the key with the passed public key from the wallet.


**Parameters**

|Type|Name|Schema|
|---|---|---|
|**Body**|**Delete Key Request**  <br>*required*|[DeleteKeyRequest](#deletekeyrequest)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Response to `DELETE /v1/key`|[APIV1DELETEKeyResponse](#apiv1deletekeyresponse)|


**Produces**

* `application/json`


<a name="exportkey"></a>
### POST /v1/key/export
Export a key
```
POST /v1/key/export
```


**Description**
Export the secret key associated with the passed public key.


**Parameters**

|Type|Name|Schema|
|---|---|---|
|**Body**|**Export Key Request**  <br>*required*|[ExportKeyRequest](#exportkeyrequest)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Response to `POST /v1/key/export`|[APIV1POSTKeyExportResponse](#apiv1postkeyexportresponse)|


**Produces**

* `application/json`


<a name="importkey"></a>
### POST /v1/key/import
Import a key
```
POST /v1/key/import
```


**Description**
Import an externally generated key into the wallet. Note that if you wish to back up the imported key, you must do so by backing up the entire wallet database, because imported keys were not derived from the wallet's master derivation key.


**Parameters**

|Type|Name|Schema|
|---|---|---|
|**Body**|**Import Key Request**  <br>*required*|[ImportKeyRequest](#importkeyrequest)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Response to `POST /v1/key/import`|[APIV1POSTKeyImportResponse](#apiv1postkeyimportresponse)|


**Produces**

* `application/json`


<a name="listkeysinwallet"></a>
### POST /v1/key/list
List keys in wallet
```
POST /v1/key/list
```


**Description**
Lists all of the public keys in this wallet. All of them have a stored private key.


**Parameters**

|Type|Name|Schema|
|---|---|---|
|**Body**|**List Keys Request**  <br>*required*|[ListKeysRequest](#listkeysrequest)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Response to `POST /v1/key/list`|[APIV1POSTKeyListResponse](#apiv1postkeylistresponse)|


**Produces**

* `application/json`


<a name="exportmasterkey"></a>
### POST /v1/master-key/export
Export the master derivation key from a wallet
```
POST /v1/master-key/export
```


**Description**
Export the master derivation key from the wallet. This key is a master "backup" key for the underlying wallet. With it, you can regenerate all of the wallets that have been generated with this wallet's `POST /v1/key` endpoint. This key will not allow you to recover keys imported from other wallets, however.


**Parameters**

|Type|Name|Schema|
|---|---|---|
|**Body**|**Export Master Key Request**  <br>*required*|[ExportMasterKeyRequest](#exportmasterkeyrequest)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Response to `POST /v1/master-key/export`|[APIV1POSTMasterKeyExportResponse](#apiv1postmasterkeyexportresponse)|


**Produces**

* `application/json`


<a name="deletemultisig"></a>
### DELETE /v1/multisig
Delete a multisig
```
DELETE /v1/multisig
```


**Description**
Deletes multisig preimage information for the passed address from the wallet.


**Parameters**

|Type|Name|Schema|
|---|---|---|
|**Body**|**Delete Multisig Request**  <br>*required*|[DeleteMultisigRequest](#deletemultisigrequest)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Response to POST /v1/multisig/delete|[APIV1DELETEMultisigResponse](#apiv1deletemultisigresponse)|


**Produces**

* `application/json`


<a name="exportmultisig"></a>
### POST /v1/multisig/export
Export multisig address metadata
```
POST /v1/multisig/export
```


**Description**
Given a multisig address whose preimage this wallet stores, returns the information used to generate the address, including public keys, threshold, and multisig version.


**Parameters**

|Type|Name|Schema|
|---|---|---|
|**Body**|**Export Multisig Request**  <br>*required*|[ExportMultisigRequest](#exportmultisigrequest)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Response to `POST /v1/multisig/export`|[APIV1POSTMultisigExportResponse](#apiv1postmultisigexportresponse)|


**Produces**

* `application/json`


<a name="importmultisig"></a>
### POST /v1/multisig/import
Import a multisig account
```
POST /v1/multisig/import
```


**Description**
Generates a multisig account from the passed public keys array and multisig metadata, and stores all of this in the wallet.


**Parameters**

|Type|Name|Schema|
|---|---|---|
|**Body**|**Import Multisig Request**  <br>*required*|[ImportMultisigRequest](#importmultisigrequest)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Response to `POST /v1/multisig/import`|[APIV1POSTMultisigImportResponse](#apiv1postmultisigimportresponse)|


**Produces**

* `application/json`


<a name="listmultisg"></a>
### POST /v1/multisig/list
List multisig accounts
```
POST /v1/multisig/list
```


**Description**
Lists all of the multisig accounts whose preimages this wallet stores


**Parameters**

|Type|Name|Schema|
|---|---|---|
|**Body**|**List Multisig Request**  <br>*required*|[ListMultisigRequest](#listmultisigrequest)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Response to `POST /v1/multisig/list`|[APIV1POSTMultisigListResponse](#apiv1postmultisiglistresponse)|


**Produces**

* `application/json`


<a name="signmultisigtransaction"></a>
### POST /v1/multisig/sign
Sign a multisig transaction
```
POST /v1/multisig/sign
```


**Description**
Start a multisig signature, or add a signature to a partially completed multisig signature object.


**Parameters**

|Type|Name|Schema|
|---|---|---|
|**Body**|**Sign Multisig Transaction Request**  <br>*required*|[SignMultisigRequest](#signmultisigrequest)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Response to `POST /v1/multisig/sign`|[APIV1POSTMultisigTransactionSignResponse](#apiv1postmultisigtransactionsignresponse)|


**Produces**

* `application/json`


<a name="signmultisigprogram"></a>
### POST /v1/multisig/signprogram
Sign a program for a multisig account
```
POST /v1/multisig/signprogram
```


**Description**
Start a multisig signature, or add a signature to a partially completed multisig signature object.


**Parameters**

|Type|Name|Schema|
|---|---|---|
|**Body**|**Sign Multisig Program Request**  <br>*required*|[SignProgramMultisigRequest](#signprogrammultisigrequest)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Response to `POST /v1/multisig/signdata`|[APIV1POSTMultisigProgramSignResponse](#apiv1postmultisigprogramsignresponse)|


**Produces**

* `application/json`


<a name="signprogram"></a>
### POST /v1/program/sign
Sign program
```
POST /v1/program/sign
```


**Description**
Signs the passed program with a key from the wallet, determined by the account named in the request.


**Parameters**

|Type|Name|Schema|
|---|---|---|
|**Body**|**Sign Program Request**  <br>*required*|[SignProgramRequest](#signprogramrequest)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Response to `POST /v1/data/sign`|[APIV1POSTProgramSignResponse](#apiv1postprogramsignresponse)|


**Produces**

* `application/json`


<a name="signtransaction"></a>
### POST /v1/transaction/sign
Sign a transaction
```
POST /v1/transaction/sign
```


**Description**
Signs the passed transaction with a key from the wallet, determined by the sender encoded in the transaction.


**Parameters**

|Type|Name|Schema|
|---|---|---|
|**Body**|**Sign Transaction Request**  <br>*required*|[SignTransactionRequest](#signtransactionrequest)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Response to `POST /v1/transaction/sign`|[APIV1POSTTransactionSignResponse](#apiv1posttransactionsignresponse)|


**Produces**

* `application/json`


<a name="createwallet"></a>
### POST /v1/wallet
Create a wallet
```
POST /v1/wallet
```


**Description**
Create a new wallet (collection of keys) with the given parameters.


**Parameters**

|Type|Name|Schema|
|---|---|---|
|**Body**|**Create Wallet Request**  <br>*required*|[CreateWalletRequest](#createwalletrequest)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Response to `POST /v1/wallet`|[APIV1POSTWalletResponse](#apiv1postwalletresponse)|


**Produces**

* `application/json`


<a name="getwalletinfo"></a>
### POST /v1/wallet/info
Get wallet info
```
POST /v1/wallet/info
```


**Description**
Returns information about the wallet associated with the passed wallet handle token. Additionally returns expiration information about the token itself.


**Parameters**

|Type|Name|Schema|
|---|---|---|
|**Body**|**Get Wallet Info Request**  <br>*required*|[WalletInfoRequest](#walletinforequest)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Response to `POST /v1/wallet/info`|[APIV1POSTWalletInfoResponse](#apiv1postwalletinforesponse)|


**Produces**

* `application/json`


<a name="initwallethandletoken"></a>
### POST /v1/wallet/init
Initialize a wallet handle token
```
POST /v1/wallet/init
```


**Description**
Unlock the wallet and return a wallet handle token that can be used for subsequent operations. These tokens expire periodically and must be renewed. You can `POST` the token to `/v1/wallet/info` to see how much time remains until expiration, and renew it with `/v1/wallet/renew`. When you're done, you can invalidate the token with `/v1/wallet/release`.


**Parameters**

|Type|Name|Schema|
|---|---|---|
|**Body**|**Initialize Wallet Handle Token Request**  <br>*required*|[InitWalletHandleTokenRequest](#initwallethandletokenrequest)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Response to `POST /v1/wallet/init`|[APIV1POSTWalletInitResponse](#apiv1postwalletinitresponse)|


**Produces**

* `application/json`


<a name="releasewallethandletoken"></a>
### POST /v1/wallet/release
Release a wallet handle token
```
POST /v1/wallet/release
```


**Description**
Invalidate the passed wallet handle token, making it invalid for use in subsequent requests.


**Parameters**

|Type|Name|Schema|
|---|---|---|
|**Body**|**Release Wallet Handle Token Request**  <br>*required*|[ReleaseWalletHandleTokenRequest](#releasewallethandletokenrequest)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Response to `POST /v1/wallet/release`|[APIV1POSTWalletReleaseResponse](#apiv1postwalletreleaseresponse)|


**Produces**

* `application/json`


<a name="renamewallet"></a>
### POST /v1/wallet/rename
Rename a wallet
```
POST /v1/wallet/rename
```


**Description**
Rename the underlying wallet to something else


**Parameters**

|Type|Name|Schema|
|---|---|---|
|**Body**|**Rename Wallet Request**  <br>*required*|[RenameWalletRequest](#renamewalletrequest)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Response to `POST /v1/wallet/rename`|[APIV1POSTWalletRenameResponse](#apiv1postwalletrenameresponse)|


**Produces**

* `application/json`


<a name="renewwallethandletoken"></a>
### POST /v1/wallet/renew
Renew a wallet handle token
```
POST /v1/wallet/renew
```


**Description**
Renew a wallet handle token, increasing its expiration duration to its initial value


**Parameters**

|Type|Name|Schema|
|---|---|---|
|**Body**|**Renew Wallet Handle Token Request**  <br>*required*|[RenewWalletHandleTokenRequest](#renewwallethandletokenrequest)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Response `POST /v1/wallet/renew`|[APIV1POSTWalletRenewResponse](#apiv1postwalletrenewresponse)|


**Produces**

* `application/json`


<a name="listwallets"></a>
### GET /v1/wallets
List wallets
```
GET /v1/wallets
```


**Description**
Lists all of the wallets that kmd is aware of.


**Parameters**

|Type|Name|Schema|
|---|---|---|
|**Body**|**List Wallet Request**  <br>*optional*|[ListWalletsRequest](#listwalletsrequest)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Response to `GET /v1/wallets`|[APIV1GETWalletsResponse](#apiv1getwalletsresponse)|


**Produces**

* `application/json`


<a name="getversion"></a>
### GET /versions
Retrieves the current version
```
GET /versions
```


**Parameters**

|Type|Name|Schema|
|---|---|---|
|**Body**|**Versions Request**  <br>*optional*|[VersionsRequest](#versionsrequest)|


**Responses**

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Response to `GET /versions`|[VersionsResponse](#versionsresponse)|


**Produces**

* `application/json`




<a name="definitions"></a>
## Definitions

<a name="apiv1deletekeyresponse"></a>
### APIV1DELETEKeyResponse
APIV1DELETEKeyResponse is the response to `DELETE /v1/key`
friendly:DeleteKeyResponse


|Name|Schema|
|---|---|
|**error**  <br>*optional*|boolean|
|**message**  <br>*optional*|string|


<a name="apiv1deletemultisigresponse"></a>
### APIV1DELETEMultisigResponse
APIV1DELETEMultisigResponse is the response to POST /v1/multisig/delete`
friendly:DeleteMultisigResponse


|Name|Schema|
|---|---|
|**error**  <br>*optional*|boolean|
|**message**  <br>*optional*|string|


<a name="apiv1getwalletsresponse"></a>
### APIV1GETWalletsResponse
APIV1GETWalletsResponse is the response to `GET /v1/wallets`
friendly:ListWalletsResponse


|Name|Schema|
|---|---|
|**error**  <br>*optional*|boolean|
|**message**  <br>*optional*|string|
|**wallets**  <br>*optional*|< [APIV1Wallet](#apiv1wallet) > array|


<a name="apiv1postkeyexportresponse"></a>
### APIV1POSTKeyExportResponse
APIV1POSTKeyExportResponse is the response to `POST /v1/key/export`
friendly:ExportKeyResponse


|Name|Schema|
|---|---|
|**error**  <br>*optional*|boolean|
|**message**  <br>*optional*|string|
|**private_key**  <br>*optional*|[PrivateKey](#privatekey)|


<a name="apiv1postkeyimportresponse"></a>
### APIV1POSTKeyImportResponse
APIV1POSTKeyImportResponse is the response to `POST /v1/key/import`
friendly:ImportKeyResponse


|Name|Schema|
|---|---|
|**address**  <br>*optional*|string|
|**error**  <br>*optional*|boolean|
|**message**  <br>*optional*|string|


<a name="apiv1postkeylistresponse"></a>
### APIV1POSTKeyListResponse
APIV1POSTKeyListResponse is the response to `POST /v1/key/list`
friendly:ListKeysResponse


|Name|Schema|
|---|---|
|**addresses**  <br>*optional*|< string > array|
|**error**  <br>*optional*|boolean|
|**message**  <br>*optional*|string|


<a name="apiv1postkeyresponse"></a>
### APIV1POSTKeyResponse
APIV1POSTKeyResponse is the response to `POST /v1/key`
friendly:GenerateKeyResponse


|Name|Schema|
|---|---|
|**address**  <br>*optional*|string|
|**error**  <br>*optional*|boolean|
|**message**  <br>*optional*|string|


<a name="apiv1postmasterkeyexportresponse"></a>
### APIV1POSTMasterKeyExportResponse
APIV1POSTMasterKeyExportResponse is the response to `POST /v1/master-key/export`
friendly:ExportMasterKeyResponse


|Name|Schema|
|---|---|
|**error**  <br>*optional*|boolean|
|**master_derivation_key**  <br>*optional*|[MasterDerivationKey](#masterderivationkey)|
|**message**  <br>*optional*|string|


<a name="apiv1postmultisigexportresponse"></a>
### APIV1POSTMultisigExportResponse
APIV1POSTMultisigExportResponse is the response to `POST /v1/multisig/export`
friendly:ExportMultisigResponse


|Name|Schema|
|---|---|
|**error**  <br>*optional*|boolean|
|**message**  <br>*optional*|string|
|**multisig_version**  <br>*optional*|integer (uint8)|
|**pks**  <br>*optional*|< [PublicKey](#publickey) > array|
|**threshold**  <br>*optional*|integer (uint8)|


<a name="apiv1postmultisigimportresponse"></a>
### APIV1POSTMultisigImportResponse
APIV1POSTMultisigImportResponse is the response to `POST /v1/multisig/import`
friendly:ImportMultisigResponse


|Name|Schema|
|---|---|
|**address**  <br>*optional*|string|
|**error**  <br>*optional*|boolean|
|**message**  <br>*optional*|string|


<a name="apiv1postmultisiglistresponse"></a>
### APIV1POSTMultisigListResponse
APIV1POSTMultisigListResponse is the response to `POST /v1/multisig/list`
friendly:ListMultisigResponse


|Name|Schema|
|---|---|
|**addresses**  <br>*optional*|< string > array|
|**error**  <br>*optional*|boolean|
|**message**  <br>*optional*|string|


<a name="apiv1postmultisigprogramsignresponse"></a>
### APIV1POSTMultisigProgramSignResponse
APIV1POSTMultisigProgramSignResponse is the response to `POST /v1/multisig/signdata`
friendly:SignProgramMultisigResponse


|Name|Description|Schema|
|---|---|---|
|**error**  <br>*optional*||boolean|
|**message**  <br>*optional*||string|
|**multisig**  <br>*optional*|**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|


<a name="apiv1postmultisigtransactionsignresponse"></a>
### APIV1POSTMultisigTransactionSignResponse
APIV1POSTMultisigTransactionSignResponse is the response to `POST /v1/multisig/sign`
friendly:SignMultisigResponse


|Name|Description|Schema|
|---|---|---|
|**error**  <br>*optional*||boolean|
|**message**  <br>*optional*||string|
|**multisig**  <br>*optional*|**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|


<a name="apiv1postprogramsignresponse"></a>
### APIV1POSTProgramSignResponse
APIV1POSTProgramSignResponse is the response to `POST /v1/data/sign`
friendly:SignProgramResponse


|Name|Description|Schema|
|---|---|---|
|**error**  <br>*optional*||boolean|
|**message**  <br>*optional*||string|
|**sig**  <br>*optional*|**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|


<a name="apiv1posttransactionsignresponse"></a>
### APIV1POSTTransactionSignResponse
APIV1POSTTransactionSignResponse is the response to `POST /v1/transaction/sign`
friendly:SignTransactionResponse


|Name|Description|Schema|
|---|---|---|
|**error**  <br>*optional*||boolean|
|**message**  <br>*optional*||string|
|**signed_transaction**  <br>*optional*|**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|


<a name="apiv1postwalletinforesponse"></a>
### APIV1POSTWalletInfoResponse
APIV1POSTWalletInfoResponse is the response to `POST /v1/wallet/info`
friendly:WalletInfoResponse


|Name|Schema|
|---|---|
|**error**  <br>*optional*|boolean|
|**message**  <br>*optional*|string|
|**wallet_handle**  <br>*optional*|[APIV1WalletHandle](#apiv1wallethandle)|


<a name="apiv1postwalletinitresponse"></a>
### APIV1POSTWalletInitResponse
APIV1POSTWalletInitResponse is the response to `POST /v1/wallet/init`
friendly:InitWalletHandleTokenResponse


|Name|Schema|
|---|---|
|**error**  <br>*optional*|boolean|
|**message**  <br>*optional*|string|
|**wallet_handle_token**  <br>*optional*|string|


<a name="apiv1postwalletreleaseresponse"></a>
### APIV1POSTWalletReleaseResponse
APIV1POSTWalletReleaseResponse is the response to `POST /v1/wallet/release`
friendly:ReleaseWalletHandleTokenResponse


|Name|Schema|
|---|---|
|**error**  <br>*optional*|boolean|
|**message**  <br>*optional*|string|


<a name="apiv1postwalletrenameresponse"></a>
### APIV1POSTWalletRenameResponse
APIV1POSTWalletRenameResponse is the response to `POST /v1/wallet/rename`
friendly:RenameWalletResponse


|Name|Schema|
|---|---|
|**error**  <br>*optional*|boolean|
|**message**  <br>*optional*|string|
|**wallet**  <br>*optional*|[APIV1Wallet](#apiv1wallet)|


<a name="apiv1postwalletrenewresponse"></a>
### APIV1POSTWalletRenewResponse
APIV1POSTWalletRenewResponse is the response to `POST /v1/wallet/renew`
friendly:RenewWalletHandleTokenResponse


|Name|Schema|
|---|---|
|**error**  <br>*optional*|boolean|
|**message**  <br>*optional*|string|
|**wallet_handle**  <br>*optional*|[APIV1WalletHandle](#apiv1wallethandle)|


<a name="apiv1postwalletresponse"></a>
### APIV1POSTWalletResponse
APIV1POSTWalletResponse is the response to `POST /v1/wallet`
friendly:CreateWalletResponse


|Name|Schema|
|---|---|
|**error**  <br>*optional*|boolean|
|**message**  <br>*optional*|string|
|**wallet**  <br>*optional*|[APIV1Wallet](#apiv1wallet)|


<a name="apiv1wallet"></a>
### APIV1Wallet
APIV1Wallet is the API's representation of a wallet


|Name|Schema|
|---|---|
|**driver_name**  <br>*optional*|string|
|**driver_version**  <br>*optional*|integer (uint32)|
|**id**  <br>*optional*|string|
|**mnemonic_ux**  <br>*optional*|boolean|
|**name**  <br>*optional*|string|
|**supported_txs**  <br>*optional*|< [TxType](#txtype) > array|


<a name="apiv1wallethandle"></a>
### APIV1WalletHandle
APIV1WalletHandle includes the wallet the handle corresponds to
and the number of number of seconds to expiration


|Name|Schema|
|---|---|
|**expires_seconds**  <br>*optional*|integer (int64)|
|**wallet**  <br>*optional*|[APIV1Wallet](#apiv1wallet)|


<a name="createwalletrequest"></a>
### CreateWalletRequest
APIV1POSTWalletRequest is the request for `POST /v1/wallet`


|Name|Schema|
|---|---|
|**master_derivation_key**  <br>*optional*|[MasterDerivationKey](#masterderivationkey)|
|**wallet_driver_name**  <br>*optional*|string|
|**wallet_name**  <br>*optional*|string|
|**wallet_password**  <br>*optional*|string|


<a name="deletekeyrequest"></a>
### DeleteKeyRequest
APIV1DELETEKeyRequest is the request for `DELETE /v1/key`


|Name|Schema|
|---|---|
|**address**  <br>*optional*|string|
|**wallet_handle_token**  <br>*optional*|string|
|**wallet_password**  <br>*optional*|string|


<a name="deletemultisigrequest"></a>
### DeleteMultisigRequest
APIV1DELETEMultisigRequest is the request for `DELETE /v1/multisig`


|Name|Schema|
|---|---|
|**address**  <br>*optional*|string|
|**wallet_handle_token**  <br>*optional*|string|
|**wallet_password**  <br>*optional*|string|


<a name="digest"></a>
### Digest
*Type* : < integer (uint8) > array


<a name="exportkeyrequest"></a>
### ExportKeyRequest
APIV1POSTKeyExportRequest is the request for `POST /v1/key/export`


|Name|Schema|
|---|---|
|**address**  <br>*optional*|string|
|**wallet_handle_token**  <br>*optional*|string|
|**wallet_password**  <br>*optional*|string|


<a name="exportmasterkeyrequest"></a>
### ExportMasterKeyRequest
APIV1POSTMasterKeyExportRequest is the request for `POST /v1/master-key/export`


|Name|Schema|
|---|---|
|**wallet_handle_token**  <br>*optional*|string|
|**wallet_password**  <br>*optional*|string|


<a name="exportmultisigrequest"></a>
### ExportMultisigRequest
APIV1POSTMultisigExportRequest is the request for `POST /v1/multisig/export`


|Name|Schema|
|---|---|
|**address**  <br>*optional*|string|
|**wallet_handle_token**  <br>*optional*|string|


<a name="generatekeyrequest"></a>
### GenerateKeyRequest
APIV1POSTKeyRequest is the request for `POST /v1/key`


|Name|Schema|
|---|---|
|**display_mnemonic**  <br>*optional*|boolean|
|**wallet_handle_token**  <br>*optional*|string|


<a name="importkeyrequest"></a>
### ImportKeyRequest
APIV1POSTKeyImportRequest is the request for `POST /v1/key/import`


|Name|Schema|
|---|---|
|**private_key**  <br>*optional*|[PrivateKey](#privatekey)|
|**wallet_handle_token**  <br>*optional*|string|


<a name="importmultisigrequest"></a>
### ImportMultisigRequest
APIV1POSTMultisigImportRequest is the request for `POST /v1/multisig/import`


|Name|Schema|
|---|---|
|**multisig_version**  <br>*optional*|integer (uint8)|
|**pks**  <br>*optional*|< [PublicKey](#publickey) > array|
|**threshold**  <br>*optional*|integer (uint8)|
|**wallet_handle_token**  <br>*optional*|string|


<a name="initwallethandletokenrequest"></a>
### InitWalletHandleTokenRequest
APIV1POSTWalletInitRequest is the request for `POST /v1/wallet/init`


|Name|Schema|
|---|---|
|**wallet_id**  <br>*optional*|string|
|**wallet_password**  <br>*optional*|string|


<a name="listkeysrequest"></a>
### ListKeysRequest
APIV1POSTKeyListRequest is the request for `POST /v1/key/list`


|Name|Schema|
|---|---|
|**wallet_handle_token**  <br>*optional*|string|


<a name="listmultisigrequest"></a>
### ListMultisigRequest
APIV1POSTMultisigListRequest is the request for `POST /v1/multisig/list`


|Name|Schema|
|---|---|
|**wallet_handle_token**  <br>*optional*|string|


<a name="listwalletsrequest"></a>
### ListWalletsRequest
APIV1GETWalletsRequest is the request for `GET /v1/wallets`

*Type* : object


<a name="masterderivationkey"></a>
### MasterDerivationKey
MasterDerivationKey is used to derive ed25519 keys for use in wallets

*Type* : < integer (uint8) > array


<a name="multisigsig"></a>
### MultisigSig
MultisigSig is the structure that holds multiple Subsigs


|Name|Schema|
|---|---|
|**Subsigs**  <br>*optional*|< [MultisigSubsig](#multisigsubsig) > array|
|**Threshold**  <br>*optional*|integer (uint8)|
|**Version**  <br>*optional*|integer (uint8)|


<a name="multisigsubsig"></a>
### MultisigSubsig
MultisigSubsig is a struct that holds a pair of public key and signatures
signatures may be empty


|Name|Schema|
|---|---|
|**Key**  <br>*optional*|[PublicKey](#publickey)|
|**Sig**  <br>*optional*|[Signature](#signature)|


<a name="privatekey"></a>
### PrivateKey
*Type* : < integer (uint8) > array


<a name="publickey"></a>
### PublicKey
*Type* : < integer (uint8) > array


<a name="releasewallethandletokenrequest"></a>
### ReleaseWalletHandleTokenRequest
APIV1POSTWalletReleaseRequest is the request for `POST /v1/wallet/release`


|Name|Schema|
|---|---|
|**wallet_handle_token**  <br>*optional*|string|


<a name="renamewalletrequest"></a>
### RenameWalletRequest
APIV1POSTWalletRenameRequest is the request for `POST /v1/wallet/rename`


|Name|Schema|
|---|---|
|**wallet_id**  <br>*optional*|string|
|**wallet_name**  <br>*optional*|string|
|**wallet_password**  <br>*optional*|string|


<a name="renewwallethandletokenrequest"></a>
### RenewWalletHandleTokenRequest
APIV1POSTWalletRenewRequest is the request for `POST /v1/wallet/renew`


|Name|Schema|
|---|---|
|**wallet_handle_token**  <br>*optional*|string|


<a name="signmultisigrequest"></a>
### SignMultisigRequest
APIV1POSTMultisigTransactionSignRequest is the request for `POST /v1/multisig/sign`


|Name|Description|Schema|
|---|---|---|
|**partial_multisig**  <br>*optional*||[MultisigSig](#multisigsig)|
|**public_key**  <br>*optional*||[PublicKey](#publickey)|
|**signer**  <br>*optional*||[Digest](#digest)|
|**transaction**  <br>*optional*|**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**wallet_handle_token**  <br>*optional*||string|
|**wallet_password**  <br>*optional*||string|


<a name="signprogrammultisigrequest"></a>
### SignProgramMultisigRequest
APIV1POSTMultisigProgramSignRequest is the request for `POST /v1/multisig/signprogram`


|Name|Description|Schema|
|---|---|---|
|**address**  <br>*optional*||string|
|**data**  <br>*optional*|**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**partial_multisig**  <br>*optional*||[MultisigSig](#multisigsig)|
|**public_key**  <br>*optional*||[PublicKey](#publickey)|
|**wallet_handle_token**  <br>*optional*||string|
|**wallet_password**  <br>*optional*||string|


<a name="signprogramrequest"></a>
### SignProgramRequest
APIV1POSTProgramSignRequest is the request for `POST /v1/program/sign`


|Name|Description|Schema|
|---|---|---|
|**address**  <br>*optional*||string|
|**data**  <br>*optional*|**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**wallet_handle_token**  <br>*optional*||string|
|**wallet_password**  <br>*optional*||string|


<a name="signtransactionrequest"></a>
### SignTransactionRequest
APIV1POSTTransactionSignRequest is the request for `POST /v1/transaction/sign`


|Name|Description|Schema|
|---|---|---|
|**public_key**  <br>*optional*||[PublicKey](#publickey)|
|**transaction**  <br>*optional*|Base64 encoding of msgpack encoding of a `Transaction` object<br>Note: SDK and goal usually generate `SignedTxn` objects<br>in that case, the field `txn` / `Transaction` of the<br>generated `SignedTxn` object needs to be used  <br>**Pattern** : `"^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==\|[A-Za-z0-9+/]{3}=)?$"`|string (byte)|
|**wallet_handle_token**  <br>*optional*||string|
|**wallet_password**  <br>*optional*||string|


<a name="signature"></a>
### Signature
*Type* : < integer (uint8) > array


<a name="txtype"></a>
### TxType
TxType is the type of the transaction written to the ledger

*Type* : string


<a name="versionsrequest"></a>
### VersionsRequest
VersionsRequest is the request for `GET /versions`

*Type* : object


<a name="versionsresponse"></a>
### VersionsResponse
VersionsResponse is the response to `GET /versions`
friendly:VersionsResponse


|Name|Schema|
|---|---|
|**versions**  <br>*optional*|< string > array|


<a name="walletinforequest"></a>
### WalletInfoRequest
APIV1POSTWalletInfoRequest is the request for `POST /v1/wallet/info`


|Name|Schema|
|---|---|
|**wallet_handle_token**  <br>*optional*|string|


<a name="ed25519privatekey"></a>
### ed25519PrivateKey
*Type* : < integer (uint8) > array


<a name="ed25519publickey"></a>
### ed25519PublicKey
*Type* : < integer (uint8) > array


<a name="ed25519signature"></a>
### ed25519Signature
*Type* : < integer (uint8) > array



