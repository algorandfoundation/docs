# search_transactions_limit.py
import json
# requires Python SDK version 1.3 or higher
from algosdk.v2client import indexer

# instantiate indexer client
myindexer = indexer.IndexerClient(indexer_token="", indexer_address="http://localhost:8980")

response = myindexer.search_transactions(
    min_amount=10, limit=2)

# Pretty Printing JSON string 
print("Transaction Info: " + json.dumps(response, indent=2, sort_keys=True))

