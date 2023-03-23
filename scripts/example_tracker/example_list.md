The Examples are critical to have in working order and up to date.

Runnable source files should be placed in the `examples` directory of the respective SDKs as well as in PyTeal and Beaker.

## SDKs

- ACCOUNT_GENERATE - Generate a new account
- ALGOD_CREATE_CLIENT - Create Algod client
- ALGOD_FETCH_ACCOUNT_INFO - Grab account information
- TRANSACTION_PAYMENT_CREATE - Create single Payment transaction
- TRANSACTION_PAYMENT_SIGN - Sign single Payment transaction
- TRANSACTION_PAYMENT_SUBMIT - Submit transaction and display result
- TRANSACTION_FEE_OVERRIDE - show how to override the fee
- SP_MIN_FEE - Show that the suggested params contain the min fee
- CONST_MIN_FEE - show that the SDK has consts that can be used

## REST APIs
### Algod Client:

- ALGOD_CREATE_CLIENT - show how to init client and connect to algod and pass api keys in header
- ALGOD_USE_CLIENT - fetch some data like suggested params

### Indexer Client:

- CREATE_INDEXER_CLIENT - show how to init client and connect to indexer and pass api keys in header ( TODO: rename to INDEXER_CREATE_CLIENT )
- INDEXER_LOOKUP_ASSET - Find asset by id, print details
- INDEXER_PAGINATE_RESULTS - Iterate over pages in results
- INDEXER_PREFIX_SEARCH - Show searching by note prefix
- INDEXER_SEARCH_MIN_AMOUNT - Show searching for accounts with min amount

### KMD:

- KMD_CREATE_CLIENT - Show how to init client and connect to kmd
- KMD_CREATE_WALLET - Show how to create a new wallet
- KMD_RECOVER_WALLET - show how to recover a wallet
- KMD_CREATE_ACCOUNT - Show how to create a new account
- KMD_IMPORT_ACCOUNT - Show how to create a new account
- KMD_EXPORT_ACCOUNT - show how to get the list of accounts in a wallet

## Accounts

- ACCOUNT_GENERATE - generate a new account with the SDK (mirrors SDK reference page)
- ACCOUNT_RECOVER_MNEMONIC - Create an account from a known mnemonic and print pubkey/privatekey
- ACCOUNT_RECOVER_PRIVATE_KEY - Create an account from a known private key and print pubkey/mnemonic
- ACCOUNT_MULTISIG_CREATE - Create a multisig account from a list of accounts
- ACCOUNT_REKEY - show how to rekey an account

### Multisig Accounts

- MULTISIG_CREATE - Create a multisig from 3 accounts
- MULTISIG_SIGN - Sign a transaction with the multisig 

### Lsig Accounts

- LSIG_COMPILE - Read a TEAL file and compile it
- LSIG_INIT - Create a LogicSig object from the compiled TEAL
- LSIG_PASS_ARGS - Pass arguments to the LogicSig
- LSIG_SIGN_FULL - Sign a transaction with the LogicSig
- LSIG_DELEGATE_FULL - Delegate the LogicSig to another account

## ASAs

- ASSET_CREATE - Create an Asset from acct1
- ASSET_OPTIN - Opt acct2 in to the newly created asset
- ASSET_XFER - Send asset from acct1 to acct2
- ASSET_FREEZE - Freeze the asset in acct2 and then unfreeze the asset in acct2
- ASSET_CONFIG - Reconfigure Asset to remove the freeze address
- ASSET_CLAWBACK - Clawback Asset from acct2
- ASSET_OPT_OUT - Opt out of the asset
- ASSET_DELETE - Delete the asset
- ASSET_INFO - Show how to query the asset info from algod

## Applications

- APP_SCHEMA - Show setting schema for the app
- APP_SOURCE - Show reading teal file
- APP_COMPILE - Show compile teal source
- APP_CREATE - App call to create the app
- APP_CALL - App call to call the app
- APP_READ_STATE - Read state for the app from the algod
- APP_OPTIN - Opt in to the app
- APP_CLEAR - Clear state for the app
- APP_CLOSEOUT - Close out of the app
- APP_NOOP - Noop call to the app
- APP_UPDATE - Update the app source
- APP_DELETE - Delete the app

## Atomic Transactions

### No ATC

- ATOMIC_CREATE_TXNS - Create a couple Payment transactions
- ATOMIC_GROUP_TXNS - Group them
- ATOMIC_GROUP_DEBUG - print the group ids and txids #TODO
- ATOMIC_GROUP_SIGN - sign them
- ATOMIC_GROUP_ASSEMBLE - re-assemble signed transaction (??)
- ATOMIC_GROUP_SEND - send them
- ATOMIC*GROUP_RESULTS - Show how to get \_all* the results back (not just the first txid) # TODO

### With ATC

- ATC_CREATE - Create a new instance of the ATC
- ATC_CREATE_SIGNER - Show how to create a TransactionSigner
- ATC_ADD_TRANSACTION - Add a simple transaction to the ATC
- ATC_CONTRACT_INIT - Create a Contract object
- ATC_ADD_METHOD_CALL - Call a method using the Contract object
- ATC_GROUP_TXIDS - Show how you can get the transaction group and print the group ids and txids
- ATC_RESULTS - Execute the group, Show how to get the results (esp for non method calls)
- ATC_BOX_REF - Show how to set a reference to the box

## Encoding/Decoding

- CODEC_ADDRESS - Address
- CODEC_BASE64 - arbitrary byte slices encoded as b64
- CODEC_UINT64 - Uint64
- CODEC_STATE - Global/local state - how to decode, what the types mean
- CODEC_ABI_TYPES - ABI types (using something like ABIType.from("(uint64,address)"))
- CODEC_TRANSACTION_UNSIGNED - Show how to read/write unsigned transactions
- CODEC_TRANSACTION_SIGNED- Show how to read/write signed transactions
- CODEC_ABI - Show how to encode/decode from an ABI type
- CODEC_BLOCKS - Blocks [not implemented yet]

## Transactions

- TRANSACTION_SIGN_OFFLINE - Show how to create and sign a transaction using cold wallet
- TRANSACTION_KEYREG_OFFLINE_CREATE - Create a keyreg txn to take an account offline
- TRANSACTION_KEYREG_ONLINE_CREATE - Create a keyreg txn to take an account online

## Debugging

- DEBUG_DRYRUN_DUMP - Create a file with a dryrun request object encoded as msgpack
- DEBUG_DRYRUN_SUBMIT - Create a dryrun request and submit it to the algod
