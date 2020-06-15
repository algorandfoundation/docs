# search_tx_address_txntype.py
import json
# requires Python SDK version 1.3 or higher
from algosdk.v2client import indexer

# instantiate indexer client
myindexer = indexer.IndexerClient(indexer_token="", indexer_address="http://localhost:8980")

response = myindexer.search_transactions_by_address(
    address="SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU",
    txn_type="acfg")

print("txn_type: acfg = " +
      json.dumps(response, indent=2, sort_keys=True))
