import os
import base64
import json

from algosdk.v2client import algod
from algosdk.future.transaction import *
from algosdk import transaction
from algosdk.atomic_transaction_composer import *
from algosdk.constants import *
from algosdk.util import *

# Note this method is defined in sandbox.py
from sandbox import get_accounts

# This code is meant for learning purposes only
# It should not be used in production

# Read a file
def load_resource(res):
    dir_path = os.path.dirname(os.path.realpath(__file__))
    path = os.path.join(dir_path, res)
    with open(path, "r") as f:
        data = f.read()
    return data


try:
    # Create an algod client, using default sandbox parameters here
    algod_token = "a" * 64
    algod_address = "http://localhost:4001"
    algod_client = algod.AlgodClient(algod_token, algod_address)

    # Get receiver account info (address and private key)
    receiver_addr, receiver_sk = get_accounts()[0]

    # Load in our program
    myprogram = "samplearg.teal"
    data = load_resource(myprogram)

    # Compile the program against the algod
    response = algod_client.compile(data)
    print(f"Response Result (base64 encoded): {response['result']}")
    print(f"Response Hash: {response['hash']}")

    # Decode the program to bytes and encode the argument as bytes
    program = base64.b64decode(response["result"])
    arg1 = (123).to_bytes(8, "big")

    # Create the lsig account, passing arg
    lsig = LogicSigAccount(program, args=[arg1])

    # Recover the account that is wanting to delegate signature
    # never use mnemonics in code, for demo purposes
    delegator_addr, delegator_sk = get_accounts()[1]
    print("Address of Sender/Delegator: " + delegator_addr)

    # Sign the logic signature with an account sk
    lsig.sign(delegator_sk)

    # Get suggested parameters
    params = algod_client.suggested_params()

    # replace with any other address or amount
    receiver = receiver_addr
    amount = 10000

    # Create a transaction
    txn = PaymentTxn(delegator_addr, params, receiver, amount)

    # Create the LogicSigTransaction with contract account LogicSigAccount
    lstx = transaction.LogicSigTransaction(txn, lsig.lsig)

    # Send raw LogicSigTransaction to network
    txid = algod_client.send_transaction(lstx)
    print("Transaction ID: " + txid)

    # Wait for confirmation
    confirmed_txn = wait_for_confirmation(algod_client, txid, 4)
    print(f"Result confirmed in round: {confirmed_txn['confirmed-round']}")
    print(f"Transaction information: {json.dumps(confirmed_txn, indent=4)}")
except Exception as e:
    print(e)
