// SearchTxAddressBlockRange.java
// requires java-algorand-sdk 1.4.0 or higher (see pom.xml)
package com.algorand.SearchTxAddressBlockRange;

import com.algorand.algosdk.v2.client.common.IndexerClient;
import com.algorand.algosdk.v2.client.common.Client;
import org.json.JSONObject;
import com.algorand.algosdk.crypto.Address;

public class SearchTxAddressBlockRange {
    public Client indexerInstance = null;
    // utility function to connect to a node
    private Client connectToNetwork(){
        final String INDEXER_API_ADDR = "localhost";
        final int INDEXER_API_PORT = 8980;       
        IndexerClient indexerClient = new IndexerClient(INDEXER_API_ADDR, INDEXER_API_PORT); 
        return indexerClient;
    }

    public static void main(String args[]) throws Exception {
        SearchTxAddressBlockRange ex = new SearchTxAddressBlockRange();
        IndexerClient indexerClientInstance = (IndexerClient) ex.connectToNetwork();
        Address account = new Address("XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4");
        Long min_round = Long.valueOf(7048876);
        Long max_round = Long.valueOf(7048878);       
        String response = indexerClientInstance.searchForTransactions().address(account)
                    .minRound(min_round).maxRound(max_round).execute().toString();
        JSONObject jsonObj = new JSONObject(response.toString());
        System.out.println("Transaction Info: " + jsonObj.toString(2)); // pretty print json
    }
 }