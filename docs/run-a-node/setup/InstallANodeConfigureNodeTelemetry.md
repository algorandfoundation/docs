title: Configure Telemetry
Algod is instrumented to provide telemetry which is used for insight into the software's performance and usage. Telemetry is disabled by default and so no data will be shared with Algorand Inc. Enabling telemetry provides data to Algorand to improve the software and help to identify issues. 

### Enable Telemetry

Telemetry can be enabled by following the command below *as the user running `algod`* (replacing `<name>` with your desired hostname, e.g. 'SarahsLaptop'):


``` 
bash
diagcfg telemetry name -n <name>
``` 


If the Debian package is used, the above command needs to be run as the `algorand` user, i.e.:


``` 
bash
sudo -u algorand -H -E diagcfg telemetry name -n <name>
``` 


(The option `-H` is necessary to ensure `$HOME` is set to the home directory of the `algorand` user.)

Telemetry can also be provided without providing a hostname:


``` 
bash
diagcfg telemetry enable
``` 


If the Debian package is used, the above command needs to be run as the `algorand` user, i.e.:


``` 
bash
sudo -u algorand -H -E diagcfg telemetry enable
``` 


!!! Note
    After enabling (or disabling) telemetry, the node needs to be restarted.

!!! Warning
Do *not* run `diagcfg` as the `root` user: it would only enable telemetry for nodes run as the `root` user (and nodes should usually not be run as the `root` user).

In particular, do *not* run `sudo diagcfg ...`.

### Disable Telemetry

Telemetry can be disabled at any time by using (*as the user running `algod`*):


``` 
bash
diagcfg telemetry disable
``` 


If the Debian package is used, the above command needs to be run as the `algorand` user, i.e.:


``` 
bash
sudo -u algorand -H -E diagcfg telemetry
``` 


### Technical Details and Checking Telemetry is Enabled

Running the `diagcfg` commands will create and update the logging configuration settings stored in `~/.algorand/logging.config` (and `data/logging.config` if the parameter `-d data` is provided).

To check if telemetry is enabled, run (*as the user running `algod`*):


``` 
bash
diagcfg telemetry
``` 


If the Debian package is used, the above command needs to be run as the `algorand` user, i.e.:


``` 
bash
sudo -u algorand -H -E diagcfg telemetry
``` 


It is also possible to check whether `algod` is connected to the telemetry server by running:


``` 
bash
sudo netstat -an | grep :9243
``` 


If a single node/`algod` is running and telemetry is enabled, the output should look like:


``` 
plain
tcp        0      0 xxx.xxx.xxx.xxx:yyyyy        18.214.74.184:9243      ESTABLISHED
``` 


When telemetry is disabled, the above command prints nothing.

For additional detail see the [reference page on telemetry](../../reference/telemetry-config).