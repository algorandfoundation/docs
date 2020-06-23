import json
import time
import base64
import os
from algosdk.v2client import algod
from algosdk import mnemonic
from algosdk.future import transaction
from algosdk import encoding
from algosdk import account

# This atomic transfer example code requires three (3) acounts:
#  - account_1 requires a user-defined mnemonic and be funded with 1001000 microAlgos
#  - account_2 requires a user-defined mnemonic and be funded with 2001000 microAlgos
#  - account_3 auto-generated within the code, 1000000 microAlgos will be transfered here
# For account_1 and account_2, replcace the string "Your 25-word mnemonic goes here" in the code below.
# For account_3, ensure you note the mnemonic generated for future.
# Faucents available for funding accounts:
#  - TestNet: https://developer.algorand.org/docs/reference/algorand-networks/testnet/#faucet
#  - BetaNet: https://developer.algorand.org/docs/reference/algorand-networks/betanet/#faucet
# Replace the algod_address and algod_token parameters below to connect to your API host.

# user declared account mnemonics for account1 and account2
mnemonic_1 = "Your 25-word mnemonic goes here"
mnemonic_2 = "Your 25-word mnemonic goes here"

# user declared algod connection parameters
algod_address = "http://localhost:4001"
algod_token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

# Function that waits for a given txId to be confirmed by the network
def wait_for_confirmation(client, txid):
    last_round = client.status().get('last-round')
    txinfo = client.pending_transaction_info(txid)
    while not (txinfo.get('confirmed-round') and txinfo.get('confirmed-round') > 0):
        print("Waiting for confirmation...")
        last_round += 1
        client.status_after_block(last_round)
        txinfo = client.pending_transaction_info(txid)
    print("Transaction {} confirmed in round {}.".format(txid, txinfo.get('confirmed-round')))
    return txinfo

# utility function to get address string
def get_address(mn) :
    pk_account_a = mnemonic.to_private_key(mn)
    address = account.address_from_private_key(pk_account_a)
    print("Address :", address)
    return address

# utility function to generate new account
def generate_new_account() :
	private_key, address = account.generate_account()
	print("Created new account: ", address)
	print("Generated mnemonic: \"{}\"".format(mnemonic.from_private_key(private_key)))
	return address    

# utility function to display account balance
def display_account_algo_balance(client, address) :
    account_info = client.account_info(address)
    print("{}: {} microAlgos".format(address, account_info["amount"]))

def group_transactions() :
	# Initialize an algodClient
    algod_client = algod.AlgodClient(algod_token, algod_address)

	# declared account1 and account2 based on user supplied mnemonics
    print("Loading two existing accounts...")
    account_1 = get_address(mnemonic_1)
    account_2 = get_address(mnemonic_2)

	# convert mnemonic1 and mnemonic2 using the mnemonic.ToPrivateKey() helper function
    sk_1 = mnemonic.to_private_key(mnemonic_1)
    sk_2 = mnemonic.to_private_key(mnemonic_2)

	# generate account3, display mnemonic, wait
	# print("Generating new account...")
    account_3 = generate_new_account()
    print("!! NOTICE !! Please retain the above generated \"25-word mnemonic passphrase\" for future use.")

	# display account balances
    print("Initial balances:")
    display_account_algo_balance(algod_client, account_1)
    display_account_algo_balance(algod_client, account_2)
    display_account_algo_balance(algod_client, account_3)

	# get node suggested parameters
    params = algod_client.suggested_params()
    # comment out the next two (2) lines to use suggested fees
    params.flat_fee = True
    params.fee = 1000

	# create transactions
    print("Creating transactions...")
	# from account 1 to account 3
    sender = account_1
    receiver = account_3
    amount = 1000000
    txn_1 = transaction.PaymentTxn(sender, params, receiver, amount)
    print("...txn_1: from {} to {} for {} microAlgos".format(sender, receiver, amount))
    print("...created txn_1: ", txn_1.get_txid())

	# from account 2 to account 1
    sender = account_2
    receiver = account_1
    amount = 2000000
    txn_2 = transaction.PaymentTxn(sender, params, receiver, amount)
    print("...txn_2: from {} to {} for {} microAlgos".format(sender, receiver, amount))
    print("...created txn_2: ", txn_2.get_txid())

	# combine transations
    print("Combining transactions...")
	# the SDK does this implicitly within grouping below
     
    print("Grouping transactions...")
	# compute group id and put it into each transaction
    group_id = transaction.calculate_group_id([txn_1, txn_2])
    print("...computed groupId: ", group_id)
    txn_1.group = group_id
    txn_2.group = group_id

	# split transaction group
    print("Splitting unsigned transaction group...")
    # this example does not use files on disk, so splitting is implicit above

	# sign transactions
    print("Signing transactions...")
    stxn_1 = txn_1.sign(sk_1)    
    print("...account1 signed txn_1: ", stxn_1.get_txid())
    stxn_2 = txn_2.sign(sk_2)
    print("...account2 signed txn_2: ", stxn_2.get_txid())

	# assemble transaction group
    print("Assembling transaction group...")
    signedGroup =  []
    signedGroup.append(stxn_1)
    signedGroup.append(stxn_2)

	# send transactions
    print("Sending transaction group...")
    tx_id = algod_client.send_transactions(signedGroup)

    # wait for confirmation
    wait_for_confirmation(algod_client, tx_id) 

	# display account balances
    print("Final balances:")
    display_account_algo_balance(algod_client, account_1)
    display_account_algo_balance(algod_client, account_2)
    display_account_algo_balance(algod_client, account_3)

	# display confirmed transaction group
	# tx1
    confirmed_txn = algod_client.pending_transaction_info(txn_1.get_txid())
    print("Transaction information: {}".format(json.dumps(confirmed_txn, indent=4)))

	# tx2
    confirmed_txn = algod_client.pending_transaction_info(txn_2.get_txid())
    print("Transaction information: {}".format(json.dumps(confirmed_txn, indent=4)))
group_transactions()