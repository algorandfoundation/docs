title: Give Permissions to User to Run as Algorand User

After installing with package manager for Debian or Red Hat distributions and want to customize a different user to be run algorand goal commands requires elevated permissions.

Remember, the package manager installs and configures `algod` as a service and place the algorand binaries in the `/usr/bin` directory. These binaries will be in the path so the `algod` and `goal` commands can be executed from anywhere. The algorand user home directory is `/var/lib/algorand`. Additionally, every node has a data directory, in this case, it will be set to `/var/lib/algorand`.

To customize a different user, for example `algodev` user, then `algodev` user needs to be given permissions to run as `algorand` user. Change `algodev` username in this documentation to be specific to your environment.

---+++ Add file to /etc/sudoers.d/ Directory

---+++ Warning
Warning: Never edit this file with a normal text editor! Always use the visudo command instead!

Because improper syntax in the /etc/sudoers file can leave you with a broken system where it is impossible to obtain elevated privileges, it is important to use the visudo command to edit the file.

So from the example above `algodev`will be given elevated permissions to run commands as algorand user.

First, check what editor is saved as visudo editor.

On ubuntu only:
``` 
sudo update-alternatives --config editor
``` 

Check the directory first before making any changes. Confirm which files are there already. Confirm after edit too.
``` 
ls -la /etc/sudoers.d/
``` 

Then vi a new file named algodevuser

The name `algodevuser` is arbitrary, but to give relevancy to you later. Change as necessary.

As root user create new file:

``` 
bash
cd /etc/sudoers.d/

sudo visudo -f /etc/sudoers.d/algodevuser
``` 

and insert the 1 line into the file.

``` 
algodev ALL=(algorand:algorand) NOPASSWD:ALL
``` 

---+++ Note
Do not attempt to create this new file if you are not comfortable with vi commands.

---+++ After Edit

To confirm the permissions were granted for `algodev` user, login as algodev user and run a goal command as algorand user.

``` 
su -l algodev
whoami
sudo -u algorand -E -s
whoami
cd /var/lib/algorand
goal node status -d /var/lib/algorand
``` 