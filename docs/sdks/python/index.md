title: Your First Transaction

This section is a quick start guide for interacting with the Algorand network using Python. This guide will help to install ***sandbox***, which provides a node for testing and development. This guide will also help to install the Python SDK, create an account and submit your first transaction on Algorand. 

# Alternative Guide

If you are a visual learner, try our submitting your first transaction [live demo](https://replit.com/@Algorand/gettingStartedPython#main.py) or watch a [full video](https://www.youtube.com/watch?v=ku2hFalMWmA) that explains the following steps.

# Sandbox Install

!!! Prerequisites
    - Docker Compose ([install guide](https://docs.docker.com/compose/install/))
    - Git ([install guide](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)) 

Algorand provides a docker instance for setting up a node, which can be used to get started developing. To install and use this instance, follow these instructions.
​
```bash
git clone https://github.com/algorand/sandbox.git
cd sandbox
./sandbox up testnet
```
[`Watch Video`](https://youtu.be/ku2hFalMWmA?t=23)  
[`More Information`](https://developer.algorand.org/articles/introducing-sandbox-20/)  

This will install a Sandbox node connected to the Algorand TestNet. To read more about Algorand networks see [Algorand Networks](https://developer.algorand.org/docs/reference/algorand-networks/). 

To use Indexer in the sandbox, start it to the default private network as follows. 
 
```bash
./sandbox up
```

!!! Info 
    The Indexer allows quick searching of the entire  blockchain for transactions, assets, applications and accounts in a timely manner. To learn more about this capability, see [Searching the Blockchain](https://developer.algorand.org/docs/features/indexer/). When running Algorand Sandbox for TestNet, BetaNet or MainNet, you will not have access to the Sandbox Algorand Indexer. 
 
!!! Warning
    The sandbox installation may take a few minutes to startup in order to catch up to the current block round. To learn more about fast catchup, see [Sync Node Network using Fast Catchup](https://developer.algorand.org/docs/run-a-node/setup/install/#sync-node-network-using-fast-catchup)

# Install SDK
Algorand provides an SDK for Python which is available as a pip package. To install the Python SDK, open a terminal and run the following command:
​
``` bash
pip3 install py-algorand-sdk
``` 
[`Watch Video`](https://youtu.be/ku2hFalMWmA?t=128)  ​

Alternatively, choose and download a [distribution file](https://pypi.org/project/py-algorand-sdk/#files), and run

``` bash
pip3 install [file name].
``` 

The [GitHub repository](https://github.com/algorand/py-algorand-sdk) contains additional documentation and examples.

See the Python SDK [reference documentation](https://py-algorand-sdk.readthedocs.io/en/latest/) for more information on methods.  
​
The SDK is installed and can now interact with the Sandbox created earlier.​ 

# Create an Account on Algorand
In order to interact with the Algorand blockchain, you must have a funded account. To quickly create a test account use the following code.
​
```python
from algosdk import account, mnemonic

def generate_algorand_keypair():
    private_key, address = account.generate_account()
    print("My address: {}".format(address))
    print("My private key: {}".format(private_key))
    print("My passphrase: {}".format(mnemonic.from_private_key(private_key)))
```
[`Watch Video`](https://youtu.be/ku2hFalMWmA?t=161)  
[`More Information`](https://developer.algorand.org/docs/features/accounts/create/#standalone)

!!! Tip
    Make sure to save your account's address and passphrase at a seperate place, as they will be used later on.
​
!!! Warning
    Never share Mnemonic and private key. Production environments require stringent private key management. For more information on key management in community Wallets, click [here](https://developer.algorand.org/docs/community/#wallets). For the [Algorand open source wallet](https://developer.algorand.org/articles/algorand-wallet-now-open-source/), click [here](https://github.com/algorand/algorand-wallet).

# Fund the Account
Before sending transactions to the Algorand network, the account must be funded to cover the minimal transaction fees that exist on Algorand. To fund the account use the [Algorand faucet](https://dispenser.testnet.aws.algodev.network/). 
​
[`Watch Video`](https://youtu.be/ku2hFalMWmA?t=257)  
​
!!! Info
    All Algorand accounts require a minimum balance to be registered in the ledger. To read more about Algorand minimum balance, see [Account Overview](https://developer.algorand.org/docs/features/accounts/#minimum-balance).  

# Connect Your Client
Client must be instantiated prior to making calls to the API endpoints. You must provide values for `<algod-address>` and `<algod-token>`. The CLI tools implement the client natively. Code beyond this point will be put into the ***first_transaction_example*** function to create a single script.

```python
from algosdk.v2client import algod

def first_transaction_example(private_key, my_address):
    algod_address = "http://localhost:4001"
    algod_token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
    algod_client = algod.AlgodClient(algod_token, algod_address)
```
[`Watch Video`](https://youtu.be/ku2hFalMWmA?t=322)  

!!! Info
    The example code connects to the sandbox Algod client. If you want to connect to a Purestake client, see Purestake's [code samples](https://developer.purestake.io/code-samples).

# Check Your Balance
Before moving on to the next step, make sure your account has been funded by the faucet.

```python
    account_info = algod_client.account_info(my_address)
    print("Account balance: {} microAlgos".format(account_info.get('amount')) + "\n")
```
[`Watch Video`](https://youtu.be/ku2hFalMWmA?t=381) 

# Build First Transaction
Communication with the Algorand network is performed using transactions. To create a payment transaction use the following code.
​
```python
# build transaction
from algosdk.future.transaction import PaymentTxn

    params = algod_client.suggested_params()
    # comment out the next two (2) lines to use suggested fees
    params.flat_fee = True
    params.fee = 1000
    receiver = "HZ57J3K46JIJXILONBBZOHX6BKPXEM2VVXNRFSUED6DKFD5ZD24PMJ3MVA"
    note = "Hello World".encode()

    unsigned_txn = PaymentTxn(my_address, params, receiver, 1000000, None, note)
```
[`Watch Video`](https://youtu.be/ku2hFalMWmA?t=398)
​
!!! Info
    Algorand supports many transaction types. To see what types are supported see [Transactions](https://developer.algorand.org/docs/features/transactions/). 

# Sign First Transaction
Before the transaction is considered valid, it must be signed by a private key. Use the following code to sign the transaction.
​
```python
	# sign transaction
    signed_txn = unsigned_txn.sign(private_key)
``` 
[`Watch Video`](https://youtu.be/ku2hFalMWmA?t=454)
​
!!! Info
    Algorand provides many ways to sign transactions. To see other ways see [Authorization](https://developer.algorand.org/docs/features/transactions/signatures/#single-signatures). 
    
# Submit the Transaction
The signed transaction can now be submitted to the network. `wait_for_confirmation` is called after the transaction is submitted to wait until the transaction is broadcast to the Algorand blockchain and is confirmed. For more information, see [Wait for Confirmation](https://developer.algorand.org/docs/build-apps/hello_world/#wait-for-confirmation)
​
```python
import json
import base64

    #submit transaction
    txid = algod_client.send_transaction(signed_txn)
    print("Successfully sent transaction with txID: {}".format(txid))

    # wait for confirmation	
    try:
        confirmed_txn = wait_for_confirmation(algod_client, txid, 4)  
    except Exception as err:
        print(err)
        return

    print("Transaction information: {}".format(
        json.dumps(confirmed_txn, indent=4)))
    print("Decoded note: {}".format(base64.b64decode(
        confirmed_txn["txn"]["txn"]["note"]).decode()))
    
    # utility function for waiting on a transaction confirmation
def wait_for_confirmation(client, transaction_id, timeout):
    """
    Wait until the transaction is confirmed or rejected, or until 'timeout'
    number of rounds have passed.
    Args:
        transaction_id (str): the transaction to wait for
        timeout (int): maximum number of rounds to wait    
    Returns:
        dict: pending transaction information, or throws an error if the transaction
            is not confirmed or rejected in the next timeout rounds
    """
    start_round = client.status()["last-round"] + 1
    current_round = start_round

    while current_round < start_round + timeout:
        try:
            pending_txn = client.pending_transaction_info(transaction_id)
        except Exception:
            return 
        if pending_txn.get("confirmed-round", 0) > 0:
            return pending_txn
        elif pending_txn["pool-error"]:  
            raise Exception(
                'pool error: {}'.format(pending_txn["pool-error"]))
        client.status_after_block(current_round)                   
        current_round += 1
    raise Exception(
        'pending tx not found in timeout rounds, timeout value = : {}'.format(timeout))
``` 
[`Watch Video`](https://youtu.be/ku2hFalMWmA?t=480)

# Complete Example
The complete example below illustrates how to quickly submit your first transaction.
​
```python
import json
import base64
from algosdk import account, mnemonic
from algosdk.v2client import algod
from algosdk.future.transaction import PaymentTxn

def generate_algorand_keypair():
    private_key, address = account.generate_account()
    print("My address: {}".format(address))
    print("My private key: {}".format(private_key))
    print("My passphrase: {}".format(mnemonic.from_private_key(private_key)))

# Write down the address, private key, and the passphrase for later usage
generate_algorand_keypair()

def first_transaction_example(private_key, my_address):
	algod_address = "http://localhost:4001"
	algod_token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	algod_client = algod.AlgodClient(algod_token, algod_address)

	print("My address: {}".format(my_address))
	account_info = algod_client.account_info(my_address)
	print("Account balance: {} microAlgos".format(account_info.get('amount')))

	# build transaction
	params = algod_client.suggested_params()
	# comment out the next two (2) lines to use suggested fees
	params.flat_fee = True
	params.fee = 1000
	receiver = "HZ57J3K46JIJXILONBBZOHX6BKPXEM2VVXNRFSUED6DKFD5ZD24PMJ3MVA"
	note = "Hello World".encode()

	unsigned_txn = PaymentTxn(my_address, params, receiver, 1000000, None, note)

	# sign transaction
	signed_txn = unsigned_txn.sign(private_key))

    # submit transaction
	txid = algod_client.send_transaction(signed_txn)
	print("Signed transaction with txID: {}".format(txid))

    # wait for confirmation	
	try:
		confirmed_txn = wait_for_confirmation(algod_client, txid, 4)  
	except Exception as err:
		print(err)
		return

	print("Transaction information: {}".format(
		json.dumps(confirmed_txn, indent=4)))
	print("Decoded note: {}".format(base64.b64decode(
		confirmed_txn["txn"]["txn"]["note"]).decode()))
    
    account_info = algod_client.account_info(my_address)
    print("Account balance: {} microAlgos".format(account_info.get('amount')) + "\n")

# utility function for waiting on a transaction confirmation
def wait_for_confirmation(client, transaction_id, timeout):
    """
    Wait until the transaction is confirmed or rejected, or until 'timeout'
    number of rounds have passed.
    Args:
        transaction_id (str): the transaction to wait for
        timeout (int): maximum number of rounds to wait    
    Returns:
        dict: pending transaction information, or throws an error if the transaction
            is not confirmed or rejected in the next timeout rounds
    """
    start_round = client.status()["last-round"] + 1
    current_round = start_round

    while current_round < start_round + timeout:
        try:
            pending_txn = client.pending_transaction_info(transaction_id)
        except Exception:
            return 
        if pending_txn.get("confirmed-round", 0) > 0:
            return pending_txn
        elif pending_txn["pool-error"]:  
            raise Exception(
                'pool error: {}'.format(pending_txn["pool-error"]))
        client.status_after_block(current_round)                   
        current_round += 1
    raise Exception(
        'pending tx not found in timeout rounds, timeout value = : {}'.format(timeout))

#replace private_key and my_address with your private key and your address
first_transaction_example(private_key, my_address)​
```
[`Run Code`](https://replit.com/@Algorand/gettingStartedPython#main.py)  
[`Watch Video`](https://youtu.be/ku2hFalMWmA?t=556)  
[`Go to Github`](https://github.com/algorand/docs/blob/staging/examples/start_building/v2/python/your_first_transaction.py)
​​
!!! Warning 
    In order for this transaction to be successful, the account must be funded. 

# Viewing the Transaction
To view the transaction, open the [AlgoExplorer](https://testnet.algoexplorer.io/) or [Goal Seeker](https://goalseeker.purestake.io/algorand/testnet) and paste the transaction ID into the search bar.  

[`Watch Video`](https://youtu.be/ku2hFalMWmA?t=618) 

# Setting Up Your Editor/Framework
The Algorand community provides many editors, frameworks, and plugins that can be used to work with the Algorand Network. Tutorials have been created for configuring each of these for use with Algorand. Select your Editor preference below.  

* [Setting Up VSCode](https://developer.algorand.org/tutorials/vs-code-javascript/)  
* [Algorand Studio](https://developer.algorand.org/articles/intro-algorand-studio-algorand-vs-code-extension/)  
* [Algorand Studio VSCode Extension](https://developer.algorand.org/articles/intro-algorand-studio-algorand-vs-code-extension/)  
* [AlgoDEA InteliJ Plugin](https://developer.algorand.org/articles/making-development-easier-algodea-intellij-plugin/)  
* [Algorand Builder Framework](https://developer.algorand.org/articles/introducing-algorand-builder/)  
