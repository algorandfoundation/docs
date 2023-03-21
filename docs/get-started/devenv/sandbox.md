title: Sandbox Private Network

As explained in the ["Set up your development environment" page](index.md), for any more advanced development on Algorand, [sandbox]() is the recommended option.

## Installing sandbox

Once you have the prerequisites detailed on the [sandbox README](https://github.com/algorand/sandbox#getting-started), you can install the Algorand sandbox:

```sh
# clone the sandbox from GitHub
git clone https://github.com/algorand/sandbox.git

# enter the sandbox folder
cd sandbox

# run the sandbox executable to start a private network
./sandbox up
```

## Running a private network on sandbox

Starting the sandbox with the `up` command will take a couple of minutes if this is your first time running the sandbox. The script will make sure to pull all required Docker images before setting up your sandbox.

A successful node installation will print a list of prefunded accounts. Here's an example of the sandbox output.

```sh
# Available accounts
./sandbox goal account list

# [offline]	HCNMMIL3MKYILOLUO74NF6OPCJU4RU7IE5PX6JYBIT5YHAMQIVO5YADHMU	HCNMMIL3MKYILOLUO74NF6OPCJU4RU7IE5PX6JYBIT5YHAMQIVO5YADHMU	1000000000000000 microAlgos

# [offline]	3KHVQUNTXBFKPTWPPLRYZY3MZIW4EB6XYWRTTIA36O6ZSMRLSEWA2J2HTA	3KHVQUNTXBFKPTWPPLRYZY3MZIW4EB6XYWRTTIA36O6ZSMRLSEWA2J2HTA	4000000000000000 microAlgos

# [online]	5FRKKWRG3UAJQNB7QIOWBW2JICZS4YUF2WFAETHGN3CBM4R3N27NY3T2KQ	5FRKKWRG3UAJQNB7QIOWBW2JICZS4YUF2WFAETHGN3CBM4R3N27NY3T2KQ	4000000000000000 microAlgos
```

## Connect to sandbox using an SDK

To connect to the sandbox using an SDK, you can use the below connection object.

```js
const algosdk = require('algosdk');

// create client object to connect to sandbox's algod client
const algodToken = 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa';
const algodServer = 'http://localhost';
const algodPort = 4001;
const algodClient = new algosdk.Algodv2(algodToken, algodServer, algodPort);
```

## Sandbox Block Explorer (Dappflow)

Now you have a private network with transactions being created, you're going to want to be able to view the transactions and data on it. This is where [Dappflow](https://dappflow.org) comes in. It's a Block Explorer that can not only access Mainnet and Testnet using public endpoint, but also lets you view your Sandbox network.

https://explorer.dappflow.org/

![Dappflow Explorer](/imgs/dappflow.png)
