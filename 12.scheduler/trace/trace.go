package main

import (
	"os"
	"runtime/trace"
	"sync"
)

// go tool trace trace.out
func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			a := 0

			for i := 0; i < 1e7; i++ {
				a += 1
			}

			wg.Done()
		}()
	}

	wg.Wait()
}
