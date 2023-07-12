# Install on Mac

> **Note**
> This method will install Python 3.10 as a dependency via Homebrew. If you already have python installed, you may prefer to use `pipx install algokit` as explained [here](#install-algokit-with-pipx-on-any-os).

1. Ensure prerequisites are installed

   - [Homebrew](https://docs.brew.sh/Installation)
   - [Git](https://github.com/git-guides/install-git#install-git-on-mac) (should already be available if `brew` is installed)
   - [Docker](https://docs.docker.com/desktop/install/mac-install/), (or `brew install --cask docker`)
     > **Note**
     > Docker requires MacOS 11+

2. Install using Homebrew `brew install algorandfoundation/tap/algokit`
3. Restart the terminal to ensure AlgoKit is available on the path
4. [Verify installation](#verify-installation)

### Maintenance

Some useful commands for updating or removing AlgoKit in the future.

- To update AlgoKit: `brew upgrade algokit`
- To remove AlgoKit: `brew uninstall algokit`

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