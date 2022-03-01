package com.algorand.javatest.multisig;


import java.math.BigInteger;
import java.util.ArrayList;
import java.util.List;

import com.algorand.algosdk.account.Account;
import com.algorand.algosdk.v2.client.common.AlgodClient;
import com.algorand.algosdk.v2.client.common.Response;
import com.algorand.algosdk.v2.client.model.PendingTransactionResponse;
import com.algorand.algosdk.v2.client.model.TransactionParametersResponse;
import com.algorand.algosdk.algod.client.ApiException;
import com.algorand.algosdk.crypto.Address;
import com.algorand.algosdk.crypto.Ed25519PublicKey;
import com.algorand.algosdk.crypto.MultisigAddress;
import com.algorand.algosdk.transaction.SignedTransaction;
import com.algorand.algosdk.transaction.Transaction;
import com.algorand.algosdk.util.Encoder;
import java.util.Scanner;
import org.json.JSONObject;
import com.algorand.algosdk.v2.client.model.PostTransactionsResponse;
import com.algorand.algosdk.v2.client.Utils;

/**
 * Test Multisignature
 *
 */
public class Multisig {

    public AlgodClient client = null;
    
    // utility function to connect to a node
    private AlgodClient connectToNetwork() {

        // Initialize an algod client
        // sandbox
        final String ALGOD_API_ADDR = "localhost";
        final Integer ALGOD_PORT = 4001;
        final String ALGOD_API_TOKEN = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";

        AlgodClient client = new AlgodClient(ALGOD_API_ADDR, ALGOD_PORT, ALGOD_API_TOKEN);
        return client;

    }

 
    
    static Scanner scan = new Scanner(System.in);
    public void multisigExample() throws Exception {
  
        if (client == null)
            this.client = connectToNetwork();
        //  never use mnemonics in production code, replace for demo purposes only
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
      ;
        System.out.println("Multisignature Address: " + msa.toString());
        System.out.println("Navigate to this link and dispense:  https://dispenser.testnet.aws.algodev.network?account=" + msa.toString());            
        System.out.println("PRESS ENTER KEY TO CONTINUE...");     
        scan.nextLine();
 
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
            BigInteger amount = BigInteger.valueOf(100000); // 100000 microAlgos = .1 Algo
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
            // Msgpack encode the signed transaction
            byte[] encodedTxBytes = Encoder.encodeToMsgPack(completeTx);

            // Submit the transaction to the network
            Response < PostTransactionsResponse > rawtxresponse = client.RawTransaction().rawtxn(encodedTxBytes).execute();
            if (!rawtxresponse.isSuccessful()) {
                throw new Exception(rawtxresponse.message());
            }
            String id = rawtxresponse.body().txId;
            // Wait for transaction confirmation
            PendingTransactionResponse pTrx = Utils.waitForConfirmation(client, id, 4);

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
        Multisig t = new Multisig();
        t.multisigExample();
    }
}