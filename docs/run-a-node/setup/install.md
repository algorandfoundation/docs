title: Install a Node
# Overview
This guide explains how to install the Algorand Node software on Linux Distributions and Mac OS. When installing on Linux, three methods are covered: *RPM installs*, *Debian installs* and *Other Linux Distributions*. The *Other Linux Distributions* method has been verified on Ubuntu, CentOS, Fedora, Raspian (Raspberry Pi 3) and allows manually setting data directories and requires manual updates. The *RPM* and *Debian* installs use fixed directories and automatically update.

A node installation consists of two folders: the binaries (bin) and the data (data) folders. The bin folder can be created anywhere, but Algorand recommends `~/node`, if not using *RPM* or *Debian* installs. This location is referenced later in the documentation. Remember to replace this location in the documentation below with the correct location. It is assumed that this folder is dedicated to Algorand binaries and is archived before each update. Note that nothing is currently deleted, the binaries for Algorand are just overwritten.

When installing for the first time a `data` directory will need to be specified unless using the *RPM* or *Debian* install. Algorand recommends using a location under the `node` folder, e.g. `~/node/data`. See [Node Artifacts](../../reference/node/artifacts.md) reference for a detailed list of all files that are installed. An environment variable can be set that points to the data directory and goal will use this location if a specific `data` folder is not specified.

```
export ALGORAND_DATA=~/node/data 
```

When installing with *Debian* or *RPM* packages the binaries will be installed in the `/usr/bin` and the data directory will be set to `/var/lib/algorand`. It is advisable with these installs to add the following export to shell config files.

```
export ALGORAND_DATA=/var/lib/algorand
```

!!! info
    When installing with the *Debian* or *RPM* packages, `kmd` related files such as the kmd token file will be written to the `${HOME}/.algorand/kmd-version` directory. These files are primarily used with the SDKs and REST endpoints. See [REST Endpoints](../../reference/sdks/index.md#rest-endpoints) for more details.
 

# Installation Overview
Installing a new node is generally a 3 to 4-step process and will depend on the operating system. Each install option is listed in this guide and is accessible from the table of contents or select from the list below:

1. Installing on a Mac
2. Installing with Debian
3. Installing on Other Linux Distros


# Installing on a Mac
Verified on OSX v10.13.4 / High Sierra.

+ Create a temporary folder to hold the install package and files.

```
mkdir ~/node
cd ~/node
```

Download the updater script.
  
```
curl https://raw.githubusercontent.com/algorand/go-algorand-doc/master/downloads/installers/update.sh -O
```  

+ Ensure that your system knows it's an executable file.
  
```
chmod 544 update.sh
```  

+ Run the installer from within your node directory.

```
./update.sh -i -c stable -p ~/node -d ~/node/data -n
```
When the installer runs, it will pull down the latest update package from S3 and install it. The `-n` option above tells the installer to not auto-start the node. If the installation succeeds the node will need to be started manually described later in this [guide](#start-node). 

!!! info
    When installing the `rel/beta` release, specify the beta channel `-c beta`

# Installing with Debian
Nodes have been verified on Ubuntu 18.04. Other Debian-based distros should work as well (use apt-get install rather than apt install).

+ Open a terminal and run the following commands.

```
sudo apt-get update
sudo apt-get install -y gnupg2 curl software-properties-common
curl -O https://releases.algorand.com/key.pub
sudo apt-key add key.pub
sudo add-apt-repository "deb https://releases.algorand.com/deb/ stable main"
sudo apt-get update
sudo apt-get install -y algorand
algod -v
```

These commands will install and configure `algod` as a service and place the algorand binaries in the `/usr/bin` directory. These binaries will be in the path so the `algod` and `goal` commands can be executed from anywhere. Additionally, every node has a data directory, in this case, it will be set to `/var/lib/algorand`. 


This install defaults to the Algorand MainNet network. See switching networks<LINK> for details on changing to another network.


# Installing with RPM
Installing on Fedora and Centos are described below. 

+ To install to CentOS, open a terminal and run the following commands.
  
```
curl -O https://releases.algorand.com/rpm/rpm_algorand.pub
sudo rpmkeys --import rpm_algorand.pub
sudo yum install yum-utils
sudo yum-config-manager --add-repo https://releases.algorand.com/rpm/stable/algorand.repo
sudo yum install algorand
```

+ To install to Fedora open a terminal and run the following commands.
  
```
dnf install -y 'dnf-command(config-manager)'
dnf config-manager --add-repo=https://releases.algorand.com/rpm/stable/algorand.repo
dnf install algorand
```

These commands will install and configure `algod` as a service and place the algorand binaries in the `/usr/bin` directory. These binaries will be in the path so the `algod` and `goal` commands can be executed from anywhere. Additionally, every node has a data directory, in this case, it will be set to `/var/lib/algorand`. 

This install defaults to the Algorand MainNet network. See switching networks<LINK> for details on changing to another network.

# Installing with Other Linux Distros
Nodes have been verified on Ubuntu, CentOS, Fedora, Raspian (Raspberry Pi 1-4). Other modern distros should work as well.

+ Create a temporary folder to hold the install package and files.

```
mkdir ~/node
cd ~/node
```

Download the updater script.
  
```
wget https://raw.githubusercontent.com/algorand/go-algorand-doc/master/downloads/installers/update.sh
```  

+ Ensure that your system knows it's an executable file.
  
```
chmod 544 update.sh
```  

+ Run the installer from within your node directory.

```
./update.sh -i -c stable -p ~/node -d ~/node/data -n
```

When the installer runs, it will pull down the latest update package from S3 and install it. The `-n` option above tells the installer to not auto-start the node. If the installation succeeds the node will need to be started manually described later in this [guide](#start-node). 


# Configure Telemetry
Algod is instrumented to provide telemetry which is used for insight into the software's performance and usage. Telemetry is disabled by default and so no data will be shared with Algorand Inc. Enabling telemetry provides data to Algorand to improve the software and help to identify issues. Telemetry can be enabled by following the commands below replacing &lt;name&gt; with your desired hostname (e.g. 'SarahsLaptop').

```
cd ~/node
./diagcfg telemetry name -n <name>
```
Telemetry can also be provided without providing a hostname:

```
./diagcfg telemetry enable
```

!!! info
    Telemetry can be disabled at any time by using the `./diagcfg telemetry disable` command.

Running the `diagcfg` commands will create and update the logging configuration settings stored in ~/.algorand/logging.config.

# Start Node 

+ The *Debian* and *RPM* installs automatically start the node. Starting and stopping a node installed with one of these packages should be done using `systemctl` commands:

```
sudo systemctl start algorand
```

```
sudo systemctl stop algorand
```


With these installs, the status of the node can be checked by running:

```
goal node status -d /var/lib/algorand
```

+ The *Mac* or *Other Linux Distros* installs require that the node be be started manually. This can be done from the `~node` directory with the following command:

```
./goal node start -d data
```

This will start the node and it can be verified by running:

```
pgrep algod
```

The node can be manually stopped by runnning:

```
./goal node stop -d data
```


# Sync Node with Network
When a node first starts, it will need to sync with the network. This process can take a while as the node is loading up the current ledger and catching up to the rest of the network. The status can be checked by running the following goal command:

```
./goal node status -d data
```

The goal node status command will return information about the node and what block number it is currently processing. When the node is caught up with the rest of the network, the "Sync Time" will be 0.0 as in the example response below (if on MainNet, some details will be different).

```
Last committed block: 125064
Time since last block: 3.1s
Sync Time: 0.0s
Last consensus protocol: https://github.com/algorandfoundation/specs/tree/5615adc36bad610c7f165fa2967f4ecfa75125f0
Next consensus protocol: https://github.com/algorandfoundation/specs/tree/5615adc36bad610c7f165fa2967f4ecfa75125f0
Round for next consensus protocol: 125065
Next consensus protocol supported: true
Genesis ID: testnet-v1.0
Genesis hash: SGO1GKSzyE7IEPItTxCByw9x8FmnrCDexi9/cOUJOiI=
```

# Updating Node
The *RPM* or *Debian* packages are updated automatically. For other installs, check for and install the latest updates by running `./update.sh -d ~/node/data` at any time from within your node directory. It will query S3 for available builds and see if there are newer builds than the currently installed version. To force an update, run `./update.sh -i -c stable -d ~/node/data`. 

If there is a newer version, it will be downloaded and unpacked. The node will shutdown, the binaries and data files will be archived, and the new binaries will be installed. If any part of the process fails, the node will restore the previous version (bin and data) and restart the node. If it succeeds, the new version is started. The automatic start can be disabled by adding the `-n` option.

Setting up a schedule to automatically check for and install updates can be done with CRON.

```
crontab -e 
```

Add a line that looks like this (run update.sh every hour, on the half-hour, of every day), where ‘user’ is the name of the account used to install / run the node:

```
30 * * * * /home/user/node/update.sh -d /home/user/node/data >/home/user/node/update.log 2>&1
```

# DNS Configuration for betanet
For the `betanet` network, when installing a new node or relay, make the following modification to the `config.json` file located in the node's data directory. 
First, if there is not a config.json, make a copy of the config.json.example file.  
```
cp config.json.example config.json
```
Then edit the config.json file and replace the line
``` 
"DNSBootstrapID": "<network>.algorand.network",
```
with 
``` 
"DNSBootstrapID": "<network>.algodev.network",
```
This modification to the `DNSBootstrapID` is only required for the `betanet` network.
