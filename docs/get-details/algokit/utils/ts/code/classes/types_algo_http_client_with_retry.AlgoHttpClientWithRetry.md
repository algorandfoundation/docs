[@algorandfoundation/algokit-utils](../index.md) / [types/algo-http-client-with-retry](../modules/types_algo_http_client_with_retry.md) / AlgoHttpClientWithRetry

# Class: AlgoHttpClientWithRetry

[types/algo-http-client-with-retry](../modules/types_algo_http_client_with_retry.md).AlgoHttpClientWithRetry

A HTTP Client that wraps the Algorand SDK HTTP Client with retries

## Hierarchy

- [`URLTokenBaseHTTPClient`](types_urlTokenBaseHTTPClient.URLTokenBaseHTTPClient.md)

  ↳ **`AlgoHttpClientWithRetry`**

## Table of contents

### Constructors

- [constructor](types_algo_http_client_with_retry.AlgoHttpClientWithRetry.md#constructor)

### Properties

- [MAX\_BACKOFF\_MS](types_algo_http_client_with_retry.AlgoHttpClientWithRetry.md#max_backoff_ms)
- [MAX\_TRIES](types_algo_http_client_with_retry.AlgoHttpClientWithRetry.md#max_tries)
- [RETRY\_ERROR\_CODES](types_algo_http_client_with_retry.AlgoHttpClientWithRetry.md#retry_error_codes)
- [RETRY\_STATUS\_CODES](types_algo_http_client_with_retry.AlgoHttpClientWithRetry.md#retry_status_codes)

### Methods

- [callWithRetry](types_algo_http_client_with_retry.AlgoHttpClientWithRetry.md#callwithretry)
- [delete](types_algo_http_client_with_retry.AlgoHttpClientWithRetry.md#delete)
- [get](types_algo_http_client_with_retry.AlgoHttpClientWithRetry.md#get)
- [post](types_algo_http_client_with_retry.AlgoHttpClientWithRetry.md#post)

## Constructors

### constructor

• **new AlgoHttpClientWithRetry**(`tokenHeader`, `baseServer`, `port?`, `defaultHeaders?`)

#### Parameters

| Name | Type |
| :------ | :------ |
| `tokenHeader` | [`TokenHeader`](../modules/types_urlTokenBaseHTTPClient.md#tokenheader) |
| `baseServer` | `string` |
| `port?` | `string` \| `number` |
| `defaultHeaders` | `Record`<`string`, `any`\> |

#### Inherited from

[URLTokenBaseHTTPClient](types_urlTokenBaseHTTPClient.URLTokenBaseHTTPClient.md).[constructor](types_urlTokenBaseHTTPClient.URLTokenBaseHTTPClient.md#constructor)

#### Defined in

[src/types/urlTokenBaseHTTPClient.ts:47](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/urlTokenBaseHTTPClient.ts#L47)

## Properties

### MAX\_BACKOFF\_MS

▪ `Static` `Private` `Readonly` **MAX\_BACKOFF\_MS**: ``10000``

#### Defined in

[src/types/algo-http-client-with-retry.ts:8](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/algo-http-client-with-retry.ts#L8)

___

### MAX\_TRIES

▪ `Static` `Private` `Readonly` **MAX\_TRIES**: ``5``

#### Defined in

[src/types/algo-http-client-with-retry.ts:7](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/algo-http-client-with-retry.ts#L7)

___

### RETRY\_ERROR\_CODES

▪ `Static` `Private` `Readonly` **RETRY\_ERROR\_CODES**: `string`[]

#### Defined in

[src/types/algo-http-client-with-retry.ts:13](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/algo-http-client-with-retry.ts#L13)

___

### RETRY\_STATUS\_CODES

▪ `Static` `Private` `Readonly` **RETRY\_STATUS\_CODES**: `number`[]

#### Defined in

[src/types/algo-http-client-with-retry.ts:12](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/algo-http-client-with-retry.ts#L12)

## Methods

### callWithRetry

▸ `Private` **callWithRetry**(`func`): `Promise`<`BaseHTTPClientResponse`\>

#### Parameters

| Name | Type |
| :------ | :------ |
| `func` | () => `Promise`<`BaseHTTPClientResponse`\> |

#### Returns

`Promise`<`BaseHTTPClientResponse`\>

#### Defined in

[src/types/algo-http-client-with-retry.ts:25](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/algo-http-client-with-retry.ts#L25)

___

### delete

▸ **delete**(`relativePath`, `data`, `query?`, `requestHeaders?`): `Promise`<`BaseHTTPClientResponse`\>

#### Parameters

| Name | Type |
| :------ | :------ |
| `relativePath` | `string` |
| `data` | `Uint8Array` |
| `query?` | `Query`<`string`\> |
| `requestHeaders` | `Record`<`string`, `string`\> |

#### Returns

`Promise`<`BaseHTTPClientResponse`\>

#### Overrides

[URLTokenBaseHTTPClient](types_urlTokenBaseHTTPClient.URLTokenBaseHTTPClient.md).[delete](types_urlTokenBaseHTTPClient.URLTokenBaseHTTPClient.md#delete)

#### Defined in

[src/types/algo-http-client-with-retry.ts:67](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/algo-http-client-with-retry.ts#L67)

___

### get

▸ **get**(`relativePath`, `query?`, `requestHeaders?`): `Promise`<`BaseHTTPClientResponse`\>

#### Parameters

| Name | Type |
| :------ | :------ |
| `relativePath` | `string` |
| `query?` | `Query`<`string`\> |
| `requestHeaders` | `Record`<`string`, `string`\> |

#### Returns

`Promise`<`BaseHTTPClientResponse`\>

#### Overrides

[URLTokenBaseHTTPClient](types_urlTokenBaseHTTPClient.URLTokenBaseHTTPClient.md).[get](types_urlTokenBaseHTTPClient.URLTokenBaseHTTPClient.md#get)

#### Defined in

[src/types/algo-http-client-with-retry.ts:54](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/algo-http-client-with-retry.ts#L54)

___

### post

▸ **post**(`relativePath`, `data`, `query?`, `requestHeaders?`): `Promise`<`BaseHTTPClientResponse`\>

#### Parameters

| Name | Type |
| :------ | :------ |
| `relativePath` | `string` |
| `data` | `Uint8Array` |
| `query?` | `Query`<`string`\> |
| `requestHeaders` | `Record`<`string`, `string`\> |

#### Returns

`Promise`<`BaseHTTPClientResponse`\>

#### Overrides

[URLTokenBaseHTTPClient](types_urlTokenBaseHTTPClient.URLTokenBaseHTTPClient.md).[post](types_urlTokenBaseHTTPClient.URLTokenBaseHTTPClient.md#post)

#### Defined in

[src/types/algo-http-client-with-retry.ts:58](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/algo-http-client-with-retry.ts#L58)
