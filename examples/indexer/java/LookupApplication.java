// LookupApplication.java
// requires java-algorand-sdk 1.4.0 or higher (see pom.xml)
package com.algorand.javatest.indexer;

import com.algorand.algosdk.v2.client.common.IndexerClient;
import com.algorand.algosdk.v2.client.common.Client;
import com.algorand.algosdk.crypto.Address;
import org.json.JSONObject;

public class LookupApplication {
    public Client indexerInstance = null;
    // utility function to connect to a node
    private Client connectToNetwork(){
        final String INDEXER_API_ADDR = "localhost";
        final int INDEXER_API_PORT = 59998;       
        IndexerClient indexerClient = new IndexerClient(INDEXER_API_ADDR, INDEXER_API_PORT); 
        return indexerClient;
    }
    public static void main(String args[]) throws Exception {
        LookupApplication ex = new LookupApplication();
        IndexerClient indexerClientInstance = (IndexerClient)ex.connectToNetwork();
        Long application_id = Long.valueOf(22);
        String response = indexerClientInstance.lookupApplicationByID(application_id).execute().toString();
        JSONObject jsonObj = new JSONObject(response.toString());
        System.out.println("Response Info: " + jsonObj.toString(2)); // pretty print json
    }
 }

// response information should look similar to this...
//  Response Info:
//  {
//   "application": {
//     "id": 22,
//     "params": {
//       "global-state": [],
//       "creator": "GHFRLVOMKJNTJ4HY3P74ZR4CNE2PB7CYAUAJ6HVAVVDX7ZKEMLJX6AAF4M",
//       "local-state-schema": {
//         "num-uint": 0,
//         "num-byte-slice": 0
//       },
//       "global-state-schema": {
//         "num-uint": 0,
//         "num-byte-slice": 0
//       }
//     }
//   },
//   "current-round": 377
// }