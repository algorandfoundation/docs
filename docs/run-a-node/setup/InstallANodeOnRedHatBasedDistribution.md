title: Overview

When installing on Red Hat based distribution the package manager also installs the systemd algod daemon service.

- The package manager method uses fixed directories and automatically updates.
- The automatic updates may restart algod service.

It has been validated on Fedora, and CentOS

To check hardware specs see [Install a Node Overview](./install-node)

---++ Install a Node on Red Hat based distributions (Fedora, CentOS, ...)

Installing on Fedora and CentOS are described below.

These commands will install and configure `algod` as a service and place the algorand binaries in the `/usr/bin` directory. These binaries will be in the path so the `algod` and `goal` commands can be executed from anywhere. The algorand user home directory is `/var/lib/algorand`. Additionally, every node has a data directory, in this case, it will be set to `/var/lib/algorand`.

- See [Node Artifacts](../../reference/artifacts) reference for a detailed list of some of files that are installed by this method.

- This install defaults to the Algorand MainNet network. See [switching networks](./switching-networks) for details on changing to another network.

The recommended method for most users on a compatible OS is to do the installation with the following package manager steps.

---++ To install to CentOS 7, open a terminal and run the following commands.


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


---++ To install to Fedora or CentOS 8 Stream, open a terminal and run the following commands.


``` 
curl -O https://releases.algorand.com/rpm/rpm_algorand.pub
sudo rpmkeys --import rpm_algorand.pub
dnf install -y 'dnf-command(config-manager)'
dnf config-manager --add-repo=https://releases.algorand.com/rpm/stable/algorand.repo
dnf install algorand
``` 


---+++ info
`kmd` related files such as the kmd token file will be written to the `${HOME}/.algorand/kmd-version` directory. These files are primarily used with the SDKs and REST endpoints. [Read more about kmd](../../clis/kmd).

---+++ Note
Most tools are included in the node binary package and do not require a separate install. There are a few additional tools (such as `pingpong`) in a separate tools package (i.e., `tools_stable_linux-amd64_2.1.6.tar.gz`).

---++ Configure Node Data Directory

The algod service is installed and configured with the default data directory `/var/lib/algorand`.

---+++ Note
Skip this section to keep the default install configuration and want to run algorand.service on MainNet.

- To customize the configuration of the preinstalled algod service follow [Edit systemd service template unit file](./Configure-algod-as-a-systemd-service#edit-template-service-file).
- To switch networks and change the data directory follow [switch networks](../operations/switch_network).


---++ Start Node
Installs by a package manager automatically start the node. Starting and stopping a node should be done using `systemctl` commands:

``` 
sudo systemctl start algorand
``` 

``` 
sudo systemctl restart algorand
``` 

``` 
sudo systemctl stop algorand
``` 

``` 
sudo systemctl status algorand
``` 

---++ Status Node

The status of the node can be checked by running:

- -u, --user=user          run command (or edit file) as specified user
- -E, --preserve-env      preserve user environment when running command
- -s, --shell                   run shell as the target user; a command may also be specified

As algorand user:
``` 
sudo -u algorand -E -s
cd /var/lib/algorand
goal node status -d /var/lib/algorand
``` 
OR

As different user (with elevated permissions set):
``` 
sudo -u algorand -E -s goal node status -d /var/lib/algorand
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

---++ Updating Node

The *RPM* or *Debian* packages are updated automatically. For installs on [Other Linux or Mac](./install-node).

If there is a newer version, it will be downloaded and unpacked. The node will shutdown, the binaries and data files will be archived, and the new binaries will be installed. If any part of the process fails, the node will restore the previous version (bin and data) and restart the node. If it succeeds, the new version is started. The automatic start can be disabled by adding the `-n` option.

Setting up a schedule to automatically check for and install updates can be done with CRON.

``` 
crontab -e
``` 

Add a line that looks like this (run update.sh every hour, on the half-hour, of every day), where &#8216;user&#8217; is the name of the account used to install / run the node:

``` 
30 * * * * /home/user/node/update.sh -d /home/user/node/data >/home/user/node/update.log 2>&1
``` 

---+++ Note
The automatic updates will restart the node, that includes the preinstalled `template unit file` algorand@.service, so if you create a custom `instance defined template unit file` to run node on TestNet or BetaNet which has a different data directory remember to disable one service before enabling the other service.

- See [Configure algod as a systemd service](./configure-algod-as-a-systemd-service)

To rule out the scenario that two services may restart check with `top` command or list systemctl service to confirm which is enabled.

``` 
systemctl list-units --type=service --all | grep algorand

systemctl list-units --type=service --state=running | grep algorand
``` 


---++ Running Scripts as Algorand User and Data Directory

After the node is installed and configured for network, the usage between goal and algod service is important to understand. The scripts can be run as algorand user or as a user who has been granted elevated permissions.

Since the data directory `/var/lib/algorand` is owned by the user `algorand` and the daemon `algod` is run as the user `algorand`, some operations such as the ones related to wallets and accounts keys (`goal account ...` and `goal wallet ...`) need to be run as the user `algorand`.

!!! Never run goal commands as root user

- Never run `goal` as `root`* (e.g., do *not* run `sudo goal ...`). Running `goal` as `root` can compromise the permissions of files in `/var/lib/algorand`.

- Always run goal commands as algorand or with `sudo -u algorand -E goal ...` where the user running the command has elevated permissions to run as algorand.

---+++ Run goal commands as algorand user

To run goal commands as algorand user, first login to algorand user.

- -u, --user=user          run command (or edit file) as specified user
- -E, --preserve-env      preserve user environment when running command
- -s, --shell                   run shell as the target user; a command may also be specified

``` 
sudo -u algorand -E -s
whoami
cd /var/lib/algorand
goal node status -d /var/lib/algorand
``` 

---+++ Run goal commands as different user

First configure user for [algorand elevated permissions](./configure-user-for-algorand-elevated-permissions). Then run `sudo -u algorand -E -s goal ...` commands.

- -u, --user=user          run command (or edit file) as specified user
- -E, --preserve-env      preserve user environment when running command
- -s, --shell                   run shell as the target user; a command may also be specified

---++ Configure ALGORAND_DATA environment variable for Algorand User

Since goal commands are run as algorand user, or from a different user, sometimes the -d argument may not be provided. The goal command will use that environment variable if no `-d` flag is specified. It is recommended to add to shell config files the following environment variable that points to the data directory:

``` 
export ALGORAND_DATA=/var/lib/algorand
``` 

Note that the environment variable set by this command is not permanent, so it is advisable to add the exports to shell config files (e.g., `~/.bashrc` or `~/.zshrc`).

Use this option when installing in the following operating systems: Debian, Ubuntu, Fedora, CentOS, and other Debian and Red Hat based Linux distributions.

For example, to list participation keys, use 

``` 
bash
sudo -u algorand -E goal account listpartkeys
``` 
(assuming the environment variable `$ALGORAND_DATA` is set to `/var/lib/algorand`) or 
``` 
bash
sudo -u algorand -E goal account listpartkey -d /var/lib/algorand
``` 
(otherwise).


---++ Installing the Devtools
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