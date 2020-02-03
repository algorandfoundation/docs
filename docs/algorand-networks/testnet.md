title: TestNet

# Version
v2.0.2.stable

# Protocol Version
https://github.com/algorandfoundation/specs/tree/4a9db6a25595c6fd097cf9cc137cc83027787eaa

# Genesis ID
testnet-v1.0

# Genesis Hash
SGO1GKSzyE7IEPItTxCByw9x8FmnrCDexi9/cOUJOiI

All of this information can be found by running a node on the network of your choice and running `algod -v` to return the Version and `goal node status` for the Protocol Version, Genesis ID and Genesis Hash.

# TestNet Dispenser

You can dispense algos to your TestNet account by visiting [Algorand TestNet Dispenser](https://bank.testnet.algorand.network/).

Click the Google sign-in link at the top of the Dispenser page and sign-in with your Google account.
Complete the reCAPTCHA.
Copy and paste the public key of the target address you would like to send money to into the empty text box.
Click "Dispense".
A 200 status code and a transaction ID means success. You can run goal account balance -a <pubkey> -d <data-dir> to view the new balance.

# TestNet Block Explorer
There are currently two block explorers to choose from: [GoalSeeker](https://goalseeker.purestake.io/algorand/mainnet) and [AlgoExplorer](https://testnet.algoexplorer.io/)
The TestNet block explorer can be found here: https://testnet.algoexplorer.io/ you can select which network you'd like to point to by making your selection on the network drop down menu to the right (for Algo Explorer) and to the left (for GoalSeeker).

# Switch Network to TestNet
Instructions on configuring your node for TestNet can be found in the <LINK> page.

