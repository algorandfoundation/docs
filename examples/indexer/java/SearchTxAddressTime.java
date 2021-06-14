// SearchTxAddressTime.java
// requires java-algorand-sdk 1.4.0 or higher (see pom.xml)
package com.algorand.javatest.indexer;

import com.algorand.algosdk.v2.client.common.IndexerClient;
import com.algorand.algosdk.v2.client.common.Client;
import com.algorand.algosdk.v2.client.common.Response;
import com.algorand.algosdk.v2.client.model.TransactionsResponse;
import org.json.JSONObject;
import com.algorand.algosdk.crypto.Address;
import java.util.Date;
import java.text.SimpleDateFormat;

public class SearchTxAddressTime {
    public Client indexerInstance = null;
    // utility function to connect to a node
    private Client connectToNetwork(){
        final String INDEXER_API_ADDR = "localhost";
        final int INDEXER_API_PORT = 8980;       
        IndexerClient indexerClient = new IndexerClient(INDEXER_API_ADDR, INDEXER_API_PORT); 
        return indexerClient;
    }

    public static void main(String args[]) throws Exception {
        SearchTxAddressTime ex = new SearchTxAddressTime();
        IndexerClient indexerClientInstance = (IndexerClient) ex.connectToNetwork();
        Address account = new Address("RBSTLLHK2NJDL3ZH66MKSEX3BE2OWQ43EUM7S7YRVBJ2PRDRCKBSDD3YD4");
    	SimpleDateFormat sdf = new SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ssXXX");
        Date start_time = sdf.parse("2020-08-31T02:35:47-05:00");
        Response<TransactionsResponse> response = indexerClientInstance
                .searchForTransactions()
                .address(account)
                .afterTime(start_time).execute();
        if (!response.isSuccessful()) {
            throw new Exception(response.message());
        }  
        JSONObject jsonObj = new JSONObject(response.body().toString());
        System.out.println("start_time: 08/31/2020 02:35:47 = " + jsonObj.toString(2)); // pretty print json
    }
 }