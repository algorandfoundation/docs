# to be updated for v2
from algosdk import algod, transaction, account, mnemonic
from algosdk.v2client import algod
from algosdk.future.transaction import PaymentTxn, LogicSig


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

try:



    # create logic sig
    # program = b"hex-encoded-program"
    # b"\x01\x20\x01\x00\x22 is `int 0`
    # see more info here: https://developer.algorand.org/docs/features/asc1/sdks/#accessing-teal-program-from-sdks

    program = b"\x01\x20\x01\x00\x22"

    # program = b"hex-encoded-program"
    lsig = LogicSig(program)
    sender = lsig.address()

    #Recover the account that is wanting to delegate signature
    passphrase = "canal enact luggage spring similar zoo couple stomach shoe laptop middle wonder eager monitor weather number heavy skirt siren purity spell maze warfare ability ten"
    # passphrase = "25-word-mnemonic<PLACEHOLDER>"
    sk = mnemonic.to_private_key(passphrase)
    addr = account.address_from_private_key(sk)
    print("Address of Sender/Delgator: " + addr)

    # sign the logic signature with an account sk
    lsig.sign(sk)

    # create an algod client
    algod_token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
    algod_address = "http://localhost:4001"
    receiver = "ATTR6RUEHHBHXKUHT4GUOYWNBVDV2GJ5FHUWCSFZLHD55EVKZWOWSM7ABQ"

    # algod_token = "algod-token" < PLACEHOLDER >
    # algod_address = "algod-address" < PLACEHOLDER >
    # receiver = "receiver-address" < PLACEHOLDER >
    algod_client = algod.AlgodClient(algod_token, algod_address)

    # get suggested parameters
    params = algod_client.suggested_params()
    # comment out the next two (2) lines to use suggested fees
    params.flat_fee = True
    params.fee = 1000

    # build transaction
    amount = 10000 
    closeremainderto = None

    # create a transaction
    txn = PaymentTxn(
        addr, params, receiver, amount, closeremainderto)
    # Create the LogicSigTransaction with contract account LogicSig
    lstx = transaction.LogicSigTransaction(txn, lsig)

    # send raw LogicSigTransaction to network
    print("This transaction is expected to fail as it is int 0 , alwasy false")
    txid = algod_client.send_transaction(lstx)
    print("Transaction ID: " + txid)
    wait_for_confirmation(algod_client, txid)
except Exception as e:
    print(e)
