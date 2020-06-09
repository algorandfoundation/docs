title: Overview

Algorand has three public networks: [**MainNet**](mainnet.md), [**TestNet**](testnet.md), and [**BetaNet**](betanet.md). This section provides details about each of these networks that will help you validate the integrity of your connection to them. Learn more about [Choosing a Network](../../build-apps/setup.md#choosing-a-network) in the **Start Building** section. 

<center>
**Diagram of the relationship between the three public networks.**

![Algorand Public Networks](../../imgs/network.png)

</center>

Each network page, contains the following information:

- **What's New** ([B](./betanet.md#whats-new))- An optional section where new features are highlighted with links to accompanying documentation.
- **Version** ([M](./mainnet.md#version), [T](./testnet.md#version), [B](./betanet.md#version)) - The latest protocol software version.  Should match `goal -v` or `GET /versions` [build version](../rest-apis/algod/v1.md#buildversion).
- **Release Version** ([M](./mainnet.md#release-version), [T](./testnet.md#release-version), [B](./betanet.md#release-version)) - A link to the official release notes where you can view all the latest changes.
- **Genesis ID** ([M](./mainnet.md#genesis-id), [T](./testnet.md#genesis-id), [B](./betanet.md#genesis-id)) - A human-readable identifier for the network. This should not be used as a unique identifier.
- **Genesis Hash** ([M](./mainnet.md#genesis-hash), [T](./testnet.md#genesis-hash), [B](./betanet.md#genesis-hash)) - The unique identifier for the network, present in every transaction. Validate that your transactions match the network you plan to submit them to.
- **FeeSink Address** ([M](./mainnet.md#feesink-address), [T](./testnet.md#feesink-address), [B](./betanet.md#feesink-address)) - Read more about special accounts [here](../../features/accounts/index.md#special-accounts).
- **RewardsPool Address** ([M](./mainnet.md#rewardspool-address), [T](./testnet.md#rewardspool-address), [B](./betanet.md#rewardspool-address)) - Read more about special accounts [here](../../features/accounts/index.md#special-accounts).
- **Faucet** ([T](./testnet.md#faucet), [B](./betanet.md#faucet)) - Link to a faucet (TestNet and BetaNet only).


