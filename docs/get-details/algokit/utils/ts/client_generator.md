# AlgoKit TypeScript client generator (algokit-client-generator-ts)

This project generates a type-safe smart contract client in TypeScript for the Algorand Blockchain that wraps the [application client](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/docs/capabilities/app-client.md) in [AlgoKit Utils](https://github.com/algorandfoundation/algokit-utils-ts) and tailors it to a specific smart contract. It does this by reading an [ARC-0032](https://github.com/algorandfoundation/ARCs/blob/main/ARCs/arc-0032.md) application spec file and generating a client which exposes methods for each ABI method in the target smart contract, along with helpers to create, update, and delete the application.

## Usage

### Prerequisites

To be able to consume the generated file you need to include it in a TypeScript project that has (at least) the following package installed:

```
npm install @algorandfoundation/algokit-utils
```

### Use

The cli can be used to generate a client via the following command.

```
npx --yes  @algorandfoundation/algokit-client-generator generate -a ./application.json -o ./client.generated.ts
```

For details on how to use the generated client see the more detailed [usage docs](../)

## Examples

There are a range of [examples](https://github.com/algorandfoundation/algokit-client-generator-ts/tree/main/examples) that you can look at to see a source smart contract (`{contract.py}`), the generated client (`client.ts`) and some tests that demonstrate how you can use the client (`client.spec.ts`).

## Contributing

If you want to contribute to this project the following information will be helpful.

### Initial setup

1. Clone this repository locally
2. Install pre-requisites:

   - Install `AlgoKit` - [Link](https://github.com/algorandfoundation/algokit-cli#install): Ensure you can execute `algokit --version`.
   - Bootstrap your local environment; run `algokit bootstrap all` within this folder, which will:
     - Install `Poetry` - [Link](https://python-poetry.org/docs/#installation): The minimum required version is `1.2`. Ensure you can execute `poetry -V` and get `1.2`+
     - Run `poetry install` in the root directory, which will set up a `.venv` folder with a Python virtual environment and also install all Python dependencies
     - Run `npm install`

3. Open the project and start debugging / developing via:
   - VS Code
     1. Open the repository root in VS Code
     2. Install recommended extensions
     3. Run tests via test explorer
   - IDEA (e.g. PyCharm)
     1. Open the repository root in the IDE
     2. It should automatically detect it's a Poetry project and set up a Python interpreter and virtual environment.
     3. Run tests
   - Other
     1. Open the repository root in your text editor of choice
     2. Run `npm run test`

### Subsequently

1. If you update to the latest source code and there are new dependencies you will need to run `algokit bootstrap all` again
2. Follow step 3 above

### Building examples

In the `examples` folder there is a series of example contracts along with their generated client. These contracts are built using [Beaker](https://beaker.algo.xyz/).

If you want to make changes to any of the smart contract examples and re-generate the ARC-0032 application.json files then change the corresponding `examples/{contract}/{contract}.py` file and then run:

```
poetry run python -m examples
```

Or in Visual Studio Code you can use the default build task (Ctrl+Shift+B).

### Continuous Integration / Continuous Deployment (CI/CD)

This project uses [GitHub Actions](https://docs.github.com/en/actions/learn-github-actions/understanding-github-actions) to define CI/CD workflows, which are located in the [`.github/workflows`](./.github/workflows) folder.

### Approval tests

Making any changes to the generated code will result in the approval tests failing. The approval tests work by generating a version of client
and outputting it to `./examples/APP_NAME/client.generated.ts` then comparing to the approved version `./examples/APP_NAME/client.ts`. If you
make a change and break the approval tests, you will need to update the approved version by overwriting it with the generated version.
You can run `npm run update-approvals` to update all approved clients in one go.
