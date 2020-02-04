title: Signatures

Every transaction must be authorized by the [sender](../reference-docs/transactions.md#sender) of before it is submitted to the network. Authorization occurs with the addition of a **signature** to a transaction object. More specifically, a transaction object is wrapped in a `SignedTxn` object that includes the [transaction](../reference-docs/transactions.md#txn) and the [signature](../reference-docs/transactions.md#sig).

In the [Transactions Guide](./transactions.md), you learned how to construct a transaction. In this section, you will learn the different ways to sign, or authorize, those transactions. 


# Single Signatures
A single signature corresponds to a signature from the private key of an [Algorand public/private key pair](./accounts.md#keys-and-addresses).

This is an example of a transaction signed by an Algorand private key displayed with `goal clerk inspect` command:

```json
{
  "sig": "ynA5Hmq+qtMhRVx63pTO2RpDrYiY1wzF/9Rnnlms6NvEQ1ezJI/Ir9nPAT6+u+K8BQ32pplVrj5NTEMZQqy9Dw==",
  "txn": {
    "amt": 10000000,
    "fee": 1000,
    "fv": 4694301,
    "gen": "testnet-v1.0",
    "gh": "SGO1GKSzyE7IEPItTxCByw9x8FmnrCDexi9/cOUJOiI=",
    "lv": 4695301,
    "rcv": "QC7XT7QU7X6IHNRJZBR67RBMKCAPH67PCSX4LYH4QKVSQ7DQZ32PG5HSVQ",
    "snd": "EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4",
    "type": "pay"
  }
}
```
This transaction sends 1 Algo from `"EW64GC..."` to `"QC7XT7..."` on TestNet. The transaction was signed with the private key that corresponds to the `"snd"` address of `"EW64GC..."`. The base64 encoded signature is shown as the value of the [`"sig"`](../reference-docs/transactions.md#sig) field.

**How-To**

- [Sign a transaction with your private key](../getting-started/tutorial.md#sign-the-transaction)

# Multisignatures

When the [sender](../reference-docs/transactions.md#sender) of a transaction is the address of a [multisignature account](./accounts.md#multisignature) then the transaction requires authorization through an `"msig"` object that contains subsignatures from the associated private keys of the addresses that multisignature account is composed of. The number of sub-signatures must be equal to or greater than the threshold value of the multisignature account. See [Accounts - Multisignature](./accounts.md#multisignature) for details on how to configure a multisignature account.

!!! important
	Upon signing, either the signing agent or the transaction needs to know the composition of the multisignature account, i.e. the ordered addresses, threshold, and version. 

This is an example of a transaction signed by the 2 of the 3 private keys displayed with `goal clerk inspect` command:


  

# Logic Signatures

Logic Signatures authorize transactions associated with Algorand Smart Contracts. 
Link to [ASC1 SDK Usage](../feature-guides/asc1/asc1_sdk_usage.md)