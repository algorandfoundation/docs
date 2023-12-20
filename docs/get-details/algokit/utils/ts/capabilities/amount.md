# Algo amount handling

Algo amount handling is one of the core capabilities provided by AlgoKit Utils. It allows you to reliably and tersely specify amounts of microAlgos and Algos and safely convert between them.

Any AlgoKit Utils function that needs an Algo amount will take an `AlgoAmount` object, which ensures that there is never any confusion about what value is being passed around. Whenever an AlgoKit Utils function calls into an underlying algosdk function, or if you need to take an `AlgoAmount` and pass it into an underlying algosdk function (per the [modularity principle](../index.md#core-principles)) you can safely and explicitly convert to microAlgos or algos.

To see some usage examples check out the [automated tests](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/amount.spec.ts). Alternatively, you see the [reference documentation](../code/classes/types_amount.AlgoAmount.md) for `AlgoAmount`.

## `AlgoAmount`

The `AlgoAmount` class provides a safe wrapper around an underlying `number` amount of microAlgos where any value entering or existing the `AlgoAmount` class must be explicitly stated to be in microAlgos or Algos. This makes it much safer to handle Algo amounts rather than passing them around as raw `number`'s where it's easy to make a (potentially costly!) mistake and not perform a conversion when one is needed (or perform one when it shouldn't be!).

### Creating an `AlgoAmount`

There are a few ways to create an `AlgoAmount`:

- Algos
  - Constructor: `new AlgoAmount({algos: 10})`
  - Static helper: `AlgoAmount.algos(10)`
  - AlgoKit Helper: `algokit.algos(10)`
  - Number coersion: `(10).algos()` 
!!! Note
    You have to wrap the number in brackets or have it in a variable or function return, a raw number value can't have a method called on it)
- microAlgos
  - Constructor: `new AlgoAmount({microAlgos: 10_000})`
  - Static helper: `AlgoAmount.algos(10)`
  - AlgoKit Helper: `algokit.microAlgos(10_000)`
  - Number coersion: `(10_000).microAlgos()` 
!!! Note
    You have to wrap the number in brackets or have it in a variable or function return, a raw number value can't have a method called on it)

!!! Note
    To use any of the versions that reference `AlgoAmount` type itself you need to import it:
    ```typescript
    import { AlgoAmount } from '@algorandfoundation/algokit-utils/types/amount'
    ```

### Extracting a value from `AlgoAmount`

The `AlgoAmount` class has methods to return algos and microAlgos:

- `amount.algos()` - Returns the value in Algos
- `amount.microAlgos()` - Returns the value in microAlgos

`AlgoAmount` will coerce to a `number` automatically (in microAlgos), which is not recommended to be used outside of allowing you to use `AlgoAmount` objects in comparison operations such as `<` and `>=` etc.

You can also call `.toString()` or use an `AlgoAmount` directly in string interpolation to convert it to a nice user-facing formatted amount expressed in microAlgos.
