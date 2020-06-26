from algosdk import algod, transaction, account, mnemonic

try:

    # create logic sig
    # program = b"hex-encoded-program"
    # b"\x01\x20\x01\x00\x22 is `int 0`
    # see more info here: https://developer.algorand.org/docs/features/asc1/sdks/#accessing-teal-program-from-sdks
    program = b"\x01\x20\x01\x00\x22"
    lsig = transaction.LogicSig(program)
    sender = lsig.address()

    # create an algod client
    algod_token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" 
    algod_address = "http://localhost:4001" 
    receiver = "ATTR6RUEHHBHXKUHT4GUOYWNBVDV2GJ5FHUWCSFZLHD55EVKZWOWSM7ABQ" 

    # algod_token = "algod-token" < PLACEHOLDER >
    # algod_address = "algod-address" < PLACEHOLDER >
    # receiver = "receiver-address" < PLACEHOLDER >

    acl = algod.AlgodClient(algod_token, algod_address)

    # get suggested parameters
    params = acl.suggested_params()
    gen = params["genesisID"]
    gh = params["genesishashb64"]
    last_round = params["lastRound"]
    fee = params["fee"]
    amount = 10000 
    closeremainderto = None

    # create a transaction
    txn = transaction.PaymentTxn(
        sender, fee, last_round, last_round+100, gh, receiver, amount, closeremainderto)
    # Create the LogicSigTransaction with contract account LogicSig
    lstx = transaction.LogicSigTransaction(txn, lsig)
    print("This transaction is expected to fail as it is int 0 , alwasy false")
    # send raw LogicSigTransaction to network
    txid = acl.send_transaction(lstx)
    print("Transaction ID: " + txid)
except Exception as e:
    print(e)
