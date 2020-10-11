package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	limit()
	fmt.Println()

	burst()
	fmt.Println()

	reservation()
	fmt.Println()
}

func limit() {
	fmt.Println("run limit")
	rl := rate.NewLimiter(100, 1)

	for i := 0; i < 10; i++ {
		start := time.Now()
		_ = rl.Wait(context.Background())
		fmt.Printf("#%d, took: %v\n", i, time.Since(start))
	}
}

func burst() {
	fmt.Println("run burst")
	rl := rate.NewLimiter(100, 2)

	for i := 0; i < 10; i++ {
		start := time.Now()
		_ = rl.WaitN(context.Background(), 2) // return err if burst > 2
		fmt.Printf("#%d, took: %v\n", i, time.Since(start))
	}
}

func reservation() {
	fmt.Println("run reservation")
	rl := rate.NewLimiter(10, 100)
	rl.WaitN(context.Background(), 100)

	re := rl.ReserveN(time.Now(), 50)
	if !re.OK() {
		return
	}

	fmt.Println("delay:", re.Delay())
}
