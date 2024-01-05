title: Technical FAQ


# Protocol Limits

How many transactions can be in an atomic group?
How large can my approval program be?
How many arguments can I pass?
What is the Minimum Balance Requirement increase for an asset opt-in?

The limits applied at the protocol level are documented [here](/docs/get-details/parameter_tables).


# Address Encoding/Decoding

An address comes in 2 forms:
    1) encoded, 58 characters long, looks something like `7K5TT4US7M3FM7L3XBJXSXLJGF2WCXPBV2YZJJO2FH46VCZOS3ICJ7E4QU`
    2) decoded, 32 byte, looks something like `0xfabb39f292fb36567d7bb853795d693175615de1aeb194a5da29f9ea8b2e96d0` as hexadecimal

You can translate from one to the other by using the SDK supplied methods.

For example, in python `encoding.encode_address` will convert the 32 byte version to the encoded 58 character long version and `encoding.decode_address` will perform the opposite translation.

All SDKs have a similarly named method.

!! Note that smart contracts operate _only_ on the 32 byte version, so any interaction where an address is used should be translated prior to passing it to the smart contract. This is handled for you automatically in some cases (e.g. sender on a transaction)

[Address Details](/docs/get-details/accounts/#keys-and-addresses)
[Encoding Details](/docs/get-details/encoding/#address-encoding)

# Application State Encoding/Decoding

When calling the API for global or local state, the result returned is in the form of an array of state values. Each entry in the array represents a key/value pair in global or local state. To decode the key or bytes value, just base64 decode the string into bytes, then encode the bytes in whatever format is required.

*examples*:

- for a simple string, just encode it as `ascii` or `utf-8`
- for an address call `encode_address` or similar depending on SDK
- for an ABI tuple like `(address,uint64,bool)` use `ABIType.from_string("(address,uint64,bool)").decode(value_bytes)`


```ts
    interface StateValue {
      // The base64 encoded key for this state value
      key: string;
      // The value as a sort of union type
      value: {
        // if set, the base64 encoded bytes for the state value
        bytes: string;
        // if set, the number stored in state
        uint: number;
        // 1 for bytes, 2 for uint
        type?: number;
        // Only set in the eval deltas, describes delete/set
        action?: number;
      };
    }
```


# Deciphering Algod Errors

A 400 error typically occurs because there was some issue with the transactions. The exact reason will depend on the circumstances, but the error message will contain more information.

Common reasons include:

- `transaction already in ledger: ...`

    This happens when resending a transaction that has already been approved or accepted into a transaction pool. To make the transaction unique, add a nonce to the note field or change the valid rounds. To check if the transaction is in the pool use `/v2/transactions/pending/{txid}` REST API endpoint.

- `underflow on subtracting SEND_AMOUNT from sender amount AVAILABLE_AMOUNT`, `account ADDRESS balance 0 below min XXX (N assets)`

    All of these happen when some account does not have enough Algos to cover the transaction (taking into account their minimum balance requirement and any other transaction in the same group). Fund the account if necessary or cover the fees from its transactions with another transaction in the group.

- `receiver error: must optin, asset ASSET_ID missing from ACCOUNT_ADDR`

    This happens when the intended receiver of some asset hasn't opted in.

- `should have been authorized by ADDR1 but was actually authorized by ADDR2`

    This happens when the wrong signer is used to sign a transaction.

- `TransactionPool.Remember: txn dead: round XXXXXXXX outside of XXXXXXXX--XXXXXXXX`

    This happens when trying to submit a transaction outside its validity window, transactions may only have up to 1000 rounds between first/last valid round and the current round must be somewhere between first/last valid.

- `fee N below threshold M`

    This happens when the fee attached to a transaction (or transaction group) is not sufficient to cover the fees for the transaction.

- `only clearing out is supported for applications that have been deleted`

    This happens when an attempt is made to call an application that no longer exists

- `this transaction should be issued by the manager. It is issued by ADDR`

    This happens when trying to issue an asset config transaction by an account that is _not_ the manager account for the asset

- `cannot close asset ID in allocating account`

    This happens when the creator of an asset tries to close out of the asset.

- `check failed on ApprovalProgram: program version N greater than protocol supported version M`

    This happens when trying to create an application with a version that is not yet supported by the protocol.
    If this happens on your local sandbox, set the protocol version to `future` to get access to the latest changes.

- `cannot clear state: ADDR is not currently opted in to app N`

    This happens when an account that is not opted in to an application tries to opt out or clear state for that app.

- `invalid : program version mismatch: N != M`

    This happens when you attempt to deploy or update a programs approval and clear state code with two different versions. Check that they both have the same version number.


# Deciphering Logic errors

How can I debug this logic error?

Generally Logic errors can be debugged using [Dryrun, Tealdbg, or Simulate](/docs/get-details/dapps/smart-contracts/debugging).

For some common errors, find an explanation below:

-  `logic eval error: invalid {Asset|App|Account|Box} reference`

    This happens when the reference is not passed in the app call transactions in the appropriate reference field. By passing a reference in the transaction, the node is able to quickly load the reference into memory and have it available for the AVM runtime to consult quickly. See [Box Storage Limitations](#box-storage) for additional information.


- `logic eval error: assert failed pc=XXX`

    An `assert` was invoked on something that evaluated to 0. The `pc` will provide a pointer to where in the program the `assert` failed. To find where in the TEAL source program this corresponds to, compile the source TEAL with [`source_map`](/docs/rest-apis/algod#post-v2tealcompile) enabled and use the result to find the line in the source program.

- `logic eval error: program logs too large.`

    At most, 1k may be logged and the `log` op may be called at most 32 times in a single app call.

- `logic eval error: write budget (N) exceeded`, `logic eval error: box read budget (N) exceeded`

    Every box reference passed allots another 1k to the read/write budget for the app call. By passing more box refs a larger IO budget is available. The budget is shared across app calls within the same group so a 32k box can be read/written to as long as 32 box references are passed (8 per txn, 4 txns). See [Box Storage Limitations](#box-storage) for additional information.

- `logic eval error: store TYPE count N exceeds schema TYPE count M`

    Schema needs to be large enough to allow storage requirements. Since the schema immutable after creation, a new application must be created if more storage is required. See [Global/Local Storage Limitations](#globallocal-storage) for additional information.

- `logic eval error: err opcode executed.`

    Typically, this is the result of the internal routing of the app finding that there are no matching routes for some conditional block. This often happens when something like the ABI method selector passed is not present in the application.

- `logic eval error: overspend`

    The account listed in the error did not have enough Algos to cover the transactions and fees of the transaction.

- `logic eval error: * overflowed.`

    The math operation would have returned a value > max uint64 (2^64 - 1). A solution is to use wide math opcodes to perform the arithmetic or byte math ops like `b*`.

- `logic eval error: - would result negative.`

    The math operation would have returned a negative value (< 0). Consider checking the values before using them within the operation.

- `logic eval error: invalid ApplicationArgs index 0.`

    An attempt was made by the contract to access an Application Arguments array element that was not set. Check that the correct Application Arguments are passed and that the flow of the program does not accidentally reference an invalid index.

# Environment Setup

How do I setup the Algod to connect to sandbox?

Instructions for [Algokit](https://github.com/algorandfoundation/algokit-cli/blob/main/docs/algokit.md)

Or

Instructions for [Sandbox](https://github.com/algorand/sandbox/README.md)


# API Providers

Which API should I use?

Below is a list of API providers that can be used to query Algod or Indexer data.

- [AlgoNode](https://algonode.io/api/) (algod and indexer available, free tier available)
- [Blockdaemon](https://www.blockdaemon.com/protocols/algorand) (algod only)
- [GetBlock.io](https://getblock.io/nodes/algo/) (algod only, free tier available)

Or find others [here](https://developer.algorand.org/ecosystem-projects/?tags=api-services)

# Smart Contract Storage Limitations

When storing data on-chain you must give consideration to the different limitations between state and box storage. Be aware the note field is not an option, as that's transactional data and is not stored directly to the ledger, nor is it accessible to other application calls outside the initial transaction group.

## Global/Local Storage

[Global](/docs/get-details/dapps/smart-contracts/apps/state/#global-storage) and [Local](/docs/get-details/dapps/smart-contracts/apps/state/#local-storage) state storage can store key-value pairs up to a maximum size of 128 bytes, inclusive of the key length. For example if the key is 8 bytes, then a maximum of 120 bytes can be stored in the value. Since each key must be unique, only a single 0 byte length key can be used to store a 128 byte value. Storing larger values requires the data to be split up and stored in multiple key-value pairs within state storage.

## Box Storage

[Box storage](/docs/get-details/dapps/smart-contracts/apps/state/#box-storage) is ideal for storing larger dynamically sized data. Since the minimum balance requirement is proportional to the amount of data you allocate to the box, it's generally more cost efficient, allowing you to allocate exactly the amount you need. Keys must again be unique per application, much like state storage, however the key is not included as part of the box size and instead must be between 1 to 64 bytes. The exact size of the box value can be up to 32KB (32768 bytes), however anything over 1024 bytes is going to require additional box references. This is due to each box reference being given a 1024 bytes of operational budget. For example if you wanted to store 2000 bytes in a box named "data", you'd need to include the "data" box reference two times in your application call transaction.


# Pending Transactions

When you run a non-participation node only transactions submitted to that particular node are seen as pending in the [transaction pool](/docs/rest-apis/algod/#get-v2accountsaddresstransactionspending). If you're interested in receiving all pending transactions within the network you'll need to set `ForceFetchTransactions` to True in the [algod configuration settings](/docs/run-a-node/reference/config/#algod-configuration-settings). Note that this will increase the bandwidth used by the node.

# Verifying Release Binaries

When downloading a release, you're able to look at the accompanying signature files to verify that you've received artifacts provided by Algorand Inc. The steps to do so vary slightly by product and package type.

## algod

The algod binaries have detached signatures for every archive file. The signatures are used as follows:
1. Download public keys from https://releases.algorand.com
2. Install one or more of the public keys with `gpg --import key.pub`.
2. Download an artifact and its .sig file.
3. Verify using a gpg tool: `gpg --verify file.tar.gz.sig file.tar.gz`

## Conduit & Indexer

Instead of detached signatures, these binaries have a single checksum file which is signed. Their public key is also hosted on a keyserver to simplify the process of installing, refreshing and revoking keys as needed. They can also be downloaded and imported from the releases page outlined above.

1. Install the Algorand signing key from keys.openpgp.org or . Using GnuPG this is done with the following command:
    > gpg --keyserver keys.openpgp.org --search-keys dev-ci+build@algorand.com

2. Refresh keys if necessary. In the event of a security breech, we will revoke the key. To account for this, simply refresh the key:
    > gpg --refresh-keys --keyserver keys.openpgp.org

3. Download `checksums.txt.sig` and `checksums.txt` and verify the signature:
    > gpg --verify checksums.txt.sig checksums.txt

4. Download one or more release archives and verify the checksum:
    > sha256sum -c < checksums.txt

