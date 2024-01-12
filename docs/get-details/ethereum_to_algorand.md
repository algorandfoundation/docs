title: From Ethereum to Algorand

This guide is meant for Ethereum **developers** who want to switch to Algorand to benefit from its numerous advantages: faster block time (less than 4s), immediate finality (a transaction can never be reverted once in a block), very low fees, and high throughput, while still being highly secure.

It is focused on **dApp developers** (as opposed to end users and node runners).

If you want to find the equivalent of any Ethereum specific notion/tool/service, just search on this page and you will most likely find the equivalent notion/tool/service on Algorand.

# Main High-Level Differences

In this section, we highlight the main differences between Ethereum and Algorand.

## Accounts and Smart Contracts

Both Ethereum and Algorand are account-based blockchains supporting smart contracts.

On Ethereum, smart contracts are actually represented by accounts, while on Algorand smart contracts and accounts are two different objects.
Concretely:

* Ethereum's **Externally-owned accounts (EOA)** correspond to Algorand accounts. They are both represented by an **address**.
    * Example of Ethereum address:
        * User-friendly representation: `0x65e9980679DE55744f386aa1999307f1687A92f9`
        * Raw address: 20 bytes
    * Example of Algorand address:
        * User-friendly representation: `QD3BO4RMWXBOZIPHTGGB3RSKSOAKOHM2HGN7QDZXH4ECBGJRIU3AHHC3JU`
        * Raw address: 32 bytes
* Ethereum's **contract accounts** correspond to Algorand **application ID**, which are 64-bit integers. (Algorand applications also have an associated application account/address, see below.)
    * Example of Ethereum contract account:
        * User-friendly representation: `0xbc4ca0eda7647a8ab7c2061c2e118a18a936f13d`
        * Raw: 20 bytes
    * Example of Algorand application ID:
        * User-friendly representation: `947957720`
        * Raw: uint64

!!! warning
    Algorand also has a notion of **contract account** which is controlled by a [smart signature](../dapps/smart-contracts/smartsigs/).
    Smart signatures are akin to [account abstraction](https://eips.ethereum.org/EIPS/eip-4337) on Ethereum and are only meant to be used in more advanced scenarios.

Algorand smart contracts are called **applications**. Each application has an associated [**application account**](../dapps/smart-contracts/apps/innertx/) (or several accounts if the rekeying feature is used) that can send and receive tokens. The application account's address can be computed from the application ID. See below how to transfer tokens to an application.

Like Ethereum smart contracts, Algorand applications can make transactions from their application accounts. These transactions are called **inner transactions** on Algorand.

## Fungible and Non-Fungible Tokens (FT and NFT)

Ethereum allows creating custom fungible and non-fungible tokens by deploying smart contracts following ERC-20, ERC-721, or ERC-1155 standards. Transacting with such tokens is very different from transacting the base cryptocurrency Ether.

On Algorand, such custom tokens are called [Algorand Standard Assets](../asa/) and are backed in the protocol itself. They do not require writing a smart contract and transferring them is similar to transferring the base cryptocurrency with one main difference: [mandatory opt-in](../asa/#receiving-an-asset). Opting in to an ASA is done by making a 0 amount asset transfer of the ASA from the account opting in to itself. This helps reducing spam from unwanted ASAs. It also has an impact on the [minimum balance](../accounts/#minimum-balance) (see below).

Tokens on Ethereum are defined by a contract address (+ an ID for ERC-1155 tokens). On Algorand they are just defined as a 64-bit unsigned integer ID.

## Gas, Transaction Fees, and Minimum Balance

### Gas (Ethereum)

On Ethereum, transaction fees are called gas fee. Gas fee is paid whether the transaction is successful or not.

### Transaction Fees (Algorand)

On Algorand, transaction fees are only paid if the transaction is included in the block. Transaction fee rules are explained in the [transaction structure documentation](../transactions/#fees). Since the max TPS of Algorand is much higher than Ethereum, congestion is much less likely, and almost all transactions can use the minimum base fee of 0.001 Algo per transaction.

This minimum transaction fee is independent of the transaction type: application call, Algo transfers, and ASA transfers for example all cost the same when there is no congestion. However, complex smart contracts may require to make "dummy"/"opup" transactions to increase their computational budget (which can be done automatically using [PyTeal OpUp](https://pyteal.readthedocs.io/en/latest/opup.html)).

### Minimum Balance

In addition to transaction fees, Algorand also has a notion of **minimum balance**. At a high-level, stored data (account balances, application states, ...) on Algorand is always associated to an account. Every time the amount of stored data increases (e.g., opt in to an ASA or application, storage of extra data as boxes in a smart contract, ...), the minimum balance requirement of the associated account increases.

The minimum balance acts like a deposit to rent space on the blockchain. If the space is liberated (e.g., opt out of the asset), the minimum balance requirement decreases. A basic account has a minimum balance requirement of 0.1 Algo. Opting in an asset for example, increases this requirement by an additional 0.1 Algo.

In more detail, here are the documentation pages discussing the minimum balance requirements:

* [in general](../accounts/#minimum-balance)
* [for ASAs](../asa/#assets-overview)
* [for smart contracts](../dapps/smart-contracts/apps/#minimum-balance-requirement-for-a-smart-contract)
* [as a summary table](../parameter_tables/)

!!! info
    Even if users never liberate space and the minimum balance requirement is essentially acting as a fee that users pay, as of March 2023, the resulting costs of transacting on Algorand is still orders of magnitude lower than the costs of transacting on Ethereum.

## Availability of Resources to Smart Contract

Algorand is designed to ensure high TPS and low latency. Every access to the blockchain state is very costly (in time). Therefore, to ensure blockchain state access is not in the critical path of block evaluation and to limit its cost, Algorand applications are restricted in the amount of resources (account balance/state, application state, ...) they can access. Furthermore, these resources need to be specified inside the transaction. This is to allow nodes to pre-fetch data.

While this may seem a strong restriction, most smart contracts do not dynamically decide which resources to access depending on the current blockchain state. In other words, it is possible to know in advance which resources will be accessed.

In the near future, the [simulate endpoint](https://medium.com/algorand/try-before-you-buy-on-algorand-5acd1b9617d1) will allow a dApp to easily and dynamically learns the [list of resources](../dapps/smart-contracts/apps/#reference-arrays) that need to be added to a transaction.

## Smart Contract Storage

One important difference between Ethereum and Algorand smart contracts is storage.

Ethereum smart contract storage is a huge array of 2^256 uint256 elements. The solidity language has higher-level types like dynamic arrays and mappings that are then mapped to this huge storage array (dynamic types use keccak to compute the location of each item).

For performances reasons, Algorand smart contracts have [three different types of storage](../dapps/smart-contracts/apps/state). While it is possible to only use boxes and essentially have a similar model as Ethereum (with the caveat that the boxes used need to be specified in the transaction), it can be more cost-effective to use local and global storage in some cases.

In particular, the common solidity pattern

```solidity
mapping (address => uint) public balances;
```

is often better replaced by local storage. (With the caveat that local storage can always be erased by the account holder using [ClearState transactions](../dapps/smart-contracts/apps/). So care should be taken that doing so would not put the smart contract in an insecure state.)

## Unique Features of Algorand: Multisig, Atomic Transfers, Rekeying

### Multisig Accounts

On Ethereum, it is possible to write smart contracts to ensure that fund transfers require approval/signatures by multiple distinct users. On Algorand, multisig accounts are first-class citizens and can be created very easily. See [the multisig account documentation](../accounts/create/#multisignature).

### Atomic Transfer / Group Transaction

Atomic transfers or group transactions allow the grouping of multiple transactions together so that they either all succeed or all fail. This can allow two users to securely exchange assets without the risk of one of the users failing to fulfill their side of the transaction.

Group transactions are also used a lot by smart contracts. For example, to send tokens to a smart contract, it is common to group a token transaction to the application account with an application call.

### Rekeying

Rekeying is a powerful protocol feature which enables an Algorand account holder to maintain a static public address while dynamically rotating the authoritative private spending key(s). See [the rekeying documentation](../accounts/rekey/).

There is no direct equivalent on Ethereum although this can be simulated using a smart contract and/or account abstraction.

## Nonces, Validity Windows, and Leases

Ethereum uses nonces to prevent transaction from being replayed.

Algorand does not have nonces.  Instead, two identical transactions cannot be committed to the blockchain. In addition, transactions have a validity window and optional leases. The [validity window (aka first/last valid rounds)](../transactions/#setting-first-and-last-valid) specifies between which rounds a transaction can be committed to the blockchain.

If the same transaction needs to be executed twice, some field needs to be changed. One option is to add a random note field or to slightly change the validity window.

[Leases](https://developer.algorand.org/articles/leased-transactions-securing-advanced-smart-contract-design) provide more fine-grained ways of preventing duplicated transactions from happening and are mostly used in conjunction with [smart signatures](../dapps/smart-contracts/smartsigs/) in very advanced scenarios.

Most dApp developers are unlikely to need to use leases and smart signatures.

## Re-Entrancy

Algorand is not susceptible to most re-entrancy attacks for multiple reasons:

1. Application calls and payment/asset transfer transactions are different. When an application transfers tokens to another application account or to a user account, it does not trigger any code execution.
2. An application cannot make (directly or indirectly) an application call to itself.

# Design Patterns

In this section, we go over common design patterns Ethereum uses and their equivalent solutions on Algorand.

## Transfer Tokens to an Application

On Ethereum, transferring tokens to a smart contract is done in two ways:

1. For Ether, the tokens are directly sent with the call to the smart contract.
2. For other tokens (ERC-20, ERC-721, ERC-1155), the user first needs to call a function (of the token smart contract) to approve the smart contract they want to call to spend tokens on their behalf.

On Algorand, transferring tokens is similar whether the tokens are the Algo or an ASA. It is also more explicit. The user typically creates a group of two transactions: the first one transfers the token to the application account and the second one calls the application.

!!! tip
    The high-level [Beaker framework](https://github.com/algorand-devrel/beaker) hides the details of those transfers. Beaker provides natural ways for a smart contract to specify that it needs to receive certain tokens, as well as client SDKs to easily make the required group of transactions.

## Proxy

Proxy smart contracts are heavily used on Ethereum as Ethereum smart contracts are not updatable.

Whereas on Algorand, applications can specify arbitrary rules for whether they can be updated or deleted.

This is strictly more general and flexible than on Ethereum: Algorand applications can indeed prevent any update and deletion like Ethereum smart contracts.

The proxy design pattern may still be useful on Algorand if you want to provide the option to the user to decide whether they only ever want to use a non-upgradable smart contracts (calling the smart contract directly) or an upgradable one (calling the proxy). A proxy can also be useful to split smart contracts that are too large.

## Pull Over Push

On Algorand like on Ethereum, you may want to consider the pull-over-push pattern whenever the smart contract needs to make multiple transfers in one application call.

While accounts on Algorand cannot reject Algo transfers, token transfers can fail for various reasons, including (but maybe not limited to):

* if the receiver account does not exist and less than 0.1 Algo are transferred to it, the transaction will fail due to the minimum balance requirement
* if the receiver account did not opt in to the ASA being transferred, the transaction will fail.

## Factory

The factory pattern is possible on Algorand though it is very rare. In general using a big application is easier.


# Glossary

## Accounts and Applications

| Ethereum                       | Algorand                          | Notes                                                                                                |
| ------------------------------ | --------------------------------- | ---------------------------------------------------------------------------------------------------- |
| externally-owned account (EOA) | account                           |                                                                                                      |
| contract account               | application / application account | Algorand applications are not accounts but have an associated application account to receive tokens. |
| smart contract                 | smart contract / application      |                                                                                                      |
| account abstraction            | smart signature contract account  |                                                                                                      |

## Data Types

| Ethereum              | Algorand                                                                                                                                              | Notes                                            |
| --------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------ |
| storage               | [global storage, local storage, boxes](../dapps/smart-contracts/apps/state/)                                                                          | See section above about storage                  |
| memory                | scratchspace ([TEAL](../dapps/avm/teal/#storing-and-loading-from-scratchspace) / [PyTeal](https://pyteal.readthedocs.io/en/stable/scratch.html))      | Much like Ethereum, the stack can also be used to store temporary values |
| environment variables | [txn](../dapps/avm/teal/opcodes/#txn-f) / [Txn](https://pyteal.readthedocs.io/en/stable/accessing_transaction_field.html#id1)                         | For data about the current transaction           |
|                       | [gtxn](../dapps/avm/teal/opcodes/#gtxn-t-f) / [Gtxn](https://pyteal.readthedocs.io/en/stable/accessing_transaction_field.html#atomic-transfer-groups) | For data about other transactions in the group   |
|                       | [global](../dapps/avm/teal/opcodes/#global-f) / [Global](https://pyteal.readthedocs.io/en/stable/accessing_transaction_field.html#global-parameters)  | For other data                                   |

## Functions, Methods, Subroutines

| Ethereum                 | Algorand                                                                                                                                                                                                                                                      | Notes                                            |
| ------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------ |
| internal function        | subroutine ([TEAL](https://developer.algorand.org/docs/get-details/dapps/avm/teal/#looping-and-subroutines) / [PyTeal](https://pyteal.readthedocs.io/en/stable/control_structures.html#creating-subroutines))                                                 |                                                  |
| external function        | [method](https://developer.algorand.org/docs/get-details/dapps/smart-contracts/ABI/#methods) ([PyTeal](https://pyteal.readthedocs.io/en/stable/abi.html#registering-methods) / [Beaker](https://algorand-devrel.github.io/beaker/html/usage.html#decorators)) | Beaker handles method seamlessly                 |
| view function            | [read-only method](https://arc.algorand.foundation/ARCs/arc-0022)                                                                                                                                                                                             |                                                  |
| constructor              | [create transaction](https://developer.algorand.org/docs/get-details/dapps/smart-contracts/apps/create/)                                                                                                                                                      | Beaker handles this seamlessly                   |
| public/private functions | n/a                                                                                                                                                                                                                                                           | No notion of derived smart contracts on Algorand |

## Misc

| Ethereum | Algorand | Notes |
| -------- | -------- | ----- |
| events / logs | logs ([TEAL](../dapps/avm/teal/opcodes/#log) / [PyTeal](https://pyteal.readthedocs.io/en/stable/api.html#pyteal.Log)) |  |

## Standards / ERC / ARC

The equivalent of ERC on Algorand are [ARC](https://arc.algorand.foundation).

| Ethereum          | Algorand              | Notes                                                               |
| ----------------- | --------------------- | ------------------------------------------------------------------- |
| ERC-20            | ASA / ARC-3 (+ ARC-19)           | ARC-3 is a convention for the metadata of ASA, ARC-19 can be used when the metadata is updatable                       |
| ERC-20            | ASA / ARC-20            | ARC-20, aka "smart ASA", defines the interface to control an ASA through a Smart Contract (the ASA is used for accounting, the Smart Contract to transfer, freeze, etc a-la ERC-20) |
| ERC-781, ERC-1155 | ASA / ARC-3 (+ ARC-19) or ARC-69 | ARC-3 and ARC-69 are two conventions for the metadata of an ASA NFT, ARC-19 can be used when the metadata is updatable |

# Tools and Services

This is a non-exhaustive list of tools and services used by Ethereum developers, with some of their equivalents on Algorand (non-exhaustive, in alphabetical order).

*Disclaimer*: The list below is not an endorsement of any of the tools, services, or wallets named or linked. As in all the developer documentation, this information is purely for educational purpose. In no event will Algorand or Algorand Foundation be liable for any damages of any kind (including, loss of revenue, income or profits, loss of use or data, or damages of any sort) arising out of or in any way connected to this information. You understand that you are fully responsible for the security and the availability of your keys.

|                              | Ethereum               | Algorand                                                                                                                                                                                                   | Notes                                                                                                 |
| ---------------------------- | ---------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| block explorer               | Etherscan              | [list of block explorers](https://developer.algorand.org/ecosystem-projects/?tags=block-explorers)                                                                                                         |
| API service                  | Infura                 | [algonode.io](https://algonode.io), [BlockDaemon](https://www.blockdaemon.com/protocols/algorand), [GetBlock.io](https://getblock.io/nodes/algo/)                       | Algorand also provides an official [indexer](../indexer/) software, that these services provide access too |
| wallet                       | Metamask               | mobile wallets with PeraConnect/WalletConnect ([Pera](https://perawallet.app/), [DeFly](https://defly.app/)), [MyAlgoWallet](https://wallet.myalgo.com/) with [MyAlgoConnect](https://connect.myalgo.com/) |                                                                                                       |
| development environment      | Truffle Suite, Hardhat | [AlgoKit](https://github.com/algorandfoundation/algokit-cli)                                                                                                                                               |
| one-click private blockchain | Ganache                | [sandbox](https://github.com/algorand/sandbox)                                                                                                                                                             | [AlgoKit](https://github.com/algorandfoundation/algokit-cli) uses sandbox and is recommended          |
