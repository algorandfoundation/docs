title: conduit importers filereader
---


### Config

<table>

<tr>

<th>key</th><th>type</th><th>description</th>



<tr><td>block-dir</td><td>string</td><td><code>block-dir</code> is the path to a directory where block data is stored.

</td></tr>



<tr><td>retry-duration</td><td>time.Duration</td><td> <code>retry-duration</code> controls the delay between checks when the importer has caught up and is waiting for new blocks to appear.<br/>

	The input duration will be interpreted in nanoseconds.

</td></tr>



<tr><td>retry-count</td><td>uint64</td><td> <code>retry-count</code> controls the number of times to check for a missing block

	before generating an error. The retry count and retry duration should

	be configured according the expected round time.

</td></tr>



<tr><td>filename-pattern</td><td>string</td><td> <code>filename-pattern</code> is the format used to find block files. It uses go string formatting and should accept one number for the round.

	The default pattern is



	"%[1]d_block.json"

</td></tr>

</table>



