package com.algorand.javatest.assets.v2;

import java.math.BigInteger;

import com.algorand.algosdk.v2.client.common.AlgodClient;
import com.algorand.algosdk.account.Account;
import com.algorand.algosdk.v2.client.model.*;
import org.json.JSONArray;
import org.json.JSONObject;
import com.algorand.algosdk.v2.client.common.*;
import com.algorand.algosdk.algod.client.ApiException;
import com.algorand.algosdk.crypto.Address;
import com.algorand.algosdk.transaction.SignedTransaction;
import com.algorand.algosdk.transaction.Transaction;
import com.algorand.algosdk.util.Encoder;

// Show Creating, modifying, sending and listing assets

public class CreateAsset {

    public AlgodClient client = null;

    // utility function to connect to a node
    private AlgodClient connectToNetwork() {

        // "f1dee49e36a82face92fdb21cd3d340a1b369925cd12f3ee7371378f1665b9b1";
        // final String ALGOD_API_ADDR = "<var>algod-address</var>";
        // final String ALGOD_API_TOKEN = "<var>algod-token</var>";
        final String ALGOD_API_ADDR = "localhost";
        final int ALGOD_PORT = 4001;
        final String ALGOD_API_TOKEN = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
        AlgodClient client = new AlgodClient(ALGOD_API_ADDR, ALGOD_PORT, ALGOD_API_TOKEN);
        return client;
    }

    // utility function to print created asset
    public void printCreatedAsset(Account account, Long assetID) throws Exception {
        if (client == null)
            this.client = connectToNetwork();
        String accountInfo = client.AccountInformation(account.getAddress()).execute().toString();
        JSONObject jsonObj = new JSONObject(accountInfo.toString());
        JSONArray jsonArray = (JSONArray) jsonObj.get("created-assets");
        if (jsonArray.length() > 0)
            try { 
                    for (Object o : jsonArray) {
                        JSONObject ca = (JSONObject) o;
                        Integer myassetIDInt = (Integer) ca.get("index");
                         if (assetID.longValue() == myassetIDInt.longValue()) {
                            System.out.println("Created Asset Info: " + ca.toString(2)); // pretty print
                            break;
                        }
                    }               
            } catch (Exception e) {
                throw (e);
            }
    }

    // utility function to print asset holding
    public void printAssetHolding(Account account, Long assetID) throws Exception {
        if (client == null)
            this.client = connectToNetwork();
        String accountInfo = client.AccountInformation(account.getAddress()).execute().toString();
        JSONObject jsonObj = new JSONObject(accountInfo.toString());
        JSONArray jsonArray = (JSONArray) jsonObj.get("assets");
        if (jsonArray.length() > 0)
            try {
                for (Object o : jsonArray) {
                    JSONObject ca = (JSONObject) o;
                    Integer myassetIDInt = (Integer) ca.get("asset-id");
                    if (assetID.longValue() == myassetIDInt.longValue()) {
                        System.out.println("Asset Holding Info: " + ca.toString(2)); // pretty print
                        break;
                    }
                }
            } catch (Exception e) {
                throw (e);
            }
    }

    // utility function to wait on a transaction to be confirmed
    public void waitForConfirmation(String txID) throws Exception {
        if (client == null)
            this.client = connectToNetwork();
        Long lastRound = client.GetStatus().execute().body().lastRound;
        while (true) {
            try {
                // Check the pending tranactions
                Response<PendingTransactionResponse> pendingInfo = client.PendingTransactionInformation(txID).execute();
                if (pendingInfo.body().confirmedRound != null && pendingInfo.body().confirmedRound > 0) {
                    // Got the completed Transaction
                    System.out.println(
                            "Transaction " + txID + " confirmed in round " + pendingInfo.body().confirmedRound);
                    break;
                }
                lastRound++;
                client.WaitForBlock(lastRound).execute();
            } catch (Exception e) {
                throw (e);
            }
        }
    }

    // Utility function for sending a raw signed transaction to the network
    public String submitTransaction(SignedTransaction signedTx) throws Exception {
        try {
            // Msgpack encode the signed transaction
            byte[] encodedTxBytes = Encoder.encodeToMsgPack(signedTx);
            String id = client.RawTransaction().rawtxn(encodedTxBytes).execute().body().txId;
            ;
            return (id);
        } catch (ApiException e) {
            throw (e);
        }
    }

    public void createAsset() throws Exception {
        if (client == null)
            this.client = connectToNetwork();
        // recover example accounts

        final String account1_mnemonic = "year crumble opinion local grid injury rug happy away castle minimum bitter upon romance federal entire rookie net fabric soft comic trouble business above talent";
        final String account2_mnemonic = "beauty nurse season autumn curve slice cry strategy frozen spy panic hobby strong goose employ review love fee pride enlist friend enroll clip ability runway";
        final String account3_mnemonic = "picnic bright know ticket purity pluck stumble destroy ugly tuna luggage quote frame loan wealth edge carpet drift cinnamon resemble shrimp grain dynamic absorb edge";

        // final String account1_mnemonic = <var>your-25-word-mnemonic</var>
        // final String account2_mnemonic = <var>your-25-word-mnemonic</var>
        // final String account3_mnemonic = <var>your-25-word-mnemonic</var>

        Account acct1 = new Account(account1_mnemonic);
        Account acct2 = new Account(account2_mnemonic);
        Account acct3 = new Account(account3_mnemonic);
        System.out.println("Account1: " + acct1.getAddress());
        System.out.println("Account2: " + acct2.getAddress());
        System.out.println("Account3: " + acct3.getAddress());

        // get changing network parameters for each transaction
        TransactionParametersResponse params = client.TransactionParams().execute().body();
        params.fee = (long) 1000;

        // Create the Asset:
        BigInteger assetTotal = BigInteger.valueOf(10000);
        boolean defaultFrozen = false;
        String unitName = "myunit";
        String assetName = "my longer asset name";
        String url = "http://this.test.com";
        String assetMetadataHash = "16efaa3924a6fd9d3a4824799a4ac65d";
        Address manager = acct2.getAddress();
        Address reserve = acct2.getAddress();
        Address freeze = acct2.getAddress();
        Address clawback = acct2.getAddress();
        Integer decimals = 0;
        Transaction tx = Transaction.AssetCreateTransactionBuilder().sender(acct1.getAddress()).assetTotal(assetTotal)
                .assetDecimals(decimals).assetUnitName(unitName).assetName(assetName).url(url)
                .metadataHashUTF8(assetMetadataHash).manager(manager).reserve(reserve).freeze(freeze)
                .defaultFrozen(defaultFrozen).clawback(clawback).suggestedParams(params).build();

        // Sign the Transaction with creator account
        SignedTransaction signedTx = acct1.signTransaction(tx);
        Long assetID = null;
        try {
            String id = submitTransaction(signedTx);
            System.out.println("Transaction ID: " + id);
            waitForConfirmation(id);
            // Read the transaction
            PendingTransactionResponse pTrx = client.PendingTransactionInformation(id).execute().body();
            // Now that the transaction is confirmed we can get the assetID
            assetID = pTrx.assetIndex;
            System.out.println("AssetID = " + assetID);         
            printCreatedAsset(acct1, assetID);
            printAssetHolding(acct1, assetID);
        } catch (Exception e) {
            e.printStackTrace();
            return;
        }




    }

    public static void main(String args[]) throws Exception {

        CreateAsset ex = new CreateAsset();
        ex.createAsset();

    }
}
// resource https://developer.algorand.org/docs/features/asa/
