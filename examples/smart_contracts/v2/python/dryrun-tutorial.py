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


if __name__ == "__main__":
    unittest.main()
