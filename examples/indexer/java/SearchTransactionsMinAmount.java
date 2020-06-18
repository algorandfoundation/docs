// SearchTransactionsMinAmount.java
// requires java-algorand-sdk 1.4.0 or higher (see pom.xml)
package com.algorand.SearchTransactionsMinAmount;

import com.algorand.algosdk.v2.client.common.IndexerClient;
import com.algorand.algosdk.v2.client.common.Client;
import org.json.JSONObject;

public class SearchTransactionsMinAmount {
    public Client indexerInstance = null;
    // utility function to connect to a node
    private Client connectToNetwork(){
        final String INDEXER_API_ADDR = "localhost";
        final int INDEXER_API_PORT = 8980;       
        IndexerClient indexerClient = new IndexerClient(INDEXER_API_ADDR, INDEXER_API_PORT); 
        return indexerClient;
    }
    public static void main(String args[]) throws Exception {
        SearchTransactionsMinAmount ex = new SearchTransactionsMinAmount();
        IndexerClient indexerClientInstance = (IndexerClient)ex.connectToNetwork();
        Long min_amount = Long.valueOf(10);     
        String response = indexerClientInstance
                .searchForTransactions()
                .currencyGreaterThan(min_amount).execute().toString();
        JSONObject jsonObj = new JSONObject(response.toString());
        System.out.println("Transaction Info: " + jsonObj.toString(2)); // pretty print json
    }
 }