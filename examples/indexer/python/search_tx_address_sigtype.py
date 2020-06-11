# search_tx_address_sigtype.py
import json
# requires Python SDK version 1.3 or higher
from algosdk.v2client import indexer

myindexer = indexer.IndexerClient(
    indexer_token="", indexer_address="http://localhost:8980")

response = myindexer.search_transactions_by_address(
    address="XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4", sig_type="msig")

print("sig_type: msig = " +
      json.dumps(response, indent=2, sort_keys=True))
