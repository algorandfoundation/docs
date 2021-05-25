package com.algorand.javatest.offline;

import java.io.Console;
import java.math.BigInteger;
import java.util.ArrayList;
import java.util.List;

import com.algorand.algosdk.account.Account;
import com.algorand.algosdk.v2.client.common.AlgodClient;
import com.algorand.algosdk.v2.client.common.Response;
import com.algorand.algosdk.v2.client.model.PendingTransactionResponse;
import com.algorand.algosdk.v2.client.model.NodeStatusResponse;
import com.algorand.algosdk.v2.client.model.TransactionParametersResponse;
import com.algorand.algosdk.algod.client.ApiException;
import com.algorand.algosdk.crypto.Address;
import com.algorand.algosdk.crypto.Ed25519PublicKey;
import com.algorand.algosdk.crypto.MultisigAddress;
import com.algorand.algosdk.crypto.MultisigSignature;
import com.algorand.algosdk.transaction.SignedTransaction;
import com.algorand.algosdk.transaction.Transaction;
import com.algorand.algosdk.util.Encoder;
import org.json.JSONObject;
import com.algorand.algosdk.v2.client.model.PostTransactionsResponse;
import java.nio.file.Files;
import java.nio.file.Paths;

/**
 * Test Multisignature
 *
 */
public class MultisigOffline {

    public AlgodClient client = null;

    public static void waitForEnter(String message) {
        Console c = System.console();
        if (c != null) {
            // printf-like arguments
            if (message != null)
                c.format(message);
            c.format("\nPress ENTER to proceed.\n");
            c.readLine();
        }
    }
    
    // utility function to connect to a node
    private AlgodClient connectToNetwork() {

        // Initialize an algod client

        // sandbox
        // Initialize an algod client
        final String ALGOD_API_ADDR = "localhost";
        final Integer ALGOD_PORT = 4001;
        final String ALGOD_API_TOKEN = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";

        AlgodClient client = new AlgodClient(ALGOD_API_ADDR, ALGOD_PORT, ALGOD_API_TOKEN);
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
        Response < NodeStatusResponse > resp = myclient.GetStatus().execute();
        if (!resp.isSuccessful()) {
            throw new Exception(resp.message());
        }
        NodeStatusResponse nodeStatusResponse = resp.body();
        Long startRound = nodeStatusResponse.lastRound + 1;
        Long currentRound = startRound;
        while (currentRound < (startRound + timeout)) {
            // Check the pending transactions                 
            Response < PendingTransactionResponse > resp2 = myclient.PendingTransactionInformation(txID).execute();
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
            resp = myclient.WaitForBlock(currentRound).execute();
            if (!resp.isSuccessful()) {
                throw new Exception(resp.message());
            }
            currentRound++;
        }
        throw new Exception("Transaction not confirmed after " + timeout + " rounds!");
    }
    

    public void writeMultisigSignedTransaction() throws Exception {
        if (client == null)
            this.client = connectToNetwork();

        final String account1_mnemonic = "predict mandate aware dizzy limit match hazard fantasy victory auto fortune hello public dragon ostrich happy blue spray parrot island odor actress only ability hurry";
        final String account2_mnemonic = "moon grid random garlic effort faculty fence gym write skin they joke govern home huge there claw skin way bid fit bean damp able only";
        final String account3_mnemonic = "mirror zone together remind rural impose balcony position minimum quick manage climb quit draft lion device pluck rug siege robust spirit fine luggage ability actual";

        // final String account1_mnemonic = <var>your-25-word-mnemonic</var>
        // final String account2_mnemonic = <var>your-25-word-mnemonic</var>
        // final String account3_mnemonic = <var>your-25-word-mnemonic</var>

        Account act1 = new Account(account1_mnemonic);
        Account act2 = new Account(account2_mnemonic);
        Account act3 = new Account(account3_mnemonic);
        System.out.println("Account1: " + act1.getAddress());
        System.out.println("Account2: " + act2.getAddress());
        System.out.println("Account3: " + act3.getAddress());

        final String DEST_ADDR = act3.getAddress().toString();

        // List for Pks for multisig account
        List<Ed25519PublicKey> publicKeys = new ArrayList<>();
        publicKeys.add(act1.getEd25519PublicKey());
        publicKeys.add(act2.getEd25519PublicKey());
        publicKeys.add(act3.getEd25519PublicKey());

        // Instantiate the Multisig Account
        MultisigAddress msa = new MultisigAddress(1, 2, publicKeys);

        System.out.println("Multisignature Address: " + msa.toString());
        // waitForEnter("Use TestNet Dispenser to add funds, wait for the transaction to be finalized and press enter");

        // setup transaction   
        try {
            Response < TransactionParametersResponse > resp = client.TransactionParams().execute();
            if (!resp.isSuccessful()) {
                throw new Exception(resp.message());
            }
            TransactionParametersResponse params = resp.body();
            if (params == null) {
                throw new Exception("Params retrieval error");
            }                      
            BigInteger amount = BigInteger.valueOf(1000000); // microAlgos
            // add some notes to the transaction
            byte[] notes = "These are some notes encoded in some way!".getBytes();
            // Setup Transaction
            Address sender = new Address(msa.toString());

            Transaction tx = Transaction.PaymentTransactionBuilder()
                    .sender(sender)
                    .amount(amount)
                    .receiver(DEST_ADDR)
                    .note(notes)
                    .suggestedParams(params).build();
            // Sign the Transaction for two accounts
            SignedTransaction signedTx = act1.signMultisigTransaction(msa, tx);
            SignedTransaction completeTx = act2.appendMultisigTransaction(msa, signedTx);
           
           // save signed transaction to a file
           Files.write(Paths.get("./signed.txn"), Encoder.encodeToMsgPack(completeTx));

        } catch (ApiException e) {
            // This is generally expected, but should give us an informative error message.
            System.err.println("Exception when calling algod#rawTransaction: " + e.getResponseBody());
        }
    }
    public void readMultisigSignedTransaction() throws Exception {
        if (client == null)
            this.client = connectToNetwork();


        // setup transaction   
        try {
             // Read the transaction from a file
             SignedTransaction decodedSignedTransaction = Encoder
             .decodeFromMsgPack(Files.readAllBytes(Paths.get("./signed.txn")), SignedTransaction.class);
            //System.out.println("Signed transaction with txid: " + decodedSignedTransaction.transactionID);
    
            
            
            // Msgpack encode the signed transaction
            byte[] encodedTxBytes = Encoder.encodeToMsgPack(decodedSignedTransaction);

            // Submit the transaction to the network
            Response < PostTransactionsResponse > rawtxresponse = client.RawTransaction().rawtxn(encodedTxBytes).execute();
            if (!rawtxresponse.isSuccessful()) {
                throw new Exception(rawtxresponse.message());
            }
            String id = rawtxresponse.body().txId;
            System.out.println("Successfully sent tx with ID: " + id);
            // Wait for transaction confirmation
            PendingTransactionResponse pTrx = waitForConfirmation(client, id, 4);

            System.out.println("Transaction " + id + " confirmed in round " + pTrx.confirmedRound);
            // Read the transaction
            JSONObject jsonObj = new JSONObject(pTrx.toString());
            System.out.println("Transaction information (with notes): " + jsonObj.toString(2));
            System.out.println("Decoded note: " + new String(pTrx.txn.tx.note));


        } catch (ApiException e) {
            // This is generally expected, but should give us an informative error message.
            System.err.println("Exception when calling algod#rawTransaction: " + e.getResponseBody());
        }
    } 

    public void writeMultisigUnsignedTransaction() throws Exception {
        if (client == null)
            this.client = connectToNetwork();

        final String account1_mnemonic = "predict mandate aware dizzy limit match hazard fantasy victory auto fortune hello public dragon ostrich happy blue spray parrot island odor actress only ability hurry";
        final String account2_mnemonic = "moon grid random garlic effort faculty fence gym write skin they joke govern home huge there claw skin way bid fit bean damp able only";
        final String account3_mnemonic = "mirror zone together remind rural impose balcony position minimum quick manage climb quit draft lion device pluck rug siege robust spirit fine luggage ability actual";

        // final String account1_mnemonic = <var>your-25-word-mnemonic</var>
        // final String account2_mnemonic = <var>your-25-word-mnemonic</var>
        // final String account3_mnemonic = <var>your-25-word-mnemonic</var>

        Account act1 = new Account(account1_mnemonic);
        Account act2 = new Account(account2_mnemonic);
        Account act3 = new Account(account3_mnemonic);
        System.out.println("Account1: " + act1.getAddress());
        System.out.println("Account2: " + act2.getAddress());
        System.out.println("Account3: " + act3.getAddress());

        final String DEST_ADDR = act3.getAddress().toString();

        // List for Pks for multisig account
        List<Ed25519PublicKey> publicKeys = new ArrayList<>();
        publicKeys.add(act1.getEd25519PublicKey());
        publicKeys.add(act2.getEd25519PublicKey());
        publicKeys.add(act3.getEd25519PublicKey());

        // Instantiate the Multisig Account
        MultisigAddress msa = new MultisigAddress(1, 2, publicKeys);

        System.out.println("Multisignature Address: " + msa.toString());

        // waitForEnter("Use TestNet Dispenser to add funds, wait for the transaction to be finalized and press enter");

        // setup transaction   
        try {
            Response < TransactionParametersResponse > resp = client.TransactionParams().execute();
            if (!resp.isSuccessful()) {
                throw new Exception(resp.message());
            }
            TransactionParametersResponse params = resp.body();
            if (params == null) {
                throw new Exception("Params retrieval error");
            }                      
            BigInteger amount = BigInteger.valueOf(1000000); // microAlgos
            // add some notes to the transaction
            byte[] notes = "These are some notes encoded in some way!".getBytes();
            // Setup Transaction
            Address sender = new Address(msa.toString());

            Transaction tx = Transaction.PaymentTransactionBuilder()
                    .sender(sender)
                    .amount(amount)
                    .receiver(DEST_ADDR)
                    .note(notes)
                    .suggestedParams(params).build();

            // NOTE: save as signed even though it has not been
         
            MultisigSignature mSig = new MultisigSignature();
            SignedTransaction stx = new SignedTransaction(tx,mSig,tx.txID());
            stx.tx = tx;

            // Save transaction to a file
            Files.write(Paths.get("./unsigned.txn"), Encoder.encodeToMsgPack(stx));
            System.out.println("Transaction written to a file");
            // Save multisig to a file
            Files.write(Paths.get("./unsigned.msig"), Encoder.encodeToMsgPack(msa));


            // // Sign the Transaction for two accounts
            // SignedTransaction signedTx = act1.signMultisigTransaction(msa, tx);
            // SignedTransaction completeTx = act2.appendMultisigTransaction(msa, signedTx);
            // // Msgpack encode the signed transaction
            // byte[] encodedTxBytes = Encoder.encodeToMsgPack(completeTx);

            // // Submit the transaction to the network
            // Response < PostTransactionsResponse > rawtxresponse = client.RawTransaction().rawtxn(encodedTxBytes).execute();
            // if (!rawtxresponse.isSuccessful()) {
            //     throw new Exception(rawtxresponse.message());
            // }
            // String id = rawtxresponse.body().txId;
            // System.out.println("Successfully sent tx with ID: " + id);
            // // Wait for transaction confirmation
            // PendingTransactionResponse pTrx = waitForConfirmation(client, id, 4);

            // System.out.println("Transaction " + id + " confirmed in round " + pTrx.confirmedRound);
            // // Read the transaction
            // JSONObject jsonObj = new JSONObject(pTrx.toString());
            // System.out.println("Transaction information (with notes): " + jsonObj.toString(2));
            // System.out.println("Decoded note: " + new String(pTrx.txn.tx.note));


        } catch (ApiException e) {
            // This is generally expected, but should give us an informative error message.
            System.err.println("Exception when calling algod#rawTransaction: " + e.getResponseBody());
        }
    }
    public void readMultisigUnsignedTransaction() throws Exception {
        if (client == null)
            this.client = connectToNetwork();

        final String account1_mnemonic = "predict mandate aware dizzy limit match hazard fantasy victory auto fortune hello public dragon ostrich happy blue spray parrot island odor actress only ability hurry";
        final String account2_mnemonic = "moon grid random garlic effort faculty fence gym write skin they joke govern home huge there claw skin way bid fit bean damp able only";
        final String account3_mnemonic = "mirror zone together remind rural impose balcony position minimum quick manage climb quit draft lion device pluck rug siege robust spirit fine luggage ability actual";

        // final String account1_mnemonic = <var>your-25-word-mnemonic</var>
        // final String account2_mnemonic = <var>your-25-word-mnemonic</var>
        // final String account3_mnemonic = <var>your-25-word-mnemonic</var>

        Account act1 = new Account(account1_mnemonic);
        Account act2 = new Account(account2_mnemonic);
        Account act3 = new Account(account3_mnemonic);
        System.out.println("Account1: " + act1.getAddress());
        System.out.println("Account2: " + act2.getAddress());
        System.out.println("Account3: " + act3.getAddress());

        //final String DEST_ADDR = act3.getAddress().toString();

        // read transaction from file
        SignedTransaction decodedTransaction = Encoder
        .decodeFromMsgPack(Files.readAllBytes(Paths.get("./unsigned.txn")), SignedTransaction.class);
        System.out.println("transaction txid: " + decodedTransaction.tx.txID());

        MultisigAddress decodedMultisig = Encoder
        .decodeFromMsgPack(Files.readAllBytes(Paths.get("./unsigned.msig")),
        MultisigAddress.class);
        System.out.println("Multisig Address: " + decodedMultisig.toString());

        MultisigAddress msig = decodedMultisig;

        System.out.println("Multisig Address: " + msig.toString());

        Transaction tx1 = decodedTransaction.tx;
   
        try {

            SignedTransaction signedTx = act1.signMultisigTransaction(msig, tx1);
            SignedTransaction completeTx = act2.appendMultisigTransaction(msig, signedTx);
            // Msgpack encode the signed transaction
            byte[] encodedTxBytes = Encoder.encodeToMsgPack(completeTx);

            // Submit the transaction to the network
            Response < PostTransactionsResponse > rawtxresponse = client.RawTransaction().rawtxn(encodedTxBytes).execute();
            if (!rawtxresponse.isSuccessful()) {
                throw new Exception(rawtxresponse.message());
            }
            String id = rawtxresponse.body().txId;
            System.out.println("Successfully sent tx with ID: " + id);
            // Wait for transaction confirmation
            PendingTransactionResponse pTrx = waitForConfirmation(client, id, 4);

            System.out.println("Transaction " + id + " confirmed in round " + pTrx.confirmedRound);
            // Read the transaction
            JSONObject jsonObj = new JSONObject(pTrx.toString());
            System.out.println("Transaction information (with notes): " + jsonObj.toString(2));
            System.out.println("Decoded note: " + new String(pTrx.txn.tx.note));


        } catch (ApiException e) {
            // This is generally expected, but should give us an informative error message.
            System.err.println("Exception when calling algod#rawTransaction: " + e.getResponseBody());
        }
    } 

    public static void main(String args[]) throws Exception {
        MultisigOffline t = new MultisigOffline();
       // t.writeMultisigUnsignedTransaction();
       // t.readMultisigUnsignedTransaction();

        t.writeMultisigSignedTransaction();
        t.readMultisigSignedTransaction();

    }
}