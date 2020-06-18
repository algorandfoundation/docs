    package com.algorand.YourFirstTransaction;

    import com.algorand.algosdk.account.Account;
    import com.algorand.algosdk.crypto.Address;
    import com.algorand.algosdk.transaction.SignedTransaction;
    import com.algorand.algosdk.transaction.Transaction;
    import com.algorand.algosdk.util.Encoder;
    import com.algorand.algosdk.v2.client.common.AlgodClient;
    import com.algorand.algosdk.v2.client.common.Response;
    import com.algorand.algosdk.v2.client.model.PendingTransactionResponse;
    import com.algorand.algosdk.v2.client.model.TransactionParametersResponse;
    public class YourFirstTransaction {
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
        // utility function to wait on a transaction to be confirmed    
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

        public void gettingStartedExample() throws Exception {

            if( client == null ) this.client = connectToNetwork();

            // Import your private key mnemonic and address
            final String PASSPHRASE = "Your 25-word mnemonic generated and displayed above";
            com.algorand.algosdk.account.Account myAccount = new Account(PASSPHRASE);
            System.out.println("My Address: " + myAccount.getAddress());

            String myAddress = myAccount.getAddress().toString();

            com.algorand.algosdk.v2.client.model.Account accountInfo = client.AccountInformation(myAccount.getAddress()).execute().body();

            System.out.println(String.format("Account Balance: %d microAlgos", accountInfo.amount));

            try {
                // Construct the transaction
                final String RECEIVER = "GD64YIY3TWGDMCNPP553DZPPR6LDUSFQOIJVFDPPXWEG3FVOJCCDBBHU5A";
                String note = "Hello World";
                TransactionParametersResponse params = client.TransactionParams().execute().body();
                Transaction txn = Transaction.PaymentTransactionBuilder()
                .sender(myAddress)
                .note(note.getBytes())
                .amount(100000)
                .receiver(new Address(RECEIVER))
                .suggestedParams(params)
                .build();


                // Sign the transaction
                SignedTransaction signedTxn = myAccount.signTransaction(txn);
                System.out.println("Signed transaction with txid: " + signedTxn.transactionID);

                // Submit the transaction to the network
                byte[] encodedTxBytes = Encoder.encodeToMsgPack(signedTxn);
                String id = client.RawTransaction().rawtxn(encodedTxBytes).execute().body().txId;
                System.out.println("Successfully sent tx with ID: " + id);

                // Wait for transaction confirmation
                waitForConfirmation(id);

                //Read the transaction
                PendingTransactionResponse pTrx = client.PendingTransactionInformation(id).execute().body();
                System.out.println("Transaction information (with notes): " + pTrx.toString());
                System.out.println("Decoded note: " + new String(pTrx.txn.tx.note));


            } catch (Exception e) {
                System.err.println("Exception when calling algod#transactionInformation: " + e.getMessage());
            }
        }

        public static void main(String args[]) throws Exception {
            YourFirstTransaction t = new YourFirstTransaction();
            t.gettingStartedExample();
        }
    }