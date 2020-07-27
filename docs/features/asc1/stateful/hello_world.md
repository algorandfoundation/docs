title: Your First Stateful Smart Contract

Use the guide below to deploy a "hello world" Algorand application using [stateful TEAL](TODO:). This requires [setting up your workspace](https://developer.algorand.org/docs/build-apps/setup) and [connecting to a node](https://developer.algorand.org/docs/build-apps/connect/). This guide covers drafting the smart contract code, deploying to the network, interacting with and removing the application.

This is a very basic application which implements a counter. Each time the application is called, the counter value is incremented. 

!!! Info
	Algorand implements both “Stateful” and “Stateless” applications, sometimes referred to as version 1 and version 2 programs respectively. Both are smart contracts and use the [TEAL Language](https://developer.algorand.org/docs/features/asc1/teal_overview/) but the available [OpCodes](https://developer.algorand.org/docs/reference/teal/opcodes/) vary by version. This guide covers only stateful applications.

## Stateful Application Overview

Before diving into the code, it’s important to first understand the structure of an Algorand stateful smart contract. The term “stateful” means the application is able to store information or “maintain state” within the ledger. The information, _data_, is structured into key/value pairs. The TEAL language defines the available OpCodes used during application execution. Transaction are used to _call_ the application and may include _arguments_ (additional data) which are evaluated at run-time. Application execution must complete with a single, non-zero, value remaining on the stack to be valid and commit all state changes to the ledger. 

## Application Components

### State Locations

Applications have access to two storage locations for stateful data: _global_ and _local_. 

#### Global Data

_Global_ data is stored within the application. The application will read its global data the same regardless of which user made the application call.

#### Local Data

Each application may store _local_ data within the user’s account. This allows an application developer to store user specific data and vary the execution results per user.

### State Variables

Both global and local data are store are as key/value pairs, where the key is a byteslice and the value may be either a byteslice or an integer.

```json
{ “key”: {
	“type”: <1 || 2>,
	“value”: <[]byte || uint64>
	}
}
```

### Programs

Every Algorand application is comprised of at least two _programs_: *approval* and *clear state*. 

#### Approval Program
The _approval program_ is the main application logic for the smart contract. 

#### Clear State Program
The _clear state_ program is used to retire an application, its global state and remove local state from the user’s account. 

### Interfaces
#### put
#### get
#### get external

## Draft Application Code

### Approval Program

Create a new file named `approval_program.teal` and add the following code:

```teal
#pragma version 2

// read global state
byte "counter"
dup
app_global_get

// increment the value
int 1
+

// update global state
app_global_put

int 1
return
```

#### Define pragma version

The first line `#pragma version 2` instructs the TEAL interpreter to use “version 2” OpCodes during execution. This is required for stateful smart contract applications. 

#### Read from Global State

- Line 4 `byte “counter”` places the bytes representing the string “counter” on the stack
- Line 5 `dup` duplicates that, so now “counter” is on the stack twice.
- Line 6 `app_global_get` pops the top of the stack, looks within _global_ for the _key_ “counter” and places the _value_ back on top of the stack.

#### Perform Program Logic

- Line 9 `int 1` places the integer 1 on the stack.
- Line 10 `+` pops two elements from the stack (the _value_ for “counter” and int 1), adds them, and places the result back on top of the stack.

#### Write to Global State

Line 13 `app_global_put` also pops two elements from the stack (_key_ “counter” from line 4 and _value_ calculated in line 10) and stores into _global_ for this application the _value_ for _key_.

#### Stack Execution

At this point the stack is empty. A valid approval program must end with a single , non-zero, value remaining on the stack. If the value is 0x00 (false) the program will fail, all state transitions will be discarded, and the calling transaction will also fail. TODO: does the caller get charged the FEE if TX fails? How is that reflected on the ledger/user history?

### Clear State Program

Create a new file named `clear_state_program.teal` and add the following code:
```teal
#pragma version 2
// This program clears program state

int 1
```

- The first line `#pragma version 2` instructs the TEAL interpreter to use “version 2” OpCodes during execution. This is required for stateful smart contract applications. 
- Line 4 `int 1` places the integer 1 on the stack. This allows the program to clear the state without any additional conditions.

## Deploy New Application

```bash
export TEAL_APPROVAL_PROG="approval_program.teal"
export TEAL_CLEAR_PROG="clear_state_program.teal"

export ADDR_CREATOR="YOURACCOUNTIDENTIFIERGOESHERE"

export GLOBAL_BYTESLICES=0
export GLOBAL_INTS=1
export LOCAL_BYTESLICES=0
export LOCAL_INTS=0
```

```bash
goal app create --creator $ADDR_CREATOR --approval-prog $TEAL_APPROVAL_PROG --clear-prog $TEAL_CLEAR_PROG --global-byteslices $GLOBAL_BYTESLICES --global-ints $GLOBAL_INTS --local-byteslices $LOCAL_BYTESLICES --local-ints $LOCAL_INTS 

goal account dump -a $ADDR_CREATOR
```

## Option Into Application (not required)

## Call Application

```bash
goal app read --global --app-id $APP_ID

goal app call --app-id $APP_ID --from $ADDR_CREATOR

goal app read --global --app-id $APP_ID
```

## Retire the Application
