title: Build with Reach

Alice and Bob are now ready to kick-off development of their Algorand-powered auction [dApp](https://developer.algorand.org/docs/get-started/dapps/)! This guide will take you through the steps to build the dApp, which includes setting up a development environment, writing the dApp code which can contain multiple smart contracts, deploying it, and writing the functions to interact with it.

A few tips before getting started. The goal of this guide is to get you up and running with a working prototype that represents a real use case, as quickly as possible. We use a hands-on example to teach you basic design principles and best practices for building dApps with Reach on Algorand. This guide does not cover all the details of the dApp backend and frontend code. This is intentional so that you can focus on solidifying higher-level concepts that will be the foundation for building any dApp on Algorand using Reach. So don’t worry if you don’t understand what everything does in the solution. This is expected! After you feel comfortable with the basics, you can head over to the detailed [Reach documentation](https://docs.reach.sh/) and work on becoming an expert in the Reach language.

Now let’s get started. 

# Organization

This guide is organized into two sections. This document is the first section which helps you launch the dApp and run an auction simulation. The second section will provide a deeper dive into the different components of a reach application.


All of the code for this guide is located [here] (https://github.com/algorand/reach-auction). The `index.mjs` and `index.rsh` files are included, so follow along!

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

Confirm the download worked by running


``` bash
$ ./reach version
```


Since Reach is Dockerized, when first used, the images it uses need to be downloaded. This will happen automatically when used for the first time, but can be done manually now by running


```
  $ ./reach update
```


You’ll know that everything is in order if you can run


```
  $ ./reach compile --help
```


To determine the current version is installed, run


```
  $ ./reach hashes
```


Output should look similar to:

```
reach: fb449c94
reach-cli: fb449c94
react-runner: fb449c94
rpc-server: fb449c94
runner: fb449c94
devnet-algo: fb449c94
devnet-cfx: fb449c94
devnet-eth: fb449c94
```


All of the hashes listed should be the same and then visit the #releases channel on the [Reach Discord Server](https://discord.gg/9kbHPfwbwn) to see the current hashes.

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

The sample app shown below is an auction. Alice and Bob think through the features for their dApp. They list off the following requirements:



1. Sellers must be able to create a new auction for each piece of artwork. The artwork must be held by the contract after the auction begins and until the auction closes.
2. The auction can be closed before it begins, in which case the artwork will be returned to the seller.
3. Sellers can specify a reserve price, which if not met, will return the artwork to them.
4. For each bid, if the new bid is higher than the previous bid, the previous bid will be refunded to the previous bidder and the new bid will be recorded and held by the contract.
5. At the end of a successful auction, where the reserve price was met, the highest bidder will receive the artwork and the seller will receive the full bid amount.

## Reach architecture

Reach programs contain all properties of a dApp. The **backend**, often found in a file such as [index.rsh](https://github.com/algorand/reach-auction/blob/main/index.rsh), controls what data is published to the consensus network (a blockchain) and defines interfaces that can interact with the **frontend**. The frontend is often contained within a file such as [index.mjs](https://github.com/algorand/reach-auction/blob/main/index.mjs) and provides a graphical representation of the interface for its participants. i.e. Alice and Bob. The user interface may be in the form of a command line, or more preferably, a web or mobile app.

In Reach, developers think about their participants and how they will interact with one another, as well as, with the contract. Developers codify the business logic of their application and the Reach compiler builds the smart contract. Reach abstracts the heavy lifting of smart contract development, allowing developers to focus their energy on participant interaction.

# The dApp 

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
Alice has 10 ALGO and 0 of the NFT
Alice decides to bid 2.508057
Bob has 10 ALGO and 0 of the NFT
Bob decides to bid 9.575844
Carla has 10 ALGO and 0 of the NFT
Carla decides to bid 3.81785
Creator has 9.999 ALGO and 1 of the NFT
Creator sets parameters of sale
Bob sees that the NFT is 6, the reserve price is 2, and that they have until 38 to bid
Alice sees that the NFT is 6, the reserve price is 2, and that they have until 38 to bid
Carla sees that the NFT is 6, the reserve price is 2, and that they have until 38 to bid
Bob bids 9.575844 against 2
Alice bids 2.508057 against 2
Carla bids 3.81785 against 2
Alice does not bid because 9.575844 is too high
Carla does not bid because 9.575844 is too high
Creator saw that 2KRZLQILBKY43A5OC56TP3SJBBSP35MV5X4GSVZDUOHYYEQK2JU5TN27BI bid 9.575844
Creator observes the auction has hit the timeout
Creator observes the auction has hit the timeout
Creator saw that 2KRZLQILBKY43A5OC56TP3SJBBSP35MV5X4GSVZDUOHYYEQK2JU5TN27BI won
Creator has 19.561862 ALGO and 0 of the NFT
Bob saw that 2KRZLQILBKY43A5OC56TP3SJBBSP35MV5X4GSVZDUOHYYEQK2JU5TN27BI won
Bob has 0.421156 ALGO and 1 of the NFT
Alice does not bid because 9.575844 is too high
Carla does not bid because 9.575844 is too high
Alice saw that 2KRZLQILBKY43A5OC56TP3SJBBSP35MV5X4GSVZDUOHYYEQK2JU5TN27BI won
Carla saw that 2KRZLQILBKY43A5OC56TP3SJBBSP35MV5X4GSVZDUOHYYEQK2JU5TN27BI won
Alice has 9.999009 ALGO and 0 of the NFT
Carla has 9.999009 ALGO and 0 of the NFT
```
# Backend

The Backend [index.rsh](https://github.com/algorand/reach-auction/blob/main/index.rsh) , defines the interface for functions coded in the frontend. A participant is an "actor" which takes part in the application (dApp). A participant is an "actor" who takes part in the application (dApp) and can have persistently stored values, called its local state. Participants are associated with an account (address) on the consensus network. A Consensus Network is a Network protocol (a blockchain) that contains network tokens (ALGO, ETH, etc.), non-network tokens (ASA, ERC-20, etc.), as well as, a set of accounts and contracts. 

## Basic Functions of the Auction

`Creator` is a Participant that has getSale, seeBid and timeout functions.  

`Bidder` is a ParticipantClass that has seeParams and getBid functions. A participant class is a category of Participant, it is like a Participant, but can occur many times in a single application. 

Imagine an application where users vote for their favorite color. There can be many voters who participate in the act of voting, and yet, collectively, they are all voters. In this example, all voters are members of the “voter participant class”. Similarly, in the auction dApp, we have a participant class of bidders. Every bidder within this class has the ability to place a bid.


`index.rsh`

```javascript
'reach 0.1';

const MUInt = Maybe(UInt);
const common = {
 showOutcome: Fun([Address], Null)
};
const Params = Tuple(Token, UInt, UInt);
// Creator is a Participant that has getSale, seeBid and timeout functions.  A participant is an “actor” who takes part in the application (dApp). Participants are associated with an account (address) on the consensus network.
export const main = Reach.App(() => {
 const Creator = Participant('Creator', {
   ...common,
   getSale: Fun([], Params),
   seeBid: Fun([Address, UInt], Null),
   timeout: Fun([], Null),
 });
 // Bidder is a ParticipantClass that has seeParams and getBid functions. 
 const Bidder = ParticipantClass('Bidder', {
   ...common,
   seeParams: Fun([Params], Null),
   getBid: Fun([UInt], MUInt),
 });
 deploy();
```
The above code simply defines a specific dApp participant, the creator of the auction, and a participant class of dApp actors, the bidders in the auction. This code also exposes the interfaces that the UI should provide as input to the backend of the dApp.

## Publish Information and Pay

The next section introduces `.pay()` and `.publish()` functionality, allowing participants to pay the contract and each other. They also enable the contract to pay out any remaining funds before exiting. Pay or Publish are often needed when the entire dApp is waiting for a single participant to act. If a participant is sharing information, then you need `publish()`. When paying a previously known amount, use `pay()`. This kind of transfer always explicitly names the acting participant.

Publish writes information to the blockchain. In this case, it is writing the nft ID, reserve price, and length in blocks.

```
Creator.publish(nftId, reservePrice, lenInBlocks);
```

Here, the creator uses pay to transfer a specific NFT(quantity of 1) to the ownership of the dApps smart contract. Later, the contract will transfer this NFT to a winning bidder of the auction.


```
Creator.pay([[amt, nftId]]);
```

The code below the creator publishes information and pays 1 NFT while the dApp is waiting for a participant to act. The `lastConsensusTime()` primitive returns the network time of the last publication of the dApp.  This may not be available if there was no such previous publication, such as at the beginning of an application before the first publication. The `Bidder.interact.seeParams` method interacts with the frontend to display the NFTId, reservePrice and the end time to bid.  Using interact communicates with the frontend to use these parameters from the UI.



```javascript
 Creator.only(() => {
   // Binding the value of getSale to the result of interacting with the participant. This happens in a local step. declassify declassifies the argument, in this case that means the value of getSale
   const [ nftId, reservePrice, lenInBlocks ] = declassify(interact.getSale());
 });
 Creator.publish(nftId, reservePrice, lenInBlocks);
 const amt = 1;
 commit();
 Creator.pay([[amt, nftId]]);
 const end = lastConsensusTime() + lenInBlocks;
 Bidder.interact.seeParams([nftId, reservePrice, end]);
```
Recall the functions, `getSale()` and `seeParams()` were created in the participant interact interface. Further into the workshop, the same functions will be used by the frontend. The participant interact interface enables the frontend to retrieve functionality from the backend. 
In the code above, `interact.getSale()` returns nftID, reservePrice and lenInBlocks. The Creator of the auction publishes the auction params for NFT id, reservePrice and lenInBlocks and commits to the blockchain. The Bidder sees the params.

## Consensus Transfer Patterns

In the next snippet, the rules for the outcome of the bidding are defined. The consensus transfer uses `parallelReduce`, which facilitates bidders repeatedly providing new bids as they compete to be the highest bidder before a time limit is reached. A consensus transfer occurs when a single participant (called the originator) makes a publication of a set of public values from its local state and transfers zero or more network tokens to the contract account. Additional consensus transfer patterns are discussed in the reach documentation [here](https://docs.reach.sh/guide-ctransfers.html). In the logic below, if the bid is greater than the currentPrice, the transfer is made to the highest bidder. It will also refund the previous high bidder. The Maybe computation can be some or none: (evaluate the return of a function). 



`index.rsh`

```javascript
// parallelReduce facilitates bidders repeatedly providing new bids as they compete to be the highest bidder before a time limit is reached
 const [ highestBidder, lastPrice, currentPrice ] =
   parallelReduce([ Creator, 0, reservePrice ])
     .invariant(balance(nftId) == amt && balance() == lastPrice)
     .while(lastConsensusTime() <= end)
     // If the bid is greater than the currentPrice, the transfer is made to the highest bidder. It will also refund the previous high bidder.
     .case(Bidder,
       (() => {
         const mbid = highestBidder != this
           ? declassify(interact.getBid(currentPrice))
           : MUInt.None();
         return ({
           // The Maybe computation can be some or none: (evaluate the return of a function).
           when: maybe(mbid, false, ((b) => b > currentPrice)),
           msg : fromSome(mbid, 0)
         });
       }),
       ((bid) => bid),
       ((bid) => {
         require(bid > currentPrice);
         // A consensus transfer occurs when a single participant (called the originator) makes a publication of a set of public values from its local state and transfers zero or more network tokens to the contract account.
         transfer(lastPrice).to(highestBidder);
         Creator.interact.seeBid(this, bid);
         return [ this, bid, bid ];
       }))

     
```
In the above code, the parallelReduce facilitates many bidders acting simultaneously as long as the auction has not been completed. First it uses the interact interface to retrieve a specific bidders offer and then if this value is higher than the current price, the contract transfers the current last price to the previous high bidder, transfers the new bid to the contract and displays this result to the front end using the interact interface’s seeBid method.

## Timeouts

Timeouts prevent participants from stalling transactions in their local steps. A timeout can be calculated in reference to the passage of time or the passage of blocks. When the contract times out, the Creator interacts with the timeout object in the frontend, publishes data to the participants and Reach returns the highest bidder, the last price and the current price. The time argument can be expressed in `absoluteTime(amt)`, `relativeTime(time)`, `relativeSecs(amt)` or `absoluteSecs(secs)`.  Also `timeremaining()` can be used in conjunction with the `makeDeadline(UInt)` function. 

`index.rsh`

```javascript
     .timeout(absoluteTime(end), () => {
       Creator.interact.timeout();
       Creator.publish();
       return [ highestBidder, lastPrice, currentPrice ];
     });
```


## Transfers

The transfer is made from the Bidder to the Creator for the bid amount lastPrice on the NFT. The NFT is transferred to the highest bidder.  Each can see the results with `showOutcome`. 

`index.rsh`

```javascript
 transfer(lastPrice).to(Creator);
 transfer(amt, nftId).to(highestBidder);
 commit();
// “each” participant listed can access the ShowOutcome method and see the outcome. 
 each([Creator, Bidder], () => interact.showOutcome(highestBidder));
 exit();
});
```

## Frontend

The **Frontend**, often created as `index.mjs`, is where test accounts for Alice, Bob, and Carla are created. 

An instance of the reach stdlib (standard library) is imported. The standard library contains utility functions that can be called from the frontend such as `parseCurrency` and `launchToken`. For testing purposes, each participant is given a balance of 10 algos.  An NFT is created. Finally, the Creator deploys the contract. 


`index.mjs`

```javascript
// Import and instance of the reach stdlib
import { loadStdlib } from '@reach-sh/stdlib/loader.mjs';
// The following line is crucial to connect to the backend
import * as backend from './build/index.main.mjs';

const N = 3;
const names = ["Creator", "Alice", "Bob", "Carla"];

(async () => {
 const stdlib = await loadStdlib(process.env);
  // Each participant is given a balance of 10 algos. Parse the native currency and return a balance
 const startingBalance = stdlib.parseCurrency(10);
 const [ accCreator, ...accBidders ] =
   await stdlib.newTestAccounts(1+N, startingBalance);
 // Create an NFT.
 const theNFT = await stdlib.launchToken(accCreator, "beepboop", "NFT", { supply: 1 });

 await Promise.all( [ accCreator, ...accBidders ].map(async (acc, i) => {
   acc.setDebugLabel(names[i]);
 }));
// Deploy the contract
 const ctcCreator = accCreator.contract(backend);
```

## Interact Logic


Interact methods are expressions that only occur in a participant’s local step. They are called from the backend for the frontend to execute. In other words, interact methods receive an evaluation of an expression and output a value. 
In this program, interact methods include: `showBalance`, `getSale`, `seeBid`, `timeout`, `showOutcome` and `showBalance`. 

The `showBalance` method uses `stdlib.balanceOf(acc)` to display the Algo balance of the interacting bidder. It can also be used to show the balance of a specific ASA or NFT token using `stdlib.balanceOf(acc, theNFT.id)`. 

The `getSale` function sets the parameters of the sale including the NFT Asset ID, reserve price, and length in Blocks. 

The `seeBid` function shows the bid and who saw it.


`index.mjs`

```javascript
 const showBalance = async (acc, i) => {
   const amt = await stdlib.balanceOf(acc);
   const amtNFT = await stdlib.balanceOf(acc, theNFT.id);
   console.log(`${names[i]} has ${stdlib.formatCurrency(amt)} ${stdlib.standardUnit} and ${amtNFT} of the NFT`);
 };

 await Promise.all([
   (async () => {
     await showBalance(accCreator, 0);
     const n = names[0];
     await backend.Creator(ctcCreator, {
       getSale: () => {
         console.log(`${n} sets parameters of sale`);
         return [ theNFT.id, stdlib.parseCurrency(2), 30 ]
       },
       seeBid: (who, bid) => {
         console.log(`${n} saw that ${stdlib.formatAddress(who)} bid ${stdlib.formatCurrency(bid)}`);
       },
       timeout: () => {
         console.log(`${n} observes the auction has hit the timeout`);
       },
       showOutcome: (winner) => {
         console.log(`${n} saw that ${stdlib.formatAddress(winner)} won`);
       },
     });
     await showBalance(accCreator, 0);
   })(),
```


The bidders attach to the contract that the Creator deployed using `acc.contract(backend, ctcCreator.getInfo())` . Console messages are displayed for the bidding process showing the bid with and random amount `stdlib.parseCurrency(Math.random() * 10)`.  The showOutcome function displays who won the bidding using `stdlib.addressEq(winner, acc)` In the getBid function, if the `currentPrice` is less than the bid, then the new price is returned as well as who placed the bid.  Finally, if the bidder did not win, they opt out of the asset using `theNFT.optOut(acc)`. 

`index.mjs`

```javascript
   ...accBidders.map(async (acc, i) => {
     await showBalance(acc, i+1);
     const n = names[i+1];
     const ctc = acc.contract(backend, ctcCreator.getInfo());
     const bid = stdlib.parseCurrency(Math.random() * 10);
     let IWon = false;
     console.log(`${n} decides to bid ${stdlib.formatCurrency(bid)}`);
     await backend.Bidder(ctc, {
       showOutcome: (winner) => {
         console.log(`${n} saw that ${stdlib.formatAddress(winner)} won`);
         IWon = stdlib.addressEq(winner, acc);
       },
       seeParams: async ([nftId, reservePrice, end]) => {
         console.log(`${n} sees that the NFT is ${nftId}, the reserve price is ${stdlib.formatCurrency(reservePrice)}, and that they have until ${end} to bid`);
         await acc.tokenAccept(nftId);
       },
       getBid: (currentPrice) => {
         if ( currentPrice.lt(bid) ) {
           console.log(`${n} bids ${stdlib.formatCurrency(bid)} against ${stdlib.formatCurrency(currentPrice)}`);
           return ['Some', bid];
         } else {
           console.log(`${n} does not bid because ${stdlib.formatCurrency(currentPrice)} is too high`);
           return ['None', null];
         }
       },
     });
     await showBalance(acc, i+1);
     if ( ! IWon ) {
       await theNFT.optOut(acc);
     }
     return;
   },
 )]);
})();
```

You have learned how to build a basic Reach app which consists of backend and frontend code. The backend defines the interface for frontend functions as well as the participants and interacting with the frontend functions and persisting local storage at the account level. The frontend deploys the dApp to the blockchain and provides the logic around account creation, funding and basic functions of your solution such as getting a bid, and creating an NFT. 


## Learn More

Verification and Remote Procedure Calls (RPC) are key features to Reach. Reach's verification engine ensures that invariants about the state of a program assumed by programmers are held by all possible executions of the program. In other words, the verification process provides smart contracts that are resilient to common logic errors that lead to mistakes.  The Reach RPC Server provides access to compiled JavaScript backends via an HTTPS-accessible JSON-based RPC protocol. The server allows frontends to be written in any programming language. Reach provides client libraries for JavaScript, Python, and Go. Example frontends written using the Reach RPC Server can be found in the Reach RPC documentation.

If you'd like to learn more about developing Algorand dApps using Reach, as well as more advanced concepts, see  [verification](https://docs.reach.sh/guide/assert/#guide-assert) and [Reach RPC](https://docs.reach.sh/rpc/) Server for more details.  For auditing, mathematical proofs, cryptographic commitment schemes, loop invariants,  timeouts and more see the Reach [documentation](https://docs.reach.sh/). 


Complete code for this auction simulation can be found [here](https://github.com/algorand/reach-auction). 
