package com.algorand.javatest.firsttransaction;

import com.algorand.algosdk.account.Account;
import com.algorand.algosdk.crypto.Address;
import com.algorand.algosdk.transaction.SignedTransaction;
import com.algorand.algosdk.transaction.Transaction;
import com.algorand.algosdk.util.Encoder;
import com.algorand.algosdk.v2.client.common.AlgodClient;
import com.algorand.algosdk.v2.client.common.Response;
import com.algorand.algosdk.v2.client.model.NodeStatusResponse;
import com.algorand.algosdk.v2.client.model.PendingTransactionResponse;
import com.algorand.algosdk.v2.client.model.TransactionParametersResponse;
import org.json.JSONArray;
import org.json.JSONObject;

public class YourFirstTransaction {
    public AlgodClient client = null;

    // utility function to connect to a node
    private AlgodClient connectToNetwork() {

        // final String ALGOD_API_ADDR = "https://testnet-algorand.api.purestake.io/ps2";
        // Initialize an algod client
        final String ALGOD_API_ADDR = "localhost";
        final Integer ALGOD_PORT = 4001;
        final String ALGOD_API_TOKEN = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";

        // AlgodClient client = (AlgodClient) new AlgodClient(ALGOD_API_ADDR, ALGOD_PORT, ALGOD_API_TOKEN);
        // hackathon - demos instance
        // final String ALGOD_API_ADDR = "http://hackathon.algodev.network";
        // final Integer ALGOD_PORT = 9100;
        // final String ALGOD_API_TOKEN = "ef920e2e7e002953f4b29a8af720efe8e4ecc75ff102b165e0472834b25832c1";
 
        AlgodClient client = new AlgodClient(ALGOD_API_ADDR,
        ALGOD_PORT, ALGOD_API_TOKEN);
 
        // AlgodClient client = new AlgodClient();
        //  client.setBasePath(ALGOD_API_ADDR);

        // client.addDefaultHeader("X-API-Key", "YOUR API KEY");

       //  AlgodApi algodApiInstance = new AlgodApi(client); 
        return client;
    }
    /**
     * utility function to wait on a transaction to be confirmed
     * the timeout parameter indicates how many rounds do you wish to check pending transactions for
     */
    public PendingTransactionResponse waitForConfirmation(AlgodClient myclient, String txID, Integer timeout)
            throws Exception {
        if (myclient == null || txID == null || timeout < 0) {
            throw new IllegalArgumentException("Bad arguments for waitForConfirmation.");
        }
        Response<NodeStatusResponse> resp = myclient.GetStatus().execute();
        if (!resp.isSuccessful()) {
            throw new Exception(resp.message());
        }
        NodeStatusResponse nodeStatusResponse = resp.body();
        Long startRound = nodeStatusResponse.lastRound+1;
        Long currentRound = startRound;
        while (currentRound < (startRound + timeout)) { 
                // Check the pending transactions                 
                Response<PendingTransactionResponse> resp2 = myclient.PendingTransactionInformation(txID).execute();
                if (resp2.isSuccessful()) {
                    PendingTransactionResponse pendingInfo = resp2.body();               
                    if (pendingInfo != null) {
                        if (pendingInfo.confirmedRound != null && pendingInfo.confirmedRound > 0) {
                            // Got the completed Transaction
                            return pendingInfo;                     
                        }
                        if (pendingInfo.poolError != null && pendingInfo.poolError.length() > 0) {
                            // If there was a pool error, then the transaction has been rejected!
                            throw new Exception("The transaction has been rejected with a pool error: " + pendingInfo.poolError);
                        }
                    }
                }
        
                Response<NodeStatusResponse> resp3 = myclient.WaitForBlock(currentRound).execute();
                if (!resp3.isSuccessful()) {
                    throw new Exception(resp3.message());
                }   
                currentRound++;                  
        }
        throw new Exception("Transaction not confirmed after " + timeout + " rounds!");
    }

    public void gettingStartedExample() throws Exception {

        if (client == null)
            this.client = connectToNetwork();

        // Import your private key mnemonic and address
        final String PASSPHRASE = "patrol target joy dial ethics flip usual fatigue bulb security prosper brand coast arch casino burger inch cricket scissors shoe evolve eternal calm absorb school";
        com.algorand.algosdk.account.Account myAccount = new Account(PASSPHRASE);
        System.out.println("My Address: " + myAccount.getAddress());

        String myAddress = printBalance(myAccount);

        try {
            // Construct the transaction
            final String RECEIVER = "L5EUPCF4ROKNZMAE37R5FY2T5DF2M3NVYLPKSGWTUKVJRUGIW4RKVPNPD4";
            String note = "Hello World";
            TransactionParametersResponse params = client.TransactionParams().execute().body();
            Transaction txn = Transaction.PaymentTransactionBuilder()
                    .sender(myAddress)
                    .note(note.getBytes())
                    .amount(100000)
                    .receiver(new Address(RECEIVER))
                    .suggestedParams(params)
                    .build();

            // Sign the transaction
            SignedTransaction signedTxn = myAccount.signTransaction(txn);
            System.out.println("Signed transaction with txid: " + signedTxn.transactionID);

            // Submit the transaction to the network
            byte[] encodedTxBytes = Encoder.encodeToMsgPack(signedTxn);
            String id = client.RawTransaction().rawtxn(encodedTxBytes).execute().body().txId;
            System.out.println("Successfully sent tx with ID: " + id);

            // Wait for transaction confirmation
            PendingTransactionResponse pTrx = waitForConfirmation(client, id, 4);

            System.out.println("Transaction " + id + " confirmed in round " + pTrx.confirmedRound);
            // Read the transaction
            JSONObject jsonObj = new JSONObject(pTrx.toString());
            System.out.println("Transaction information (with notes): " + jsonObj.toString(2));
            System.out.println("Decoded note: " + new String(pTrx.txn.tx.note));
            printBalance(myAccount);
            } catch (Exception e) {
                System.err.println("Exception when calling algod#transactionInformation: " + e.getMessage());
        }
    }

    private String printBalance(com.algorand.algosdk.account.Account myAccount) throws Exception {
        String myAddress = myAccount.getAddress().toString();

        com.algorand.algosdk.v2.client.model.Account accountInfo = client.AccountInformation(myAccount.getAddress())
                .execute().body();

        System.out.println(String.format("Account Balance: %d microAlgos", accountInfo.amount));
        return myAddress;
    }

    public static void main(String args[]) throws Exception {
        YourFirstTransaction t = new YourFirstTransaction();
        t.gettingStartedExample();
    }
}