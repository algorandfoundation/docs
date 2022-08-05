title: Build with Reach

Alice and Bob are now ready to kick-off development of their Algorand-powered auction [dApp](https://developer.algorand.org/docs/get-started/dapps/)! This guide will take you through the steps to build the dApp, which includes setting up a development environment, writing the dApp code which can contain multiple smart contracts, deploying it.

A few tips before getting started. The goal of this guide is to get you up and running with a working prototype that represents a real use case, as quickly as possible. We use a hands-on example to teach you basic design principles and best practices for building dApps with Reach on Algorand. This guide does not cover all the details of the dApp backend and frontend code. This is intentional so that you can focus on solidifying higher-level concepts that will be the foundation for building any dApp on Algorand using Reach. So don’t worry if you don’t understand what everything does in the solution. This is expected! After you feel comfortable with the basics, you can head over to the detailed [Reach documentation](https://docs.reach.sh/tut/#tuts) and work on becoming an expert in the Reach language.

All of the code for this guide is located [here](https://github.com/algorand/reach-auction). The `index.mjs` and `index.rsh` files are included, so follow along!

Now let’s get started. 

## Reach architecture

Using high level languages to build dApps instead of low level assembly language is attractive for many professional developers. Frontends can be built in languages such as Python, Go, JavaScript and C# utilizing Reach's RPC Server. Or they can directly use the Algorand SDK and interact with their application using the standard Algorand ABI. The focus for a developer using Reach is the business logic. Reach takes care of the internals on contract storage, protocol diagrams, state validation and network details in general. 

Reach programs contain all properties of a dApp. The **backend**, often found in a file such as [index.rsh](https://github.com/algorand/reach-auction/blob/main/index.rsh), controls what data is published to the consensus network (a blockchain) and defines interfaces that can interact with the **frontend**. The frontend is often contained within a file such as [index.mjs](https://github.com/algorand/reach-auction/blob/main/index.mjs) and provides a graphical representation of the interface for its participants. i.e. Alice and Bob. The user interface may be in the form of a command line, or more preferably, a web or mobile app.

In Reach, developers think about their participants and how they will interact with one another, as well as, with the contract. Developers codify the business logic of their application and the Reach compiler builds the smart contract. Reach abstracts the heavy lifting of smart contract development, allowing developers to focus their energy on participant interaction. With Reach, the automatic verification of a dApp, alone, is one great reason to consider using Reach vs PyTeal. This protects against blockchain attacks, guarantees contract balance is zero and prevents locked away tokens by the contract that are inaccessible. Reach facilitates blockchain properties, like token linearity. As such, auditing also becomes very easy to accomplish.


# Organization

This guide is organized into two sections. This document is the first section which helps you launch the dApp and run an auction simulation. The second section will provide a deeper dive into the different components of a reach application.




## Install Reach

Reach is designed to work on POSIX systems with [make](https://en.wikipedia.org/wiki/Make_(software)), [Docker](https://www.docker.com/get-started), and [Docker Compose](https://docs.docker.com/compose/install/) installed. The best way to install Docker on Mac and Windows is with [Docker Desktop](https://www.docker.com/products/docker-desktop). 


To confirm everything is installed try to run the following three commands and see no errors.


```
  $ make --version
  $ docker --version
  $ docker-compose --version
```


If you’re using Windows, consult [the guide to using Reach on Windows](https://docs.reach.sh/guide-windows.html).

Once confirmed that reach prerequisite are installed, choose a directory and clone the repo example so you can follow along:

```
  $ mkdir -p ~/reach/ && cd ~/reach/
  $ git clone https://github.com/algorand/reach-auction
```

Navigate to the project folder

``` bash
cd reach-auction
```

Next, download Reach by running
``` bash
$ curl https://docs.reach.sh/reach -o reach ; chmod +x reach
```


Since Reach is Dockerized, when first used, the images it uses need to be downloaded. This will happen automatically when used for the first time, but can be done manually now by running


```
  $ ./reach update
```


More information: Detailed Reach install instructions can be found in the [reach docs](https://reach.sh/). 



## Setup environments and run tests


Sometimes it may be convenient to use the reach run command, preceded by setting the **REACH_CONNECTOR_MODE **, especially when testing multiple blockchain deployments.

```
  REACH_CONNECTOR_MODE=ALGO-devnet ./reach run
```


Set an environment variable to use the Algorand Blockchain. 


```
export REACH_CONNECTOR_MODE="ALGO-devnet"
```


# Application Overview


## Alice and Bob’s auction design

The sample dApp shown below is an auction. This involves the potential buyer sending the bid in algos. If the bid is successful, the dApp will hold the algos. Bidding is only allowed between the beginning and end times of the auction. If the bid supplants a previous bid, the contract automatically refunds the previous higher bidder’s algos. The requirements are listed [here](https://developer.algorand.org/docs/get-started/dapps/#alice-and-bobs-auction-design).

<center>
![Auction close](../../imgs/dapp-close.png){: style="width:500px" align=center }
<figcaption style="font-size:12px">Close out the auction: Carla receives the NFT and Alice gets paid 2000 Algos.</figcaption>
</center>

## The dApp 

A build folder is created in your project by the Reach compiler when using 

`./reach run`

This folder contains a file called `index.main.mjs`. This file contains the smart contract bytecode that the blockchain is able to read. Additionally, `index.main.mjs` is a backend JavaScript file that contains an asynchronous function for each participant, as defined in `index.mjs`. 

You can imagine that `index.mjs` mirrors `index.main.mjs`. The `index.mjs` provides frontend interactivity and `index.main.mjs` enables functionality in the backend. For example, if a Reach program contains a participant named 'Alice' in the Reach.App, then the JavaScript backend will include a function named Alice (i.e. backend.Alice).

## Run the Reach Application

Run the auction


```
$ ./reach run
```


Output should be similar to below,  showing a winner of the auction:


```
Creating test account for Creator
Having creator create testing NFT
Creator sets parameters of sale: {
  nftId: BigNumber { _hex: '0x03', _isBigNumber: true },
  minBid: BigNumber { _hex: '0x1e8480', _isBigNumber: true },
  lenInBlocks: 10
}
Alice decides to bid 5.011883.
Alice balance before is 99.999
Alice out bid 0xcef1e5570c827650b927ac62d3342d722af32e72bed67ae80e79796922ba2c52 who bid 2.
Alice balance after is 94.985117
Bob decides to bid 5.999184.
Bob balance before is 99.999
Bob out bid 0x7b104010a9355fc63015fd8543813fed4e6af8345e371e71cd480d9e3fa1678e who bid 5.011883.
Bob balance after is 93.996816
Claire decides to bid 6.22388.
Claire balance before is 99.999
Claire failed to bid, because the auction is over
Claire balance after is 99.999
Creator saw that PMIEAEFJGVP4MMAV7WCUHAJ75VHGV6BULY3R44ONJAGZ4P5BM6HJSDRA74 bid 5.011883.
Creator saw that FX7A54UVE6YAD2C5PPSK65VKKDK3DKG2YB3OD6NEZXMHPRV62BRF7G6PJM bid 5.999184.
Creator saw that FX7A54UVE6YAD2C5PPSK65VKKDK3DKG2YB3OD6NEZXMHPRV62BRF7G6PJM won with 5.999184
Alice has 99.997198 ALGO and 0 of the NFT
Bob has 93.997002 ALGO and 1 of the NFT
Claire has 99.999198 ALGO and 0 of the NFT
```
# Backend

The Backend [index.rsh](https://github.com/algorand/reach-auction/blob/main/index.rsh) , defines the interface for functions coded in the frontend. A participant is an "actor" who takes part in the application (dApp) and can have persistently stored values, called its local state. Participants are associated with an account (address) on the consensus network. A Consensus Network is a Network protocol (a blockchain) that contains network tokens (ALGO, ETH, etc.), non-network tokens (ASA, ERC-20, etc.), as well as a set of accounts and contracts. 

## Basic Functions of the Auction

`Creator` is a Participant that has getSale, seeBid and timeout functions.  

`Bidder.bid` is an API Participant that determines the highest bid and calls the seeBid function. API participants refer to a class of participants who share the same functionality. They are essentially the same type of user who share the same functionality.

Imagine an application where users vote for their favorite color. There can be unlimited voters who participate in the act of voting. Although each voter is an independent user, they all share the same functionality. In this example, all voters are members of the “voter API”. Similarly, in this dApp, we have an API of Bidders. Every bidder within this API has the ability to place a bid, but only one can win the auction.

An API is unique because it represents the frontend pinging the backend only as needed. Using APIs allows Reach developers to create quieter and more efficient programs that are not as computationally expensive.

However, at least one user must always be a Participant because APIs are not able to deploy a Reach application.

`index.rsh`

```javascript
'reach 0.1';

// Creator is a Participant that has getSale, auctionReady, seeBid and 
// showOutcome functions.  A participant is an “actor” who 
// takes part in the application (dApp). 
// Participants are associated with an account (address) 
// on the consensus network.
export const main = Reach.App(() => {
    const Creator = Participant('Creator', {
        getSale: Fun([], Object({
            nftId: Token,
            minBid: UInt,
            lenInBlocks: UInt,
        })),
        auctionReady: Fun([], Null),
        seeBid: Fun([Address, UInt], Null),
        showOutcome: Fun([Address, UInt], Null),
    });
     // Bidder is an API that has a bid function.
    const Bidder = API('Bidder', {
        bid: Fun([UInt], Tuple(Address, UInt)),
    });
    init();
```
The above code simply defines a specific dApp participant, the creator of the auction, and a bidder API for bidders in the auction. This code also exposes the interfaces that the UI should provide as input to the backend of the dApp.

## Publish Information and Pay

The next section introduces `.pay()` and `.publish()` functionality, allowing participants to pay the contract and each other. They also enable the contract to pay out any remaining funds before exiting. Pay or Publish are often needed when the entire dApp is waiting for a single participant to act. If a participant is sharing information, then you need `publish()`. When paying a previously known amount, use `pay()`. This kind of transfer always explicitly names the acting participant.

Publish writes information to the blockchain. In this case, it is writing the nft ID, minimum bid, and length in blocks.

```
Creator.publish(nftId, minBid, lenInBlocks);
```

Here, the creator uses pay to transfer a specific NFT(quantity of 1) to the ownership of the dApps smart contract. Later, the contract will transfer this NFT to a winning bidder of the auction.


```
Creator.pay([[amt, nftId]]);
```

The code below the creator publishes information and pays 1 NFT while the dApp is waiting for a participant to act. The `lastConsensusTime()` primitive returns the network time of the last publication of the dApp.  This may not be available if there was no such previous publication, such as at the beginning of an application before the first publication. Using interact communicates with the frontend to use these parameters from the UI.



```javascript
    Creator.only(() => {
         // Binding the value of getSale to the result of interacting
         // with the participant. This happens in a local step. 
         // declassify declassifies the argument, 
         // in this case that means the value of getSale
        const {nftId, minBid, lenInBlocks} = declassify(interact.getSale());
    });
    Creator.publish(nftId, minBid, lenInBlocks);
    const amt = 1;
    commit();
    Creator.pay([[amt, nftId]]);
    Creator.interact.auctionReady();
    assert(balance(nftId) == amt, "balance of NFT is wrong");
    const end = lastConsensusTime() + lenInBlocks;
```
Recall the functions, `getSale()` and `auctionReady()` were created in the participant interact interface. Further into the workshop, the same functions will be used by the frontend. The participant interact interface enables the frontend to retrieve functionality from the backend. 
In the code above, `interact.getSale()` returns nftID, minBid and lenInBlocks. The Creator of the auction publishes the auction params for NFT id, minBid and lenInBlocks and commits to the blockchain. 

## Consensus Transfer Patterns

In the next snippet, the rules for the outcome of the bidding are defined. The consensus transfer uses `parallelReduce`, which facilitates bidders repeatedly providing new bids as they compete to be the highest bidder before a time limit is reached. The parallelReduce name comes from the fact that participants are trying to produce a new state from the current values in parallel. Parallel reduce is essentially a fork+case loop. So while Alice is producing a new state, Bob might also be producing one. Reach determines who tried to update the state first (fork) and depending on the winner of the race runs a specified logic (case). A consensus transfer occurs when a single participant (called the originator) makes a publication of a set of public values from its local state and transfers zero or more network tokens to the contract account. Additional consensus transfer patterns are discussed in the reach documentation [here](https://docs.reach.sh/guide-ctransfers.html). In the logic below, if the bid is greater than the currentPrice, the transfer is made to the highest bidder. It will also refund the previous high bidder. 

The code below contains `api_`, pronounced, "API Macro,". It is a special macro of `api`. The `api_` form can be used when the dynamic assertions, assume and require, are identical. One advantage of `api_` is that it allows Reach developers to write less code, thereby reducing opportunities to introduce bugs. It's also a quick indication that the dynamic assertions will not change between local and consensus steps.



`index.rsh`

```javascript
    // Create the parallelReduce; align the Creator as the first highestBidder, 
    // the minBid as the lastPrice, and set isFirstBid to the boolean, "true"
    const [
        highestBidder,
        lastPrice,
        isFirstBid,
    ] = parallelReduce([Creator, minBid, true])
    // The first invariant, asserts that the balance 
    // of the nftID will be equal to the amount of the NFTs
        .invariant(balance(nftId) == amt)
    // The second invariant which states that the balance 
    // will either be 0 if it is the first bid or equal to 
    // the most recent price    
        .invariant(balance() == (isFirstBid ? 0 : lastPrice))
        .while(lastConsensusTime() <= end)
        // If the bid is greater than the currentPrice, 
        // the transfer is made to the highest bidder. 
        // It will also refund the previous high bidder. 

        .api_(Bidder.bid, (bid) => {
            check(bid > lastPrice, "bid is too low");
            // Return the bid and a created "notify" argument to update 
            // other APIs with the newest high bidder and their price.
            return [ bid, (notify) => {
                notify([highestBidder, lastPrice]);
                // If this bid is not the first bid then the api_ 
                // will return the lastPrice to the address of the 
                // previous highest bidder.
                if ( ! isFirstBid ) {
                    transfer(lastPrice).to(highestBidder);
                }
                const who = this;
                // Creator interacts with the seeBid function, 
                // which accepts who (an API Bidder) and their bid.
                Creator.interact.seeBid(who, bid);
                // Return highestBidder, lastPrice, and isFirstBid. 
                // The parallelReduce will now update the variables in the race. 
                // As a result, new Bidders will see the updated highestBidder and the lastPrice.
                return [who, bid, false];
            }];
        })     
```


## Timeouts

Timeouts prevent participants from stalling transactions in their local steps. A timeout can be calculated in reference to the passage of time or the passage of blocks. When the contract times out, the Creator interacts with the timeout object in the frontend, publishes data to the participants and Reach returns the highest bidder, the last price and isFirstBid. The time argument can be expressed in `absoluteTime(amt)`, `relativeTime(time)`, `relativeSecs(amt)` or `absoluteSecs(secs)`.  Also `timeremaining()` can be used in conjunction with the `makeDeadline(UInt)` function. 

`index.rsh`

```javascript
        // A timeout method is provided to escape the parallelReduce. 
        // This timeout is triggered when the end of absoluteTime resolves to true. 
        // Once timeout is triggered, Creator will move the DApp to a consensus step with publish 
        // and the DApp will return the values of highestBidder, lastPrice, and isFirstBid
        .timeout(absoluteTime(end), () => {
            Creator.publish();
            return [highestBidder, lastPrice, isFirstBid];
        });
```


## Transfers

The transfer is made from the Bidder to the Creator for the bid amount lastPrice on the NFT. The NFT is transferred to the highest bidder.  Each can see the results with `showOutcome`. 

`index.rsh`

```javascript
        transfer(amt, nftId).to(highestBidder);
        if ( ! isFirstBid ) { transfer(lastPrice).to(Creator); }
        Creator.interact.showOutcome(highestBidder, lastPrice);
    commit();
    exit();
});
```

# Frontend

The **Frontend**, often created as `index.mjs`, is where test accounts for Alice, Bob, and Claire are created. 

An instance of the reach stdlib (standard library) is imported. The standard library contains utility functions that can be called from the frontend such as `parseCurrency` and `launchToken`. For testing purposes, each participant is given a balance of 100 algos.  An NFT is created. Finally, the Creator deploys the contract. 


`index.mjs`

```javascript
// Import and instance of the reach stdlib
import { loadStdlib } from '@reach-sh/stdlib';
// The following line is crucial to connect to the backend
import * as backend from './build/index.main.mjs';

const stdlib = loadStdlib();
  // Each participant is given a balance of 100 algos. Parse the native currency and return a balance
const startingBalance = stdlib.parseCurrency(100);
console.log(`Creating test account for Creator`);
const accCreator = await stdlib.newTestAccount(startingBalance);

console.log(`Having creator create testing NFT`);
// Create an NFT.
const theNFT = await stdlib.launchToken(accCreator, "bumple", "NFT", { supply: 1 });
const nftId = theNFT.id;
const minBid = stdlib.parseCurrency(2);
const lenInBlocks = 10;
const params = { nftId, minBid, lenInBlocks };
// Establish done as a let, rather than a const so that the boolean 
// is able to mutate to true when it is time to exit the parallelReduce
let done = false;
const bidders = [];
// Create startBidders(), a function that will be triggered by auctionReady()
const startBidders = async () => {
    // Store the value in minBid as bid. This value is also mutable 
    // so that the bid value can be updated as bidders race for the NFT
    let bid = minBid;
    // Create the runBidder() function, which will be called 
    // when API Participants place a bid
    const runBidder = async (who) => {
        const inc = stdlib.parseCurrency(Math.random() * 10);
        bid = bid.add(inc);
        // Initialize a new account with devnet tokens
        const acc = await stdlib.newTestAccount(startingBalance);
        // Set the account address as a distinguishing label in debug logs
        acc.setDebugLabel(who);
        // Return a promise when the NFT is ready to be accepted by the contract
        await acc.tokenAccept(nftId);
        // Using standard JavaScript, this statement pushes the bidder 
        // and their account to the bidders array.
        bidders.push([who, acc]);
        // Attach the API Participant to the contract
        const ctc = acc.contract(backend, ctcCreator.getInfo());
        // Store the balance of the Bidder in getBal
        const getBal = async () => stdlib.formatCurrency(await stdlib.balanceOf(acc));

        console.log(`${who} decides to bid ${stdlib.formatCurrency(bid)}.`);
        console.log(`${who} balance before is ${await getBal()}`);
        try {
            // Store the Bidder's address and bid amount in a constant
            const [ lastBidder, lastBid ] = await ctc.apis.Bidder.bid(bid);
            // Print a message stating who has been outbid and for how much
            console.log(`${who} out bid ${lastBidder} who bid ${stdlib.formatCurrency(lastBid)}.`);
        } catch (e) {
            // Print a message that a bid failed because the auction ended
            console.log(`${who} failed to bid, because the auction is over`);
        }
        // Print the Bidder's balance after their bid
        console.log(`${who} balance after is ${await getBal()}`);
    };
    // Create three bidders, Alice, Bob, and Claire who will wait to race until runBidder is called.
    await runBidder('Alice');
    await runBidder('Bob');
    await runBidder('Claire');
    // A while loop that increments devnet transaction blocks forward until done is true.
    while ( ! done ) {
        await stdlib.wait(1);
    }
};

```
In the code above, the bidders attach to the contract that the Creator deployed using `const ctc = acc.contract(backend, ctcCreator.getInfo())` . Console messages are displayed for the bidding process showing the bid with and random amount `stdlib.parseCurrency(Math.random() * 10)`.  

We've created the first half of the Bidders communication model. Bidders now have logic that dictates their behavior when interacting with the contract. This communication model allows Bidders to enter a race with one another and ensures that they are aware of the highest bidder and the bid to beat.

A clearly defined communication model is extremely important because in web3 there is no middle man. Decentralized applications create an opportunity to reduce fees and increase the velocity of communications and funding. In this case, instead of waiting for the highest bidder to wire credit, the network tokens are escrowed the moment the bid is placed, and transferred the moment the auction closes.

## Interact Logic

Interact methods are expressions that only occur in a participant’s local step. They are called from the backend for the frontend to execute. In other words, interact methods receive an evaluation of an expression and output a value. 

In this program, interact methods include:  `getSale`,`auctionReady`,  `seeBid`, and `showOutcome`. 

The `getSale` function sets the parameters of the sale including the NFT Asset ID, min bid, and length in Blocks. 

The `auctionReady` method starts the bidding action. 

The `seeBid` function shows the bid and who saw it.

The `showOutcome` function shows who won. 

This code snippet starts with creating the contract by the Creator. 

```javascript
// The creator account creates the contract
const ctcCreator = accCreator.contract(backend);
await ctcCreator.participants.Creator({
    getSale: () => {
        console.log(`Creator sets parameters of sale:`, params);
        return params;
    },
    auctionReady: () => {
        startBidders();
    },
    seeBid: (who, amt) => {
        console.log(`Creator saw that ${stdlib.formatAddress(who)} bid ${stdlib.formatCurrency(amt)}.`);
    },
    showOutcome: (winner, amt) => {
        console.log(`Creator saw that ${stdlib.formatAddress(winner)} won with ${stdlib.formatCurrency(amt)}`);
    },
});
```

Finally, we have the for loop which sumarizes the status of each bidder.

```javascript
for ( const [who, acc] of bidders ) {
    // Store the balance of each bidder's account and how many NFTs the account holds
    const [ amt, amtNFT ] = await stdlib.balancesOf(acc, [null, nftId]);
    // Print the log of the user's balance and number of NFTs
    console.log(`${who} has ${stdlib.formatCurrency(amt)} ${stdlib.standardUnit} and ${amtNFT} of the NFT`);
}
done = true;
```

You have learned how to build a basic Reach app which consists of backend and frontend code. The backend defines the interface for frontend functions as well as the participants and interacting with the frontend functions and persisting local storage at the account level. The frontend deploys the dApp to the blockchain and provides the logic around account creation, funding and basic functions of your solution such as getting a bid, and creating an NFT. 


## Learn More

Verification and Remote Procedure Calls (RPC) are key features to Reach. Reach's verification engine ensures that invariants about the state of a program assumed by programmers are held by all possible executions of the program. In other words, the verification process provides smart contracts that are resilient to common logic errors that lead to mistakes.  The Reach RPC Server provides access to compiled JavaScript backends via an HTTPS-accessible JSON-based RPC protocol. The server allows frontends to be written in any programming language. Reach provides client libraries for JavaScript, Python, and Go. Example frontends written using the Reach RPC Server can be found in the Reach RPC documentation.

If you'd like to learn more about developing Algorand dApps using Reach, as well as more advanced concepts, see  [verification](https://docs.reach.sh/guide/assert/#guide-assert) and [Reach RPC](https://docs.reach.sh/rpc/) Server for more details. For auditing, mathematical proofs, cryptographic commitment schemes, loop invariants, timeouts and more see the Reach [documentation](https://docs.reach.sh/). 


Complete code for this auction simulation can be found [here](https://github.com/algorand/reach-auction). 
