title: Configure Node as a Relay

A benefit of Algorand's decentralized network implementation is that a relay is effectively the same as any other node. The distinction currently is made by configuring a node to actively listen for connections from other nodes and having itself advertised using SRV records available through DNS. 

It is possible to set up a relay for a personal network that does not require DNS entries. This is done using the following steps.

# Install Node 
Follow the [install instructions](install.md) for the specific operating system that the relay will run on.

# Edit the Configuration File
Edit the configuration file for the node as described in the [configuration](config.md) guide. Set the property `NetAddress` to `":4161"` and save the file. Make sure the file is named `config.json`.

!!! warning
	As a precaution, it is not recommended that relay nodes interact with accounts or participate in consensus.

# Start the Node
Start the node as described in the [install](install.md) guide. The node will now listen for incoming traffic on port 4161. Other nodes can now connect to this relay.

# Connect a Node to Relay
Any node can connect to this relay by specifying it in the `goal node start` command. 

```
./goal node start -d data -p "ipaddress:4161"
```

The node can also be set up to connect to multiple relays using a `;` separated list of relays.

```
./goal node start -d data -p "ipaddress-1:4161;ipaddress-2:4161"
```

!!! warning
	Using the above process will prevent the node from connecting to any of the Algorand networks. See the [Phonebook](../../../reference-docs/node_files/#phonebookjson) documentation for more information on how nodes connect to relays.



