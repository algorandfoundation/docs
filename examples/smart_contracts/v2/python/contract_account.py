import os
import base64
import json

from algosdk.v2client import algod
from algosdk.future.transaction import *
from algosdk import transaction, account, mnemonic
from algosdk.atomic_transaction_composer import *
from algosdk.constants import *
from algosdk.util import *

# Note this method is defined in sandbox.py
from sandbox import get_accounts

# Read a file
def load_resource(res):
    dir_path = os.path.dirname(os.path.realpath(__file__))
    path = os.path.join(dir_path, res)
    with open(path, "r") as f:
        data = f.read()
    return data


def contract_account_example():
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
    sender = lsig.address()
    print(f"Address of lsig: {sender}")

    # Get suggested parameters
    params = algod_client.suggested_params()

    # replace with any other address or amount
    receiver = receiver_addr
    amount = 10000

    # fund the contract account to test payment from the contract
    atc = AtomicTransactionComposer()
    signer = AccountTransactionSigner(receiver_sk)
    ptxn = TransactionWithSigner(
        PaymentTxn(
            receiver_addr,
            params,
            lsig.address(),
            algos_to_microalgos(0.1)
            + MIN_TXN_FEE
            + amount,  # Amount = min balance + min txn fees + payment amount
        ),
        signer,
    )
    atc.add_transaction(ptxn)
    result = atc.execute(algod_client, 2)
    for res in result.abi_results:
        print(res.return_value)

    # Create a transaction
    txn = PaymentTxn(sender, params, receiver, amount)

    # Create the LogicSigTransaction with contract account LogicSigAccount
    lstx = transaction.LogicSigTransaction(txn, lsig.lsig)

    # Send raw LogicSigTransaction to network
    txid = algod_client.send_transaction(lstx)
    print("Transaction ID: " + txid)

    # wait for confirmation
    confirmed_txn = wait_for_confirmation(algod_client, txid, 4)
    print(f"Result confirmed in round: {confirmed_txn['confirmed-round']}")
    print(f"Transaction information: {json.dumps(confirmed_txn, indent=4)}")


contract_account_example()
