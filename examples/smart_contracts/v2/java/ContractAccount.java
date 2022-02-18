package com.algorand.javatest.smart_contracts;

import com.algorand.algosdk.account.Account;
import com.algorand.algosdk.algod.client.ApiException;

import com.algorand.algosdk.crypto.Address;
import com.algorand.algosdk.crypto.LogicsigSignature;
import com.algorand.algosdk.transaction.SignedTransaction;
import com.algorand.algosdk.transaction.Transaction;
import com.algorand.algosdk.util.Encoder;
import com.algorand.algosdk.v2.client.common.AlgodClient;
import com.algorand.algosdk.v2.client.common.Response;
import com.algorand.algosdk.v2.client.model.PendingTransactionResponse;
import com.algorand.algosdk.v2.client.model.TransactionParametersResponse;
import java.nio.file.Files;
import java.nio.file.Paths;
import org.json.JSONObject;
import java.util.ArrayList;
import java.util.Base64;
import com.algorand.algosdk.v2.client.model.CompileResponse;
import com.algorand.algosdk.v2.client.Utils;

public class ContractAccount {
    // Utility function to update changing block parameters
    public AlgodClient client = null;

    // utility function to connect to a node
    private AlgodClient connectToNetwork() {

        // Initialize an algod client
        final String ALGOD_API_ADDR = "localhost";
        final Integer ALGOD_PORT = 4001;
        final String ALGOD_API_TOKEN = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
        // final String ALGOD_API_ADDR = "<algod-address>";
        // final Integer ALGOD_PORT = <algod-port>;
        // final String ALGOD_API_TOKEN = "<algod-token>";

        AlgodClient client = new AlgodClient(ALGOD_API_ADDR, ALGOD_PORT, ALGOD_API_TOKEN);
        return client;
    }

    public void contractAccountExample() throws Exception {
        // Initialize an algod client
        if (client == null)
            this.client = connectToNetwork();
        // Set the receiver
        // final String RECEIVER = "<receiver-address>";
        final String RECEIVER = "QUDVUXBX4Q3Y2H5K2AG3QWEOMY374WO62YNJFFGUTMOJ7FB74CMBKY6LPQ";
        // Read program from file samplearg.teal
        // This code is meant for learning purposes only
        // It should not be used in production
        // arg_0
        // btoi
        // int 123
        // ==
        byte[] source = Files.readAllBytes(Paths.get("./samplearg.teal"));   
        //byte[] source = Files.readAllBytes(Paths.get("./<filename">));
        // compile
        Response < CompileResponse > compileresponse = client.TealCompile().source(source).execute();
        if (!compileresponse.isSuccessful()) {
            throw new Exception(compileresponse.message());
        }
        CompileResponse response = compileresponse.body();
        // print results
        System.out.println("response: " + response);
        System.out.println("Hash: " + response.hash);
        System.out.println("Result: " + response.result);
        byte[] program = Base64.getDecoder().decode(response.result.toString());

        // create logic sig

        // string parameter
        // ArrayList<byte[]> teal_args = new ArrayList<byte[]>();
        // String orig = "<my string>";
        // teal_args.add(orig.getBytes());
        // LogicsigSignature lsig = new LogicsigSignature(program, teal_args);

        // integer parameter
        ArrayList<byte[]> teal_args = new ArrayList<byte[]>();
        byte[] arg1 = { 123 };
        teal_args.add(arg1);
        LogicsigSignature lsig = new LogicsigSignature(program, teal_args);
        // For no args use null as second param
        // LogicsigSignature lsig = new LogicsigSignature(program, null);     
        TransactionParametersResponse params = client.TransactionParams().execute().body();
        // create a transaction
        String note = "Hello World";
        Transaction txn = Transaction.PaymentTransactionBuilder()
                .sender(lsig
                .toAddress())
                .note(note.getBytes())
                .amount(100000)
                .receiver(new Address(RECEIVER))
                .suggestedParams(params)
                .build();   
        try {
            // create the LogicSigTransaction with contract account LogicSig
            SignedTransaction stx = Account.signLogicsigTransaction(lsig, txn);
            // send raw LogicSigTransaction to network
            byte[] encodedTxBytes = Encoder.encodeToMsgPack(stx);
            // logic signature transaction can be written to a file
            // try {
            //     String FILEPATH = "./simple.stxn";
            //     File file = new File(FILEPATH);
            //     OutputStream os = new FileOutputStream(file);
            //     os.write(encodedTxBytes);
            //     os.close();
            // } catch (Exception e) {
            //     System.out.println("Exception: " + e);
            // }
            String id = client.RawTransaction().rawtxn(encodedTxBytes).execute().body().txId;
            // Wait for transaction confirmation
            PendingTransactionResponse pTrx = Utils.waitForConfirmation(client,id,4);          
            System.out.println("Transaction " + id + " confirmed in round " + pTrx.confirmedRound);

            JSONObject jsonObj = new JSONObject(pTrx.toString());
            System.out.println("Transaction information (with notes): " + jsonObj.toString(2)); // pretty print
            System.out.println("Decoded note: " + new String(pTrx.txn.tx.note));
        } catch (ApiException e) {
            System.err.println("Exception when calling algod#rawTransaction: " + e.getResponseBody());
        }
    }
    public static void main(final String args[]) throws Exception {
        ContractAccount t = new ContractAccount();
        t.contractAccountExample();
    }
}

