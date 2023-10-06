Title: Creation methods

This section describes the three primary methods to create accounts on Algorand, how to use them in the SDKs, `goal`, and `algokey`, and the reasons you might want to choose one method over another for your application. 

# Quick start videos

If you prefer videos, take a look at this 10 minute guide to getting started with creating accounts, which also includes using the Pera Algo Wallet.

<iframe width="100%" style="aspect-ratio:16/9" src="https://www.youtube-nocookie.com/embed/TnpGO0P0BA0" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>

The three primary ways to create accounts on Algorand are as [wallet-derived accounts](#wallet-derived-kmd) (using [kmd](../../../run-a-node/reference/artifacts#kmd)), as [standalone](#standalone), or as [multisignature accounts](#multisignature) (which entails one of the prior methods).

!!! info
    Remember that accounts participating in transactions are required to maintain a minimum balance of 100,000 micro Algos. Prior to using a newly created account in transactions, make sure that it has a sufficient balance by transferring at least 100,000 micro Algos to it.  An initial transfer of under that amount will fail due to the minimum balance constraint.

!!! info
	The Algorand community provides many wallets that can be used to create an Algorand account as well. See [Wallets](https://developer.algorand.org/ecosystem-projects/?tags=wallets)	for more details.

# Wallet-derived (kmd)

The Key Management Daemon is a process that runs on [Algorand nodes](../../../run-a-node/reference/artifacts#kmd), so if you are using a [third-party API service](../../../archive/build-apps/setup#1-use-a-third-party-service), this process likely will not be available to you. kmd is the underlying key storage mechanism used with `goal`.  The SDKs also connect to kmd through a REST endpoint and access token. 

**Reasons you might want to use kmd**

Public/private key pairs are generated from a single master derivation key. You only need to remember the single mnemonic that represents this master derivation key (i.e. the wallet passphrase/mnemonic) to regenerate all of the accounts in that wallet. 

There is no way for someone else to determine that two addresses are generated from the same master derivation key. This provides a potential avenue for applications to implement anonymous spending for end users without requiring users to store multiple passphrases.

**Reasons you might not want to use kmd**

Using kmd requires running a process and storing keys on disk. If you do not have access to a node or you require a more lightweight solution, [Standalone Accounts](#standalone) may be a better suited option.


## How-to use kmd
### Start the kmd process

To initiate the kmd process and generate the required `kmd.net` and `kmd.token` files use [`goal kmd`](../../../clis/goal/kmd/kmd) or [`kmd`](../../../clis/kmd) command line utilities. 

Start kmd with a 3600 second timeout.

=== "goal"
	```zsh
	$ goal kmd start -t 3600
	Successfully started kmd
	```

=== "kmd"
	```zsh
	$ kmd -d data/kmd-v<version>/ -t 3600
	```

Retrieve the kmd IP address and access token:

```zsh
$ echo "kmd IP address: " `cat $ALGORAND_DATA/kmd-v<version>/kmd.net`
kmd IP address:  [ip-address]:[port]

$ echo "kmd token: " `cat $ALGORAND_DATA/kmd-v<version>/kmd.token`
kmd token:  [token]
```
### Create a wallet and generate an account

Create a new wallet and generate an account. In the SDKs, connect to kmd through a kmd client then create a new wallet. With the wallet handle, generate an account. 

=== "JavaScript"
	<!-- ===JSSDK_KMD_CREATE_CLIENT=== -->
	```javascript
	const kmdToken = 'a'.repeat(64);
	const kmdServer = 'http://localhost';
	const kmdPort = 4002;
	
	const kmdClient = new algosdk.Kmd(kmdToken, kmdServer, kmdPort);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/kmd.ts#L10-L15)
	<!-- ===JSSDK_KMD_CREATE_CLIENT=== -->
	<!-- ===JSSDK_KMD_CREATE_WALLET=== -->
	```javascript
	const walletName = 'testWallet1';
	const password = 'testpassword';
	// MDK is undefined since we are creating a completely new wallet
	const masterDerivationKey = undefined;
	const driver = 'sqlite';
	
	const wallet = await kmdClient.createWallet(
	  walletName,
	  password,
	  masterDerivationKey,
	  driver
	);
	const walletID = wallet.wallet.id;
	console.log('Created wallet:', walletID);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/kmd.ts#L23-L37)
	<!-- ===JSSDK_KMD_CREATE_WALLET=== -->
	<!-- ===JSSDK_KMD_CREATE_ACCOUNT=== -->
	```javascript
	// wallet handle is used to establish a session with the wallet
	const wallethandle = (
	  await kmdClient.initWalletHandle(walletID, 'testpassword')
	).wallet_handle_token;
	console.log('Got wallet handle:', wallethandle);
	
	const { address } = await kmdClient.generateKey(wallethandle);
	console.log('Created new account:', address);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/kmd.ts#L40-L48)
	<!-- ===JSSDK_KMD_CREATE_ACCOUNT=== -->

=== "Python"
	<!-- ===PYSDK_KMD_CREATE_CLIENT=== -->
	```python
	kmd_address = "http://localhost:4002"
	kmd_token = "a" * 64
	
	kmd_client = kmd.KMDClient(kmd_token=kmd_token, kmd_address=kmd_address)
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/kmd.py#L5-L9)
	<!-- ===PYSDK_KMD_CREATE_CLIENT=== -->
	<!-- ===PYSDK_KMD_CREATE_WALLET=== -->
	```python
	# create a wallet object which, if not available yet, also creates the wallet in the KMD
	wlt = wallet.Wallet("MyNewWallet", "supersecretpassword", kmd_client)
	# get wallet information
	info = wlt.info()
	print(f"Wallet name: {info['wallet']['name']}")
	
	backup = wlt.get_mnemonic()
	print(f"mnemonic for master derivation key: {backup}")
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/kmd.py#L29-L37)
	<!-- ===PYSDK_KMD_CREATE_WALLET=== -->
	<!-- ===PYSDK_KMD_CREATE_ACCOUNT=== -->
	```python
	# create an account using the wallet object
	address = wlt.generate_key()
	print(f"New account: {address}")
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/kmd.py#L40-L43)
	<!-- ===PYSDK_KMD_CREATE_ACCOUNT=== -->

=== "Java"
	<!-- ===JAVASDK_KMD_CREATE_CLIENT=== -->
	```java
	String kmdHost = "http://localhost:4002";
	String kmdToken = "a".repeat(64);
	
	KmdClient kmdClient = new KmdClient();
	kmdClient.setBasePath(kmdHost);
	kmdClient.setApiKey(kmdToken);
	
	KmdApi kmd = new KmdApi(kmdClient);
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/KMDExamples.java#L26-L34)
	<!-- ===JAVASDK_KMD_CREATE_CLIENT=== -->
	<!-- ===JAVASDK_KMD_CREATE_WALLET=== -->
	```java
	// create a new CreateWalletRequest and set parameters
	CreateWalletRequest cwr = new CreateWalletRequest();
	cwr.setWalletName(walletName);
	cwr.setWalletPassword(password);
	cwr.setWalletDriverName("sqlite"); // other option is `ledger`
	// using our client, pass the request
	APIV1POSTWalletResponse result = kmd.createWallet(cwr);
	APIV1Wallet wallet = result.getWallet();
	System.out.printf("Wallet name: %s\n", wallet.getName());
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/KMDExamples.java#L40-L49)
	<!-- ===JAVASDK_KMD_CREATE_WALLET=== -->
	<!-- ===JAVASDK_KMD_CREATE_ACCOUNT=== -->
	```java
	// create a request to generate a new key, using the handle token
	GenerateKeyRequest gkr = new GenerateKeyRequest();
	gkr.setWalletHandleToken(handleToken);
	APIV1POSTKeyResponse generatedKey = kmd.generateKey(gkr);
	String addr = generatedKey.getAddress();
	System.out.printf("New account: %s\n", addr);
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/KMDExamples.java#L74-L80)
	<!-- ===JAVASDK_KMD_CREATE_ACCOUNT=== -->

=== "Go"
	<!-- ===GOSDK_KMD_CREATE_CLIENT=== -->
	```go
	// Create a new kmd client, configured to connect to out local sandbox
	var kmdAddress = "http://localhost:4002"
	var kmdToken = strings.Repeat("a", 64)
	kmdClient, err := kmd.MakeClient(
		kmdAddress,
		kmdToken,
	)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/kmd/main.go#L23-L30)
	<!-- ===GOSDK_KMD_CREATE_CLIENT=== -->
	<!-- ===GOSDK_KMD_CREATE_WALLET=== -->
	```go
	// Create the example wallet, if it doesn't already exist
	createResponse, err := kmdClient.CreateWallet(
		"DemoWallet",
		"password",
		kmd.DefaultWalletDriver,
		types.MasterDerivationKey{},
	)
	if err != nil {
		fmt.Printf("error creating wallet: %s\n", err)
		return
	}
	
	// We need the wallet ID in order to get a wallet handle, so we can add accounts
	exampleWalletID = createResponse.Wallet.ID
	fmt.Printf("Created wallet '%s' with ID: %s\n", createResponse.Wallet.Name, exampleWalletID)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/kmd/main.go#L40-L55)
	<!-- ===GOSDK_KMD_CREATE_WALLET=== -->
	<!-- ===GOSDK_KMD_CREATE_ACCOUNT=== -->
	```go
	// Get a wallet handle.
	initResponse, _ = kmdClient.InitWalletHandle(
		exampleWalletID,
		"password",
	)
	exampleWalletHandleToken = initResponse.WalletHandleToken
	
	// Generate a new address from the wallet handle
	genResponse, err = kmdClient.GenerateKey(exampleWalletHandleToken)
	if err != nil {
		fmt.Printf("Error generating key: %s\n", err)
		return
	}
	accountAddress := genResponse.Address
	fmt.Printf("New Account: %s\n", accountAddress)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/kmd/main.go#L58-L73)
	<!-- ===GOSDK_KMD_CREATE_ACCOUNT=== -->

=== "goal"
	```zsh
	$ goal wallet new testwallet
	Please choose a password for wallet 'testwallet': 
	Please confirm the password: 
	Creating wallet...
	Created wallet 'testwallet'
	Your new wallet has a backup phrase that can be used for recovery.
	Keeping this backup phrase safe is extremely important.
	Would you like to see it now? (Y/n): y
	Your backup phrase is printed below.
	Keep this information safe -- never share it with anyone!

	[25-word mnemonic]

	$ goal account new
	Created new account with address [address]
	```

### Recover wallet and regenerate account
To recover a wallet and any previously generated accounts, use the wallet backup phrase (also called the wallet mnemonic or passphrase). The master derivation key for the wallet will always generate the same addresses in the same order. Therefore the process of recovering an account within the wallet looks exactly like generating a new account. 

!!! info
	An offline wallet may not accurately reflect account balances, but the state for those accounts (e.g. its balance, online status) are safely stored on the blockchain. kmd will repopulate those balances when connected to a node.

=== "JavaScript"
	<!-- ===JSSDK_KMD_RECOVER_WALLET=== -->
	```javascript
	const exportedMDK = (
	  await kmdClient.exportMasterDerivationKey(wallethandle, 'testpassword')
	).master_derivation_key;
	const recoveredWallet = await kmdClient.createWallet(
	  'testWallet2',
	  'testpassword',
	  exportedMDK,
	  'sqlite'
	);
	const recoeveredWalletID = recoveredWallet.wallet.id;
	
	console.log('Created wallet: ', recoeveredWalletID);
	
	const recoveredWalletHandle = (
	  await kmdClient.initWalletHandle(recoeveredWalletID, 'testpassword')
	).wallet_handle_token;
	console.log('Got wallet handle: ', recoveredWalletHandle);
	
	const recoveredAddr = (await kmdClient.generateKey(recoveredWalletHandle))
	  .address;
	console.log('Recovered account: ', recoveredAddr);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/kmd.ts#L67-L88)
	<!-- ===JSSDK_KMD_RECOVER_WALLET=== -->

=== "Python"
	<!-- ===PYSDK_KMD_RECOVER_WALLET=== -->
	```python
	# Create the master derivation key from our backed up mnemonic
	mdk = mnemonic.to_master_derivation_key(backup)
	
	# recover the wallet by passing mdk during creation
	new_wallet = wallet.Wallet(
	    "MyNewWalletCopy", "testpassword", kmd_client, mdk=mdk
	)
	
	info = new_wallet.info()
	wallet_id = info["wallet"]["id"]
	print(f"Created Wallet: {wallet_id}")
	
	rec_addr = wlt.generate_key()
	print("Recovered account:", rec_addr)
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/kmd.py#L47-L61)
	<!-- ===PYSDK_KMD_RECOVER_WALLET=== -->

=== "Java"
	<!-- ===JAVASDK_KMD_RECOVER_WALLET=== -->
	```java
	// create a new CreateWalletRequest and set parameters
	CreateWalletRequest recoverRequest = new CreateWalletRequest();
	recoverRequest.setWalletName("Recovered:" + walletName);
	recoverRequest.setWalletPassword(password);
	recoverRequest.setWalletDriverName("sqlite");
	// Pass the specific derivation key we want to use
	// to recover the wallet
	recoverRequest.setMasterDerivationKey(backupKey);
	APIV1POSTWalletResponse recoverResponse = kmd.createWallet(recoverRequest);
	APIV1Wallet recoveredWallet = recoverResponse.getWallet();
	System.out.printf("Wallet name: %s\n", recoveredWallet.getName());
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/KMDExamples.java#L60-L71)
	<!-- ===JAVASDK_KMD_RECOVER_WALLET=== -->

=== "Go"
	<!-- ===GOSDK_KMD_RECOVER_WALLET=== -->
	```go
	keyBytes, err := mnemonic.ToKey(backupPhrase)
	if err != nil {
		fmt.Printf("failed to get key: %s\n", err)
		return
	}
	
	var mdk types.MasterDerivationKey
	copy(mdk[:], keyBytes)
	recoverResponse, err := kmdClient.CreateWallet(
		"RecoveryWallet",
		"password",
		kmd.DefaultWalletDriver,
		mdk,
	)
	if err != nil {
		fmt.Printf("error creating wallet: %s\n", err)
		return
	}
	
	// We need the wallet ID in order to get a wallet handle, so we can add accounts
	exampleWalletID = recoverResponse.Wallet.ID
	fmt.Printf("Created wallet '%s' with ID: %s\n", recoverResponse.Wallet.Name, exampleWalletID)
	
	// Get a wallet handle. The wallet handle is used for things like signing transactions
	// and creating accounts. Wallet handles do expire, but they can be renewed
	initResponse, err = kmdClient.InitWalletHandle(
		exampleWalletID,
		"password",
	)
	if err != nil {
		fmt.Printf("Error initializing wallet handle: %s\n", err)
		return
	}
	
	// Extract the wallet handle
	exampleWalletHandleToken = initResponse.WalletHandleToken
	fmt.Printf("Got wallet handle: '%s'\n", exampleWalletHandleToken)
	
	// Generate a new address from the wallet handle
	genResponse, err = kmdClient.GenerateKey(exampleWalletHandleToken)
	if err != nil {
		fmt.Printf("Error generating key: %s\n", err)
		return
	}
	fmt.Printf("Recovered address %s\n", genResponse.Address)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/kmd/main.go#L118-L163)
	<!-- ===GOSDK_KMD_RECOVER_WALLET=== -->

=== "goal"
	<!-- ===GOAL_KMD_RECOVER_WALLET=== -->
	```zsh
	$ goal wallet new -r <recovered-wallet-name>
	Please type your recovery mnemonic below, and hit return when you are done: 
	[25-word wallet mnemonic]
	Please choose a password for wallet [RECOVERED_WALLET_NAME]: 
	Please confirm the password: 
	Creating wallet...
	Created wallet [RECOVERED_WALLET_NAME]

	$ goal account new -w <recovered-wallet-name>
	Created new account with address [RECOVERED_ADDRESS]
	```
	<!-- ===GOAL_KMD_RECOVER_WALLET=== -->

### Export an account
Use this to retrieve the 25-word mnemonic for the account.

=== "JavaScript"
	<!-- ===JSSDK_KMD_EXPORT_ACCOUNT=== -->
	```javascript
	const accountKey = await kmdClient.exportKey(wallethandle, password, address);
	const accountMnemonic = algosdk.secretKeyToMnemonic(accountKey.private_key);
	console.log('Account Mnemonic: ', accountMnemonic);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/kmd.ts#L51-L54)
	<!-- ===JSSDK_KMD_EXPORT_ACCOUNT=== -->

=== "Python"
	<!-- ===PYSDK_KMD_EXPORT_ACCOUNT=== -->
	```python
	# Get the id for the wallet we want to export an account from
	wallet_id = get_wallet_id_from_name("MyNewWallet")
	# Get a session handle for the wallet after providing password
	wallethandle = kmd_client.init_wallet_handle(wallet_id, "supersecretpassword")
	# Export the account key for the address passed
	accountkey = kmd_client.export_key(
	    wallethandle, "supersecretpassword", address
	)
	# Print the mnemonic for the accounts private key
	mn = mnemonic.from_private_key(accountkey)
	print(f"Account mnemonic: {mn}")
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/kmd.py#L64-L75)
	<!-- ===PYSDK_KMD_EXPORT_ACCOUNT=== -->

=== "Java"
	<!-- ===JAVASDK_KMD_EXPORT_ACCOUNT=== -->
	```java
	ExportKeyRequest ekr = new ExportKeyRequest();
	ekr.setAddress(addr);
	ekr.setWalletHandleToken(handleToken);
	ekr.setWalletPassword(password);
	APIV1POSTKeyExportResponse exportedKeyResp = kmd.exportKey(ekr);
	byte[] exportedKey = exportedKeyResp.getPrivateKey();
	String mn = Mnemonic.fromKey(Arrays.copyOfRange(exportedKey, 0, 32));
	System.out.printf("Exported mnemonic: %s\n", mn);
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/KMDExamples.java#L83-L91)
	<!-- ===JAVASDK_KMD_EXPORT_ACCOUNT=== -->

=== "Go"
	<!-- ===GOSDK_KMD_EXPORT_ACCOUNT=== -->
	```go
	// Extract the account sk
	accountKeyResponse, _ := kmdClient.ExportKey(
		exampleWalletHandleToken,
		"password",
		accountAddress,
	)
	accountKey := accountKeyResponse.PrivateKey
	// Convert sk to mnemonic
	mn, err := mnemonic.FromPrivateKey(accountKey)
	if err != nil {
		fmt.Printf("Error getting backup phrase: %s\n", err)
		return
	}
	fmt.Printf("Account Mnemonic: %v ", mn)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/kmd/main.go#L76-L90)
	<!-- ===GOSDK_KMD_EXPORT_ACCOUNT=== -->

### Import an account
Use these methods to import a 25-word account-level mnemonic.

!!! warning
	For compatibility with other developer tools, `goal` provides functions to import and export accounts into kmd wallets, however, keep in mind that an imported account can **not** be recovered/derived from the wallet-level mnemonic. You must always keep track of the account-level mnemonics that you import into kmd wallets.

=== "JavaScript"
	<!-- ===JSSDK_KMD_IMPORT_ACCOUNT=== -->
	```javascript
	const newAccount = algosdk.generateAccount();
	console.log('Account: ', newAccount.addr);
	const importedAccount = await kmdClient.importKey(
	  wallethandle,
	  newAccount.sk
	);
	console.log('Account successfully imported: ', importedAccount);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/kmd.ts#L57-L64)
	<!-- ===JSSDK_KMD_IMPORT_ACCOUNT=== -->

=== "Python"
	<!-- ===PYSDK_KMD_IMPORT_ACCOUNT=== -->
	```python
	wallet_id = get_wallet_id_from_name("MyNewWallet")
	
	# Generate a new account client side
	new_private_key, new_address = account.generate_account()
	mn = mnemonic.from_private_key(new_private_key)
	print(f"Account: {new_address} Mnemonic: {mn}")
	
	# Import the account to the wallet in KMD
	wallethandle = kmd_client.init_wallet_handle(wallet_id, "supersecretpassword")
	importedaccount = kmd_client.import_key(wallethandle, new_private_key)
	print("Account successfully imported: ", importedaccount)
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/kmd.py#L78-L89)
	<!-- ===PYSDK_KMD_IMPORT_ACCOUNT=== -->

=== "Java"
	<!-- ===JAVASDK_KMD_IMPORT_ACCOUNT=== -->
	```java
	
	String recoveredWalletHandleToken = getHandle(kmd, recoveredWallet, password);
	
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/KMDExamples.java#L92-L95)
	<!-- ===JAVASDK_KMD_IMPORT_ACCOUNT=== -->

=== "Go"
	<!-- ===GOSDK_KMD_IMPORT_ACCOUNT=== -->
	```go
	account := crypto.GenerateAccount()
	fmt.Println("Account Address: ", account.Address)
	mn, err = mnemonic.FromPrivateKey(account.PrivateKey)
	if err != nil {
		fmt.Printf("Error getting backup phrase: %s\n", err)
		return
	}
	fmt.Printf("Account Mnemonic: %s\n", mn)
	importedAccount, _ := kmdClient.ImportKey(
		exampleWalletHandleToken,
		account.PrivateKey,
	)
	fmt.Println("Account Successfully Imported: ", importedAccount.Address)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/kmd/main.go#L93-L106)
	<!-- ===GOSDK_KMD_IMPORT_ACCOUNT=== -->

# Standalone 

A standalone account is an Algorand address and private key pair that is _not_ stored on disk. The
private key is most often in the [25-word mnemonic form](../#transformation-private-key-to-25-word-mnemonic).


**Reasons you might want to use standalone accounts**

Standalone accounts have a low setup cost as you do not need to connect to a separate client that depends on separate hardware. All you need is the 25-word human-readable mnemonic of the relevant account. 

Since keys are not stored on disk, standalone accounts can be used in [secure offline signing procedures](../../transactions/offline_transactions) where hardware constraints may make using kmd more difficult.

Standalone account mnemonics are widely used across developer tools and services within the Algorand ecosystem. However, this should not limit developers who prefer to use kmd since [import](#import-account) and [export](#export-account) functions exist with kmd to ensure compatibility.

!!! info
	Algorand's mobile wallet (Android, iOS) uses standalone accounts. Use the 25-word mnemonic to import accounts into the mobile wallet.

**Reasons you might _not_ want to use standalone accounts**

If you prefer storing your keys encrypted on disk instead of storing human-readable 25-word mnemonics, kmd may be a better option. 

### How to generate a standalone account


=== "JavaScript"
	<!-- ===JSSDK_ACCOUNT_GENERATE=== -->
	```javascript
	const generatedAccount = algosdk.generateAccount();
	const passphrase = algosdk.secretKeyToMnemonic(generatedAccount.sk);
	console.log(`My address: ${generatedAccount.addr}`);
	console.log(`My passphrase: ${passphrase}`);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/accounts.ts#L80-L84)
	<!-- ===JSSDK_ACCOUNT_GENERATE=== -->

=== "Python"
	<!-- ===PYSDK_ACCOUNT_GENERATE=== -->
	```python
	private_key, address = account.generate_account()
	print(f"address: {address}")
	print(f"private key: {private_key}")
	print(f"mnemonic: {mnemonic.from_private_key(private_key)}")
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/account.py#L5-L9)
	<!-- ===PYSDK_ACCOUNT_GENERATE=== -->

=== "Java"
	<!-- ===JAVASDK_ACCOUNT_GENERATE=== -->
	```java
	Account acct = new Account();
	System.out.println("Address: " + acct.getAddress());
	System.out.println("Passphrase: " + acct.toMnemonic());
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/Overview.java#L76-L79)
	<!-- ===JAVASDK_ACCOUNT_GENERATE=== -->

=== "Go"
	<!-- ===GOSDK_ACCOUNT_GENERATE=== -->
	```go
	account := crypto.GenerateAccount()
	mn, err := mnemonic.FromPrivateKey(account.PrivateKey)
	
	if err != nil {
		log.Fatalf("failed to generate account: %s", err)
	}
	
	log.Printf("Address: %s\n", account.Address)
	log.Printf("Mnemonic: %s\n", mn)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/account/main.go#L16-L25)
	<!-- ===GOSDK_ACCOUNT_GENERATE=== -->


=== "algokey"
	```bash
	$ algokey generate
	Private key mnemonic: [PASSPHRASE]
	Public key: [ADDRESS]
	```

# Multisignature

# Quick start videos - Multisignature

If you prefer videos, take a look at this 4 minute guide to Learn About Multisig Algorand Account.

<iframe width="100%" style="aspect-ratio:16/9" src="https://www.youtube-nocookie.com/embed/kOugNkk4HuE" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>

Multisignature accounts are a logical representation of an ordered set of addresses with a threshold and version. Multisignature accounts can perform the same operations as other accounts, including sending transactions and participating in consensus. The address for a multisignature account is essentially a hash of the _ordered_ list of accounts, the threshold and version values. The threshold determines how many signatures are required to process any transaction from this multisignature account. 

Multisignature accounts cannot nest other multisignature accounts. 

Creating a multisignature account with Address A, Address B, and Address C will not produce the same address as one with Address B, Address A, and Address C. However, signing a multisignature transaction does not require any specific order. 

!!! tip
	You can use the fact that order matters to generate multiple addresses that can be signed by the same set of keys.

Send Algos to the address to [initialize its state on the blockchain](#accounts) as you would any other address.


**Reasons you might want to use multisignature accounts**

Since every transaction requires a threshold of signatures you can create an extra layer of security on an account by requiring multiple signatures to authorize spending. The total accounts and threshold can be tailored to fit your security model.

The keys that can sign for the multisignature account can be stored in separate locations and they can be generated with kmd, as standalone accounts, or with a mixture of both.

Multisignature accounts can also be used to create cryptographically secure governance structures for an account, where keys can be owned by multiple users and spending is authorized by a subset of those users. Pair this with Algorand Smart Contract functionality for the potential to realize even more complex governance structures such as authorizing spending from an account given a _specific_ subset of signatures. Read more about [TEAL](../../dapps/avm/teal) and [Algorand Smart Contracts](../../dapps/smart-contracts/).

**Reasons you might _not_ want to use multisignature accounts**

Multisignature accounts trade off convenience for security. Every transaction requires multiple signatures which can be overly complex for a scenario where security or governance is not critical. 

## How to generate a multisignature account
The following code shows how to generate a multisignature account composed of three Algorand addresses, with a signing threshold of 2, and using version 1 of the software (currently the only version). Hardcode the mnemonics in the code samples below to recreate a specific multisignature address or create [new accounts](#how-to-generate-a-standalone-account).

!!! tip
	Since multisignature accounts are just logical representations of the data defined above, anyone can "create" the same Algorand address if they know how it is composed. This information is public and included in a signed transaction from a multisignature account. See [how multisignatures look in a signed transaction](../../transactions/signatures#multisignatures).

=== "JavaScript"
	<!-- ===JSSDK_MULTISIG_CREATE=== -->
	```javascript
	const signerAccounts: algosdk.Account[] = [];
	signerAccounts.push(algosdk.generateAccount());
	signerAccounts.push(algosdk.generateAccount());
	signerAccounts.push(algosdk.generateAccount());
	
	// multiSigParams is used when creating the address and when signing transactions
	const multiSigParams = {
	  version: 1,
	  threshold: 2,
	  addrs: signerAccounts.map((a) => a.addr),
	};
	const multisigAddr = algosdk.multisigAddress(multiSigParams);
	
	console.log('Created MultiSig Address: ', multisigAddr);
	```
	[Snippet Source](https://github.com/algorand/js-algorand-sdk/blob/examples/examples/accounts.ts#L27-L41)
	<!-- ===JSSDK_MULTISIG_CREATE=== -->


=== "Python"
	<!-- ===PYSDK_MULTISIG_CREATE=== -->
	```python
	version = 1  # multisig version
	threshold = 2  # how many signatures are necessary
	# create a Multisig given the set of participants and threshold
	msig = transaction.Multisig(
	    version,
	    threshold,
	    [account_1.address, account_2.address, account_3.address],
	)
	print("Multisig Address: ", msig.address())
	```
	[Snippet Source](https://github.com/algorand/py-algorand-sdk/blob/examples/examples/account.py#L25-L34)
	<!-- ===PYSDK_MULTISIG_CREATE=== -->

=== "Java"
	<!-- ===JAVASDK_MULTISIG_CREATE=== -->
	```java
	int version = 1; // no other versions at the time of writing
	int threshold = 2; // we're making a 2/3 msig
	
	// Populate a list of Ed25519 pubkeys
	List<Ed25519PublicKey> accts = new ArrayList<>();
	accts.add(addr1.getEd25519PublicKey());
	accts.add(addr2.getEd25519PublicKey());
	accts.add(addr3.getEd25519PublicKey());
	// create the MultisigAddress object
	MultisigAddress msig = new MultisigAddress(version, threshold, accts);
	System.out.printf("msig address: %s\n", msig.toAddress().toString());
	```
	[Snippet Source](https://github.com/algorand/java-algorand-sdk/blob/examples/examples/src/main/java/com/algorand/examples/AccountExamples.java#L77-L88)
	<!-- ===JAVASDK_MULTISIG_CREATE=== -->

=== "Go"
	<!-- ===GOSDK_MULTISIG_CREATE=== -->
	```go
	// Get pre-defined set of keys for example
	_, pks := loadAccounts()
	addr1, _ := types.DecodeAddress(pks[1])
	addr2, _ := types.DecodeAddress(pks[2])
	addr3, _ := types.DecodeAddress(pks[3])
	
	ma, err := crypto.MultisigAccountWithParams(1, 2, []types.Address{
		addr1,
		addr2,
		addr3,
	})
	
	if err != nil {
		panic("invalid multisig parameters")
	}
	fromAddr, _ := ma.Address()
	// Print multisig account
	fmt.Printf("Multisig address : %s \n", fromAddr)
	```
	[Snippet Source](https://github.com/algorand/go-algorand-sdk/blob/examples/examples/kmd/main.go#L178-L196)
	<!-- ===GOSDK_MULTISIG_CREATE=== -->

=== "goal"
	<!-- ===GOAL_MULTISIG_CREATE=== -->
	```zsh
		$ ADDRESS1=$(goal account new | awk '{ print $6 }')
		$ ADDRESS2=$(goal account new | awk '{ print $6 }')
		$ ADDRESS3=$(goal account new | awk '{ print $6 }')

		$ goal account multisig new $ADDRESS1 $ADDRESS2 $ADDRESS3 -T 2
		Created new account with address [MULTISIG_ADDRESS]
	```
	<!-- ===GOAL_MULTISIG_CREATE=== -->

Multisignature accounts may also be referred to as multisig accounts and a multisig account composed of 3 addresses with a threshold of 2 is often referred to as a 2 out of 3 (i.e. 2/3) multisig account.