title: Catchup And Status

# Catchup
When first starting a node it will process all blocks in the blockchain, even if it does not store all blocks locally. The node does so to verify every block in the blockchain thereby checking the validity of the chain. The process can be time-consuming but is essential when running a trusted node.  

!!! info
    If interested in getting a build that provides quick access to begin development take a look at [Fast Catchup](../../setup/install#sync-node-network-using-fast-catchup),  [Docker Sandbox](../../../archive/build-apps/setup#2-use-docker-sandbox) or use a [Third-party](../../../archive/build-apps/setup#1-use-a-third-party-service) service. 

!!! info    
    If a node is stopped it will stop processing blocks. Once the node is restarted, it will start processing blocks where it left off.

# Node Status 
It is possible to check the status of the catchup process by checking a node's status.

```
goal node status [-d <data directory>]
```

After running this status check, monitor the `Sync Time:` property that is returned. If this value is incrementing, the node is still synching. The `Sync Time:` will display `Sync Time: 0.0s` when the node is fully caught up. The status also reports the last block process by the node in the `Last committed block:` property. Comparing this block number to what is shown using an Algorand [Block Explorer](../../community.md#block-explorers) will indicate how much more time catchup will take.




