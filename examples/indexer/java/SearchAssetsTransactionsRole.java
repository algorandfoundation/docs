// SearchAssetsTransactionsRole.java
// requires java-algorand-sdk 1.4.0 or higher (see pom.xml)
package com.algorand.javatest.indexer;

import com.algorand.algosdk.v2.client.common.IndexerClient;
import com.algorand.algosdk.v2.client.common.Response;
import com.algorand.algosdk.v2.client.model.Enums.AddressRole;
import com.algorand.algosdk.v2.client.common.Client;
import com.algorand.algosdk.crypto.Address;
import com.algorand.algosdk.v2.client.model.TransactionsResponse;
import org.json.JSONObject;

public class SearchAssetsTransactionsRole {
    public Client indexerInstance = null;
    // utility function to connect to a node
    private Client connectToNetwork(){
        final String INDEXER_API_ADDR = "localhost";
        final int INDEXER_API_PORT = 8980;       
        IndexerClient indexerClient = new IndexerClient(INDEXER_API_ADDR, INDEXER_API_PORT); 
        return indexerClient;
    }
    public static void main(String args[]) throws Exception {
        SearchAssetsTransactionsRole ex = new SearchAssetsTransactionsRole();
        IndexerClient indexerClientInstance = (IndexerClient)ex.connectToNetwork();
        Long asset_id = Long.valueOf(408947);
        AddressRole addressRole = AddressRole.RECEIVER;
        Address account = new Address("G26NNWKJUPSTGVLLDHCUQ7LFJHMZP2UUAQG2HURLI6LOEI235YCQUNPQEI");           
        Response<TransactionsResponse> response = indexerClientInstance
            .searchForTransactions()
            .address(account)
            .assetId(asset_id)
            .addressRole(addressRole).execute();
        if (!response.isSuccessful()) {
            throw new Exception(response.message());
        }  
        JSONObject jsonObj = new JSONObject(response.body().toString());
        System.out.println("Asset Info for Name: " + jsonObj.toString(2)); // pretty print json

    }
 }