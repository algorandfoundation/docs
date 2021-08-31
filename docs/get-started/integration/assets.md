title: Algos and assets

!!! warning "TO DO"
    This is placeholder content for this page. Needs to be updated.
 
This guide maps core technical concepts related to the Algo to Algorand Standard Assets -- Algorand's Layer-1 implementation for third-party fungible and non-fungible tokens. 

## Transactions and Minimum Balances

Transferring ASAs is very similar to transferring Algos, with just a few important differences.


!!! note "Algos to ASAs"
    **Transaction Types**: Similar to a `pay` transaction, which exclusively transfers **Algos**, an `axfer` transaction exclusively transfers **Algorand Standard Assets**.
    
    **Minimum Balances**: Accounts on Algorand require a minimum balance of 100,000 microAlgos. This balance requirement increases by 100,000 microAlgos for _each_ asset holding.
    

    
The key structural differences between **Algo Payments** and **ASA Transfers** are summarized in the table below. 
    

|Field Description|Payment (Algos)|Transfer (ASAs)|
|-----|----|------|
|Transaction `"type"`|`"pay"`|`"axfer"`|
|Total amount to transfer|`"amt"`|`"aamt"`|
|Sender|`"snd"`|`"snd"`|
|Receiver|`"rcv"`|`"arcv"`|
|Asset Identifier|No need to specify an ID since Algos is implied by the transaction type.|The value of `"xaid"` determines the specific asset to be transferred.|

**Full Documentation**

[**Transaction Fields**](https://developer.algorand.org/docs/features/transactions/) 
The full list of required and optional transaction fields.
[Payment (Algos)](https://developer.algorand.org/docs/reference/transactions/#payment-transaction) * [Transfer (ASAs)](https://developer.algorand.org/docs/reference/transactions/#asset-transfer-transaction)

[**Transaction Structure**](https://developer.algorand.org/docs/features/transactions/) 
More detailed structural views and explanations of various transaction types.
[Payment (Algos)](https://developer.algorand.org/docs/features/transactions/#payment-transaction) * [Transfer (ASAs)](https://developer.algorand.org/docs/features/transactions/#transfer-an-asset) * [Opt-In (ASAs)](https://developer.algorand.org/docs/features/transactions/#opt-in-to-an-asset)

[**Asset Transaction How-Tos**](https://developer.algorand.org/docs/features/asa/)
How-to use the SDKs and `goal` to create asset-related transactions.
[Payment (Algos)](https://developer.algorand.org/docs/build-apps/hello_world/) * [Transfer (ASAs)](https://developer.algorand.org/docs/features/asa/#transferring-an-asset) * [Opt-In (ASAs)](https://developer.algorand.org/docs/features/asa/#receiving-an-asset)


----------

## Opting In and Out of ASAs

!!! note "Algos to ASAs"
    Any account can receive Algos, however a potential recipient of a specific ASA must first opt-in to the asset so that the account holder does not see their minimum balance requirement increase without their knowing.    
    
**"Opting In"** to an asset is simply an asset transfer of 0, to and from the opting in account as shown in the table below.

|Field Description|Transfer (ASAs)| Opt-In (ASAs)|
|-----|------|-----|
|Transaction `"type"`|`"axfer"`|`"axfer"`|
|Total amount to transfer|`"aamt"`|`"aamt"` with value of `0`|
|Sender|`"snd"`|`"snd"` (must be same as `"arcv"`)|
|Receiver|`"arcv"`|`"arcv"` (must be same as `"snd"`)|
|Asset Identifier|The value of `"xaid"` determines the specific asset to be transferred.|The value of `"xaid"` determines the specific asset to opt-in to.|

**"Opting out"** of an asset requires specifying an `AssetCloseTo` (`"aclose"`) field in the [Asset Transfer transaction type](https://developer.algorand.org/docs/reference/transactions/#asset-transfer-transaction). This will result in removal of the ASA holding from the sender's account and a decrease in the minimum balance requirement for that account by 100,000 microAlgos.

----------

## Closing Accounts

[In the Algo `"pay"` transaction](https://developer.algorand.org/docs/features/transactions/#close-an-account), you can close-out an account by specifying an address in the optional `CloseRemainderTo` (`"close"`) field. The result of this action is that the remaining account balance will be sent to the specified address and the `"snd"` address will be effectively removed from the ledger.

An account must opt-out of all asset holdings _before_ closing out its Algo balance.


----------

## Account Balances

Algo and ASA balances are both located in an [account's balance record](https://developer.algorand.org/docs/reference/rest-apis/algod/v2/#account).

[`GET /v2/accounts/{address}`](https://developer.algorand.org/docs/reference/rest-apis/algod/v2/#get-v2accountsaddress)

The Algo balance can be found at the top level under `"amount"`, while multiple ASA balances are contained in an array under `"assets"`. 

|Algo Balance | ASA Balance |
|---|---|
|`"amount": <micro-algo-balance>`|`"assets":[{"amount": <asa-balance>, "asset-id": <asset-id>, ...}, ...]`|


Algorand provides [an advanced tool, called the Indexer](https://developer.algorand.org/docs/features/indexer/), that offers an API to efficiently search historical transaction data, including multiple ways to search for ASA data. A few relevant examples include: [Searching Assets by Name](https://developer.algorand.org/docs/features/indexer/#search-assets), [Searching for Accounts Opted In to an Asset](https://developer.algorand.org/docs/features/indexer/#searching-for-accounts-based-on-asset), and [Searching for Transactions that Involve a Specific ASA](https://developer.algorand.org/docs/features/indexer/#searching-for-transactions-based-on-asset).


-----
[_Sign up for the Algorand Developer Newsletter_](https://developer.algorand.org/pages/newsletter/) _to get the latest on new developer tools and features._