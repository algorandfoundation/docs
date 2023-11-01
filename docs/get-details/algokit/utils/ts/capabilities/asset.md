# Assets

The asset management functions include opting in and out of assets, which are fundamental to asset interaction in a blockchain environment.
To see some usage examples check out the [automated tests](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/asset.spec.ts).

## `optIn`

Before an account can receive a specific asset, it must `opt-in` to receive it. An opt-in transaction places an asset holding of 0 into the account and increases its minimum balance by [100,000 microAlgos](https://developer.algorand.org/docs/get-details/asa/#assets-overview).
The `optIn` function facilitates the opt-in process for an account to multiple assets, allowing the account to receive and hold those assets.

## `optOut`

An account can opt out of an asset at any time. This means that the account will no longer hold the asset, and the account will no longer be able to receive the asset. The account also recovers the Minimum Balance Requirement for the asset (100,000 microAlgos)
The `optOut` function manages the opt-out process, permitting the account to discontinue holding a group of assets.

!!! Note 
    It's essential to note that an account can only opt-out of an asset if its balance of that asset is zero.
