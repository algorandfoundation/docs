title: Test


## two spaces

=== "Windows"
  Windows

=== "macOS"
  macOS

=== "Linux"
  Linux

=== "OS agnostic"
  OS agnostic

## three spaces

=== "Windows"
   Windows

=== "macOS"
   macOS

=== "Linux"
   Linux

=== "OS agnostic"
   OS agnostic

## four spaces

=== "Windows"
    Windows

=== "macOS"
    macOS

=== "Linux"
    Linux

=== "OS agnostic"
    OS agnostic

## one tab

=== "Windows"
	Windows

=== "macOS"
	macOS

=== "Linux"
	Linux

=== "OS agnostic"
	OS agnostic

## Install AlgoKit A

=== "Windows"
   > **Note**
   > This method will install the most recent python3 version [via winget](https://learn.microsoft.com/en-us/windows/package-manager/winget/). If you already have python 3.10+ installed, you may you may prefer to use `pipx install algokit` as explained within the pipx on any OS section so you can control the python version used.

   1. Ensure prerequisites are installed
      - [Git](https://github.com/git-guides/install-git#install-git-on-windows) (or `winget install git.git`)
      - [Docker](https://docs.docker.com/desktop/install/windows-install/) (or `winget install docker.dockerdesktop`)
        > **Note**
        > See [our LocalNet documentation](https://github.com/algorandfoundation/algokit-cli/blob/main/docs/features/localnet.md#prerequisites) for more tips on installing Docker on Windows
   2. Install Python3 using WinGet
      1. Install python: `winget install python.python.3.11`
      2. Restart the terminal to ensure Python and pip are available on the path
         > **Note**
         > Windows has a feature called **App Execution Aliases** that provides redirects for the Python command that guide users to the
         > Windows Store. Unfortunately these aliases can prevent normal execution of Python if Python is installed via other means, to disable them
         > search for **Manage app execution aliases** from the start menu, and then turn off entries listed as
         > **App Installer python.exe** or **App Installer python3.exe**.
      3. Install pipx:
         ```
         pip install --user pipx
         python -m pipx ensurepath
         ```
      4. Restart the terminal to ensure pipx is available on the path
      5. Install AlgoKit via pipx: `pipx install algokit`
      6. Restart the terminal to ensure AlgoKit is available on the path
   <iframe width="100%" style="aspect-ratio:16/9" src="https://www.youtube-nocookie.com/embed/22RvINnZsRo" title="Installing AlgoKit on Windows" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>

=== "macOS"
   > **Note**
   > This method will install Python 3.10 as a dependency via Homebrew. If you already have python installed, you may prefer to use `pipx install algokit` as explained within the pipx on any OS section so you can control the python version used.
   
   1. Ensure prerequisites are installed

      - [Homebrew](https://docs.brew.sh/Installation)
      - [Git](https://github.com/git-guides/install-git#install-git-on-mac) (should already be available if `brew` is installed)
      - [Docker](https://docs.docker.com/desktop/install/mac-install/), (or `brew install --cask docker`)
        > **Note**
        > Docker requires MacOS 11+
   2. Install using Homebrew `brew install algorandfoundation/tap/algokit`
   3. Restart the terminal to ensure AlgoKit is available on the path
   4. [Verify installation](#verify-installation)
   <iframe width="100%" style="aspect-ratio:16/9" src="https://www.youtube-nocookie.com/embed/zsurtpCGmgE" title="Installing AlgoKit on macOS" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>

=== "Linux"
   1. Ensure prerequisites are installed
      - [Python 3.10+](https://www.python.org/downloads/)
        > **Note**
        > There is probably a better way to install Python than to download it directly, e.g. your local Linux package manager

      - [pipx](https://pypa.github.io/pipx/#on-linux-install-via-pip-requires-pip-190-or-later)
      - [Git](https://github.com/git-guides/install-git#install-git-on-linux)
      - [Docker](https://docs.docker.com/desktop/install/linux-install/)

   2. Continue with step 2 in the following section to install via `pipx` on any OS 

=== "OS agnostic"

   To install AlgoKit, run the following command from a terminal.

   ```shell
   pipx install algokit
   ```

   After the installation completes, **restart the terminal**.

Additional AlgoKit videos are available on the [@AlgoDevs YouTube channel](https://youtube.com/@AlgoDevs).

## Install AlgoKit B

=== "Windows"
   1. Ensure prerequisites are installed
      - [Git](https://github.com/git-guides/install-git#install-git-on-windows) (or `winget install git.git`)
      - [Docker](https://docs.docker.com/desktop/install/windows-install/) (or `winget install docker.dockerdesktop`)
        > **Note**
        > See [our LocalNet documentation](https://github.com/algorandfoundation/algokit-cli/blob/main/docs/features/localnet.md#prerequisites) for more tips on installing Docker on Windows
   2. Install Python3 using WinGet
      1. Install python: `winget install python.python.3.11`
      2. Restart the terminal to ensure Python and pip are available on the path
         > **Note**
         > Windows has a feature called **App Execution Aliases** that provides redirects for the Python command that guide users to the
         > Windows Store. Unfortunately these aliases can prevent normal execution of Python if Python is installed via other means, to disable them
         > search for **Manage app execution aliases** from the start menu, and then turn off entries listed as
         > **App Installer python.exe** or **App Installer python3.exe**.
      3. Install pipx:
         ```
         pip install --user pipx
         python -m pipx ensurepath
         ```
      4. Restart the terminal to ensure pipx is available on the path
      5. Install AlgoKit via pipx: `pipx install algokit`
      6. Restart the terminal to ensure AlgoKit is available on the path
   <iframe width="100%" style="aspect-ratio:16/9" src="https://www.youtube-nocookie.com/embed/22RvINnZsRo" title="Installing AlgoKit on Windows" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>

=== "macOS"
   1. Ensure prerequisites are installed

      - [Homebrew](https://docs.brew.sh/Installation)
      - [Git](https://github.com/git-guides/install-git#install-git-on-mac) (should already be available if `brew` is installed)
      - [Docker](https://docs.docker.com/desktop/install/mac-install/), (or `brew install --cask docker`)
        > **Note**
        > Docker requires MacOS 11+
   2. Install using Homebrew `brew install algorandfoundation/tap/algokit`
   3. Restart the terminal to ensure AlgoKit is available on the path
   4. [Verify installation](#verify-installation)
   <iframe width="100%" style="aspect-ratio:16/9" src="https://www.youtube-nocookie.com/embed/zsurtpCGmgE" title="Installing AlgoKit on macOS" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>

=== "Linux"
   1. Ensure prerequisites are installed
      - [Python 3.10+](https://www.python.org/downloads/)
        > **Note**
        > There is probably a better way to install Python than to download it directly, e.g. your local Linux package manager

      - [pipx](https://pypa.github.io/pipx/#on-linux-install-via-pip-requires-pip-190-or-later)
      - [Git](https://github.com/git-guides/install-git#install-git-on-linux)
      - [Docker](https://docs.docker.com/desktop/install/linux-install/)

   2. Continue with step 2 in the following section to install via `pipx` on any OS 

=== "OS agnostic"
   To install AlgoKit, run the following command from a terminal.

   ```shell
   pipx install algokit
   ```

   After the installation completes, **restart the terminal**.

Additional AlgoKit videos are available on the [@AlgoDevs YouTube channel](https://youtube.com/@AlgoDevs).

## Install AlgoKit C

=== "Windows"
   <iframe width="100%" style="aspect-ratio:16/9" src="https://www.youtube-nocookie.com/embed/22RvINnZsRo" title="Installing AlgoKit on Windows" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>

=== "macOS"
   <iframe width="100%" style="aspect-ratio:16/9" src="https://www.youtube-nocookie.com/embed/zsurtpCGmgE" title="Installing AlgoKit on macOS" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>

=== "Linux"
   1. Ensure prerequisites are installed
      - [Python 3.10+](https://www.python.org/downloads/)
        > **Note**
        > There is probably a better way to install Python than to download it directly, e.g. your local Linux package manager

      - [pipx](https://pypa.github.io/pipx/#on-linux-install-via-pip-requires-pip-190-or-later)
      - [Git](https://github.com/git-guides/install-git#install-git-on-linux)
      - [Docker](https://docs.docker.com/desktop/install/linux-install/)

   2. Continue with step 2 in the following section to install via `pipx` on any OS 

=== "OS agnostic"
   To install AlgoKit, run the following command from a terminal.

   ```shell
   pipx install algokit
   ```

   After the installation completes, **restart the terminal**.

Additional AlgoKit videos are available on the [@AlgoDevs YouTube channel](https://youtube.com/@AlgoDevs).

## Next steps

- To learn more about AlgoKit and what you can do with it, checkout the [AlgoKit documentation](../index.md).
- To learn more about Beaker, take a look at the [Beaker documentation](https://beaker.algo.xyz/).
- More information on Algorand smart contracts is also available in the [smart contract documentation](https://developer.algorand.org/docs/get-details/dapps/smart-contracts/).
