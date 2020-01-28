title: kmd
---
## kmd



Key Management Daemon (kmd)



### Synopsis



The Key Management Daemon (kmd) is a low level wallet and key management

tool. It works in conjunction with algod and goal to keep secrets safe. An

optional timeout flag will automatically terminate kmd after a number of

seconds has elapsed, allowing a simple way to ensure kmd will be shutdown in

a timely manner. This is a blocking command.



```

kmd [flags]

```



### Options



```

  -d, --data-dir string    kmd data directory.

  -h, --help               help for kmd

  -t, --timout-secs uint   Number of seconds that kmd will run for before termination.

```




