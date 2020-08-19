import json
from algosdk.v2client import indexer

# instantiate indexer client
myindexer = indexer.IndexerClient(indexer_token="", indexer_address="http://localhost:59998")
response = myindexer.search_applications()

print("Response Info: " + json.dumps(response, indent=2, sort_keys=True))

# response should look similar to this...
# Response Info: {
# "applications": [
#     {
#         "id": 20,
#         "params": {
#             "approval-program": "ASABASI=",
#             "clear-state-program": "ASABASI=",
#             "creator": "DQ5PMCTEBZLM4UJEDSGZLKAV6ZGXRK2C5WYAFC63RSHI54ASQSJHDMMTUM",
#             "global-state-schema": {
#                 "num-byte-slice": 0,
#                 "num-uint": 0
#             },
#             "local-state-schema": {
#                 "num-byte-slice": 0,
#                 "num-uint": 0
#             }
#         }
#     },
#     {
#         "id": 22,
#         "params": {
#             "approval-program": null,
#             "clear-state-program": null,
#             "creator": "GHFRLVOMKJNTJ4HY3P74ZR4CNE2PB7CYAUAJ6HVAVVDX7ZKEMLJX6AAF4M",
#             "global-state-schema": {
#                 "num-byte-slice": 0,
#                 "num-uint": 0
#             },
#             "local-state-schema": {
#                 "num-byte-slice": 0,
#                 "num-uint": 0
#             }
#         }
#     },
#     ...
