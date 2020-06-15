# search_tx_address_block_range.py
import json
# requires Python SDK version 1.3 or higher
from algosdk.v2client import indexer

# instantiate indexer client
myindexer = indexer.IndexerClient(indexer_token="", indexer_address="http://localhost:8980")

response = myindexer.search_transactions_by_address(
    address="XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4", min_round=7048876, max_round=7048878)

print("min-max rounds: 7048876-7048878 = " +
      json.dumps(response, indent=2, sort_keys=True))
