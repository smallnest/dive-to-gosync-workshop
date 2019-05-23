package main

import (
	"sync"
	"time"
)

// https://groups.google.com/forum/#!topic/golang-nuts/bv_Qac26fDc
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		wg.Add(1)
	}()
	wg.Wait()
}
