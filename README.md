# learn-io

follow [Ardan Labs](https://github.com/ardanlabs/gotraining/blob/master/topics/go/profiling/memcpu/stream.go)


```shell script
$ go test -run none -bench . -benchtime 3s -benchmem -cpuprofile p.out
$ go tool pprof p.out
(pprof) list method1
(pprof) web list method1
$ go tool pprof -http :3000 p.out
$ go tool pprof -http :3000 memcpu.test p.out
```

```shell script
$ go test -run none -bench . -benchtime 3s -benchmem -memprofile p.out
$ go tool pprof -<OPTIONAL_PICK_MEM_PROFILE> p.out
$ go tool pprof -<OPTIONAL_PICK_MEM_PROFILE> -http :3000 memcpu.test p.out

// Useful to see pressure on heap over time.
-alloc_space  : All allocations happened since program start ** default
-alloc_objects: Number of object allocated at the time of profile

// Useful to see current status of heap.
-inuse_space  : Allocations live at the time of profile
-inuse_objects: Number of bytes allocated at the time of profile

```