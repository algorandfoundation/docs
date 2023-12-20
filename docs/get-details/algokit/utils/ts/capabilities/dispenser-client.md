# TestNet Dispenser Client

The TestNet Dispenser Client is a utility for interacting with the AlgoKit TestNet Dispenser API. It provides methods to fund an account, register a refund for a transaction, and get the current limit for an account.

## Creating a Dispenser Client

To create a Dispenser Client, you need to provide an authorization token. This can be done in two ways:

1. Pass the token directly to the client constructor as `authToken`.
2. Set the token as an environment variable `ALGOKIT_DISPENSER_ACCESS_TOKEN` (see [docs](https://github.com/algorandfoundation/algokit/blob/main/docs/testnet_api.md#error-handling) on how to obtain the token).

If both methods are used, the constructor argument takes precedence.

```ts
import * as algokit from '@algorandfoundation/algokit-utils'

// Using constructor argument
const client = algokit.getTestNetDispenserApiClient({ authToken: 'your_auth_token' })

// Using environment variable
process.env['ALGOKIT_DISPENSER_ACCESS_TOKEN'] = 'your_auth_token'
const client = algokit.getTestNetDispenserApiClient()
```

## Funding an Account

To fund an account with Algos from the dispenser API, use the `fund` method. This method requires the receiver's address, the amount to be funded, and the asset ID.

```ts
const response = await client.fund('receiver_address', 1000)
```

The `fund` method returns a `DispenserFundResponse` object, which contains the transaction ID (`txId`) and the amount funded.

## Registering a Refund

To register a refund for a transaction with the dispenser API, use the `refund` method. This method requires the transaction ID of the refund transaction.

```ts
await client.refund('transaction_id')
```

> Keep in mind, to perform a refund you need to perform a payment transaction yourself first by sending funds back to TestNet Dispenser, then you can invoke this refund endpoint and pass the txn_id of your refund txn. You can obtain dispenser address by inspecting the sender field of any issued fund transaction initiated via [fund](#funding-an-account).

## Getting Current Limit

To get the current limit for an account with Algos from the dispenser API, use the `getLimit` method. This method requires the account address.

```ts
const response = await client.getLimit()
```

The `limit` method returns a `DispenserLimitResponse` object, which contains the current limit amount.

## Error Handling

If an error occurs while making a request to the dispenser API, an exception will be raised with a message indicating the type of error. Refer to [Error Handling docs](https://github.com/algorandfoundation/algokit/blob/main/docs/testnet_api.md#error-handling) for details on how you can handle each individual error `code`.
