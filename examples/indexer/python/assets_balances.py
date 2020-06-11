# assets_balances.py
import json
# requires Python SDK version 1.3 or higher
from algosdk.v2client import indexer

# instantiate indexer client
myindexer = indexer.IndexerClient(indexer_token="", indexer_address="http://localhost:8980")

response = myindexer.asset_balances(
    asset_id=2044572)
print("Asset Balance: " + json.dumps(response, indent=2, sort_keys=True))
