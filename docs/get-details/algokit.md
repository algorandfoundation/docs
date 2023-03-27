title: AlgoKit

AlgoKit is a collection of tools and libraries that make it easy to build applications on Algorand.

# CLI Tool

The [AlgoKit CLI Tool](https://github.com/algorandfoundation/algokit-cli) provides a command line interface developers building applications on Algorand.

Using the CLI tool a developer can, among other things, spin up or manage a [local network](#localnet) or initialize a new project from a [Template](#templates).

## LocalNet

The AlgoKit `localnet` sub-command allows a developer to quickly spin up a private network in a local docker container. This is useful for testing and development. This command replaces the need to install a local node manually or with the Sandbox.

Simply running
```sh
$ algokit localnet start
```

Will download the required docker containers and spin up a network listening on the same ports as Sandbox and create and fund a set of accounts for the developer to use.

## Templates

The AlgoKit `init` sub-command allows a developer to instantly initialize a new project from a template. A default set of templates is provided by the AlgoKit CLI tool but any template can be used by passing the `--template-url` flag.  

Simply running 
```sh
$ algokit init 
```

Will take the developer through a set of prompts to initialize a new project. The developer can choose from a set of templates and the CLI tool will download the template and initialize a the project in a new directory.

# Utilities

There are a set of libraries that `AlgoKit` provides to make it easier to build applications on Algorand. These libraries are written with developer experience in mind and abstract the most common tasks away from the developer. 

 - [Python](https://github.com/algorandfoundation/algokit-utils-py)
 - [TypeScript/JS](https://github.com/algorandfoundation/algokit-utils-ts)