title: BetaNet

# Version
v2.0.10.beta

# Release Version
https://github.com/algorand/go-algorand/releases/tag/v2.0.10-beta

# Genesis ID
betanet-v1.0

# Genesis Hash
mFgazF+2uRS1tMiL9dsj01hJGySEmPN28B/TjjvpVW0

All of this information can be found by running a node on the network of your choice and running `algod -v` to return the Version and `goal node status` for the Protocol Version, Genesis ID and Genesis Hash.

# BetaNet Dispenser

You can dispense algos to your TestNet account by visiting [Algorand BetaNet Dispenser](https://bank.betanet.algodev.network/).

Click the Google sign-in link at the top of the Dispenser page and sign-in with your Google account.
Complete the reCAPTCHA.
Copy and paste the public key of the target address you would like to send money to into the empty text box.
Click "Dispense".
A 200 status code and a transaction ID means success. You can run goal account balance -a <pubkey> -d <data-dir> to view the new balance.

# BetaNet Block Explorer
The BetaNet block explorer can be found here: https://betanet.algoexplorer.io/ you can select which network you'd like to point to by making your selection on the network drop down menu to the right.

# Switch Network to BetaNet
Instructions on configuring your node for BetaNet can be found in the <LINK> page.



