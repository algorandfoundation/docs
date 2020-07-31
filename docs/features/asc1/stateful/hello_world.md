title: Your First Application, a Stateful Smart Contract

Use the guide below to deploy a "hello world" Algorand application, a [stateful smart contract](/features/asc1/#stateful-smart-contracts) using the [Transaction Execution Approval Language (TEAL)](/features/asc1/teal/). This requires [setting up your workspace](/build-apps/setup/) and [connecting to a node](/build-apps/connect/). This guide covers drafting the smart contract code, deploying to the network, interacting with and removing the application.

This is a very basic application which implements a counter. Each time the application is called, the counter value is incremented. 

!!! Info
	[Algorand Smart Contracts](http://127.0.0.1:8000/features/asc1/) (ASC1) are deployed as either “Stateful” and “Stateless” programs. Both use the [Transaction Execution Approval Language (TEAL)](/features/asc1/teal/) but the available [OpCodes](/reference/teal/opcodes/) and therefore their function vary by type. This guide covers only [stateful smart contracts](/features/asc1/#stateful-smart-contracts); learn more about [stateless smart contracts](/features/asc1/#stateful-smart-contracts).

## Stateful Application Overview

Before diving into the code, it’s important to first understand the components of an Algorand application. The term “stateful” means the application is able to store information or “maintain state” within the ledger. The information ("data") is structured into _key/value pairs_. The [Transaction Execution Approval Language (TEAL)](/features/asc1/teal/) defines the available [OpCodes](/reference/teal/opcodes/) for use during program execution. _Application Call Transactions_ are used to interact with the application and may include _arguments_ (additional data) which are evaluated by the program at run-time. Every program execution must complete with a single non-zero unit64 value remaining on the stack to be valid and thus commit all state changes to the ledger. 

## Application Components

### Programs

Every Algorand stateful application is comprised of at least two _programs_: *approval* and *clear state*. 

#### Approval Program
The _approval program_ is the main application logic for the smart contract. 

#### Clear State Program
The _clear state_ program is used to retire an application, its global state and remove local state from the user’s account. 

### State Storage Locations

Applications have access to three storage locations for stateful data: _global_, _local_ and _external_. 

#### Global

_Global_ data are stored within the application itself. The program will read from and write to its global storage the same regardless of which account submitted the _application call transaction_.

#### Local

Each program may read from and write to _local_ storage within the account object of the "calling account". This allows an application developer to store user specific data and vary program execution logic at run-time per user.

#### External

Each program may read both the _global_ and _local_ state storage locations for limited number of other _external_ programs and accounts. However, the program is prohibited from writing to an _external_ state storage location.

### State Data

Both global and local data are stored as _key/value pairs_, where the _key_ is bytes and the _value_ may be either a `bytes` or `uint64`.

```json tab="Data Structure Template"
{ “key”: {
	“type”: <1 || 2>,
	“value”: <[]byte || uint64>
	}
}
```
### Program Execution
Application users will

#### Application Call Transaction

Each _application call transaction_ from a user account instructs the _approval application_ to execute. Arguments may be specified within the call transaction and consumed by the program during execution. 

### State Access and Updates

TEAL provides OpCodes allowing the program to _get_ (read) and _put_ (write) data within state storage locations. 

#### get
Programs may implicitly read their own _global_ storage and the _local_ storage of the account submitting the _application call transaction_. 

#### get_external
Reading from _global_ and _local_ storage of an _external_ program or account is allowed by explicitly passing the address as an argument within the _application call transaction_. Programs may read from _global_ storage of up to four _external_ programs. Additionally, programs may read from _local_ storage of up to two _external_ accounts.

#### put
Writing data is restricted to _global_ storage of the "called" program and the _local_ storage of the "calling" account, both specified within the _application call transaction_ (note: _external_ locations may only be read from).

## Draft Application Code

### Approval Program

Create a new file named `approval_program.teal` and add the following code:

```teal
// define stateful TEAL version
#pragma version 2

// read global state
byte "counter"
dup
app_global_get

// increment the value
int 1
+

// store to scratch space
dup
store 0

// update global state
app_global_put

// load return value as approval
load 0
return
```

#### Define TEAL Version

The _approval program_ is holding stateful data so if must instruct the TEAL interpreter to use “version 2” OpCodes during execution.

- `#pragma version 2`  

#### Read from Global State

This program has a single _global_ _key/value pair_ to store the number of times the program was called. The _key_ is the string "counter" and the _value_ will be defined as an integer when the application is created below.

- `byte “counter”` places the bytes representing the string “counter” on the stack
- `dup` duplicates that, so now “counter” is on the stack twice
- `app_global_get` pops the top of the stack, looks within _global_ for the _key_ “counter” and places the _value_ back on top of the stack.

#### Perform Program Logic

The program now has the current value for "counter" and will increment this value. A copy of this value will be stored in scratch space for use below. Note: scratch space is accessible by the program only during run-time and cannot persist any values after execution completes.

- `int 1` places the integer 1 on the stack
- `+` pops two elements from the stack (the _value_ for “counter” and int 1), sums them, and places the result back on top of the stack
- `dup` duplicates the to top most value on the stack
- `store 0` pops the duplicate value from the stack and moves it to the first position within scratch space.

#### Write to Global State

Now the program will store the incremented _value_ into _global_ storage at _key_ "counter". This value will persist on the ledger if the program completes successfully.

- `app_global_put` pops two elements from the stack (_key_ “counter” from line 5 and _value_ calculated in line 11) and stores into _global_ the _value_ for _key_.

#### Approval

At this point the stack is empty. A valid approval program must complete with a single non-zero uint64 value remaining on the stack. If the value is 0x00 (false) the program will fail, all state transitions will be discarded, and the calling transaction will also fail. The scratch space still holds a duplicate copy the counter value. Loading that ensures the program will complete successfully.

- `load 0` copies the first value from scratch space and places it on top of the stack
- `return` use last value on stack as success value; end

### Clear State Program

The _clear state program_ is used by the application to clear the _global_ and _local_ state storage locations for the application and accounts. The "hello world" application does not evaluate any conditions and simply approves the call.

Create a new file named `clear_state_program.teal` and add the following code:
```teal
#pragma version 2
// This program clears program state

int 1
```

- `#pragma version 2` instructs the TEAL interpreter to use “version 2” OpCodes during execution. This is required for stateful smart contract applications
- `int 1` places the integer 1 on the stack, signaling approval.

## Deploy New Application

With the two program files completed, the application may be created on the blockchain using the `bash` script below. The script uses `goal` and requires a funded account `$ADDR_CREATOR` to pay the transaction fee. The flags are rather descriptive. Notice that only `$GLOBAL_INTS` has a value of `1` and is used to allocate a single _global_ state storage element using the `--global-ints ` flag. 

```bash tab="bash"
export ADDR_CREATOR="YOURACCOUNTIDENTIFIERGOESHERE"

export TEAL_APPROVAL_PROG="approval_program.teal"
export TEAL_CLEAR_PROG="clear_state_program.teal"

export GLOBAL_BYTESLICES=0
export GLOBAL_INTS=1
export LOCAL_BYTESLICES=0
export LOCAL_INTS=0

goal app create --creator $ADDR_CREATOR \
                --approval-prog $TEAL_APPROVAL_PROG \
				--clear-prog $TEAL_CLEAR_PROG \
				--global-byteslices $GLOBAL_BYTESLICES \
				--global-ints $GLOBAL_INTS \
				--local-byteslices $LOCAL_BYTESLICES \
				--local-ints $LOCAL_INTS 
```

Results:
```bash
Attempting to create app (approval size 25, hash L4N6WP75R2G6M3TMLWSLA5S4PNHQIMGYTFMSOWNU6Q6X3R5LOU5A; clear size 5, hash YOE6C22GHCTKAN3HU4SE5PGIPN5UKXAJTXCQUPJ3KKF5HOAH646A)
Issued transaction from account TZIWJ3BKGLGTFDA4YAZPHIXGIEOVA5XLHWOCINUQY2PPHZ6XVFGKGMWVGA, txid L7A7HHCLWV3KBUHKMSVEVBMRPWRLHTXEUVYVVLFKWUDOWJEAWLTQ (fee 1000)
Transaction L7A7HHCLWV3KBUHKMSVEVBMRPWRLHTXEUVYVVLFKWUDOWJEAWLTQ still pending as of round 260564
Transaction L7A7HHCLWV3KBUHKMSVEVBMRPWRLHTXEUVYVVLFKWUDOWJEAWLTQ committed in round 260566
Created app with app index 33
```

Application ID 33 was created above. Check yours using:

```bash
goal account dump --address $ADDR_CREATOR
```
Results:
```json
{
  "addr": "TZIWJ3BKGLGTFDA4YAZPHIXGIEOVA5XLHWOCINUQY2PPHZ6XVFGKGMWVGA",
  "algo": 308000,
  "appp": {
    "33": {
      "approv": "AiABASYBB2NvdW50ZXIoSWQiCEk1AGc0AA==",
      "clearp": "AiABASI=",
      "gs": {
        "counter": {
          "tt": 2,
          "ui": 1
        }
      },
      "gsch": {
        "nui": 1
      }
    }
  },
  "ebase": 6516,
  "tsch": {
    "nui": 1
  }
}
```

Notice the "appp" section has "33" which is the application ID found within the creator's account. Export your specific value for use within future scripts using:
```bash
export APP_ID=<your_app_id>
```

## Application Calls

There are a number of application calls available to interact with the application.

### OptIn to Application

Most applications require each user [OptIn]() prior to interacting with the application. However, this "hello world" application did not allocate any _local_ state when created so it will not require an OptIn _application call_ by any accounts. 

### Read State

Both _global_ and _local_ state may be read by anyone accessing the ledger. The following displays the current state of your application:

```bash
goal app read --global --app-id $APP_ID
```
Result:
```json
{
  "counter": {
    "tt": 2,
    "ui": 1
}
```
Here the single _global_ _key/value pair_ for your application is displayed. The _key_ "counter" was assigned the type "tt" of integer at creation (1 for bytes, 2 for uint64). The current integer _value_ "ui" is 1. The _approval program_ executed once during creation. 

### Call Application

Calling the application again requires knowing the application id and paying the transaction fee. Anyone on the network may call this application, so modifying the `--from` flag will succeed if that account is funded.
```bash
goal app call --app-id $APP_ID --from $ADDR_CREATOR
```

Check the application state again:
```bash
goal app read --global --app-id $APP_ID
```
Result:
```json
{
  "counter": {
    "tt": 2,
    "ui": 2
}
```
The "hello world" application has been called again as observed by the incremented _value_ "ui" for _key_ "counter". Continue making calls and reading the updated state.

## Retire the Application
Each account is limited to creating or opting into a total of 10 applications. Removing an application may be accomplished using:
```bash
goal app delete --app-id $APP_ID --from $ADDR_CREATOR
```
Results for both `goal account dump` and `goal app read` will not include the application. You may reuse the existing program source files to create the application again. However, that will result in a new application id and a reset of your _global_ "counter" state back to 1. 