# Install on Linux

1. Ensure prerequisites are installed

   - [Python 3.10+](https://www.python.org/downloads/)

     > **Note**
     > There is probably a better way to install Python than to download it directly, e.g. your local Linux package manager

   - [pipx](https://pypa.github.io/pipx/#on-linux-install-via-pip-requires-pip-190-or-later)
   - [Git](https://github.com/git-guides/install-git#install-git-on-linux)
   - [Docker](https://docs.docker.com/desktop/install/linux-install/)

2. Continue with step 2 in the following section to install via [pipx](#install-algokit-with-pipx-on-any-os)

## Verify installation

Verify AlgoKit is installed correctly by running `algokit --version` and you should see output similar to:

```
algokit, version 1.0.1
```

> **Note**
> If you get receive one of the following errors:
>
> - `command not found: algokit` (bash/zsh)
> - `The term 'algokit' is not recognized as the name of a cmdlet, function, script file, or operable program.` (PowerShell)
>
> Then ensure that `algokit` is available on the PATH by running `pipx ensurepath` and restarting the terminal.

It is also recommended that you run `algokit doctor` to verify there are no issues in your local environment and to diagnose any problems if you do have difficulties running AlgoKit. The output of this command will look similar to:

```
timestamp: 2023-03-27T01:23:45+00:00
AlgoKit: 1.0.1
AlgoKit Python: 3.11.1 (main, Dec 23 2022, 09:28:24) [Clang 14.0.0 (clang-1400.0.29.202)] (location: /Users/algokit/.local/pipx/venvs/algokit)
OS: macOS-13.1-arm64-arm-64bit
docker: 20.10.21
docker compose: 2.13.0
git: 2.37.1
python: 3.10.9 (location:  /opt/homebrew/bin/python)
python3: 3.10.9 (location:  /opt/homebrew/bin/python3)
pipx: 1.1.0
poetry: 1.3.2
node: 18.12.1
npm: 8.19.2
brew: 3.6.18

If you are experiencing a problem with AlgoKit, feel free to submit an issue via:
https://github.com/algorandfoundation/algokit-cli/issues/new
Please include this output, if you want to populate this message in your clipboard, run `algokit doctor -c`
```

Per the above output, the doctor command output is a helpful tool if you need to ask for support or [raise an issue](https://github.com/algorandfoundation/algokit-cli/issues/new).