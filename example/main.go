package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/alxmsl/ttracker"
)

func init() {
	ttracker.DefaultHandler = func(tracker *ttracker.TimeoutTracker) {
		fmt.Println(tracker.Name, "tracker timings:")
		for _, t := range tracker.Timeouts {
			fmt.Println(t.Label, t.Duration, t.Elapsed)
		}
	}
}

func main() {
	t := ttracker.Start("my_tracker", 100*time.Millisecond)

	// ...goto remote storage, for example
	time.Sleep(time.Duration(rand.Int31n(100)) * time.Millisecond)

	t.Track("storage")

	// ...make complex calculations
	time.Sleep(time.Duration(rand.Int31n(100)) * time.Millisecond)

	t.Track("calculations")

	// Stop timing tracking process
	t.Stop()
}
