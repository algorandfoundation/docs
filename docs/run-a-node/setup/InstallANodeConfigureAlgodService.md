title: Installing algod as a systemd service

After installing using the Updater script method, there are several shell scripts that are bundled into the tarball that are helpful in running `algod`. One of those is the `systemd-setup.sh` script to create a system service.

- username, user run command that will run algorand.service
- group, the user's group found in /etc/group
- [bindir], an optional binary directory (`bindir`) parameter.

``` 
Usage: ./systemd-setup.sh username group [bindir]
``` 

---+++ Note
Skip this section if you already installed using package manager for Debian or Red Hat distributions and want to customize the default install configuration of algorand.service.

### Installing system-wide

This section explains how to install system-wide service assuming you installed algorand with Updater script method. To install `algod` as a system-wide service, run the script with root privileges:

``` 
cd ~/node
sudo ./systemd-setup.sh algorand algorand
``` 

This will create the service in `/lib/systemd/system/algorand@.service` and will have used the template `algorand@.service.template` (downloaded in the same tarball) to create the service. It includes a lot of helpful information at the top of the file and is worth perusing.

The location of the binaries is needed by the template to tell `systemd` where to find `algod`. This can be controlled by the `bindir` parameter, which is the third parameter when calling the shell script, and is expected to be an absolute path.

Here is a snippet of the template:

<pre>
``` 
[Service]
ExecStart=@@BINDIR@@/algod -d %I
User=@@USER@@
Group=@@GROUP@@
...
``` 
</pre>

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