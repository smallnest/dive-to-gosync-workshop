package main

import (
	"fmt"
	"sync"
	"time"
)

func or(chans ...<-chan any) <-chan any {
	out := make(chan any)
	go func() {
		var once sync.Once
		for _, c := range chans {
			go func(c <-chan any) {
				select {
				case <-c:
					once.Do(func() { close(out) })
				case <-out:
				}
			}(c)
		}
	}()
	return out
}

func main() {
	sig := func(after time.Duration) <-chan any {
		c := make(chan any)
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()

	<-or(
		sig(10*time.Second),
		sig(10*time.Second),
		sig(10*time.Second),
		sig(10*time.Second),
		sig(10*time.Second),
		sig(01*time.Minute),
	)

	fmt.Printf("done after %v", time.Since(start))
}
