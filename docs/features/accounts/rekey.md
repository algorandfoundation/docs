Title: Rekeying 🔷

## Overview

Rekeying is a powerful protocol feature which enables an Algorand account holder to maintain a static public address while dynamically rotating the authoritative private spending key(s). This is accomplished by issuing a "rekey-to transaction" which sets the _authorized address_ field within the account object. Future transaction authorization using the account's public address must be provided by the spending key(s) associated with the _authorized address_ which may be a single key address, MultiSig address or LogicSig program address. Key management is an important concept to understand and Algorand provides tools to accomplish relevant tasks securely. 

!!! Info
    The term "spending key(s)" is used throughout this document to signify that generally either a single key or a set of keys from a MultiSig account may authorize from a given public address. The address itself cannot distinguish how many spending keys are specifically required.

### Account Review

The [account overview](../index.md#keys-and-addresses) page introduces _keys_, _addresses_ and _accounts_. During initial account generation, a public key and corresponding private spending key are created and used to derive the Algorand address. This public address is commonly displayed within wallet software and remains static for each account. When you receive Algos or other assets, they will be sent to your public Algorand address. When you send from your account, the transaction must be authorized using the appropriate private spending key(s).  

### Authorized Addresses

The _balance record_ of every account includes the "auth-addr" field which, when populated, defines the required _authorized address_ to be evaluated during transaction validation. Initially, the "auth-addr" field is implicitly set to the account's "address" field and the only valid _private spending key_ is the one created during account generation. To conserve resources, the "auth-addr" field is only stored and displayed after an authorized _rekey-to transaction_ is confirmed by the network. 

Conceptually illustrated in the image below, a "standard" account uses its _private spending key_ to authorize from its _public address_. A "rekeyed" account defines the _authorized address_ which references a distinct "foreign" address and thus requires the _private spending key(s)_ thereof to authorize future transactions.

![Accounts](../../imgs/accounts.png)

#### Standard Account

Use the following code sample to view a _standard_ account on BetaNet:

```bash tab="goal"
$ goal account dump --address NFFMZJC6H52JLEAITTJ7OIML3XCJFKIRXYRJLO4WLWIJZB7N6CTWESRAZU
```

Response:

```json hl_lines="2"
{
  "addr": "NFFMZJC6H52JLEAITTJ7OIML3XCJFKIRXYRJLO4WLWIJZB7N6CTWESRAZU",
  "algo": 100000,
  [...]
}
```
Notice the response includes the "addr" field which is the public address. Implicitly, only the spending key associated with this address may authorize transactions for this account.

#### Rekeyed Account

Next, modify your command slightly to display results for this _rekeyed_ account: `L42DW7MSHP4PMIAZSDAXYTZVHTE756KGXCJYGFKCET5XHIAWLBYYNSMZQU`.

Response:

```json hl_lines="2 4"
{
  "addr": "L42DW7MSHP4PMIAZSDAXYTZVHTE756KGXCJYGFKCET5XHIAWLBYYNSMZQU",
  "algo": 100000,
  "spend": "NFFMZJC6H52JLEAITTJ7OIML3XCJFKIRXYRJLO4WLWIJZB7N6CTWESRAZU",
  [...]
}
```

This response includes the addition of the "spend" field. This is the "auth-addr" within the _account object_ and signifies any transactions from `L42DW7MSHP4PMIAZSDAXYTZVHTE756KGXCJYGFKCET5XHIAWLBYYNSMZQU` must now be authorized by `NFFMZJC6H52JLEAITTJ7OIML3XCJFKIRXYRJLO4WLWIJZB7N6CTWESRAZU` to be confirmed by the network. 

### Rekey-to Transaction

A _rekey-to transaction_ is a `payment` type transaction which includes the `rekey-to` parameter set to a well-formed Algorand address. Authorization for this transaction must be provided by the existing _authorized address_. As shown in the first example account above, the _authorized address_ is implicitly the "addr" field of this account even though the "auth-addr" field is not explicitly defined. Only the private spending key of this "addr" address may be used to authorize a _rekey-to transaction_.

The _rekey-to transaction_ workflow is as follows: 

- Construct a payment transaction which specifies an _address_ for the `rekey-to` parameter
- Add required signature(s) from the **current** _authorized address_
- Send and confirm the transaction on the network

#### Construct Transaction

The following commands will construct an unsigned transaction file `rekey.txn` and inspect the contents:

```bash tab="goal"
$ goal clerk send --from L42DW7MSHP4PMIAZSDAXYTZVHTE756KGXCJYGFKCET5XHIAWLBYYNSMZQU \
                  --to L42DW7MSHP4PMIAZSDAXYTZVHTE756KGXCJYGFKCET5XHIAWLBYYNSMZQU \
                  --amount 0 \
                  --rekey-to NFFMZJC6H52JLEAITTJ7OIML3XCJFKIRXYRJLO4WLWIJZB7N6CTWESRAZU \
                  --out rekey.txn
$ goal clerk inspect rekey.txn
```

Response:

```json hl_lines="11"
rekey.txn[0]
{
  "txn": {
    "fee": 1000,
    "fv": 4921687,
    "gen": "betanet-v1.0",
    "gh": "mFgazF+2uRS1tMiL9dsj01hJGySEmPN28B/TjjvpVW0=",
    "lv": 4922687,
    "note": "bbD6hjNZNdg=",
    "rcv": "L42DW7MSHP4PMIAZSDAXYTZVHTE756KGXCJYGFKCET5XHIAWLBYYNSMZQU",
    "rekey": "NFFMZJC6H52JLEAITTJ7OIML3XCJFKIRXYRJLO4WLWIJZB7N6CTWESRAZU",
    "snd": "L42DW7MSHP4PMIAZSDAXYTZVHTE756KGXCJYGFKCET5XHIAWLBYYNSMZQU",
    "type": "pay"
  }
}
```

Construction of the _rekey-to transaction_ includes the `rekey-to` parameter and the value `"NFFMZJC6H52JLEAITTJ7OIML3XCJFKIRXYRJLO4WLWIJZB7N6CTWESRAZU"`. Notice the resulting unsigned transaction output includes the "rekey" field and this value.

#### Add Authorized Signature(s) 

Adding the **current** authorized signature(s) to a _rekey-to transaction_ is required prior to sending to the network for confirmation. The "snd" field provides the address to the _account object_ where the "auth-addr" field is defined. 

!!! Info
    Examples provided below demonstrate the commands in detail and allow you to _rekey_ accounts in various scenarios. 

#### Send and Confirm

Once all of the required signatures are gathered into a single signed transaction, it may be sent to the network for confirmation. The result for the sample account is:

```json hl_lines="4"
{
  "addr": "L42DW7MSHP4PMIAZSDAXYTZVHTE756KGXCJYGFKCET5XHIAWLBYYNSMZQU",
  "algo": 100000,
  "spend": "NFFMZJC6H52JLEAITTJ7OIML3XCJFKIRXYRJLO4WLWIJZB7N6CTWESRAZU",
  [...]
}
```

### Conclusion

The result of a confirmed _rekey-to transaction_ will be the "auth-addr" field of the _account object_ is defined, modified or removed. Defining or modifying means only the _private spending key(s)_ of the corresponding _authorized address_ may authorize future transactions for this _public address_. Removing the "auth-addr" field is really an explicit assignment of the _authorized address_ back to the "addr" field of the _account object_ (observed implicitly because the field is not displayed). 

The "auth-addr" may be specified within a _rekey-to transaction_ as a distinct foreign address representing a single key address, MultiSig address or LogicSig program address to provide maximum flexibility in key management options. 

!!! Warning
    The protocol does not validate control of the required spending key(s) associated with the _authorized address_ defined by `--rekey-to` parameter when the _rekey-to transaction_ is sent. This is by design and affords additional privacy features to the new _authorized address_. It is incumbent upon the user to ensure proper key management practices and `--rekey-to` assignments.


## Use Case Scenarios

!!! Info
    Below are a series of potential use cases for rekeying various accounts. 

## 1 - Rekey to Single Address

The first scenario rekeys a single key account with address "A" to a distinct single key account with address "B". This requires two single key accounts at time _t0_. The result from time _t1_ is transactions for address "A" must be authorized by address "B". 

![Rekey-to Single Key Address](../../imgs/rekey-single-single.png)

### Generate and Fund Accounts

Refer to the Getting Started guide to learn how to generate two accounts and fund their respective address from the Faucet. This example uses the following public addresses:

```bash tab="bash"
ADDR_A="UGAGADYHIUGFGRBEPHXRFI6Z73HUFZ25QP32P5FV4H6B3H3DS2JII5ZF3Q"
ADDR_B="LOWE5DE25WOXZB643JSNWPE6MGIJNBLRPU2RBAVUNI4ZU22E3N7PHYYHSY"
```

View the initial _authorized address_ for Account A using `goal`:

```
$ goal account dump --address $ADDR_A
```

Response:

```json hl_lines="2"
{
  "addr": "UGAGADYHIUGFGRBEPHXRFI6Z73HUFZ25QP32P5FV4H6B3H3DS2JII5ZF3Q",
  "algo": 100000,
  [...]
}
```

Implicitly, the _authorized address_ is the defined as the "addr" field displayed here.

### Rekey to Single Address

Account A intends to _rekey_ its _authorized address_ to `$ADDR_B` which is the public address of Account "B". This can be accomplished in a single `goal` command:

```bash tab="goal"
$ goal clerk send --from $ADDR_A --to $ADDR_A --amount 0 --rekey-to $ADDR_B
```

Results of `goal account dump --address $ADDR_A` will now display:

```json hl_lines="5"
{
  "addr": "UGAGADYHIUGFGRBEPHXRFI6Z73HUFZ25QP32P5FV4H6B3H3DS2JII5ZF3Q",
  "algo": 199000,
  [...]
  "spend": "LOWE5DE25WOXZB643JSNWPE6MGIJNBLRPU2RBAVUNI4ZU22E3N7PHYYHSY"
}
```

The populated "spend" field instructs the validation protocol to only approve transactions for this _account object_ when authorized by the _spending key(s)_ of that address. Validators will ignore all other attempted authorizations, including those from the _public address_ defined in the "addr" field. 

#### TEST: Send with Auth A

The following transaction will **fail** because, by default, `goal` attempts to add the authorization using the `--from` parameter. However, the protocol will reject this because it is expecting the authorization from `$ADDR_B` due to the confirmed _rekeying transaction_ above.

```bash tab="goal"
$ goal clerk send --from $ADDR_A --to $ADDR_B --amount 1000
```

### Send from Authorized Address

Sending from the _authorized address_ of Account "A" requires:

- Construct an unsigned transaction from `$ADDR_A`
- Sign using _authorized address_ `$ADDR_B`
- Send authorized transaction 

#### Construct Unsigned Transaction

First, construct an unsigned transaction using `goal` with the `--outfile` flag to write the unsigned transction to a file:

```bash tab="goal"
$ goal clerk send --from $ADDR_A --to $ADDR_B --amount 1000 --out send-single.txn
```

#### Sign Using Authorized Address

Next, locate the wallet containing the _private spending key_ for Account "B". The `goal clerk sign` command provides the flag `--signer` which allows specifying the proper required _authorized address_ `$ADDR_B`. Notice the `infile` flag reads in the unsigned transaction file from above and the `--outfile` flag writes the signed transaction to a separate file.

```bash tab="goal"
$ goal clerk sign --signer $ADDR_B --infile send-single.txn --outfile send-single.stxn
```

#### TEST: Send with Auth B

Finally, send the the signed transaction file using `goal`:

```bash tab="goal"
$ goal clerk rawsend --filename send-single.stxn
```

This will succeed, sending the 1000 microAlgos from `$ADDR_A` to `$ADDR_B` using the _private spending key_ of Account "B".

## 2 - Rekey to MultiSig Address

The second scenario _rekeys_ a single key account with _public address_ "A" to a MultiSig address "BC_T1". This scenario reuses both Accounts "A" and "B", adds a third Account "C" and creates a MultiSig Account "BC_T1" comprised of addresses "B" and "C" with a threshold of 1. The result will be the _private spending key_ for `$ADDR_B` or `$ADDR_C` may authorize transaction from `$ADDR_A`.

![Rekey-to MultiSig Address](../../imgs/rekey-single-multisig.png)

### Generate Additional Account

Follow the same procedure as above to generate a third account for use as "C".

### Generate New MultiSig Account

Reference the documentation to generate MultiSig account. Ensure it uses both `$ADDR_B` and the new `$ADDR_C` with a threshold of 1  (so either "B" or "C" may authorize). Set the resulting account address to the `$ADDR_BC_T1` environment variable for use below.

### Rekey to MultiSig Address

Recall from scenario 1 that Account "A" has already _rekeyed_ to `$ADDR_B`. 

### Construct Unsigned Transaction

The _rekey transaction_ constructed for this scenario requires authorize from `$ADDR_B`.

```bash tab="goal"
$ goal clerk send --from $ADDR_A --to $ADDR_A --amount 0 --rekey-to $ADDR_BC_T1
```

### Sign Rekey Transaction

```bash tab="goal"
$ goal clerk sign --signer $ADDR_B --infile rekey-multisig.txn --outfile rekey-multisig.stxn
```

### Send and Confirm Rekey to MultiSig

```bash tab="goal"
$ goal clerk rawsend --filename rekey-multisig.stxn
$ goal account dump --address $ADDR_A
```

The _rekey transaction_ will confirm, resulting in the "spend" field update within the _account object_:

```json hl_lines="5"
{
  "addr": "UGAGADYHIUGFGRBEPHXRFI6Z73HUFZ25QP32P5FV4H6B3H3DS2JII5ZF3Q",
  "algo": 199000,
  [...]
  "spend": "NEWMULTISIGADDRESSBCT1..."
}
```

#### TEST: Send with Auth BC_T1

Use the established pattern:

- Construct unsigned transaction
- Sign transaction
- Confirm transaction

```bash tab="goal"
$ goal clerk send --from $ADDR_A --to $ADDR_B --amount 1000 --out send-multisig-bct1.txn
$ goal clerk sign --signer $ADDR_C --infile send-multisig-bct1.txn --outfile send-multisig-bct1.stxn
$ goal clerk rawsend --filename send-multisig-bct1.stxn
```

This transaction will succeed as _private spending key_ for `$ADDR_C` provided the authorization and meets the threshold requirement for the MultiSig account.



