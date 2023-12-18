# Account management

Account management is one of the core capabilities provided by AlgoKit Utils. It allows you to create mnemonic, rekeyed, multisig, transaction signer, idempotent KMD and environment variable injected accounts that can be used to sign transactions as well as representing a sender address at the same time. This significantly simplifies passing around sender/signer within and into AlgoKit Utils.

## `SendTransactionFrom`

Any AlgoKit Utils function that needs to sign/send a transaction will take a [`SendTransactionFrom`](../code/modules/types_transaction.md#sendtransactionfrom) object, which represents an account that combined a sender and signer and is a type union between the following types:

- `Account` - An in-built algosdk `Account` object
- [`SigningAccount`](../code/classes/types_account.SigningAccount.md) - An abstraction around `algosdk.Account` that supports rekeyed accounts
- `LogicSigAccount` - An in-built algosdk `algosdk.LogicSigAccount` object
- [`MultisigAccount`](../code/classes/types_account.MultisigAccount.md) - An abstraction around `algosdk.MultisigMetadata`, `algosdk.makeMultiSigAccountTransactionSigner`, `algosdk.multisigAddress`, `algosdk.signMultisigTransaction` and `algosdk.appendSignMultisigTransaction` that supports multisig accounts with one or more signers present
- [`TransactionSignerAccount`](../code/interfaces/types_account.TransactionSignerAccount.md) - An interface that provides a sender address alongside a transaction signer (e.g. for use with `AtomicTransactionComposer` or [useWallet](https://github.com/TxnLab/use-wallet))

The use of in-built algosdk types like `Account`, `LogicSigAccount` and `TransactionSigner` is aligned to the [Modularity](../index.md#core-principles) principle. Allowing you to co-exist non AlgoKit Utils code with AlgoKit Utils functions.

## Using `SendTransactionFrom`

AlgoKit Utils provides a few helper methods to take one of these `SendTransactionFrom` objects:

- [`algokit.getSenderAddress`](../code/modules/index.md#getsenderaddress) - Returns the public address of the sender the account represents
- [`algokit.getSenderTransactionSigner`](../code/modules/index.md#getsendertransactionsigner) - Returns a `TransactionSigner` to represent the signer of the account
!!! Note
    this is memoized so multiple calls to this for the same account will safely return the same `TransactionSigner` instance; this works nicely with `AtomicTransactionComposer`
- [`algokit.signTransaction`](../code/modules/index.md#signtransaction) - Signs a single `algosdk.Transaction` object with the given account

## Accounts

In order to get the accounts you can use the underlying algosdk methods where relevant, or you can use the following AlgoKit Utils functions (all of which return a type compatible with `SendTransactionFrom`):

- [`algokit.mnemonicAccountFromEnvironment(account, algod, kmd?)`](../code/modules/index.md#mnemonicaccountfromenvironment) - Returns an Algorand account with private key loaded by convention based on the given name identifier - either by idempotently creating the account in KMD or from environment variable via `process.env['{NAME}_MNEMONIC']` and (optionally) `process.env['{NAME}_SENDER']` (if account is rekeyed)
  - This allows you to have powerful code that will automatically create and fund an account by name locally and when deployed against TestNet/MainNet will automatically resolve from environment variables, without having to have different code
  
!!! Note
    `account` can either be a string name, or an object with `{name: string, fundWith?: AlgoAmount}`, where `fundWith` allows you to control how many ALGOs are seeded into an account created in KMD
- [`algokit.mnemonicAccount(mnemonicSecret)`](../code/modules/index.md#mnemonicaccount) - Returns an Algorand account (`algosdk.Account`) with secret key loaded (i.e. that can sign transactions) by taking the mnemonic secret.
- [`algokit.multisigAccount(multisigParams, signingAccounts)`](../code/modules/index.md#multisigaccount) - Returns a multisig account with one or more signing keys loaded.
- [`algokit.rekeyedAccount(signer, sender)`](../code/modules/index.md#rekeyedaccount) - Returns a `SigningAccount` representing the given rekeyed sender/signer combination
- [`algokit.transactionSignerAccount(signer, sender)`](../code/modules/index.md#transactionsigneraccount) - Returns a `TransactionSigner` along with its sender address
- [`algokit.randomAccount()`](../code/modules/index.md#randomaccount) - Returns a new, cryptographically randomly generated account with private key loaded.

### Dispenser

- [`algokit.getDispenserAccount(algod, kmd?)`](../code/modules/index.md#getdispenseraccount) - Returns an account that can act as a dispenser to fund other accounts either via Kmd (when targeting LocalNet) or by convention from environment variable via `process.env.DISPENSER_MNEMONIC` (and optionally `process.env.DISPENSER_SENDER` if rekeyed)

## Rekey account

One of the unique features of Algorand is the ability to change the private key that can authorise transactions for an account. This is called [rekeying](https://developer.algorand.org/docs/get-details/accounts/rekey/).

You can issue a transaction to rekey an account by using the `algokit.rekeyAccount(rekey, algod)` function. The `rekey` parameter is an [`AlgoRekeyParams`](../code/interfaces/types_transfer.AlgoRekeyParams.md) object with the following properties:

- All properties in [`SendTransactionParams`](./transaction.md#sendtransactionparams)
- `from: SendTransactionFrom` - The account that will be rekeyed
- `rekeyTo: SendTransactionFrom | string` - The address of the account that will be used to authorise transactions for the rekeyed account going forward
- `transactionParams?: SuggestedParams` - The optional [transaction parameters](./transaction.md#transaction-params)
- `note?: TransactionNote` - The [transaction note](./transaction.md#transaction-notes)
- `lease?: string | Uint8Array`: A [lease](https://developer.algorand.org/articles/leased-transactions-securing-advanced-smart-contract-design/) to assign to the transaction to enforce a mutually exclusive transaction (useful to prevent double-posting and other scenarios)

```typescript
// Example
await algokit.rekeyAccount(
  {
    from: account,
    rekeyTo: newAccount,
    // Optionally specify transactionParams, note, lease and transaction sending parameters
  },
  algod,
)

const rekeyedAccount = algokit.rekeyedAccount(newAccount, account.addr)
// rekeyedAccount can be used to sign transactions on behalf of account...
```
