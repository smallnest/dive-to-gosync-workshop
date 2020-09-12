package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	dosomething(100, &wg)
	dosomething(110, &wg)
	dosomething(120, &wg)
	dosomething(130, &wg)

	wg.Wait()
	fmt.Println("Done")
}

func dosomething(millisecs time.Duration, wg *sync.WaitGroup) {
	wg.Add(1)

	go func() {
		duration := millisecs * time.Millisecond
		time.Sleep(duration)

		fmt.Println("后台执行, duration:", duration)
		wg.Done()
	}()
}
