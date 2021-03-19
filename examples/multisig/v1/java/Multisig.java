package com.algorand.javatest.multisig;

import java.io.Console;
import java.io.IOException;
import java.math.BigInteger;
import java.util.ArrayList;
import java.util.List;

import com.algorand.algosdk.account.Account;
import com.algorand.algosdk.algod.client.AlgodClient;
import com.algorand.algosdk.algod.client.ApiException;
import com.algorand.algosdk.algod.client.api.AlgodApi;
import com.algorand.algosdk.algod.client.auth.ApiKeyAuth;
import com.algorand.algosdk.algod.client.model.TransactionID;
import com.algorand.algosdk.algod.client.model.TransactionParams;
import com.algorand.algosdk.crypto.Address;
import com.algorand.algosdk.crypto.Digest;
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
    public AlgodApi algodApiInstance = null;

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
    private AlgodApi connectToNetwork() {

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
        final String ALGOD_API_ADDR = "http://localhost:4001";
        final String ALGOD_API_TOKEN = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
 
        AlgodClient client = (AlgodClient) new AlgodClient().setBasePath(ALGOD_API_ADDR);
        ApiKeyAuth api_key = (ApiKeyAuth) client.getAuthentication("api_key");
        api_key.setApiKey(ALGOD_API_TOKEN);
        algodApiInstance = new AlgodApi(client);
        return algodApiInstance;
    }

    // utility function to wait on a transaction to be confirmed
    public void waitForConfirmation(String txID) throws Exception {
        if (algodApiInstance == null)
            connectToNetwork();
        while (true) {
            try {
                // Check the pending tranactions
                com.algorand.algosdk.algod.client.model.Transaction pendingInfo = algodApiInstance
                        .pendingTransactionInformation(txID);
                if (pendingInfo.getRound() != null && pendingInfo.getRound().longValue() > 0) {
                    // Got the completed Transaction
                    System.out.println("Transaction " + pendingInfo.getTx() + " confirmed in round "
                            + pendingInfo.getRound().longValue());
                    break;
                }
                algodApiInstance
                        .waitForBlock(BigInteger.valueOf(algodApiInstance.getStatus().getLastRound().longValue() + 1));
            } catch (Exception e) {
                throw (e);
            }
        }
    }

    public void multisigExample() throws Exception {


        if (algodApiInstance == null)
            connectToNetwork();
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
        waitForEnter("Use TestNet Dispenser to add funds, wait for the transaction to be finalized and press enter");
        // setup transaction   
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
        // add some notes to the transaction
        byte[] notes = "These are some notes encoded in some way!".getBytes();
        // Setup Transaction
        Address sender = new Address(msa.toString());
          
        Transaction tx = Transaction.PaymentTransactionBuilder().sender(sender).amount(amount)
                .receiver(DEST_ADDR).fee(fee).firstValid(firstValidRound).lastValid(lastValidRound)
                .genesisHash(genesisHash).note(notes).genesisID(genesisID).build();
        // Sign the Transaction for two accounts
        SignedTransaction signedTx = act1.signMultisigTransaction(msa, tx);
        SignedTransaction completeTx = act2.appendMultisigTransaction(msa, signedTx);

        // send the transaction to the network
        try {
            // Msgpack encode the signed transaction
            byte[] encodedTxBytes = Encoder.encodeToMsgPack(completeTx);
            TransactionID id = algodApiInstance.rawTransaction(encodedTxBytes);
            // Wait for transaction confirmation
            waitForConfirmation(id.getTxId());         
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