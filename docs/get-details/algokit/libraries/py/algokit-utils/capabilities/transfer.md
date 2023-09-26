# Algo transfers

Algo transfers is a higher-order use case capability provided by AlgoKit Utils allows you to easily initiate algo transfers between accounts, including dispenser management and 
idempotent account funding.

To see some usage examples check out the [automated tests](https://github.com/algorandfoundation/algokit-utils-py/blob/main/tests/test_transfer.py).

## Transferring Algos

The key function to facilitate Algo transfers is `algokit.transfer(algod_client, transfer_parameters)`, which returns the underlying `PaymentTxn` and takes a `TransferParameters`

The following fields on `TransferParameters` are required to transfer ALGOs:
* `from_account`: The account or signer that will send the ALGOs
* `to_address`: The address of the account that will receive the ALGOs
* `micro_algos`: The amount of micro ALGOs to send

## Ensuring minimum Algos

The ability to automatically fund an account to have a minimum amount of disposable ALGOs to spend is incredibly useful for automation and deployment scripts. 
The function to facilitate this is `ensure_funded(algod, parameters)`, which takes a `EnsureBalanceParameters` and returns the underlying `PaymentTxn` if a payment was made, or None otherwise

The following fields on `EnsureBalanceParameters` are required to ensure minimum ALGOs:

* `account_to_fund`: The account address that will receive the ALGOs
* `min_spending_balance_micro_algos`: The minimum balance of micro ALGOs that the account should have available to spend (i.e. on top of minimum balance requirement)
* `min_funding_increment_micro_algos`: When issuing a funding amount, the minimum amount to transfer (avoids many small transfers if this gets called often on an active account)
* `funding_source`: The account or signer that will send the ALGOs, if not set will use `get_dispenser_account`

The function calls Algod to find the current balance and minimum balance requirement, gets the difference between those two numbers and checks to see if it's more than the 
`min_spending_balance_micro_algos` and if so then it will send the difference, or the `min_funding_increment_micro_algos` if that is specified.

## Dispenser

If you want to programmatically send funds then you will often need a "dispenser" account that has a store of ALGOs that can be sent and a private key available for that dispenser account.

There is a standard AlgoKit Utils function to get access to a [dispenser account](./account.md#account): `get_dispenser_account`. When running against 
[LocalNet](https://github.com/algorandfoundation/algokit-cli/blob/main/docs/features/localnet.md), the dispenser account can be automatically determined using the 
[Kmd API](https://developer.algorand.org/docs/rest-apis/kmd). When running against other networks like TestNet or MainNet the mnemonic of the dispenser account can be provided via environment 
variable `DISPENSER_MNEMONIC`
