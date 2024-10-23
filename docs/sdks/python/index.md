title: Python SDK: Your first transaction

This section is a quick start guide for interacting with the Algorand network using Python. This guide will help to install [Algorand sandbox](https://github.com/algorand/sandbox){target=blank}, which provides a node for testing and development. This guide will also help to install the Python SDK, create an account and submit your first transaction on Algorand.  

# Install Sandbox

!!! info
    This step is only required if you are not using AlgoKit. If you are using AlgoKit, you can spin up a sandbox using the LocalNet, see [AlgoKit getting started guide](/docs/get-started/algokit/#start-a-localnet) for more information. 

!!! Prerequisites
    - Docker Compose ([install guide](https://docs.docker.com/compose/install/){target=_blank})
    - Git ([install guide](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git){target=_blank}) 

Algorand provides a docker instance for setting up a node, which can be used to get started developing quickly. To install and use this instance, follow these instructions.
​
```bash
git clone https://github.com/algorand/sandbox.git
cd sandbox
./sandbox up dev 
```

This will install and start private network. To read more about Algorand networks see [Algorand Networks](../../get-details/algorand-networks/index.md){target=_blank}. 

[More Information about the sandbox](https://developer.algorand.org/articles/introducing-sandbox-20/) and [how to use](https://developer.algorand.org/tutorials/exploring-the-algorand-sandbox/) it.

 
# Install Python SDK
Algorand provides an SDK for Python which is available as a pip package. To install the Python SDK, open a terminal and run the following command:
​
``` bash
pip3 install py-algorand-sdk
``` 

The [GitHub repository](https://github.com/algorand/py-algorand-sdk){target=_blank} contains additional documentation and examples.

See the Python SDK [reference documentation](https://py-algorand-sdk.readthedocs.io/en/latest/){target=_blank} for more information on methods.  
​
The SDK is installed and can now interact with the Sandbox created earlier.​ 

# Create an Account
In order to interact with the Algorand blockchain, you must have a funded account. To quickly create a test account use the following code.

​
<!-- ===PYSDK_ACCOUNT_GENERATE=== -->
```python
private_key, address = account.generate_account()
print(f"address: {address}")
print(f"private key: {private_key}")
print(f"mnemonic: {mnemonic.from_private_key(private_key)}")
```
[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/account.py#L5-L9)
<!-- ===PYSDK_ACCOUNT_GENERATE=== -->

[`More Information`](../../get-details/accounts/create.md#standalone){target=_blank}  
​
!!! Warning
    Never share mnemonic and private key. Production environments require stringent private key management. For more information on key management in community Wallets, click [here](https://developer.algorand.org/docs/community/#wallets){target=_blank}. For the [Algorand open source wallet](https://developer.algorand.org/articles/algorand-wallet-now-open-source/){target=_blank}, click [here](https://github.com/algorand/algorand-wallet){target=_blank}.

# Fund the Account
Before sending transactions to the Algorand network, the account must be funded to cover the minimal transaction fees that exist on Algorand. In this example, we'll be using prefunded accounts available in the Sandbox. To fund an account on Testnet account use the [Algorand faucet](https://dispenser.testnet.aws.algodev.network/){target=_blank}. 
​
!!! Info
    All Algorand accounts require a minimum balance to be registered in the ledger. To read more about Algorand minimum balance, see [Account Overview](../../get-details/accounts/index.md#minimum-balance){target=_blank}.  

# Connect Your Client
An Algod client must be instantiated prior to making calls to the API endpoints. You must provide values for `<algod-address>` and `<algod-token>`. The CLI tools implement the client natively. By default, the `algod_token` for each [sandbox](https://github.com/algorand/sandbox) is set to its `aaa...` value and the `algod_address` corresponds to `http://localhost:4001`.


<!-- ===PYSDK_ALGOD_CREATE_CLIENT=== -->
```python
# Create a new algod client, configured to connect to our local sandbox
algod_address = "http://localhost:4001"
algod_token = "a" * 64
algod_client = algod.AlgodClient(algod_token, algod_address)

# Or, if necessary, pass alternate headers

# Create a new client with an alternate api key header
special_algod_client = algod.AlgodClient(
    "", algod_address, headers={"X-API-Key": algod_token}
)
```
[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/overview.py#L10-L21)
<!-- ===PYSDK_ALGOD_CREATE_CLIENT=== -->

!!! Info
    The example code connects to the sandbox Algod client. If you want to connect to a public API client, change the host, port, and token parameters to match the API service. See some service available [here](https://developer.algorand.org/ecosystem-projects/?tags=api-services)

!!! Info
    If you are connecting to the Testnet, a dispenser is available [here](https://dispenser.testnet.aws.algodev.network/){target=_blank}


# Check Your Balance
Before moving on to the next step, make sure your account has been funded.

<!-- ===PYSDK_ALGOD_FETCH_ACCOUNT_INFO=== -->
```python
account_info: Dict[str, Any] = algod_client.account_info(address)
print(f"Account balance: {account_info.get('amount')} microAlgos")
```
[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/overview.py#L33-L35)
<!-- ===PYSDK_ALGOD_FETCH_ACCOUNT_INFO=== -->

# Build First Transaction
Transactions are used to interact with the Algorand network. To create a payment transaction use the following code.
​
<!-- ===PYSDK_TRANSACTION_PAYMENT_CREATE=== -->
```python
# grab suggested params from algod using client
# includes things like suggested fee and first/last valid rounds
params = algod_client.suggested_params()
unsigned_txn = transaction.PaymentTxn(
    sender=address,
    sp=params,
    receiver=address2,
    amt=1000000,
    note=b"Hello World",
)
```
[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/overview.py#L39-L49)
<!-- ===PYSDK_TRANSACTION_PAYMENT_CREATE=== -->
​
!!! Info
    Algorand supports many transaction types. To see what types are supported see [Transactions](../../get-details/transactions/index.md#transaction-types){target=_blank}. 

# Sign First Transaction
Before the transaction is considered valid, it must be signed by a private key. Use the following code to sign the transaction.
​
<!-- ===PYSDK_TRANSACTION_PAYMENT_SIGN=== -->
```python
# sign the transaction
signed_txn = unsigned_txn.sign(private_key)
```
[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/overview.py#L52-L54)
<!-- ===PYSDK_TRANSACTION_PAYMENT_SIGN=== -->
​
!!! Info
    Algorand provides many ways to sign transactions. To see other ways see [Authorization](../../get-details/transactions/signatures.md#single-signatures){target=_blank}. 
    
# Submit the Transaction
The signed transaction can now be submitted to the network. `wait_for_confirmation` SDK Method is called after the transaction is submitted to wait until the transaction is broadcast to the Algorand blockchain and is confirmed. 

<!-- ===PYSDK_TRANSACTION_PAYMENT_SUBMIT=== -->
```python
# submit the transaction and get back a transaction id
txid = algod_client.send_transaction(signed_txn)
print("Successfully submitted transaction with txID: {}".format(txid))

# wait for confirmation
txn_result = transaction.wait_for_confirmation(algod_client, txid, 4)

print(f"Transaction information: {json.dumps(txn_result, indent=4)}")
print(f"Decoded note: {b64decode(txn_result['txn']['txn']['note'])}")
```
[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/overview.py#L57-L66)
<!-- ===PYSDK_TRANSACTION_PAYMENT_SUBMIT=== -->

# View the Transaction

To view the transaction we submitted to the sandbox Algod, open [Lora](https://lora.algokit.io/localnet){target=_blank} and choose `LocalNet` configuration option, then search for the transaction ID. 

To view a transaction submitted to public network like testnet, open [Lora](https://lora.algokit.io/testnet){target=_blank} or [Pera Explorer](https://testnet.explorer.perawallet.app/){target=blank} and paste the transaction ID into the search bar.
