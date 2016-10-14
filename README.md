Timeout tracker
===============

TTracker is made to find breakdown (by timings) in any parts of your code 

Usage example
-------------

For example, you have program which is going to request database and 
 then is making complexity calculations. You have to know where you will 
 be catching latency for calls that are more than 100ms

```go
// You have to define default handler for breakdown timeouts
func init() {
	ttracker.DefaultHandler = func(tracker *ttracker.TimeoutTracker) {
		for _, t := range tracker.Timeouts {
			fmt.Println(t.Label, t.Duration, t.Elapsed)
		}
	}
}

func handleRequest() {
    // Start timing tracking process
    t := ttracker.Start("my_tracker", 100*time.Millisecond)

    // ...goto remote storage, for example
    
    t.Track("storage")
    
    // ...make complex calculations
    
    t.Track("calculations")
    
    // Stop timing tracking process
    t.Stop()
}
```

This code will write to standard output tracker name, tracking label and 
 label timeout for each `handleRequest` call which more than 100ms
Example output will be like this

```
my_tracker storage 85.147435ms
my_tracker calculations 176.854402ms
```

See the source code [here](example/main.go)

Benchmarks & Tests
------------------

```
$: go test
OK: 4 passed
PASS
ok  	test/tracker	0.042s
$: go test -check.b -check.bmem
PASS: ttracker_test.go:74: TrackerSuite.BenchmarkTracking	 5000000	       596 ns/op	      98 B/op	       0 allocs/op
OK: 1 passed
PASS
ok  	test/tracker	4.162s
```

License
-------

Apache 2.0
