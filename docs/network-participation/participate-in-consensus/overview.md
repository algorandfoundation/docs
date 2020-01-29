title: Overview

This section covers the procedure for registering an account to participate in the Algorand Consensus Protocol. It begins with a description of important concept, then it provides an overview of the process with links to detailed procedures. 

!!! tldr
	- Accounts participate in the Algorand consensus protocol by generating a valid participation key and then registering that key online with a special online registration transaction.
	- Accounts mark themselves online by submitting an “online” key registration transaction for a valid participation key.
	- Accounts mark themselves offline by submitting an “offline” key registration transaction. 
	- It is important to mark your account offline if it is not actually participating. Not doing so is bad network behavior and will decrease the honest/dishonest user ratio that underpins the liveness of the agreement protocol.
	- It is important to ensure that only the latest update of the participation key ever exists at any time (up to allowed backups dating less than one hour, see above).
	- In the event of node migration, hardware swap, or other similar events, it is preferable to have your participation key offline for a few rounds than to have it present on multiple nodes at the same time.


# Important Concepts

## Online Participation
In the context of this section, participation means participation in the Algorand consensus protocol. An account that participates in the Algorand consensus protocol is eligible and available to be selected to propose and vote on new blocks in the Algorand blockchain. We will explain what this means in more detail later on. Note that the likelihood that this account will be selected to propose and vote on blocks is proportional to its amount of stake. 

## Online/Offline Status
In order for an account to participate in consensus, it must first mark itself online. Marking an account online requires the account to have a valid participation key, which has been registered with the network by way of an “online” key registration transaction (aka, “changeonlinestatus” transaction), signed by the private key of the participating account. Marking an account offline requires an “offline” key registration transaction also signed by the private key of the participating account.

# Best Practices

## Good Network Behavior

### Ensure that Online Accounts are Participating
If an account registers itself online, it is important that its participation key actually is online. A participation key is online if there is a single fully-synchronized node on the Algorand network that has that key in the ledger directory. You should always mark an account offline if it is not actually available to participate. Why? The network uses the online/offline status of an account to calculate block vote thresholds. Thus, if you are not actually participating, but your status shows as online you would be considered a dishonest user and will negatively impact the voting threshold. Furthermore, if your node experiences issues you are not able to solve promptly (say, within an hour), you shall put your account offline as soon as you can.

### Renew participation keys before they expire
Participation keys are valid for a specific range of rounds. You must monitor your participation keys and renew them before they expire. You can renew a participation key anytime before it expires, and we recommend to do it at least two weeks in advance. The validity ranges of participation keys can overlap. For any account, at any time, at most one participation key is registered, namely the one included in the last “online” key registration transaction for this account. (Note: An “online” or “offline” key registration actually takes effect 320 rounds after the transaction is committed to the blockchain.) 

## Ensure that Participation Nodes are running
You are required to monitor your participation node to ensure it is properly working. [TO DO]

An important additional monitoring for participation nodes is to check that your node clock is on time, as blocks in the blockchains are timestamped by the block proposer (i.e., the node proposing the block).

## Security of Participation Keys

!!! warning
	The 

To ensure the security of the blockchain, participation keys are regularly updated, so that they cannot be used to vote on earlier rounds. It is of paramount importance that there only exists a single instance of the participation key (files `*.partkey` in the ledger directory) at any time in the system, and prior instances must be securely erased. This means in particular that once you upload a participation key to a node, you must completely and securely erase it from everywhere else. We stress again that you must ensure complete secure erasure of all previous versions of the participation key files from the node, and ensure that no backup of them is left.

The only exception to the above is short-lived backup for the purpose of recovery from node failures. However, even in this case backups must be securely purged every hour. You are responsible to design a secure process, but we require the following minimum process: each backup is encrypted using a different AES symmetric encryption key, and that key is held in an HSM with FIPS-140.2 level 2 (or more) or a Cloud HSM, and both the backup and the key in the HSM are purged every hour. Failure to comply is considered a dishonest behaviour and may undermine the security of the blockchain. 

Note that only the private keys (the files `*.partkey`) on the nodes are updated. The public part of the participation key (which is included in the key registration transaction) does not change and is not erased.


