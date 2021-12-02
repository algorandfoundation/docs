title: Schema

# Background

- [WalletConnect](https://docs.walletconnect.org/) is an open protocol to communicate securely between mobile wallets and decentralized applications (dApps) using QR code scanning (desktop) or deep linking (mobile). It’s main use case allows users to sign transactions on web apps using a mobile wallet.
- v1 of WC (currently deployed) has first-party support for Ethereum chains, but it can be extended to other chains through custom message schema. [Binance Chain does this with their custom schema](https://docs.binance.org/guides/concepts/walletconnect.html#protocol-differences).

# Purpose

The purpose of this document is to define a custom WalletConnect schema for Algorand. The schema in this document is based on the [Algorand Foundation’s Wallet Transaction Signing API](https://github.com/algorandfoundation/ARCs/blob/main/ARCs/arc-0001.md), and it attempts to be as similar to that as possible. The main contribution of this document is to define how that schema can be used with WalletConnect.

# Schema

!!! Note
    All interfaces are defined in TypeScript. These interfaces are designed to be serializable to and from valid JSON objects.

A WalletConnect schema is a set of JSON-RPC 2.0 requests and responses. WalletConnect will send requests to the Algorand Wallet and will receive either signed transactions or failures as responses. All requests adhere to the following structure:

```
interface JsonRpcRequest {
  id: number;
  jsonrpc: "2.0";
  method: string;
  params: any[];
}
```

A successful request will return a response that adheres to the following structure:

```
interface JsonRpcResponse {
  id: number; // will be the same as the request's id
  jsonrpc: "2.0";
  result: any;
}
```

The Algorand schema consists of the requests and responses below.

## algo_signTxn

This request is used to ask a wallet to sign one or more transactions in an atomic transaction group.

### Request

This request adheres to the following structure:

```
interface AlgoSignTxnRequest {
  id: number;
  jsonrpc: "2.0";
  method: "algo_signTxn";
  params: SignTxnParams;
}
```

The parameters `SignTxnParams`, are defined as:
`type SignTxnParams = [WalletTransaction[], SignTxnOpts?];`

!!! Note 
    `SignTxnParams` is a [tuple with an optional element](https://www.typescriptlang.org/docs/handbook/release-notes/typescript-3-0.html#optional-elements-in-tuple-types), meaning its length can be 1 or 2.

The first element in the tuple is an array of `WalletTransaction` objects. The length of this array must be between 1 and 16 (inclusive on both ends). Each transaction in the group (even ones not being signed by the wallet) must be an element in this array.

The second element in the tuple is an `SignTxnOpts` object, which is optional.

The `WalletTransaction` and `SignTxnOpts` types are defined as:

```
interface WalletTransaction {
  /**
   * Base64 encoding of the canonical msgpack encoding of a     
   * Transaction.
   */
  txn: string;
 
  /**
   * Optional authorized address used to sign the transaction when 
   * the account is rekeyed. Also called the signor/sgnr.
   */
  authAddr?: AlgorandAddress;
 
  /**
   * Optional multisig metadata used to sign the transaction
   */
  msig?: MultisigMetadata;
 
  /**
   * Optional list of addresses that must sign the transactions
   */
  signers?: AlgorandAddress[];
 
  /**
   * Optional message explaining the reason of the transaction
   */
  message?: string;
}
 
interface SignTxnOpts {
  /**
   * Optional message explaining the reason of the group of 
   * transactions.
   */
  message?: string;
  
  // other options may be present, but are not standard
}
```

The above interfaces reference AlgorandAddress and MultisigMetadata types. These are defined as:

```
/**
 * AlgorandAddress is a 58-character base32 string that represents an
 * Algorand address with a checksum.
 */
type AlgorandAddress = string;
 
/**
 * Options for creating and using a multisignature account.
 */
interface MultisigMetadata {
  /**
   * Multisig version.
   */
  version: number;
 
  /**
   * Multisig threshold value. Authorization requires a subset of 
   * signatures, equal to or greater than the threshold value.
   */
  threshold: number;
 
  /**
   * List of Algorand addresses of possible signers for this
   * multisig. Order is important.
   */
  addrs: AlgorandAddress[];
}
```

A description for each parameter in `WalletTransaction` can be found here: https://github.com/algorandfoundation/ARCs/blob/main/ARCs/arc-0001.md#interface-wallettransaction

### Response

If the wallet approves the request, it will send back the following response:

```
interface AlgoSignTxnResponse {
  id: number;
  jsonrpc: "2.0";
  result: Array<string | null>;
}
```

In this response, result is an array with the same length as the request params. For every integer `i` such that `0 <= i < result.length`:

- If the transaction at index `i` in the group should be signed by the wallet (i.e. `params[0][i].signers` is not an empty array) : `result[i]` will be a base64-encoded string containing the msgpack-encoded signed transaction `params[i].txn`.
- Otherwise: `result[i]` will be null, since the wallet was not requested to sign this transaction.

If the wallet does not approve signing every transaction whose signature is being requested, the request must fail. A failure like this should be indicated in the rejection message as described [here](https://github.com/algorandfoundation/ARCs/blob/main/ARCs/arc-0001.md#error-standards).

# Future Additions

Possible future additions to the schema may include:

- A request type for wallets to sign a LogicSig program, resulting in a delegated LogicSig spending program for an account.