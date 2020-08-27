package testing;

import com.algorand.algosdk.util.Encoder;
import com.algorand.algosdk.account.Account;
import com.algorand.algosdk.crypto.Address;
import com.algorand.algosdk.logic.StateSchema;
import com.algorand.algosdk.v2.client.algod.TealCompile;
import com.algorand.algosdk.v2.client.common.AlgodClient;
import com.algorand.algosdk.v2.client.common.Response;
import com.algorand.algosdk.v2.client.model.Application;
import com.algorand.algosdk.v2.client.model.CompileResponse;
import com.algorand.algosdk.v2.client.model.PendingTransactionResponse;
import com.algorand.algosdk.transaction.SignedTransaction;
import com.algorand.algosdk.transaction.Transaction;
import com.algorand.algosdk.v2.client.model.TransactionParametersResponse;
import com.algorand.algosdk.v2.client.model.TransactionsResponse;

public class StatefulSmartContract {
    public AlgodClient client = null;

    // utility function to connect to a node
    private AlgodClient connectToNetwork() {

        // Initialize an algod client
        final String ALGOD_API_ADDR = "localhost";
        // TODO: final Integer ALGOD_PORT = 4001;
        // TODO: final String ALGOD_API_TOKEN = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
        final Integer ALGOD_PORT = 8080;
        final String ALGOD_API_TOKEN = "f73ee5dac477f8ce7f7ac7599b6e77d5bb0c3a43e52712d151922d35a881fc5c";

        AlgodClient client = (AlgodClient) new AlgodClient(ALGOD_API_ADDR, ALGOD_PORT, ALGOD_API_TOKEN);
        return client;
    }

    // utility function to wait on a transaction to be confirmed
    public void waitForConfirmation(String txID) throws Exception {
        if (client == null)
            this.client = connectToNetwork();
        Long lastRound = client.GetStatus().execute().body().lastRound;
        while (true) {
            try {
                // Check the pending transactions
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

    // helper function to compile program source
    public String compileProgram(AlgodClient client, byte[] programSource) {
        Response<CompileResponse> compileResponse = null;
        try {
            compileResponse = client.TealCompile().source(programSource).execute();
        } catch (Exception e) {
            e.printStackTrace();
        }
        System.out.println(compileResponse.body().result);
        return compileResponse.body().result;
    }

    public int createApp(AlgodClient client, Account creator, String approvalProgramSource, String clearProgramSource, int globalInts, 
            int globalBytes, int localInts, int localBytes) {
        // define sender as creator
        Address sender = creator.getAddress();

        // get node suggested parameters
        try {
            TransactionParametersResponse params = client.TransactionParams().execute().body();
        } catch (Exception e) {
            e.printStackTrace();
        }

        // create unsigned transaction TODO:
        Transaction txn = Transaction.ApplicationCreateTransactionBuilder().
                             .approvalProgram(approvalProgramSource)
                             .clearStateProgram(clearProgramSource)
                             .globalStateSchema(new StateSchema(globalInts, globalBytes))
                             .localStateSchema(new StateSchema(localInts, localBytes));

    
        // // sign transaction
        // SignedTransaction signedTxn = creator.signTransaction(txn);
        // System.out.println("Signed transaction with txid: " + signedTxn.transactionID);

        // // send to network
        // byte[] encodedTxBytes = Encoder.encodeToMsgPack(signedTxn);
        // String id = client.RawTransaction().rawtxn(encodedTxBytes).execute().body().txId;
        // System.out.println("Successfully sent tx with ID: " + id);

        // // await confirmation
        // waitForConfirmation(id);

        // // display results
        // PendingTransactionResponse pTrx = client.PendingTransactionInformation(id).execute().body();
        // Long appId = pTrx.txn.tx.applicationId;
        // System.out.println("Created new app-id: " + appId);    
    
        return 1; //appId;
    }

    public void statefulSmartContract() throws Exception {

        // user declared account mnemonics
        String creatorMnemonic = "Your 25-word mnemonic goes here";
        String userMnemonic = "A second distinct 25-word mnemonic goes here";

        // TODO: REMOVE:
        creatorMnemonic = "domain tomato skate earth donor twice sorry pave grid image review grunt edit news jelly grain act soldier barely daring master soft skate absorb exhaust";
        userMnemonic  = "fury paddle knife situate six enjoy praise sea ketchup present quiz east rail ability scout loud lens iron update daughter food task supply above avoid";

        // declare application state storage (immutable)
        int localInts = 1;
        int localBytes = 1;
        int globalInts = 1;
        int globalBytes = 0;

        // user declared approval program (initial)
        String approvalProgramSourceInitial = "#pragma version 2\n" +
        "///// Handle each possible OnCompletion type. We don't have to worry about\n" +
        "//// handling ClearState, because the ClearStateProgram will execute in that\n" +
        "//// case, not the ApprovalProgram.\n" +

        "txn OnCompletion\n" +
        "int NoOp\n" +
        "==\n" +
        "bnz handle_noop\n" +

        "txn OnCompletion\n" +
        "int OptIn\n" +
        "==\n" +
        "bnz handle_optin\n" +

        "txn OnCompletion\n" +
        "int CloseOut\n" +
        "==\n" +
        "bnz handle_closeout\n" +

        "txn OnCompletion\n" +
        "int UpdateApplication\n" +
        "==\n" +
        "bnz handle_updateapp\n" +

        "txn OnCompletion\n" +
        "int DeleteApplication\n" +
        "==\n" +
        "bnz handle_deleteapp\n" +

        "//// Unexpected OnCompletion value. Should be unreachable.\n" +
        "err\n" +

        "handle_noop:\n" +
        "//// Handle NoOp\n" +
        "//// Check for creator\n" +
        "addr 5XWY6RBNYHCSY2HK5HCTO62DUJJ4PT3G4L77FQEBUKE6ZYRGQAFTLZSQQ4\n" +
        "txn Sender\n" +
        "==\n" +
        "bnz handle_optin\n" +

        "//// read global state\n" +
        "byte \"counter\"\n" +
        "dup\n" +
        "app_global_get\n" +

        "//// increment the value\n" +
        "int 1\n" +
        "+\n" +

        "//// store to scratch space\n" +
        "dup\n" +
        "store 0\n" +

        "//// update global state\n" +
        "app_global_put\n" +

        "//// read local state for sender\n" +
        "int 0\n" +
        "byte \"counter\"\n" +
        "app_local_get\n" +

        "//// increment the value\n" +
        "int 1\n" +
        "+\n" +
        "store 1\n" +

        "//// update local state for sender\n" +
        "int 0\n" +
        "byte \"counter\"\n" +
        "load 1\n" +
        "app_local_put\n" +

        "//// load return value as approval\n" +
        "load 0\n" +
        "return\n" +

        "handle_optin:\n" +
        "//// Handle OptIn\n" +
        "//// approval\n" +
        "int 1\n" +
        "return\n" +

        "handle_closeout:\n" +
        "//// Handle CloseOut\n" +
        "////approval\n" +
        "int 1\n" +
        "return\n" +

        "handle_deleteapp:\n" +
        "//// Check for creator\n" +
        "addr 5XWY6RBNYHCSY2HK5HCTO62DUJJ4PT3G4L77FQEBUKE6ZYRGQAFTLZSQQ4\n" +
        "txn Sender\n" +
        "==\n" +
        "return\n" +

        "handle_updateapp:\n" +
        "//// Check for creator\n" +
        "addr 5XWY6RBNYHCSY2HK5HCTO62DUJJ4PT3G4L77FQEBUKE6ZYRGQAFTLZSQQ4\n" +
        "txn Sender\n" +
        "==\n" +
        "return\n";

        // user declared approval program (refactored)
        String approvalProgramSourceRefactored = "#pragma version 2\n" +
        "//// Handle each possible OnCompletion type. We don't have to worry about\n" +
        "//// handling ClearState, because the ClearStateProgram will execute in that\n" +
        "//// case, not the ApprovalProgram.\n" +

        "txn OnCompletion\n" +
        "int NoOp\n" +
        "==\n" +
        "bnz handle_noop\n" +

        "txn OnCompletion\n" +
        "int OptIn\n" +
        "==\n" +
        "bnz handle_optin\n" +

        "txn OnCompletion\n" +
        "int CloseOut\n" +
        "==\n" +
        "bnz handle_closeout\n" +

        "txn OnCompletion\n" +
        "int UpdateApplication\n" +
        "==\n" +
        "bnz handle_updateapp\n" +

        "txn OnCompletion\n" +
        "int DeleteApplication\n" +
        "==\n" +
        "bnz handle_deleteapp\n" +

        "//// Unexpected OnCompletion value. Should be unreachable.\n" +
        "err\n" +

        "handle_noop:\n" +
        "//// Handle NoOp\n" +
        "//// Check for creator\n" +
        "addr 5XWY6RBNYHCSY2HK5HCTO62DUJJ4PT3G4L77FQEBUKE6ZYRGQAFTLZSQQ4\n" +
        "txn Sender\n" +
        "==\n" +
        "bnz handle_optin\n" +

        "//// read global state\n" +
        "byte \"counter\"\n" +
        "dup\n" +
        "app_global_get\n" +

        "//// increment the value\n" +
        "int 1\n" +
        "+\n" +

        "//// store to scratch space\n" +
        "dup\n" +
        "store 0\n" +

        "//// update global state\n" +
        "app_global_put\n" +

        "//// read local state for sender\n" +
        "int 0\n" +
        "byte \"counter\"\n" +
        "app_local_get\n" +

        "//// increment the value\n" +
        "int 1\n" +
        "+\n" +
        "store 1\n" +

        "//// update local state for sender\n" +
        "//// update \"counter\"\n" +
        "int 0\n" +
        "byte \"counter\"\n" +
        "load 1\n" +
        "app_local_put\n" +

        "//// update \"timestamp\"\n" +
        "int 0\n" +
        "byte \"timestamp\"\n" +
        "txn ApplicationArgs 0\n" +
        "app_local_put\n" +

        "//// load return value as approval\n" +
        "load 0\n" +
        "return\n" +

        "handle_optin:\n" +
        "//// Handle OptIn\n" +
        "//// approval\n" +
        "int 1\n" +
        "return\n" +

        "handle_closeout:\n" +
        "//// Handle CloseOut\n" +
        "////approval\n" +
        "int 1\n" +
        "return\n" +

        "handle_deleteapp:\n" +
        "//// Check for creator\n" +
        "addr 5XWY6RBNYHCSY2HK5HCTO62DUJJ4PT3G4L77FQEBUKE6ZYRGQAFTLZSQQ4\n" +
        "txn Sender\n" +
        "==\n" +
        "return\n" +

        "handle_updateapp:\n" +
        "//// Check for creator\n" +
        "addr 5XWY6RBNYHCSY2HK5HCTO62DUJJ4PT3G4L77FQEBUKE6ZYRGQAFTLZSQQ4\n" +
        "txn Sender\n" +
        "==\n" +
        "return\n";
        
        // declare clear state program source
        String clearProgramSource = "#pragma version 2\n" +
        "int 1\n";

        try {
            // Create an algod client
            if( client == null ) this.client = connectToNetwork();

            // get accounts from mnemonic
            Account creator = new Account(creatorMnemonic);
            Account user = new Account(userMnemonic);
        
            // compile programs NOT IMPLEMENTED IN JAVA SDK
            String approvalProgram = compileProgram(client, approvalProgramSourceInitial.getBytes("UTF-8"));

            // create new application
            int appId = createApp(client, creator, approvalProgramSourceInitial, clearProgramSource, globalInts, globalBytes, localInts, localBytes);
            
            // opt-in to application
        
            // call application without arguments
        
            // read local state of application from user account
        
            // read global state of application
        
            // update application
        
            // call application with arguments
        
            // read local state of application from user account
        
            // close-out from application
        
            // opt-in again to application
        
            // call application with arguments
        
            // read local state of application from user account
        
            // delete application
        
            // clear application from user account

        } catch (Exception e) {
            System.err.println("Exception raised: " + e.getMessage());
        }
    }

    public static void main(String args[]) throws Exception {
        StatefulSmartContract t = new StatefulSmartContract();
        t.statefulSmartContract();
    }
}