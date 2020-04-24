title: Register Online

This section assumes that you have already [generated a participation key](generate_keys.md) for the account you plan to mark online. 

Registering an account online requires authorizing a [key registration transaction](../../features/transactions/index.md#key-registration-transaction) with details of the participation key that will vote on the account's behalf. Once the transaction is processed by the blockchain, the Verifiable Random Function public key (referred to as the VRF public the key) is written into the account’s data and the account will start participating in consensus with that key. This VRF public key is how the account is associated with the specific participation keys.

!!! info "Important"
	The moment a key registration transaction is confirmed by the network it takes 320 rounds for the change to take effect. In other words, if a key registration is confirmed in round 1000, the account will not start participating until round 1320.

# Create an online key registration transaction

Create a key registration transaction for the address: `EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4` by inserting the following code snippet into the construction portion of the example shown in [Authorizing Transactions Offline](../../features/transactions/offline_transactions.md#unsigned-transaction-file-operations). The file produced and displayed with `goal clerk inspect` should look almost exactly the same as the output shown in the [constructing a register online transaction example](../../features/transactions/index.md#register-account-online). 


```python tab="Python"
from algosdk import encoding
...
def write_unsigned():
	# setup connectionon
    algod_client = connect_to_network()

    # get suggested parameters
    params = algod_client.suggested_params()

    b64votekey = "eXq34wzh2UIxCZaI1leALKyAvSz/+XOe0wqdHagM+bw="
    votekey_addr = encoding.encode_address(base64.b64decode(b64votekey))
    b64selkey = "X84ReKTmp+yfgmMCbbokVqeFFFrKQeFZKEXG89SXwm4="
    selkey_addr = encoding.encode_address(base64.b64decode(b64selkey))

    # create transaction
    data = {
        "sender": "EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4",
        "votekey": votekey_addr,
        "selkey": selkey_addr,
        "votefst": 6000000,
        "votelst":9000000,
        "votekd": 1730,
        "fee": 2000,
        "flat_fee": True,
        "first": 6002000,
        "last": 6003000,
        "gen": params.get('genesisID'),
        "gh": params.get('genesishashb64')
    }
    txn = transaction.KeyregTxn(**data)
...
```

```java tab="Java"
...
import java.util.Base64;
...
import com.algorand.algosdk.crypto.ParticipationPublicKey;
import com.algorand.algosdk.crypto.VRFPublicKey;
...
    public void writeUnsignedTransaction(){

        // connect to node
        if( algodApiInstance == null ) connectToNetwork();

        final String SRC_ADDR = "EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4";
        String selKeyEncoded = new String("X84ReKTmp+yfgmMCbbokVqeFFFrKQeFZKEXG89SXwm4=");
        byte[] decodedSelKey = Base64.getDecoder().decode(selKeyEncoded);
        String voteKeyEncoded = new String("eXq34wzh2UIxCZaI1leALKyAvSz/+XOe0wqdHagM+bw=");
        byte[] decodedVoteKey = Base64.getDecoder().decode(voteKeyEncoded);

        try {
            // Get suggested parameters from the node
            TransactionParams params = algodApiInstance.transactionParams();

            // create transaction
            String genId = params.getGenesisID();
            Digest genesisHash = new Digest(params.getGenesishashb64());
            ParticipationPublicKey voteKey = new ParticipationPublicKey(decodedSelKey);
            VRFPublicKey selKey = new VRFPublicKey(decodedVoteKey);
            BigInteger voteFst = BigInteger.valueOf(6000000);
            BigInteger voteLst = BigInteger.valueOf(9000000);
            BigInteger firstRound = BigInteger.valueOf(6002000);
            BigInteger lastRound = BigInteger.valueOf(6003000);
            BigInteger fee = BigInteger.valueOf(2000);
            BigInteger voteKd = BigInteger.valueOf(1730);
            Transaction tx = new Transaction(new Address(SRC_ADDR), fee, firstRound, lastRound,
                    null, genId, genesisHash, voteKey, selKey,  voteFst, voteLst, voteKd);
```

```go tab="Go"
func saveUnsignedTransaction() {

	// setup connection
	algodClient := setupConnection()

	// get network suggested parameters
	txParams, err := algodClient.SuggestedParams()
	if err != nil {
		fmt.Printf("error getting suggested tx params: %s\n", err)
		return
	}

	// create transaction
	fromAddr := "EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4"
	genID := txParams.GenesisID
	genesisHash := base64.StdEncoding.EncodeToString(txParams.GenesisHash)
	voteKey := "eXq34wzh2UIxCZaI1leALKyAvSz/+XOe0wqdHagM+bw="
	selKey := "X84ReKTmp+yfgmMCbbokVqeFFFrKQeFZKEXG89SXwm4="
	voteFirst := uint64(6000000)
	voteLast := uint64(9000000)
	keyDilution := uint64(1730)
	tx, err := transaction.MakeKeyRegTxnWithFlatFee(fromAddr, 2000, 6002000,
		6003000, nil, genID, genesisHash, voteKey, selKey, voteFirst, voteLast,
		keyDilution)
	if err != nil {
		fmt.Printf("Error creating transaction: %s\n", err)
		return
	}
	unsignedTx := types.SignedTxn{
		Txn: tx,
	}
```

```zsh tab="goal"
# WARNING: This command must be run on the node where the partkey lives and the node
# must only have a single partkey for the account. Otherwise the command will
# choose one at random.
$ goal account changeonlinestatus --address=EW64GC6F24M7NDSC5R3ES4YUVE3ZXXNMARJHDCCCLIHZU6TBEOC7XRSBG4 --fee=2000 --firstvalid=6002000 --lastvalid=6003000 --online=true --txfile=online.txn
```

# Authorize and Send the Transaction
Use the appropriate [authorization method](../../features/transactions/signatures.md) to sign the transaction. 

!!! tip
    It is recommended that you authorize the transaction offline to protect the security of your private keys, especially for high-value participating accounts. 

Once the transaction is signed, wait for the network to reach the transaction's first valid round and then [submit](../../build-apps/hello_world.md#submit-the-transaction) and [wait for confirmation](../../build-apps/hello_world.md#wait-for-confirmation). 

# Check that the node is participating

At any time, you can validate whether your node is participating by `grep`-ing the `node.log` file in your data directory, looking for a `"VoteBroadcast"` messages where the `"Sender"` is your public key participation address.

```
$ grep 'VoteBroadcast' node.log
...
{"Context":"Agreement","Hash":"QJADVNJZDXYEQUPHITB6REFDGBY4AHBPPBIPVXLOPOASZA4T3PIA","ObjectPeriod":0,"ObjectRound":896659,"ObjectStep":2,"Period":0,"Round":0,"Sender":"3IE2GDYYSI56U53AQ6UUWRGAIGG5D4RHWLMCXJOPWQJA2ABF2X2A","Step":0,"Type":"VoteBroadcast","Weight":1,"WeightTotal":1,"file":"pseudonode.go","function":"github.com/algorand/go-algorand/agreement.pseudonodeVotesTask.execute","level":"info","line":344,"msg":"vote created for broadcast (weight 1, total weight 1)","time":"2019-05-10T18:38:54.137592-04:00"}
...
```

**See also**

- [Key Registration Transactions](../../features/transactions/index.md#key-registration-transaction)
- [Register account online](../../features/transactions/index.md#register-account-online)
