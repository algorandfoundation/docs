title: conduit processors blockevaluator
---
## The <code>block_processor</code> Plugin

This plugin runs a local ledger, processing blocks passed to it, and adding



### Config

<table>

<tr>

<th>key</th><th>type</th><th>description</th>



<tr><td>catchpoint</td><td>string</td><td> <code>catchpoint</code> to initialize the local ledger to.<br/>

	For more data on ledger catchpoints, see the

	<a hre=https://developer.algorand.org/docs/run-a-node/operations/catchup/>Algorand developer docs</a>

</td></tr>



<tr><td>ledger-dir</td><td>string</td><td><code>ledger-dir</code> is the directory which contains the ledger.

</td></tr>



<tr><td>algod-data-dir</td><td>string</td><td><code>algod-data-dir</code> is the algod data directory.

</td></tr>



<tr><td>algod-token</td><td>string</td><td><code>algod-token</code> is the API token for Algod usage.

</td></tr>



<tr><td>algod-addr</td><td>string</td><td><code>algod-addr</code> is the address of the Algod server

</td></tr>

</table>





### Example Configs



```yaml

config:

  - any:

    - tag: ""

	  expression: ""

	  expression-type: ""

```

