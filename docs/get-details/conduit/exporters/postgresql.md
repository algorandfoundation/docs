title: conduit exporters postgresql
---


### ExporterConfig

<table>

<tr>

<th>key</th><th>type</th><th>description</th>



<tr><td>connection-string</td><td>string</td><td> <code>connectionstring</code> is the Postgresql connection string<br/>

	See https://github.com/jackc/pgconn for more details

</td></tr>



<tr><td>max-conn</td><td>uint32</td><td> <code>max-conn</code> specifies the maximum connection number for the connection pool.<br/>

	This means the total number of active queries that can be running concurrently can never be more than this.

</td></tr>



<tr><td>test</td><td>bool</td><td> <code>test</code> will replace an actual DB connection being created via the connection string,

	with a mock DB for unit testing.

</td></tr>



<tr><td>delete-task</td><td>util.PruneConfigurations</td><td><code>delete-task</code> is the configuration for data pruning.

</td></tr>

</table>



