// SearchAssetsName.java
// requires java-algorand-sdk 1.4.0 or higher (see pom.xml)
package com.algorand.javatest.indexer;

import com.algorand.algosdk.v2.client.common.IndexerClient;
import com.algorand.algosdk.v2.client.common.Client;
import org.json.JSONObject;

public class SearchAssetsName {
    public Client indexerInstance = null;
    // utility function to connect to a node
    private Client connectToNetwork(){
        final String INDEXER_API_ADDR = "localhost";
        final int INDEXER_API_PORT = 8980;       
        IndexerClient indexerClient = new IndexerClient(INDEXER_API_ADDR, INDEXER_API_PORT); 
        return indexerClient;
    }
    public static void main(String args[]) throws Exception {
        SearchAssetsName ex = new SearchAssetsName();
        IndexerClient indexerClientInstance = (IndexerClient)ex.connectToNetwork();
        String name = "DevDocsCoin";        
        String response = indexerClientInstance.searchForAssets()
                        .name(name).execute().toString();
        JSONObject jsonObj = new JSONObject(response.toString());
        System.out.println("Asset Info for Name: " + jsonObj.toString(2)); // pretty print json
    }
 }