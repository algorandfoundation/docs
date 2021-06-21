// SearchTransactionsPaging.java
// requires java-algorand-sdk 1.4.0 or higher (see pom.xml)
package com.algorand.javatest.indexer;

import com.algorand.algosdk.v2.client.common.IndexerClient;
import com.algorand.algosdk.v2.client.common.Client;
import com.algorand.algosdk.v2.client.common.Response;
import com.algorand.algosdk.v2.client.model.TransactionsResponse;
import org.json.JSONArray;
import org.json.JSONObject;

public class SearchTransactionsPaging {
    public Client indexerInstance = null;
    // utility function to connect to a node
    private Client connectToNetwork(){
        // final String INDEXER_API_ADDR = "localhost";
        // final int INDEXER_API_PORT = 8980;
        final String INDEXER_API_ADDR = "localhost";
        // final String INDEXER_TOKEN = "";
        final int INDEXER_API_PORT = 8980;            
        IndexerClient indexerClient = new IndexerClient(INDEXER_API_ADDR, INDEXER_API_PORT); 
        return indexerClient;
    }

    public static void main(String args[]) throws Exception {
        SearchTransactionsPaging ex = new SearchTransactionsPaging();
        IndexerClient indexerClientInstance = (IndexerClient) ex.connectToNetwork();
        String nexttoken = "";
        Integer numtx = 1;  
        Long maxround=Long.valueOf(30000);           
        // loop until there are no more transactions in the response
        // for the limit (max limit is 1000 per request)
        while (numtx > 0) {
            Long min_amount = Long.valueOf(500000000000L);
            Long limit = Long.valueOf(4);
            String next_page = nexttoken;
            Response<TransactionsResponse> response = indexerClientInstance
                .searchForTransactions()
                .next(next_page)
                .currencyGreaterThan(min_amount)
                .maxRound(maxround)
                .limit(limit).execute();
            if (!response.isSuccessful()) {
                throw new Exception(response.message());
            }        
            JSONObject jsonObj = new JSONObject(response.body().toString());

            JSONArray jsonArray = (JSONArray) jsonObj.get("transactions");
            numtx = jsonArray.length();
            if (numtx > 0) {

                nexttoken = jsonObj.get("next-token").toString();
                JSONObject jsonObjAll = new JSONObject(response.body().toString());
                System.out.println("Transaction Info: " + jsonObjAll.toString(2)); // pretty print json
            }
        }

    }
 }