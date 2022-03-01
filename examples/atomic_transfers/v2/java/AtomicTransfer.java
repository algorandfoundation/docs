package com.algorand.javatest.atomictransfer;


import com.algorand.algosdk.v2.client.model.PostTransactionsResponse;
import java.io.ByteArrayOutputStream;

import com.algorand.algosdk.account.Account;
import com.algorand.algosdk.v2.client.common.AlgodClient;
import com.algorand.algosdk.v2.client.common.Response;
import com.algorand.algosdk.v2.client.model.PendingTransactionResponse;
import com.algorand.algosdk.v2.client.model.TransactionParametersResponse;
import com.algorand.algosdk.crypto.Digest;
import com.algorand.algosdk.transaction.SignedTransaction;
import com.algorand.algosdk.transaction.Transaction;
import com.algorand.algosdk.transaction.TxGroup;
import com.algorand.algosdk.util.Encoder;
import org.json.JSONObject;
import com.algorand.algosdk.v2.client.Utils;


public class AtomicTransfer {

    public AlgodClient client = null;

    // utility function to connect to a node
    private AlgodClient connectToNetwork() {

        // Initialize an algod client
        final String ALGOD_API_ADDR = "localhost";
        final Integer ALGOD_PORT = 4001;
        final String ALGOD_API_TOKEN = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";

        AlgodClient client = new AlgodClient(ALGOD_API_ADDR, ALGOD_PORT, ALGOD_API_TOKEN);
        return client;
    }


    public void AtomicTransfer() throws Exception {

        if (client == null)
            this.client = connectToNetwork();

        // final String account1_mnemonic = "Your 25-word mnemonic goes here";
        // final String account2_mnemonic = "Your 25-word mnemonic goes here";
        // final String account3_mnemonic = "Your 25-word mnemonic goes here";
        // never use mnemonics in code, for demo purposes only

        // recover account A, B, C
        Account acctA = new Account(account1_mnemonic);
        Account acctB = new Account(account2_mnemonic);
        Account acctC = new Account(account3_mnemonic);
        System.out.println("AccountA: " + acctA.getAddress());
        System.out.println("AccountB: " + acctB.getAddress());
        System.out.println("AccountC: " + acctC.getAddress());
        // get node suggested parameters
        TransactionParametersResponse params = client.TransactionParams().execute().body();

        // Create the first transaction
        Transaction tx1 = Transaction.PaymentTransactionBuilder()
                .sender(acctA.getAddress())
                .amount(10000)
                .receiver(acctC.getAddress())
                .suggestedParams(params)
                .build();

        // Create the second transaction
        Transaction tx2 = Transaction.PaymentTransactionBuilder()
                .sender(acctB.getAddress())
                .amount(20000)
                .receiver(acctA.getAddress())
                .suggestedParams(params)
                .build();
        // group transactions an assign ids
        Digest gid = TxGroup.computeGroupID(new Transaction[] { tx1, tx2 });
        tx1.assignGroupID(gid);
        tx2.assignGroupID(gid);

        // sign individual transactions
        SignedTransaction signedTx1 = acctA.signTransaction(tx1);
        ;
        SignedTransaction signedTx2 = acctB.signTransaction(tx2);
        ;

        try {
            // put both transaction in a byte array
            ByteArrayOutputStream byteOutputStream = new ByteArrayOutputStream();
            byte[] encodedTxBytes1 = Encoder.encodeToMsgPack(signedTx1);
            byte[] encodedTxBytes2 = Encoder.encodeToMsgPack(signedTx2);
            byteOutputStream.write(encodedTxBytes1);
            byteOutputStream.write(encodedTxBytes2);
            byte groupTransactionBytes[] = byteOutputStream.toByteArray();

            // send transaction group
            Response<PostTransactionsResponse> response = client.RawTransaction().rawtxn(groupTransactionBytes).execute();
            if (!response.isSuccessful()){
                throw new Exception(response.message());
            }
            String id = response.body().txId;
            // wait for confirmation
                        // Wait for transaction confirmation
            PendingTransactionResponse pTrx = Utils.waitForConfirmation(client, id, 4);
            System.out.println("Transaction " + id + " confirmed in round " + pTrx.confirmedRound);
            // Read the transaction
            JSONObject jsonObj = new JSONObject(pTrx.toString());
            System.out.println("Transaction information (with notes): " + jsonObj.toString(2));

        } catch (Exception e) {
            System.out.println("Submit Exception: " + e);
        }
    }

    public static void main(String args[]) throws Exception {
        AtomicTransfer mn = new AtomicTransfer();
        mn.AtomicTransfer();
    }
}