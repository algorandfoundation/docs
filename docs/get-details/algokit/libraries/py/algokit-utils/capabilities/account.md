# Account management

Account management is one of the core capabilities provided by AlgoKit Utils. It allows you to create mnemonic, idempotent KMD and environment variable injected accounts 
that can be used to sign transactions as well as representing a sender address at the same time.

(account)=
## `Account`

Encapsulates a private key with convenience properties for `address`, `signer` and `public_key`.

There are various methods of obtaining an `Account` instance

* `get_account`: Returns an `Account` instance with the private key loaded by convention based on the given name identifier:
  * from an environment variable containing a mnemonic `{NAME}_MNEMONIC` OR
  * loading the account from KMD ny name if it exists (LocalNet only) OR
  * creating the account in KMD with associated name (LocalNet only)

  This allows you to have powerful code that will automatically create and fund an account by name locally and when deployed against 
  TestNet/MainNet will automatically resolve from environment variables
 
* `Account.new_account`: Returns a new `Account` using `algosdk.account.generate_account()`
* `Account(private_key)`: Load an existing account from a private key
* `Account(private_key, address)`: Load an existing account from a private key and address, useful for re-keyed accounts
* `get_account_from_mnemonic`: Load an existing account from a mnemonic
* `get_dispenser_account`: Gets a dispenser account that is funded by either:
  * Using the LocalNet default account (LocalNet only) OR
  * Loading an account from `DISPENSER_MNEMONIC`

If working with a LocalNet instance, there are some additional functions that rely on a KMD service being exposed:
* `create_kmd_wallet_account`, `get_kmd_wallet_account` or `get_or_create_kmd_wallet_account`: These functions allow retrieving a KMD wallet account by name,
* `get_localnet_default_account`: Gets default localnet account that is funded with algos
