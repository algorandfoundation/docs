from algosdk import transaction, account, mnemonic
from algosdk.v2client import algod
from algosdk.future.transaction import *
import os
import base64
import json

# Read a file
def load_resource(res):
    dir_path = os.path.dirname(os.path.realpath(__file__))
    path = os.path.join(dir_path, res)
    with open(path, "rb") as fin:
        data = fin.read()
    return data

def contract_account_example():

    # Create an algod client
    algod_token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" 
    algod_address = "http://localhost:4001" 

    # algod_token = "<algod-token>" 
    # algod_address = "<algod-address>" 
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

    # string parameter
    # arg_str = "<my string>"
    # arg1 = arg_str.encode()
    # lsig = transaction.LogicSigAccount(program, args=[arg1])

    # see more info here: https://developer.algorand.org/docs/features/asc1/sdks/#accessing-teal-program-from-sdks

    # Create arg to pass if TEAL program requires an arg
    # if not, omit args param
    arg1 = (123).to_bytes(8, 'big')
    lsig = LogicSigAccount(program, args=[arg1])
    sender = lsig.address()

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
        sender, params, receiver, amount, closeremainderto)

    # Create the LogicSigTransaction with contract account LogicSigAccount
    lstx = transaction.LogicSigTransaction(txn, lsig)
    # transaction.write_to_file([lstx], "simple.stxn")

    # Send raw LogicSigTransaction to network
    txid = algod_client.send_transaction(lstx)
    print("Transaction ID: " + txid) 
    # wait for confirmation	
    try:
        confirmed_txn = wait_for_confirmation(algod_client, txid, 4)  
        print("TXID: ", txid)
        print("Result confirmed in round: {}".format(confirmed_txn['confirmed-round']))

    except Exception as err:
        print(err)

    print("Transaction information: {}".format(
    json.dumps(confirmed_txn, indent=4)))
 
contract_account_example()
