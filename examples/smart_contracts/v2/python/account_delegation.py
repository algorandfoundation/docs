from algosdk import algod, transaction, account, mnemonic
from algosdk.v2client import algod

import os
import base64
from algosdk.future.transaction import *

# Read a file
def load_resource(res):
    dir_path = os.path.dirname(os.path.realpath(__file__))
    path = os.path.join(dir_path, res)
    with open(path, "rb") as fin:
        data = fin.read()
    return data
try:

    # Create an algod client
    algod_token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
    algod_address = "http://localhost:4001"

    # algod_token = "<algod-token>"
    # algod_address = "<algod-address:port>"
    # receiver = "<receiver-address>"
    receiver = "NQMDAY2QKOZ4ZKJLE6HEO6LTGRJHP3WQVZ5C2M4HKQQLFHV5BU5AW4NVRY"
    algod_client = algod.AlgodClient(algod_token, algod_address)

    myprogram = "samplearg.teal"
    # myprogram = "<filename>"
    # Read TEAL program
    data = load_resource(myprogram)
    source = data.decode('utf-8')
    # Compile TEAL program
    # // This code is meant for learning purposes only
    # // It should not be used in production
    # // sample.teal

    # arg_0
    # btoi
    # int 123
    # ==

    # // bto1
    # // Opcode: 0x17
    # // Pops: ... stack, []byte
    # // Pushes: uint64
    # // converts bytes X as big endian to uint64
    # // btoi panics if the input is longer than 8 bytes

    response = algod_client.compile(source)
    # Print(response)
    print("Response Result = ", response['result'])
    print("Response Hash = ", response['hash'])

    # Create logic sig
    programstr = response['result']
    t = programstr.encode("ascii")
    # program = b"hex-encoded-program"
    program = base64.decodebytes(t)
    print(program)
    print(len(program) * 8)
    # Create arg to pass
    # string parameter
    # arg_str = "<my string>"
    # arg1 = arg_str.encode()
    # lsig = transaction.LogicSig(program, args=[arg1])

    # integer parameter
    # arg1 = (123).to_bytes(8, 'big')
    # lsig = transaction.LogicSig(program, args=[arg1])
    # see more info here: https://developer.algorand.org/docs/features/asc1/sdks/#accessing-teal-program-from-sdks

    # if TEAL program requires an arg,
    # if not, omit args param on LogicSig
    # lsig = LogicSig(program)
    arg1 = (123).to_bytes(8, 'big')
    lsig = LogicSig(program, args=[arg1])

    # Recover the account that is wanting to delegate signature
    # never use mnemonics in code, for demo purposes    
    passphrase = "<25-word-mnemonic>"
    sk = mnemonic.to_private_key(passphrase)
    addr = account.address_from_private_key(sk)
    print("Address of Sender/Delegator: " + addr)

    # Sign the logic signature with an account sk
    lsig.sign(sk)

    # Get suggested parameters
    params = algod_client.suggested_params()
    # Comment out the next two (2) lines to use suggested fees
    # params.flat_fee = True
    # params.fee = 1000

    # Build transaction
    amount = 10000
    closeremainderto = None

    # Create a transaction
    txn = PaymentTxn(
        addr, params, receiver, amount, closeremainderto)
    # Create the LogicSigTransaction with contract account LogicSig
    lstx = transaction.LogicSigTransaction(txn, lsig)
    txns = [lstx]
    transaction.write_to_file(txns, "simple.stxn")
    # Send raw LogicSigTransaction to network
    txid = algod_client.send_transaction(lstx)
    print("Transaction ID: " + txid)

    confirmed_txn = wait_for_confirmation(algod_client, txid, 4)
    print("TXID: ", txid)
    print("Result confirmed in round: {}".format(confirmed_txn['confirmed-round']))    
except Exception as e:
    print(e)
