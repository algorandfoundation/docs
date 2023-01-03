import json
import time
import base64
from algosdk import account

from algosdk.v2client import algod
from algosdk import transaction


def getting_started_example():
	algod_address = "http://localhost:4001"
	algod_token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	algod_client = algod.AlgodClient(algod_token, algod_address)

    # Generate new account for this transaction
	secret_key, my_address = account.generate_account()
    
	print("My address: {}".format(my_address))

    # Check your balance. It should be 0 microAlgos

	account_info = algod_client.account_info(my_address)
	print("Account balance: {} microAlgos".format(account_info.get('amount')) + "\n")

    # Fund the created account
	print('Fund the created account using testnet faucet: \n https://dispenser.testnet.aws.algodev.network/?account=' + format(my_address))

	completed = ""
	while completed.lower() != 'yes':
		completed = input("Type 'yes' once you funded the account: ");

	print('Fund transfer in process...')
    # Wait for the faucet to transfer funds
	time.sleep(10)	

	account_info = algod_client.account_info(my_address)
	print("Account balance: {} microAlgos".format(account_info.get('amount')))

	# build transaction
	params = algod_client.suggested_params()
	# comment out the next two (2) lines to use suggested fees
	params.flat_fee = True
	params.fee = 1000
	receiver = "HZ57J3K46JIJXILONBBZOHX6BKPXEM2VVXNRFSUED6DKFD5ZD24PMJ3MVA"
	note = "Hello World".encode()
	amount = 1000000
	unsigned_txn = transaction.PaymentTxn(my_address, params, receiver, amount, None, note)

	# sign transaction
	signed_txn = unsigned_txn.sign(secret_key)
	txid = algod_client.send_transaction(signed_txn)
	print("Signed transaction with txID: {}".format(txid))

    # wait for confirmation	
	try:
		confirmed_txn = transaction.wait_for_confirmation(algod_client, txid, 4)  
	except Exception as err:
		print(err)
		return

	print("Transaction information: {}".format(
		json.dumps(confirmed_txn, indent=4)))
	print("Decoded note: {}".format(base64.b64decode(
		confirmed_txn["txn"]["txn"]["note"]).decode()))
	print("Starting Account balance: {} microAlgos".format(account_info.get('amount')) )
	print("Amount transfered: {} microAlgos".format(amount) )    
	print("Fee: {} microAlgos".format(params.fee) ) 
	account_info = algod_client.account_info(my_address)
	print("Final Account balance: {} microAlgos".format(account_info.get('amount')) + "\n")


getting_started_example()
