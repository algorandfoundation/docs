// SearchTxAddresstxntype.java
// requires java-algorand-sdk 1.4.0 or higher (see pom.xml)
package com.algorand.SearchTxAddresstxntype;

import com.algorand.algosdk.v2.client.common.IndexerClient;
import com.algorand.algosdk.v2.client.model.Enums.TxType;
import com.algorand.algosdk.v2.client.common.Client;
import org.json.JSONObject;
import com.algorand.algosdk.crypto.Address;

public class SearchTxAddresstxntype {
    public Client indexerInstance = null;
    // utility function to connect to a node
    private Client connectToNetwork(){
        final String INDEXER_API_ADDR = "localhost";
        final int INDEXER_API_PORT = 8980;       
        IndexerClient indexerClient = new IndexerClient(INDEXER_API_ADDR, INDEXER_API_PORT); 
        return indexerClient;
    }

    public static void main(String args[]) throws Exception {
        SearchTxAddresstxntype ex = new SearchTxAddresstxntype();
        IndexerClient indexerClientInstance = (IndexerClient) ex.connectToNetwork();
        Address account = new Address("SWOUICD7Y5PQBWWEYC4XZAQZI7FJRZLD5O3CP4GU2Y7FP3QFKA7RHN2WJU");
        TxType txType = TxType.ACFG;
        String response = indexerClientInstance.searchForTransactions().address(account).txType(txType).execute().toString();
        JSONObject jsonObj = new JSONObject(response.toString());
        System.out.println("txn_type: acfg = " + jsonObj.toString(2)); // pretty print json
    }
 }