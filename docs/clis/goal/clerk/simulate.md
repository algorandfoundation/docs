title: goal clerk simulate
---
## goal clerk simulate



Simulate a transaction or transaction group with algod's simulate REST endpoint



### Synopsis



Simulate a transaction or transaction group with algod's simulate REST endpoint under various configurations.



```

goal clerk simulate [flags]

```



### Options



```

      --allow-empty-signatures     Allow transactions without signatures to be simulated as if they had correct signatures

      --allow-more-logging         Lift the limits on log opcode during simulation

      --allow-more-opcode-budget   Apply max extra opcode budget for apps per transaction group (default 320000) during simulation

      --allow-unnamed-resources    Allow access to unnamed resources during simulation

      --extra-opcode-budget uint   Apply extra opcode budget for apps per transaction group during simulation

      --full-trace                 Enable all options for simulation execution trace

  -h, --help                       help for simulate

      --request string             Simulate request object to run. Mutually exclusive with --txfile

      --request-only-out string    Filename for writing simulate request object. If provided, the command will only write the request object and exit. No simulation will happen

  -o, --result-out string          Filename for writing simulation result

      --scratch                    Report scratch change during simulation time

      --stack                      Report stack change during simulation time

      --state                      Report application state changes during simulation time

      --trace                      Enable simulation time execution trace of app calls

  -t, --txfile string              Transaction or transaction-group to test. Mutually exclusive with --request

```



### Options inherited from parent commands



```

  -d, --datadir stringArray   Data directory for the node

  -k, --kmddir string         Data directory for kmd

  -w, --wallet string         Set the wallet to be used for the selected operation

```



### SEE ALSO



* [goal clerk](../../clerk/clerk/)	 - Provides the tools to control transactions 



