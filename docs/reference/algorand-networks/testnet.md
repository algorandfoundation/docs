title: TestNet

# What's New
**algod v2** was deployed to TestNet on June 16, 2020. Visit the new docs to get started:

- [Connect with the v2 Client](../../build-apps/connect.md)
- [Your First Transaction](../../build-apps/hello_world.md)
- [v2 SDK Migration Guide](../sdks/migration.md)
- [v2 REST API](../rest-apis/algod/v2.md)
- [Indexer V2 Feature Guide](../../features/indexer.md)
- [Installing the Indexer](../../run-a-node/setup/indexer.md)
- [Indexer REST API](../rest-apis/indexer.md)
  
# Version
`v2.0.8.stable`

# Release Version
https://github.com/algorand/go-algorand/releases/tag/v2.0.8-stable

# Genesis ID
`testnet-v1.0`

# Genesis Hash
`SGO1GKSzyE7IEPItTxCByw9x8FmnrCDexi9/cOUJOiI=`

# FeeSink Address
`A7NMWS3NT3IUDMLVO26ULGXGIIOUQ3ND2TXSER6EBGRZNOBOUIQXHIBGDE`

# RewardsPool Address
`7777777777777777777777777777777777777777777777777774MSJUVU`

# Faucet

You can dispense algos to your TestNet account by visiting [Algorand TestNet Dispenser](https://bank.testnet.algorand.network/).

Click the Google sign-in link at the top of the Dispenser page and sign-in with your Google account. Complete the reCAPTCHA. Then copy and paste the address you would like to send money to into the empty text box. Click "Dispense". A `200` status code and a transaction ID means success. Run `goal account balance -a <algorand-address>` to view the new balance.
