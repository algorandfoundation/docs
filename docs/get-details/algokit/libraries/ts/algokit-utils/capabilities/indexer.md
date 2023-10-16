# Indexer lookups / searching

Indexer lookups / searching is a higher-order use case capability provided by AlgoKit Utils that builds on top of the core capabilities. It provides type-safe indexer API wrappers (no more `Record<string, any>` pain), including automatic pagination control.

To see some usage examples check out the [automated tests](../../src/indexer-lookup.spec.ts).

## Indexer wrapper functions

There is a subset of [indexer API calls](https://developer.algorand.org/docs/rest-apis/indexer) that are exposed as easy to use methods with correct typing exposed and automatic pagination for multi item returns.

- [`algokit.lookupTransactionById(transactionId, indexer)`](../code/modules/index.md#lookuptransactionbyid) - Finds a transaction by ID
- [`algokit.lookupAccountByAddress(accountAddress, indexer)`](../code/modules/index.md#lookupaccountbyaddress) - Finds an account by address
- [`algokit.lookupAccountCreatedApplicationByAddress(indexer, address, getAll?, paginationLimit?)`](../code/modules/index.md#lookupaccountcreatedapplicationbyaddress) - Finds all applications created for an account
- [`algokit.searchTransactions(indexer, searchCriteria, paginationLimit?)`](../code/modules/index.md#searchtransactions) - Search for transactions with a given set of criteria
- [`algokit.executePaginatedRequest(extractItems, buildRequest)`](../code/modules/index.md#executepaginatedrequest) - Execute the given indexer request with automatic pagination

### Search transactions example

To use the `algokit.searchTransaction` method, you can follow this example as a starting point:

```typescript
const transactions = await algokit.searchTransactions(indexer, (s) => s.txType('pay').addressRole('sender').address(myAddress))
```

### Automatic pagination example

To use the `algokit.executePaginatedRequest` method, you can follow this example as a starting point:

```typescript
const transactions = await executePaginatedRequest(
  (response: TransactionSearchResults) => {
    return response.transactions
  },
  (nextToken) => {
    let s = indexer.searchForTransactions().txType('pay').address(myAddress).limit(1000)
    if (nextToken) {
      s = s.nextToken(nextToken)
    }
    return s
  },
)
```

It takes the first lambda to translate the raw response into the array that should keep getting appended as the pagination is followed and the second lambda constructs the request (without the `.do()` call), including populating the pagination token.

## Indexer API response types

The response model type definitions for the majority of [indexer API](https://developer.algorand.org/docs/rest-apis/indexer) are exposed from the `types/indexer` namespace in AlgoKit Utils. This is so that you can have a much better experience than the default response type of `Record<string, any>` from the indexer client in `algosdk`. If there is a type you want to use that is missing feel free to [submit a pull request](https://github.com/algorandfoundation/algokit-utils-ts/pulls) to [add the type(s)](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/indexer.ts).

To access these types you can import them:

```typescript
import { /* ... */ } '@algorandfoundation/algokit-utils/types/indexer'
```

As a general convention, the response types are named `{TypeName}Result` for a single item result and `{TypeName}Results` for a multiple item result where `{TypeName}` is:

- `{Entity}Lookup` for an API call response that returns a lookup for a single entity e.g. `AssetLookupResult`
- `{Entity}Search` for an API call response that searches for a type of entity e.g. `TransactionSearchResults`
- The `UpperCamelCase` name of a given model type as specified in the [official documentation](https://developer.algorand.org/docs/rest-apis/indexer) for any sub-types within a response e.g. `ApplicationResult`

The reason `Result/Results` is suffixed to the type is to avoid type name clashes for commonly used types from `algosdk` like `Transaction`.

To use these types with an indexer call you simply need to find the right result type and cast the response from `.do()` for the call in question, e.g.:

```typescript
import { TransactionLookupResult } from '@algorandfoundation/algokit-utils/types/indexer'

...

const transaction = (await indexer.lookupTransactionByID(transactionId).do()) as TransactionLookupResult
```
