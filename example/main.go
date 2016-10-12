package main

import (
	"fmt"
	"time"

	"test/tracker"
	"math/rand"
)

func init() {
	ttracker.DefaultHandler = func(name, label string, timeout time.Duration) {
		fmt.Println(name, label, timeout)
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
