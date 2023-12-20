# TestNet Dispenser Client

The TestNet Dispenser Client is a utility for interacting with the AlgoKit TestNet Dispenser API. It provides methods to fund an account, register a refund for a transaction, and get the current limit for an account.

## Creating a Dispenser Client

To create a Dispenser Client, you need to provide an authorization token. This can be done in two ways:

1. Pass the token directly to the client constructor as `auth_token`.
2. Set the token as an environment variable `ALGOKIT_DISPENSER_ACCESS_TOKEN` (see [docs](https://github.com/algorandfoundation/algokit/blob/main/docs/testnet_api.md#error-handling) on how to obtain the token).

If both methods are used, the constructor argument takes precedence.

```py
from algokit_utils import TestNetDispenserApiClient

# Using constructor argument

client = TestNetDispenserApiClient(auth_token="your_auth_token")

# Using environment variable

import os
os.environ["ALGOKIT_DISPENSER_ACCESS_TOKEN"] = "your_auth_token"
client = TestNetDispenserApiClient()
```

## Funding an Account

To fund an account with Algos from the dispenser API, use the `fund` method. This method requires the receiver's address, the amount to be funded, and the asset ID.

```py
response = client.fund(address="receiver_address", amount=1000, asset_id=0)
```

The `fund` method returns a `FundResponse` object, which contains the transaction ID (`tx_id`) and the amount funded.

## Registering a Refund

To register a refund for a transaction with the dispenser API, use the `refund` method. This method requires the transaction ID of the refund transaction.

```py
client.refund(refund_txn_id="transaction_id")
```

> Keep in mind, to perform a refund you need to perform a payment transaction yourself first by send funds back to TestNet Dispenser, then you can invoke this `refund` endpoint and pass the txn_id of your refund txn. You can obtain dispenser address by inspecting the `sender` field of any issued `fund` transaction initiated via [`fund`](#funding-an-account).

## Getting Current Limit

To get the current limit for an account with Algos from the dispenser API, use the `get_limit` method. This method requires the account address.

```py
response = client.get_limit(address="account_address")
```

The `get_limit` method returns a `LimitResponse` object, which contains the current limit amount.

## Error Handling

If an error occurs while making a request to the dispenser API, an exception will be raised with a message indicating the type of error. Refer to [Error Handling docs](https://github.com/algorandfoundation/algokit/blob/main/docs/testnet_api.md#error-handling) for details on how you can handle each individual error `code`.
