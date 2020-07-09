package com.algorand.algosdk.GroupedTransaction;

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



public class GroupedTransaction {   

    public AlgodClient client = null;

    // utility function to connect to a node
    private AlgodClient connectToNetwork(){

        // Initialize an algod client
        final String ALGOD_API_ADDR = "localhost";
        final Integer ALGOD_PORT = 4001;
        final String ALGOD_API_TOKEN = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";

        AlgodClient client = (AlgodClient) new AlgodClient(ALGOD_API_ADDR, ALGOD_PORT, ALGOD_API_TOKEN);
        return client;
    }

    public void waitForConfirmation( String txID ) throws Exception{
        if( client == null ) this.client = connectToNetwork();
        Long lastRound = client.GetStatus().execute().body().lastRound;
        while(true) {
            try {
                //Check the pending tranactions
                Response<PendingTransactionResponse> pendingInfo = client.PendingTransactionInformation(txID).execute();
                if (pendingInfo.body().confirmedRound != null && pendingInfo.body().confirmedRound > 0) {
                    //Got the completed Transaction
                    System.out.println("Transaction " + txID + " confirmed in round " + pendingInfo.body().confirmedRound);
                    break;
                } 
                lastRound++;
                client.WaitForBlock(lastRound).execute();
            } catch (Exception e) {
                throw( e );
            }
        }
    }

    public void AtomicTransfer() throws Exception {

        if( client == null ) this.client = connectToNetwork();

        final String account1_mnemonic = "Your 25-word mnemonic goes here";
        final String account2_mnemonic = "Your 25-word mnemonic goes here";
        final String account3_mnemonic = "Your 25-word mnemonic goes here";

        // recover account A, B, C
        Account acctA  = new Account(account1_mnemonic); 
        Account acctB  = new Account(account2_mnemonic);
        Account acctC  = new Account(account3_mnemonic); 

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
        Digest gid = TxGroup.computeGroupID(new Transaction[]{tx1, tx2});
        tx1.assignGroupID(gid);
        tx2.assignGroupID(gid);

        // sign individual transactions
        SignedTransaction signedTx1 = acctA.signTransaction(tx1);;
        SignedTransaction signedTx2 = acctB.signTransaction(tx2);;

        try {
            // put both transaction in a byte array 
            ByteArrayOutputStream byteOutputStream = new ByteArrayOutputStream( );
            byte[] encodedTxBytes1 = Encoder.encodeToMsgPack(signedTx1);
            byte[] encodedTxBytes2 = Encoder.encodeToMsgPack(signedTx2);
            byteOutputStream.write(encodedTxBytes1);
            byteOutputStream.write(encodedTxBytes2);
            byte groupTransactionBytes[] = byteOutputStream.toByteArray();

            // send transaction group
            String id = client.RawTransaction().rawtxn(groupTransactionBytes).execute().body().txId;
            System.out.println("Successfully sent tx with ID: " + id);

            // wait for confirmation
            waitForConfirmation(id);

        } catch (Exception e) {
            System.out.println("Submit Exception: " + e); 
        }
    }
    public static void main(String args[]) throws Exception {
        GroupedTransaction mn = new GroupedTransaction();
        mn.AtomicTransfer();
    }
}