import json
import time
import base64
from algosdk import mnemonic
from algosdk.v2client import algod
from algosdk.future.transaction import PaymentTxn

def wait_for_confirmation(client, txid):
	"""
	Utility function to wait until the transaction is
	confirmed before proceeding.
	"""
	last_round = client.status().get('last-round')
	txinfo = client.pending_transaction_info(txid)
	while not (txinfo.get('confirmed-round') and txinfo.get('confirmed-round') > 0):
		print("Waiting for confirmation")
		last_round += 1
		client.status_after_block(last_round)
		txinfo = client.pending_transaction_info(txid)
	print("Transaction {} confirmed in round {}.".format(txid, txinfo.get('confirmed-round')))
	return txinfo


def getting_started_example():
	algod_address = "http://localhost:4001"
	algod_token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	algod_client = algod.AlgodClient(algod_token, algod_address)

	passphrase = "Your 25-word mnemonic generated and displayed above"

	# generate a public/private key pair
	private_key = mnemonic.to_private_key(passphrase)
	my_address = mnemonic.to_public_key(passphrase)
	print("My address: {}".format(my_address))

	account_info = algod_client.account_info(my_address)
	print("Account balance: {} microAlgos".format(account_info.get('amount')))

	# build transaction
	params = algod_client.suggested_params()
	# comment out the next two (2) lines to use suggested fees
	params.flat_fee = True
	params.fee = 1000
	receiver = "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A"
	note = "Hello World".encode()

	unsigned_txn = PaymentTxn(my_address, params, receiver, 1000000, None, note)

	# sign transaction
	signed_txn = unsigned_txn.sign(mnemonic.to_private_key(passphrase))
	txid = algod_client.send_transaction(signed_txn)
	print("Signed transaction with txID: {}".format(txid))

	# wait for confirmation
	wait_for_confirmation(algod_client, txid) 

	# read transction
	try:
		confirmed_txn = algod_client.pending_transaction_info(txid)
	except Exception as err:
		print(err)
	print("Transaction information: {}".format(json.dumps(confirmed_txn, indent=4)))
	print("Decoded note: {}".format(base64.b64decode(confirmed_txn["txn"]["txn"]["note"]).decode()))

getting_started_example()