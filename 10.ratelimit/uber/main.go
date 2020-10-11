package main

import (
	"fmt"
	"time"

	"go.uber.org/ratelimit"
)

func main() {
	// limit()
	// fmt.Println()

	// unlimit()
	// fmt.Println()

	slack()
	fmt.Println()

	withoutSlack()
	fmt.Println()
}

func limit() {
	fmt.Println("run at 100/s")
	rl := ratelimit.New(100) // 每秒100个，均匀漏出

	takenPrev := time.Now()
	for i := 0; i < 10; i++ {
		start := time.Now()
		rl.Take()
		fmt.Printf("#%d, interval: %v, took: %v\n", i, time.Since(takenPrev), time.Since(start))
		takenPrev = time.Now()
	}
}

func unlimit() {
	fmt.Println("run as unlimited")
	rl := ratelimit.NewUnlimited()

	takenPrev := time.Now()
	for i := 0; i < 10; i++ {
		start := time.Now()
		rl.Take()
		fmt.Printf("#%d, interval: %v, took: %v\n", i, time.Since(takenPrev), time.Since(start))
		takenPrev = time.Now()
	}
}

func slack() {
	fmt.Println("run with slack")
	rl := ratelimit.New(10)
	rl.Take()
	time.Sleep(160 * time.Millisecond)
	rl.Take()
	time.Sleep(20 * time.Millisecond)
	start := time.Now()
	rl.Take()
	fmt.Printf("took: %v\n", time.Since(start))
}

func withoutSlack() {
	fmt.Println("run without slack")
	rl := ratelimit.New(10, ratelimit.WithoutSlack)
	rl.Take()
	time.Sleep(160 * time.Millisecond)
	rl.Take()
	time.Sleep(20 * time.Millisecond)
	start := time.Now()
	rl.Take()
	fmt.Printf("took: %v\n", time.Since(start))
}
