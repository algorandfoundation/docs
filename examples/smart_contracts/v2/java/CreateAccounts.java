package com.algorand.javatest.smart_contracts.v2;
import com.algorand.algosdk.account.Account;

public class CreateAccounts
 {
     public static void main(String args[]) throws Exception {
     try {
        Account myAccount1 = new Account();
        System.out.println("My Address1: " + myAccount1.getAddress());
        System.out.println("My Passphrase1: " + myAccount1.toMnemonic());
        Account myAccount2 = new Account();
        System.out.println("My Address2: " + myAccount2.getAddress());
        System.out.println("My Passphrase2: " + myAccount2.toMnemonic());
        Account myAccount3 = new Account();
        System.out.println("My Address3: " + myAccount3.getAddress());
        System.out.println("My Passphrase3: " + myAccount3.toMnemonic());

        //Copy off accounts and mnemonics    
        //Dispense TestNet Algos to each account: https://bank.testnet.algorand.network/
        // resource:
        // https://developer.algorand.org/docs/features/accounts/create/#standalone
    } catch (Exception e) {
        e.printStackTrace();
        return;
    }
    }
}