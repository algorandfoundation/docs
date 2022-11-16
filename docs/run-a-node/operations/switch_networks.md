title: Switch Networks
By default, an Algorand installation is configured to run on MainNet. For most users, this is the desired outcome. Developers, however, need access to [TestNet or BetaNet networks](../../../archive/build-apps/setup/#choosing-a-network). This guide describes switching between networks.


# Replacing the Genesis file
Every Algorand node has a data directory that is used to store the ledger and other configuration information. As part of this configuration, a `genesis.json` file is used. The `genesis.json` file specifies the initial state of the blockchain - its ‘genesis block’. This is a JSON formatted file with the schema for the blockchain. It contains the network name and id, the protocol version and list of allocated addresses to start the chain with. Each address contains a list of things like address status and the amount of Algos owned by the address.

As part of the installer, a `genesisfiles` directory is created under the node's installed location for binaries. This directory contains additional directories for each of the Algorand networks: BetaNet, TestNet, and MainNet. These directories contain the `genesis.json` file for each of the Algorand networks (eg `~/node/genesisfiles/mainnet/genesis.json`). 

!!! info
    The genesis file for *Debian* and *RPM* installs are stored in the `/var/lib/algorand/genesis/` directory.

The network can be switched by either replacing the current genesis file located in the `data` directory with the specific network `genesis.json` or by creating a new `data` directory and copying the specific network `genesis.json` file to the new `data` directory. Replacing the current genesis file will not destroy the current network data, but will prevent running multiple networks on the same node. To run multiple networks at the same time multiple data directories are required.

# Using a new Data Directory
To construct a new data directory follow the steps described below that depend on the type of install that was used with the current node.

## Mac OS or Other Linux Distros
Assume the node binaries are currently installed in `~/node` and the data directory is set to `~/node/data`. Create the new data directory for the network that is being switched to. (eg `~/node/testnetdata`). 

```
cd ~/node
./goal node stop -d data
mkdir testnetdata 
cp ~/node/genesisfiles/testnet/genesis.json ~/node/testnetdata
```
Startup the new node using the `-d` option to point at the new network.

```
./goal node start -d ~/node/testnetdata
```

The node will restart and begin communicating with the TestNet network. It will need to sync with the network which will take time. Run the following command to check the current sync status.

```
./goal node status -d ~/node/testnetdata
```

At this point, the original network can be stared as well using the standard data directory.

## Debian or RPM installs
For *Debian* and *RPM* installs the data directory is set to `/var/lib/algorand`. With these installs, the genesis files are stored in `/var/lib/algorand/genesis/`. Create the new data directory for the network that is being switched to. (eg `/var/lib/algorand_testnet`). 

```
ALGORAND_DATA=/var/lib/algorand_testnet
sudo mkdir -p ${ALGORAND_DATA}
```
Copy the specific network genesis file to the new directory and set the ownership to the algorand account.

```
sudo cp -p /var/lib/algorand/genesis/testnet/genesis.json ${ALGORAND_DATA}/genesis.json
sudo cp -p /var/lib/algorand/system.json ${ALGORAND_DATA}/system.json
sudo chown -R algorand ${ALGORAND_DATA}
```

Use `systemctl` to enable the process and start it.

```
sudo systemctl enable algorand@$(systemd-escape ${ALGORAND_DATA})
sudo systemctl start algorand@$(systemd-escape ${ALGORAND_DATA})
```
The process can be disabled or stopped using `systemctl` commands.

```
sudo systemctl stop algorand@$(systemd-escape ${ALGORAND_DATA})
sudo systemctl disable algorand@$(systemd-escape ${ALGORAND_DATA})
```

# DNS Configuration for BetaNet

For the BetaNet network, when installing a new node or relay, make the following modification to the `config.json` file located in the node's data directory. 

Edit the `config.json` file and replace the line
``` 
"DNSBootstrapID": "<network>.algorand.network",
```
with 
``` 
"DNSBootstrapID": "<network>.algodev.network",
```
If former line is not present, just add the latter line.
If there is no `config.json` in the Algorand data folder, just create a new one with the following content:
```json
{
    "DNSBootstrapID": "<network>.algodev.network"
}
```

This modification to the `DNSBootstrapID` is only for the BetaNet network.
