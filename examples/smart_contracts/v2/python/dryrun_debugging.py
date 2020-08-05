from algosdk import algod, transaction, account, mnemonic
from algosdk.v2client import algod
from algosdk.future.transaction import PaymentTxn, LogicSig
import os
import base64
from algosdk.v2client.models import DryrunRequest, DryrunSource
# from algosdk.testing import dryrun
import json

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

# dryrun source if provided, else dryrun compiled
def dryrun_debug(lstx, mysource):
    sources = []
    if (mysource != None):
        # source
        sources = [DryrunSource(field_name="lsig", source=mysource, txn_index=0)]
    drr = DryrunRequest(txns=[lstx], sources=sources)
    dryrun_response = algod_client.dryrun(drr)
    return dryrun_response


try:

    # Create an algod client
    algod_token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
    algod_address = "http://localhost:4001"

    # algod_token = "<algod-token>"
    # algod_address = "<algod-address:port>"
    # receiver = "<receiver-address>"
    receiver = "ATTR6RUEHHBHXKUHT4GUOYWNBVDV2GJ5FHUWCSFZLHD55EVKZWOWSM7ABQ"
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
    passphrase = "canal enact luggage spring similar zoo couple stomach shoe laptop middle wonder eager monitor weather number heavy skirt siren purity spell maze warfare ability ten"
    # passphrase = "<25-word-mnemonic>"
    sk = mnemonic.to_private_key(passphrase)
    addr = account.address_from_private_key(sk)
    print("Address of Sender/Delegator: " + addr)

    # Sign the logic signature with an account sk
    lsig.sign(sk)

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
        addr, params, receiver, amount, closeremainderto)
    # Create the LogicSigTransaction with contract account LogicSig
    lstx = transaction.LogicSigTransaction(txn, lsig)
    # transaction.write_to_file([lstx], "./simple.stxn")

    # compile
    dryrun_response_compiled = dryrun_debug(lstx, None)
    print ("COMPILED Dryrun results...")
    print(json.dumps(dryrun_response_compiled, indent=2))
   
    # source   
    dryrun_respone_source = dryrun_debug(lstx, source)
    print("SOURCE Dryrun results...")
    print(json.dumps(dryrun_respone_source, indent=2))
   
    # Send raw LogicSigTransaction to network

    txid = algod_client.send_transaction(lstx)
    print("Transaction ID: " + txid)
    wait_for_confirmation(algod_client, txid)
except Exception as e:
    print(e)

# output should look similar to this
# COMPILED Dryrun results...
# {
#     "error": "",
#     "protocol-version": "https://github.com/algorandfoundation/specs/tree/e5f565421d720c6f75cdd186f7098495caf9101f",
#     "txns": [
#         {
#             "disassembly": [
#                 "// version 1",
#                 "intcblock 123",
#                 "arg_0",
#                 "btoi",
#                 "intc_0",
#                 "==",
#                 ""
#             ],
#             "logic-sig-messages": [
#                 "PASS"
#             ],
#             "logic-sig-trace": [
#                 {
#                     "line": 1,
#                     "pc": 1,
#                     "stack": []
#                 },
#                 {
#                     "line": 2,
#                     "pc": 4,
#                     "stack": []
#                 },
#                 {
#                     "line": 3,
#                     "pc": 5,
#                     "stack": [
#                         {
#                             "bytes": "AAAAAAAAAHs=",
#                             "type": 1,
#                             "uint": 0
#                         }
#                     ]
#                 },
#                 {
#                     "line": 4,
#                     "pc": 6,
#                     "stack": [
#                         {
#                             "bytes": "",
#                             "type": 2,
#                             "uint": 123
#                         }
#                     ]
#                 },
#                 {
#                     "line": 5,
#                     "pc": 7,
#                     "stack": [
#                         {
#                             "bytes": "",
#                             "type": 2,
#                             "uint": 123
#                         },
#                         {
#                             "bytes": "",
#                             "type": 2,
#                             "uint": 123
#                         }
#                     ]
#                 },
#                 {
#                     "line": 6,
#                     "pc": 8,
#                     "stack": [
#                         {
#                             "bytes": "",
#                             "type": 2,
#                             "uint": 1
#                         }
#                     ]
#                 }
#             ]
#         }
#     ]
# }
# SOURCE Dryrun results...
# {
#     "error": "",
#     "protocol-version": "https://github.com/algorandfoundation/specs/tree/e5f565421d720c6f75cdd186f7098495caf9101f",
#     "txns": [
#         {
#             "disassembly": [
#                 "// version 1",
#                 "intcblock 123",
#                 "arg_0",
#                 "btoi",
#                 "intc_0",
#                 "==",
#                 ""
#             ],
#             "logic-sig-messages": [
#                 "PASS"
#             ],
#             "logic-sig-trace": [
#                 {
#                     "line": 1,
#                     "pc": 1,
#                     "stack": []
#                 },
#                 {
#                     "line": 2,
#                     "pc": 4,
#                     "stack": []
#                 },
#                 {
#                     "line": 3,
#                     "pc": 5,
#                     "stack": [
#                         {
#                             "bytes": "AAAAAAAAAHs=",
#                             "type": 1,
#                             "uint": 0
#                         }
#                     ]
#                 },
#                 {
#                     "line": 4,
#                     "pc": 6,
#                     "stack": [
#                         {
#                             "bytes": "",
#                             "type": 2,
#                             "uint": 123
#                         }
#                     ]
#                 },
#                 {
#                     "line": 5,
#                     "pc": 7,
#                     "stack": [
#                         {
#                             "bytes": "",
#                             "type": 2,
#                             "uint": 123
#                         },
#                         {
#                             "bytes": "",
#                             "type": 2,
#                             "uint": 123
#                         }
#                     ]
#                 },
#                 {
#                     "line": 6,
#                     "pc": 8,
#                     "stack": [
#                         {
#                             "bytes": "",
#                             "type": 2,
#                             "uint": 1
#                         }
#                     ]
#                 }
#             ]
#         }
#     ]
# }
