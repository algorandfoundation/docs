package com.algorand.javatest.smart_contracts;

import com.algorand.algosdk.v2.client.common.AlgodClient;
import java.nio.file.Files;
import java.nio.file.Paths;
import com.algorand.algosdk.v2.client.model.CompileResponse;

public class CompileTeal {
    // Utility function to update changing block parameters
    public AlgodClient client = null;

    // utility function to connect to a node
    private AlgodClient connectToNetwork() {

        // Initialize an algod client
        // final Integer ALGOD_PORT = <algod-port>;
        // final String ALGOD_API_ADDR = "<algod-address>";
        // final String ALGOD_API_TOKEN = "<algod-token>";
        final Integer ALGOD_PORT = 4001;
        final String ALGOD_API_ADDR = "localhost";
        final String ALGOD_API_TOKEN = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";

        AlgodClient client = new AlgodClient(ALGOD_API_ADDR, ALGOD_PORT, ALGOD_API_TOKEN);
        return client;
    }

    public void compileTealSource() throws Exception {
        // Initialize an algod client
        if (client == null)
            this.client = connectToNetwork();
 
        // read file - int 0
        byte[] data = Files.readAllBytes(Paths.get("./sample.teal"));
        // byte[] data = Files.readAllBytes(Paths.get("<./filename>"));
 
        // compile
        CompileResponse response = client.TealCompile().source(data).execute().body();
        // print results
        System.out.println("response: " + response);
        System.out.println("Hash: " + response.hash); 
        System.out.println("Result: " + response.result); 
    }

    public static void main(final String args[]) throws Exception {
        CompileTeal t = new CompileTeal();
        t.compileTealSource();
    }

}
// Output should look similar to this... 
// response:
// {"hash":"KI4DJG2OOFJGUERJGSWCYGFZWDNEU2KWTU56VRJHITP62PLJ5VYMBFDBFE","result":"ASABACI="}
// Hash: KI4DJG2OOFJGUERJGSWCYGFZWDNEU2KWTU56VRJHITP62PLJ5VYMBFDBFE 
// Result: ASABACI=

// resource
// https://developer.algorand.org/docs/features/asc1/sdks/#account-delegation-sdk-usage
