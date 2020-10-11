package main

import (
	"fmt"
	"time"

	"github.com/juju/ratelimit"
)

func main() {
	limit()
	fmt.Println()

	limitWithQuantum()
	fmt.Println()

	limitWithRate()
	fmt.Println()
}

func limit() {
	rl := ratelimit.NewBucket(10*time.Millisecond, 1) // 1/10ms
	fmt.Println("run limit, rate:", rl.Rate())

	for i := 0; i < 10; i++ {
		start := time.Now()
		rl.Wait(1)
		fmt.Printf("#%d, took: %v\n", i, time.Since(start))
	}
}

func limitWithQuantum() {
	rl := ratelimit.NewBucketWithQuantum(100*time.Millisecond, 1, 10) // 10/100ms, refill per 100*time.Millisecond
	fmt.Println("run limitWithQuantum, rate:", rl.Rate())

	for i := 0; i < 10; i++ {
		start := time.Now()
		rl.Wait(1)
		fmt.Printf("#%d, took: %v\n", i, time.Since(start))
	}
}

func limitWithRate() {
	rl := ratelimit.NewBucketWithRate(100, 1) // 100/s
	fmt.Println("run limitWithRate, rate:", rl.Rate())

	for i := 0; i < 10; i++ {
		start := time.Now()
		rl.Wait(1)
		fmt.Printf("#%d, took: %v\n", i, time.Since(start))
	}
}
