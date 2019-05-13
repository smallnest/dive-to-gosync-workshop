package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx1, c1 := context.WithCancel(context.Background())
	go func() {
		fmt.Println("g1 start")
		<-ctx1.Done()
		fmt.Println("g1 done, err:", ctx1.Err())
	}()

	ctx2, c2 := context.WithDeadline(ctx1, time.Now().Add(time.Hour))
	go func() {
		fmt.Println("g2 start")
		<-ctx2.Done()
		fmt.Println("g2 done, err:", ctx2.Err())
	}()

	time.Sleep(time.Second)
	c1()

	time.Sleep(10 * time.Second)
	c2()
}
