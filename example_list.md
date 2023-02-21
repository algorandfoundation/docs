The Examples are critical to have in working order and up to date.

We'll put runnable source files in the `examples` directory of the respective SDKs as well as in PyTeal and Beaker.

Goal and TEAL source location tbd
## SDKs

Show simple examples of how to use the SDK 

[Reference Page for Python](./docs/sdks/python/index.md) 
[Reference Page for JS](./docs/sdks/javascript/index.md) 
[Reference Page for Go](./docs/sdks/go/index.md) 
[Reference Page for Java](./docs/sdks/java/index.md) 


* ACCOUNT_GENERATE - Generate a new account
* ALGOD_CREATE_CLIENT - Create Algod client
* FETCH_ACCOUNT_INFO - Grab account information
* SIMPLE_PAYMENT_TRANSACTION_CREATE - Create single Payment transaction
* SIMPLE_PAYMENT_TRANSACTION_SIGN - Sign single Payment transaction
* SIMPLE_PAYMENT_TRANSACTION_SUBMIT - Submit transaction and display result


## REST APIs

We should show how to initialize a new client, set key header and do some very
simple calls against them. Refer to SDK docs or REST API docs for full functionality 

### Algod Client:
    
[Reference page](./docs/get-details/algod.md) (new page)

* ALGOD_CREATE_CLIENT - show how to init client and connect to algod and pass api keys in header
* ALGOD_USE_CLIENT - fetch some data like suggested params

### Indexer Client:

[Reference page](./docs/get-details/indexer.md)

* INDEXER_CREATE_CLIENT - show how to init client and connect to indexer and pass api keys in header
* INDEXER_USE_CLIENT - show how to get status/transactions 
* INDEXER_CLIENT_PAGINATE - show how to paginate through results

### KMD:

[Reference page](./docs/get-details/accounts/create.md)

* KMD_CREATE_CLIENT - Show how to init client and connect to kmd
* KMD_CREATE_WALLET - Show how to create a new wallet
* KMD_RECOVER_WALLET - show how to recover a wallet
* KMD_CREATE_ACCOUNT - Show how to create a new account
* KMD_IMPORT_ACCOUNT - Show how to create a new account
* KMD_EXPORT_ACCOUNT - show how to get the list of accounts in a wallet


## Accounts

[Reference Page](./docs/get-details/accounts/create.md)

* ACCOUNT_GENERATE - generate a new account with the SDK (mirrors SDK reference page)
* ACCOUNT_RECOVER_MNEMONIC - Create an account from a known mnemonic and print pubkey/privatekey
* ACCOUNT_RECOVER_PRIVATE_KEY - Create an account from a known private key and print pubkey/mnemonic
* ACCOUNT_MULTISIG_CREATE - Create a multisig account from a list of accounts
* ACCOUNT_REKEY - show how to rekey an account

## ASAs

[Reference Page](./docs/get-details/asa.md)

* ASSET_CREATE - Create an Asset from acct1 
* ASSET_OPTIN - Opt acct2 in to the newly created asset
* ASSET_XFER - Send asset from acct1 to acct2 
* ASSET_FREEZE - Freeze the asset in acct2 and then unfreeze the asset in acct2 
* ASSET_CONFIG - Reconfigure Asset to remove the freeze address 
* ASSET_CLAWBACK - Clawback Asset from acct2
* ASSET_DELETE - Delete the asset
* ASSET_INFO - Show how to query the asset info from algod

## Atomic Transactions

### No ATC

[Reference Page](./docs/get-details/atomic-transfers.md)

* ATOMIC_CREATE_TXNS - Create a couple Payment transactions
* ATOMIC_GROUP_TXNS - Group them
* ATOMIC_GROUP_DEBUG - print the group ids and txids #TODO
* ATOMIC_GROUP_SIGN - sign them
* ATOMIC_GROUP_ASSEMBLE - re-assemble signed transaction (??) 
* ATOMIC_GROUP_SEND - send them
* ATOMIC_GROUP_RESULTS - Show how to get _all_ the results back (not just the first txid) # TODO


### With ATC

[Reference Page](./docs/get-details/atc.md)

* ATC_CREATE - Create a new instance of the ATC
* ATC_CREATE_SIGNER - Show how to create a TransactionSigner
* ATC_ADD_TRANSACTION - Add a simple transaction to the ATC
* ATC_CONTRACT_INIT - Create a Contract object
* ATC_ADD_METHOD_CALL - Call a method using the Contract object 
* ATC_GROUP_TXIDS - Show how you can get the transaction group and print the group ids and txids
* ATC_RESULTS - Execute the group, Show how to get the results (esp for non method calls)

## Encoding/Decoding

We should show common encoding/decoding on [this page](./docs/get-details/encoding.md)

* CODEC_ADDRESS - Address
* CODEC_UINT64 - Uint64
* CODEC_STATE - Global/local state - how to decode, what the types mean 
* CODEC_ABI_TYPES - ABI types (using something like ABIType.from("(uint64,address)"))
* CODEC_TRANSACTION- Transactions (Signed and unsigned)
* CODEC_BLOCKS - Blocks

##  