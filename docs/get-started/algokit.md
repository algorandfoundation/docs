title: AlgoKit quick start

AlgoKit is the primary tool used by the Algorand community to develop smart contracts on the Algorand blockchain. It provides the capabilities to develop, test and deploy Algorand smart contracts within minutes! This guide is intended to help you setup AlgoKit and to start developing your application.

# Quick start videos
If you prefer videos, take a look at this 10 minute guide to getting started.
<center>
[![Learn How to Build on Algorand in 10 Minutes](https://i.ytimg.com/vi/dow6U8DxOGc/hqdefault.jpg)](https://www.youtube.com/embed/dow6U8DxOGc)
</center>

Detailed video guides for both [Windows](https://www.youtube.com/embed/22RvINnZsRo) and [Mac](https://www.youtube.com/embed/zsurtpCGmgE) are also available.


# Prequisites
This guide presents installing AlgoKit using an OS agnostic proceeddure. For OS specific instructions take a look that the [AlgoKit install](https://github.com/algorandfoundation/algokit-cli/blob/main/README.md#install) guide.

Using this proceedure requires the the following components be installed already

* [Python 3.10](https://www.python.org/downloads/) or higher
* [PipX](https://pypa.github.io/pipx/#on-linux-install-via-pip-requires-pip-190-or-later)
* [Git](https://github.com/git-guides/install-git#install-git)
* [Docker](https://docs.docker.com/desktop/install/mac-install/)
* [VSCode](https://code.visualstudio.com/download)


# Install AlogKit

To install AlgoKit, run the following command from a terminal.

```shell
pipx install algokit
```

After the installation completes, <b>restart the terminal</b>.

# Verify the Installation

To verify AlgoKit Installed correctly run the following.

```shell
algokit --version
```

Output similar to the following shoud be displayed

`algokit, version 0.5.0`

Should be displayed.

# Start a LocalNet
AlgoKit supports using a local version of the Algorand blockchain. To start an instance of this LocalNet run the following command from the terminal.

```shell
algokit loclanet start
```

This should start an instance of the LocalNet within docker. If you open the Docker Desktop application you should something simliar to the following.

<center>
![Docker Desktop LocalNet Instance](../../imgs/localnet.png){: style="width:700px" align=center }
</center>

# Create an AlgoKit project
Now that AlgoKit is installed, a new project can be created. This can be done by running:

```shell
algokit init
```

This will lauch a guided menu system to create a specific project. After first naming your project, you will be prompted to select a specific template. Templates are basic starter applications for Algorand. To read more about templates checkout AlgoKit detailed documentation. For now, use the down arrow to selet the playground template. Select defaults for the rest of the prompts.

Once finished, VSCode should automatically be opened with a starter project named `playground` and you will be prompted to install appropriate VSCode plugins. This starter app will contain one smart contract (built using [Beaker](../get-details/dapps/writing-contracts/beaker.md)) named `helloworld.py` with one method (`hello`) that takes a `String` and returns a `String`.

<center>
![AlgoKit Playground Contract](../../imgs/algokitplayground.png){: style="width:700px" align=center }
</center>

# Run the Demo Application

Once the playground application is created, you will notice in the `helloworld` folder a file named `demo.py` which is a simple example of using AlgoKit to deploy and make a call to the `helloworld` smart contract that is running on the LocalNet instance started earlier.

<center>
![AlgoKit Playground Demo](../../imgs/algokitdemo.png){: style="width:700px" align=center }
</center>

Right clicking on this file and selecting `Run Python File in Terminal` will deploy the `HellowWorldApp` smart contract and then call it passing the parameter `name` with a value of `Beaker`.

This should produce something similiar to the following in the VSCode terminal.

```shell
(playground-py3.11) User@Algo-User-MBP myalgokit project % "/Users/user/code/algokit/myalgokit pro
ject/.venv/bin/python" "/Users/user/code/algokit/myalgokit project/playground/hello_world/demo.py"
Deployed app in txid SBNIJYZUOXVPXDFJHAVDDCO5TQ4WZ7P37QOOQM6CSVSMIURETHNQ
        App ID: 11 
        Address: 377KDIVHB7K2LFFJIPUPQFWJGVD36MAP4EI6Y2HAGRAW6JPK4MDUBC3YVY 
    
Hello, Beaker
```

The App ID of of the deployed contract and its Algorand address is displayed, followed by the message `Hello, Beaker`.

At this point you have deployed a simple contract to an Algorand network and called it successfully!

Additionally, you can find the native TEAL smart contract code and the appropriate smart contract manifest JSON files located in the `artifacts` folder.

<center>
![AlgoKit Playground Demo](../../imgs/algokitartifacts.png){: style="width:700px" align=center }
</center>

Running the `build.py` python file will also generate these artifacts. These files can be used by tools like Dappflow to deploy your smart contract to the various Algorand networks. 

# Using Dappflow 

Dappflow is a web application that can be used to easily deploy and test a contract. Dappflow can be launched by running the following command from the VSCode terminal.

```shell
algokit explore
```

This command will launch the browser and load the Dappflow web application. It should automatically connect to your LocalNet running network. It will be displayed as `sandnet-v1` in the upper left hand corner.

<center>
![AlgoKit Dappflow](../../imgs/dappflow1.png){: style="width:700px" align=center }
</center>

## Create test account

Select `Dev Wallets` from the left menu and click on the `Create wallet` button. This will create an account on the LocalNet and fund it with 100 Algos that can be used to test with.

<center>
![AlgoKit Dappflow](../../imgs/dappflow2.png){: style="width:700px" align=center }
</center>

In the bottom left hand corner of the menu, select `Connect wallet` and you will be prompted with several wallet choices. Choose the `Dev Wallet` option. This will connect the wallet you just created to Dappflow.

## Deploy the Hello World application

To deploy the application created, select the `Beaker studio` menu and click on the import beaker app. Seletct `File` and `Upload file`, browse to the artifacts created in the previous section of this guide. Select the `application.json` manifest file. This will load the specific manifest file for the Hello World sample application.

<center>
![AlgoKit Dappflow](../../imgs/dappflow3.png){: style="width:700px" align=center }
</center>

To deploy this application again, select the `Creat app` button followed by the `Create` button from the popup. You should get a `Transaction successful` messation with the option to view the specific transaction in the explorer. Close out the popup and then scroll down to the `ABI` seciton of the page. The hello method should be displayed with an execute button beside it.

<center>
![AlgoKit Dappflow](../../imgs/dappflow4.png){: style="width:700px" align=center }
</center>

Click on the Execute button and a popup will be displayed allowing you to enter the parameter that we defined in the `HelloWorldApp` smart contract.

<center>
![AlgoKit Dappflow](../../imgs/dappflow5.png){: style="width:700px" align=center }
</center>

Enter a string in the parameter and click on `Execute`. You should get get a confirmation that the method executed properly and what the smart contract returned.

<center>
![AlgoKit Dappflow](../../imgs/dappflow6.png){: style="width:700px" align=center }
</center>

You have now successfully deployed and executed a smart contract method call using Dappflow!


# Next steps

- To learn more about AlgoKit and what you can do with it, checkout the [AlgoKit documentation](../get-details/algokit.md).
- To learn more about Beaker, take a look at the [Beaker documentation](../get-details/dapps/writing-contracts/beaker.md).
- To get detailed AlgoKit documenation, see the [AlgoKit repository](https://github.com/algorandfoundation/algokit-cli).
- More information on Algorand smart contracts is also availble in the [smart contract documentation](../get-details/dapps/smart-contracts/apps/index.md).
