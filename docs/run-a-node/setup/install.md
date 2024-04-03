title: Install a node
# Overview

This guide explains how to install the Algorand Node software on Linux distributions and Mac OS. When installing on Linux, two installation methods are covered: by package manager and by updater script.

[The package manager method](#installation-with-a-package-manager) uses fixed directories and automatically updates. It has been validated on Debian, Ubuntu, Fedora, and CentOS.

[The updater script method](#installation-with-the-updater-script) allows manually setting data directories and requires manual updates. It has been tested on the same Linux distributions from above, as well as on openSUSE Leap, Manjaro, Mageia, Alpine, and Solus.

!!! warning
    Do not mix and match installation methods as this can lead to hard-to-debug issues. If the package manager method is available in your environment, we strongly recommend using only this method.

!!! tip
    Windows users may choose to build their own native binaries. Rand Labs provides a [repository](https://github.com/randlabs/algorand-windows-node/) with instructions on how to accomplish this, along with an [accompanying tutorial](https://developer.algorand.org/tutorials/compile-and-run-the-algorand-node-natively-windows/).

!!! tip
    If you are a developer and want to use a private network, [AlgoKit](/docs/get-started/algokit) is often simpler than installing a node manually. However, it is still recommended to install the Algorand software without running a node, to get access to the developer tools such as `msgpacktool` and `algokey`.

### Hardware Requirements

(Last update to this section: April 3, 2024.)

Due to the higher TPS on MainNet, to successfully run an Algorand MainNet node, the following hardware is necessary:

* 8GB of RAM (Some node operators successfully use 4GB, but it requires additional memory management and may still lag behind during extremely high network activity)
* A not-too-slow SSD: HDD and SD cards are too slow for a MainNet node and will most likely prevent the node to sync
* at least 100Mbps connection (1Gbps recommended)

Participation nodes (especially those with high stake) and relays have higher requirements to ensure the performance of the overall blockchain.

Recommended system specification for **participation** nodes is:

* 8 vCPU
* 16 GB RAM
* 100 GB NVMe SSD or equivalent
* 1 Gbps connection with low latency

Recommended system specification for **archival non-relay** nodes is:

* 8 vCPU
* 16 GB RAM
* 3 TB SSD (blocks and catchpoint storage - `ColdDataDir` in the [configuration file](https://developer.algorand.org/docs/run-a-node/reference/config/#algod-configuration-settings))
* 100 GB NVMe SSD (accounts - `HotDataDir` in the [configuration file](https://developer.algorand.org/docs/run-a-node/reference/config/#algod-configuration-settings))
* 5 TB/month egress
* 1 Gbps connection with low latency

Recommended system specification for **non-archival relay** (aka light relay) nodes is:

* 16 vCPU
* 32 GB RAM
* 256 GB NVMe SSD or equivalent
* 30 TB/month egress
* 1 Gbps connection with very low latency

Recommended system specification for **archival relay** nodes is:

* 16 vCPU
* 32 GB RAM
* 3 TB SSD (blocks and catchpoint storage - `ColdDataDir` in the [configuration file](https://developer.algorand.org/docs/run-a-node/reference/config/#algod-configuration-settings))
* 100 GB NVMe SSD (accounts - `HotDataDir` in the [configuration file](https://developer.algorand.org/docs/run-a-node/reference/config/#algod-configuration-settings))
* 30 TB/month egress
* 1 Gbps connection with very low latency

While directly-attached NVMe SSD are recommended, in October 2022, the use of AWS EBS gp3 was able to provide sufficient performance. Users choosing this option should constantly monitor performance of the node, and may need to upgrade to faster storage solutions in the future (e.g., in case of important increase of TPS support).


!!! info
    Private networks used for development require much lower specs as they are usually achieving much lower TPS. A Raspberry Pi with 2GB of RAM is sufficient for low-TPS private networks.
    

### Docker Images

Official Docker images are available from [Docker Hub](https://hub.docker.com/r/algorand/algod).


### Package manager installation overview

See [Node Artifacts](../../reference/artifacts) reference for a detailed list of some of files that are installed by this method. An environment variable can be set that points to the data directory and goal will use that variable if no `-d` flag is specified. The binaries will be installed in the `/usr/bin` and the data directory will be set to `/var/lib/algorand`. It is recommended to add to shell config files the following environment variable that points to the data directory:

```
export ALGORAND_DATA=/var/lib/algorand
```

Note that the environment variable set by this command is not permanent, so it is advisable to add the exports to shell config files (e.g., `~/.bashrc` or `~/.zshrc`).

Use this option when installing in the following operating systems: Debian, Ubuntu, Fedora, CentOS, and other Debian and Red Hat based Linux distributions.

!!! info
    `kmd` related files such as the kmd token file will be written to the `${HOME}/.algorand/kmd-version` directory. These files are primarily used with the SDKs and REST endpoints.


### Updater script installation overview
A node installation consists of two folders: the binaries (bin) and the data (data) folders. The bin folder can be created anywhere, but Algorand recommends `~/node`. This location is referenced later in the documentation. Remember to replace this location in the documentation below with the correct location. It is assumed that this folder is dedicated to Algorand binaries and is archived before each update. Note that nothing is currently deleted, the binaries for Algorand are just overwritten.

When installing for the first time a `data` directory will need to be specified. Algorand recommends using a location under the `node` folder, e.g. `~/node/data`. See [Node Artifacts](../../reference/artifacts) reference for a detailed list of all files that are installed. An environment variable can be set that points to the data directory and goal will use that variable if no `-d` flag is specified.

Additionally, it is convenient to add `~/node` to `PATH` so `goal` becomes directly executable, instead of having to constantly reference it as `./goal` in the `node` directory.

```
export ALGORAND_DATA="$HOME/node/data"
export PATH="$HOME/node:$PATH"
```

Note that the environment variables set by these commands are not permanent, so it is advisable to add the exports to shell config files (e.g., `~/.bashrc` or `~/.zshrc`).

Use this option when installing in the following operating systems: macOS, openSUSE Leap, Manjaro, Mageia, Alpine, Solus, etc. 
Also, use this method for the Linux distributions listed in the previous section if you want full control of the installation process (this is not recommended for most users).


# Installation with a package manager

This is the recommended method for most users on a compatible OS.

## Debian based distributions (Debian, Ubuntu, Linux Mint, ...)
Nodes have been verified on Ubuntu 18.04 and 20.04, as well as on Debian 11. Other Debian-based distros should work as well (use apt-get install rather than apt install).

+ Open a terminal and run the following commands.

```
sudo apt-get update
sudo apt-get install -y gnupg2 curl software-properties-common
curl -o - https://releases.algorand.com/key.pub | sudo tee /etc/apt/trusted.gpg.d/algorand.asc
sudo add-apt-repository "deb [arch=amd64] https://releases.algorand.com/deb/ stable main"
sudo apt-get update

# To get both algorand and the devtools:
sudo apt-get install -y algorand-devtools

# Or, to only install algorand:
sudo apt-get install -y algorand

algod -v
```

These commands will install and configure `algod` as a service and place the algorand binaries in the `/usr/bin` directory. These binaries will be in the path so the `algod` and `goal` commands can be executed from anywhere. Additionally, every node has a data directory, in this case, it will be set to `/var/lib/algorand`.

This install defaults to the Algorand MainNet network. See [switching networks](../operations/switch_networks.md) for details on changing to another network.

!!! Note
    Most tools are included in the node binary package and do not require a separate install. There are a few additional tools (such as `pingpong`) in a separate tools package (i.e., `tools_stable_linux-amd64_2.1.6.tar.gz`).

!!! Note 
    Since the data directory `/var/lib/algorand` is owned by the user `algorand` and the daemon `algod` is run as the user `algorand`, some operations such as the ones related to wallets and accounts keys (`goal account ...` and `goal wallet ...`) need to be run as the user `algorand`. For example, to list participation keys, use 
    ```bash
    sudo -u algorand -E goal account listpartkeys
    ``` 
    (assuming the environment variable `$ALGORAND_DATA` is set to `/var/lib/algorand`) or 
    ```bash
    sudo -u algorand -E goal account listpartkey -d /var/lib/algorand
    ``` 
    (otherwise).
    
!!! Warning
     *Never run `goal` as `root`* (e.g., do *not* run `sudo goal account listpartkeys`). Running `goal` as `root` can compromise the permissions of files in `/var/lib/algorand`.


## Red Hat based distributions (Fedora, CentOS, ...)
Installing on Fedora and CentOS are described below.

+ To install to CentOS 7, open a terminal and run the following commands.

```
curl -O https://releases.algorand.com/rpm/rpm_algorand.pub
sudo rpmkeys --import rpm_algorand.pub
sudo yum install yum-utils
sudo yum-config-manager --add-repo https://releases.algorand.com/rpm/stable/algorand.repo

# To get both algorand and the devtools:
sudo yum install algorand-devtools

# Or, to only install algorand:
sudo yum install algorand
```

+ To install to Fedora or CentOS 8 Stream, open a terminal and run the following commands.

```
curl -O https://releases.algorand.com/rpm/rpm_algorand.pub
sudo rpmkeys --import rpm_algorand.pub
dnf install -y 'dnf-command(config-manager)'
dnf config-manager --add-repo=https://releases.algorand.com/rpm/stable/algorand.repo
dnf install algorand
```

These commands will install and configure `algod` as a service and place the algorand binaries in the `/usr/bin` directory. These binaries will be in the path so the `algod` and `goal` commands can be executed from anywhere. Additionally, every node has a data directory, in this case, it will be set to `/var/lib/algorand`.

This install defaults to the Algorand MainNet network. See switching networks<LINK> for details on changing to another network.

!!! Note
    Most tools are included in the node binary package and do not require a separate install. There are a few additional tools (such as `pingpong`) in a separate tools package (i.e., `tools_stable_linux-amd64_2.1.6.tar.gz`).


## Installing the Devtools
Beginning with the 2.1.5 release, there is now a new package called `algorand-devtools` that contains the developer tools.  The package contains the following binaries, some of which are new (as of 2.1.5) and some of which have been removed from the `algorand` package to decrease its size:

- carpenter
- catchupsrv
- msgpacktool
- tealcut
- tealdbg

Installing the devtools is simple and no additional entries need to be added for either `apt` or `yum` to be aware of them.  Simply install the tools as usual via the respective package manager. Since the `algorand` package is a dependency of `algorand-devtools` and the two former cannot be older than the latter, one of two possible scenarios will occur upon downloading the devtools:

- If `algorand` has not been previously installed, it will automatically download it.
- If `algorand` is installed but older than the devtools, it will automatically upgrade it.

See the examples above to understand how to install the deb and rpm packages.


## Start Node
Installs by a package manager automatically start the node. Starting and stopping a node should be done using `systemctl` commands:

```
sudo systemctl start algorand
```

```
sudo systemctl stop algorand
```

The status of the node can be checked by running:

```
goal node status -d /var/lib/algorand
```


# Installation with the updater script

## Installing on a Mac
Verified on OSX v12.3.1 (Monterey).

+ Create a folder to hold the install package and files.

```
mkdir ~/node
cd ~/node
```

Download the updater script.

```
curl https://raw.githubusercontent.com/algorand/go-algorand/rel/stable/cmd/updater/update.sh -O
```

+ Ensure that your system knows it's an executable file.

```
chmod 744 update.sh
```

+ Run the installer from within your node directory.

```
./update.sh -i -c stable -p ~/node -d ~/node/data -n
```
When the installer runs, it will pull down the latest update package from S3 and install it. The `-n` option above tells the installer to not auto-start the node. If the installation succeeds the node will need to be started manually described later in this [guide](#start-node).

!!! info
    When installing the `rel/beta` release, specify the beta channel `-c beta`

!!! info
    Add the following exports to shell config files. Hereafter, goal will default to using `$ALGORAND_DATA` as the data directory, removing the need to specify `-d ~/node/data` in every command.

    ```
    export ALGORAND_DATA="$HOME/node/data"
    export PATH="$HOME/node:$PATH"
    ```

!!! Note
    If installing using the updater script then all the binaries are downloaded together, i.e., there is not a separate devtools archive file or package.


## Installing on Linux
Nodes have been verified on Ubuntu, CentOS, Fedora, openSUSE Leap, Manjaro, Mageia, Alpine, and Solus. Other modern distros should work as well.

+ Create a temporary folder to hold the install package and files.

```
mkdir ~/node
cd ~/node
```

Download the updater script.

```
wget https://raw.githubusercontent.com/algorand/go-algorand/rel/stable/cmd/updater/update.sh
```

+ Ensure that your system knows it's an executable file.

```
chmod 744 update.sh
```

+ Run the installer from within your node directory.

```
./update.sh -i -c stable -p ~/node -d ~/node/data -n
```

When the installer runs, it will pull down the latest update package from S3 and install it. The `-n` option above tells the installer to not auto-start the node. If the installation succeeds the node will need to be started manually described later in this [guide](#start-node).

!!! info
    Add the following exports to shell config files. Hereafter, goal will default to using `$ALGORAND_DATA` as the data directory, removing the need to specify `-d ~/node/data` in every command.

    ```
    export ALGORAND_DATA="$HOME/node/data"
    export PATH="$HOME/node:$PATH"
    ```

!!! Note
    If installing using the updater script then all the binaries are downloaded together, i.e., there is not a separate devtools archive file or package.


## Start Node
Installs by the updater script require that the node be started manually. This can be done with the following command:

```
goal node start
```

This will start the node and it can be verified by running:

```
pgrep algod
```

The node can be manually stopped by running:

```
goal node stop
```


## Installing algod as a systemd service
When installing using the updater script, there are several shell scripts that are bundled into the tarball that will are helpful in running `algod`. One of those is the `systemd-setup.sh` script to create a system service.

```
Usage: ./systemd-setup.sh username group [bindir]
```

Note that this takes an optional binary directory (`bindir`) parameter. This will be discussed more in the following sections.

### Installing system-wide

To install `algod` as a system-wide service, run the script with root privileges:

```
sudo ./systemd-setup.sh algorand algorand
```

This will create the service in `/lib/systemd/system/algorand@.service` and will have used the template `algorand@.service.template` (downloaded in the same tarball) to create the service. It includes a lot of helpful information at the top of the file and is worth perusing.

The location of the binaries is needed by the template to tell `systemd` where to find `algod`. This can be controlled by the `bindir` parameter, which is the third parameter when calling the shell script, and is expected to be an absolute path.

Here is a snippet of the template:

```
[Service]
ExecStart=@@BINDIR@@/algod -d %I
User=@@USER@@
Group=@@GROUP@@
...
```

!!! Note
    If `bindir` is not provided, the script will assume the current working directory.

After installing, the script will also make `systemd` aware that the script is present on the system. However, if making changes after installation, be sure to run the following command to register those changes:

```
sudo systemctl daemon-reload
```

All that's left now is to start the service using `systemctl`. If preferred, it can also be enabled to start on system startup.

```
sudo systemctl start algorand@$(systemd-escape $ALGORAND_DATA)
```


# Configure Telemetry
Algod is instrumented to provide telemetry which is used for insight into the software's performance and usage. Telemetry is disabled by default and so no data will be shared with Algorand Inc. Enabling telemetry provides data to Algorand to improve the software and help to identify issues. 

## Enable Telemetry

Telemetry can be enabled by following the command below *as the user running `algod`* (replacing `<name>` with your desired hostname, e.g. 'SarahsLaptop'):

```bash
diagcfg telemetry name -n <name>
```

If the Debian package is used, the above command needs to be run as the `algorand` user, i.e.:

```bash
sudo -u algorand -H -E diagcfg telemetry name -n <name>
```

(The option `-H` is necessary to ensure `$HOME` is set to the home directory of the `algorand` user.)

Telemetry can also be provided without providing a hostname:

```bash
diagcfg telemetry enable
```

If the Debian package is used, the above command needs to be run as the `algorand` user, i.e.:

```bash
sudo -u algorand -H -E diagcfg telemetry enable
```

!!! Note
    After enabling (or disabling) telemetry, the node needs to be restarted.

!!! Warning
    Do *not* run `diagcfg` as the `root` user: it would only enable telemetry for nodes run as the `root` user (and nodes should usually not be run as the `root` user). In particular, do *not* run `sudo diagcfg ...`.

## Disable Telemetry

Telemetry can be disabled at any time by using (*as the user running `algod`*):

```bash
diagcfg telemetry disable
```

If the Debian package is used, the above command needs to be run as the `algorand` user, i.e.:

```bash
sudo -u algorand -H -E diagcfg telemetry
```

## Technical Details and Checking Telemetry is Enabled

Running the `diagcfg` commands will create and update the logging configuration settings stored in `~/.algorand/logging.config` (and `data/logging.config` if the parameter `-d data` is provided).

To check if telemetry is enabled, run (*as the user running `algod`*):

```bash
diagcfg telemetry
```

If the Debian package is used, the above command needs to be run as the `algorand` user, i.e.:

```bash
sudo -u algorand -H -E diagcfg telemetry
```

It is also possible to check whether `algod` is connected to the telemetry server by running:

```bash
sudo netstat -an | grep :9243
```

If a single node/`algod` is running and telemetry is enabled, the output should look like:

```plain
tcp        0      0 xxx.xxx.xxx.xxx:yyyyy        18.214.74.184:9243      ESTABLISHED
```

When telemetry is disabled, the above command prints nothing.

For additional detail see the [reference page on telemetry](../../reference/telemetry-config).

# Sync Node with Network
When a node first starts, it will need to sync with the network. This process can take a while as the node is loading up the current ledger and catching up to the rest of the network. See the section below a [Fast Catchup](#sync-node-network-using-fast-catchup) option. The status can be checked by running the following goal command:

```
goal node status
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

# Sync Node Network using Fast Catchup

Fast Catchup is a new feature and will rapidly update a node using catchpoint  snapshots. A new command on goal node is now available for catchup. The entire process should sync a node in minutes rather than hours or days. As an example, the results for a BetaNet fast catchup, at the time of writing this, was a couple minutes to get to the sync point and a few more minutes to sync the remaining blocks since the snapshot. The total blocks synced was around 4.2 million blocks and it finished syncing in under 6 minutes. Actual sync times may vary depending on the number of accounts, number of blocks and the network.  Here are the links to get the most recent catchup point snapshot per network. The results  include a round to catchup to and the provided catchpoint. Paste into the `goal node catchup` command.

BetaNet
https://algorand-catchpoints.s3.us-east-2.amazonaws.com/channel/betanet/latest.catchpoint

TestNet
https://algorand-catchpoints.s3.us-east-2.amazonaws.com/channel/testnet/latest.catchpoint

MainNet
https://algorand-catchpoints.s3.us-east-2.amazonaws.com/channel/mainnet/latest.catchpoint

The results will look similar to this:
`4420000#Q7T2RRTDIRTYESIXKAAFJYFQWG4A3WRA3JIUZVCJ3F4AQ2G2HZRA`

!!! warning
    Do **NOT** use fast catchup on an *archival* or relay node. If you ever do it, you need to reset your node and start from scratch.

!!! warning
    Fast catchup requires trust in the entity providing the catchpoint. An attacker controlling enough relays and proposing a malicious catchpoint can in theory allow a node to sync to an incorrect state of the blockchain. For full decentralization and no trust requirement, either use a catchpoint generated by one of your archival node (you can read the catchpoint using `goal node status`) or catch up from scratch without fast catchup.

Steps:

1) Start the node, if not started already, and run a status.

`goal node start`

`goal node status`

Results should look something like this...

```
Last committed block: 308
Time since last block: 0.0s
Sync Time: 6.5s
Last consensus protocol: https://github.com/algorand/spec/tree/a26ed78ed8f834e2b9ccb6eb7d3ee9f629a6e622
Next consensus protocol: https://github.com/algorand/spec/tree/a26ed78ed8f834e2b9ccb6eb7d3ee9f629a6e622
Round for next consensus protocol: 309
Next consensus protocol supported: true
Last Catchpoint:
Genesis ID: betanet-v1.0
Genesis hash: mFgazF+2uRS1tMiL9dsj01hJGySEmPN28B/TjjvpVW0=
```
2) Use the sync point captured above and paste into the catchup option

`goal node catchup 4420000#Q7T2RRTDIRTYESIXKAAFJYFQWG4A3WRA3JIUZVCJ3F4AQ2G2HZRA`

3) Run another status and results should look something like this showing a Catchpoint status:
`goal node status`

Results should show 5 Catchpoint status lines for Catchpoint, total accounts, accounts processed, total blocks , downloaded blocks.

```
Last committed block: 4453
Sync Time: 15.8s
Catchpoint: 4420000#Q7T2RRTDIRTYESIXKAAFJYFQWG4A3WRA3JIUZVCJ3F4AQ2G2HZRA
Catchpoint total accounts: 1146
Catchpoint accounts processed: 1146
Catchpoint total blocks: 1000
Catchpoint downloaded blocks: 81
Genesis ID: betanet-v1.0
Genesis hash: mFgazF+2uRS1tMiL9dsj01hJGySEmPN28B/TjjvpVW0=
```

!!! Note
    Only the first 3 status (`Catchpoint`, `total accounts`, `accounts processed`) will show right after catchup begins. Status on `total blocks` and `downloaded blocks` will only show after catchup processes the total number of accounts, which takes several minutes.

4) A new option can facilitate a status watch, -w which takes a parameter of time, in milliseconds, between two successive status updates. This will eliminate the need to repeatedly issue a status manually. Press ^c to exit the watch.

`goal node status -w 1000`

5) Notice that the 5 Catchpoint status lines will disappear when completed, and then only a few more minutes are needed to sync from that point to the current block. **Once there is a Sync Time of 0, the node is synced and if fully usable. **

```
Last committed block: 4431453
Time since last block: 3.9s
Sync Time: 0.0s
Last consensus protocol: https://github.com/algorandfoundation/specs/tree/e5f565421d720c6f75cdd186f7098495caf9101f
Next consensus protocol: https://github.com/algorandfoundation/specs/tree/e5f565421d720c6f75cdd186f7098495caf9101f
Round for next consensus protocol: 4431454
Next consensus protocol supported: true
Last Catchpoint: 4430000#UAQPNY32LP3K5ARGFUQEFTBGELI5ZAQOMBGE7YL5ZFXL2MXWTO2A
Genesis ID: betanet-v1.0
Genesis hash: mFgazF+2uRS1tMiL9dsj01hJGySEmPN28B/TjjvpVW0=
```

## Troubleshooting for fast catchup

If fast catchup fails, check the following:

* the node is *not* archival (and is not a relay).
* the software is up-to-date: `goal version -v` should report the latest `Algorand` version in https://github.com/algorand/go-algorand/releases, ignoring the `Algorand BetaNet` releases.
* the catch point matches the network used by the node and reported as `Genesis ID` by `goal node status`.
* the hardware requirements above are satisfied, in particular a not-too-slow SSD is used.
* the computer does not run out of memory.

If you used fast catchup on an archival node, you need to stop the node, remove all files in the data folder except the configuration files (`*.json`) and the genesis files (`genesis*`), and restart the node. This operation will also remove all keys stored in the node. Proceed with caution.

# Updating Node
The *RPM* or *Debian* packages are updated automatically. For other installs, check for and install the latest updates by running `./update.sh -d ~/node/data` at any time from within your node directory. Note that the `-d` argument has to be specified when updating. It will query S3 for available builds and see if there are newer builds than the currently installed version. To force an update, run `./update.sh -i -c stable -d ~/node/data`.

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
