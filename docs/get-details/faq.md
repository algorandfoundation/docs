title: FAQ

# Encode address
 
How do I decode/encode an address?

    An address comes in 2 forms:
        1) encoded, 58 byte, looks something like `7K5TT4US7M3FM7L3XBJXSXLJGF2WCXPBV2YZJJO2FH46VCZOS3ICJ7E4QU`
        2) decoded, 32 byte, looks something like `0xfabb39f292fb36567d7bb853795d693175615de1aeb194a5da29f9ea8b2e96d0` as hexadecimal

    You can translate from one to the other by using the SDK supplied methods. 
    
    For example, in python `encoding.encode_address` will convert the 32 byte version to the encoded 58 byte version and `encoding.decode_address` will perform the opposite translation.

    All SDKs have a similarly named method.

    !! Note that smart contracts operate _only_ on the 32 byte version, so any interaction where an address is used should be translated prior to passing it to the smart contract. This is handled for you automatically in some cases (e.g. sender on a transaction)

    Resources:
    https://developer.algorand.org/docs/get-details/accounts/#keys-and-addresses
    https://developer.algorand.org/docs/get-details/encoding/#address


# Decode state

How do I decode state values?

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


# What is this 400 error?

A 400 error typically occurs because there was some issue with the transactions. The exact reason will depend on the circumstances but the error message will contain more information.

Common reasons include:

- `transaction already in ledger: ...`

    This happens when resending a transaction that has already been approved. To make the transaction unique, add a nonce to the notefield or change the valid rounds.

- `underflow on subtracting SEND_AMOUNT from sender amount AVAILABLE_AMOUNT`, `account ADDRESS balance 0 below min XXX (N assets)`

    All of these happen when some account does not have enough algos to cover the transaction (taking into account their minimum balance requirement and any other transaction in the same group). Fund the account if necessary or cover the fees from its transactions with another transaction in the group.

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

    This happens when you attempt to deploy or update a programs approval and clearstate code with two different versions. Check that they both have the same version number.


# Logic errors

How can I debug this logic error?

Generally Logic errors can be debugged using [Dryrun, Tealdbg, or Simulate](https://developer.algorand.org/docs/get-details/dapps/smart-contracts/debugging).

For some common errors, find an explanation below:  

-  `logic eval error: invalid {Asset|App|Account|Box} reference` 

    This happens when the reference is not passed in the app call transactions in the appropriate reference field. By passing a reference in the transaction, the node is able to quickly load the reference into memory and have it available for the AVM runtime to consult quickly.


- `logic eval error: assert failed pc=XXX`

    An assert was called on something that evaluated to 0. The `pc` will provide a pointer to where in the program the assert failed. To find where in the TEAL source program this corresponds to, compile the source TEAL with (`source_map`)[TODO: link] enabled and use the result to find the line in the source program. 

- `logic eval error: program logs too large.`

    At most, 1k may be logged and the `log` op may be called at most 32 times in a single app call. 

- `logic eval error: write budget (N) exceeded`, `logic eval error: box read budget (N) exceeded`

    Every box reference passed allots another 1k to the read/write budget for the app call. By passing more box refs a larger IO budget is available. The budget is shared across app calls within the same group so a 32k box can be read/written to as long as 32 box references are passed (8 per txn, 4 txns).

- `logic eval error: store TYPE count N exceeds schema TYPE count M`

    Schema needs to be large enough to allow storage requirements. immutable so need to recreate app

- `logic eval error: err opcode executed.`

    Unless you added an `err` to your program explicitly, it's likely this is the result of the internal routing of the app finding that there are no matching routes for some conditional block. Typically this happens when something like the ABI method selector passed is not present in the application.

- `logic eval error: overspend` 

    The account listed in the error did not have enough Algos to cover the transactions and fees of the transaction. 

- `logic eval error: * overflowed.`

    The math operation would have returned a value > max uint64. Use wide math ops or byte math ops like `b*` .

- `logic eval error: - would result negative.`

    The math operation would have returned a negative value < 0. Check the values before using them within the operation.

- `logic eval error: invalid ApplicationArgs index 0.`

    An attempt was made by the contract to access an Application Arguments array element that was not set. Check that the correct app args are passed and that the flow of the program does not accidentally reference an invalid index. 

# Environment Setup

How do I setup the Algod to connect to sandbox?

    [Algokit](TODO: link)
    [VideoSeries](TODO: link)

How do I connect to an API? Which API should I use?

    Link to all api providers
        AlgoNode
        AlgoExplorer
        PureStake - note requires diff api key header

    Discuss tradeoffs and fail overs?

# Passing around data

How do I pass a (un)signed transaction from server to front end or vice versa?
    On one side: msgpack encode it to bytes => base64 encode the bytes to a string => ship the string over the wire 
    On the other side: base64 decode the string to bytes => msgpack decode the bytes into a signed/unsigned transaction
# Querying chain state

How do I get an accounts state for an app?

    On chain or off chain?
    Off chain: use sdks to query  account data from the algod, itll be in there keyed by app id
    On Chain: make sure the account and the app id for the app that set the local storage are passed in the transaction then `App.localGetEx()` with [PyTEAL](TODO: link)