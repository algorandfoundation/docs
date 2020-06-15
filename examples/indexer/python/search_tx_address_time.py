# search_tx_address_time.py
import json
# requires Python SDK version 1.3 or higher
from algosdk.v2client import indexer

# instantiate indexer client
myindexer = indexer.IndexerClient(indexer_token="", indexer_address="http://localhost:8980")

# gets transactions for an account after a timestamp
response = myindexer.search_transactions_by_address(
    address="XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4", start_time="2020-06-03T10:00:00-05:00")

print("Transaction Start Time 2020-06-03T10:00:00-05:00 = " +
      json.dumps(response, indent=2, sort_keys=True))


