title: Overview

This section covers the procedure for registering an account to participate in the Algorand Consensus Protocol. It begins with a description of important concepts and best practices before getting started. 

!!! tldr
	- Accounts participate in the Algorand consensus protocol by generating a valid participation key and then registering that key online with a special online registration transaction.
	- Accounts mark themselves online by submitting an [online key registration transaction](../../get-details/transactions#register-account-online) for a valid participation key.
	- Accounts mark themselves offline by submitting an [offline key registration transaction](../../get-details/transactions#register-account-offline). 
	- It is important to mark your account offline if it is not participating. Not doing so is bad network behavior and will decrease the honest/dishonest user ratio that underpins the liveness of the agreement protocol.
	- It is important to ensure that only the latest update of the participation key ever exists at any time.
	- In the event of node migration, hardware swap, or other similar events, it is preferable to have your participation key offline for a few rounds than to have it present on multiple nodes at the same time.

# Important Concepts

## Online Participation
In the context of this section, participation means participation in the Algorand consensus protocol. An account that participates in the Algorand consensus protocol is eligible and available to be selected to propose and vote on new blocks in the Algorand blockchain. Note that the likelihood that this account will be selected to propose and vote on blocks is proportional to its amount of stake. Read more about voting and block proposals in the [Algorand Consensus Protocol overview](../../get-details/algorand_consensus#the-algorand-consensus-protocol).

!!! info
    Participation in consensus currently does not affect the proportion of participation rewards an account receives. Participation rewards are [based on the amount](https://algorand.foundation/rewards-faq) of stake an account has, irrespective of whether the account is marked online or offline. Furthermore, currently there are no participation rewards on MainNet.

## Online/Offline Status
In order for an account to participate in consensus, it must first mark itself online. Marking an account online requires the account to have a valid participation key, which has been registered with the network by way of an [online key registration transaction](../../get-details/transactions#register-account-online), authorized by the participating account. Marking an account offline requires an [offline key registration transaction](../../get-details/transactions#register-account-offline) also authorized by the private key of the participating account.

# Guidelines for a Healthy Network

## Ensure that Online Accounts are Participating
If an account registers itself online, it is important that its participation key is online. _A participation key is online if there is a single fully-synchronized node on the Algorand network that has that key in its ledger directory._ You should always mark an account offline if it is not actually available to participate, since the network uses the online/offline status of an account to calculate block vote thresholds. If you are marked online but you are not participating, you would be considered a dishonest user and will negatively impact the voting threshold. Furthermore, if your node experiences issues you are not able to solve promptly, it is recommended that you register the account offline as soon as possible.

!!! tip
	If you keep your private keys in cold storage, it is recommended that you generate and sign enough offline transactions to be able to take the account offline in case of emergencies. Learn more about generating transactions for future rounds in the [Transactions guide](../../get-details/transactions#). Learn how to create and authorize transactions offline in the [Authorizing Transactions Offline](../../get-details/transactions/offline_transactions) guide.

## Renew participation keys before they expire
Participation keys are valid for a specific round range. Make sure to renew participation keys or mark the account offline before the current participation key expires. Your account will _not_ automatically be marked offline.

_Visit the [Renew Participation Keys section](./renew.md) for detailed instructions._

## Ensure that Participation Nodes are working properly

Monitor your participation node to ensure high performance and consistent access to your registered participation key. 
The following should be monitored:

* last committed block (`goal node status` or API) matches a third-party API service
* CPU / RAM / disk use are within thresholds
* clock is accurate (blocks are timestamped using the clock time from the block proposer's node, so keep your node clock accurate and on time)
* the participation node is sending votes and proposing blocks at the expected frequency.

## Securely Store Participation Keys

Registered participation keys that are in operation are regularly updated through the protocol so that they cannot be used to vote on earlier rounds. Essentially, the set of keys corresponding to earlier rounds are deleted after the round passes to ensure that the compromise of a participation key by a bad actor does not give the bad actor the potential to rewrite history. Because of this, it is important that there only exists a single instance of the participation key (files ending in `*.partkey`) at any time in the system. 

!!! warning
	Because of this, holding backups of participation keys is highly discouraged, unless appropriate procedures are setup to purge those backups on a regular basis.

# Technical Procedure

The rest of this section will describe how to [generate a participation key](./generate_keys.md) for an account, [mark or register the account online](./online.md) (with the generated participation key), [mark or register an account offline](./offline.md), and [renew a participation key](./renew.md) for an account.
