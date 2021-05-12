title: URI Scheme

This URI specification represents a standardized way for applications and websites to send requests and information through deeplinks, QR codes, etc. It is heavily based on Bitcoin’s [BIP-0021](https://github.com/bitcoin/bips/blob/master/bip-0021.mediawiki) and should be seen as derivative of it. The decision to base it on BIP-0021 was made to make it easy and compatible as possible for any other application.

# Specifications

## General format

Algorand URIs follow the general format for URIs as set forth in RFC 3986. The path component consists of an Algorand operation, and the query component provides additional parameters for that operation, if necessary.

Elements of the query component may contain characters outside the valid range. These must first be encoded according to UTF-8, and then each octet of the corresponding UTF-8 sequence must be percent-encoded as described in RFC 3986.

## ABNF Grammar

```
algorandurn     = "algorand://" ( payrequest / optinrequest )

payrequest      = [ "pay/" ] algorandaddress [ "?" payparams ]
algorandaddress = *base32
payparams       = payparam [ "&" payparams ]
payparam        = [ amountparam / labelparam / noteparam / assetparam / otherparam ]
amountparam     = "amount=" *digit
labelparam      = "label=" *qchar
assetparam      = "asset=" *digit
noteparam       = xnote / note
xnote           = "xnote=" *qchar
note            = "note=" *qchar
otherparam      = qchar *qchar [ "=" *qchar ]

optinrequest    = "optin/" creatableid
creatableid     = digit *digit
```

Here, "qchar" corresponds to valid characters of an RFC 3986 URI query component, excluding the "=" and "&" characters, which this specification takes as separators.

The scheme component ("algorand:") is case-insensitive, and implementations must accept any combination of uppercase and lowercase letters. The rest of the URI is case-sensitive, including the query parameter keys.

!!! Caveat
    When it comes to generation of an address' QR,  many exchanges and wallets encodes the address w/o the scheme component (“algorand:”). This is not a URI so it is OK.

# Operations

Algorand URIs can convey requests for different operations. All supported operations are listed below.

## Identify

The identify operation is used to share a user's address with others. Unlike the other operations, this operation
should have no on-chain effects.

The general format for this operation is `algorand://{address}?label={name}`.

### Query Params

- label (optional): Label for the address. This may be used to identify who the address belongs to.

### Examples

Sharing just an address:

```
algorand://TMTAD6N22HCS2LKH7677L2KFLT3PAQWY6M4JFQFXQS32ECBFC23F57RYX4
```

Sharing an address with a label:

```
algorand://TMTAD6N22HCS2LKH7677L2KFLT3PAQWY6M4JFQFXQS32ECBFC23F57RYX4?label=Silvio
```

## Payment

This operation is a request for a payment of Algos or an ASA to a recipient address.

The general format for this operation is `algorand://pay/{address}?{params}`.

!!! Note
    For backwards compatibility, including `pay/` before the recipient's address in the URI is optional, but
    highly encouraged.

### Query Params

- amount: microAlgos or smallest unit of asset 

- asset (optional): The asset id this request refers to (if Algos, simply omit this parameter) 

- xnote (optional): A URL-encoded notes field value that must not be modifiable by the user when displayed to users.

- note (optional): A URL-encoded default notes field value that the the user interface may optionally make editable by the user.

- label (optional): Label for that address (e.g. name of receiver)

### Transfer amount/size

!!! Note
    This is DIFFERENT than Bitcoin’s BIP-0021

If an amount is provided, it MUST be specified in basic unit of the asset. For example, if it’s Algos (Algorand native unit), the amount should be specified in microAlgos. All amounts MUST NOT contain commas nor a period (.) Strictly non negative integers. 

e.g. for 100 Algos, the amount needs to be 100000000, for 54.1354 Algos the amount needs to be 54135400. 

Algorand Clients should display the amount in whole Algos. Where needed, microAlgos can be used as well. In any case, the units shall be clear for the user. 

### Examples

Request 150.5 Algos from an address:

```
algorand://pay/TMTAD6N22HCS2LKH7677L2KFLT3PAQWY6M4JFQFXQS32ECBFC23F57RYX4?amount=150500000
```

Request 150 units of Asset ID 45 from an address:

```
algorand://pay/TMTAD6N22HCS2LKH7677L2KFLT3PAQWY6M4JFQFXQS32ECBFC23F57RYX4?amount=150&asset=45
```

## Opt-in

This operation is a request for opting-in to a specific asset or application. This operation does not have any query
parameters.

The general format for this operation is `algorand://optin/{id}`.

### Examples

Request opt-in to asset/application 45:

```
algorand://optin/45
```

!!! Note
    There is no different way to encode a URI for opting into an asset or application. In order to differentiate
    between the two, you need to look up whether the ID represents an asset or an application on-chain.
