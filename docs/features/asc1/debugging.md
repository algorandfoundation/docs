title: Smart Contract Debugging

# Writing Logic Signature Transaction for Debugging
The SDKs and `goal` command-line tool provide functionality to do a test run of a TEAL program. The `goal` command-line tool can be used with [stateless smart contracts.](stateless/walkthrough.md) 

TEAL programs can be written with any editor and are compiled using the goal command-line tool or [SDK](stateless/sdks.md). They can also be built using python with the [PyTeal](teal/pyteal.md) Library. 

The command-line tool provides the ability to use these compiled programs within transactions. The `goal clerk dryrun` command is described in the [goal TEAL Walkthrough](stateless/walkthrough.md) documentation. From the SDK a logic signature transaction can be written to a file to be used with the `goal clerk dryrun` command. The following code details how this is done. The goal tab illustrates run the `dryrun` on the generated file.


```javascript tab="JavaScript"
    let rawSignedTxn = algosdk.signLogicSigTransaction(txn, lsig);
    fs.writeFileSync("simple.stxn", rawSignedTxn.blob);
```

```python tab="Python"
    logicsig_txn = transaction.LogicSigTransaction(txn, lsig)
    transaction.write_to_file([logicsig_txn], "simple.stxn")
```

```java tab="Java"
    SignedTransaction stx = Account.signLogicsigTransaction(lsig, tx);
    byte[] outBytes = Encoder.encodeToMsgPack(stx);
    try {
        String FILEPATH = "./simple.stxn";
        File file = new File(FILEPATH);
        OutputStream os = new FileOutputStream(file);
        os.write(outBytes);
        os.close();
    }
    catch (Exception e) {
        System.out.println("Exception: " + e);
    }    
```

```go tab="Go"
	txid, stx, err := crypto.SignLogicsigTransaction(lsig, tx)
	if err != nil {
        ...
    }
    f, err := os.Create("simple.stxn")
    if err != nil {
        ...
    }
    defer f.Close()
    if _, err := f.Write(stx); err != nil {
        ...
    }
    if err := f.Sync(); err != nil {
        ...
    }    
```

```text tab="goal"
$ goal clerk dryrun -t simple.stxn
tx[0] cost=2 trace:
  1 intcblock => <empty stack>
  4 intc_0 => 0 0x0

REJECT
```


# Using Dryrun for Debugging a TEAL Program
When using the `goal` command-line tool, the `-o` option is used to write a signed transaction out to a file and the transaction will not be submitted to the network. This allows testing of the TEAL logic with the `goal clerk dryrun` command which shows how the TEAL is processed and approved or rejected.

SDK supports debugging with dryrun using source or the compiled TEAL program. The dryrun response includes disassembly, logic sig messages w PASS/REJECT, a sig trace, app call messages and app call trace.  

```javascript tab="JavaScript"
// dryrunDebugging returns a response with disassembly, logic-sig-messages w PASS/REJECT and sig-trace
async function dryrunDebugging(lsig, txn, data) {
    if (data == null)
    {
        //compile
        txns = [{
            lsig: lsig,
            txn: txn,
        }];        
    }
    else
    {
        // source
        txns = [{
            txn: txn,
        }];
        sources = [new algosdk.modelsv2.DryrunSource("lsig", data.toString("utf8"), 0)];
    }
    const dr = new algosdk.modelsv2.DryrunRequest({
        txns: txns,
        sources: sources,
    });
    dryrunResponse = await algodclient.dryrun(dr).do();
    return dryrunResponse;
}

// the dryrun response should look similar to this
// {
// "error": "",
// "protocol-version": "https://github.com/algorandfoundation/specs/tree/e5f565421d720c6f75cdd186f7098495caf9101f",
// "txns": [
// {
//     "disassembly": [
//         "// version 1",
//         "intcblock 123",
//         "arg_0",
//         "btoi",
//         "intc_0",
//         "==",
//         ""
//     ],
//     "logic-sig-messages": [
//         "PASS"
//     ],
//     "logic-sig-trace": [
//         {
//             "line": 1,
//             "pc": 1,
//             "stack": []
//         },
//         {
//             "line": 2,
//             "pc": 4,
//             "stack": []
//         },
//         {
//             "line": 3,
//             "pc": 5,
//             "stack": [
//                 {
//                     "bytes": "ew==",
//                     "type": 1,
//                     "uint": 0
//                 }
//             ]
//         },
//         {
//             "line": 4,
//             "pc": 6,
//             "stack": [
//                 {
//                     "bytes": "",
//                     "type": 2,
//                     "uint": 123
//                 }
//             ]
//         },
//         {
//             "line": 5,
//             "pc": 7,
//             "stack": [
//                 {
//                     "bytes": "",
//                     "type": 2,
//                     "uint": 123
//                 },
//                 {
//                     "bytes": "",
//                     "type": 2,
//                     "uint": 123
//                 }
//             ]
//         },
//         {
//             "line": 6,
//             "pc": 8,
//             "stack": [
//                 {
//                     "bytes": "",
//                     "type": 2,
//                     "uint": 1
//                 }
//             ]
//         }
//     ]
// }
// ]
// }
```

```python tab="Python"
# dryrun_debug returns a response with disassembly, logic-sig-messages w PASS/REJECT and sig-trace
# dryrun source if provided, else dryrun compiled
def dryrun_debug(lstx, mysource):
    sources = []
    if (mysource != None):
        # source
        sources = [DryrunSource(field_name="lsig", source=mysource, txn_index=0)]
    drr = DryrunRequest(txns=[lstx], sources=sources)
    dryrun_response = algod_client.dryrun(drr)
    return dryrun_response

# dryrun response should look similar to this
# {
# "error": "",
# "protocol-version": "https://github.com/algorandfoundation/specs/tree/e5f565421d720c6f75cdd186f7098495caf9101f",
# "txns": [
# {
# "disassembly": [
#     "// version 1",
#     "intcblock 123",
#     "arg_0",
#     "btoi",
#     "intc_0",
#     "==",
#     ""
# ],
# "logic-sig-messages": [
#     "PASS"
# ],
# "logic-sig-trace": [
#     {
#         "line": 1,
#         "pc": 1,
#         "stack": []
#     },
#     {
#         "line": 2,
#         "pc": 4,
#         "stack": []
#     },
#     {
#         "line": 3,
#         "pc": 5,
#         "stack": [
#             {
#                 "bytes": "AAAAAAAAAHs=",
#                 "type": 1,
#                 "uint": 0
#             }
#         ]
#     },
#     {
#         "line": 4,
#         "pc": 6,
#         "stack": [
#             {
#                 "bytes": "",
#                 "type": 2,
#                 "uint": 123
#             }
#         ]
#     },
#     {
#         "line": 5,
#         "pc": 7,
#         "stack": [
#             {
#                 "bytes": "",
#                 "type": 2,
#                 "uint": 123
#             },
#             {
#                 "bytes": "",
#                 "type": 2,
#                 "uint": 123
#             }
#         ]
#     },
#     {
#         "line": 6,
#         "pc": 8,
#         "stack": [
#             {
#                 "bytes": "",
#                 "type": 2,
#                 "uint": 1
#             }
#         ]
#     }
# ]
# }
# ]
# }
```

```java tab="Java"
// getDryrun returns a response with disassembly, logic-sig-messages w PASS/REJECT and sig-trace
private Response<DryrunResponse> getDryrunResponse(SignedTransaction stxn, byte[] source)
        throws Exception{
    List<DryrunSource> sources = new ArrayList<DryrunSource>();
    List<SignedTransaction> stxns = new ArrayList<SignedTransaction>();
    // compiled 
    if (source == null) {
        stxns.add(stxn);
    }
    // source
    else if (source != null) {
        DryrunSource drs = new DryrunSource();
        drs.fieldName = "lsig";
        drs.source = new String(source);
        drs.txnIndex = 0l;
        sources.add(drs);
        stxns.add(stxn);
    }
    Response<DryrunResponse> dryrunResponse;
    DryrunRequest dr = new DryrunRequest();
    dr.txns = stxns;
    dr.sources = sources;
    dryrunResponse = client.TealDryrun().request(dr).execute();
    return dryrunResponse;
} 
// response should look similar to this
// Dryrun repsonse : {
//   "protocol-version": "https://github.com/algorandfoundation/specs/tree/e5f565421d720c6f75cdd186f7098495caf9101f",
//   "error": "",
//   "txns": [{
//     "global-delta": [],
//     "local-deltas": [],
//     "disassembly": [
//       "// version 1",
//       "intcblock 123",
//       "arg_0",
//       "btoi",
//       "intc_0",
//       "==",
//       ""
//     ],
//     "logic-sig-messages": ["PASS"],
//     "app-call-messages": [],
//     "app-call-trace": [],
//     "logic-sig-trace": [
//       {
//         "stack": [],
//         "pc": 1,
//         "line": 1,
//         "scratch": []
//       },
//       {
//         "stack": [],
//         "pc": 4,
//         "line": 2,
//         "scratch": []
//       },
//       {
//         "stack": [{
//           "bytes": "ew==",
//           "type": 1,
//           "uint": 0
//         }],
//         "pc": 5,
//         "line": 3,
//         "scratch": []
//       },
//       {
//         "stack": [{
//           "bytes": "",
//           "type": 2,
//           "uint": 123
//         }],
//         "pc": 6,
//         "line": 4,
//         "scratch": []
//       },
//       {
//         "stack": [
//           {
//             "bytes": "",
//             "type": 2,
//             "uint": 123
//           },
//           {
//             "bytes": "",
//             "type": 2,
//             "uint": 123
//           }
//         ],
//         "pc": 7,
//         "line": 5,
//         "scratch": []
//       },
//       {
//         "stack": [{
//           "bytes": "",
//           "type": 2,
//           "uint": 1
//         }],
//         "pc": 8,
//         "line": 6,
//         "scratch": []
//       }
//     ]
//   }]
// }
```

```go tab="Go"
// dryrunDebugging returns a response with disassembly, logic-sig-messages w PASS/REJECT and sig-trace
func dryrunDebugging(lsig types.LogicSig, args [][]byte,tealFile []byte, client *algod.Client) modelsV2.DryrunResponse {
    // if tealFile is nil, use compile
	txns := []types.SignedTxn{{}}
    sources := []modelsV2.DryrunSource{}
    if (tealFile == nil){
        // complile
    	txns[0].Lsig = lsig    
    }else{
        // source
        txns[0].Lsig.Args = args
		sources = append(sources, modelsV2.DryrunSource{
			FieldName: "lsig",
			Source:    string(tealFile),
            TxnIndex:  0,
        })
    }

	ddr := modelsV2.DryrunRequest{
		Txns:    txns,
		Sources: sources,
	}

	result, err := client.TealDryrun(ddr).Do(context.Background())
	if err != nil {
		return result
    }
    return  result
}
// the response should look similar to this...
// Result = {
// "protocol-version": "https://github.com/algorandfoundation/specs/tree/e5f565421d720c6f75cdd186f7098495caf9101f",
// "txns": [
// {
//     "disassembly": [
//         "// version 1",
//         "intcblock 123",
//         "arg_0",
//         "btoi",
//         "intc_0",
//         "==",
//         ""
//     ],
//     "logic-sig-messages": [
//         "PASS"
//     ],
//     "logic-sig-trace": [
//         {
//             "line": 1,
//             "pc": 1
//         },
//         {
//             "line": 2,
//             "pc": 4
//         },
//         {
//             "line": 3,
//             "pc": 5,
//             "stack": [
//                 {
//                     "bytes": "AAAAAAAAAHs=",
//                     "type": 1
//                 }
//             ]
//         },
//         {
//             "line": 4,
//             "pc": 6,
//             "stack": [
//                 {
//                     "type": 2,
//                     "uint": 123
//                 }
//             ]
//         },
//         {
//             "line": 5,
//             "pc": 7,
//             "stack": [
//                 {
//                     "type": 2,
//                     "uint": 123
//                 },
//                 {
//                     "type": 2,
//                     "uint": 123
//                 }
//             ]
//         },
//         {
//             "line": 6,
//             "pc": 8,
//             "stack": [
//                 {
//                     "type": 2,
//                     "uint": 1
//                 }
//             ]
//         }
//     ]
// }

// ]
// }   
```

```text tab="goal"
$ goal clerk send -f C3MKH24QL3GHSD5CDQ47ZNQZMNZRX4MUTV6LVPAXMWAXMIISYSOWPGH674 -a 1000000 -t STF6TH6PKINM4CDIQHNSC7QEA4DM5OJKKSACAPWGTG776NWSQOMAYVGOQE -L mydelegatedsig.lsig -d ~/node/data -o out.stxn
$ goal clerk dryrun -t out.stxn
tx[0] cost=2 trace:
  1 intcblock => <empty stack>
  4 intc_0 => 0 0x0

REJECT
```


!!! info
    The example code snippets are provided throughout this page and are abbreviated for conciseness and clarity. Full running code examples for each SDK are available within the GitHub repo at [/examples/smart_contracts](https://github.com/algorand/docs/tree/master/examples/smart_contracts) and for [download](https://github.com/algorand/docs/blob/master/examples/smart_contracts/smart_contracts.zip?raw=true) (.zip)


