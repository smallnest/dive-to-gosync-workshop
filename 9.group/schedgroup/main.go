package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/mdlayher/schedgroup"
)

func main() {
	sg := schedgroup.New(context.Background())

	n := 5
	spread := 100 * time.Millisecond

	timeC := make(chan time.Time, n)

	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()

	go func() {
		defer func() {
			close(timeC)
			wg.Done()
		}()

		// Produce the current time when a task is fired.
		for i := 0; i < n; i++ {
			sg.Delay(time.Duration(i+1)*spread, func() {
				timeC <- time.Now()
			})
		}

		if err := sg.Wait(); err != nil {
			panicf("failed to wait: %v", err)
		}
	}()

	var last time.Time
	var recv int

	for tv := range timeC {
		recv++

		if !last.IsZero() {
			diff := tv.Sub(last)
			fmt.Printf("executed a task since %v\n", diff)
		}

		last = tv
	}
}

func panicf(format string, a ...any) {
	panic(fmt.Sprintf(format, a...))
}
