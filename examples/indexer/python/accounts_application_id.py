# accounts_assetid.py
import json
# requires Python SDK version 1.3 or higher
from algosdk.v2client import indexer

# instantiate indexer client

# myindexer_token = 'B3SU4KcVKi94Jap2VXkK83xx38bsv95K5UZm2lab'
# myindexer_address = 'https://betanet-algorand.api.purestake.io/idx2/'
# myindexer_header = {'X-Api-key': myindexer_token}

myindexer_token = 'B3SU4KcVKi94Jap2VXkK83xx38bsv95K5UZm2lab'
myindexer_address = 'https://betanet-algorand.api.purestake.io/idx2/'
myindexer_header = {'X-Api-key': myindexer_token}

# myindexer_token = 'YddOUGbAjHLr1uPZtZwHOvMDmXvR1Zvw1f3Roj2PT1ufenXbNyIxIz0IeznrLbDsF'
# myindexer_address = 'https://indexer-internal-betanet.aws.algodev.network:443'
# myindexer_header = {'X-Indexer-API-Token': myindexer_token}

# https://indexer-internal-beta//net.aws.algodev.network

# myindexer = indexer.IndexerClient(indexer_token="", indexer_address="http://localhost:8980")
# myindexer = indexer.IndexerClient(
#     indexer_token="", indexer_address=myindexer_address, headers=myindexer_header)
myindexer = indexer.IndexerClient(
    indexer_token="", indexer_address=myindexer_address, headers=myindexer_header)

response = myindexer.accounts(
    application_id=2672020)
print("Account Info: " + json.dumps(response, indent=2, sort_keys=True))

