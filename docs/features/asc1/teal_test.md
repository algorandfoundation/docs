title: Testing TEAL programs

This guide covers testing smart contracts without deploying them to a network by using Python SDK test harness.

## Prerequisites:
  - Python SDK installed.
  - **algod** running at `address:port` and token `token`.

Latest versions of **algod** come with so-called **Developer API** for TEAL program compilation and dry-running. The API is disabled by default, to enable set `EnableDeveloperAPI` to `true` in `config.json`. The following snapshot does the job for `goal node`-managed node:

```bash
export NETWORK_DIR=network
export NODE_NAME=Primary
goal node stop -d "$NETWORK_DIR/$NODE_NAME"

cp "$NETWORK_DIR/$NODE_NAME/config.json" "$NETWORK_DIR/$NODE_NAME/config.json.bak"
jq '. += {"EnableDeveloperAPI": true}' \
    < "$NETWORK_DIR/$NODE_NAME/config.json" \
    > "$NETWORK_DIR/$NODE_NAME/config.json.new"
rm "$NETWORK_DIR/$NODE_NAME/config.json"
mv "$NETWORK_DIR/$NODE_NAME/config.json.new" "$NETWORK_DIR/$NODE_NAME/config.json"

goal node start -d "$NETWORK_DIR/$NODE_NAME"
```

## Basic setup and simple tests

The test harness consist of `DryrunTestCaseMixin` class that is supposed to be used as a mixin to `unittest`-based user-defined tests and `Helper` class with various utilities. Look at example below.

```python tab="Python"
import base64
import unittest

from algosdk.constants import payment_txn, appcall_txn
from algosdk.encoding import decode_address, checksum
from algosdk.future import transaction
from algosdk.testing.dryrun import DryrunTestCaseMixin, Helper
from algosdk.v2client import algod
from algosdk.v2client.models import Account, Application, \
    ApplicationParams, ApplicationLocalState, ApplicationStateSchema, \
    TealKeyValue, TealValue

algod_host = "localhost"
algod_port = 60000
algod_token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"


class ExampleTestCase(DryrunTestCaseMixin, unittest.TestCase):
    def setUp(self):
        algod_address = f"http://{algod_host}:{algod_port}"
        self.algo_client = algod.AlgodClient(algod_token, algod_address)

    def test_simple(self):
        self.assertPass("int 1")
        self.assertReject("int 0")
        self.assertNoError("int 0")
        self.assertError("byte 0", "byte arg did not parse: 0")


if __name__ == "__main__":
    unittest.main()
```

There are a few things to pay attention to:
1. `algo_client` is expected to be initialized prior using any asserts or helpers from `DryrunTestCaseMixin`
2. Basic asserts `assertPass`, `assertReject`, `assertError` and `assertNoError` to check if the program return `true`, `false` or does \[not\] err during compilation or execution. These assert functions accept source program (`string` data type), compiled program (`byte` data type) or raw **dryrun response** object (either typed `DryrunResponse` or `dict` instance).

Further in the tutorial the same `ExampleTestCase` is used to show additional functionality.

`DryrunTestCaseMixin` provides helpers for testing both `LogicSig` and `Application` smart contracts.

## Testing `LogicSig` with arguments

Consider `test_logic_sig` function below added to `ExampleTestCase` test case.

```python tab="Python"
    def test_logic_sig(self):
        """Shows how to test logic sig with parameters"""
        source = """
arg 0
btoi
int 0x31
==
"""
        self.assertError(source, "cannot load arg[0]")
        self.assertReject(source)
        self.assertPass(source, lsig=dict(args=[b"1", b"2"]))

        drr = self.dryrun_request(source, lsig=dict(args=[b"\x31", b"2"]))
        self.assertPass(drr)
```

This example demonstrates how to pass `LogicSig` parameters - they needs to be a list of `bytes` items in `args` key of `lsig` parameter to any assert function. In general, specifying `lsig` non-none parameter forces to use `LogicSig` run mode.

If one needs more information about execution, use `dryrun_request` helper to get a raw **dryrun response** object as shown below. Note, `assert` functions are able to handle it as well.

```python tab="Python"
    def test_logic_sig_ex(self):
        """Shows how to user and examine raw dryrun response"""
        source = """
arg 0
btoi
int 0x31
==
"""
        drr = self.dryrun_request(source, lsig=dict(args=[b"\x31", b"2"]))
        self.assertPass(drr)
```

Printing out this raw **dryrun response** object:

```json
{"error": "", "protocol-version": "future", "txns": [{"disassembly": ["// version 1", "intcblock 49", "arg_0", "btoi", "intc_0", "==", ""], "logic-sig-messages": ["PASS"], "logic-sig-trace": [{"line": 1, "pc": 1, "stack": []}, {"line": 2, "pc": 4, "stack": []}, {"line": 3, "pc": 5, "stack": [{"bytes": "MQ==", "type": 1, "uint": 0}]}, {"line": 4, "pc": 6, "stack": [{"bytes": "", "type": 2, "uint": 49}]}, {"line": 5, "pc": 7, "stack": [{"bytes": "", "type": 2, "uint": 49}, {"bytes": "", "type": 2, "uint": 49}]}, {"line": 6, "pc": 8, "stack": [{"bytes": "", "type": 2, "uint": 1}]}]}]}
```

Note, the raw response contains entire execution trace in `logic-sig-trace` (or `app-call-trace` in case of **app call** execution). This info can be pretty printed with `Helper.pprint` call (see below).

## Testing `Application` global state changes

Now let's test an application and check what does it write to the global state. Example below is an initialization prologue of **voting app** (TODO: full code ref).

First, note `#pragma version 2` at the beginning of the source code. This forces the assembler to produce **TEAL v2** program with extended instructions set (`bz`, `app_global_put` and so on).
Second, `app` parameter in assert functions. It allows setting application-specific fields like `OnCompletion`, `ApplicationID`, `ApplicationArgs`, `Accounts` (more on this in the next section).

```python tab="Python"
    def test_app_global_state(self):
        """Use voting app as example to check app initialization"""
        source = """#pragma version 2
int 0
txn ApplicationID
==
bz not_creation
byte "Creator"
txn Sender
app_global_put
txn NumAppArgs
int 4
==
bz failed
byte "RegBegin"
txna ApplicationArgs 0
btoi
app_global_put
byte "RegEnd"
txna ApplicationArgs 1
btoi
app_global_put
byte "VoteBegin"
txna ApplicationArgs 2
btoi
app_global_put
byte "VoteEnd"
txna ApplicationArgs 3
btoi
app_global_put
int 1
return
not_creation:
int 0
return
failed:
int 0
return
"""
        self.assertReject(source, app=dict(app_idx=0))
        self.assertReject(source, app=dict(app_idx=1, args=[b"\x01", b"\xFF", b"\x01\x00", b"\x01\xFF"]))
        self.assertPass(source, app=dict(app_idx=0, args=[b"\x01", b"\xFF", b"\x01\x00", b"\x01\xFF"]))

        sender = "42NJMHTPFVPXVSDGA6JGKUV6TARV5UZTMPFIREMLXHETRKIVW34QFSDFRE"
        drr = self.dryrun_request(
            source,
            sender=sender,
            app=dict(
                app_idx=0,
                args=[
                    (0x01).to_bytes(1, byteorder="big"),
                    (0xFF).to_bytes(1, byteorder="big"),
                    (0x0100).to_bytes(2, byteorder="big"),
                    (0x01FF).to_bytes(2, byteorder="big"),
                ]
        ))
        self.assertPass(drr)

        value = dict(
            key="Creator",
            value=dict(action=1, bytes=base64.b64encode(decode_address(sender)).decode("utf-8"))
        )
        self.assertGlobalStateContains(drr, value)

        value = dict(key="RegBegin", value=dict(action=2, uint=0x01))
        self.assertGlobalStateContains(drr, value)

        value = dict(key="RegEnd", value=dict(action=2, uint=0xFF))
        self.assertGlobalStateContains(drr, value)

        value = dict(key="VoteBegin", value=dict(action=2, uint=0x0100))
        self.assertGlobalStateContains(drr, value)

        value = dict(key="VoteEnd", value=dict(action=2, uint=0x01FF))
        self.assertGlobalStateContains(drr, value)
```

Two `assertReject` statements in the beginning on the test check prerequisites: application call in **creation** mode (`app_idx = 0`) and number of required initialization parameters.
Then `dryrun_request` helper is used to obtain execution result and written values with `assertGlobalStateContains`. Changes are reported as `EvalDelta` type with `key`, `value`, `action`, `uint` or `bytes` properties. `bytes` values are base64-encoded, and `action` is explained in the table below:

| action | description     |
|--------|-----------------|
| 1      | set bytes value |
| 2      | set uint value  |
| 3      | delete value    |

Having this information, `assertGlobalStateContains` validates that **Creator** global key is set to txn sender address, and all the **RegBegin**, **RegEnd**, **VoteBegin** and **VoteEnd** are properly initialized.

## Testing `Application` with existing global state

In this example let's verify correctness of application update. **Voting app** uses transaction sender field to authenticate application creator that is written in the global state on creation.
In order to check this behavior we supply **global state** using `global_state` keyword in app parameters. Each entry must have `TealKeyValue` type that consist of string `key` and `TealValue` value. Table below explains this datatype:

| type | bytes | uint  | description |
|------|-------|-------|-------------|
| 1    | value | 0     | bytes value |
| 2    | empty | value | uint value  |

To pass the update check in the our **voting app**, the **global state** needs "Creator" key and some pre-defined value that matches txn sender.

```python tab="Python"
    def test_app_global_state_existing(self):
        """Use voting app as example to check app update"""
        source = """#pragma version 2
int 0
txn ApplicationID
==
bz not_creation
// fail on creation in this test scenario
int 0
return
not_creation:
int UpdateApplication
txn OnCompletion
==
bz failed
byte "Creator"
app_global_get
txn Sender
==
bz failed
int 1
return
failed:
int 0
"""
        sender = self.default_address()

        self.assertReject(source, app=dict(app_idx=0))
        self.assertReject(source, app=dict(app_idx=1))
        self.assertReject(
            source,
            app=dict(
                app_idx=1,
                accounts=[sender]
        ))

        app = dict(
            app_idx=1,
            global_state=[TealKeyValue(
                key="Creator",
                value=TealValue(type=1, bytes=b""),
            )]
        )
        self.assertReject(source, app=app)

        app["on_complete"] = transaction.OnComplete.UpdateApplicationOC
        self.assertReject(source, app=app)

        app["global_state"][0].value.bytes = decode_address(sender)
        self.assertPass(source, app=app)
```

Initial calls runs with `assertReject` ensure the program fails if pre-conditions not met: not an update call or missed values in the global state. The last `assertPass` run under the correct conditions:
1. tnx.on_complete = `UpdateApplicationOC`
2. txn.index != 0 (not an app create call)
3. global state of the app has "Creator" key with the value matching tnx sender.

!!! info
    If sender is not set, then default zero address is used.


For more sophisticated testing scenarios see **Testing with transaction objects** section below.

## Testing `Application` local state changes

It's time to check **voting** function. The app reads and writes own **local state** in sender's **balance record**. When someone **opts-in** to an app, local state is allocated and can be used. For safety operations the app needs to ensure the caller opted in before with `app_opted_in` opcode.
To make this work in tests, we supply balance records using `accounts` keyword in app parameters. The account can be either a `string` address or `Account` type. In the former case test harness creates an empty `Account` entry.
To pass the update check in the our **voting app**, two records need to exist:
1. App creator's record with global values `VoteBegin` and `VoteEnds`
2. Caller's record with pre-allocated local state for this app.

Look at the example below. It adds more state on every step to ensure the app passes only if all pre-condition are met.

```python tab="Python"
    def test_app_local_state(self):
        """Use voting app as example to check local state writes"""
        source = """#pragma version 2
txna ApplicationArgs 0
byte "vote"
==
bnz vote
int 0
return
vote:
global Round
byte "VoteBegin"
app_global_get
>=
global Round
byte "VoteEnd"
app_global_get
<=
&&
bz failed
int 0
txn ApplicationID
app_opted_in
bz failed
int 0
txn ApplicationID
byte "voted"
app_local_get_ex
bnz voted
//read existing vote candidate
txna ApplicationArgs 1
app_global_get
bnz increment_existing
pop
int 0
increment_existing:
int 1
+
store 1
txna ApplicationArgs 1
load 1
app_global_put
int 0 //sender
byte "voted"
txna ApplicationArgs 1
app_local_put
int 1
return
voted:
pop
int 1
return
failed:
int 0
return
"""
        drr = self.dryrun_request(source, app=dict(app_idx=1))
        self.assertReject(drr)
        self.assertError(drr, "invalid ApplicationArgs index 0")

        drr = self.dryrun_request(source, app=dict(app_idx=1, args=[b"vote"]))
        self.assertReject(drr)
        self.assertNoError(drr)

        sender = "42NJMHTPFVPXVSDGA6JGKUV6TARV5UZTMPFIREMLXHETRKIVW34QFSDFRE"
        creator = "DFPKC2SJP3OTFVJFMCD356YB7BOT4SJZTGWLIPPFEWL3ZABUFLTOY6ILYE"
        creator_data = Account(
            address=creator,
            status="Offline",
            created_apps=[Application(
                id=1,
                params=ApplicationParams(
                    global_state=[
                        TealKeyValue(
                            key="VoteBegin",
                            value=TealValue(type=2, uint=1)),
                        TealKeyValue(
                            key="VoteEnd",
                            value=TealValue(type=2, uint=1000)),
                    ]
                )
            )]
        )

        accounts = [creator_data]

        drr = self.dryrun_request(source, app=dict(
            app_idx=1,
            args=[b"vote"],
            round=3,
            creator=creator,
            accounts=accounts,
        ))
        self.assertReject(drr)
        self.assertNoError(drr)

        sender_data = Account(
            address=sender,
            status="Offline",
            apps_local_state=[
                ApplicationLocalState(
                    id=1
                )
            ]
        )

        accounts = [creator_data, sender_data]
        drr = self.dryrun_request(source, sender=sender, app=dict(
            app_idx=1,
            creator=creator,
            args=[b"vote"],
            round=3,
            accounts=accounts,
        ))
        self.assertError(drr, "invalid ApplicationArgs index 1")

        accounts = [creator_data, sender_data]
        drr = self.dryrun_request(source, sender=sender, app=dict(
            app_idx=1,
            creator=creator,
            args=[b"vote", "test"],
            round=3,
            accounts=accounts,
        ))
        self.assertPass(drr)

        value = dict(
            key="voted",
            value=Helper.build_bytes_delta_value("test")
        )
        self.assertLocalStateContains(drr, sender, value)
        with self.assertRaises(AssertionError):
            self.assertLocalStateContains(drr, creator, value)

        value = dict(
            key="test",
            value=dict(action=2, uint=1)
        )
        self.assertGlobalStateContains(drr, value)
```

At the beginning of the test `assertReject` calls the `id=1` without any arguments and ensure it fails on parameters access. Next, the test supplies `vote` argument. This time no errors, and the call rejected because there is not enough state. We create a **balance record** as in `creator_data` variable saying it's app creator by providing `address=creator` and `created_apps=[Application()]` parameters. Then we call the app with:

```python tab="Python"
        drr = self.dryrun_request(source, app=dict(
            app_idx=1,
            args=[b"vote"],
            round=3,
            creator=creator,
            accounts=accounts,
        ))
```

Important to note, `app_idx` equals to application id in `created_apps` in the balance record initialized above. `creator` parameter matches to `address` in the balance record, current execution `round = 3` is set in the interval `[VoteBegin, VoteEnd]` from the global state, and finally, balance records are provided to the executor with `accounts` argument.

The program is still rejected but now for different reason. It can be easily seen with `Helper.pprint(drr)`. Now the app needs sender's local state, that is later set as `sender_data` variable.

At the very end of the test `assertLocalStateContains` and `assertGlobalStateContains` verify that local and global states have been properly updated:

```python tab="Python"
        value = dict(
            key="voted",
            value=Helper.build_bytes_delta_value("test")
        )
        self.assertLocalStateContains(drr, sender, value)

        value = dict(
            key="test",
            value=dict(action=2, uint=1)
        )
        self.assertGlobalStateContains(drr, value)
```

## Testing with transaction objects

Although examples above provide testing tools for "create and run" scenarios when testing single programs, sometimes transaction interactions also need to tested.
In this example we consider how stateful application can offload some computations to stateless logicsig program, and ensure the logic sig is the right one.

Suppose our logic sig program computes the following hash `h(h(a1) + h(a2) + h(a3) + h(a4))` where `+` is string concatenation, and then verifies it against some provided proof.
Suppose our application approves only if the calculation is made correctly. To achieve this, create a txn group where:
1. Txn 1 is an escrow logic sig txn with hash calculation approval program (see `logic_source below`).
2. Txn 2 is app call txn with txn 1 checker (see `app_source` below).
    * Ensure txn 1 sender is know-ahead escrow address.
    * Ensure txn 1 Note field is set to proof that is needed to be confirmed.
3. Input data `a1`, `a2`, `a3`, `a4` are set as `ApplicationArgs` for txn 2 and accessed from txn 1 (logic sig args can be used as well, since both the logic hash and the proof checked in the app call program).

```python tab="Python"
    def test_transactions(self):
        """Test app call and logic sig transactions interaction"""
        logic_source = """#pragma version 2
gtxna 1 ApplicationArgs 0
sha512_256
gtxna 1 ApplicationArgs 1
sha512_256
concat
gtxna 1 ApplicationArgs 2
sha512_256
gtxna 1 ApplicationArgs 3
sha512_256
concat
concat
sha512_256
txn Note
==
"""
        # compile the logic sig program
        logic_compiled = self.algo_client.compile(logic_source)
        self.assertIn("hash", logic_compiled)
        self.assertIn("result", logic_compiled)
        logic = base64.b64decode(logic_compiled["result"])
        logic_hash = logic_compiled["hash"]

        # compute proof from parameters
        args = [b"this", b"is", b"a", b"test"]
        parts = []
        for arg in args:
            parts.append(checksum(arg))

        proof = checksum(b''.join(parts))

        # create and compile app call program
        app_source = f"""#pragma version 2
gtxn 0 Sender
addr {logic_hash}
==
gtxn 0 Note
byte {"0x" + proof.hex()}
==
&&
"""
        app_compiled = self.algo_client.compile(app_source)
        self.assertIn("result", app_compiled)
        app = base64.b64decode(app_compiled["result"])

        # create transactions
        txn1 = Helper.sample_txn(logic_hash, payment_txn)
        txn1.note = proof
        logicsig = transaction.LogicSig(logic, None)
        stxn1 = transaction.LogicSigTransaction(txn1, logicsig)

        app_idx = 1
        txn2 = Helper.sample_txn(self.default_address(), appcall_txn)
        txn2.index = app_idx
        txn2.app_args = args
        stxn2 = transaction.SignedTransaction(txn2, None)

        # create a balance record with the application
        # creator address is a random one
        creator = "DFPKC2SJP3OTFVJFMCD356YB7BOT4SJZTGWLIPPFEWL3ZABUFLTOY6ILYE"
        creator_data = Account(
            address=creator,
            status="Offline",
            created_apps=[Application(
                id=1,
                params=ApplicationParams(
                    approval_program=app,
                    local_state_schema=ApplicationStateSchema(64, 64),
                    global_state_schema=ApplicationStateSchema(64, 64),
                )
            )]
        )

        drr = self.dryrun_request_from_txn(
            [stxn1, stxn2],
            app=dict(
                accounts=[creator_data]
        ))
        self.assertPass(drr)

        # now check the verification logic
        # wrong creator
        txn1.sender = creator
        drr = self.dryrun_request_from_txn(
            [stxn1, stxn2],
            app=dict(
                accounts=[creator_data]
        ))
        self.assertPass(drr, txn_index=0)
        self.assertReject(drr, txn_index=1)

        # wrong proof
        txn1.sender = logic_hash
        txn1.note = b'wrong'
        drr = self.dryrun_request_from_txn(
            [stxn1, stxn2],
            app=dict(
                accounts=[creator_data]
        ))
        self.assertReject(drr, txn_index=0)
        self.assertReject(drr, txn_index=1)
        self.assertReject(drr)
```

## Debugging failed tests

### Trace pretty print

Let's take a program from `test_app_local_state` and print its execution trace with `Helper.pprint`.

```python tab="Python"
        # ... same as above
        accounts = [creator_data, sender_data]
        drr = self.dryrun_request(source, sender=sender, app=dict(
            app_idx=1,
            creator=creator,
            args=[b"vote", "test"],
            round=3,
            accounts=accounts,
        ))
        self.assertPass(drr)
        Helper.pprint(drr)  # print the trace
```

When the code executed it prints out:
```
txn[0] messages:
ApprovalProgram
PASS
txn[0] trace:
   1 (0001): intcblock 0 1             []
   2 (0005): bytecblock 0x766f7465 0x566f7465426567696e 0x566f7465456e64 0x766f746564 []
   3 (0036): txn ApplicationArgs 0     []
   4 (0039): bytec_0                   ["vote"]
   5 (0040): ==                        ["vote" "vote"]
   6 (0041): bnz label1                [1]
  10 (0046): global Round              []
  11 (0048): bytec_1                   [3]
  12 (0049): app_global_get            [3 "VoteBegin"]
  13 (0050): >=                        [3 1]
  14 (0051): global Round              [1]
  15 (0053): bytec_2                   [1 3]
  16 (0054): app_global_get            [1 3 "VoteEnd"]
  17 (0055): <=                        [1 3 1000]
  18 (0056): &&                        [1 1]
  19 (0057): bz label2                 [1]
  20 (0060): intc_0                    []
  21 (0061): txn ApplicationID         [0]
  22 (0063): app_opted_in              [0 1]
  23 (0064): bz label2                 [1]
  24 (0067): intc_0                    []
  25 (0068): txn ApplicationID         [0]
  26 (0070): bytec_3                   [0 1]
  27 (0071): app_local_get_ex          [0 1 "voted"]
  28 (0072): bnz label3                [0 0]
  29 (0075): txn ApplicationArgs 1     [0]
  30 (0078): app_global_get            [0 "test"]
  31 (0079): bnz label4                [0 0]
  32 (0082): pop                       [0]
  33 (0083): intc_0                    []
  35 (0084): intc_1                    [0]
  36 (0085): +                         [0 1]
  37 (0086): store 1                   [1]
  38 (0088): txn ApplicationArgs 1     []
  39 (0091): load 1                    ["test"]
  40 (0093): app_global_put            ["test" 1]
  41 (0094): intc_0                    []
  42 (0095): bytec_3                   [0]
  43 (0096): txn ApplicationArgs 1     [0 "voted"]
  44 (0099): app_local_put             [0 "voted" "test"]
  45 (0100): intc_1                    []
  46 (0101): return                    [1]
  54 (0107):                           [1]
```

The first column contains line numbers in the disassembly, the second - `pc` values, the third - opcode and arguments, and the last - stack values (top on right) _before_ the opcode execution.

### Use debugger

Again, let's use the program from `test_app_local_state`. This time we use `Helper.build_dryrun_request` (with the same interface as `dryrun_request`) to create request and `Helper.save_dryrun_request` to save to a file:

```python tab="Python"
        # ... same as above but replaced self.dryrun_request by Helper.build_dryrun_request
        accounts = [creator_data, sender_data]
        req = Helper.build_dryrun_request(source, sender=sender, app=dict(
            app_idx=1,
            creator=creator,
            args=[b"vote", "test"],
            round=3,
            accounts=accounts,
        ))
        Helper.save_dryrun_request("dr.msgp", req)
```

Then run the debugger with the command below and open the `devtools://` url in Google Chrome:

```bash
tealdbg debug --dryrun-req dr.msgp

 Using proto: https://github.com/algorandfoundation/specs/tree/3a83c4c743f8b17adfd73944b4319c25722a6782
 ------------------------------------------------
 CDT debugger listening on: ws://localhost:9392/bbbe040c1529c7dbac116caa0687030c735b22fd674b1
 Or open in Chrome:
 devtools://devtools/bundled/js_app.html?experiments=true&v8only=false&abb56885fe7b73a6d5bbbe040c1529c7dbac116caa0687030c735b22fd674b1
 ------------------------------------------------
```

For more information about debugging TEAL read [the documentation](https://github.com/algorand/go-algorand/tree/master/cmd/tealdbg#algorand-teal-debugger) or TEAL debugging guide.