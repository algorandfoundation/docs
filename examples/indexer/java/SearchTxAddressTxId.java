// SearchTxAddressTxId.java
// requires java-algorand-sdk 1.4.0 or higher (see pom.xml)
package com.algorand.javatest.indexer;

import com.algorand.algosdk.v2.client.common.IndexerClient;
import com.algorand.algosdk.v2.client.common.Client;
import com.algorand.algosdk.v2.client.common.Response;
import com.algorand.algosdk.v2.client.model.TransactionsResponse;
import org.json.JSONObject;
import com.algorand.algosdk.crypto.Address;

public class SearchTxAddressTxId {
    public Client indexerInstance = null;
    // utility function to connect to a node
    private Client connectToNetwork(){
        final String INDEXER_API_ADDR = "localhost";
        final int INDEXER_API_PORT = 8981;       
        IndexerClient indexerClient = new IndexerClient(INDEXER_API_ADDR, INDEXER_API_PORT); 
        return indexerClient;
    }

    public static void main(String args[]) throws Exception {
        SearchTxAddressTxId ex = new SearchTxAddressTxId();
        IndexerClient indexerClientInstance = (IndexerClient) ex.connectToNetwork();
        Address account = new Address("RBSTLLHK2NJDL3ZH66MKSEX3BE2OWQ43EUM7S7YRVBJ2PRDRCKBSDD3YD4");
        String txid = "GZKSVXCVQASQ2KLI7JUWZ2LP54WAKPVA3TC2ZGQVAHKIWEMMFNEQ";
        Response<TransactionsResponse> response = indexerClientInstance
            .searchForTransactions()
            .address(account)
            .txid(txid).execute();        
        if (!response.isSuccessful()) {
            throw new Exception(response.message());
        }  
        JSONObject jsonObj = new JSONObject(response.body().toString());
        System.out.println("txid: GZKSVXCVQASQ2KLI7JUWZ2LP54WAKPVA3TC2ZGQVAHKIWEMMFNEQ = " + jsonObj.toString(2)); // pretty print json
    }
 }