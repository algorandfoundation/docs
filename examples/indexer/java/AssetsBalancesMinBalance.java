// AssetsBalancesMinBalance.java
// requires java-algorand-sdk 1.4.0 or higher (see pom.xml)
package com.algorand.javatest.indexer;

import com.algorand.algosdk.v2.client.common.IndexerClient;
import com.algorand.algosdk.v2.client.common.Client;
import com.algorand.algosdk.v2.client.common.Response;
import com.algorand.algosdk.v2.client.model.AssetBalancesResponse;
import org.json.JSONObject;

public class AssetsBalancesMinBalance {
    public Client indexerInstance = null;
    // utility function to connect to a node
    private Client connectToNetwork(){
        final String INDEXER_API_ADDR = "localhost";
        final int INDEXER_API_PORT = 8981;       
        IndexerClient indexerClient = new IndexerClient(INDEXER_API_ADDR, INDEXER_API_PORT); 
        return indexerClient;
    }
    public static void main(String args[]) throws Exception {
        AssetsBalancesMinBalance ex = new AssetsBalancesMinBalance();
        IndexerClient indexerClientInstance = (IndexerClient)ex.connectToNetwork();
        Long asset_id = Long.valueOf(440307);
        // searches for asset greater than currencyGreaterThan
        Long currencyGreaterThan = Long.valueOf(0);        
        Response<AssetBalancesResponse> response = indexerClientInstance
            .lookupAssetBalances(asset_id)
            .currencyGreaterThan(currencyGreaterThan).execute();
        if (!response.isSuccessful()) {
            throw new Exception(response.message());
        } 
        JSONObject jsonObj = new JSONObject(response.body().toString());
        System.out.println("Asset Info: " + jsonObj.toString(2)); // pretty print json                        

    }
 }