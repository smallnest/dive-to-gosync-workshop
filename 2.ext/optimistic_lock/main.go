package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/reugn/async"
)

func main() {

	var sv = &syncValue{l: async.NewOptimisticLock()}
	var wg sync.WaitGroup
	wg.Add(51)

	// read
	for i := 0; i < 50; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				fmt.Println("read: ", sv.read())
			}
			wg.Done()
		}()
	}

	//write
	go func() {
		for i := 0; i < 50; i++ {
			sv.incr()
			time.Sleep(time.Millisecond)
		}
		wg.Done()
	}()

	wg.Wait()
}

type syncValue struct {
	v int
	l *async.OptimisticLock
}

func (sv *syncValue) read() int {
	var v int
	ok := false
	for !ok {
		stamp := sv.l.OptLock()
		v = sv.v
		ok = sv.l.OptUnlock(stamp)
	}

	return v
}

func (sv *syncValue) incr() {
	sv.l.Lock()
	sv.v++
	sv.l.Unlock()
}
