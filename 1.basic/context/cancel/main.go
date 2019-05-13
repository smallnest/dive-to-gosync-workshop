package main

import (
	"context"
	"fmt"
	_ "net/http/pprof"
	"time"
)

// parent的cancel可以取消子context.
// 子context只能取消自己和它的子.
// 子会propagateCancel，一直上上找，直到找到一个可cancel的context, 把自己加到它的children中.
func main() {
	ctx := context.Background()

	ctx1, c1 := context.WithCancel(ctx)
	go func() {
		fmt.Println("g1 start")
		<-ctx1.Done()
		fmt.Println("g1 done, err:", ctx1.Err())
	}()

	ctx2, c2 := context.WithCancel(ctx1)
	go func() {
		fmt.Println("g2 start")
		<-ctx2.Done()
		fmt.Println("g2 done, err:", ctx1.Err())
	}()

	ctx3, c3 := context.WithCancel(ctx2)
	go func() {
		fmt.Println("g3 start")
		<-ctx3.Done()
		fmt.Println("g3 done, err:", ctx1.Err())
	}()

	time.Sleep(1e9)
	c1()
	time.Sleep(5 * time.Second)
	c2()
	time.Sleep(5 * time.Second)
	c3()

	// time.Sleep(1e9)
	// c3()
	// time.Sleep(5 * time.Second)
	// c2()
	// time.Sleep(5 * time.Second)
	// c1()

}
