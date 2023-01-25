title: Installation with the updater script

### Installing on a Mac
Verified on OSX v12.3.1 (Monterey).

!!! Create a folder to hold the install package and files.

<pre>
``` 
mkdir ~/node
cd ~/node
``` 
</pre>

Download the updater script.

``` 
curl https://raw.githubusercontent.com/algorand/go-algorand/rel/stable/cmd/updater/update.sh -O
``` 

!!! Ensure that your system knows it's an executable file.

``` 
chmod 744 update.sh
``` 

!!! Run the installer from within your node directory.

``` 
./update.sh -i -c stable -p ~/node -d ~/node/data -n
``` 

When the installer runs, it will pull down the latest update package from S3 and install it. The `-n` option above tells the installer to not auto-start the node. If the installation succeeds the node will need to be started manually described later in this [guide](#start-node).

!!! info
    When installing the `rel/beta` release, specify the beta channel `-c beta`

!!! info
    Add the following exports to shell config files. Hereafter, goal will default to using `$ALGORAND_DATA` as the data directory, removing the need to specify `-d ~/node/data` in every command.

<pre>
``` 
    export ALGORAND_DATA="$HOME/node/data"
    export PATH="$HOME/node:$PATH"
``` 
</pre>

!!! Note
    If installing using the updater script then all the binaries are downloaded together, i.e., there is not a separate devtools archive file or package.