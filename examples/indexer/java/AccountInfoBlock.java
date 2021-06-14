// AccountInfoBlock.java
// requires java-algorand-sdk 1.4.0 or higher (see pom.xml)
package com.algorand.javatest.indexer;
import com.algorand.algosdk.v2.client.common.IndexerClient;
import com.algorand.algosdk.v2.client.common.Response;
import com.algorand.algosdk.v2.client.model.AccountResponse;
import com.algorand.algosdk.v2.client.common.Client;
import com.algorand.algosdk.crypto.Address;
import org.json.JSONObject;

public class AccountInfoBlock {
    public Client indexerInstance = null;
    // utility function to connect to a node
    private Client connectToNetwork(){
        final String INDEXER_API_ADDR = "localhost";
        final int INDEXER_API_PORT = 8980;       
        IndexerClient indexerClient = new IndexerClient(INDEXER_API_ADDR, INDEXER_API_PORT); 
        return indexerClient;
    }
    public static void main(String args[]) throws Exception {
        AccountInfoBlock ex = new AccountInfoBlock();
        IndexerClient indexerClientInstance = (IndexerClient)ex.connectToNetwork();
        Address account = new Address("ZYATCNPXNCPI4RE7VPKH7OBLFDFWJGNWY2RPE35DH5IPZFHWBWRL7TS3LU");
        Long round = Long.valueOf(14653225);
        Response<AccountResponse>  response = indexerClientInstance
            .lookupAccountByID(account)
            .round(round).execute();
        if (!response.isSuccessful()) {
            throw new Exception(response.message());
        } 

        JSONObject jsonObj = new JSONObject(response.body().toString());
        System.out.println("Account Info for block: " + jsonObj.toString(2)); // pretty print json

    }
 }