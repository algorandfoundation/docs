// LookupApplication.java
// requires java-algorand-sdk 1.4.0 or higher (see pom.xml)
package com.algorand.javatest.indexer;

import com.algorand.algosdk.v2.client.common.IndexerClient;
import com.algorand.algosdk.v2.client.common.Client;
import com.algorand.algosdk.crypto.Address;
import org.json.JSONObject;
import com.algorand.algosdk.v2.client.model.ApplicationResponse;
import com.algorand.algosdk.v2.client.common.Response;

public class LookupApplication {
    public Client indexerInstance = null;
    // utility function to connect to a node
    private Client connectToNetwork(){
        final String INDEXER_API_ADDR = "localhost";
        final int INDEXER_API_PORT = 8981;
        IndexerClient indexerClient = new IndexerClient(INDEXER_API_ADDR, INDEXER_API_PORT);
        return indexerClient;
    }

    public static void main(String args[]) throws Exception {
        LookupApplication ex = new LookupApplication();
        IndexerClient indexerClientInstance = (IndexerClient)ex.connectToNetwork();
        Long application_id = Long.valueOf(15974179);   
        Response<ApplicationResponse> response = indexerClientInstance
            .lookupApplicationByID(application_id).execute(); 
        if (!response.isSuccessful()) {
            throw new Exception(response.message());
        } 
        JSONObject jsonObj = new JSONObject(response.body().toString());
        System.out.println("Response Info: " + jsonObj.toString(2)); // pretty print json          

    }
 }
// response should look similar to this...
//  Response Info: {
//     "application": {
//       "deleted": false,
//       "created-at-round": 14412247,
//       "id": 15974179,
//       "params": {
//         "clear-state-program": "AiABASJD",
//         "global-state": [
//           {
//             "value": {
//               "bytes": "",
//               "type": 2,
//               "uint": 200
//             },
//             "key": "YXNzZXRfY29lZmZpY2llbnQ="
//           },
//           {
//             "value": {
//               "bytes": "",
//               "type": 2,
//               "uint": 13164495
//             },
//             "key": "YXNzZXRfaWQ="
//           },
//           {
//             "value": {
//               "bytes": "",
//               "type": 2,
//               "uint": 4294967316
//             },
//             "key": "b3B0aW9uX2ZpcnN0"
//           },
//           {
//             "value": {
//               "bytes": "",
//               "type": 2,
//               "uint": 4294967326
//             },
//             "key": "b3B0aW9uX3NlY29uZA=="
//           },
//           {
//             "value": {
//               "bytes": "v5WyR1XaHaNu97Sngrp+bcpo3BLYGZNcFgFUUXK59Lg=",
//               "type": 1,
//               "uint": 0
//             },
//             "key": "Q3JlYXRvcg=="
//           },
//           {
//             "value": {
//               "bytes": "bXVjaGRlY2lzaW9u",
//               "type": 1,
//               "uint": 0
//             },
//             "key": "TmFtZQ=="
//           },
//           {
//             "value": {
//               "bytes": "",
//               "type": 2,
//               "uint": 4294967296
//             },
//             "key": "b3B0aW9uX2Fub3RoZXJvbmU="
//           },
//           {
//             "value": {
//               "bytes": "",
//               "type": 2,
//               "uint": 4294967296
//             },
//             "key": "b3B0aW9uX2hleQ=="
//           },
//           {
//             "value": {
//               "bytes": "",
//               "type": 2,
//               "uint": 4294967296
//             },
//             "key": "b3B0aW9uX3RoaXJk"
//           },
//           {
//             "value": {
//               "bytes": "",
//               "type": 2,
//               "uint": 1623000921
//             },
//             "key": "dm90aW5nX2VuZF90aW1l"
//           },
//           {
//             "value": {
//               "bytes": "",
//               "type": 2,
//               "uint": 1622136921
//             },
//             "key": "dm90aW5nX3N0YXJ0X3RpbWU="
//           },
//           {
//             "value": {
//               "bytes": "",
//               "type": 2,
//               "uint": 4294967296
//             },
//             "key": "b3B0aW9uX3NpeA=="
//           }
//         ],
//         "creator": "X6K3ER2V3IO2G3XXWSTYFOT6NXFGRXAS3AMZGXAWAFKFC4VZ6S4K4JLOFY",
//         "local-state-schema": {
//           "num-uint": 1,
//           "num-byte-slice": 1
//         },
//         "approval-program": "AiAFAAUEAYCAgIAQJgwEdm90ZQthZGRfb3B0aW9ucwdDcmVhdG9yBE5hbWUIYXNzZXRfaWQRYXNzZXRfY29lZmZpY2llbnQRdm90aW5nX3N0YXJ0X3RpbWUPdm90aW5nX2VuZF90aW1lC05VTExfT1BUSU9OB29wdGlvbl8UUVZvdGVEZWNpc2lvbkNyZWRpdHMBLTEYIhJAACYxGSMSQADJMRkkEkAAxzEZJRJAAMU2GgAoEkAA6zYaACkSQAFcACoxAGcrNhoAZycENhoGF2cnBTYaBxdnJwY2GggXZycHNhoJF2cyBycGZA1BAAIiQzYaAScIE0AABSVDQgAJJwk2GgFQIQRnNhoCJwgTQAAFJUNCAAknCTYaAlAhBGc2GgMnCBNAAAUlQ0IACScJNhoDUCEEZzYaBCcIE0AABSVDQgAJJwk2GgRQIQRnNhoFJwgTQAAFJUNCAAknCTYaBVAhBGclQyVDQgExIkNCASwiQ0IBJzIHJwZkDUEAAiJDIicEZHAANQA1ATQAQAAFIkNCAAwiJwo0AScFZAtmJUNCAPkyBycGZAxBAAIiQzIHJwdkDUEAAiJDIicJNhoBUGU1AjUDNAJAAAUiQ0IASiInCiInCmI2GgIXNhoCFwsJZiInCmIiD0AABSJDQgAoNhoDJwsSQAARJwk2GgFQNAM2GgIXCGdCAA4nCTYaAVA0AzYaAhcJZyVDQgCAMgcnBmQNQQACIkM2GgEnCBNAAAUlQ0IACScJNhoBUCEEZzYaAicIE0AABSVDQgAJJwk2GgJQIQRnNhoDJwgTQAAFJUNCAAknCTYaA1AhBGc2GgQnCBNAAAUlQ0IACScJNhoEUCEEZzYaBScIE0AABSVDQgAJJwk2GgVQIQRnJUM=",
//         "global-state-schema": {
//           "num-uint": 10,
//           "num-byte-slice": 3
//         }
//       }
//     },
//     "current-round": 14694513
//   }