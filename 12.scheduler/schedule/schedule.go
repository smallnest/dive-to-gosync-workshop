package main

import "sync"

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {
			a := 0

			for i := 0; i < 1e8; i++ {
				a += 1
			}

			wg.Done()
		}()
	}

	wg.Wait()
}
