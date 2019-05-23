package main

import (
	"sync"
)

// https://groups.google.com/forum/#!topic/golang-nuts/bv_Qac26fDc
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		go func() {
			for {
				wg.Add(1)
				wg.Done()
			}
		}()
	}
	for i := 0; i < 100; i++ {
		go func() {
			for {
				wg.Wait()
			}
		}()
	}

	select {}
}
