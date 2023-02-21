Title: Rekeying

## Overview

Rekeying is a powerful protocol feature which enables an Algorand account holder to maintain a static public address while dynamically rotating the authoritative private spending key(s). This is accomplished by issuing a "rekey-to transaction" which sets the _authorized address_ field within the account object. Future transaction authorization using the account's public address must be provided by the spending key(s) associated with the _authorized address_ which may be a single key address, MultiSig address or LogicSig program address. Key management is an important concept to understand and Algorand provides tools to accomplish relevant tasks securely. 

!!! Info
    The term "spending key(s)" is used throughout this document to signify that generally either a single key or a set of keys from a MultiSig account may authorize from a given public address. The address itself cannot distinguish how many spending keys are specifically required.

!!! warning
    Using the `--close-to` parameter on any transaction from a _rekeyed account_ will remove the **auth-addr** field, thus reverting signing authority to the original address. The `--close-to` parameter should be used with caution by keyholder(s) of **auth-addr** as the effects remove their authority to access this account thereafter.

### Account Review

The [account overview](../#keys-and-addresses) page introduces _keys_, _addresses_ and _accounts_. During initial account generation, a public key and corresponding private spending key are created and used to derive the Algorand address. This public address is commonly displayed within wallet software and remains static for each account. When you receive Algos or other assets, they will be sent to your public Algorand address. When you send from your account, the transaction must be authorized using the appropriate private spending key(s).  

### Authorized Addresses

The _balance record_ of every account includes the "auth-addr" field which, when populated, defines the required _authorized address_ to be evaluated during transaction validation. Initially, the "auth-addr" field is implicitly set to the account's "address" field and the only valid _private spending key_ is the one created during account generation. To conserve resources, the "auth-addr" field is only stored and displayed after an authorized _rekey-to transaction_ is confirmed by the network. 

Conceptually illustrated in the image below, a "standard" account uses its _private spending key_ to authorize from its _public address_. A "rekeyed" account defines the _authorized address_ which references a distinct "foreign" address and thus requires the _private spending key(s)_ thereof to authorize future transactions.

![Accounts](../../imgs/accounts.png)

#### Standard Account

Use the following code sample to view a _standard_ account on BetaNet:

=== "goal"
  ```bash
  goal account dump --address NFFMZJC6H52JLEAITTJ7OIML3XCJFKIRXYRJLO4WLWIJZB7N6CTWESRAZU
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

=== "goal"
  ```bash
  goal clerk send --from L42DW7MSHP4PMIAZSDAXYTZVHTE756KGXCJYGFKCET5XHIAWLBYYNSMZQU \
                    --to L42DW7MSHP4PMIAZSDAXYTZVHTE756KGXCJYGFKCET5XHIAWLBYYNSMZQU \
                    --amount 0 \
                    --rekey-to NFFMZJC6H52JLEAITTJ7OIML3XCJFKIRXYRJLO4WLWIJZB7N6CTWESRAZU \
                    --out rekey.txn
  goal clerk inspect rekey.txn
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

```bash
  ADDR_A="UGAGADYHIUGFGRBEPHXRFI6Z73HUFZ25QP32P5FV4H6B3H3DS2JII5ZF3Q"
  ADDR_B="LOWE5DE25WOXZB643JSNWPE6MGIJNBLRPU2RBAVUNI4ZU22E3N7PHYYHSY"
```

View the initial _authorized address_ for Account A using `goal`:

```bash
goal account dump --address $ADDR_A
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

=== "goal"
  ```bash
  goal clerk send --from $ADDR_A --to $ADDR_A --amount 0 --rekey-to $ADDR_B
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

=== "goal"
  ```bash
  goal clerk send --from $ADDR_A --to $ADDR_B --amount 100000
  ```

### Send from Authorized Address

Sending from the _authorized address_ of Account "A" requires:

- Construct an unsigned transaction from `$ADDR_A`
- Sign using _authorized address_ `$ADDR_B`
- Send authorized transaction 

#### Construct Unsigned Transaction

First, construct an unsigned transaction using `goal` with the `--outfile` flag to write the unsigned transction to a file:

=== "goal"
  ```bash
  goal clerk send --from $ADDR_A --to $ADDR_B --amount 100000 --out send-single.txn
  ```

#### Sign Using Authorized Address

Next, locate the wallet containing the _private spending key_ for Account "B". The `goal clerk sign` command provides the flag `--signer` which allows specifying the proper required _authorized address_ `$ADDR_B`. Notice the `infile` flag reads in the unsigned transaction file from above and the `--outfile` flag writes the signed transaction to a separate file.

=== "goal"
  ```bash
  goal clerk sign --signer $ADDR_B --infile send-single.txn --outfile send-single.stxn
  ```

#### TEST: Send with Auth B

Finally, send the the signed transaction file using `goal`:

=== "goal"
  ```bash
  goal clerk rawsend --filename send-single.stxn
  ```

This will succeed, sending the 100000 microAlgos from `$ADDR_A` to `$ADDR_B` using the _private spending key_ of Account "B".

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

=== "goal"
  ```bash
  goal clerk send --from $ADDR_A --to $ADDR_A --amount 0 --rekey-to $ADDR_BC_T1 --out rekey-multisig.txn
  ```

### Sign Rekey Transaction

=== "goal"
  ```bash
  goal clerk sign --signer $ADDR_B --infile rekey-multisig.txn --outfile rekey-multisig.stxn
  ```

### Send and Confirm Rekey to MultiSig

=== "goal"
  ```bash
  goal clerk rawsend --filename rekey-multisig.stxn
  goal account dump --address $ADDR_A
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

=== "goal"
  ```bash
  goal clerk send --from $ADDR_A --to $ADDR_B --amount 100000 --msig-params="1 $ADDR_B $ADDR_C" --out send-multisig-bct1.txn
  goal clerk multisig sign --tx send-multisig-bct1.txn --address $ADDR_C
  goal clerk rawsend --filename send-multisig-bct1.txn
  ```

This transaction will succeed as _private spending key_ for `$ADDR_C` provided the authorization and meets the threshold requirement for the MultiSig account.

## SDK Example:
In part 1, rekey from Account 3 to allow to sign from Account 1. Then in part 2, send from account 3 to account 2 and sign from Account 1.


=== "Python"
<!-- ===PYSDK_ACCOUNT_REKEY=== -->
  ```python
    import json
    from algosdk import account, mnemonic
    from algosdk.v2client import algod
    from algosdk.transaction import *


    def getting_started_example():
        algod_address = "http://localhost:4001"
        algod_token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
        algod_client = algod.AlgodClient(algod_token, algod_address)

        # Part 1
        # rekey from Account 3 to allow to sign from Account 1

        # Part 2
        # send from account 3 to account 2 and sign from Account 1

        # never use mnemonics in production code, replace for demo purposes only

        account1_passphrase = "PASTE your phrase for account 1"
        account2_passphrase = "PASTE your phrase for account 2"
        account3_passphrase = "PASTE your phrase for account 3"

        private_key1 = mnemonic.to_private_key(account1_passphrase)
        private_key2 = mnemonic.to_private_key(account2_passphrase)
        private_key3 = mnemonic.to_private_key(account3_passphrase)
        
        account1 = account.address_from_private_key(private_key1)
        account2 = account.address_from_private_key(private_key2)    
        account3 = account.address_from_private_key(private_key3)

        print("Account 1 : {}".format(account1))
        print("Account 2 : {}".format(account2))
        print("Account 3 : {}".format(account3))
        
        # Part 1
        # build transaction
        params = algod_client.suggested_params()
        # comment out the next two (2) lines to use suggested fees
        # params.flat_fee = True
        # params.fee = 1000


        # opt-in send tx to same address as sender and use 0 for amount w rekey account
        # to account 1
        amount = int(0)   
        rekeyaccount = account1
        sender = account3
        receiver = account3    
        unsigned_txn = PaymentTxn(
          sender, params, receiver, amount, None, None, None, rekeyaccount)

        # sign transaction with account 3
        signed_txn = unsigned_txn.sign(
          mnemonic.to_private_key(account3_passphrase))
        txid = algod_client.send_transaction(signed_txn)
        print("Signed transaction with txID: {}".format(txid))

        # wait for confirmation

        confirmed_txn = wait_for_confirmation(algod_client, txid, 4)
        print("TXID: ", txid)
        print("Result confirmed in round: {}".format(confirmed_txn['confirmed-round']))

        # read transction
        try:
            confirmed_txn = algod_client.pending_transaction_info(txid)
            account_info = algod_client.account_info(account3)
            
        except Exception as err:
            print(err)
        print("Transaction information: {}".format(
            json.dumps(confirmed_txn, indent=4)))
        print("Account 3 information : {}".format(
            json.dumps(account_info, indent=4)))

        #  Part 2
        #  send payment from account 3
        #  to acct 2 and signed by account 1


        private_key_account1 = mnemonic.to_private_key(account1_passphrase)  
        account1 = account.address_from_private_key(private_key_account1)

        private_key_account2 = mnemonic.to_private_key(account2_passphrase)  
        account2 = account.address_from_private_key(private_key_account2)

        private_key_account3 = mnemonic.to_private_key(account3_passphrase)  
        account3 = account.address_from_private_key(private_key_account3)

        amount = int(1000000)
        receiver = account2
        unsigned_txn = PaymentTxn(
          account3, params, receiver, amount, None, None, None, account1)
        # sign transaction
        signed_txn = unsigned_txn.sign(
          mnemonic.to_private_key(account1_passphrase))
        txid = algod_client.send_transaction(signed_txn)
        print("Signed transaction with txID: {}".format(txid))

        # wait for confirmation

        confirmed_txn = wait_for_confirmation(algod_client, txid, 4)
        print("TXID: ", txid)
        print("Result confirmed in round: {}".format(confirmed_txn['confirmed-round']))
        account_info_rekey = algod_client.account_info(account3)
        print("Account 3 information (from) : {}".format(
            json.dumps(account_info_rekey, indent=4)))
        account_info_rekey = algod_client.account_info(account2)
        print("Account 2 information (to) : {}".format(
            json.dumps(account_info_rekey, indent=4)))


    getting_started_example()

  ```
<!-- ===PYSDK_ACCOUNT_REKEY=== -->

=== "JavaScript"
<!-- ===JSSDK_ACCOUNT_REKEY=== -->
<!-- ===JSSDK_ACCOUNT_REKEY=== -->


=== "Java"
<!-- ===JAVASDK_ACCOUNT_REKEY=== -->
<!-- ===JAVASDK_ACCOUNT_REKEY=== -->

=== "Go"
<!-- ===GOSDK_ACCOUNT_REKEY=== -->
<!-- ===GOSDK_ACCOUNT_REKEY=== -->