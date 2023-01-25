title: Switch Networks

This install defaults to the Algorand MainNet? network. See [switching networks](../operations/switch_networks.md) for details on changing to another network. 

---++ DNS Configuration for betanet

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