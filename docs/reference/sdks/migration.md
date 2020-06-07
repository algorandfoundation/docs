title: API V2 Migration Guide

In June 2020, Algorand introduced the V2 API for `algod` and deprecated the V1 API. Both APIs remain functional to allow developers time to transition their application code to the fully supported V2 endpoints. Simultaneously, Algorand introduced `algorand-indexer` with only a V2 API. Use this guide to update your preferred SDK for V2 client support and then transition your application to use V2 clients for `algod` and `algorand-indexer` where applicable.

!!! information 
    The `kmd` API remains unchanged at V1. There are no changes required for application code using the V1 `kmd` client. 

!!! warning
    The deprecation of the V1 API for `algod` will become a breaking change in a future release of the `algod` software. Please ensure your application code is migrated to the new V2 endpoints at this time.

# Update your SDK

## JavaScript
### Update

```bash
# update using npm
npm install algosdk

# verify installed version
npm list algosdk
```
### Supported Clients by Version

| SDK Version | Supported V1 Clients | Supported V2 Clients |
| ----------- | -------------------- | -------------------- |
| thru algod@1.5.0 | `algod`, `kmd`       | n/a                  |
| from algod@1.6.1 | `kmd`                | `algod`, `algorand-indexer` |


### Supported Library Imports
<table>
    <tr>
        <th>API</td>
        <th>V1 Import</td>
        <th>V2 Import</td>
    </tr>
    <tr>
        <td>
            algod
        </td>
        <td>
            ```javascript
            const algosdk = require('algosdk'); // deprecated
            ```
        </td>
        <td>
            ```javascript
            const algosdk = require('algosdk');
            ```
        </td>
    </tr>
    <tr>
        <td>
            kmd
        </td>
        <td>
            ```javascript
            from algosdk import kmd
            ```
        </td>
        <td>
            ```javascript
            // Not implemented, use V1
            ```
        </td>
    </tr>
    <tr>
        <td>
            indexer
        </td>
        <td>
            ```javascript
            // Not implemented, use V2
            ```
        </td>
        <td>
            ```javascript
            const algosdk = require('algosdk');
            ```
        </td>
    </tr>
</table>

### Recommended Client Instantiations
<table>
    <tr>
        <th>API</td>
        <th>Client Instantiation</td>
    </tr>
    <tr>
        <td>
            algod
        </td>
        <td>
            ```javascript
            let algodClient = new algosdk.Algodv2(algod_token, algod_server, algod_port);
            ```
        </td>
    </tr>
    <tr>
        <td>
            kmd
        </td>
        <td>
            ```javascript
            let kmdClient = new algosdk.Kmd(kmd_token, kmd_server, kmd_port);
            ```
        </td>
    </tr>
    <tr>
        <td>
            indexer
        </td>
        <td>
            ```javascript
            let indexerClient = new algosdk.Indexer(indexer_token, indexer_server, indexer_port);
            ```
        </td>
    </tr>
</table>

## Python
### Update

```bash
# update using pip
pip3 install py-algorand-sdk

# verify installed version
pip3 list | grep "py-algorand-sdk"
```

### Supported Clients by Version

| SDK Version | Supported V1 Clients | Supported V2 Clients |
| ----------- | -------------------- | -------------------- |
| thru py-algorand-sdk 1.2.1 | `algod`, `kmd`       | n/a                  |
| from py-algorand-sdk 1.3.0 | `kmd`                | `algod`, `algorand-indexer` |


### Supported Library Imports
<table>
    <tr>
        <th>API</td>
        <th>V1 Import</td>
        <th>V2 Import</td>
    </tr>
    <tr>
        <td>
            algod
        </td>
        <td>
            ```python
            from algosdk import algod # deprecated
            ```
        </td>
        <td>
            ```python
            from algosdk.v2client import algod
            ```
        </td>
    </tr>
    <tr>
        <td>
            kmd
        </td>
        <td>
            ```python
            from algosdk import kmd
            ```
        </td>
        <td>
            ```python
            # Not implemented, use V1
            ```
        </td>
    </tr>
    <tr>
        <td>
            indexer
        </td>
        <td>
            ```python
            # Not implemented, use V2
            ```
        </td>
        <td>
            ```python
            from algosdk.v2client import indexer
            ```
        </td>
    </tr>
</table>

### Recommended Client Instantiations
<table>
    <tr>
        <th>API</td>
        <th>Client Instantiation</td>
    </tr>
    <tr>
        <td>
            algod
        </td>
        <td>
            ```python
            algod_client = algod.AlgodClient(algod_token, algod_address)
            ```
        </td>
    </tr>
    <tr>
        <td>
            kmd
        </td>
        <td>
            ```python
            algod_client = algod.AlgodClient(algod_token, algod_address)
            ```
        </td>
    </tr>
    <tr>
        <td>
            indexer
        </td>
        <td>
            ```python
            indexer_client = indexer.IndexerClient(indexer_token, indexer_address)
            ```
        </td>
    </tr>
</table>

## Java
### Update

```
Maven:
<dependency>
    <groupId>com.algorand</groupId>
    <artifactId>algosdk</artifactId>
    <version>1.4.0</version>
</dependency>
```
### Supported Clients by Version

| SDK Version | Supported V1 Clients | Supported V2 Clients |
| ----------- | -------------------- | -------------------- |
| thru java-algorand-sdk 1.3.1 | `algod`, `kmd`       | n/a                  |
| from java-algorand-sdk 1.4.0 | `kmd`                | `algod`, `algorand-indexer` |


### Supported Library Imports
<table>
    <tr>
        <th>API</td>
        <th>V1 Import</td>
        <th>V2 Import</td>
    </tr>
    <tr>
        <td>
            algod
        </td>
        <td>
            ```java
            import com.algorand.algosdk.algod.client.AlgodClient;
            import com.algorand.algosdk.algod.client.ApiException;
            import com.algorand.algosdk.algod.client.api.AlgodApi;
            ```
        </td>
        <td>
            ```java
            import com.algorand.algosdk.v2.client.common.AlgodClient;
            ```
        </td>
    </tr>
    <tr>
        <td>
            kmd
        </td>
        <td>
            ```java
            import com.algorand.algosdk.kmd.client.KmdClient;
            import com.algorand.algosdk.kmd.client.ApiException;
            import com.algorand.algosdk.kmd.client.api.KmdApi;
            ```
        </td>
        <td>
            ```java
            // Not implemented, use V1
            ```
        </td>
    </tr>
    <tr>
        <td>
            indexer
        </td>
        <td>
            ```java
            // Not implemented, use V2
            ```
        </td>
        <td>
            ```java
            import com.algorand.algosdk.v2.client.common.IndexerClient;
            ```
        </td>
    </tr>
</table>

### Recommended Client Instantiations
<table>
    <tr>
        <th>API</td>
        <th>Client Instantiation</td>
    </tr>
    <tr>
        <td>
            algod
        </td>
        <td>
            ```java
            let algodClient = new algosdk.Algodv2(algod_token, algod_server, algod_port);
            ```
        </td>
    </tr>
    <tr>
        <td>
            kmd
        </td>
        <td>
            ```java
            let kmdClient = new algosdk.Kmd(kmd_token, kmd_server, kmd_port);
            ```
        </td>
    </tr>
    <tr>
        <td>
            indexer
        </td>
        <td>
            ```java
            let indexerClient = new algosdk.Indexer(indexer_token, indexer_server, indexer_port);
            ```
        </td>
    </tr>
</table>

## Go
### Update

```bash
go get -u github.com/algorand/go-algorand-sdk/...
make build
```

### Supported Clients by Version

| SDK Version | Supported V1 Clients | Supported V2 Clients |
| ----------- | -------------------- | -------------------- |
| thru go-algorand-sdk 1.3.0 | `algod`, `kmd`       | n/a                  |
| from go-algorand-sdk 1.4.0 | `kmd`                | `algod`, `algorand-indexer` |


### Supported Library Imports
<table>
    <tr>
        <th>API</td>
        <th>V1 Import</td>
        <th>V2 Import</td>
    </tr>
    <tr>
        <td>
            algod
        </td>
        <td>
            ```go
            import ( 
                "github.com/algorand/go-algorand-sdk/client/algod" 
            ) // deprecated
            ```
        </td>
        <td>
            ```go
            import ( 
                "github.com/algorand/go-algorand-sdk/client/v2/algod" 
            )
            ```
        </td>
    </tr>
    <tr>
        <td>
            kmd
        </td>
        <td>
            ```go
            import ( 
                "github.com/algorand/go-algorand-sdk/client/kmd" 
            )
            ```
        </td>
        <td>
            ```go
            // Not implemented, use V1
            ```
        </td>
    </tr>
    <tr>
        <td>
            indexer
        </td>
        <td>
            ```go
            // Not implemented, use V2
            ```
        </td>
        <td>
            ```go
            import ( 
                "github.com/algorand/go-algorand-sdk/client/v2/indexer" 
            )
            ```
        </td>
    </tr>
</table>

### Recommended Client Instantiations
<table>
    <tr>
        <th>API</td>
        <th>Client Instantiation</td>
    </tr>
    <tr>
        <td>
            algod
        </td>
        <td>
            ```go
            algodClient, err := algod.MakeClient(algodAddress, algodToken)
            ```
        </td>
    </tr>
    <tr>
        <td>
            kmd
        </td>
        <td>
            ```go
            kmdClient, err := kmd.MakeClient(kmdAddress, kmdToken)
            ```
        </td>
    </tr>
    <tr>
        <td>
            indexer
        </td>
        <td>
            ```go
            indexerClient, err := indexer.MakeClient(indexerAddress, indexerToken)
            ```
        </td>
    </tr>
</table>

# Update your Application Code


