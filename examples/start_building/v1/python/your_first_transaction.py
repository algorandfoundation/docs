import json
import time
import base64
from algosdk import algod
from algosdk import mnemonic
from algosdk import transaction

# utility for waiting on a transaction confirmation
def wait_for_confirmation( algod_client, txid ):
    while True:
        txinfo = algod_client.pending_transaction_info(txid)
        if txinfo.get('round') and txinfo.get('round') > 0:
            print("Transaction {} confirmed in round {}.".format(txid, txinfo.get('round')))
            break
        else:
            print("Waiting for confirmation...")
            algod_client.status_after_block(algod_client.status().get('lastRound') +1)

def gettingStartedExample():
    algod_address = <algod-address>
    algod_token = <algod-token>
    algod_client = algod.AlgodClient(algod_token, algod_address)

    passphrase = <25-word-mnemonic>

    private_key = mnemonic.to_private_key(passphrase)
    my_address = mnemonic.to_public_key(passphrase)
    print("My address: {}".format(my_address))

    account_info = algod_client.account_info(my_address)
    print("Account balance: {} microAlgos".format(account_info.get('amount')))

    params = algod_client.suggested_params()
    note = "Hello World".encode()
    receiver = "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A"

    data = {
        "sender": my_address,
        "receiver": receiver,
        "fee": params.get('minFee'),
        "flat_fee": True,
        "amt": 1000000,
        "first": params.get('lastRound'),
        "last": params.get('lastRound') + 1000,
        "note": note,
        "gen": params.get('genesisID'),
        "gh": params.get('genesishashb64')
    }

    txn = transaction.PaymentTxn(**data)
    signed_txn = txn.sign(private_key)
    txid = signed_txn.transaction.get_txid()
    print("Signed transaction with txID: {}".format(txid))

    algod_client.send_transaction(signed_txn)

    # wait for confirmation
    wait_for_confirmation( algod_client, txid) 

    # Read the transction
    try:
        confirmed_txn = algod_client.transaction_info(my_address, txid)
    except Exception as err:
        print(err)
    print("Transaction information: {}".format(json.dumps(confirmed_txn, indent=4)))
    print("Decoded note: {}".format(base64.b64decode(confirmed_txn.get('noteb64')).decode()))


gettingStartedExample()