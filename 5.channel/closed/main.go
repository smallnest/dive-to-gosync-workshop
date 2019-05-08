package main

import (
	"fmt"
	"time"
)

func main() {
	receive()
}

func receive() {
	ch := make(chan int, 100)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	// ch <- 0
	close(ch) // !!!!!!

	for {
		i, ok := <-ch
		fmt.Println(i, ok)
		time.Sleep(time.Second)
	}
}
