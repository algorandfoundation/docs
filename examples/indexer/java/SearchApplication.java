// SearchApplication.java
// requires java-algorand-sdk 1.4.0 or higher (see pom.xml)
package com.algorand.javatest.indexer;

import com.algorand.algosdk.v2.client.common.IndexerClient;
import com.algorand.algosdk.v2.client.common.Client;
import com.algorand.algosdk.crypto.Address;
import com.algorand.algosdk.v2.client.common.Response;
import com.algorand.algosdk.v2.client.model.ApplicationsResponse;
import org.json.JSONObject;

public class SearchApplication {
    public Client indexerInstance = null;
    // utility function to connect to a node
    private Client connectToNetwork(){
        final String INDEXER_API_ADDR = "localhost";
        final int INDEXER_API_PORT = 8980;       
        IndexerClient indexerClient = new IndexerClient(INDEXER_API_ADDR, INDEXER_API_PORT); 
        return indexerClient;
    }
    public static void main(String args[]) throws Exception {
        SearchApplication ex = new SearchApplication();
        IndexerClient indexerClientInstance = (IndexerClient)ex.connectToNetwork();
        Long limit = 4L;
        Response<ApplicationsResponse> response = indexerClientInstance
            .searchForApplications()
            .limit(limit).execute();
        if (!response.isSuccessful()) {
            throw new Exception(response.message());
        } 

        JSONObject jsonObj = new JSONObject(response.body().toString());
        System.out.println("Response Info: " + jsonObj.toString(2)); // pretty print json

    }
 }

// response should look similar to this...
//  Response Info:
//  {
//   "next-token": "142",
//   "current-round": 377,
//   "applications": [
//     {
//       "id": 20,
//       "params": {
//         "clear-state-program": "ASABASI=",
//         "global-state": [],
//         "creator": "DQ5PMCTEBZLM4UJEDSGZLKAV6ZGXRK2C5WYAFC63RSHI54ASQSJHDMMTUM",
//         "local-state-schema": {
//           "num-uint": 0,
//           "num-byte-slice": 0
//         },
//         "approval-program": "ASABASI=",
//         "global-state-schema": {
//           "num-uint": 0,
//           "num-byte-slice": 0
//         }
//       }
//     },
//     {
//       "id": 22,
//       "params": {
//         "global-state": [],
//         "creator": "GHFRLVOMKJNTJ4HY3P74ZR4CNE2PB7CYAUAJ6HVAVVDX7ZKEMLJX6AAF4M",
//         "local-state-schema": {
//           "num-uint": 0,
//           "num-byte-slice": 0
//         },
//         "global-state-schema": {
//           "num-uint": 0,
//           "num-byte-slice": 0
//         }
//       }
//     },
//     ...