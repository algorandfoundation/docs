title: Install a node

### Overview

This guide explains how to install the Algorand Node software on Linux distributions and Mac OS. 

When installing on Debian based distribution or Red Hat based distribution the package manager configures the systemd algod daemon service. The package manager method uses fixed directories and automatically updates. It has been validated on Debian, Ubuntu, Fedora, and CentOS.

When installing on Other Linux or Mac the updater script method will download and requires manually setting data directories and requires manual updates. It has been tested on the same Linux distributions from above, as well as on openSUSE Leap, Manjaro, Mageia, Alpine, and Solus.

Read the hardware requirements below to confirm specs:

- [Hardware requirements](#hardware-requirements)

Then install using one of the following:

- [AlgorandDocumentation.InstallANodeOnDebianBasedDistribution](Install a Node on Debian based distributions)
- [AlgorandDocumentation.InstallANodeOnRedHatBasedDistribution](Install a Node on Red Hat based distributions)
- [AlgorandDocumentation.InstallANodeOnOtherLinux](Install a Node on Other Linux)
- [AlgorandDocumentation.InstallANodeOnMac](Install a Node on a Mac)

After installation is complete, the following configurations are available for customization:

- [AlgorandDocumentation.GivePermissionsRunAsAlgorandUser](Give user permissions to run as algorand user)
- [AlgorandDocumentation.InstallANodeConfigureAlgodService](Configure algod as a systemd service)
- [AlgorandDocumentation.InstallANodeConfigureNodeTelemetry](Configure node to use Telemetry)
- [AlgorandDocumentation.InstallANodeSyncNodeFastCatchup](Sync node using Fast Catchup)
- [AlgorandDocumentation.InstallANodeSwitchNetworks](Switching Networks - BetaNet, TestNet, MainNet)

!!! warning
Do not mix and match installation methods as this can lead to hard-to-debug issues. If the package manager method is available in your environment, we strongly recommend using only this method.

!!! tip
Windows users may choose to use [Rand Labs](https://github.com/randlabs/algorand-windows-node/) installation binaries.

!!! tip
If you are a developer and want to use a private network, [sandbox](../../get-started/devenv/sandbox.md) is often simpler than installing a node manually. However, it is still recommended to install the Algorand software without running a node, to get access to the developer tools such as `msgpacktool` and `algokey`.


### Hardware requirements

(Last update to this section: November 3, 2022.)

Due to the higher TPS on MainNet, to successfully run an Algorand MainNet node, the following hardware is necessary:

- at least 4GB of RAM (8GB strongly recommended)
- a not-too-slow SSD: HDD and SD cards are too slow for a MainNet node and will most likely prevent the node to sync
- at least 100Mbps connection (1Gbps recommended)

Participation nodes (especially those with high stake) and relays have higher requirements to ensure the performance of the overall blockchain.

Recommended system specification for participation nodes is:

- 8 vCPU
- 16 GB RAM
- 100 GB NVMe SSD or equivalent
- 1 Gbps connection with low latency

Recommended system specification for relay nodes is:

- 16 vCPU
- 32 GB RAM
- 3 TB NVMe SSD or equivalent
- 30 TB/month egress
- 1 Gbps connection with very low latency

While directly-attached NVMe SSD are recommended, in October 2022, the use of AWS EBS gp3 was able to provide sufficient performance. Users choosing this option should constantly monitor performance of the node, and may need to upgrade to faster storage solutions in the future (e.g., in case of important increase of TPS support).

The third-party website [Algoscan Analytics](https://developer.algoscan.app/) indicates the current size of the data folder for MainNet/TestNet/BetaNet archival nodes.

!!! info
Private networks used for development require much lower specs as they are usually achieving much lower TPS. A Raspberry Pi with 2GB of RAM is sufficient for low-TPS private networks.