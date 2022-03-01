package com.algorand.javatest.assets;

import java.math.BigInteger;
import com.algorand.algosdk.v2.client.common.AlgodClient;
import com.algorand.algosdk.account.Account;
import org.json.JSONArray;
import org.json.JSONObject;
import com.algorand.algosdk.v2.client.common.*;
import com.algorand.algosdk.algod.client.ApiException;
import com.algorand.algosdk.crypto.Address;
import com.algorand.algosdk.transaction.SignedTransaction;
import com.algorand.algosdk.transaction.Transaction;
import com.algorand.algosdk.util.Encoder;
import com.algorand.algosdk.v2.client.Utils;
import com.algorand.algosdk.v2.client.model.PendingTransactionResponse;
import com.algorand.algosdk.v2.client.model.PostTransactionsResponse;
import com.algorand.algosdk.v2.client.model.TransactionParametersResponse;

// Show Creating, modifying, sending and listing assets

public class AssetExample {

    public AlgodClient client = null;

    // utility function to connect to a node
    private AlgodClient connectToNetwork() {

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
        if (jsonArray.length() > 0) {
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
    }

    // utility function to print asset holding
    public void printAssetHolding(Account account, Long assetID) throws Exception {
        if (client == null)
            this.client = connectToNetwork();
        String accountInfo = client.AccountInformation(account.getAddress()).execute().toString();
        JSONObject jsonObj = new JSONObject(accountInfo.toString());
        JSONArray jsonArray = (JSONArray) jsonObj.get("assets");
        if (jsonArray.length() > 0) {
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
    }




    // Utility function for sending a raw signed transaction to the network
    public String submitTransaction(SignedTransaction signedTx) throws Exception {
        try {
            // Msgpack encode the signed transaction
            byte[] encodedTxBytes = Encoder.encodeToMsgPack(signedTx);
            String[] headers = {"Content-Type"};
            String[] values = {"application/x-binary"};
            Response < PostTransactionsResponse > rawtxresponse = 
                client.RawTransaction().rawtxn(encodedTxBytes).execute(headers, values);
            if (!rawtxresponse.isSuccessful()) {
                throw new Exception(rawtxresponse.message());
            }
            String id = rawtxresponse.body().txId;

            // String id = client.RawTransaction().rawtxn(encodedTxBytes).execute().body().txId;
            // ;
            return (id);
        } catch (ApiException e) {
            throw (e);
        }
    }

    public void assetExample() throws Exception {
        if (client == null)
            this.client = connectToNetwork();
        // recover example accounts
	    // Shown for demonstration purposes. NEVER reveal secret mnemonics in practice in code.

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

        // CREATE ASSET
        // get changing network parameters for each transaction
        Response < TransactionParametersResponse > resp = client.TransactionParams().execute();
        if (!resp.isSuccessful()) {
            throw new Exception(resp.message());
        }
        TransactionParametersResponse params = resp.body();
        if (params == null) {
            throw new Exception("Params retrieval error");
        }
        // params.fee = (long) 1000;        
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
        Transaction tx = Transaction.AssetCreateTransactionBuilder()
                .sender(acct1.getAddress())
                .assetTotal(assetTotal)
                .assetDecimals(decimals)
                .assetUnitName(unitName)
                .assetName(assetName).url(url)
                .metadataHashUTF8(assetMetadataHash)
                .manager(manager)
                .reserve(reserve)
                .freeze(freeze)
                .defaultFrozen(defaultFrozen)
                .clawback(clawback)
                .suggestedParams(params).build();

        // Sign the Transaction with creator account
        SignedTransaction signedTx = acct1.signTransaction(tx);
        Long assetID = null;
        try {
            String id = submitTransaction(signedTx);
            System.out.println("Transaction ID: " + id);
            PendingTransactionResponse pTrx = Utils.waitForConfirmation(client,id,4);          
            System.out.println("Transaction " + id + " confirmed in round " + pTrx.confirmedRound);
            // Now that the transaction is confirmed we can get the assetID
            assetID = pTrx.assetIndex;
            System.out.println("AssetID = " + assetID);
            printCreatedAsset(acct1, assetID);
            printAssetHolding(acct1, assetID);

        } catch (Exception e) {
            e.printStackTrace();
            return;
        }

        // CHANGE MANAGER
        // Change Asset Configuration:
        // assetID = Long.valueOf((your asset id));
        // get changing network parameters for each transaction
        resp = client.TransactionParams().execute();
        if (!resp.isSuccessful()) {
            throw new Exception(resp.message());
        }
        params = resp.body();
        if (params == null) {
            throw new Exception("Params retrieval error");
        }
  
        // params.fee = (long) 1000;
        // configuration changes must be done by
        // the manager account - changing manager of the asset

        tx = Transaction.AssetConfigureTransactionBuilder()
                .sender(acct2.getAddress())
                .assetIndex(assetID)
                .manager(acct1.getAddress())
                .reserve(reserve)
                .freeze(freeze)
                .clawback(clawback)
                .suggestedParams(params)
                .build();

        // the transaction must be signed by the current manager account
        signedTx = acct2.signTransaction(tx);
        // send the transaction to the network
        try {
            String id = submitTransaction(signedTx);
            System.out.println("Transaction ID: " + id);
            PendingTransactionResponse pTrx = Utils.waitForConfirmation(client,id,4);          
            System.out.println("Transaction " + id + " confirmed in round " + pTrx.confirmedRound);

            // the manager should now be the same as the creator
            System.out.println("AssetID = " + assetID);
            printCreatedAsset(acct1, assetID);

        } catch (Exception e) {
            e.printStackTrace();
            return;
        }


        // OPT-IN
        // Opt in to Receiving the Asset
        // assetID = Long.valueOf((your asset id));
        // get changing network parameters for each transaction
        resp = client.TransactionParams().execute();
        if (!resp.isSuccessful()) {
            throw new Exception(resp.message());
        }
        params = resp.body();
        if (params == null) {
            throw new Exception("Params retrieval error");
        }
        // params.fee = (long) 1000;
        // configuration changes must be done by
        // the manager account - changing manager of the asset
        tx = Transaction.AssetAcceptTransactionBuilder()
                .acceptingAccount(acct3.getAddress())
                .assetIndex(assetID)
                .suggestedParams(params)
                .build();
        // The transaction must be signed by the current manager account
        signedTx = acct3.signTransaction(tx);
        // send the transaction to the network and
        try {
            String id = submitTransaction(signedTx);
            System.out.println("Transaction ID: " + id);
            PendingTransactionResponse pTrx = Utils.waitForConfirmation(client,id,4);          
            System.out.println("Transaction " + id + " confirmed in round " + pTrx.confirmedRound);

            // We can now list the account information for acct3
            // and see that it can accept the new asset
            System.out.println("Account 3 = " + acct3.getAddress().toString());
            printAssetHolding(acct3, assetID);
        } catch (Exception e) {
            e.printStackTrace();
            return;
        }

        // TRANSFER ASSET
        // Transfer the Asset:
        // assetID = Long.valueOf((your asset id));
        // get changing network parameters for each transaction
        resp = client.TransactionParams().execute();
        if (!resp.isSuccessful()) {
            throw new Exception(resp.message());
        }
        params = resp.body();
        if (params == null) {
            throw new Exception("Params retrieval error");
        }
        // params.fee = (long) 1000;
        // set asset xfer specific parameters
        BigInteger assetAmount = BigInteger.valueOf(10);
        Address sender = acct1.getAddress();
        Address receiver = acct3.getAddress();
        tx = Transaction.AssetTransferTransactionBuilder()
                .sender(sender)
                .assetReceiver(receiver)
                .assetAmount(assetAmount)
                .assetIndex(assetID)
                .suggestedParams(params)
                .build();
        // The transaction must be signed by the sender account
        signedTx = acct1.signTransaction(tx);
        // send the transaction to the network
        try {
            String id = submitTransaction(signedTx);
            System.out.println("Transaction ID: " + id);
            PendingTransactionResponse pTrx = Utils.waitForConfirmation(client,id,4);          
            System.out.println("Transaction " + id + " confirmed in round " + pTrx.confirmedRound);

            // list the account information for acct1 and acct3
            System.out.println("Account 3  = " + acct3.getAddress().toString());
            printAssetHolding(acct3, assetID);
            System.out.println("Account 1  = " + acct1.getAddress().toString());
            printAssetHolding(acct1, assetID);
        } catch (Exception e) {
            e.printStackTrace();
            return;
        }

        // FREEZE
        // Freeze the Asset:
        // assetID = Long.valueOf((your asset id));
        // get changing network parameters for each transaction
        resp = client.TransactionParams().execute();
        if (!resp.isSuccessful()) {
            throw new Exception(resp.message());
        }
        params = resp.body();
        if (params == null) {
            throw new Exception("Params retrieval error");
        }
        // params.fee = (long) 1000;
        // The asset was created and configured to allow freezing an account
        // set asset specific parameters
        boolean freezeState = true;
        // The sender should be freeze account
        tx = Transaction.AssetFreezeTransactionBuilder()
                .sender(acct2.getAddress())
                .freezeTarget(acct3.getAddress())
                .freezeState(freezeState)
                .assetIndex(assetID)
                .suggestedParams(params)
                .build();
        // The transaction must be signed by the freeze account
        signedTx = acct2.signTransaction(tx);
        // send the transaction to the network
        try {
            String id = submitTransaction(signedTx);
            System.out.println("Transaction ID: " + id);
            PendingTransactionResponse pTrx = Utils.waitForConfirmation(client,id,4);          
            System.out.println("Transaction " + id + " confirmed in round " + pTrx.confirmedRound);
            System.out.println("Account 3 = " + acct3.getAddress().toString());
            printAssetHolding(acct3, assetID);

        } catch (Exception e) {
            e.printStackTrace();
            return;
        }

        // REVOKE (or clawback)
        // Revoke the asset:
        // The asset was also created with the ability for it to be revoked by
        // clawbackaddress.
        // assetID = Long.valueOf((your asset id));
        // get changing network parameters for each transaction
        resp = client.TransactionParams().execute();
        if (!resp.isSuccessful()) {
            throw new Exception(resp.message());
        }
        params = resp.body();
        if (params == null) {
            throw new Exception("Params retrieval error");
        }
        // params.fee = (long) 1000;

        // set asset specific parameters
        assetAmount = BigInteger.valueOf(10);
        tx = Transaction.AssetClawbackTransactionBuilder()
                .sender(acct2.getAddress())
                .assetClawbackFrom(acct3.getAddress())
                .assetReceiver(acct1.getAddress())
                .assetAmount(assetAmount)
                .assetIndex(assetID)
                .suggestedParams(params)
                .build();
        // The transaction must be signed by the clawback account
        signedTx = acct2.signTransaction(tx);
        // send the transaction to the network and
        // wait for the transaction to be confirmed
        try {
            String id = submitTransaction(signedTx);
            System.out.println("Transaction ID: " + id);
            PendingTransactionResponse pTrx = Utils.waitForConfirmation(client,id,4);          
            System.out.println("Transaction " + id + " confirmed in round " + pTrx.confirmedRound);

            // list the account information for acct1 and acct3
            System.out.println("Account 3  = " + acct3.getAddress().toString());
            printAssetHolding(acct3, assetID);
            System.out.println("Account 1  = " + acct1.getAddress().toString());
            printAssetHolding(acct1, assetID);
        } catch (Exception e) {
            e.printStackTrace();
            return;
        }

        // DESTROY

        // Destroy the Asset:
        // All assets should now be back in
        // creators account
        // assetID = Long.valueOf((your asset id));
        // get changing network parameters for each transaction
        resp = client.TransactionParams().execute();
        if (!resp.isSuccessful()) {
            throw new Exception(resp.message());
        }
        params = resp.body();
        if (params == null) {
            throw new Exception("Params retrieval error");
        }
        // params.fee = (long) 1000;

        // set destroy asset specific parameters
        // The manager must sign and submit the transaction
        tx = Transaction.AssetDestroyTransactionBuilder()
                .sender(acct1.getAddress())
                .assetIndex(assetID)
                .suggestedParams(params)
                .build();
        // The transaction must be signed by the manager account
        signedTx = acct1.signTransaction(tx);
        // send the transaction to the network
        try {
            String id = submitTransaction(signedTx);
            System.out.println("Transaction ID: " + id);
            PendingTransactionResponse pTrx = Utils.waitForConfirmation(client,id,4);          
            System.out.println("Transaction " + id + " confirmed in round " + pTrx.confirmedRound);

            // We list the account information for acct1
            // and check that the asset is no longer exist
            System.out.println("Account 3 must do a transaction for an amount of 0, ");
            System.out.println("with a assetCloseTo to the creator account, to clear it from its accountholdings");
            System.out.println("Account 1  = " + acct1.getAddress().toString());            
            System.out.println("Nothing should print after this, Account 1 asset is sucessfully deleted");
            printAssetHolding(acct1, assetID);
            printCreatedAsset(acct1, assetID);
        } catch (Exception e) {
            e.printStackTrace();
            return;
        }
    }

    public static void main(String args[]) throws Exception {

        AssetExample ex = new AssetExample();
        ex.assetExample();
    }
}
