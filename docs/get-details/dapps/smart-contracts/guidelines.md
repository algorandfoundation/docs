title: Modern guidelines for smart contracts and smart signatures on Algorand

## Disclaimer

The content provided here is not an exhaustive list of security considerations. Before deploying any application in production, it is strongly advised to get the contract audited by one more third-party security companies.

## Use smart signatures very SPARINGLY

Algorand supports two types of programs on the blockchain: **smart signatures** and **smart contracts**. Smart signatures themselves can be used both as **delegated (logic) signatures** or **contract account**.

Most dApps should **only** use **smart contracts** and should **not** use any **smart signatures**. 

See the [section "More about smart signatures vs smart contracts"](#more-about-smart-signatures-vs-smart-contracts) for additional details.

!!! info
    As in the section below, smart signatures were necessary in the past. If you see an example (on the [developer portal](https://developer.algorand.org) or anywhere else online) that uses smart signatures, please take some time to see if it can be migrated following the instructions below.


## DO proper fee management

### DO NOT hard-code 1,000 microAlgos as minimum fee

The minimum fee per transaction is a consensus parameter. It **CAN** be changed over time. As such, **NEVER** hard-code its value anywhere.

If you hard-code 1,000 microAlgos instead of using one of the method below, your smart contract or code will fail, if the minimum fee is increased.

#### Get minimum fee in smart contract

=== "PyTeal"
    ```py
    Global.min_txn_fee()
    ```

=== "TEAL"
    ```teal
    global MinTxnFee
    ```

#### Get minimum fee off-chain with SDK

There are two options:

* using SDK constants:

=== "Python"
    <!-- ===PYSDK_CONST_MIN_FEE=== -->
	```python
	from algosdk import constants
	
	print(constants.MIN_TXN_FEE)
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/overview.py#L83-L86)
    <!-- ===PYSDK_CONST_MIN_FEE=== -->

=== "JavaScript"
    <!-- ===JSSDK_CONST_MIN_FEE=== -->
	```javascript
	const minFee = algosdk.ALGORAND_MIN_TX_FEE;
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/atomics.ts#L53-L54)
    <!-- ===JSSDK_CONST_MIN_FEE=== -->

=== "Java"
    <!-- ===JAVASDK_CONST_MIN_FEE=== -->
	```java
	System.out.printf("Min fee from const: %d\n", Account.MIN_TX_FEE_UALGOS);
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/Overview.java#L70-L71)
    <!-- ===JAVASDK_CONST_MIN_FEE=== -->
    
=== "Go"
    <!-- ===GOSDK_CONST_MIN_FEE=== -->
	```go
	log.Printf("Min fee const: %d", transaction.MinTxnFee)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/overview/main.go#L101-L102)
    <!-- ===GOSDK_CONST_MIN_FEE=== -->
    
* using an algod API:

=== "Python"
    <!-- ===PYSDK_SP_MIN_FEE=== -->
	```python
	suggested_params = algod_client.suggested_params()
	print(suggested_params.min_fee)
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/overview.py#L78-L80)
    <!-- ===PYSDK_SP_MIN_FEE=== -->

=== "JavaScript"
    <!-- ===JSSDK_SP_MIN_FEE=== -->
	```javascript
	// Not supported because getTransactionParams erases the information
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/atomics.ts#L63-L64)
    <!-- ===JSSDK_SP_MIN_FEE=== -->

=== "Java"
    <!-- ===JAVASDK_SP_MIN_FEE=== -->
	```java
	Response<TransactionParametersResponse> tpr = algodClient.TransactionParams().execute();
	TransactionParametersResponse sp = tpr.body();
	System.out.printf("Min fee from suggested params: %d\n", sp.minFee);
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/Overview.java#L50-L53)
    <!-- ===JAVASDK_SP_MIN_FEE=== -->
    
=== "Go"
    <!-- ===GOSDK_SP_MIN_FEE=== -->
	```go
	suggestedParams, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		log.Fatalf("failed to %s", err)
	}
	log.Printf("Min fee from suggested params: %d", suggestedParams.MinFee)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/overview/main.go#L93-L98)
    <!-- ===GOSDK_SP_MIN_FEE=== -->

### DO bound fees for smart signatures

In most scenarios, if you are using smart signatures, the transaction fee should be constrained by the smart signature, otherwise an attacker may be able to replay/send transactions with a much higher fee.

Often, the best solution is to require the fee to be 0 and rely on [fee pooling](../../../transactions/#pooled-transaction-fees): the other transactions (that are not signed by a smart signature) should provide enough fees to cover all transactions.

There are a few cases where forcing a 0 fee would not be appropriate: for example when a contract account is created to allow users to execute certain transactions.
In this case, the transaction fee needs to be reasonably bounded depending on the exact use case (e.g., five times the minimum fee).

### DO use 0 fee for inner transactions when possible

The simplest solution for inner transaction fee is to set them to 0 and rely on [fee pooling](../../../transactions/#pooled-transaction-fees): the other transactions (that are not signed by a smart signature) should provide enough fees to cover all transactions.

!!! tip
    Setting the inner transaction fee to 0 is different from not setting it. When it is not set, the inner transaction fee is set to the minimum allowable, taking into account the minimum transaction fee and credit from overpaying in earlier transactions. It is recommended to always set the transaction fee.

### DO handle congestion smoothly and securely

In case of congestion, the blockchain may require higher fees.

This is a complex subject with [a discussion in progress on the forum](https://forum.algorand.org/t/brainstorm-on-default-valid-window-length-and-lease/8208).

!!! tip
    Inner transaction fees are not subject to congestion fees. Inner transaction fees can always be set to the minimum fee (see guideline above how to access it) or to 0 (the recommended solution). When set to 0, the inner transaction will take the minimum fee from [fee pooling](../../../transactions/#pooled-transaction-fees).

## DO validate carefully incoming transactions

When validating an incoming transaction to an application account to validate payment for the action of a smart contract, do not forget to:

* Always check the `Type` or `TypeEnum` field to ensure the transaction is of the expected type.
* Always check the asset ID for asset transfer.
    * Smart contract developers may be tempted to optimize the contract by remarking that asset transfers can only work for assets that have been opted in by the application account. However, using the opt-in mechanism for that purpose is extremely error-prone. This is extremely strong discouraged.
* Always check the recipient is the application account.
* Always check the amount.


These guidelines are not exhaustive.

Similar guidelines apply to validate incoming transactions for smart signatures, in which case the recipient account may not be the application account.


## DO manage key carefully

If the smart contract is not updatable and the dApp is fully web3 (users interacting with it have their own wallet), then the only account needed for deployment is the creator account.
The simplest solution is most likely to deploy the smart contract with a throw away account.

The account may then rekey to an account without a known secret key (e.g., `737777777777777777777777777777777777777777777777777UFEJ2CI`, the reward pool) so that the account essentially is locked forever.
However, note that you must be sure the creator account will never be used for any other purpose (like validating the dApp before rekeying).

If the smart contract is updatable, then extreme care should be taken to secure the key allowing for example.
This document is not meant to provide exhaustive key management guidelines as key management is an extremely complex topic that requires an expert to find the proper trade-off for the specific applications.
However, here are some points:

* Do *NOT* use a hardware or software wallet without multisig. 
    * Note that using a hardware or software wallet with multisigs and multiple people holding a subkey may *NOT* be appropriate either. Appropriateness needs to be assesed by an expert considering all the aspects of the project
* Do *NOT* store secret keys (or mnemonics) in unencrypted/unsecure location such as environment variables.
* *NEVER* store secret keys (or mnemonics) in code (e.g., Github).

If the dApp requires the use of custodial accounts, then key management is even more complex.
In some cases, technical solutions can be used to limit risks, see for example [this forum post](https://forum.algorand.org/t/could-i-implement-an-asa-with-its-own-custom-wallet/3139/2).
However, as for the keys to update smart contract, it is strongly recommended to consult a security expert.

## DO follow these additional guidelines for smart contracts (and smart signatures)

### General

* Use [Beaker](https://github.com/algorand-devrel/beaker) instead of plain PyTeal or TEAL, unless you need to go in production very quickly. Note however that Beaker may introduce breaking changes.
* If using plain TEAL or plain PyTEAL (without Beaker):
    * Explicitly handling all `OnComplete` actions, including delete/update.
* Follow ARCs:
    * Strictly follow [ARC-4](https://arc.algorand.foundation/ARCs/arc-0004)
    * Check if there exists an [ARC](https://arc.algorand.foundation) standardizing what the smart contract will do.

* DO NOT hard-code the minimum balance requirements: those are consensus parameters that can change with consensus upgrades:
    * Ensure that your smart contract can work if the minimum balance requirements change.
    * Use the proper API endpoints or TEAL/PyTEAL opcode to read the minimum balance requirement of an account.
* DO NOT enforce group size as this prevents composability. But BE extremely careful to not allow your smart contract to consider twice the same payment transaction for the same method call. (An insecure example if to have a method A that looks for a payment transaction just before its call and a method B that looks for a payment transaction two transactions before. then the same payment transaction can be used to make both call A and B, which is in most cases insecure.) 

### Updatability and deletability

* Carefully decide whether the smart contracts are updatable / deletable:
    * Pros of updatable smart contracts
        * Can add features after creation
        * Can fix bugs after creation
    * Cons of updatable smart contracts
        * Can introduce bugs in updates
        * Generally means some sense of trust is involved
* Ensure that the smart contract cannot be deleted or made unusable while the application account still has funds.


### Local state, boxes, clear state

* Do **NOT** use local state for data that cannot be lost:
    * Local state can always be cleared by the user. There is no trustless way to recover data.
    * Ideally make clear to the user consequences of clearing user data (often this implies loss of funds).
    * Use boxes to store data that should never be cleared.
    * Be sure to allow the user opting out easily of the smart contract and recover their minimum balance.
* Write a clear program that properly handle clear state by users and that can never fail:
    * Do not use boxes or any foreign account/asset/application when clearing, as the user may not provide them in the transaction.
    * Remember that the state will be cleared even if the TEAL code fails.
    * A usual example is when the user sends some funds to the application account and the amount is recorded in local storage (e.g., for retrieval). In this case, a solution can keep in global storage a counter of the "lost funds" by users that cleared they state.
* Delete all boxes before deleting a smart contract:
    * Each box increases the minimum balance requirement of the application account. If the smart contract is deleted, the minimum balance is forever lost.
    * The above is true even if the application account is rekeyed to another account. Rekeying the application account to another account will allow to transfer all the funds out of the application account even if the smart contract is deleted, **except** for the minimum balance due to boxes.

### Oracle, interactions, and randomness

* Use randomness beacon such as [this one](https://developer.algorand.org/articles/usage-and-best-practices-for-randomness-beacon) to generate random values and follow strictly the [best practices](https://developer.algorand.org/articles/usage-and-best-practices-for-randomness-beacon). In addition:
    * Do **NOT** use the block hash or a PRG/Hash function (such as SHA256 or SHA3) on any block or transaction field information. This is **NOT** secure. For more information read the [randomness on Algorand article](https://developer.algorand.org/articles/randomness-on-algorand/).
    * Take some time to read very carefully and thoroughly the documents linked above to fully understand what is happening and what the trust model is. Before using a randomness oracle, check whether the required trust assumption is sufficient for your purpose. (As an extreme example, it may not be wise to use a randomness oracle for a 10B$ lottery.)
* Remember that smart contracts and smart signatures **CANNOT KEEP SECRETS** and **MUST NOT USE PASSWORDS**:
    * Everything that a smart contract operates on is made public to everyone on the blockchain. In particular, any input and output of the smart contract is public. (This is the case for all major blockchains.)
    * Oracles might be used but like for randomness beacons above, the exact trust model need to be studied in a lot of details. For example, Intel SGX may or may not be appropriate in certain cases, due to its known vulnerabilities.
    * Direct or indirect uses of passwords by smart contracts or smart signatures are 99.9% of the time insecure. If your dApp uses password, we strongly recommend that a cryptography expert fully review the design. Most of the documentation and examples on the Internet are wrong or misleading. (As an example amongst many others, looking for PBKDF2 will give examples with number of iterations is only 1,024 which is completely inadequate.)
* Remember that smart contract and smart signatures **CANNOT** interact by themselves with the outside world:
    * Smart contracts and smart signatures can only read on-chain data.
    * Oracles may provide a way for them to interact with the outside world. But as for the randomness beacon above, it is necessary to study the trust model. 

### Overflows and underflows

* Be particularly careful with **integer arithmetic**: overflow and underflow can cause issues. See for example the [(since corrected) TinyMan pool overflow errors](https://docs.tinyman.org/known-issues/2021-11-12-pool-overflow-errors). Any time integer arithmetic is used, do a semi-formal analysis proving that overflow and underflow can never happen.
    * A semi-formal analysis works by manually writing invariants such as "the value XYZ is always between 1 and 42" and proving (with pen and paper) these invariants hold whatever operation is done on the smart contract. "Proving" here really means "making a proper mathematical proof". "Handwaving" is strongly not recommended.
* Similarly as integer arithmetic overflow/underflow, do an analysis that that byte arrays never go beyond the limits of the AVM. This is a rarer issue but this still needs to be manually checked.



## DO follow these additional guidelines for dApps beyond smart contract and smart signature code

### Algod, indexer, and caching

* If using third-party API services:
    * Understand the trust model: if you do not run the node or indexer, you need to fully trust the provider to operate honestly. A malicious provider can return invalid answers.
    * Understand the SLA and maximum number load that the provider allows: free tiers may have limited SLA or limited throughput. If your dApp attracts many users, this may prevent proper functioning.
* For web3 application, use the algod/indexer endpoints provided by the user wallet if the user wallet provides it.
* Consider whether adding cache is necessary. Like for direct access to database in web2, direct access to blockchain will never be as cheap and as fast as serving cached answers.
    * Note that cache does not make a dApp non-web3. For example, an NFT galleries may use a cache to quickly display NFTs but then check the hashes of the retrieved cache items against the blockchain (ideally asynchronously).

### Standard security guidelines

Blockchains do not automatically make something secure.
Traditional security guidelines still need to be strictly followed.
Actually, due to the nature of blockchains, those guidelines are even more important: in a web3 project, a bug can easily cost millions of dollars without possibility to recover the money through the justice system.

Here we highlight some of them (the list is non-exhaustive):

* **NEVER** store any key or mnemonic in source code. It is almost impossible to securely remove a key or mnemonic from a git repo.
* Perform regular audits.
* Have a security expert review any security critical part of the system, particularly key management and anything cryptography related (including any use of password).
* Ensure highest quality of code, following best practices (including outstanding documentation, full test coverage, proper CI). Review [classical security issues](https://owasp.org/www-project-top-ten/).
* Ensure proper safeguards preventing a single employees to merge and deploy code for example. 
* Actively monitor both the dApp on-chain and all the systems around the dApps. Have a response plan as well as a person available 24/7 to mitigate any risk (especially for high value DeFi apps).
* Have mitigation systems in place:
    * Some projects have some form of kill switch for example at the beginning. The kill switch is sometimes disabled after some time.
* Actively limit dependencies and attack surface. Constantly monitor used libraries for any exploit and update as soon as possible. 
* If using servers, follow [NIST recommendation](https://nvlpubs.nist.gov/nistpubs/Legacy/SP/nistspecialpublication800-123.pdf ).
* Operations should typically fail rather than default to insecure modes. For example if the keys are unavailable, it is usually better to fail than to store sensitive material in the clear. When failing is not an option, at the very least there has to be a very visible warning displayed.

## More about smart signatures vs smart contracts

### Issues with smart signatures

* Smart signatures are much harder to audit than smart contracts in most settings.
* Smart signatures cannot be used in inner transactions from smart contracts, which strongly limits composability of smart contracts (at least for TEAL v8 and before).
* Smart signatures are much less flexible than smart contracts. While some simple dApps could be based on smart signatures and using smart signatures could slightly reduce fees, adding any feature would become problematic and any upgrade would most likely be impossible.
* Contrary to smart contract, smart signatures do not have a standardized ABI (Application Binary Interface). Smart contracts have [ARC-4](https://arc.algorand.foundation/ARCs/arc-0004).
* (For delegated logic signatures only) Most wallets do not support signing delegated logic signatures, as this is a potentially dangerous operation: a malicious delegated logic signature can remain dormant for years and can allow to siphon out funds of an account much later. There is no way to check whether an account signed a malicious delegated logic signature until the actual fund siphoning occurs.

#### Examples comparing smart signature and applications

To highlight how easier applications are to analyze comparent to smart signatures, see the following PyTEAL examples.

A smart signature like

```python
return Seq(
  Assert(Txn.Amount() == Int(1)),
  Assert(Txn.TypeEnum() == TxnType.Payment),
  Assert(Txn.Receiver() == intended_receiver)
)
```

has the following issues:
* This contract can get drained via `CloseRemainderTo`.
* This contract can lose authorization via `RekeyTo`.
* This contract can get drained via fees.

Compared with the following application

```python
return Seq(
  InnerTxnBuilder.Begin(),
  InnerTxnBuilder.SetFields({
      TxnField.type_enum: TxnType.AssetTransfer,
      TxnField.receiver: intended_receiver,
      TxnField.amount: Int(1)
  }),
  InnerTxnBuilder.Submit(),
)
```

where:
* Contract account can't be closed
* Can't be rekeyed
* Fee will always be minimum fee required

### (Niche) use cases for smart signatures

There are three main (niche) use cases for smart signatures:

* As a contract account, smart signatures can be used when very costly operations need to be performed, such as `ed25519verify` or `vrf_verify`. A smart contract needing to perform such operations would require to be called (in a group of transactions) just after a transaction from the contract account. 
    * Importantly, this mechanism can only be used for smart contract methods that are not meant to be called from another smart contract, as smart contracts currently cannot issue inner transaction signed by a logic signature.
    * An alternative solution is to [increase the opcode budget](../../../dapps/avm/teal/#dynamic-operational-cost-of-teal-opcodes) of the smart contract using group transactions and inner transactions. PyTeal actually includes a [budget increase utility](https://pyteal.readthedocs.io/en/stable/opup.html). This method however is more expensive in terms of transaction fees, compared to using a smart signature contract account. ON the other hand, it has the advantages of being easier to code and of not breaking composability.
* As a contract account, smart signatures can be used to allow any user (even without a blockchain account) to send certain transactions for free:
    * This is only meaningful if there are only a limited number of transactions that can be executed in a given amount of time. 
    * A concrete example would be a smart signature allowing to call a smart contract acting as a third-party blockchain light client. Only when the third-party blockchain has new valid blocks can the smart contract be called succesfully.
* As a delegated logic signature, smart signatures can be used in complex systems to delegate certain operations on an account to another signature key or another smart contract. Here are a few examples:
    * Delegation of the issuance of "change offline" transactions to a lower-security key to give the ability to very quickly send "change offline" transactions in case of a participation node issue.
    * Delegation of ASA transfers to a "hot key." This allows using the Algorand blockchain to record token ownership in a custodial system (e.g., for a loyalty reward program). The actual keys of the custodial accounts can be kept in cold storage, while the system uses hot keys whose only ability is to transfer tokens between custodial accounts. This way, the system is no less secure than a centralized system with a centralized DB: in case of leak of keys, the only thing the attacker can do is to transfer the tokens incorrectly, but the attacker is not able to rekey accounts or steal Algos (if Algos are used for any other purpose).

Note that the second use case can often be done using a smart contract and inner transactions. However, a delegated logic signature is often easier to write and use for that purpose.

### Migration from smart signatures to smart contracts

In the past, smart signatures were used as companions of smart contracts in the following scenarios:

* **Escrow accounts** for smart contracts before inner transactions (AVM 1.0 / TEAL v5). Now inner transactions should be used. Notes:
    * Since AVM 1.0 / TEAL v5, escrow accounts can be either a smart signature contract account or an application account (associated to a smart contract). Application accounts use inner transactions and are preferred.
    * There are a few rare corner cases where inner transactions cannot be used. For example TEAL v8 does not allow making an inner transaction from a (smart signature) contract account nor allow making an application call inner transaction for a TEAL v2/v3/v4 smart contract. In those cases, it may be needed to use a contract account as escrow account. However, it is strongly recommended to restrict the use of such an escrow account to the rare method calls that need it. Indeed, such method calls would themselves not be callable from inner transactions perpetuating the above issue.
    * In case multiple escrow accounts are needed (instead of a globa escrow account), the simplest solution is to rekey accounts to the application account.
    * Before AVM 1.1 / TEAL v6, only a small subsets of inner transactions were supported, making use of contract accounts for escrow accounts sometimes necessary. This is no more an issue.
* **Un-erasable local storage** for smart contracts before boxes (TEAL v8). Local storage can be erased at any time by the account associated to the local storage. Sometimes, this is not a desirable property. In the past, special contract accounts holding local storages may have been used. Boxes are now the proper solution.
* **Infinite storage** for smart contracts before boxes (TEAL v8). Global storage is limited and the solution to get more global storage in the past was to use local storage of contract accounts (or accounts rekeyed to the application account). Boxes are now the proper solution.

### Note on the terminology

Smart signatures are also called **logic signature** or **LogicSig**.

Smart signatures were also called in the past **stateless smart contracts** or just "Algorand smart contracts (ASC1)" before TEAL version 2. 
In particular, any TEAL program that does not start with `#pragma version <N>` is a TEAL v1 smart signature and should not be used at all.

## Additional guidelines specific to smart signatures

!!! info
    As mentioned above, smart signatures should only be used in very few niche use cases. Most likely you should not be using them. These guidelines are only there for these very few niche use cases.
   
These guidelines are specific to smart signatures. They complement the other guidelines in this document. Smart signatures should also obey the other guidelines.
   
* **Only use smart signatures when absolutely required.** See section above for details.
* Do not use TEAL v1 smart signatures (i.e., TEAL codes not starting with any `#pragma version ...`):
    * If you need to migrate a TEAL v1 smart signature to a more recent version, be extra careful. In particular, add a test checking that `RekeyTo` is the `ZeroAddress`. Not doing so most likely allows any attacker to 
* Check as many transaction fields (accessible via `txn`/`Txn` in TEAL/PyTEAL) as possible to constrain the transactions approved by the smart signature as much as possible to avoid security issues. In particular (this is not an exhaustive list):
    * Always check the `Type` or `TypeEnum` field to ensure the transaction is of the expected type.
    * Always check the following fields: `RekeyTo`, `CloseRemainderTo` (for payments), and `AssetCloseTo` (for asset transfers).
        * For most cases, these fields must be set to the `ZeroAddress`.
    * Follow the guidelines for transaction validation and for fee management above.
* Be careful that the same smart signature can be used for multiple transactions in a group:
    * Either check the current transaction fields using `txn` / `Txn.` (instead of `gtxn` / `Gtxn`);
    * Or use `gtxn` / `Gtxn` for an index `i` and verify that the current transaction is the `i`-th ().
    * Not doing the above can allow an attacker to call twice the smart signature in the same group: the first transaction obeys the rule of the smart signature and is safe, but the second one is unconstrained and can allow the attacker to siphon out all funds.
* If the use case permits it, set up an expiration mechanism:
    * Contract accounts should allow the sender to reclaim their assets after a certain round/time (`txn FirstValid > NNNN` or `txn FirstValidTime > TTTT`).
    * Delegated logic signatures should specify an expiration round in the future after which they will no longer function.
* Be extra careful when combining multiple contract accounts in a group of transaction and to perform multiple checks. The simplest solution is to enforce the transaction group size 
* Remember that, as for smart contracts, the smart signature code, the transaction fields, and the arguments of the smart signature are all **public**. In particular:
    * An attacker can replay a transaction signed by a smart signature (whether in contract account or logic signature mode).
    * For example, a delegated logic signature allowing approving any change offline transaction is insecure. The first time it is used, an attacker can replay it many times and empty the account (using all the fees)! The delegated logic signature needs to 
* Be careful about the fact that, contrary to arguments of smart contracts, arguments of smart signatures are **NOT** signed by the sender account and are **NOT** part of the computation of the group ID. In other words, an attacker can intercept a transaction signed by a smart signature and change its arguments (as long as it is still accepted by the smart signature).
* Be careful about the fact that the same smart signature can be used in multiple networks.
  * For instance, if a smart signature is signed with the intent of using it on TestNet, that same transaction can be sent to MainNet with that same smart signature.
  * Always use new accounts when using TestNet (or any other network) to avoid reusing the same account that signed a smart signature.
  * Make sure to write a smart signature that checks which network it is being run on. You can use:

=== "PyTeal"
    ```py
    Global.genesis_hash()
    ```

=== "TEAL"
    ```teal
    global GenesisHash
    ```

## Additional resources

* [Trails of Bits guideslines](https://github.com/crytic/building-secure-contracts/tree/master/not-so-smart-contracts/algorand)
* [ASC Security Guidelines from Joe Polny](https://github.com/joe-p/algo-edu/blob/master/resources/en-US/asc_security.md) on which this document is partially based