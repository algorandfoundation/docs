import json
import time
import base64
from algosdk import mnemonic
from algosdk.v2client import algod
from algosdk.future.transaction import PaymentTxn
from algosdk.future.transaction import wait_for_confirmation

def getting_started_example():
	algod_address = "http://localhost:4001"
	algod_token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	algod_client = algod.AlgodClient(algod_token, algod_address)

	passphrase = "price clap dilemma swim genius fame lucky crack torch hunt maid palace ladder unlock symptom rubber scale load acoustic drop oval cabbage review abstract embark"

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
	try:
		confirmed_txn = wait_for_confirmation(algod_client, txid, 4)  
	except Exception as err:
		print(err)
		return

	print("Transaction information: {}".format(
		json.dumps(confirmed_txn, indent=4)))
	print("Decoded note: {}".format(base64.b64decode(
		confirmed_txn["txn"]["txn"]["note"]).decode()))


getting_started_example()
