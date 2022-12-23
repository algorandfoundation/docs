title: conduit exporters filewriter
---


### Config

<table>

<tr>

<th>key</th><th>type</th><th>description</th>



<tr><td>block-dir</td><td>string</td><td> <code>blocks-dir</code> is an optional path to a directory where block data should be

	stored.<br/>

	The directory is created if it doesn't exist.<br/>

	If no directory is provided the default plugin data directory is used.

</td></tr>



<tr><td>filename-pattern</td><td>string</td><td> <code>filename-pattern</code> is the format used to write block files. It uses go

	string formatting and should accept one number for the round.<br/>

	If the file has a '.gz' extension, blocks will be gzipped.

	Default:



		"%[1]d_block.json"

</td></tr>



<tr><td>drop-certificate</td><td>bool</td><td><code>drop-certificate</code> is used to remove the vote certificate from the block data before writing files.

</td></tr>

</table>



