package com.algorand.YourFirstTransaction;

import java.math.BigInteger;

import java.util.concurrent.TimeUnit;

import com.algorand.algosdk.account.Account;
import com.algorand.algosdk.algod.client.AlgodClient;
import com.algorand.algosdk.algod.client.ApiException;
import com.algorand.algosdk.algod.client.api.AlgodApi;
import com.algorand.algosdk.algod.client.auth.ApiKeyAuth;
import com.algorand.algosdk.algod.client.model.*;
import com.algorand.algosdk.crypto.Address;
import com.algorand.algosdk.crypto.Digest;
import com.algorand.algosdk.transaction.Transaction;
import com.algorand.algosdk.transaction.SignedTransaction;
import com.algorand.algosdk.util.Encoder;

public class YourFirstTransaction {
    public AlgodApi algodApiInstance = null;
    // utility function to connect to a node
    private AlgodApi connectToNetwork(){

        // Initialize an algod client
        final String ALGOD_API_ADDR = "algod-address<PLACEHOLDER>";
        final String ALGOD_API_TOKEN = "algod-token<PLACEHOLDER>";
        AlgodClient client = (AlgodClient) new AlgodClient().setBasePath(ALGOD_API_ADDR);
        ApiKeyAuth api_key = (ApiKeyAuth) client.getAuthentication("api_key");
        api_key.setApiKey(ALGOD_API_TOKEN);
        algodApiInstance = new AlgodApi(client);   
        return algodApiInstance;
    }
    // utility function to wait on a transaction to be confirmed    
    public void waitForConfirmation( String txID ) throws Exception{
        if( algodApiInstance == null ) connectToNetwork();
        while(true) {
            try {
                //Check the pending tranactions
                com.algorand.algosdk.algod.client.model.Transaction pendingInfo = algodApiInstance.pendingTransactionInformation(txID);
                if (pendingInfo.getRound() != null && pendingInfo.getRound().longValue() > 0) {
                    //Got the completed Transaction
                    System.out.println("Transaction " + pendingInfo.getTx() + " confirmed in round " + pendingInfo.getRound().longValue());
                    break;
                } 
                algodApiInstance.waitForBlock(BigInteger.valueOf( algodApiInstance.getStatus().getLastRound().longValue() +1 ) );
            } catch (Exception e) {
                throw( e );
            }
        }
    }

    public void gettingStartedExample() throws Exception {

        if( algodApiInstance == null ) connectToNetwork();

        // Import your private key mnemonic and address
        final String PASSPHRASE = <25-word-mnemonic>;
        com.algorand.algosdk.account.Account myAccount = new Account(PASSPHRASE);
        System.out.println("My Address: " + myAccount.getAddress());

        String myAddress = myAccount.getAddress().toString();
        com.algorand.algosdk.algod.client.model.Account accountInfo = 
            algodApiInstance.accountInformation(myAddress);
        System.out.println(String.format("Account Balance: %d microAlgos", accountInfo.getAmount()));

        // Construct the transaction
        final String RECEIVER = "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A";
        BigInteger fee;
        String genesisID;
        Digest genesisHash;
        BigInteger firstValidRound;
        fee = BigInteger.valueOf(1000);
        try {
            TransactionParams params = algodApiInstance.transactionParams();
            genesisHash = new Digest(params.getGenesishashb64());
            genesisID = params.getGenesisID();
            System.out.println("Minimum Fee: " + fee);
            firstValidRound = params.getLastRound();
            System.out.println("Current Round: " + firstValidRound);
        } catch (ApiException e) {
            throw new RuntimeException("Could not get params", e);
        }
        BigInteger amount = BigInteger.valueOf(1000000); // microAlgos
        BigInteger lastValidRound = firstValidRound.add(BigInteger.valueOf(1000)); // 1000 is the max tx window
        String note = "Hello World";


        Transaction txn = new Transaction(myAccount.getAddress(), fee, firstValidRound,
                lastValidRound, note.getBytes(), amount, new Address(RECEIVER),
                genesisID, genesisHash);

        // Sign the transaction
        SignedTransaction signedTxn = myAccount.signTransaction(txn);
        System.out.println("Signed transaction with txid: " + signedTxn.transactionID);

        // Submit the transaction to the network
        try {
            byte[] encodedTxBytes = Encoder.encodeToMsgPack(signedTxn);
            TransactionID id = algodApiInstance.rawTransaction(encodedTxBytes);
            System.out.println("Successfully sent tx with ID: " + id);

            // Wait for transaction confirmation
            waitForConfirmation(id.getTxId());
        } catch (ApiException e) {
            System.err.println("Exception when calling algod#rawTransaction: " + e.getResponseBody());
        }


        //Read the transaction from the blockchain
        try {
            com.algorand.algosdk.algod.client.model.Transaction confirmedTxn =
                    algodApiInstance.transactionInformation(RECEIVER, signedTxn.transactionID);
            System.out.println("Transaction information (with notes): " + confirmedTxn.toString());
            System.out.println("Decoded note: " + new String(confirmedTxn.getNoteb64()));
        } catch (ApiException e) {
            System.err.println("Exception when calling algod#transactionInformation: " + e.getCode());
        }
    }    public static void main(String args[]) throws Exception {
        YourFirstTransaction t = new YourFirstTransaction();
        t.gettingStartedExample();
    }
}