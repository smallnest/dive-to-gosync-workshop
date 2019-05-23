package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan int, 3)

	var done = make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				return
			case i := <-ch:
				fmt.Print(i)
			}
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			select {
			case <-done:
				return
			case ch <- i:
			}

			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(2 * time.Second)
}
