package com.algorand.javatest.smart_contracts.v1;

import java.math.BigInteger;
import java.util.ArrayList;

import com.algorand.algosdk.account.Account;
import com.algorand.algosdk.algod.client.AlgodClient;
import com.algorand.algosdk.algod.client.ApiException;
import com.algorand.algosdk.algod.client.api.AlgodApi;
import com.algorand.algosdk.algod.client.auth.ApiKeyAuth;
import com.algorand.algosdk.algod.client.model.TransactionID;
import com.algorand.algosdk.algod.client.model.TransactionParams;
import com.algorand.algosdk.crypto.Address;
import com.algorand.algosdk.crypto.Digest;
import com.algorand.algosdk.crypto.LogicsigSignature;
import com.algorand.algosdk.transaction.SignedTransaction;
import com.algorand.algosdk.transaction.Transaction;
import com.algorand.algosdk.util.Encoder;

public class AccountDeligation {
    // Utility function to update changing block parameters

    public static void main(final String args[]) throws Exception {
        // Initialize an algod client
        // your own node
        // final String ALGOD_API_ADDR = "http://127.0.0.1:8080";
        // final String ALGOD_API_TOKEN = "your token in your node/data/algod.token
        // file";

        // sandbox
        final String ALGOD_API_ADDR = "http://localhost:4001";
        final String ALGOD_API_TOKEN = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";

        // Purestake
        // final String ALGOD_API_ADDR =
        // "https://testnet-algorand.api.purestake.io/ps1";
        // final String ALGOD_API_TOKEN = "B3SU4KcVKi94Jap2VXkK83xx38bsv95K5UZm2lab";


        //Create an instance of the algod API client
        AlgodClient client = (AlgodClient) new AlgodClient().setBasePath(ALGOD_API_ADDR);
        ApiKeyAuth api_key = (ApiKeyAuth) client.getAuthentication("api_key");
        api_key.setApiKey(ALGOD_API_TOKEN);
        AlgodApi algodApiInstance = new AlgodApi(client);

        // import your private key mnemonic and address
        final String SRC_ACCOUNT = "buzz genre work meat fame favorite rookie stay tennis demand panic busy hedgehog snow morning acquire ball grain grape member blur armor foil ability seminar";
        Account src = new Account(SRC_ACCOUNT);
        // Set the receiver
        final String DEST_ADDR = "QUDVUXBX4Q3Y2H5K2AG3QWEOMY374WO62YNJFFGUTMOJ7FB74CMBKY6LPQ";

        // get suggested parameters
        BigInteger suggestedFeePerByte = BigInteger.valueOf(1);
        BigInteger firstRound = BigInteger.valueOf(301);
        String genId = null;
        Digest genesisHash = null;
        try {
            // Get suggested parameters from the node
            TransactionParams params = algodApiInstance.transactionParams();
            suggestedFeePerByte = params.getFee();
            firstRound = params.getLastRound();
            genId = params.getGenesisID();
            genesisHash = new Digest(params.getGenesishashb64());

        } catch (ApiException e) {
            System.err.println("Exception when calling algod#transactionParams");
            e.printStackTrace();
        }


        // create logic sig int 0
        // hex example 0x01, 0x20, 0x01, 0x00, 0x22
        // byte[] program = {
        // hex-encoded-program<PLACEHOLDER>
        // };
        byte[] program = {
            0x01, 0x20, 0x01, 0x00, 0x22  // int 0, returns false, so rawTransaction will fail below
        };
        LogicsigSignature lsig = new LogicsigSignature(program, null);        
       
        // sign the logic signature with an account sk
        src.signLogicsig(lsig);

        // create a transaction
        BigInteger amount = BigInteger.valueOf(1000000);
        BigInteger lastRound = firstRound.add(BigInteger.valueOf(1000)); 
        Transaction tx = new Transaction(src.getAddress(), new Address(DEST_ADDR), BigInteger.valueOf(1000), amount, firstRound, lastRound, genId, genesisHash);
        
        
        try {
            // create the LogicSigTransaction with contract account LogicSig
            SignedTransaction stx = Account.signLogicsigTransaction(lsig, tx);

            // send raw LogicSigTransaction to network
            byte[] encodedTxBytes = Encoder.encodeToMsgPack(stx);
            TransactionID id = algodApiInstance.rawTransaction(encodedTxBytes);
            System.out.println("Successfully sent tx with id: " + id);

        } catch (ApiException e) {
            System.err.println("Exception when calling algod#rawTransaction: " + e.getResponseBody());
        }

    }
}

// resource
// https://developer.algorand.org/docs/features/asc1/sdks/#account-delegation-sdk-usage
