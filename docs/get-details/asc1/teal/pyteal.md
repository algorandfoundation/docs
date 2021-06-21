title: PyTeal Overview

[PyTeal](https://github.com/algorand/pyteal) is a python language binding for Algorand Smart Contracts (ASC) that abstracts away the complexities in writing smart contracts. 

PyTeal allows stateless and stateful smart contracts to be written in python and then compiled to TEAL. Note that the TEAL code is not automatically compiled to byte code, but that can be done with any of the SDKs, including Python. The TEAL can also be compiled using the `goal` command-line tool.

See the complete code [documentation for PyTeal](https://pyteal.readthedocs.io/en/latest/) for more information.

# Installing PyTeal
Complete install instructions are available in the PyTeal [documentation](https://pyteal.readthedocs.io/en/latest/). Run the following to install PyTeal with the Python package manager.

```python
$ pip3 install pyteal --upgrade
```

# Simple PyTeal Example
As an example, a simple stateless smart contract can be created with PyTeal. This example creates a stateless contract account that only allows withdrawals from a specific account. It also checks to verify that the `RekeyTo`, `CloseRemainderTo`, and `AssetCloseTo` transaction properties are not set, to prevent malicious transactions (see [guidelines](../../../reference/teal/guidelines)) for more details.

```python
# simple.py
#!/usr/bin/env python3
from pyteal import *

def bank_for_account(receiver):
    """Only allow receiver to withdraw funds from this contract account."""
    is_payment = Txn.type_enum() == Int(1)
    is_correct_receiver = Txn.receiver() == Addr(receiver)
    is_one_tx = Global.group_size() == Int(1)

    safety_cond = And(
        Txn.asset_close_to() == Global.zero_address(),
        Txn.close_remainder_to() == Global.zero_address(),
        Txn.rekey_to() == Global.zero_address(),
    )

    return And(is_payment, is_correct_receiver, is_one_tx, safety_cond)

program = bank_for_account("ZZAF5ARA4MEC5PVDOP64JM5O5MQST63Q2KOY2FLYFLXXD3PFSNJJBYAFZM")
print(compileTeal(program, Mode.Signature, version=3))
```

When executed this python will return the following TEAL.

```bash
$ python3 simple.py
#pragma version 3
txn TypeEnum
int 1
==
txn Receiver
addr ZZAF5ARA4MEC5PVDOP64JM5O5MQST63Q2KOY2FLYFLXXD3PFSNJJBYAFZM
==
&&
global GroupSize
int 1
==
&&
txn AssetCloseTo
global ZeroAddress
==
txn CloseRemainderTo
global ZeroAddress
==
&&
txn RekeyTo
global ZeroAddress
==
&&
&&
```

This TEAL code can then be compiled and used with either SDKs or the `goal` command-line tools.
