# CLI Reference



AlgoKit is your one-stop shop to develop applications on the Algorand blockchain.

If you are getting started, please see the [quick start tutorial](tutorials/intro.md).

```shell
algokit [OPTIONS] COMMAND [ARGS]...
```

### Options


### --version
Show the version and exit.


### -v, --verbose
Enable logging of DEBUG messages to the console.


### --color, --no-color
Force enable or disable of console output styling.


### --skip-version-check
Skip version checking and prompting.

## bootstrap

Expedited initial setup for any developer by installing and configuring dependencies and other
key development environment setup activities.

```shell
algokit bootstrap [OPTIONS] COMMAND [ARGS]...
```

### Options


### --force
Continue even if minimum AlgoKit version is not met

### all

Runs all bootstrap sub-commands in the current directory and immediate sub directories.

```shell
algokit bootstrap all [OPTIONS]
```

### Options


### --interactive, --non-interactive, --ci
Enable/disable interactive prompts. If the CI environment variable is set, defaults to non-interactive

### env

Copies .env.template file to .env in the current working directory and prompts for any unspecified values.

```shell
algokit bootstrap env [OPTIONS]
```

### Options


### --interactive, --non-interactive, --ci
Enable/disable interactive prompts. If the CI environment variable is set, defaults to non-interactive

### npm

Runs npm install in the current working directory to install Node.js dependencies.

```shell
algokit bootstrap npm [OPTIONS]
```

### poetry

Installs Python Poetry (if not present) and runs poetry install in the current working directory to install Python dependencies.

```shell
algokit bootstrap poetry [OPTIONS]
```

## completions

Install and Uninstall AlgoKit shell integrations.

```shell
algokit completions [OPTIONS] COMMAND [ARGS]...
```

### install

Install shell completions, this command will attempt to update the interactive profile script
for the current shell to support algokit completions. To specify a specific shell use –shell.

```shell
algokit completions install [OPTIONS]
```

### Options


### --shell <shell>
Specify shell to install algokit completions for.


* **Options**

    bash | zsh


### uninstall

Uninstall shell completions, this command will attempt to update the interactive profile script
for the current shell to remove any algokit completions that have been added.
To specify a specific shell use –shell.

```shell
algokit completions uninstall [OPTIONS]
```

### Options


### --shell <shell>
Specify shell to install algokit completions for.


* **Options**

    bash | zsh


## config

Configure settings used by AlgoKit

```shell
algokit config [OPTIONS] COMMAND [ARGS]...
```

### version-prompt

Controls whether AlgoKit checks and prompts for new versions.
Set to [disable] to prevent AlgoKit performing this check permanently, or [enable] to resume checking.
If no argument is provided then outputs current setting.

Also see –skip-version-check which can be used to disable check for a single command.

```shell
algokit config version-prompt [OPTIONS] [[enable|disable]]
```

### Arguments


### ENABLE
Optional argument

## deploy

Deploy smart contracts from AlgoKit compliant repository.

```shell
algokit deploy [OPTIONS] [ENVIRONMENT_NAME]
```

### Options


### -C, --command <command>
Custom deploy command. If not provided, will load the deploy command from .algokit.toml file.


### --interactive, --non-interactive, --ci
Enable/disable interactive prompts. If the CI environment variable is set, defaults to non-interactive


### -P, --path <path>
Specify the project directory. If not provided, current working directory will be used.

### Arguments


### ENVIRONMENT_NAME
Optional argument

## doctor

Diagnose potential environment issues that may affect AlgoKit

Will search the system for AlgoKit dependencies and show their versions, as well as identifying any
potential issues.

```shell
algokit doctor [OPTIONS]
```

### Options


### -c, --copy-to-clipboard
Copy the contents of the doctor message (in Markdown format) in your clipboard.

## explore

Explore the specified network in the browser using Dappflow.

```shell
algokit explore [OPTIONS] [[localnet|testnet|mainnet]]
```

### Arguments


### NETWORK
Optional argument

## generate

Generate code for an Algorand project.

```shell
algokit generate [OPTIONS] COMMAND [ARGS]...
```

### client

Create a typed ApplicationClient from an ARC-32 application.json

Supply the path to an application specification file or a directory to recursively search
for “application.json” files

```shell
algokit generate client [OPTIONS] APP_SPEC_PATH_OR_DIR
```

### Options


### -o, --output <output_path_pattern>
Path to the output file. The following tokens can be used to substitute into the output path: {contract_name}, {app_spec_dir}


### -l, --language <language>
Programming language of the generated client code


* **Options**

    python | typescript


### Arguments


### APP_SPEC_PATH_OR_DIR
Required argument

## goal

Run the Algorand goal CLI against the AlgoKit LocalNet.

Look at [https://developer.algorand.org/docs/clis/goal/goal/](https://developer.algorand.org/docs/clis/goal/goal/) for more information.

```shell
algokit goal [OPTIONS] [GOAL_ARGS]...
```

### Options


### --console
Open a Bash console so you can execute multiple goal commands and/or interact with a filesystem.

### Arguments


### GOAL_ARGS
Optional argument(s)

## init

Initializes a new project from a template, including prompting
for template specific questions to be used in template rendering.

Templates can be default templates shipped with AlgoKit, or custom
templates in public Git repositories.

Includes ability to initialise Git repository, run algokit bootstrap and
automatically open Visual Studio Code.

This should be run in the parent directory that you want the project folder
created in.

```shell
algokit init [OPTIONS]
```

### Options


### -n, --name <directory_name>
Name of the project / directory / repository to create.


### -t, --template <template_name>
Name of an official template to use. To see a list of descriptions, run this command with no arguments.


* **Options**

    beaker | react | fullstack | playground



### --template-url <URL>
URL to a git repo with a custom project template.


### --template-url-ref <URL>
Specific tag, branch or commit to use on git repo specified with –template-url. Defaults to latest.


### --UNSAFE-SECURITY-accept-template-url
Accept the specified template URL, acknowledging the security implications of arbitrary code execution trusting an unofficial template.


### --git, --no-git
Initialise git repository in directory after creation.


### --defaults
Automatically choose default answers without asking when creating this template.


### --bootstrap, --no-bootstrap
Whether to run algokit bootstrap to install and configure the new project’s dependencies locally.


### --ide, --no-ide
Whether to open an IDE for you if the IDE and IDE config are detected. Supported IDEs: VS Code.


### -a, --answer <key> <value>
Answers key/value pairs to pass to the template.

## localnet

Manage the AlgoKit LocalNet.

```shell
algokit localnet [OPTIONS] COMMAND [ARGS]...
```

### console

Run the Algorand goal CLI against the AlgoKit LocalNet via a Bash console so you can execute multiple goal commands and/or interact with a filesystem.

```shell
algokit localnet console [OPTIONS]
```

### explore

Explore the AlgoKit LocalNet using Dappflow

```shell
algokit localnet explore [OPTIONS]
```

### logs

See the output of the Docker containers

```shell
algokit localnet logs [OPTIONS]
```

### Options


### --follow, -f
Follow log output.


### --tail <tail>
Number of lines to show from the end of the logs for each container.


* **Default**

    `all`


### reset

Reset the AlgoKit LocalNet.

```shell
algokit localnet reset [OPTIONS]
```

### Options


### --update, --no-update
Enable or disable updating to the latest available LocalNet version, default: don’t update

### start

Start the AlgoKit LocalNet.

```shell
algokit localnet start [OPTIONS]
```

### status

Check the status of the AlgoKit LocalNet.

```shell
algokit localnet status [OPTIONS]
```

### stop

Stop the AlgoKit LocalNet.

```shell
algokit localnet stop [OPTIONS]
```
