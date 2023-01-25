title: Updater script Installation Overview

A node installation consists of two folders: the binaries (bin) and the data (data) folders.

- The bin folder can be created anywhere, but Algorand recommends `~/node`. This location is referenced later in the documentation.
- The data folder `~/node/data` is defined by the user.

Remember to replace this location in the documentation below with the correct location. It is assumed that this folder is dedicated to Algorand binaries and is archived before each update.

Nodes have been verified on Ubuntu, CentOS, Fedora, openSUSE Leap, Manjaro, Mageia, Alpine, and Solus. Other modern distros should work as well.

---+++Note
    If installing using the updater script nothing is currently deleted, the binaries for Algorand are just overwritten.

!!! Note
    If installing using the updater script then all the binaries are downloaded together, i.e., there is not a separate devtools archive file or package.

---+++ info
`kmd` related files such as the kmd token file will be written to the `${HOME}/.algorand/kmd-version` directory. These files are primarily used with the SDKs and REST endpoints. [Read more about kmd](../../clis/kmd).

### Installing on Linux

!!! Create a temporary folder to hold the install package and files.

``` 
mkdir ~/node
cd ~/node
``` 

Download the updater script.

``` 
wget https://raw.githubusercontent.com/algorand/go-algorand/rel/stable/cmd/updater/update.sh
``` 

!!! Ensure that your system knows it's an executable file.

``` 
chmod 744 update.sh
``` 

!!! Run updater script to install node

``` 
cd ~/node
./update.sh -i -c stable -p ~/node -d ~/node/data -n
``` 

---+++ Configure Node Data Directory

When installing for the first time a `data` directory will need to be specified. Algorand recommends using a location under the `node` folder, e.g. `~/node/data`. See [Node Artifacts](../../reference/artifacts) reference for a detailed list of all files that are installed. An environment variable can be set that points to the data directory and goal will use that variable if no `-d` flag is specified.

Additionally, it is convenient to add `~/node` to `PATH` so `goal` becomes directly executable, instead of having to constantly reference it as `./goal` in the `node` directory.

<pre>
``` 
export ALGORAND_DATA="$HOME/node/data"
export PATH="$HOME/node:$PATH"
``` 
</pre>

Note that the environment variables set by these commands are not permanent, so it is advisable to add the exports to shell config files (e.g., `~/.bashrc` or `~/.zshrc`).

Use this option when installing in the following operating systems: macOS, openSUSE Leap, Manjaro, Mageia, Alpine, Solus, etc.

Also, use this method for the Debian and Red Hat Linux distributions listed in [Install Node](./install-node) section if you want full control of the installation process (this is not recommended for most users).

!!! info
The following exports to shell config files. Hereafter, goal will default to using `$ALGORAND_DATA` as the data directory, removing the need to specify `-d ~/node/data` in every command.

---++ Start Node
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

---++ Status Node

The status of the node can be checked by running:

``` 
goal node status

goal node status -d ~/node/data

goal node status -d $HOME/node/data

goal node status -d $ALGORAND_DATA
``` 

Should see output 
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


### Updating Node

Since the Updater script is the manual approach to installing and running node there is no automatic updates. 

For other installs, check for and install the latest updates by running `./update.sh -d ~/node/data` at any time from within your node directory. Note that the `-d` argument has to be specified when updating. It will query S3 for available builds and see if there are newer builds than the currently installed version. To force an update, run `./update.sh -i -c stable -d ~/node/data`.

To update run the Updater script installer from within your node directory.

``` 
./update.sh -i -c stable -p ~/node -d ~/node/data -n
``` 

When the installer runs, it will pull down the latest update package from S3 and install it. The `-n` option above tells the installer to not auto-start the node. If the installation succeeds the node will need to be started manually described later in this [guide](#start-node).

---++ Running Scripts as Algorand User and Data Directory

After the node is installed and configured for network, the usage between goal and algod service is important to understand.

Since the data directory `~/node/data` is owned by the user `algod` is run by the user, some operations such as the ones related to wallets and accounts keys (`goal account ...` and `goal wallet ...`) need to be run as the user.

!!! Never run goal commands as root user

- Never run `goal` as `root`* (e.g., do *not* run `sudo goal ...`). Running `goal` as `root` can compromise the permissions of files in `~/node/data`.

---+++ Run goal commands as different user

First configure user for [algorand elevated permissions](./configure-user-for-algorand-elevated-permissions).

Then, if user_A installed node then always run goal commands as user_A OR run goal commands with `sudo -u user_A -E goal ...` where the different user_B running the command has elevated permissions to run as user_A.

To run goal commands as user_A, first login to user_A.

- -u, --user=user          run command (or edit file) as specified user
- -E, --preserve-env      preserve user environment when running command
- -s, --shell                   run shell as the target user; a command may also be specified

``` 
sudo -u user_A -E -s
whoami
cd ~/node
goal node status -d $ALGORAND_DATA
``` 