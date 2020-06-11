# assets_balances_min_balance.py
import json
# requires Python SDK version 1.3 or higher
from algosdk.v2client import indexer

# instantiate indexer client
myindexer = indexer.IndexerClient(indexer_token="", indexer_address="http://localhost:8980")

# gets assets with a min balance of 200 for AssetID
response = myindexer.asset_balances(
    asset_id=2044572, min_balance=200)

print("Asset Balances :" + json.dumps(response, indent=2, sort_keys=True))
