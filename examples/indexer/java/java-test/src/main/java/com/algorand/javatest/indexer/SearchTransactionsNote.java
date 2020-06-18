// SearchTransactionsNote.java
// requires java-algorand-sdk 1.4.0 or higher (see pom.xml)
package com.algorand.javatest.indexer;
import com.algorand.algosdk.v2.client.common.IndexerClient;
import com.algorand.algosdk.v2.client.common.Client;
import org.json.JSONObject;
import com.algorand.algosdk.util.Encoder;

public class SearchTransactionsNote {
    public Client indexerInstance = null;
    // utility function to connect to a node
    private Client connectToNetwork(){
        final String INDEXER_API_ADDR = "localhost";
        final int INDEXER_API_PORT = 8980;       
        IndexerClient indexerClient = new IndexerClient(INDEXER_API_ADDR, INDEXER_API_PORT); 
        return indexerClient;
    }
    public static void main(String args[]) throws Exception {
        SearchTransactionsNote ex = new SearchTransactionsNote();
        IndexerClient indexerClientInstance = (IndexerClient)ex.connectToNetwork();
        byte[] encodednote = Encoder.decodeFromBase64("c2hvd2luZyBwcmVmaXg="); // "showing prefix"
        String response = indexerClientInstance
                .searchForTransactions()
                .notePrefix(encodednote).execute().toString();
        JSONObject jsonObj = new JSONObject(response.toString());
        System.out.println("Transaction Info: " + jsonObj.toString(2)); // pretty print json
    }
 }