// SearchTxAddressTime.java
// requires java-algorand-sdk 1.4.0 or higher (see pom.xml)
package com.algorand.javatest.indexer;

import com.algorand.algosdk.v2.client.common.IndexerClient;
import com.algorand.algosdk.v2.client.common.Client;
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
        Address account = new Address("XIU7HGGAJ3QOTATPDSIIHPFVKMICXKHMOR2FJKHTVLII4FAOA3CYZQDLG4");
    	SimpleDateFormat sdf = new SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ssXXX");
        Date start_time = sdf.parse("2020-06-03T10:00:00-05:00");
        String response = indexerClientInstance
                .searchForTransactions()
                .address(account)
                .afterTime(start_time).execute().toString();
        JSONObject jsonObj = new JSONObject(response.toString());
        System.out.println("start_time: 06/03/2020 11:00:00-05:00 = " + jsonObj.toString(2)); // pretty print json
    }
 }