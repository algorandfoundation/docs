import os
import base64
import json

from algosdk.v2client import algod
from algosdk.future.transaction import *
from algosdk import transaction, account, mnemonic
from algosdk.atomic_transaction_composer import *

# // This code is meant for learning purposes only
# // It should not be used in production

# Read a file
def load_resource(res):
    dir_path = os.path.dirname(os.path.realpath(__file__))
    path = os.path.join(dir_path, res)
    with open(path, "rb") as fin:
        data = fin.read()
    return data


try:
    # Create an algod client, using default sandbox parameters here
    algod_token = "a" * 64
    algod_address = "http://localhost:4001"
    algod_client = algod.AlgodClient(algod_token, algod_address)

    # Get account info
    addr1_mnemonic = "<Your 25-word-mnemonic>"
    private_key = mnemonic.to_private_key(addr1_mnemonic)
    addr1 = account.address_from_private_key(private_key)

    # Load in our program
    myprogram = "sample.teal"
    data = load_resource(myprogram)
    source = data.decode("utf-8")

    # Compile the program against the algod
    response = algod_client.compile(source)
    print(f"Response Result (base64 encoded): {response['result']}")
    print(f"Response Hash: {response['hash']}")

    # Decode the program to bytes and encode the argument as bytes
    program = base64.b64decode(response["result"])
    arg1 = (123).to_bytes(8, "big")

    # Create the lsig account, passing arg
    lsig = LogicSigAccount(program, args=[arg1])

    # Recover the account that is wanting to delegate signature
    # never use mnemonics in code, for demo purposes
    passphrase = "<Your 25-word-mnemonic>"
    sk = mnemonic.to_private_key(passphrase)
    addr = account.address_from_private_key(sk)
    print("Address of Sender/Delegator: " + addr)

    # Sign the logic signature with an account sk
    lsig.sign(sk)

    # Get suggested parameters
    params = algod_client.suggested_params()

    # replace with any other address or amount
    receiver = addr1
    amount = 10000

    # fund the contract account to test payment from the contract
    atc = AtomicTransactionComposer()
    signer = AccountTransactionSigner(private_key)
    ptxn = TransactionWithSigner(
        PaymentTxn(addr1, params, lsig.address(), 1000000), signer
    )
    atc.add_transaction(ptxn)
    result = atc.execute(algod_client, 2)
    for res in result.abi_results:
        print(res.return_value)

    # Create a transaction
    txn = PaymentTxn(addr, params, receiver, amount)

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
