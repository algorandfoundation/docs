# Algo transfers

Algo transfers is a higher-order use case capability provided by AlgoKit Utils allows you to easily initiate algo transfers between accounts, including dispenser management and
idempotent account funding.

To see some usage examples check out the [automated tests](https://github.com/algorandfoundation/algokit-utils-py/blob/main/tests/test_transfer.py).

## Transferring Algos

The key function to facilitate Algo transfers is `algokit.transfer(algod_client, transfer_parameters)`, which returns the underlying `EnsureFundedResponse` and takes a `TransferParameters`

The following fields on `TransferParameters` are required to transfer ALGOs:

- `from_account`: The account or signer that will send the ALGOs
- `to_address`: The address of the account that will receive the ALGOs
- `micro_algos`: The amount of micro ALGOs to send

## Ensuring minimum Algos

The ability to automatically fund an account to have a minimum amount of disposable ALGOs to spend is incredibly useful for automation and deployment scripts.
The function to facilitate this is `ensure_funded(client, parameters)`, which takes an `EnsureBalanceParameters` instance and returns the underlying `EnsureFundedResponse` if a payment was made, a string if the dispenser API was used, or None otherwise.

The following fields on `EnsureBalanceParameters` are required to ensure minimum ALGOs:

- `account_to_fund`: The account address that will receive the ALGOs. This can be an `Account` instance, an `AccountTransactionSigner` instance, or a string.
- `min_spending_balance_micro_algos`: The minimum balance of micro ALGOs that the account should have available to spend (i.e. on top of minimum balance requirement).
- `min_funding_increment_micro_algos`: When issuing a funding amount, the minimum amount to transfer (avoids many small transfers if this gets called often on an active account). Default is 0.
- `funding_source`: The account (with private key) or signer that will send the ALGOs. If not set, it will use `get_dispenser_account`. This can be an `Account` instance, an `AccountTransactionSigner` instance, [`TestNetDispenserApiClient`](dispenser-client.md) instance, or None.
- `suggested_params`: (optional) Transaction parameters, an instance of `SuggestedParams`.
- `note`: (optional) The transaction note, default is "Funding account to meet minimum requirement".
- `fee_micro_algos`: (optional) The flat fee you want to pay, useful for covering extra fees in a transaction group or app call.
- `max_fee_micro_algos`: (optional) The maximum fee that you are happy to pay (default: unbounded). If this is set it's possible the transaction could get rejected during network congestion.

The function calls Algod to find the current balance and minimum balance requirement, gets the difference between those two numbers and checks to see if it's more than the `min_spending_balance_micro_algos`. If so, it will send the difference, or the `min_funding_increment_micro_algos` if that is specified. If the account is on TestNet and `use_dispenser_api` is True, the [AlgoKit TestNet Dispenser API](../../../../features/dispenser) will be used to fund the account.

!!! Note
	If you are attempting to fund via Dispenser API, make sure to set `ALGOKIT_DISPENSER_ACCESS_TOKEN` environment variable prior to invoking `ensure_funded`. To generate the token refer to [AlgoKit CLI documentation](../../../../features/dispenser#login)

## Transfering Assets

The key function to facilitate asset transfers is `transfer_asset(algod_client, transfer_parameters)`, which returns a `AssetTransferTxn` and takes a `TransferAssetParameters`:

The following fields on `TransferAssetParameters` are required to transfer assets:

- `from_account`: The account or signer that will send the ALGOs
- `to_address`: The address of the account that will receive the ALGOs
- `asset_id`: The asset id that will be transfered
- `amount`: The amount to send as the smallest divisible unit value

## Dispenser

If you want to programmatically send funds then you will often need a "dispenser" account that has a store of ALGOs that can be sent and a private key available for that dispenser account.

There is a standard AlgoKit Utils function to get access to a [dispenser account](./account.md#account): `get_dispenser_account`. When running against
[LocalNet](../../../../features/localnet), the dispenser account can be automatically determined using the
[Kmd API](https://developer.algorand.org/docs/rest-apis/kmd). When running against other networks like TestNet or MainNet the mnemonic of the dispenser account can be provided via environment
variable `DISPENSER_MNEMONIC`

Please note that this does not refer to the [AlgoKit TestNet Dispenser API](./dispenser-client.md) which is a separate abstraction that can be used to fund accounts on TestNet via dedicated API service.
