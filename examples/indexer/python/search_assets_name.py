# search_assets_name
import json
# requires Python SDK version 1.3 or higher
from algosdk.v2client import indexer

# instantiate indexer client
myindexer = indexer.IndexerClient(indexer_token="", indexer_address="http://localhost:8980")

response = myindexer.search_assets(
    name="DevDocsCoin")
print("Asset Name Info: " + json.dumps(response, indent=2, sort_keys=True))
