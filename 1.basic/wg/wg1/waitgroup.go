package main

import (
	"fmt"
	"sync"
	"time"
)

// https://go101.org/article/concurrent-synchronization-more.html
// http://goinbigdata.com/golang-wait-for-all-goroutines-to-finish/
// https://golang.org/pkg/sync/#example_WaitGroup

// correct or wrong usage

type Counter struct {
	mu    sync.Mutex
	count uint64
}

func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func worker(c *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	c.Incr()
}

func main() {
	var counter Counter
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go worker(&counter, &wg)
	}
	wg.Wait()

	fmt.Println(counter.Count())
}
