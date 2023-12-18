# Debugger

The AlgoKit TypeScript Utilities package provides a set of debugging tools that can be used to simulate and trace transactions on the Algorand blockchain. These tools and methods are optimized for developers who are building applications on Algorand and need to test and debug their smart contracts via [AlgoKit AVM Debugger extension](https://github.com/algorandfoundation/algokit-avm-vscode-debugger).

## Configuration

The `config.ts` file contains the `UpdatableConfig` class which manages and updates configuration settings for the AlgoKit project. The class has the following attributes:

- `debug`: Indicates whether debug mode is enabled.
- `projectRoot`: The path to the project root directory. Can be ignored if you are using `algokit-utils` inside an `algokit` compliant project (containing `.algokit.toml` file). For non algokit compliant projects, simply provide the path to the folder where you want to store sourcemaps and traces to be used with [`AlgoKit AVM Debugger`](https://github.com/algorandfoundation/algokit-avm-vscode-debugger). Alternatively you can also set the value via the `ALGOKIT_PROJECT_ROOT` environment variable.
- `traceAll`: Indicates whether to trace all operations. Defaults to false, this means that when debug mode is enabled, any (or all) application client calls performed via `algokit-utils` will store responses from `simulate` endpoint. These files are called traces, and can be used with [AlgoKit AVM Debugger](https://marketplace.visualstudio.com/items?itemName=algorandfoundation.algokit-avm-vscode-debugger) to debug TEAL source codes, and a large variety of transaction types in the atomic groups. Default behaviour will perform persistence of traces only on scenarios where any of the transaction groups emitted by an ApplicationClient or `sendAtomicTransactionComposer` method fails.
- `traceBufferSizeMb`: The size of the trace buffer in megabytes. By default uses 256 megabytes. When output folder containing debug trace files exceedes the size, oldest files are removed to optimize for storage consumption. This is useful when you are running a long running application and want to keep the trace files for debugging purposes but also be mindful of storage consumption.
- `maxSearchDepth`: The maximum depth to search for a an `algokit` config file. By default it will traverse at most `10` folders searching for `.algokit.toml` file which will be used to determine algokit compliant project root path.

The `configure` method can be used to set these attributes.

To enable debug mode in your project you can configure it as follows:

```ts
import { Config } from '@algorandfoundation/algokit-utils'
Config.configure({
  debug: true,
})
```

## Debugging Utilities

Debugging utilities can be used to simplify gathering artifacts to be used with [AlgoKit AVM Debugger](https://github.com/algorandfoundation/algokit-avm-vscode-debugger) in non algokit compliant projects. The following methods are provided:

- `persistSourceMaps`: This method persists the sourcemaps for the given sources as AlgoKit AVM Debugger compliant artifacts. It accepts an array of `PersistSourceMapInput` objects. Each object can either contain `rawTeal`, in which case the function will execute a compile to obtain byte code, or it can accept an object of type `CompiledTeal` provided by algokit, which is used for source codes that have already been compiled and contain the traces. It also accepts the root directory of the project, an `Algodv2` client to perform the compilation, and a boolean indicating whether to include the source files in the output.
- `simulateAndPersistResponse`: This method simulates the atomic transactions using the provided `AtomicTransactionComposer` object and `Algodv2` object, and persists the simulation response to an AlgoKit AVM Debugger compliant JSON file. It accepts the `AtomicTransactionComposer` with transaction(s) loaded, an `Algodv2` client to perform the simulation, the root directory of the project, and the buffer size in megabytes.

To enable debug mode with extra trace persistence for AVM VSCode Debugger, you can configure it as follows:

```ts
import { Config } from '@algorandfoundation/algokit-utils'
Config.configure({
  debug: true,
  traceAll: true, // if ignored, will only trace failed atomic transactions and application client calls
  projectRoot: '/path/to/project/root', // if ignored, will try to find the project root automatically
})
```

### Trace filename format

The trace files are named in a specific format to provide useful information about the transactions they contain. The format is as follows:

```ts
;`${timestamp}_lr${lastRound}_${transactionTypes}.trace.avm.json`
```

Where:

- `timestamp`: The time when the trace file was created, in ISO 8601 format, with colons and periods removed.
- `lastRound`: The last round when the simulation was performed.
- `transactionTypes`: A string representing the types and counts of transactions in the atomic group. Each transaction type is represented as `${count}#${type}`, and different transaction types are separated by underscores.

For example, a trace file might be named `20220301T123456Z_lr1000_2#pay_1#axfer.trace.avm.json`, indicating that the trace file was created at `2022-03-01T12:34:56Z`, the last round was `1000`, and the atomic group contained 2 payment transactions and 1 asset transfer transaction.
