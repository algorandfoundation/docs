from algosdk import transaction, account, mnemonic
from algosdk.v2client import algod
from algosdk.future.transaction import PaymentTxn, LogicSig
import os
import base64

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
    print("Transaction {} confirmed in round {}.".format(
        txid, txinfo.get('confirmed-round')))
    return txinfo

# Read a file
def load_resource(res):
    dir_path = os.path.dirname(os.path.realpath(__file__))
    path = os.path.join(dir_path, res)
    with open(path, "rb") as fin:
        data = fin.read()
    return data

try:

    # Create an algod client
    # algod_token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" 
    # algod_address = "http://localhost:4001" 
 

    algod_token = "6b3a2ae3896f23be0a1f0cdd083b6d6d046fbeb594a3ce31f2963b717f74ad43"
    algod_address = "http://127.0.0.1:54746"

    # algod_token = "algod-token" < PLACEHOLDER >
    # algod_address = "algod-address" < PLACEHOLDER >
    # receiver = "receiver-address" < PLACEHOLDER >
    receiver = "ATTR6RUEHHBHXKUHT4GUOYWNBVDV2GJ5FHUWCSFZLHD55EVKZWOWSM7ABQ"
    algod_client = algod.AlgodClient(algod_token, algod_address)
   
    myprogram = "samplearg.teal"
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
    # arg_str = "my string"
    # arg1 = arg_str.encode()
    # lsig = transaction.LogicSig(program, args=[arg1])

    # integer parameter
    # arg1 = (123).to_bytes(8, 'big')
    # lsig = transaction.LogicSig(program, args=[arg1])

    # program = b"hex-encoded-program"
    # for example, b"\x01\x20\x01\x00\x22 is `int 0`
    # see more info here: https://developer.algorand.org/docs/features/asc1/sdks/#accessing-teal-program-from-sdks

    # Create arg to pass if TEAL program requires an arg,
    # if not, omit args param
    arg1 = (123).to_bytes(8, 'big')
    lsig = LogicSig(program, args=[arg1])
    sender = lsig.address()

    # Get suggested parameters
    params = algod_client.suggested_params()
    # Comment out the next two (2) lines to use suggested fees
    params.flat_fee = True
    params.fee = 1000
    
    # Build transaction  
    amount = 10000 
    closeremainderto = None

    # Create a transaction
    txn = PaymentTxn(
        sender, params, receiver, amount, closeremainderto)

    # Create the LogicSigTransaction with contract account LogicSig
    lstx = transaction.LogicSigTransaction(txn, lsig)
    # transaction.write_to_file([lstx], "simple.stxn")

    # Send raw LogicSigTransaction to network
    txid = algod_client.send_transaction(lstx)
    print("Transaction ID: " + txid)    
    wait_for_confirmation(algod_client, txid) 

except Exception as e:
    print(e)
