# search_transactions_paging.py
import json
# requires Python SDK version 1.3 or higher
from algosdk.v2client import indexer

# instantiate indexer client
myindexer = indexer.IndexerClient(indexer_token="", indexer_address="http://localhost:8980")

nexttoken = ""
numtx = 1

# loop using next_page to paginate until there are no more transactions in the response
# for the limit (max is 1000  per request)

while (numtx > 0):

    response = myindexer.search_transactions(
        min_amount=100000000000000, limit=2, next_page=nexttoken) 
    transactions = response['transactions']
    numtx = len(transactions)
    if (numtx > 0):
        nexttoken = response['next-token']
        # Pretty Printing JSON string 
        print("Tranastion Info: " + json.dumps(response, indent=2, sort_keys=True))

