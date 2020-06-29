package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gammazero/workerpool"
)

func main() {
	wp := workerpool.New(20)

	fn := func(i int) {
		fmt.Println("Start Job", i)
		time.Sleep(time.Duration(i) * time.Second)
		fmt.Println("End Job", i)
	}

	for i := 0; i < 50; i++ {
		i := i
		wp.Submit(func() {
			fn(i)
		})
	}
	log.Println("Submitted!")

	wp.StopWait()
}
