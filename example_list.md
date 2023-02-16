The Examples are critical to have in working order and up to date.

We'll put runnable source files in the `examples` directory of the respective SDKs as well as in PyTeal and Beaker.

Goal and TEAL source location tbd

## REST APIs

We should show how to initialize a new client, set key header and do some very
simple calls against them. Refer to SDK docs or REST API docs for full functionality 

### Algod Client:
    
    show how to init client and connect to algod
    show how to set different api keys
    show how to get status/suggested params

### Indexer Client:

[Reference page](./docs/get-details/indexer.md)

    show how to init client and connect to indexer 
    show how to set different api keys
    show how to get status/transactions 
    show how to paginate through results

### KMD:

    Show how to init client and connect to kmd
    Show how to get a wallet handle 
    Show how to grab accounts using the handle



<!-- ===PYSDK_CREATE_ALGOD_CLIENT=== -->

<!-- ===PYSDK_CREATE_ALGOD_CLIENT=== -->

## Accounts

We should show how to create, save, restore accounts on [this page](./docs/get-details/accounts/create.md)

    Create a new account with the SDK
    Create an account from a known mnemonic and print pubkey/privatekey
    Create an account from a known private key and print pubkey/mnemonic

## ASAs

We should show how to do all the things with ASAs on [this page](./docs/get-details/asa.md)

### All the things

    Create an Asset from acct1 
    Opt acct2 in to the newly created asset
    Send asset from acct1 to acct2 
    Freeze the asset in acct2 
    Unfreeze the asset in acct2 
    Reconfigure Asset to remove the freeze address 
    Clawback Asset from acct2
    Delete the asset


## Atomic Transactions

### No ATC

We should show how to group transactions on [this page](./docs/get-details/atomic-transfers.md) without using the ATC

    Create a couple Payment transactions
    Group them     
    Print the group ids and txids
    Sign them
    Send them
    Show how to get _all_ the results back (not just the first txid)


### With ATC

We should show how to group transactions on [this page](./docs/get-details/atc.md) using the ATC 

    Create a contract object
    Create a couple TransactionWithSigners (1 payment, 1 method call using the contract object)
    Show how you can get the transaction group and print the group ids and txids
    Execute the group 
    Show how to get the results (esp for non method calls)

## Encoding/Decoding

We should show common encoding/decoding on [this page](./docs/get-details/encoding.md)

    Address
    Uint64
    Global/local state - how to decode, what the types mean 
    ABI types (using something like ABIType.from("(uint64,address)"))
    Transactions (Signed and unsigned)
    Blocks

## 