package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// https://go101.org/article/concurrent-synchronization-more.html
// http://goinbigdata.com/golang-wait-for-all-goroutines-to-finish/
// https://golang.org/pkg/sync/#example_WaitGroup

// correct or wrong usage

func main() {
	var count int64
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		go func() {
			wg.Add(1)
			atomic.AddInt64(&count, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(atomic.LoadInt64(&count))
}
