# search_asset_transactions_role.py
import json
# requires Python SDK version 1.3 or higher
from algosdk.v2client import indexer

# instantiate indexer client
myindexer = indexer.IndexerClient(indexer_token="", indexer_address="http://localhost:8980")

response = myindexer.search_asset_transactions(
    asset_id=2044572, address_role="receiver", address="UF7ATOM6PBLWMQMPUQ5QLA5DZ5E35PXQ2IENWGZQLEJJAAPAPGEGC3ZYNI")

print("Asset Transaction Info: " + json.dumps(response, indent=2, sort_keys=True))
