import json
# requires Python SDK version 1.3 or higher
from algosdk.v2client import indexer

# instantiate indexer client

# myindexer_token = 'B3SU4KcVKi94Jap2VXkK83xx38bsv95K5UZm2lab'
# myindexer_address = 'https://testnet-algorand.api.purestake.io/idx2/'
# myindexer_header = {'X-Api-key': myindexer_token}

myindexer_address = 'http://localhost:59998'

myindexer = indexer.IndexerClient(
    indexer_token="", indexer_address=myindexer_address)
# myindexer = indexer.IndexerClient(
#     indexer_token="", indexer_address=myindexer_address, headers=myindexer_header)

nexttoken = ""
num_accounts = 1
# loop using next_page to paginate until there are no more accounts
# in the response
# (max is 100 default
# unless limit is used for max 1000 per request on accounts)
while (num_accounts > 0):
    response = myindexer.accounts(
        application_id=70, limit=2,next_page=nexttoken)
    accounts = response['accounts']
    num_accounts = len(accounts)
    if (num_accounts > 0):
        nexttoken = response['next-token']
        # Pretty Printing JSON string
        print("Account Info for Application ID: " + json.dumps(response, indent=2, sort_keys=True))
