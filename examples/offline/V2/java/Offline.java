package com.algorand.javatest.offline;

import com.algorand.algosdk.account.Account;
import com.algorand.algosdk.crypto.Address;
import com.algorand.algosdk.transaction.SignedTransaction;
import com.algorand.algosdk.transaction.Transaction;
import com.algorand.algosdk.util.Encoder;
import com.algorand.algosdk.v2.client.common.AlgodClient;
import com.algorand.algosdk.v2.client.common.Response;
import com.algorand.algosdk.v2.client.model.NodeStatusResponse;
import com.algorand.algosdk.v2.client.model.PendingTransactionResponse;
import com.algorand.algosdk.v2.client.model.PostTransactionsResponse;
import com.algorand.algosdk.v2.client.model.TransactionParametersResponse;
import org.json.JSONObject;
import java.nio.file.Files;
import java.nio.file.Paths;
import com.algorand.algosdk.v2.client.Utils;


public class Offline {
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
        return client;
    }
   


    public void writeUnsignedTransaction() throws Exception {

        if (client == null)
            this.client = connectToNetwork();
        // Import your private key mnemonic and address
        // Shown for demonstration purposes. NEVER reveal secret mnemonics in practice.
        final String PASSPHRASE = "<var>your-25-word-mnemonic</var>";
        com.algorand.algosdk.account.Account myAccount = new Account(PASSPHRASE);
        System.out.println("My Address: " + myAccount.getAddress());

        String myAddress = printBalance(myAccount);

        try {
            // Construct the transaction
            final String RECEIVER = "L5EUPCF4ROKNZMAE37R5FY2T5DF2M3NVYLPKSGWTUKVJRUGIW4RKVPNPD4";
            String note = "Hello World";
            Response < TransactionParametersResponse > resp = client.TransactionParams().execute();
            if (!resp.isSuccessful()) {
                throw new Exception(resp.message());
            }
            TransactionParametersResponse params = resp.body();
            if (params == null) {
                throw new Exception("Params retrieval error");
            }
            System.out.println("Algorand suggested parameters: " + params);
            Transaction tx = Transaction.PaymentTransactionBuilder()
                .sender(myAddress)
                .note(note.getBytes())
                .amount(100000)
                .receiver(new Address(RECEIVER))
                .suggestedParams(params)
                .build();

            // save as signed even though it has not been
            SignedTransaction stx = myAccount.signTransaction(tx);
            System.out.println("Signed transaction with txid: " + stx.transactionID);
            stx.tx = tx;  
            // Save transaction to a file 
            Files.write(Paths.get("./unsigned.txn"), Encoder.encodeToMsgPack(stx));
            System.out.println("Transaction written to a file");
        } catch (Exception e) {
            System.err.println("Exception when calling algod#transactionInformation: " + e.getMessage());
        }
    }
    public void readUnsignedTransaction() throws Exception {

        if (client == null)
            this.client = connectToNetwork();
        // Import your private key mnemonic and address
        // Shown for demonstration purposes. NEVER reveal secret mnemonics in practice.
        final String PASSPHRASE = "<var>your-25-word-mnemonic</var>";
        com.algorand.algosdk.account.Account myAccount = new Account(PASSPHRASE);
        System.out.println("My Address: " + myAccount.getAddress());

        try {

            // read transaction from file
            SignedTransaction decodedTransaction = Encoder.decodeFromMsgPack(
                Files.readAllBytes(Paths.get("./unsigned.txn")), SignedTransaction.class);            
            Transaction tx = decodedTransaction.tx;           
        

            // Sign the transaction
            SignedTransaction signedTxn = myAccount.signTransaction(tx);
 
            // Submit the transaction to the network
            byte[] encodedTxBytes = Encoder.encodeToMsgPack(signedTxn);
            Response < PostTransactionsResponse > rawtxresponse = client.RawTransaction().rawtxn(encodedTxBytes).execute();
            if (!rawtxresponse.isSuccessful()) {
                throw new Exception(rawtxresponse.message());
            }
            String id = rawtxresponse.body().txId;
            System.out.println("Successfully sent tx with ID: " + id);

            // Wait for transaction confirmation
            PendingTransactionResponse pTrx = Utils.waitForConfirmation(client, id, 4);

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
    public void writeSignedTransaction() throws Exception {

        if (client == null)
            this.client = connectToNetwork();
        // Import your private key mnemonic and address
        // Shown for demonstration purposes. NEVER reveal secret mnemonics in practice.
        final String PASSPHRASE = "<var>your-25-word-mnemonic</var>";
        com.algorand.algosdk.account.Account myAccount = new Account(PASSPHRASE);
        System.out.println("My Address: " + myAccount.getAddress());

        String myAddress = printBalance(myAccount);

        try {
            // Construct the transaction
            final String RECEIVER = "L5EUPCF4ROKNZMAE37R5FY2T5DF2M3NVYLPKSGWTUKVJRUGIW4RKVPNPD4";
            String note = "Hello World";
            Response < TransactionParametersResponse > resp = client.TransactionParams().execute();
            if (!resp.isSuccessful()) {
                throw new Exception(resp.message());
            }
            TransactionParametersResponse params = resp.body();
            if (params == null) {
                throw new Exception("Params retrieval error");
            }
            System.out.println("Algorand suggested parameters: " + params);
            Transaction txn = Transaction.PaymentTransactionBuilder()
                .sender(myAddress)
                .note(note.getBytes())
                .amount(100000)
                .receiver(new Address(RECEIVER))
                .suggestedParams(params)
                .build();

            // Sign the transaction
            SignedTransaction signedTx = myAccount.signTransaction(txn);
            System.out.println("Signed transaction with txid: " + signedTx.transactionID);
            // save signed transaction to  a file 
            Files.write(Paths.get("./signed.txn"), Encoder.encodeToMsgPack(signedTx));

        } catch (Exception e) {
            System.err.println("Exception when calling algod#transactionInformation: " + e.getMessage());
        }
    }
    public void readSignedTransaction() throws Exception {

        if (client == null)
            this.client = connectToNetwork();
        // Import your private key mnemonic and address
        // Shown for demonstration purposes. NEVER reveal secret mnemonics in practice.
        final String PASSPHRASE = "<var>your-25-word-mnemonic</var>";
        com.algorand.algosdk.account.Account myAccount = new Account(PASSPHRASE);
        System.out.println("My Address: " + myAccount.getAddress());
        // read signed transaction
        try {
            SignedTransaction decodedSignedTransaction = Encoder.decodeFromMsgPack(
                Files.readAllBytes(Paths.get("./signed.txn")), SignedTransaction.class);     

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
            PendingTransactionResponse pTrx = Utils.waitForConfirmation(client, id, 4);

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
        Response < com.algorand.algosdk.v2.client.model.Account > respAcct = client.AccountInformation(myAccount.getAddress()).execute();
        if (!respAcct.isSuccessful()) {
            throw new Exception(respAcct.message());
        }
        com.algorand.algosdk.v2.client.model.Account accountInfo = respAcct.body();
        System.out.println(String.format("Account Balance: %d microAlgos", accountInfo.amount));
        return myAddress;
    }

    public static void main(String args[]) throws Exception {
        Offline t = new Offline();
        t.writeUnsignedTransaction();
        t.readUnsignedTransaction();

        // t.writeSignedTransaction();
        // t.readSignedTransaction();

    }
}