Title: Rekeying

## Overview

Rekeying is a powerful protocol feature which enables an Algorand account holder to maintain a static receiving public address while dynamically rotating the authorized spending key(s). This is accomplished by defining an _authorized address_ for each Algorand account object which references a distinct address, MultiSig address or LogicSig address. Key management is an important concept to understand and Algorand provide tools to accomplish relevant tasks securely. 

!!! Warning
    This documentation is based on pre-release information and considered neither authoritative nor final. Critiques are welcome to the PR or @ryanRfox on Slack. Testing is _only_ available (to my knowledge) on a private network built from the [maxj/applications](https://github.com/justicz/go-algorand/tree/maxj/applications) repo. Unfortunately, the code samples provided will not (yet) work on BetaNet or TestNet. There are many "TODO:" items reminding me to change them prior to publication.

### Quick Review

The [account overview](https://staging.new-dev-site.algorand.org/docs/features/accounts/#keys-and-addresses0) page introduced _keys_, _addresses_ and _accounts_. Initially, each Algorand account is comprised of a public key (pk) and the spending key (sk), which derive the address. The address is commonly displayed within wallet software and remains static for each account. When you receive Algos or other assets, they will be sent to your address. When you send from your account, the transaction must be authorized using the spending key (sk).  

### Introducing Authorized Address

The balance record of every account includes the _auth-addr_ field which, when populated, defines the authorized address (aa) to sign transactions from this account. Initially, the _auth-addr_ field is implicitly set to the account's _address_ field and the only valid (aa) is the (sk) created during account generation. To conserve resources, the _auth-addr_ field is only stored and displayed when an authorized `rekey-to` transaction is confirmed by the network. 

TODO: Figure 1. Insert drawing of account object which includes the terms and relationships noted above.

#### Initial Account

Use the following code samples to view a brand new account on TestNet:

```bash tab="goal"
# TODO: Change ACCOUNT_ID to match BetaNet, then TestNet later.
# Set the appropriate values for your environment. These are the defaults for Sandbox on TestNet
$ HEADER="x-algo-api-token:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
$ curl "localhost:4001/v2/accounts/2OOTR3IDG57QCHTCFWHMN6YEEQUMGHCT3HIBMPH5A3VQCDZG5GK7DIKDXI" -H $HEADER
```

```bash tab="JavaScript"

``` 

```bash tab="Python"

``` 

```bash tab="Java"

``` 

```bash tab="Go"

``` 

Response:
Notice the account object lacks the _auth-addr_ field. The (sk) for _address_ is the implicit (aa) for an initial account.

```json hl_lines="2"
{
    "address": "2OOTR3IDG57QCHTCFWHMN6YEEQUMGHCT3HIBMPH5A3VQCDZG5GK7DIKDXI",
    "amount": 5000020000000000,
    "amount-without-pending-rewards": 5000020000000000,
    "apps-local-state": [],
    "apps-total-schema": {
        "num-byte-slice": 0,
        "num-uint": 0
    },
    "assets": [],
    "created-apps": [],
    "created-assets": [],
    "pending-rewards": 0,
    "reward-base": 0,
    "rewards": 0,
    "round": 1602,
    "status": "Offline"
}
```

#### Authorized Account

Next, modify your code slightly to display results for this account `YQP5SHSZOGOUUXIXMBM445HM5N67SCV3XZQ7TXNFOTLC7ZY5QW24DHBDOY`.

Response:
Notice the _auth-addr_ field is populated, which means any transactions from `YQP5SHSZOGOUUXIXMBM445HM5N67SCV3XZQ7TXNFOTLC7ZY5QW24DHBDOY` must now be signed by `2OOTR3IDG57QCHTCFWHMN6YEEQUMGHCT3HIBMPH5A3VQCDZG5GK7DIKDXI` to become authorized. 

```json hl_lines="2 11"
{
    "address": "YQP5SHSZOGOUUXIXMBM445HM5N67SCV3XZQ7TXNFOTLC7ZY5QW24DHBDOY",
    "amount": 5000020000000000,
    "amount-without-pending-rewards": 5000020000000000,
    "apps-local-state": [],
    "apps-total-schema": {
        "num-byte-slice": 0,
        "num-uint": 0
    },
    "assets": [],
    "auth-addr": "2OOTR3IDG57QCHTCFWHMN6YEEQUMGHCT3HIBMPH5A3VQCDZG5GK7DIKDXI",
    "created-apps": [],
    "created-assets": [],
    "pending-rewards": 0,
    "reward-base": 0,
    "rewards": 0,
    "round": 1602,
    "status": "Offline"
}
```

#### rekey-to Event

The only way to change the (aa) is to confirming a "rekey-to" transaction on the network. Initially, the (aa) is implicitly the (sk), as shown in the first example above. The second example account above completed at least one successful "rekey-to" event, perhaps many, but we only observe the most recent and authoritative result. The "rekey-to" event is comprised of: 

- Construct a payment transaction which specifies an _address_ for the `rekey-to` parameter
- Sign the transaction with the current (aa) 
- Confirm the transaction on the network

The result will be the _auth-addr_ field of the account object is defined, modified or removed. Defining or modifying means only the (sk) of the corresponding _auth-addr_ may authorize future transactions for this _address_. Removing the _auth-addr_ is really an explicit redefining of the (aa) back to the (sk) of this _address_ (observed implicitly). 

The _auth-addr_ may be specified as a distinct address, MultiSig address or LogicSig address to provide maximum flexibility in key management options. The defined _auth-addr_ may be nested as well.

!!! Warning
    The protocol does not validate control of the (sk) defined for the _auth-addr_ during a "rekey-to" event. It is incumbent upon the user to ensure proper key management practices.


!!! Info
    Scenarios with code samples will demonstrate key rotation below.

### Potential Use Cases

## Scenario 1 - Rekey to Single Address

### Generate New Account

### Rekey to Single Address

#### TEST: Send from Initial Account

### Send from Authorized Account (Single Address)

### Generate Unsigned Transaction

### Sign Using Authorized Account

### Broadcast Signed Transaction

## Scenario 2 - Remove Authorized Account

### Load Existing Account with AuthAddr Set

### Rekey to This Account Address

### Send from This Account

## Scenario 3 - Rekey to MultiSig Address

### Generate Two New Accounts

### Generate New MultiSig Account

### Rekey to MultiSig Address

### Generate Unsigned Transaction

### Sign using MultiSig Component Accounts

### Broadcast Signed Transaction

## Scenario 4 - Update MultiSig Membership




