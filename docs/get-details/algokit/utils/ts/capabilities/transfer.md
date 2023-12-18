# Algo transfers

Algo transfers is a higher-order use case capability provided by AlgoKit Utils that builds on top of the core capabilities, particularly [Algo amount handling](./amount.md) and [Transaction management](./transaction.md). It allows you to easily initiate algo transfers between accounts, including dispenser management and idempotent account funding.

To see some usage examples check out the [automated tests](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/transfer.spec.ts).

## `transferAlgos`

The key function to facilitate Algo transfers is `algokit.transferAlgos(transfer, algod)`, which returns a [`SendTransactionResult`](./transaction.md#sendtransactionresult) and takes a [`AlgoTransferParams`](../code/interfaces/types_transfer.AlgoTransferParams.md):

- All properties in [`SendTransactionParams`](./transaction.md#sendtransactionparams)
- `from: SendTransactionFrom` - The account that will send the ALGOs
- `to: SendTransactionFrom | string` - The address of the account that will receive the ALGOs
- `amount: AlgoAmount` - The [amount](./amount.md) of ALGOs to send
- `transactionParams?: SuggestedParams` - The optional [transaction parameters](./transaction.md#transaction-params)
- `note?: TransactionNote` - The [transaction note](./transaction.md#transaction-notes)
- `lease?: string | Uint8Array`: A [lease](https://developer.algorand.org/articles/leased-transactions-securing-advanced-smart-contract-design/) to assign to the transaction to enforce a mutually exclusive transaction (useful to prevent double-posting and other scenarios)

## `ensureFunded`

The `ensureFunded` function automatically funds an account to maintain a minimum amount of disposable ALGOs. This is particularly useful for automation and deployment scripts. The function is defined as `algokit.ensureFunded(funding, algod, kmd?)` and returns a [`EnsureFundedReturnType`](../code/interfaces/types_transfer.EnsureFundedReturnType.md) if a transaction was needed, or `undefined` if no transaction was required. The function takes a [`EnsureFundedParams`](../code/interfaces/types_transfer.EnsureFundedParams.md) object as an argument:

- All properties in [`SendTransactionParams`](./transaction.md#sendtransactionparams)
- `accountToFund: SendTransactionFrom | string` - The account that is to be funded
- `fundingSource?: SendTransactionFrom | TestNetDispenserApiClient` - The account that is the source of funds or a dispenser API client. If not specified, it will use the [dispenser](./account.md#dispenser)
- `minSpendingBalance: AlgoAmount` - The minimum balance of ALGOs that the account should have available to spend (i.e., on top of the minimum balance requirement)
- `minFundingIncrement?: AlgoAmount` - When issuing a funding amount, the minimum amount to transfer. This avoids many small transfers if this function gets called often on an active account
- `transactionParams?: SuggestedParams` - The optional [transaction parameters](./transaction.md#transaction-params)
- `note?: TransactionNote` - The [transaction note](./transaction.md#transaction-notes)
- `lease?: string | Uint8Array`: A [lease](https://developer.algorand.org/articles/leased-transactions-securing-advanced-smart-contract-design/) to assign to the transaction to enforce a mutually exclusive transaction (useful to prevent double-posting and other scenarios)

The function calls Algod to find the current balance and minimum balance requirement, calculates the difference between those two numbers, and checks to see if it's more than the `minSpendingBalance`. If so, it will send the difference, or the `minFundingIncrement` if that is specified. If the `fundingSource` is an instance of `TestNetDispenserApiClient`, the function will use the dispenser API to fund the account. Refer to [algokit-cli documentation](https://github.com/algorandfoundation/algokit-cli/blob/main/docs/features/dispenser.md#ci-access-token) for details on obtaining an access token for AlgoKit TestNet Dispenser API.

## `transferAsset`

The key function to facilitate asset transfers is `transferAsset(transfer, algod)`, which returns a [`SendTransactionResult`](./transaction.md#sendtransactionresult) and takes a [`TransferAssetParams`](../code/interfaces/types_transfer.TransferAssetParams.md):

- All properties in [`SendTransactionParams`](./transaction.md#sendtransactionparams)
- `from: SendTransactionFrom` - The account that will send the asset
- `to: SendTransactionFrom | string` - The account / account address that will receive the asset
- `assetId: number` - The asset id that will be transfered
- `amount: number | bigint` - The amount to send in the smallest divisible unit
- `transactionParams?: SuggestedParams` - The optional [transaction parameters](./transaction.md#transaction-params)
- `clawbackFrom: SendTransactionFrom | string` - An optional address of a target account from which to perform a clawback operation. Please note, in such cases senderAccount must be equal to clawback field on ASA metadata.
- `note?: TransactionNote` - The [transaction note](./transaction.md#transaction-notes)
- `lease?: string | Uint8Array`: A [lease](https://developer.algorand.org/articles/leased-transactions-securing-advanced-smart-contract-design/) to assign to the transaction to enforce a mutually exclusive transaction (useful to prevent double-posting and other scenarios)

## Dispenser

If you want to programmtically send funds then you will often need a "dispenser" account that has a store of ALGOs that can be sent and a private key available for that dispenser account.

There is a standard AlgoKit Utils function to get access to a [dispenser account](./account.md#accounts): [`algokit.getDispenserAccount(algod, kmd?)`](../code/modules/index.md#getdispenseraccount). When running against [LocalNet](https://github.com/algorandfoundation/algokit-cli/blob/main/docs/features/localnet.md), the dispenser account can be automatically determined using the [Kmd API](https://developer.algorand.org/docs/rest-apis/kmd). When running against other networks like TestNet or MainNet the mnemonic (and optionally sender address if it's been rekeyed) of the dispenser account can be provided via environment variables (`process.env.DISPENSER_MNEMONIC` and optionally `process.env.DISPENSER_SENDER` if rekeyed).

!!! Note 
    This does not refer to the [AlgoKit TestNet Dispenser API](./dispenser-client.md) which is a separate abstraction that can be used to fund accounts on TestNet via dedicated API service.
