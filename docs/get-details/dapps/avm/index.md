title: AVM

The Algorand virtual machine (AVM) runs on every node in the Algorand blockchain. This virtual machine contains a stack engine that evaluates smart contracts and smart signatures against the transactions they're called with. These programs either fail and reject the transaction or succeed and apply changes according to the logic and contents of the transactions. 


# Quick start videos

If you prefer videos, take a look at this playlist to learn about AVM. Most of the videos in the list are under 5 minutes each.

<iframe width="100%" style="aspect-ratio:16/9" src="https://www.youtube-nocookie.com/embed/96pwBo5jqnk" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>

Smart contracts, also referred to as stateful smart contracts, contain logic that is deployed and can be remotely called from any node on the Algorand blockchain. These contracts are called by issuing an Application Call transaction. The AVM evaluates the contract logic against this transaction and ultimately results in success or failure. If a call to a smart contract fails, any changes produced by that call will not be committed to the blockchain. If the call is successful, the changes will be recorded to the blockchain when the block is committed. In addition to changes to balances, the logic within a smart contract can modify key/value data associated with the contract on a global or per-account basis. 

Smart signatures, also referred to as stateless contracts, contain logic that is used to sign transactions, commonly for signature delegation. The logic of the smart signature is submitted with the transaction. While the logic in the smart signature is stored on the chain as part of the transaction approval process, the logic is not remotely callable. Any new transaction that relies on the same smart signature would resubmit the logic. When the logic is submitted to a node the AVM evaluates the logic and results in success or failure. If a smart signature’s logic fails when executed by the AVM, the associated transaction will not be applied.

The AVM interprets an assembler-like language called [Transaction Execution Approval Language (TEAL)](teal/index.md). TEAL can be thought of as syntactic sugar for AVM bytecode and the [full specification](teal/specification.md) of the TEAL language is available in the developer documentation. In addition, a simple [overview guide](teal/index.md) explains many of the features of the language. TEAL programs are comprised of a set of operation codes (opcodes). These opcodes are used to implement the logic of smart contracts and smart signatures. The full list of [opcodes](teal/opcodes) is available in the developer documentation.  

While it is possible to write TEAL directly, a developer may prefer to use the PyTeal Python library, which provides a more familiar syntax. For more information on using PyTeal to write smart signatures and smart contracts see [Build with Python](/docs/get-details/dapps/writing-contracts/pyteal). 
