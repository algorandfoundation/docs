package com.algorand.javatest.multisig.v2;

import java.io.Console;
import java.io.IOException;
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

/**
 * Test Multisignature
 *
 */
public class Multisig {

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

        // Algorand Hackathon
        // final String ALGOD_API_ADDR = "http://hackathon.algodev.network:9100";
        // final String ALGOD_API_TOKEN =
        // "ef920e2e7e002953f4b29a8af720efe8e4ecc75ff102b165e0472834b25832c1";

        // your own node
        // final String ALGOD_API_ADDR = "http://localhost:8080";
        // final String ALGOD_API_TOKEN = "your ALGOD_API_TOKEN";

        // Purestake
        // final String ALGOD_API_ADDR =
        // "https://testnet-algorand.api.purestake.io/ps1";
        // final String ALGOD_API_TOKEN = "B3SU4KcVKi94Jap2VXkK83xx38bsv95K5UZm2lab";

        // Initialize an algod client
        // final String ALGOD_API_ADDR = "algod-address<PLACEHOLDER>";
        // final String ALGOD_API_TOKEN = "algod-token<PLACEHOLDER>";
        // sandbox
        // Initialize an algod client
        final String ALGOD_API_ADDR = "localhost";
        final Integer ALGOD_PORT = 4001;
        final String ALGOD_API_TOKEN = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";

        AlgodClient client = new AlgodClient(ALGOD_API_ADDR, ALGOD_PORT, ALGOD_API_TOKEN);
        return client;

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
    

    public void multisigExample() throws Exception {
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

        // Instantiate the the Multisig Accout
        MultisigAddress msa = new MultisigAddress(1, 2, publicKeys);

        System.out.println("Multisignature Address: " + msa.toString());
        // waitForEnter("Use TestNet Dispenser to add funds, wait for the transaction to be finalized and press enter");

        // setup transaction   
        try {
            TransactionParametersResponse params = client.TransactionParams().execute().body();            
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
            // Msgpack encode the signed transaction
            byte[] encodedTxBytes = Encoder.encodeToMsgPack(completeTx);
            // send the transaction to the network           
            String id = client.RawTransaction().rawtxn(encodedTxBytes).execute().body().txId;
            // Wait for transaction confirmation
            waitForConfirmation(id);         
            System.out.println("Successfully sent tx with id: " + id);
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