title: Configure node as a relay

A benefit of Algorand's decentralized network implementation is that a relay is effectively the same as any other node. The distinction currently is made by configuring a node to actively listen for connections from other nodes and having itself advertised using SRV records available through DNS. 

It is possible to set up a relay for a personal network that does not require DNS entries. This is done using the following steps.



# Install Node 
See this page for [node hardware requirements](../../setup/install/#hardware-requirements).
Follow the [install instructions](../../run-a-node/setup/install.md) for the specific operating system that the relay will run on.

# Edit the Configuration File
Edit the configuration file for the node as described in the [configuration](../config) guide. Set the property `NetAddress` to `":4161"` for TestNet and to `":4160"` for MainNet. Then the file. Make sure the file is named `config.json`.

Concretely, your `config.json` file should look like:

```json
{
    "NetAddress": ":4161"
}
```

for TestNet.

!!! warning
	As a precaution, it is not recommended that relay nodes interact with accounts or participate in consensus.

# Start the Node
Start the node as described in the [install](../../run-a-node/setup/install.md) guide. The node will now listen for incoming traffic on port 4161 (for TestNet) or on port 4160 (for MainNet). Other nodes can now connect to this relay.

# Connect a Node to Relay
Any node can connect to this relay by specifying it in the `goal node start` command. 

```
goal node start -p "ipaddress:4161"
```

The node can also be set up to connect to multiple relays using a `;` separated list of relays.

```
goal node start -p "ipaddress-1:4161;ipaddress-2:4161"
```

(4161 needs to be replaced by 4160 for MainNet.)

!!! warning
	Using the above process will prevent the node from connecting to any of the Algorand networks. See the [Phonebook](../artifacts#phonebookjson) documentation for more information on how nodes connect to relays.

See the [Algorand Foundation FAQ](https://www.algorand.foundation/general-faq#04-faq) for additional details about MainNet relays.

