// AccountInfo.java
// requires java-algorand-sdk 1.4.0 or higher (see pom.xml)
package com.algorand.AccountInfo;

import com.algorand.algosdk.v2.client.common.IndexerClient;
import com.algorand.algosdk.v2.client.common.Client;
import com.algorand.algosdk.crypto.Address;
import org.json.JSONObject;

public class AccountInfo {
    public Client indexerInstance = null;
    // utility function to connect to a node
    private Client connectToNetwork(){
        final String INDEXER_API_ADDR = "localhost";
        final int INDEXER_API_PORT = 8980;       
        IndexerClient indexerClient = new IndexerClient(INDEXER_API_ADDR, INDEXER_API_PORT); 
        return indexerClient;
    }
    public static void main(String args[]) throws Exception {
        AccountInfo ex = new AccountInfo();
        IndexerClient indexerClientInstance = (IndexerClient)ex.connectToNetwork();
        Address account = new Address("7WENHRCKEAZHD37QMB5T7I2KWU7IZGMCC3EVAO7TQADV7V5APXOKUBILCI");
         String response = indexerClientInstance.lookupAccountByID(account).execute().toString();
        JSONObject jsonObj = new JSONObject(response.toString());
        System.out.println("Account Info: " + jsonObj.toString(2)); // pretty print json
    }
 }