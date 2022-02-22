import json
import time
import base64
import os
from algosdk import mnemonic
from algosdk.v2client import algod
from algosdk.future.transaction import PaymentTxn
from algosdk.future.transaction import *


def connect_to_network():
    algod_address = "http://localhost:4001"
    algod_token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
    algod_client = algod.AlgodClient(algod_token, algod_address)
    return algod_client
    



def write_signed_transaction_to_file():
    algod_client = connect_to_network()
    # Shown for demonstration purposes. NEVER reveal secret mnemonics in practice.
    passphrase = "PASTE phrase for account"
    # generate a public/private key pair
    my_address = mnemonic.to_public_key(passphrase)
    print("My address: {}".format(my_address))

    account_info = algod_client.account_info(my_address)
    print("Account balance: {} microAlgos".format(account_info.get('amount')))

    # build transaction
    params = algod_client.suggested_params()
    # comment out the next two (2) lines to use suggested fees
    # params.flat_fee = True
    # params.fee = 1000
    receiver = "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A"
    note = "Hello World".encode()
    unsigned_txn = PaymentTxn(
        my_address, params, receiver, 100000, None, note)
    # sign transaction
    signed_txn = unsigned_txn.sign(mnemonic.to_private_key(passphrase))
    # write to file
    dir_path = os.path.dirname(os.path.realpath(__file__))
    transaction.write_to_file([signed_txn], dir_path + "/signed.txn")


def read_signed_transaction_from_file():
    algod_client = connect_to_network()
    # read signed transaction from file
    dir_path = os.path.dirname(os.path.realpath(__file__))
    txns = transaction.retrieve_from_file(dir_path + "/signed.txn")
    signed_txn = txns[0]
    try:
        txid = algod_client.send_transaction(signed_txn)
        print("Signed transaction with txID: {}".format(txid))
    # wait for confirmation
        confirmed_txn = wait_for_confirmation(algod_client, txid, 4)
        print("TXID: ", txid)
        print("Result confirmed in round: {}".format(confirmed_txn['confirmed-round']))

    except Exception as err:
        print(err)
        return

    print("Transaction information: {}".format(
        json.dumps(confirmed_txn, indent=4)))
    print("Decoded note: {}".format(base64.b64decode(
        confirmed_txn["txn"]["txn"]["note"]).decode()))


def write_unsigned_transaction_to_file():
    algod_client = connect_to_network()
    passphrase = "price clap dilemma swim genius fame lucky crack torch hunt maid palace ladder unlock symptom rubber scale load acoustic drop oval cabbage review abstract embark"
    # generate a public/private key pair
    my_address = mnemonic.to_public_key(passphrase)
    print("My address: {}".format(my_address))
    account_info = algod_client.account_info(my_address)
    print("Account balance: {} microAlgos".format(account_info.get('amount')))
    # build transaction
    params = algod_client.suggested_params()
    # comment out the next two (2) lines to use suggested fees
    # params.flat_fee = True
    # params.fee = 1000
    receiver = "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A"
    note = "Hello World".encode()
    unsigned_txn = PaymentTxn(
        my_address, params, receiver, 1000000, None, note)
    # write to file
    dir_path = os.path.dirname(os.path.realpath(__file__))
    transaction.write_to_file([unsigned_txn], dir_path + "/unsigned.txn")


def read_unsigned_transaction_from_file():
    algod_client = connect_to_network()

    passphrase = "price clap dilemma swim genius fame lucky crack torch hunt maid palace ladder unlock symptom rubber scale load acoustic drop oval cabbage review abstract embark"

    my_address = mnemonic.to_public_key(passphrase)
    print("My address: {}".format(my_address))
    account_info = algod_client.account_info(my_address)
    print("Account balance: {} microAlgos".format(account_info.get('amount')))

    dir_path = os.path.dirname(os.path.realpath(__file__))
    txns = transaction.retrieve_from_file(dir_path + "/unsigned.txn")
    # sign and submit transaction
    txn = txns[0]
    signed_txn = txn.sign(mnemonic.to_private_key(passphrase))
    txid = signed_txn.transaction.get_txid()
    print("Signed transaction with txID: {}".format(txid))	

    try:
        txid = algod_client.send_transaction(signed_txn)
        # wait for confirmation              
        confirmed_txn = wait_for_confirmation(algod_client, txid, 4)
        print("TXID: ", txid)
        print("Result confirmed in round: {}".format(confirmed_txn['confirmed-round']))

    except Exception as err:
        print(err)
        return

    print("Transaction information: {}".format(
        json.dumps(confirmed_txn, indent=4)))
    print("Decoded note: {}".format(base64.b64decode(
        confirmed_txn["txn"]["txn"]["note"]).decode()))


def test_signed():
    write_signed_transaction_to_file()
    read_signed_transaction_from_file()


def test_unsigned():
    write_unsigned_transaction_to_file()
    read_unsigned_transaction_from_file()


test_signed()
# test_unsigned()



