# Debugger

The AlgoKit Python Utilities package provides a set of debugging tools that can be used to simulate and trace transactions on the Algorand blockchain. These tools and methods are optimized for developers who are building applications on Algorand and need to test and debug their smart contracts via [AVM Debugger extension](link to vscode extension).

## Configuration

The `config.py` file contains the `UpdatableConfig` class which manages and updates configuration settings for the AlgoKit project. The class has the following attributes:

-   `debug`: Indicates whether debug mode is enabled.
-   `project_root`: The path to the project root directory. Can be ignored if you are using `algokit_utils` inside an `algokit` compliant project (containing `.algokit.toml` file). For non algokit compliant projects, simply provide the path to the folder where you want to store sourcemaps and traces to be used with [`AVM Debugger`](links to extension). Alternatively you can also set the value via the `ALGOKIT_PROJECT_ROOT` environment variable.
-   `trace_all`: Indicates whether to trace all operations. Defaults to false, this means that when debug mode is enabled, any (or all) application client calls performed via `algokit_utils` will store responses from `simulate` endpoint. These files are called traces, and can be used with [AVM Debugger](link to vscode extension) to debug TEAL source codes, transactions in the atomic group and etc.
-   `trace_buffer_size_mb`: The size of the trace buffer in megabytes. By default uses 256 megabytes. When output folder containing debug trace files exceedes the size, oldest files are removed to optimize for storage consumption.
-   `max_search_depth`: The maximum depth to search for a an `algokit` config file. By default it will traverse at most 10 folders searching for `.algokit.toml` file which will be used to assume algokit compliant project root path.

The `configure` method can be used to set these attributes.

To enable debug mode in your project you can configure it as follows:

```py
from algokit_utils.config import config

config.configure(debug=True)
```

## Debugging Utilities

Debugging utilities can be used to simplify gathering artifacts to be used with [AlgoKit AVM Debugger](https://github.com/algorandfoundation/algokit-avm-vscode-debugger) in non algokit compliant projects. The following methods are provided:

-   `persist_sourcemaps`: This method persists the sourcemaps for the given sources as AVM Debugger compliant artifacts. It takes a list of `PersistSourceMapInput` objects, a `Path` object representing the root directory of the project, an `AlgodClient` object for interacting with the Algorand blockchain, and a boolean indicating whether to dump teal source files along with sourcemaps.
-   `simulate_and_persist_response`: This method simulates the atomic transactions using the provided `AtomicTransactionComposer` object and `AlgodClient` object, and persists the simulation response to an AVM Debugger compliant JSON file. It takes an `AtomicTransactionComposer` object representing the atomic transactions to be simulated and persisted, a `Path` object representing the root directory of the project, an `AlgodClient` object representing the Algorand client, and a float representing the size of the trace buffer in megabytes.

### Trace filename format

The trace files are named in a specific format to provide useful information about the transactions they contain. The format is as follows:

```ts
`${timestamp}_lr${last_round}_${transaction_types}.trace.avm.json`;
```

Where:

-   `timestamp`: The time when the trace file was created, in ISO 8601 format, with colons and periods removed.
-   `last_round`: The last round when the simulation was performed.
-   `transaction_types`: A string representing the types and counts of transactions in the atomic group. Each transaction type is represented as `${count}#${type}`, and different transaction types are separated by underscores.

For example, a trace file might be named `20220301T123456Z_lr1000_2#pay_1#axfer.trace.avm.json`, indicating that the trace file was created at `2022-03-01T12:34:56Z`, the last round was `1000`, and the atomic group contained 2 payment transactions and 1 asset transfer transaction.
