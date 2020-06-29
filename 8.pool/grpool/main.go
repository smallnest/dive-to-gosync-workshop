package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ivpusic/grpool"
)

func main() {
	pool := grpool.NewPool(20, 50)
	pool.WaitCount(50)

	fn := func(i int) {
		fmt.Println("Start Job", i)
		time.Sleep(time.Duration(i) * time.Second)
		fmt.Println("End Job", i)
	}

	for i := 0; i < 50; i++ {
		i := i
		pool.JobQueue <- func() {
			fn(i)
		}
	}
	log.Println("Submitted!")

	// wait until we call JobDone for all jobs
	pool.WaitAll()

	pool.Release()
}
