import json
from algosdk.v2client import algod
from algosdk import account, encoding, mnemonic
from algosdk.future.transaction import Multisig, PaymentTxn, MultisigTransaction
import base64
import os
from algosdk.future.transaction import *

def connect_to_network():
    # Specify your node address and token. This must be updated.
    # algod_address = ""  # ADD ADDRESS
    # algod_token = ""  # ADD TOKEN
    algod_address = "http://localhost:4001"
    algod_token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
    algod_client = algod.AlgodClient(algod_token, algod_address)
    return algod_client


def write_multisig_unsigned_transaction_to_file():
    algod_client = connect_to_network()
    # Change these values with mnemonics
    # mnemonic1 = "PASTE phrase for account 1"
    # mnemonic2 = "PASTE phrase for account 2"
    # mnemonic3 = "PASTE phrase for account 3"

    mnemonic1 = "patrol target joy dial ethics flip usual fatigue bulb security prosper brand coast arch casino burger inch cricket scissors shoe evolve eternal calm absorb school"
    mnemonic2 = "genius inside turtle lock alone blame parent civil depend dinosaur tag fiction fun skill chief use damp daughter expose pioneer today weasel box about silly"
    mnemonic3 = "off canyon mystery cable pluck emotion manual legal journey grit lunch include friend social monkey approve lava steel school mango auto cactus huge ability basket"

    # For ease of reference, add account public and private keys to
    # an accounts dict.

    private_key_1 = mnemonic.to_private_key(mnemonic1)
    account_1 = mnemonic.to_public_key(mnemonic1)

    private_key_2 = mnemonic.to_private_key(mnemonic2)
    account_2 = mnemonic.to_public_key(mnemonic2)

    private_key_3 = mnemonic.to_private_key(mnemonic3)
    account_3 = mnemonic.to_public_key(mnemonic3)

    print("Account 1 address: {}".format(account_1))
    print("Account 2 address: {}".format(account_2))
    print("Account 3 address: {}".format(account_3))

    # create a multisig account
    version = 1  # multisig version
    threshold = 2  # how many signatures are necessary
    msig = Multisig(version, threshold, [account_1, account_2])

    print("Multisig Address: ", msig.address())
    print("Please go to: https://bank.testnet.algorand.network/ to fund multisig account.", msig.address())
    # input("Please go to: https://bank.testnet.algorand.network/ to fund multisig account." + '\n' + "Press Enter to continue...")

    # get suggested parameters
    params = algod_client.suggested_params()
    # comment out the next two (2) lines to use suggested fees
    # params.flat_fee = True
    # params.fee = 1000

    # create a transaction
    sender = msig.address()
    recipient = account_3
    amount = 10000
    note = "Hello Multisig".encode()
    txn = PaymentTxn(sender, params, recipient, amount, None, note, None)

    # create a SignedTransaction object
    mtx = MultisigTransaction(txn, msig)

        # write to file
    dir_path = os.path.dirname(os.path.realpath(__file__))
    transaction.write_to_file(
        [mtx], dir_path + "/unsigned.mtx")

def read_multisig_unsigned_transaction_from_file():
    algod_client = connect_to_network()
    # Change these values with mnemonics
    # mnemonic1 = "PASTE phrase for account 1"
    # mnemonic2 = "PASTE phrase for account 2"
    # mnemonic3 = "PASTE phrase for account 3"

    mnemonic1 = "patrol target joy dial ethics flip usual fatigue bulb security prosper brand coast arch casino burger inch cricket scissors shoe evolve eternal calm absorb school"
    mnemonic2 = "genius inside turtle lock alone blame parent civil depend dinosaur tag fiction fun skill chief use damp daughter expose pioneer today weasel box about silly"
    mnemonic3 = "off canyon mystery cable pluck emotion manual legal journey grit lunch include friend social monkey approve lava steel school mango auto cactus huge ability basket"

    # For ease of reference, add account public and private keys to
    # an accounts dict.

    private_key_1 = mnemonic.to_private_key(mnemonic1)
    account_1 = mnemonic.to_public_key(mnemonic1)

    private_key_2 = mnemonic.to_private_key(mnemonic2)
    account_2 = mnemonic.to_public_key(mnemonic2)

    private_key_3 = mnemonic.to_private_key(mnemonic3)
    account_3 = mnemonic.to_public_key(mnemonic3)

    print("Account 1 address: {}".format(account_1))
    print("Account 2 address: {}".format(account_2))
    print("Account 3 address: {}".format(account_3))

# read from file
    dir_path = os.path.dirname(os.path.realpath(__file__))
    msigs = transaction.retrieve_from_file(dir_path + "/unsigned.mtx")
    mtx = msigs[0]


    # sign the transaction
    mtx.sign(private_key_1)
    mtx.sign(private_key_2)

    # send the transaction
    txid = algod_client.send_raw_transaction(
        encoding.msgpack_encode(mtx))
        # wait for confirmation	
    try:
        confirmed_txn = wait_for_confirmation(algod_client, txid, 4) 
        print("TXID: ", txid)
        print("Result confirmed in round: {}".format(confirmed_txn['confirmed-round'])) 
        print("Transaction information: {}".format(
            json.dumps(confirmed_txn, indent=4)))
        print("Decoded note: {}".format(base64.b64decode(
            confirmed_txn["txn"]["txn"]["note"]).decode()))
    except Exception as err:
        print(err)


def write_multisig_signed_transaction_to_file():
    algod_client = connect_to_network()
    # Change these values with mnemonics
    # Shown for demonstration purposes. NEVER reveal secret mnemonics in practice.
    mnemonic1 = "PASTE phrase for account 1"
    mnemonic2 = "PASTE phrase for account 2"
    mnemonic3 = "PASTE phrase for account 3"

 
    # For ease of reference, add account public and private keys to
    # an accounts dict.

    private_key_1 = mnemonic.to_private_key(mnemonic1)
    account_1 = mnemonic.to_public_key(mnemonic1)

    private_key_2 = mnemonic.to_private_key(mnemonic2)
    account_2 = mnemonic.to_public_key(mnemonic2)

    private_key_3 = mnemonic.to_private_key(mnemonic3)
    account_3 = mnemonic.to_public_key(mnemonic3)
    # create a multisig account
    version = 1  # multisig version
    threshold = 2  # how many signatures are necessary
    msig = Multisig(version, threshold, [account_1, account_2])

    print("Multisig Address: ", msig.address())
    print("Please go to: https://dispenser.testnet.aws.algodev.network/ to fund multisig account.", msig.address())
    # input("Please go to: https://dispenser.testnet.aws.algodev.network/ to fund multisig account." + '\n' + "Press Enter to continue...")

    # get suggested parameters
    params = algod_client.suggested_params()
    # comment out the next two (2) lines to use suggested fees
    # params.flat_fee = True
    # params.fee = 1000

    # create a transaction
    sender = msig.address()
    recipient = account_3
    amount = 10000
    note = "Hello Multisig".encode()
    txn = PaymentTxn(sender, params, recipient, amount, None, note, None)

    # create a SignedTransaction object
    mtx = MultisigTransaction(txn, msig)

    # sign the transaction
    mtx.sign(private_key_1)
    mtx.sign(private_key_2)

   # print encoded transaction
    print(encoding.msgpack_encode(mtx))

    # write to file
    dir_path = os.path.dirname(os.path.realpath(__file__))
    transaction.write_to_file(
        [mtx], dir_path + "/signed.mtx")
    print("Signed mtx file saved!")

def read_multisig_signed_transaction_from_file():
    algod_client = connect_to_network()
    
    # read from file
    dir_path = os.path.dirname(os.path.realpath(__file__))
    msigs = transaction.retrieve_from_file(dir_path + "/signed.mtx")
    mtx = msigs[0]

    try:
    # send the transaction
        txid = algod_client.send_raw_transaction(
        encoding.msgpack_encode(mtx))        
         # wait for confirmation	       
        confirmed_txn = wait_for_confirmation(algod_client, txid, 4) 
        print("TXID: ", txid) 
        print("Result confirmed in round: {}".format(confirmed_txn['confirmed-round']))
        print("Transaction information: {}".format(
            json.dumps(confirmed_txn, indent=4)))
        print("Decoded note: {}".format(base64.b64decode(
            confirmed_txn["txn"]["txn"]["note"]).decode()))
    except Exception as err:
        print(err)



def test_signed():
    write_multisig_signed_transaction_to_file()
    read_multisig_signed_transaction_from_file()


def test_unsigned():
    write_multisig_unsigned_transaction_to_file()
    read_multisig_unsigned_transaction_from_file()
    

# test_signed()
test_unsigned()