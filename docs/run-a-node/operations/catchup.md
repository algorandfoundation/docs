title: Catchup And status

# Catchup

When first starting a node it will process all blocks in the blockchain, even if it does not store all blocks locally. The node does so to verify every block in the blockchain thereby checking the validity of the chain. The process can be time-consuming but is essential when running a trusted node.  

With adequate [hardware and Internet connection](../../setup/install#hardware-requirements), syncing MainNet takes between 2 and 4 weeks in October 2022.

If you cannot wait for catchup, there are multiple options:
* [AlgoKit](/docs/get-started/algokit) allows quick setup of private network and public networks. For public networks, the node will be non-archival and use fast catchup. Sandbox should only be used for development purposes.
* [Fast Catchup](../../setup/install#sync-node-network-using-fast-catchup) can be used to quickly sync a **non-archival** node, but requires trust in the entity providing the catchpoint.
* Third-party snapshots such [AlgoNode.io snapshots](https://algonode.io/extras/) may be used for **archival** nodes, but requires trust in the third party. Algorand denies any responsibility if any such snapshot is used.

!!! info    
    If a node is stopped it will stop processing blocks. Once the node is restarted, it will start processing blocks where it left off.

!!! info
    The terms "sync" and "catchup/catch up" are used interchangeably.

# Node Status

It is possible to check the status of the catchup process by checking a node's status.

```
goal node status [-d <data directory>]
```

After running this status check, monitor the `Sync Time:` property that is returned. If this value is incrementing, the node is still synching. The `Sync Time:` will display `Sync Time: 0.0s` when the node is fully caught up. The status also reports the last block process by the node in the `Last committed block:` property. Comparing this block number to what is shown using an Algorand [Block Explorer](https://developer.algorand.org/ecosystem-projects/#block-explorers) will indicate how much more time catchup will take.

!!! info
    Progress of last committed block is not linear. Earlier blocks usually have much fewer transactions than more recent blocks and catchup usually slows down the closer to the current block you are.


