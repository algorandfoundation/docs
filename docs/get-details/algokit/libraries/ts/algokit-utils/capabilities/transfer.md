# Algo transfers

Algo transfers is a higher-order use case capability provided by AlgoKit Utils that builds on top of the core capabilities, particularly [Algo amount handling](./amount.md) and [Transaction management](./transaction.md). It allows you to easily initiate algo transfers between accounts, including dispenser management and idempotent account funding.

To see some usage examples check out the [automated tests](../../src/transfer.spec.ts).

## `transferAlgos`

The key function to facilitate Algo transfers is `algokit.transferAlgos(transfer, algod)`, which returns a [`SendTransactionResult`](./transaction.md#sendtransactionresult) and takes a [`AlgoTransferParams`](../code/interfaces/types_transfer.AlgoTransferParams.md):

- All properties in [`SendTransactionParams`](./transaction.md#sendtransactionparams)
- `from: SendTransactionFrom` - The account that will send the ALGOs
- `to: string` - The address of the account that will receive the ALGOs
- `amount: AlgoAmount` - The [amount](./amount.md) of ALGOs to send
- `transactionParams?: SuggestedParams` - The optional [transaction parameters](./transaction.md#transaction-params)
- `note?: TransactionNote` - The [transaction note](./transaction.md#transaction-notes)

## `ensureFunded`

The ability to automatically fund an account to have a minimum amount of disposable ALGOs to spend is incredibly useful for automation and deployment scripts. The function to facilitate this is `algokit.ensureFunded(funding, algod, kmd?)`, which returns a [`SendTransactionResult`](./transaction.md#sendtransactionresult) (or undefined if it didn't need to send a transaction) and takes a [`EnsureFundedParams`](../code/interfaces/types_transfer.EnsureFundedParams.md):

- All properties in [`SendTransactionParams`](./transaction.md#sendtransactionparams)
- `accountToFund: SendTransactionFrom | string` - The account that is to be funded
- `fundingSource?: SendTransactionFrom` - The account that is the source of funds, if not specified then it will use the [dispenser](./account.md#dispenser)
- `minSpendingBalance: AlgoAmount` - The minimum balance of ALGOs that the account should have available to spend (i.e. on top of minimum balance requirement)
- `minFundingIncrement?: AlgoAmount` - When issuing a funding amount, the minimum amount to transfer (avoids many small transfers if this gets called often on an active account)
- `amount: AlgoAmount` - The [amount](./amount.md) of ALGOs to send
- `transactionParams?: SuggestedParams` - The optional [transaction parameters](./transaction.md#transaction-params)
- `note?: TransactionNote` - The [transaction note](./transaction.md#transaction-notes)

The function calls Algod to find the current balance and minimum balance requirement, gets the difference between those two numbers and checks to see if it's more than the `minSpendingBalance` and if so then it will send the difference, or the `minFundingIncrement` if that is specified.

## Dispenser

If you want to programmtically send funds then you will often need a "dispenser" account that has a store of ALGOs that can be sent and a private key available for that dispenser account.

There is a standard AlgoKit Utils function to get access to a [dispenser account](./account.md#accounts): [`algokit.getDispenserAccount(algod, kmd?)`](../code/modules/index.md#getdispenseraccount). When running against [LocalNet](https://github.com/algorandfoundation/algokit-cli/blob/main/docs/features/localnet.md), the dispenser account can be automatically determined using the [Kmd API](https://developer.algorand.org/docs/rest-apis/kmd). When running against other networks like TestNet or MainNet the mnemonic (and optionally sender address if it's been rekeyed) of the dispenser account can be provided via environment variables (`process.env.DISPENSER_MNEMONIC` and optionally `process.env.DISPENSER_SENDER` if rekeyed).
