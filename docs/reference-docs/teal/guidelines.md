title: Guidelines

The following list of guidelines should be followed when working with TEAL.

* Use or modify example templates which are designed to embody best practices.

* TEAL code should validate as many `txn` fields as possible. Things not checked could become set to anything.

* `CloseRemainderTo` or `AssetCloseTo` should be the intended recipient or equal to global `ZeroAddress`. An unchecked address could steal all the value!

* When not closing out a contract account, `Amount` or `AssetAmount` should be checked to be equal to the intended transfer amount.

* `Fee` should be less than some reasonable amount (in micro-Algos). An `unchecked` Fee could burn the entire value of the contract account!

* `Type` or `TypeEnum` can be checked to ensure something is a `Payment` or `AssetXfer` or another transaction type.

* `Sender` doesn’t need to be checked. It gets checked automatically.

* Using a contract account limits your exposure as compared to a delegated account. (Mistakes and bugs can only lose as much as you put into that account. If you can close out that account it closes out any exposure.)

* Counterpoint: using a regular private/public key account with delegated logic allows you to always use the private key to sign any transaction and not be locked into only what the logic allows.

* Contract accounts should have an expiration in the future (`txn` `FirstValid` > NNNN, or `txn` `FirstValidTime` > TTTT) after which the originator can reclaim their assets.

* Delegated LogicSigs may wish to specify an expiration round in the future after which they will no longer function.

* Transaction groups must be designed with ‘slots’ such that each transaction will know what position it is in. A transaction’s logic can check (`txn` `GroupIndex` == K && `global` `GroupSize` == M) and then also check specific other transactions by their index through `gtxn` N field

Example:

`txn[0]` sends N Algos from A to B; program checks that (M/N >= desired exchange rate) and (`gtxn` 1 `Receiver` == A)

`txn[1]` sends M AltCoin from B to A; program checks that (N/M >= desired exchange rate) and (`gtxn` 0 `Receiver` == B)

* Corollary to checking as many `txn` fields as possible: be strict with other transaction properties in a group.

* Verify the `GroupSize` to make sure the size corresponds to the number of transactions the logic is expecting. 

# Tips and Patterns
In addition to following the guidelines, the following tips and patterns may also be useful. 

* A contract account that requires `CloseRemainderTo` or `AssetCloseTo` being equal to the intended recipient is good for a conditional all-or-nothing payment.

* A recurring payment contract should combine (`FirstValid` % N == 0) and (`Lease` == xxxxx), where N is the period of the recurring payment in rounds, at least 1000.

* Some operations panic and exit the transaction validation with an immediate fail. For example a divide-by-zero. It’s possible to test for these and skip panicking instructions with a `bnz` branch-if-not-zero instruction.